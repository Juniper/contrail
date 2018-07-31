package cluster

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/flosch/pongo2"
	"github.com/mattn/go-shellwords"
	"github.com/siddontang/go/log"
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

func (a *ansibleProvisioner) getAnsibleRepoDir() (ansibleRepoDir string) {
	return filepath.Join(defaultAnsibleRepoDir, defaultAnsibleRepo)
}

func (a *ansibleProvisioner) fetchAnsibleDeployer() error {
	repoDir := a.getAnsibleRepoDir()

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
	repoDir := a.getAnsibleRepoDir()
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
	repoDir := a.getAnsibleRepoDir()
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
			log.Errorf("Error while deleting tmpfile: %s", err)
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
	return a.createInstancesFile(a.getInstanceFile())
}

// nolint: gocyclo
func (a *ansibleProvisioner) getOpenstackDerivedVars() *openstackVariables {
	openstackVars := openstackVariables{}
	cluster := a.clusterData.getOpenstackClusterInfo()
	// Enable haproxy when multiple openstack control nodes present in cluster
	if len(cluster.OpenstackControlNodes) > 1 {
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
	context := pongo2.Context{
		"cluster":          a.clusterData.clusterInfo,
		"openstackCluster": a.clusterData.getOpenstackClusterInfo(),
		"k8sCluster":       a.clusterData.getK8sClusterInfo(),
		"nodes":            a.clusterData.getAllNodesInfo(),
		"openstack":        a.getOpenstackDerivedVars(),
	}
	content, err := a.applyTemplate(a.getInstanceTemplate(), context)
	if err != nil {
		return err
	}

	// strip empty lines in instances yml content
	regex, _ := regexp.Compile("\n[ \r\n\t]*\n")
	contentString := regex.ReplaceAllString(string(content), "\n")
	content = []byte(contentString)
	err = a.writeToFile(destination, content)
	if err != nil {
		return err
	}
	a.log.Info("Created instance.yml input file for ansible deployer")
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
	if a.cluster.config.Test {
		return a.mockPlay(ansibleArgs)
	}
	repoDir := a.getAnsibleRepoDir()
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
	err := a.play(ansibleArgs)
	return err
}

func (a *ansibleProvisioner) playInstancesConfig(ansibleArgs []string) error {
	// play instances configuration playbook
	ansibleArgs = append(ansibleArgs, defaultInstanceConfPlay)
	err := a.play(ansibleArgs)
	return err
}

func (a *ansibleProvisioner) playOrchestratorProvision(ansibleArgs []string) error {
	// play orchestrator provisioning playbook
	switch a.clusterData.clusterInfo.Orchestrator {
	case "openstack":
		ansibleArgs = append(ansibleArgs, "-e force_checkout=yes")
		switch a.clusterData.clusterInfo.ProvisioningAction {
		case "ADD_COMPUTE":
			ansibleArgs = append(ansibleArgs, "--tags nova")
		}
		ansibleArgs = append(ansibleArgs, defaultOpenstackProvPlay)
	case "kubernetes":
		ansibleArgs = append(ansibleArgs, defaultKubernetesProvPlay)
	}
	err := a.play(ansibleArgs)
	return err
}

func (a *ansibleProvisioner) playContrailProvision(ansibleArgs []string) error {
	// play contrail provisioning playbook
	ansibleArgs = append(ansibleArgs, defaultContrailProvPlay)
	err := a.play(ansibleArgs)
	return err
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
	case "UPGRADE":
		if err := a.playContrailProvision(args); err != nil {
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
	case "create":
		if a.isCreated() {
			return nil
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
