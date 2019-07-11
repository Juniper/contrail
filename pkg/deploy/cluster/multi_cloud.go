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

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/osutil"
	"github.com/Juniper/contrail/pkg/retry"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
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
	bgpMCRoutesForController  = "False"
	debugMCRoutes             = "False"
	torBGPSecret              = "contrail_secret"
	torOSPFSecret             = "contrail_secret"

	testTemplate = "./../../cloud/test_data/test_cmd.tmpl"

	openstack = "openstack"

	addCloud    = "ADD_CLOUD"
	updateCloud = "UPDATE_CLOUD"
	deleteCloud = "DELETE_CLOUD"

	aws    = "aws"
	azure  = "azure"
	gcp    = "gcp"
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

// nolint: gocyclo
func (m *multiCloudProvisioner) Deploy() error {
	m.updateMCWorkDir()
	switch m.clusterData.ClusterInfo.ProvisioningAction {
	case addCloud:
		updated, err := m.isMCUpdated()
		if err != nil {
			m.Log.Errorf("add cloud failed with err: %s", err)
			return err
		}
		if updated {
			return nil
		}
		err = m.createMCCluster()
		if err != nil {
			m.Log.Errorf("add cloud failed with err: %s", err)
			return err
		}
	case updateCloud:
		updated, err := m.isMCUpdated()
		if err != nil {
			m.Log.Errorf("update cloud failed with err: %s", err)
			return err
		}
		if updated {
			return nil
		}
		err = m.updateMCCluster()
		if err != nil {
			m.Log.Errorf("update cloud failed with err: %s", err)
			return err
		}
	case deleteCloud:
		if m.clusterData.ClusterInfo.ProvisioningState != statusNoState {
			return nil
		}
		m.Reporter.ReportStatus(context.Background(), statusUpdateProgress, defaultResource)

		if err := m.deleteMCCluster(); err != nil {
			m.Reporter.ReportStatus(context.Background(), statusUpdateFailed, defaultResource)
			m.Log.Errorf("delete cloud failed with err: %s", err)
			return err
		}
		m.Reporter.ReportStatus(context.Background(), statusUpdated, defaultResource)
	}
	return nil
}

// TODO: Take care of this
// nolint: gocyclo
func (m *multiCloudProvisioner) createMCCluster() error {
	m.Log.Infof("Starting %s of contrail cluster: %s",
		m.clusterData.ClusterInfo.ProvisioningAction,
		m.clusterData.ClusterInfo.FQName)

	var failedStatus string
	if m.action == createAction {
		m.Reporter.ReportStatus(context.Background(), statusCreateProgress, defaultResource)
		failedStatus = statusCreateFailed
	} else {
		m.Reporter.ReportStatus(context.Background(), statusUpdateProgress, defaultResource)
		failedStatus = statusUpdateFailed
	}

	if err := m.verifyCloudStatus(); err != nil {
		m.Reporter.ReportStatus(context.Background(), failedStatus, defaultResource)
		return err
	}

	// TODO: Unwrap this part of code from if statement after changing test framework
	if !m.cluster.config.Test {
		if err := m.updateAnsibleDeployer(); err != nil {
			m.Reporter.ReportStatus(context.Background(), failedStatus, defaultResource)
			return err
		}
	}

	if err := m.createFiles(); err != nil {
		m.Reporter.ReportStatus(context.Background(), failedStatus, defaultResource)
		return err
	}

	if err := m.manageSSHAgent(createAction); err != nil {
		m.Reporter.ReportStatus(context.Background(), failedStatus, defaultResource)
		return err
	}

	if err := m.multiCloudCLIRunAll(); err != nil {
		m.Reporter.ReportStatus(context.Background(), failedStatus, defaultResource)
		return err
	}

	if m.action == createAction {
		m.Reporter.ReportStatus(context.Background(), statusCreated, defaultResource)
	} else {
		m.Reporter.ReportStatus(context.Background(), statusUpdated, defaultResource)
	}
	return nil
}

func (m *multiCloudProvisioner) updateAnsibleDeployer() error {
	if m.cluster.config.AnsibleFetchURL != "" {
		if err := m.fetchAnsibleDeployer(); err != nil {
			return err
		}
	}
	if m.cluster.config.AnsibleCherryPickRevision != "" {
		if err := m.cherryPickAnsibleDeployer(); err != nil {
			return err
		}
	}
	if m.cluster.config.AnsibleRevision != "" {
		if err := m.resetAnsibleDeployer(); err != nil {
			return err
		}
	}
	return nil
}

