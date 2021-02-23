package goblin

import "os"

func notEmpty(params []string) error {
	for s := range params {
		if params[s] == "" {
			return nil
		}
	}
	return nil
}

func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path)
	}
	return
}
