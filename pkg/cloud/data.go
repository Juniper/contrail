package cloud

import (
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
	users          []*models.CloudUser
	subnets        []*subnetData
	securityGroups []*sgData
	providers      []*providerData
	instances      []*instanceData
}

type apiServer struct {
	client *client.HTTP
}

type providerData struct {
	parentCloud *Data
	info        *models.CloudProvider
	regions     []*regionData
	apiServer
}

type regionData struct {
	parentProvider *providerData
	info           *models.CloudRegion
	virtualClouds  []*virtualCloudData
	apiServer
}

type virtualCloudData struct {
	parentRegion *regionData
	info         *models.VirtualCloud
	sgs          []*sgData
	instances    []*instanceData
	subnets      []*subnetData
	apiServer
}

type instanceData struct {
	parentVC      *virtualCloudData
	info          *models.Node
	roles         []string
	protocolsMode []string
	provision     string
	pvtIntf       *models.Port
	tags          map[string]string
	apiServer
}

type subnetData struct {
	parentVC *virtualCloudData
	info     *models.CloudPrivateSubnet
	apiServer
}

type sgData struct {
	parentVC *virtualCloudData
	info     *models.CloudSecurityGroup
	apiServer
}

func (s *subnetData) getPvtSubnetObject() (*models.CloudPrivateSubnet, error) {

	response := new(services.GetCloudPrivateSubnetResponse)
	_, err := s.client.Read("/cloud-private-subnet/"+s.info.UUID, response)
	if err != nil {
		return nil, err
	}
	return response.GetCloudPrivateSubnet(), nil
}

