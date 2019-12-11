package cluster

import (
	"bytes"
	"context"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/fileutil/template"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/Juniper/asf/pkg/retry"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	defaultContrailCommonTemplate = "contrail_common.tmpl"
	defaultGatewayCommonTemplate  = "gateway_common.tmpl"
	defaultTORCommonTemplate      = "tor_common.tmpl"

	mcWorkDir                 = "multi-cloud"
	mcAnsibleRepo             = "contrail-multi-cloud"
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
	bgpMCRoutesForController  = "false"
	debugMCRoutes             = "false"
	torBGPSecret              = "contrail_secret"
	torOSPFSecret             = "contrail_secret"

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

type contrailCommon struct {
	User            string                    `yaml:"contrail_user"`
	Password        string                    `yaml:"contrail_password"`
	Port            string                    `yaml:"contrail_port"`
	Tenant          string                    `yaml:"contrail_tenant"`
	Configuration   contrailConfiguration     `yaml:"contrail_configuration"`
	ProviderConfigs map[string]providerConfig `yaml:"provider_config"`
}

type providerConfig struct {
	SSHUser        string `yaml:"ssh_user"`
	SSHPassword    string `yaml:"ssh_pwd"`
	SSHPublicKey   string `yaml:"ssh_public_key,omitempty"`
	DomainSuffix   string `yaml:"domainsuffix,omitempty"`
	NTPServer      string `yaml:"ntpserver"`
	ManageEtcHosts bool   `yaml:"manage_etc_hosts"`
}

type contrailConfiguration struct {
	Version               string `yaml:"CONTRAIL_VERSION,omitempty"`
	EncapsulationPriority string `yaml:"ENCAP_PRIORITY"`
	CloudOrchestrator     string `yaml:"CLOUD_ORCHESTRATOR"`
	VrouterEncryption     string `yaml:"VROUTER_ENCRYPTION,omitempty"`
}

type gatewayCommon struct {
	ConfigPath                string `yaml:"PATH_CONFIG"`
	SSLConfigLocalPath        string `yaml:"PATH_SSL_CONFIG_LOCAL"`
	SSLConfigPath             string `yaml:"PATH_SSL_CONFIG"`
	OpenVPNConfigPath         string `yaml:"PATH_OPENVPN_CONFIG"`
	BirdConfigPath            string `yaml:"PATH_BIRD_CONFIG"`
	StrongSwanConfigPath      string `yaml:"PATH_STRONGSWAN_CONFIG"`
	VRRPConfigPath            string `yaml:"PATH_VRRP_CONFIG"`
	AWSConfigPath             string `yaml:"PATH_AWS_CONFIG"`
	InterfaceConfigPath       string `yaml:"PATH_INTERFACE_CONFIG"`
	FwConfigPath              string `yaml:"PATH_FW_CONFIG"`
	GCPConfigPath             string `yaml:"PATH_GCP_CONFIG"`
	SecretConfigPath          string `yaml:"PATH_SECRET_CONFIG"`
	VxLanConfigPath           string `yaml:"PATH_VXLAN_CONFIG"`
	ContainerRegistry         string `yaml:"CONTAINER_REGISTRY"`
	ContrailMulticloudVersion string `yaml:"CONTRAIL_MULTICLOUD_VERSION,omitempty"`
	UpgradeKernel             string `yaml:"UPGRADE_KERNEL"`
	KernelVersion             string `yaml:"KERNEL_VERSION"`
	AutonomousSystem          string `yaml:"AS"`
	VPNLoNetwork              string `yaml:"vpn_lo_network"`
	VPNNetwork                string `yaml:"vpn_network"`
	OpenVPNPort               string `yaml:"openvpn_port"`
	BfdInterval               string `yaml:"bfd_interval"`
	BfdMultiplier             string `yaml:"bfd_multiplier"`
	BfdIntervalMultihop       string `yaml:"bfd_interval_multihop"`
	BfdMultiplierMultihop     string `yaml:"bfd_multiplier_multihop"`
	CoreBgpSecret             string `yaml:"core_bgp_secret"`
	RegistryPrivateInsecure   string `yaml:"registry_private_insecure"`
}

type torCommon struct {
	BGPSecret                        string `yaml:"tor_bgp_secret"`
	OspfSecret                       string `yaml:"tor_ospf_secret"`
	DebugMulticloudRoutes            string `yaml:"DEBUG_MULTI_CLOUD_ROUTES"`
	BGPMulticloudRoutesForController string `yaml:"BGP_MULTI_CLOUD_ROUTES_FOR_CONTROLLER"`
}

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
			return errors.Wrapf(err, "%s failed", pa)
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
		return errors.Wrapf(err, "%s failed", pa)
	} else if updated {
		return nil
	}

	if m.action == createAction {
		m.reportStatus(statusCreateProgress)
	} else {
		m.reportStatus(statusUpdateProgress)
	}

	if err := m.manageMCCluster(); err != nil {
		if m.action == createAction {
			m.reportStatus(statusCreateFailed)
		} else {
			m.reportStatus(statusUpdateFailed)
		}
		return errors.Wrapf(err, "%s failed", pa)
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
	return m.clusterData.ClusterInfo.ProvisioningState == statusNoState &&
		m.clusterData.ClusterInfo.ProvisioningAction == deleteCloud
}

