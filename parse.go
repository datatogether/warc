package warc

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"
)

// Parse parses the sql and returns a set of Records, which
// is the AST representation of the query.
func ParseAll(r io.Reader) ([]Record, error) {
	tokenizer := NewTokenizer(r)
	if yyParse(tokenizer) != 0 {
		return nil, errors.New(tokenizer.LastError)
	}
	return tokenizer.Records, nil
}

// ParseRecord parses a single record from r
func ParseRecord(r io.Reader) (Record, error) {
	tokenizer := NewTokenizer(r)
	if yyParse(tokenizer) != 0 {
		return nil, errors.New(tokenizer.LastError)
	}
	return tokenizer.Record.Record()
}

// TODO
// func ParseRecord(r io.Reader) (Record, error) {}

// parseRecord is an internal struct for lexing values into.
// because we don't know what kind of record we're working with
// until the WarcType header field is encountered, this struct
// serves as an intermediary for accumulating values into
type parseRecord struct {
	ContentLength             string
	ContentType               string
	Version                   string
	WARCType                  RecordType
	WARCRecordId              string
	WARCDate                  string
	WARCConcurrentTo          string
	WARCBlockDigest           string
	WARCPayloadDigest         string
	WARCIPAddress             string
	WARCRefersTo              string
	WARCTargetURI             string
	WARCTruncated             string
	WARCWarcinfoID            string
	WARCFilename              string
	WARCProfile               string
	WARCIdentifiedPayloadType string
	WARCSegmentOriginID       string
	WARCSegmentNumber         string
	WARCSegmentTotalLength    string
	CustomFields              map[string]string
	Content                   []byte
}

func newParseRecord() *parseRecord {
	return &parseRecord{
		CustomFields: map[string]string{},
	}
}

