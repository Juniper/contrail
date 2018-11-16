package deploy

import (
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/deploy/cluster"
	"github.com/Juniper/contrail/pkg/deploy/deploy"
	pkglog "github.com/Juniper/contrail/pkg/log"
)

// manager inerface to manage resources
type manager interface {
	manage() error
}

type oneShotManager struct {
	deploy *Command
	log    *logrus.Entry
}

func (o *oneShotManager) manage() error {
	deployer, err := newDeployer(o.deploy)
	if err != nil {
		return err
	}
	return deployer.Deploy()
}

func newOneShotManager(deploy *Command) (*oneShotManager, error) {
	// create logger for oneshot manager
	logger := pkglog.NewLogger("oneshot-manager")
	pkglog.SetLogLevel(logger, deploy.config.LogLevel)

	return &oneShotManager{
		deploy: deploy,
		log:    logger,
	}, nil
}

func newManager(deploy *Command) (manager, error) {
	switch deploy.managerType {
	case "oneshot":
		return newOneShotManager(deploy)
	}
	//TODO(ijohnson) Support daemon manager with etcd
	return nil, errors.New("unsupported manager type")
}

func newDeployer(deployer *Command) (deploy.Deployer, error) {
	switch deployer.config.ResourceType {
	case "contrail_cluster":
		c := &cluster.Config{
			ClusterID:                 deployer.config.ResourceID,
			Action:                    deployer.config.Action,
			TemplateRoot:              deployer.config.TemplateRoot,
			LogLevel:                  deployer.config.LogLevel,
			LogFile:                   deployer.config.LogFile,
			AnsibleSudoPass:           deployer.config.AnsibleSudoPass,
			AnsibleFetchURL:           deployer.config.AnsibleFetchURL,
			AnsibleCherryPickRevision: deployer.config.AnsibleCherryPickRevision,
			AnsibleRevision:           deployer.config.AnsibleRevision,
			Test:                      deployer.config.Test,
		}
		cluster, err := cluster.NewCluster(c)
		if err != nil {
			return nil, err
		}
		return cluster.GetDeployer()
	}
	return nil, errors.New("unsupported resource type")
}
