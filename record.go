package warc

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"time"
)

// RecordType enumerates different types of WARC Records
type RecordType int

const (
	// RecordTypeUnknown is the default type of record, which shouldn't be
	// accepted by anything that wants to know a type of record.
	RecordTypeUnknown RecordType = iota
	// RecordTypeWarcInfo describes the records that follow it, up through end
	// of file, end of input, or until next 'warcinfo' record. Typically, this
	// appears once and at the beginning of a WARC file. For a web archive, it
	// often contains information about the web crawl which generated the
	// following records.
	// The format of this descriptive record block may vary, though the use of
	// the "application/warc-fields" content-type is recommended. Allowable
	// fields include, but are not limited to, all \[DCMI\] plus the following
	// field definitions. All fields are optional.
	RecordTypeWarcInfo
	// RecordTypeResponse should contain a complete scheme-specific response,
	// including network protocol information where possible. The exact
	// contents of a 'response' record are determined not just by the record
	// type but also by the URI scheme of the record's target-URI, as described
	// below.
	RecordTypeResponse
	// RecordTypeResource contains a resource, without full protocol response
	// information. For example: a file directly retrieved from a locally
	// accessible repository or the result of a networked retrieval where the
	// protocol information has been discarded. The exact contents of a
	// 'resource' record are determined not just by the record type but also by
	// the URI scheme of the record's target-URI, as described below.
	// For all 'resource' records, the payload is defined as the record block.
	// A 'resource' record, with a synthesized target-URI, may also be used to
	// archive other artefacts of a harvesting process inside WARC files.
	RecordTypeResource
	// RecordTypeRequest holds the details of a complete scheme-specific
	// request, including network protocol information where possible. The
	// exact contents of a 'request' record are determined not just by the
	// record type but also by the URI scheme of the record's target-URI, as
	// described below.
	RecordTypeRequest
	// RecordTypeMetadata contains content created in order to further
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
	RecordTypeMetadata
	// RecordTypeRevisit describes the revisitation of content already
	// archived, and might include only an abbreviated content body which has
	// to be interpreted relative to a previous record. Most typically, a
	// 'revisit' record is used instead of a 'response' or 'resource' record to
	// indicate that the content visited was either a complete or substantial
	// duplicate of material previously archived.
	// Using a 'revisit' record instead of another type is optional, for when
	// benefits of reduced storage size or improved cross-referencing of
	// material are desired.
	RecordTypeRevisit
	// RecordTypeConversion shall contain an alternative version of another
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
	RecordTypeConversion
	// RecordTypeContinuation blocks from 'continuation' records must be appended to
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
	RecordTypeContinuation
)

// RecordType satisfies the stringer interface
func (r RecordType) String() string {
	switch r {
	case RecordTypeWarcInfo:
		return "warcinfo"
	case RecordTypeResponse:
		return "response"
	case RecordTypeResource:
		return "resource"
	case RecordTypeRequest:
		return "request"
	case RecordTypeMetadata:
		return "metadata"
	case RecordTypeRevisit:
		return "revisit"
	case RecordTypeConversion:
		return "conversion"
	case RecordTypeContinuation:
		return "continuation"
	default:
		return ""
	}
}

// ParseRecordType parses a RecordType from a string
func ParseRecordType(s string) RecordType {
	switch s {
	case RecordTypeWarcInfo.String():
		return RecordTypeWarcInfo
	case RecordTypeResponse.String():
		return RecordTypeResponse
	case RecordTypeResource.String():
		return RecordTypeResource
	case RecordTypeRequest.String():
		return RecordTypeRequest
	case RecordTypeMetadata.String():
		return RecordTypeMetadata
	case RecordTypeRevisit.String():
		return RecordTypeRevisit
	case RecordTypeConversion.String():
		return RecordTypeConversion
	case RecordTypeContinuation.String():
		return RecordTypeContinuation
	default:
		return RecordTypeUnknown
	}
}

// A Record consists of a version indicator (eg: WARC/1.0), zero or more headers,
// and possibly a content block.
// Upgrades to specific types of records can be done using type assertions
// and/or the Type method.
type Record struct {
	Format  RecordFormat
	Type    RecordType
	Headers Header
	Content *bytes.Buffer
}

// ID gives The ID for this record
func (r *Record) ID() string {
	return strings.TrimSuffix(strings.TrimPrefix(r.Headers[FieldNameWARCRecordID], "<urn:uuid:"), ">")
}

// TargetURI is a convenience method for getting the uri
// that this record is targeting
func (r *Record) TargetURI() string {
	return r.Headers[FieldNameWARCTargetURI]
}

// Date gives the time.Time of record creation, returns empty (zero) time if
// no Warc-Date header is present, or if the header is an
// invalid timestamp
func (r *Record) Date() time.Time {
	t, err := time.Parse(time.RFC3339, r.Headers[FieldNameWARCDate])
	if err != nil {
		return time.Time{}
	}
	return t
}

// ContentLength of content block in bytes, returns 0 if
// Content-Length header is missing or invalid
func (r *Record) ContentLength() int {
	len, err := strconv.ParseInt(r.Headers[FieldNameContentLength], 10, 64)
	if err != nil {
		return 0
	}
	return int(len)
}

// Write this record to the given writer.
func (r *Record) Write(w io.Writer) error {
	r.Headers[FieldNameContentLength] = strconv.FormatInt(int64(r.Content.Len()), 10)
	r.Headers[FieldNameWARCType] = r.Type.String()
	switch r.Type {
	case RecordTypeResponse, RecordTypeRevisit:
		r.Headers[FieldNameWARCBlockDigest] = Sha1Digest(r.Content.Bytes())
	}

	if err := writeHeader(w, r); err != nil {
		return err
	}
	return writeBlock(w, bytes.NewReader(r.Content.Bytes()))
}

// Bytes returns the record formatted as a byte slice
func (r *Record) Bytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := r.Write(buf)
	return buf.Bytes(), err
}

// Body returns a record's body with any HTTP headers omitted
func (r *Record) Body() ([]byte, error) {
	// TODO - actually remove headers
	// buf := &bytes.Buffer{}
	// err := writeBlock(buf, r.Content)
	return readBlockBody(r.Content.Bytes())
}

// SetBody sets the body of the record, leaving any written
// http headers in record
func (r *Record) SetBody(body []byte) error {
	repl, err := replaceBlockBody(r.Content.Bytes(), body)
	if err != nil {
		return err
	}
	r.Content = bytes.NewBuffer(repl)
	return nil
}

// RecordFormat determines different formats for records, this is
// for any later support of ARC files, should we need to add it.
type RecordFormat int

const (
	// RecordFormatWarc default is the Warc Format 1.0
	RecordFormatWarc RecordFormat = iota
	// RecordFormatUnknown reporesents unknown / errored record format
	RecordFormatUnknown
)

func (r RecordFormat) String() string {
	switch r {
	case RecordFormatWarc:
		return "WARC/1.0"
	default:
		return ""
	}
}

func recordFormat(s string) RecordFormat {
	switch s {
	case "WARC/1.0":
		return RecordFormatWarc
	default:
		return RecordFormatUnknown
	}
}
