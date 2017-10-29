package rewrite

import (
	"github.com/datatogether/cdxj"
	"github.com/datatogether/warc/warc"
	"strings"
)

type WarcRecordRewriter struct {
	Index    cdxj.Writer
	Urlrw    *UrlRewriter
	Cookierw *CookieRewriter

	rules        map[string]RewriteRule
	rewriters    map[RewriterType]Rewriter
	rewriteTypes map[string]string
	contentTypes map[string]string
}

func NewWarcRecordRewriter(config ...func(o *Config)) *WarcRecordRewriter {
	// c := makeConfig(config...)
	return &WarcRecordRewriter{
	// rewriters: o.Rewriters,
	}
}

func (wrr *WarcRecordRewriter) Rewrite(rec *warc.Record) (*warc.Record, error) {
	rwinfo := wrr.rewriteInfo(rec)
	// var contentRewriter Rewriter

	if !rwinfo.shouldRewriteContent() {
		return rec, nil
	}

	// rule := wrr.rule(wrr.Index)
	// contentRewriter := wrr.createRewriter(textType, rule, rwinfo, cdx)
	return nil, ErrNotFinished
}

func (wrr *WarcRecordRewriter) rewriteInfo(rec *warc.Record) rewriteInfo {
	return rewriteInfo{
		record:              rec,
		urlRw:               wrr.Urlrw,
		cookieRw:            wrr.Cookierw,
		isContentRw:         false, // TODO - determine this value
		rewriteTypes:        wrr.rewriteTypes,
		defaultContentTypes: wrr.contentTypes,
	}
}

func (wrr *WarcRecordRewriter) rule(rec cdxj.Writer) RewriteRule {
	// TODO - huh?
	return Keep
}

// func (wrr *WarcRecordRewriter) createRewriter(textType string) {

// }

// func (wrr *WarcRecordRewriter) rwClass(rule, textType string) (string, string) {
// 	if textType == "js" {
// 	}
// }

type rewriteInfo struct {
	record              *warc.Record
	urlRw               Rewriter
	cookieRw            Rewriter
	isContentRw         bool
	rewriteTypes        map[string]string
	defaultContentTypes map[string]string
}

func (rwi rewriteInfo) shouldRewriteContent() bool {
	return true
}

func (rwi rewriteInfo) textTypeAndCharset() (string, string) {
	ct := rwi.record.Headers.Get(warc.FieldNameContentType)
	parts := strings.Split(ct, ";")
	mime := parts[0]
	ogTextType := rwi.rewriteTypes[strings.ToLower(mime)]
	textType := rwi.resolveTextType(ogTextType)
	charset := ""

	if textType == "guess-text" || textType == "guess-bin" {
		textType = ""
	}
	if textType == "js" {
		// if 'callback=jQuery' in self.url_rewriter.wburl.url or '.json?' in self.url_rewriter.wburl.url:
		//                 text_type = 'json'
	}
	if (textType != "" && ogTextType != textType) || textType == "html" {
		newMime := rwi.defaultContentTypes[textType]

		if newMime != "" && newMime != mime {
			// newContentType := ct.
			// new_content_type = content_type.replace(mime, new_mime)
			// self.record.http_headers.replace_header('Content-Type', new_content_type)
		}

		if len(parts) == 2 {
			parts = strings.Split(strings.ToLower(parts[1]), "charset=")
			if len(parts) == 2 {
				charset = strings.TrimSpace(parts[1])
			}
		}
	}

	return textType, charset
}

func (rwi rewriteInfo) resolveTextType(textType string) string {
	// mod = self.url_rewriter.wburl.mod
	mod := ""
	if textType == "css" && mod == "js_" {
		textType = "css"
	}

	isCssOrJs := mod == "js_" || mod == "cs_"

	if textType == "guess-text" {
		if !isCssOrJs && !(mod == "if_" || mod == "mp_" || mod == "") {
			return ""
		}
	} else if textType == "guess-bin" || textType == "html" {
		if !isCssOrJs {
			return textType
		}
	}

	return textType

	// http.DetectContentType(data)
}

func (rwi rewriteInfo) IsUrlRw() bool {
	return true
}

var DefaultWarcRecordRewriters = map[string]Rewriter{
	// "header": DefaultHeaderRewriter,
	"header": NoopRewriter,
	// "cookie": HostScopeCookieRewriter,
	"cookie": NoopRewriter,
	// "html":   HTMLRewriter,
	"html": NoopRewriter,
	// "html-banner-only": HTMLInsertOnlyRewriter,
	"html-banner-only": NoopRewriter,
	// "css":              CSSRewriter,
	"css": NoopRewriter,
	// "js":               JSLocationOnlyRewriter,
	"js": NoopRewriter,
	// "js-proxy":         JSNoneRewriter,
	"js-proxy": NoopRewriter,
	// "json":             JSONPRewriter,
	"json": NoopRewriter,
	// "xml":              XMLRewriter,
	"xml": NoopRewriter,
	// "dash":             RewriteDASH,
	"dash": NoopRewriter,
	// "hls":              RewriteHLS,
	"hls": NoopRewriter,
	// "amf":              RewriteAMF,
	"amf": NoopRewriter,
}

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
