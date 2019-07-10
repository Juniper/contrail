package cluster

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/Juniper/contrail/pkg/ansible"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/osutil"
	"github.com/flosch/pongo2"

	shellwords "github.com/mattn/go-shellwords"
)

const (
	createAction = "create"
	updateAction = "update"
	deleteAction = "delete"

	provisionProvisioningAction     = "PROVISION"
	upgradeProvisioningAction       = "UPGRADE"
	importProvisioningAction        = "IMPORT"
	addComputeProvisioningAction    = "ADD_COMPUTE"
	deleteComputeProvisioningAction = "DELETE_COMPUTE"
	addCSNProvisioningAction        = "ADD_CSN"

	enable  = "yes"
	disable = "no"
)

type openstackVariables struct {
	enableHaproxy string
}

type contrailAnsibleDeployer struct {
	deployCluster
	ansibleClient *ansible.CLIClient
}

func (a *contrailAnsibleDeployer) installVenv(installDir string) error {
	a.Log.Info("installing virtualenv for ", installDir)

	args := []string{"install",  "virtualenv==16.4.3"}
	err := osutil.ExecCmdAndWait(a.Reporter, "pip", args, "/usr/bin/")
	if err != nil {
		a.Log.Errorf("Error: pip install virtualenv: %s", err)
		return err
	}
	a.Log.Info("pip install virtualenv successful")

	args = []string{"-m", "virtualenv",
	                defaultAppformixAnsibleRepoDir +
			"appformix-ansible-deployer/" +
			installDir +
			"/venv"}
	err = osutil.ExecCmdAndWait(a.Reporter, "python", args, "/usr/bin")
	if err != nil {
		a.Log.Errorf("Error: virtualenv creation: %s", err)
		return err
	}
	a.Log.Info("virtualenv created")

	args = []string{"install",
		        "-r",
	                defaultAppformixAnsibleRepoDir +
			"appformix-ansible-deployer/" +
			installDir +
			"/requirements.txt"}
	err = osutil.ExecCmdAndWait(a.Reporter,
		    defaultAppformixAnsibleRepoDir + "appformix-ansible-deployer/" + installDir + "/venv/bin/pip",
		    args, "")
	if err != nil {
		a.Log.Errorf("Error: pip install requirements.txt: %s", err)
		return err
	}
	a.Log.Info("pip install requirements.txt successful")

	return nil
}

// nolint: gocyclo
func (a *contrailAnsibleDeployer) untar(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		er := f.Close()
		if er != nil {
			a.Log.Errorf("Error while untar file: %s", er)
		}
	}()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer func() {
		er := gzr.Close()
		if er != nil {
			a.Log.Errorf("Error while untar file: %s", er)
		}
	}()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

			// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			er := f.Close()
			if er != nil {
				a.Log.Errorf("Error while untar file: %s", er)
			}
		}
	}
}

func (a *contrailAnsibleDeployer) getInstanceTemplate() (instanceTemplate string) {
	return filepath.Join(a.getTemplateRoot(), defaultInstanceTemplate)
}

func (a *contrailAnsibleDeployer) getInstanceFile() (instanceFile string) {
	return filepath.Join(a.getWorkingDir(), defaultInstanceFile)
}

func (a *contrailAnsibleDeployer) getVcenterFile() (instanceFile string) {
	return filepath.Join(a.getWorkingDir(), defaultVcenterFile)
}

func (a *contrailAnsibleDeployer) getVcenterTemplate() (instanceTemplate string) {
	return filepath.Join(a.getTemplateRoot(), defaultVcenterTemplate)
}

func (a *contrailAnsibleDeployer) getInventoryTemplate() (inventoryTemplate string) {
	return filepath.Join(a.getTemplateRoot(), defaultInventoryTemplate)
}

func (a *contrailAnsibleDeployer) getInventoryFile() (inventoryFile string) {
	return filepath.Join(a.getWorkingDir(), defaultInventoryFile)
}

func (a *contrailAnsibleDeployer) getAnsibleDeployerRepoDir() (ansibleRepoDir string) {
	return filepath.Join(defaultAnsibleRepoDir, defaultAnsibleRepo)
}

