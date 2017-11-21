package warc

import (
	"bytes"
	"compress/gzip"
)

// Sanitize removes any data from a warc record body
// that may interfere with parsing
func Sanitize(contentSniff string, body []byte) (sanitized []byte, err error) {
	switch contentSniff {
	case "application/pdf", "application/zip":
		// default to gzipping content
		buf := &bytes.Buffer{}
		w := gzip.NewWriter(buf)
		if _, err := w.Write(body); err != nil {
			return nil, err
		}
		if err := w.Close(); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	default:
		return bytes.Replace(body, crlf, []byte("\n"), -1), nil
	}
}
