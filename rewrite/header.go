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
	Urlrw  Rewriter
}

func NewHeaderRewriter(configs ...func(cfg *Config)) *HeaderRewriter {
	c := makeConfig(configs...)
	return &HeaderRewriter{
		prefix: c.HeaderPrefix,
		rules:  c.HeaderRules,
		Urlrw:  c.Defmod,
	}
}

func (hrw *HeaderRewriter) Rewrite(p []byte) ([]byte, error) {
	return nil, ErrNotFinished
}

func (hrw *HeaderRewriter) rewriteHeader(name, value string) (string, string) {
	switch hrw.rules[name] {
	case Keep:
		return name, value
	case UrlRewrite:

	case PrefixIfContentRewrite:
	case PrefixIfUrlRewrite:
	case ContentLength:
	case Cookie:
	case Prefix:

		// if rule == 'keep':
		//           return (name, value)

		//       elif rule == 'url-rewrite':
		//           if self.rwinfo.is_url_rw():
		//               return (name, self.rwinfo.url_rewriter.rewrite(value))
		//           else:
		//               return (name, value)

		//       elif rule == 'prefix-if-content-rewrite':
		//           if self.rwinfo.is_content_rw:
		//               return (self.header_prefix + name, value)
		//           else:
		//               return (name, value)

		//       elif rule == 'prefix-if-url-rewrite':
		//           if self.rwinfo.is_url_rw():
		//               return (self.header_prefix + name, value)
		//           else:
		//               return (name, value)

		//       elif rule == 'content-length':
		//           if value == '0':
		//               return (name, value)

		//           if not self.rwinfo.is_content_rw:
		//               try:
		//                   if int(value) >= 0:
		//                       return (name, value)
		//               except:
		//                   pass

		//           return (self.header_prefix + name, value)

		//       elif rule == 'cookie':
		//           if self.rwinfo.cookie_rewriter:
		//               return self.rwinfo.cookie_rewriter.rewrite(value)
		//           else:
		//               return (name, value)

		//       elif rule == 'prefix':
		//           return (self.header_prefix + name, value)

		//       return (name, value)
	}
	return name, value
}

func (hrw *HeaderRewriter) AddCacheHeaders(headers map[string]string) {

}
