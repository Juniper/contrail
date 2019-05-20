package cloud

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	gatewayRole = "gateway"
	computeRole = "compute"
)

type dataInterface interface {
	hasInfo() bool
}

type dataList []dataInterface

//Data for cloud provider data
type Data struct {
	info           *models.Cloud
	credentials    []*models.Credential
	users          []*models.CloudUser
	subnets        []*subnetData
	securityGroups []*sgData
	providers      []*providerData
	instances      []*instanceData
	tors           []*torData
	delRequest     bool
	apiClient
}

type apiClient struct {
	client *client.HTTP
	ctx    context.Context
}

type providerData struct {
	parentCloud *Data
	info        *models.CloudProvider
	regions     []*regionData
	apiClient
}

type regionData struct {
	parentProvider *providerData
	info           *models.CloudRegion
	virtualClouds  []*virtualCloudData
	apiClient
}

type virtualCloudData struct {
	parentRegion *regionData
	info         *models.VirtualCloud
	sgs          []*sgData
	instances    []*instanceData
	tors         []*torData
	subnets      []*subnetData
	apiClient
}

type instanceData struct {
	parentVC      *virtualCloudData
	info          *models.Node
	roles         []string
	protocolsMode []string
	provision     string
	pvtIntf       *models.Port
	gateway       string
	services      []string
	username      string
	apiClient
}

type torData struct {
	parentVC               *virtualCloudData
	info                   *models.PhysicalRouter
	provision              string
	autonomousSystemNumber int
	interfaceNames         []string
	privateSubnets         []string
	apiClient
}

type subnetData struct {
	parentVC *virtualCloudData
	info     *models.CloudPrivateSubnet
	apiClient
}

type sgData struct {
	parentVC *virtualCloudData
	info     *models.CloudSecurityGroup
	apiClient
}

// GetCloudData gets the entire tree data accociated with the cloud
func GetCloudData(ctx context.Context, cloudID string, httpClient *client.HTTP, isDelRequest bool) (*Data, error) {

	ac := &apiClient{
		client: httpClient,
		ctx:    ctx,
	}

	// create new cloud data field
	cloudData, err := newCloudData(cloudID, ac)
	if err != nil {
		return nil, err
	}

	err = cloudData.update(isDelRequest)
	if err != nil {
		return nil, err
	}

	return cloudData, nil

}

// newCloudData create a new data struct and fills data associated
// with cloud's entire data
func newCloudData(cloudID string, ac *apiClient) (*Data, error) {

	data := &Data{
		apiClient: apiClient{
			client: ac.client,
			ctx:    ac.ctx,
		},
	}

	cloudObject, err := GetCloud(ac.ctx, ac.client, cloudID)
	if err != nil {
		return nil, err
	}

	data.info = cloudObject
	return data, nil

}

// update updates all the fields of data
func (d *Data) update(isDelRequest bool) error {

	d.delRequest = isDelRequest
	// update providers field of data struct
	err := d.updateProviders()
	if err != nil {
		return err
	}

	// update user filed of data struct
	err = d.updateUsers()
	if err != nil {
		return err
	}

	return nil
}

// getPvtSubnetObject gets cloud-private-subnet object from db using method provided by api-server
func (s *subnetData) getPvtSubnetObject() (*models.CloudPrivateSubnet, error) {

	request := new(services.GetCloudPrivateSubnetRequest)
	request.ID = s.info.UUID

	subnetResp, err := s.client.GetCloudPrivateSubnet(s.ctx, request)
	if err != nil {
		return nil, err
	}
	return subnetResp.GetCloudPrivateSubnet(), nil
}

// hasInfo checks if subnetData has info
func (s *subnetData) hasInfo() bool {
	return s.info != nil
}

// newSubnet creates new subnetData
func (v *virtualCloudData) newSubnet(subnet *models.CloudPrivateSubnet) (*subnetData, error) {

	s := &subnetData{
		parentVC: v,
		info:     subnet,
		apiClient: apiClient{
			client: v.client,
			ctx:    v.ctx,
		},
	}

	// get cloud-pvt-subnet object
	subnetObj, err := s.getPvtSubnetObject()
	if err != nil {
		return nil, err
	}

	s.info = subnetObj
	return s, nil
}

