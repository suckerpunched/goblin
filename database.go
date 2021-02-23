package goblin

import (
	"os"
	"path/filepath"
	"sync"
)

// Database ...
type Database struct {
	Driver  *Driver
	Options Options
}

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
	}

	driver.configureBackend(opts.Backend)
	driver.configureCompression(opts.Compression)
	driver.configureFormatter(opts.Format)

	database := Database{
		Driver:  &driver,
		Options: opts,
	}

	if _, err := os.Stat(path); err == nil {
		return &database, nil
	}

	return &database, os.MkdirAll(path, 0755)
}
