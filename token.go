package warc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const eofChar = 0x100

var recordTypes = map[string]int{
	"conversion":   CONVERSION,
	"continuation": CONTINUATION,
	"metadata":     METADATA,
	"resource":     RESOURCE,
	"response":     RESPONSE,
	"request":      REQUEST,
	"revisit":      REVISIT,
	"warcinfo":     WARCINFO,
}

var definedFields = map[string]int{
	"content-length":               CONTENT_LENGTH,
	"content-type":                 CONTENT_TYPE,
	"warc-block-digest":            WARC_BLOCK_DIGEST,
	"warc-concurrent-to":           WARC_CONCURRENT_TO,
	"warc-filename":                WARC_FILENAME,
	"warc-date":                    WARC_DATE,
	"warc-identified-payload-type": WARC_IDENTIFIED_PAYLOAD_TYPE,
	"warc-ip-address":              WARC_IP_ADDRESS,
	"warc-payload-digest":          WARC_PAYLOAD_DIGEST,
	"warc-profile":                 WARC_PROFILE,
	"warc-record-id":               WARC_RECORD_ID,
	"warc-refers-to":               WARC_REFERS_TO,
	"warc-segment-origin-id":       WARC_SEGMENT_ORIGIN_ID,
	"warc-segment-number":          WARC_SEGMENT_NUMBER,
	"warc-segment-total-length":    WARC_SEGMENT_TOTAL_LENGTH,
	"warc-target-uri":              WARC_TARGET_URI,
	"warc-truncated":               WARC_TRUNCATED,
	"warc-type":                    WARC_TYPE,
	"warc-warcinfo-id":             WARC_WARCINFO_ID,
}

type scanPhase int

const (
	scanPhaseVersion scanPhase = iota
	scanPhaseHeader
	scanPhaseContent
)

// NewTokenizer creates a new Tokenizer to read input from r
func NewTokenizer(r io.Reader) *Tokenizer {
	tkn := &Tokenizer{
		Phase: scanPhaseVersion,
		// Records: make([]Record, 0),
		Record: &parseRecord{},
	}
	s := bufio.NewScanner(r)
	// s.Split(tkn.ScanToken)
	tkn.InStream = s
	return tkn
}

// Tokenizer generates WARC tokens by scanning
// from InStream
type Tokenizer struct {
	InStream         *bufio.Scanner
	ForceEOF         bool
	Phase            scanPhase
	Position         int
	lastToken        []byte
	nextToken        int
	nextBytes        []byte
	LastError        string
	RecordsRemaining int
	Records          []Record
	Record           *parseRecord
}

// Future plans for one day
// func (tkn *Tokenizer) ScanToken(data []byte, atEOF bool) (advance int, token []byte, err error) {
// 	if atEOF && len(data) == 0 {
// 		return 0, nil, nil
// 	}
// 	if i := bytes.IndexByte(data, '\n'); i >= 0 {
// 		// We have a full newline-terminated line.
// 		return i + 1, dropCR(data[0:i]), nil
// 	}
// 	// If we're at EOF, we have a final, non-terminated line. Return it.
// 	if atEOF {
// 		return len(data), dropCR(data), nil
// 	}
// 	// Request more data.
// 	return 0, nil, nil
// }

// func dropCR(data []byte) []byte {
// 	if len(data) > 0 && data[len(data)-1] == '\r' {
// 		return data[0 : len(data)-1]
// 	}
// 	return data
// }

// Scan is called by the parser to lex the instream into
// a series of tokens.
// This implementation is really silly & brittle.
// But will be a good start to get happy-path tests to pass
// At some point this will need much more attention
// TODO - make this not silly.
func (tkn *Tokenizer) Scan() (int, []byte) {
	if tkn.nextToken != 0 {
		defer func() {
			tkn.nextToken = 0
			tkn.nextBytes = nil
		}()
		return tkn.nextToken, tkn.nextBytes
	}

	tkn.next()

	if tkn.ForceEOF {
		return 0, nil
	}
	switch tkn.Phase {
	case scanPhaseVersion:
		tkn.Phase = scanPhaseHeader
		v := tkn.InStream.Bytes()
		// if bytes is empty file is over
		if len(v) == 0 {
			return 0, nil
		}
		return WARC_VERSION, v
	case scanPhaseHeader:
		// text should be blank when we encounter CLRF CLRF
		// advance the phase &
		if tkn.InStream.Text() == "" {
			tkn.Phase = scanPhaseContent
			return tkn.Scan()
		}

		// split header along first ':' character
		header := strings.SplitAfterN(tkn.InStream.Text(), ":", 2)
		if len(header) != 2 {
			return 0, nil
		}

		key := strings.TrimSpace(strings.ToLower(header[0]))
		key = key[:len(key)-1]
		// TODO - multiline values! - handle with a custom scanner func
		value := strings.TrimSpace(header[1])

		tkn.nextToken = FIELD_VALUE
		tkn.nextBytes = []byte(value)

		// check for definedFields
		field := definedFields[key]
		switch field {
		case WARC_TYPE:
			// parse value now
			rt := recordTypes[strings.ToLower(value)]
			if rt == 0 {
				return 0, nil
			}
			tkn.nextToken = rt
			tkn.nextBytes = nil
			return WARC_TYPE, nil
		case 0:
			// if definedFields maps to zero, it's a custom
			// field key
			return FIELD_KEY, []byte(key)
		default:
			// fmt.Println("returning field", field, key, tkn.nextToken, value)
			return field, []byte(key)
		}
	case scanPhaseContent:
		tkn.Phase = scanPhaseVersion
		// block := tkn.InStream.Bytes()
		block := []byte{}
		for {
			next := tkn.InStream.Bytes()
			if len(next) == 0 {
				break
			}
			// TODO - this can/will break hashes b/c the internal scanner
			// removes \r from \r\n
			block = append(block, '\n')
			block = append(block, next...)
			tkn.next()
		}
		if tkn.RecordsRemaining == 1 {
			tkn.ForceEOF = true
		}
		return BLOCK, block
	}

	// Header fields can be extended over multiple lines by
	// preceding each extra line with at least one space or tab character.
	return 0, nil
}

// next advances the token scanner
func (tkn *Tokenizer) next() {
	if !tkn.InStream.Scan() || tkn.InStream.Err() != nil {
		if tkn.InStream.Err() != nil {
			fmt.Println("instream scan error:", tkn.InStream.Err())
		}
		tkn.ForceEOF = true
	}
	tkn.Position++
}

// Lex returns the next token form the Tokenizer.
// This function is used by go yacc.
func (tkn *Tokenizer) Lex(lval *yySymType) int {
	typ, val := tkn.Scan()
	switch typ {
	case FIELD_KEY, FIELD_VALUE, BLOCK:
		lval.bytes = val
	}
	tkn.lastToken = val
	return typ
}

func (tkn *Tokenizer) Error(err string) {
	buf := &bytes.Buffer{}
	if tkn.lastToken != nil {
		fmt.Fprintf(buf, "%s at position %v near '%s'", err, tkn.Position, tkn.lastToken)
	} else {
		fmt.Fprintf(buf, "%s at position %v", err, tkn.Position)
	}
	tkn.LastError = buf.String()
}
