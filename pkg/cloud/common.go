package cloud

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/retry"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	testTemplate        = "./test_data/test_cmd.tmpl"
	executedCmdTestFile = "executed_cmd.yml"

	statusRetryInterval = 3 * time.Second
)

// GetCloudDir gets directory of cloud
func GetCloudDir(cloudID string) string {
	return filepath.Join(defaultWorkRoot, cloudID)
}

// getTemplateRoot returns path to the directory where cloud templates are stored
func (c *Cloud) getTemplateRoot() string {
	templateRoot := c.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}

// GetMultiCloudRepodir returns path to multi-cloud directory
func GetMultiCloudRepodir() string {
	return filepath.Join(defaultMultiCloudDir, defaultMultiCloudRepo)
}

// getGenerateTopologyCmd returns generate topology command of mc deployer
func getGenerateTopologyCmd(mcDir string) string {
	return filepath.Join(mcDir, defaultGenTopoScript)
}

// GetGenInventoryCmd get generate inventory command
func GetGenInventoryCmd(mcDir string) string {
	return filepath.Join(mcDir, defaultGenInventoryScript)
}

// TestCmdHelper helps to write cmd to a file (instead of executing)
func TestCmdHelper(cmd string, args []string, workDir string, testTemplate string) error {
	context := pongo2.Context{
		"cmd":  cmd,
		"args": args,
	}

	content, err := template.Apply(testTemplate, context)
	if err != nil {
		return err
	}

	destPath := filepath.Join(workDir, executedCmdTestFile)
	return fileutil.AppendToFile(destPath, content, defaultRWOnlyPerm)
}

// GetCloud gets cloud data for a given cloud UUID
func GetCloud(ctx context.Context, client *client.HTTP, cloudID string) (*models.Cloud, error) {

	request := new(services.GetCloudRequest)
	request.ID = cloudID

	cloudResp, err := client.GetCloud(ctx, request)
	if err != nil {
		return nil, err
	}

	return cloudResp.GetCloud(), nil
}

// GetTopoFile gets topology file for a cloud
func GetTopoFile(cloudID string) string {
	return filepath.Join(GetCloudDir(cloudID), defaultTopologyFile)
}

// GetSecretFile gets secret file for a cloud
func GetSecretFile(cloudID string) string {
	return filepath.Join(GetCloudDir(cloudID), defaultSecretFile)
}

// GetTFStateFile get terraform state file
func GetTFStateFile(cloudID string) string {
	return filepath.Join(GetCloudDir(cloudID), defaultTFStateFile)
}

// deleteNodeObjects deletes all the node objects given as input
// returns list of error
func deleteNodeObjects(ctx context.Context,
	client *client.HTTP, nodeList []*instanceData) []string {

	var errList []string
	// Delete Node related dependencies and node itself
	for _, node := range nodeList {
		if node.info.PortGroups != nil {
			for _, portGroup := range node.info.PortGroups {
				_, err := client.DeletePortGroup(ctx,
					&services.DeletePortGroupRequest{
						ID: portGroup.UUID,
					},
				)
				if err != nil {
					errList = append(errList, fmt.Sprintf(
						"failed deleting PortGroup %s err_msg: %s",
						portGroup.UUID, err))
				}
			}
		}
		if node.info.Ports != nil {
			for _, port := range node.info.Ports {
				_, err := client.DeletePort(ctx,
					&services.DeletePortRequest{
						ID: port.UUID,
					},
				)
				if err != nil {
					errList = append(errList, fmt.Sprintf(
						"failed deleting Port %s err_msg: %s",
						port.UUID, err))
				}
			}
		}
		_, err := client.DeleteNode(ctx,
			&services.DeleteNodeRequest{
				ID: node.info.UUID,
			},
		)
		if err != nil {
			errList = append(errList, fmt.Sprintf(
				"failed deleting Node %s err_msg: %s",
				node.info.UUID, err))
		}
	}
	return errList
}

