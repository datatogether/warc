package warc

// Beginning of a WARC record, consisting of one first line declaring the
// record to be in the WARC format with a given version number, followed by
// lines of named fields up to a blank line.
type Header struct {
	Version                   string
	WARCType                  RecordType
	WARCRecordId              string
	WARCDate                  string
	ContentLength             string
	ContentType               string
	WARCConcurrentTo          string
	WARCBlockDigest           string
	WARCPayloadDigest         string
	WARCIPAddress             string
	WARCRefersTo              string
	WARCTargetUri             string
	WARCTruncated             string
	WARCInfoId                string
	WARCFilename              string
	WARCProfile               string
	WARCIdentifiedPayloadType string
	WARCSegmentOriginID       string
	WARCSegmentNumber         string
	WARCSegmentTotalLength    string
	CustomFields              map[string]string
}

type NamedField struct {
	Name  string
	Value string
}
