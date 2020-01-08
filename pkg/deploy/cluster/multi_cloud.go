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
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/Juniper/asf/pkg/retry"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

type multiCloudAction string

const (
	mcWorkDir              = "multi-cloud"
	mcRepository           = "contrail-multi-cloud"
	defaultMCInventoryFile = "inventories/inventory.yml"
	defaultTopologyFile    = "topology.yml"
	defaultSecretFile      = "secret.yml"

	createMC multiCloudAction = "CREATE_CLOUD"
	updateMC multiCloudAction = "UPDATE_CLOUD"
	deleteMC multiCloudAction = "DELETE_CLOUD"

	aws    = "aws"
	azure  = "azure"
	gcp    = "gcp"
	onPrem = "private"

	statusRetryInterval = 3 * time.Second
)

// MultiCloudProvisioner is a structure for performing MultiCloud tasks.
type MultiCloudProvisioner struct {
	clusterUUID            string
	isClusterSetToNonState bool
	action                 multiCloudAction
	publicCloudID          string
	privateCloudID         string
	workDir                string
	authorizedRegistry     *cloud.AuthorizedRegistry
	reporter               *report.Reporter
	log                    *logrus.Entry
	exec                   CommandExecutor
	apiServer              *client.HTTP
	test                   bool
}

// NewMultiCloudProvisioner returns new MultiCloudProvisioner
func NewMultiCloudProvisioner(
	clusterID, logFile, clusterWorkRoot string, apiServer *client.HTTP, isTest bool, exec CommandExecutor,
) (*MultiCloudProvisioner, error) {
	resp, err := apiServer.GetContrailCluster(context.Background(), &services.GetContrailClusterRequest{
		ID: clusterID,
	})
	if err != nil {
		return nil, err
	}

	provisioner := &MultiCloudProvisioner{
		clusterUUID:            clusterID,
		isClusterSetToNonState: resp.GetContrailCluster().GetProvisioningState() == statusNoState,
		workDir:                getMCWorkRoot(clusterWorkRoot),
		reporter:               newMCReporter(apiServer, clusterID, logFile),
		log:                    logutil.NewFileLogger("multi-cloud-provisioner", logFile),
		apiServer:              apiServer,
		test:                   isTest,
		exec:                   exec,
	}

	if err = provisioner.setMultiCloudAction(resp.GetContrailCluster().GetProvisioningAction()); err != nil {
		return nil, err
	}

	if err = provisioner.updateCloudIDs(resp.GetContrailCluster()); err != nil {
		return nil, errors.Wrap(err, "cannot update MultiCloudProvisioner cloud refs")
	}

	reg, err := cloud.NewAuthorizedRegistry(resp.GetContrailCluster())
	if err != nil {
		return nil, errors.Wrap(err, "cannot update MultiCloudProvisioner authorized registry")
	}
	provisioner.authorizedRegistry = reg

	return provisioner, nil
}

func (m *MultiCloudProvisioner) setMultiCloudAction(action string) error {
	switch multiCloudAction(action) {
	case createMC, updateMC, deleteMC:
		m.action = multiCloudAction(action)
		return nil
	}
	return errors.Errorf("unknown MultiCloud action: %s", string(action))
}

func getMCWorkRoot(clusterWorkRoot string) string {
	if clusterWorkRoot == "" {
		clusterWorkRoot = defaultWorkRoot
	}
	return filepath.Join(clusterWorkRoot, mcWorkDir)
}

func newMCReporter(apiServer *client.HTTP, clusterID, logFile string) *report.Reporter {
	return report.NewReporter(
		apiServer, filepath.Join(defaultResourcePath, clusterID), logutil.NewFileLogger("reporter", logFile),
	)
}

func (m *MultiCloudProvisioner) updateCloudIDs(cluster *models.ContrailCluster) error {
	if len(cluster.CloudRefs) != 2 {
		return errors.Errorf("cluster %v should refer only two clouds (private and public). It refers to %v clouds",
			cluster.UUID, len(cluster.CloudRefs))
	}
	for _, c := range cluster.CloudRefs {
		cloudType, err := m.cloudType(c.UUID)
		if err != nil {
			return errors.Wrap(err, "cannot resolve type of the cloud")
		}
		switch cloudType {
		case onPrem:
			m.privateCloudID = c.GetUUID()
		case aws, azure, gcp:
			m.publicCloudID = c.GetUUID()
		default:
			return errors.Errorf("Cloud %s has invalid provider type: %s", c.UUID, cloudType)
		}
	}
	if m.privateCloudID == "" {
		return errors.Errorf("cluster %v is missing private Cloud ID", m.clusterUUID)
	}
	if m.publicCloudID == "" {
		return errors.Errorf("cluster %v is missing public Cloud ID", m.clusterUUID)
	}
	return nil
}

