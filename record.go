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
	Write(io.Writer) error
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
func (r WARCInfo) Write(w io.Writer) error {
	err := WriteHeader(w, map[int]string{
		WARC_RECORD_ID:      r.WARCRecordId,
		WARC_DATE:           r.WARCDate.Format(time.RFC3339),
		CONTENT_LENGTH:      int64String(r.ContentLength),
		CONTENT_TYPE:        r.ContentType,
		WARC_BLOCK_DIGEST:   r.WARCBlockDigest,
		WARC_PAYLOAD_DIGEST: r.WARCPayloadDigest,
		WARC_TRUNCATED:      r.WARCTruncated,
		WARC_FILENAME:       r.WARCFilename,
	})
	if err != nil {
		return err
	}
	return WriteBlock(w, r.content)
}

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
func (r Response) Write(w io.Writer) error {
	err := WriteHeader(w, map[int]string{
		WARC_RECORD_ID: r.WARCRecordId,
		WARC_DATE:      r.WARCDate.Format(time.RFC3339),
		CONTENT_LENGTH: int64String(r.ContentLength),
		CONTENT_TYPE:   r.ContentType,
		// TODO - add fields
	})
	if err != nil {
		return err
	}
	return WriteBlock(w, r.content)
}

// A 'resource' record contains a resource, without full protocol response
// information. For example: a file directly retrieved from a locally
// accessible repository or the result of a networked retrieval where the
// protocol information has been discarded. The exact contents of a
// 'resource' record are determined not just by the record type but also by
// the URI scheme of the record's target-URI, as described below.
// For all 'resource' records, the payload is defined as the record block.
// A 'resource' record, with a synthesized target-URI, may also be used to
// archive other artefacts of a harvesting process inside WARC files.
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
func (r Resource) Write(w io.Writer) error {
	err := WriteHeader(w, map[int]string{
		WARC_RECORD_ID: r.WARCRecordId,
		WARC_DATE:      r.WARCDate.Format(time.RFC3339),
		CONTENT_LENGTH: int64String(r.ContentLength),
		CONTENT_TYPE:   r.ContentType,
		// TODO - add fields
	})
	if err != nil {
		return err
	}
	return WriteBlock(w, r.content)
}

// A 'request' record holds the details of a complete scheme-specific
// request, including network protocol information where possible. The
// exact contents of a 'request' record are determined not just by the
// record type but also by the URI scheme of the record's target-URI, as
// described below.
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
func (r Request) Write(w io.Writer) error {
	err := WriteHeader(w, map[int]string{
		WARC_RECORD_ID: r.WARCRecordId,
		WARC_DATE:      r.WARCDate.Format(time.RFC3339),
		CONTENT_LENGTH: int64String(r.ContentLength),
		CONTENT_TYPE:   r.ContentType,
		// TODO - add fields
	})
	if err != nil {
		return err
	}
	return WriteBlock(w, r.content)
}

// A 'metadata' record contains content created in order to further
// describe, explain, or accompany a harvested resource, in ways not
// covered by other record types. A 'metadata' record will almost always
// refer to another record of another type, with that other record holding
// original harvested or transformed content. (However, it is allowable for
// a 'metadata' record to refer to any record type, including other
// 'metadata' records.) Any number of metadata records may reference one
// specific other record.
// The format of the metadata record block may vary. The
// "application/warc-fields" format, defined earlier, may be used.
// Allowable fields include all \[DCMI\] plus the following field
// definitions. All fields are optional.
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
func (r Metadata) Write(w io.Writer) error {
	err := WriteHeader(w, map[int]string{
		WARC_RECORD_ID: r.WARCRecordId,
		WARC_DATE:      r.WARCDate.Format(time.RFC3339),
		CONTENT_LENGTH: int64String(r.ContentLength),
		CONTENT_TYPE:   r.ContentType,
		// TODO - add fields
	})
	if err != nil {
		return err
	}
	return WriteBlock(w, r.content)
}

// A 'revisit' record describes the revisitation of content already
// archived, and might include only an abbreviated content body which has
// to be interpreted relative to a previous record. Most typically, a
// 'revisit' record is used instead of a 'response' or 'resource' record to
// indicate that the content visited was either a complete or substantial
// duplicate of material previously archived.
// Using a 'revisit' record instead of another type is optional, for when
// benefits of reduced storage size or improved cross-referencing of
// material are desired.
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
func (r Revisit) Write(w io.Writer) error {
	err := WriteHeader(w, map[int]string{
		WARC_RECORD_ID: r.WARCRecordId,
		WARC_DATE:      r.WARCDate.Format(time.RFC3339),
		CONTENT_LENGTH: int64String(r.ContentLength),
		CONTENT_TYPE:   r.ContentType,
		// TODO - add fields
	})
	if err != nil {
		return err
	}
	return WriteBlock(w, r.content)
}

// A 'conversion' record shall contain an alternative version of another
// record's content that was created as the result of an archival process.
// Typically, this is used to hold content transformations that maintain
// viability of content after widely available rendering tools for the
// originally stored format disappear. As needed, the original content may
// be migrated (transformed) to a more viable format in order to keep the
// information usable with current tools while minimizing loss of
// information (intellectual content, look and feel, etc). Any number of
// 'conversion' records may be created that reference a specific source
// record, which may itself contain transformed content. Each
// transformation should result in a freestanding, complete record, with no
// dependency on survival of the original record.
// Metadata records may be used to further describe transformation records.
// Wherever practical, a 'conversion' record should contain a
// 'WARC-Refers-To' field to identify the prior material converted.
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
func (r Conversion) Write(w io.Writer) error {
	err := WriteHeader(w, map[int]string{
		WARC_RECORD_ID: r.WARCRecordId,
		WARC_DATE:      r.WARCDate.Format(time.RFC3339),
		CONTENT_LENGTH: int64String(r.ContentLength),
		CONTENT_TYPE:   r.ContentType,
		// TODO - add fields
	})
	if err != nil {
		return err
	}
	return WriteBlock(w, r.content)
}

// Record blocks from 'continuation' records must be appended to
// corresponding prior record block(s) (e.g., from other WARC files) to
// create the logically complete full-sized original record. That is,
// 'continuation' records are used when a record that would otherwise cause
// a WARC file size to exceed a desired limit is broken into segments. A
// continuation record shall contain the named fields
// 'WARC-Segment-Origin-ID' and 'WARC-Segment-Number', and the last
// 'continuation' record of a series shall contain a
// 'WARC-Segment-Total-Length' field. The full details of WARC record
// segmentation are described in the below section Record Segmentation. See
// also annex C.8 below for an example of a ‘continuation’ record.
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
func (r Continuation) Write(w io.Writer) error {
	err := WriteHeader(w, map[int]string{
		WARC_RECORD_ID: r.WARCRecordId,
		WARC_DATE:      r.WARCDate.Format(time.RFC3339),
		CONTENT_LENGTH: int64String(r.ContentLength),
		// TODO - add fields
	})
	if err != nil {
		return err
	}
	return WriteBlock(w, r.content)
}
