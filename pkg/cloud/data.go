package cloud

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

type dataInterface interface {
	hasInfo() bool
}

type dataList []dataInterface

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
	delRequest     bool
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

type instanceRole string

const (
	computeNodeInstanceRole instanceRole = "compute_node"
	vRouterInstanceRole     instanceRole = "vrouter"
	controllerInstanceRole  instanceRole = "controller"
	k8sMasterInstanceRole   instanceRole = "k8s_master"
	gatewayInstanceRole     instanceRole = "gateway"
	bareInstanceRole        instanceRole = "bare_node"
)

type instanceData struct {
	parentVC      *virtualCloudData
	info          *models.Node
	roles         []instanceRole
	protocolsMode []string
	provision     bool
	pvtIntf       *models.Port
	gateway       string
	services      []string
	username      string
	apiServer
}

type torData struct {
	parentVC               *virtualCloudData
	info                   *models.PhysicalRouter
	provision              bool
	autonomousSystemNumber int
	interfaceNames         []string
	privateSubnets         []string
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
	request := &services.GetCloudPrivateSubnetRequest{ID: s.info.UUID}

	subnetResp, err := s.client.GetCloudPrivateSubnet(s.ctx, request)
	return subnetResp.GetCloudPrivateSubnet(), err
}

func (s *subnetData) hasInfo() bool {
	return s.info != nil
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
	var unSortedSubnet dataList
	for _, subnet := range v.info.CloudPrivateSubnets {
		newSubnet, err := v.newSubnet(subnet)
		if err != nil {
			return err
		}
		unSortedSubnet = append(unSortedSubnet, newSubnet)
	}

	sort.Sort(unSortedSubnet)
	for _, sortedSubnet := range unSortedSubnet {
		v.subnets = append(v.subnets, sortedSubnet.(*subnetData))
	}

	data := v.parentRegion.parentProvider.parentCloud
	data.subnets = append(data.subnets, v.subnets...)
	return nil
}

func (i *instanceData) getNodeObject() (*models.Node, error) {
	request := &services.GetNodeRequest{ID: i.info.UUID}

	instResp, err := i.client.GetNode(i.ctx, request)
	return instResp.GetNode(), err
}

func (i *instanceData) updateInstType(instance *models.Node) error {
	if instance.Type != "" {
		return nil
	}
	instance.Type = "private"
	_, err := i.client.UpdateNode(i.ctx,
		&services.UpdateNodeRequest{
			Node: instance,
		},
	)
	return err
}

func (i *torData) getTorObject() (*models.PhysicalRouter, error) {
	request := &services.GetPhysicalRouterRequest{ID: i.info.UUID}

	torResp, err := i.client.GetPhysicalRouter(i.ctx, request)
	return torResp.GetPhysicalRouter(), err
}

func (i *torData) hasInfo() bool {
	return i.info != nil
}

func (d *Data) providerNames() []string {
	names := make([]string, 0, len(d.providers))
	for _, provider := range d.providers {
		names = append(names, provider.info.Type)
	}
	return names
}

// nolint: gocyclo
func (v *virtualCloudData) newInstance(instance *models.Node, isDelRequest bool) (*instanceData, error) {
	inst := &instanceData{
		parentVC:  v,
		info:      instance,
		provision: true,
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
		updatePrivateInstanceRolesAndProvisionState(inst)
		if err := inst.updatePvtIntf(isDelRequest); err != nil {
			return nil, err
		}
		if err := inst.updateVrouterGW(isDelRequest, inst.info); err != nil {
			return nil, err
		}
	} else {
		updatePublicInstanceRolesAndProvisionState(inst)
		if err := inst.updateInstType(instObj); err != nil {
			return nil, err
		}
		if err := inst.updateInstanceUsername(v.parentRegion.parentProvider.info.Type); err != nil {
			return nil, err
		}
	}

	if inst.info.ContrailMulticloudGWNodeBackRefs != nil {
		if err := inst.updateProtoModes(isDelRequest); err != nil {
			return nil, err
		}
		if err := inst.updateMCGWServices(); err != nil {
			return nil, err
		}
	}

	return inst, nil
}

