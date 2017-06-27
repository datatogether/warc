package warc

import (
	"os"
	"testing"
)

func TestWarcWrite(t *testing.T) {
	f, err := os.Open("testdata/test.warc")
	// data, err := ioutil.ReadFile("testdata/test.warc")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer f.Close()

	records, err := NewReader(f).ReadAll()
	if err != nil {
		t.Error(err)
		return
	}

	out, err := os.Create("testdata/out.warc")
	if err != nil {
		t.Error(err)
		return
	}
	defer out.Close()
	if err := WriteRecords(out, records); err != nil {
		t.Error(err)
		return
	}

}
