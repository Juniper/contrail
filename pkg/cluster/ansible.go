package cluster

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
	shellwords "github.com/mattn/go-shellwords"
)

const (
	CreateAction = "create"
	UpdateAction = "update"
	DeleteAction = "delete"

	provisionProvisioningAction  = "PROVISION"
	upgradeProvisioningAction    = "UPGRADE"
	importProvisioningAction     = "IMPORT"
	addComputeProvisioningAction = "ADD_COMPUTE"
	addCSNProvisioningAction     = "ADD_CSN"

	enable  = "yes"
	disable = "no"
)

type openstackVariables struct {
	enableHaproxy string
}

type ansibleProvisioner struct {
	provisionCommon
}

// nolint: gocyclo
func (a *ansibleProvisioner) untar(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		er := f.Close()
		if er != nil {
			a.log.Errorf("Error while untar file: %s", er)
		}
	}()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer func() {
		er := gzr.Close()
		if er != nil {
			a.log.Errorf("Error while untar file: %s", er)
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
				a.log.Errorf("Error while untar file: %s", er)
			}
		}
	}
}

func (a *ansibleProvisioner) getInstanceTemplate() (instanceTemplate string) {
	return filepath.Join(a.getTemplateRoot(), defaultInstanceTemplate)
}

func (a *ansibleProvisioner) getInstanceFile() (instanceFile string) {
	return filepath.Join(a.getWorkingDir(), defaultInstanceFile)
}

func (a *ansibleProvisioner) getInventoryTemplate() (inventoryTemplate string) {
	return filepath.Join(a.getTemplateRoot(), defaultInventoryTemplate)
}

func (a *ansibleProvisioner) getInventoryFile() (inventoryFile string) {
	return filepath.Join(a.getWorkingDir(), defaultInventoryFile)
}

func (a *ansibleProvisioner) getAnsibleDeployerRepoDir() (ansibleRepoDir string) {
	return filepath.Join(defaultAnsibleRepoDir, defaultAnsibleRepo)
}

func (a *ansibleProvisioner) getAppformixAnsibleDeployerRepoDir() (ansibleRepoDir string) {
	return filepath.Join(defaultAppformixAnsibleRepoDir, defaultAppformixAnsibleRepo)
}

func (a *ansibleProvisioner) getAnsibleDatapathEncryptionRepoDir() (ansibleRepoDir string) {
	return filepath.Join(defaultAnsibleRepoDir, defaultAnsibleDatapathEncryptionRepo)
}

func (a *ansibleProvisioner) fetchAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()

	a.log.Infof("Fetching :%s", a.cluster.config.AnsibleFetchURL)
	args, err := shellwords.Parse(a.cluster.config.AnsibleFetchURL)
	if err != nil {
		return err
	}
	args = append([]string{"fetch"}, args...)
	err = a.execCmd("git", args, repoDir)
	if err != nil {
		return err
	}
	a.log.Info("git fetch completed")

	return nil
}

func (a *ansibleProvisioner) cherryPickAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()
	a.log.Infof("Cherry-picking :%s", a.cluster.config.AnsibleCherryPickRevision)
	args := []string{"cherry-pick", a.cluster.config.AnsibleCherryPickRevision}
	err := a.execCmd("git", args, repoDir)
	if err != nil {
		return err
	}
	a.log.Info("Cherry-pick completed")

	return nil
}

func (a *ansibleProvisioner) resetAnsibleDeployer() error {
	repoDir := a.getAnsibleDeployerRepoDir()
	a.log.Infof("Git reset to %s", a.cluster.config.AnsibleRevision)
	args := []string{"reset", "--hard", a.cluster.config.AnsibleRevision}
	err := a.execCmd("git", args, repoDir)
	if err != nil {
		return err
	}
	a.log.Info("Git reset completed")

	return nil
}

