package rewrite

type RewriteRule int

const (
	Keep RewriteRule = iota
	PrefixIfUrlRewrite
	Prefix
	UrlRewrite
	PrefixIfContentRewrite
	ContentLength
	Cookie
)

type HeaderRewriter struct {
	prefix string
	rules  map[string]RewriteRule
}

func NewHeaderRewriter(configs ...func(cfg *Config)) *HeaderRewriter {
	c := makeConfig(configs...)

	return &HeaderRewriter{
		prefix: c.HeaderPrefix,
		rules:  c.HeaderRules,
	}
}

func (hrw *HeaderRewriter) Rewrite(p []byte) ([]byte, error) {
	return nil, ErrNotFinished
}

func (hrw *HeaderRewriter) rewriteHeader(name, value string) {
	switch hrw.rules[name] {
	case Keep:
	case UrlRewrite:
	case PrefixIfContentRewrite:
	case PrefixIfUrlRewrite:
	case ContentLength:
	case Cookie:
	case Prefix:
	}
}

func (hrw *HeaderRewriter) addCacheHeaders(headers map[string]string) {

}
