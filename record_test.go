package warc

import (
	"testing"
)

func TestRecordId(t *testing.T) {
	r := &Record{}
	if r.Id() != "" {
		t.Errorf("id mismatch. expected '', got: '%s'", r.Id())
	}
}
