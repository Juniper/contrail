package cloud

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

//Data for cloud provider data
type Data struct {
	cloud          *Cloud
	info           *models.Cloud
	credentials    []*models.Credential
	subnets        []*models.CloudPrivateSubnet
	securityGroups []*models.CloudSecurityGroup
	users          []*models.CloudUser
	providers      []*providerData
	instances      []*instanceData
}

type apiServer struct {
	client *client.HTTP
	ctx    context.Context
}

type providerData struct {
	info    *models.CloudProvider
	regions []*regionData
	apiServer
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
	info          *models.Node
	roles         []string
	protocolsMode []string
	provision     string
	pvtIntf       *models.Port
	tags          map[string]string
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

func (s *subnetData) getPvtSubnetObject() (*models.CloudPrivateSubnet, error) {

	request := new(services.GetCloudPrivateSubnetRequest)
	request.ID = s.info.UUID

	subnetResp, err := s.client.GetCloudPrivateSubnet(s.ctx, request)
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
			ctx:    v.ctx,
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

	request := new(services.GetNodeRequest)
	request.ID = i.info.UUID

	instResp, err := i.client.GetNode(i.ctx, request)
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
			ctx:    v.ctx,
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

		v.instances = append(v.instances, newI)
	}
	return nil
}

func (sg *sgData) getSGObject() (*models.CloudSecurityGroup, error) {

	request := new(services.GetCloudSecurityGroupRequest)
	request.ID = sg.info.UUID

	sgResp, err := sg.client.GetCloudSecurityGroup(sg.ctx, request)
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
			ctx:    v.ctx,
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

	request := new(services.GetVirtualCloudRequest)
	request.ID = v.info.UUID

	vCloudResp, err := v.client.GetVirtualCloud(v.ctx, request)
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
			ctx:    r.ctx,
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

	request := new(services.GetCloudRegionRequest)
	request.ID = r.info.UUID

	regResp, err := r.client.GetCloudRegion(r.ctx, request)
	if err != nil {
		return nil, err
	}

	return regResp.GetCloudRegion(), nil

}

func (p *providerData) newRegion(region *models.CloudRegion) (*regionData, error) {

	reg := &regionData{
		info: region,
		apiServer: apiServer{
			client: p.client,
			ctx:    p.ctx,
		},
	}

	regObj, err := reg.getRegionObject()
	reg.info = regObj

	if err != nil {
		return nil, err
	}

	return reg, nil
}

func (p *providerData) getProviderObject() (*models.CloudProvider, error) {

	request := new(services.GetCloudProviderRequest)
	request.ID = p.info.UUID

	provResp, err := p.client.GetCloudProvider(p.ctx, request)
	if err != nil {
		return nil, err
	}

	return provResp.GetCloudProvider(), nil

}

func (d *Data) newProvider(provider *models.CloudProvider) (*providerData, error) {

	prov := &providerData{
		info: provider,
		apiServer: apiServer{
			client: d.cloud.APIServer,
			ctx:    d.cloud.ctx,
		},
	}

	provObj, err := prov.getProviderObject()
	prov.info = provObj

	if err != nil {
		return nil, err
	}

	return prov, nil
}

func (p *providerData) updateRegions() error {

	for _, region := range p.info.CloudRegions {
		newRegion, err := p.newRegion(region)
		if err != nil {
			return err
		}

		err = newRegion.updateVClouds()
		if err != nil {
			return err
		}

		p.regions = append(p.regions, newRegion)
	}
	return nil
}

func (d *Data) updateProviders() error {
	for _, provider := range d.info.CloudProviders {
		newProvider, err := d.newProvider(provider)
		if err != nil {
			return err
		}

		err = newProvider.updateRegions()
		if err != nil {
			return err
		}

		d.providers = append(d.providers, newProvider)
	}
	return nil
}

func getUserObject(ctx context.Context, uuid string, apiClient *client.HTTP) (*models.CloudUser, error) {

	request := new(services.GetCloudUserRequest)
	request.ID = uuid

	userResp, err := apiClient.GetCloudUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return userResp.GetCloudUser(), nil

}

