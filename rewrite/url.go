package rewrite

import (
	"net/url"
)

type UrlRewriter struct {
	fromHost string
	to       *url.URL
}

func NewUrlRewriter(from, to string) *UrlRewriter {
	f, err := url.Parse(from)
	if err != nil {
		// TODO - ugh.
		panic(err)
	}

	t, err := url.Parse(to)
	if err != nil {
		// TODO
		panic(err)
	}

	return &UrlRewriter{
		fromHost: f.Host,
		to:       t,
	}
}

func (urw *UrlRewriter) Rewrite(p []byte) ([]byte, error) {
	// call to rewrite with empty slice is a no-op
	if len(p) == 0 {
		return nil, nil
	}

	u, err := urw.to.Parse(string(p))
	if err != nil {
		return nil, err
	}

	// fmt.Println(u.Host, urw.fromHost)
	if u.Host == urw.fromHost {
		u.Host = urw.to.Host
		if u.Scheme != urw.to.Scheme {
			u.Scheme = urw.to.Scheme
		}
	}

	return []byte(u.String()), nil
}
