package formatter

import (
	"bytes"
	"encoding/json"
)

// JSON ...
type JSON struct{}

// Encode ...
func (j *JSON) Encode(v interface{}) ([]byte, error) {
	n := new(bytes.Buffer)
	e := json.NewEncoder(n)

	if err := e.Encode(v); err != nil {
		return nil, err
	}
	return n.Bytes(), nil
}

// Decode ...
func (j *JSON) Decode(b []byte, v interface{}) error {
	n := new(bytes.Buffer)
	n.Write(b)

	e := json.NewDecoder(n)

	if err := e.Decode(v); err != nil {
		return err
	}
	return nil
}
