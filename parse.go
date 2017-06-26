package warc

import (
	"fmt"
	"strconv"
	"time"
)

const (
	content_length               = "content-length"
	content_type                 = "content-type"
	warc_block_digest            = "warc-block-digest"
	warc_concurrent_to           = "warc-concurrent-to"
	warc_filename                = "warc-filename"
	warc_date                    = "warc-date"
	warc_identified_payload_type = "warc-identified-payload-type"
	warc_ip_address              = "warc-ip-address"
	warc_payload_digest          = "warc-payload-digest"
	warc_profile                 = "warc-profile"
	warc_record_id               = "warc-record-id"
	warc_refers_to               = "warc-refers-to"
	warc_segment_origin_id       = "warc-segment-origin-id"
	warc_segment_number          = "warc-segment-number"
	warc_segment_total_length    = "warc-segment-total-length"
	warc_target_uri              = "warc-target-uri"
	warc_truncated               = "warc-truncated"
	warc_type                    = "warc-type"
	warc_warcinfo_id             = "warc-warcinfo-id"
)

func recordType(headers map[string]string) (t RecordType) {
	switch headers[warc_type] {
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
	}
	return
}

func newRecord(h map[string]string, content []byte) (Record, error) {
	warcDate, err := time.Parse(time.RFC3339, h[warc_date])
	if err != nil {
		return nil, err
	}

	contentLength, err := strconv.ParseInt(h[content_length], 10, 64)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(content))
	// fmt.Println("* * * * * * * * *")

	switch recordType(h) {
	case RecordTypeWarcInfo:
		return &WARCInfo{
			WARCRecordId:      h[warc_record_id],
			WARCDate:          warcDate,
			ContentLength:     contentLength,
			ContentType:       h[content_type],
			WARCBlockDigest:   h[warc_block_digest],
			WARCPayloadDigest: h[warc_payload_digest],
			WARCTruncated:     h[warc_truncated],
			WARCFilename:      h[warc_filename],
			Content:           content,
		}, nil
	case RecordTypeResponse:
		return &Response{
			WARCRecordId:              h[warc_record_id],
			WARCDate:                  warcDate,
			ContentLength:             contentLength,
			ContentType:               h[content_type],
			WARCConcurrentTo:          h[warc_concurrent_to],
			WARCBlockDigest:           h[warc_block_digest],
			WARCPayloadDigest:         h[warc_payload_digest],
			WARCIPAddress:             h[warc_ip_address],
			WARCTargetURI:             h[warc_target_uri],
			WARCTruncated:             h[warc_truncated],
			WARCWarcinfoID:            h[warc_warcinfo_id],
			WARCIdentifiedPayloadType: h[warc_identified_payload_type],
			Content:                   content,
		}, nil
	case RecordTypeResource:
		return &Resource{
			WARCRecordId:              h[warc_record_id],
			WARCDate:                  warcDate,
			ContentLength:             contentLength,
			ContentType:               h[content_type],
			WARCConcurrentTo:          h[warc_concurrent_to],
			WARCPayloadDigest:         h[warc_payload_digest],
			WARCBlockDigest:           h[warc_block_digest],
			WARCIPAddress:             h[warc_ip_address],
			WARCTargetURI:             h[warc_target_uri],
			WARCTruncated:             h[warc_truncated],
			WARCWarcinfoID:            h[warc_warcinfo_id],
			WARCIdentifiedPayloadType: h[warc_identified_payload_type],
			Content:                   content,
		}, nil
	case RecordTypeRequest:
		return &Request{
			WARCRecordId:              h[warc_record_id],
			WARCDate:                  warcDate,
			ContentLength:             contentLength,
			ContentType:               h[content_type],
			WARCConcurrentTo:          h[warc_concurrent_to],
			WARCBlockDigest:           h[warc_block_digest],
			WARCPayloadDigest:         h[warc_payload_digest],
			WARCIPAddress:             h[warc_ip_address],
			WARCTargetURI:             h[warc_target_uri],
			WARCTruncated:             h[warc_truncated],
			WARCWarcinfoID:            h[warc_warcinfo_id],
			WARCIdentifiedPayloadType: h[warc_identified_payload_type],
			Content:                   content,
		}, nil
	case RecordTypeMetadata:
		return &Metadata{
			WARCRecordId:     h[warc_record_id],
			WARCDate:         warcDate,
			ContentLength:    contentLength,
			ContentType:      h[content_type],
			WARCConcurrentTo: h[warc_concurrent_to],
			WARCBlockDigest:  h[warc_block_digest],
			WARCIPAddress:    h[warc_ip_address],
			WARCRefersTo:     h[warc_refers_to],
			WARCTargetURI:    h[warc_target_uri],
			WARCTruncated:    h[warc_truncated],
			WARCWarcinfoID:   h[warc_warcinfo_id],
			Content:          content,
		}, nil
	case RecordTypeRevisit:
		return &Revisit{
			WARCRecordId:      h[warc_record_id],
			WARCDate:          warcDate,
			ContentLength:     contentLength,
			ContentType:       h[content_type],
			WARCConcurrentTo:  h[warc_concurrent_to],
			WARCBlockDigest:   h[warc_block_digest],
			WARCPayloadDigest: h[warc_payload_digest],
			WARCIPAddress:     h[warc_ip_address],
			WARCRefersTo:      h[warc_refers_to],
			WARCTargetURI:     h[warc_target_uri],
			WARCTruncated:     h[warc_truncated],
			WARCWarcinfoID:    h[warc_warcinfo_id],
			WARCProfile:       h[warc_profile],
			Content:           content,
		}, nil
	case RecordTypeConversion:
		return &Conversion{
			WARCRecordId:      h[warc_record_id],
			WARCDate:          warcDate,
			ContentLength:     contentLength,
			ContentType:       h[content_type],
			WARCBlockDigest:   h[warc_block_digest],
			WARCPayloadDigest: h[warc_payload_digest],
			WARCRefersTo:      h[warc_refers_to],
			WARCTruncated:     h[warc_truncated],
			WARCWarcinfoID:    h[warc_warcinfo_id],
			Content:           content,
		}, nil
	case RecordTypeContinuation:
		seg, err := strconv.ParseInt(h[warc_segment_number], 10, 0)
		if err != nil {
			return nil, fmt.Errorf("error parsing WARCSegmentNumber: %s", err.Error())
		}
		length, err := strconv.ParseInt(h[warc_segment_total_length], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing WARCSegmentTotalLength: %s", err.Error())
		}
		return &Continuation{
			WARCRecordId:           h[warc_record_id],
			WARCDate:               warcDate,
			ContentLength:          contentLength,
			WARCBlockDigest:        h[warc_block_digest],
			WARCPayloadDigest:      h[warc_payload_digest],
			WARCTruncated:          h[warc_truncated],
			WARCWarcinfoID:         h[warc_warcinfo_id],
			WARCSegmentNumber:      int(seg),
			WARCSegmentOriginID:    h[warc_segment_origin_id],
			WARCSegmentTotalLength: length,
			Content:                content,
		}, nil
	default:
		return nil, fmt.Errorf("unrecognized record format")
	}
}
