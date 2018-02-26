package cluster

import (
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/sirupsen/logrus"
)

// Reporter reports provisioing status and log
type Reporter struct {
	api      *apisrv.Client
	resource string
	log      *logrus.Entry
}

func (r *Reporter) reportStatus(status string) {
	//TODO(ijohnson) Implement status update
}

func (r *Reporter) reportLog(logMsg string) error {
	//TODO(ijohnson) Implement status update
	return nil
}
