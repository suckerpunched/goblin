package goblin

import (
	"fmt"
	"os"
)

func notEmpty(params map[string]string) error {
	for s := range params {
		if params[s] == "" {
			return fmt.Errorf("unable to process, %s must be provided", s)
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
