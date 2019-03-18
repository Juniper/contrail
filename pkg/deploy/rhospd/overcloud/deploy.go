package overcloud

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/logutil/report"
)

const (
	defaultTemplateRoot = "./pkg/deploy/rhospd/overcloud/templates"
)

type deployOverCloud struct {
	base.Deploy
	overcloud     *OverCloud
	overcloudID   string
	action        string
	overcloudData *base.Data
}

func (p *deployOverCloud) isCreated() bool {
	state := p.overcloudData.ClusterInfo.ProvisioningState
	if p.action == createAction && (state == statusNoState || state == "") {
		return false
	}
	p.Log.Infof("OverCloud %s already deployed, STATE: %s", p.overcloudID, state)
	return true
}

func (p *deployOverCloud) getTemplateRoot() string {
	templateRoot := p.overcloud.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}

func (p *deployOverCloud) getOverCloudHomeDir() string {
	dir := filepath.Join(defaultWorkRoot, p.overcloudID)
	return dir
}

func (p *deployOverCloud) getWorkingDir() string {
	dir := filepath.Join(p.getOverCloudHomeDir())
	return dir
}

func (p *deployOverCloud) createWorkingDir() error {
	return os.MkdirAll(p.getWorkingDir(), os.ModePerm)
}

func (p *deployOverCloud) deleteWorkingDir() error {
	return os.RemoveAll(p.getOverCloudHomeDir())
}

func newContrailCloudDeployer(overcloud *OverCloud, cData *base.Data) (base.Deployer, error) {
	overcloudID := overcloud.config.ResourceID
	// create logger for reporter
	logger := logutil.NewFileLogger("reporter", overcloud.config.LogFile)

	r := report.NewReporter(overcloud.APIServer,
		fmt.Sprintf("%s/%s", defaultResourcePath, overcloudID), logger)

	// create logger for contrail-cloud deployer
	logger = logutil.NewFileLogger("contrail-cloud-deployer", overcloud.config.LogFile)

	return &contrailCloudDeployer{deployOverCloud{
		overcloud:     overcloud,
		overcloudID:   overcloudID,
		action:        overcloud.config.Action,
		overcloudData: cData,
		Deploy: base.Deploy{
			Reporter: r,
			Log:      logger,
		},
	}}, nil
}

func newDeployerByID(overcloud *OverCloud) (base.Deployer, error) {
	var cData *base.Data
	var err error
	if overcloud.config.Action == deleteAction {
		cData = &base.Data{Reader: overcloud.APIServer}
	} else {
		r := base.NewResourceManager(overcloud.APIServer, overcloud.config.LogFile)
		cData, err = r.GetClusterDetails(overcloud.config.ResourceID)
	}
	if err != nil {
		return nil, err
	}
	return newContrailCloudDeployer(overcloud, cData)
}