func (m *MultiCloudProvisioner) cloudType(uuid string) (string, error) {
	cloudResp, err := m.apiServer.GetCloud(context.Background(), &services.GetCloudRequest{ID: uuid})
	if err != nil {
		return "=", err
	}
	if len(cloudResp.GetCloud().GetCloudProviders()) == 0 {
		return "", errors.Errorf("cloud %s has no cloud providers", uuid)
	}
	provResp, err := m.apiServer.GetCloudProvider(context.Background(), &services.GetCloudProviderRequest{
		ID: cloudResp.GetCloud().GetCloudProviders()[0].GetUUID(),
	})
	if err != nil {
		return "", err
	}
	return provResp.GetCloudProvider().GetType(), nil
}

// Deploy performs Multicloud provisioning.
func (m *MultiCloudProvisioner) Deploy() error {
	if !m.isClusterSetToNonState {
		m.log.Infof("Got MC Request with state other than: %s", statusNoState)
		return nil
	}

	m.reportStatus(statusUpdateProgress)
	deployErr := m.deploy()
	if err := m.removeVulnerableFiles(); err != nil {
		deployErr = errors.Errorf("failed to delete vulnerable files: %v; deploy error (if any): %v", err, deployErr)
	}
	if err := m.unsetMulticloudProvisioningFlag(); err != nil {
		deployErr = errors.Errorf(
			"failed to unset cloud.is_multicloud_provisioning flag: %v; deploy error (if any): %v", err, deployErr,
		)
	}
	if deployErr != nil {
		m.reportStatus(statusUpdateFailed)
		m.log.Errorf("Multi-Cloud provisioning failed with error: %v", deployErr)
		return deployErr
	}
	m.reportStatus(statusUpdated)
	return nil
}

func (m *MultiCloudProvisioner) reportStatus(status string) {
	m.reporter.ReportStatus(context.Background(), map[string]interface{}{statusField: status}, defaultResource)
}

func (m *MultiCloudProvisioner) deploy() error {
	if m.action == deleteMC {
		return m.deleteMCCluster()
	}

	if err := m.verifyCloudStatus(); err != nil {
		return err
	}

	if updated, err := m.isMCUpdated(); err != nil {
		return err
	} else if updated {
		return nil
	}

	return m.manageMCCluster()
}

