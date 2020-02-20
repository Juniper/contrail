package cluster

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/Juniper/asf/pkg/retry"
	"github.com/Juniper/contrail/pkg/ansible"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

// Multicloud related constants.
const (
	MCWorkDir              = "multi-cloud"
	mcRepositoryPath       = "/root/contrail-multi-cloud"
	DefaultMCInventoryFile = "inventories/inventory.yml"
	DefaultTopologyFile    = "topology.yml"
	DefaultSecretFile      = "secret.yml"

	addCloud    = "ADD_CLOUD"
	updateCloud = "UPDATE_CLOUD"
	deleteCloud = "DELETE_CLOUD"

	authSockFile = "agent"
	mcStateFile  = "state.yml"

	aws    = "aws"
	azure  = "azure"
	gcp    = "gcp"
	onPrem = "private"

	statusRetryInterval = 3 * time.Second
)

type multiCloudProvisioner struct {
	ContrailAnsibleDeployer
	workDir string
}

// Deploy performs Multicloud provisioning.
func (m *multiCloudProvisioner) Deploy() error {
	deployErr := m.deploy()
	if deployErr != nil {
		m.setStatusToFail()
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
	if m.clusterData.ClusterInfo.ProvisioningState != StatusNoState {
		return nil
	}

	m.updateMCWorkDir()
	pa := m.clusterData.ClusterInfo.ProvisioningAction

	if m.isMCDeleteRequest() {
		m.reportStatus(StatusUpdateProgress)
		if err := m.deleteMCCluster(); err != nil {
			return errors.Wrapf(err, "%s failed", pa)
		}
		m.reportStatus(StatusUpdated)
		return nil
	}

	if pa != addCloud && pa != updateCloud {
		return nil
	}

	if pa == addCloud {
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

	if err := m.waitForCloudsToBeUpdated(); err != nil {
		return err
	}

	if updated, err := m.isMCUpdated(); err != nil {
		return errors.Wrapf(err, "%s failed", pa)
	} else if updated {
		return nil
	}

	if m.action == CreateAction {
		m.reportStatus(StatusCreateProgress)
	} else {
		m.reportStatus(StatusUpdateProgress)
	}

	if err := m.manageMCCluster(); err != nil {
		return errors.Wrapf(err, "%s failed", pa)
	}

	if m.action == CreateAction {
		m.reportStatus(StatusCreated)
	} else {
		m.reportStatus(StatusUpdated)
	}
	return nil
}

func (m *multiCloudProvisioner) updateMCWorkDir() {
	m.workDir = m.getMCWorkingDir(m.getWorkingDir())
}

func (m *multiCloudProvisioner) isMCDeleteRequest() bool {
	return m.clusterData.ClusterInfo.ProvisioningState == StatusNoState &&
		m.clusterData.ClusterInfo.ProvisioningAction == deleteCloud
}

func (m *multiCloudProvisioner) reportStatus(status string) {
	m.Reporter.ReportStatus(context.Background(), map[string]interface{}{StatusField: status}, DefaultResource)
}

func (m *multiCloudProvisioner) deleteMCCluster() error {
	if err := m.createFiles(); err != nil {
		return err
	}
	if err := m.cleanupProvisioning(); err != nil {
		return err
	}
	if m.action != DeleteAction {
		if err := m.removeCloudRefFromCluster(); err != nil {
			return err
		}
	}
	// nolint: errcheck
	defer os.RemoveAll(m.workDir)
	return nil
}

func (m *multiCloudProvisioner) createFiles() error {
	if err := m.createClusterTopologyFile(m.workDir); err != nil {
		return err
	}
	return m.createClusterSecretFile()
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
	return fileutil.AppendToFile(dest, content, DefaultFilePermRWOnly)
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

	reg, err := cloud.NewAuthorizedRegistry(m.clusterData.ClusterInfo)
	if err != nil {
		return err
	}
	sfc.AuthorizedRegistries = []*cloud.AuthorizedRegistry{reg}

	marshaled, err := yaml.Marshal(sfc)
	if err != nil {
		return errors.Wrapf(err, "couldn't marshal secret to yaml for cluster %s", m.cluster.config.ClusterID)
	}
	err = ioutil.WriteFile(m.getClusterSecretFile(), marshaled, DefaultFilePermRWOnly)
	return errors.Wrapf(err, "couldn't create secret.yml file for cluster %s", m.cluster.config.ClusterID)
}

func (m *multiCloudProvisioner) isMCUpdated() (bool, error) {
	if _, err := os.Stat(m.getClusterTopoFile(m.workDir)); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errors.Wrapf(err, "couldn't read old topology file: %s", m.getClusterTopoFile(m.workDir))
	}
	if ok, err := m.compareClusterTopologyFile(); err != nil {
		return true, err
	} else if ok {
		m.Log.Infof("%s topology file is already up-to-date", DefaultResource)
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

	if err := m.createFiles(); err != nil {
		return err
	}
	return m.provision()
}

func (m *multiCloudProvisioner) waitForCloudsToBeUpdated() error {
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
		case StatusCreateProgress:
			return true, fmt.Errorf("waiting for create cloud %s to complete", cloudUUID)
		case StatusUpdateProgress:
			return true, fmt.Errorf("waiting for update cloud %s to complete", cloudUUID)
		case StatusNoState:
			return true, fmt.Errorf("waiting for cloud %s to complete processing", cloudUUID)
		case StatusCreated:
			return false, nil
		case StatusUpdated:
			return false, nil
		case StatusCreateFailed:
			return false, fmt.Errorf("cloud %s status has failed in creating", cloudUUID)
		case StatusUpdateFailed:
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
		"--secret", m.getClusterSecretFile(), "--tf_state", m.getTFStateFile(), "--state", m.getMCStateFile(),
	}
	if m.clusterData.ClusterInfo.ProvisioningAction == updateCloud {
		args = append(args, "--update")
	}
	agent, stopAgent, err := m.startSSHAgent(filepath.Join(m.workDir, authSockFile))
	if err != nil {
		return errors.Wrap(err, "cannot start SSH Agent")
	}
	defer stopAgent()
	return m.runProvisionActionInContainer(append([]string{cmd}, args...), agent)
}

func (m *multiCloudProvisioner) runProvisionActionInContainer(cmdWithArgs []string, agent *sshAgent) error {
	env := []string{fmt.Sprintf("%s=%s", sshAuthSock, agent.AuthenticationSocket())}

	imgRef, err := ansible.ImageReference(m.clusterData.ClusterInfo.ContainerRegistry,
		cloud.MultiCloudContainer, ansible.GetContrailVersion(m.clusterData.ClusterInfo, m.Log))
	if err != nil {
		return err
	}

	return m.containerPlayer.StartExecuteAndRemove(context.Background(), &ansible.ContainerParameters{
		ImageRef:            imgRef,
		ImageRefUsername:    m.clusterData.ClusterInfo.ContainerRegistryUsername,
		ImageRefPassword:    m.clusterData.ClusterInfo.ContainerRegistryPassword,
		HostVolumes:         m.getMCVolumes(agent),
		ForceContainerRecreate:        true,
		WorkingDirectory:    mcRepositoryPath,
		Env:                 env,
		ContainerPrefix:     "multicloud-deployer",
		HostNetwork:         true,
		OverwriteEntrypoint: true,
		RemoveContainer:     true,
		Privileged: true,
	}, cmdWithArgs)
}

// These volumes are loaded from host - not from Contrail Command container.
func (m *multiCloudProvisioner) getMCVolumes(agent *sshAgent) []ansible.Volume {
	keyPaths := services.NewKeyFileDefaults()
	return []ansible.Volume{
		{
			Source: m.workDir,
			Target: m.workDir,
		},
		{
			Source: m.getPublicCloudWorkDir(),
			Target: m.getPublicCloudWorkDir(),
		},
		{
			Source: keyPaths.KeyHomeDir,
			Target: keyPaths.KeyHomeDir,
		},
	}
}

func (m *multiCloudProvisioner) cleanupProvisioning() error {
	cmd := multicloudCLI
	args := []string{
		"all", "clean", "--topology", m.getClusterTopoFile(m.workDir),
		"--secret", m.getClusterSecretFile(), "--tf_state", m.getTFStateFile(),
		"--state", m.getMCStateFile(),
	}
	agent, stopAgent, err := m.startSSHAgent(filepath.Join(m.workDir, authSockFile))
	if err != nil {
		return err
	}
	defer stopAgent()
	return m.runProvisionActionInContainer(append([]string{cmd}, args...), agent)
}

func (m *multiCloudProvisioner) startSSHAgent(authSockDst string) (*sshAgent, func(), error) {
	// SSH Agent authentication socket must be located in the volume shared between
	// host and contrail_command container. Current container exec function does not
	// allow mounting volumes from container - not from host.
	kp, err := m.getPublicCloudKeyPair()
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot resolve public cloud keypair")
	}
	if kp.GetSSHKeyDirPath() == "" || kp.GetName() == "" {
		return nil, nil, errors.Errorf("credentials %s do not contain SSHKeyDirectory/SSHKeyName", kp.GetUUID())
	}

	agent := &sshAgent{}
	stopAgent := func() {
		if stopErr := agent.Stop(); stopErr != nil {
			m.Log.Errorf("cannot stop SSH Agent: %v", stopErr)
		}
	}

	if err = agent.Run(filepath.Join(kp.GetSSHKeyDirPath(), kp.GetName()), authSockDst); err != nil {
		stopAgent()
		return nil, nil, err
	}
	return agent, stopAgent, nil
}

