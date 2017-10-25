package rewrite

type Config struct {
	DestUrl      string
	Defmod       Rewriter
	Rewriters    []RewriterType
	HeaderPrefix string
	HeaderRules  map[string]RewriteRule
}

func DefaultConfig() *Config {
	return &Config{
		DestUrl:      "",
		Defmod:       NoopRewriter,
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
