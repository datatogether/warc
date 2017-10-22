package rewrite

import (
	"testing"
)

const noChangeCss = `
  html, head, body{
    color: black;
    background: black;
    margin: 0;
    padding: 0;
  }
`

func TestCssRewriter(t *testing.T) {
	urlrw := NewUrlRewriter("http://a.com", "http://b.tv")
	rw := NewCssRewriter(urlrw)
	testRewriteCases(t, rw, stringTestCases([]stringTestCase{
		{"", "", nil},
		{noChangeCss, noChangeCss, nil},
		{`@import("http://a.com/path/to/css")`, `@import("http://t.tv/path/to/css")`, nil},
	}))
}
