package rewrite

import (
	"testing"
)

func TestUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"", ""},
		{"http://youtube.com", "http://youtube.com"},
		{"https://a.com", "http://b.tv"},
		{"http://a.com/", "http://b.tv/"},
		{"/relative/url", "http://b.tv/relative/url"},
		{"http://a.com/path?query=a", "http://b.tv/path?query=a"},
	})

	rw := NewUrlRewriter("http://a.com", "http://b.tv")
	testRewriteCases(t, rw, cases)
}

func TestSchemeChangeUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"", ""},
		{"http://youtube.com", "http://youtube.com"},
		{"http://a.com", "https://b.tv"},
		{"/relative/url", "https://b.tv/relative/url"},
		{"https://a.com/path?query=a", "https://b.tv/path?query=a"},
	})

	rw := NewUrlRewriter("http://a.com", "https://b.tv")
	testRewriteCases(t, rw, cases)
}

func TestRelativeUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"", ""},
		{"http://youtube.com", "http://youtube.com"},
		{"http://a.com", "/"},
		{"/relative/url", "/relative/url"},
		{"https://a.com/path?query=a", "/path?query=a"},
	})

	rw := NewRelativeUrlRewriter("http://a.com")
	testRewriteCases(t, rw, cases)
}
