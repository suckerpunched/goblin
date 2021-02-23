package compression

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// GZIP ...
type GZIP struct{}

// Compress ...
func (g *GZIP) Compress(b []byte) ([]byte, error) {
	n := new(bytes.Buffer)
	gz := gzip.NewWriter(n)

	defer gz.Close()

	if _, err := gz.Write(b); err != nil {
		return nil, err
	}

	if err := gz.Flush(); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return n.Bytes(), nil
}

// Decompress ...
func (g *GZIP) Decompress(b []byte) ([]byte, error) {
	n := bytes.NewBuffer(b)
	gz, err := gzip.NewReader(n)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	d, err := ioutil.ReadAll(gz)
	if err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return d, nil
}
