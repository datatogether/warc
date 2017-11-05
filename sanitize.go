package warc

import (
	"bytes"
)

// Sanitize removes any data from a warc record body
// that may interfere with parsing
func Sanitize(contentSniff string, body []byte) (sanitized []byte) {
	switch contentSniff {
	case "application/html; charset=utf-8":
		return bytes.Replace(body, crlf, []byte("\n"), -1)
	case "application/zip":
		return body
	default:
		return bytes.Replace(body, crlf, []byte("\n"), -1)
	}
}
