package rewrite

type CookieRewriter struct {
}

func NewCookieRewriter(configs ...func(*Config)) *CookieRewriter {
	// c := makeConfig(configs...)
	return &CookieRewriter{}
}

func (crw *CookieRewriter) Rewrite(p []byte) []byte {
	// TODO
	return p
}
