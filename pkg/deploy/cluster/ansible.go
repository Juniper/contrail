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
	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/flosch/pongo2"

	shellwords "github.com/mattn/go-shellwords"
	yaml "gopkg.in/yaml.v2"
)

// Ansible related constants.
const (
	CreateAction = "create"
	UpdateAction = "update"
	DeleteAction = "delete"

	ProvisionProvisioningAction     = "PROVISION"
	UpgradeProvisioningAction       = "UPGRADE"
	ImportProvisioningAction        = "IMPORT"
	AddComputeProvisioningAction    = "ADD_COMPUTE"
	DeleteComputeProvisioningAction = "DELETE_COMPUTE"
	AddCSNProvisioningAction        = "ADD_CSN"
	AddCVFMProvisioningAction       = "ADD_CVFM"
	DestroyAction                   = "DESTROY"

	enable  = "yes"
	disable = "no"
)

// Player runs ansible playbook in a container
type Player interface {
	Play(
		ctx context.Context,
		imageRef string,
		imageRefUsername string,
		imageRefPassword string,
		workRoot string,
		ansibleBinaryRepo string,
		ansibleArgs []string,
		keepContainerAlive bool,
	) error
}

type openstackVariables struct {
	enableHaproxy string
}

// ContrailAnsibleDeployer is a deployer using CAD.
type ContrailAnsibleDeployer struct {
	deployCluster
	ansibleClient   *ansible.CLIClient
	containerPlayer Player
}

// nolint: gocyclo
func (a *ContrailAnsibleDeployer) untar(src, dst string) error {
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

func (a *ContrailAnsibleDeployer) getInstanceTemplate() (instanceTemplate string) {
	return filepath.Join(a.getTemplateRoot(), DefaultInstanceTemplate)
}

func (a *ContrailAnsibleDeployer) getInstanceFile() (instanceFile string) {
	return filepath.Join(a.getWorkingDir(), DefaultInstanceFile)
}

func (a *ContrailAnsibleDeployer) getVcenterFile() (instanceFile string) {
	return filepath.Join(a.getWorkingDir(), DefaultVcenterFile)
}

func (a *ContrailAnsibleDeployer) getVcenterTemplate() (instanceTemplate string) {
	return filepath.Join(a.getTemplateRoot(), DefaultVcenterTemplate)
}

func (a *ContrailAnsibleDeployer) getInventoryTemplate() (inventoryTemplate string) {
	return filepath.Join(a.getTemplateRoot(), DefaultInventoryTemplate)
}

func (a *ContrailAnsibleDeployer) getInventoryFile() (inventoryFile string) {
	return filepath.Join(a.getWorkingDir(), DefaultInventoryFile)
}

func (a *ContrailAnsibleDeployer) getAnsibleDeployerRepoDir() (ansibleRepoDir string) {
	return filepath.Join(DefaultAnsibleRepoDir, DefaultAnsibleRepo)
}

func (a *ContrailAnsibleDeployer) getAnsibleDeployerRepoInContainer() string {
	return DefaultAnsibleRepoInContainer
}

func (a *ContrailAnsibleDeployer) getAppformixAnsibleDeployerRepoDir() (ansibleRepoDir string) {
	if a.cluster.config.Test {
		return filepath.Join(a.cluster.config.WorkRoot, defaultAppformixAnsibleRepo, defaultAppformixDir)
	}
	return filepath.Join(defaultAppformixAnsibleRepoDir, defaultAppformixAnsibleRepo, defaultAppformixDir)
}

func (a *ContrailAnsibleDeployer) getXflowDeployerDir() (xflowDir string) {
	return filepath.Join(defaultAppformixAnsibleRepoDir, defaultAppformixAnsibleRepo, defaultXflowDir)
}

func (a *ContrailAnsibleDeployer) getAnsibleDatapathEncryptionRepoDir() (ansibleRepoDir string) {
	return filepath.Join(DefaultAnsibleRepoDir, DefaultAnsibleDatapathEncryptionRepo)
}

func (a *ContrailAnsibleDeployer) fetchAnsibleDeployer() error {
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

func (a *ContrailAnsibleDeployer) cherryPickAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()
	a.Log.Infof("Cherry-picking :%s", a.cluster.config.AnsibleCherryPickRevision)
	args := []string{"cherry-pick", a.cluster.config.AnsibleCherryPickRevision}
	if err := a.cluster.commandExecutor.ExecCmdAndWait(a.Reporter, "git", args, repoDir); err != nil {
		return err
	}
	a.Log.Info("Cherry-pick completed")

	return nil
}

func (a *ContrailAnsibleDeployer) resetAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()
	a.Log.Infof("Git reset to %s", a.cluster.config.AnsibleRevision)
	args := []string{"reset", "--hard", a.cluster.config.AnsibleRevision}
	if err := a.cluster.commandExecutor.ExecCmdAndWait(a.Reporter, "git", args, repoDir); err != nil {
		return err
	}
	a.Log.Info("Git reset completed")

	return nil
}

