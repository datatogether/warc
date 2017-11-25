package warc

import (
	"bufio"
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"io"
	"io/ioutil"
)

// Reader parses WARC records from an underlying scanner.
// Create a new reader with NewReader
type Reader struct {
	rc      io.ReadCloser  // raw io.readerCloser
	scanner *bufio.Scanner // scanner to pull tokens from
	phase   scanPhase
}

// NewReader creates a new WARC reader from an io.Reader
// Always use NewReader, (instead of manually allocating a reader)
func NewReader(r io.Reader) (*Reader, error) {
	rc, err := decompress(r)
	if err != nil {
		return nil, err
	}

	rdr := &Reader{
		rc:      rc,
		scanner: bufio.NewScanner(rc),
	}
	rdr.scanner.Split(rdr.split)
	return rdr, nil
}

// Read a record, will return nil, io.EOF to signal
// no more records
func (r *Reader) Read() (Record, error) {
	// rec, err := r.readRecord()
	// if err == nil {
	// 	fmt.Println(string(rec.(*Resource).Content))
	// }
	return r.readRecord()
}

// ReadAll Consumes the entire reader, returning a slice of records
func (r *Reader) ReadAll() (records Records, err error) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}
}

// scanphase denotes different "modes" for scanning
type scanPhase int

const (
	scanPhaseVersion scanPhase = iota
	scanPhaseHeaderKey
	scanPhaseHeaderValue
	scanPhaseContent
)

func (r *Reader) readRecord() (rec Record, err error) {
	var key string
	rec = Record{
		Headers: map[string]string{},
	}

	for r.scanner.Scan() {
		token := r.scanner.Bytes()

		switch r.phase {
		case scanPhaseVersion:
			rec.Format = recordFormat(string(bytes.TrimSpace(token)))
			r.phase = scanPhaseHeaderKey
		case scanPhaseHeaderKey:
			if bytes.Equal(token, crlf) {
				r.phase = scanPhaseContent
			} else {
				key = CanonicalKey(string(token))
				r.phase = scanPhaseHeaderValue
			}
		case scanPhaseHeaderValue:
			rec.Headers[key] = string(bytes.TrimSpace(token))
			if key == FieldNameWARCType {
				rec.Type = ParseRecordType(rec.Headers[key])
			}
			r.phase = scanPhaseHeaderKey
		case scanPhaseContent:
			// need to copy here b/c the underlying bytes shift as the buffer
			// moves through the file
			buf := make([]byte, len(r.scanner.Bytes()))
			copy(buf, r.scanner.Bytes())
			rec.Content = bytes.NewBuffer(buf)
			r.phase = scanPhaseVersion
			return
		}
	}

	return rec, io.EOF
}

func (r *Reader) split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	switch r.phase {
	case scanPhaseVersion:
		return splitLine(data, atEOF)
	case scanPhaseHeaderKey:
		return splitKey(data, atEOF)
	case scanPhaseHeaderValue:
		return splitValue(data, atEOF)
	default: // default to scanPhaseContent
		return splitBlock(data, atEOF)
	}
}

var crlf = []byte("\r\n")
var doubleCrlf = []byte("\r\n\r\n")

func splitLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func splitKey(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if bytes.Index(data, crlf) == 0 {
		return len(crlf), crlf, nil
	}
	if i := bytes.IndexByte(data, ':'); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func splitValue(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// TODO - MULTILINE VALUES
	if i := bytes.Index(data, crlf); i == 0 {
		// if we hit double clrf return
		return len(crlf), nil, nil
	} else if i > 0 {
		return i + len(crlf), data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func splitBlock(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.Index(data, doubleCrlf); i >= 0 {
		return i + len(doubleCrlf), data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// readBlockBody
func readBlockBody(data []byte) ([]byte, error) {
	start := bytes.LastIndex(data, crlf)
	if start == -1 {
		return data, nil
	}
	r := bytes.NewReader(data[start+len(crlf):])
	res, err := decompress(r)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return ioutil.ReadAll(res)
}

const (
	compressionNone = iota
	compressionBZIP
	compressionGZIP
)

// guessCompression returns the compression type of a data stream by matching
// the first two bytes with the magic numbers of compression formats.
func guessCompression(b *bufio.Reader) (int, error) {
	magic, err := b.Peek(2)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return compressionNone, err
	}
	switch {
	case magic[0] == 0x42 && magic[1] == 0x5a:
		return compressionBZIP, nil
	case magic[0] == 0x1f && magic[1] == 0x8b:
		return compressionGZIP, nil
	}
	return compressionNone, nil
}

// decompress automatically decompresses data streams and makes sure the result
// obeys the io.ReadCloser interface. This way callers don't need to check
// whether the underlying reader has a Close() function or not, they just call
// defer Close() on the result.
func decompress(r io.Reader) (res io.ReadCloser, err error) {
	// Create a buffered reader to peek the stream's magic number.
	dataReader := bufio.NewReader(r)
	compr, err := guessCompression(dataReader)
	if err != nil {
		return nil, err
	}
	switch compr {
	case compressionGZIP:
		gzipReader, err := gzip.NewReader(dataReader)
		if err != nil {
			return nil, err
		}
		res = gzipReader
	case compressionBZIP:
		bzipReader := bzip2.NewReader(dataReader)
		res = ioutil.NopCloser(bzipReader)
	case compressionNone:
		res = ioutil.NopCloser(dataReader)
	}
	return res, err
}
