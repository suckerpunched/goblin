package formatter

import (
	"bytes"
	"encoding/gob"
)

type GOB struct{}

func (g *GOB) Encode(v interface{}) ([]byte, error) {
	n := new(bytes.Buffer)
	e := gob.NewEncoder(n)

	if err := e.Encode(v); err != nil {
		return nil, err
	}

	return n.Bytes(), nil
}

func (g *GOB) Decode(b []byte, v interface{}) error {
	n := new(bytes.Buffer)
	n.Write(b)

	e := gob.NewDecoder(n)

	if err := e.Decode(v); err != nil {
		return err
	}

	return nil
}
