package cluster

import (
	"context"

	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// DataStore interface to store cluster data
type DataStore interface {
	updateClusterDetails(string, *Cluster) error
	updateNodeDetails(*Cluster) error
	addNode(*models.Node)
	addCredential(*models.Credential)
	addKeypair(*models.Keypair)
}

// OpenstackData is the representation of openstack cluster details.
type OpenstackData struct {
	clusterInfo  *models.OpenstackCluster
	nodesInfo    []*models.Node
	keypairsInfo []*models.Keypair
	credsInfo    []*models.Credential
	Reader       services.ReadService
}

// KubernetesData is the representation of kubernetes cluster details.
type KubernetesData struct {
	clusterInfo  *models.KubernetesCluster
	nodesInfo    []*models.Node
	keypairsInfo []*models.Keypair
	credsInfo    []*models.Credential
	Reader       services.ReadService
}

// VCenterData is the representation of VCenter details.
type VCenterData struct {
	clusterInfo  *models.VCenter
	nodesInfo    []*models.Node
	keypairsInfo []*models.Keypair
	credsInfo    []*models.Credential
	Reader       services.ReadService
}

// AppformixData is the representation of appformix cluster details.
type AppformixData struct {
	clusterInfo  *models.AppformixCluster
	nodesInfo    []*models.Node
	keypairsInfo []*models.Keypair
	credsInfo    []*models.Credential
	Reader       services.ReadService
}

// XflowData is the representation of Appformix Flows cluster details.
type XflowData struct {
	ClusterInfo *models.AppformixFlows
	NodesInfo   map[string]*models.Node
}

// NewXflowData creates empty XflowData instance.
func NewXflowData() *XflowData {
	return &XflowData{
		NodesInfo: make(map[string]*models.Node),
	}
}

// Data is the representation of cluster details.
type Data struct {
	clusterInfo           *models.ContrailCluster
	nodesInfo             []*models.Node
	keypairsInfo          []*models.Keypair
	credsInfo             []*models.Credential
	cloudInfo             []*models.Cloud
	openstackClusterData  []*OpenstackData
	vcenterData           []*VCenterData
	kubernetesClusterData []*KubernetesData
	appformixClusterData  []*AppformixData
	xflowData             []*XflowData
	Reader                services.ReadService
	// TODO (ijohnson): Add gce/aws/kvm info
}

func (a *AppformixData) addKeypair(keypair *models.Keypair) {
	a.keypairsInfo = append(a.keypairsInfo, keypair)
}

func (a *AppformixData) addCredential(cred *models.Credential) {
	a.credsInfo = append(a.credsInfo, cred)
}

func (a *AppformixData) addNode(node *models.Node) {
	a.nodesInfo = append(a.nodesInfo, node)
}

func (a *AppformixData) interfaceToAppformixControllerNode(
	appformixControllerNodes []*models.AppformixControllerNode, c *Cluster) error {
	a.clusterInfo.AppformixControllerNodes = nil
	for _, appformixControllerNode := range appformixControllerNodes {
		appformixControllerNodeInfo := models.InterfaceToAppformixControllerNode(appformixControllerNode)
		// Read appformixController role node to get the node refs information
		appformixControllerNodeData, err := c.getResource(
			defaultAppformixControllerNodeResPath, appformixControllerNodeInfo.UUID)
		if err != nil {
			return err
		}
		appformixControllerNodeInfo = models.InterfaceToAppformixControllerNode(
			appformixControllerNodeData)
		a.clusterInfo.AppformixControllerNodes = append(
			a.clusterInfo.AppformixControllerNodes, appformixControllerNodeInfo)
	}
	return nil
}

func (a *AppformixData) interfaceToAppformixBareHostNode(
	appformixBareHostNodes []*models.AppformixBareHostNode, c *Cluster) error {
	a.clusterInfo.AppformixBareHostNodes = nil
	for _, appformixBareHostNode := range appformixBareHostNodes {
		appformixBareHostNodeInfo := models.InterfaceToAppformixBareHostNode(appformixBareHostNode)
		// Read appformixBareHost role node to get the node refs information
		appformixBareHostNodeData, err := c.getResource(
			defaultAppformixBareHostNodeResPath, appformixBareHostNodeInfo.UUID)
		if err != nil {
			return err
		}
		appformixBareHostNodeInfo = models.InterfaceToAppformixBareHostNode(
			appformixBareHostNodeData)
		a.clusterInfo.AppformixBareHostNodes = append(
			a.clusterInfo.AppformixBareHostNodes, appformixBareHostNodeInfo)
	}
	return nil
}

func (a *AppformixData) interfaceToAppformixOpenstackNode(
	appformixOpenstackNodes []*models.AppformixOpenstackNode, c *Cluster) error {
	a.clusterInfo.AppformixOpenstackNodes = nil
	for _, appformixOpenstackNode := range appformixOpenstackNodes {
		appformixOpenstackNodeInfo := models.InterfaceToAppformixOpenstackNode(appformixOpenstackNode)
		// Read appformixOpenstack role node to get the node refs information
		appformixOpenstackNodeData, err := c.getResource(
			defaultAppformixOpenstackNodeResPath, appformixOpenstackNodeInfo.UUID)
		if err != nil {
			return err
		}
		appformixOpenstackNodeInfo = models.InterfaceToAppformixOpenstackNode(
			appformixOpenstackNodeData)
		a.clusterInfo.AppformixOpenstackNodes = append(
			a.clusterInfo.AppformixOpenstackNodes, appformixOpenstackNodeInfo)
	}
	return nil
}

func (a *AppformixData) interfaceToAppformixComputeNode(
	appformixComputeNodes []*models.AppformixComputeNode, c *Cluster) error {
	a.clusterInfo.AppformixComputeNodes = nil
	for _, appformixComputeNode := range appformixComputeNodes {
		appformixComputeNodeInfo := models.InterfaceToAppformixComputeNode(appformixComputeNode)
		// Read appformixCompute role node to get the node refs information
		appformixComputeNodeData, err := c.getResource(
			defaultAppformixComputeNodeResPath, appformixComputeNodeInfo.UUID)
		if err != nil {
			return err
		}
		appformixComputeNodeInfo = models.InterfaceToAppformixComputeNode(
			appformixComputeNodeData)
		a.clusterInfo.AppformixComputeNodes = append(
			a.clusterInfo.AppformixComputeNodes, appformixComputeNodeInfo)
	}
	return nil
}

// nolint: gocyclo
func (a *AppformixData) updateNodeDetails(c *Cluster) error {
	m := make(map[string]bool)
	for _, node := range a.clusterInfo.AppformixControllerNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, a); err != nil {
				return err
			}
		}
	}
	for _, node := range a.clusterInfo.AppformixBareHostNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, a); err != nil {
				return err
			}
		}
	}
	for _, node := range a.clusterInfo.AppformixOpenstackNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, a); err != nil {
				return err
			}
		}
	}
	for _, node := range a.clusterInfo.AppformixComputeNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, a); err != nil {
				return err
			}
		}
	}
	return nil
}

