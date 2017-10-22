package rewrite

import (
	"io"
)

var headTags = []string{"html", "head", "base", "link", "meta", "title", "style", "script", "object", "bgsound"}
var beforeHeadTags = []string{"html", "head"}
var dataRwProtocols = []string{"http://", "https://", "//"}

type HtmlRewriter struct {
	urlRewriter   *UrlRewriter
	jsRewriter    Rewriter
	cssRewriter   Rewriter
	url           string
	defmod        string
	parseComments bool
	rewriteTags   map[string]map[string]string
}

func NewHtmlRewriter(configs ...func(*Config)) *HtmlRewriter {
	c := makeConfig(configs...)
	return &HtmlRewriter{}
}

func rewriteTags(defmod string) map[string]map[string]string {
	return map[string]map[string]string{
		"a":          {"href": defmod},
		"applet":     {"codebase": "oe_", "archive": "oe_"},
		"area":       {"href": defmod},
		"audio":      {"src": "oe_"},
		"base":       {"href": defmod},
		"blockquote": {"cite": defmod},
		"body":       {"background": "im_"},
		"button":     {"formaction": defmod},
		"command":    {"icon": "im_"},
		"del":        {"cite": defmod},
		"embed":      {"src": "oe_"},
		"head":       {"": defmod}, // for head rewriting
		"iframe":     {"src": "if_"},
		"image":      {"src": "im_", "xlink:href": "im_"},
		"img":        {"src": "im_", "srcset": "im_"},
		"ins":        {"cite": defmod},
		"input":      {"src": "im_", "formaction": defmod},
		"form":       {"action": defmod},
		"frame":      {"src": "fr_"},
		"link":       {"href": "oe_"},
		"meta":       {"content": defmod},
		"object":     {"codebase": "oe_", "data": "oe_"},
		"param":      {"value": "oe_"},
		"q":          {"cite": defmod},
		"ref":        {"href": "oe_"},
		"script":     {"src": "js_"},
		"source":     {"src": "oe_"},
		"video":      {"src": "oe_", "poster": "im_"},
	}
}
