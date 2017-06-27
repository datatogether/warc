package warc

import (
	"bytes"
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

// Record is the common interface for all WARC Record Types
// A Record consists of a version indicator (eg: WARC/1.0), zero or more headers,
// and possibly a content block.
// Upgrades to specific types of records can be done using type assertions
// and/or the Type method.
type Record interface {
	// Return the type of record
	Type() RecordType
	// The ID for this record
	GetRecordID() string
	// Datestamp of record creation
	GetDate() time.Time
	// Length of content block in bytes
	GetContentLength() int64
	// Reader for content itself
	GetContent() io.Reader
	// Write this record to a given writer
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
	Content           []byte
}

func (r WARCInfo) Type() RecordType        { return RecordTypeWarcInfo }
func (r WARCInfo) GetRecordID() string     { return r.WARCRecordId }
func (r WARCInfo) GetDate() time.Time      { return r.WARCDate }
func (r WARCInfo) GetContentLength() int64 { return r.ContentLength }
func (r WARCInfo) GetContent() io.Reader   { return bytes.NewReader(r.Content) }
func (r WARCInfo) Write(w io.Writer) error {
	err := writeHeader(w, r.Type(), map[string]string{
		warc_record_id:      r.WARCRecordId,
		warc_date:           r.WARCDate.Format(time.RFC3339),
		content_length:      int64String(r.ContentLength),
		content_type:        r.ContentType,
		warc_block_digest:   r.WARCBlockDigest,
		warc_payload_digest: r.WARCPayloadDigest,
		warc_truncated:      r.WARCTruncated,
		warc_filename:       r.WARCFilename,
	})
	if err != nil {
		return err
	}
	return writeBlock(w, r.Content)
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
	Content                   []byte
}

func (r Response) Type() RecordType        { return RecordTypeResponse }
func (r Response) GetRecordID() string     { return r.WARCRecordId }
func (r Response) GetDate() time.Time      { return r.WARCDate }
func (r Response) GetContentLength() int64 { return r.ContentLength }
func (r Response) GetContent() io.Reader   { return bytes.NewReader(r.Content) }
func (r Response) Write(w io.Writer) error {
	err := writeHeader(w, r.Type(), map[string]string{
		warc_record_id:               r.WARCRecordId,
		warc_date:                    r.WARCDate.Format(time.RFC3339),
		content_length:               int64String(r.ContentLength),
		content_type:                 r.ContentType,
		warc_concurrent_to:           r.WARCConcurrentTo,
		warc_block_digest:            r.WARCBlockDigest,
		warc_payload_digest:          r.WARCPayloadDigest,
		warc_ip_address:              r.WARCIPAddress,
		warc_target_uri:              r.WARCTargetURI,
		warc_truncated:               r.WARCTruncated,
		warc_warcinfo_id:             r.WARCWarcinfoID,
		warc_identified_payload_type: r.WARCIdentifiedPayloadType,
	})
	if err != nil {
		return err
	}
	return writeBlock(w, r.Content)
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
	Content                   []byte
}

func (r Resource) Type() RecordType        { return RecordTypeResource }
func (r Resource) GetRecordID() string     { return r.WARCRecordId }
func (r Resource) GetDate() time.Time      { return r.WARCDate }
func (r Resource) GetContentLength() int64 { return r.ContentLength }
func (r Resource) GetContent() io.Reader   { return bytes.NewReader(r.Content) }
func (r Resource) Write(w io.Writer) error {
	err := writeHeader(w, r.Type(), map[string]string{
		warc_record_id:               r.WARCRecordId,
		warc_date:                    r.WARCDate.Format(time.RFC3339),
		content_length:               int64String(r.ContentLength),
		content_type:                 r.ContentType,
		warc_concurrent_to:           r.WARCConcurrentTo,
		warc_block_digest:            r.WARCBlockDigest,
		warc_payload_digest:          r.WARCPayloadDigest,
		warc_ip_address:              r.WARCIPAddress,
		warc_target_uri:              r.WARCTargetURI,
		warc_truncated:               r.WARCTruncated,
		warc_warcinfo_id:             r.WARCWarcinfoID,
		warc_identified_payload_type: r.WARCIdentifiedPayloadType,
	})
	if err != nil {
		return err
	}
	return writeBlock(w, r.Content)
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
	Content                   []byte
}

func (r Request) Type() RecordType        { return RecordTypeRequest }
func (r Request) GetRecordID() string     { return r.WARCRecordId }
func (r Request) GetDate() time.Time      { return r.WARCDate }
func (r Request) GetContentLength() int64 { return r.ContentLength }
func (r Request) GetContent() io.Reader   { return bytes.NewReader(r.Content) }
func (r Request) Write(w io.Writer) error {
	err := writeHeader(w, r.Type(), map[string]string{
		warc_record_id:               r.WARCRecordId,
		warc_date:                    r.WARCDate.Format(time.RFC3339),
		content_length:               int64String(r.ContentLength),
		content_type:                 r.ContentType,
		warc_concurrent_to:           r.WARCConcurrentTo,
		warc_block_digest:            r.WARCBlockDigest,
		warc_payload_digest:          r.WARCPayloadDigest,
		warc_ip_address:              r.WARCIPAddress,
		warc_target_uri:              r.WARCTargetURI,
		warc_truncated:               r.WARCTruncated,
		warc_warcinfo_id:             r.WARCWarcinfoID,
		warc_identified_payload_type: r.WARCIdentifiedPayloadType,
	})
	if err != nil {
		return err
	}
	return writeBlock(w, r.Content)
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
	Content          []byte
}

func (r Metadata) Type() RecordType        { return RecordTypeMetadata }
func (r Metadata) GetRecordID() string     { return r.WARCRecordId }
func (r Metadata) GetDate() time.Time      { return r.WARCDate }
func (r Metadata) GetContentLength() int64 { return r.ContentLength }
func (r Metadata) GetContent() io.Reader   { return bytes.NewReader(r.Content) }
func (r Metadata) Write(w io.Writer) error {
	err := writeHeader(w, r.Type(), map[string]string{
		warc_record_id:     r.WARCRecordId,
		warc_date:          r.WARCDate.Format(time.RFC3339),
		content_length:     int64String(r.ContentLength),
		content_type:       r.ContentType,
		warc_concurrent_to: r.WARCConcurrentTo,
		warc_block_digest:  r.WARCBlockDigest,
		warc_ip_address:    r.WARCIPAddress,
		warc_refers_to:     r.WARCRefersTo,
		warc_target_uri:    r.WARCTargetURI,
		warc_truncated:     r.WARCTruncated,
		warc_warcinfo_id:   r.WARCWarcinfoID,
	})
	if err != nil {
		return err
	}
	return writeBlock(w, r.Content)
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
	Content           []byte
}

func (r Revisit) Type() RecordType        { return RecordTypeRevisit }
func (r Revisit) GetRecordID() string     { return r.WARCRecordId }
func (r Revisit) GetDate() time.Time      { return r.WARCDate }
func (r Revisit) GetContentLength() int64 { return r.ContentLength }
func (r Revisit) GetContent() io.Reader   { return bytes.NewReader(r.Content) }
func (r Revisit) Write(w io.Writer) error {
	err := writeHeader(w, r.Type(), map[string]string{
		warc_record_id:      r.WARCRecordId,
		warc_date:           r.WARCDate.Format(time.RFC3339),
		content_length:      int64String(r.ContentLength),
		content_type:        r.ContentType,
		warc_concurrent_to:  r.WARCConcurrentTo,
		warc_block_digest:   r.WARCBlockDigest,
		warc_payload_digest: r.WARCPayloadDigest,
		warc_ip_address:     r.WARCIPAddress,
		warc_refers_to:      r.WARCRefersTo,
		warc_target_uri:     r.WARCTargetURI,
		warc_truncated:      r.WARCTruncated,
		warc_warcinfo_id:    r.WARCWarcinfoID,
		warc_profile:        r.WARCProfile,
	})
	if err != nil {
		return err
	}
	return writeBlock(w, r.Content)
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
	Content           []byte
}

func (r Conversion) Type() RecordType        { return RecordTypeConversion }
func (r Conversion) GetRecordID() string     { return r.WARCRecordId }
func (r Conversion) GetDate() time.Time      { return r.WARCDate }
func (r Conversion) GetContentLength() int64 { return r.ContentLength }
func (r Conversion) GetContent() io.Reader   { return bytes.NewReader(r.Content) }
func (r Conversion) Write(w io.Writer) error {
	err := writeHeader(w, r.Type(), map[string]string{
		warc_record_id:      r.WARCRecordId,
		warc_date:           r.WARCDate.Format(time.RFC3339),
		content_length:      int64String(r.ContentLength),
		content_type:        r.ContentType,
		warc_block_digest:   r.WARCBlockDigest,
		warc_payload_digest: r.WARCPayloadDigest,
		warc_refers_to:      r.WARCRefersTo,
		warc_truncated:      r.WARCTruncated,
		warc_warcinfo_id:    r.WARCWarcinfoID,
	})
	if err != nil {
		return err
	}
	return writeBlock(w, r.Content)
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
	Content                []byte
}

func (r Continuation) Type() RecordType        { return RecordTypeContinuation }
func (r Continuation) GetRecordID() string     { return r.WARCRecordId }
func (r Continuation) GetDate() time.Time      { return r.WARCDate }
func (r Continuation) GetContentLength() int64 { return r.ContentLength }
func (r Continuation) GetContent() io.Reader   { return bytes.NewReader(r.Content) }
func (r Continuation) Write(w io.Writer) error {
	err := writeHeader(w, r.Type(), map[string]string{
		warc_record_id:            r.WARCRecordId,
		warc_date:                 r.WARCDate.Format(time.RFC3339),
		content_length:            int64String(r.ContentLength),
		warc_block_digest:         r.WARCBlockDigest,
		warc_payload_digest:       r.WARCPayloadDigest,
		warc_truncated:            r.WARCTruncated,
		warc_warcinfo_id:          r.WARCWarcinfoID,
		warc_segment_number:       intString(r.WARCSegmentNumber),
		warc_segment_origin_id:    r.WARCSegmentOriginID,
		warc_segment_total_length: int64String(r.WARCSegmentTotalLength),
	})
	if err != nil {
		return err
	}
	return writeBlock(w, r.Content)
}
