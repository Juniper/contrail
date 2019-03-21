package cluster

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/osutil"
	"github.com/Juniper/contrail/pkg/retry"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	defaultContrailCommonTemplate  = "contrail_common.tmpl"
	defaultGatewayCommonTemplate   = "gateway_common.tmpl"
	defaultTORCommonTemplate       = "tor_common.tmpl"
	defaultOpenstackConfigTemplate = "openstack_config.tmpl"
	defaultAuthRegistryTemplate    = "authorized_registries.tmpl"

	mcWorkDir                 = "multi-cloud"
	mcAnsibleRepo             = "contrail-multi-cloud"
	defaultSSHKeyRepo         = "keypair"
	defaultContrailCommonFile = "ansible/contrail/common.yml"
	defaultTORCommonFile      = "ansible/tor/common.yml"
	defaultGatewayCommonFile  = "ansible/gateway/common.yml"
	defaultMCInventoryFile    = "inventories/inventory.yml"
	defaultTopologyFile       = "topology.yml"
	defaultSecretFile         = "secret.yml"
	defaultSSHAgentFile       = "ssh-agent-config.yml"

	defaultContrailUser       = "admin"
	defaultContrailPassword   = "c0ntrail123"
	defaultContrailConfigPort = "8082"
	defaultContrailTenant     = "default-project"
	pathConfig                = "/etc/multicloud"
	bgpSecret                 = "bgp_secret"
	debugMCRoutes             = "False"
	torBGPSecret              = "contrail_secret"
	torOSPFSecret             = "contrail_secret"

	mcState = "state.yml"

	defaultMCGWDeployPlay        = "ansible/gateway/playbooks/deploy_and_run_all.yml"
	defaultMCInstanceConfPlay    = "ansible/contrail/playbooks/configure.yml"
	defaultMCKubernetesProvPlay  = "ansible/contrail/playbooks/orchestrator.yml"
	defaultMCDeployContrail      = "ansible/contrail/playbooks/deploy.yml"
	defaultMCSetupContrailRoutes = "ansible/contrail/playbooks/add_tunnel_routes.yml"
	defaultMCFixComputeDNS       = "ansible/contrail/playbooks/fix_compute_dns.yml"
	defaultMCContrailCleanup     = "ansible/contrail/playbooks/cleanup.yml"
	defaultMCGatewayCleanup      = "ansible/gateway/playbooks/cleanup.yml"
	testTemplate                 = "./../../cloud/test_data/test_cmd.tmpl"

	openstack            = "openstack"
	setupRoutesPlayRetry = 3

	addCloud    = "ADD_CLOUD"
	updateCloud = "UPDATE_CLOUD"
	deleteCloud = "DELETE_CLOUD"

	aws    = "aws"
	azure  = "azure"
	onPrem = "private"

	statusRetryInterval = 3 * time.Second
)

type multiCloudProvisioner struct {
	contrailAnsibleDeployer
	workDir string
}

// SSHAgentConfig related to ssh-agent process
type SSHAgentConfig struct {
	AuthSock string `yaml:"auth_sock"`
	PID      string `yaml:"pid"`
}

// PubKeyConfig to read secret file
type PubKeyConfig struct {
	Info         map[string]string   `yaml:"public_key"`
	AwsAccessKey string              `yaml:"aws_access_key"`
	AwsSecretKey string              `yaml:"aws_secret_key"`
	AuthReg      []map[string]string `yaml:"authorized_registries"`
}

func (m *multiCloudProvisioner) Deploy() error {

	m.updateMCWorkDir()
	switch m.clusterData.ClusterInfo.ProvisioningAction {
	case addCloud:
		err := m.verifyCloudStatus()
		if err != nil {
			return err
		}
		err = m.createMCCluster()
		if err != nil {
			return err
		}
		return m.createEndpoints()
	case updateCloud:
		err := m.verifyCloudStatus()
		if err != nil {
			return err
		}
		updated, err := m.isMCUpdated()
		if err != nil {
			return err
		}
		if updated {
			return nil
		}
		return m.updateMCCluster()
	case deleteCloud:
		return m.deleteMCCluster()
	}
	return nil
}