func (m *multiCloudProvisioner) reportStatus(status string) {
	m.Reporter.ReportStatus(context.Background(), map[string]interface{}{statusField: status}, defaultResource)
}

func (m *multiCloudProvisioner) deleteMCCluster() error {
	if err := m.createFiles(); err != nil {
		return err
	}
	if err := m.manageSSHAgent(); err != nil {
		return err
	}
	if err := m.cleanupProvisioning(); err != nil {
		return err
	}
	if m.action != deleteAction {
		if err := m.removeCloudRefFromCluster(); err != nil {
			return err
		}
	}
	// nolint: errcheck
	defer m.stopSSHAgent()
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
	if err := m.createContrailCommonFile(); err != nil {
		return err
	}
	if err := m.createGatewayCommonFile(); err != nil {
		return err
	}
	return m.createTORCommonFile()
}

// nolint: gocyclo
func (m *multiCloudProvisioner) createClusterTopologyFile(workDir string) error {
	pubCloudID, pvtCloudID, err := m.publicAndPrivateCloudIDs()
	if err != nil {
		return err
	}
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
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return fileutil.AppendToFile(dest, content, defaultFilePermRWOnly)
}

func (m *multiCloudProvisioner) createClusterSecretFile() error {
	pubCloudKeyPair, err := m.getPublicCloudKeyPair()
	if err != nil {
		return errors.Wrap(err, "failed to to get public Cloud keypair")
	}

	sfc := cloud.SecretFileConfig{}
	if err = sfc.Update(pubCloudKeyPair); err != nil {
		return errors.Wrap(err, "failed to to update secret file config")
	}

	err = template.ApplyToFile(
		m.secretTemplatePath(), m.getClusterSecretFile(),
		pongo2.Context{
			"secret":  sfc,
			"cluster": m.clusterData.ClusterInfo,
			"tag":     m.clusterData.ClusterInfo.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG"),
		},
		defaultFilePermRWOnly,
	)
	return errors.Wrap(err, "failed to to create secret file")
}

