package agent

import (
	"os"

	"github.com/pkg/errors"
)

func directoryHandler(action, dir string) error {
	switch action {
	case "create":
		err := os.MkdirAll(dir, 0744)
		if err != nil {
			return errors.Wrapf(err, "creation of %s directory failed", dir)
		}
	case "delete":
		err := os.RemoveAll(dir)
		if err != nil {
			return errors.Wrapf(err, "deletion of %s directory failed", dir)
		}
	}

	return nil
}
