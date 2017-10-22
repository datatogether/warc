package rewrite

import (
	"regexp"
	"testing"
)

func TestCharsetRegex(t *testing.T) {
	testRegexCases(t, CharsetRegex, []reTestCase{
		{"nothing", "", "nothing"},
	})
}
func TestCssUrlRegex(t *testing.T) {
	testRegexCases(t, CssUrlRegex, []reTestCase{
		{"nothing", "", "nothing"},
	})
}
func TestCssImportNoUrlRegex(t *testing.T) {
	testRegexCases(t, CssImportNoUrlRegex, []reTestCase{
		{"nothing", "", "nothing"},
	})
}
func TestHttpxMatchString(t *testing.T) {
	testRegexCases(t, HttpxMatchString, []reTestCase{
		{"nothing", "", "nothing"},
	})
}

func TestJsHttpx(t *testing.T) {
	testRegexCases(t, JsHttpx, []reTestCase{
		{"nothing", "", "nothing"},
	})
}

type reTestCase struct {
	in, repl, out string
}

func testRegexCases(t *testing.T, re *regexp.Regexp, cases []reTestCase) {
	for i, c := range cases {
		got := re.ReplaceAll([]byte(c.in), []byte(c.repl))
		if string(got) != c.out {
			t.Errorf("case %d mismatch. expected: '%s', got: '%s'", i, c.out, got)
		}
	}
}
