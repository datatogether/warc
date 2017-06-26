package warc

import (
	"fmt"
	"io"
	"strconv"
)

// definedFieldNames maps tokens back to their respsective
// field name strings
var definedFieldNames = map[int]string{
	CONTENT_LENGTH:               "content-length",
	CONTENT_TYPE:                 "content-type",
	WARC_BLOCK_DIGEST:            "warc-block-digest",
	WARC_CONCURRENT_TO:           "warc-concurrent-to",
	WARC_FILENAME:                "warc-filename",
	WARC_DATE:                    "warc-date",
	WARC_IDENTIFIED_PAYLOAD_TYPE: "warc-identified-payload-type",
	WARC_IP_ADDRESS:              "warc-ip-address",
	WARC_PAYLOAD_DIGEST:          "warc-payload-digest",
	WARC_PROFILE:                 "warc-profile",
	WARC_RECORD_ID:               "warc-record-id",
	WARC_REFERS_TO:               "warc-refers-to",
	WARC_SEGMENT_ORIGIN_ID:       "warc-segment-origin-id",
	WARC_SEGMENT_NUMBER:          "warc-segment-number",
	WARC_SEGMENT_TOTAL_LENGTH:    "warc-segment-total-length",
	WARC_TARGET_URI:              "warc-target-uri",
	WARC_TRUNCATED:               "warc-truncated",
	WARC_TYPE:                    "warc-type",
	WARC_WARCINFO_ID:             "warc-warcinfo-id",
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
func writeHeader(w io.Writer, t RecordType, fields map[int]string) error {
	if err := writeWarcVersion(w); err != nil {
		return err
	}
	if err := writeField(w, definedFieldNames[WARC_TYPE], t.String()); err != nil {
		return err
	}
	if err := writeDefinedFields(w, fields); err != nil {
		return err
	}
	// if _, err := io.WriteString(w, "\r\n"); err != nil {
	// 	return err
	// }
	return nil
}

// WriteBlock writes all of reader (record content) to w, followed by 2 CRLF's
func writeBlock(w io.Writer, r []byte) error {
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
func writeDefinedFields(w io.Writer, fields map[int]string) error {
	for field, value := range fields {
		key := definedFieldNames[field]
		if key == "" {
			return fmt.Errorf("no defined field name with integer %d exists for value %s", field, value)
		}

		// don't write empty fields
		if value == "" {
			continue
		}

		if err := writeField(w, key, value); err != nil {
			return err
		}
	}
	return nil
}

func writeField(w io.Writer, key, value string) error {
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
