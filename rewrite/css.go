package rewrite

type CssRewriter struct {
	Rw *UrlRewriter
}

func NewCssRewriter(urlrw *UrlRewriter) *CssRewriter {
	return &CssRewriter{
		Rw: urlrw,
	}
}

func (rerw *CssRewriter) Rewrite(p []byte) ([]byte, error) {
	rep := ReplaceAllSubmatchFunc(CssUrlRegex, p, func(i []byte) []byte {
		o, err := rerw.Rw.Rewrite(i)
		if err != nil {
			return i
		}
		return append([]byte("url(\""), append(o, []byte("\")")...)...)
	})

	return rep, nil
}
