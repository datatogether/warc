package rewrite

import (
	"testing"
)

func TestUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"", "", nil},
		{"http://youtube.com", "http://youtube.com", nil},
		{"https://a.com", "https://b.tv", nil},
		{"http://a.com/", "http://b.tv/", nil},
		{"/relative/url", "http://b.tv/relative/url", nil},
		{"http://a.com/path?query=a", "http://b.tv/path?query=a", nil},
	})

	rw := NewUrlRewriter("http://a.com", "http://b.tv")
	testRewriteCases(t, rw, cases)
}