func (m *multiCloudProvisioner) isMCUpdated() (bool, error) {
	if m.clusterData.ClusterInfo.ProvisioningState == statusNoState {
		return false, nil
	}
	if _, err := os.Stat(m.getClusterTopoFile(m.workDir)); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		m.reportStatus(statusUpdateFailed)
		return false, errors.Wrapf(err, "couldn't read old topology file: %s", m.getClusterTopoFile(m.workDir))
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

	if err := m.verifyCloudStatus(); err != nil {
		return err
	}
	if err := m.createFiles(); err != nil {
		return err
	}
	if err := m.manageSSHAgent(); err != nil {
		return err
	}
	return m.provision()
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

func (m *multiCloudProvisioner) publicAndPrivateCloudIDs() (string, string, error) {
	pubCloudID, err := m.publicCloudID()
	if err != nil {
		return "", "", err
	}
	pvtCloudID, err := m.privateCloudID()
	if err != nil {
		return "", "", err
	}

	return pubCloudID, pvtCloudID, nil
}

func (m *multiCloudProvisioner) privateCloudID() (string, error) {
	for _, cloudRef := range m.clusterData.CloudInfo {
		for _, p := range cloudRef.CloudProviders {
			if p.Type == onPrem {
				return cloudRef.UUID, nil
			}
		}
	}
	return "", errors.New("no private cloud in Cluster")
}

func (m *multiCloudProvisioner) publicCloudID() (string, error) {
	for _, cloudRef := range m.clusterData.CloudInfo {
		for _, p := range cloudRef.CloudProviders {
			if p.Type == aws || p.Type == azure || p.Type == gcp {
				return cloudRef.UUID, nil
			}
		}
	}
	return "", errors.New("no public cloud in Cluster")
}

func (m *multiCloudProvisioner) provision() error {
	cmd := multicloudCLI
	args := []string{
		"all", "provision", "--topology", m.getClusterTopoFile(m.workDir),
		"--secret", m.getClusterSecretFile(), "--tf_state", m.getTFStateFile(),
	}
	if m.clusterData.ClusterInfo.ProvisioningAction == updateCloud {
		args = append(args, "--update")
	}
	if m.contrailAnsibleDeployer.ansibleClient.IsTest() {
		return m.mockCLI(strings.Join(append([]string{cmd}, args...), " "))
	}
	vars := []string{
		fmt.Sprintf("SSH_AUTH_SOCK=%s", os.Getenv("SSH_AUTH_SOCK")),
		fmt.Sprintf("SSH_AGENT_PID=%s", os.Getenv("SSH_AGENT_PID")),
	}
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
		filepath.Join(m.contrailAnsibleDeployer.ansibleClient.WorkingDirectory(), "executed_cmd.yml"),
		content,
		defaultFilePermRWOnly,
	)
}

func (m *multiCloudProvisioner) cleanupProvisioning() error {
	cmd := multicloudCLI
	mcRepoDir := m.getMCDeployerRepoDir()
	args := []string{
		"all", "clean", "--topology", m.getClusterTopoFile(m.workDir),
		"--secret", m.getClusterSecretFile(), "--tf_state", m.getTFStateFile(),
		"--state", filepath.Join(mcRepoDir, "state.yml"),
	}
	if m.contrailAnsibleDeployer.ansibleClient.IsTest() {
		return m.mockCLI(strings.Join(append([]string{cmd}, args...), " "))
	}
	// TODO: Change inventory path after specifying work dir during provisioning.
	return osutil.ExecCmdAndWait(m.Reporter, cmd, args, mcRepoDir)
}

