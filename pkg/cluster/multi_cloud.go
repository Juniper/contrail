package cluster

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/common"

	"github.com/flosch/pongo2"
)

const (
	defaultContrailCommonTemplate  = "contrail_common.tmpl"
	defaultGatewayCommonTemplate   = "gateway_common.tmpl"
	defaultTORCommonTemplate       = "tor_common.tmpl"
	defaultOpenstackConfigTemplate = "openstack_config.tmpl"

	mcWorkDir                 = "multi-cloud"
	defaultContrailCommonFile = "ansible/contrail/common.yml"
	defaultTORCommonFile      = "ansible/tor/common.yml"
	defaultGatewayCommonFile  = "ansible/gateway/common.yml"
	defaultMCInventoryFile    = "inventories/inventory.yml"
	defaultTopologyFile       = "topology.yml"
	defaultSecretFile         = "secret.yml"

	defaultContrailUser       = "admin"
	defaultContrailPassword   = "c0ntrail123"
	defaultContrailConfigPort = "8082"
	defaultContrailTenant     = "default-project"
	pathConfig                = "/etc/multicloud"
	upgradeKernel             = "True"
	kernelVersion             = "3.10.0-862.3.2.el7.x86_64"
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
	defaultMCContrailCleanup     = "ansible/contrail/playbooks/cleanup.yml"
	defaultMCGatewayCleanup      = "ansible/gateway/playbooks/cleanup.yml"
	testTemplate                 = "./../cloud/test_data/test_cmd.tmpl"

	openstack = "openstack"

	addCloud    = "ADD_CLOUD"
	updateCloud = "UPDATE_CLOUD"
	deleteCloud = "DELETE_CLOUD"

	aws    = "aws"
	azure  = "azure"
	onPrem = "private"
	public = "public"
)

type multiCloudProvisioner struct {
	ansibleProvisioner
	workDir string
}

func (m *multiCloudProvisioner) provision() error {

	m.updateMCWorkDir()
	switch m.clusterData.clusterInfo.ProvisioningAction {
	case addCloud:
		err := m.createMCCluster()
		if err != nil {
			return err
		}
		return m.createEndpoints()
	case updateCloud:
		updated, err := m.isMCUpdated()
		if err != nil {
			return err
		}
		if updated {
			return nil
		}
		err = m.updateMCCluster()
		if err != nil {
			return err
		}
	case deleteCloud:
		err := m.deleteMCCluster()
		if err != nil {
			return err
		}
	}
	return nil
}

// nolint: gocyclo
func (m *multiCloudProvisioner) createMCCluster() error {

	m.log.Infof("Starting %s of contrail cluster: %s",
		m.clusterData.clusterInfo.ProvisioningAction,
		m.clusterData.clusterInfo.FQName)

	status := map[string]interface{}{statusField: statusUpdateProgress}
	if m.action == createAction {
		status = map[string]interface{}{statusField: statusCreateProgress}
	}
	m.reporter.ReportStatus(status, defaultResource)

	status[statusField] = statusUpdateFailed
	if m.action == createAction {
		status[statusField] = statusCreateFailed
	}

	if !m.cluster.config.Test {
		if m.cluster.config.AnsibleFetchURL != "" {
			err := m.fetchAnsibleDeployer()
			if err != nil {
				m.reporter.ReportStatus(status, defaultResource)
				return err
			}
		}
		if m.cluster.config.AnsibleCherryPickRevision != "" {
			err := m.cherryPickAnsibleDeployer()
			if err != nil {
				m.reporter.ReportStatus(status, defaultResource)
				return err
			}
		}
		if m.cluster.config.AnsibleRevision != "" {
			err := m.resetAnsibleDeployer()
			if err != nil {
				m.reporter.ReportStatus(status, defaultResource)
				return err
			}
		}
	}

	err := m.createFiles(m.workDir)
	if err != nil {
		m.reporter.ReportStatus(status, defaultResource)
		return err
	}

	err = m.runGenerateInventory(m.workDir)
	if err != nil {
		m.reporter.ReportStatus(status, defaultResource)
		return err
	}

	err = m.mcPlayBook()
	if err != nil {
		m.reporter.ReportStatus(status, defaultResource)
		return err
	}
	return nil
}