// removePvtSubnetRefFromNodes remove cloud-private-subnet refs from node
func removePvtSubnetRefFromNodes(ctx context.Context,
	client *client.HTTP, nodeList []*instanceData) error {

	for _, node := range nodeList {
		for _, cloudPvtSubnetRef := range node.info.CloudPrivateSubnetRefs {
			_, err := client.DeleteNodeCloudPrivateSubnetRef(ctx,
				&services.DeleteNodeCloudPrivateSubnetRefRequest{
					ID:                        node.info.UUID,
					NodeCloudPrivateSubnetRef: cloudPvtSubnetRef,
				},
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// deleteContrailMCGWRole deletes contrail-contrail-multicloud-node role object associated with the instances
func deleteContrailMCGWRole(ctx context.Context,
	client *client.HTTP, nodeList []*instanceData) []string {

	var errList []string
	for _, node := range nodeList {
		if node.info.ContrailMulticloudGWNodeBackRefs != nil {
			for _, mcGWNodeBackRef := range node.info.ContrailMulticloudGWNodeBackRefs {
				_, err := client.DeleteContrailMulticloudGWNode(ctx,
					&services.DeleteContrailMulticloudGWNodeRequest{
						ID: mcGWNodeBackRef.UUID,
					},
				)
				if err != nil {
					errList = append(errList, fmt.Sprintf(
						"failed deleting ContrailMulticloudGWNode %s err_msg: %s",
						mcGWNodeBackRef.UUID, err))
				}
			}
		}
	}
	return errList
}

// deleteSGObjects deletes given cloud-security-group objects
func deleteSGObjects(ctx context.Context,
	client *client.HTTP, sgList []*sgData) []string {

	errList := []string{}
	// Delete CloudSecurityGroup related dependencies and CloudSecurityGroup itself
	for _, sg := range sgList {
		for _, sgRule := range sg.info.CloudSecurityGroupRules {
			_, err := client.DeleteCloudSecurityGroupRule(ctx,
				&services.DeleteCloudSecurityGroupRuleRequest{
					ID: sgRule.UUID,
				},
			)
			if err != nil {
				errList = append(errList, fmt.Sprintf(
					"failed deleting CloudSecurityGroupRule %s err_msg: %s",
					sgRule.UUID, err))
			}
		}
		_, err := client.DeleteCloudSecurityGroup(ctx,
			&services.DeleteCloudSecurityGroupRequest{
				ID: sg.info.UUID,
			},
		)
		if err != nil {
			errList = append(errList, fmt.Sprintf(
				"failed deleting CloudSecurityGroup %s err_msg: %s",
				sg.info.UUID, err))
		}
	}
	return errList
}

// deletePvtSubnetObjects deletes cloud-private-subnet-objecs
func deletePvtSubnetObjects(ctx context.Context,
	client *client.HTTP, subnetList []*subnetData) []string {

	errList := []string{}
	// Delete CloudPrivateSubnet related dependencies and CloudPrivateSubnet itself
	for _, pvtsubnet := range subnetList {
		_, err := client.DeleteCloudPrivateSubnet(ctx,
			&services.DeleteCloudPrivateSubnetRequest{
				ID: pvtsubnet.info.UUID,
			},
		)
		if err != nil {
			errList = append(errList, fmt.Sprintf(
				"failed deleting CloudPrivateSubnet %s err_msg: %s",
				pvtsubnet.info.UUID, err))
		}
	}
	return errList
}

// deleteCloudProviderAndDeps deletes cloud-providers and its dependencies
func deleteCloudProviderAndDeps(ctx context.Context,
	client *client.HTTP, providerList []*providerData) []string {

	errList := []string{}
	// Delete Provider dependencies and iteslf
	for _, provider := range providerList {
		for _, region := range provider.regions {
			for _, vc := range region.virtualClouds {

				retErrList := deleteSGObjects(ctx, client, vc.sgs)
				if retErrList != nil {
					errList = append(errList, retErrList...)
				}

				retErrList = deletePvtSubnetObjects(ctx, client, vc.subnets)
				if retErrList != nil {
					errList = append(errList, retErrList...)
				}

				_, err := client.DeleteVirtualCloud(ctx,
					&services.DeleteVirtualCloudRequest{
						ID: vc.info.UUID,
					},
				)
				if err != nil {
					errList = append(errList, fmt.Sprintf(
						"failed deleting VirtualCloud %s err_msg: %s",
						vc.info.UUID, err))
				}
			}
			_, err := client.DeleteCloudRegion(ctx,
				&services.DeleteCloudRegionRequest{
					ID: region.info.UUID,
				},
			)
			if err != nil {
				errList = append(errList, fmt.Sprintf(
					"failed deleting CloudRegion %s err_msg: %s",
					region.info.UUID, err))
			}
		}
		_, err := client.DeleteCloudProvider(ctx,
			&services.DeleteCloudProviderRequest{
				ID: provider.info.UUID,
			},
		)
		if err != nil {
			errList = append(errList, fmt.Sprintf(
				"failed deleting CloudProvider %s err_msg: %s",
				provider.info.UUID, err))
		}
	}
	return errList
}

// deleteCloudUsers deletes given cloud users
func deleteCloudUsers(ctx context.Context,
	client *client.HTTP, userList []*models.CloudUser) []string {

	var errList []string
	// Delete user & its dependencies
	for _, u := range userList {
		_, err := client.DeleteCloudUser(ctx,
			&services.DeleteCloudUserRequest{
				ID: u.UUID,
			},
		)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}
	return errList
}

// deleteCredentialAndDeps delete credentials and dependencies associated with it
func deleteCredentialAndDeps(ctx context.Context,
	client *client.HTTP, credList []*models.Credential) []string {

	var errList []string
	// Delete credential & its dependencies
	for _, cred := range credList {

		_, err := client.DeleteCredential(ctx,
			&services.DeleteCredentialRequest{
				ID: cred.UUID,
			},
		)
		if err != nil {
			errList = append(errList, err.Error())
		}

		for _, kp := range cred.KeypairRefs {
			_, err := client.DeleteKeypair(ctx,
				&services.DeleteKeypairRequest{
					ID: kp.UUID,
				},
			)
			if err != nil {
				errList = append(errList, err.Error())
			}
		}
	}
	return errList
}

// nolint: gocyclo
// deleteAPIObjects deletes all the schema objects associated with cloud (children and refs and there deps)
func (c *Cloud) deleteAPIObjects(d *Data) error {

	if d.isCloudPrivate() {
		err := removePvtSubnetRefFromNodes(c.ctx, c.APIServer, d.getGatewayNodes())
		if err != nil {
			return err
		}
	}

	var errList, warnList []string

	retErrList := deleteContrailMCGWRole(c.ctx,
		c.APIServer, d.getGatewayNodes())

	if retErrList != nil {
		errList = append(errList, retErrList...)
	}

	if d.isCloudPublic() {
		retErrList = deleteNodeObjects(c.ctx, c.APIServer, d.instances)
		if retErrList != nil {
			errList = append(errList, retErrList...)
		}
	}

	retErrList = deleteCloudProviderAndDeps(c.ctx,
		c.APIServer, d.providers)
	if retErrList != nil {
		errList = append(errList, retErrList...)
	}

	_, err := c.APIServer.DeleteCloud(c.ctx,
		&services.DeleteCloudRequest{
			ID: d.info.UUID,
		},
	)
	if err != nil {
		errList = append(errList, fmt.Sprintf(
			"failed deleting Cloud %s err_msg: %s",
			d.info.UUID, err))
	}

	cloudUserErrList := deleteCloudUsers(c.ctx, c.APIServer, d.users)
	if cloudUserErrList != nil {
		warnList = append(warnList, cloudUserErrList...)
	}

	if d.isCloudPublic() {
		credErrList := deleteCredentialAndDeps(c.ctx, c.APIServer, d.credentials)
		warnList = append(warnList, credErrList...)
	}

	// log the warning messages
	if len(warnList) > 0 {
		c.log.Warnf("could not delete cloud refs deps because of errors: %s",
			strings.Join(warnList, "\n"))
	}
	// join all the errors and return it
	if len(errList) > 0 {
		return errors.New(strings.Join(errList, "\n"))
	}
	return nil
}

// genKeyPair generates ssh keypair
func genKeyPair(bits int) ([]byte, []byte, error) {

	// creating private key
	pvtKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	// encoding private key with PEM format
	pvtKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pvtKey),
	}
	var encodedPvtKey bytes.Buffer
	if err = pem.Encode(&encodedPvtKey, pvtKeyPEM); err != nil {
		return nil, nil, err
	}

	// creating public key
	pubKey, err := ssh.NewPublicKey(&pvtKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}

	pub := ssh.MarshalAuthorizedKey(pubKey)
	return pub, encodedPvtKey.Bytes(), nil
}

// tfStateOutputExists checks if terrafrom state file exists in cloud work dir
func tfStateOutputExists(cloudID string) bool {

	tfState, err := readStateFile(GetTFStateFile(cloudID))
	if err != nil {
		return false
	}

	mState := tfState.RootModule()
	if len(mState.Outputs) == 0 && len(mState.Resources) == 0 {
		return false
	}
	return true
}

// waitForClusterStatusToBeUpdated waits for contrail-cluster processing to complete
// only waits if cluster's status is showing as "UPDATE_IN_PROGRESS" or "CREATE_IN_PROGRESS"
func waitForClusterStatusToBeUpdated(ctx context.Context, log *logrus.Entry,
	httpClient *client.HTTP, clusterUUID string) error {

	return retry.Do(func() (retry bool, err error) {

		clusterResp, err := httpClient.GetContrailCluster(ctx,
			&services.GetContrailClusterRequest{
				ID: clusterUUID,
			},
		)
		if err != nil {
			return false, err
		}

		if clusterResp.ContrailCluster.ProvisioningAction == "DELETE_CLOUD" {
			switch clusterResp.ContrailCluster.ProvisioningState {
			case statusCreateProgress:
				return true, fmt.Errorf("waiting for create cluster %s to complete",
					clusterUUID)
			case statusUpdateProgress:
				return true, fmt.Errorf("waiting for update cluster %s to complete",
					clusterUUID)
			case statusNoState:
				return true, fmt.Errorf("waiting for cluster %s to complete processing",
					clusterUUID)
			case statusCreated:
				return false, nil
			case statusUpdated:
				return false, nil
			case statusCreateFailed:
				return false, fmt.Errorf("cluster %s status has failed in creating",
					clusterUUID)
			case statusUpdateFailed:
				return false, fmt.Errorf("cluster %s status has failed in updating",
					clusterUUID)
			}
			return false, fmt.Errorf("unknown cluster status %s for cluster %s",
				clusterResp.ContrailCluster.ProvisioningState, clusterUUID)

		}

		return false, fmt.Errorf("contrail-cluster %s does not have provisioning_action set to DELETE_CLOUD",
			clusterUUID)
	}, retry.WithLog(log.Logger),
		retry.WithInterval(statusRetryInterval))
}