// nolint: gocyclo
func (a *AppformixData) updateClusterDetails(clusterID string, c *Cluster) error {
	ctx := context.Background()
	resp, err := c.APIServer.GetAppformixCluster(ctx, &services.GetAppformixClusterRequest{ID: clusterID})
	if err != nil {
		return err
	}
	a.clusterInfo = resp.AppformixCluster

	// Expand appformix_controller back ref

	if err = a.interfaceToAppformixControllerNode(a.clusterInfo.AppformixControllerNodes, c); err != nil {
		return err
	}

	// Expand appformix_bare_host back ref

	if err = a.interfaceToAppformixBareHostNode(a.clusterInfo.AppformixBareHostNodes, c); err != nil {
		return err
	}

	// Expand appformix_openstack back ref

	if err = a.interfaceToAppformixOpenstackNode(a.clusterInfo.AppformixOpenstackNodes, c); err != nil {
		return err
	}

	// Expand appformix_compute back ref

	if err = a.interfaceToAppformixComputeNode(a.clusterInfo.AppformixComputeNodes, c); err != nil {
		return err
	}

	// get all nodes information
	if err = a.updateNodeDetails(c); err != nil {
		return err
	}
	return nil
}

func (k *KubernetesData) addKeypair(keypair *models.Keypair) {
	k.keypairsInfo = append(k.keypairsInfo, keypair)
}

func (k *KubernetesData) addCredential(cred *models.Credential) {
	k.credsInfo = append(k.credsInfo, cred)
}

func (k *KubernetesData) addNode(node *models.Node) {
	k.nodesInfo = append(k.nodesInfo, node)
}

func (k *KubernetesData) updateNodeDetails(c *Cluster) error {
	m := make(map[string]bool)
	for _, node := range k.clusterInfo.KubernetesMasterNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, k); err != nil {
				return err
			}
		}
	}
	for _, node := range k.clusterInfo.KubernetesNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, k); err != nil {
				return err
			}
		}
	}
	for _, node := range k.clusterInfo.KubernetesKubemanagerNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, k); err != nil {
				return err
			}
		}
	}
	return nil
}

func (k *KubernetesData) interfaceToKubernetesNode(kubernetesNodes interface{}, c *Cluster) error {
	n, ok := kubernetesNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, kubernetesNode := range n {
		kubernetesNodeInfo := models.InterfaceToKubernetesNode(kubernetesNode)
		// Read kubernetes role node to get the node refs information
		kubernetesNodeData, err := c.getResource(
			defaultKubernetesNodeResPath, kubernetesNodeInfo.UUID)
		if err != nil {
			return err
		}
		kubernetesNodeInfo = models.InterfaceToKubernetesNode(
			kubernetesNodeData)
		k.clusterInfo.KubernetesNodes = append(
			k.clusterInfo.KubernetesNodes, kubernetesNodeInfo)
	}
	return nil
}

func (k *KubernetesData) interfaceToKubernetesMasterNode(kubernetesMasterNodes interface{}, c *Cluster) error {
	n, ok := kubernetesMasterNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, kubernetesMasterNode := range n {
		kubernetesMasterNodeInfo := models.InterfaceToKubernetesMasterNode(kubernetesMasterNode)
		// Read kubernetesMaster role node to get the node refs information
		kubernetesMasterNodeData, err := c.getResource(
			defaultKubernetesMasterNodeResPath, kubernetesMasterNodeInfo.UUID)
		if err != nil {
			return err
		}
		kubernetesMasterNodeInfo = models.InterfaceToKubernetesMasterNode(
			kubernetesMasterNodeData)
		k.clusterInfo.KubernetesMasterNodes = append(
			k.clusterInfo.KubernetesMasterNodes, kubernetesMasterNodeInfo)
	}
	return nil
}

func (k *KubernetesData) interfaceToKubernetesKubemanagerNode(
	kubernetesKubemanagerNodes interface{}, c *Cluster,
) error {
	n, ok := kubernetesKubemanagerNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, kubernetesKubemanagerNode := range n {
		kubernetesKubemanagerNodeInfo := models.InterfaceToKubernetesKubemanagerNode(kubernetesKubemanagerNode)
		// Read kubernetesKubemanager role node to get the node refs information
		kubernetesKubemanagerNodeData, err := c.getResource(
			defaultKubernetesKubemanagerNodeResPath, kubernetesKubemanagerNodeInfo.UUID)
		if err != nil {
			return err
		}
		kubernetesKubemanagerNodeInfo = models.InterfaceToKubernetesKubemanagerNode(
			kubernetesKubemanagerNodeData)
		k.clusterInfo.KubernetesKubemanagerNodes = append(
			k.clusterInfo.KubernetesKubemanagerNodes, kubernetesKubemanagerNodeInfo)
	}
	return nil
}

func (k *KubernetesData) updateClusterDetails(clusterID string, c *Cluster) error {
	rData, err := c.getResource(defaultK8sResourcePath, clusterID)
	if err != nil {
		return err
	}
	k.clusterInfo = models.InterfaceToKubernetesCluster(rData)

	// Expand kubernetes node back ref
	if kubernetesNodes, ok := rData["kubernetes_nodes"]; ok {
		if err = k.interfaceToKubernetesNode(kubernetesNodes, c); err != nil {
			return err
		}
	}

	// Expand kubernetes_master back ref
	if kubernetesMasterNodes, ok := rData["kubernetes_master_nodes"]; ok {
		if err = k.interfaceToKubernetesMasterNode(kubernetesMasterNodes, c); err != nil {
			return err
		}
	}

	// Expand kubernetes_kubemanager back ref
	if kubernetesKubemanagerNodes, ok := rData["kubernetes_kubemanager_nodes"]; ok {
		if err = k.interfaceToKubernetesKubemanagerNode(kubernetesKubemanagerNodes, c); err != nil {
			return err
		}
	}

	// get all nodes information
	if err = k.updateNodeDetails(c); err != nil {
		return err
	}
	return nil
}

