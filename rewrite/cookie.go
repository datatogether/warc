package rewrite

import (
	"bytes"
	"fmt"
)

type CookieRewriter struct {
	buf *bytes.Buffer
}

func NewCookieRewriter(configs ...func(*Config)) *CookieRewriter {
	// c := makeConfig(configs...)
	return &CookieRewriter{
		buf: &bytes.Buffer{},
	}
}

func (crw *CookieRewriter) Rewrite(p []byte) ([]byte, error) {
	// TODO
	return nil, fmt.Errorf("not finished")
}

func (crw *CookieRewriter) Read(p []byte) (int, error) {
	return crw.buf.Read(p)
}

func (crw *CookieRewriter) Write(p []byte) (int, error) {
	rw, err := crw.Rewrite(p)
	if err != nil {
		return 0, err
	}
	return crw.buf.Write(rw)
}
