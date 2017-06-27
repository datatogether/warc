package warc

import (
	"os"
	"testing"
)

func TestReadAll(t *testing.T) {
	f, err := os.Open("testdata/test.warc")
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

	if len(records) <= 0 {
		t.Errorf("record length mismatch: %d isn't enough records", len(records))
		return
	}

	// for _, r := range records {
	// 	fmt.Println(r.Type().String())
	// }
}
