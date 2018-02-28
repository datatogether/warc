package warc

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestResponseRecords(t *testing.T) {
	const (
		hdr1         = "custom header content fde7073d7b95"
		hdr2         = "custom header content 25b1be4eb31c"
		path         = "/4f3f2471fd8d"
		responseBody = "Response body\n40f9fcaa4120"
		requestStr   = "Request body\n1af849d49e58"
	)
	var warcinfoID = NewUUID()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Custom-Header", hdr2)
		fmt.Fprint(w, responseBody)
	}))
	defer srv.Close()

	requestBody := strings.Repeat(requestStr, 50)
	body := strings.NewReader(requestBody)
	helper := CaptureHelper{
		WarcinfoID:        warcinfoID,
		ReqBodyReadSeeker: body,
	}
	req, err := http.NewRequest("PUT", srv.URL+path, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Custom-Header", hdr1)
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: helper.DialContext(nil),
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	reqRecord, respRecord, err := NewRequestResponseRecords(helper, req, resp)
	if err != nil {
		t.Fatal(err)
	}

	if reqRecord.Headers.Get(FieldNameWARCPayloadDigest) != Sha1Digest([]byte(requestBody)) {
		t.Error("Bad request body digest")
	}
	if respRecord.Headers.Get(FieldNameWARCPayloadDigest) != Sha1Digest([]byte(responseBody)) {
		t.Error("Bad response body digest")
	}

	var buf bytes.Buffer
	reqRecord.Write(&buf)
	str := buf.String()
	// fmt.Println(str)
	if !strings.Contains(str, path) {
		t.Error("Path not found in request record")
	}
	if !strings.Contains(str, hdr1) {
		t.Error("Headers not found in request record")
	}
	if !strings.Contains(str, requestBody) {
		t.Error("Body not found in request record")
	}
	if strings.Contains(str, "Transfer-Encoding: chunked") {
		t.Error("Request written with chunked Transfer-Encoding")
	}

	buf.Reset()
	respRecord.Write(&buf)
	str = buf.String()
	if !strings.Contains(str, path) {
		t.Error("Path not found in response record")
	}
	if !strings.Contains(str, srv.URL) {
		t.Error("Hostname (WARC-Request-URI) not found in response record")
	}
	if !strings.Contains(str, hdr2) {
		t.Error("Headers not found in response record")
	}
	if !strings.Contains(str, responseBody) {
		t.Error("Body not found in response record")
	}
	if !strings.Contains(str, warcinfoID) {
		t.Error("Warcinfo-ID not found in response record")
	}

	// t.Logf("%#v", respRecord.Content)
}
