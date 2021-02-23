package backend

import (
	"fmt"
	"io/ioutil"
)

type Local struct{}

func (l *Local) Write(path string, b []byte) error {
	err := ioutil.WriteFile(path, b, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (l *Local) Read(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}