func (a *contrailAnsibleDeployer) getAppformixAnsibleDeployerRepoDir() (ansibleRepoDir string) {
	return filepath.Join(defaultAppformixAnsibleRepoDir, defaultAppformixAnsibleRepo, defaultAppformixDir)
}

func (a *contrailAnsibleDeployer) getXflowDeployerDir() (xflowDir string) {
	return filepath.Join(defaultAppformixAnsibleRepoDir, defaultAppformixAnsibleRepo, defaultXflowDir)
}

func (a *contrailAnsibleDeployer) getAnsibleDatapathEncryptionRepoDir() (ansibleRepoDir string) {
	return filepath.Join(defaultAnsibleRepoDir, defaultAnsibleDatapathEncryptionRepo)
}

func (a *contrailAnsibleDeployer) fetchAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()

	a.Log.Infof("Fetching :%s", a.cluster.config.AnsibleFetchURL)
	args, err := shellwords.Parse(a.cluster.config.AnsibleFetchURL)
	if err != nil {
		return err
	}
	args = append([]string{"fetch"}, args...)
	err = osutil.ExecCmdAndWait(a.Reporter, "git", args, repoDir)
	if err != nil {
		return err
	}
	a.Log.Info("git fetch completed")

	return nil
}

func (a *contrailAnsibleDeployer) cherryPickAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()
	a.Log.Infof("Cherry-picking :%s", a.cluster.config.AnsibleCherryPickRevision)
	args := []string{"cherry-pick", a.cluster.config.AnsibleCherryPickRevision}
	err := osutil.ExecCmdAndWait(a.Reporter, "git", args, repoDir)
	if err != nil {
		return err
	}
	a.Log.Info("Cherry-pick completed")

	return nil
}

func (a *contrailAnsibleDeployer) resetAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()
	a.Log.Infof("Git reset to %s", a.cluster.config.AnsibleRevision)
	args := []string{"reset", "--hard", a.cluster.config.AnsibleRevision}
	err := osutil.ExecCmdAndWait(a.Reporter, "git", args, repoDir)
	if err != nil {
		return err
	}
	a.Log.Info("Git reset completed")

	return nil
}

func (a *contrailAnsibleDeployer) compareInventory() (identical bool, err error) {
	tmpfile, err := ioutil.TempFile("", "instances")
	if err != nil {
		return false, err
	}
	tmpFileName := tmpfile.Name()
	defer func() {
		if err = os.Remove(tmpFileName); err != nil {
			a.Log.Errorf("Error while deleting tmpfile: %s", err)
		}
	}()

	a.Log.Debugf("Creating temporary inventory %s", tmpFileName)
	err = a.createInstancesFile(tmpFileName)
	if err != nil {
		return false, err
	}

	newInventory, err := ioutil.ReadFile(tmpFileName)
	if err != nil {
		return false, err
	}
	oldInventory, err := ioutil.ReadFile(a.getInstanceFile())
	if err != nil {
		return false, err
	}

	return bytes.Equal(oldInventory, newInventory), nil
}

func (a *contrailAnsibleDeployer) createInventory() error {
	if err := a.createInstancesFile(a.getInstanceFile()); err != nil {
		return err
	}
	if a.clusterData.ClusterInfo.Orchestrator == orchestratorVcenter {
		if err := a.createVcenterVarsFile(a.getVcenterFile()); err != nil {
			return err
		}
	}
	if a.clusterData.ClusterInfo.DatapathEncryption {
		return a.createDatapathEncryptionInventory(a.getInventoryFile())
	}
	return nil
}

