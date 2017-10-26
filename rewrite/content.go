package rewrite

import (
	"github.com/datatogether/cdxj"
	"github.com/datatogether/warc/warc"
)

type ContentRewriter struct {
	rules        map[string]RewriteRule
	rewriters    map[RewriterType]Rewriter
	contentTypes map[string]string
}

func NewContentRewriter(options ...func(o *Config)) *ContentRewriter {
	// o := makeConfig(options...)
	return &ContentRewriter{
	// rewriters: o.Rewriters,
	}
}

// func (crw *ContentRewriter) Rewrite() {

// }

type RewriteContentParams struct {
	Record   *warc.Record
	Cdxj     cdxj.Writer
	Urlrw    *UrlRewriter
	Cookierw *CookieRewriter
}

func (crw *ContentRewriter) RewriteContent(req *RewriteContentParams) {

}

func (crw *ContentRewriter) rewriteInfo(req *RewriteContentParams) rewriteInfo {
	return rewriteInfo{
		req.Record,
		req.Urlrw,
		req.Cookierw,
		false,
	}
}

type rewriteInfo struct {
	record      *warc.Record
	urlRw       Rewriter
	cookieRw    Rewriter
	isContentRw bool
	// rewriteTypes
}

func (rwi rewriteInfo) shouldRewriteContent() bool {
	return true
}

// func (crw *ContentRewriter) rwClass(rule, textType string) (string, string) {
// 	if textType == "js" {

// 	}
// }

// var DefaultContentRewriters = map[string]Rewriter{
// 	"header":           DefaultHeaderRewriter,
// 	"cookie":           HostScopeCookieRewriter,
// 	"html":             HTMLRewriter,
// 	"html-banner-only": HTMLInsertOnlyRewriter,
// 	"css":              CSSRewriter,
// 	"js":               JSLocationOnlyRewriter,
// 	"js-proxy":         JSNoneRewriter,
// 	"json":             JSONPRewriter,
// 	"xml":              XMLRewriter,
// 	"dash":             RewriteDASH,
// 	"hls":              RewriteHLS,
// 	"amf":              RewriteAMF,
// }

var RewriteTypes = map[string]string{
	// HTML
	"text/html":             "html",
	"application/xhtml":     "html",
	"application/xhtml+xml": "html",
	// CSS
	"text/css": "css",
	// JS
	"text/javascript":          "js",
	"application/javascript":   "js",
	"application/x-javascript": "js",
	// JSON
	"application/json": "json",
	// HLS
	"application/x-mpegURL":         "hls",
	"application/vnd.apple.mpegurl": "hls",
	// DASH
	"application/dash+xml": "dash",
	// AMF
	"application/x-amf": "amf",
	// XML -- don"t rewrite xml
	//"text/xml": "xml",
	//"application/xml": "xml",
	//"application/rss+xml": "xml",
	// PLAIN
	"text/plain": "guess-text",
	// DEFAULT or octet-stream
	"": "guess-text",
	"application/octet-stream": "guess-bin",
}

var defaultContentTypes = map[string]string{
	"html": "text/html",
	"css":  "text/css",
	"js":   "text/javascript",
}
