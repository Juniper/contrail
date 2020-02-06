package undercloud

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/fileutil/template"
	"github.com/Juniper/asf/pkg/osutil"
)

const (
	CreateAction = "create"
	UpdateAction = "update"
	DeleteAction = "delete"

	filePermRWOnly = 0600

	ProvisionProvisioningAction = "PROVISION"
	ImportProvisioningAction    = "IMPORT"
	TestIntrospectionDir        = "./test_data/introspection"
)

type contrailCloudDeployer struct {
	deployUnderCloud
}

func (c *contrailCloudDeployer) getInventoryTemplate() (inventoryTemplate string) {
	return filepath.Join(c.getTemplateRoot(), defaultInventoryTemplate)
}

func (c *contrailCloudDeployer) getInventoryFile() (inventoryFile string) {
	return filepath.Join(c.getWorkingDir(), defaultInventoryFile)
}

func (c *contrailCloudDeployer) getSiteTemplate() (siteTemplate string) {
	return filepath.Join(c.getTemplateRoot(), defaultSiteTemplate)
}

func (c *contrailCloudDeployer) getSiteFile() (siteFile string) {
	return filepath.Join(c.getWorkingDir(), defaultSiteFile)
}

func (c *contrailCloudDeployer) getIntrospectionDir() (introspectionDir string) {
	introspectionDir = defaultContrailCloudIntrospectDir
	if c.undercloud.config.Test {
		introspectionDir = TestIntrospectionDir
	}
	return introspectionDir
}

func (c *contrailCloudDeployer) createInputFile(
	pContext pongo2.Context, templateFile, destination, destinationFile string) error {

	content, err := template.Apply(templateFile, pContext)
	if err != nil {
		return err
	}

	err = fileutil.WriteToFile(destination, content, filePermRWOnly)
	if err != nil {
		return err
	}
	c.Log.Infof("Created %s input file for rhospd contrail cloud deployer", destinationFile)

	contrailCloudDir := defaultContrailCloudConfigDir
	if c.undercloud.config.Test {
		contrailCloudDir = c.getWorkingDir() + "/config"
	}
	siteDestination := strings.Join(
		[]string{contrailCloudDir, destinationFile}, "/")
	err = fileutil.WriteToFile(siteDestination, content, filePermRWOnly)
	if err != nil {
		return err
	}
	c.Log.Infof("Copied %s to contrail cloud deployer config directory", destinationFile)
	return nil
}

func (c *contrailCloudDeployer) createSiteFile(destination string) error {
	c.Log.Info("Creating site.yml input file for contrail cloud deployer")
	cloudManager := c.undercloudData.cloudManagerInfo
	context := pongo2.Context{
		"jumphost":     cloudManager.RhospdJumphostNodes[0],
		"cloudManager": cloudManager,
		"undercloud":   cloudManager.RhospdUndercloudNodes[0],
		"overcloud":    cloudManager.RhospdOvercloudNodes[0],
		"networks":     c.undercloudData.overcloudNetworks,
	}
	return c.createInputFile(
		context, c.getSiteTemplate(), destination, defaultSiteFile)
}

func (c *contrailCloudDeployer) createInventoryFile(destination string) error {
	if len(c.undercloudData.overcloudNodes) == 0 {
		c.Log.Info("skip creating inventory.yml for contrail cloud deployer")
		return nil
	}
	c.Log.Info("Creating inventory.yml input file for contrail cloud deployer")
	context := pongo2.Context{
		"nodes": c.undercloudData.overcloudNodes,
	}
	return c.createInputFile(
		context, c.getInventoryTemplate(), destination, defaultInventoryFile)
}
func (c *contrailCloudDeployer) mockExec(cmd string, args []string) error {
	destination := filepath.Join(c.getWorkingDir(), "executed_command.yml")
	if len(args) > 0 {
		cmd = cmd + " " + strings.Join(args, " ")
	}
	err := fileutil.AppendToFile(destination, []byte(cmd+"\n"), filePermRWOnly)
	return err
}

func (c *contrailCloudDeployer) compareInputFile(
	fileType string) (identical bool, err error) {

	tmpfile, err := ioutil.TempFile("", fileType)
	if err != nil {
		return false, err
	}
	tmpFileName := tmpfile.Name()
	defer func() {
		if err = os.Remove(tmpFileName); err != nil {
			c.Log.Errorf("Error while deleting tmpfile: %s", err)
		}
	}()

	c.Log.Debugf("Creating temporary %s %s", fileType, tmpFileName)
	switch fileType {
	case "site":
		err = c.createSiteFile(tmpFileName)
	case "inventory":
		err = c.createInventoryFile(tmpFileName)
	}
	if err != nil {
		return false, err
	}

	newInput, err := ioutil.ReadFile(tmpFileName)
	if err != nil {
		return false, err
	}
	var oldInput []byte
	switch fileType {
	case "site":
		oldInput, err = ioutil.ReadFile(c.getSiteFile())
	case "inventory":
		oldInput, err = ioutil.ReadFile(c.getInventoryFile())
	}
	if err != nil {
		return false, err
	}

	return bytes.Equal(oldInput, newInput), nil
}