func updatePrivateInstanceRolesAndProvisionState(i *instanceData) {
	if i.info.ContrailVrouterNodeBackRefs != nil && i.info.KubernetesNodeBackRefs != nil {
		i.roles = append(i.roles, computeNodeInstanceRole)
	} else if i.info.ContrailVrouterNodeBackRefs != nil {
		i.roles = append(i.roles, vRouterInstanceRole)
	}
	if i.info.ContrailConfigNodeBackRefs != nil || i.info.ContrailControlNodeBackRefs != nil {
		i.roles = append(i.roles, controllerInstanceRole)
		i.provision = false
	}
	if i.info.KubernetesMasterNodeBackRefs != nil {
		i.roles = append(i.roles, k8sMasterInstanceRole)
	}
	if i.info.ContrailMulticloudGWNodeBackRefs != nil {
		i.roles = append(i.roles, gatewayInstanceRole)
	}
	if i.info.OpenstackComputeNodeBackRefs != nil {
		i.provision = false
	}
}
func updatePublicInstanceRolesAndProvisionState(i *instanceData) {
	if hasCloudRole(i.info.CloudInfo.Roles, bareInstanceRole) {
		i.roles = []instanceRole{bareInstanceRole}
	}
}

func hasCloudRole(roles []string, nodeRole instanceRole) bool {
	for _, role := range roles {
		if role == string(nodeRole) {
			return true
		}
	}
	return false
}

func (i *instanceData) updateInstanceUsername(providerType string) error {
	switch i.info.CloudInfo.OperatingSystem {
	case "ubuntu18":
		i.username = "ubuntu"
	case "centos7":
		i.username = "centos"
	case "rhel7":
		if providerType == AWS {
			i.username = "ec2-user"
		} else {
			i.username = "redhat"
		}
	default:
		return fmt.Errorf("instance %s operating system %s is not valid",
			i.info.UUID, i.info.CloudInfo.OperatingSystem)
	}
	return nil
}

func (i *instanceData) hasInfo() bool {
	return i.info != nil
}

func (v *virtualCloudData) newTorInstance(p *models.PhysicalRouter) (tor *torData, err error) {
	if !v.parentRegion.parentProvider.parentCloud.isCloudPrivate() {
		return nil, nil
	}
	tor = &torData{
		parentVC:  v,
		info:      p,
		provision: true,
		apiServer: apiServer{
			client: v.client,
			ctx:    v.ctx,
		},
	}
	if tor.info, err = tor.getTorObject(); err != nil {
		return nil, err
	}
	for _, keyValuePair := range tor.info.GetAnnotations().GetKeyValuePair() {
		switch keyValuePair.Key {
		case "autonomous_system":
			if tor.autonomousSystemNumber, err = strconv.Atoi(keyValuePair.Value); err != nil {
				return nil, errors.Wrap(err, "fail to parse autonomous_system annotation")
			}
		case "interface":
			tor.interfaceNames = strings.Split(keyValuePair.Value, ",")
		case "private_subnet":
			tor.privateSubnets = strings.Split(keyValuePair.Value, ",")
		}
	}

	return tor, nil
}

func (v *virtualCloudData) updateInstances(isdelRequest bool) error {
	var unsortedInstances dataList
	nodes, err := v.getInstancesWithTag(v.info.TagRefs, isdelRequest)
	if err != nil {
		return err
	}

	for _, instance := range nodes {
		newI, err := v.newInstance(instance, isdelRequest)
		if err != nil {
			return err
		}
		unsortedInstances = append(unsortedInstances, newI)
	}
	sort.Sort(unsortedInstances)
	for _, sortedI := range unsortedInstances {
		v.instances = append(v.instances, sortedI.(*instanceData))
	}

	data := v.parentRegion.parentProvider.parentCloud
	data.instances = append(data.instances, v.instances...)
	return nil
}

