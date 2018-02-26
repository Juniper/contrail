package cluster

import (
	"errors"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

type Reporter struct {
	api      *apisrv.Client
	resource string
	log      *logrus.Entry
}

func (r *Reporter) reportStatus(status string) error {
	//TODO(ijohnson) Implement status update
	return nil
}

func (r *Reporter) reportLog(logMsg string) error {
	//TODO(ijohnson) Implement status update
	return nil
}
