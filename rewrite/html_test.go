package rewrite

import (
	"testing"
)

func TestRewriteHtml(t *testing.T) {
	rw := NewHtmlRewriter()
	testRewriteCases(t, rw, []rewriteTestCase{
		{htmlNoChange, htmlNoChange, nil},
	})
}

var htmlNoChange = []byte(`<!DOCTYPE html>
<html>
<head>
  <title></title>
</head>
<body>

</body>
</html>`)
