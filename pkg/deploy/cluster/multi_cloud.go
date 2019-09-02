package cluster

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	defaultContrailCommonFile = "ansible-multicloud/contrail/common.yml"
	defaultTORCommonFile      = "ansible-multicloud/tor/common.yml"
	defaultGatewayCommonFile  = "ansible-multicloud/gateway/common.yml"
	defaultMCInventoryFile    = "inventories/inventory.yml"
	defaultTopologyFile       = "topology.yml"
	defaultSecretFile         = "secret.yml"
	defaultSecretTemplate     = "secret.tmpl"
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

	mcState = "state.yml"

	defaultMCGWDeployPlay        = "ansible-multicloud/gateway/playbooks/deploy_and_run_all.yml"
	defaultMCInstanceConfPlay    = "ansible-multicloud/contrail/playbooks/configure.yml"
	defaultMCKubernetesProvPlay  = "ansible-multicloud/contrail/playbooks/orchestrator.yml"
	defaultMCTORPlay             = "ansible-multicloud/tor/playbooks/deploy_and_run_all.yml"
	defaultMCDeployContrail      = "ansible-multicloud/contrail/playbooks/deploy.yml"
	defaultMCSetupContrailRoutes = "ansible-multicloud/contrail/playbooks/add_tunnel_routes.yml"
	defaultMCFixComputeDNS       = "ansible-multicloud/contrail/playbooks/fix_compute_dns.yml"
	defaultMCContrailCleanup     = "ansible-multicloud/contrail/playbooks/cleanup.yml"
	defaultMCGatewayCleanup      = "ansible-multicloud/gateway/playbooks/cleanup.yml"
	testTemplate                 = "./../../cloud/test_data/test_cmd.tmpl"

	openstack            = "openstack"
	setupRoutesPlayRetry = 3

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

// Deploy performs Multicloud provisioning.
func (m *multiCloudProvisioner) Deploy() error {
	deployErr := m.deploy()
	if deployErr != nil {
		m.Log.Errorf("Multi-Cloud provisioning failed with error: %v", deployErr)
	}
	if err := m.removeVulnerableFiles(); err != nil {
		return errors.Errorf("failed to delete vulnerable files: %v; deploy error (if any): %v", err, deployErr)
	}
	if err := m.unsetMulticloudProvisioningFlag(); err != nil {
		return errors.Errorf(
			"failed to unset cloud.is_multicloud_provisioning flag: %v; deploy error (if any): %v", err, deployErr,
		)
	}
	return deployErr
}

// nolint: gocyclo
func (m *multiCloudProvisioner) deploy() error {
	m.updateMCWorkDir()
	pa := m.clusterData.ClusterInfo.ProvisioningAction

	if m.isMCDeleteRequest() {
		m.reportStatus(statusUpdateProgress)
		if err := m.deleteMCCluster(); err != nil {
			m.reportStatus(statusUpdateFailed)
			m.Log.Errorf("%s failed with err: %s", pa, err)
			return err
		}
		m.reportStatus(statusUpdated)
		return nil
	}

	if pa != addCloud && pa != updateCloud {
		return nil
	}

	if pa == addCloud && !m.cluster.config.Test {
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
	}

	updated, err := m.isMCUpdated()
	if err != nil {
		m.Log.Errorf("%s failed with err: %s", pa, err)
		return err
	}
	if updated {
		return nil
	}

	if err = m.manageMCCluster(); err != nil {
		if m.action == createAction {
			m.reportStatus(statusCreateFailed)
		} else {
			m.reportStatus(statusUpdateFailed)
		}
		m.Log.Errorf("%s failed with err: %s", pa, err)
		return err
	}

	if m.action == createAction {
		m.reportStatus(statusCreated)
	} else {
		m.reportStatus(statusUpdated)
	}
	return nil
}

func (m *multiCloudProvisioner) reportStatus(status string) {
	m.Reporter.ReportStatus(context.Background(), map[string]interface{}{statusField: status}, defaultResource)
}