// nolint: gocyclo
func (a *contrailAnsibleDeployer) getOpenstackDerivedVars() *openstackVariables {
	openstackVars := openstackVariables{}
	cluster := a.clusterData.GetOpenstackClusterInfo()
	// Enable haproxy when multiple openstack control nodes present in cluster
	if (cluster != nil) && (len(cluster.OpenstackControlNodes) > 1) {
		openstackVars.enableHaproxy = enable
		return &openstackVars
	}
	// get CONTROL_NODES from contrail configuration
	openstackControlNodes := []string{}
	if c := a.clusterData.ClusterInfo.GetContrailConfiguration(); c != nil {
		for _, keyValuePair := range c.GetKeyValuePair() {
			if keyValuePair.Key == "OPENSTACK_NODES" {
				openstackControlNodes = strings.Split(keyValuePair.Value, ",")
				break
			}
		}
	}
	// get openstack control node ip when single node is present
	if len(openstackControlNodes) != 0 {
		openstackManagementIP := ""
		for _, node := range a.clusterData.NodesInfo {
			if node.UUID == cluster.OpenstackControlNodes[0].NodeRefs[0].UUID {
				openstackManagementIP = node.IPAddress
				break
			}
		}
		openstackIP := net.ParseIP(openstackManagementIP)
		if openstackIP == nil {
			openstackVars.enableHaproxy = disable
			return &openstackVars
		}
		for _, openstackControlNode := range openstackControlNodes {
			// user error
			// do not enable haproxy if openstack ip is specified
			// as openstack control nodes(OPENSTACK_NODES) as well
			openstackControlNodeIP := net.ParseIP(openstackControlNode)
			if bytes.Equal(openstackIP, openstackControlNodeIP) {
				openstackVars.enableHaproxy = disable
				return &openstackVars
			}
		}
		// enable haproxy if openstack ip is different from
		// openstack control nodes(OPENSTACK_NODES)
		openstackVars.enableHaproxy = enable
		return &openstackVars
	}
	// do not enable haproxy on single openstack control node and
	//single interface setups
	openstackVars.enableHaproxy = disable
	return &openstackVars
}

func (a *contrailAnsibleDeployer) createInstancesFile(destination string) error {
	a.Log.Info("Creating instance.yml input file for ansible deployer")
	context := pongo2.Context{
		"cluster":            a.clusterData.ClusterInfo,
		"openstackCluster":   a.clusterData.GetOpenstackClusterInfo(),
		"k8sCluster":         a.clusterData.GetK8sClusterInfo(),
		"vcenter":            a.clusterData.GetVCenterClusterInfo(),
		"appformixCluster":   a.clusterData.GetAppformixClusterInfo(),
		"xflowCluster":       a.clusterData.GetXflowData(),
		"nodes":              a.clusterData.GetAllNodesInfo(),
		"credentials":        a.clusterData.GetAllCredsInfo(),
		"keypairs":           a.clusterData.GetAllKeypairsInfo(),
		"openstack":          a.getOpenstackDerivedVars(),
		"defaultSSHUser":     a.clusterData.DefaultSSHUser,
		"defaultSSHPassword": a.clusterData.DefaultSSHPassword,
		"defaultSSHKey":      a.clusterData.DefaultSSHKey,
	}
	content, err := template.Apply(a.getInstanceTemplate(), context)
	if err != nil {
		return err
	}

	err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	a.Log.Info("Created instance.yml input file for ansible deployer")
	return nil
}

func (a *contrailAnsibleDeployer) createDatapathEncryptionInventory(destination string) error {
	a.Log.Info("Creating inventory.yml input file for datapath encryption ansible deployer")
	context := pongo2.Context{
		"cluster": a.clusterData.ClusterInfo,
		"nodes":   a.clusterData.GetAllNodesInfo(),
	}
	content, err := template.Apply(a.getInventoryTemplate(), context)
	if err != nil {
		return err
	}
	err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	a.Log.Info("Created inventory.yml input file for datapath encryption ansible deployer")
	return nil
}

func (a *contrailAnsibleDeployer) createVcenterVarsFile(destination string) error {
	a.Log.Info("Creating vcenter_vars.yml input file for vcenter ansible deployer")
	context := pongo2.Context{
		"cluster": a.clusterData.ClusterInfo,
		"vcenter": a.clusterData.GetVCenterClusterInfo(),
		"nodes":   a.clusterData.GetAllNodesInfo(),
	}
	content, err := template.Apply(a.getVcenterTemplate(), context)
	if err != nil {
		return err
	}
	err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	a.Log.Info("Created vcenter_vars.yml input file for vcenter ansible deployer")
	return nil
}

