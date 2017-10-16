package warc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// DoRequest is a stand-in for performing an archival http request
// while we work on the API for this package, this may be moved
// into a package of it's own
func DoRequest(req *http.Request) (Records, error) {
	reqr := RequestRecord(req)

	reqr.Headers[warcDate] = time.Now().Format(time.RFC3339)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resr, err := HttpResponseRecord(res)
	if err != nil {
		return nil, err
	}

	return Records{
		reqr,
		resr,
	}, nil
}

func RequestRecord(req *http.Request) *Record {
	body := contentFromHttpRequest(req)

	return &Record{
		Type: RecordTypeRequest,
		Headers: map[string]string{
			contentType:  "application/http; msgtype=request",
			warcRecordId: NewUuid(),
		},
		Content: bytes.NewBuffer(body),
	}
}

func contentFromHttpRequest(req *http.Request) []byte {
	buf := &bytes.Buffer{}

	if req.Method == "" {
		req.Method = "GET"
	}

	buf.WriteString(fmt.Sprintf("%s / %s\r\n", req.Method, req.Proto))
	buf.WriteString(fmt.Sprintf("Host: %s\r\n", req.Host))
	if err := writeHttpHeaders(buf, req.Header); err != nil {
		fmt.Println("error writing to buffer? strange:", err.Error())
	}

	// buf.WriteString(fmt.Sprintf("User-Agent: %s\r\n", req.UserAgent()))
	// TODO - finish

	return buf.Bytes()
}

// HttpResponseRecord creates a record from an HTTP response
func HttpResponseRecord(res *http.Response) (*Record, error) {
	raw, sanitized, err := SanitizeResponse(res)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	writeHttpHeaders(buf, res.Header)
	buf.WriteString("\r\n")
	buf.Write(sanitized)

	resr := &Record{
		Type: RecordTypeResponse,
		Headers: map[string]string{
			warcPayloadDigest: sha1Digest(raw),
			contentType:       "application/http; msgtype=response",
			warcRecordId:      NewUuid(),
		},
		Content: buf,
	}
	return resr, nil
}

func SanitizeResponse(res *http.Response) (raw, sanitized []byte, err error) {
	defer res.Body.Close()

	raw, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	// TODO - lololol finish
	sanitized = bytes.Replace(raw, crlf, []byte("CRLF"), -1)
	return
}
