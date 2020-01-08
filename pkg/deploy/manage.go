package deploy

import (
	"os/exec"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/deploy/cluster"
	"github.com/Juniper/contrail/pkg/deploy/rhospd/undercloud"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// manager inerface to manage resources
type manager interface {
	manage() error
}

type oneShotManager struct {
	deploy *Deploy
	log    *logrus.Entry
}

// TODO(apasyniuk) Export this into asf/osutils
type osCommandExecutor struct{}

func (e *osCommandExecutor) ExecCmdAndWait(
	r *report.Reporter, cmd string, args []string, dir string, envVars ...string,
) error {
	return osutil.ExecCmdAndWait(r, cmd, args, dir, envVars...)
}

func (e *osCommandExecutor) ExecAndWait(r *report.Reporter, cmd *exec.Cmd) error {
	return osutil.ExecAndWait(r, cmd)
}

func (o *oneShotManager) manage() error {
	deployer, err := newDeployer(o.deploy)
	if err != nil {
		return err
	}
	return deployer.Deploy()
}

func newOneShotManager(deploy *Deploy) (*oneShotManager, error) {
	return &oneShotManager{
		deploy: deploy,
		log:    logutil.NewLogger("oneshot-manager"),
	}, nil
}

func newManager(deploy *Deploy) (manager, error) {
	if deploy.managerType == oneShot {
		return newOneShotManager(deploy)
	}
	//TODO(ijohnson) Support daemon manager with etcd
	return nil, errors.New("unsupported manager type")
}

func newDeployer(deploy *Deploy) (base.Deployer, error) {
	switch deploy.config.ResourceType {
	case "contrail_cluster":
		c, err := cluster.NewCluster(&cluster.Config{
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
			ServiceUserID:             deploy.config.ServiceUserID,
			ServiceUserPassword:       deploy.config.ServiceUserPassword,
		}, &osCommandExecutor{})
		if err != nil {
			return nil, err
		}
		return c.GetDeployer()
	case "rhospd_cloud_manager":
		u, err := undercloud.NewUnderCloud(&undercloud.Config{
			APIServer:    deploy.APIServer,
			ResourceID:   deploy.config.ResourceID,
			Action:       deploy.config.Action,
			TemplateRoot: deploy.config.TemplateRoot,
			LogLevel:     deploy.config.LogLevel,
			LogFile:      deploy.config.LogFile,
		})
		if err != nil {
			return nil, err
		}
		return u.GetDeployer()
	}
	return nil, errors.New("unsupported resource type")
}