func (v *virtualCloudData) updateTorInstances() error {
	var unSortedTOR dataList
	physicalRouters, err := v.getTorInstancesWithTag(v.info.TagRefs)
	if err != nil {
		return err
	}

	for _, physicalRouter := range physicalRouters {
		newI, err := v.newTorInstance(physicalRouter)
		if err != nil {
			return err
		}
		unSortedTOR = append(unSortedTOR, newI)
	}

	sort.Sort(unSortedTOR)
	for _, sortedTOR := range unSortedTOR {
		v.tors = append(v.tors, sortedTOR.(*torData))
	}
	data := v.parentRegion.parentProvider.parentCloud
	data.tors = append(data.tors, v.tors...)
	return nil
}

func (sg *sgData) getSGObject() (*models.CloudSecurityGroup, error) {
	request := &services.GetCloudSecurityGroupRequest{ID: sg.info.UUID}

	sgResp, err := sg.client.GetCloudSecurityGroup(sg.ctx, request)
	return sgResp.GetCloudSecurityGroup(), err
}

func (v *virtualCloudData) getInstancesWithTag(
	tagRefs []*models.VirtualCloudTagRef, isDelRequest bool,
) ([]*models.Node, error) {
	var nodesOfVC []*models.Node

	for _, tag := range tagRefs {
		tagResp, err := v.client.GetTag(v.ctx, &services.GetTagRequest{ID: tag.UUID})
		if err != nil {
			return nil, err
		}
		nodesOfVC = append(nodesOfVC, tagResp.Tag.NodeBackRefs...)
	}
	if len(nodesOfVC) == 0 && !isDelRequest {
		return nil, errors.New("virtual cloud tag is not used by any nodes")
	}

	for i, node := range nodesOfVC {
		nodeResp, err := v.client.GetNode(v.ctx,
			&services.GetNodeRequest{
				ID: node.UUID,
			},
		)
		if err != nil {
			return nil, err
		}
		nodesOfVC[i] = nodeResp.Node
	}
	return nodesOfVC, nil
}

