package warc

import (
	"bytes"
	// "fmt"
	// "io"
	"io/ioutil"
	"os"
	"testing"
)

func TestWarcWrite(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.warc")
	if err != nil {
		t.Error(err.Error())
		return
	}

	records, err := ParseAll(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
		return
	}

	buf := &bytes.Buffer{}

	for _, r := range records {
		if r.(*Resource).WARCBlockDigest == "sha1:28ee620ee6d9ed280505fa9faca0ba357db82ffd" {
			// io.Copy(os.Stdout, r.GetContent())
		}
	}

	if err := WriteRecords(buf, records); err != nil {
		t.Error(err)
		return
	}

	// fmt.Println(len(data), len(buf.Bytes()))
	ioutil.WriteFile("testdata/out.warc", buf.Bytes(), os.ModePerm)
}
