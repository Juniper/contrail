package cloud

import (
	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

//Data for cloud provider data
type Data struct {
	cloud               *Cloud
	cloudInfo           *models.Cloud
	projectName         string
	credentials         []*models.Credential
	cloudSubnets        []*models.CloudPrivateSubnet
	cloudSecurityGroups []*models.CloudSecurityGroup
	regions             []*regionData
	accounts            []*accountData
}

type apiServer struct {
	client *client.HTTP
}

type regionData struct {
	info          *models.CloudRegion
	virtualClouds []*virtualCloudData
	apiServer
}

type virtualCloudData struct {
	info      *models.VirtualCloud
	sgs       []*sgData
	instances []*instanceData
	subnets   []*subnetData
	apiServer
}

type instanceData struct {
	info *models.Node
	apiServer
}

type subnetData struct {
	info *models.CloudPrivateSubnet
	apiServer
}

type sgData struct {
	info *models.CloudSecurityGroup
	apiServer
}

type accountData struct {
	info     *models.CloudAccount
	projects []*projectData
	apiServer
}

type projectData struct {
	info  *models.CloudProject
	users []*userData
	apiServer
}

type userData struct {
	info *models.CloudUser
	apiServer
}

func (p *projectData) getProjectObject() (*models.CloudProject, error) {

	ctx := returnContext()
	request := new(services.GetCloudProjectRequest)
	request.ID = p.info.UUID

	projResp, err := p.client.GetCloudProject(ctx, request)
	if err != nil {
		return nil, err
	}
	return projResp.GetCloudProject(), nil

}

func (a *accountData) newProject(project *models.CloudProject) (*projectData, error) {
	proj := &projectData{
		info: project,
		apiServer: apiServer{
			client: a.client,
		},
	}

	projObj, err := proj.getProjectObject()
	if err != nil {
		return nil, err
	}

	proj.info = projObj
	return proj, nil
}

func (a *accountData) updateProjects() error {

	for _, proj := range a.info.CloudProjects {
		newProj, err := a.newProject(proj)
		if err != nil {
			return err
		}

		err = newProj.updateUsers()
		if err != nil {
			return err
		}

		a.projects = append(a.projects, newProj)

	}
	return nil
}

func (a *accountData) getAccountObject() (*models.CloudAccount, error) {

	ctx := returnContext()
	request := new(services.GetCloudAccountRequest)
	request.ID = a.info.UUID

	accResp, err := a.client.GetCloudAccount(ctx, request)
	if err != nil {
		return nil, err
	}
	return accResp.GetCloudAccount(), nil

}

func (d *Data) newAccount(account *models.CloudAccount) (*accountData, error) {

	acc := &accountData{
		info: account,
		apiServer: apiServer{
			client: d.cloud.APIServer,
		},
	}

	accObj, err := acc.getAccountObject()
	if err != nil {
		return nil, err
	}

	acc.info = accObj
	return acc, nil
}

func (d *Data) updateAccounts() error {
	for _, account := range d.cloudInfo.CloudAccounts {
		newAccount, err := d.newAccount(account)
		if err != nil {
			return err
		}

		err = newAccount.updateProjects()
		if err != nil {
			return err
		}

		d.accounts = append(d.accounts, newAccount)
	}
	return nil
}

func (s *subnetData) getPvtSubnetObject() (*models.CloudPrivateSubnet, error) {

	ctx := returnContext()
	request := new(services.GetCloudPrivateSubnetRequest)
	request.ID = s.info.UUID

	subnetResp, err := s.client.GetCloudPrivateSubnet(ctx, request)
	if err != nil {
		return nil, err
	}
	return subnetResp.GetCloudPrivateSubnet(), nil
}

func (v *virtualCloudData) newSubnet(subnet *models.CloudPrivateSubnet) (*subnetData, error) {

	s := &subnetData{
		info: subnet,
		apiServer: apiServer{
			client: v.client,
		},
	}

	subnetObj, err := s.getPvtSubnetObject()
	if err != nil {
		return nil, err
	}

	s.info = subnetObj
	return s, nil
}

func (v *virtualCloudData) updateSubnets() error {

	for _, subnet := range v.info.CloudPrivateSubnets {
		newSubnet, err := v.newSubnet(subnet)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		v.subnets = append(v.subnets, newSubnet)
	}
	return nil
}

func (i *instanceData) getNodeObject() (*models.Node, error) {

	ctx := returnContext()
	request := new(services.GetNodeRequest)
	request.ID = i.info.UUID

	instResp, err := i.client.GetNode(ctx, request)
	if err != nil {
		return nil, err
	}
	return instResp.GetNode(), nil
}

func (v *virtualCloudData) newInstance(instance *models.Node) (*instanceData, error) {

	inst := &instanceData{
		info: instance,
		apiServer: apiServer{
			client: v.client,
		},
	}

	instObj, err := inst.getNodeObject()
	if err != nil {
		return nil, err
	}

	inst.info = instObj
	return inst, nil
}

func (v *virtualCloudData) updateInstances() error {

	for _, instance := range v.info.Nodes {
		newI, err := v.newInstance(instance)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		v.instances = append(v.instances, newI)
	}
	return nil
}

func (sg *sgData) getSGObject() (*models.CloudSecurityGroup, error) {

	ctx := returnContext()
	request := new(services.GetCloudSecurityGroupRequest)
	request.ID = sg.info.UUID

	sgResp, err := sg.client.GetCloudSecurityGroup(ctx, request)
	if err != nil {
		return nil, err
	}
	return sgResp.GetCloudSecurityGroup(), nil
}

func (v *virtualCloudData) newSG(mSG *models.CloudSecurityGroup) (*sgData, error) {

	sg := &sgData{
		info: mSG,
		apiServer: apiServer{
			client: v.client,
		},
	}

	sgObj, err := sg.getSGObject()
	if err != nil {
		return nil, err
	}

	sg.info = sgObj
	return sg, nil
}

func (v *virtualCloudData) updateSGs() error {

	for _, sg := range v.info.CloudSecurityGroups {
		newSG, err := v.newSG(sg)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		v.sgs = append(v.sgs, newSG)
	}
	return nil
}

func (v *virtualCloudData) getVCloudObject() (*models.VirtualCloud, error) {

	ctx := returnContext()
	request := new(services.GetVirtualCloudRequest)
	request.ID = v.info.UUID

	vCloudResp, err := v.client.GetVirtualCloud(ctx, request)
	if err != nil {
		return nil, err
	}
	return vCloudResp.GetVirtualCloud(), nil
}

func (r *regionData) newVCloud(vCloud *models.VirtualCloud) (*virtualCloudData, error) {

	vc := &virtualCloudData{
		info: vCloud,
		apiServer: apiServer{
			client: r.client,
		},
	}

	vCloudObj, err := vc.getVCloudObject()
	vc.info = vCloudObj

	if err != nil {
		return nil, err
	}

	return vc, nil
}

func (r *regionData) updateVClouds() error {

	for _, vc := range r.info.VirtualClouds {
		newVC, err := r.newVCloud(vc)
		if err != nil {
			return err
		}

		err = newVC.updateSGs()
		if err != nil {
			return err
		}

		err = newVC.updateInstances()
		if err != nil {
			return err
		}

		err = newVC.updateSubnets()
		if err != nil {
			return err
		}

		r.virtualClouds = append(r.virtualClouds, newVC)
	}
	return nil
}

func (r *regionData) getRegionObject() (*models.CloudRegion, error) {

	ctx := returnContext()
	request := new(services.GetCloudRegionRequest)
	request.ID = r.info.UUID

	regResp, err := r.client.GetCloudRegion(ctx, request)
	if err != nil {
		return nil, err
	}

	return regResp.GetCloudRegion(), nil

}

func (d *Data) newRegion(region *models.CloudRegion) (*regionData, error) {

	reg := &regionData{
		info: region,
		apiServer: apiServer{
			client: d.cloud.APIServer,
		},
	}

	regObj, err := reg.getRegionObject()
	reg.info = regObj

	if err != nil {
		return nil, err
	}

	return reg, nil
}

func (d *Data) updateRegions() error {

	for _, region := range d.cloudInfo.CloudRegions {
		newRegion, err := d.newRegion(region)
		if err != nil {
			return err
		}

		err = newRegion.updateVClouds()
		if err != nil {
			return err
		}

		d.regions = append(d.regions, newRegion)
	}
	return nil
}

func (u *userData) getUserObject() (*models.CloudUser, error) {

	ctx := returnContext()
	request := new(services.GetCloudUserRequest)
	request.ID = u.info.UUID

	userResp, err := u.client.GetCloudUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return userResp.GetCloudUser(), nil

}

func (p *projectData) newUser(userRef *models.CloudProjectCloudUserRef) (*userData, error) {

	u := &userData{
		info: &models.CloudUser{
			UUID: userRef.UUID,
		},
		apiServer: apiServer{
			client: p.client,
		},
	}

	userObj, err := u.getUserObject()
	if err != nil {
		return nil, err
	}

	u.info = userObj
	return u, nil
}

func (p *projectData) updateUsers() error {

	for _, u := range p.info.CloudUserRefs {
		newUser, err := p.newUser(u)
		if err != nil {
			return err
		}

		p.users = append(p.users, newUser)
	}
	return nil
}

func (c *Cloud) listCloudSecurityGroup() ([]*models.CloudSecurityGroup, error) {

	request := new(services.ListCloudSecurityGroupRequest)
	ctx := returnContext()
	response, err := c.APIServer.ListCloudSecurityGroup(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.CloudSecurityGroups, nil
}

func (c *Cloud) listCloudPrivatesubnet() ([]*models.CloudPrivateSubnet, error) {

	request := new(services.ListCloudPrivateSubnetRequest)

	ctx := returnContext()
	response, err := c.APIServer.ListCloudPrivateSubnet(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.CloudPrivateSubnets, nil
}

func (c *Cloud) listCredential() ([]*models.Credential, error) {

	request := new(services.ListCredentialRequest)
	ctx := returnContext()

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

	ctx := returnContext()
	request := new(services.GetCloudRequest)
	request.ID = c.config.CloudID

	cloudResp, err := c.APIServer.GetCloud(ctx, request)
	if err != nil {
		return nil, err
	}

	return cloudResp.GetCloud(), nil

}

func (c *Cloud) newCloudData() (*Data, error) {

	data := Data{}
	data.cloud = c

	cloudObject, err := c.getCloudObject()
	if err != nil {
		return nil, err
	}

	data.cloudInfo = cloudObject
	return &data, nil

}

func (d *Data) updateData() error {

	subnets, err := d.cloud.listCloudPrivatesubnet()
	if err != nil {
		return err
	}
	d.cloudSubnets = subnets

	sgs, err := d.cloud.listCloudSecurityGroup()
	if err != nil {
		return err
	}
	d.cloudSecurityGroups = sgs

	creds, err := d.cloud.listCredential()
	if err != nil {
		return err
	}
	d.credentials = creds

	err = d.updateRegions()
	if err != nil {
		return err
	}

	err = d.updateAccounts()
	if err != nil {
		return err
	}

	// TODO(madhukar) - Needs to handle multiple projects
	for _, account := range d.accounts {
		for _, project := range account.projects {
			if project.info.Name != "" {
				d.projectName = project.info.Name
				break
			}
		}
	}
	return nil
}