func (a *contrailAnsibleDeployer) play(ansibleArgs []string) error {
	return a.playFromDirectory(a.getAnsibleDeployerRepoDir(), ansibleArgs)
}

func (a *contrailAnsibleDeployer) playFromDirectory(directory string, ansibleArgs []string) error {
	a.Log.WithField("directory", directory).Info("Running playbook")
	return a.ansibleClient.Play(directory, ansibleArgs, "")
}

func (a *contrailAnsibleDeployer) playInstancesProvision(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultInstanceProvPlay)
	return a.play(ansibleArgs)
}

func (a *contrailAnsibleDeployer) playInstancesConfig(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultInstanceConfPlay)
	return a.play(ansibleArgs)
}

func (a *contrailAnsibleDeployer) playOrchestratorProvision(ansibleArgs []string) error {
	switch a.clusterData.ClusterInfo.Orchestrator {
	case orchestratorOpenstack:
		ansibleArgs = append(ansibleArgs, "-e force_checkout=yes")
		switch a.clusterData.ClusterInfo.ProvisioningAction {
		case addComputeProvisioningAction, deleteComputeProvisioningAction:
			ansibleArgs = append(ansibleArgs, "--tags=nova")
		}
		ansibleArgs = append(ansibleArgs, defaultOpenstackProvPlay)
	case orchestratorKubernetes:
		ansibleArgs = append(ansibleArgs, defaultKubernetesProvPlay)
	case orchestratorVcenter:
		ansibleArgs = append(ansibleArgs, defaultvCenterProvPlay)
	}

	return a.play(ansibleArgs)
}

func (a *contrailAnsibleDeployer) playContrailProvision(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultContrailProvPlay)
	return a.play(ansibleArgs)
}

func (a *contrailAnsibleDeployer) playContrailDatapathEncryption() error {
	if a.clusterData.ClusterInfo.DatapathEncryption {
		inventory := filepath.Join(a.getWorkingDir(), "inventory.yml")
		ansibleArgs := []string{"-i", inventory, defaultContrailDatapathEncryptionPlay}
		return a.playFromDirectory(a.getAnsibleDatapathEncryptionRepoDir(), ansibleArgs)
	}
	return nil
}

func (a *contrailAnsibleDeployer) appformixVenvDir() string {
	return filepath.Join(a.getAppformixAnsibleDeployerRepoDir(), "venv")
}

func (a *contrailAnsibleDeployer) xflowVenvDir() string {
	return filepath.Join(a.getXflowDeployerDir(), "venv")
}

func (a *contrailAnsibleDeployer) playAppformixProvision() error {
	if a.clusterData.GetAppformixClusterInfo() != nil {
		err := a.installVenv("appformix")
		if err != nil {
			a.Log.Errorf("Error in installVenv: %s", err)
			return err
		}
		AppformixUsername := a.clusterData.GetAppformixClusterInfo().AppformixUsername
		AppformixPassword := a.clusterData.GetAppformixClusterInfo().AppformixPassword
		if AppformixUsername != "" {
			err := os.Setenv("APPFORMIX_USERNAME", AppformixUsername)
			if err != nil {
				return err
			}
		}
		if AppformixPassword != "" {
			err := os.Setenv("APPFORMIX_PASSWORD", AppformixPassword)
			if err != nil {
				return err
			}
		}
		AppformixVersion := a.clusterData.GetAppformixClusterInfo().AppformixVersion
		ansibleArgs := []string{"-e", "config_file=" + a.getInstanceFile(),
			"-e", "appformix_version=" + AppformixVersion,
			"--skip-tags=install_docker"}
		ansibleArgs = append(ansibleArgs, defaultAppformixProvPlay)

		srcFile := "appformix-" + AppformixVersion + ".tar.gz"
		err = a.untar(defaultAppformixImageDir+srcFile, defaultAppformixImageDir)
		if err != nil {
			a.Log.Errorf("Error while untar file: %s", err)
		}
		repoDir := a.getAppformixAnsibleDeployerRepoDir()
		return a.ansibleClient.Play(repoDir, ansibleArgs, a.appformixVenvDir())
	}
	return nil
}

