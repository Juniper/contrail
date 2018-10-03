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
	"github.com/flosch/pongo2"
)

const (
	defaultContrailCommonTemplate = "contrail_common.tmpl"
	defaultGatewayCommonTemplate  = "gateway_common.tmpl"
	defaultTORCommonTemplate      = "tor_common.tmpl"
	defaultContrailUser           = "admin"
	defaultContrailPassword       = "c0ntrail123"
	defaultContrailConfigPort     = "8082"
	defaultContrailTenant         = "default-project"
	pathConfig                    = "/etc/multicloud"
	upgradeKernel                 = "True"
	kernelVersion                 = "3.10.0-862.3.2.el7.x86_64"
	asn                           = "65000"
	vpnLoNetwork                  = "100.65.0.0/16"
	vpnNetwork                    = "100.64.0.0/16"
	openvpnPort                   = "443"
	bfdInterval                   = "200ms"
	bfdMultiplier                 = "5"
	bfdIntervalMultihop           = "500ms"
	bfdMultiplierMultihop         = "5"
	bgpSecret                     = "bgp_secret"
	debugMCRoutes                 = "False"
	torBGPSecret                  = "contrail_secret"
	torOSPFSecret                 = "contrail_secret"

	mcState = "state.yml"

	defaultMCGWDeployPlay        = "ansible/gateway/playbooks/deploy_and_run_all.yml"
	defaultMCInstanceConfPlay    = "ansible/contrail/playbooks/configure.yml"
	defaultMCKubernetesProvPlay  = "ansible/contrail/playbooks/orchestrator.yml"
	defaultMCDeployContrail      = "ansible/contrail/playbooks/deploy.yml"
	defaultMCSetupContrailRoutes = "ansible/contrail/playbooks/add_tunnel_routes.yml"
	defaultMCContrailCleanup     = "ansible/contrail/playbooks/cleanup.yml"
	defaultMCGatewayCleanup      = "ansible/gateway/playbooks/cleanup.yml"
)

type multiCloudProvisioner struct {
	ansibleProvisioner
	workDir string
}

