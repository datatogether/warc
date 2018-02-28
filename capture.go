package warc

import (
	"bytes"
	"context"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// CaptureHelper is used for the NewRequestResponseRecords() method. Additional
// fields may be added in the future.
type CaptureHelper struct {
	WarcinfoID string
	RemoteAddr string

	// The request body will need to be read multiple times, so please provide
	// one of the following.  (note: bytes.Reader and strings.Reader are
	// ReadSeekers.)
	ReqBodyReadSeeker  io.ReadSeeker
	ReqBodyBytesBuffer *bytes.Buffer
}

// DialContext returns a wrapper around net.DialContext that saves the
// connected-to IP in the CaptureHelper.
func (c *CaptureHelper) DialContext(dialer *net.Dialer) func(ctx context.Context, network, addr string) (net.Conn, error) {
	if dialer == nil {
		dialer = &net.Dialer{}
	}
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := dialer.DialContext(ctx, network, addr)
		if err != nil {
			c.RemoteAddr = conn.RemoteAddr().String()
		}
		return conn, err
	}
}

func (c *CaptureHelper) resetRequestBody(req *http.Request) io.Reader {
	if req.Body == nil || req.Body == http.NoBody {
		return http.NoBody
	}
	if req.GetBody != nil {
		body, err := req.GetBody()
		if err != nil {
			panic(err)
		}
		return body
	}
	if c.ReqBodyReadSeeker != nil {
		c.ReqBodyReadSeeker.Seek(0, io.SeekStart)
		return c.ReqBodyReadSeeker
	}
	if c.ReqBodyBytesBuffer != nil {
		return c.ReqBodyBytesBuffer
	}
	panic("capturehelper: unable to rewind request body")
}

// NewRequestResponseRecords creates a new request/response record pair for the
// provided HTTP request and response.
//
// Make sure to provide the request Body in the CaptureHelper so it can be read
// from again.  The response Body should not yet have been used; if the caller
// needs the body, replace it with an ioutil.NopCloser(io.TeeReader) (the
// caller is then responsible for calling body.Close()).
func NewRequestResponseRecords(info CaptureHelper, req *http.Request, resp *http.Response) (Record, Record, error) {
	reqRec := Record{Format: RecordFormatWarc, Type: RecordTypeRequest, Headers: make(Header)}
	respRec := Record{Format: RecordFormatWarc, Type: RecordTypeResponse, Headers: make(Header)}
	reqUID := NewUUID()
	respUID := NewUUID()
	reqRec.Headers.Set(FieldNameWARCRecordID, reqUID)
	respRec.Headers.Set(FieldNameWARCRecordID, respUID)
	respRec.Headers.Set(FieldNameWARCConcurrentTo, reqUID)
	eventStamp := time.Now().Format(time.RFC3339)
	reqRec.Headers.Set(FieldNameWARCDate, eventStamp)
	respRec.Headers.Set(FieldNameWARCDate, eventStamp)
	u2 := new(url.URL)
	*u2 = *req.URL
	if u2.Host == "" {
		u2.Host = req.Host
	}
	reqRec.Headers.Set(FieldNameWARCTargetURI, u2.String())
	respRec.Headers.Set(FieldNameWARCTargetURI, u2.String())
	if info.WarcinfoID != "" {
		reqRec.Headers.Set(FieldNameWARCWarcinfoID, info.WarcinfoID)
		respRec.Headers.Set(FieldNameWARCWarcinfoID, info.WarcinfoID)
	}
	if info.RemoteAddr != "" {
		ip, _, err := net.SplitHostPort(info.RemoteAddr)
		if err != nil {
			if strings.Contains(err.Error(), "missing port in address") {
				ip = info.RemoteAddr
				err = nil
			}
		}
		if err != nil {
			return reqRec, respRec, errors.Wrap(err, "Bad RemoteAddr value in CaptureHelper")
		}
		reqRec.Headers.Set(FieldNameWARCIPAddress, ip)
		respRec.Headers.Set(FieldNameWARCIPAddress, ip)
	}

	// Write request using stdlib
	reqDigester := sha1.New()
	reqRec.Content = new(bytes.Buffer)
	clonedBody := info.resetRequestBody(req)
	if clonedBody != nil {
		teedBody := io.TeeReader(clonedBody, reqDigester)
		req.Body = ioutil.NopCloser(teedBody)
	}
	err := req.Write(reqRec.Content)
	reqRec.Headers.Set(FieldNameWARCPayloadDigest, formatDigest(reqDigester.Sum(nil)))
	if err != nil {
		return reqRec, respRec, errors.Wrap(err, "writing request body")
	}
	reqRec.Headers[FieldNameWARCBlockDigest] = Sha1Digest(reqRec.Content.Bytes())

	// Write response
	// Can't use stdlib, as it does extra processing of Content-Length, transfer encodings, etc
	respDigester := sha1.New()
	respRec.Content = new(bytes.Buffer)
	teedBody := io.TeeReader(resp.Body, respDigester)

	text := resp.Status
	text = strings.TrimPrefix(text, strconv.Itoa(resp.StatusCode)+" ")
	fmt.Fprintf(respRec.Content, "HTTP/%d.%d %03d %s\r\n", resp.ProtoMajor, resp.ProtoMinor, resp.StatusCode, text)
	resp.Header.Write(respRec.Content)
	io.WriteString(respRec.Content, "\r\n")
	_, err = io.Copy(respRec.Content, teedBody)
	respRec.Headers.Set(FieldNameWARCPayloadDigest, formatDigest(respDigester.Sum(nil)))
	if err != nil {
		return reqRec, respRec, errors.Wrap(err, "writing response body")
	}
	// block digest will be set in Record.Write()

	return reqRec, respRec, nil
}
