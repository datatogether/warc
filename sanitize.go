package warc

import (
	"bytes"
)

// Sanitize removes any data from a warc record body
// that may interfere with parsing
func Sanitize(body []byte) (sanitized []byte) {
	// TODO - lololol finish
	return bytes.Replace(body, crlf, []byte("\n"), -1)
}
