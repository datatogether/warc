package rewrite

type CssRewriter struct {
	Rw *UrlRewriter
}

func NewCssRewriter(urlrw *UrlRewriter) *CssRewriter {
	return &CssRewriter{
		Rw: urlrw,
	}
}

func (rerw *CssRewriter) Rewrite(p []byte) []byte {
	return ReplaceAllSubmatchFunc(CssUrlRegex, p, func(i []byte) []byte {
		o := rerw.Rw.Rewrite(i)
		return append([]byte("url(\""), append(o, []byte("\")")...)...)
	})
}