func (m *multiCloudProvisioner) updateMCCluster() error {
	m.log.Infof("Starting %s of contrail cluster: %s", m.action, m.clusterData.clusterInfo.FQName)

	status := map[string]interface{}{}
	status[statusField] = statusUpdateProgress
	m.reporter.ReportStatus(status, defaultResource)
	status[statusField] = statusUpdateFailed

	err := m.createFiles(m.workDir)
	if err != nil {
		m.reporter.ReportStatus(status, defaultResource)
		return err
	}

	err = m.runGenerateInventory(m.workDir)
	if err != nil {
		m.reporter.ReportStatus(status, defaultResource)
		return err
	}

	err = m.mcPlayBook()
	if err != nil {
		m.reporter.ReportStatus(status, defaultResource)
		return err
	}
	return nil
}

func (m *multiCloudProvisioner) deleteMCCluster() error {

	// nolint: errcheck
	defer os.RemoveAll(m.workDir)
	// best effort cleaning of nodes
	// nolint: errcheck
	_ = m.mcPlayBook()
	return nil
}

func (m *multiCloudProvisioner) isMCUpdated() (bool, error) {
	status := map[string]interface{}{}
	if _, err := os.Stat(m.getMCInventoryFile(m.workDir)); err == nil {
		ok, err := m.compareMCInventoryFile()
		if err != nil {
			status[statusField] = statusUpdateFailed
			m.reporter.ReportStatus(status, defaultResource)
			return true, err
		}
		if ok {
			m.log.Infof("%s inventory file is already up-to-date", defaultResource)
			return true, nil
		}
	}
	return false, nil
}

func (m *multiCloudProvisioner) compareMCInventoryFile() (bool, error) {

	tmpDir, err := ioutil.TempDir("", "inventory")
	if err != nil {
		return false, err
	}

	// nolint: errcheck
	defer os.RemoveAll(tmpDir)
	m.log.Debugf("Creating temperory inventory at dir %s", tmpDir)

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

	m.log.Info("Generating inventory file multi-cloud provisioner")
	m.log.Debugf("Command executed: %s %s", cmd,
		strings.Join(args, " "))

	if m.cluster.config.Test {
		return cloud.TestCmdHelper(cmd, args, workDir, testTemplate)
	}

	err := common.ExecCmdAndWait(m.reporter, cmd, args, workDir)
	if err != nil {
		return err
	}

	if m.isOrchestratorOpenstack() {
		err := m.appendOpenStackConfigToInventory(m.getMCInventoryFile(workDir))
		if err != nil {
			return err
		}
	}

	m.log.Infof("Successfully generated inventory file")

	return nil
}

