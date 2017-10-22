package rewrite

type RewriteRule int

const (
	PrefixIfUrlRewrite RewriteRule = iota
	Keep
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

func Rewrite(p []byte) (int, error) {
	return 0, nil
}

func (hrw *HeaderRewriter) RewriteHeader(name, value string) {
	switch rule {
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
