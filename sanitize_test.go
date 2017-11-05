package warc

import (
	"bytes"
	"testing"

	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

func TestSanitize(t *testing.T) {
	wordDoc, err := readTestFile("test-doc.docx")
	if err != nil {
		t.Error(err.Error())
		return
	}

	cases := []struct {
		mime   string
		body   []byte
		expect []byte
	}{
		{"application/zip", wordDoc, wordDoc},
	}

	for i, c := range cases {
		got := Sanitize(c.mime, c.body)

		if !bytes.Equal(got, c.expect) {
			dmp := dmp.New()
			diffs := dmp.DiffMain(string(c.expect), string(got), true)
			if len(diffs) == 0 {
				t.Logf("case %d bytes were unequal but computed no difference between results")
				continue
			}

			t.Error("case %d byte mismatch")
			continue
		}

		if bytes.Contains(got, doubleCrlf) {
			t.Errorf("case %d can't contain crlf", i)
			continue
		}
	}
}