func (m *MultiCloudProvisioner) deleteMCCluster() error {
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

func (m *MultiCloudProvisioner) createFiles() error {
	if err := m.createClusterTopologyFile(m.workDir); err != nil {
		return err
	}
	return m.createClusterSecretFile()
}

func (m *MultiCloudProvisioner) createClusterTopologyFile(workDir string) error {
	topologyFile := m.getClusterTopoFile(workDir)
	if err := fileutil.CopyFile(cloud.GetTopoFile(m.privateCloudID), topologyFile, true); err != nil {
		return errors.Wrap(err, "topology file is not created for onprem cloud")
	}
	if err := appendFile(cloud.GetTopoFile(m.publicCloudID), topologyFile); err != nil {
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

func (m *MultiCloudProvisioner) createClusterSecretFile() error {
	pubCloudKeyPair, err := m.getPublicCloudKeyPair()
	if err != nil {
		return errors.Wrap(err, "failed to to get public Cloud keypair")
	}

	sfc := cloud.SecretFileConfig{}
	if err = sfc.Update(pubCloudKeyPair); err != nil {
		return errors.Wrap(err, "failed to to update secret file config")
	}

	sfc.AuthorizedRegistries = []*cloud.AuthorizedRegistry{m.authorizedRegistry}

	marshaled, err := yaml.Marshal(sfc)
	if err != nil {
		return errors.Wrapf(err, "couldn't marshal secret to yaml for cluster %s", m.clusterUUID)
	}
	err = ioutil.WriteFile(m.getClusterSecretFile(), marshaled, defaultFilePermRWOnly)
	return errors.Wrapf(err, "couldn't create secret.yml file for cluster %s", m.clusterUUID)
}

func (m *MultiCloudProvisioner) isMCUpdated() (bool, error) {
	if _, err := os.Stat(m.getClusterTopoFile(m.workDir)); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errors.Wrapf(err, "couldn't read old topology file: %s", m.getClusterTopoFile(m.workDir))
	}
	if ok, err := m.compareClusterTopologyFile(); err != nil {
		return true, err
	} else if ok {
		m.log.Infof("%s topology file is already up-to-date", defaultResource)
		return true, nil
	}
	return false, nil
}

func (m *MultiCloudProvisioner) compareClusterTopologyFile() (bool, error) {
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

func (m *MultiCloudProvisioner) manageMCCluster() error {
	m.log.Infof("Starting %s of MultiCloud: %s", m.action, m.clusterUUID)

	if err := m.createFiles(); err != nil {
		return err
	}
	return m.provision()
}

func (m *MultiCloudProvisioner) verifyCloudStatus() error {
	if err := m.waitForCloudStatusToBeUpdated(context.Background(), m.publicCloudID); err != nil {
		return err
	}
	return m.waitForCloudStatusToBeUpdated(context.Background(), m.privateCloudID)
}

func (m *MultiCloudProvisioner) waitForCloudStatusToBeUpdated(ctx context.Context, cloudUUID string) error {
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

func (m *MultiCloudProvisioner) provision() error {
	cmd := multicloudCLI
	args := []string{
		"all", "provision", "--topology", m.getClusterTopoFile(m.workDir),
		"--secret", m.getClusterSecretFile(), "--tf_state", cloud.GetTFStateFile(m.publicCloudID),
	}
	if m.action == updateMC {
		args = append(args, "--update")
	}
	agent, stopAgent, err := m.startSSHAgent()
	if err != nil {
		return errors.Wrap(err, "cannot start SSH Agent")
	}
	defer stopAgent()
	return m.exec.ExecCmdAndWait(m.reporter, cmd, args, m.getMCDeployerRepoDir(), agent.GetExportVars()...)
}

func (m *MultiCloudProvisioner) startSSHAgent() (*sshAgent, func(), error) {
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

func (m *MultiCloudProvisioner) setStatusToFail() {
	if m.action == createMC {
		m.reportStatus(statusCreateFailed)
	} else {
		m.reportStatus(statusUpdateFailed)
	}
}

func (m *MultiCloudProvisioner) cleanupProvisioning() error {
	cmd := multicloudCLI
	mcRepoDir := m.getMCDeployerRepoDir()
	args := []string{
		"all", "clean", "--topology", m.getClusterTopoFile(m.workDir),
		"--secret", m.getClusterSecretFile(), "--tf_state", cloud.GetTFStateFile(m.publicCloudID),
		"--state", filepath.Join(mcRepoDir, "state.yml"),
	}
	agent, stopAgent, err := m.startSSHAgent()
	if err != nil {
		return err
	}
	defer stopAgent()
	// TODO: Change inventory path after specifying work dir during provisioning.
	if err = m.exec.ExecCmdAndWait(m.reporter, cmd, args, mcRepoDir, agent.GetExportVars()...); err != nil {
		return err
	}
	return m.unsetIsMultiCloudFlag()
}

func (m *MultiCloudProvisioner) unsetIsMultiCloudFlag() error {
	_, err := m.apiServer.UpdateContrailCluster(context.Background(), &services.UpdateContrailClusterRequest{
		ContrailCluster: &models.ContrailCluster{
			UUID:         m.clusterUUID,
			IsMulticloud: false,
		},
	})
	return err
}

func (m *MultiCloudProvisioner) removeVulnerableFiles() error {
	removeErr := osutil.ForceRemoveFiles(m.filesToRemove(), m.log)
	return errors.Wrap(removeErr, "failed to remove credential files")
}

func (m *MultiCloudProvisioner) filesToRemove() []string {
	kfd := services.NewKeyFileDefaults()

	return []string{
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
}

// unsetMulticloudProvisioningFlag unsets cloud.is_multicloud_provisioning flag.
// See pkg/cloud.Cloud.removeVulnerableFiles() comment regarding cloud.is_multicloud_provisioning flag purpose.
func (m *MultiCloudProvisioner) unsetMulticloudProvisioningFlag() error {
	if _, err := m.apiServer.UpdateCloud(
		context.Background(),
		&services.UpdateCloudRequest{
			Cloud: &models.Cloud{
				UUID:                     m.publicCloudID,
				IsMulticloudProvisioning: false,
			},
		},
	); err != nil {
		return errors.Wrapf(err, "failed to update Cloud with UUID: %s", m.publicCloudID)
	}

	m.log.WithField("public-cloud-id", m.publicCloudID).Debug("Successfully unset cloud.is_multicloud_provisioning flag")
	return nil
}

func (m *MultiCloudProvisioner) removeCloudRefFromCluster() error {
	for _, cloudUUID := range []string{m.privateCloudID, m.publicCloudID} {
		if _, err := m.apiServer.DeleteContrailClusterCloudRef(
			context.Background(), &services.DeleteContrailClusterCloudRefRequest{
				ID: m.clusterUUID,
				ContrailClusterCloudRef: &models.ContrailClusterCloudRef{
					UUID: cloudUUID,
				},
			},
		); err != nil {
			return err
		}
	}
	return nil
}

func (m *MultiCloudProvisioner) getPublicCloudKeyPair() (*models.Keypair, error) {
	ctx := context.Background()
	cloudObj, err := cloud.GetCloud(ctx, m.apiServer, m.publicCloudID)
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

func (m *MultiCloudProvisioner) getMCInventoryFile(workDir string) string {
	return filepath.Join(workDir, defaultMCInventoryFile)
}

// use topology constant from cloud pkg
func (m *MultiCloudProvisioner) getClusterTopoFile(workDir string) string {
	return filepath.Join(workDir, defaultTopologyFile)
}

func (m *MultiCloudProvisioner) getClusterSecretFile() string {
	return filepath.Join(m.workDir, defaultSecretFile)
}

func (m *MultiCloudProvisioner) getMCDeployerRepoDir() string {
	return filepath.Join(defaultAnsibleRepoDir, mcRepository)
}