func (m *multiCloudProvisioner) multiCloudCLIRunAll() error {
	topology := m.getFilePath(defaultTopologyFile)
	secret := m.getFilePath(defaultSecretFile)
	// TODO: Remove this after refactoring test framework.
	if m.contrailAnsibleDeployer.ansibleClient.IsTest() {
		return m.mockCLI("deployer all run --topology " + topology + " --secret " + secret + " --retry 10")
	}
	return m.ansibleClient.PlayViaDeployer("all", "run", "--topology", topology, "--secret", secret, "--retry", "10")
}

func (m *multiCloudProvisioner) mockCLI(cliCommand string) error {
	content, err := template.Apply("./test_data/test_mc_cli_command.tmpl", pongo2.Context{
		"command": cliCommand,
	})
	if err != nil {
		return err
	}

	return fileutil.AppendToFile(
		filepath.Join(m.contrailAnsibleDeployer.ansibleClient.GetWorkingDirectory(), "executed_ansible_playbook.yml"),
		content,
		// TODO: use const
		0600,
	)
}

func (m *multiCloudProvisioner) updateMCCluster() error {
	m.Log.Infof("Starting %s of contrail cluster: %s", m.action, m.clusterData.ClusterInfo.FQName)
	m.Reporter.ReportStatus(context.Background(), statusUpdateProgress, defaultResource)

	if err := m.verifyCloudStatus(); err != nil {
		m.Reporter.ReportStatus(context.Background(), statusUpdateFailed, defaultResource)
		return err
	}

	if err := m.createFiles(); err != nil {
		m.Reporter.ReportStatus(context.Background(), statusUpdateFailed, defaultResource)
		return err
	}

	if err := m.manageSSHAgent(updateAction); err != nil {
		m.Reporter.ReportStatus(context.Background(), statusUpdateFailed, defaultResource)
		return err
	}

	// TODO: Ensure playMCSetupControllerGWRoutes ordering isn't important
	if err := m.multiCloudCLIRunAll(); err != nil {
		m.Reporter.ReportStatus(context.Background(), statusUpdateFailed, defaultResource)
		return err
	}

	m.Reporter.ReportStatus(context.Background(), statusUpdated, defaultResource)
	return nil
}

func (m *multiCloudProvisioner) deleteMCCluster() error {
	// need to export ssh-agent variables before killing the process
	if err := m.manageSSHAgent(updateAction); err != nil {
		return err
	}

	if err := m.multiCloudCLICleanup(); err != nil {
		return err
	}

	// TODO: Check this action.
	if m.action != deleteAction {
		if err := m.removeCloudRefsFromCluster(); err != nil {
			return err
		}
	}
	// nolint: errcheck
	defer m.manageSSHAgent(deleteAction)
	// nolint: errcheck
	defer os.RemoveAll(m.workDir)
	return nil
}

func (m *multiCloudProvisioner) multiCloudCLICleanup() error {
	invPath := m.getFilePath(defaultMCInventoryFile)
	// TODO: Remove this after refactoring test framework.
	if m.contrailAnsibleDeployer.ansibleClient.IsTest() {
		return m.mockCLI(fmt.Sprintf("deployer all clean --inventory %v --retry %v", invPath, 5))
	}
	return m.ansibleClient.PlayViaDeployer("all", "clean", "--inventory", invPath, "--retry", "5")
}

func (m *multiCloudProvisioner) isMCUpdated() (bool, error) {
	if m.clusterData.ClusterInfo.ProvisioningState == statusNoState {
		return false, nil
	}
	if _, err := os.Stat(m.getFilePath(defaultTopologyFile)); err != nil {
		return false, nil
	}
	// TODO: Refactor compareClusterTopologyFile
	ok, err := m.compareClusterTopologyFile()
	if err != nil {
		m.Reporter.ReportStatus(context.Background(), statusUpdateFailed, defaultResource)
		return true, err
	}
	if ok {
		m.Log.Infof("%s topology file is already up-to-date", defaultResource)
	}
	return ok, nil
}

