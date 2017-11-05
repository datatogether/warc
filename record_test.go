package warc

import (
	"os"
	"testing"
)

func TestRecordId(t *testing.T) {
	r := &Record{}
	if r.Id() != "" {
		t.Errorf("id mismatch. expected '', got: '%s'", r.Id())
	}
}

func TestRecordBody(t *testing.T) {
	// TODO
	f, err := os.Open("testdata/response.warc")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer f.Close()

	rdr, err := NewReader(f)
	if err != nil {
		t.Error(err)
		return
	}

	records, err := rdr.ReadAll()
	if err != nil {
		t.Error(err)
		return
	}
	// fmt.Println(records[1].Content.String())

	_, err = records[1].Body()
	if err != nil {
		t.Error(err)
		return
	}

	// fmt.Println(string(b))
}