func (m *multiCloudProvisioner) isMCUpdated() (bool, error) {
	if m.clusterData.ClusterInfo.ProvisioningState == statusNoState {
		return false, nil
	}
	if _, err := os.Stat(m.getClusterTopoFile(m.workDir)); err != nil {
		m.Log.Errorf("couldn't load topology file: %s", m.getClusterTopoFile(m.workDir))
		return false, nil
	}
	ok, err := m.compareClusterTopologyFile()
	if err != nil {
		m.reportStatus(statusUpdateFailed)
		return true, err
	}
	if ok {
		m.Log.Infof("%s topology file is already up-to-date", defaultResource)
		return true, nil
	}
	return false, nil
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

	newTopology, err := ioutil.ReadFile(m.getClusterTopoFile(tmpDir))
	if err != nil {
		return false, err
	}

	oldTopology, err := ioutil.ReadFile(m.getClusterTopoFile(m.workDir))
	if err != nil {
		return false, err
	}
	return bytes.Equal(oldTopology, newTopology), nil
}

// nolint: gocyclo
func (m *multiCloudProvisioner) createClusterTopologyFile(workDir string) error {
	pubCloudID, pvtCloudID, err := m.getPubPvtCloudID()
	if err != nil {
		return err
	}
	// Remove old topology file
	topo := m.getClusterTopoFile(workDir)
	if err := os.Remove(topo); !os.IsNotExist(err) {
		return errors.Wrapf(err, "couldn't remove old topology file: %s", topo)
	}
	onPremTopoFile := cloud.GetTopoFile(pvtCloudID)
	if err := appendFile(onPremTopoFile, topo); err != nil {
		return errors.Wrap(err, "topology file is not created for onprem cloud")
	}

	publicTopoFile := cloud.GetTopoFile(pubCloudID)
	if err := appendFile(publicTopoFile, topo); err != nil {
		return errors.Wrap(err, "topology file is not created for public cloud")
	}

	return nil
}

func appendFile(src, dest string) error {
	if _, err := os.Stat(src); err != nil {
		return err
	}
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return fileutil.AppendToFile(dest, content, defaultFilePermRWOnly)
}

func (m *multiCloudProvisioner) manageMCCluster() error {
	m.Log.Infof("Starting %s of contrail cluster: %s",
		m.clusterData.ClusterInfo.ProvisioningAction,
		m.clusterData.ClusterInfo.FQName)

	if m.action == createAction {
		m.reportStatus(statusCreateProgress)
	} else {
		m.reportStatus(statusUpdateProgress)
	}

	if err := m.verifyCloudStatus(); err != nil {
		return err
	}

	if err := m.createFiles(); err != nil {
		return err
	}
	if err := m.manageSSHAgent(m.action); err != nil {
		return err
	}
	if err := m.runGenerateInventory(m.clusterData.ClusterInfo.ProvisioningAction); err != nil {
		return err
	}
	if err := m.mcPlayBook(); err != nil {
		return err
	}

	return nil
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

func waitForCloudStatusToBeUpdated(
	ctx context.Context, log *logrus.Entry, httpClient *client.HTTP, cloudUUID string,
) error {

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
	}, retry.WithLog(log.Logger), retry.WithInterval(statusRetryInterval))
}

func (m *multiCloudProvisioner) createFiles() error {
	if err := m.createClusterTopologyFile(m.workDir); err != nil {
		return err
	}
	if err := m.createClusterSecretFile(); err != nil {
		return err
	}
	if err := m.createContrailCommonFile(m.getContrailCommonFile()); err != nil {
		return err
	}
	if err := m.createGatewayCommonFile(m.getGatewayCommonFile()); err != nil {
		return err
	}
	if err := m.createTORCommonFile(m.getTORCommonFile()); err != nil {
		return err
	}
	return nil
}