func (a *contrailAnsibleDeployer) playXflowProvision() error {
	if a.clusterData.GetXflowData() != nil && a.clusterData.GetXflowData().ClusterInfo != nil {
		err := a.installVenv("xflow")
		if err != nil {
			a.Log.Errorf("Error in installVenv: %s", err)
			return err
		}
		venvDir := a.xflowVenvDir()

		xflowDir := a.getXflowDeployerDir()
		if _, err := os.Stat(xflowDir); os.IsNotExist(err) {
			return err
		}

		cmd := "bash"
		cmdArgs := []string{"deploy_xflow.sh", a.getInstanceFile()}
		a.Log.Infof("provisioning xflow: %s %s", cmd, strings.Join(cmdArgs, " "))
		command, err := osutil.VenvCommand(venvDir, cmd, cmdArgs...)
		if err != nil {
			a.Log.Errorf("Error when creating preparing command to run in venv %s: %s", venvDir, err)
			return err
		}

		command.Dir = xflowDir

		if err := osutil.ExecAndWait(a.Reporter, command); err != nil {
			a.Log.Errorf("Error when running command in venv %s: %s", venvDir, err)
			return err
		}
		a.Log.Infof("Finished provisioning xflow")
	}
	return nil
}

// nolint: gocyclo
func (a *contrailAnsibleDeployer) playBook() error {
	args := []string{"-i", "inventory/", "-e",
		"config_file=" + a.getInstanceFile(),
		"-e orchestrator=" + a.clusterData.ClusterInfo.Orchestrator}
	if a.cluster.config.AnsibleSudoPass != "" {
		sudoArg := "-e ansible_sudo_pass=" + a.cluster.config.AnsibleSudoPass
		args = append(args, sudoArg)
	}

	switch a.clusterData.ClusterInfo.ProvisioningAction {
	case provisionProvisioningAction, "":
		if err := a.playInstancesProvision(args); err != nil {
			return err
		}
		if a.clusterData.ClusterInfo.Orchestrator == orchestratorVcenter {
			if err := a.playOrchestratorProvision(args); err != nil {
				return err
			}
			if err := a.playInstancesConfig(args); err != nil {
				return err
			}
		} else {
			if err := a.playInstancesConfig(args); err != nil {
				return err
			}
			if err := a.playOrchestratorProvision(args); err != nil {
				return err
			}
		}
		if err := a.playContrailProvision(args); err != nil {
			return err
		}
		if err := a.playContrailDatapathEncryption(); err != nil {
			return err
		}
		if err := a.playAppformixProvision(); err != nil {
			return err
		}
		if err := a.playXflowProvision(); err != nil {
			return err
		}
	case upgradeProvisioningAction:
		if err := a.playContrailProvision(args); err != nil {
			return err
		}
		if err := a.playContrailDatapathEncryption(); err != nil {
			return err
		}
		if err := a.playAppformixProvision(); err != nil {
			return err
		}
		if err := a.playXflowProvision(); err != nil {
			return err
		}
	case addComputeProvisioningAction:
		if a.clusterData.ClusterInfo.Orchestrator == orchestratorVcenter {
			if err := a.playOrchestratorProvision(args); err != nil {
				return err
			}
			if err := a.playInstancesConfig(args); err != nil {
				return err
			}
		} else {
			if err := a.playInstancesConfig(args); err != nil {
				return err
			}
			if err := a.playOrchestratorProvision(args); err != nil {
				return err
			}
		}
		if err := a.playContrailProvision(args); err != nil {
			return err
		}
		if err := a.playContrailDatapathEncryption(); err != nil {
			return err
		}
		if err := a.playAppformixProvision(); err != nil {
			return err
		}
	case deleteComputeProvisioningAction:
		if err := a.playInstancesConfig(args); err != nil {
			return err
		}
		if err := a.playOrchestratorProvision(args); err != nil {
			return err
		}
		if err := a.playContrailProvision(args); err != nil {
			return err
		}
		if err := a.playAppformixProvision(); err != nil {
			return err
		}
	case addCSNProvisioningAction:
		if err := a.playInstancesConfig(args); err != nil {
			return err
		}
		if err := a.playContrailProvision(args); err != nil {
			return err
		}
		if err := a.playAppformixProvision(); err != nil {
			return err
		}
	}
	return nil
}

