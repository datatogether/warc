package warc

import (
	"bytes"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
)

// Sha1Digest calculates the shasum of a slice of bytes
func Sha1Digest(data []byte) string {
	hash := sha1.Sum(data)
	buf := &bytes.Buffer{}
	base32.NewEncoder(base32.StdEncoding, buf).Write(hash[:])
	return fmt.Sprintf("sha1:%s", buf.String())
}
