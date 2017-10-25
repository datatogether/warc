package rewrite

import (
	"net/http"
	"testing"
)

func TestHeaderRewriter(t *testing.T) {
	defaultHeaderRw := NewHeaderRewriter(func(c *Config) {
		c.HeaderPrefix = "Test-"
		c.Defmod = NewUrlRewriter("http://a.com", "https://b.tv")
	})

	cases := []struct {
		hrw     *HeaderRewriter
		in, out map[string]string
	}{
		{defaultHeaderRw, input, output},
	}

	for i, c := range cases {
		headers := http.Header{}
		for key, value := range c.in {
			headers.Add(key, value)
		}

		out := c.hrw.RewriteHeaders(headers)

		for key, _ := range out {
			value := out.Get(key)
			expect := c.out[key]
			if expect == "" {
				t.Errorf("case %d generated unexpected header: %s, value: %s", i, key, value)
				continue
			}

			if expect != value {
				t.Errorf("case %d header: %s, value mismatch. expected: '%s' got: '%s'", i, key, expect, value)
				continue
			}
		}
	}
}

var input = map[string]string{
	"Access-Control-Allow-Origin":         "prefix_if_url_rewrite",
	"Access-Control-Allow-Credentials":    "prefix_if_url_rewrite",
	"Access-Control-Expose-Headers":       "prefix_if_url_rewrite",
	"Access-Control-Max-Age":              "prefix_if_url_rewrite",
	"Access-Control-Allow-Methods":        "prefix_if_url_rewrite",
	"Access-Control-Allow-Headers":        "prefix_if_url_rewrite",
	"Accept-Patch":                        "keep",
	"Accept-Ranges":                       "keep",
	"Age":                                 "prefix",
	"Allow":                               "keep",
	"Alt-Svc":                             "prefix",
	"Cache-Control":                       "prefix",
	"Connection":                          "prefix",
	"Content-Base":                        "http://a.com/path",
	"Content-Disposition":                 "keep",
	"Content-Encoding":                    "prefix_if_content_rewrite",
	"Content-Language":                    "keep",
	"Content-Length":                      "100",
	"Content-Location":                    "http://a.com/path",
	"Content-Md5":                         "prefix",
	"Content-Range":                       "keep",
	"Content-Security-Policy":             "prefix",
	"Content-Security-Policy-Report-Only": "prefix",
	"Content-Type":                        "keep",
	"Date":                                "keep",
	"Etag":                                "prefix",
	"Expires":                             "prefix",
	"Last-Modified":                       "prefix",
	"Link":                                "keep",
	"Location":                            "http://a.com/path",
	"P3p":                                 "prefix",
	"Pragma":                              "prefix",
	"Proxy-Authenticate":                  "keep",
	"Public-Key-Pins":                     "prefix",
	"Retry-After":                         "prefix",
	"Server":                              "prefix",
	"Set-Cookie":                          "cookie",
	"Strict-Transport-Security":           "prefix",
	"Trailer":                             "prefix",
	"Transfer-Encoding":                   "prefix",
	"Tk":                                  "prefix",
	"Upgrade":                             "prefix",
	"Upgrade-Insecure-Requests":           "prefix",
	"Vary":             "prefix",
	"Via":              "prefix",
	"Warning":          "prefix",
	"Www-Authenticate": "keep",
	"X-Frame-Options":  "prefix",
	"X-Xss-Protection": "prefix",
}

var output = map[string]string{
	"Test-Access-Control-Allow-Origin":         "prefix_if_url_rewrite",
	"Test-Access-Control-Allow-Credentials":    "prefix_if_url_rewrite",
	"Test-Access-Control-Expose-Headers":       "prefix_if_url_rewrite",
	"Test-Access-Control-Max-Age":              "prefix_if_url_rewrite",
	"Test-Access-Control-Allow-Methods":        "prefix_if_url_rewrite",
	"Test-Access-Control-Allow-Headers":        "prefix_if_url_rewrite",
	"Accept-Patch":                             "keep",
	"Accept-Ranges":                            "keep",
	"Test-Age":                                 "prefix",
	"Allow":                                    "keep",
	"Test-Alt-Svc":                             "prefix",
	"Test-Cache-Control":                       "prefix",
	"Test-Connection":                          "prefix",
	"Content-Base":                             "https://b.tv/path",
	"Content-Disposition":                      "keep",
	"Content-Encoding":                         "prefix_if_content_rewrite",
	"Content-Language":                         "keep",
	"Content-Length":                           "100",
	"Content-Location":                         "https://b.tv/path",
	"Test-Content-Md5":                         "prefix",
	"Content-Range":                            "keep",
	"Test-Content-Security-Policy":             "prefix",
	"Test-Content-Security-Policy-Report-Only": "prefix",
	"Content-Type":                             "keep",
	"Date":                                     "keep",
	"Test-Etag":                                "prefix",
	"Test-Expires":                             "prefix",
	"Test-Last-Modified":                       "prefix",
	"Link":                                     "keep",
	"Location":                                 "https://b.tv/path",
	"Test-P3p":                                 "prefix",
	"Test-Pragma":                              "prefix",
	"Proxy-Authenticate":                       "keep",
	"Test-Public-Key-Pins":                     "prefix",
	"Test-Retry-After":                         "prefix",
	"Test-Server":                              "prefix",
	"Set-Cookie":                               "cookie",
	"Test-Strict-Transport-Security":           "prefix",
	"Test-Trailer":                             "prefix",
	"Test-Transfer-Encoding":                   "prefix",
	"Test-Tk":                                  "prefix",
	"Test-Upgrade":                             "prefix",
	"Test-Upgrade-Insecure-Requests":           "prefix",
	"Test-Vary":                                "prefix",
	"Test-Via":                                 "prefix",
	"Test-Warning":                             "prefix",
	"Www-Authenticate":                         "keep",
	"Test-X-Frame-Options":                     "prefix",
	"Test-X-Xss-Protection":                    "prefix",
}
