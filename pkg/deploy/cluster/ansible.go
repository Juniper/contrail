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

	"github.com/Juniper/asf/pkg/fileutil/template"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/Juniper/asf/pkg/retry"
	"github.com/Juniper/contrail/pkg/ansible"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/flosch/pongo2"

	shellwords "github.com/mattn/go-shellwords"
	yaml "gopkg.in/yaml.v2"
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
	addCVFMProvisioningAction       = "ADD_CVFM"
	destroyAction                   = "DESTROY"

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
	if a.cluster.config.Test {
		return filepath.Join(a.cluster.config.WorkRoot, defaultAppformixAnsibleRepo, defaultAppformixDir)
	}
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
	if err = a.cluster.commandExecutor.ExecCmdAndWait(a.Reporter, "git", args, repoDir); err != nil {
		return err
	}
	a.Log.Info("git fetch completed")

	return nil
}

func (a *contrailAnsibleDeployer) cherryPickAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()
	a.Log.Infof("Cherry-picking :%s", a.cluster.config.AnsibleCherryPickRevision)
	args := []string{"cherry-pick", a.cluster.config.AnsibleCherryPickRevision}
	if err := a.cluster.commandExecutor.ExecCmdAndWait(a.Reporter, "git", args, repoDir); err != nil {
		return err
	}
	a.Log.Info("Cherry-pick completed")

	return nil
}

