package rewrite

import (
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
