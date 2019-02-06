package base

import (
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/logutil/report"
)

// Deployer interface
type Deployer interface {
	Deploy() error
}

// Deploy base deploy
type Deploy struct {
	Log      *logrus.Entry
	Reporter *report.Reporter
}
