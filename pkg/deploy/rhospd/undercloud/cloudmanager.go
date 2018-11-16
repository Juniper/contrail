package undercloud

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/osutil"
)

const (
	createAction   = "create"
	updateAction   = "update"
	deleteAction   = "delete"
	filePermRWOnly = 0600

	provisionProvisioningAction = "PROVISION"
	importProvisioningAction    = "IMPORT"
)

type contrailCloudDeployer struct {
	deployUnderCloud
}

func (a *contrailCloudDeployer) getSiteTemplate() (siteTemplate string) {
	return filepath.Join(a.getTemplateRoot(), defaultSiteTemplate)
}

func (a *contrailCloudDeployer) getSiteFile() (siteFile string) {
	return filepath.Join(a.getWorkingDir(), defaultSiteFile)
}

func (a *contrailCloudDeployer) createSiteFile(destination string) error {
	a.Log.Info("Creating site.yml input file for contrail cloud deployer")
	cloudManager := a.undercloudData.cloudManagerInfo
	context := pongo2.Context{
		"jumphost":     cloudManager.RhospdJumphostNodes[0],
		"cloudManager": cloudManager,
		"undercloud":   cloudManager.RhospdUndercloudNodes[0],
		"overcloud":    cloudManager.RhospdOverclouds[0],
		"networks":     a.undercloudData.overcloudNetworks,
	}
	content, err := template.Apply(a.getSiteTemplate(), context)
	if err != nil {
		return err
	}

	err = fileutil.WriteToFile(destination, content, filePermRWOnly)
	if err != nil {
		return err
	}
	a.Log.Info("Created instance.yml input file for ansible deployer")
	return nil
}

func (a *contrailCloudDeployer) mockExec(cmd string) error {
	destination := filepath.Join(a.getWorkingDir(), "executed_command.yml")
	err := fileutil.AppendToFile(destination, []byte(cmd), filePermRWOnly)
	return err
}

func (a *contrailCloudDeployer) compareSite() (identical bool, err error) {
	tmpfile, err := ioutil.TempFile("", "site")
	if err != nil {
		return false, err
	}
	tmpFileName := tmpfile.Name()
	defer func() {
		if err = os.Remove(tmpFileName); err != nil {
			a.Log.Errorf("Error while deleting tmpfile: %s", err)
		}
	}()

	a.Log.Debugf("Creating temperory site %s", tmpFileName)
	err = a.createSiteFile(tmpFileName)
	if err != nil {
		return false, err
	}

	newSite, err := ioutil.ReadFile(tmpFileName)
	if err != nil {
		return false, err
	}
	oldSite, err := ioutil.ReadFile(a.getSiteFile())
	if err != nil {
		return false, err
	}

	return bytes.Equal(oldSite, newSite), nil
}

func (a *contrailCloudDeployer) execFromDir(cmdline string) error {
	if a.undercloud.config.Test {
		return a.mockExec(cmdline)
	}
	a.Log.Infof("Executintg command: %s", cmdline)
	if err := osutil.ExecCmdAndWait(
		a.Reporter, cmdline, []string{}, defaultWorkingDir); err != nil {
		return err
	}
	a.Log.Infof("Finished executing command: %s", cmdline)

	return nil
}

func (a *contrailCloudDeployer) playBook() error {
	switch a.undercloudData.cloudManagerInfo.ProvisioningAction {
	case provisionProvisioningAction, "":
		return a.execFromDir(installContrailCloudCmd + "\n")
	}
	return nil
}

func (a *contrailCloudDeployer) createUndercloud() error {
	a.Log.Infof("Starting %s of contrail undercloud: %s", a.action,
		a.undercloudData.cloudManagerInfo.RhospdUndercloudNodes[0].FQName)
	status := map[string]interface{}{statusField: statusCreateProgress}
	a.Reporter.ReportStatus(context.Background(), status, defaultResource)

	status[statusField] = statusCreateFailed
	err := a.createWorkingDir()
	if err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = a.createSiteFile(a.getSiteFile())
	if err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = a.playBook()
	if err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	status[statusField] = statusCreated
	a.Reporter.ReportStatus(context.Background(), status, defaultResource)
	return nil
}

func (a *contrailCloudDeployer) isUpdated() (updated bool, err error) {
	if a.undercloudData.cloudManagerInfo.ProvisioningState == statusNoState {
		return false, nil
	}
	status := map[string]interface{}{}
	if _, err := os.Stat(a.getSiteFile()); err == nil {
		ok, err := a.compareSite()
		if err != nil {
			status[statusField] = statusUpdateFailed
			a.Reporter.ReportStatus(context.Background(), status, defaultResource)
			return false, err
		}
		if ok {
			a.Log.Infof("contrail undercloud: %s is already up-to-date",
				a.undercloudData.cloudManagerInfo.RhospdUndercloudNodes[0].FQName)
			return true, nil
		}
	}
	return false, nil
}

func (a *contrailCloudDeployer) updateUndercloud() error {
	a.Log.Infof("Starting %s of contrail undercloud: %s", a.action,
		a.undercloudData.cloudManagerInfo.RhospdUndercloudNodes[0].FQName)
	status := map[string]interface{}{}
	status[statusField] = statusUpdateProgress
	a.Reporter.ReportStatus(context.Background(), status, defaultResource)

	status[statusField] = statusUpdateFailed
	err := a.createSiteFile(a.getSiteFile())
	if err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = a.playBook()
	if err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	status[statusField] = statusUpdated
	a.Reporter.ReportStatus(context.Background(), status, defaultResource)
	return nil
}

func (a *contrailCloudDeployer) deleteUndercloud() error {
	a.Log.Infof("Starting %s of contrail undercloud: %s",
		a.action, a.undercloud.config.ResourceID)
	return a.deleteWorkingDir()
}

func (a *contrailCloudDeployer) Deploy() error {
	switch a.action {
	case createAction:
		err := a.createUndercloud()
		if err != nil {
			return err
		}
		return nil
	case updateAction:
		updated, err := a.isUpdated()
		if err != nil {
			return err
		}
		if updated {
			return nil
		}
		err = a.createUndercloud()
		if err != nil {
			return err
		}
		return nil
	case deleteAction:
		err := a.deleteUndercloud()
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
