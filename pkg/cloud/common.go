package cloud

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	testTemplate        = "./test_data/test_cmd.tmpl"
	executedCmdTestFile = "executed_cmd.yml"
)

// GetCloudDir gets directory of cloud
func GetCloudDir(cloudID string) string {
	return filepath.Join(defaultWorkRoot, cloudID)
}

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

func deleteNodeObjects(ctx context.Context,
	client *client.HTTP, nodeList []*instanceData) error {

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
					return err
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
					return err
				}
			}
		}
		_, err := client.DeleteNode(ctx,
			&services.DeleteNodeRequest{
				ID: node.info.UUID,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteSGObjects(ctx context.Context,
	client *client.HTTP, sgList []*sgData) error {

	// Delete CloudSecurityGroup related dependencies and CloudSecurityGroup itself
	for _, sg := range sgList {
		for _, sgRule := range sg.info.CloudSecurityGroupRules {
			_, err := client.DeleteCloudSecurityGroupRule(ctx,
				&services.DeleteCloudSecurityGroupRuleRequest{
					ID: sgRule.UUID,
				},
			)
			if err != nil {
				return err
			}
		}
		_, err := client.DeleteCloudSecurityGroup(ctx,
			&services.DeleteCloudSecurityGroupRequest{
				ID: sg.info.UUID,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func deletePvtSubnetObjects(ctx context.Context,
	client *client.HTTP, subnetList []*subnetData) error {

	// Delete CloudPrivateSubnet related dependencies and CloudPrivateSubnet itself
	for _, pvtsubnet := range subnetList {
		_, err := client.DeleteCloudPrivateSubnet(ctx,
			&services.DeleteCloudPrivateSubnetRequest{
				ID: pvtsubnet.info.UUID,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteCloudProviderAndDeps(ctx context.Context,
	client *client.HTTP, providerList []*providerData) error {

	// Delete Provider dependencies and iteslf
	for _, provider := range providerList {
		for _, region := range provider.regions {
			for _, vc := range region.virtualClouds {

				err := deleteSGObjects(ctx, client, vc.sgs)
				if err != nil {
					return err
				}

				err = deletePvtSubnetObjects(ctx, client, vc.subnets)
				if err != nil {
					return err
				}

				_, err = client.DeleteVirtualCloud(ctx,
					&services.DeleteVirtualCloudRequest{
						ID: vc.info.UUID,
					},
				)
				if err != nil {
					return err
				}
			}
			_, err := client.DeleteCloudRegion(ctx,
				&services.DeleteCloudRegionRequest{
					ID: region.info.UUID,
				},
			)
			if err != nil {
				return err
			}
		}
		_, err := client.DeleteCloudProvider(ctx,
			&services.DeleteCloudProviderRequest{
				ID: provider.info.UUID,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

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

func (c *Cloud) deleteAPIObjects(d *Data) error {

	if d.isCloudPublic() {
		err := deleteNodeObjects(c.ctx, c.APIServer, d.instances)
		if err != nil {
			return err
		}
	}

	err := deleteCloudProviderAndDeps(c.ctx,
		c.APIServer, d.providers)
	if err != nil {
		c.log.Error("error while trying to delete cloud provider & deps")
		return err
	}

	_, err = c.APIServer.DeleteCloud(c.ctx,
		&services.DeleteCloudRequest{
			ID: d.cloud.config.CloudID,
		},
	)
	if err != nil {
		c.log.Errorf("error while trying to delete cloud: %s",
			d.cloud.config.CloudID)
		return err
	}

	var errList []string

	cloudUserErrList := deleteCloudUsers(c.ctx, c.APIServer, d.users)
	if cloudUserErrList != nil {
	}
	errList = append(errList, cloudUserErrList...)

	if d.isCloudPublic() {
		credErrList := deleteCredentialAndDeps(c.ctx, c.APIServer, d.credentials)
		errList = append(errList, credErrList...)
	}

	// join all the errors and return it
	if len(errList) > 0 {
		return errors.New(strings.Join(errList, "\n"))
	}
	return nil
}

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
