package goblin

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
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

	Log zerolog.Logger
}

func (D *Driver) configureLogger(level string, formatted bool) {

	if formatted {
		D.Log = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		D.Log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	}

	switch level {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
}

func (D *Driver) configureFormatter(opt string) {
	switch opt {
	case "json":
		D.Log.Debug().Str("format", "json").Msg("formatter")
		D.Formatter = &formatter.JSON{}
	case "gob":
		D.Log.Debug().Str("format", "gob").Msg("formatter")
		D.Formatter = &formatter.GOB{}
	default:
		D.Log.Debug().Str("format", "json").Msg("formatter")
		D.Formatter = &formatter.JSON{}
	}
}

func (D *Driver) configureCompression(opt string) {
	switch opt {
	case "gzip":
		D.Log.Debug().Str("compression", "gzip").Msg("compression")
		D.Compression = &compression.GZIP{}
	default:
		D.Compression = nil
	}
}

func (D *Driver) configureBackend(opt string) {
	switch opt {
	case "local":
		D.Log.Debug().Str("backend", "local").Msg("backend")
		D.Backend = &backend.Local{}
	default:
		D.Log.Debug().Str("backend", "local").Msg("backend")
		D.Backend = &backend.Local{}
	}
}

func (D *Driver) obtainMutex(collection string) *sync.Mutex {
	D.Log.Debug().Str("collection", collection).Msg("obtaining mutex")

	D.Mutex.Lock()
	defer D.Mutex.Unlock()

	m, ok := D.Engaged[collection]
	if !ok {
		m = &sync.Mutex{}
		D.Engaged[collection] = m
	}

	return m
}
