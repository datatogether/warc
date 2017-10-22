package rewrite

type CssRewriter struct {
	urlrw *UrlRewriter
}

func NewCssRewriter(urlrw *UrlRewriter) *CssRewriter {
	return &CssRewriter{
		urlrw: urlrw,
	}
}

func (crw *CssRewriter) Rewrite(p []byte) ([]byte, error) {
	repl := CssImportNoUrlRegex.ReplaceAllFunc(p, func(match []byte) []byte {
		rep, err := crw.urlrw.Rewrite(match)
		if err != nil {
			return match
		}
		return rep
	})
	return repl, nil
}
