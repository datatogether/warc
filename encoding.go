package warc

import (
	"bytes"
)

// TODO - part of a planned support for golang standard
// encoding / decoding.
// WARCMarshaler follows other golang encoding / decoding
// interfaces. See encoding/json.JSONMarshaler as an example.
//
// type WARCMarshaler interface {
// 	MarshalWARC() (data []byte, err error)
// }
//
// type WARCUnmarshaler interface {
// 	UnmarshalWARC(data []byte) (err error)
// }

// UnmarshalRecord reads a single record from data
func UnmarshalRecord(data []byte) (Record, error) {
	r, err := NewReader(bytes.NewReader(data))
	if err != nil {
		return Record{}, err
	}
	return r.readRecord()
}

// UnmarshalRecords reads a slice of records from a slice of bytes
func UnmarshalRecords(data []byte) (Records, error) {
	r, err := NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return r.ReadAll()
}
