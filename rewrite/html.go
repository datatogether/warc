package rewrite

import (
	"bytes"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

var headTags = []string{"html", "head", "base", "link", "meta", "title", "style", "script", "object", "bgsound"}
var beforeHeadTags = []string{"html", "head"}
var dataRwProtocols = []string{"http://", "https://", "//"}

type HtmlRewriter struct {
	urlRewriter   *UrlRewriter
	jsRewriter    Rewriter
	cssRewriter   Rewriter
	url           string
	defmod        Rewriter
	parseComments bool
	rewriteTags   map[string]map[string]Rewriter
}

func NewHtmlRewriter(configs ...func(*Config)) *HtmlRewriter {
	c := makeConfig(configs...)
	return &HtmlRewriter{
		rewriteTags: rewriteTags(c.Defmod),
	}
}

func (hrw *HtmlRewriter) Rewrite(p []byte) ([]byte, error) {
	rdr := bytes.NewReader(p)
	tokenizer := html.NewTokenizer(rdr)
	w := &bytes.Buffer{}

	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()
		switch tt {
		// case html.TextToken:
		// case html.CommentToken:
		case html.ErrorToken:
			// ErrorToken means that an error occurred during tokenization.
			// most common is end-of-file (EOF)
			if tokenizer.Err().Error() == "EOF" {
				return w.Bytes(), nil
			}
			return nil, tokenizer.Err()
		case html.DoctypeToken:
		case html.StartTagToken, html.SelfClosingTagToken:
			name, hasAttr := tokenizer.TagName()
			if hasAttr {
				hrw.rewriteToken(string(name), &token)
			}
		}

		w.WriteString(token.String())
	}

	return p, nil
}

func (hrw *HtmlRewriter) rewriteMetaRefresh(p []byte, metaRefresh *regexp.Regexp) {

}

func (hrw *HtmlRewriter) rewriteToken(name string, tok *html.Token) error {
	attrs := hrw.rewriteTags[name]
	if attrs != nil {
		for _, a := range tok.Attr {
			repl := attrs[strings.ToLower(a.Key)]
			if repl != nil {
				rw, err := repl.Rewrite([]byte(a.Val))
				if err != nil {
					return err
				}
				a.Val = string(rw)
			}
		}
	}
	return nil
}

func rewriteTags(defmod Rewriter) map[string]map[string]Rewriter {
	oe := PrefixRewriter{Prefix: []byte("oe_")}
	im := PrefixRewriter{Prefix: []byte("im_")}
	if_ := PrefixRewriter{Prefix: []byte("if_")}
	fr_ := PrefixRewriter{Prefix: []byte("fr_")}
	js_ := PrefixRewriter{Prefix: []byte("js_")}

	return map[string]map[string]Rewriter{
		"a":          {"href": defmod},
		"applet":     {"codebase": oe, "archive": oe},
		"area":       {"href": defmod},
		"audio":      {"src": oe},
		"base":       {"href": defmod},
		"blockquote": {"cite": defmod},
		"body":       {"background": im},
		"button":     {"formaction": defmod},
		"command":    {"icon": im},
		"del":        {"cite": defmod},
		"embed":      {"src": oe},
		"head":       {"": defmod}, // for head rewriting
		"iframe":     {"src": if_},
		"image":      {"src": im, "xlink:href": im},
		"img":        {"src": im, "srcset": im},
		"ins":        {"cite": defmod},
		"input":      {"src": im, "formaction": defmod},
		"form":       {"action": defmod},
		"frame":      {"src": fr_},
		"link":       {"href": oe},
		"meta":       {"content": defmod},
		"object":     {"codebase": oe, "data": oe},
		"param":      {"value": oe},
		"q":          {"cite": defmod},
		"ref":        {"href": oe},
		"script":     {"src": js_},
		"source":     {"src": oe},
		"video":      {"src": oe, "poster": im},
	}
}
