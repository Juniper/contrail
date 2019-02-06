package undercloud

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/logutil/report"
)

const (
	defaultTemplateRoot = "./pkg/deploy/rhospd/undercloud/templates"
)

type deployUnderCloud struct {
	base.Deploy
	undercloud     *UnderCloud
	undercloudID   string
	action         string
	undercloudData *Data
}

func (p *deployUnderCloud) isCreated() bool {
	state := p.undercloudData.cloudManagerInfo.ProvisioningState
	if p.action == createAction && (state == statusNoState || state == "") {
		return false
	}
	p.Log.Infof("UnderCloud %s already deployed, STATE: %s", p.undercloudID, state)
	return true
}

func (p *deployUnderCloud) getTemplateRoot() string {
	templateRoot := p.undercloud.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}

func (p *deployUnderCloud) getUnderCloudHomeDir() string {
	dir := filepath.Join(defaultWorkRoot, p.undercloudID)
	return dir
}

func (p *deployUnderCloud) getWorkingDir() string {
	dir := filepath.Join(p.getUnderCloudHomeDir())
	return dir
}

func (p *deployUnderCloud) createWorkingDir() error {
	return os.MkdirAll(p.getWorkingDir(), os.ModePerm)
}

func (p *deployUnderCloud) deleteWorkingDir() error {
	return os.RemoveAll(p.getUnderCloudHomeDir())
}

func newContrailCloudDeployer(undercloud *UnderCloud, cData *Data) (base.Deployer, error) {
	return &contrailCloudDeployer{deployUnderCloud{
		undercloud:     undercloud,
		undercloudID:   undercloud.config.ResourceID,
		action:         undercloud.config.Action,
		undercloudData: cData,
		Deploy: base.Deploy{
			Reporter: report.NewReporter(
				undercloud.APIServer,
				fmt.Sprintf("%s/%s", defaultResourcePath, undercloud.config.ResourceID),
				logutil.NewFileLogger("reporter", undercloud.config.LogFile),
			),
			Log: logutil.NewFileLogger("contrail-cloud-deployer", undercloud.config.LogFile),
		},
	}}, nil
}

func newDeployerByID(undercloud *UnderCloud) (base.Deployer, error) {
	var cData *Data
	var err error
	if undercloud.config.Action == deleteAction {
		cData = &Data{}
	} else {
		cData = NewData(undercloud.APIServer)
		err = cData.getCloudManagerDetails(undercloud.config.ResourceID)
	}
	if err != nil {
		return nil, err
	}
	return newContrailCloudDeployer(undercloud, cData)
}
