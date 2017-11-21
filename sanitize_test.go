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
	wordDocGz, err := readTestFile("test-doc.docx.gz")
	if err != nil {
		t.Error(err.Error())
		return
	}

	cases := []struct {
		mime   string
		body   []byte
		expect []byte
		err    string
	}{
		{"application/zip", wordDoc, wordDocGz, ""},
	}

	for i, c := range cases {
		got, err := Sanitize(c.mime, c.body)

		if !(err == nil && c.err == "" || err != nil && err.Error() == c.err) {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, err.Error())
			continue
		}

		if !bytes.Equal(got, c.expect) {
			dmp := dmp.New()
			diffs := dmp.DiffMain(string(c.expect), string(got), true)
			if len(diffs) == 0 {
				t.Logf("case %d bytes were unequal but computed no difference between results", i)
				continue
			}
			t.Errorf("case %d mismatch:\n%s", i, dmp.DiffPrettyText(diffs))
			continue
		}

		if bytes.Contains(got, doubleCrlf) {
			t.Errorf("case %d can't contain double crlf", i)
			continue
		}
	}
}
