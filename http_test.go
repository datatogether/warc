package warc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoRequest(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello!")
	}))

	req, _ := http.NewRequest("GET", s.URL, nil)
	records, err := DoRequest(req)
	if err != nil {
		t.Error(err.Error())
		return
	}

	buf := bytes.NewBuffer(nil)
	for _, rec := range records {
		if err := rec.Write(buf); err != nil {
			t.Error(err.Error())
		}
	}

	fmt.Println(buf.String())
}
