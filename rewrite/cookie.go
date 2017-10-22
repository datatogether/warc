package rewrite

import (
	"fmt"
)

type CookieRewriter struct {
}

func NewCookieRewriter(configs ...func(*Config)) *CookieRewriter {
	// c := makeConfig(configs...)
	return &CookieRewriter{}
}

func (crw *CookieRewriter) Rewrite(p []byte) ([]byte, error) {
	// TODO
	return nil, fmt.Errorf("not finished")
}
