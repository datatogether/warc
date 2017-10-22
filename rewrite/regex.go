package rewrite

import (
	"bytes"
	"regexp"
)

var (
	CharsetRegex = regexp.MustCompile(`<meta[^>]*?[\s;"\']charset\s*=[\s"\']*([^\s"\'/>]*)`)
	CssUrlRegex  = regexp.MustCompile(`url\\s*\\(\\s*(?:[\\\\\"']|(?:&.{1,4};))*\\s*([^)'\"]+)\\s*(?:[\\\\\"']|(?:&.{1,4};))*\\s*\\)`)
	// CssImportNoUrlRegex = regexp.MustCompile(`@import\\s+(?!url)\\(?\\s*['\"]?(?!url[\\s\\(])([\w.:/\\\\-]+)`)
	CssImportNoUrlRegex = regexp.MustCompile(``)
	HttpxMatchString    = regexp.MustCompile(`https?:\\?/\\?/[A-Za-z0-9:_@.-]+`)
	// JsHttpx             = regexp.MustCompile(`(?:(?<=["\';])https?:|(?<=["\']))\\{0,4}/\\{0,4}/[A-Za-z0-9:_@%.\\-]+/`)
	JsHttpx = regexp.MustCompile(``)
)

var NoopRewriter = PrefixRewriter{}

type RegexRewriter struct {
	Re      *regexp.Regexp
	Replace []byte
}

func (rerw *RegexRewriter) Rewrite(p []byte) ([]byte, error) {
	repl := CssImportNoUrlRegex.ReplaceAll(p, rerw.Replace)
	return repl, nil
}

// PrefixRewriter adds a prefix if not present
type PrefixRewriter struct {
	Prefix []byte
}

func (prw PrefixRewriter) Rewrite(p []byte) ([]byte, error) {
	if !bytes.HasPrefix(p, prw.Prefix) {
		return append(prw.Prefix, p...), nil
	}
	return p, nil
}
