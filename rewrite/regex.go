package rewrite

import (
	"regexp"
)

var (
	CharsetRegex = regexp.MustCompile(`<meta[^>]*?[\s;"\']charset\s*=[\s"\']*([^\s"\'/>]*)`)
)
