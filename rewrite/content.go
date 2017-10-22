package rewrite

type ContentRewriter struct {
	rules     map[string]RewriteRule
	rewriters map[RewriterType]Rewriter
}

func NewContentRewriter(options ...func(o *Config)) {
	o := DefaultConfig()
	for _, option := range options {
		option(o)
	}

}

func (crw *ContentRewriter) rewriter() {

}

func (crw *ContentRewriter) rwClass(rule, textType string) (string, string) {
	if textType == "js" {

	}

}