func (m *multiCloudProvisioner) createClusterSecretFile() error {
	pubCloudID, _, err := m.getPubPvtCloudID()
	if err != nil {
		return errors.Wrap(err, "failed to to get public Cloud ID")
	}

	pubCloudProviders, err := m.providers()
	if err != nil {
		return errors.Wrap(err, "failed to to get public Cloud providers")
	}

	pubCloudKeyPair, err := m.getPublicCloudKeyPair()
	if err != nil {
		return errors.Wrap(err, "failed to to get public Cloud keypair")
	}

	sfc := cloud.SecretFileConfig{}
	if err = sfc.Update(pubCloudID, pubCloudProviders, pubCloudKeyPair); err != nil {
		return errors.Wrap(err, "failed to to update secret file config")
	}
	content, err := template.Apply(
		m.secretTemplatePath(),
		pongo2.Context{
			"secret": sfc,
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to to create secret file")
	}

	if err = fileutil.WriteToFile(m.getClusterSecretFile(), content, defaultFilePermRWOnly); err != nil {
		return err
	}

	authRegcontent, err := m.getAuthRegistryContent(m.clusterData.ClusterInfo)
	if err != nil {
		return err
	}
	return fileutil.AppendToFile(m.getClusterSecretFile(), authRegcontent, defaultFilePermRWOnly)
}

func (m *multiCloudProvisioner) createContrailCommonFile(destination string) error {
	m.Log.Info("Creating contrail/common.yml input file for multi-cloud deployer")
	contrailPassword := m.getContrailPassword()

	pContext := pongo2.Context{
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
	content, err := template.Apply(m.contrailCommonTemplatePath(), pContext)
	if err != nil {
		return err
	}

	if err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly); err != nil {
		return err
	}
	m.Log.Info("Created contrail/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createGatewayCommonFile(destination string) error {
	m.Log.Info("Creating gateway/common.yml input file for multi-cloud deployer")
	pContext := pongo2.Context{
		"cluster":    m.clusterData.ClusterInfo,
		"pathConfig": pathConfig,
		"bgpSecret":  bgpSecret,
	}
	content, err := template.Apply(m.gatewayCommonTemplatePath(), pContext)
	if err != nil {
		return err
	}

	if err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly); err != nil {
		return err
	}
	m.Log.Info("Created gateway/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createTORCommonFile(destination string) error {
	m.Log.Info("Creating tor/common.yml input file for multi-cloud deployer")
	pContext := pongo2.Context{
		"torBGPSecret":             torBGPSecret,
		"torOSPFSecret":            torOSPFSecret,
		"debugMCRoutes":            debugMCRoutes,
		"bgpMCRoutesForController": bgpMCRoutesForController,
	}

	content, err := template.Apply(m.torCommonTemplatePath(), pContext)
	if err != nil {
		return err
	}

	if err = fileutil.WriteToFile(destination, content, defaultFilePermRWOnly); err != nil {
		return err
	}
	m.Log.Info("Created tor/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) deleteMCCluster() error {
	// nolint: errcheck
	defer func() {
		if err := m.removeVulnerableFiles(); err != nil {
			m.Log.WithError(err).Error("Failed to remove vulnerable files post delete")
		}
	}()

	if err := m.createFiles(); err != nil {
		return err
	}
	if err := m.manageSSHAgent(updateAction); err != nil {
		return err
	}
	if err := m.runGenerateInventory(updateCloud); err != nil {
		return err
	}

	// nolint: errcheck
	_ = m.mcPlayBook()

	if m.action != deleteAction {
		if err := m.removeCloudRefFromCluster(); err != nil {
			return err
		}
	}
	// nolint: errcheck
	defer m.manageSSHAgent(deleteAction)
	// nolint: errcheck
	defer os.RemoveAll(m.workDir)
	return nil
}

func (m *multiCloudProvisioner) isMCDeleteRequest() bool {
	return m.clusterData.ClusterInfo.ProvisioningState == statusNoState &&
		m.clusterData.ClusterInfo.ProvisioningAction == deleteCloud
}

func (m *multiCloudProvisioner) runGenerateInventory(cloudAction string) error {
	cmd := cloud.GetGenInventoryCmd(cloud.GetMultiCloudRepodir())
	var args []string

	switch cloudAction {
	case addCloud:
		args = []string{
			"-t", m.getClusterTopoFile(m.workDir),
			"-s", m.getClusterSecretFile(),
			"-ts", m.getTFStateFile(),
		}
	case updateCloud:
		args = []string{
			"-t", m.getClusterTopoFile(m.workDir),
			"-s", m.getClusterSecretFile(),
			"-ts", m.getTFStateFile(),
			"--instate", mcState,
			"--outstate", mcState,
		}
	default:
		return fmt.Errorf("%s action not supported by generate inventory", cloudAction)
	}

	m.Log.Info("Generating inventory file multi-cloud provisioner")
	m.Log.Debugf("Command executed: %s %s", cmd, strings.Join(args, " "))

	if m.cluster.config.Test {
		return cloud.TestCmdHelper(cmd, args, m.workDir, testTemplate)
	}

	if err := osutil.ExecCmdAndWait(m.Reporter, cmd, args, cloud.GetMultiCloudRepodir()); err != nil {
		return err
	}

	// enabling openstack playbook only for test purposes
	// because its not yet supported by multi-cloud-deployer
	if m.isOrchestratorOpenstack() && m.cluster.config.Test {
		if err := m.appendOpenStackConfigToInventory(m.getMCInventoryFile(m.workDir)); err != nil {
			return err
		}
	}
	m.Log.Infof("Successfully generated inventory file")
	return nil
}

// nolint: gocyclo
func (m *multiCloudProvisioner) mcPlayBook() error {
	args := []string{"-i", m.getMCInventoryFile(m.getMCDeployerRepoDir())}
	if m.cluster.config.AnsibleSudoPass != "" {
		sudoArg := "-e ansible_sudo_pass=" + m.cluster.config.AnsibleSudoPass
		args = append(args, sudoArg)
	}

	switch m.clusterData.ClusterInfo.ProvisioningAction {
	case addCloud:
		if err := m.copySSHKeyPairToMC(); err != nil {
			return err
		}
		if err := m.playDeployMCGW(args); err != nil {
			return err
		}
		if err := m.playMCInstancesConfig(args); err != nil {
			return err
		}

		torExists, err := m.checkIfTORExists()
		if err != nil {
			return err
		}
		if torExists {
			if err = m.playTORConfig(args); err != nil {
				return err
			}
		}

		k8sArgs := append(args, fmt.Sprintf("-e orchestrator=%s", orchestratorKubernetes))
		if err := m.playMCK8SProvision(k8sArgs); err != nil {
			return err
		}

		contrailArgs := append(args, strings.Split("--limit all:!tors", " ")...)
		skipRoles := []string{"vrouter"}
		if err := m.playMCDeployContrail(contrailArgs, skipRoles); err != nil {
			return err
		}

		skipRoles = []string{"config_database", "config,control", "webui", "analytics", "analytics_database", "k8s"}
		if err := m.playMCDeployContrail(contrailArgs, skipRoles); err != nil {
			return err
		}

		for i := 0; i < setupRoutesPlayRetry; i++ {
			m.Log.Debugf("TRY %d of running setup controller gw route play", i+1)
			if err = m.playMCSetupControllerGWRoutes(args); err == nil {
				break
			}
			if i == setupRoutesPlayRetry-1 {
				return err
			}
		}

		for i := 0; i < setupRoutesPlayRetry; i++ {
			m.Log.Debugf("TRY %d of running setup remote gw route play", i+1)
			if err = m.playMCSetupRemoteGWRoutes(args); err == nil {
				break
			}
			if i == setupRoutesPlayRetry-1 {
				return err
			}
		}
		return m.playMCFixComputeDNS(args)

	case updateCloud:
		if err := m.copySSHKeyPairToMC(); err != nil {
			return err
		}
		if err := m.playDeployMCGW(args); err != nil {
			return err
		}

		for i := 0; i < setupRoutesPlayRetry; i++ {
			var err error
			m.Log.Debugf("TRY %d of running setup controller gw route play", i+1)
			if err = m.playMCSetupControllerGWRoutes(args); err == nil {
				break
			}
			if i == setupRoutesPlayRetry-1 {
				return err
			}
		}

		if err := m.playMCInstancesConfig(args); err != nil {
			return err
		}

		torExists, err := m.checkIfTORExists()
		if err != nil {
			return err
		}
		if torExists {
			if err = m.playTORConfig(args); err != nil {
				return err
			}
		}

		k8sArgs := append(args, fmt.Sprintf("-e orchestrator=%s", orchestratorKubernetes))
		if err := m.playMCK8SProvision(k8sArgs); err != nil {
			return err
		}

		contrailArgs := append(args, strings.Split("--limit all:!tors", " ")...)
		skipRoles := []string{"vrouter"}
		if err := m.playMCDeployContrail(contrailArgs, skipRoles); err != nil {
			return err
		}

		skipRoles = []string{"config_database", "config,control", "webui", "analytics", "analytics_database", "k8s"}
		if err := m.playMCDeployContrail(contrailArgs, skipRoles); err != nil {
			return err
		}

		for i := 0; i < setupRoutesPlayRetry; i++ {
			m.Log.Debugf("TRY %d of running setup remote gw route play", i+1)
			if err = m.playMCSetupRemoteGWRoutes(args); err == nil {
				break
			}
			if i == setupRoutesPlayRetry-1 {
				return err
			}
		}

		return m.playMCFixComputeDNS(args)

	case deleteCloud:
		//best effort cleaning up
		// nolint: errcheck
		_ = m.playMCContrailCleanup(args)
		// nolint: errcheck
		_ = m.playMCGatewayCleanup(args)

		return nil
	}

	return nil
}

func (m *multiCloudProvisioner) removeVulnerableFiles() error {
	f, filesErr := m.filesToRemove()
	removeErr := osutil.ForceRemoveFiles(f, m.Log)
	if filesErr != nil {
		return errors.Errorf(
			"failed to calculate paths of files to be deleted: %v ; other files remove error (if any): %v",
			filesErr,
			removeErr,
		)
	}
	return errors.Wrap(removeErr, "failed to remove credential files")
}

func (m *multiCloudProvisioner) filesToRemove() ([]string, error) {
	f := []string{
		m.getMCInventoryFile(m.getMCDeployerRepoDir()),
		m.getClusterSecretFile(),
	}
	kfd, err := services.NewKeyFileDefaults()
	if err != nil {
		return f, errors.Wrap(
			err,
			"failed to create KeyFileDefaults - "+
				"Secret file, Azure profile, Azure access token, Google account, AWS access, AWS secret "+
				"will not be deleted",
		)
	}
	f = append(
		f,
		kfd.GetAzureProfilePath(),
		kfd.GetAzureAccessTokenPath(),
		kfd.GetGoogleAccountPath(),
		kfd.GetAWSAccessPath(),
		kfd.GetAWSSecretPath(),
	)
	pubCloudID, _, err := m.getPubPvtCloudID()
	if err != nil {
		return f, errors.Wrap(
			err, "failed to retrieve public cloud UUID - Secret file, AWS access, AWS secret will not be deleted",
		)
	}
	f = append(f, cloud.GetSecretFile(pubCloudID))
	return f, nil
}

// unsetMulticloudProvisioningFlag unsets cloud.is_multicloud_provisioning flag.
// See pkg/cloud.Cloud.removeVulnerableFiles() comment regarding cloud.is_multicloud_provisioning flag purpose.
func (m *multiCloudProvisioner) unsetMulticloudProvisioningFlag() error {
	pubCloudID, _, err := m.getPubPvtCloudID()
	if err != nil {
		return errors.Wrap(err, "failed to to get public Cloud ID")
	}

	if _, err = m.cluster.APIServer.UpdateCloud(
		context.Background(),
		&services.UpdateCloudRequest{
			Cloud: &models.Cloud{
				UUID:                     pubCloudID,
				IsMulticloudProvisioning: false,
			},
		},
	); err != nil {
		return errors.Wrapf(err, "failed to update Cloud with UUID: %s", pubCloudID)
	}

	m.Log.WithField("public-cloud-id", pubCloudID).Debug("Successfully unset cloud.is_multicloud_provisioning flag")
	return nil
}

func (m *multiCloudProvisioner) removeCloudRefFromCluster() error {
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

func (m *multiCloudProvisioner) secretTemplatePath() string {
	return filepath.Join(m.getTemplateRoot(), defaultSecretTemplate)
}

func (m *multiCloudProvisioner) providers() ([]string, error) {
	pubCloudID, _, err := m.getPubPvtCloudID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to to get public Cloud ID")
	}

	cloudObj, err := cloud.GetCloud(context.Background(), m.cluster.APIServer, pubCloudID)
	if err != nil {
		return nil, err
	}

	return providerNames(cloudObj.GetCloudProviders()), nil
}

func providerNames(providers []*models.CloudProvider) []string {
	result := []string{}
	for _, provider := range providers {
		result = append(result, provider.Type)
	}
	return result
}

func (m *multiCloudProvisioner) getPublicCloudKeyPair() (*models.Keypair, error) {
	pubCloudID, _, err := m.getPubPvtCloudID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to to get public Cloud ID")
	}

	ctx := context.Background()
	cloudObj, err := cloud.GetCloud(ctx, m.cluster.APIServer, pubCloudID)
	if err != nil {
		return nil, err
	}

	cloudUserObj, err := m.cluster.APIServer.GetCloudUser(
		ctx,
		&services.GetCloudUserRequest{
			ID: cloudObj.CloudUserRefs[0].UUID,
		},
	)
	if err != nil {
		return nil, err
	}

	credentialObj, err := m.cluster.APIServer.GetCredential(
		ctx,
		&services.GetCredentialRequest{
			ID: cloudUserObj.CloudUser.CredentialRefs[0].UUID,
		},
	)
	if err != nil {
		return nil, err
	}

	keypairObj, err := m.cluster.APIServer.GetKeypair(
		ctx,
		&services.GetKeypairRequest{
			ID: credentialObj.Credential.KeypairRefs[0].UUID,
		},
	)
	if err != nil {
		return nil, err
	}

	return keypairObj.Keypair, nil
}

func (m *multiCloudProvisioner) getAuthRegistryContent(cluster *models.ContrailCluster) ([]byte, error) {
	pContext := pongo2.Context{
		"cluster": cluster,
		"tag":     cluster.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG"),
	}

	content, err := template.Apply(m.authRegistryTemplatePath(), pContext)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (m *multiCloudProvisioner) appendOpenStackConfigToInventory(destination string) error {
	m.Log.Info("Appending openstack config to inventory file")

	pContext := pongo2.Context{
		"openstackCluster": m.clusterData.GetOpenstackClusterInfo(),
	}
	content, err := template.Apply(m.openstackConfigTemplatePath(), pContext)
	if err != nil {
		return err
	}

	if err = fileutil.AppendToFile(destination, content, defaultFilePermRWOnly); err != nil {
		return err
	}
	m.Log.Info("Appended openstack config to inventory file")
	return nil
}

func (m *multiCloudProvisioner) checkIfTORExists() (bool, error) {
	_, pvtCloudID, err := m.getPubPvtCloudID()
	if err != nil {
		return false, err
	}

	tag, err := getTagOfPvtCloud(context.Background(), m.cluster.APIServer, pvtCloudID)
	if err != nil {
		return false, err
	}
	return tag.PhysicalRouterBackRefs != nil, nil

}

func getTagOfPvtCloud(ctx context.Context, client *client.HTTP, cloudID string) (*models.Tag, error) {
	cloudObj, err := cloud.GetCloud(ctx, client, cloudID)
	if err != nil {
		return nil, err
	}

	if cloudObj.CloudProviders != nil {
		cloudProvObj, err := client.GetCloudProvider(ctx,
			&services.GetCloudProviderRequest{
				ID: cloudObj.CloudProviders[0].UUID,
			},
		)
		if err != nil {
			return nil, err
		}

		if cloudProvObj.CloudProvider.CloudRegions != nil {
			cloudRegObj, err := client.GetCloudRegion(ctx,
				&services.GetCloudRegionRequest{
					ID: cloudProvObj.CloudProvider.CloudRegions[0].UUID,
				},
			)
			if err != nil {
				return nil, err
			}

			if cloudRegObj.CloudRegion.VirtualClouds != nil {
				vCloudObj, err := client.GetVirtualCloud(ctx,
					&services.GetVirtualCloudRequest{
						ID: cloudRegObj.CloudRegion.VirtualClouds[0].UUID,
					},
				)
				if err != nil {
					return nil, err
				}

				if vCloudObj.VirtualCloud.TagRefs != nil {
					tagObj, err := client.GetTag(ctx,
						&services.GetTagRequest{
							ID: vCloudObj.VirtualCloud.TagRefs[0].UUID,
						},
					)
					if err != nil {
						return nil, err
					}
					return tagObj.Tag, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("tag not found in cloud %s", cloudID)
}

func (m *multiCloudProvisioner) playTORConfig(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultMCTORPlay)
	return m.play(ansibleArgs)

}

func (m *multiCloudProvisioner) contrailCommonTemplatePath() string {
	return filepath.Join(m.getTemplateRoot(), defaultContrailCommonTemplate)
}

func (m *multiCloudProvisioner) authRegistryTemplatePath() string {
	return filepath.Join(m.getTemplateRoot(), defaultAuthRegistryTemplate)
}

func (m *multiCloudProvisioner) gatewayCommonTemplatePath() string {
	return filepath.Join(m.getTemplateRoot(), defaultGatewayCommonTemplate)
}

func (m *multiCloudProvisioner) torCommonTemplatePath() string {
	return filepath.Join(m.getTemplateRoot(), defaultTORCommonTemplate)
}

func (m *multiCloudProvisioner) openstackConfigTemplatePath() string {
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

func (m *multiCloudProvisioner) getClusterSecretFile() string {
	return filepath.Join(m.workDir, defaultSecretFile)
}

func (m *multiCloudProvisioner) getSSHAgentFile() string {
	return filepath.Join(m.workDir, defaultSSHAgentFile)
}

func (m *multiCloudProvisioner) getMCDeployerRepoDir() string {
	return filepath.Join(defaultAnsibleRepoDir, mcAnsibleRepo)
}

func (m *multiCloudProvisioner) getContrailCommonFile() string {
	return filepath.Join(m.workDir, defaultContrailCommonFile)
}

func (m *multiCloudProvisioner) getGatewayCommonFile() string {
	return filepath.Join(m.workDir, defaultGatewayCommonFile)
}

func (m *multiCloudProvisioner) getTORCommonFile() string {
	return filepath.Join(m.workDir, defaultTORCommonFile)
}

func (m *multiCloudProvisioner) getMCWorkingDir(clusterWorkDir string) string {
	return filepath.Join(clusterWorkDir, mcWorkDir)
}

func (m *multiCloudProvisioner) getContrailPassword() string {
	if m.isOrchestratorOpenstack() {
		openStackClusterInfo := m.clusterData.GetOpenstackClusterInfo()
		if o := openStackClusterInfo.KollaPasswords.GetValue("keystone_admin_password"); o != "" {
			return o
		}
	}
	return defaultContrailPassword
}

func (m *multiCloudProvisioner) isOrchestratorOpenstack() bool {
	return strings.ToLower(m.clusterData.ClusterInfo.Orchestrator) == openstack
}

func (m *multiCloudProvisioner) play(ansibleArgs []string) error {
	return m.playFromDirectory(m.getMCDeployerRepoDir(), ansibleArgs)
}

func (m *multiCloudProvisioner) playDeployMCGW(ansibleArgs []string) error {
	ansibleArgs = append(ansibleArgs, defaultMCGWDeployPlay)
	return m.play(ansibleArgs)
}

func (m *multiCloudProvisioner) playMCInstancesConfig(ansibleArgs []string) error {
	// configure compute node instances first
	computeArgs := append(ansibleArgs, strings.Split("--limit all:!gateways", " ")...)
	computeArgs = append(computeArgs, defaultMCInstanceConfPlay)
	if err := m.play(computeArgs); err != nil {
		return err
	}

	// now configure gateway node instances
	gatewayArgs := append(ansibleArgs, strings.Split("--limit all:!compute_nodes", " ")...)
	gatewayArgs = append(gatewayArgs, defaultMCInstanceConfPlay)
	return m.play(gatewayArgs)
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
