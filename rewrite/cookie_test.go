package rewrite

import (
	"testing"
)

const noChangeCookie = `
COOKIE DATA HERE
`

func TestCookieRewriter(t *testing.T) {
	rw := NewCookieRewriter()
	testRewriteCases(t, rw, stringTestCases([]stringTestCase{
		{"", ""},
		{noChangeCookie, noChangeCookie},
	}))
}