// nolint: gocyclo
func (m *multiCloudProvisioner) mcPlayBook() error {
	args := []string{"-i", m.getMCInventoryFile(m.workDir)}
	if m.cluster.config.AnsibleSudoPass != "" {
		sudoArg := "-e ansible_sudo_pass=" + m.cluster.config.AnsibleSudoPass
		args = append(args, sudoArg)
	}

	switch m.clusterData.clusterInfo.ProvisioningAction {
	case addCloud:
		if err := m.playDeployMCGW(args); err != nil {
			return err
		}
		if err := m.playMCInstancesConfig(args); err != nil {
			return err
		}
		// add tor playbook as well

		if m.isOrchestratorOpenstack() {
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
		if err := m.playMCSetupControllerGWRoutes(args); err != nil {
			return err
		}
		if err := m.playMCSetupRemoteGWRoutes(args); err != nil {
			return err
		}
	case updateCloud:

		if err := m.playDeployMCGW(args); err != nil {
			return err
		}
		if err := m.playMCSetupControllerGWRoutes(args); err != nil {
			return err
		}
		if err := m.playMCInstancesConfig(args); err != nil {
			return err
		}

		if m.isOrchestratorOpenstack() {
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
		if err := m.playMCSetupRemoteGWRoutes(args); err != nil {
			return err
		}
	case deleteCloud:

		//best effort cleaning up
		m.playMCContrailCleanup(args)
		m.playMCGatewayCleanup(args)
		return nil
	}

	if m.cluster.config.Action == deleteAction {
		// nolint: errcheck
		_ = os.RemoveAll(m.getWorkingDir())
	}

	return nil
}

func (m *multiCloudProvisioner) createFiles(workDir string) error {

	err := m.createClusterTopologyFile(workDir)
	if err != nil {
		return err
	}

	err = m.createClusterSecretFile(workDir)
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

// nolint: gocyclo
func (m *multiCloudProvisioner) createClusterTopologyFile(workDir string) error {

	cloudInfo := make(map[string]string)
	for _, cloudRef := range m.clusterData.cloudInfo {
		for _, p := range cloudRef.CloudProviders {
			if p.Type == onPrem {
				cloudInfo[onPrem] = cloudRef.UUID
				break
			}
			if p.Type == aws || p.Type == azure {
				cloudInfo[public] = cloudRef.UUID
				break
			}
		}
	}

	onPremCloudID, ok := cloudInfo[onPrem]
	if !ok {
		return errors.New("private cloud object is not created and attached to cluster object")
	}
	onPremTopoFile := cloud.GetTopoFile(onPremCloudID)

	if _, err := os.Stat(onPremTopoFile); err == nil {
		onPremTopoContent, err := ioutil.ReadFile(onPremTopoFile)
		if err != nil {
			return err
		}
		err = common.WriteToFile(m.getClusterTopoFile(workDir), onPremTopoContent, defaultFilePermRWOnly)
		if err != nil {
			return err
		}
	} else {
		return errors.New("topology file is not created for onprem cloud")
	}

	publicCloudID, ok := cloudInfo[public]
	if !ok {
		return errors.New("public cloud object is not created and attached to cluster object")
	}
	publicTopoFile := cloud.GetTopoFile(publicCloudID)

	if _, err := os.Stat(publicTopoFile); err == nil {
		publicTopoContent, err := ioutil.ReadFile(publicTopoFile)
		if err != nil {
			return err
		}
		err = common.AppendToFile(m.getClusterTopoFile(workDir), publicTopoContent, defaultFilePermRWOnly)
		if err != nil {
			return err
		}
	} else {
		return errors.New("topology file is not created for public cloud")
	}

	return nil
}

func (m *multiCloudProvisioner) createClusterSecretFile(workDir string) error {

	var publicCloudID string
	for _, cloudObj := range m.clusterData.cloudInfo {
		for _, p := range cloudObj.CloudProviders {
			if p.Type != onPrem {
				publicCloudID = cloudObj.UUID
				break
			}
		}
		if publicCloudID != "" {
			break
		}
	}

	if publicCloudID == "" {
		return errors.New("public cloud object is not created or object is not attached to cluster object")
	}

	secretFile := cloud.GetSecretFile(publicCloudID)
	if _, err := os.Stat(secretFile); err == nil {
		secretFileContent, err := ioutil.ReadFile(secretFile)
		if err != nil {
			return err
		}
		err = common.WriteToFile(m.getClusterSecretFile(workDir), secretFileContent, defaultFilePermRWOnly)
		if err != nil {
			return err
		}
	} else {
		return errors.New("secret file is not created for public cloud")
	}

	return nil
}

func (m *multiCloudProvisioner) createContrailCommonFile(destination string) error {
	m.log.Info("Creating contrail/common.yml input file for multi-cloud deployer")
	SSHUser, SSHPassword, SSHKey, err := m.cluster.getDefaultCredential()
	if err != nil {
		return err
	}
	contrailPassword, err := m.getContrailPassword()
	if err != nil {
		return err
	}

	context := pongo2.Context{
		"cluster":                   m.clusterData.clusterInfo,
		"k8sCluster":                m.clusterData.getK8sClusterInfo(),
		"defaultSSHUser":            SSHUser,
		"defaultSSHPassword":        SSHPassword,
		"defaultSSHKey":             SSHKey,
		"defaultContrailUser":       defaultContrailUser,
		"defaultContrailPassword":   contrailPassword,
		"defaultContrailConfigPort": defaultContrailConfigPort,
		"defaultContrailTenant":     defaultContrailTenant,
	}
	content, err := common.Apply(m.getContrailCommonTemplate(), context)
	if err != nil {
		return err
	}

	err = common.WriteToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	m.log.Info("Created contrail/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createGatewayCommonFile(destination string) error {
	m.log.Info("Creating gateway/common.yml input file for multi-cloud deployer")
	context := pongo2.Context{
		"cluster":       m.clusterData.clusterInfo,
		"pathConfig":    pathConfig,
		"upgradeKernel": upgradeKernel,
		"kernelVersion": kernelVersion,
		"bgpSecret":     bgpSecret,
	}
	content, err := common.Apply(m.getGatewayCommonTemplate(), context)
	if err != nil {
		return err
	}

	err = common.WriteToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	m.log.Info("Created gateway/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createTORCommonFile(destination string) error {
	m.log.Info("Creating tor/common.yml input file for multi-cloud deployer")
	context := pongo2.Context{
		"torBGPSecret":  torBGPSecret,
		"torOSPFSecret": torOSPFSecret,
		"debugMCRoutes": debugMCRoutes,
	}
	content, err := common.Apply(m.getTORCommonTemplate(), context)
	if err != nil {
		return err
	}

	err = common.WriteToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	m.log.Info("Created tor/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) appendOpenStackConfigToInventory(destination string) error {
	m.log.Info("Appending openstack config to inventory file")

	context := pongo2.Context{
		"openstackCluster": m.clusterData.getOpenstackClusterInfo(),
	}
	content, err := common.Apply(m.getOpenstackConfigTemplate(), context)
	if err != nil {
		return err
	}

	err = common.AppendToFile(destination, content, defaultFilePermRWOnly)
	if err != nil {
		return err
	}
	m.log.Info("Appended openstack config to inventory file")
	return nil
}

func (m *multiCloudProvisioner) getContrailCommonTemplate() string {
	return filepath.Join(m.getTemplateRoot(), defaultContrailCommonTemplate)
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

	for _, c := range m.clusterData.cloudInfo {
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
		openStackClusterInfo := m.clusterData.getOpenstackClusterInfo()
		for _, v := range openStackClusterInfo.KollaPasswords.KeyValuePair {
			if v.Key == "keystone_admin_password" {
				return v.Value, nil
			}
		}
	}
	return defaultContrailPassword, nil
}

func (m *multiCloudProvisioner) isOrchestratorOpenstack() bool {
	return strings.ToLower(m.clusterData.clusterInfo.Orchestrator) == openstack
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

func (m *multiCloudProvisioner) playMCContrailCleanup(ansibleArgs []string) {
	ansibleArgs = append(ansibleArgs, defaultMCContrailCleanup)
	// nolint: errcheck
	_ = m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCGatewayCleanup(ansibleArgs []string) {
	ansibleArgs = append(ansibleArgs, defaultMCGatewayCleanup)
	// nolint: errcheck
	_ = m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) updateMCWorkDir() {
	m.workDir = m.getMCWorkingDir(m.getWorkingDir())
}

func getSkipTagArgs(tagsToBeSkipped []string) []string {
	args := []string{"--skip-tags"}
	return append(args, strings.Join(tagsToBeSkipped, ","))
}
