package warc

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

const testRecordId = "<urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>"

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
		Version: WARC_VERSION,
		Headers: map[string]string{
			warcRecordId:  testRecordId,
			warcType:      RecordTypeWarcInfo.String(),
			warcFilename:  "testfile.warc.gz",
			warcDate:      "2000-01-01T00:00:00Z",
			contentType:   "application/warc-fields",
			contentLength: "86",
		},
		Content: []byte("software: recorder test\r\n" +
			"format: WARC File Format 1.0\r\n" +
			"json-metadata: {\"foo\": \"bar\"}\r\n"),
	}

	if err := testWriteRecord(rec, WARCINFO_RECORD); err != nil {
		t.Error(err)
	}
}

func TestRequestRecord(t *testing.T) {
	rec := &Record{
		Version: WARC_VERSION,
		Headers: map[string]string{
			warcType:          RecordTypeRequest.String(),
			warcRecordId:      testRecordId,
			warcTargetUri:     "http://example.com/",
			warcDate:          "2000-01-01T00:00:00Z",
			warcPayloadDigest: "sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ",
			warcBlockDigest:   "sha1:ONEHF6PTXPTTHE3333XHTD2X45TZ3DTO",
			contentType:       "application/http; msgtype=request",
			contentLength:     "54",
		},
		Content: []byte("GET / HTTP/1.0\r\n" +
			"User-Agent: foo\r\n" +
			"Host: example.com\r\n" +
			"\r\n"),
	}

	if err := testWriteRecord(rec, REQUEST_RECORD); err != nil {
		t.Error(err)
	}
}

func TestResponseRecord(t *testing.T) {
	rec := &Record{
		Version: WARC_VERSION,
		Headers: map[string]string{
			contentLength:     "97",
			contentType:       "application/http; msgtype=response",
			warcBlockDigest:   "sha1:OS3OKGCWQIJOAOC3PKXQOQFD52NECQ74",
			warcDate:          "2000-01-01T00:00:00Z",
			warcPayloadDigest: "sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O",
			warcRecordId:      "<urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>",
			warcTargetUri:     "http://example.com/",
			warcType:          RecordTypeResponse.String(),
		},
		Content: []byte("HTTP/1.0 200 OK\r\n" +
			"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
			"Custom-Header: somevalue\r\n" +
			"\r\n" +
			"some\n" +
			"text"),
	}

	if err := testWriteRecord(rec, RESPONSE_RECORD); err != nil {
		t.Error(err)
	}
}

func testWriteRecord(r *Record, expect []byte) error {
	buf := &bytes.Buffer{}
	if err := r.Write(buf); err != nil {
		return fmt.Errorf("error writing record: %s", err.Error())
	}

	if r.ContentLength() != len(r.Content) {
		return fmt.Errorf("Record Content-Length mistmatch: %d != %d", r.ContentLength(), len(r.Content))
	}

	if len(buf.Bytes()) != len(expect) {
		dmp := dmp.New()
		diffs := dmp.DiffMain(buf.String(), string(expect), true)
		fmt.Println("error diff output:")
		fmt.Println(dmp.DiffPrettyText(diffs))

		for i, b := range buf.Bytes() {
			if b != expect[i] {
				return fmt.Errorf("byte length mismatch. expected: %d, got: %d. first error at index %d: '%#v'", len(expect), len(buf.Bytes()), i, b)
			}
		}

		return fmt.Errorf("byte length mismatch. expected: %d, got: %d, ", len(expect), len(buf.Bytes()))
	}

	if !bytes.Equal(buf.Bytes(), expect) {
		return fmt.Errorf("byte mismatch: %s != %s", buf.String(), string(expect))
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

var WARCINFO_RECORD = []byte("WARC/1.0\r\n" +
	"Content-Length: 86\r\n" +
	"Content-Type: application/warc-fields\r\n" +
	"Warc-Date: 2000-01-01T00:00:00Z\r\n" +
	"Warc-Filename: testfile.warc.gz\r\n" +
	"Warc-Record-Id: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n" +
	"Warc-Type: warcinfo\r\n" +
	"\r\n" +
	"software: recorder test\r\n" +
	"format: WARC File Format 1.0\r\n" +
	"json-metadata: {\"foo\": \"bar\"}\r\n" +
	"\r\n" +
	"\r\n")

var REQUEST_RECORD = []byte("WARC/1.0\r\n" +
	"Content-Length: 54\r\n" +
	"Content-Type: application/http; msgtype=request\r\n" +
	"Warc-Block-Digest: sha1:ONEHF6PTXPTTHE3333XHTD2X45TZ3DTO\r\n" +
	"Warc-Date: 2000-01-01T00:00:00Z\r\n" +
	"Warc-Payload-Digest: sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ\r\n" +
	"Warc-Record-Id: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n" +
	"Warc-Target-Uri: http://example.com/\r\n" +
	"Warc-Type: request\r\n" +
	"\r\n" +
	"GET / HTTP/1.0\r\n" +
	"User-Agent: foo\r\n" +
	"Host: example.com\r\n" +
	"\r\n" +
	"\r\n" +
	"\r\n")

var RESPONSE_RECORD = []byte("WARC/1.0\r\n" +
	"Content-Length: 97\r\n" +
	"Content-Type: application/http; msgtype=response\r\n" +
	"Warc-Block-Digest: sha1:OS3OKGCWQIJOAOC3PKXQOQFD52NECQ74\r\n" +
	"Warc-Date: 2000-01-01T00:00:00Z\r\n" +
	"Warc-Payload-Digest: sha1:B6QJ6BNJ3R4B23XXMRKZKHLPGJY2VE4O\r\n" +
	"Warc-Record-Id: <urn:uuid:12345678-feb0-11e6-8f83-68a86d1772ce>\r\n" +
	"Warc-Target-Uri: http://example.com/\r\n" +
	"Warc-Type: response\r\n" +
	"\r\n" +
	"HTTP/1.0 200 OK\r\n" +
	"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
	"Custom-Header: somevalue\r\n" +
	"\r\n" +
	"some\n" +
	"text\r\n" +
	"\r\n")
