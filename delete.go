package goblin

import (
	"fmt"
	"os"
	"path/filepath"
)

// Delete ...
func (D *Database) Delete(collection, resource string) error {
	if err := notEmpty(map[string]string{
		"collection": collection,
	}); err != nil {
		return err
	}

	mutex := D.Driver.obtainMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	path := filepath.Join(D.Driver.Path, collection)
	if resource != "" {
		path = filepath.Join(path, resource+"."+D.Options.ext)
	}

	switch fi, err := stat(path); {
	case fi == nil, err != nil:
		D.Driver.Log.Error().Msgf("unable to find file or directory, %v", path)
		return fmt.Errorf("unable to find file or directory, %v", path)

	case fi.Mode().IsDir():
		D.Driver.Log.Info().Str("collection", path).Msg("deleting collection")
		return os.RemoveAll(path)

	case fi.Mode().IsRegular():
		D.Driver.Log.Info().Str("resource", path).Msg("deleting resource")
		return os.RemoveAll(path)
	}

	D.Driver.Log.Error().Msg("unable to process delete")
	return fmt.Errorf("unable to process delete")
}