func (a *contrailAnsibleDeployer) resetAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()
	a.Log.Infof("Git reset to %s", a.cluster.config.AnsibleRevision)
	args := []string{"reset", "--hard", a.cluster.config.AnsibleRevision}
	if err := a.cluster.commandExecutor.ExecCmdAndWait(a.Reporter, "git", args, repoDir); err != nil {
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
	if err = a.createInstancesFile(tmpFileName); err != nil {
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
	if cluster != nil && len(cluster.OpenstackControlNodes) > 1 {
		openstackVars.enableHaproxy = enable
		return &openstackVars
	}
	// get CONTROL_NODES from contrail configuration
	var openstackControlNodes []string
	if c := a.clusterData.ClusterInfo.GetContrailConfiguration(); c != nil {
		if v := c.GetValue("OPENSTACK_NODES"); v != "" {
			openstackControlNodes = strings.Split(v, ",")
		}
	}

	// do not enable haproxy on single openstack control node and
	//single interface setups
	if len(openstackControlNodes) == 0 {
		openstackVars.enableHaproxy = disable
		return &openstackVars
	}

	// get openstack control node ip when single node is present
	openstackManagementIP := ""
	if cluster != nil && len(cluster.OpenstackControlNodes) > 0 && len(cluster.OpenstackControlNodes[0].NodeRefs) > 0 {
		for _, node := range a.clusterData.NodesInfo {
			if node.UUID == cluster.OpenstackControlNodes[0].NodeRefs[0].UUID {
				openstackManagementIP = node.IPAddress
				break
			}
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

func (a *contrailAnsibleDeployer) createInstancesFile(destination string) error {
	a.Log.Info("Creating instance.yml input file for ansible deployer")
	pContext := pongo2.Context{
		"cluster":            a.clusterData.ClusterInfo,
		"openstackCluster":   a.clusterData.GetOpenstackClusterInfo(),
		"k8sCluster":         a.clusterData.GetK8sClusterInfo(),
		"vcenter":            a.clusterData.GetVCenterClusterInfo(),
		"appformixCluster":   a.clusterData.GetAppformixClusterInfo(),
		"monitoredNodes":     a.clusterData.GetAppformixMonitoredNodes(),
		"xflowCluster":       a.clusterData.GetXflowData(),
		"nodes":              a.clusterData.GetAllNodesInfo(),
		"credentials":        a.clusterData.GetAllCredsInfo(),
		"keypairs":           a.clusterData.GetAllKeypairsInfo(),
		"openstack":          a.getOpenstackDerivedVars(),
		"defaultSSHUser":     a.clusterData.DefaultSSHUser,
		"defaultSSHPassword": a.clusterData.DefaultSSHPassword,
		"defaultSSHKey":      a.clusterData.DefaultSSHKey,
	}
	if err := template.ApplyToFile(a.getInstanceTemplate(), destination, pContext, defaultFilePermRWOnly); err != nil {
		return err
	}
	a.Log.Info("Created instance.yml input file for ansible deployer")
	return nil
}

func (a *contrailAnsibleDeployer) createDatapathEncryptionInventory(destination string) error {
	a.Log.Info("Creating inventory.yml input file for datapath encryption ansible deployer")
	pContext := pongo2.Context{
		"cluster": a.clusterData.ClusterInfo,
		"nodes":   a.clusterData.GetAllNodesInfo(),
	}
	if err := template.ApplyToFile(a.getInventoryTemplate(), destination, pContext, defaultFilePermRWOnly); err != nil {
		return err
	}
	a.Log.Info("Created inventory.yml input file for datapath encryption ansible deployer")
	return nil
}

func (a *contrailAnsibleDeployer) createVcenterVarsFile(destination string) error {
	a.Log.Info("Creating vcenter_vars.yml input file for vcenter ansible deployer")
	pContext := pongo2.Context{
		"cluster": a.clusterData.ClusterInfo,
		"vcenter": a.clusterData.GetVCenterClusterInfo(),
		"nodes":   a.clusterData.GetAllNodesInfo(),
	}

	if err := template.ApplyToFile(a.getVcenterTemplate(), destination, pContext, defaultFilePermRWOnly); err != nil {
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

func (a *contrailAnsibleDeployer) playOrchestratorDestroy(ansibleArgs []string) error {
	destroyAnsibleArgs := ansibleArgs[:2]
	switch a.clusterData.ClusterInfo.Orchestrator {
	case orchestratorOpenstack:
		destroyAnsibleArgs = append(destroyAnsibleArgs, defaultOpenstackDestoryPlay)
	}
	return a.play(destroyAnsibleArgs)
}

func (a *contrailAnsibleDeployer) playContrailProvision(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultContrailProvPlay)
	return a.play(ansibleArgs)
}

func (a *contrailAnsibleDeployer) playContrailDestroy(ansibleArgs []string) error {
	destroyAnsibleArgs := ansibleArgs[:2]
	destroyAnsibleArgs = append(destroyAnsibleArgs, defaultContrailDestoryPlay)
	return a.play(destroyAnsibleArgs)
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

// AppformixConfig is saved defaults
type AppformixConfig struct {
	// AppformixVersion is AppFormix Version in deployment
	AppformixVersion string `yaml:"appformix_version"`
}

func (a *contrailAnsibleDeployer) playAppformixProvision() error {
	if a.clusterData.GetAppformixClusterInfo() != nil {
		repoDir := a.getAppformixAnsibleDeployerRepoDir()
		configFile := repoDir + "/" + "config.yml"
		data, ioerr := ioutil.ReadFile(configFile)
		if ioerr != nil {
			return ioerr
		}

		var config AppformixConfig
		if err := yaml.UnmarshalStrict(data, &config); err != nil {
			return ioerr
		}

		AppformixVersion := config.AppformixVersion
		ansibleArgs := []string{"-e", "config_file=" + a.getInstanceFile(),
			"-e", "appformix_version=" + AppformixVersion,
			"--skip-tags=install_docker"}
		ansibleArgs = append(ansibleArgs, defaultAppformixProvPlay)

		srcFile := "appformix-" + AppformixVersion + ".tar.gz"
		if err := a.untar(defaultAppformixImageDir+srcFile, defaultAppformixImageDir); err != nil {
			a.Log.Errorf("Error while untar file: %s", err)
		}
		return a.ansibleClient.Play(repoDir, ansibleArgs, a.appformixVenvDir())
	}
	return nil
}

func (a *contrailAnsibleDeployer) playXflowProvision() error {
	if a.clusterData.GetXflowData() != nil && a.clusterData.GetXflowData().ClusterInfo != nil {
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

		if err := a.cluster.commandExecutor.ExecAndWait(a.Reporter, command); err != nil {
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
	case provisionProvisioningAction, "", addCVFMProvisioningAction:
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
	case destroyAction:
		if err := a.playOrchestratorDestroy(args); err != nil {
			return err
		}
		if err := a.playContrailDestroy(args); err != nil {
			return err
		}
		/*
			if err := a.cleanupResources(args); err != nil {
				return err
			}
		*/
	}
	return nil
}

// nolint: gocyclo
func (a *contrailAnsibleDeployer) createCluster() error {
	a.Log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.ClusterInfo.FQName)
	status := map[string]interface{}{statusField: statusCreateProgress}
	a.Reporter.ReportStatus(context.Background(), status, defaultResource)

	status[statusField] = statusCreateFailed
	if err := a.createWorkingDir(); err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	if !a.cluster.config.Test {
		if a.cluster.config.AnsibleFetchURL != "" {
			if err := a.fetchAnsibleDeployer(); err != nil {
				a.Reporter.ReportStatus(context.Background(), status, defaultResource)
				return err
			}
		}
		if a.cluster.config.AnsibleCherryPickRevision != "" {
			if err := a.cherryPickAnsibleDeployer(); err != nil {
				a.Reporter.ReportStatus(context.Background(), status, defaultResource)
				return err
			}
		}
		if a.cluster.config.AnsibleRevision != "" {
			if err := a.resetAnsibleDeployer(); err != nil {
				a.Reporter.ReportStatus(context.Background(), status, defaultResource)
				return err
			}
		}
	}
	if err := a.createInventory(); err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}
	if err := a.playBook(); err != nil {
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

	if err := a.createInventory(); err != nil {
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}
	if err := a.playBook(); err != nil {
		status[statusField] = statusUpdateFailed
		a.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	status[statusField] = statusUpdated
	a.Reporter.ReportStatus(context.Background(), status, defaultResource)
	return nil
}

func (a *contrailAnsibleDeployer) deleteCluster() error {
	a.Log.Infof("Starting %s of contrail cluster: %s", a.action, a.cluster.config.ClusterID)
	return a.deleteWorkingDir()
}

func (a *contrailAnsibleDeployer) handleCreate() error {
	if a.isCreated() {
		if err := a.createInventory(); err != nil {
			return err
		}
		if err := a.updateEndpoints(); err != nil {
			return err
		}
	} else {
		if err := a.createCluster(); err != nil {
			return err
		}
		if err := a.createEndpoints(); err != nil {
			return err
		}
	}
	// after setting the endpoints we don't know when the proxy will read them, so we retry
	times := 3
	return retry.Do(func() (retry bool, err error) {
		times--
		return times > 0, a.ensureServiceUserCreated()
	}, retry.WithInterval(apisrv.ProxySyncInterval))
}

func (a *contrailAnsibleDeployer) handleUpdate() error {
	updated, err := a.isUpdated()
	if err != nil {
		return err
	}
	if updated {
		return nil
	}
	if err = a.updateCluster(); err != nil {
		return err
	}
	return a.updateEndpoints()
}

func (a *contrailAnsibleDeployer) handleDelete() error {
	if err := a.deleteCluster(); err != nil {
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
