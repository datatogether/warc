package warc

import (
	"bytes"
	"compress/gzip"
	stdErrors "errors"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

// NewUUID generates a new version 4 uuid
func NewUUID() string {
	return fmt.Sprintf("<urn:uuid:%s>", uuid.New())
}

type flusher interface {
	Flush() error
}

type closeResetWriter interface {
	Close() error
	Reset(w io.Writer)
}

// Writer provides functionality for writing WARC files in compressed and
// uncompressed formats.
//
// To construct a Writer, call NewWriterCompressed or NewWriterRaw.
type Writer struct {
	seekW io.WriteSeeker
	wr    io.Writer
	cmprs bool

	// RecordCallback will be called after each record is written to the file.
	// If a WriteSeeker was not provided, the provided positions will be
	// invalid.
	RecordCallback func(r *Record, startPos, endPos int64)
}

// NewWriterCompressed initializes a WARC Writer writing to a compressed
// stream.  The first parameter should be the "backing stream" of the
// compression.  The second parameter is a compress/gzip writer writing to the
// rawFile parameter.
//
// Seek will only be called with whence == io.SeekCurrent and offset == 0.
//
// See also CountWriter() if you need a "fake" Seek implementation.
func NewWriterCompressed(rawFile io.WriteSeeker, cmprsWriter *gzip.Writer) (*Writer, error) {
	w := &Writer{
		seekW: rawFile,
		wr:    cmprsWriter,
		cmprs: true,
	}
	return w, nil
}

// NewWriterRaw initializes a WARC Writer writing to an uncompressed stream.
// If the provided Writer implements io.Seeker, the RecordCallback will be
// available.  If the provided Writer implements interface{Flush() error}, it
// will be flushed after every written Record.
//
// See also CountWriter() if you need a "fake" Seek implementation.
func NewWriterRaw(out io.Writer) (*Writer, error) {
	w := &Writer{
		wr: out,
	}
	if wseeker, ok := out.(io.WriteSeeker); ok {
		w.seekW = wseeker
	}
	return w, nil
}

type countWriter struct {
	count int64
	w     io.Writer
}

// CountWriter implements a limited version of io.Seeker around the provided
// Writer.  It only supports offset == 0 and whence == io.SeekCurrent or
// io.SeekEnd, and returns the current number of written bytes in both cases.
func CountWriter(w io.Writer) io.WriteSeeker {
	return &countWriter{count: 0, w: w}
}

// implements io.Writer
func (c *countWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	if n >= 0 {
		c.count += int64(n)
	}
	return n, err
}

var errCountWriterNotImplemented = stdErrors.New("unsupported seek operation")

// implements io.Seeker
func (c *countWriter) Seek(offset int64, whence int) (int64, error) {
	if offset != 0 || !(whence == io.SeekCurrent || whence == io.SeekEnd) {
		return 0, errCountWriterNotImplemented
	}
	return c.count, nil
}

// WriteRecord adds the record to the WARC file and returns the file offsets
// the record was written at.
//
// No processing is done to the Record contents beyond those mentioned in
// Record.Write.  If clients want extra processing (e.g. setting the
// Warcinfo-Id header) they are encouraged to create a wrapper.
func (w *Writer) WriteRecord(rec *Record) (startPos, endPos int64, err error) {
	if w.seekW != nil {
		startPos, err = w.seekW.Seek(0, io.SeekCurrent)
		err = errors.Wrap(err, "warc writer: seek 0")
		if err != nil {
			return
		}
	}

	err = rec.Write(w.wr)
	err = errors.Wrap(err, "warc writer: write record")
	if err != nil {
		return
	}

	// flush is not sufficient for gzip writer, need to Close/Reset
	closeReset, crOK := w.wr.(closeResetWriter)
	crOK = crOK && w.cmprs
	if flusher, ok := w.wr.(flusher); ok && !crOK {
		err = errors.Wrap(flusher.Flush(), "warc writer: flush")
		if err != nil {
			return
		}
	}
	if crOK {
		err = errors.Wrap(closeReset.Close(), "warc writer: flush")
		if err != nil {
			return
		}
	}
	// check the position BETWEEN close / reset
	if w.seekW != nil {
		endPos, err = w.seekW.Seek(0, io.SeekCurrent)
		err = errors.Wrap(err, "warc writer: seek 0")
		if err != nil {
			return
		}
	}
	if crOK {
		closeReset.Reset(w.seekW)
	}

	return
}

// Close cleans up any resources the warc.Writer might be holding on to.
func (w *Writer) Close() error {
	return nil
}

// WriteRecords calls Write on each record to w.
// Deprecated: see Writer type
func WriteRecords(w io.Writer, records Records) error {
	for _, rec := range records {
		if err := rec.Write(w); err != nil {
			return err
		}
	}
	return nil
}

// WriteHeader writes a fully formed header with version to w
func writeHeader(w io.Writer, r *Record) error {
	if err := writeWarcVersion(w, r); err != nil {
		return err
	}
	if err := writeFields(w, r.Headers); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "\r\n"); err != nil {
		return err
	}
	return nil
}

// WriteBlock writes all of reader (record content) to w, followed by 2 CRLF's
func writeBlock(w io.Writer, r io.Reader) error {
	if _, err := io.Copy(w, r); err != nil {
		return err
	}
	// write 2xCRLF
	_, err := io.WriteString(w, "\r\n\r\n")
	return err
}

// writeWarcVersion writes the warc version header
func writeWarcVersion(w io.Writer, r *Record) error {
	_, err := io.WriteString(w, r.Format.String()+"\r\n")
	return err
}

// WriteRequestMethodAndHeaders calls req.Write(w). (deprecated, see
// NewRequestResponseRecords)
func WriteRequestMethodAndHeaders(w io.Writer, req *http.Request) error {
	return req.Write(w)
}

// WriteHTTPHeaders writes all http headers to an io.Writer, separated by newlines
// Used to add http headers to a record
func WriteHTTPHeaders(w io.Writer, headers http.Header) error {
	for k := range headers {
		if _, err := io.WriteString(w, fmt.Sprintf("%s: %s\n", k, headers.Get(k))); err != nil {
			return err
		}
	}
	return nil
}

// replaceBlockBody replaces the body of a warc record, leaving
// and written headers in place
func replaceBlockBody(data, repl []byte) ([]byte, error) {
	start := bytes.LastIndex(data, crlf)
	if start == -1 {
		return repl, nil
	}
	return append(data[start:], repl...), nil
}

// writeDefinedFields takes a map of token constants to values, and writes them to w
// it skips fields who's value is ""
func writeFields(w io.Writer, fields map[string]string) error {
	keys := make([]string, len(fields))
	i := 0
	for field := range fields {
		keys[i] = field
		i++
	}

	// sort fields alphabetically
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, key := range keys {
		if err := writeField(w, key, fields[key]); err != nil {
			return err
		}
	}
	return nil
}

func writeField(w io.Writer, key, value string) error {
	// don't write empty fields
	if value == "" {
		return nil
	}
	// format entry
	ln := fmt.Sprintf("%s: %s\r\n", key, value)
	_, err := io.WriteString(w, ln)
	return err
}
