package warc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func DoRequest(req *http.Request) (Records, error) {
	reqr := RequestRecord(req)

	reqr.Headers[warcDate] = time.Now().Format(time.RFC3339)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resr, err := ResponseRecord(res)
	if err != nil {
		return nil, err
	}

	return Records{
		reqr,
		resr,
	}, nil
}

func ResponseRecord(res *http.Response) (*Record, error) {
	raw, sanitized, err := SanitizeResponse(res)
	if err != nil {
		return nil, err
	}

	resr := &Record{
		Type: RecordTypeResponse,
		Headers: map[string]string{
			warcPayloadDigest: sha1Digest(raw),
			warcBlockDigest:   sha1Digest(sanitized),
		},
		Content: bytes.NewBuffer(sanitized),
	}
	return resr, nil
}

func RequestRecord(req *http.Request) *Record {
	body := BodyFromRequest(req)

	return &Record{
		Type: RecordTypeRequest,
		Headers: map[string]string{
			contentType:     "application/http; msgtype=request",
			warcBlockDigest: sha1Digest(body),
		},
		Content: bytes.NewBuffer(body),
	}
}

func BodyFromRequest(req *http.Request) []byte {
	buf := &bytes.Buffer{}

	if req.Method == "" {
		req.Method = "GET"
	}

	buf.WriteString(fmt.Sprintf("%s / %s\r\n", req.Method, req.Proto))
	buf.WriteString(fmt.Sprintf("User-Agent: %s\r\n", req.UserAgent()))
	buf.WriteString(fmt.Sprintf("Host: %s\r\n", req.Host))
	// TODO - finish

	return buf.Bytes()
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