func (d *Data) updateUsers() error {
	for _, user := range d.info.CloudUserRefs {
		userObj, err := getUserObject(d.cloud.ctx, user.UUID, d.cloud.APIServer)
		if err != nil {
			return err
		}
		d.users = append(d.users, userObj)
	}
	return nil
}

func (d *Data) updateInstances() error {

	for _, prov := range d.providers {
		for _, region := range prov.regions {
			for _, vcloud := range region.virtualClouds {
				for _, instance := range vcloud.instances {

					if d.isCloudPrivate() {
						err := instance.updateRoles()
						if err != nil {
							return err
						}
						err = instance.updatePvtIntf()
						if err != nil {
							return err
						}

						if instance.info.OpenstackComputeNodeBackRefs != nil {
							instance.provision = strconv.FormatBool(false)
						}
					}

					if instance.provision == "" {
						instance.provision = strconv.FormatBool(true)
					}

					if instance.info.ContrailMulticloudGWNodeBackRefs != nil {
						err := instance.updateProtoModes()
						if err != nil {
							return err
						}
						err = instance.updateMCGWTags()
						if err != nil {
							return err
						}
					}
				}
				d.instances = append(d.instances, vcloud.instances...)
			}
		}
	}
	return nil
}

func (d *Data) updatePvtSubnets() error {
	for _, prov := range d.providers {
		for _, region := range prov.regions {
			for _, vcloud := range region.virtualClouds {
				for _, subnet := range vcloud.subnets {
					d.subnets = append(d.subnets, subnet.info)
				}
			}
		}
	}
	return nil
}

func (d *Data) updateSGs() error {
	for _, prov := range d.providers {
		for _, region := range prov.regions {
			for _, vcloud := range region.virtualClouds {
				for _, sg := range vcloud.sgs {
					d.securityGroups = append(d.securityGroups, sg.info)
				}
			}
		}
	}
	return nil
}

func (d *Data) updateCredentials() error {
	for _, user := range d.users {
		for _, cred := range user.CredentialRefs {
			credObj, err := getCredObject(d.cloud.ctx, d.cloud.APIServer, cred.UUID)
			if err != nil {
				return err
			}
			d.credentials = append(d.credentials, credObj)
		}
	}
	return nil
}

func (i *instanceData) updateRoles() error {

	//if i.info.ContrailAnalyticsNodeBackRefs != nil {
	//	i.roles = append(i.roles, "analytics")
	//}

	if i.info.ContrailVrouterNodeBackRefs != nil {
		i.roles = append(i.roles, "vrouter")
	}

	//if i.info.ContrailServiceNodeBackRefs != nil {
	//	i.roles = append(i.roles, "csn")
	//}

	//if i.info.ContrailControlNodeBackRefs != nil {
	//	i.roles = append(i.roles, "control")
	//}

	//if i.info.ContrailConfigDatabaseNodeBackRefs != nil {
	//	i.roles = append(i.roles, "configdb")
	//}

	//if i.info.ContrailWebuiNodeBackRefs != nil {
	//	i.roles = append(i.roles, "webui")
	//}

	if i.info.ContrailConfigNodeBackRefs != nil {
		i.roles = append(i.roles, "controller")
	}

	//if i.info.ContrailAnalyticsDatabaseNodeBackRefs != nil {
	//	i.roles = append(i.roles, "analyticsdb")
	//}

	if i.info.KubernetesNodeBackRefs != nil {
		i.roles = append(i.roles, "k8s_node")
	}

	// role has to be kubemanager

	//if i.info.KubernetesKubemanagerNodeBackRefs != nil {
	//	i.roles = append(i.roles, "k8s_master")
	//}

	if i.info.KubernetesMasterNodeBackRefs != nil {
		i.roles = append(i.roles, "k8s_master")
	}

	if i.info.ContrailMulticloudGWNodeBackRefs != nil {
		i.roles = append(i.roles, "gateway")
	}

	return nil

}

