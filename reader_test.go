package warc

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestReadAll(t *testing.T) {
	f, err := os.Open("testdata/test.warc")
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

	if len(records) <= 0 {
		t.Errorf("record length mismatch: %d isn't enough records", len(records))
		return
	}

	// for _, r := range records {
	// 	fmt.Println(r.Type().String())
	// }
}

func readTestFile(path string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join("testdata", path))
}