func (a *ansibleProvisioner) compareInventory() (identical bool, err error) {
	tmpfile, err := ioutil.TempFile("", "instances")
	if err != nil {
		return false, err
	}
	tmpFileName := tmpfile.Name()
	defer func() {
		if err = os.Remove(tmpFileName); err != nil {
			a.log.Errorf("Error while deleting tmpfile: %s", err)
		}
	}()

	a.log.Debugf("Creating temperory inventory %s", tmpFileName)
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

func (a *ansibleProvisioner) createInventory() error {
	if err := a.createInstancesFile(a.getInstanceFile()); err != nil {
		return err
	}
	if a.clusterData.clusterInfo.DatapathEncryption {
		return a.createDatapathEncryptionInventory(a.getInventoryFile())
	}
	return nil
}

// nolint: gocyclo
func (a *ansibleProvisioner) getOpenstackDerivedVars() *openstackVariables {
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

func (a *ansibleProvisioner) createInstancesFile(destination string) error {
	a.log.Info("Creating instance.yml input file for ansible deployer")
	SSHUser, SSHPassword, SSHKey, err := a.cluster.getDefaultCredential()
	if err != nil {
		return err
	}
	context := pongo2.Context{
		"cluster":            a.clusterData.clusterInfo,
		"openstackCluster":   a.clusterData.getOpenstackClusterInfo(),
		"k8sCluster":         a.clusterData.getK8sClusterInfo(),
		"appformixCluster":   a.clusterData.getAppformixClusterInfo(),
		"nodes":              a.clusterData.getAllNodesInfo(),
		"credentials":        a.clusterData.getAllCredsInfo(),
		"keypairs":           a.clusterData.getAllKeypairsInfo(),
		"openstack":          a.getOpenstackDerivedVars(),
		"defaultSSHUser":     SSHUser,
		"defaultSSHPassword": SSHPassword,
		"defaultSSHKey":      SSHKey,
	}
	content, err := a.applyTemplate(a.getInstanceTemplate(), context)
	if err != nil {
		return err
	}

	err = a.writeToFile(destination, content)
	if err != nil {
		return err
	}
	a.log.Info("Created instance.yml input file for ansible deployer")
	return nil
}

func (a *ansibleProvisioner) createDatapathEncryptionInventory(destination string) error {
	a.log.Info("Creating inventory.yml input file for datapath encryption ansible deployer")
	context := pongo2.Context{
		"cluster": a.clusterData.clusterInfo,
		"nodes":   a.clusterData.getAllNodesInfo(),
	}
	content, err := a.applyTemplate(a.getInventoryTemplate(), context)
	if err != nil {
		return err
	}
	err = a.writeToFile(destination, content)
	if err != nil {
		return err
	}
	a.log.Info("Created inventory.yml input file for datapath encryption ansible deployer")
	return nil
}

func (a *ansibleProvisioner) mockPlay(ansibleArgs []string) error {
	playBookIndex := len(ansibleArgs) - 1
	context := pongo2.Context{
		"playBook":    ansibleArgs[playBookIndex],
		"ansibleArgs": strings.Join(ansibleArgs[:playBookIndex], " "),
	}
	content, err := a.applyTemplate("./test_data/test_ansible_playbook.tmpl", context)
	if err != nil {
		return err
	}
	destination := filepath.Join(a.getWorkingDir(), "executed_ansible_playbook.yml")
	err = a.appendToFile(destination, content)
	return err
}

func (a *ansibleProvisioner) play(ansibleArgs []string) error {
	repoDir := a.getAnsibleDeployerRepoDir()
	return a.playFromDir(repoDir, ansibleArgs)
}

func (a *ansibleProvisioner) playFromDir(
	repoDir string, ansibleArgs []string) error {
	if a.cluster.config.Test {
		return a.mockPlay(ansibleArgs)
	}
	cmdline := "ansible-playbook"
	a.log.Infof("Playing playbook: %s %s",
		cmdline, strings.Join(ansibleArgs, " "))
	cmd := exec.Command(cmdline, ansibleArgs...)
	cmd.Dir = repoDir
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}

	// Report progress log periodically to stdout
	go a.reporter.reportLog(stdout)
	go a.reporter.reportLog(stderr)

	if err = cmd.Wait(); err != nil {
		return err
	}
	a.log.Infof("Finished playing playbook: %s %s",
		cmdline, strings.Join(ansibleArgs, " "))

	return nil
}

func (a *ansibleProvisioner) playInstancesProvision(ansibleArgs []string) error {
	// play instances provisioning playbook
	ansibleArgs = append(ansibleArgs, defaultInstanceProvPlay)
	return a.play(ansibleArgs)
}

func (a *ansibleProvisioner) playInstancesConfig(ansibleArgs []string) error {
	// play instances configuration playbook
	ansibleArgs = append(ansibleArgs, defaultInstanceConfPlay)
	return a.play(ansibleArgs)
}

func (a *ansibleProvisioner) playOrchestratorProvision(ansibleArgs []string) error {
	// play orchestrator provisioning playbook
	switch a.clusterData.clusterInfo.Orchestrator {
	case orchestratorOpenstack:
		ansibleArgs = append(ansibleArgs, "-e force_checkout=yes")
		switch a.clusterData.clusterInfo.ProvisioningAction {
		case addComputeProvisioningAction:
			ansibleArgs = append(ansibleArgs, "--tags=nova")
		}
		ansibleArgs = append(ansibleArgs, defaultOpenstackProvPlay)
	case orchestratorKubernetes:
		ansibleArgs = append(ansibleArgs, defaultKubernetesProvPlay)
	}
	return a.play(ansibleArgs)
}

func (a *ansibleProvisioner) playContrailProvision(ansibleArgs []string) error {
	// play contrail provisioning playbook
	ansibleArgs = append(ansibleArgs, defaultContrailProvPlay)
	return a.play(ansibleArgs)
}

func (a *ansibleProvisioner) playContrailDatapathEncryption() error {
	if a.clusterData.clusterInfo.DatapathEncryption {
		inventory := filepath.Join(a.getWorkingDir(), "inventory.yml")
		ansibleArgs := []string{"-i", inventory, defaultContrailDatapathEncryptionPlay}
		return a.playFromDir(a.getAnsibleDatapathEncryptionRepoDir(), ansibleArgs)
	}
	return nil
}

func (a *ansibleProvisioner) playAppformixProvision() error {
	if a.clusterData.getAppformixClusterInfo() != nil {
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
		if a.clusterData.clusterInfo.Orchestrator == orchestratorOpenstack {
			ansibleArgs = append(ansibleArgs,
				"-e @/etc/kolla/external/admin-openrc.yml")
		}
		ansibleArgs = append(ansibleArgs, defaultAppformixProvPlay)

		repoDir := a.getAppformixAnsibleDeployerRepoDir()
		if _, err := os.Stat(repoDir); os.IsNotExist(err) {
			srcPath := a.clusterData.getAppformixClusterInfo().AppformixImageDir
			srcFile := "/appformix-" + AppformixVersion + ".tar.gz"
			er := a.untar(srcPath+srcFile, repoDir)
			if er != nil {
				a.log.Errorf("Error while untar file: %s", er)
			}
		}
		return a.playFromDir(repoDir, ansibleArgs)
	}
	return nil
}

// nolint: gocyclo
func (a *ansibleProvisioner) playBook() error {
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
		if err := a.playInstancesConfig(args); err != nil {
			return err
		}
		if err := a.playOrchestratorProvision(args); err != nil {
			return err
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
		if err := a.playInstancesConfig(args); err != nil {
			return err
		}
		if err := a.playOrchestratorProvision(args); err != nil {
			return err
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
func (a *ansibleProvisioner) createCluster() error {
	a.log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.clusterInfo.FQName)
	status := map[string]interface{}{statusField: statusCreateProgress}
	a.reporter.reportStatus(status)

	status[statusField] = statusCreateFailed
	err := a.createWorkingDir()
	if err != nil {
		a.reporter.reportStatus(status)
		return err
	}

	if !a.cluster.config.Test {
		if a.cluster.config.AnsibleFetchURL != "" {
			err = a.fetchAnsibleDeployer()
			if err != nil {
				a.reporter.reportStatus(status)
				return err
			}
		}
		if a.cluster.config.AnsibleCherryPickRevision != "" {
			err = a.cherryPickAnsibleDeployer()
			if err != nil {
				a.reporter.reportStatus(status)
				return err
			}
		}
		if a.cluster.config.AnsibleRevision != "" {
			err = a.resetAnsibleDeployer()
			if err != nil {
				a.reporter.reportStatus(status)
				return err
			}
		}
	}
	err = a.createInventory()
	if err != nil {
		a.reporter.reportStatus(status)
		return err
	}

	err = a.playBook()
	if err != nil {
		a.reporter.reportStatus(status)
		return err
	}

	status[statusField] = statusCreated
	a.reporter.reportStatus(status)
	return nil
}

func (a *ansibleProvisioner) isUpdated() (updated bool, err error) {
	if a.clusterData.clusterInfo.ProvisioningState == statusNoState {
		return false, nil
	}
	status := map[string]interface{}{}
	if _, err := os.Stat(a.getInstanceFile()); err == nil {
		ok, err := a.compareInventory()
		if err != nil {
			status[statusField] = statusUpdateFailed
			a.reporter.reportStatus(status)
			return false, err
		}
		if ok {
			a.log.Infof("contrail cluster: %s is already up-to-date", a.clusterData.clusterInfo.FQName)
			return true, nil
		}
	}
	return false, nil
}

func (a *ansibleProvisioner) updateCluster() error {
	a.log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.clusterInfo.FQName)
	status := map[string]interface{}{}
	status[statusField] = statusUpdateProgress
	a.reporter.reportStatus(status)

	err := a.createInventory()
	if err != nil {
		a.reporter.reportStatus(status)
		return err
	}
	err = a.playBook()
	if err != nil {
		status[statusField] = statusUpdateFailed
		a.reporter.reportStatus(status)
		return err
	}

	status[statusField] = statusUpdated
	a.reporter.reportStatus(status)
	return nil
}

func (a *ansibleProvisioner) deleteCluster() error {
	a.log.Infof("Starting %s of contrail cluster: %s", a.action, a.cluster.config.ClusterID)
	return a.deleteWorkingDir()
}

func (a *ansibleProvisioner) provision() error {
	switch a.action {
	case CreateAction:
		if a.isCreated() {
			return a.updateEndpoints()
		}
		err := a.createCluster()
		if err != nil {
			return err
		}
		return a.createEndpoints()
	case UpdateAction:
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
	case DeleteAction:
		err := a.deleteCluster()
		if err != nil {
			return err
		}
		return a.deleteEndpoints()
	}
	return nil
}