// nolint: gocyclo
func (m *multiCloudProvisioner) createMCCluster() error {

	m.Log.Infof("Starting %s of contrail cluster: %s",
		m.clusterData.ClusterInfo.ProvisioningAction,
		m.clusterData.ClusterInfo.FQName)

	status := map[string]interface{}{statusField: statusUpdateProgress}
	if m.action == createAction {
		status = map[string]interface{}{statusField: statusCreateProgress}
	}
	m.Reporter.ReportStatus(context.Background(), status, defaultResource)

	status[statusField] = statusUpdateFailed
	if m.action == createAction {
		status[statusField] = statusCreateFailed
	}

	if !m.cluster.config.Test {
		if m.cluster.config.AnsibleFetchURL != "" {
			err := m.fetchAnsibleDeployer()
			if err != nil {
				m.Reporter.ReportStatus(context.Background(), status, defaultResource)
				return err
			}
		}
		if m.cluster.config.AnsibleCherryPickRevision != "" {
			err := m.cherryPickAnsibleDeployer()
			if err != nil {
				m.Reporter.ReportStatus(context.Background(), status, defaultResource)
				return err
			}
		}
		if m.cluster.config.AnsibleRevision != "" {
			err := m.resetAnsibleDeployer()
			if err != nil {
				m.Reporter.ReportStatus(context.Background(), status, defaultResource)
				return err
			}
		}
	}

	err := m.createFiles(m.workDir)
	if err != nil {
		m.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = m.manageSSHAgent(m.workDir, createAction)
	if err != nil {
		m.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = m.runGenerateInventory(m.workDir)
	if err != nil {
		m.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = m.mcPlayBook()
	if err != nil {
		m.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}
	status[statusField] = statusUpdated
	if m.action == createAction {
		status[statusField] = statusCreated
	}
	m.Reporter.ReportStatus(context.Background(), status, defaultResource)
	return nil
}

func (m *multiCloudProvisioner) updateMCCluster() error {
	m.Log.Infof("Starting %s of contrail cluster: %s", m.action, m.clusterData.ClusterInfo.FQName)

	status := map[string]interface{}{}
	status[statusField] = statusUpdateProgress
	m.Reporter.ReportStatus(context.Background(), status, defaultResource)
	status[statusField] = statusUpdateFailed

	err := m.createFiles(m.workDir)
	if err != nil {
		m.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = m.manageSSHAgent(m.workDir, updateAction)
	if err != nil {
		m.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = m.runGenerateInventory(m.workDir)
	if err != nil {
		m.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}

	err = m.mcPlayBook()
	if err != nil {
		m.Reporter.ReportStatus(context.Background(), status, defaultResource)
		return err
	}
	status[statusField] = statusUpdated
	m.Reporter.ReportStatus(context.Background(), status, defaultResource)
	return nil
}

func (m *multiCloudProvisioner) deleteMCCluster() error {

	// nolint: errcheck
	defer os.RemoveAll(m.workDir)
	// nolint: errcheck
	defer m.manageSSHAgent(m.workDir, deleteAction)
	// best effort cleaning of nodes
	// need to export ssh-agent variables before killing the process
	err := m.manageSSHAgent(m.workDir, updateAction)
	if err != nil {
		return err
	}
	return m.mcPlayBook()
}

func (m *multiCloudProvisioner) isMCUpdated() (bool, error) {
	status := map[string]interface{}{}
	if _, err := os.Stat(m.getMCInventoryFile(m.workDir)); err == nil {
		ok, err := m.compareMCInventoryFile()
		if err != nil {
			status[statusField] = statusUpdateFailed
			m.Reporter.ReportStatus(context.Background(), status, defaultResource)
			return true, err
		}
		if ok {
			m.Log.Infof("%s inventory file is already up-to-date", defaultResource)
			return true, nil
		}
	}
	return false, nil
}

func (m *multiCloudProvisioner) verifyCloudStatus() error {
	for _, cloudRef := range m.clusterData.ClusterInfo.CloudRefs {
		err := waitForCloudStatusToBeUpdated(context.Background(), m.cluster.log,
			m.cluster.APIServer, cloudRef.UUID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *multiCloudProvisioner) compareMCInventoryFile() (bool, error) {

	tmpDir, err := ioutil.TempDir("", "inventory")
	if err != nil {
		return false, err
	}

	// nolint: errcheck
	defer os.RemoveAll(tmpDir)
	m.Log.Debugf("Creating temperory inventory at dir %s", tmpDir)

	err = m.createFiles(tmpDir)
	if err != nil {
		return false, err
	}

	err = m.runGenerateInventory(tmpDir)
	if err != nil {
		return false, err
	}
	newInventory, err := ioutil.ReadFile(m.getMCInventoryFile(tmpDir))
	if err != nil {
		return false, err
	}

	oldInventory, err := ioutil.ReadFile(m.getMCInventoryFile(m.workDir))
	if err != nil {
		return false, err
	}
	return bytes.Equal(oldInventory, newInventory), nil
}

func (m *multiCloudProvisioner) runGenerateInventory(workDir string) error {

	cmd := cloud.GetGenInventoryCmd(cloud.GetMultiCloudRepodir())

	args := strings.Split(fmt.Sprintf("-t %s -s %s -ts %s --instate %s --outstate %s",
		m.getClusterTopoFile(workDir), m.getClusterSecretFile(workDir), m.getTFStateFile(),
		mcState, mcState), " ")

	m.Log.Info("Generating inventory file multi-cloud provisioner")
	m.Log.Debugf("Command executed: %s %s", cmd,
		strings.Join(args, " "))

	if m.cluster.config.Test {
		return cloud.TestCmdHelper(cmd, args, workDir, testTemplate)
	}

	err := osutil.ExecCmdAndWait(m.Reporter, cmd, args, workDir)
	if err != nil {
		return err
	}

	// enabling openstack playbook only for test purposes
	// because its not yet supported by multi-cloud-deployer
	if m.isOrchestratorOpenstack() && m.cluster.config.Test {
		err := m.appendOpenStackConfigToInventory(m.getMCInventoryFile(workDir))
		if err != nil {
			return err
		}
	}

	m.Log.Infof("Successfully generated inventory file")

	return nil
}

// nolint: gocyclo
func (m *multiCloudProvisioner) mcPlayBook() error {
	args := []string{"-i", m.getMCInventoryFile(m.workDir)}
	if m.cluster.config.AnsibleSudoPass != "" {
		sudoArg := "-e ansible_sudo_pass=" + m.cluster.config.AnsibleSudoPass
		args = append(args, sudoArg)
	}

	switch m.clusterData.ClusterInfo.ProvisioningAction {
	case addCloud:
		if err := m.playDeployMCGW(args); err != nil {
			return err
		}
		if err := m.playMCInstancesConfig(args); err != nil {
			return err
		}
		// add tor playbook as well

		// enabling openstack playbook only for test purposes
		// because its not supported by multi-cloud-deployer
		if m.isOrchestratorOpenstack() && m.cluster.config.Test {
			openstackArgs := append(args, fmt.Sprintf("-e orchestrator=%s", openstack))
			if err := m.playOrchestratorProvision(openstackArgs); err != nil {
				return err
			}
		}
		if err := m.playMCK8SProvision(args); err != nil {
			return err
		}

		skipRoles := []string{"vrouter"}
		if err := m.playMCDeployContrail(args, skipRoles); err != nil {
			return err
		}
		skipRoles = []string{"config_database", "config,control", "webui",
			"analytics", "analytics_database", "k8s"}

		if err := m.playMCDeployContrail(args, skipRoles); err != nil {
			return err
		}

		for i := 0; i < setupRoutesPlayRetry; i++ {
			m.Log.Debugf("TRY %d of running setup controller gw route play", i+1)
			if err := m.playMCSetupControllerGWRoutes(args); err != nil {
				if i == setupRoutesPlayRetry-1 {
					return err
				}
				continue
			} else {
				break
			}
		}
		for i := 0; i < setupRoutesPlayRetry; i++ {
			m.Log.Debugf("TRY %d of running setup remote gw route play", i+1)
			if err := m.playMCSetupRemoteGWRoutes(args); err != nil {
				if i == setupRoutesPlayRetry-1 {
					return err
				}
				continue
			} else {
				break
			}
		}

		return m.playMCFixComputeDNS(args)

	case updateCloud:

		if err := m.playDeployMCGW(args); err != nil {
			return err
		}

		for i := 0; i < setupRoutesPlayRetry; i++ {
			m.Log.Debugf("TRY %d of running setup controller gw route play", i+1)
			if err := m.playMCSetupControllerGWRoutes(args); err != nil {
				if i == setupRoutesPlayRetry-1 {
					return err
				}
				continue
			} else {
				break
			}
		}

		if err := m.playMCInstancesConfig(args); err != nil {
			return err
		}

		// enabling openstack playbook only for test purposes
		// because its not supported by multi-cloud-deployer
		if m.isOrchestratorOpenstack() && m.cluster.config.Test {
			openstackArgs := append(args, fmt.Sprintf("-e orchestrator=%s", openstack))
			if err := m.playOrchestratorProvision(openstackArgs); err != nil {
				return err
			}
		}

		// add tor playbook as well
		if err := m.playMCK8SProvision(args); err != nil {
			return err
		}
		skipRoles := []string{"vrouter"}
		if err := m.playMCDeployContrail(args, skipRoles); err != nil {
			return err
		}
		skipRoles = []string{"config_database", "config,control", "webui",
			"analytics", "analytics_database", "k8s"}
		if err := m.playMCDeployContrail(args, skipRoles); err != nil {
			return err
		}

		for i := 0; i < setupRoutesPlayRetry; i++ {
			m.Log.Debugf("TRY %d of running setup remote gw route play", i+1)
			if err := m.playMCSetupRemoteGWRoutes(args); err != nil {
				if i == setupRoutesPlayRetry-1 {
					return err
				}
				continue
			} else {
				break
			}
		}

		return m.playMCFixComputeDNS(args)

	case deleteCloud:

		status := make(map[string]interface{})
		if m.action != deleteAction {
			status[statusField] = statusUpdateProgress
			m.Reporter.ReportStatus(context.Background(), status, defaultResource)
			status[statusField] = statusUpdated
		}
		//best effort cleaning up
		// nolint: errcheck
		_ = m.playMCContrailCleanup(args)
		// nolint: errcheck
		_ = m.playMCGatewayCleanup(args)

		if m.action != deleteAction {
			m.Reporter.ReportStatus(context.Background(), status, defaultResource)
		}
		return nil
	}

	return nil
}

func (m *multiCloudProvisioner) createFiles(workDir string) error {

	err := m.createClusterTopologyFile(workDir)
	if err != nil {
		return err
	}

	err = m.manageClusterSecret(workDir)
	if err != nil {
		return err
	}

	if err := m.createContrailCommonFile(m.getContrailCommonFile(workDir)); err != nil {
		return err
	}

	if err := m.createGatewayCommonFile(m.getGatewayCommonFile(workDir)); err != nil {
		return err
	}

	//tor file
	if err := m.createTORCommonFile(m.getTORCommonFile(workDir)); err != nil {
		return err
	}
	return nil

}

func readSSHAgentConfig(path string) (*SSHAgentConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	agentConfig := &SSHAgentConfig{}
	err = yaml.UnmarshalStrict(data, agentConfig)
	if err != nil {
		return nil, err
	}
	return agentConfig, nil
}

func (m *multiCloudProvisioner) readPubKeyConfig() (*PubKeyConfig, error) {

	secretFilePath := m.getClusterSecretFile(m.workDir)
	data, err := ioutil.ReadFile(secretFilePath)
	if err != nil {
		return nil, err
	}

	pubKeyConfig := &PubKeyConfig{}
	err = yaml.UnmarshalStrict(data, pubKeyConfig)
	if err != nil {
		return nil, err
	}
	return pubKeyConfig, nil

}

func (m *multiCloudProvisioner) writeSSHAgentConfig(conf *SSHAgentConfig,
	dest string) error {

	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	return fileutil.WriteToFile(dest, data, defaultFilePermRWOnly)

}

func (m *multiCloudProvisioner) runSSHAgent(workDir string, sshAgentPath string) (*SSHAgentConfig, error) {

	cmd := "ssh-agent"
	args := []string{"-s"}
	cmdline := exec.Command(cmd, args...)

	cmdOutput := &bytes.Buffer{}
	cmdline.Stdout = cmdOutput

	if err := cmdline.Run(); err != nil {
		return nil, err
	}

	stdout := string(cmdOutput.Bytes())

	stdoutList := strings.SplitAfterN(stdout, "=", 3)
	sockIDList := strings.SplitAfterN(stdoutList[1], ";", 2)
	pidList := strings.SplitAfterN(stdoutList[2], ";", 2)

	sshAgentConf := &SSHAgentConfig{
		AuthSock: strings.TrimSuffix(sockIDList[0], ";"),
		PID:      strings.TrimSuffix(pidList[0], ";"),
	}

	err := m.writeSSHAgentConfig(sshAgentConf, sshAgentPath)
	if err != nil {
		return nil, err
	}
	return sshAgentConf, nil
}

// nolint: gocyclo
func (m *multiCloudProvisioner) manageSSHAgent(workDir string,
	action string) error {

	// To-Do: to use ssh/agent library to create agent unix process
	sshAgentPath := m.getSSHAgentFile(workDir)
	_, err := os.Stat(sshAgentPath)
	if err != nil {
		if action == deleteAction {
			return nil
		}
		_, err = m.runSSHAgent(workDir, sshAgentPath)
		if err != nil {
			return err
		}
	} else {
		process, err := isSSHAgentProcessRunning(sshAgentPath) //nolint: govet
		if err != nil {
			_, err = m.runSSHAgent(workDir, sshAgentPath)
			if err != nil {
				return err
			}
		}

		if action == deleteAction && err == nil {
			return process.Kill()
		}
	}
	if m.cluster.config.Test {
		return nil
	}

	sshAgentConf, err := readSSHAgentConfig(sshAgentPath)
	if err != nil {
		return err
	}

	err = os.Setenv("SSH_AUTH_SOCK", sshAgentConf.AuthSock)
	if err != nil {
		return err
	}

	err = os.Setenv("SSH_AGENT_PID", sshAgentConf.PID)
	if err != nil {
		return err
	}

	m.Log.Debugf("SSH_AUTH_SOCK: %s", os.Getenv("SSH_AUTH_SOCK"))
	m.Log.Debugf("SSH_AGENT_PID: %s", os.Getenv("SSH_AGENT_PID"))

	pubKey, err := m.readPubKeyConfig()
	if err != nil {
		return err
	}

	pubCloudID, _, err := m.getPubPvtCloudID()
	if err != nil {
		return err
	}

	pubCloudSSHKeyDir := filepath.Join(cloud.GetCloudDir(pubCloudID),
		defaultSSHKeyRepo)

	clusterSSHKeyDir := filepath.Join(m.getMCWorkingDir(m.getWorkingDir()), defaultSSHKeyRepo)
	sshKeyName, ok := pubKey.Info["name"]

	err = copySSHKeyPair(pubCloudSSHKeyDir, clusterSSHKeyDir, sshKeyName)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("secret file format is not valid")
	}
	sshPvtKeyPath := filepath.Join(clusterSSHKeyDir, sshKeyName)

	return m.addPvtKeyToSSHAgent(sshPvtKeyPath)
}

func (m *multiCloudProvisioner) addPvtKeyToSSHAgent(keyPath string) error {

	cmd := "ssh-add"
	args := []string{"-d", fmt.Sprintf("%s", keyPath)}
	m.Log.Debugf("Executing command: %s", fmt.Sprintf("%s %s", cmd, keyPath))

	if m.cluster.config.Test {
		return cloud.TestCmdHelper(cmd, args, m.workDir, testTemplate)
	}
	// ignore if there is an error while deleting key,
	// this step is to make sure that updated keys are loaded to agent
	_ = osutil.ExecCmdAndWait(m.Reporter, cmd, args, m.workDir) // nolint: errcheck

	// readd the key
	args = []string{fmt.Sprintf("%s", keyPath)}
	m.Log.Debugf("Executing command: %s", fmt.Sprintf("%s %s", cmd, keyPath))

	if m.cluster.config.Test {
		return cloud.TestCmdHelper(cmd, args, m.workDir, testTemplate)
	}
	return osutil.ExecCmdAndWait(m.Reporter, cmd, args, m.workDir)
}

func copySSHKeyPair(srcSSHKeyDir string,
	destSSHKeyDir string, keyName string) error {

	srcSSHPvtKeyPath := filepath.Join(srcSSHKeyDir, keyName)
	srcSSHPubKeyPath := filepath.Join(srcSSHKeyDir, keyName+".pub")

	pvtKeyFileInfo, err := os.Stat(srcSSHPvtKeyPath)
	if err != nil {
		return err
	}
	pvtKey, err := fileutil.GetContent("file://" + srcSSHPvtKeyPath)
	if err != nil {
		return err
	}

	pubKeyFileInfo, err := os.Stat(srcSSHPubKeyPath)
	if err != nil {
		return err
	}
	pubKey, err := fileutil.GetContent("file://" + srcSSHPubKeyPath)
	if err != nil {
		return err
	}

	destSSHPvtKeyPath := filepath.Join(destSSHKeyDir, keyName)
	destSSHPubKeyPath := filepath.Join(destSSHKeyDir, keyName+".pub")

	if err = fileutil.WriteToFile(destSSHPvtKeyPath, pvtKey,
		pvtKeyFileInfo.Mode()); err != nil {
		return err
	}

	if err = fileutil.WriteToFile(destSSHPubKeyPath, pubKey,
		pubKeyFileInfo.Mode()); err != nil {
		return err
	}
	return nil
}

// nolint: gocyclo
func (m *multiCloudProvisioner) createClusterTopologyFile(workDir string) error {

	pubCloudID, pvtCloudID, err := m.getPubPvtCloudID()
	if err != nil {
		return err
	}

	onPremTopoFile := cloud.GetTopoFile(pvtCloudID)

	if _, err := os.Stat(onPremTopoFile); err == nil {
		onPremTopoContent, err := ioutil.ReadFile(onPremTopoFile)
		if err != nil {
			return err
		}
		err = fileutil.WriteToFile(m.getClusterTopoFile(workDir), onPremTopoContent, defaultFilePermRWOnly)
		if err != nil {
			return err
		}
	} else {
		return errors.New("topology file is not created for onprem cloud")
	}

	publicTopoFile := cloud.GetTopoFile(pubCloudID)

	if _, err := os.Stat(publicTopoFile); err == nil {
		publicTopoContent, err := ioutil.ReadFile(publicTopoFile)
		if err != nil {
			return err
		}
		err = fileutil.AppendToFile(m.getClusterTopoFile(workDir), publicTopoContent, defaultFilePermRWOnly)
		if err != nil {
			return err
		}
	} else {
		return errors.New("topology file is not created for public cloud")
	}
	return nil
}

func (m *multiCloudProvisioner) manageClusterSecret(workDir string) error {

	pubCloudID, _, err := m.getPubPvtCloudID()
	if err != nil {
		return err
	}

	secretFile := cloud.GetSecretFile(pubCloudID)

	return m.createClusterSecretFile(secretFile, workDir)
}

func (m *multiCloudProvisioner) createClusterSecretFile(secretFile string,
	workDir string) error {

	if _, err := os.Stat(secretFile); err == nil {
		secretFileContent, err := ioutil.ReadFile(secretFile)
		if err != nil {
			return err
		}
		err = fileutil.WriteToFile(m.getClusterSecretFile(workDir), secretFileContent, defaultFilePermRWOnly)
		if err != nil {
			return err
		}
		authRegcontent, err := m.getAuthRegistryContent(m.clusterData.ClusterInfo)
		if err != nil {
			return err
		}
		return fileutil.AppendToFile(m.getClusterSecretFile(workDir), authRegcontent, defaultFilePermRWOnly)
	}

	return errors.New("secret file is not created for public cloud")

}

func (m *multiCloudProvisioner) getAuthRegistryContent(cluster *models.ContrailCluster) ([]byte, error) {
	context := pongo2.Context{
		"cluster": cluster,
	}

	content, err := template.Apply(m.getAuthRegistryTemplate(), context)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (m *multiCloudProvisioner) createContrailCommonFile(destination string) error {
	m.Log.Info("Creating contrail/common.yml input file for multi-cloud deployer")
	contrailPassword, err := m.getContrailPassword()
	if err != nil {
		return err
	}

	context := pongo2.Context{
		"cluster":                   m.clusterData.ClusterInfo,
		"k8sCluster":                m.clusterData.GetK8sClusterInfo(),
		"defaultSSHUser":            m.clusterData.DefaultSSHUser,
		"defaultSSHPassword":        m.clusterData.DefaultSSHPassword,
		"defaultSSHKey":             m.clusterData.DefaultSSHKey,
		"defaultContrailUser":       defaultContrailUser,
		"defaultContrailPassword":   contrailPassword,
		"defaultContrailConfigPort": defaultContrailConfigPort,
		"defaultContrailTenant":     defaultContrailTenant,
	}
	content, err := template.Apply(m.getContrailCommonTemplate(), context)
	if err != nil {
		return err
	}

	err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	m.Log.Info("Created contrail/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createGatewayCommonFile(destination string) error {
	m.Log.Info("Creating gateway/common.yml input file for multi-cloud deployer")
	context := pongo2.Context{
		"cluster":    m.clusterData.ClusterInfo,
		"pathConfig": pathConfig,
		"bgpSecret":  bgpSecret,
	}
	content, err := template.Apply(m.getGatewayCommonTemplate(), context)
	if err != nil {
		return err
	}

	err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	m.Log.Info("Created gateway/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createTORCommonFile(destination string) error {
	m.Log.Info("Creating tor/common.yml input file for multi-cloud deployer")
	context := pongo2.Context{
		"torBGPSecret":  torBGPSecret,
		"torOSPFSecret": torOSPFSecret,
		"debugMCRoutes": debugMCRoutes,
	}

	content, err := template.Apply(m.getTORCommonTemplate(), context)
	if err != nil {
		return err
	}

	err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	m.Log.Info("Created tor/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) appendOpenStackConfigToInventory(destination string) error {
	m.Log.Info("Appending openstack config to inventory file")

	context := pongo2.Context{
		"openstackCluster": m.clusterData.GetOpenstackClusterInfo(),
	}
	content, err := template.Apply(m.getOpenstackConfigTemplate(), context)
	if err != nil {
		return err
	}

	err = fileutil.AppendToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	m.Log.Info("Appended openstack config to inventory file")
	return nil
}

func (m *multiCloudProvisioner) getContrailCommonTemplate() string {
	return filepath.Join(m.getTemplateRoot(), defaultContrailCommonTemplate)
}

func (m *multiCloudProvisioner) getAuthRegistryTemplate() string {
	return filepath.Join(m.getTemplateRoot(), defaultAuthRegistryTemplate)
}

func (m *multiCloudProvisioner) getGatewayCommonTemplate() string {
	return filepath.Join(m.getTemplateRoot(), defaultGatewayCommonTemplate)
}

func (m *multiCloudProvisioner) getTORCommonTemplate() string {
	return filepath.Join(m.getTemplateRoot(), defaultTORCommonTemplate)
}

func (m *multiCloudProvisioner) getOpenstackConfigTemplate() string {
	return filepath.Join(m.getTemplateRoot(), defaultOpenstackConfigTemplate)
}

func (m *multiCloudProvisioner) getTFStateFile() string {

	for _, c := range m.clusterData.CloudInfo {
		for _, prov := range c.CloudProviders {
			if prov.Type != onPrem {
				return cloud.GetTFStateFile(c.UUID)
			}
		}
	}
	return ""
}

func (m *multiCloudProvisioner) getMCInventoryFile(workDir string) string {
	return filepath.Join(workDir, defaultMCInventoryFile)
}

// use topology constant from cloud pkg
func (m *multiCloudProvisioner) getClusterTopoFile(workDir string) string {
	return filepath.Join(workDir, defaultTopologyFile)
}

func (m *multiCloudProvisioner) getClusterSecretFile(workDir string) string {
	return filepath.Join(workDir, defaultSecretFile)
}

func (m *multiCloudProvisioner) getSSHAgentFile(workDir string) string {
	return filepath.Join(workDir, defaultSSHAgentFile)
}

func (m *multiCloudProvisioner) getMCDeployerRepoDir() (ansibleRepoDir string) {
	return filepath.Join(defaultAnsibleRepoDir, mcAnsibleRepo)
}

func (m *multiCloudProvisioner) play(ansibleArgs []string) error {
	m.Log.Info("playing from mc deployer")
	repoDir := m.getMCDeployerRepoDir()
	m.Log.Infof("playing from mc repo: %s", repoDir)
	return m.playFromDir(repoDir, ansibleArgs)
}

func (m *multiCloudProvisioner) getContrailCommonFile(workDir string) string {
	return filepath.Join(workDir, defaultContrailCommonFile)
}

func (m *multiCloudProvisioner) getGatewayCommonFile(workDir string) string {
	return filepath.Join(workDir, defaultGatewayCommonFile)
}

func (m *multiCloudProvisioner) getTORCommonFile(workDir string) string {
	return filepath.Join(workDir, defaultTORCommonFile)
}

func (m *multiCloudProvisioner) getMCWorkingDir(clusterWorkDir string) string {
	return filepath.Join(clusterWorkDir, mcWorkDir)
}

func (m *multiCloudProvisioner) getContrailPassword() (string, error) {
	if m.isOrchestratorOpenstack() {
		openStackClusterInfo := m.clusterData.GetOpenstackClusterInfo()
		for _, v := range openStackClusterInfo.KollaPasswords.KeyValuePair {
			if v.Key == "keystone_admin_password" {
				return v.Value, nil
			}
		}
	}
	return defaultContrailPassword, nil
}

func (m *multiCloudProvisioner) isOrchestratorOpenstack() bool {
	if strings.ToLower(m.clusterData.ClusterInfo.Orchestrator) == openstack {
		return true
	}
	return false
}

func (m *multiCloudProvisioner) playDeployMCGW(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultMCGWDeployPlay)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCInstancesConfig(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultMCInstanceConfPlay)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCK8SProvision(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultMCKubernetesProvPlay)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCDeployContrail(ansibleArgs, skipRoles []string) error {
	skipRoleArgs := getSkipTagArgs(skipRoles)
	ansibleArgs = append(ansibleArgs, append(skipRoleArgs, defaultMCDeployContrail)...)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCSetupControllerGWRoutes(ansibleArgs []string) error {
	skipRoles := []string{"remotegw"}
	skipRoleArgs := getSkipTagArgs(skipRoles)
	ansibleArgs = append(ansibleArgs, append(skipRoleArgs, defaultMCSetupContrailRoutes)...)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCSetupRemoteGWRoutes(ansibleArgs []string) error {
	skipRoles := []string{"controllergw"}
	skipRoleArgs := getSkipTagArgs(skipRoles)
	ansibleArgs = append(ansibleArgs, append(skipRoleArgs, defaultMCSetupContrailRoutes)...)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCFixComputeDNS(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultMCFixComputeDNS)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCContrailCleanup(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultMCContrailCleanup)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCGatewayCleanup(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultMCGatewayCleanup)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) updateMCWorkDir() {
	m.workDir = m.getMCWorkingDir(m.getWorkingDir())
}

func getSkipTagArgs(tagsToBeSkipped []string) []string {
	args := []string{"--skip-tags"}
	return append(args, strings.Join(tagsToBeSkipped, ","))
}

func (m *multiCloudProvisioner) getPubPvtCloudID() (string, string, error) {

	var publicCloudID, onPremCloudID string
	for _, cloudRef := range m.clusterData.CloudInfo {
		for _, p := range cloudRef.CloudProviders {
			if p.Type == onPrem {
				onPremCloudID = cloudRef.UUID
				continue
			}
			if p.Type == aws || p.Type == azure {
				publicCloudID = cloudRef.UUID
				continue
			}
		}
	}
	if publicCloudID == "" || onPremCloudID == "" {
		return "", "", errors.New("public or OnPrem cloud is not added to cluster")
	}
	return publicCloudID, onPremCloudID, nil

}

func isSSHAgentProcessRunning(sshAgentPath string) (*os.Process, error) {

	sshAgentConf, err := readSSHAgentConfig(sshAgentPath)
	if err != nil {
		return nil, err
	}
	pid, err := strconv.Atoi(sshAgentConf.PID) // nolint: govet
	if err != nil {
		return nil, err
	}
	// check if process is running
	process, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return nil, err
	}
	return process, nil
}

func waitForCloudStatusToBeUpdated(ctx context.Context, log *logrus.Entry,
	httpClient *client.HTTP, cloudUUID string) error {

	return retry.Do(func() (retry bool, err error) {

		cloudResp, err := httpClient.GetCloud(ctx, &services.GetCloudRequest{
			ID: cloudUUID,
		},
		)
		if err != nil {
			return false, err
		}

		switch cloudResp.Cloud.ProvisioningState {
		case statusCreateProgress:
			return true, fmt.Errorf("waiting for create cloud %s to complete",
				cloudUUID)
		case statusUpdateProgress:
			return true, fmt.Errorf("waiting for update cloud %s to complete",
				cloudUUID)
		case statusNoState:
			return true, fmt.Errorf("waiting for cloud %s to complete processing",
				cloudUUID)
		case statusCreated:
			return false, nil
		case statusUpdated:
			return false, nil
		case statusCreateFailed:
			return false, fmt.Errorf("cloud %s status has failed in creating",
				cloudUUID)
		case statusUpdateFailed:
			return false, fmt.Errorf("cloud %s status has failed in updating",
				cloudUUID)
		}

		return false, fmt.Errorf("unknown cloud status %s for cloud %s",
			cloudResp.Cloud.ProvisioningState, cloudUUID)
	}, retry.WithLog(log.Logger),
		retry.WithInterval(statusRetryInterval))
}