func (v *virtualCloudData) newSubnet(subnet *models.CloudPrivateSubnet) (*subnetData, error) {

	s := &subnetData{
		parentVC: v,
		info:     subnet,
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

	data := v.parentRegion.parentProvider.parentCloud
	data.subnets = append(data.subnets, v.subnets...)
	return nil
}

func (i *instanceData) getNodeObject() (*models.Node, error) {

	response := new(services.GetNodeResponse)
	_, err := i.client.Read("/node/"+i.info.UUID, response)
	if err != nil {
		return nil, err
	}
	return response.GetNode(), nil
}

func (v *virtualCloudData) newInstance(instance *models.Node) (*instanceData, error) {

	inst := &instanceData{
		parentVC: v,
		info:     instance,
		apiServer: apiServer{
			client: v.client,
		},
	}

	instObj, err := inst.getNodeObject()
	if err != nil {
		return nil, err
	}
	inst.info = instObj

	data := v.parentRegion.parentProvider.parentCloud
	if data.isCloudPrivate() {

		i := inst

		if i.info.ContrailVrouterNodeBackRefs != nil && i.info.KubernetesNodeBackRefs != nil {
			i.roles = append(i.roles, "compute_node")
		} else if i.info.ContrailVrouterNodeBackRefs != nil {
			i.roles = append(i.roles, "vrouter")
		}
		if i.info.ContrailConfigNodeBackRefs != nil {
			i.roles = append(i.roles, "controller")
		}

		if i.info.KubernetesMasterNodeBackRefs != nil {
			i.roles = append(i.roles, "k8s_master")
		}

		if i.info.ContrailMulticloudGWNodeBackRefs != nil {
			i.roles = append(i.roles, "gateway")
		}

		err = inst.updatePvtIntf()
		if err != nil {
			return nil, err
		}

		if inst.info.OpenstackComputeNodeBackRefs != nil {
			inst.provision = strconv.FormatBool(false)
		}
	}

	if inst.provision == "" {
		inst.provision = strconv.FormatBool(true)
	}

	if inst.info.ContrailMulticloudGWNodeBackRefs != nil {
		err := inst.updateProtoModes()
		if err != nil {
			return nil, err
		}
		err = inst.updateMCGWTags()
		if err != nil {
			return nil, err
		}
	}

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

	data := v.parentRegion.parentProvider.parentCloud
	data.instances = append(data.instances, v.instances...)
	return nil
}

func (sg *sgData) getSGObject() (*models.CloudSecurityGroup, error) {

	response := new(services.GetCloudSecurityGroupResponse)
	_, err := sg.client.Read("/cloud-security-group/"+sg.info.UUID, response)
	if err != nil {
		return nil, err
	}
	return response.GetCloudSecurityGroup(), nil
}

func (v *virtualCloudData) newSG(mSG *models.CloudSecurityGroup) (*sgData, error) {

	sg := &sgData{
		parentVC: v,
		info:     mSG,
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

	data := v.parentRegion.parentProvider.parentCloud
	data.securityGroups = append(data.securityGroups, v.sgs...)
	return nil
}

func (v *virtualCloudData) getVCloudObject() (*models.VirtualCloud, error) {

	response := new(services.GetVirtualCloudResponse)
	_, err := v.client.Read("/virtual-cloud/"+v.info.UUID, response)
	if err != nil {
		return nil, err
	}
	return response.GetVirtualCloud(), nil
}

func (r *regionData) newVCloud(vCloud *models.VirtualCloud) (*virtualCloudData, error) {

	vc := &virtualCloudData{
		parentRegion: r,
		info:         vCloud,
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

	response := new(services.GetCloudRegionResponse)
	_, err := r.client.Read("/cloud-region/"+r.info.UUID, response)
	if err != nil {
		return nil, err
	}

	return response.GetCloudRegion(), nil

}

func (p *providerData) newRegion(region *models.CloudRegion) (*regionData, error) {

	reg := &regionData{
		parentProvider: p,
		info:           region,
		apiServer: apiServer{
			client: p.client,
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

	response := new(services.GetCloudProviderResponse)
	_, err := p.client.Read("/cloud-provider/"+p.info.UUID, response)
	if err != nil {
		return nil, err
	}

	return response.GetCloudProvider(), nil

}

func (d *Data) newProvider(provider *models.CloudProvider) (*providerData, error) {

	prov := &providerData{
		parentCloud: d,
		info:        provider,
		apiServer: apiServer{
			client: d.cloud.APIServer,
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

func getUserObject(uuid string,
	apiClient *client.HTTP) (*models.CloudUser, error) {

	response := new(services.GetCloudUserResponse)
	_, err := apiClient.Read("/cloud-user/"+uuid, response)
	if err != nil {
		return nil, err
	}

	return response.GetCloudUser(), nil

}

func (d *Data) updateUsers() error {
	for _, user := range d.info.CloudUserRefs {
		userObj, err := getUserObject(user.UUID, d.cloud.APIServer)
		if err != nil {
			return err
		}

		// Adding logic to handle a ssh key generation if not added as cred ref
		if userObj.CredentialRefs != nil {
			for _, cred := range userObj.CredentialRefs {
				credObj, err := getCredObject(d.cloud.ctx, d.cloud.APIServer, cred.UUID)
				if err != nil {
					return err
				}
				d.credentials = append(d.credentials, credObj)
			}
		}
		d.users = append(d.users, userObj)
	}
	return nil
}

func (i *instanceData) updateProtoModes() error {
	for _, gwNodeRef := range i.info.ContrailMulticloudGWNodeBackRefs {
		response := new(services.GetContrailMulticloudGWNodeResponse)
		_, err := i.client.Read("/contrail-multicloud-gw-node/"+gwNodeRef.UUID,
			response)
		if err != nil {
			return err
		}

		i.protocolsMode = response.GetContrailMulticloudGWNode().ProtocolsMode
		return nil
	}
	return nil
}

func (i *instanceData) updatePvtIntf() error {
	for _, port := range i.info.Ports {
		i.pvtIntf = port
		return nil
	}
	return fmt.Errorf("onprem node %s should have private ip address",
		i.info.Name)
}

func (i *instanceData) updateMCGWTags() error {
	for _, gwNodeRef := range i.info.ContrailMulticloudGWNodeBackRefs {
		response := new(services.GetContrailMulticloudGWNodeResponse)
		_, err := i.client.Read("/contrail-multicloud-gw-node/"+gwNodeRef.UUID,
			response)
		if err != nil {
			return err
		}
		gwNode := response.GetContrailMulticloudGWNode()
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

	cloudObject, err := GetCloud(c.APIServer, c.config.CloudID)
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

	return d.updateUsers()

}

func (d *Data) isCloudCreated() bool {

	status := d.info.ProvisioningState
	if d.cloud.config.Action == createAction && (status == statusNoState || status == "") {
		return false
	}
	d.cloud.log.Infof("Cloud %s already provisioned, STATE: %s", d.info.UUID, status)
	return true
}

func (d *Data) isCloudUpdateRequest() bool {

	status := d.info.ProvisioningState
	if d.cloud.config.Action == updateAction && (status == statusNoState) {
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

	return !d.isCloudPrivate()
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
	return nil, errors.New("cloudUser ref not found with cloud object")
}
