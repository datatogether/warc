package warc

import (
	"bytes"
	"os"
	"testing"
)

func TestRecordID(t *testing.T) {
	r := &Record{}
	if r.ID() != "" {
		t.Errorf("id mismatch. expected '', got: '%s'", r.ID())
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

	body, err := records[1].Body()
	if err != nil {
		t.Error(err)
		return
	}

	if bytes.HasPrefix(body, crlf) {
		t.Errorf("content shouldn't have CRLF prefix")
		return
	}

	if bytes.HasSuffix(body, crlf) {
		t.Errorf("content shouldn't have CRLF suffix")
		return
	}

	// if len(body) != records[1].ContentLength() {
	// 	t.Errorf("content length mistmatch: %d != %d", records[1].ContentLength(), len(body))
	// 	return
	// }
	// fmt.Println(string(b))
}
