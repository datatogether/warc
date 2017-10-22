package rewrite

import (
	"net/url"
)

type UrlRewriter struct {
	Host, Scheme string
}

func NewUrlRewriter(base string) *UrlRewriter {
	u, err := url.Parse(base)
	if err != nil {
		// TODO - ugh.
		panic(err)
	}

	return &UrlRewriter{
		Scheme: u.Scheme,
		Host:   u.Host,
	}
}

func (urw *UrlRewriter) Rewrite(p []byte) ([]byte, error) {
	// call to rewrite with empty slice is a no-op
	if len(p) == 0 {
		return nil, nil
	}

	u, err := url.Parse(string(p))
	if err != nil {
		return nil, err
	}

	u.Host = urw.Host
	u.Scheme = urw.Scheme

	return []byte(u.String()), nil
}

func (urw *UrlRewriter) rewriteBase(p []byte, url, mod string) {

}

func (urw *UrlRewriter) writeDefaultBase() {

}

func (urw *UrlRewriter) ensureUrlHasPath() {

}

func (urw *UrlRewriter) tryUnescape() {

}