func (v *VCenterData) addKeypair(keypair *models.Keypair) {
	v.keypairsInfo = append(v.keypairsInfo, keypair)
}

func (v *VCenterData) addCredential(cred *models.Credential) {
	v.credsInfo = append(v.credsInfo, cred)
}

func (v *VCenterData) addNode(node *models.Node) {
	v.nodesInfo = append(v.nodesInfo, node)
}

func (v *VCenterData) updateNodeDetails(c *Cluster) error {
	m := make(map[string]bool)
	for _, node := range v.clusterInfo.VCenterPluginNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, v); err != nil {
				return err
			}
		}
	}
	for _, node := range v.clusterInfo.VCenterComputes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, v); err != nil {
				return err
			}
		}
	}
	for _, node := range v.clusterInfo.VCenterManagerNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, v); err != nil {
				return err
			}
		}
	}

	return nil
}

func (v *VCenterData) interfaceToVCenterPluginNode(
	vcenterPluginNodes interface{}, c *Cluster) error {
	for _, vcenterPluginNode := range vcenterPluginNodes.([]interface{}) {
		vcenterPluginNodeInfo := models.InterfaceToVCenterPluginNode(
			vcenterPluginNode.(map[string]interface{}))
		// Read vcenter_plugin role node to get the node refs information
		vcenterPluginNodeData, err := c.getResource(
			defaultVCenterPluginNodeResPath, vcenterPluginNodeInfo.UUID)
		if err != nil {
			return err
		}
		vcenterPluginNodeInfo = models.InterfaceToVCenterPluginNode(
			vcenterPluginNodeData)
		v.clusterInfo.VCenterPluginNodes = append(
			v.clusterInfo.VCenterPluginNodes, vcenterPluginNodeInfo)
	}
	return nil
}

func (v *VCenterData) interfaceToVCenterCompute(
	vcenterComputes interface{}, c *Cluster) error {
	for _, vcenterCompute := range vcenterComputes.([]interface{}) {
		vcenterComputeInfo := models.InterfaceToVCenterCompute(
			vcenterCompute.(map[string]interface{}))
		// Read vcenter_compute role node to get the node refs information
		vcenterComputeData, err := c.getResource(
			defaultVCenterComputeResPath, vcenterComputeInfo.UUID)
		if err != nil {
			return err
		}
		vcenterComputeInfo = models.InterfaceToVCenterCompute(
			vcenterComputeData)
		v.clusterInfo.VCenterComputes = append(
			v.clusterInfo.VCenterComputes, vcenterComputeInfo)
	}
	return nil
}

func (v *VCenterData) interfaceToVCenterManagerNode(
	vcenterManagerNodes interface{}, c *Cluster) error {
	for _, vcenterManagerNode := range vcenterManagerNodes.([]interface{}) {
		vcenterManagerNodeInfo := models.InterfaceToVCenterManagerNode(
			vcenterManagerNode.(map[string]interface{}))
		// Read vcenter_manager role node to get the node refs information
		vcenterManagerNodeData, err := c.getResource(
			defaultVCenterManagerNodeResPath, vcenterManagerNodeInfo.UUID)
		if err != nil {
			return err
		}
		vcenterManagerNodeInfo = models.InterfaceToVCenterManagerNode(
			vcenterManagerNodeData)
		v.clusterInfo.VCenterManagerNodes = append(
			v.clusterInfo.VCenterManagerNodes, vcenterManagerNodeInfo)
	}
	return nil
}

func (v *VCenterData) updateClusterDetails(clusterID string, c *Cluster) error {
	rData, err := c.getResource(defaultVCenterResourcePath, clusterID)
	if err != nil {
		return err
	}
	v.clusterInfo = models.InterfaceToVCenter(rData)

	// Expand vcenter_plugin back ref
	if vcenterPluginNodes, ok := rData["vCenter_plugin_nodes"]; ok {
		if err = v.interfaceToVCenterPluginNode(vcenterPluginNodes, c); err != nil {
			return err
		}
	}
	// Expand vcenter_compute back ref
	if vcenterComputes, ok := rData["vCenter_computes"]; ok {
		if err = v.interfaceToVCenterCompute(vcenterComputes, c); err != nil {
			return err
		}
	}

	// Expand vcenter_manager back ref
	if vcenterManagerNodes, ok := rData["vCenter_manager_nodes"]; ok {
		if err = v.interfaceToVCenterManagerNode(vcenterManagerNodes, c); err != nil {
			return err
		}
	}

	// get all nodes information
	if err = v.updateNodeDetails(c); err != nil {
		return err
	}
	return nil
}

func (o *OpenstackData) addKeypair(keypair *models.Keypair) {
	o.keypairsInfo = append(o.keypairsInfo, keypair)
}

func (o *OpenstackData) addCredential(cred *models.Credential) {
	o.credsInfo = append(o.credsInfo, cred)
}

func (o *OpenstackData) addNode(node *models.Node) {
	o.nodesInfo = append(o.nodesInfo, node)
}

