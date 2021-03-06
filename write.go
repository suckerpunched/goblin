package goblin

import (
	"os"
	"path/filepath"
)

// -------
func (D *Database) Write(collection, resource string, v interface{}) error {
	if err := notEmpty(map[string]string{
		"collection": collection,
		"resource":   resource,
	}); err != nil {
		return err
	}

	mutex := D.Driver.obtainMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(D.Driver.Path, collection)
	path := filepath.Join(dir, resource+"."+D.Options.ext)

	if err := os.MkdirAll(dir, 0755); err != nil {
		D.Driver.Log.Error().Msgf("unable to make directory, %v", dir)
		return err
	}

	D.Driver.Log.Info().Str("path", path).Msgf("write")

	b, _ := D.Driver.Formatter.Encode(v)

	if D.Driver.Compression != nil {
		b, _ = D.Driver.Compression.Compress(b)
	}

	return D.Driver.Backend.Write(path, b)
}