func (c *contrailCloudDeployer) compareSite() (identical bool, err error) {
	return c.compareInputFile("site")
}

func (c *contrailCloudDeployer) compareInventory() (identical bool, err error) {
	if len(c.undercloudData.overcloudNodes) == 0 {
		c.Log.Info("skip comparing inventory.yml for contrail cloud deployer")
		return true, nil
	}
	return c.compareInputFile("inventory")
}

func (c *contrailCloudDeployer) execFromDir(cmdline string, args []string) error {
	if c.undercloud.config.Test {
		return c.mockExec(cmdline, args)
	}
	c.Log.Infof("Executing command: %s", cmdline)
	if err := osutil.ExecCmdAndWait(
		c.Reporter, cmdline, args, c.getWorkingDir()); err != nil {
		return err
	}
	c.Log.Infof("Finished executing command: %s", cmdline)

	return nil
}

func (c *contrailCloudDeployer) playBook(args []string) error {
	provisioningAction := ProvisionProvisioningAction
	if c.undercloudData.cloudManagerInfo != nil {
		provisioningAction = c.undercloudData.cloudManagerInfo.ProvisioningAction
	}
	switch provisioningAction {
	case ProvisionProvisioningAction, "":
		if err := c.execFromDir(addKnownHostsCmd, args); err != nil {
			return err
		}
		if err := c.execFromDir(installContrailCloudCmd, args); err != nil {
			return err
		}
		if len(c.undercloudData.overcloudNodes) > 0 {
			if err := c.execFromDir(inventoryAssignCmd, args); err != nil {
				return err
			}
			return c.introspectAndUpdateNodes()
		}
	}
	return nil
}

func (c *contrailCloudDeployer) createUndercloud() error {
	c.Log.Infof("Starting %s of contrail undercloud: %s", c.action,
		c.undercloudData.cloudManagerInfo.RhospdUndercloudNodes[0].FQName)
	status := map[string]interface{}{statusField: statusCreateProgress}
	c.Reporter.ReportStatus(context.Background(), status, defaultResource)

	status[statusField] = statusCreateFailed
	err := c.createWorkingDir()
	if err != nil {
		c.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = c.createInventoryFile(c.getInventoryFile())
	if err != nil {
		c.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = c.createSiteFile(c.getSiteFile())
	if err != nil {
		c.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = c.playBook([]string{})
	if err != nil {
		c.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	status[statusField] = statusCreated
	c.Reporter.ReportStatus(context.Background(), status, defaultResource)
	return nil
}

func (c *contrailCloudDeployer) isUpdated() (updated bool, err error) {
	if c.undercloudData.cloudManagerInfo.ProvisioningState == statusNoState {
		return false, nil
	}
	status := map[string]interface{}{}
	if _, err := os.Stat(c.getSiteFile()); err == nil {
		ok, err := c.compareSite()
		if err != nil {
			status[statusField] = statusUpdateFailed
			c.Reporter.ReportStatus(context.Background(), status, defaultResource)
			return false, err
		}
		if ok {
			c.Log.Infof("contrail undercloud: %s is already up-to-date",
				c.undercloudData.cloudManagerInfo.RhospdUndercloudNodes[0].FQName)
			return true, nil
		}
	}
	return false, nil
}

func (c *contrailCloudDeployer) updateUndercloud() error {
	c.Log.Infof("Starting %s of contrail undercloud: %s", c.action,
		c.undercloudData.cloudManagerInfo.RhospdUndercloudNodes[0].FQName)
	status := map[string]interface{}{}
	status[statusField] = statusUpdateProgress
	c.Reporter.ReportStatus(context.Background(), status, defaultResource)

	status[statusField] = statusUpdateFailed
	err := c.createInventoryFile(c.getInventoryFile())
	if err != nil {
		c.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = c.createSiteFile(c.getSiteFile())
	if err != nil {
		c.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = c.playBook([]string{})
	if err != nil {
		c.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	status[statusField] = statusUpdated
	c.Reporter.ReportStatus(context.Background(), status, defaultResource)
	return nil
}

func (c *contrailCloudDeployer) deleteUndercloud() error {
	c.Log.Infof("Starting %s of contrail undercloud: %s",
		c.action, c.undercloud.config.ResourceID)
	err := c.playBook([]string{"-c"})
	if err != nil {
		return err
	}
	return c.deleteWorkingDir()
}

func (c *contrailCloudDeployer) Deploy() error {
	switch c.action {
	case CreateAction:
		err := c.createUndercloud()
		if err != nil {
			return err
		}
		return nil
	case UpdateAction:
		updated, err := c.isUpdated()
		if err != nil {
			return err
		}
		if updated {
			return nil
		}
		err = c.createUndercloud()
		if err != nil {
			return err
		}
		return nil
	case DeleteAction:
		err := c.deleteUndercloud()
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
