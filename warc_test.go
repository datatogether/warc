package warc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestWarcParse(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.warc")
	if err != nil {
		t.Error(err.Error())
		return
	}

	records, err := Parse(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
		return
	}

	if len(records) <= 0 {
		t.Errorf("recod length mismatch: %d isn't enough records", len(records))
		return
	}
}
