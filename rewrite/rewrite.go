package rewrite

import (
	"bytes"
	"io"
)

type Rewriter interface {
	// io.ReadWriter
	Rewrite([]byte) (int, error)
}

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

type RewriteBuffer struct {
	bytes.Buffer
	rw Rewriter
}

func (rwb *RewriteBuffer) Write(p []byte) (int, error) {
	if b, err := rwb.rw.Rewrite(p); err != nil {
		return b, err
	}

	return rwb.Buffer.Write(p)
}
