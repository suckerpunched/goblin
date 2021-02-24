package backend

import (
	"io/ioutil"
	"os"
)

type Local struct{}

func (l *Local) Write(path string, b []byte) error {
	temp := path + ".temp"
	err := ioutil.WriteFile(temp, b, 0644)
	if err != nil {
		return err
	}
	return os.Rename(temp, path)
}

func (l *Local) Read(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}
