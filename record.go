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
	Id() string
	Date() time.Time
	Header() *Header
	ContentLength() int64
	Content() io.Reader
}

type record struct {
	Version string
	Header  *Header
	Content io.Reader
}

// WARCInfo record type
type Info struct {
	Header  *Header
	Content io.Reader
}

func (i Info) Type() RecordType { return RecordTypeWarcInfo }

type Response struct {
	Header  *Header
	Content io.Reader
}

func (r Response) Type() RecordType { return RecordTypeResponse }

type Resource struct {
	Header  *Header
	Content io.Reader
}

func (r Resource) Type() RecordType { return RecordTypeResource }

type Request struct {
	Header  *Header
	Content io.Reader
}

func (r Request) Type() RecordType { return RecordTypeRequest }

type Metadata struct {
	Header  *Header
	Content io.Reader
}

func (r Metadata) Type() RecordType { return RecordTypeMetadata }

type Revisit struct {
	Header  *Header
	Content io.Reader
}

func (r Revisit) Type() RecordType { return RecordTypeRevisit }

type Conversion struct {
	Header  *Header
	Content io.Reader
}

func (r Conversion) Type() RecordType { return RecordTypeConversion }

type Continuation struct {
	Header  *Header
	Content io.Reader
}

func (r Continuation) Type() RecordType { return RecordTypeContinuation }