func (m *multiCloudProvisioner) provision() error {

	m.updateMCWorkDir()
	switch m.action {
	case createAction:
		err := m.createMCCluster()
		if err != nil {
			return err
		}
		return m.createEndpoints()
	case updateAction:
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
	case deleteAction:
		err := m.deleteMCCluster()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *multiCloudProvisioner) createMCCluster() error {

	m.log.Infof("Starting %s of contrail cluster: %s",
		m.clusterData.clusterInfo.ProvisioningAction,
		m.clusterData.clusterInfo.FQName)

	status := map[string]interface{}{statusField: statusUpdateProgress}
	if m.clusterData.clusterInfo.ProvisioningState == createAction {
		status = map[string]interface{}{statusField: statusCreateProgress}
	}
	m.reporter.ReportStatus(status, defaultResource)

	status[statusField] = statusUpdateFailed
	if m.clusterData.clusterInfo.ProvisioningState == createAction {
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

	defer os.RemoveAll(m.workDir)
	// best effort cleaning of nodes
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

	if !m.cluster.config.Test {
		err := cloud.ExecCmd(m.reporter, cmd, args, workDir)
		if err != nil {
			return err
		}
		m.log.Infof("Successfully generated inventory file")
	}
	return nil
}

func (m *multiCloudProvisioner) mcPlayBook() error {
	args := []string{"-i", m.getMCInventoryFile(m.workDir)}
	if m.cluster.config.AnsibleSudoPass != "" {
		sudoArg := "-e ansible_sudo_pass=" + m.cluster.config.AnsibleSudoPass
		args = append(args, sudoArg)
	}

	switch m.action {
	case createAction:
		if err := m.playDeployMCGW(args); err != nil {
			return err
		}
		if err := m.playMCInstancesConfig(args); err != nil {
			return err
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
		if err := m.playMCSetupControllerGWRoutes(args); err != nil {
			return err
		}
		if err := m.playMCSetupRemoteGWRoutes(args); err != nil {
			return err
		}
	case updateAction:
		if err := m.playDeployMCGW(args); err != nil {
			return err
		}
		if err := m.playMCSetupControllerGWRoutes(args); err != nil {
			return err
		}
		if err := m.playMCInstancesConfig(args); err != nil {
			return err
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
	case deleteAction:

		//best effort cleaning up
		_ = m.playMCContrailCleanup(args)
		_ = m.playMCGatewayCleanup(args)
		return nil
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

	if err := m.createContrailCommonFile(workDir); err != nil {
		return err
	}

	if err := m.createGatewayCommonFile(workDir); err != nil {
		return err
	}

	//tor file
	if err := m.createTORCommonFile(workDir); err != nil {
		return err
	}
	return nil

}

func (m *multiCloudProvisioner) createClusterTopologyFile(workDir string) error {

	// this logic needs to be changed when cloud_provisioner object is added
	cloudInfo := make(map[string]string)
	for _, cloudRef := range m.clusterData.cloudInfo {
		if cloudRef.Type == "private" {
			cloudInfo[cloudRef.Type] = cloudRef.UUID
			continue
		}
		cloudInfo["public"] = cloudRef.UUID
	}

	onPremCloudID, ok := cloudInfo["private"]
	if !ok {
		return errors.New("private cloud object is not created and attached to cluster object")
	}
	onPremTopoFile := cloud.GetTopoFile(onPremCloudID)

	if _, err := os.Stat(onPremTopoFile); err == nil {
		onPremTopoContent, err := ioutil.ReadFile(onPremTopoFile)
		if err != nil {
			return err
		}
		err = cloud.AppendToFile(m.getClusterTopoFile(workDir), onPremTopoContent)
		if err != nil {
			return err
		}
	} else {
		return errors.New("topology file is not created for onprem cloud")
	}

	publicCloudID, ok := cloudInfo["public"]
	if !ok {
		return errors.New("public cloud object is not created and attached to cluster object")
	}
	publicTopoFile := cloud.GetTopoFile(publicCloudID)

	if _, err := os.Stat(publicTopoFile); err == nil {
		publicTopoContent, err := ioutil.ReadFile(publicTopoFile)
		if err != nil {
			return err
		}
		err = cloud.AppendToFile(m.getClusterTopoFile(workDir), publicTopoContent)
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
	for _, cloud := range m.clusterData.cloudInfo {
		if cloud.Type != "private" {
			publicCloudID = cloud.UUID
			break
		}
	}

	if publicCloudID == "" {
		return errors.New("public cloud object is not created and attached to cluster object")
	}

	secretFile := cloud.GetSecretFile(publicCloudID)
	if _, err := os.Stat(secretFile); err == nil {
		secretFileContent, err := ioutil.ReadFile(secretFile)
		if err != nil {
			return err
		}
		err = cloud.AppendToFile(m.getClusterSecretFile(workDir), secretFileContent)
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
	context := pongo2.Context{
		"cluster":                   m.clusterData.clusterInfo,
		"k8sCluster":                m.clusterData.getK8sClusterInfo(),
		"defaultSSHUser":            SSHUser,
		"defaultSSHPassword":        SSHPassword,
		"defaultSSHKey":             SSHKey,
		"defaultContrailUser":       defaultContrailUser,
		"defaultContrailPassword":   defaultContrailPassword,
		"defaultContrailConfigPort": defaultContrailConfigPort,
		"defaultContrailTenant":     defaultContrailTenant,
	}
	content, err := m.applyTemplate(m.getContrailCommonTemplate(), context)
	if err != nil {
		return err
	}

	err = m.writeToFile(destination, content)
	if err != nil {
		return err
	}
	m.log.Info("Created contrail/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createGatewayCommonFile(destination string) error {
	m.log.Info("Creating gateway/common.yml input file for multi-cloud deployer")
	context := pongo2.Context{
		"cluster":               m.clusterData.clusterInfo,
		"pathConfig":            pathConfig,
		"upgradeKernel":         upgradeKernel,
		"kernelVersion":         kernelVersion,
		"asn":                   asn,
		"vpnLoNetwork":          vpnLoNetwork,
		"vpnNetwork":            vpnNetwork,
		"openvpnPort":           openvpnPort,
		"bfdInterval":           bfdInterval,
		"bfdMultiplier":         bfdMultiplier,
		"bfdIntervalMultihop":   bfdIntervalMultihop,
		"bfdMultiplierMultihop": bfdMultiplierMultihop,
		"bgpSecret":             bgpSecret,
	}
	content, err := m.applyTemplate(m.getGatewayCommonTemplate(), context)
	if err != nil {
		return err
	}

	err = m.writeToFile(destination, content)
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
	content, err := m.applyTemplate(m.getTORCommonTemplate(), context)
	if err != nil {
		return err
	}

	err = m.writeToFile(destination, content)
	if err != nil {
		return err
	}
	m.log.Info("Created tor/common.yml input file for multi-cloud deployer")
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

// use topology constant from cloud pkg
func (m *multiCloudProvisioner) getClusterTopoFile(workDir string) string {
	return filepath.Join(workDir, "topology.yml")
}

func (m *multiCloudProvisioner) getTFStateFile() string {

	for _, c := range m.clusterData.cloudInfo {
		if c.Type != "private" {
			return cloud.GetTFStateFile(c.UUID)
		}
	}
	return ""
}

func (m *multiCloudProvisioner) getMCInventoryFile(workDir string) string {
	return filepath.Join(workDir, "inventories/inventory.yml")
}

func (m *multiCloudProvisioner) getClusterSecretFile(workDir string) string {
	return filepath.Join(workDir, "secret.yml")
}

func (m *multiCloudProvisioner) getContrailCommonFile(workDir string) string {
	return filepath.Join(workDir, "contrail/common.yml")
}

func (m *multiCloudProvisioner) getGatewayCommonFile(workDir string) string {
	return filepath.Join(workDir, "gateway/common.yml")
}

func (m *multiCloudProvisioner) getTORCommonFile(workDir string) string {
	return filepath.Join(workDir, "tor/common.yml")
}

func (m *multiCloudProvisioner) getMCWorkingDir(clusterWorkDir string) string {
	return filepath.Join(clusterWorkDir, "multi-cloud")
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
	ansibleArgs = append(ansibleArgs, append(skipRoleArgs, defaultMCKubernetesProvPlay)...)
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
