package rewrite

import (
	"regexp"
)

var (
	CharsetRegex = regexp.MustCompile(`<meta[^>]*?[\s;"\']charset\s*=[\s"\']*([^\s"\'/>]*)`)
	CssUrlRegex  = regexp.MustCompile(`(?m)url\s*\(\s*(?:[\"']|(?:&.{1,4};))*\s*([^)'\"]+)\s*(?:["']|(?:&.{1,4};))*\s*\)`)
	// CssImportNoUrlRegex = regexp.MustCompile(`@import\\s+(?!url)\\(?\\s*['\"]?(?!url[\\s\\(])([\w.:/\\\\-]+)`)
	CssImportNoUrlRegex = regexp.MustCompile(``)
	HttpxMatchString    = regexp.MustCompile(`https?:\\?/\\?/[A-Za-z0-9:_@.-]+`)
	// JsHttpx             = regexp.MustCompile(`(?:(?<=["\';])https?:|(?<=["\']))\\{0,4}/\\{0,4}/[A-Za-z0-9:_@%.\\-]+/`)
	JsHttpx = regexp.MustCompile(``)
)

type RegexRewriter struct {
	Re *regexp.Regexp
	Rw Rewriter
}

func (rerw *RegexRewriter) Rewrite(p []byte) []byte {
	return rerw.Re.ReplaceAllFunc(p, rerw.Rw.Rewrite)
}

// Shameless copy pasta from Stack Overflow
// https://stackoverflow.com/questions/28000832/how-to-access-a-capturing-group-from-regexp-replaceallfunc
func ReplaceAllSubmatchFunc(re *regexp.Regexp, b []byte, f func(s []byte) []byte) []byte {
	idxs := re.FindAllSubmatchIndex(b, -1)
	if len(idxs) == 0 {
		return b
	}
	l := len(idxs)
	ret := append([]byte{}, b[:idxs[0][0]]...)
	for i, pair := range idxs {
		// replace internal submatch with result of user supplied function
		ret = append(ret, f(b[pair[2]:pair[3]])...)
		if i+1 < l {
			ret = append(ret, b[pair[1]:idxs[i+1][0]]...)
		}
	}
	ret = append(ret, b[idxs[len(idxs)-1][1]:]...)
	return ret
}
