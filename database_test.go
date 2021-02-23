package goblin

import (
	"testing"
)

func TestNew(t *testing.T) {
	if _, err := New(database, nil); err != nil {
		t.Error("create new database failed.", err)
	}

	if _, err := New(database, &Options{
		Format:      "gob",
		Compression: "gzip",
	}); err != nil {
		t.Error("create new database /w options failed.", err)
	}
}
