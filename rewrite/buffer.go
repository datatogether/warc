package rewrite

import (
	"bytes"
)

// Buffer behaves just like a bytes.Buffer, but
// uses a rewriter to adjust any bytes written
// with buffer.Write
type Buffer struct {
	bytes.Buffer
	rw Rewriter
}

// NewBuffer allocates a new rewriting buffer.
// Unlike bytes.Buffer, users should always use
// NewBuffer, even if passing nil for data
func NewBuffer(data []byte, rw Rewriter) *Buffer {
	return &Buffer{
		Buffer: *bytes.NewBuffer(data),
		rw:     rw,
	}
}

func (rwb *Buffer) Write(p []byte) (int, error) {
	rw := rwb.rw.Rewrite(p)
	return rwb.Buffer.Write(rw)
}
