package cloud

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/common"
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

func getCloudUser(d *Data) (*models.CloudUser, error) {

	if d.users != nil {
		return d.users[0], nil
	}
	return nil, fmt.Errorf("cloud user not found")
}

func getUserCred(user *models.CloudUser) (username string, password string, err error) {

	username = user.AzureCredential.Username
	password = user.AzureCredential.Password

	if username == "" || password == "" {
		return username, password, fmt.Errorf("username or password not found for user uuid: %s", user.UUID)
	}

	return username, password, nil
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

	content, err := common.Apply(testTemplate, context)
	if err != nil {
		return err
	}

	destPath := filepath.Join(workDir, executedCmdTestFile)
	return common.AppendToFile(destPath, content, defaultRWOnlyPerm)
}

// GetCloud gets cloud data for a given cloud UUID
func GetCloud(client *client.HTTP, cloudID string) (*models.Cloud, error) {

	response := new(services.GetCloudResponse)

	_, err := client.Read("/cloud/"+cloudID, response)
	if err != nil {
		return nil, err
	}

	return response.GetCloud(), nil
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

func deleteNodeObjects(client *client.HTTP,
	nodeList []*instanceData) []string {

	var errList []string
	// Delete Node related dependencies and node itself
	for _, node := range nodeList {
		if node.info.PortGroups != nil {
			for _, portGroup := range node.info.PortGroups {
				_, err := client.Delete("/port-group/"+portGroup.UUID, nil)
				if err != nil {
					errList = append(errList, err.Error())
				}
			}
		}
		if node.info.Ports != nil {
			for _, port := range node.info.Ports {
				_, err := client.Delete("/port/"+port.UUID, nil)
				if err != nil {
					errList = append(errList, err.Error())
				}
			}
		}
		_, err := client.Delete("/node/"+node.info.UUID, nil)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}
	return errList
}

func deleteSGObjects(client *client.HTTP, sgList []*sgData) []string {

	var errList []string

	// Delete CloudSecurityGroup related dependencies and CloudSecurityGroup itself
	for _, sg := range sgList {
		for _, sgRule := range sg.info.CloudSecurityGroupRules {
			_, err := client.Delete("/cloud-security-group-rule/"+sgRule.UUID, nil)
			if err != nil {
				errList = append(errList, err.Error())
			}
		}
		_, err := client.Delete("/cloud-security-group/"+sg.info.UUID, nil)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}
	return errList
}

func deletePvtSubnetObjects(client *client.HTTP,
	subnetList []*subnetData) []string {

	var errList []string

	// Delete CloudPrivateSubnet related dependencies and CloudPrivateSubnet itself
	for _, pvtsubnet := range subnetList {
		_, err := client.Delete("/cloud-private-subnet/"+pvtsubnet.info.UUID, nil)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}
	return errList
}

func deleteCloudProviderAndDeps(client *client.HTTP,
	providerList []*providerData) []string {

	var errList []string

	// Delete Provider dependencies and iteslf
	for _, provider := range providerList {
		for _, region := range provider.regions {
			for _, vc := range region.virtualClouds {

				sgErrList := deleteSGObjects(client, vc.sgs)
				errList = append(errList, sgErrList...)

				pvtSubnetErrList := deletePvtSubnetObjects(client, vc.subnets)
				errList = append(errList, pvtSubnetErrList...)

				_, err := client.Delete("/virtual-cloud/"+vc.info.UUID, nil)
				if err != nil {
					errList = append(errList, err.Error())
				}
			}
			_, err := client.Delete("/cloud-region/"+region.info.UUID, nil)
			if err != nil {
				errList = append(errList, err.Error())
			}
		}
		_, err := client.Delete("/cloud-provider/"+provider.info.UUID, nil)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}
	return errList
}

func deleteCloudUsers(client *client.HTTP,
	userList []*models.CloudUser) []string {

	var errList []string
	// Delete user & its dependencies
	for _, u := range userList {
		_, err := client.Delete("/cloud-user/"+u.UUID, nil)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}
	return errList
}

func deleteCredentialAndDeps(client *client.HTTP,
	credList []*models.Credential) []string {

	var errList []string
	// Delete credential & its dependencies
	for _, cred := range credList {
		_, err := client.Delete("/credential/"+cred.UUID, nil)
		if err != nil {
			errList = append(errList, err.Error())
		}

		for _, kp := range cred.KeypairRefs {
			_, err := client.Delete("/keypair/"+kp.UUID, nil)
			if err != nil {
				errList = append(errList, err.Error())
			}
		}
	}
	return errList
}

func (c *Cloud) deleteAPIObjects(d *Data) error {

	var errList []string

	nodeErrList := deleteNodeObjects(c.APIServer, d.instances)
	errList = append(errList, nodeErrList...)

	providerErrList := deleteCloudProviderAndDeps(c.APIServer, d.providers)
	errList = append(errList, providerErrList...)

	_, err := c.APIServer.Delete("/cloud/"+d.cloud.config.CloudID, nil)
	if err != nil {
		errList = append(errList, err.Error())
	}

	cloudUserErrList := deleteCloudUsers(c.APIServer, d.users)
	errList = append(errList, cloudUserErrList...)

	if d.isCloudPublic() {
		credErrList := deleteCredentialAndDeps(c.APIServer, d.credentials)
		errList = append(errList, credErrList...)
	}

	// join all the errors and return it
	if len(errList) > 0 {
		return errors.New(strings.Join(errList, "\n"))
	}
	return nil
}