func (a *ContrailAnsibleDeployer) compareInventory() (identical bool, err error) {
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

// CreateInventory creates "instances" file and "inventory" file.
func (a *ContrailAnsibleDeployer) CreateInventory() error {
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
func (a *ContrailAnsibleDeployer) getOpenstackDerivedVars() *openstackVariables {
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

func (a *ContrailAnsibleDeployer) createInstancesFile(destination string) error {
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
	if err := template.ApplyToFile(a.getInstanceTemplate(), destination, pContext, DefaultFilePermRWOnly); err != nil {
		return err
	}
	a.Log.Info("Created instance.yml input file for ansible deployer")
	return nil
}

func (a *ContrailAnsibleDeployer) createDatapathEncryptionInventory(destination string) error {
	a.Log.Info("Creating inventory.yml input file for datapath encryption ansible deployer")
	pContext := pongo2.Context{
		"cluster": a.clusterData.ClusterInfo,
		"nodes":   a.clusterData.GetAllNodesInfo(),
	}
	if err := template.ApplyToFile(a.getInventoryTemplate(), destination, pContext, DefaultFilePermRWOnly); err != nil {
		return err
	}
	a.Log.Info("Created inventory.yml input file for datapath encryption ansible deployer")
	return nil
}

func (a *ContrailAnsibleDeployer) createVcenterVarsFile(destination string) error {
	a.Log.Info("Creating vcenter_vars.yml input file for vcenter ansible deployer")
	pContext := pongo2.Context{
		"cluster": a.clusterData.ClusterInfo,
		"vcenter": a.clusterData.GetVCenterClusterInfo(),
		"nodes":   a.clusterData.GetAllNodesInfo(),
	}

	if err := template.ApplyToFile(a.getVcenterTemplate(), destination, pContext, DefaultFilePermRWOnly); err != nil {
		return err
	}
	a.Log.Info("Created vcenter_vars.yml input file for vcenter ansible deployer")
	return nil
}

func (a *ContrailAnsibleDeployer) playInContainer(ansibleArgs []string) error {
	a.Log.WithField("directory", a.getAnsibleDeployerRepoInContainer()).Info("Running playbook in container")
	return a.containerPlayer.Play(
		context.Background(),
		a.clusterData.ClusterInfo.ContainerRegistry+
			"/contrail-kolla-ansible-deployer:"+
			a.clusterData.ClusterInfo.ContrailVersion,
		a.clusterData.ClusterInfo.ContainerRegistryUsername,
		a.clusterData.ClusterInfo.ContainerRegistryPassword,
		a.getWorkRoot(),
		a.getAnsibleDeployerRepoInContainer(),
		ansibleArgs,
		false,
	)
}

func (a *ContrailAnsibleDeployer) playFromDirectory(directory string, ansibleArgs []string) error {
	a.Log.WithField("directory", directory).Info("Running playbook")
	return a.ansibleClient.Play(directory, ansibleArgs, "")
}

func (a *ContrailAnsibleDeployer) playInstancesProvision(ansibleArgs []string) error {
	return a.playInContainer(append(ansibleArgs, DefaultInstanceProvPlay))
}

func (a *ContrailAnsibleDeployer) playInstancesConfig(ansibleArgs []string) error {
	return a.playInContainer(append(ansibleArgs, DefaultInstanceConfPlay))
}

func (a *ContrailAnsibleDeployer) playOrchestratorProvision(ansibleArgs []string) error {
	switch a.clusterData.ClusterInfo.Orchestrator {
	case orchestratorOpenstack:
		ansibleArgs = append(ansibleArgs, "-e force_checkout=yes")
		switch a.clusterData.ClusterInfo.ProvisioningAction {
		case AddComputeProvisioningAction, DeleteComputeProvisioningAction:
			ansibleArgs = append(ansibleArgs, "--tags=nova")
		}
		ansibleArgs = append(ansibleArgs, DefaultOpenstackProvPlay)
	case orchestratorKubernetes:
		ansibleArgs = append(ansibleArgs, DefaultKubernetesProvPlay)
	case orchestratorVcenter:
		ansibleArgs = append(ansibleArgs, DefaultvCenterProvPlay)
	}

	return a.playInContainer(ansibleArgs)
}

func (a *ContrailAnsibleDeployer) playOrchestratorDestroy(ansibleArgs []string) error {
	destroyAnsibleArgs := ansibleArgs[:2]
	switch a.clusterData.ClusterInfo.Orchestrator {
	case orchestratorOpenstack:
		destroyAnsibleArgs = append(destroyAnsibleArgs, DefaultOpenstackDestoryPlay)
	}

	return a.playInContainer(destroyAnsibleArgs)
}

func (a *ContrailAnsibleDeployer) playContrailProvision(ansibleArgs []string) error {
	return a.playInContainer(append(ansibleArgs, DefaultContrailProvPlay))
}

func (a *ContrailAnsibleDeployer) playContrailDestroy(ansibleArgs []string) error {
	return a.playInContainer(append(ansibleArgs[:2], DefaultContrailDestoryPlay))
}

// TODO(dji): use ansible.container_deployer in the future
func (a *ContrailAnsibleDeployer) playContrailDatapathEncryption() error {
	if a.clusterData.ClusterInfo.DatapathEncryption {
		inventory := filepath.Join(a.getWorkingDir(), "inventory.yml")
		ansibleArgs := []string{"-i", inventory, DefaultContrailDatapathEncryptionPlay}
		return a.playFromDirectory(a.getAnsibleDatapathEncryptionRepoDir(), ansibleArgs)
	}
	return nil
}

func (a *ContrailAnsibleDeployer) appformixVenvDir() string {
	return filepath.Join(a.getAppformixAnsibleDeployerRepoDir(), "venv")
}

func (a *ContrailAnsibleDeployer) xflowVenvDir() string {
	return filepath.Join(a.getXflowDeployerDir(), "venv")
}

// AppformixConfig is saved defaults
type AppformixConfig struct {
	// AppformixVersion is AppFormix Version in deployment
	AppformixVersion string `yaml:"appformix_version"`
}

// TODO(dji): use ansible.container_deployer in the future
func (a *ContrailAnsibleDeployer) playAppformixProvision() error {
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

// TODO(dji): use ansible.container_deployer in the future
func (a *ContrailAnsibleDeployer) playXflowProvision() error {
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
func (a *ContrailAnsibleDeployer) playBook() error {
	args := []string{"-i", "inventory/", "-e",
		"config_file=" + a.getInstanceFile(),
		"-e orchestrator=" + a.clusterData.ClusterInfo.Orchestrator}
	if a.cluster.config.AnsibleSudoPass != "" {
		sudoArg := "-e ansible_sudo_pass=" + a.cluster.config.AnsibleSudoPass
		args = append(args, sudoArg)
	}

	switch a.clusterData.ClusterInfo.ProvisioningAction {
	case ProvisionProvisioningAction, "", AddCVFMProvisioningAction:
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
	case UpgradeProvisioningAction:
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
	case AddComputeProvisioningAction:
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
	case DeleteComputeProvisioningAction:
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
	case AddCSNProvisioningAction:
		if err := a.playInstancesConfig(args); err != nil {
			return err
		}
		if err := a.playContrailProvision(args); err != nil {
			return err
		}
		if err := a.playAppformixProvision(); err != nil {
			return err
		}
	case DestroyAction:
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
func (a *ContrailAnsibleDeployer) createCluster() error {
	a.Log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.ClusterInfo.FQName)
	status := map[string]interface{}{StatusField: StatusCreateProgress}
	a.Reporter.ReportStatus(context.Background(), status, DefaultResource)

	status[StatusField] = StatusCreateFailed
	if err := a.createWorkingDir(); err != nil {
		a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
		return err
	}

	if !a.cluster.config.Test {
		if a.cluster.config.AnsibleFetchURL != "" {
			if err := a.fetchAnsibleDeployer(); err != nil {
				a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
				return err
			}
		}
		if a.cluster.config.AnsibleCherryPickRevision != "" {
			if err := a.cherryPickAnsibleDeployer(); err != nil {
				a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
				return err
			}
		}
		if a.cluster.config.AnsibleRevision != "" {
			if err := a.resetAnsibleDeployer(); err != nil {
				a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
				return err
			}
		}
	}
	if err := a.CreateInventory(); err != nil {
		a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
		return err
	}
	if err := a.playBook(); err != nil {
		a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
		return err
	}

	status[StatusField] = StatusCreated
	a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
	return nil
}

func (a *ContrailAnsibleDeployer) isUpdated() (updated bool, err error) {
	if a.clusterData.ClusterInfo.ProvisioningState == StatusNoState {
		return false, nil
	}
	status := map[string]interface{}{}
	if _, err := os.Stat(a.getInstanceFile()); err == nil {
		ok, err := a.compareInventory()
		if err != nil {
			status[StatusField] = StatusUpdateFailed
			a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
			return false, err
		}
		if ok {
			a.Log.Infof("contrail cluster: %s is already up-to-date", a.clusterData.ClusterInfo.FQName)
			return true, nil
		}
	}
	return false, nil
}

func (a *ContrailAnsibleDeployer) updateCluster() error {
	a.Log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.ClusterInfo.FQName)
	status := map[string]interface{}{}
	status[StatusField] = StatusUpdateProgress
	a.Reporter.ReportStatus(context.Background(), status, DefaultResource)

	if err := a.CreateInventory(); err != nil {
		a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
		return err
	}
	if err := a.playBook(); err != nil {
		status[StatusField] = StatusUpdateFailed
		a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
		return err
	}

	status[StatusField] = StatusUpdated
	a.Reporter.ReportStatus(context.Background(), status, DefaultResource)
	return nil
}

func (a *ContrailAnsibleDeployer) deleteCluster() error {
	a.Log.Infof("Starting %s of contrail cluster: %s", a.action, a.cluster.config.ClusterID)
	return a.deleteWorkingDir()
}

func (a *ContrailAnsibleDeployer) handleCreate() error {
	if a.isCreated() {
		if err := a.CreateInventory(); err != nil {
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
	times := 2
	if err := retry.Do(func() (retry bool, err error) {
		times--
		return times > 0, a.ensureServiceUserCreated()
	}, retry.WithInterval(proxy.SyncInterval)); err != nil {
		// TODO(mblotniak): Fail instead of logging
		a.Log.Warnf("Skipping service user creation: %v", err)
	}
	return nil
}

func (a *ContrailAnsibleDeployer) handleUpdate() error {
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

func (a *ContrailAnsibleDeployer) handleDelete() error {
	if err := a.deleteCluster(); err != nil {
		return err
	}
	return a.deleteEndpoints()
}

// Deploy handles create/update/delete deployment action.
func (a *ContrailAnsibleDeployer) Deploy() error {
	switch a.action {
	case CreateAction:
		return a.handleCreate()
	case UpdateAction:
		return a.handleUpdate()
	case DeleteAction:
		return a.handleDelete()
	}
	return nil
}
