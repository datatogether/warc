package rewrite

import (
	"bytes"
	"errors"
)

var ErrNotFinished = errors.New("not finished")

// Rewriter takes an input byte slice of and returns an output
// slice of rewritten bytes, the length of input & output will
// not necessarily match, implementations *may* alter input bytes
type Rewriter interface {
	Rewrite(i []byte) (o []byte, err error)
}

// RewriterType enumerates rewriters that operate on different
// types of content
type RewriterType int

const (
	RwTypeUnknown RewriterType = iota
	RwTypeUrl
	RwTypeHeader
	RwTypeContent
	RwTypeCookie
	RwTypeHtml
	RwTypeJavascript
	RwTypeCss
)

func (rwt RewriterType) String() string {
	return map[RewriterType]string{
		RwTypeUnknown:    "",
		RwTypeUrl:        "url",
		RwTypeHeader:     "header",
		RwTypeContent:    "content",
		RwTypeCookie:     "cookie",
		RwTypeHtml:       "html",
		RwTypeJavascript: "javascript",
		RwTypeCss:        "css",
	}[rwt]
}

var NoopRewriter = PrefixRewriter{}

// PrefixRewriter adds a prefix if not present
type PrefixRewriter struct {
	Prefix []byte
}

func (prw PrefixRewriter) Rewrite(p []byte) ([]byte, error) {
	if !bytes.HasPrefix(p, prw.Prefix) {
		return append(prw.Prefix, p...), nil
	}
	return p, nil
}
