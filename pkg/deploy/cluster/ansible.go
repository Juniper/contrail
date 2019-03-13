package cluster

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
	"github.com/mattn/go-shellwords"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/osutil"
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
	return filepath.Join(defaultAppformixAnsibleRepoDir, defaultAppformixAnsibleRepo)
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
	if a.clusterData.clusterInfo.Orchestrator == orchestratorVcenter {
		if err := a.createVcenterVarsFile(a.getVcenterFile()); err != nil {
			return err
		}
	}
	if a.clusterData.clusterInfo.DatapathEncryption {
		return a.createDatapathEncryptionInventory(a.getInventoryFile())
	}
	return nil
}

// nolint: gocyclo
func (a *contrailAnsibleDeployer) getOpenstackDerivedVars() *openstackVariables {
	openstackVars := openstackVariables{}
	cluster := a.clusterData.getOpenstackClusterInfo()
	// Enable haproxy when multiple openstack control nodes present in cluster
	if (cluster != nil) && (len(cluster.OpenstackControlNodes) > 1) {
		openstackVars.enableHaproxy = enable
		return &openstackVars
	}
	// get CONTROL_NODES from contrail configuration
	openstackControlNodes := []string{}
	if c := a.clusterData.clusterInfo.GetContrailConfiguration(); c != nil {
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
		for _, node := range a.clusterData.nodesInfo {
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
	SSHUser, SSHPassword, SSHKey, err := a.cluster.getDefaultCredential()
	if err != nil {
		return err
	}
	context := pongo2.Context{
		"cluster":            a.clusterData.clusterInfo,
		"openstackCluster":   a.clusterData.getOpenstackClusterInfo(),
		"k8sCluster":         a.clusterData.getK8sClusterInfo(),
		"vcenter":            a.clusterData.getVCenterClusterInfo(),
		"appformixCluster":   a.clusterData.getAppformixClusterInfo(),
		"xflowCluster":       a.clusterData.getXflowData(),
		"nodes":              a.clusterData.getAllNodesInfo(),
		"credentials":        a.clusterData.getAllCredsInfo(),
		"keypairs":           a.clusterData.getAllKeypairsInfo(),
		"openstack":          a.getOpenstackDerivedVars(),
		"defaultSSHUser":     SSHUser,
		"defaultSSHPassword": SSHPassword,
		"defaultSSHKey":      SSHKey,
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
		"cluster": a.clusterData.clusterInfo,
		"nodes":   a.clusterData.getAllNodesInfo(),
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
		"cluster": a.clusterData.clusterInfo,
		"vcenter": a.clusterData.getVCenterClusterInfo(),
		"nodes":   a.clusterData.getAllNodesInfo(),
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

func (a *contrailAnsibleDeployer) mockPlay(ansibleArgs []string) error {
	playBookIndex := len(ansibleArgs) - 1
	context := pongo2.Context{
		"playBook":    ansibleArgs[playBookIndex],
		"ansibleArgs": strings.Join(ansibleArgs[:playBookIndex], " "),
	}
	content, err := template.Apply("./test_data/test_ansible_playbook.tmpl", context)
	if err != nil {
		return err
	}
	destination := filepath.Join(a.getWorkingDir(), "executed_ansible_playbook.yml")
	err = fileutil.AppendToFile(destination, content, defaultFilePermRWOnly)
	return err
}

func (a *contrailAnsibleDeployer) play(ansibleArgs []string) error {
	repoDir := a.getAnsibleDeployerRepoDir()
	return a.playFromDir(repoDir, ansibleArgs)
}

func (a *contrailAnsibleDeployer) playFromDir(
	repoDir string, ansibleArgs []string) error {
	return a.playFromDirInVenv(repoDir, ansibleArgs, "")
}

func (a *contrailAnsibleDeployer) playFromDirInVenv(
	repoDir string,
	ansibleArgs []string,
	venvDir string) error {

	if a.cluster.config.Test {
		return a.mockPlay(ansibleArgs)
	}

	var venvLogString string
	if venvDir != "" {
		venvLogString = fmt.Sprintf(" in venv %s", venvDir)
	}

	cmdline := "ansible-playbook"
	a.Log.Infof("Playing playbook: %s %s%s",
		cmdline, strings.Join(ansibleArgs, " "), venvLogString)

	var cmd *exec.Cmd

	if venvDir != "" {
		var err error
		cmd = &exec.Cmd{Path: cmdline, Args: append([]string{cmdline}, ansibleArgs...)}
		cmd.Env = os.Environ()
		cmd, err = osutil.Venv(cmd, venvDir)

		if err != nil {
			return err
		}
	} else {
		cmd = exec.Command(cmdline, ansibleArgs...)
	}

	cmd.Dir = repoDir

	err := osutil.ExecAndWait(a.Reporter, cmd)
	if err != nil {
		a.Log.Errorf("error when running playbook: %s", err)
		return err
	}

	a.Log.Infof("Finished playing playbook: %s %s%s",
		cmdline, strings.Join(ansibleArgs, " "), venvLogString)

	return nil
}

func (a *contrailAnsibleDeployer) playInstancesProvision(ansibleArgs []string) error {
	// play instances provisioning playbook
	ansibleArgs = append(ansibleArgs, defaultInstanceProvPlay)
	a.Log.Warn("NOT PLAYING INSTANCES PROVISION")
	return nil
	// return a.play(ansibleArgs) // todo: uncomment
}

func (a *contrailAnsibleDeployer) playInstancesConfig(ansibleArgs []string) error {
	// play instances configuration playbook
	ansibleArgs = append(ansibleArgs, defaultInstanceConfPlay)
	a.Log.Warn("NOT PLAYING INSTANCES CONFIG")
	return nil
	// return a.play(ansibleArgs) //todo: uncomment
}

func (a *contrailAnsibleDeployer) playOrchestratorProvision(ansibleArgs []string) error {
	// play orchestrator provisioning playbook
	switch a.clusterData.clusterInfo.Orchestrator {
	case orchestratorOpenstack:
		ansibleArgs = append(ansibleArgs, "-e force_checkout=yes")
		switch a.clusterData.clusterInfo.ProvisioningAction {
		case addComputeProvisioningAction, deleteComputeProvisioningAction:
			ansibleArgs = append(ansibleArgs, "--tags=nova")
		}
		ansibleArgs = append(ansibleArgs, defaultOpenstackProvPlay)
	case orchestratorKubernetes:
		ansibleArgs = append(ansibleArgs, defaultKubernetesProvPlay)
	case orchestratorVcenter:
		ansibleArgs = append(ansibleArgs, defaultvCenterProvPlay)
	}
	a.Log.Warn("NOT PLAYING ORCHERSTRATOR PROVISION")
	return nil
	//return a.play(ansibleArgs) //todo: uncommnet
}

func (a *contrailAnsibleDeployer) playContrailProvision(ansibleArgs []string) error {
	// play contrail provisioning playbook
	ansibleArgs = append(ansibleArgs, defaultContrailProvPlay)
	a.Log.Warn("JUST FOR TESTING - NOT PLAYING CONTRAIL PROVISION") // todo: remove
	return nil                                                      //todo: remove
	// return a.play(ansibleArgs) //todo: uncomment
}

func (a *contrailAnsibleDeployer) playContrailDatapathEncryption() error {
	if a.clusterData.clusterInfo.DatapathEncryption {
		inventory := filepath.Join(a.getWorkingDir(), "inventory.yml")
		ansibleArgs := []string{"-i", inventory, defaultContrailDatapathEncryptionPlay}
		return a.playFromDir(a.getAnsibleDatapathEncryptionRepoDir(), ansibleArgs)
	}
	return nil
}

func (a *contrailAnsibleDeployer) apprormixVenvDir() string {
	return filepath.Join(a.getWorkingDir(), "appformix-venv")
}

func (a *contrailAnsibleDeployer) createAppfromixVenv() error {
	err := osutil.ExecCmdAndWait(
		a.Reporter,
		"pip",
		[]string{"install", "virtualenv"},
		"",
	) //todo - make virtualenv program available in docker

	if err != nil {
		return err
	}

	venvDir := a.apprormixVenvDir()

	createVenvCmd := exec.Command("virtualenv", venvDir)
	err = osutil.ExecAndWait(a.Reporter, createVenvCmd)

	if err != nil {
		return err
	}

	// not using full path to pip here nor exec.Command call (which will try to resolve the path for us)
	// so that later call to Venv will find the one inside the virtual env
	// also - the first arg is the name of the program (that's a different convention than with exec.Command call)
	installAnsibleReqsInVenvCmd := &exec.Cmd{Path: "pip", Args: []string{
		"pip",
		"install",
		"ansible==2.4.2.0",
		"requests==2.21.0",
	}}

	installAnsibleReqsInVenvCmd.Env = os.Environ()
	installAnsibleReqsInVenvCmd, err = osutil.Venv(installAnsibleReqsInVenvCmd, venvDir)

	if err != nil {
		return err
	}

	err = osutil.ExecAndWait(a.Reporter, installAnsibleReqsInVenvCmd)

	return err
}

func (a *contrailAnsibleDeployer) playAppformixProvision() error {
	/*	if a.clusterData.getAppformixClusterInfo() != nil {

		err := a.createAppfromixVenv()

		if err != nil {
			return err
		}

		AppformixUsername := a.clusterData.getAppformixClusterInfo().AppformixUsername
		AppformixPassword := a.clusterData.getAppformixClusterInfo().AppformixPassword
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
		AppformixVersion := a.clusterData.getAppformixClusterInfo().AppformixVersion
		ansibleArgs := []string{"-e", "config_file=" + a.getInstanceFile(),
			"-e", "appformix_version=" + AppformixVersion}
		ansibleArgs = append(ansibleArgs, defaultAppformixProvPlay)

		imageDir := a.clusterData.getAppformixClusterInfo().AppformixImageDir
		if _, err := os.Stat(imageDir); os.IsNotExist(err) {
			a.Log.Errorf("imageDir %s does not exist, %s", imageDir, err)
		}
		srcFile := "/appformix-" + AppformixVersion + ".tar.gz"
		err = a.untar(imageDir+srcFile, imageDir)
		if err != nil {
			a.Log.Errorf("Error while untar file: %s", err)
		}
		repoDir := a.getAppformixAnsibleDeployerRepoDir()
		return a.playFromDirInVenv(repoDir, ansibleArgs, a.apprormixVenvDir())
	} */
	a.Log.Warn("NOT PLAYING APPFORMIX PROVISION") //todo: revert

	return nil
}

func (a *contrailAnsibleDeployer) playXflowProvision() error {
	if a.clusterData.getXflowData() != nil && a.clusterData.getXflowData().ClusterInfo != nil {
		xflowDir := a.getXflowDeployerDir()
		if _, err := os.Stat(xflowDir); os.IsNotExist(err) {
			return err
		}

		cmd := "bash"
		cmdArgs := []string{"deploy_xflow.sh", a.getInstanceFile()}
		a.Log.Infof("provisioning xflow: %s %s", cmd, strings.Join(cmdArgs, " "))
		if err := osutil.ExecCmdAndWait(a.Reporter, cmd, cmdArgs, xflowDir); err != nil {
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
		"-e orchestrator=" + a.clusterData.clusterInfo.Orchestrator}
	if a.cluster.config.AnsibleSudoPass != "" {
		sudoArg := "-e ansible_sudo_pass=" + a.cluster.config.AnsibleSudoPass
		args = append(args, sudoArg)
	}

	switch a.clusterData.clusterInfo.ProvisioningAction {
	case provisionProvisioningAction, "":
		if err := a.playInstancesProvision(args); err != nil {
			return err
		}
		if a.clusterData.clusterInfo.Orchestrator == orchestratorVcenter {
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
	case addComputeProvisioningAction:
		if a.clusterData.clusterInfo.Orchestrator == orchestratorVcenter {
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
	a.Log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.clusterInfo.FQName)
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
	if a.clusterData.clusterInfo.ProvisioningState == statusNoState {
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
			a.Log.Infof("contrail cluster: %s is already up-to-date", a.clusterData.clusterInfo.FQName)
			return true, nil
		}
	}
	return false, nil
}

func (a *contrailAnsibleDeployer) updateCluster() error {
	a.Log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.clusterInfo.FQName)
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

func (a *contrailAnsibleDeployer) Deploy() error {
	switch a.action {
	case createAction:
		if a.isCreated() {
			return a.updateEndpoints()
		}
		err := a.createCluster()
		if err != nil {
			return err
		}
		return a.createEndpoints()
	case updateAction:
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
	case deleteAction:
		err := a.deleteCluster()
		if err != nil {
			return err
		}
		return a.deleteEndpoints()
	}
	return nil
}