func (m *multiCloudProvisioner) setStatusToFail() {
	if m.action == CreateAction {
		m.reportStatus(StatusCreateFailed)
	} else {
		m.reportStatus(StatusUpdateFailed)
	}
}

func (m *multiCloudProvisioner) removeVulnerableFiles() error {
	removeErr := osutil.ForceRemoveFiles(m.filesToRemove(), m.Log)
	return errors.Wrap(removeErr, "failed to remove credential files")
}

func (m *multiCloudProvisioner) filesToRemove() []string {
	kfd := services.NewKeyFileDefaults()

	// inventory.yml is another vulnerable file that should be deleted hovewer current
	// contrail-multi-cloud repository creates this file in the container which is removed
	// after performing an action.
	f := []string{
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

func (m *multiCloudProvisioner) getPublicCloudWorkDir() string {
	for _, c := range m.clusterData.CloudInfo {
		for _, prov := range c.CloudProviders {
			if prov.Type != onPrem {
				return cloud.GetCloudDir(c.UUID)
			}
		}
	}
	return ""
}

func (m *multiCloudProvisioner) getMCInventoryFile(workDir string) string {
	return filepath.Join(workDir, DefaultMCInventoryFile)
}

func (m *multiCloudProvisioner) getMCStateFile() string {
	return filepath.Join(m.workDir, mcStateFile)
}

// use topology constant from cloud pkg
func (m *multiCloudProvisioner) getClusterTopoFile(workDir string) string {
	return filepath.Join(workDir, DefaultTopologyFile)
}

func (m *multiCloudProvisioner) getClusterSecretFile() string {
	return filepath.Join(m.workDir, DefaultSecretFile)
}

func (m *multiCloudProvisioner) getMCWorkingDir(clusterWorkDir string) string {
	return filepath.Join(clusterWorkDir, MCWorkDir)
}
