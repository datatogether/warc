package rewrite

import (
	"bytes"
	"testing"

	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

func TestRewriteHtml(t *testing.T) {
	cases := []struct {
		in, out []byte
		err     error
	}{
		{noChange, noChange, nil},
	}

	for i, c := range cases {
		rw := NewHtmlRewriter()
		got, err := rw.Rewrite(c.in)
		if err != nil && err != c.err {
			t.Errorf("case %d error mismatch: %s != %s", i, err, c.err)
			continue
		}

		if !bytes.Equal(got, c.out) {
			dmp := dmp.New()
			diffs := dmp.DiffMain(string(got), string(c.out), true)

			t.Errorf("case %d mismatch:\n%s", i, dmp.DiffPrettyText(diffs))
		}
	}
}

var noChange = []byte(`<!DOCTYPE html>
<html>
<head>
  <title></title>
</head>
<body>

</body>
</html>`)
