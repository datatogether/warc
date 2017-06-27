package warc

import (
	"bufio"
	"bytes"
	"io"
)

// Reader parses WARC records from an underlying scanner.
// Create a new reader with NewReader
type Reader struct {
	scanner *bufio.Scanner    // scanner to pull tokens from
	phase   scanPhase         // current phase of record parsing
	version string            // current record verion
	headers map[string]string // current record headers
	key     string            // current header key to find the value of
	content []byte            // current record content
}

// NewReader creates a new WARC reader from an io.Reader
// Always use NewReader, (instead of manually allocating a reader)
func NewReader(r io.Reader) *Reader {
	rdr := &Reader{
		scanner: bufio.NewScanner(r),
		headers: map[string]string{},
	}
	rdr.scanner.Split(rdr.split)
	return rdr
}

// Read a record, will return nil, io.EOF to signal
// no more records
func (r *Reader) Read() (Record, error) {
	// rec, err := r.parseRecord()
	// if err == nil {
	// 	fmt.Println(string(rec.(*Resource).Content))
	// }
	return r.parseRecord()
}

// Consume the entire reader, returning a slice of records
func (r *Reader) ReadAll() (records []Record, err error) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return
}

// scanphase denotes different "modes" for scanning
type scanPhase int

const (
	scanPhaseVersion scanPhase = iota
	scanPhaseHeaderKey
	scanPhaseHeaderValue
	scanPhaseContent
)

func (r *Reader) parseRecord() (Record, error) {
	for r.scanner.Scan() {
		// need to copy here. trust.
		token := make([]byte, len(r.scanner.Bytes()))
		copy(token, r.scanner.Bytes())

		switch r.phase {
		case scanPhaseVersion:
			r.version = string(bytes.TrimSpace(token))
			r.phase = scanPhaseHeaderKey
		case scanPhaseHeaderKey:
			if bytes.Equal(token, crlf) {
				r.phase = scanPhaseContent
			} else {
				r.key = string(bytes.ToLower(bytes.TrimSpace(token)))
				r.phase = scanPhaseHeaderValue
			}
		case scanPhaseHeaderValue:
			r.headers[r.key] = string(bytes.TrimSpace(token))
			r.key = ""
			r.phase = scanPhaseHeaderKey
		case scanPhaseContent:
			r.content = token
			return r.record()
		}
	}

	return nil, io.EOF
}

// Generate a record from current reader state & reset the record
func (r *Reader) record() (Record, error) {
	defer r.reset()
	if r.scanner.Err() != nil {
		return nil, r.scanner.Err()
	}
	return newRecord(r.headers, r.content)
}

// reset the reader for another record
func (r *Reader) reset() {
	r.phase = scanPhaseVersion
	r.key = ""
	r.version = ""
	r.headers = map[string]string{}
	r.content = nil
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
