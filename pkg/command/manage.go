package command

import (
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/command/cluster"
	"github.com/Juniper/contrail/pkg/command/deploy"
	pkglog "github.com/Juniper/contrail/pkg/log"
)

// manager inerface to manage resources
type manager interface {
	manage() error
}

type oneShotManager struct {
	command *Command
	log     *logrus.Entry
}

func (o *oneShotManager) manage() error {
	deployer, err := newDeployer(o.command)
	if err != nil {
		return err
	}
	return deployer.Deploy()
}

func newOneShotManager(command *Command) (*oneShotManager, error) {
	// create logger for oneshot manager
	logger := pkglog.NewLogger("oneshot-manager")
	pkglog.SetLogLevel(logger, command.config.LogLevel)

	return &oneShotManager{
		command: command,
		log:     logger,
	}, nil
}

func newManager(command *Command) (manager, error) {
	switch command.managerType {
	case "oneshot":
		return newOneShotManager(command)
	}
	//TODO(ijohnson) Support daemon manager with etcd
	return nil, errors.New("unsupported manager type")
}

func newDeployer(commander *Command) (deploy.Deployer, error) {
	switch commander.config.ResourceType {
	case "contrail_cluster":
		c := &cluster.Config{
			ClusterID:                 commander.config.ResourceID,
			Action:                    commander.config.Action,
			TemplateRoot:              commander.config.TemplateRoot,
			LogLevel:                  commander.config.LogLevel,
			LogFile:                   commander.config.LogFile,
			AnsibleSudoPass:           commander.config.AnsibleSudoPass,
			AnsibleFetchURL:           commander.config.AnsibleFetchURL,
			AnsibleCherryPickRevision: commander.config.AnsibleCherryPickRevision,
			AnsibleRevision:           commander.config.AnsibleRevision,
			Test:                      commander.config.Test,
		}
		cluster, err := cluster.NewCluster(c)
		if err != nil {
			return nil, err
		}
		return cluster.GetDeployer()
	}
	return nil, errors.New("unsupported resource type")
}
