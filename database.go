package goblin

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

// Database ...
type Database struct {
	Driver  *Driver
	Options Options
}

// New ...
func New(path string, options *Options) (*Database, error) {
	path = filepath.Clean(path)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Format == "" {
		opts.Format = "json"
	}

	switch opts.Compression {
	case "gzip":
		opts.ext = opts.Format + ".gz"
	default:
		opts.ext = opts.Format
	}

	driver := Driver{
		Path:    path,
		Engaged: make(map[string]*sync.Mutex),
		Log:     logrus.New(),
	}

	driver.configureLogger()
	driver.configureBackend(opts.Backend)
	driver.configureCompression(opts.Compression)
	driver.configureFormatter(opts.Format)

	database := Database{
		Driver:  &driver,
		Options: opts,
	}

	if _, err := os.Stat(path); err == nil {
		database.Driver.Log.WithFields(logrus.Fields{"path": path}).Info("loading previous database")
		return &database, nil
	}

	database.Driver.Log.WithFields(logrus.Fields{"path": path}).Info("creating new database")
	return &database, os.MkdirAll(path, 0755)
}