func (m *multiCloudProvisioner) createContrailCommonFile() error {
	m.Log.Info("Creating contrail/common.yml input file for multi-cloud deployer")

	encapPriority := m.clusterData.ClusterInfo.EncapPriority
	if encapPriority == "" {
		encapPriority = "MPLSoGRE,MPLSoUDP,VXLAN"
	}

	// TODO: Check if not nil and if is not empty value
	vrouterEncryption := m.clusterData.ClusterInfo.ContrailConfiguration.GetValue("VROUTER_ENCRYPTION")

	// TODO: Check if not nil and if is not empty value
	contrailVersion := m.clusterData.ClusterInfo.ContrailVersion
	if contrailVersion == "" {
		contrailVersion = m.clusterData.ClusterInfo.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG")
	}

	common := contrailCommon{
		User:     defaultContrailUser,
		Password: m.getContrailPassword(),
		Tenant:   defaultContrailTenant,
		Port:     defaultContrailConfigPort,
		Configuration: contrailConfiguration{
			CloudOrchestrator:     "kubernetes",
			VrouterEncryption:     vrouterEncryption,
			EncapsulationPriority: encapPriority,
			Version:               contrailVersion,
		},
		ProviderConfigs: map[string]providerConfig{
			"bms": providerConfig{
				SSHUser:      m.clusterData.DefaultSSHKey,
				SSHPassword:  m.clusterData.DefaultSSHPassword,
				SSHPublicKey: m.clusterData.DefaultSSHKey,
				NTPServer:    m.clusterData.ClusterInfo.NTPServer,
				DomainSuffix: m.clusterData.ClusterInfo.DomainSuffix,
			},
		},
	}

	data, err := yaml.Marshal(common)
	if err != nil {
		return errors.Wrap(err, "cannot marshal contrail common yml file")
	}

	if err = fileutil.WriteToFile(m.getContrailCommonFile(), data, defaultFilePermRWOnly); err != nil {
		return errors.Wrap(err, "cannot create contrail common yml file")
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

	if err := template.ApplyToFile(
		m.gatewayCommonTemplatePath(),
		destination,
		pContext,
		defaultFilePermRWOnly,
	); err != nil {
		return err
	}
	m.Log.Info("Created gateway/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createGatewayCommonFile() error {
	m.Log.Info("Creating gateway/common.yml input file for multi-cloud deployer")
	pathConfigVar := "{{ PATH_CONFIG }}"
	// TODO: Czy tu musi być ten {{ PATH_CONFIG }}. Askuj Szymoneł
	common := gatewayCommon{
		ConfigPath:           pathConfig,
		SSLConfigPath:        filepath.Join(pathConfigVar, "ssl"),
		OpenVPNConfigPath:    filepath.Join(pathConfigVar, "openvpn"),
		BirdConfigPath:       filepath.Join(pathConfigVar, "bird"),
		StrongSwanConfigPath: filepath.Join(pathConfigVar, "strongswan"),
		VRRPConfigPath:       filepath.Join(pathConfigVar, "vrrp"),
		AWSConfigPath:        filepath.Join(pathConfigVar, "aws"),
		InterfaceConfigPath:  "/etc/network/interfaces.d",
		FwConfigPath:         filepath.Join(pathConfigVar, "firewall"),
		GCPConfigPath:        filepath.Join(pathConfigVar, "gcp"),
		SecretConfigPath:     filepath.Join(pathConfigVar, "secret"),
		VxLanConfigPath:      filepath.Join(pathConfigVar, "vxlan"),
		ContainerRegistry:    m.clusterData.ClusterInfo.ContainerRegistry,
	}

	data, err := yaml.Marshal(common)
	if err != nil {
		return errors.Wrap(err, "cannot marshal tor common yml file")
	}

	if err = fileutil.WriteToFile(m.getGatewayCommonFile(), data, defaultFilePermRWOnly); err != nil {
		return errors.Wrap(err, "cannot create tor common yml file")
	}
	m.Lo
	m.Log.Info("Created gateway/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createTORCommonFile() error {
	m.Log.Info("Creating tor/common.yml input file for multi-cloud deployer")

	common := torCommon{
		BGPMulticloudRoutesForController: bgpMCRoutesForController,
		BGPSecret:                        torBGPSecret,
		DebugMulticloudRoutes:            debugMCRoutes,
		OspfSecret:                       torOSPFSecret,
	}
	data, err := yaml.Marshal(common)
	if err != nil {
		return errors.Wrap(err, "cannot marshal tor common yml file")
	}

	if err = fileutil.WriteToFile(m.getTORCommonFile(), data, defaultFilePermRWOnly); err != nil {
		return errors.Wrap(err, "cannot create tor common yml file")
	}
	m.Log.Info("Created tor/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) removeVulnerableFiles() error {
	removeErr := osutil.ForceRemoveFiles(m.filesToRemove(), m.Log)
	return errors.Wrap(removeErr, "failed to remove credential files")
}

func (m *multiCloudProvisioner) filesToRemove() []string {
	kfd := services.NewKeyFileDefaults()

	f := []string{
		m.getMCInventoryFile(m.getMCDeployerRepoDir()),
		m.getClusterSecretFile(),
		kfd.GetAzureSubscriptionIDPath(),
		kfd.GetAzureClientIDPath(),
		kfd.GetAzureClientSecretPath(),
		kfd.GetAzureTenantIDPath(),
		kfd.GetGoogleAccountPath(),
		kfd.GetAWSAccessPath(),
		kfd.GetAWSSecretPath(),
	}
	return f
}

// unsetMulticloudProvisioningFlag unsets cloud.is_multicloud_provisioning flag.
// See pkg/cloud.Cloud.removeVulnerableFiles() comment regarding cloud.is_multicloud_provisioning flag purpose.
func (m *multiCloudProvisioner) unsetMulticloudProvisioningFlag() error {
	pubCloudID, err := m.publicCloudID()
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
	pubCloudID, err := m.publicCloudID()
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
	if strings.ToLower(m.clusterData.ClusterInfo.Orchestrator) == openstack {
		openStackClusterInfo := m.clusterData.GetOpenstackClusterInfo()
		if o := openStackClusterInfo.KollaPasswords.GetValue("keystone_admin_password"); o != "" {
			return o
		}
	}
	return defaultContrailPassword
}
