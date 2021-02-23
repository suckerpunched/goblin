package goblin

import (
	"os"
	"testing"
)

type goblin struct {
	Size int
}

var (
	db         *Database
	database   = "./gold"
	collection = "bank"

	x goblin
	y goblin
)

func setupTestingData() {
	x = goblin{Size: 5}
	y = goblin{}
}

func createDatabase(format string, compression string) error {
	var err error
	if db, err = New(database, &Options{
		Format:      format,
		Compression: compression,
	}); err != nil {
		return err
	}

	return nil
}

func TestMain(m *testing.M) {
	os.RemoveAll("./gold")

	code := m.Run()

	os.RemoveAll("./gold")
	os.Exit(code)
}

func TestWriteReadGOB(t *testing.T) {
	createDatabase("gob", "")
	setupTestingData()

	var err error

	err = db.Write(collection, "write-read", &x)
	if err != nil {
		t.Error("database write failed,", err)
	}

	err = db.Read(collection, "write-read", &y)
	if err != nil {
		t.Error("database read failed,", err)
	}

	if y.Size != x.Size {
		t.Errorf("database flow failed, got: \"%d\", expected: \"%d\".", y.Size, x.Size)
	}
}

func TestWriteReadCompressedGOB(t *testing.T) {
	createDatabase("gob", "gzip")
	setupTestingData()

	var err error

	err = db.Write(collection, "write-read", &x)
	if err != nil {
		t.Error("database write failed,", err)
	}

	err = db.Read(collection, "write-read", &y)
	if err != nil {
		t.Error("database read failed,", err)
	}

	if y.Size != x.Size {
		t.Errorf("database flow failed, got: \"%d\", expected: \"%d\".", y.Size, x.Size)
	}
}

func TestWriteReadJSON(t *testing.T) {
	createDatabase("json", "")
	setupTestingData()

	var err error

	err = db.Write(collection, "write-read", &x)
	if err != nil {
		t.Error("database write failed,", err)
	}

	err = db.Read(collection, "write-read", &y)
	if err != nil {
		t.Error("database read failed,", err)
	}

	if y.Size != x.Size {
		t.Errorf("database flow failed, got: \"%d\", expected: \"%d\".", y.Size, x.Size)
	}
}

func TestWriteReadCompressedJSON(t *testing.T) {
	createDatabase("json", "gzip")
	setupTestingData()

	var err error

	err = db.Write(collection, "write-read", &x)
	if err != nil {
		t.Error("database write failed,", err)
	}

	err = db.Read(collection, "write-read", &y)
	if err != nil {
		t.Error("database read failed,", err)
	}

	if y.Size != x.Size {
		t.Errorf("database flow failed, got: \"%d\", expected: \"%d\".", y.Size, x.Size)
	}
}
