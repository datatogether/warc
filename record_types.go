package warc

// RecordType enumerates different types of WARC Records
type RecordType int

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
		return "unknown"
	}
}
