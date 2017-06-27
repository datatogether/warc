# warc
--
    import "github.com/archivers-space/warc"


## Usage

#### func  WriteRecords

```go
func WriteRecords(w io.Writer, records []Record) error
```
WriteRecords calls Write on each record to w

#### type Continuation

```go
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
```

Record blocks from 'continuation' records must be appended to corresponding
prior record block(s) (e.g., from other WARC files) to create the logically
complete full-sized original record. That is, 'continuation' records are used
when a record that would otherwise cause a WARC file size to exceed a desired
limit is broken into segments. A continuation record shall contain the named
fields 'WARC-Segment-Origin-ID' and 'WARC-Segment-Number', and the last
'continuation' record of a series shall contain a 'WARC-Segment-Total-Length'
field. The full details of WARC record segmentation are described in the below
section Record Segmentation. See also annex C.8 below for an example of a
‘continuation’ record.

#### func (Continuation) GetContent

```go
func (r Continuation) GetContent() io.Reader
```

#### func (Continuation) GetContentLength

```go
func (r Continuation) GetContentLength() int64
```

#### func (Continuation) GetDate

```go
func (r Continuation) GetDate() time.Time
```

#### func (Continuation) GetRecordID

```go
func (r Continuation) GetRecordID() string
```

#### func (Continuation) Type

```go
func (r Continuation) Type() RecordType
```

#### func (Continuation) Write

```go
func (r Continuation) Write(w io.Writer) error
```

#### type Conversion

```go
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
```

A 'conversion' record shall contain an alternative version of another record's
content that was created as the result of an archival process. Typically, this
is used to hold content transformations that maintain viability of content after
widely available rendering tools for the originally stored format disappear. As
needed, the original content may be migrated (transformed) to a more viable
format in order to keep the information usable with current tools while
minimizing loss of information (intellectual content, look and feel, etc). Any
number of 'conversion' records may be created that reference a specific source
record, which may itself contain transformed content. Each transformation should
result in a freestanding, complete record, with no dependency on survival of the
original record. Metadata records may be used to further describe transformation
records. Wherever practical, a 'conversion' record should contain a
'WARC-Refers-To' field to identify the prior material converted.

#### func (Conversion) GetContent

```go
func (r Conversion) GetContent() io.Reader
```

#### func (Conversion) GetContentLength

```go
func (r Conversion) GetContentLength() int64
```

#### func (Conversion) GetDate

```go
func (r Conversion) GetDate() time.Time
```

#### func (Conversion) GetRecordID

```go
func (r Conversion) GetRecordID() string
```

#### func (Conversion) Type

```go
func (r Conversion) Type() RecordType
```

#### func (Conversion) Write

```go
func (r Conversion) Write(w io.Writer) error
```

#### type Metadata

```go
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
```

A 'metadata' record contains content created in order to further describe,
explain, or accompany a harvested resource, in ways not covered by other record
types. A 'metadata' record will almost always refer to another record of another
type, with that other record holding original harvested or transformed content.
(However, it is allowable for a 'metadata' record to refer to any record type,
including other 'metadata' records.) Any number of metadata records may
reference one specific other record. The format of the metadata record block may
vary. The "application/warc-fields" format, defined earlier, may be used.
Allowable fields include all \[DCMI\] plus the following field definitions. All
fields are optional.

#### func (Metadata) GetContent

```go
func (r Metadata) GetContent() io.Reader
```

#### func (Metadata) GetContentLength

```go
func (r Metadata) GetContentLength() int64
```

#### func (Metadata) GetDate

```go
func (r Metadata) GetDate() time.Time
```

#### func (Metadata) GetRecordID

```go
func (r Metadata) GetRecordID() string
```

#### func (Metadata) Type

```go
func (r Metadata) Type() RecordType
```

#### func (Metadata) Write

```go
func (r Metadata) Write(w io.Writer) error
```

#### type Reader

```go
type Reader struct {
}
```

Reader parses WARC records from an underlying scanner. Create a new reader with
NewReader

#### func  NewReader

```go
func NewReader(r io.Reader) *Reader
```
NewReader creates a new WARC reader from an io.Reader Always use NewReader,
(instead of manually allocating a reader)

#### func (*Reader) Read

```go
func (r *Reader) Read() (Record, error)
```
Read a record, will return nil, io.EOF to signal no more records

#### func (*Reader) ReadAll

```go
func (r *Reader) ReadAll() (records []Record, err error)
```
Consume the entire reader, returning a slice of records

#### type Record

```go
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
```

Record is the common interface for all WARC Record Types A Record consists of a
version indicator (eg: WARC/1.0), zero or more headers, and possibly a content
block. Upgrades to specific types of records can be done using type assertions
and/or the Type method.

#### type RecordType

```go
type RecordType int
```

RecordType enumerates different types of WARC Records

```go
const (
	RecordTypeUnknown RecordType = iota
	RecordTypeWarcInfo
	RecordTypeResponse
	RecordTypeResource
	RecordTypeRequest
	RecordTypeMetadata
	RecordTypeRevisit
	RecordTypeConversion
	RecordTypeContinuation
)
```

