package warc

import (
	"fmt"
	"github.com/pborman/uuid"
	"io"
	"net/http"
	"sort"
)

func NewUuid() string {
	return fmt.Sprintf("<urn:uuid:%s>", uuid.New())
}

// WriteRecords calls Write on each record to w
func WriteRecords(w io.Writer, records []Record) error {
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

func writeHttpHeaders(w io.Writer, headers http.Header) error {
	for k, _ := range headers {
		if _, err := io.WriteString(w, fmt.Sprintf("%s: %s\r\n", k, headers.Get(k))); err != nil {
			return err
		}
	}
	return nil
}

// writeDefinedFields takes a map of token constants to values, and writes them to w
// it skips fields who's value is ""
func writeFields(w io.Writer, fields map[string]string) error {
	keys := make([]string, len(fields))
	i := 0
	for field, _ := range fields {
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