func (m *multiCloudProvisioner) verifyCloudStatus() error {
	for _, cloudRef := range m.clusterData.ClusterInfo.CloudRefs {
		if err := waitForCloudStatusToBeUpdated(
			context.Background(), m.cluster.log, m.cluster.APIServer, cloudRef.UUID,
		); err != nil {
			return err
		}
	}
	return nil
}

func (m *multiCloudProvisioner) compareClusterTopologyFile() (bool, error) {
	tmpDir, err := ioutil.TempDir("", "topology")
	if err != nil {
		return false, err
	}

	// nolint: errcheck
	defer os.RemoveAll(tmpDir)
	m.Log.Debugf("Creating temperory topology at dir %s", tmpDir)

	if err = m.createClusterTopologyFile(tmpDir); err != nil {
		return false, err
	}

	newTopology, err := ioutil.ReadFile(filepath.Join(tmpDir, defaultTopologyFile))
	if err != nil {
		return false, err
	}

	oldTopology, err := ioutil.ReadFile(m.getFilePath(defaultTopologyFile))
	if err != nil {
		return false, err
	}
	return bytes.Equal(oldTopology, newTopology), nil
}

func (m *multiCloudProvisioner) createFiles() error {
	if err := m.createClusterTopologyFile(m.workDir); err != nil {
		return err
	}
	if err := m.createClusterSecretFile(); err != nil {
		return err
	}
	if err := m.createContrailCommonFile(m.getFilePath(defaultContrailCommonFile)); err != nil {
		return err
	}
	if err := m.createGatewayCommonFile(m.getFilePath(defaultGatewayCommonFile)); err != nil {
		return err
	}
	if err := m.createTORCommonFile(m.getFilePath(defaultTORCommonFile)); err != nil {
		return err
	}
	return nil
}

func (m *multiCloudProvisioner) removeCloudRefsFromCluster() error {
	for _, cloudRef := range m.clusterData.ClusterInfo.CloudRefs {
		if _, err := m.cluster.APIServer.DeleteContrailClusterCloudRef(
			context.Background(), &services.DeleteContrailClusterCloudRefRequest{
				ID:                      m.clusterData.ClusterInfo.UUID,
				ContrailClusterCloudRef: cloudRef,
			},
		); err != nil {
			return err
		}
	}
	return nil
}

func readSSHAgentConfig(path string) (*SSHAgentConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	agentConfig := &SSHAgentConfig{}
	if err = yaml.UnmarshalStrict(data, agentConfig); err != nil {
		return nil, err
	}
	return agentConfig, nil
}

func (m *multiCloudProvisioner) readPubKeyConfig() (*PubKeyConfig, error) {
	data, err := ioutil.ReadFile(m.getFilePath(defaultSecretFile))
	if err != nil {
		return nil, err
	}

	pubKeyConfig := &PubKeyConfig{}
	if err = yaml.Unmarshal(data, pubKeyConfig); err != nil {
		return nil, err
	}
	return pubKeyConfig, nil
}

func (m *multiCloudProvisioner) writeSSHAgentConfig(conf *SSHAgentConfig, dest string) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return fileutil.WriteToFile(dest, data, defaultFilePermRWOnly)
}

func (m *multiCloudProvisioner) runSSHAgent(sshAgentPath string) (*SSHAgentConfig, error) {
	cmd := "ssh-agent"
	args := []string{"-s"}
	cmdline := exec.Command(cmd, args...)

	cmdOutput := &bytes.Buffer{}
	cmdline.Stdout = cmdOutput
	// set pgid enables creation of all child process with pgid of parent process
	cmdline.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

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

	if err := m.writeSSHAgentConfig(sshAgentConf, sshAgentPath); err != nil {
		return nil, err
	}
	return sshAgentConf, nil
}

