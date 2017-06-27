package warc

import (
	"fmt"
	"io"
	"strconv"
)

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
func writeHeader(w io.Writer, t RecordType, fields map[string]string) error {
	if err := writeWarcVersion(w); err != nil {
		return err
	}
	if err := writeField(w, warc_type, t.String()); err != nil {
		return err
	}
	if err := writeFields(w, fields); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "\r\n"); err != nil {
		return err
	}
	return nil
}

// WriteBlock writes all of reader (record content) to w, followed by 2 CRLF's
func writeBlock(w io.Writer, r []byte) error {
	// fmt.Println(string(r))
	// fmt.Println("------")
	if _, err := w.Write(r); err != nil {
		return err
	}
	// write 2xCRLF
	_, err := io.WriteString(w, "\r\n\r\n")
	return err
}

// writeWarcVersion writes the warc version header
func writeWarcVersion(w io.Writer) error {
	_, err := io.WriteString(w, "WARC/1.0\r\n")
	return err
}

// writeDefinedFields takes a map of token constants to values, and writes them to w
// it skips fields who's value is ""
func writeFields(w io.Writer, fields map[string]string) error {
	for field, value := range fields {
		if err := writeField(w, field, value); err != nil {
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

// convenience func to convert int64s to a string
func int64String(i int64) string {
	return strconv.FormatInt(i, 10)
}

func intString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
