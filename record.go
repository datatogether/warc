package warc

import (
	"io"
	"time"
)

// A WARC format file is the simple concatenation of one or more WARC
// records. The first record usually describes the records to follow. In
// general, record content is either the direct result of a retrieval
// attempt — web pages, inline images, URL redirection information, DNS
// hostname lookup results, standalone files, etc. — or is synthesized
// material (e.g., metadata, transformed content) that provides additional
// information about archived content.
type Records []Record

// Basic constituent of a WARC file, consisting of a sequence of WARC
// records.
type Record interface {
	Type() RecordType
	// Id() string
	// Date() time.Time
	// Header() *Header
	// ContentLength() int64
	Content() io.Reader
}

// A 'warcinfo' record describes the records that follow it, up through end
// of file, end of input, or until next 'warcinfo' record. Typically, this
// appears once and at the beginning of a WARC file. For a web archive, it
// often contains information about the web crawl which generated the
// following records.
// The format of this descriptive record block may vary, though the use of
// the "application/warc-fields" content-type is recommended. Allowable
// fields include, but are not limited to, all \[DCMI\] plus the following
// field definitions. All fields are optional.
type WARCInfo struct {
	WARCRecordId      string
	WARCDate          time.Time
	ContentLength     int64
	ContentType       string
	WARCBlockDigest   string
	WARCPayloadDigest string
	WARCTruncated     string
	WARCFilename      string
	content           io.Reader
}

func (r WARCInfo) Type() RecordType   { return RecordTypeWarcInfo }
func (r WARCInfo) Content() io.Reader { return r.content }

// A 'response' record should contain a complete scheme-specific response,
// including network protocol information where possible. The exact
// contents of a 'response' record are determined not just by the record
// type but also by the URI scheme of the record's target-URI, as described
// below.
type Response struct {
	WARCRecordId              string
	WARCDate                  time.Time
	ContentLength             int64
	ContentType               string
	WARCConcurrentTo          string
	WARCBlockDigest           string
	WARCPayloadDigest         string
	WARCIPAddress             string
	WARCTargetURI             string
	WARCTruncated             string
	WARCWarcinfoID            string
	WARCIdentifiedPayloadType string
	content                   io.Reader
}

func (r Response) Type() RecordType   { return RecordTypeResponse }
func (r Response) Content() io.Reader { return r.content }

type Resource struct {
	WARCRecordId              string
	WARCDate                  time.Time
	ContentLength             int64
	ContentType               string
	WARCConcurrentTo          string
	WARCBlockDigest           string
	WARCPayloadDigest         string
	WARCIPAddress             string
	WARCTargetURI             string
	WARCTruncated             string
	WARCWarcinfoID            string
	WARCIdentifiedPayloadType string
	content                   io.Reader
}

func (r Resource) Type() RecordType   { return RecordTypeResource }
func (r Resource) Content() io.Reader { return r.content }

type Request struct {
	WARCRecordId              string
	WARCDate                  time.Time
	ContentLength             int64
	ContentType               string
	WARCConcurrentTo          string
	WARCBlockDigest           string
	WARCPayloadDigest         string
	WARCIPAddress             string
	WARCTargetURI             string
	WARCTruncated             string
	WARCWarcinfoID            string
	WARCIdentifiedPayloadType string
	content                   io.Reader
}

func (r Request) Type() RecordType   { return RecordTypeRequest }
func (r Request) Content() io.Reader { return r.content }

type Metadata struct {
	WARCRecordId     string
	WARCDate         time.Time
	ContentLength    int64
	ContentType      string
	WARCConcurrentTo string
	WARCBlockDigest  string
	WARCIPAddress    string
	WARCRefersTo     string
	WARCTargetURI    string `json:"omitempty"`
	WARCTruncated    string
	WARCWarcinfoID   string
	content          io.Reader
}

func (r Metadata) Type() RecordType   { return RecordTypeMetadata }
func (r Metadata) Content() io.Reader { return r.content }

type Revisit struct {
	WARCRecordId      string
	WARCDate          time.Time
	ContentLength     int64
	ContentType       string
	WARCConcurrentTo  string
	WARCBlockDigest   string
	WARCPayloadDigest string
	WARCIPAddress     string
	WARCRefersTo      string
	WARCTargetURI     string
	WARCTruncated     string
	WARCWarcinfoID    string
	WARCProfile       string
	content           io.Reader
}

func (r Revisit) Type() RecordType   { return RecordTypeRevisit }
func (r Revisit) Content() io.Reader { return r.content }

type Conversion struct {
	WARCRecordId      string
	WARCDate          time.Time
	ContentLength     int64
	ContentType       string
	WARCBlockDigest   string
	WARCPayloadDigest string
	WARCRefersTo      string
	WARCTruncated     string
	WARCWarcinfoID    string
	content           io.Reader
}

func (r Conversion) Type() RecordType   { return RecordTypeConversion }
func (r Conversion) Content() io.Reader { return r.content }

type Continuation struct {
	WARCRecordId           string
	WARCDate               time.Time
	ContentLength          int64
	WARCBlockDigest        string
	WARCPayloadDigest      string
	WARCTruncated          string
	WARCWarcinfoID         string
	WARCSegmentNumber      int
	WARCSegmentOriginID    string
	WARCSegmentTotalLength int64
	content                io.Reader
}

func (r Continuation) Type() RecordType   { return RecordTypeContinuation }
func (r Continuation) Content() io.Reader { return r.content }
