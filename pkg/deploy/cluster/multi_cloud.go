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
	defaultContrailCommonTemplate = "contrail_common.tmpl"
	defaultGatewayCommonTemplate  = "gateway_common.tmpl"
	defaultTORCommonTemplate      = "tor_common.tmpl"
	defaultAuthRegistryTemplate   = "authorized_registries.tmpl"

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

	if updated, err := m.isMCUpdated(); err != nil {
		m.Log.Errorf("%s failed with err: %s", pa, err)
		return err
	} else if updated {
		return nil
	}

	if err := m.manageMCCluster(); err != nil {
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

func (m *multiCloudProvisioner) updateMCWorkDir() {
	m.workDir = m.getMCWorkingDir(m.getWorkingDir())
}

func (m *multiCloudProvisioner) isMCDeleteRequest() bool {
	if m.clusterData.ClusterInfo.ProvisioningState == statusNoState &&
		m.clusterData.ClusterInfo.ProvisioningAction == deleteCloud {
		return true
	}
	return false
}

func (m *multiCloudProvisioner) reportStatus(status string) {
	m.Reporter.ReportStatus(context.Background(), map[string]interface{}{statusField: status}, defaultResource)
}

func (m *multiCloudProvisioner) deleteMCCluster() error {
	if err := m.createFiles(); err != nil {
		return err
	}
	if err := m.manageSSHAgent(updateAction); err != nil {
		return err
	}
	if err := m.multiCloudCLICleanup(); err != nil {
		return err
	}
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

// nolint: gocyclo
func (m *multiCloudProvisioner) createClusterTopologyFile(workDir string) error {
	pubCloudID, pvtCloudID, err := m.getPubPvtCloudID()
	if err != nil {
		return err
	}
	// Remove old topology file
	topologyFile := m.getClusterTopoFile(workDir)
	if err = fileutil.CopyFile(cloud.GetTopoFile(pvtCloudID), topologyFile, true); err != nil {
		return errors.Wrap(err, "topology file is not created for onprem cloud")
	}
	if err = appendFile(cloud.GetTopoFile(pubCloudID), topologyFile); err != nil {
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

func (m *multiCloudProvisioner) createClusterSecretFile() error {
	pubCloudID, err := m.getPubCloudID()
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

func (m *multiCloudProvisioner) getAuthRegistryContent(cluster *models.ContrailCluster) ([]byte, error) {
	pContext := pongo2.Context{
		"cluster": cluster,
		"tag":     cluster.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG"),
	}

	return template.Apply(m.authRegistryTemplatePath(), pContext)
}

func (m *multiCloudProvisioner) isMCUpdated() (bool, error) {
	if m.clusterData.ClusterInfo.ProvisioningState == statusNoState {
		return false, nil
	}
	if _, err := os.Stat(m.getClusterTopoFile(m.workDir)); err != nil {
		m.Log.Errorf("couldn't load topology file: %s", m.getClusterTopoFile(m.workDir))
		return false, nil
	}
	if ok, err := m.compareClusterTopologyFile(); err != nil {
		m.reportStatus(statusUpdateFailed)
		return true, err
	} else if ok {
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
	if err := m.multiCloudCLIProvisioning(); err != nil {
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

func (m *multiCloudProvisioner) getPubPvtCloudID() (string, string, error) {
	pubCloudID, err := m.getPubCloudID()
	if err != nil {
		return "", "", err
	}
	pvtCloudID, err := m.getPvtCloudID()
	if err != nil {
		return "", "", err
	}

	return pubCloudID, pvtCloudID, nil
}

func (m *multiCloudProvisioner) getPvtCloudID() (string, error) {
	for _, cloudRef := range m.clusterData.CloudInfo {
		for _, p := range cloudRef.CloudProviders {
			if p.Type == onPrem {
				return cloudRef.UUID, nil
			}
		}
	}
	return "", errors.New("no private cloud in Cluster")
}

func (m *multiCloudProvisioner) getPubCloudID() (string, error) {
	for _, cloudRef := range m.clusterData.CloudInfo {
		for _, p := range cloudRef.CloudProviders {
			if p.Type == aws || p.Type == azure || p.Type == gcp {
				return cloudRef.UUID, nil
			}
		}
	}
	return "", errors.New("no public cloud in Cluster")
}
func (m *multiCloudProvisioner) providers() ([]string, error) {
	pubCloudID, err := m.getPubCloudID()
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

func (m *multiCloudProvisioner) multiCloudCLIProvisioning() error {
	topology := m.getClusterTopoFile(m.workDir)
	secret := m.getClusterSecretFile()
	stateTF := m.getTFStateFile()
	// TODO: Remove this after refactoring test framework.
	cmd := "deployer"
	args := []string{"all", "provision", "--topology", topology, "--secret", secret, "--tf_state", stateTF}
	if m.contrailAnsibleDeployer.ansibleClient.IsTest() {
		return m.mockCLI(strings.Join(append([]string{cmd}, args...), " "))
	}
	sshAuthSock := fmt.Sprintf("SSH_AUTH_SOCK=%s", os.Getenv("SSH_AUTH_SOCK"))
	sshAgentPid := fmt.Sprintf("SSH_AGENT_PID=%s", os.Getenv("SSH_AGENT_PID"))
	vars := []string{sshAuthSock, sshAgentPid}
	return osutil.ExecCmdAndWait(m.Reporter, cmd, args, m.getMCDeployerRepoDir(), vars...)
}

func (m *multiCloudProvisioner) mockCLI(cliCommand string) error {
	content, err := template.Apply("./test_data/test_mc_cli_command.tmpl", pongo2.Context{
		"command": cliCommand,
	})
	if err != nil {
		return err
	}

	return fileutil.AppendToFile(
		filepath.Join(m.contrailAnsibleDeployer.ansibleClient.GetWorkingDirectory(), "executed_cmd.yml"),
		content,
		defaultFilePermRWOnly,
	)
}

func (m *multiCloudProvisioner) multiCloudCLICleanup() error {
	topology := m.getClusterTopoFile(m.workDir)
	secret := m.getClusterSecretFile()
	stateTF := m.getTFStateFile()

	commands := []struct {
		cmd  string
		args []string
	}{{
		cmd:  "deployer",
		args: []string{"inventory", "create", "--topology", topology, "--secret", secret, "--tf_state", stateTF},
	}, {
		cmd:  "deployer",
		args: []string{"contrail", "clean"},
	}, {
		cmd:  "deployer",
		args: []string{"multicloud", "clean"},
	}}

	for _, c := range commands {
		// TODO: Remove this after refactoring test framework.
		if m.contrailAnsibleDeployer.ansibleClient.IsTest() {
			if err := m.mockCLI(strings.Join(append([]string{c.cmd}, c.args...), " ")); err != nil {
				return err
			}
			continue
		}
		// TODO: Change inventory path after specifying work dir during provisioning.
		if err := osutil.ExecCmdAndWait(m.Reporter, c.cmd, c.args, m.getMCDeployerRepoDir()); err != nil {
			return err
		}
	}
	return nil
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
	pubCloudID, err := m.getPubCloudID()
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
	pubCloudID, err := m.getPubCloudID()
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

func (m *multiCloudProvisioner) getPublicCloudKeyPair() (*models.Keypair, error) {
	pubCloudID, err := m.getPubCloudID()
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

func (m *multiCloudProvisioner) secretTemplatePath() string {
	return filepath.Join(m.getTemplateRoot(), defaultSecretTemplate)
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
	if strings.ToLower(m.clusterData.ClusterInfo.Orchestrator) == openstack {
		return true
	}
	return false
}