// nolint: gocyclo
func (m *multiCloudProvisioner) manageSSHAgent(action string) error {
	// To-Do: to use ssh/agent library to create agent unix process
	sshAgentPath := m.getFilePath(defaultSSHAgentFile)
	_, err := os.Stat(sshAgentPath)
	if err != nil {
		if action == deleteAction {
			return nil
		}
		if _, err = m.runSSHAgent(sshAgentPath); err != nil {
			return err
		}
	} else {
		process, err := isSSHAgentProcessRunning(sshAgentPath) //nolint: govet
		if err != nil {
			_, err = m.runSSHAgent(sshAgentPath)
			if err != nil {
				return err
			}
		}

		if action == deleteAction {
			// sends signal to pgid if pid is a negative number
			// this helps in killing all the child (zombie leftover process)
			return syscall.Kill(-process.Pid, syscall.SIGKILL)
		}
	}
	if m.cluster.config.Test {
		return nil
	}

	sshAgentConf, err := readSSHAgentConfig(sshAgentPath)
	if err != nil {
		return err
	}

	if err = os.Setenv("SSH_AUTH_SOCK", sshAgentConf.AuthSock); err != nil {
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
	sshPrivName, ok := pubKey.Info["name"]
	sshPubName := sshPrivName + ".pub"

	if err := copyKey(pubCloudSSHKeyDir, clusterSSHKeyDir, sshPrivName); err != nil {
		return err
	}

	if err := copyKey(pubCloudSSHKeyDir, clusterSSHKeyDir, sshPubName); err != nil {
		return err
	}

	if !ok {
		return errors.New("secret file format is not valid")
	}
	sshPvtKeyPath := filepath.Join(clusterSSHKeyDir, sshPrivName)

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

func copyKey(srcDir, dstDir, keyName string) error {
	srcKeyPath := filepath.Join(srcDir, keyName)

	keyFileInfo, err := os.Stat(srcKeyPath)
	if err != nil {
		return err
	}
	keyContent, err := fileutil.GetContent("file://" + srcKeyPath)
	if err != nil {
		return err
	}

	dstKeyPath := filepath.Join(dstDir, keyName)
	return fileutil.WriteToFile(dstKeyPath, keyContent, keyFileInfo.Mode())
}

// nolint: gocyclo
func (m *multiCloudProvisioner) createClusterTopologyFile(destinationDir string) error {
	pubCloudID, pvtCloudID, err := m.getPubPvtCloudID()
	if err != nil {
		return err
	}
	topologyFile := filepath.Join(destinationDir, defaultTopologyFile)

	// remove old topology file
	if err = os.Remove(topologyFile); err != nil && !os.IsNotExist(err) {
		return errors.Wrapf(err, "unexpected error occurred during old topology file deletion %v", topologyFile)
	}

	if err = m.appendFileToFile(cloud.GetTopoFile(pvtCloudID), topologyFile); err != nil {
		return errors.Wrap(err, "topology file is not created for onprem cloud")
	}

	err = m.appendFileToFile(cloud.GetTopoFile(pubCloudID), topologyFile)
	return errors.Wrap(err, "topology file is not created for public cloud")
}

func (m *multiCloudProvisioner) appendFileToFile(source, destination string) error {
	content, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	return fileutil.AppendToFile(destination, content, defaultFilePermRWOnly)
}

func (m *multiCloudProvisioner) createClusterSecretFile() error {
	pubCloudID, _, err := m.getPubPvtCloudID()
	if err != nil {
		return err
	}
	secretFile := m.getFilePath(defaultSecretFile)

	// remove old secret file
	if err = os.Remove(secretFile); err != nil && !os.IsNotExist(err) {
		return errors.Wrapf(err, "unexpected error occurred during old secret file deletion %v", secretFile)
	}

	if err = m.appendFileToFile(cloud.GetSecretFile(pubCloudID), secretFile); err != nil {
		return errors.Wrap(err, "secret file is not created for public cloud")
	}

	authRegcontent, err := m.getAuthRegistryContent(m.clusterData.ClusterInfo)
	if err != nil {
		return err
	}
	return fileutil.AppendToFile(secretFile, authRegcontent, defaultFilePermRWOnly)
}

func (m *multiCloudProvisioner) getAuthRegistryContent(cluster *models.ContrailCluster) ([]byte, error) {
	ctx := pongo2.Context{
		"cluster": cluster,
	}

	return template.Apply(m.getTemplatePath(defaultAuthRegistryTemplate), ctx)
}

func (m *multiCloudProvisioner) createContrailCommonFile(destination string) error {
	m.Log.Info("Creating contrail/common.yml input file for multi-cloud deployer")
	contrailPassword, err := m.getContrailPassword()
	if err != nil {
		return err
	}
	ctx := pongo2.Context{
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
	if err := createFileFromTemplate(
		m.getTemplatePath(defaultContrailCommonTemplate), destination, ctx,
	); err != nil {
		return err
	}
	m.Log.Info("Created contrail/common.yml input file for multi-cloud deployer")
	return nil
}

func createFileFromTemplate(templatePath, destination string, ctx pongo2.Context) error {
	content, err := template.Apply(templatePath, ctx)
	if err != nil {
		return err
	}

	if err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly); err != nil {
		return err
	}
	return nil
}

func (m *multiCloudProvisioner) createGatewayCommonFile(destination string) error {
	m.Log.Info("Creating gateway/common.yml input file for multi-cloud deployer")
	ctx := pongo2.Context{
		"cluster":    m.clusterData.ClusterInfo,
		"pathConfig": pathConfig,
		"bgpSecret":  bgpSecret,
	}
	if err := createFileFromTemplate(
		m.getTemplatePath(defaultGatewayCommonTemplate), destination, ctx,
	); err != nil {
		return err
	}
	m.Log.Info("Created gateway/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createTORCommonFile(destination string) error {
	m.Log.Info("Creating tor/common.yml input file for multi-cloud deployer")
	ctx := pongo2.Context{
		"torBGPSecret":             torBGPSecret,
		"torOSPFSecret":            torOSPFSecret,
		"debugMCRoutes":            debugMCRoutes,
		"bgpMCRoutesForController": bgpMCRoutesForController,
	}
	if err := createFileFromTemplate(m.getTemplatePath(defaultTORCommonTemplate), destination, ctx); err != nil {
		return err
	}
	m.Log.Info("Created tor/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) appendOpenStackConfigToInventory(destination string) error {
	m.Log.Info("Appending openstack config to inventory file")

	ctx := pongo2.Context{
		"openstackCluster": m.clusterData.GetOpenstackClusterInfo(),
	}
	content, err := template.Apply(m.getTemplatePath(defaultOpenstackConfigTemplate), ctx)
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

func (m *multiCloudProvisioner) getTemplatePath(templateFile string) string {
	return filepath.Join(m.getTemplateRoot(), templateFile)
}

func (m *multiCloudProvisioner) getFilePath(file string) string {
	return filepath.Join(m.workDir, file)
}

func (m *multiCloudProvisioner) getMCDeployerRepoDir() (ansibleRepoDir string) {
	return filepath.Join(defaultAnsibleRepoDir, mcAnsibleRepo)
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

func (m *multiCloudProvisioner) updateMCWorkDir() {
	m.workDir = m.getMCWorkingDir(m.getWorkingDir())
}

func (m *multiCloudProvisioner) getPubPvtCloudID() (string, string, error) {
	var publicCloudID, onPremCloudID string
	for _, cloudRef := range m.clusterData.CloudInfo {
		for _, p := range cloudRef.CloudProviders {
			if p.Type == onPrem {
				onPremCloudID = cloudRef.UUID
				continue
			}
			if p.Type == aws || p.Type == azure || p.Type == gcp {
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
		cloudResp, err := httpClient.GetCloud(ctx, &services.GetCloudRequest{ID: cloudUUID})
		if err != nil {
			return false, err
		}

		switch cloudResp.Cloud.ProvisioningState {
		case statusCreateProgress:
			return true, fmt.Errorf("waiting for create cloud %s to complete", cloudUUID)
		case statusUpdateProgress:
			return true, fmt.Errorf("waiting for update cloud %s to complete", cloudUUID)
		case statusNoState:
			return true, fmt.Errorf("waiting for cloud %s to complete processing", cloudUUID)
		case statusCreated:
			return false, nil
		case statusUpdated:
			return false, nil
		case statusCreateFailed:
			return false, fmt.Errorf("cloud %s status has failed in creating", cloudUUID)
		case statusUpdateFailed:
			return false, fmt.Errorf("cloud %s status has failed in updating", cloudUUID)
		}

		return false, fmt.Errorf("unknown cloud status %s for cloud %s", cloudResp.Cloud.ProvisioningState, cloudUUID)
	},
		retry.WithLog(log.Logger),
		retry.WithInterval(statusRetryInterval),
	)
}
