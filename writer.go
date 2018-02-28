package warc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/pborman/uuid"
)

// NewUUID generates a new version 4 uuid
func NewUUID() string {
	return fmt.Sprintf("<urn:uuid:%s>", uuid.New())
}

// WriteRecords calls Write on each record to w.
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
