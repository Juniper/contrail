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
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

const (
	mcWorkDir              = "multi-cloud"
	mcRepository           = "contrail-multi-cloud"
	defaultMCInventoryFile = "inventories/inventory.yml"
	defaultTopologyFile    = "topology.yml"
	defaultSecretFile      = "secret.yml"

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
	clusterUUID string
	clusterProvisioningState string
	mcAction string
	clouds []cloudRef
	workDir string
	authorizedRegistry *cloud.AuthorizedRegistry
	reporter *report.Reporter
	log *logrus.Entry
	exec CommandExecutor
	apiServer *client.HTTP
	test bool
}

// Deploy performs Multicloud provisioning.
func (m *multiCloudProvisioner) Deploy() error {
	m.reportStatus(statusUpdateProgress)
	deployErr := m.deploy()
	if deployErr != nil {
		m.reportStatus(statusUpdateFailed)
		m.log.Errorf("Multi-Cloud provisioning failed with error: %v", deployErr)
	}
	m.reportStatus(statusUpdated)
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
	if m.clusterProvisioningState != statusNoState {
		return nil
	}

	if m.mcAction == deleteAction {
		if err := m.deleteMCCluster(); err != nil {
			return err
		}
		return nil
	}

	if updated, err := m.isMCUpdated(); err != nil {
		return err
	} else if updated {
		return nil
	}

	return m.manageMCCluster()
}

func (m *multiCloudProvisioner) reportStatus(status string) {
	m.reporter.ReportStatus(context.Background(), map[string]interface{}{statusField: status}, defaultResource)
}

func (m *multiCloudProvisioner) deleteMCCluster() error {
	if err := m.createFiles(); err != nil {
		return err
	}
	if err := m.cleanupProvisioning(); err != nil {
		return err
	}
	if err := m.removeCloudRefFromCluster(); err != nil {
		return err
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
	return fileutil.AppendToFile(dest, content, defaultFilePermRWOnly)
}

type cloudRef struct {
	uuid string
	provider string
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

	// reg, err := cloud.NewAuthorizedRegistry(m.clusterData.ClusterInfo)
	// if err != nil {
	// 	return err
	// }
	// sfc.AuthorizedRegistries = []*cloud.AuthorizedRegistry{reg}

	marshaled, err := yaml.Marshal(sfc)
	if err != nil {
		return errors.Wrapf(err, "couldn't marshal secret to yaml for cluster %s", m.clusterUUID)
	}
	err = ioutil.WriteFile(m.getClusterSecretFile(), marshaled, defaultFilePermRWOnly)
	return errors.Wrapf(err, "couldn't create secret.yml file for cluster %s", m.clusterUUID)
}

func (m *multiCloudProvisioner) isMCUpdated() (bool, error) {
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
		m.log.Infof("%s topology file is already up-to-date", defaultResource)
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
	m.log.Debugf("Creating temperory topology at dir %s", tmpDir)

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
	m.log.Infof("Starting %s of contrail cluster: %s", m.mcAction, m.clusterUUID)

	if err := m.verifyCloudStatus(); err != nil {
		return err
	}
	if err := m.createFiles(); err != nil {
		return err
	}
	return m.provision()
}

func (m *multiCloudProvisioner) verifyCloudStatus() error {
	for _, cloudRef := range m.clouds {
		if err := m.waitForCloudStatusToBeUpdated(context.Background(), cloudRef.uuid); err != nil {
			return err
		}
	}
	return nil
}

func (m *multiCloudProvisioner) waitForCloudStatusToBeUpdated(ctx context.Context, cloudUUID string) error {
	return retry.Do(func() (retry bool, err error) {
		cloudResp, err := m.apiServer.GetCloud(ctx, &services.GetCloudRequest{ID: cloudUUID})
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
	}, retry.WithLog(m.log.Logger), retry.WithInterval(statusRetryInterval))
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
	for _, cloudRef := range m.clouds {
		if cloudRef.provider == onPrem {
			return cloudRef.uuid, nil
		}
	}
	return "", errors.New("no private cloud in Cluster")
}

func (m *multiCloudProvisioner) publicCloudID() (string, error) {
	for _, cloudRef := range m.clouds {
		if cloudRef.provider == aws || cloudRef.provider == azure || cloudRef.provider == gcp {
			return cloudRef.uuid, nil
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
	if m.mcAction == updateCloud {
		args = append(args, "--update")
	}
	agent, stopAgent, err := m.startSSHAgent()
	if err != nil {
		return errors.Wrap(err, "cannot start SSH Agent")
	}
	defer stopAgent()
	return m.exec.ExecCmdAndWait(m.reporter, cmd, args, m.getMCDeployerRepoDir(), agent.GetExportVars()...)
}

func (m *multiCloudProvisioner) cleanupProvisioning() error {
	cmd := multicloudCLI
	mcRepoDir := m.getMCDeployerRepoDir()
	args := []string{
		"all", "clean", "--topology", m.getClusterTopoFile(m.workDir),
		"--secret", m.getClusterSecretFile(), "--tf_state", m.getTFStateFile(),
		"--state", filepath.Join(mcRepoDir, "state.yml"),
	}
	agent, stopAgent, err := m.startSSHAgent()
	if err != nil {
		return err
	}
	defer stopAgent()
	// TODO: Change inventory path after specifying work dir during provisioning.
	return m.exec.ExecCmdAndWait(m.reporter, cmd, args, mcRepoDir, agent.GetExportVars()...)
}

func (m *multiCloudProvisioner) startSSHAgent() (*sshAgent, func(), error) {
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
			m.log.Errorf("cannot stop SSH Agent: %v", stopErr)
		}
	}

	if err = agent.Run(filepath.Join(kp.GetSSHKeyDirPath(), kp.GetName())); err != nil {
		stopAgent()
		return nil, nil, err
	}
	return agent, stopAgent, nil
}

func (m *multiCloudProvisioner) removeVulnerableFiles() error {
	removeErr := osutil.ForceRemoveFiles(m.filesToRemove(), m.log)
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

	if _, err = m.apiServer.UpdateCloud(
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

	m.log.WithField("public-cloud-id", pubCloudID).Debug("Successfully unset cloud.is_multicloud_provisioning flag")
	return nil
}

func (m *multiCloudProvisioner) removeCloudRefFromCluster() error {
	for _, cloudRef := range m.clouds {
		if _, err := m.apiServer.DeleteContrailClusterCloudRef(
			context.Background(), &services.DeleteContrailClusterCloudRefRequest{
				ID:                      m.clusterUUID,
				ContrailClusterCloudRef: &models.ContrailClusterCloudRef {
					UUID: cloudRef.uuid,
				},
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
	cloudObj, err := cloud.GetCloud(ctx, m.apiServer, pubCloudID)
	if err != nil {
		return nil, err
	}

	cloudUserObj, err := m.apiServer.GetCloudUser(
		ctx,
		&services.GetCloudUserRequest{
			ID: cloudObj.CloudUserRefs[0].UUID,
		},
	)
	if err != nil {
		return nil, err
	}

	credentialObj, err := m.apiServer.GetCredential(
		ctx,
		&services.GetCredentialRequest{
			ID: cloudUserObj.CloudUser.CredentialRefs[0].UUID,
		},
	)
	if err != nil {
		return nil, err
	}

	keypairObj, err := m.apiServer.GetKeypair(
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
	for _, cloudRef := range m.clouds {
		if cloudRef.provider != onPrem {
			return cloud.GetTFStateFile(cloudRef.uuid)
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
	return filepath.Join(defaultAnsibleRepoDir, mcRepository)
}
