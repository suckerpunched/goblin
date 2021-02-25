package goblin

import "path/filepath"

func (D *Database) Read(collection, resource string, v interface{}) error {
	if err := notEmpty(map[string]string{
		"collection": collection,
		"resource":   resource,
	}); err != nil {
		return err
	}

	dir := filepath.Join(D.Driver.Path, collection)
	path := filepath.Join(dir, resource+"."+D.Options.ext)

	if _, err := stat(path); err != nil {
		D.Driver.Log.Error().Msgf("unable to find file or directory, %v", path)
		return err
	}

	D.Driver.Log.Info().Str("path", path).Msgf("read")

	b, _ := D.Driver.Backend.Read(path)

	if D.Driver.Compression != nil {
		b, _ = D.Driver.Compression.Decompress(b)
	}

	return D.Driver.Formatter.Decode(b, v)
}
