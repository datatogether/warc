package rewrite

import (
	"testing"
)

func TestRewriteHtml(t *testing.T) {
	urlrw := NewUrlRewriter("http://a.com", "https://b.tv")
	rw := NewHtmlRewriter(urlrw)
	cases := stringTestCases([]stringTestCase{
		{"", ""},
		{htmlNoChange, htmlNoChange},
		{basicHtmlRewriteIn, basicHtmlRewriteOut},
	})
	testRewriteCases(t, rw, cases)
}

var htmlNoChange = `<!DOCTYPE html>
<html>
<head>
  <title></title>
</head>
<body>
</body>
</html>`

var basicHtmlRewriteIn = `<!DOCTYPE html>
<html>
<head>
  <title></title>
  <meta></meta>
</head>
<body background="background">
  <a href="/apples" nochange="leave/me/alone">link</a>
  <applet codebase="http://a.com/codebase" archive="http://appletarchive.com"></applet>
  <area href="http://a.com/path" />
  <audio src="http://a.com/audio/path" />
  <base href="/b"></base>
  <blockquote cite="http://a.com"></blockquote>
  <button formaction="/">
    <p>no touch</p>
  </button>
  <command icon="word">
    <del cite="citation"></del>
    <embed src="/b/v/c"></embed>
  </command>
  <iframe src="/word"></iframe>
  <image src="huh" xlink:href="im"></image>
  <img src="stuff" srcset="srcset"></img>
  <ins cite="/1234"></ins>
  <input src="/huh" formaction="/word"></input>
  <form action="/form/action">
    <frame src="/frame/src"></frame>
  </form>
  <link href="/link"></link>
  <script src="/static/js/script.js"></script>
  <source src="/turn/left"></source>
  <video src="/yep" poster="poster"></video>
  <h3 href="http://a.com/stuff"></h3>
</body>
</html>`

var basicHtmlRewriteOut = `<!DOCTYPE html>
<html>
<head>
  <title></title>
  <meta></meta>
</head>
<body background="im_background">
  <a href="https://b.tv/apples" nochange="leave/me/alone">link</a>
  <applet codebase="oe_http://a.com/codebase" archive="oe_http://appletarchive.com"></applet>
  <area href="https://b.tv/path"/>
  <audio src="oe_http://a.com/audio/path"/>
  <base href="https://b.tv/b"></base>
  <blockquote cite="https://b.tv"></blockquote>
  <button formaction="https://b.tv/">
    <p>no touch</p>
  </button>
  <command icon="im_word">
    <del cite="https://b.tv/citation"></del>
    <embed src="oe_/b/v/c"></embed>
  </command>
  <iframe src="if_/word"></iframe>
  <image src="im_huh" xlink:href="im_im"></image>
  <img src="im_stuff" srcset="im_srcset"></img>
  <ins cite="https://b.tv/1234"></ins>
  <input src="im_/huh" formaction="https://b.tv/word"></input>
  <form action="https://b.tv/form/action">
    <frame src="fr_/frame/src"></frame>
  </form>
  <link href="oe_/link"></link>
  <script src="js_/static/js/script.js"></script>
  <source src="oe_/turn/left"></source>
  <video src="oe_/yep" poster="im_poster"></video>
  <h3 href="http://a.com/stuff"></h3>
</body>
</html>`
