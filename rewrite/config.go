package rewrite

type Config struct {
	Defmod       string
	Rewriters    []RewriterType
	HeaderPrefix string
	HeaderRules  map[string]RewriteRule
}

func DefaultConfig() *Config {
	return &Config{
		Rewriters: []RewriterType{
			RwTypeCookie,
			RwTypeHeader,
			RwTypeJavascript,
			RwTypeCss,
		},
		Defmod:       "",
		HeaderPrefix: "X-Archive-Orig-",
		HeaderRules:  DefaultHeaderRewriters,
	}
}

func makeConfig(configs ...func(*Config)) *Config {
	cfg := DefaultConfig()
	for _, config := range configs {
		config(cfg)
	}
	return cfg
}

func makeRewriters(cfg *Config) map[RewriterType]Rewriter {
	rws := map[RewriterType]Rewriter{}
	for _, rwt := range cfg.Rewriters {
		switch rwt {
		case RwTypeCookie:
		}
	}

	return rws
}

var DefaultHeaderRewriters = map[string]RewriteRule{
	"access-control-allow-origin":      PrefixIfUrlRewrite,
	"access-control-allow-credentials": PrefixIfUrlRewrite,
	"access-control-expose-headers":    PrefixIfUrlRewrite,
	"access-control-max-age":           PrefixIfUrlRewrite,
	"access-control-allow-methods":     PrefixIfUrlRewrite,
	"access-control-allow-headers":     PrefixIfUrlRewrite,

	"accept-patch":  Keep,
	"accept-ranges": Keep,

	"age": Prefix,

	"allow": Keep,

	"alt-svc":       Prefix,
	"cache-control": Prefix,

	"connection": Prefix,

	"content-base":                        UrlRewrite,
	"content-disposition":                 Keep,
	"content-encoding":                    PrefixIfContentRewrite,
	"content-language":                    Keep,
	"content-length":                      ContentLength,
	"content-location":                    UrlRewrite,
	"content-md5":                         Prefix,
	"content-range":                       Keep,
	"content-security-policy":             Prefix,
	"content-security-policy-report-only": Prefix,
	"content-type":                        Keep,

	"date": Keep,

	"etag":    Prefix,
	"expires": Prefix,

	"last-modified": Prefix,
	"link":          Keep,
	"location":      UrlRewrite,

	"p3p":    Prefix,
	"pragma": Prefix,

	"proxy-authenticate": Keep,

	"public-key-pins": Prefix,
	"retry-after":     Prefix,
	"server":          Prefix,

	"set-cookie": Cookie,

	"strict-transport-security": Prefix,

	"trailer":           Prefix,
	"transfer-encoding": Prefix,
	"tk":                Prefix,

	"upgrade":                   Prefix,
	"upgrade-insecure-requests": Prefix,

	"vary": Prefix,

	"via": Prefix,

	"warning": Prefix,

	"www-authenticate": Keep,

	"x-frame-options":  Prefix,
	"x-xss-protection": Prefix,
}
