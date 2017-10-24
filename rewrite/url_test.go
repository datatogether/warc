package rewrite

import (
	"testing"
)

func TestUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"", "", nil},
		{"http://youtube.com", "http://youtube.com", nil},
		{"https://a.com", "http://b.tv", nil},
		{"http://a.com/", "http://b.tv/", nil},
		{"/relative/url", "http://b.tv/relative/url", nil},
		{"http://a.com/path?query=a", "http://b.tv/path?query=a", nil},
	})

	rw := NewUrlRewriter("http://a.com", "http://b.tv")
	testRewriteCases(t, rw, cases)

	cases = stringTestCases([]stringTestCase{
		{"", "", nil},
		{"http://youtube.com", "http://youtube.com", nil},
		{"http://a.com", "https://b.tv", nil},
		{"/relative/url", "https://b.tv/relative/url", nil},
		{"https://a.com/path?query=a", "https://b.tv/path?query=a", nil},
	})

	rw = NewUrlRewriter("http://a.com", "https://b.tv")
	testRewriteCases(t, rw, cases)
}
