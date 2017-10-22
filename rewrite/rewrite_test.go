package rewrite

import (
	"bytes"
	"testing"

	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

type rewriteTestCase struct {
	in, out []byte
	err     error
}

type stringTestCase struct {
	in, out string
	err     error
}

func stringTestCases(in []stringTestCase) (cases []rewriteTestCase) {
	for _, c := range in {
		cases = append(cases, rewriteTestCase{
			in:  []byte(c.in),
			out: []byte(c.out),
			err: c.err,
		})
	}
	return
}

func testRewriteCases(t *testing.T, rw Rewriter, cases []rewriteTestCase) {
	for i, c := range cases {
		got, err := rw.Rewrite(c.in)
		if err != nil && err != c.err {
			t.Errorf("case %d error mismatch: %s != %s", i, err, c.err)
			continue
		}

		if !bytes.Equal(got, c.out) {
			dmp := dmp.New()
			diffs := dmp.DiffMain(string(got), string(c.out), true)

			t.Errorf("case %d mismatch:\n%s", i, dmp.DiffPrettyText(diffs))
			t.Errorf("expected: %s, got: %s", string(c.out), string(got))
		}
	}
}
