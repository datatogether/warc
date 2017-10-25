package rewrite

import (
	"bytes"
	"fmt"
	"regexp"

	"golang.org/x/net/html"
)

var headTags = []string{"html", "head", "base", "link", "meta", "title", "style", "script", "object", "bgsound"}
var beforeHeadTags = []string{"html", "head"}
var dataRwProtocols = []string{"http://", "https://", "//"}

type HtmlRewriter struct {
	urlrw         Rewriter
	jsRewriter    Rewriter
	cssRewriter   Rewriter
	url           string
	defmod        Rewriter
	parseComments bool
	rewriteTags   map[string]map[string]Rewriter
}

func NewHtmlRewriter(urlrw Rewriter, configs ...func(*Config)) *HtmlRewriter {
	// c := makeConfig(configs...)
	return &HtmlRewriter{
		urlrw:       urlrw,
		rewriteTags: rewriteTags(urlrw),
	}
}

func (hrw *HtmlRewriter) Rewrite(p []byte) []byte {
	rdr := bytes.NewReader(p)
	tokenizer := html.NewTokenizer(rdr)
	w := &bytes.Buffer{}
	var token html.Token

	for {
		tt := tokenizer.Next()
		// token := tokenizer.Token()
		switch tt {
		// case html.TextToken:
		// case html.CommentToken:
		case html.ErrorToken:
			// ErrorToken means that an error occurred during tokenization.
			// most common is end-of-file (EOF)
			if tokenizer.Err().Error() == "EOF" {
				return w.Bytes()
			}

			fmt.Println(tokenizer.Err().Error())
			return p
		case html.StartTagToken:
			name, hasAttr := tokenizer.TagName()
			token := html.Token{
				Type: html.StartTagToken,
				Data: string(name),
			}
			if hasAttr {
				hrw.rewriteToken(&token, tokenizer)
			}
			w.WriteString(token.String())
			continue
		case html.SelfClosingTagToken:
			name, hasAttr := tokenizer.TagName()
			token := html.Token{
				Type: html.SelfClosingTagToken,
				Data: string(name),
			}
			if hasAttr {
				hrw.rewriteToken(&token, tokenizer)
			}
			w.WriteString(token.String())
			continue
		}

		token = tokenizer.Token()
		w.WriteString(token.String())
	}

	return w.Bytes()
}

func (hrw *HtmlRewriter) rewriteMetaRefresh(p []byte, metaRefresh *regexp.Regexp) {

}

func (hrw *HtmlRewriter) rewriteToken(t *html.Token, tok *html.Tokenizer) {
	attrs := hrw.rewriteTags[t.Data]
	for {
		key, val, more := tok.TagAttr()
		repl := attrs[string(bytes.ToLower(key))]
		if repl != nil {
			val = repl.Rewrite(val)
		}

		t.Attr = append(t.Attr, html.Attribute{
			Key: string(key),
			Val: string(val),
		})

		if !more {
			return
		}
	}
	return
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
