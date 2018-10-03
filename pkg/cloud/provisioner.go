package cloud

import (
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

//Provisioner provides an interface to provision cloud
type Provisioner interface {
	modifyTopology() error
	selectCloudCredential() error
	invokeTerraform() error
	reportStatus() error
}

type provisionData struct {
	cloud *Cloud
	log   *logrus.Entry
}

func newProvisioner(cloud *Cloud) (*provisionData, error) {
	// create logger for oneshot manager
	logger := pkglog.NewLogger("cloud-provisioner")
	pkglog.SetLogLevel(logger, cloud.config.LogLevel)

	return &provisionData{
		cloud: cloud,
		log:   logger,
	}, nil
}
