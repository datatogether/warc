package warc

import (
	"os"
	"testing"
)

func TestWarcWrite(t *testing.T) {
	f, err := os.Open("testdata/test.warc")
	// data, err := ioutil.ReadFile("testdata/test.warc")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer f.Close()

	rdr, err := NewReader(f)
	if err != nil {
		t.Error(err)
		return
	}

	records, err := rdr.ReadAll()
	if err != nil {
		t.Error(err)
		return
	}

	out, err := os.Create("testdata/out.warc")
	if err != nil {
		t.Error(err)
		return
	}
	defer out.Close()
	if err := WriteRecords(out, records); err != nil {
		t.Error(err)
		return
	}

}

const WARCINFO_RECORD = `
WARC/1.0\r\n
WARC-Type: warcinfo\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Filename: testfile.warc.gz\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
Content-Type: application/warc-fields\r\n
Content-Length: 86\r\n
\r\n
software: recorder test\r\n
format: WARC File Format 1.0\r\n
json-metadata: {"foo": "bar"}\r\n
\r\n
\r\n
`

const RESPONSE_RECORD = `
WARC/1.0\r\n
WARC-Type: response\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: http://example.com/\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r\n
WARC-Block-Digest: sha1:OS3OKGCWQIJOAOC3PKXQOQFD52NECQ74\r\n
Content-Type: application/http; msgtype=response\r\n
Content-Length: 97\r\n
\r\n
HTTP/1.0 200 OK\r\n
Content-Type: text/plain; charset="UTF-8"\r\n
Custom-Header: somevalue\r\n
\r\n
some\n
text\r\n
\r\n
`

const RESPONSE_RECORD_2 = `
WARC/1.0\r\n
WARC-Type: response\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: http://example.com/\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r\n
WARC-Block-Digest: sha1:U6KNJY5MVNU3IMKED7FSO2JKW6MZ3QUX\r\n
Content-Type: application/http; msgtype=response\r\n
Content-Length: 145\r\n
\r\n
HTTP/1.0 200 OK\r\n
Content-Type: text/plain; charset="UTF-8"\r\n
Content-Length: 9\r\n
Custom-Header: somevalue\r\n
Content-Encoding: x-unknown\r\n
\r\n
some\n
text\r\n
\r\n
`

const REQUEST_RECORD = `
WARC/1.0\r\n
WARC-Type: request\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: http://example.com/\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ\r\n
WARC-Block-Digest: sha1:ONEHF6PTXPTTHE3333XHTD2X45TZ3DTO\r\n
Content-Type: application/http; msgtype=request\r\n
Content-Length: 54\r\n
\r\n
GET / HTTP/1.0\r\n
User-Agent: foo\r\n
Host: example.com\r\n
\r\n
\r\n
\r\n
`

const REQUEST_RECORD_2 = `
WARC/1.0\r\n
WARC-Type: request\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: http://example.com/\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:R5VZAKIE53UW5VGK43QJIFYS333QM5ZA\r\n
WARC-Block-Digest: sha1:L7SVBUPPQ6RH3ANJD42G5JL7RHRVZ5DV\r\n
Content-Type: application/http; msgtype=request\r\n
Content-Length: 92\r\n
\r\n
POST /path HTTP/1.0\r\n
Content-Type: application/json\r\n
Content-Length: 17\r\n
\r\n
{"some": "value"}\r\n
\r\n
`

const REVISIT_RECORD_1 = `
WARC/1.0\r\n
WARC-Type: revisit\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: http://example.com/\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Profile: http://netpreserve.org/warc/1.0/revisit/identical-payload-digest\r\n
WARC-Refers-To-Target-URI: http://example.com/foo\r\n
WARC-Refers-To-Date: 1999-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r\n
WARC-Block-Digest: sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ\r\n
Content-Type: application/http; msgtype=response\r\n
Content-Length: 0\r\n
\r\n
\r\n
\r\n
`

const REVISIT_RECORD_2 = `
WARC/1.0\r\n
WARC-Type: revisit\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: http://example.com/\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Profile: http://netpreserve.org/warc/1.0/revisit/identical-payload-digest\r\n
WARC-Refers-To-Target-URI: http://example.com/foo\r\n
WARC-Refers-To-Date: 1999-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r\n
WARC-Block-Digest: sha1:A6J5UTI2QHHCZFCFNHQHCDD3JJFKP53V\r\n
Content-Type: application/http; msgtype=response\r\n
Content-Length: 88\r\n
\r\n
HTTP/1.0 200 OK\r\n
Content-Type: text/plain; charset="UTF-8"\r\n
Custom-Header: somevalue\r\n
\r\n
\r\n
\r\n
`

const RESOURCE_RECORD = `
WARC/1.0\r\n
WARC-Type: resource\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: ftp://example.com/\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r\n
WARC-Block-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r\n
Content-Type: text/plain\r\n
Content-Length: 9\r\n
\r\n
some\n
text\r\n
\r\n
`

const METADATA_RECORD = `
WARC/1.0\r\n
WARC-Type: metadata\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: http://example.com/\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:ZOLBLKAQVZE5DXH56XE6EH6AI6ZUGDPT\r\n
WARC-Block-Digest: sha1:ZOLBLKAQVZE5DXH56XE6EH6AI6ZUGDPT\r\n
Content-Type: application/json\r\n
Content-Length: 67\r\n
\r\n
{"metadata": {"nested": "obj", "list": [1, 2, 3], "length": "123"}}\r\n
\r\n
`

const DNS_RESPONSE_RECORD = `
WARC/1.0\r\n
WARC-Type: response\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: dns:google.com\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:2AAVJYKKIWK5CF6EWE7PH63EMNLO44TH\r\n
WARC-Block-Digest: sha1:2AAVJYKKIWK5CF6EWE7PH63EMNLO44TH\r\n
Content-Type: application/http; msgtype=response\r\n
Content-Length: 147\r\n
\r\n
20170509000739\n
google.com.     185 IN  A   209.148.113.239\n
google.com.     185 IN  A   209.148.113.238\n
google.com.     185 IN  A   209.148.113.250\n
\r\n\r\n
`
const DNS_RESOURCE_RECORD = `
WARC/1.0\r\n
WARC-Type: resource\r\n
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n
WARC-Target-URI: dns:google.com\r\n
WARC-Date: 2000-01-01T00:00:00Z\r\n
WARC-Payload-Digest: sha1:2AAVJYKKIWK5CF6EWE7PH63EMNLO44TH\r\n
WARC-Block-Digest: sha1:2AAVJYKKIWK5CF6EWE7PH63EMNLO44TH\r\n
Content-Type: application/warc-record\r\n
Content-Length: 147\r\n
\r\n
20170509000739\n
google.com.     185 IN  A   209.148.113.239\n
google.com.     185 IN  A   209.148.113.238\n
google.com.     185 IN  A   209.148.113.250\n
\r\n\r\n
`
