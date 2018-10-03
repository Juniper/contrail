package cloud

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/flosch/pongo2"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	defaultWorkRoot           = "/var/tmp/cloud/config"
	defaultTemplateRoot       = "./pkg/cloud/configs"
	defaultGenTopoScript      = "transform/generate_topology.py"
	defaultGenInventoryScript = "transform/generate_inventories.py"
	defaultTFStateFile        = "terraform.tfstate"
	testTemplate              = "./test_data/test_cmd.tmpl"
	executedCmdTestFile       = "executed_cmd.yml"
)

// GetCloudDir gets directory of cloud
func GetCloudDir(cloudID string) string {
	return filepath.Join(getWorkRoot(), cloudID)
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

func getWorkRoot() string {
	return defaultWorkRoot
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
func GetCloud(ctx context.Context, httpServer *client.HTTP, cloudID string) (*models.Cloud, error) {

	request := new(services.GetCloudRequest)
	request.ID = cloudID

	cloudResp, err := httpServer.GetCloud(ctx, request)
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

type mockMetadataGetter basemodels.Metadata

func (m *mockMetadataGetter) GetMetadata(_ context.Context,
	_ basemodels.Metadata) (*basemodels.Metadata, error) {
	return (*basemodels.Metadata)(m), nil
}

func (m *mockMetadataGetter) ListMetadata(ctx context.Context,
	metadataSlice []*basemodels.Metadata) ([]*basemodels.Metadata, error) {
	return []*basemodels.Metadata{(*basemodels.Metadata)(m)}, nil
}

func (c *Cloud) deleteAPIObjects(d *Data) error {

	var errList []string
	// Delete Node related dependencies and node itself
	for _, instance := range d.instances {
		if instance.info.PortGroups != nil {
			for _, portGroup := range instance.info.PortGroups {
				_, err := c.APIServer.DeletePortGroup(c.ctx,
					&services.DeletePortGroupRequest{
						ID: portGroup.UUID,
					},
				)
				if err != nil {
					errList = append(errList, err.Error())
				}
			}
		}
		if instance.info.Ports != nil {
			for _, port := range instance.info.Ports {
				_, err := c.APIServer.DeletePort(c.ctx,
					&services.DeletePortRequest{
						ID: port.UUID,
					},
				)
				if err != nil {
					errList = append(errList, err.Error())
				}
			}
		}
		_, err := c.APIServer.DeleteNode(c.ctx,
			&services.DeleteNodeRequest{
				ID: instance.info.UUID,
			},
		)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}

	// Delete CloudSecurityGroup related dependencies and CloudSecurityGroup itself
	for _, sg := range d.securityGroups {
		for _, sgRule := range sg.CloudSecurityGroupRules {
			_, err := c.APIServer.DeleteCloudSecurityGroupRule(c.ctx,
				&services.DeleteCloudSecurityGroupRuleRequest{
					ID: sgRule.UUID,
				},
			)
			if err != nil {
				errList = append(errList, err.Error())
			}
		}
		_, err := c.APIServer.DeleteCloudSecurityGroup(c.ctx,
			&services.DeleteCloudSecurityGroupRequest{
				ID: sg.UUID,
			},
		)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}

	// Delete CloudPrivateSubnet related dependencies and CloudPrivateSubnet itself
	for _, pvtsubnet := range d.subnets {
		_, err := c.APIServer.DeleteCloudPrivateSubnet(c.ctx,
			&services.DeleteCloudPrivateSubnetRequest{
				ID: pvtsubnet.UUID,
			},
		)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}

	// Delete Provider dependencies and iteslf
	for _, provider := range d.providers {
		for _, region := range provider.regions {
			for _, vc := range region.virtualClouds {

				_, err := c.APIServer.DeleteVirtualCloud(c.ctx,
					&services.DeleteVirtualCloudRequest{
						ID: vc.info.UUID,
					},
				)
				if err != nil {
					errList = append(errList, err.Error())
				}
			}
			_, err := c.APIServer.DeleteCloudRegion(c.ctx,
				&services.DeleteCloudRegionRequest{
					ID: region.info.UUID,
				},
			)
			if err != nil {
				errList = append(errList, err.Error())
			}
		}
		_, err := c.APIServer.DeleteCloudProvider(c.ctx,
			&services.DeleteCloudProviderRequest{
				ID: provider.info.UUID,
			},
		)
		if err != nil {
			errList = append(errList, err.Error())
		}

		// Delete credential & its dependencies
		for _, cred := range d.credentials {
			for _, kp := range cred.KeypairRefs {
				_, err := c.APIServer.DeleteKeypair(c.ctx,
					&services.DeleteKeypairRequest{
						ID: kp.UUID,
					},
				)
				if err != nil {
					errList = append(errList, err.Error())
				}
			}

			_, err := c.APIServer.DeleteCredential(c.ctx,
				&services.DeleteCredentialRequest{
					ID: cred.UUID,
				},
			)
			if err != nil {
				errList = append(errList, err.Error())
			}
		}

		// Delete user & its dependencies
		for _, u := range d.users {
			_, err := c.APIServer.DeleteCloudUser(c.ctx,
				&services.DeleteCloudUserRequest{
					ID: u.UUID,
				},
			)
			if err != nil {
				errList = append(errList, err.Error())
			}
		}
	}

	return nil
}

//func (c *Cloud) deleteAPIObjects(d *Data) error {
//
//	var events []*services.Event
//
//	for _, u := range d.users {
//		events = append(events, userDeleteEvent(u.UUID))
//	}
//
//	for _, p := range d.providers {
//		events = append(events, providerDeleteEvent(p.info.UUID))
//		for _, r := range p.regions {
//			events = append(events, regionDeleteEvent(r.info.UUID))
//			for _, vc := range r.virtualClouds {
//				events = append(events, vCloudDeleteEvent(vc.info.UUID))
//				for _, sg := range vc.sgs {
//					events = append(events, sgDeleteEvent(sg.info.UUID))
//					for _, sgRule := range sg.info.CloudSecurityGroupRules {
//						events = append(events, sgRuleDeleteEvent(sgRule.UUID))
//					}
//				}
//				for _, subnet := range vc.subnets {
//					events = append(events, subnetDeleteEvent(subnet.info.UUID))
//				}
//			}
//		}
//	}
//
//	for _, i := range d.instances {
//		events = append(events, instanceDeleteEvent(i.info.UUID))
//	}
//
//	events = append(events, cloudDeleteEvent(d.info.UUID))
//
//	tv, err := models.NewTypeValidatorWithFormat()
//	if err != nil {
//		return err
//	}
//
//	eventList := services.EventList{events}
//	err = eventList.Sort()
//	if err != nil {
//		return err
//	}
//
//	service := &services.ContrailService{
//		BaseService:    services.BaseService{},
//		MetadataGetter: (*mockMetadataGetter)(new(basemodels.Metadata)),
//		TypeValidator:  tv,
//	}
//
//	responses, err := eventList.Process(c.ctx, service)
//	if err != nil {
//		return errutil.ToHTTPError(err)
//	}
//
//	return err
//}

func userDeleteEvent(uuid string) *services.Event {

	return &services.Event{
		Request: &services.Event_DeleteCloudUserRequest{
			&services.DeleteCloudUserRequest{
				ID: uuid,
			},
		},
	}
}

func providerDeleteEvent(uuid string) *services.Event {

	return &services.Event{
		Request: &services.Event_DeleteCloudProviderRequest{
			&services.DeleteCloudProviderRequest{
				ID: uuid,
			},
		},
	}

}

func regionDeleteEvent(uuid string) *services.Event {

	return &services.Event{
		Request: &services.Event_DeleteCloudRegionRequest{
			&services.DeleteCloudRegionRequest{
				ID: uuid,
			},
		},
	}

}

func vCloudDeleteEvent(uuid string) *services.Event {

	return &services.Event{
		Request: &services.Event_DeleteVirtualCloudRequest{
			&services.DeleteVirtualCloudRequest{
				ID: uuid,
			},
		},
	}

}

func sgDeleteEvent(uuid string) *services.Event {

	return &services.Event{
		Request: &services.Event_DeleteCloudSecurityGroupRequest{
			&services.DeleteCloudSecurityGroupRequest{
				ID: uuid,
			},
		},
	}

}

func sgRuleDeleteEvent(uuid string) *services.Event {

	return &services.Event{
		Request: &services.Event_DeleteCloudSecurityGroupRuleRequest{
			&services.DeleteCloudSecurityGroupRuleRequest{
				ID: uuid,
			},
		},
	}

}

func subnetDeleteEvent(uuid string) *services.Event {

	return &services.Event{
		Request: &services.Event_DeleteCloudPrivateSubnetRequest{
			&services.DeleteCloudPrivateSubnetRequest{
				ID: uuid,
			},
		},
	}

}

func instanceDeleteEvent(uuid string) *services.Event {

	return &services.Event{
		Request: &services.Event_DeleteNodeRequest{
			&services.DeleteNodeRequest{
				ID: uuid,
			},
		},
	}

}

func cloudDeleteEvent(uuid string) *services.Event {

	return &services.Event{
		Request: &services.Event_DeleteCloudRequest{
			&services.DeleteCloudRequest{
				ID: uuid,
			},
		},
	}
}
