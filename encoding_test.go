package warc

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestUnmarshalRecord(t *testing.T) {
	f, err := os.Open("testdata/test.warc")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err.Error())
		return
	}

	rec, err := UnmarshalRecord(data)
	if err != nil {
		t.Errorf("unexpected UnmarshalRecord error: %s", err.Error())
		return
	}

	if rec.Type != RecordTypeResource {
		t.Errorf("wrong record type. expected: %s, got: %s", RecordTypeResource, rec.Type)
	}
}

func TestUnmarshalRecords(t *testing.T) {
	f, err := os.Open("testdata/test.warc")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err.Error())
		return
	}

	recs, err := UnmarshalRecords(data)
	if err != nil {
		t.Errorf("unexpected UnmarshalRecord error: %s", err.Error())
		return
	}

	if len(recs) != 10 {
		t.Errorf("wrong number of records. expected: %d, got: %d", 10, len(recs))
	}
}
