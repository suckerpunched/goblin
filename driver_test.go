package goblin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDriverFlow(t *testing.T) {
	createDatabase("gob", "gzip")
	setupTestingData()

	dir := filepath.Join(db.Driver.Path, collection)
	path := filepath.Join(dir, "encode_compress"+"."+db.Options.ext)

	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error("mkdirall failed", err)
	}

	driverEncode(t, path)
	driverDecode(t, path)

	if y.Size != x.Size {
		t.Errorf("driver flow failed, got: \"%d\", expected: \"%d\".", y.Size, x.Size)
	}
}

func driverEncode(t *testing.T, path string) {
	b, err := db.Driver.Formatter.Encode(&x)
	if err != nil {
		t.Error("driver encode failed", err)
	}

	b, err = db.Driver.Compression.Compress(b)
	if err != nil {
		t.Error("driver compress failed", err)
	}

	err = db.Driver.Backend.Write(path, b)
	if err != nil {
		t.Error("driver write failed", err)
	}
}

func driverDecode(t *testing.T, path string) {
	b, err := db.Driver.Backend.Read(path)
	if err != nil {
		t.Error("driver read failed", err)
	}

	b, err = db.Driver.Compression.Decompress(b)
	if err != nil {
		t.Error("driver decompress failed", err)
	}

	err = db.Driver.Formatter.Decode(b, &y)
	if err != nil {
		t.Error("driver decode failed", err)
	}
}
