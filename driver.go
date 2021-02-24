package goblin

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
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

	Log *logrus.Logger
}

func (D *Driver) configureLogger() {
	// D.Log.SetFormatter(&logrus.JSONFormatter{})
	D.Log.SetOutput(os.Stdout)
	D.Log.SetLevel(logrus.DebugLevel)
}

func (D *Driver) configureFormatter(opt string) {
	switch opt {
	case "json":
		D.Log.WithFields(logrus.Fields{"fmt": "json"}).Debug("formatter configured")
		D.Formatter = &formatter.JSON{}
	case "gob":
		D.Log.WithFields(logrus.Fields{"fmt": "gob"}).Debug("formatter configured")
		D.Formatter = &formatter.GOB{}
	default:
		D.Log.WithFields(logrus.Fields{"fmt": "json"}).Debug("formatter configured")
		D.Formatter = &formatter.JSON{}
	}
}

func (D *Driver) configureCompression(opt string) {
	switch opt {
	case "gzip":
		D.Log.WithFields(logrus.Fields{"compression": "gzip"}).Debug("compression configured")
		D.Compression = &compression.GZIP{}
	default:
		D.Compression = nil
	}
}

func (D *Driver) configureBackend(opt string) {
	switch opt {
	case "local":
		D.Log.WithFields(logrus.Fields{"backend": "local"}).Debug("backend configured")
		D.Backend = &backend.Local{}
	default:
		D.Log.WithFields(logrus.Fields{"backend": "local"}).Debug("backend configured")
		D.Backend = &backend.Local{}
	}
}

func (D *Driver) obtainMutex(collection string) *sync.Mutex {
	D.Log.WithFields(logrus.Fields{"collection": collection}).Debug("obtaining mutex")

	D.Mutex.Lock()
	defer D.Mutex.Unlock()

	m, ok := D.Engaged[collection]
	if !ok {
		m = &sync.Mutex{}
		D.Engaged[collection] = m
	}

	return m
}
