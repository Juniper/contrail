package command

import (
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/command/cluster"
	"github.com/Juniper/contrail/pkg/command/provision"
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
	provisioner, err := newProvisioner(o.command)
	if err != nil {
		return err
	}
	return provisioner.Provision()
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

func newProvisioner(commander *Command) (provision.Provisioner, error) {
	switch commander.config.ResourceType {
	case "contrail_cluster":
		c := &cluster.Config{
			ClusterID:                 commander.Config.ResourceID,
			Action:                    commander.Config.Action,
			TemplateRoot:              commander.Config.TemplateRoot,
			LogLevel:                  commander.Config.LogLevel,
			LogFile:                   commander.Config.LogFile,
			AnsibleSudoPass:           commander.Config.AnsibleSudoPass,
			AnsibleFetchURL:           commander.Config.AnsibleFetchURL,
			AnsibleCherryPickRevision: commander.Config.AnsibleCherryPickRevision,
			AnsibleRevision:           commander.Config.AnsibleRevision,
			Test:                      commander.Config.Test,
		}
		cluster := cluster.NewCluster(c)
		return cluster.GetProvisioner()
	}
}
