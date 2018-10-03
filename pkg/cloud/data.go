package cloud

import (
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/labstack/echo"
)

//Data for cloud provider data
type Data struct {
	cloud               *Cloud
	cloudProviderInfo   *models.Cloud
	projectName         string
	credentials         []*models.Credential
	cloudSubnets        []*models.CloudPrivateSubnet
	cloudSecurityGroups []*models.CloudSecurityGroup
}

func (d *Data) updateData() error {

	c := d.cloud.echo.AcquireContext()

	subnets, err := d.cloud.listCloudPrivatesubnet(c)
	if err != nil {
		return err
	}
	d.cloudSubnets = subnets

	sgs, err := d.cloud.listCloudSecurityGroup(c)
	if err != nil {
		return err
	}
	d.cloudSecurityGroups = sgs

	creds, err := d.cloud.listCredential(c)
	if err != nil {
		return err
	}
	d.credentials = creds

	// TODO(madhukar) - Needs to handle multiple projects
	for _, account := range d.cloudProviderInfo.CloudAccounts {
		for _, project := range account.CloudProjects {
			if project.Name != "" {
				d.projectName = project.Name
				break
			}
		}
	}
	return nil
}

func (c *Cloud) listCloudSecurityGroup(e echo.Context) ([]*models.CloudSecurityGroup, error) {
	spec := baseservices.GetListSpec(e)
	request := &services.ListCloudSecurityGroupRequest{
		Spec: spec,
	}
	ctx := returnContext(e)
	response, err := c.APIServer.ListCloudSecurityGroup(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.CloudSecurityGroups, nil
}

func (c *Cloud) listCloudPrivatesubnet(e echo.Context) ([]*models.CloudPrivateSubnet, error) {

	spec := baseservices.GetListSpec(e)
	request := &services.ListCloudPrivateSubnetRequest{
		Spec: spec,
	}

	ctx := returnContext(e)
	response, err := c.APIServer.ListCloudPrivateSubnet(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.CloudPrivateSubnets, nil
}

func (c *Cloud) listCredential(e echo.Context) ([]*models.Credential, error) {

	spec := baseservices.GetListSpec(e)
	request := &services.ListCredentialRequest{
		Spec: spec,
	}

	ctx := returnContext(e)
	response, err := c.APIServer.ListCredential(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Credentials, nil
}

func (c *Cloud) getCloudData() (*Data, error) {

	cloudData, err := c.newCloudData()
	if err != nil {
		return nil, err
	}

	err = cloudData.updateData()
	if err != nil {
		return nil, err
	}

	return cloudData, nil

}

func (c *Cloud) getCloudObject() (*models.Cloud, error) {

	ce := c.echo.AcquireContext()
	ctx := returnContext(ce)
	request := new(services.GetCloudRequest)
	request.ID = c.config.CloudID

	cloudResp, err := c.APIServer.GetCloud(ctx, request)
	if err != nil {
		return nil, err
	}

	c.echo.ReleaseContext(ce)
	return cloudResp.GetCloud(), nil

}

func (c *Cloud) newCloudData() (*Data, error) {

	data := Data{}
	data.cloud = c

	cloudObject, err := c.getCloudObject()
	if err != nil {
		return nil, err
	}

	data.cloudProviderInfo = cloudObject
	return &data, nil

}
