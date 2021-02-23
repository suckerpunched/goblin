package goblin

import (
	"sync"

	"./backend"
	"./compression"
	"./formatter"
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

// -------
// func (D *Driver) Write(collection, resource string, v interface{}) {
// 	if err := notEmpty([]string{collection, resource}); err != nil {
// 		return err
// 	}

// 	mutex := D.Driver.obtainMutex(collection)
// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	dir := filepath.Join(D.Driver.Path, collection)
// 	path := filepath.Join(dir, resource+"."+D.Options.ext)

// 	if err := os.MkdirAll(dir, 0755); err != nil {
// 		return err
// 	}

// 	b, _ := D.Formatter.Encode(v)

// 	if D.Compression != nil {
// 		b, _ = D.Compression.Compress(b)
// 	}

// 	return D.Backend.Write(path, b)
// }

// func (D *Driver) Read(path string, v interface{}) {
// 	b, _ := D.Backend.Read(path)
// 	b, _ = D.Compression.Decompress(b)
// 	D.Formatter.Decode(b, v)
// }