// updateSubnets update subnets field in the virtualCloudData
func (v *virtualCloudData) updateSubnets() error {

	var unSortedSubnet dataList
	for _, subnet := range v.info.CloudPrivateSubnets {
		// creates new subnetData
		newSubnet, err := v.newSubnet(subnet)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		unSortedSubnet = append(unSortedSubnet, newSubnet)
	}
	// sort subnet data
	sort.Sort(unSortedSubnet)
	for _, sortedSubnet := range unSortedSubnet {
		v.subnets = append(v.subnets, sortedSubnet.(*subnetData))
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

// updateInstType updates instance(node) type for public cloud nodes
func (i *instanceData) updateInstType(instance *models.Node) error {

	if instance.Type == "" {
		instance.Type = "private"
		_, err := i.client.UpdateNode(i.ctx,
			&services.UpdateNodeRequest{
				Node: instance,
			},
		)
		return err
	}
	return nil
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

// hasInfo checks if torData has info
func (i *torData) hasInfo() bool {
	return i.info != nil
}

// nolint: gocyclo
// newInstance creates new instanceData
func (v *virtualCloudData) newInstance(instance *models.Node,
	isDelRequest bool) (*instanceData, error) {

	inst := &instanceData{
		parentVC: v,
		info:     instance,
		apiClient: apiClient{
			client: v.client,
			ctx:    v.ctx,
		},
	}

	// get node schema object
	instObj, err := inst.getNodeObject()
	if err != nil {
		return nil, err
	}
	inst.info = instObj

	// for pvt cloud depending upon the cluster role it assigns a topology understandable role
	// |cluster_role  	  |  topology_role | provision   |
	// |------------------|----------------|-------------|
	// |vrouter & k8snode | compute_node   | true        |
	// |config || control | controller     | false       |
	// |k8s master        | k8s_master     | true        |
	// |mc gateway        | gateway        | true        |
	data := v.parentRegion.parentProvider.parentCloud
	if data.isCloudPrivate() {

		if inst.info.ContrailVrouterNodeBackRefs != nil && inst.info.KubernetesNodeBackRefs != nil {
			inst.roles = append(inst.roles, "compute_node")
		} else if inst.info.ContrailVrouterNodeBackRefs != nil {
			inst.roles = append(inst.roles, "vrouter")
		}
		if inst.info.ContrailConfigNodeBackRefs != nil || inst.info.ContrailControlNodeBackRefs != nil {
			inst.roles = append(inst.roles, "controller")
			inst.provision = strconv.FormatBool(false)
		}

		if inst.info.KubernetesMasterNodeBackRefs != nil {
			inst.roles = append(inst.roles, "k8s_master")
		}

		if inst.info.ContrailMulticloudGWNodeBackRefs != nil {
			inst.roles = append(inst.roles, "gateway")
		}

		err = inst.updatePvtIntf(isDelRequest)
		if err != nil {
			return nil, err
		}

		if inst.info.OpenstackComputeNodeBackRefs != nil {
			inst.provision = strconv.FormatBool(false)
		}
	}

	// default if provision is not set by now, then its true
	if inst.provision == "" {
		inst.provision = strconv.FormatBool(true)
	}

	// update mcgw role paremeters like proto modes, services and vrouter gw
	if inst.info.ContrailMulticloudGWNodeBackRefs != nil {
		err := inst.updateProtoModes(isDelRequest) //nolint: govet
		if err != nil {
			return nil, err
		}
		err = inst.updateMCGWServices()
		if err != nil {
			return nil, err
		}
		if v.parentRegion.parentProvider.parentCloud.isCloudPrivate() {
			err = inst.updateVrouterGW(gatewayRole, isDelRequest)
			if err != nil {
				return nil, err
			}
		}
	}

	// for public cloud nodes, update instance type, username, roles and provision
	if data.isCloudPublic() {
		err = inst.updateInstType(instObj)
		if err != nil {
			return nil, err
		}
		err = inst.updateInstanceUsername()
		if err != nil {
			return nil, err
		}

		// below logis is only for public cloud nodes with role as none
		// to provision non-vrouter instance on public cloud
		if hasCloudRole(inst.info.CloudInfo.Roles, "none") {
			inst.roles = []string{"compute_node"}
			inst.provision = strconv.FormatBool(false)
		}
	}

	// update vrouter gw for compute nodes
	if inst.info.ContrailVrouterNodeBackRefs != nil {
		if v.parentRegion.parentProvider.parentCloud.isCloudPrivate() {
			err = inst.updateVrouterGW(computeRole, isDelRequest)
			if err != nil {
				return nil, err
			}
		}
	}

	return inst, nil
}

// hasCloudRole check if given roles has given noderole
func hasCloudRole(roles []string, nodeRole string) bool {
	for _, role := range roles {
		if role == nodeRole {
			return true
		}
	}
	return false
}

// updateInstanceUsername updates instance username below is the matrix it uses
// operating_system | username |
// -----------------|----------|
// ubuntu16         | ubuntu   |
// ubuntu18         | ubuntu   |
// centos7          | centos   |
// redhat           | redhat   |
// rhel75           | ec2-user |
func (i *instanceData) updateInstanceUsername() error {

	switch i.info.CloudInfo.OperatingSystem {
	case "ubuntu16":
		i.username = "ubuntu"
	case "ubuntu18":
		i.username = "ubuntu"
	case "centos7":
		i.username = "centos"
	case "redhat":
		i.username = "redhat"
	case "rhel75":
		i.username = "ec2-user"
	default:
		return fmt.Errorf("instance %s operating system %s is not valid",
			i.info.UUID, i.info.CloudInfo.OperatingSystem)
	}
	return nil
}

// hasInfo checks if instanceData has info
func (i *instanceData) hasInfo() bool {
	return i.info != nil
}

func (v *virtualCloudData) newTorInstance(p *models.PhysicalRouter) (tor *torData, err error) {
	data := v.parentRegion.parentProvider.parentCloud
	if !data.isCloudPrivate() {
		return nil, nil
	}
	tor = &torData{
		parentVC: v,
		info:     p,
		apiClient: apiClient{
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
			tor.interfaceNames = strings.Split(keyValuePair.Value, ",")
		case "private_subnet":
			tor.privateSubnets = strings.Split(keyValuePair.Value, ",")
		}
	}

	if tor.provision == "" {
		tor.provision = strconv.FormatBool(true)
	}
	return tor, nil
}

// updateInstances updates instances field of virtualCloudData
func (v *virtualCloudData) updateInstances(isdelRequest bool) error {

	var unsortedInstances dataList
	nodes, err := v.getInstancesWithTag(v.info.TagRefs, isdelRequest)
	if err != nil {
		return err
	}

	for _, instance := range nodes {
		// create new instances data
		newI, err := v.newInstance(instance, isdelRequest)
		if err != nil {
			return err
		}
		unsortedInstances = append(unsortedInstances, newI)
	}

	// sort instances data
	sort.Sort(unsortedInstances)
	for _, sortedI := range unsortedInstances {
		v.instances = append(v.instances, sortedI.(*instanceData))
	}

	data := v.parentRegion.parentProvider.parentCloud
	data.instances = append(data.instances, v.instances...)
	return nil
}

// updateTorInstances update tors field of virtualCloudData
func (v *virtualCloudData) updateTorInstances() error {

	var unSortedTOR dataList
	physicalRouters, err := v.getTorInstancesWithTag(v.info.TagRefs)
	if err != nil {
		return err
	}

	for _, physicalRouter := range physicalRouters {
		// create new torData
		newI, err := v.newTorInstance(physicalRouter)
		if err != nil {
			return err
		}
		unSortedTOR = append(unSortedTOR, newI)
	}

	// sort the tor instance
	sort.Sort(unSortedTOR)
	for _, sortedTOR := range unSortedTOR {
		v.tors = append(v.tors, sortedTOR.(*torData))
	}
	data := v.parentRegion.parentProvider.parentCloud
	data.tors = append(data.tors, v.tors...)
	return nil
}

// getSGObject get cloud-security-group schema object from db using method provided by api-server
func (sg *sgData) getSGObject() (*models.CloudSecurityGroup, error) {

	request := new(services.GetCloudSecurityGroupRequest)
	request.ID = sg.info.UUID

	sgResp, err := sg.client.GetCloudSecurityGroup(sg.ctx, request)
	if err != nil {
		return nil, err
	}
	return sgResp.GetCloudSecurityGroup(), nil
}

// getInstancesWithTag get all instances(nodes) which are back referenced by the tags
func (v *virtualCloudData) getInstancesWithTag(tagRefs []*models.VirtualCloudTagRef,
	isDelRequest bool) ([]*models.Node, error) {

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
	tagRefs []*models.VirtualCloudTagRef) ([]*models.PhysicalRouter, error) {
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

// newSG creates new sgData
func (v *virtualCloudData) newSG(mSG *models.CloudSecurityGroup) (*sgData, error) {

	sg := &sgData{
		parentVC: v,
		info:     mSG,
		apiClient: apiClient{
			client: v.client,
			ctx:    v.ctx,
		},
	}

	// get cloud-security-group schema object
	sgObj, err := sg.getSGObject()
	if err != nil {
		return nil, err
	}

	sg.info = sgObj
	return sg, nil
}

// hasInfo checks if sgData has info
func (sg *sgData) hasInfo() bool {
	return sg.info != nil
}

// updateSGs updates sgs field in virtualCloudData
func (v *virtualCloudData) updateSGs() error {

	var unSortedSG dataList
	for _, sg := range v.info.CloudSecurityGroups {
		// create new sgData
		newSG, err := v.newSG(sg)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		unSortedSG = append(unSortedSG, newSG)
	}

	// sort the sgs list
	sort.Sort(unSortedSG)
	for _, sortedVC := range unSortedSG {
		v.sgs = append(v.sgs, sortedVC.(*sgData))
	}
	data := v.parentRegion.parentProvider.parentCloud
	data.securityGroups = append(data.securityGroups, v.sgs...)
	return nil
}

// getSGObject get virtual-cloud schema object from db using method provided by api-server
func (v *virtualCloudData) getVCloudObject() (*models.VirtualCloud, error) {

	request := new(services.GetVirtualCloudRequest)
	request.ID = v.info.UUID

	vCloudResp, err := v.client.GetVirtualCloud(v.ctx, request)
	if err != nil {
		return nil, err
	}
	return vCloudResp.GetVirtualCloud(), nil
}

// updateNodeWithTag updates input node with tag refs which are alo provided as input
func (v *virtualCloudData) updateNodeWithTag(
	nodeUUID string, nTagRefs []*models.NodeTagRef) error {

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

// updateControlNodeWithTag updates control node with tag ref
func (v *virtualCloudData) updateControlNodeWithTag(
	controlNodes []*models.ContrailControlNode) error {

	if controlNodes == nil {
		return fmt.Errorf("cluster does not have control nodes")
	}

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
				nodeTagRef := new(models.NodeTagRef)
				nodeTagRef.UUID = vTagRef.UUID
				nodeTagRef.To = vTagRef.To
				nodeTagRef.Href = vTagRef.Href
				nodeTagRefs = append(nodeTagRefs, nodeTagRef)
			}
			err := v.updateNodeWithTag(nodeRef.UUID, nodeTagRefs)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// updateConfigNodeWithTag updates config node with tag ref
func (v *virtualCloudData) updateConfigNodeWithTag(
	configNodes []*models.ContrailConfigNode) error {

	if configNodes == nil {
		return fmt.Errorf("cluster does not have config nodes")
	}

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
				nodeTagRef := new(models.NodeTagRef)
				nodeTagRef.UUID = vTagRef.UUID
				nodeTagRef.To = vTagRef.To
				nodeTagRef.Href = vTagRef.Href
				nodeTagRefs = append(nodeTagRefs, nodeTagRef)
			}
			err := v.updateNodeWithTag(nodeRef.UUID, nodeTagRefs)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// updateK8sClusterNodesWithTag updates k8s cluster nodes with tags
func (v *virtualCloudData) updateK8sClusterNodesWithTag(
	k8sClusterUUID string) error {

	k8sClusterObj, err := v.client.GetKubernetesCluster(v.ctx,
		&services.GetKubernetesClusterRequest{
			ID: k8sClusterUUID,
		},
	)
	if err != nil {
		return err
	}

	k8sCluster := k8sClusterObj.KubernetesCluster

	// updates kubernetes master node with tag ref
	if k8sCluster.KubernetesMasterNodes != nil {
		err = v.updateK8sMasterNodeWithTag(k8sCluster.GetKubernetesMasterNodes())
		if err != nil {
			return err
		}
	}

	return nil
}

// updateK8sMasterNodeWithTag updates kubernetes master node with tag ref
func (v *virtualCloudData) updateK8sMasterNodeWithTag(
	k8sMasterNodes []*models.KubernetesMasterNode) error {

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
				nodeTagRef := new(models.NodeTagRef)
				nodeTagRef.UUID = vTagRef.UUID
				nodeTagRef.To = vTagRef.To
				nodeTagRef.Href = vTagRef.Href
				nodeTagRefs = append(nodeTagRefs, nodeTagRef)
			}
			err := v.updateNodeWithTag(nodeRef.UUID, nodeTagRefs)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// updateVrouterNodeWithTag updates vrouter node with tag
func (v *virtualCloudData) updateVrouterNodeWithTag(
	vrouterNodes []*models.ContrailVrouterNode) error {

	if vrouterNodes == nil {
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
				nodeTagRef := new(models.NodeTagRef)
				nodeTagRef.UUID = vTagRef.UUID
				nodeTagRef.To = vTagRef.To
				nodeTagRef.Href = vTagRef.Href
				nodeTagRefs = append(nodeTagRefs, nodeTagRef)
			}
			err := v.updateNodeWithTag(nodeRef.UUID, nodeTagRefs)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// updateClusterNodeWithTag updates cluster nodes(with roles config, control, vrouter & k8s)
// with tag refs. The same tags are also referred by virtual-cloud
// nolint: gocyclo
func (v *virtualCloudData) updateClusterNodeWithTag(
	clusterUUID string) error {

	// get contrail-cluster schema object
	ccResp, err := v.client.GetContrailCluster(v.ctx,
		&services.GetContrailClusterRequest{
			ID: clusterUUID,
		},
	)

	if err != nil {
		return err
	}

	contrailCluster := ccResp.GetContrailCluster()
	if contrailCluster.ContrailControlNodes != nil {
		// update control node with tag ref
		err = v.updateControlNodeWithTag(contrailCluster.GetContrailControlNodes())
		if err != nil {
			return err
		}
	}
	if contrailCluster.ContrailConfigNodes != nil {
		// update config node with tag ref
		err = v.updateConfigNodeWithTag(contrailCluster.GetContrailConfigNodes())
		if err != nil {
			return err
		}
	}
	if contrailCluster.ContrailConfigNodes == nil &&
		contrailCluster.ContrailControlNodes == nil {
		return fmt.Errorf("cluster %s does not have control nodes or config nodes",
			contrailCluster.UUID)
	}

	if contrailCluster.KubernetesClusterRefs != nil {
		for _, k8sCluster := range contrailCluster.KubernetesClusterRefs {
			// update k8s roles with tag ref
			err = v.updateK8sClusterNodesWithTag(k8sCluster.UUID)
			if err != nil {
				return err
			}
		}
	}

	return v.updateVrouterNodeWithTag(contrailCluster.GetContrailVrouterNodes())

}

// getMCGWNodeRole returns mcGWNodeRole object back referenced by one of these instances(nodes)
func (v *virtualCloudData) getMCGWNodeRole(
	instances []*models.Node) (*models.ContrailMulticloudGWNode, error) {

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
	return nil, fmt.Errorf(
		"instances list does not have multicloud gw node back refs")
}

// getVCTagsAndUpdateClusterNodes gets tags from virtual-cloud schema object
// and then updated all cluster nodes with that tag
func (v *virtualCloudData) getVCTagsAndUpdateClusterNodes() error {

	// get instances(nodes) associated with the virtual-cloud's tag
	instances, err := v.getInstancesWithTag(v.info.TagRefs,
		v.parentRegion.parentProvider.parentCloud.delRequest)
	if err != nil {
		return err
	}

	// returns mcGWNodeRole objects referrd by one of these instances(nodes)
	mcGWNodeRole, err := v.getMCGWNodeRole(instances)
	if err != nil {
		return err
	}

	// update cluster nodes with the virtual-cloud's tag
	return v.updateClusterNodeWithTag(mcGWNodeRole.GetParentUUID())

}

// hasInfo checks if virtualCloudData has info
func (v *virtualCloudData) hasInfo() bool {
	return v.info != nil
}

// newVCloud creates new virtualCloudData
func (r *regionData) newVCloud(vCloud *models.VirtualCloud) (*virtualCloudData, error) {

	vc := &virtualCloudData{
		parentRegion: r,
		info:         vCloud,
		apiClient: apiClient{
			client: r.client,
			ctx:    r.ctx,
		},
	}

	// get virtual-cloud schema object
	vCloudObj, err := vc.getVCloudObject()
	vc.info = vCloudObj

	if err != nil {
		return nil, err
	}

	return vc, nil
}

// nolint: gocyclo
// updateVClouds updates virtualClouds field in regionData
func (r *regionData) updateVClouds() error {

	var unSortedVCloud dataList
	for _, vc := range r.info.VirtualClouds {
		// create new virtualCloudData
		newVC, err := r.newVCloud(vc)
		if err != nil {
			return err
		}

		// update sgs associated with the virtualCloudData
		err = newVC.updateSGs()
		if err != nil {
			return err
		}

		isDelRequest := r.parentProvider.parentCloud.delRequest

		// In case of private cloud and for non-delete requests, cloud package takes
		// responsibility of associating the virtual cloud tag with nodes tag
		if r.parentProvider.parentCloud.isCloudPrivate() && !isDelRequest {
			err = newVC.getVCTagsAndUpdateClusterNodes()
			if err != nil {
				return err
			}
		}

		// update instances field of virtualCloudData with lates values
		err = newVC.updateInstances(isDelRequest)
		if err != nil {
			return err
		}

		// update TOR istance for pvt cloud
		if r.parentProvider.parentCloud.isCloudPrivate() {
			err = newVC.updateTorInstances()
			if err != nil {
				return err
			}
		}

		// update subnets
		err = newVC.updateSubnets()
		if err != nil {
			return err
		}

		unSortedVCloud = append(unSortedVCloud, newVC)

	}
	// sort virtualcloud data
	sort.Sort(unSortedVCloud)
	for _, sortedVC := range unSortedVCloud {
		r.virtualClouds = append(r.virtualClouds, sortedVC.(*virtualCloudData))
	}
	return nil
}

// getRegionObject gets cloud-region object from db using method provided by api-server
func (r *regionData) getRegionObject() (*models.CloudRegion, error) {

	request := new(services.GetCloudRegionRequest)
	request.ID = r.info.UUID

	regResp, err := r.client.GetCloudRegion(r.ctx, request)
	if err != nil {
		return nil, err
	}

	return regResp.GetCloudRegion(), nil

}

// hasInfo checks if regionData has info
func (r *regionData) hasInfo() bool {
	return r.info != nil
}

// newRegion creates new regionData
func (p *providerData) newRegion(region *models.CloudRegion) (*regionData, error) {

	reg := &regionData{
		parentProvider: p,
		info:           region,
		apiClient: apiClient{
			client: p.client,
			ctx:    p.ctx,
		},
	}

	// get cloud-region schema object
	regObj, err := reg.getRegionObject()
	reg.info = regObj

	if err != nil {
		return nil, err
	}

	return reg, nil
}

// getProviderObject gets cloud-provider object from db using method provided by api-server
func (p *providerData) getProviderObject() (*models.CloudProvider, error) {

	request := new(services.GetCloudProviderRequest)
	request.ID = p.info.UUID

	provResp, err := p.client.GetCloudProvider(p.ctx, request)
	if err != nil {
		return nil, err
	}

	return provResp.GetCloudProvider(), nil

}

// hasInfo checks if providerData has info
func (p *providerData) hasInfo() bool {
	return p.info != nil
}

// newProvider creates new providerData
func (d *Data) newProvider(provider *models.CloudProvider) (*providerData, error) {

	prov := &providerData{
		parentCloud: d,
		info:        provider,
		apiClient: apiClient{
			client: d.client,
			ctx:    d.ctx,
		},
	}

	// get cloud-provider schema object
	provObj, err := prov.getProviderObject()
	prov.info = provObj

	if err != nil {
		return nil, err
	}

	return prov, nil
}

// updateRegions updates regions field in providerData
func (p *providerData) updateRegions() error {

	var unSortedRegion dataList
	for _, region := range p.info.CloudRegions {
		// creates new RegionData
		newRegion, err := p.newRegion(region)
		if err != nil {
			return err
		}

		// update virtual clouds associated with the region
		err = newRegion.updateVClouds()
		if err != nil {
			return err
		}
		unSortedRegion = append(unSortedRegion, newRegion)
	}
	// sort the region list
	sort.Sort(unSortedRegion)
	for _, sortedReg := range unSortedRegion {
		p.regions = append(p.regions, sortedReg.(*regionData))
	}
	return nil
}

// updateProviders updates providers field of data struct
func (d *Data) updateProviders() error {
	var unSortedProvider dataList
	for _, provider := range d.info.CloudProviders {
		// create new providerData
		newProvider, err := d.newProvider(provider)
		if err != nil {
			return err
		}

		// update regions associated with the provider
		err = newProvider.updateRegions()
		if err != nil {
			return err
		}

		unSortedProvider = append(unSortedProvider, newProvider)
	}
	// sort the provider list
	sort.Sort(unSortedProvider)
	for _, sortedProv := range unSortedProvider {
		d.providers = append(d.providers, sortedProv.(*providerData))
	}
	return nil
}

// getUserObject gets cloud-user object from db using method provided by api-server
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

// updateUsers updates users field in data struct
func (d *Data) updateUsers() error {
	for _, user := range d.info.CloudUserRefs {
		// get cloud-user object schema
		userObj, err := getUserObject(d.ctx, user.UUID, d.client)
		if err != nil {
			return err
		}

		// Adding logic to handle a ssh key generation if not added as cred ref
		if userObj.CredentialRefs != nil {
			for _, cred := range userObj.CredentialRefs {
				// get credential object schema
				credObj, err := getCredObject(d.ctx, d.client, cred.UUID)
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

// updateProtoModes updates protocol mode for mc gw instance
// it does by reading the ProtocolsMode parameter from ContrailMulticloudGWNode schema
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

// updateVrouterGW updates vrouter's default gw according to the role associated with it
func (i *instanceData) updateVrouterGW(role string, isDelRequest bool) error {

	switch role {
	case gatewayRole:
		return i.setMultiCloudGWNodeDefaultGW(isDelRequest)
	case computeRole:
		return i.setVrouterNodeDefaultGW(isDelRequest)
	}

	return fmt.Errorf("instance does not have a %s ref", role)

}

// setMultiCloudGWNodeDefaultGW sets vrouter_gateway for mcgw node
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

// setVrouterNodeDefaultGW sets vrouter_gateway for vrouter node
func (i *instanceData) setVrouterNodeDefaultGW(isDelRequest bool) error {

	for _, vrouterNodeRef := range i.info.ContrailVrouterNodeBackRefs {
		response, err := i.client.GetContrailVrouterNode(i.ctx,
			&services.GetContrailVrouterNodeRequest{
				ID: vrouterNodeRef.UUID,
			},
		)
		if err != nil {
			return err
		}

		if response != nil {
			vrouterNode := response.ContrailVrouterNode
			if vrouterNode.DefaultGateway != "" {
				i.gateway = vrouterNode.DefaultGateway
				return nil
			}
			response, err := i.client.GetContrailCluster(i.ctx,
				&services.GetContrailClusterRequest{
					ID: vrouterNode.ParentUUID,
				},
			)
			if err != nil {
				return err
			}
			i.gateway = response.ContrailCluster.DefaultGateway

			if i.gateway == "" && !isDelRequest {
				return fmt.Errorf(
					`default gateway is neither set for vrouter_node uuid: %s
					nor for contrail_cluster uuid: %s`,
					vrouterNodeRef.UUID, vrouterNode.ParentUUID)
			}
			return nil
		}
	}
	if isDelRequest {
		return nil
	}
	return fmt.Errorf(
		"contrailVrouterNodeBackRefs are not present for instance: %s",
		i.info.UUID)
}

// updatePvtIntf updates private intf field,
// by getting it from ports associated with the node
func (i *instanceData) updatePvtIntf(isDelRequest bool) error {
	for _, port := range i.info.Ports {
		i.pvtIntf = port
		return nil
	}
	if isDelRequest {
		return nil
	}
	return fmt.Errorf("onprem node %s should have private ip address",
		i.info.Name)
}

// updateMCGWServices updates mcgw services by reading the ContrailMulticloudGWNode schema
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

// isCloudCreated checks if cloud is already created
func isCloudCreated(action string, d *Data) bool {

	status := d.info.ProvisioningState
	if action == createAction && (status == statusNoState || status == "") {
		return false
	}
	return true
}

// isCloudUpdateRequest checks if request is a valid update request
func isCloudUpdateRequest(action string, d *Data) bool {

	status := d.info.ProvisioningState
	if action == updateAction && (status == statusNoState) {
		return true
	}
	return false
}

// isCloudPrivate checks if cloud's provider is of type private
func (d *Data) isCloudPrivate() bool {

	for _, provider := range d.info.CloudProviders {
		if provider.Type == onPrem {
			return true
		}
	}
	return false
}

// isCloudPublic checks if cloud's provider is of type public(aws/azure/gcp)
func (d *Data) isCloudPublic() bool {

	if !d.isCloudPrivate() {
		return true
	}
	return false
}

// hasProviderAWS checks if provider for cloud is aws
func (d *Data) hasProviderAWS() bool {

	for _, prov := range d.providers {
		if prov.info.Type == aws {
			return true
		}
	}
	return false

}

// hasProviderAzure checks if provider for cloud is azure
func (d *Data) hasProviderAzure() bool {
	for _, prov := range d.providers {
		if prov.info.Type == azure {
			return true
		}
	}
	return false
}

// getDefaultCloudUser returns a default user for cloud
func (d *Data) getDefaultCloudUser() (*models.CloudUser, error) {
	for _, user := range d.users {
		return user, nil
	}
	return nil, errors.New("cloudUser ref not found with cloud object")
}

// getGatewayNodes get nodes with role mc gateway
func (d *Data) getGatewayNodes() []*instanceData {
	gwNodes := []*instanceData{}
	for _, inst := range d.instances {
		if inst.info.ContrailMulticloudGWNodeBackRefs != nil {
			gwNodes = append(gwNodes, inst)
		}
	}
	return gwNodes
}

// Swap compare the items on the given index of dataList and swipes it
// method to satisfy sort interface
func (l dataList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Len length of the dataList
// method to satisfy sort interface
func (l dataList) Len() int {
	return len(l)
}

// Less compares the values of dataList at given index and return a bool
// method to satisfy sort interface
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
