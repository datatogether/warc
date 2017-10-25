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
	prefix           string
	rules            map[string]RewriteRule
	Urlrw            Rewriter
	Cookierw         Rewriter
	RewritingContent bool
}

func NewHeaderRewriter(configs ...func(cfg *Config)) *HeaderRewriter {
	c := makeConfig(configs...)
	return &HeaderRewriter{
		prefix: c.HeaderPrefix,
		rules:  c.HeaderRules,
		Urlrw:  c.Defmod,
		// Cookierw: c.CookieRewriter,
		// RewritingContent: c.ContentRewriter != nil,
	}
}

func (hrw *HeaderRewriter) Rewrite(p []byte) ([]byte, error) {
	return nil, ErrNotFinished
}

func (hrw *HeaderRewriter) RewriteHeader(name, value string) (string, string) {
	switch hrw.rules[name] {
	case Keep:
		return name, value
	case UrlRewrite:
		if hrw.Urlrw != nil {
			return name, string(hrw.Urlrw.Rewrite([]byte(value)))
		}
		return name, value
	case PrefixIfContentRewrite:
		if hrw.RewritingContent {
			return hrw.prefix + name, value
		}
		return name, value
	case PrefixIfUrlRewrite:
		if hrw.Urlrw != nil {
			return hrw.prefix + name, value
		}
		return name, value
	case ContentLength:
		if value == "0" {
			return name, value
		}
		//           if not self.rwinfo.is_content_rw:
		//               try:
		//                   if int(value) >= 0:
		//                       return (name, value)
		//               except:
		//                   pass
		return hrw.prefix + name, value
	case Cookie:
		if hrw.Cookierw != nil {
			//               return self.rwinfo.cookie_rewriter.rewrite(value)
		}
		return name, value
	case Prefix:
		return hrw.prefix + name, value
	}
	return name, value
}

func (hrw *HeaderRewriter) AddCacheHeaders(headers map[string]string) {

}
