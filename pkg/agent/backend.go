package agent

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type backend interface {
	write(string, []byte) error
	remove(string) error
}

type fileBackend struct {
}

func newBackend(t string) (backend, error) {
	switch t {
	case "file":
		return &fileBackend{}, nil
	}
	return nil, errors.New("unsupported backend type")
}

func (b *fileBackend) write(path string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, content, os.FileMode(0600))
	return err
}

func (b *fileBackend) remove(path string) error {
	return os.Remove(path)
}
