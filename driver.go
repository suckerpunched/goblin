package goblin

import (
	"sync"

	"github.com/suckerpunched/goblin/backend"
	"github.com/suckerpunched/goblin/compression"
	"github.com/suckerpunched/goblin/formatter"
)

// Driver ...
type Driver struct {
	Mutex   sync.Mutex
	Engaged map[string]*sync.Mutex
	Path    string

	Formatter interface {
		Encode(v interface{}) ([]byte, error)
		Decode(b []byte, v interface{}) error
	}

	Backend interface {
		Write(string, []byte) error
		Read(string) ([]byte, error)
	}

	Compression interface {
		Compress(b []byte) ([]byte, error)
		Decompress(b []byte) ([]byte, error)
	}
}

func (D *Driver) configureFormatter(opt string) {
	switch opt {
	case "json":
		D.Formatter = &formatter.JSON{}
	case "gob":
		D.Formatter = &formatter.GOB{}
	default:
		D.Formatter = &formatter.JSON{}
	}
}

func (D *Driver) configureCompression(opt string) {
	switch opt {
	case "gzip":
		D.Compression = &compression.GZIP{}
	default:
		D.Compression = nil
	}
}

func (D *Driver) configureBackend(opt string) {
	switch opt {
	case "local":
		D.Backend = &backend.Local{}
	default:
		D.Backend = &backend.Local{}
	}
}

func (D *Driver) obtainMutex(collection string) *sync.Mutex {
	D.Mutex.Lock()
	defer D.Mutex.Unlock()

	m, ok := D.Engaged[collection]
	if !ok {
		m = &sync.Mutex{}
		D.Engaged[collection] = m
	}

	return m
}