func (v *virtualCloudData) getTorInstancesWithTag(
	tagRefs []*models.VirtualCloudTagRef,
) ([]*models.PhysicalRouter, error) {
	var torOfVC []*models.PhysicalRouter

	for _, tag := range tagRefs {
		tagResp, err := v.client.GetTag(v.ctx, &services.GetTagRequest{ID: tag.UUID})
		if err != nil {
			return nil, err
		}
		torOfVC = append(torOfVC, tagResp.Tag.PhysicalRouterBackRefs...)
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

func (sg *sgData) hasInfo() bool {
	return sg.info != nil
}

func (v *virtualCloudData) updateSGs() error {
	var unSortedSG dataList
	for _, sg := range v.info.CloudSecurityGroups {
		newSG, err := v.newSG(sg)
		if err != nil {
			return err
		}
		unSortedSG = append(unSortedSG, newSG)
	}

	sort.Sort(unSortedSG)
	for _, sortedVC := range unSortedSG {
		v.sgs = append(v.sgs, sortedVC.(*sgData))
	}
	data := v.parentRegion.parentProvider.parentCloud
	data.securityGroups = append(data.securityGroups, v.sgs...)
	return nil
}

func (v *virtualCloudData) getVCloudObject() (*models.VirtualCloud, error) {
	request := &services.GetVirtualCloudRequest{ID: v.info.UUID}

	vCloudResp, err := v.client.GetVirtualCloud(v.ctx, request)
	return vCloudResp.GetVirtualCloud(), err
}

func (v *virtualCloudData) updateNodeWithTag(nodeUUID string, nTagRefs []*models.NodeTagRef) error {
	getNodeResp, err := v.client.GetNode(v.ctx,
		&services.GetNodeRequest{
			ID: nodeUUID,
		},
	)
	if err != nil {
		return err
	}

	for _, nTagRef := range nTagRefs {
		getNodeResp.Node.AddTagRef(nTagRef)
	}

	_, err = v.client.UpdateNode(v.ctx,
		&services.UpdateNodeRequest{
			Node: getNodeResp.Node,
		},
	)
	return err
}

func (v *virtualCloudData) updateControlNodeWithTag(controlNodes []*models.ContrailControlNode) error {
	for _, controlNode := range controlNodes {
		getControlResp, err := v.client.GetContrailControlNode(v.ctx,
			&services.GetContrailControlNodeRequest{
				ID: controlNode.UUID,
			},
		)
		if err != nil {
			return err
		}
		for _, nodeRef := range getControlResp.ContrailControlNode.NodeRefs {
			var nodeTagRefs []*models.NodeTagRef
			for _, vTagRef := range v.info.TagRefs {
				nodeTagRefs = append(nodeTagRefs, newNodeTagRef(vTagRef))
			}
			if err := v.updateNodeWithTag(nodeRef.UUID, nodeTagRefs); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *virtualCloudData) updateConfigNodeWithTag(configNodes []*models.ContrailConfigNode) error {
	for _, configNode := range configNodes {
		getConfigResp, err := v.client.GetContrailConfigNode(v.ctx,
			&services.GetContrailConfigNodeRequest{
				ID: configNode.UUID,
			},
		)
		if err != nil {
			return err
		}
		for _, nodeRef := range getConfigResp.ContrailConfigNode.NodeRefs {
			var nodeTagRefs []*models.NodeTagRef
			for _, vTagRef := range v.info.TagRefs {
				nodeTagRefs = append(nodeTagRefs, newNodeTagRef(vTagRef))
			}
			if err := v.updateNodeWithTag(nodeRef.UUID, nodeTagRefs); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *virtualCloudData) updateK8sClusterNodesWithTag(k8sClusterUUID string) error {
	k8sClusterObj, err := v.client.GetKubernetesCluster(v.ctx,
		&services.GetKubernetesClusterRequest{
			ID: k8sClusterUUID,
		},
	)
	if err != nil {
		return err
	}

	return v.updateK8sMasterNodeWithTag(k8sClusterObj.KubernetesCluster.GetKubernetesMasterNodes())
}

func (v *virtualCloudData) updateK8sMasterNodeWithTag(k8sMasterNodes []*models.KubernetesMasterNode) error {
	for _, k8sMaster := range k8sMasterNodes {
		getK8sMasterNodeResp, err := v.client.GetKubernetesMasterNode(v.ctx,
			&services.GetKubernetesMasterNodeRequest{
				ID: k8sMaster.UUID,
			},
		)
		if err != nil {
			return err
		}
		for _, nodeRef := range getK8sMasterNodeResp.KubernetesMasterNode.NodeRefs {
			var nodeTagRefs []*models.NodeTagRef
			for _, vTagRef := range v.info.TagRefs {
				nodeTagRefs = append(nodeTagRefs, newNodeTagRef(vTagRef))
			}
			if err := v.updateNodeWithTag(nodeRef.UUID, nodeTagRefs); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *virtualCloudData) updateVrouterNodeWithTag(vrouterNodes []*models.ContrailVrouterNode) error {
	if len(vrouterNodes) == 0 {
		return fmt.Errorf("cluster does not have vrouter nodes")
	}

	for _, vrouterNode := range vrouterNodes {
		getVrouterResp, err := v.client.GetContrailVrouterNode(v.ctx,
			&services.GetContrailVrouterNodeRequest{
				ID: vrouterNode.UUID,
			},
		)
		if err != nil {
			return err
		}
		for _, nodeRef := range getVrouterResp.ContrailVrouterNode.NodeRefs {
			var nodeTagRefs []*models.NodeTagRef
			for _, vTagRef := range v.info.TagRefs {
				nodeTagRefs = append(nodeTagRefs, newNodeTagRef(vTagRef))
			}
			if err := v.updateNodeWithTag(nodeRef.UUID, nodeTagRefs); err != nil {
				return err
			}
		}
	}
	return nil
}

func newNodeTagRef(v *models.VirtualCloudTagRef) *models.NodeTagRef {
	return &models.NodeTagRef{UUID: v.UUID, To: v.To, Href: v.Href}
}

// nolint: gocyclo
func (v *virtualCloudData) updateClusterNodeWithTag(mcGWNode *models.ContrailMulticloudGWNode) error {
	ccResp, err := v.client.GetContrailCluster(v.ctx,
		&services.GetContrailClusterRequest{
			ID: mcGWNode.ParentUUID,
		},
	)

	if err != nil {
		return err
	}

	contrailCluster := ccResp.GetContrailCluster()
	if err = v.updateControlNodeWithTag(contrailCluster.GetContrailControlNodes()); err != nil {
		return err
	}
	if err = v.updateConfigNodeWithTag(contrailCluster.GetContrailConfigNodes()); err != nil {
		return err
	}
	if len(contrailCluster.ContrailConfigNodes) == 0 &&
		len(contrailCluster.ContrailControlNodes) == 0 {
		return fmt.Errorf("cluster %s does not have control nodes or config nodes",
			contrailCluster.UUID)
	}

	if contrailCluster.KubernetesClusterRefs != nil {
		for _, k8sCluster := range contrailCluster.KubernetesClusterRefs {
			if err = v.updateK8sClusterNodesWithTag(k8sCluster.UUID); err != nil {
				return err
			}
		}
	}

	return v.updateVrouterNodeWithTag(contrailCluster.GetContrailVrouterNodes())
}

func (v *virtualCloudData) getMCGWNodeRole(instances []*models.Node) (*models.ContrailMulticloudGWNode, error) {
	for _, i := range instances {
		if i.GetContrailMulticloudGWNodeBackRefs() != nil {
			mcGWNodeRefs := i.GetContrailMulticloudGWNodeBackRefs()
			for _, m := range mcGWNodeRefs {
				getMCGWResp, err := v.client.GetContrailMulticloudGWNode(v.ctx,
					&services.GetContrailMulticloudGWNodeRequest{
						ID: m.UUID,
					},
				)
				return getMCGWResp.GetContrailMulticloudGWNode(), err
			}
		}
	}
	return nil, fmt.Errorf("instances list does not have multicloud gw node back refs")
}

func (v *virtualCloudData) getTagsAndUpdateClusterNodes() error {
	instances, err := v.getInstancesWithTag(v.info.TagRefs, v.parentRegion.parentProvider.parentCloud.delRequest)
	if err != nil {
		return err
	}

	mcGWNodeRole, err := v.getMCGWNodeRole(instances)
	if err != nil {
		return err
	}

	return v.updateClusterNodeWithTag(mcGWNodeRole)
}

func (v *virtualCloudData) hasInfo() bool {
	return v.info != nil
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
	if err != nil {
		return nil, err
	}

	vc.info = vCloudObj

	return vc, nil
}

// nolint: gocyclo
func (r *regionData) updateVClouds() error {
	var unSortedVCloud dataList
	for _, vc := range r.info.VirtualClouds {
		newVC, err := r.newVCloud(vc)
		if err != nil {
			return err
		}

		if err = newVC.updateSGs(); err != nil {
			return err
		}

		isDelRequest := r.parentProvider.parentCloud.delRequest

		if r.parentProvider.parentCloud.isCloudPrivate() && !isDelRequest {
			if err = newVC.getTagsAndUpdateClusterNodes(); err != nil {
				return err
			}
		}

		if err = newVC.updateInstances(isDelRequest); err != nil {
			return err
		}

		if r.parentProvider.parentCloud.isCloudPrivate() {
			err = newVC.updateTorInstances()
			if err != nil {
				return err
			}
		}

		if err := newVC.updateSubnets(); err != nil {
			return err
		}

		unSortedVCloud = append(unSortedVCloud, newVC)
	}
	sort.Sort(unSortedVCloud)
	for _, sortedVC := range unSortedVCloud {
		r.virtualClouds = append(r.virtualClouds, sortedVC.(*virtualCloudData))
	}
	return nil
}

func (r *regionData) getRegionObject() (*models.CloudRegion, error) {
	request := &services.GetCloudRegionRequest{ID: r.info.UUID}

	regResp, err := r.client.GetCloudRegion(r.ctx, request)
	if err != nil {
		return nil, err
	}

	return regResp.GetCloudRegion(), nil
}

func (r *regionData) hasInfo() bool {
	return r.info != nil
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

	return reg, err
}

func (p *providerData) getProviderObject() (*models.CloudProvider, error) {
	request := &services.GetCloudProviderRequest{ID: p.info.UUID}

	provResp, err := p.client.GetCloudProvider(p.ctx, request)
	return provResp.GetCloudProvider(), err
}

func (p *providerData) hasInfo() bool {
	return p.info != nil
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

	return prov, err
}

func (p *providerData) updateRegions() error {
	var unSortedRegion dataList
	for _, region := range p.info.CloudRegions {
		newRegion, err := p.newRegion(region)
		if err != nil {
			return err
		}

		if err := newRegion.updateVClouds(); err != nil {
			return err
		}
		unSortedRegion = append(unSortedRegion, newRegion)
	}

	sort.Sort(unSortedRegion)
	for _, sortedReg := range unSortedRegion {
		p.regions = append(p.regions, sortedReg.(*regionData))
	}
	return nil
}

func (d *Data) updateProviders() error {
	var unSortedProvider dataList
	for _, provider := range d.info.CloudProviders {
		newProvider, err := d.newProvider(provider)
		if err != nil {
			return err
		}

		if err = newProvider.updateRegions(); err != nil {
			return err
		}

		unSortedProvider = append(unSortedProvider, newProvider)
	}
	sort.Sort(unSortedProvider)
	for _, sortedProv := range unSortedProvider {
		d.providers = append(d.providers, sortedProv.(*providerData))
	}
	return nil
}

func getUserObject(ctx context.Context, uuid string, apiClient *client.HTTP) (*models.CloudUser, error) {
	request := &services.GetCloudUserRequest{ID: uuid}

	userResp, err := apiClient.GetCloudUser(ctx, request)
	return userResp.GetCloudUser(), err
}

func (d *Data) updateUsers() error {
	for _, user := range d.info.CloudUserRefs {
		userObj, err := getUserObject(d.cloud.ctx, user.UUID, d.cloud.APIServer)
		if err != nil {
			return err
		}

		// Adding logic to handle a ssh key generation if not added as cred ref
		for _, cred := range userObj.CredentialRefs {
			credObj, err := getCredObject(d.cloud.ctx, d.cloud.APIServer, cred.UUID)
			if err != nil {
				return err
			}
			d.credentials = append(d.credentials, credObj)
		}
		d.users = append(d.users, userObj)
	}
	return nil
}

func (i *instanceData) updateProtoModes(isDelRequest bool) error {
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
	if isDelRequest {
		return nil
	}
	return errors.New("instance does not have a contrail-multicloud-gw-node ref")
}

func (i *instanceData) updateVrouterGW(isDelRequest bool, info *models.Node) error {
	if info.ContrailMulticloudGWNodeBackRefs != nil {
		if err := i.setMultiCloudGWNodeDefaultGW(isDelRequest); err != nil {
			return err
		}
	}
	if info.ContrailVrouterNodeBackRefs != nil {
		if err := i.setVrouterNodeDefaultGW(isDelRequest); err != nil {
			return err
		}
	}
	return nil
}

func (i *instanceData) setMultiCloudGWNodeDefaultGW(isDelRequest bool) error {
	for _, gwNodeRef := range i.info.ContrailMulticloudGWNodeBackRefs {
		response, err := i.client.GetContrailMulticloudGWNode(i.ctx,
			&services.GetContrailMulticloudGWNodeRequest{
				ID: gwNodeRef.UUID,
			},
		)
		if err != nil {
			return err
		}

		if response != nil {
			i.gateway = response.ContrailMulticloudGWNode.DefaultGateway
		}
		if i.gateway == "" && !isDelRequest {
			return fmt.Errorf(
				"default gateway is not set for contrail_multicloud_gw_node uuid: %s",
				gwNodeRef.UUID)
		}
		return nil
	}
	if isDelRequest {
		return nil
	}
	return fmt.Errorf(
		"contrailMulticloudGWNodeBackRefs are not present for instance: %s",
		i.info.UUID)
}

func (i *instanceData) setVrouterNodeDefaultGW(isDelRequest bool) error {
	for _, vrouterNodeRef := range i.info.ContrailVrouterNodeBackRefs {
		vrouterNodeResponse, err := i.client.GetContrailVrouterNode(i.ctx,
			&services.GetContrailVrouterNodeRequest{
				ID: vrouterNodeRef.UUID,
			},
		)
		if err != nil {
			return err
		}

		if vrouterNodeResponse == nil {
			continue
		}
		vrouterNode := vrouterNodeResponse.ContrailVrouterNode
		if vrouterNode.DefaultGateway != "" {
			i.gateway = vrouterNode.DefaultGateway
			return nil
		}
		contrailClusterResponse, err := i.client.GetContrailCluster(i.ctx,
			&services.GetContrailClusterRequest{
				ID: vrouterNode.ParentUUID,
			},
		)
		if err != nil {
			return err
		}
		i.gateway = contrailClusterResponse.ContrailCluster.DefaultGateway

		if i.gateway == "" && !isDelRequest {
			return fmt.Errorf(
				`default gateway is neither set for vrouter_node uuid: %s
			nor for contrail_cluster uuid: %s`,
				vrouterNodeRef.UUID, vrouterNode.ParentUUID)
		}
		return nil
	}
	if isDelRequest {
		return nil
	}
	return fmt.Errorf(
		"contrailVrouterNodeBackRefs are not present for instance: %s",
		i.info.UUID)
}

func (i *instanceData) updatePvtIntf(isDelRequest bool) error {
	if len(i.info.Ports) > 0 {
		i.pvtIntf = i.info.Ports[0]
		return nil
	}
	if isDelRequest {
		return nil
	}
	return fmt.Errorf("onprem node %s should have private ip address",
		i.info.Name)
}

func (i *instanceData) updateMCGWServices() error {
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
		if gwNode.Services != nil {
			for _, v := range gwNode.Services {
				i.services = append(i.services, v)
			}
		}
	}
	return nil
}

func (d *Data) update(isDelRequest bool) error {
	d.delRequest = isDelRequest
	err := d.updateProviders()
	if err != nil {
		return err
	}
	return d.updateUsers()
}

func (d *Data) isCloudPrivate() bool {
	for _, provider := range d.info.CloudProviders {
		if provider.Type == onPrem {
			return true
		}
	}
	return false
}

func (d *Data) getGatewayNodes() []*instanceData {
	gwNodes := []*instanceData{}
	for _, inst := range d.instances {
		if inst.info.ContrailMulticloudGWNodeBackRefs != nil {
			gwNodes = append(gwNodes, inst)
		}
	}
	return gwNodes
}

func (l dataList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l dataList) Len() int {
	return len(l)
}

func (l dataList) Less(i, j int) bool {
	switch l[i].(type) {
	case *providerData:
		return l[i].(*providerData).info.Name < l[j].(*providerData).info.Name
	case *regionData:
		return l[i].(*regionData).info.Name < l[j].(*regionData).info.Name
	case *virtualCloudData:
		return l[i].(*virtualCloudData).info.Name < l[j].(*virtualCloudData).info.Name
	case *instanceData:
		return l[i].(*instanceData).info.Name < l[j].(*instanceData).info.Name
	case *torData:
		return l[i].(*torData).info.Name < l[j].(*torData).info.Name
	case *subnetData:
		return l[i].(*subnetData).info.Name < l[j].(*subnetData).info.Name
	case *sgData:
		return l[i].(*sgData).info.Name < l[j].(*sgData).info.Name
	default:
		return false
	}
}