// nolint: gocyclo
func (o *OpenstackData) updateNodeDetails(c *Cluster) error {
	m := make(map[string]bool)
	for _, node := range o.clusterInfo.OpenstackControlNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	for _, node := range o.clusterInfo.OpenstackNetworkNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	for _, node := range o.clusterInfo.OpenstackStorageNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	for _, node := range o.clusterInfo.OpenstackMonitoringNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	for _, node := range o.clusterInfo.OpenstackComputeNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackControlNode(openstackControlNodes interface{}, c *Cluster) error {
	n, ok := openstackControlNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackControlNode := range n {
		openstackControlNodeInfo := models.InterfaceToOpenstackControlNode(openstackControlNode)
		// Read openstackControl role node to get the node refs information
		openstackControlNodeData, err := c.getResource(
			defaultOpenstackControlNodeResPath, openstackControlNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackControlNodeInfo = models.InterfaceToOpenstackControlNode(
			openstackControlNodeData)
		o.clusterInfo.OpenstackControlNodes = append(
			o.clusterInfo.OpenstackControlNodes, openstackControlNodeInfo)
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackMonitoringNode(openstackMonitoringNodes interface{}, c *Cluster) error {
	n, ok := openstackMonitoringNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackMonitoringNode := range n {
		openstackMonitoringNodeInfo := models.InterfaceToOpenstackMonitoringNode(openstackMonitoringNode)
		// Read openstackMonitoring role node to get the node refs information
		openstackMonitoringNodeData, err := c.getResource(
			defaultOpenstackMonitoringNodeResPath, openstackMonitoringNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackMonitoringNodeInfo = models.InterfaceToOpenstackMonitoringNode(
			openstackMonitoringNodeData)
		o.clusterInfo.OpenstackMonitoringNodes = append(

			o.clusterInfo.OpenstackMonitoringNodes, openstackMonitoringNodeInfo)
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackNetworkNode(openstackNetworkNodes interface{}, c *Cluster) error {
	n, ok := openstackNetworkNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackNetworkNode := range n {
		openstackNetworkNodeInfo := models.InterfaceToOpenstackNetworkNode(openstackNetworkNode)
		// Read openstackNetwork role node to get the node refs information
		openstackNetworkNodeData, err := c.getResource(
			defaultOpenstackNetworkNodeResPath, openstackNetworkNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackNetworkNodeInfo = models.InterfaceToOpenstackNetworkNode(
			openstackNetworkNodeData)
		o.clusterInfo.OpenstackNetworkNodes = append(
			o.clusterInfo.OpenstackNetworkNodes, openstackNetworkNodeInfo)
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackStorageNode(openstackStorageNodes interface{}, c *Cluster) error {
	n, ok := openstackStorageNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackStorageNode := range n {
		openstackStorageNodeInfo := models.InterfaceToOpenstackStorageNode(openstackStorageNode)
		// Read openstackStorage role node to get the node refs information
		openstackStorageNodeData, err := c.getResource(
			defaultOpenstackStorageNodeResPath, openstackStorageNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackStorageNodeInfo = models.InterfaceToOpenstackStorageNode(
			openstackStorageNodeData)
		o.clusterInfo.OpenstackStorageNodes = append(
			o.clusterInfo.OpenstackStorageNodes, openstackStorageNodeInfo)
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackComputeNode(openstackComputeNodes interface{}, c *Cluster) error {
	n, ok := openstackComputeNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackComputeNode := range n {
		openstackComputeNodeInfo := models.InterfaceToOpenstackComputeNode(openstackComputeNode)
		// Read openstackCompute role node to get the node refs information
		openstackComputeNodeData, err := c.getResource(
			defaultOpenstackComputeNodeResPath, openstackComputeNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackComputeNodeInfo = models.InterfaceToOpenstackComputeNode(
			openstackComputeNodeData)
		o.clusterInfo.OpenstackComputeNodes = append(
			o.clusterInfo.OpenstackComputeNodes, openstackComputeNodeInfo)
	}
	return nil
}

// nolint: gocyclo
func (o *OpenstackData) updateClusterDetails(clusterID string, c *Cluster) error {
	rData, err := c.getResource(defaultOpenstackResourcePath, clusterID)
	if err != nil {
		return err
	}
	o.clusterInfo = models.InterfaceToOpenstackCluster(rData)

	// Expand openstack_compute back ref
	if openstackComputeNodes, ok := rData["openstack_compute_nodes"]; ok {
		if err = o.interfaceToOpenstackComputeNode(openstackComputeNodes, c); err != nil {
			return err
		}
	}
	// Expand openstack_storage node back ref
	if openstackStorageNodes, ok := rData["openstack_storage_nodes"]; ok {
		if err = o.interfaceToOpenstackStorageNode(openstackStorageNodes, c); err != nil {
			return err
		}
	}
	// Expand openstack_network node back ref
	if openstackNetworkNodes, ok := rData["openstack_network_nodes"]; ok {
		if err = o.interfaceToOpenstackNetworkNode(openstackNetworkNodes, c); err != nil {
			return err
		}
	}
	// Expand openstack_monitoring node back ref
	if openstackMonitoringNodes, ok := rData["openstack_monitoring_nodes"]; ok {
		if err = o.interfaceToOpenstackMonitoringNode(openstackMonitoringNodes, c); err != nil {
			return err
		}
	}
	// Expand openstack_control node back ref
	if openstackControlNodes, ok := rData["openstack_control_nodes"]; ok {
		if err = o.interfaceToOpenstackControlNode(openstackControlNodes, c); err != nil {
			return err
		}
	}
	// get all nodes information
	if err = o.updateNodeDetails(c); err != nil {
		return err
	}
	return nil
}

func (o *OpenstackData) getControlNodeIPs() (nodeIPs []string) {
	for _, controlNode := range o.clusterInfo.OpenstackControlNodes {
		for _, nodeRef := range controlNode.NodeRefs {
			for _, node := range o.nodesInfo {
				if nodeRef.UUID == node.UUID {
					nodeIPs = append(nodeIPs, node.IPAddress)
				}
			}
		}
	}
	return nodeIPs
}

func (o *OpenstackData) getOpenstackControlPorts() (nodePorts map[string]interface{}) {
	nodePorts = make(map[string]interface{})
	for _, controlNode := range o.clusterInfo.OpenstackControlNodes {
		for _, nodeRef := range controlNode.NodeRefs {
			nodeIPAddress := getNodeIPAddress(o.Reader, nodeRef.UUID)
			portMap := make(map[string]int64)
			if _, ok := nodePorts[nodeIPAddress]; !ok {
				nodePorts[nodeIPAddress] = portMap
			}
			format.InterfaceToInt64Map(nodePorts[nodeIPAddress])[identity] = controlNode.KeystonePublicPort
			format.InterfaceToInt64Map(nodePorts[nodeIPAddress])[nova] = controlNode.NovaPublicPort
			format.InterfaceToInt64Map(nodePorts[nodeIPAddress])[glance] = controlNode.GlancePublicPort
			format.InterfaceToInt64Map(nodePorts[nodeIPAddress])[ironic] = controlNode.IronicPublicPort
		}
	}
	return nodePorts
}

func (o *OpenstackData) getStorageNodeIPs() (nodeIPs []string) {
	for _, storageNode := range o.clusterInfo.OpenstackStorageNodes {
		for _, nodeRef := range storageNode.NodeRefs {
			for _, node := range o.nodesInfo {
				if nodeRef.UUID == node.UUID {
					nodeIPs = append(nodeIPs, node.IPAddress)
				}
			}
		}
	}
	return nodeIPs
}

func (o *OpenstackData) getOpenstackStoragePorts() (nodePorts map[string]interface{}) {
	nodePorts = make(map[string]interface{})
	for _, storageNode := range o.clusterInfo.OpenstackStorageNodes {
		for _, nodeRef := range storageNode.NodeRefs {
			nodeIPAddress := getNodeIPAddress(o.Reader, nodeRef.UUID)
			portMap := make(map[string]int64)
			if _, ok := nodePorts[nodeIPAddress]; !ok {
				nodePorts[nodeIPAddress] = portMap
			}
			format.InterfaceToInt64Map(nodePorts[nodeIPAddress])[swift] = storageNode.SwiftPublicPort
		}
	}
	return nodePorts
}

func (d *Data) addKeypair(keypair *models.Keypair) {
	d.keypairsInfo = append(d.keypairsInfo, keypair)
}

func (d *Data) addCredential(cred *models.Credential) {
	d.credsInfo = append(d.credsInfo, cred)
}

func (d *Data) addNode(node *models.Node) {
	d.nodesInfo = append(d.nodesInfo, node)
}

func (d *Data) interfaceToContrailVrouterNode(contrailVrouterNodes interface{}, c *Cluster) error {
	n, ok := contrailVrouterNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailVrouterNode := range n {
		contrailVrouterNodeInfo := models.InterfaceToContrailVrouterNode(contrailVrouterNode)
		// Read contrailVrouter role node to get the node refs information
		contrailVrouterNodeData, err := c.getResource(
			defaultContrailVrouterNodeResPath, contrailVrouterNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailVrouterNodeInfo = models.InterfaceToContrailVrouterNode(
			contrailVrouterNodeData)
		d.clusterInfo.ContrailVrouterNodes = append(
			d.clusterInfo.ContrailVrouterNodes, contrailVrouterNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailMCGWNode(contrailMCGWNodes interface{}, c *Cluster) error {
	n, ok := contrailMCGWNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailMCGWNode := range n {
		contrailMCGWNodeInfo := models.InterfaceToContrailMulticloudGWNode(contrailMCGWNode)
		// Read ContrailMulticloudGW role node to get the node refs information
		contrailMCGWNodeData, err := c.getResource(
			defaultContrailMCGWNodeResPath, contrailMCGWNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailMCGWNodeInfo = models.InterfaceToContrailMulticloudGWNode(
			contrailMCGWNodeData)
		d.clusterInfo.ContrailMulticloudGWNodes = append(
			d.clusterInfo.ContrailMulticloudGWNodes, contrailMCGWNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailZTPTFTPNode(contrailZTPTFTPNodes interface{}, c *Cluster) error {
	n, ok := contrailZTPTFTPNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailZTPTFTPNode := range n {
		contrailZTPTFTPNodeInfo := models.InterfaceToContrailZTPTFTPNode(contrailZTPTFTPNode)
		// Read contrailZTPTFTP role node to get the node refs information
		contrailZTPTFTPNodeData, err := c.getResource(
			defaultContrailZTPTFTPNodeResPath, contrailZTPTFTPNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailZTPTFTPNodeInfo = models.InterfaceToContrailZTPTFTPNode(
			contrailZTPTFTPNodeData)
		d.clusterInfo.ContrailZTPTFTPNodes = append(
			d.clusterInfo.ContrailZTPTFTPNodes, contrailZTPTFTPNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailZTPDHCPNode(contrailZTPDHCPNodes interface{}, c *Cluster) error {
	n, ok := contrailZTPDHCPNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailZTPDHCPNode := range n {
		contrailZTPDHCPNodeInfo := models.InterfaceToContrailZTPDHCPNode(contrailZTPDHCPNode)
		// Read contrailZTPDHCP role node to get the node refs information
		contrailZTPDHCPNodeData, err := c.getResource(
			defaultContrailZTPDHCPNodeResPath, contrailZTPDHCPNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailZTPDHCPNodeInfo = models.InterfaceToContrailZTPDHCPNode(
			contrailZTPDHCPNodeData)
		d.clusterInfo.ContrailZTPDHCPNodes = append(
			d.clusterInfo.ContrailZTPDHCPNodes, contrailZTPDHCPNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailServiceNode(contrailServiceNodes interface{}, c *Cluster) error {
	n, ok := contrailServiceNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailServiceNode := range n {
		contrailServiceNodeInfo := models.InterfaceToContrailServiceNode(contrailServiceNode)
		// Read contrailService role node to get the node refs information
		contrailServiceNodeData, err := c.getResource(
			defaultContrailServiceNodeResPath, contrailServiceNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailServiceNodeInfo = models.InterfaceToContrailServiceNode(
			contrailServiceNodeData)
		d.clusterInfo.ContrailServiceNodes = append(
			d.clusterInfo.ContrailServiceNodes, contrailServiceNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailAnalyticsDatabaseNode(contrailAnalyticsDatabaseNodes interface{}, c *Cluster) error {
	n, ok := contrailAnalyticsDatabaseNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailAnalyticsDatabaseNode := range n {
		contrailAnalyticsDatabaseNodeInfo := models.InterfaceToContrailAnalyticsDatabaseNode(contrailAnalyticsDatabaseNode)
		// Read contrailAnalyticsDatabase role node to get the node refs information
		contrailAnalyticsDatabaseNodeData, err := c.getResource(
			defaultContrailAnalyticsDatabaseNodeResPath, contrailAnalyticsDatabaseNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailAnalyticsDatabaseNodeInfo = models.InterfaceToContrailAnalyticsDatabaseNode(
			contrailAnalyticsDatabaseNodeData)
		d.clusterInfo.ContrailAnalyticsDatabaseNodes = append(
			d.clusterInfo.ContrailAnalyticsDatabaseNodes, contrailAnalyticsDatabaseNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailAnalyticsNode(contrailAnalyticsNodes interface{}, c *Cluster) error {
	n, ok := contrailAnalyticsNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailAnalyticsNode := range n {
		contrailAnalyticsNodeInfo := models.InterfaceToContrailAnalyticsNode(contrailAnalyticsNode)
		// Read contrailAnalytics role node to get the node refs information
		contrailAnalyticsNodeData, err := c.getResource(
			defaultContrailAnalyticsNodeResPath, contrailAnalyticsNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailAnalyticsNodeInfo = models.InterfaceToContrailAnalyticsNode(
			contrailAnalyticsNodeData)
		d.clusterInfo.ContrailAnalyticsNodes = append(
			d.clusterInfo.ContrailAnalyticsNodes, contrailAnalyticsNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailAnalyticsAlarmNode(contrailAnalyticsAlarmNodes interface{}, c *Cluster) error {
	n, ok := contrailAnalyticsAlarmNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailAnalyticsAlarmNode := range n {
		contrailAnalyticsAlarmNodeInfo := models.InterfaceToContrailAnalyticsAlarmNode(contrailAnalyticsAlarmNode)
		// Read contrailAnalyticsAlarm role node to get the node refs information
		contrailAnalyticsAlarmNodeData, err := c.getResource(
			defaultContrailAnalyticsAlarmNodeResPath, contrailAnalyticsAlarmNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailAnalyticsAlarmNodeInfo = models.InterfaceToContrailAnalyticsAlarmNode(
			contrailAnalyticsAlarmNodeData)
		d.clusterInfo.ContrailAnalyticsAlarmNodes = append(
			d.clusterInfo.ContrailAnalyticsAlarmNodes, contrailAnalyticsAlarmNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailAnalyticsSNMPNode(contrailAnalyticsSNMPNodes interface{}, c *Cluster) error {
	n, ok := contrailAnalyticsSNMPNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailAnalyticsSNMPNode := range n {
		contrailAnalyticsSNMPNodeInfo := models.InterfaceToContrailAnalyticsSNMPNode(contrailAnalyticsSNMPNode)
		// Read contrailAnalytics role node to get the node refs information
		contrailAnalyticsSNMPNodeData, err := c.getResource(
			defaultContrailAnalyticsSNMPNodeResPath, contrailAnalyticsSNMPNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailAnalyticsSNMPNodeInfo = models.InterfaceToContrailAnalyticsSNMPNode(
			contrailAnalyticsSNMPNodeData)
		d.clusterInfo.ContrailAnalyticsSNMPNodes = append(
			d.clusterInfo.ContrailAnalyticsSNMPNodes, contrailAnalyticsSNMPNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailWebuiNode(contrailWebuiNodes interface{}, c *Cluster) error {
	n, ok := contrailWebuiNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailWebuiNode := range n {
		contrailWebuiNodeInfo := models.InterfaceToContrailWebuiNode(contrailWebuiNode)
		// Read contrailWebui role node to get the node refs information
		contrailWebuiNodeData, err := c.getResource(
			defaultContrailWebuiNodeResPath, contrailWebuiNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailWebuiNodeInfo = models.InterfaceToContrailWebuiNode(
			contrailWebuiNodeData)
		d.clusterInfo.ContrailWebuiNodes = append(
			d.clusterInfo.ContrailWebuiNodes, contrailWebuiNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailControlNode(contrailControlNodes interface{}, c *Cluster) error {
	n, ok := contrailControlNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailControlNode := range n {
		contrailControlNodeInfo := models.InterfaceToContrailControlNode(contrailControlNode)
		// Read contrailControl role node to get the node refs information
		contrailControlNodeData, err := c.getResource(
			defaultContrailControlNodeResPath, contrailControlNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailControlNodeInfo = models.InterfaceToContrailControlNode(

			contrailControlNodeData)
		d.clusterInfo.ContrailControlNodes = append(
			d.clusterInfo.ContrailControlNodes, contrailControlNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailConfigDatabaseNode(contrailConfigDatabaseNodes interface{}, c *Cluster) error {
	n, ok := contrailConfigDatabaseNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailConfigDatabaseNode := range n {
		contrailConfigDatabaseNodeInfo := models.InterfaceToContrailConfigDatabaseNode(contrailConfigDatabaseNode)
		// Read contrailConfigDatabase role node to get the node refs information
		contrailConfigDatabaseNodeData, err := c.getResource(
			defaultContrailConfigDatabaseNodeResPath, contrailConfigDatabaseNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailConfigDatabaseNodeInfo = models.InterfaceToContrailConfigDatabaseNode(
			contrailConfigDatabaseNodeData)
		d.clusterInfo.ContrailConfigDatabaseNodes = append(
			d.clusterInfo.ContrailConfigDatabaseNodes, contrailConfigDatabaseNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailConfigNode(contrailConfigNodes interface{}, c *Cluster) error {
	n, ok := contrailConfigNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailConfigNode := range n {
		contrailConfigNodeInfo := models.InterfaceToContrailConfigNode(contrailConfigNode)
		// Read contrailConfig role node to get the node refs information
		contrailConfigNodeData, err := c.getResource(
			defaultContrailConfigNodeResPath, contrailConfigNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailConfigNodeInfo = models.InterfaceToContrailConfigNode(
			contrailConfigNodeData)
		d.clusterInfo.ContrailConfigNodes = append(
			d.clusterInfo.ContrailConfigNodes, contrailConfigNodeInfo)
	}
	return nil
}

// nolint: gocyclo
func (d *Data) updateNodeDetails(c *Cluster) error {
	m := make(map[string]bool)
	for _, node := range d.clusterInfo.ContrailConfigNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailConfigDatabaseNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailControlNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailWebuiNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailAnalyticsNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailAnalyticsDatabaseNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailAnalyticsAlarmNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailAnalyticsSNMPNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailVrouterNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailMulticloudGWNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.clusterInfo.ContrailServiceNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := c.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Data) updateCloudDetails(c *Cluster) error {
	for _, cloudRef := range d.clusterInfo.CloudRefs {
		cloudObject, err := cloud.GetCloud(context.Background(), c.APIServer, cloudRef.UUID)
		if err != nil {
			return err
		}
		d.cloudInfo = append(d.cloudInfo, cloudObject)
	}
	return nil
}

// nolint: gocyclo
func (d *Data) updateClusterDetails(clusterID string, c *Cluster) error {
	rData, err := c.getResource(defaultResourcePath, clusterID)
	if err != nil {
		return err
	}
	d.clusterInfo = models.InterfaceToContrailCluster(rData)

	// Expand config node back ref
	if configNodes, ok := rData["contrail_config_nodes"]; ok {
		if err = d.interfaceToContrailConfigNode(configNodes, c); err != nil {
			return err
		}
	}
	// Expand config database node back ref
	if configDBNodes, ok := rData["contrail_config_database_nodes"]; ok {
		if err = d.interfaceToContrailConfigDatabaseNode(configDBNodes, c); err != nil {
			return err
		}
	}
	// Expand control node back ref
	if controlNodes, ok := rData["contrail_control_nodes"]; ok {
		if err = d.interfaceToContrailControlNode(controlNodes, c); err != nil {
			return err
		}
	}
	// Expand webui node back ref
	if webuiNodes, ok := rData["contrail_webui_nodes"]; ok {
		if err = d.interfaceToContrailWebuiNode(webuiNodes, c); err != nil {
			return err
		}
	}
	// Expand analytics node back ref
	if analyticsNodes, ok := rData["contrail_analytics_nodes"]; ok {
		if err = d.interfaceToContrailAnalyticsNode(analyticsNodes, c); err != nil {
			return err
		}
	}
	// Expand analytics database node back ref
	if analyticsDBNodes, ok := rData["contrail_analytics_database_nodes"]; ok {
		if err = d.interfaceToContrailAnalyticsDatabaseNode(analyticsDBNodes, c); err != nil {
			return err
		}
	}
	// Expand analytics alarm node back ref
	if analyticsAlarmNodes, ok := rData["contrail_analytics_alarm_nodes"]; ok {
		if err = d.interfaceToContrailAnalyticsAlarmNode(analyticsAlarmNodes, c); err != nil {
			return err
		}
	}
	// Expand analytics snmp node back ref
	if analyticsSNMPNodes, ok := rData["contrail_analytics_snmp_nodes"]; ok {
		if err = d.interfaceToContrailAnalyticsSNMPNode(analyticsSNMPNodes, c); err != nil {
			return err
		}
	}
	// Expand vouter node back ref
	if vrouterNodes, ok := rData["contrail_vrouter_nodes"]; ok {
		if err = d.interfaceToContrailVrouterNode(vrouterNodes, c); err != nil {
			return err
		}
	}
	if mcGWNodes, ok := rData["contrail_multicloud_gw_nodes"]; ok {
		if err = d.interfaceToContrailMCGWNode(mcGWNodes, c); err != nil {
			return err
		}
	}
	// Expand csn node back ref
	if csnNodes, ok := rData["contrail_service_nodes"]; ok {
		if err = d.interfaceToContrailServiceNode(csnNodes, c); err != nil {
			return err
		}
	}
	// Expand tftp node back ref
	if tftpNodes, ok := rData["contrail_ztp_tftp_nodes"]; ok {
		if err = d.interfaceToContrailZTPTFTPNode(tftpNodes, c); err != nil {
			return err
		}
	}
	// Expand dhcp node back ref
	if dhcpNodes, ok := rData["contrail_ztp_dhcp_nodes"]; ok {
		if err = d.interfaceToContrailZTPDHCPNode(dhcpNodes, c); err != nil {
			return err
		}
	}
	// get all nodes information
	if err := d.updateNodeDetails(c); err != nil {
		return err
	}

	// get all cloud information
	if err := d.updateCloudDetails(c); err != nil {
		return err
	}
	return nil
}

func (d *Data) getK8sClusterData() *KubernetesData {
	// One kubernetes cluster is the supported topology
	if len(d.kubernetesClusterData) > 0 {
		return d.kubernetesClusterData[0]
	}
	return nil
}

func (d *Data) getK8sClusterInfo() *models.KubernetesCluster {
	if d.getK8sClusterData() != nil {
		return d.getK8sClusterData().clusterInfo
	}
	return nil
}

func (d *Data) getVCenterClusterData() *VCenterData {

	if len(d.vcenterData) > 0 {
		return d.vcenterData[0]
	}
	return nil
}

func (d *Data) getVCenterClusterInfo() *models.VCenter {
	if d.getVCenterClusterData() != nil {
		return d.getVCenterClusterData().clusterInfo
	}
	return nil
}

func (d *Data) getOpenstackClusterData() *OpenstackData {
	// One openstack cluster is the supported topology
	if len(d.openstackClusterData) > 0 {
		return d.openstackClusterData[0]
	}
	return nil
}

func (d *Data) getOpenstackClusterInfo() *models.OpenstackCluster {
	if d.getOpenstackClusterData() != nil {
		return d.getOpenstackClusterData().clusterInfo
	}
	return nil
}

func (d *Data) getAllKeypairsInfo() []*models.Keypair {
	keypairs := d.keypairsInfo
	if d.getOpenstackClusterData() != nil {
		keypairs = append(keypairs, d.getOpenstackClusterData().keypairsInfo...)
	}
	if d.getK8sClusterData() != nil {
		keypairs = append(keypairs, d.getK8sClusterData().keypairsInfo...)
	}

	var uniqueKeypairs []*models.Keypair
	m := make(map[string]bool)

	for _, keypair := range keypairs {
		if _, ok := m[keypair.UUID]; !ok {
			m[keypair.UUID] = true
			uniqueKeypairs = append(uniqueKeypairs, keypair)
		}
	}

	return uniqueKeypairs
}

func (d *Data) getAllCredsInfo() []*models.Credential {
	creds := d.credsInfo
	if d.getOpenstackClusterData() != nil {
		creds = append(creds, d.getOpenstackClusterData().credsInfo...)
	}
	if d.getK8sClusterData() != nil {
		creds = append(creds, d.getK8sClusterData().credsInfo...)
	}

	var uniqueCreds []*models.Credential
	m := make(map[string]bool)

	for _, cred := range creds {
		if _, ok := m[cred.UUID]; !ok {
			m[cred.UUID] = true
			uniqueCreds = append(uniqueCreds, cred)
		}
	}

	return uniqueCreds
}

func (d *Data) getAllNodesInfo() []*models.Node {
	nodes := d.nodesInfo
	if d.getOpenstackClusterData() != nil {
		nodes = append(nodes, d.getOpenstackClusterData().nodesInfo...)
	}
	if d.getK8sClusterData() != nil {
		nodes = append(nodes, d.getK8sClusterData().nodesInfo...)
	}
	if d.getAppformixClusterData() != nil {
		nodes = append(nodes, d.getAppformixClusterData().nodesInfo...)
	}

	if d.getXflowData() != nil {
		nodes = append(nodes, d.getXflowData().getNodes()...)
	}

	var uniqueNodes []*models.Node
	m := make(map[string]bool)

	for _, node := range nodes {
		if _, ok := m[node.UUID]; !ok {
			m[node.UUID] = true
			uniqueNodes = append(uniqueNodes, node)
		}
	}

	return uniqueNodes
}

func (d *Data) getConfigNodeIPs() (nodeIPs []string) {
	for _, configNode := range d.clusterInfo.ContrailConfigNodes {
		for _, nodeRef := range configNode.NodeRefs {
			for _, node := range d.nodesInfo {
				if nodeRef.UUID == node.UUID {
					nodeIPs = append(nodeIPs, node.IPAddress)
				}
			}
		}
	}
	return nodeIPs
}

func (d *Data) getConfigNodePorts() (nodePorts map[string]interface{}) {
	nodePorts = make(map[string]interface{})
	for _, configNode := range d.clusterInfo.ContrailConfigNodes {
		for _, nodeRef := range configNode.NodeRefs {
			nodeIPAddress := getNodeIPAddress(d.Reader, nodeRef.UUID)
			portMap := make(map[string]int64)
			if _, ok := nodePorts[nodeIPAddress]; !ok {
				nodePorts[nodeIPAddress] = portMap
			}
			format.InterfaceToInt64Map(nodePorts[nodeIPAddress])[config] = configNode.APIPublicPort
		}
	}
	return nodePorts
}

func (d *Data) getAnalyticsNodeIPs() (nodeIPs []string) {
	for _, analyticsNode := range d.clusterInfo.ContrailAnalyticsNodes {
		for _, nodeRef := range analyticsNode.NodeRefs {
			for _, node := range d.nodesInfo {
				if nodeRef.UUID == node.UUID {
					nodeIPs = append(nodeIPs, node.IPAddress)
				}
			}
		}
	}
	return nodeIPs
}

func (d *Data) getAnalyticsNodePorts() (nodePorts map[string]interface{}) {
	nodePorts = make(map[string]interface{})
	for _, analyticsNode := range d.clusterInfo.ContrailAnalyticsNodes {
		for _, nodeRef := range analyticsNode.NodeRefs {
			nodeIPAddress := getNodeIPAddress(d.Reader, nodeRef.UUID)
			portMap := make(map[string]int64)
			if _, ok := nodePorts[nodeIPAddress]; !ok {
				nodePorts[nodeIPAddress] = portMap
			}
			format.InterfaceToInt64Map(nodePorts[nodeIPAddress])[analytics] = analyticsNode.APIPublicPort
		}
	}
	return nodePorts
}

func (d *Data) getWebuiNodeIPs() (nodeIPs []string) {
	for _, webuiNode := range d.clusterInfo.ContrailWebuiNodes {
		for _, nodeRef := range webuiNode.NodeRefs {
			for _, node := range d.nodesInfo {
				if nodeRef.UUID == node.UUID {
					nodeIPs = append(nodeIPs, node.IPAddress)
				}
			}
		}
	}
	return nodeIPs
}

func (d *Data) getWebuiNodePorts() (nodePorts map[string]interface{}) {
	nodePorts = make(map[string]interface{})
	for _, webuiNode := range d.clusterInfo.ContrailWebuiNodes {
		for _, nodeRef := range webuiNode.NodeRefs {
			nodeIPAddress := getNodeIPAddress(d.Reader, nodeRef.UUID)
			portMap := make(map[string]int64)
			if _, ok := nodePorts[nodeIPAddress]; !ok {
				nodePorts[nodeIPAddress] = portMap
			}
			format.InterfaceToInt64Map(nodePorts[nodeIPAddress])[webui] = webuiNode.PublicPort
		}
	}
	return nodePorts
}

func (d *Data) getAppformixClusterData() *AppformixData {
	// One appformix cluster is the supported topology
	if len(d.appformixClusterData) > 0 {
		return d.appformixClusterData[0]
	}
	return nil
}

func (d *Data) getAppformixClusterInfo() *models.AppformixCluster {
	if d.getAppformixClusterData() != nil {
		return d.getAppformixClusterData().clusterInfo
	}
	return nil
}

func (d *Data) getXflowData() *XflowData {
	if len(d.xflowData) > 0 {
		return d.xflowData[0]
	}

	return nil
}

func (x *XflowData) getNodes() []*models.Node {
	res := make([]*models.Node, 0, len(x.NodesInfo))
	for _, n := range x.NodesInfo {
		res = append(res, n)
	}

	return res
}

func (d *Data) getAppformixControllerNodeIPs() (nodeIPs []string) {
	appformixClusterInfo := d.getAppformixClusterInfo()
	if appformixClusterInfo == nil {
		return nodeIPs
	}
	for _, appformixControllerNode := range appformixClusterInfo.AppformixControllerNodes {
		for _, nodeRef := range appformixControllerNode.NodeRefs {
			for _, node := range d.getAppformixClusterData().nodesInfo {
				if nodeRef.UUID == node.UUID {
					nodeIPs = append(nodeIPs, node.IPAddress)
				}
			}
		}
	}
	return nodeIPs
}

func getNodeIPAddress(readService services.ReadService, nodeID string) (IPAddress string) {
	resp, err := readService.GetNode(context.Background(), &services.GetNodeRequest{
		ID:     nodeID,
		Fields: []string{"ip_address"},
	})
	if err != nil {
		return ""
	}
	return resp.GetNode().IPAddress
}

func (d *Data) getAppformixControllerNodePorts() (nodePorts map[string]interface{}) {
	nodePorts = make(map[string]interface{})
	appformixClusterInfo := d.getAppformixClusterInfo()
	if appformixClusterInfo == nil {
		return nodePorts
	}
	for _, appformixControllerNode := range appformixClusterInfo.AppformixControllerNodes {
		for _, nodeRef := range appformixControllerNode.NodeRefs {
			nodeIPAddress := getNodeIPAddress(d.Reader, nodeRef.UUID)
			portMap := make(map[string]int64)
			if _, ok := nodePorts[nodeIPAddress]; !ok {
				nodePorts[nodeIPAddress] = portMap
			}
			format.InterfaceToInt64Map(nodePorts[nodeIPAddress])[appformix] = appformixControllerNode.PublicPort
		}
	}
	return nodePorts
}

func (d *Data) getCloudRefs() ([]*models.Cloud, error) {
	return nil, nil
}

func (x *XflowData) updateClusterDetails(ctx context.Context, uuid string, c *Cluster) error {
	resp, err := c.APIServer.GetAppformixFlows(ctx, &services.GetAppformixFlowsRequest{ID: uuid})

	if err != nil {
		return err
	}

	x.ClusterInfo = resp.AppformixFlows

	for _, appformixFlowsNode := range x.ClusterInfo.AppformixFlowsNodes {
		err = x.updateAppformixFlowsNode(ctx, appformixFlowsNode.UUID, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (x *XflowData) updateAppformixFlowsNode(ctx context.Context, uuid string, c *Cluster) error {
	resp, err := c.APIServer.GetAppformixFlowsNode(ctx, &services.GetAppformixFlowsNodeRequest{ID: uuid})

	if err != nil {
		return err
	}

	for _, nodeRef := range resp.AppformixFlowsNode.NodeRefs {
		err = x.updateNodeInfo(ctx, nodeRef.UUID, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (x *XflowData) updateNodeInfo(ctx context.Context, nodeUUID string, c *Cluster) error {
	resp, err := c.APIServer.GetNode(ctx, &services.GetNodeRequest{ID: nodeUUID})

	if err != nil {
		return err
	}

	x.NodesInfo[nodeUUID] = resp.Node

	return nil
}
