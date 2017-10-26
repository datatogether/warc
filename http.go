package warc

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/datatogether/warc/warc"
)

// DoRequest is a stand-in for performing an archival http request
// while we work on the API for this package, this may be moved
// into a package of it's own
func DoRequest(req *http.Request) (warc.Records, error) {
	reqr := RequestRecord(req)

	reqr.Headers[warc.FieldNameWarcDate] = time.Now().Format(time.RFC3339)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resr, err := HttpResponseRecord(res)
	if err != nil {
		return nil, err
	}

	return warc.Records{
		reqr,
		resr,
	}, nil
}

func RequestRecord(req *http.Request) *warc.Record {
	body := contentFromHttpRequest(req)
	return &warc.Record{
		Type: warc.RecordTypeRequest,
		Headers: map[string]string{
			warc.FieldNameContentType:  "application/http; msgtype=request",
			warc.FieldNameWarcRecordId: warc.NewUuid(),
		},
		Content: bytes.NewBuffer(body),
	}
}

func contentFromHttpRequest(req *http.Request) []byte {
	buf := &bytes.Buffer{}

	if err := warc.WriteRequestStatusAndHeaders(buf, req); err != nil {
		return buf.Bytes()
	}

	// buf.WriteString(fmt.Sprintf("%s / %s\r\n", req.Method, req.Proto))
	// buf.WriteString(fmt.Sprintf("Host: %s\r\n", req.Host))
	// if err := writeHttpHeaders(buf, req.Header); err != nil {
	// 	fmt.Println("error writing to buffer? strange:", err.Error())
	// }

	// buf.WriteString(fmt.Sprintf("User-Agent: %s\r\n", req.UserAgent()))
	// TODO - finish

	return buf.Bytes()
}

// HttpResponseRecord creates a record from an HTTP response
func HttpResponseRecord(res *http.Response) (*warc.Record, error) {
	raw, sanitized, err := SanitizeResponse(res)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	warc.WriteHttpHeaders(buf, res.Header)
	buf.WriteString("\r\n")
	buf.Write(sanitized)

	resr := &warc.Record{
		Type: warc.RecordTypeResponse,
		Headers: map[string]string{
			warc.FieldNameWarcPayloadDigest: warc.Sha1Digest(raw),
			warc.FieldNameContentType:       "application/http; msgtype=response",
			warc.FieldNameWarcRecordId:      warc.NewUuid(),
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

	sanitized = warc.Sanitize(raw)
	return
}
