package warc

import (
	"errors"
	"io"
)

// Parse parses the sql and returns a set of Records, which
// is the AST representation of the query.
func Parse(r io.Reader) ([]*record, error) {
	tokenizer := NewTokenizer(r)
	if yyParse(tokenizer) != 0 {
		return nil, errors.New(tokenizer.LastError)
	}
	return tokenizer.Records, nil
}
