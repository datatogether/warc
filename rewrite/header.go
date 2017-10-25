package rewrite

import (
	"net/http"
	"strconv"
)

type RewriteRule int

const (
	Keep RewriteRule = iota
	PrefixIfUrlRewrite
	Prefix
	UrlRewrite
	PrefixIfContentRewrite
	ContentLength
	Cookie
)

type HeaderRewriter struct {
	Prefix           string
	Rules            map[string]RewriteRule
	Urlrw            Rewriter
	Cookierw         Rewriter
	RewritingContent bool
}

func NewHeaderRewriter(configs ...func(cfg *Config)) *HeaderRewriter {
	c := makeConfig(configs...)
	return &HeaderRewriter{
		Prefix: c.HeaderPrefix,
		Rules:  c.HeaderRules,
		Urlrw:  c.Defmod,
		// Cookierw: c.CookieRewriter,
		// RewritingContent: c.ContentRewriter != nil,
	}
}

func (hrw HeaderRewriter) RewriteHeaders(headers http.Header) http.Header {
	rewritten := http.Header{}
	for key, _ := range headers {
		newkey, newval := hrw.rewriteHeader(key, headers.Get(key))
		rewritten.Add(newkey, newval)
	}
	return rewritten
}

func (hrw HeaderRewriter) rewriteHeader(name, value string) (string, string) {
	switch hrw.Rules[name] {
	case Keep:
		return name, value
	case UrlRewrite:
		if hrw.Urlrw != nil {
			return name, string(hrw.Urlrw.Rewrite([]byte(value)))
		}
		return name, value
	case PrefixIfContentRewrite:
		if hrw.RewritingContent {
			return hrw.Prefix + name, value
		}
		return name, value
	case PrefixIfUrlRewrite:
		if hrw.Urlrw != nil {
			return hrw.Prefix + name, value
		}
		return name, value
	case ContentLength:
		if value == "0" {
			return name, value
		}
		// if not rewriting content, attempt to use the
		// length value
		if !hrw.RewritingContent {
			if lenth, err := strconv.Atoi(value); err == nil {
				return name, strconv.FormatInt(int64(lenth), 10)
			}
		}
		return hrw.Prefix + name, value
	case Cookie:
		if hrw.Cookierw != nil {
			//               return self.rwinfo.cookie_rewriter.rewrite(value)
		}
		return name, value
	case Prefix:
		return hrw.Prefix + name, value
	}
	return name, value
}

// func (hrw *HeaderRewriter) AddCacheHeaders(headers map[string]string) {
// }

var DefaultHeaderRewriters = map[string]RewriteRule{
	"Access-Control-Allow-Origin":      PrefixIfUrlRewrite,
	"Access-Control-Allow-Credentials": PrefixIfUrlRewrite,
	"Access-Control-Expose-Headers":    PrefixIfUrlRewrite,
	"Access-Control-Max-Age":           PrefixIfUrlRewrite,
	"Access-Control-Allow-Methods":     PrefixIfUrlRewrite,
	"Access-Control-Allow-Headers":     PrefixIfUrlRewrite,

	"Accept-Patch":  Keep,
	"Accept-Ranges": Keep,

	"Age": Prefix,

	"Allow": Keep,

	"Alt-Svc":       Prefix,
	"Cache-Control": Prefix,

	"Connection": Prefix,

	"Content-Base":                        UrlRewrite,
	"Content-Disposition":                 Keep,
	"Content-Encoding":                    PrefixIfContentRewrite,
	"Content-Language":                    Keep,
	"Content-Length":                      ContentLength,
	"Content-Location":                    UrlRewrite,
	"Content-Md5":                         Prefix,
	"Content-Range":                       Keep,
	"Content-Security-Policy":             Prefix,
	"Content-Security-Policy-Report-Only": Prefix,
	"Content-Type":                        Keep,

	"Date": Keep,

	"Etag":    Prefix,
	"Expires": Prefix,

	"Last-Modified": Prefix,
	"Link":          Keep,
	"Location":      UrlRewrite,

	"P3p":    Prefix,
	"Pragma": Prefix,

	"Proxy-Authenticate": Keep,

	"Public-Key-Pins": Prefix,
	"Retry-After":     Prefix,
	"Server":          Prefix,

	"Set-Cookie": Cookie,

	"Strict-Transport-Security": Prefix,

	"Trailer":           Prefix,
	"Transfer-Encoding": Prefix,
	"Tk":                Prefix,

	"Upgrade":                   Prefix,
	"Upgrade-Insecure-Requests": Prefix,

	"Vary": Prefix,

	"Via": Prefix,

	"Warning": Prefix,

	"Www-Authenticate": Keep,

	"X-Frame-Options":  Prefix,
	"X-Xss-Protection": Prefix,
}
