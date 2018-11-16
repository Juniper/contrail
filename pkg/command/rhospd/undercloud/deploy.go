package undercloud

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Juniper/contrail/pkg/command/deploy"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/log/report"
)

const (
	defaultTemplateRoot = "./pkg/command/rhospd/undercloud/templates"
)

type deployUnderCloud struct {
	deploy.Deploy
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

func newContrailCloudDeployer(undercloud *UnderCloud, cData *Data) (deploy.Deployer, error) {
	undercloudID := undercloud.config.ResourceID
	// create logger for reporter
	logger := pkglog.NewFileLogger("reporter", undercloud.config.LogFile)
	pkglog.SetLogLevel(logger, undercloud.config.LogLevel)

	r := report.NewReporter(undercloud.APIServer,
		fmt.Sprintf("%s/%s", defaultResourcePath, undercloudID), logger)

	// create logger for contrail-cloud deployer
	logger = pkglog.NewFileLogger("contrail-cloud-deployer", undercloud.config.LogFile)
	pkglog.SetLogLevel(logger, undercloud.config.LogLevel)

	return &contrailCloudDeployer{deployUnderCloud{
		undercloud:     undercloud,
		undercloudID:   undercloudID,
		action:         undercloud.config.Action,
		undercloudData: cData,
		Deploy: deploy.Deploy{
			Reporter: r,
			Log:      logger,
		},
	}}, nil
}

func newDeployerByID(undercloud *UnderCloud) (deploy.Deployer, error) {
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
