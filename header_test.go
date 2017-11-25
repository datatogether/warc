package warc

import (
	"testing"
)

func TestHeader(t *testing.T) {
	h := Header{}
	if h.Get("") != "" {
		t.Errorf("expected empty string for empty string get")
		return
	}
	h.Set("warc-record-id", "test_id")
	if h.Get("WARC-Record-ID") != "test_id" {
		t.Errorf("expected get WARC-Record-ID to return %s", "test_id")
		return
	}
}

func TestCanonicalKey(t *testing.T) {
	cases := []struct {
		in, expect string
	}{
		{"warc-record-id", FieldNameWARCRecordID},
		{"WARC-DATE", FieldNameWARCDate},
		{"Warc-TYPE", FieldNameWARCType},
		{"warc-CONCURRENt-to", FieldNameWARCConcurrentTo},
		{"warC-block-digest", FieldNameWARCBlockDigest},
		{"Warc-payload-Digest", FieldNameWARCPayloadDigest},
		{"warc-ip-Address", FieldNameWARCIPAddress},
		{"warc-refers-To", FieldNameWARCRefersTo},
		{"warc-target-Uri", FieldNameWARCTargetURI},
		{"warc-truncated", FieldNameWARCTruncated},
		{"warc-warcinfo-Id", FieldNameWARCWarcinfoID},
		{"warc-filename", FieldNameWARCFilename},
		{"warc-profile", FieldNameWARCProfile},
		{"warc-identified-payload-Type", FieldNameWARCIdentifiedPayloadType},
		{"warc-segment-Number", FieldNameWARCSegmentNumber},
		{"warc-segment-origin-Id", FieldNameWARCSegmentOriginID},
		{"warc-segment-total-Length", FieldNameWARCSegmentTotalLength},
	}

	for i, c := range cases {
		got := CanonicalKey(c.in)
		if got != c.expect {
			t.Errorf("case %d mismatch. expected: '%s', got: '%s'", i, c.expect, got)
			continue
		}
	}
}
