package warc

import (
	"net/textproto"
)

// Header mimics net/http's header package,
// but with single values, and provides exceptions
// for capitalized "WARC" header keys. The WARC 1.0 spec
// calls for case-insensitive header keys, but the spec
// token diagrams list headers as being case-sensitive,
// so we'll honor case any case on read, but write records
// that match the spec token diagrams.
type Header map[string]string

func (h Header) Get(key string) string {
	return h[CanonicalKey(key)]
}

func (h Header) Set(key, value string) {
	h[CanonicalKey(key)] = value
}

func CanonicalKey(key string) string {
	key = textproto.CanonicalMIMEHeaderKey(key)
	if warcFieldMimeMap[key] != "" {
		key = warcFieldMimeMap[key]
	}
	return key
}

var warcFieldMimeMap = map[string]string{
	"Warc-Record-Id":               FieldNameWARCRecordID,
	"Warc-Date":                    FieldNameWARCDate,
	"Warc-Type":                    FieldNameWARCType,
	"Warc-Concurrent-To":           FieldNameWARCConcurrentTo,
	"Warc-Block-Digest":            FieldNameWARCBlockDigest,
	"Warc-Payload-Digest":          FieldNameWARCPayloadDigest,
	"Warc-Ip-Address":              FieldNameWARCIPAddress,
	"Warc-Refers-To":               FieldNameWARCRefersTo,
	"Warc-Target-Uri":              FieldNameWARCTargetURI,
	"Warc-Truncated":               FieldNameWARCTruncated,
	"Warc-Warcinfo-Id":             FieldNameWARCWarcinfoID,
	"Warc-Filename":                FieldNameWARCFilename,
	"Warc-Profile":                 FieldNameWARCProfile,
	"Warc-Identified-Payload-Type": FieldNameWARCIdentifiedPayloadType,
	"Warc-Segment-Number":          FieldNameWARCSegmentNumber,
	"Warc-Segment-Origin-Id":       FieldNameWARCSegmentOriginID,
	"Warc-Segment-Total-Length":    FieldNameWARCSegmentTotalLength,
}