// nolint: gocyclo
func (a *contrailAnsibleDeployer) createCluster() error {
	a.Log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.ClusterInfo.FQName)
	status := map[string]interface{}{statusField: statusCreateProgress}
	a.Reporter.ReportStatus(context.Background(), status, defaultResource)

	status[statusField] = statusCreateFailed
	err := a.createWorkingDir()
	if err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	if !a.cluster.config.Test {
		if a.cluster.config.AnsibleFetchURL != "" {
			err = a.fetchAnsibleDeployer()
			if err != nil {
				a.Reporter.ReportStatus(context.Background(), status, defaultResource)
				return err
			}
		}
		if a.cluster.config.AnsibleCherryPickRevision != "" {
			err = a.cherryPickAnsibleDeployer()
			if err != nil {
				a.Reporter.ReportStatus(context.Background(), status, defaultResource)
				return err
			}
		}
		if a.cluster.config.AnsibleRevision != "" {
			err = a.resetAnsibleDeployer()
			if err != nil {
				a.Reporter.ReportStatus(context.Background(), status, defaultResource)
				return err
			}
		}
	}
	err = a.createInventory()
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

func (a *contrailAnsibleDeployer) isUpdated() (updated bool, err error) {
	if a.clusterData.ClusterInfo.ProvisioningState == statusNoState {
		return false, nil
	}
	status := map[string]interface{}{}
	if _, err := os.Stat(a.getInstanceFile()); err == nil {
		ok, err := a.compareInventory()
		if err != nil {
			status[statusField] = statusUpdateFailed
			a.Reporter.ReportStatus(context.Background(), status, defaultResource)
			return false, err
		}
		if ok {
			a.Log.Infof("contrail cluster: %s is already up-to-date", a.clusterData.ClusterInfo.FQName)
			return true, nil
		}
	}
	return false, nil
}

func (a *contrailAnsibleDeployer) updateCluster() error {
	a.Log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.ClusterInfo.FQName)
	status := map[string]interface{}{}
	status[statusField] = statusUpdateProgress
	a.Reporter.ReportStatus(context.Background(), status, defaultResource)

	err := a.createInventory()
	if err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}
	err = a.playBook()
	if err != nil {
		status[statusField] = statusUpdateFailed
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	status[statusField] = statusUpdated
	a.Reporter.ReportStatus(context.Background(), status, defaultResource)
	return nil
}

func (a *contrailAnsibleDeployer) deleteCluster() error {
	a.Log.Infof("Starting %s of contrail cluster: %s",
		a.action, a.cluster.config.ClusterID)
	return a.deleteWorkingDir()
}

func (a *contrailAnsibleDeployer) handleCreate() error {
	if a.isCreated() {
		if err := a.createInventory(); err != nil {
			return err
		}
		return a.updateEndpoints()
	}
	err := a.createCluster()
	if err != nil {
		return err
	}
	return a.createEndpoints()
}

func (a *contrailAnsibleDeployer) handleUpdate() error {
	updated, err := a.isUpdated()
	if err != nil {
		return err
	}
	if updated {
		return nil
	}
	err = a.updateCluster()
	if err != nil {
		return err
	}
	return a.updateEndpoints()
}

func (a *contrailAnsibleDeployer) handleDelete() error {
	err := a.deleteCluster()
	if err != nil {
		return err
	}
	return a.deleteEndpoints()
}

func (a *contrailAnsibleDeployer) Deploy() error {
	switch a.action {
	case createAction:
		return a.handleCreate()
	case updateAction:
		return a.handleUpdate()
	case deleteAction:
		return a.handleDelete()
	}
	return nil
}
