package warc

import (
	"bytes"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"os"
	"strings"
	"testing"

	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

const testRecordID = "<urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>"

func init() {
	for _, t := range []*[]byte{
		&WARCInfoRecord,
		&ResponseRecord,
		&ResponseRecord2,
		&RequestRecord,
		&RequestRecord2,
		&RevisitRecord1,
		&RevisitRecord2,
		&ResourceRecord,
		&MetadataRecord,
		&DNSResponseRecord,
		&DNSResourceRecord,
	} {
		// need to replace '\r' from raw string literals with actual
		// carriage return character
		*t = bytes.Replace(*t, []byte{'\\', 'r'}, []byte{0x0d}, -1)
	}
}

func TestNewUUID(t *testing.T) {
	id := NewUUID()
	if !strings.HasPrefix(id, "<urn:uuid:") {
		t.Errorf("expected prefix: '%s', got: %s", "<urn:uuid:", id)
	}
}

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

func TestWarcinfoRecord(t *testing.T) {
	rec := &Record{
		Format: RecordFormatWarc,
		Type:   RecordTypeWarcInfo,
		Headers: map[string]string{
			FieldNameWARCRecordID:  testRecordID,
			FieldNameWARCType:      RecordTypeWarcInfo.String(),
			FieldNameWARCFilename:  "testfile.warc.gz",
			FieldNameWARCDate:      "2000-01-01T00:00:00Z",
			FieldNameContentType:   "application/warc-fields",
			FieldNameContentLength: "86",
		},
		Content: bytes.NewBuffer([]byte("software: recorder test\r\n" +
			"format: WARC File Format 1.0\r\n" +
			"json-metadata: {\"foo\": \"bar\"}\r\n")),
	}

	if err := testWriteRecord(rec, WARCInfoRecord); err != nil {
		t.Error(err)
	}
}

func TestRequestRecord(t *testing.T) {
	rec := &Record{
		Format: RecordFormatWarc,
		Type:   RecordTypeRequest,
		Headers: map[string]string{
			FieldNameWARCType:          RecordTypeRequest.String(),
			FieldNameWARCRecordID:      testRecordID,
			FieldNameWARCTargetURI:     "http://example.com/",
			FieldNameWARCDate:          "2000-01-01T00:00:00Z",
			FieldNameWARCPayloadDigest: "sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ",
			FieldNameWARCBlockDigest:   "sha1:ONEHF6PTXPTTHE3333XHTD2X45TZ3DTO",
			FieldNameContentType:       "application/http; msgtype=request",
			FieldNameContentLength:     "54",
		},
		Content: bytes.NewBuffer([]byte("GET / HTTP/1.0\r\n" +
			"User-Agent: foo\r\n" +
			"Host: example.com\r\n" +
			"\r\n")),
	}

	if err := testWriteRecord(rec, RequestRecord); err != nil {
		t.Error(err)
	}
}

func TestResponseRecord(t *testing.T) {
	rec := &Record{
		Format: RecordFormatWarc,
		Type:   RecordTypeResponse,
		Headers: map[string]string{
			FieldNameContentLength:     "97",
			FieldNameContentType:       "application/http; msgtype=response",
			FieldNameWARCBlockDigest:   "sha1:OS3OKGCWQIJOAOC3PKXQOQFD52NECQ74",
			FieldNameWARCDate:          "2000-01-01T00:00:00Z",
			FieldNameWARCPayloadDigest: "sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O",
			FieldNameWARCRecordID:      "<urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>",
			FieldNameWARCTargetURI:     "http://example.com/",
			FieldNameWARCType:          RecordTypeResponse.String(),
		},
		Content: bytes.NewBuffer([]byte("HTTP/1.0 200 OK\r\n" +
			"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
			"Custom-Header: somevalue\r\n" +
			"\r\n" +
			"some\n" +
			"text")),
	}

	if err := testWriteRecord(rec, ResponseRecord); err != nil {
		t.Error(err)
	}
}

func testWriteRecord(r *Record, expect []byte) error {
	if r.ContentLength() != r.Content.Len() {
		return fmt.Errorf("Record Content-Length mistmatch: %d != %d", r.ContentLength(), r.Content.Len())
	}

	buf := &bytes.Buffer{}
	if err := r.Write(buf); err != nil {
		return fmt.Errorf("error writing record: %s", err.Error())
	}

	if len(buf.Bytes()) != len(expect) {
		dmp := dmp.New()
		diffs := dmp.DiffMain(buf.String(), string(expect), true)
		fmt.Println("error diff output:")
		fmt.Println(dmp.DiffPrettyText(diffs))

		for i, b := range buf.Bytes() {
			if i >= len(expect) || b != expect[i] {
				return fmt.Errorf("byte length mismatch. expected: %d, got: %d. first error at index %d: '%#v'", len(expect), len(buf.Bytes()), i, b)
			}
		}

		return fmt.Errorf("byte length mismatch. expected: %d, got: %d, ", len(expect), len(buf.Bytes()))
	}

	if !bytes.Equal(buf.Bytes(), expect) {
		return fmt.Errorf("byte mismatch: %s != %s", buf.String(), string(expect))
	}

	if r.Headers[FieldNameWARCBlockDigest] != "" {
		checkSha1Hash(r.Content.Bytes(), r.Headers[FieldNameWARCBlockDigest])
	}

	return nil
}

func checkSha1Hash(content []byte, hashstr string) error {
	hash := sha1.Sum(content)
	buf := &bytes.Buffer{}
	base32.NewEncoder(base32.StdEncoding, buf).Write(hash[:])
	s := fmt.Sprintf("sha1:%s", buf.String())
	if s != hashstr {
		return fmt.Errorf("hash mismatch. expected '%s'. got: '%s'", hashstr, s)
	}
	return nil
}

// func testRequestResponseConcur(t *testing.T) {
// }

// func testReadFromStreamNoContentLength(t *testing.T) {
// }

func validateResponse(r *Record) error {
	return nil
}

var WARCInfoRecord = []byte(`WARC/1.0\r
Content-Length: 86\r
Content-Type: application/warc-fields\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Filename: testfile.warc.gz\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Type: warcinfo\r
\r
software: recorder test\r
format: WARC File Format 1.0\r
json-metadata: {"foo": "bar"}\r
\r
\r
`)

var ResponseRecord = []byte(`WARC/1.0\r
Content-Length: 97\r
Content-Type: application/http; msgtype=response\r
WARC-Block-Digest: sha1:OS3OKGCWQIJOAOC3PKXQOQFD52NECQ74\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: http://example.com/\r
WARC-Type: response\r
\r
HTTP/1.0 200 OK\r
Content-Type: text/plain; charset="UTF-8"\r
Custom-Header: somevalue\r
\r
some
text\r
\r
`)

var ResponseRecord2 = []byte(`
WARC/1.0\r
WARC-Type: response\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: http://example.com/\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r
WARC-Block-Digest: sha1:U6KNJY5MVNU3IMKED7FSO2JKW6MZ3QUX\r
Content-Type: application/http; msgtype=response\r
Content-Length: 145\r
\r
HTTP/1.0 200 OK\r
Content-Type: text/plain; charset="UTF-8"\r
Content-Length: 9\r
Custom-Header: somevalue\r
Content-Encoding: x-unknown\r
\r
some
text\r
\r
`)

var RequestRecord = []byte(`WARC/1.0\r
Content-Length: 54\r
Content-Type: application/http; msgtype=request\r
WARC-Block-Digest: sha1:ONEHF6PTXPTTHE3333XHTD2X45TZ3DTO\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: http://example.com/\r
WARC-Type: request\r
\r
GET / HTTP/1.0\r
User-Agent: foo\r
Host: example.com\r
\r
\r
\r
`)

var RequestRecord2 = []byte(`
WARC/1.0\r
WARC-Type: request\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: http://example.com/\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:R5VZAKIE53UW5VGK43QJIFYS333QM5ZA\r
WARC-Block-Digest: sha1:L7SVBUPPQ6RH3ANJD42G5JL7RHRVZ5DV\r
Content-Type: application/http; msgtype=request\r
Content-Length: 92\r
\r
POST /path HTTP/1.0\r
Content-Type: application/json\r
Content-Length: 17\r
\r
{"some": "value"}\r
\r
`)

var RevisitRecord1 = []byte(`
WARC/1.0\r
WARC-Type: revisit\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: http://example.com/\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Profile: http://netpreserve.org/warc/1.0/revisit/identical-payload-digest\r
WARC-Refers-To-Target-URI: http://example.com/foo\r
WARC-Refers-To-Date: 1999-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r
WARC-Block-Digest: sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ\r
Content-Type: application/http; msgtype=response\r
Content-Length: 0\r
\r
\r
\r
`)

var RevisitRecord2 = []byte(`
WARC/1.0\r
WARC-Type: revisit\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: http://example.com/\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Profile: http://netpreserve.org/warc/1.0/revisit/identical-payload-digest\r
WARC-Refers-To-Target-URI: http://example.com/foo\r
WARC-Refers-To-Date: 1999-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r
WARC-Block-Digest: sha1:A6J5UTI2QHHCZFCFNHQHCDD3JJFKP53V\r
Content-Type: application/http; msgtype=response\r
Content-Length: 88\r
\r
HTTP/1.0 200 OK\r
Content-Type: text/plain; charset="UTF-8"\r
Custom-Header: somevalue\r
\r
\r
\r
`)

var ResourceRecord = []byte(`
WARC/1.0\r
WARC-Type: resource\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: ftp://example.com/\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r
WARC-Block-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r
Content-Type: text/plain\r
Content-Length: 9\r
\r
some
text\r
\r
`)

var MetadataRecord = []byte(`
WARC/1.0\r
WARC-Type: metadata\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: http://example.com/\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:ZOLBLKAQVZE5DXH56XE6EH6AI6ZUGDPT\r
WARC-Block-Digest: sha1:ZOLBLKAQVZE5DXH56XE6EH6AI6ZUGDPT\r
Content-Type: application/json\r
Content-Length: 67\r
\r
{"metadata": {"nested": "obj", "list": [1, 2, 3], "length": "123"}}\r
\r
`)

var DNSResponseRecord = []byte(`
WARC/1.0\r
WARC-Type: response\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: dns:google.com\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:2AAVJYKKIWK5CF6EWE7PH63EMNLO44TH\r
WARC-Block-Digest: sha1:2AAVJYKKIWK5CF6EWE7PH63EMNLO44TH\r
Content-Type: application/http; msgtype=response\r
Content-Length: 147\r
\r
20170509000739
google.com.     185 IN  A   209.148.113.239
google.com.     185 IN  A   209.148.113.238
google.com.     185 IN  A   209.148.113.250
\r\r
`)

var DNSResourceRecord = []byte(`
WARC/1.0\r
WARC-Type: resource\r
WARC-Record-ID: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r
WARC-Target-URI: dns:google.com\r
WARC-Date: 2000-01-01T00:00:00Z\r
WARC-Payload-Digest: sha1:2AAVJYKKIWK5CF6EWE7PH63EMNLO44TH\r
WARC-Block-Digest: sha1:2AAVJYKKIWK5CF6EWE7PH63EMNLO44TH\r
Content-Type: application/warc-record\r
Content-Length: 147\r
\r
20170509000739
google.com.     185 IN  A   209.148.113.239
google.com.     185 IN  A   209.148.113.238
google.com.     185 IN  A   209.148.113.250
\r\r
`)