#### func (RecordType) String

```go
func (r RecordType) String() string
```
RecordType satisfies the stringer interface

#### type Records

```go
type Records []Record
```

A WARC format file is the simple concatenation of one or more WARC records. The
first record usually describes the records to follow. In general, record content
is either the direct result of a retrieval attempt — web pages, inline images,
URL redirection information, DNS hostname lookup results, standalone files, etc.
— or is synthesized material (e.g., metadata, transformed content) that provides
additional information about archived content.

#### type Request

```go
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
```

A 'request' record holds the details of a complete scheme-specific request,
including network protocol information where possible. The exact contents of a
'request' record are determined not just by the record type but also by the URI
scheme of the record's target-URI, as described below.

#### func (Request) GetContent

```go
func (r Request) GetContent() io.Reader
```

#### func (Request) GetContentLength

```go
func (r Request) GetContentLength() int64
```

#### func (Request) GetDate

```go
func (r Request) GetDate() time.Time
```

#### func (Request) GetRecordID

```go
func (r Request) GetRecordID() string
```

#### func (Request) Type

```go
func (r Request) Type() RecordType
```

#### func (Request) Write

```go
func (r Request) Write(w io.Writer) error
```

#### type Resource

```go
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
```

A 'resource' record contains a resource, without full protocol response
information. For example: a file directly retrieved from a locally accessible
repository or the result of a networked retrieval where the protocol information
has been discarded. The exact contents of a 'resource' record are determined not
just by the record type but also by the URI scheme of the record's target-URI,
as described below. For all 'resource' records, the payload is defined as the
record block. A 'resource' record, with a synthesized target-URI, may also be
used to archive other artefacts of a harvesting process inside WARC files.

#### func (Resource) GetContent

```go
func (r Resource) GetContent() io.Reader
```

#### func (Resource) GetContentLength

```go
func (r Resource) GetContentLength() int64
```

#### func (Resource) GetDate

```go
func (r Resource) GetDate() time.Time
```

#### func (Resource) GetRecordID

```go
func (r Resource) GetRecordID() string
```

#### func (Resource) Type

```go
func (r Resource) Type() RecordType
```

#### func (Resource) Write

```go
func (r Resource) Write(w io.Writer) error
```

#### type Response

```go
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
```

A 'response' record should contain a complete scheme-specific response,
including network protocol information where possible. The exact contents of a
'response' record are determined not just by the record type but also by the URI
scheme of the record's target-URI, as described below.

#### func (Response) GetContent

```go
func (r Response) GetContent() io.Reader
```

#### func (Response) GetContentLength

```go
func (r Response) GetContentLength() int64
```

#### func (Response) GetDate

```go
func (r Response) GetDate() time.Time
```

#### func (Response) GetRecordID

```go
func (r Response) GetRecordID() string
```

#### func (Response) Type

```go
func (r Response) Type() RecordType
```

#### func (Response) Write

```go
func (r Response) Write(w io.Writer) error
```

#### type Revisit

```go
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
```

A 'revisit' record describes the revisitation of content already archived, and
might include only an abbreviated content body which has to be interpreted
relative to a previous record. Most typically, a 'revisit' record is used
instead of a 'response' or 'resource' record to indicate that the content
visited was either a complete or substantial duplicate of material previously
archived. Using a 'revisit' record instead of another type is optional, for when
benefits of reduced storage size or improved cross-referencing of material are
desired.

#### func (Revisit) GetContent

```go
func (r Revisit) GetContent() io.Reader
```

#### func (Revisit) GetContentLength

```go
func (r Revisit) GetContentLength() int64
```

#### func (Revisit) GetDate

```go
func (r Revisit) GetDate() time.Time
```

#### func (Revisit) GetRecordID

```go
func (r Revisit) GetRecordID() string
```

#### func (Revisit) Type

```go
func (r Revisit) Type() RecordType
```

#### func (Revisit) Write

```go
func (r Revisit) Write(w io.Writer) error
```

#### type WARCInfo

```go
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
```

A 'warcinfo' record describes the records that follow it, up through end of
file, end of input, or until next 'warcinfo' record. Typically, this appears
once and at the beginning of a WARC file. For a web archive, it often contains
information about the web crawl which generated the following records. The
format of this descriptive record block may vary, though the use of the
"application/warc-fields" content-type is recommended. Allowable fields include,
but are not limited to, all \[DCMI\] plus the following field definitions. All
fields are optional.

#### func (WARCInfo) GetContent

```go
func (r WARCInfo) GetContent() io.Reader
```

#### func (WARCInfo) GetContentLength

```go
func (r WARCInfo) GetContentLength() int64
```

#### func (WARCInfo) GetDate

```go
func (r WARCInfo) GetDate() time.Time
```

#### func (WARCInfo) GetRecordID

```go
func (r WARCInfo) GetRecordID() string
```

#### func (WARCInfo) Type

```go
func (r WARCInfo) Type() RecordType
```

#### func (WARCInfo) Write

```go
func (r WARCInfo) Write(w io.Writer) error
```
