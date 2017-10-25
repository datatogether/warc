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
	urlrw := NewUrlRewriter("http://a.com", "https://b.tv")
	rw := NewCssRewriter(urlrw)
	testRewriteCases(t, rw, stringTestCases([]stringTestCase{
		{"", ""},
		{noChangeCss, noChangeCss},
		{`@import url("http://a.com/path/to/css")`, `@import url("https://b.tv/path/to/css")`},
	}))
}
