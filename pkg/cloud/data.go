package cloud

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	gatewayRole = "gateway"
	computeRole = "compute"
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
	tors           []*torData
}

type apiServer struct {
	client *client.HTTP
	ctx    context.Context
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
	tors         []*torData
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
	gateway       string
	tags          map[string]string
	apiServer
}

type torData struct {
	parentVC               *virtualCloudData
	info                   *models.PhysicalRouter
	provision              string
	autonomousSystemNumber int
	interfaceName          string
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
		parentVC: v,
		info:     subnet,
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

	data := v.parentRegion.parentProvider.parentCloud
	data.subnets = append(data.subnets, v.subnets...)
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

func (i *torData) getTorObject() (*models.PhysicalRouter, error) {

	request := new(services.GetPhysicalRouterRequest)
	request.ID = i.info.UUID

	torResp, err := i.client.GetPhysicalRouter(i.ctx, request)
	if err != nil {
		return nil, err
	}
	return torResp.GetPhysicalRouter(), nil
}

// nolint: gocyclo
func (v *virtualCloudData) newInstance(instance *models.Node) (*instanceData, error) {

	inst := &instanceData{
		parentVC: v,
		info:     instance,
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
		err := inst.updateProtoModes() //nolint: govet
		if err != nil {
			return nil, err
		}
		err = inst.updateMCGWTags()
		if err != nil {
			return nil, err
		}
		err = inst.updateVrouterGW(gatewayRole)
		if err != nil {
			return nil, err
		}
	}

	if inst.info.ContrailVrouterNodeBackRefs != nil {
		err = inst.updateVrouterGW(computeRole)
		if err != nil {
			return nil, err
		}
	}

	return inst, nil
}

func (v *virtualCloudData) newTorInstance(p *models.PhysicalRouter) (tor *torData, err error) {
	data := v.parentRegion.parentProvider.parentCloud
	if !data.isCloudPrivate() {
		return nil, nil
	}
	tor = &torData{
		parentVC: v,
		info:     p,
		apiServer: apiServer{
			client: v.client,
			ctx:    v.ctx,
		},
	}
	tor.info, err = tor.getTorObject()
	if err != nil {
		return nil, err
	}
	var k []*models.KeyValuePair
	if a := tor.info.GetAnnotations(); a != nil {
		k = a.GetKeyValuePair()
	}
	for _, keyValuePair := range k {
		switch keyValuePair.Key {
		case "autonomous_system":
			tor.autonomousSystemNumber, err = strconv.Atoi(keyValuePair.Value)
		case "interface":
			tor.interfaceName = keyValuePair.Value
		}
	}

	if tor.provision == "" {
		tor.provision = strconv.FormatBool(true)
	}
	return tor, nil
}

func (v *virtualCloudData) updateInstances() error {

	nodes, err := v.getInstancesWithTag(v.info.TagRefs)
	if err != nil {
		return err
	}

	for _, instance := range nodes {
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

func (v *virtualCloudData) updateTorInstances() error {

	physicalRouters, err := v.getTorInstancesWithTag(v.info.TagRefs)
	if err != nil {
		return err
	}

	for _, physicalRouter := range physicalRouters {
		newI, err := v.newTorInstance(physicalRouter)
		if err != nil {
			return err
		}
		v.tors = append(v.tors, newI)
	}

	data := v.parentRegion.parentProvider.parentCloud
	data.tors = append(data.tors, v.tors...)
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

func (v *virtualCloudData) getInstancesWithTag(tagRefs []*models.VirtualCloudTagRef) ([]*models.Node, error) {
	var nodesOfVC []*models.Node

	for _, tag := range tagRefs {
		tagResp, err := v.client.GetTag(v.ctx, &services.GetTagRequest{ID: tag.UUID})
		if err != nil {
			return nil, err
		}
		nodesOfVC = append(nodesOfVC, tagResp.Tag.NodeBackRefs...)
	}
	if len(nodesOfVC) == 0 {
		return nil, errors.New("virtual cloud tag is not used by any nodes")
	}
	return nodesOfVC, nil
}

func (v *virtualCloudData) getTorInstancesWithTag(
	tagRefs []*models.VirtualCloudTagRef) ([]*models.PhysicalRouter, error) {
	var torOfVC []*models.PhysicalRouter

	for _, tag := range tagRefs {
		tagResp, err := v.client.GetTag(v.ctx, &services.GetTagRequest{ID: tag.UUID})
		if err != nil {
			return nil, err
		}
		torOfVC = append(torOfVC, tagResp.Tag.PhysicalRouterBackRefs...)
	}
	if len(torOfVC) == 0 {
		return nil, errors.New("virtual cloud tag is not used by any Tor")
	}
	return torOfVC, nil
}

func (v *virtualCloudData) newSG(mSG *models.CloudSecurityGroup) (*sgData, error) {

	sg := &sgData{
		parentVC: v,
		info:     mSG,
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

	data := v.parentRegion.parentProvider.parentCloud
	data.securityGroups = append(data.securityGroups, v.sgs...)
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
		parentRegion: r,
		info:         vCloud,
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

		err = newVC.updateTorInstances()
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
		parentProvider: p,
		info:           region,
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
		parentCloud: d,
		info:        provider,
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

func getUserObject(ctx context.Context, uuid string,
	apiClient *client.HTTP) (*models.CloudUser, error) {

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
	return errors.New("instance does not have a contrail-multicloud-gw-node ref")
}

func (i *instanceData) updateVrouterGW(role string) error {
	if role == gatewayRole {
		for _, gwNodeRef := range i.info.ContrailMulticloudGWNodeBackRefs {
			response := new(services.GetContrailMulticloudGWNodeResponse)
			_, err := i.client.GetContrailMulticloudGWNode(i.ctx,
				&services.GetContrailMulticloudGWNodeRequest{
					ID: gwNodeRef.UUID,
				},
			)
			if err != nil {
				return err
			}

			i.gateway = response.GetContrailMulticloudGWNode().DefaultGateway
			return nil
		}
	}
	if role == computeRole {
		for _, vrouterNodeRef := range i.info.ContrailVrouterNodeBackRefs {
			response := new(services.GetContrailVrouterNodeResponse)
			_, err := i.client.GetContrailVrouterNode(i.ctx,
				&services.GetContrailVrouterNodeRequest{
					ID: vrouterNodeRef.UUID,
				},
			)
			if err != nil {
				return err
			}

			vrouterNode := response.ContrailVrouterNode
			if vrouterNode.DefaultGateway != "" {
				i.gateway = vrouterNode.DefaultGateway
			} else {
				response := new(services.GetContrailClusterResponse)
				_, err := i.client.GetContrailCluster(i.ctx,
					&services.GetContrailClusterRequest{
						ID: vrouterNode.ParentUUID,
					},
				)
				if err != nil {
					return err
				}
				i.gateway = response.ContrailCluster.DefaultGateway
			}
			return nil
		}
	}
	return errors.New("instance does not have a contrail-multicloud-gw-node ref")

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

	err = d.updateUsers()
	if err != nil {
		return err
	}

	return nil
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
	return nil, errors.New("cloudUser ref not found with cloud object")
}