func (p *parseRecord) Record() (Record, error) {
	// TODO - parse these at scan-time
	warcDate, err := time.Parse(time.RFC3339, p.WARCDate)
	if err != nil {
		return nil, err
	}

	contentLength, err := strconv.ParseInt(p.ContentLength, 10, 64)
	if err != nil {
		return nil, err
	}

	switch p.WARCType {
	case RecordTypeWarcInfo:
		return &WARCInfo{
			WARCRecordId:      p.WARCRecordId,
			WARCDate:          warcDate,
			ContentLength:     contentLength,
			ContentType:       p.ContentType,
			WARCBlockDigest:   p.WARCBlockDigest,
			WARCPayloadDigest: p.WARCPayloadDigest,
			WARCTruncated:     p.WARCTruncated,
			WARCFilename:      p.WARCFilename,
			Content:           p.Content,
		}, nil
	case RecordTypeResponse:
		return &Response{
			WARCRecordId:              p.WARCRecordId,
			WARCDate:                  warcDate,
			ContentLength:             contentLength,
			ContentType:               p.ContentType,
			WARCConcurrentTo:          p.WARCConcurrentTo,
			WARCBlockDigest:           p.WARCBlockDigest,
			WARCPayloadDigest:         p.WARCPayloadDigest,
			WARCIPAddress:             p.WARCIPAddress,
			WARCTargetURI:             p.WARCTargetURI,
			WARCTruncated:             p.WARCTruncated,
			WARCWarcinfoID:            p.WARCWarcinfoID,
			WARCIdentifiedPayloadType: p.WARCIdentifiedPayloadType,
			Content:                   p.Content,
		}, nil
	case RecordTypeResource:
		return &Resource{
			WARCRecordId:              p.WARCRecordId,
			WARCDate:                  warcDate,
			ContentLength:             contentLength,
			ContentType:               p.ContentType,
			WARCConcurrentTo:          p.WARCConcurrentTo,
			WARCPayloadDigest:         p.WARCPayloadDigest,
			WARCBlockDigest:           p.WARCBlockDigest,
			WARCIPAddress:             p.WARCIPAddress,
			WARCTargetURI:             p.WARCTargetURI,
			WARCTruncated:             p.WARCTruncated,
			WARCWarcinfoID:            p.WARCWarcinfoID,
			WARCIdentifiedPayloadType: p.WARCIdentifiedPayloadType,
			Content:                   p.Content,
		}, nil
	case RecordTypeRequest:
		return &Request{
			WARCRecordId:              p.WARCRecordId,
			WARCDate:                  warcDate,
			ContentLength:             contentLength,
			ContentType:               p.ContentType,
			WARCConcurrentTo:          p.WARCConcurrentTo,
			WARCBlockDigest:           p.WARCBlockDigest,
			WARCPayloadDigest:         p.WARCPayloadDigest,
			WARCIPAddress:             p.WARCIPAddress,
			WARCTargetURI:             p.WARCTargetURI,
			WARCTruncated:             p.WARCTruncated,
			WARCWarcinfoID:            p.WARCWarcinfoID,
			WARCIdentifiedPayloadType: p.WARCIdentifiedPayloadType,
			Content:                   p.Content,
		}, nil
	case RecordTypeMetadata:
		return &Metadata{
			WARCRecordId:     p.WARCRecordId,
			WARCDate:         warcDate,
			ContentLength:    contentLength,
			ContentType:      p.ContentType,
			WARCConcurrentTo: p.WARCConcurrentTo,
			WARCBlockDigest:  p.WARCBlockDigest,
			WARCIPAddress:    p.WARCIPAddress,
			WARCRefersTo:     p.WARCRefersTo,
			WARCTargetURI:    p.WARCTargetURI,
			WARCTruncated:    p.WARCTruncated,
			WARCWarcinfoID:   p.WARCWarcinfoID,
			Content:          p.Content,
		}, nil
	case RecordTypeRevisit:
		return &Revisit{
			WARCRecordId:      p.WARCRecordId,
			WARCDate:          warcDate,
			ContentLength:     contentLength,
			ContentType:       p.ContentType,
			WARCConcurrentTo:  p.WARCConcurrentTo,
			WARCBlockDigest:   p.WARCBlockDigest,
			WARCPayloadDigest: p.WARCPayloadDigest,
			WARCIPAddress:     p.WARCIPAddress,
			WARCRefersTo:      p.WARCRefersTo,
			WARCTargetURI:     p.WARCTargetURI,
			WARCTruncated:     p.WARCTruncated,
			WARCWarcinfoID:    p.WARCWarcinfoID,
			WARCProfile:       p.WARCProfile,
			Content:           p.Content,
		}, nil
	case RecordTypeConversion:
		return &Conversion{
			WARCRecordId:      p.WARCRecordId,
			WARCDate:          warcDate,
			ContentLength:     contentLength,
			ContentType:       p.ContentType,
			WARCBlockDigest:   p.WARCBlockDigest,
			WARCPayloadDigest: p.WARCPayloadDigest,
			WARCRefersTo:      p.WARCRefersTo,
			WARCTruncated:     p.WARCTruncated,
			WARCWarcinfoID:    p.WARCWarcinfoID,
			Content:           p.Content,
		}, nil
	case RecordTypeContinuation:
		seg, err := strconv.ParseInt(p.WARCSegmentNumber, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("error parsing WARCSegmentNumber: %s", err.Error())
		}
		length, err := strconv.ParseInt(p.WARCSegmentTotalLength, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing WARCSegmentTotalLength: %s", err.Error())
		}
		return &Continuation{
			WARCRecordId:           p.WARCRecordId,
			WARCDate:               warcDate,
			ContentLength:          contentLength,
			WARCBlockDigest:        p.WARCBlockDigest,
			WARCPayloadDigest:      p.WARCPayloadDigest,
			WARCTruncated:          p.WARCTruncated,
			WARCWarcinfoID:         p.WARCWarcinfoID,
			WARCSegmentNumber:      int(seg),
			WARCSegmentOriginID:    p.WARCSegmentOriginID,
			WARCSegmentTotalLength: length,
			Content:                p.Content,
		}, nil
	default:
		// TODO - handle missing type field
		return nil, fmt.Errorf("unrecognized WARC type: '%s'", p.WARCType)
	}
}
