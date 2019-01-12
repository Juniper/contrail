package deploy

import (
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/deploy/cluster"
	pkglog "github.com/Juniper/contrail/pkg/log"
)

// manager inerface to manage resources
type manager interface {
	manage() error
}

type oneShotManager struct {
	deploy *Deploy
	log    *logrus.Entry
}

func (o *oneShotManager) manage() error {
	deployer, err := newDeployer(o.deploy)
	if err != nil {
		return err
	}
	return deployer.Deploy()
}

func newOneShotManager(deploy *Deploy) (*oneShotManager, error) {
	// create logger for oneshot manager
	logger := pkglog.NewLogger("oneshot-manager")
	pkglog.SetLogLevel(logger, deploy.config.LogLevel)

	return &oneShotManager{
		deploy: deploy,
		log:    logger,
	}, nil
}

func newManager(deploy *Deploy) (manager, error) {
	switch deploy.managerType {
	case "oneshot":
		return newOneShotManager(deploy)
	}
	//TODO(ijohnson) Support daemon manager with etcd
	return nil, errors.New("unsupported manager type")
}

func newDeployer(deploy *Deploy) (base.Deployer, error) {
	switch deploy.config.ResourceType {
	case "contrail_cluster":
		c := &cluster.Config{
			APIServer:                 deploy.APIServer,
			ClusterID:                 deploy.config.ResourceID,
			Action:                    deploy.config.Action,
			TemplateRoot:              deploy.config.TemplateRoot,
			LogLevel:                  deploy.config.LogLevel,
			LogFile:                   deploy.config.LogFile,
			AnsibleSudoPass:           deploy.config.AnsibleSudoPass,
			AnsibleFetchURL:           deploy.config.AnsibleFetchURL,
			AnsibleCherryPickRevision: deploy.config.AnsibleCherryPickRevision,
			AnsibleRevision:           deploy.config.AnsibleRevision,
			Test:                      deploy.config.Test,
		}
		cluster, err := cluster.NewCluster(c)
		if err != nil {
			return nil, err
		}
		return cluster.GetDeployer()
	}
	return nil, errors.New("unsupported resource type")
}
