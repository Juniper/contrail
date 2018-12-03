package cluster

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
	shellwords "github.com/mattn/go-shellwords"
	"github.com/siddontang/go/log"

	"github.com/Juniper/contrail/pkg/common"
)

const (
	enable  = "yes"
	disable = "no"
)

type openstackVariables struct {
	enableHaproxy string
}

type ansibleProvisioner struct {
	provisionCommon
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
	err = common.ExecCmdAndWait(a.reporter, "git", args, repoDir)
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
	err := common.ExecCmdAndWait(a.reporter, "git", args, repoDir)
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
	err := common.ExecCmdAndWait(a.reporter, "git", args, repoDir)
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
			log.Errorf("Error while deleting tmpfile: %s", err)
		}
	}()

	a.log.Debugf("Creating temporary inventory %s", tmpFileName)
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
		"nodes":              a.clusterData.getAllNodesInfo(),
		"credentials":        a.clusterData.getAllCredsInfo(),
		"keypairs":           a.clusterData.getAllKeypairsInfo(),
		"openstack":          a.getOpenstackDerivedVars(),
		"defaultSSHUser":     SSHUser,
		"defaultSSHPassword": SSHPassword,
		"defaultSSHKey":      SSHKey,
	}
	content, err := common.Apply(a.getInstanceTemplate(), context)
	if err != nil {
		return err
	}

	err = common.WriteToFile(destination, content, defaultFilePermRWOnly)
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
	content, err := common.Apply(a.getInventoryTemplate(), context)
	if err != nil {
		return err
	}
	err = common.WriteToFile(destination, content, defaultFilePermRWOnly)
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
	content, err := common.Apply("./test_data/test_ansible_playbook.tmpl", context)
	if err != nil {
		return err
	}
	destination := filepath.Join(a.getWorkingDir(), "executed_ansible_playbook.yml")
	err = common.AppendToFile(destination, content, defaultFilePermRWOnly)
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

	err := common.ExecCmdAndWait(a.reporter, cmdline, ansibleArgs, repoDir)
	if err != nil {
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
	case "openstack":
		ansibleArgs = append(ansibleArgs, "-e force_checkout=yes")
		switch a.clusterData.clusterInfo.ProvisioningAction {
		case "ADD_COMPUTE":
			ansibleArgs = append(ansibleArgs, "--tags=nova")
		}
		ansibleArgs = append(ansibleArgs, defaultOpenstackProvPlay)
	case "kubernetes":
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
	case "PROVISION", "":
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
	case "UPGRADE":
		if err := a.playContrailProvision(args); err != nil {
			return err
		}
		if err := a.playContrailDatapathEncryption(); err != nil {
			return err
		}
	case "ADD_COMPUTE":
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
	case "ADD_CSN":
		if err := a.playInstancesConfig(args); err != nil {
			return err
		}
		if err := a.playContrailProvision(args); err != nil {
			return err
		}
	}
	return nil
}

// nolint: gocyclo
func (a *ansibleProvisioner) createCluster() error {
	a.log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.clusterInfo.FQName)
	status := map[string]interface{}{statusField: statusCreateProgress}
	a.reporter.ReportStatus(status, defaultResource)

	status[statusField] = statusCreateFailed
	err := a.createWorkingDir()
	if err != nil {
		a.reporter.ReportStatus(status, defaultResource)
		return err
	}

	if !a.cluster.config.Test {
		if a.cluster.config.AnsibleFetchURL != "" {
			err = a.fetchAnsibleDeployer()
			if err != nil {
				a.reporter.ReportStatus(status, defaultResource)
				return err
			}
		}
		if a.cluster.config.AnsibleCherryPickRevision != "" {
			err = a.cherryPickAnsibleDeployer()
			if err != nil {
				a.reporter.ReportStatus(status, defaultResource)
				return err
			}
		}
		if a.cluster.config.AnsibleRevision != "" {
			err = a.resetAnsibleDeployer()
			if err != nil {
				a.reporter.ReportStatus(status, defaultResource)
				return err
			}
		}
	}
	err = a.createInventory()
	if err != nil {
		a.reporter.ReportStatus(status, defaultResource)
		return err
	}

	err = a.playBook()
	if err != nil {
		a.reporter.ReportStatus(status, defaultResource)
		return err
	}

	status[statusField] = statusCreated
	a.reporter.ReportStatus(status, defaultResource)
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
			a.reporter.ReportStatus(status, defaultResource)
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
	a.reporter.ReportStatus(status, defaultResource)

	err := a.createInventory()
	if err != nil {
		a.reporter.ReportStatus(status, defaultResource)
		return err
	}
	err = a.playBook()
	if err != nil {
		status[statusField] = statusUpdateFailed
		a.reporter.ReportStatus(status, defaultResource)
		return err
	}

	status[statusField] = statusUpdated
	a.reporter.ReportStatus(status, defaultResource)
	return nil
}

func (a *ansibleProvisioner) deleteCluster() error {
	a.log.Infof("Starting %s of contrail cluster: %s", a.action, a.cluster.config.ClusterID)
	return a.deleteWorkingDir()
}

func (a *ansibleProvisioner) provision() error {
	switch a.action {
	case "create":
		if a.isCreated() {
			return a.updateEndpoints()
		}
		err := a.createCluster()
		if err != nil {
			return err
		}
		return a.createEndpoints()
	case "update":
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
	case "delete":
		err := a.deleteCluster()
		if err != nil {
			return err
		}
		return a.deleteEndpoints()
	}
	return nil
}