func (i *instanceData) updateProtoModes() error {
	for _, gwNodeRef := range i.info.ContrailMulticloudGWNodeBackRefs {
		gwNodeResp, err := i.client.GetContrailMulticloudGWNode(i.ctx,
			&services.GetContrailMulticloudGWNodeRequest{
				ID: gwNodeRef.UUID,
			},
		)
		if err != nil {
			return err
		}

		i.protocolsMode = gwNodeResp.GetContrailMulticloudGWNode().ProtocolsMode
		return nil
	}
	return nil
}

func (i *instanceData) updatePvtIntf() error {
	for _, port := range i.info.Ports {
		i.pvtIntf = port
		return nil
	}
	return fmt.Errorf("Onprem node %s should have private ip address",
		i.info.Name)
}

func (i *instanceData) updateMCGWTags() error {
	for _, gwNodeRef := range i.info.ContrailMulticloudGWNodeBackRefs {
		gwNodeResp, err := i.client.GetContrailMulticloudGWNode(i.ctx,
			&services.GetContrailMulticloudGWNodeRequest{
				ID: gwNodeRef.UUID,
			},
		)
		if err != nil {
			return err
		}
		gwNode := gwNodeResp.GetContrailMulticloudGWNode()
		if gwNode.Services.KeyValuePair != nil {
			for _, v := range gwNode.Services.KeyValuePair {
				i.tags[v.Key] = v.Value
			}
		}
	}
	return nil
}

func (c *Cloud) getCloudData() (*Data, error) {

	cloudData, err := c.newCloudData()
	if err != nil {
		return nil, err
	}

	err = cloudData.update()
	if err != nil {
		return nil, err
	}

	return cloudData, nil

}

func (c *Cloud) newCloudData() (*Data, error) {

	data := Data{}
	data.cloud = c

	cloudObject, err := GetCloud(c.ctx, c.APIServer, c.config.CloudID)
	if err != nil {
		return nil, err
	}

	data.info = cloudObject
	return &data, nil

}

func (d *Data) update() error {

	err := d.updateProviders()
	if err != nil {
		return err
	}

	err = d.updateInstances()
	if err != nil {
		return err
	}

	err = d.updatePvtSubnets()
	if err != nil {
		return err
	}

	err = d.updateSGs()
	if err != nil {
		return err
	}

	err = d.updateUsers()
	if err != nil {
		return err
	}

	err = d.updateCredentials()
	if err != nil {
		return err
	}

	return nil
}

func (d *Data) isCloudCreated() bool {

	status := d.info.ProvisioningState
	if d.cloud.config.Action == createAction && (status == "NOSTATE" || status == "") {
		return false
	}
	d.cloud.log.Infof("Cloud %s already provisioned, STATE: %s", d.info.UUID, status)
	return true
}

func (d *Data) isCloudUpdated() bool {

	status := d.info.ProvisioningState
	if d.cloud.config.Action == updateAction && (status == "NOSTATE") {
		return false
	}
	return true
}

func (d *Data) isCloudDeleteRequest() bool {
	status := d.info.ProvisioningAction
	if d.cloud.config.Action == updateAction && (status == "DELETE_CLOUD") {
		return true
	}
	return false
}

func (d *Data) isCloudPrivate() bool {

	for _, provider := range d.info.CloudProviders {
		if provider.Type == onPrem {
			return true
		}
	}
	return false
}

func (d *Data) isCloudPublic() bool {

	if !d.isCloudPrivate() {
		return true
	}
	return false
}

func (d *Data) hasProviderAWS() bool {

	for _, prov := range d.providers {
		if prov.info.Type == aws {
			return true
		}
	}
	return false

}

func (d *Data) hasProviderAzure() bool {
	for _, prov := range d.providers {
		if prov.info.Type == azure {
			return true
		}
	}
	return false
}

func (d *Data) getDefaultCloudUser() (*models.CloudUser, error) {
	for _, user := range d.users {
		return user, nil
	}
	return nil, errors.New("CloudUser ref not found with cloud object")
}
