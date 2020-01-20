package base

import (
	"context"
	"strings"

	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// DataStore interface to store cluster data
type DataStore interface {
	updateClusterDetails(string, *ResourceManager) error
	updateNodeDetails(*ResourceManager) error
	addNode(*models.Node)
	addCredential(*models.Credential)
	addKeypair(*models.Keypair)
}

// OpenstackData is the representation of openstack cluster details.
type OpenstackData struct {
	ClusterInfo  *models.OpenstackCluster
	nodesInfo    []*models.Node
	keypairsInfo []*models.Keypair
	credsInfo    []*models.Credential
	Reader       services.ReadService
}

// KubernetesData is the representation of kubernetes cluster details.
type KubernetesData struct {
	ClusterInfo  *models.KubernetesCluster
	nodesInfo    []*models.Node
	keypairsInfo []*models.Keypair
	credsInfo    []*models.Credential
	Reader       services.ReadService
}

// VCenterData is the representation of VCenter details.
type VCenterData struct {
	ClusterInfo  *models.VCenter
	nodesInfo    []*models.Node
	keypairsInfo []*models.Keypair
	credsInfo    []*models.Credential
	Reader       services.ReadService
}

// AppformixData is the representation of appformix cluster details.
type AppformixData struct {
	ClusterInfo  *models.AppformixCluster
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
	ClusterInfo           *models.ContrailCluster
	NodesInfo             []*models.Node
	keypairsInfo          []*models.Keypair
	credsInfo             []*models.Credential
	CloudInfo             []*models.Cloud
	openstackClusterData  []*OpenstackData
	vcenterData           []*VCenterData
	kubernetesClusterData []*KubernetesData
	appformixClusterData  []*AppformixData
	xflowData             []*XflowData
	DefaultSSHUser        string
	DefaultSSHPassword    string
	DefaultSSHKey         string
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
	appformixControllerNodes []*models.AppformixControllerNode, r *ResourceManager) error {
	a.ClusterInfo.AppformixControllerNodes = nil
	for _, appformixControllerNode := range appformixControllerNodes {
		appformixControllerNodeInfo := models.InterfaceToAppformixControllerNode(appformixControllerNode)
		// Read appformixController role node to get the node refs information
		appformixControllerNodeData, err := r.getResource(
			defaultAppformixControllerNodeResPath, appformixControllerNodeInfo.UUID)
		if err != nil {
			return err
		}
		appformixControllerNodeInfo = models.InterfaceToAppformixControllerNode(
			appformixControllerNodeData)
		a.ClusterInfo.AppformixControllerNodes = append(
			a.ClusterInfo.AppformixControllerNodes, appformixControllerNodeInfo)
	}
	return nil
}

func (a *AppformixData) interfaceToAppformixBareHostNode(
	appformixBareHostNodes []*models.AppformixBareHostNode, r *ResourceManager) error {
	a.ClusterInfo.AppformixBareHostNodes = nil
	for _, appformixBareHostNode := range appformixBareHostNodes {
		appformixBareHostNodeInfo := models.InterfaceToAppformixBareHostNode(appformixBareHostNode)
		// Read appformixBareHost role node to get the node refs information
		appformixBareHostNodeData, err := r.getResource(
			defaultAppformixBareHostNodeResPath, appformixBareHostNodeInfo.UUID)
		if err != nil {
			return err
		}
		appformixBareHostNodeInfo = models.InterfaceToAppformixBareHostNode(
			appformixBareHostNodeData)
		a.ClusterInfo.AppformixBareHostNodes = append(
			a.ClusterInfo.AppformixBareHostNodes, appformixBareHostNodeInfo)
	}
	return nil
}

func (a *AppformixData) interfaceToAppformixOpenstackNode(
	appformixOpenstackNodes []*models.AppformixOpenstackNode, r *ResourceManager) error {
	a.ClusterInfo.AppformixOpenstackNodes = nil
	for _, appformixOpenstackNode := range appformixOpenstackNodes {
		appformixOpenstackNodeInfo := models.InterfaceToAppformixOpenstackNode(appformixOpenstackNode)
		// Read appformixOpenstack role node to get the node refs information
		appformixOpenstackNodeData, err := r.getResource(
			defaultAppformixOpenstackNodeResPath, appformixOpenstackNodeInfo.UUID)
		if err != nil {
			return err
		}
		appformixOpenstackNodeInfo = models.InterfaceToAppformixOpenstackNode(
			appformixOpenstackNodeData)
		a.ClusterInfo.AppformixOpenstackNodes = append(
			a.ClusterInfo.AppformixOpenstackNodes, appformixOpenstackNodeInfo)
	}
	return nil
}

func (a *AppformixData) interfaceToAppformixComputeNode(
	appformixComputeNodes []*models.AppformixComputeNode, r *ResourceManager) error {
	a.ClusterInfo.AppformixComputeNodes = nil
	for _, appformixComputeNode := range appformixComputeNodes {
		appformixComputeNodeInfo := models.InterfaceToAppformixComputeNode(appformixComputeNode)
		// Read appformixCompute role node to get the node refs information
		appformixComputeNodeData, err := r.getResource(
			defaultAppformixComputeNodeResPath, appformixComputeNodeInfo.UUID)
		if err != nil {
			return err
		}
		appformixComputeNodeInfo = models.InterfaceToAppformixComputeNode(
			appformixComputeNodeData)
		a.ClusterInfo.AppformixComputeNodes = append(
			a.ClusterInfo.AppformixComputeNodes, appformixComputeNodeInfo)
	}
	return nil
}

func (a *AppformixData) interfaceToAppformixNetworkAgentsNode(
	appformixNetworkAgentsNodes []*models.AppformixNetworkAgentsNode, r *ResourceManager) error {
	a.ClusterInfo.AppformixNetworkAgentsNodes = nil
	for _, appformixNetworkAgentsNode := range appformixNetworkAgentsNodes {
		appformixNetworkAgentsNodeInfo := models.InterfaceToAppformixNetworkAgentsNode(appformixNetworkAgentsNode)
		// Read appformixNetworkAgent role node to get the node refs information
		appformixNetworkAgentsNodeData, err := r.getResource(
			defaultAppformixNetworkAgentsNodeResPath, appformixNetworkAgentsNodeInfo.UUID)
		if err != nil {
			return err
		}
		appformixNetworkAgentsNodeInfo = models.InterfaceToAppformixNetworkAgentsNode(
			appformixNetworkAgentsNodeData)
		a.ClusterInfo.AppformixNetworkAgentsNodes = append(
			a.ClusterInfo.AppformixNetworkAgentsNodes, appformixNetworkAgentsNodeInfo)
	}
	return nil
}

// nolint: gocyclo
func (a *AppformixData) updateNodeDetails(r *ResourceManager) error {
	m := make(map[string]bool)
	for _, node := range a.ClusterInfo.AppformixControllerNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, a); err != nil {
				return err
			}
		}
	}
	for _, node := range a.ClusterInfo.AppformixBareHostNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, a); err != nil {
				return err
			}
		}
	}
	for _, node := range a.ClusterInfo.AppformixOpenstackNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, a); err != nil {
				return err
			}
		}
	}
	for _, node := range a.ClusterInfo.AppformixComputeNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, a); err != nil {
				return err
			}
		}
	}
	for _, node := range a.ClusterInfo.AppformixNetworkAgentsNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, a); err != nil {
				return err
			}
		}
	}
	return nil
}

// nolint: gocyclo
func (a *AppformixData) updateClusterDetails(clusterID string, r *ResourceManager) error {
	ctx := context.Background()
	resp, err := r.APIServer.GetAppformixCluster(ctx, &services.GetAppformixClusterRequest{ID: clusterID})
	if err != nil {
		return err
	}
	a.ClusterInfo = resp.AppformixCluster

	// Expand appformix_controller back ref

	if err = a.interfaceToAppformixControllerNode(a.ClusterInfo.AppformixControllerNodes, r); err != nil {
		return err
	}

	// Expand appformix_bare_host back ref

	if err = a.interfaceToAppformixBareHostNode(a.ClusterInfo.AppformixBareHostNodes, r); err != nil {
		return err
	}

	// Expand appformix_openstack back ref

	if err = a.interfaceToAppformixOpenstackNode(a.ClusterInfo.AppformixOpenstackNodes, r); err != nil {
		return err
	}

	// Expand appformix_compute back ref

	if err = a.interfaceToAppformixComputeNode(a.ClusterInfo.AppformixComputeNodes, r); err != nil {
		return err
	}

	// Expand appformix_network_agents back ref

	if err = a.interfaceToAppformixNetworkAgentsNode(a.ClusterInfo.AppformixNetworkAgentsNodes, r); err != nil {
		return err
	}

	// get all nodes information
	if err = a.updateNodeDetails(r); err != nil {
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

func (k *KubernetesData) updateNodeDetails(r *ResourceManager) error {
	m := make(map[string]bool)
	for _, node := range k.ClusterInfo.KubernetesMasterNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, k); err != nil {
				return err
			}
		}
	}
	for _, node := range k.ClusterInfo.KubernetesNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, k); err != nil {
				return err
			}
		}
	}
	for _, node := range k.ClusterInfo.KubernetesKubemanagerNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, k); err != nil {
				return err
			}
		}
	}
	return nil
}

func (k *KubernetesData) interfaceToKubernetesNode(kubernetesNodes interface{}, r *ResourceManager) error {
	n, ok := kubernetesNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, kubernetesNode := range n {
		kubernetesNodeInfo := models.InterfaceToKubernetesNode(kubernetesNode)
		// Read kubernetes role node to get the node refs information
		kubernetesNodeData, err := r.getResource(
			defaultKubernetesNodeResPath, kubernetesNodeInfo.UUID)
		if err != nil {
			return err
		}
		kubernetesNodeInfo = models.InterfaceToKubernetesNode(
			kubernetesNodeData)
		k.ClusterInfo.KubernetesNodes = append(
			k.ClusterInfo.KubernetesNodes, kubernetesNodeInfo)
	}
	return nil
}

func (k *KubernetesData) interfaceToKubernetesMasterNode(kubernetesMasterNodes interface{}, r *ResourceManager) error {
	n, ok := kubernetesMasterNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, kubernetesMasterNode := range n {
		kubernetesMasterNodeInfo := models.InterfaceToKubernetesMasterNode(kubernetesMasterNode)
		// Read kubernetesMaster role node to get the node refs information
		kubernetesMasterNodeData, err := r.getResource(
			defaultKubernetesMasterNodeResPath, kubernetesMasterNodeInfo.UUID)
		if err != nil {
			return err
		}
		kubernetesMasterNodeInfo = models.InterfaceToKubernetesMasterNode(
			kubernetesMasterNodeData)
		k.ClusterInfo.KubernetesMasterNodes = append(
			k.ClusterInfo.KubernetesMasterNodes, kubernetesMasterNodeInfo)
	}
	return nil
}

func (k *KubernetesData) interfaceToKubernetesKubemanagerNode(
	kubernetesKubemanagerNodes interface{}, r *ResourceManager,
) error {
	n, ok := kubernetesKubemanagerNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, kubernetesKubemanagerNode := range n {
		kubernetesKubemanagerNodeInfo := models.InterfaceToKubernetesKubemanagerNode(kubernetesKubemanagerNode)
		// Read kubernetesKubemanager role node to get the node refs information
		kubernetesKubemanagerNodeData, err := r.getResource(
			defaultKubernetesKubemanagerNodeResPath, kubernetesKubemanagerNodeInfo.UUID)
		if err != nil {
			return err
		}
		kubernetesKubemanagerNodeInfo = models.InterfaceToKubernetesKubemanagerNode(
			kubernetesKubemanagerNodeData)
		k.ClusterInfo.KubernetesKubemanagerNodes = append(
			k.ClusterInfo.KubernetesKubemanagerNodes, kubernetesKubemanagerNodeInfo)
	}
	return nil
}

func (k *KubernetesData) updateClusterDetails(clusterID string, r *ResourceManager) error {
	rData, err := r.getResource(defaultK8sResourcePath, clusterID)
	if err != nil {
		return err
	}
	k.ClusterInfo = models.InterfaceToKubernetesCluster(rData)

	// Expand kubernetes node back ref
	if kubernetesNodes, ok := rData["kubernetes_nodes"]; ok {
		if err = k.interfaceToKubernetesNode(kubernetesNodes, r); err != nil {
			return err
		}
	}

	// Expand kubernetes_master back ref
	if kubernetesMasterNodes, ok := rData["kubernetes_master_nodes"]; ok {
		if err = k.interfaceToKubernetesMasterNode(kubernetesMasterNodes, r); err != nil {
			return err
		}
	}

	// Expand kubernetes_kubemanager back ref
	if kubernetesKubemanagerNodes, ok := rData["kubernetes_kubemanager_nodes"]; ok {
		if err = k.interfaceToKubernetesKubemanagerNode(kubernetesKubemanagerNodes, r); err != nil {
			return err
		}
	}

	// get all nodes information
	if err = k.updateNodeDetails(r); err != nil {
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

func (v *VCenterData) updateNodeDetails(r *ResourceManager) error {
	m := make(map[string]bool)
	for _, node := range v.ClusterInfo.VCenterPluginNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, v); err != nil {
				return err
			}
		}
	}
	for _, node := range v.ClusterInfo.VCenterComputes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, v); err != nil {
				return err
			}
		}
	}
	for _, node := range v.ClusterInfo.VCenterManagerNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, v); err != nil {
				return err
			}
		}
	}

	return nil
}

func (v *VCenterData) interfaceToVCenterPluginNode(
	vcenterPluginNodes interface{}, r *ResourceManager) error {
	for _, vcenterPluginNode := range vcenterPluginNodes.([]interface{}) {
		vcenterPluginNodeInfo := models.InterfaceToVCenterPluginNode(
			vcenterPluginNode.(map[string]interface{}))
		// Read vcenter_plugin role node to get the node refs information
		vcenterPluginNodeData, err := r.getResource(
			defaultVCenterPluginNodeResPath, vcenterPluginNodeInfo.UUID)
		if err != nil {
			return err
		}
		vcenterPluginNodeInfo = models.InterfaceToVCenterPluginNode(
			vcenterPluginNodeData)
		v.ClusterInfo.VCenterPluginNodes = append(
			v.ClusterInfo.VCenterPluginNodes, vcenterPluginNodeInfo)
	}
	return nil
}

func (v *VCenterData) interfaceToVCenterCompute(
	vcenterComputes interface{}, r *ResourceManager) error {
	for _, vcenterCompute := range vcenterComputes.([]interface{}) {
		vcenterComputeInfo := models.InterfaceToVCenterCompute(
			vcenterCompute.(map[string]interface{}))
		// Read vcenter_compute role node to get the node refs information
		vcenterComputeData, err := r.getResource(
			defaultVCenterComputeResPath, vcenterComputeInfo.UUID)
		if err != nil {
			return err
		}
		vcenterComputeInfo = models.InterfaceToVCenterCompute(
			vcenterComputeData)
		v.ClusterInfo.VCenterComputes = append(
			v.ClusterInfo.VCenterComputes, vcenterComputeInfo)
	}
	return nil
}

func (v *VCenterData) interfaceToVCenterManagerNode(
	vcenterManagerNodes interface{}, r *ResourceManager) error {
	for _, vcenterManagerNode := range vcenterManagerNodes.([]interface{}) {
		vcenterManagerNodeInfo := models.InterfaceToVCenterManagerNode(
			vcenterManagerNode.(map[string]interface{}))
		// Read vcenter_manager role node to get the node refs information
		vcenterManagerNodeData, err := r.getResource(
			defaultVCenterManagerNodeResPath, vcenterManagerNodeInfo.UUID)
		if err != nil {
			return err
		}
		vcenterManagerNodeInfo = models.InterfaceToVCenterManagerNode(
			vcenterManagerNodeData)
		v.ClusterInfo.VCenterManagerNodes = append(
			v.ClusterInfo.VCenterManagerNodes, vcenterManagerNodeInfo)
	}
	return nil
}

func (v *VCenterData) updateClusterDetails(clusterID string, r *ResourceManager) error {
	rData, err := r.getResource(defaultVCenterResourcePath, clusterID)
	if err != nil {
		return err
	}
	v.ClusterInfo = models.InterfaceToVCenter(rData)

	// Expand vcenter_plugin back ref
	if vcenterPluginNodes, ok := rData["vCenter_plugin_nodes"]; ok {
		if err = v.interfaceToVCenterPluginNode(vcenterPluginNodes, r); err != nil {
			return err
		}
	}
	// Expand vcenter_compute back ref
	if vcenterComputes, ok := rData["vCenter_computes"]; ok {
		if err = v.interfaceToVCenterCompute(vcenterComputes, r); err != nil {
			return err
		}
	}

	// Expand vcenter_manager back ref
	if vcenterManagerNodes, ok := rData["vCenter_manager_nodes"]; ok {
		if err = v.interfaceToVCenterManagerNode(vcenterManagerNodes, r); err != nil {
			return err
		}
	}

	// get all nodes information
	if err = v.updateNodeDetails(r); err != nil {
		return err
	}
	return nil
}

func (o *OpenstackData) getOpenstackPublicVip() (vip string) {
	vip = ""
	if o.ClusterInfo.OpenstackExternalVip != "" {
		vip = o.ClusterInfo.OpenstackExternalVip
	} else if o.ClusterInfo.OpenstackInternalVip != "" {
		vip = o.ClusterInfo.OpenstackInternalVip
	}

	return vip
}

func (o *OpenstackData) getOpenstackVipsFromAnnotation() (vips map[string]string) {
	vips = make(map[string]string)
	if a := o.ClusterInfo.GetAnnotations(); a != nil {
		for _, keyValuePair := range a.GetKeyValuePair() {
			switch keyValuePair.Key {
			case "keystone_vip":
				vips[identity] = keyValuePair.Value
			case "swift_vip":
				vips[swift] = keyValuePair.Value
			case "glance_vip":
				vips[glance] = keyValuePair.Value
			case "nova_vip":
				vips[nova] = keyValuePair.Value
			case "ironic_vip":
				vips[ironic] = keyValuePair.Value
			}
		}
	}
	return vips
}
func (o *OpenstackData) isSSLEnabled() bool {
	if g := o.ClusterInfo.GetKollaGlobals(); g != nil {
		for _, keyValuePair := range g.GetKeyValuePair() {
			switch keyValuePair.Key {
			case "kolla_enable_tls_external":
				yamlBoolTrue := []string{"yes", "true", "y", "on"}
				if format.ContainsString(
					yamlBoolTrue,
					strings.ToLower(keyValuePair.Value)) {
					return true
				}
			}
		}
	}
	return false
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
func (o *OpenstackData) updateNodeDetails(r *ResourceManager) error {
	m := make(map[string]bool)
	for _, node := range o.ClusterInfo.OpenstackControlNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	for _, node := range o.ClusterInfo.OpenstackNetworkNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	for _, node := range o.ClusterInfo.OpenstackStorageNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	for _, node := range o.ClusterInfo.OpenstackMonitoringNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	for _, node := range o.ClusterInfo.OpenstackComputeNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, o); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackControlNode(openstackControlNodes interface{}, r *ResourceManager) error {
	n, ok := openstackControlNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackControlNode := range n {
		openstackControlNodeInfo := models.InterfaceToOpenstackControlNode(openstackControlNode)
		// Read openstackControl role node to get the node refs information
		openstackControlNodeData, err := r.getResource(
			defaultOpenstackControlNodeResPath, openstackControlNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackControlNodeInfo = models.InterfaceToOpenstackControlNode(
			openstackControlNodeData)
		o.ClusterInfo.OpenstackControlNodes = append(
			o.ClusterInfo.OpenstackControlNodes, openstackControlNodeInfo)
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackMonitoringNode(
	openstackMonitoringNodes interface{},
	r *ResourceManager,
) error {
	n, ok := openstackMonitoringNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackMonitoringNode := range n {
		openstackMonitoringNodeInfo := models.InterfaceToOpenstackMonitoringNode(openstackMonitoringNode)
		// Read openstackMonitoring role node to get the node refs information
		openstackMonitoringNodeData, err := r.getResource(
			defaultOpenstackMonitoringNodeResPath, openstackMonitoringNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackMonitoringNodeInfo = models.InterfaceToOpenstackMonitoringNode(
			openstackMonitoringNodeData)
		o.ClusterInfo.OpenstackMonitoringNodes = append(

			o.ClusterInfo.OpenstackMonitoringNodes, openstackMonitoringNodeInfo)
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackNetworkNode(openstackNetworkNodes interface{}, r *ResourceManager) error {
	n, ok := openstackNetworkNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackNetworkNode := range n {
		openstackNetworkNodeInfo := models.InterfaceToOpenstackNetworkNode(openstackNetworkNode)
		// Read openstackNetwork role node to get the node refs information
		openstackNetworkNodeData, err := r.getResource(
			defaultOpenstackNetworkNodeResPath, openstackNetworkNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackNetworkNodeInfo = models.InterfaceToOpenstackNetworkNode(
			openstackNetworkNodeData)
		o.ClusterInfo.OpenstackNetworkNodes = append(
			o.ClusterInfo.OpenstackNetworkNodes, openstackNetworkNodeInfo)
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackStorageNode(openstackStorageNodes interface{}, r *ResourceManager) error {
	n, ok := openstackStorageNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackStorageNode := range n {
		openstackStorageNodeInfo := models.InterfaceToOpenstackStorageNode(openstackStorageNode)
		// Read openstackStorage role node to get the node refs information
		openstackStorageNodeData, err := r.getResource(
			defaultOpenstackStorageNodeResPath, openstackStorageNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackStorageNodeInfo = models.InterfaceToOpenstackStorageNode(
			openstackStorageNodeData)
		o.ClusterInfo.OpenstackStorageNodes = append(
			o.ClusterInfo.OpenstackStorageNodes, openstackStorageNodeInfo)
	}
	return nil
}

func (o *OpenstackData) interfaceToOpenstackComputeNode(openstackComputeNodes interface{}, r *ResourceManager) error {
	n, ok := openstackComputeNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, openstackComputeNode := range n {
		openstackComputeNodeInfo := models.InterfaceToOpenstackComputeNode(openstackComputeNode)
		// Read openstackCompute role node to get the node refs information
		openstackComputeNodeData, err := r.getResource(
			defaultOpenstackComputeNodeResPath, openstackComputeNodeInfo.UUID)
		if err != nil {
			return err
		}
		openstackComputeNodeInfo = models.InterfaceToOpenstackComputeNode(
			openstackComputeNodeData)
		o.ClusterInfo.OpenstackComputeNodes = append(
			o.ClusterInfo.OpenstackComputeNodes, openstackComputeNodeInfo)
	}
	return nil
}

// nolint: gocyclo
func (o *OpenstackData) updateClusterDetails(clusterID string, r *ResourceManager) error {
	rData, err := r.getResource(defaultOpenstackResourcePath, clusterID)
	if err != nil {
		return err
	}
	o.ClusterInfo = models.InterfaceToOpenstackCluster(rData)

	// Expand openstack_compute back ref
	if openstackComputeNodes, ok := rData["openstack_compute_nodes"]; ok {
		if err = o.interfaceToOpenstackComputeNode(openstackComputeNodes, r); err != nil {
			return err
		}
	}
	// Expand openstack_storage node back ref
	if openstackStorageNodes, ok := rData["openstack_storage_nodes"]; ok {
		if err = o.interfaceToOpenstackStorageNode(openstackStorageNodes, r); err != nil {
			return err
		}
	}
	// Expand openstack_network node back ref
	if openstackNetworkNodes, ok := rData["openstack_network_nodes"]; ok {
		if err = o.interfaceToOpenstackNetworkNode(openstackNetworkNodes, r); err != nil {
			return err
		}
	}
	// Expand openstack_monitoring node back ref
	if openstackMonitoringNodes, ok := rData["openstack_monitoring_nodes"]; ok {
		if err = o.interfaceToOpenstackMonitoringNode(openstackMonitoringNodes, r); err != nil {
			return err
		}
	}
	// Expand openstack_control node back ref
	if openstackControlNodes, ok := rData["openstack_control_nodes"]; ok {
		if err = o.interfaceToOpenstackControlNode(openstackControlNodes, r); err != nil {
			return err
		}
	}
	// get all nodes information
	if err = o.updateNodeDetails(r); err != nil {
		return err
	}
	return nil
}

func (o *OpenstackData) getControlNodeIPs() (nodeIPs []string) {
	for _, controlNode := range o.ClusterInfo.OpenstackControlNodes {
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
	for _, controlNode := range o.ClusterInfo.OpenstackControlNodes {
		if vip := o.getOpenstackPublicVip(); vip != "" {
			nodePorts[vip] = make(map[string]int64)
			format.InterfaceToInt64Map(nodePorts[vip])[identity] = controlNode.KeystonePublicPort
			format.InterfaceToInt64Map(nodePorts[vip])[nova] = controlNode.NovaPublicPort
			format.InterfaceToInt64Map(nodePorts[vip])[glance] = controlNode.GlancePublicPort
			format.InterfaceToInt64Map(nodePorts[vip])[ironic] = controlNode.IronicPublicPort
			return nodePorts
		}
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
	for _, storageNode := range o.ClusterInfo.OpenstackStorageNodes {
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
	for _, storageNode := range o.ClusterInfo.OpenstackStorageNodes {
		if vip := o.getOpenstackPublicVip(); vip != "" {
			nodePorts[vip] = make(map[string]int64)
			format.InterfaceToInt64Map(nodePorts[vip])[swift] = storageNode.SwiftPublicPort
			return nodePorts
		}
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

func (d *Data) getContrailExternalVip() (vip string) {
	if a := d.ClusterInfo.GetAnnotations(); a != nil {
		for _, keyValuePair := range a.GetKeyValuePair() {
			switch keyValuePair.Key {
			case "contrail_external_vip":
				vip = keyValuePair.Value
			}
		}
	}
	return vip
}

func (d *Data) isSSLEnabled() bool {
	if c := d.ClusterInfo.GetContrailConfiguration(); c != nil {
		for _, keyValuePair := range c.GetKeyValuePair() {
			switch keyValuePair.Key {
			case "SSL_ENABLE":
				yamlBoolTrue := []string{"yes", "true", "y", "on"}
				if format.ContainsString(
					yamlBoolTrue,
					strings.ToLower(keyValuePair.Value)) {
					return true
				}
			}
		}
	}
	return false
}

func (d *Data) addKeypair(keypair *models.Keypair) {
	d.keypairsInfo = append(d.keypairsInfo, keypair)
}

func (d *Data) addCredential(cred *models.Credential) {
	d.credsInfo = append(d.credsInfo, cred)
}

func (d *Data) addNode(node *models.Node) {
	d.NodesInfo = append(d.NodesInfo, node)
}

func (d *Data) interfaceToContrailVrouterNode(contrailVrouterNodes interface{}, r *ResourceManager) error {
	n, ok := contrailVrouterNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailVrouterNode := range n {
		contrailVrouterNodeInfo := models.InterfaceToContrailVrouterNode(contrailVrouterNode)
		// Read contrailVrouter role node to get the node refs information
		contrailVrouterNodeData, err := r.getResource(
			defaultContrailVrouterNodeResPath, contrailVrouterNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailVrouterNodeInfo = models.InterfaceToContrailVrouterNode(
			contrailVrouterNodeData)
		d.ClusterInfo.ContrailVrouterNodes = append(
			d.ClusterInfo.ContrailVrouterNodes, contrailVrouterNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailMCGWNode(contrailMCGWNodes interface{}, r *ResourceManager) error {
	n, ok := contrailMCGWNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailMCGWNode := range n {
		contrailMCGWNodeInfo := models.InterfaceToContrailMulticloudGWNode(contrailMCGWNode)
		// Read ContrailMulticloudGW role node to get the node refs information
		contrailMCGWNodeData, err := r.getResource(
			defaultContrailMCGWNodeResPath, contrailMCGWNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailMCGWNodeInfo = models.InterfaceToContrailMulticloudGWNode(
			contrailMCGWNodeData)
		d.ClusterInfo.ContrailMulticloudGWNodes = append(
			d.ClusterInfo.ContrailMulticloudGWNodes, contrailMCGWNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToVcenterFabricManagerNode(
	contrailVcenterFabricManagerNodes interface{},
	r *ResourceManager,
) error {
	n, ok := contrailVcenterFabricManagerNodes.([]interface{})
	if !ok {
		return nil
	}

	var vcenterUUIDs map[string]string
	vcenterUUIDs = make(map[string]string)
	for _, contrailVCFabricManagerNode := range n {
		contrailVCFabricManagerNodeInfo := models.InterfaceToContrailVcenterFabricManagerNode(contrailVCFabricManagerNode)
		// Read contrailVcenterFabricManager role node to get the node refs information
		contrailVCFabricManagerNodeData, err := r.getResource(
			defaultContrailVCFabricManagerNodeResPath, contrailVCFabricManagerNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailVCFabricManagerNodeInfo = models.InterfaceToContrailVcenterFabricManagerNode(
			contrailVCFabricManagerNodeData)
		d.ClusterInfo.ContrailVcenterFabricManagerNodes = append(
			d.ClusterInfo.ContrailVcenterFabricManagerNodes, contrailVCFabricManagerNodeInfo)

		// Collect vcenter UUIDs
		for _, vcenterRef := range contrailVCFabricManagerNodeInfo.VCenterRefs {
			vcenterUUIDs[vcenterRef.UUID] = vcenterRef.Href
		}
	}

	// vcenter data should be inferred from the references of contrail_vcenter_fabric_manager_node
	for vcenterUUID := range vcenterUUIDs {
		vCenterData := &VCenterData{Reader: r.APIServer}
		if err := vCenterData.updateClusterDetails(vcenterUUID, r); err != nil {
			return err
		}
		d.vcenterData = append(d.vcenterData, vCenterData)
	}

	return nil
}

func (d *Data) interfaceToContrailZTPTFTPNode(contrailZTPTFTPNodes interface{}, r *ResourceManager) error {
	n, ok := contrailZTPTFTPNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailZTPTFTPNode := range n {
		contrailZTPTFTPNodeInfo := models.InterfaceToContrailZTPTFTPNode(contrailZTPTFTPNode)
		// Read contrailZTPTFTP role node to get the node refs information
		contrailZTPTFTPNodeData, err := r.getResource(
			defaultContrailZTPTFTPNodeResPath, contrailZTPTFTPNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailZTPTFTPNodeInfo = models.InterfaceToContrailZTPTFTPNode(
			contrailZTPTFTPNodeData)
		d.ClusterInfo.ContrailZTPTFTPNodes = append(
			d.ClusterInfo.ContrailZTPTFTPNodes, contrailZTPTFTPNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailZTPDHCPNode(contrailZTPDHCPNodes interface{}, r *ResourceManager) error {
	n, ok := contrailZTPDHCPNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailZTPDHCPNode := range n {
		contrailZTPDHCPNodeInfo := models.InterfaceToContrailZTPDHCPNode(contrailZTPDHCPNode)
		// Read contrailZTPDHCP role node to get the node refs information
		contrailZTPDHCPNodeData, err := r.getResource(
			defaultContrailZTPDHCPNodeResPath, contrailZTPDHCPNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailZTPDHCPNodeInfo = models.InterfaceToContrailZTPDHCPNode(
			contrailZTPDHCPNodeData)
		d.ClusterInfo.ContrailZTPDHCPNodes = append(
			d.ClusterInfo.ContrailZTPDHCPNodes, contrailZTPDHCPNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailServiceNode(contrailServiceNodes interface{}, r *ResourceManager) error {
	n, ok := contrailServiceNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailServiceNode := range n {
		contrailServiceNodeInfo := models.InterfaceToContrailServiceNode(contrailServiceNode)
		// Read contrailService role node to get the node refs information
		contrailServiceNodeData, err := r.getResource(
			defaultContrailServiceNodeResPath, contrailServiceNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailServiceNodeInfo = models.InterfaceToContrailServiceNode(
			contrailServiceNodeData)
		d.ClusterInfo.ContrailServiceNodes = append(
			d.ClusterInfo.ContrailServiceNodes, contrailServiceNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailAnalyticsDatabaseNode(
	contrailAnalyticsDatabaseNodes interface{},
	r *ResourceManager,
) error {
	n, ok := contrailAnalyticsDatabaseNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailAnalyticsDatabaseNode := range n {
		contrailAnalyticsDatabaseNodeInfo := models.InterfaceToContrailAnalyticsDatabaseNode(contrailAnalyticsDatabaseNode)
		// Read contrailAnalyticsDatabase role node to get the node refs information
		contrailAnalyticsDatabaseNodeData, err := r.getResource(
			defaultContrailAnalyticsDatabaseNodeResPath, contrailAnalyticsDatabaseNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailAnalyticsDatabaseNodeInfo = models.InterfaceToContrailAnalyticsDatabaseNode(
			contrailAnalyticsDatabaseNodeData)
		d.ClusterInfo.ContrailAnalyticsDatabaseNodes = append(
			d.ClusterInfo.ContrailAnalyticsDatabaseNodes, contrailAnalyticsDatabaseNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailAnalyticsNode(contrailAnalyticsNodes interface{}, r *ResourceManager) error {
	n, ok := contrailAnalyticsNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailAnalyticsNode := range n {
		contrailAnalyticsNodeInfo := models.InterfaceToContrailAnalyticsNode(contrailAnalyticsNode)
		// Read contrailAnalytics role node to get the node refs information
		contrailAnalyticsNodeData, err := r.getResource(
			defaultContrailAnalyticsNodeResPath, contrailAnalyticsNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailAnalyticsNodeInfo = models.InterfaceToContrailAnalyticsNode(
			contrailAnalyticsNodeData)
		d.ClusterInfo.ContrailAnalyticsNodes = append(
			d.ClusterInfo.ContrailAnalyticsNodes, contrailAnalyticsNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailAnalyticsAlarmNode(
	contrailAnalyticsAlarmNodes interface{},
	r *ResourceManager,
) error {
	n, ok := contrailAnalyticsAlarmNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailAnalyticsAlarmNode := range n {
		contrailAnalyticsAlarmNodeInfo := models.InterfaceToContrailAnalyticsAlarmNode(contrailAnalyticsAlarmNode)
		// Read contrailAnalyticsAlarm role node to get the node refs information
		contrailAnalyticsAlarmNodeData, err := r.getResource(
			defaultContrailAnalyticsAlarmNodeResPath, contrailAnalyticsAlarmNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailAnalyticsAlarmNodeInfo = models.InterfaceToContrailAnalyticsAlarmNode(
			contrailAnalyticsAlarmNodeData)
		d.ClusterInfo.ContrailAnalyticsAlarmNodes = append(
			d.ClusterInfo.ContrailAnalyticsAlarmNodes, contrailAnalyticsAlarmNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailAnalyticsSNMPNode(contrailAnalyticsSNMPNodes interface{}, r *ResourceManager) error {
	n, ok := contrailAnalyticsSNMPNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailAnalyticsSNMPNode := range n {
		contrailAnalyticsSNMPNodeInfo := models.InterfaceToContrailAnalyticsSNMPNode(contrailAnalyticsSNMPNode)
		// Read contrailAnalytics role node to get the node refs information
		contrailAnalyticsSNMPNodeData, err := r.getResource(
			defaultContrailAnalyticsSNMPNodeResPath, contrailAnalyticsSNMPNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailAnalyticsSNMPNodeInfo = models.InterfaceToContrailAnalyticsSNMPNode(
			contrailAnalyticsSNMPNodeData)
		d.ClusterInfo.ContrailAnalyticsSNMPNodes = append(
			d.ClusterInfo.ContrailAnalyticsSNMPNodes, contrailAnalyticsSNMPNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailWebuiNode(contrailWebuiNodes interface{}, r *ResourceManager) error {
	n, ok := contrailWebuiNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailWebuiNode := range n {
		contrailWebuiNodeInfo := models.InterfaceToContrailWebuiNode(contrailWebuiNode)
		// Read contrailWebui role node to get the node refs information
		contrailWebuiNodeData, err := r.getResource(
			defaultContrailWebuiNodeResPath, contrailWebuiNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailWebuiNodeInfo = models.InterfaceToContrailWebuiNode(
			contrailWebuiNodeData)
		d.ClusterInfo.ContrailWebuiNodes = append(
			d.ClusterInfo.ContrailWebuiNodes, contrailWebuiNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailControlNode(contrailControlNodes interface{}, r *ResourceManager) error {
	n, ok := contrailControlNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailControlNode := range n {
		contrailControlNodeInfo := models.InterfaceToContrailControlNode(contrailControlNode)
		// Read contrailControl role node to get the node refs information
		contrailControlNodeData, err := r.getResource(
			defaultContrailControlNodeResPath, contrailControlNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailControlNodeInfo = models.InterfaceToContrailControlNode(

			contrailControlNodeData)
		d.ClusterInfo.ContrailControlNodes = append(
			d.ClusterInfo.ContrailControlNodes, contrailControlNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailConfigDatabaseNode(
	contrailConfigDatabaseNodes interface{},
	r *ResourceManager,
) error {
	n, ok := contrailConfigDatabaseNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailConfigDatabaseNode := range n {
		contrailConfigDatabaseNodeInfo := models.InterfaceToContrailConfigDatabaseNode(contrailConfigDatabaseNode)
		// Read contrailConfigDatabase role node to get the node refs information
		contrailConfigDatabaseNodeData, err := r.getResource(
			defaultContrailConfigDatabaseNodeResPath, contrailConfigDatabaseNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailConfigDatabaseNodeInfo = models.InterfaceToContrailConfigDatabaseNode(
			contrailConfigDatabaseNodeData)
		d.ClusterInfo.ContrailConfigDatabaseNodes = append(
			d.ClusterInfo.ContrailConfigDatabaseNodes, contrailConfigDatabaseNodeInfo)
	}
	return nil
}

func (d *Data) interfaceToContrailConfigNode(contrailConfigNodes interface{}, r *ResourceManager) error {
	n, ok := contrailConfigNodes.([]interface{})
	if !ok {
		return nil
	}

	for _, contrailConfigNode := range n {
		contrailConfigNodeInfo := models.InterfaceToContrailConfigNode(contrailConfigNode)
		// Read contrailConfig role node to get the node refs information
		contrailConfigNodeData, err := r.getResource(
			defaultContrailConfigNodeResPath, contrailConfigNodeInfo.UUID)
		if err != nil {
			return err
		}
		contrailConfigNodeInfo = models.InterfaceToContrailConfigNode(
			contrailConfigNodeData)
		d.ClusterInfo.ContrailConfigNodes = append(
			d.ClusterInfo.ContrailConfigNodes, contrailConfigNodeInfo)
	}
	return nil
}

// nolint: gocyclo
func (d *Data) updateNodeDetails(r *ResourceManager) error {
	m := make(map[string]bool)
	for _, node := range d.ClusterInfo.ContrailConfigNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailConfigDatabaseNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailControlNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailWebuiNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailAnalyticsNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailAnalyticsDatabaseNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailAnalyticsAlarmNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailAnalyticsSNMPNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailVcenterFabricManagerNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailVrouterNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailMulticloudGWNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	for _, node := range d.ClusterInfo.ContrailServiceNodes {
		for _, nodeRef := range node.NodeRefs {
			if err := r.getNode(nodeRef.UUID, m, d); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Data) updateCloudDetails(r *ResourceManager) error {
	for _, cloudRef := range d.ClusterInfo.CloudRefs {
		cloudObject, err := cloud.GetCloud(context.Background(), r.APIServer, cloudRef.UUID)
		if err != nil {
			return err
		}
		d.CloudInfo = append(d.CloudInfo, cloudObject)
	}
	return nil
}

// nolint: gocyclo
func (d *Data) updateClusterDetails(clusterID string, r *ResourceManager) error {
	rData, err := r.getResource(defaultResourcePath, clusterID)
	if err != nil {
		return err
	}
	d.ClusterInfo = models.InterfaceToContrailCluster(rData)

	// Expand config node back ref
	if configNodes, ok := rData["contrail_config_nodes"]; ok {
		if err = d.interfaceToContrailConfigNode(configNodes, r); err != nil {
			return err
		}
	}
	// Expand config database node back ref
	if configDBNodes, ok := rData["contrail_config_database_nodes"]; ok {
		if err = d.interfaceToContrailConfigDatabaseNode(configDBNodes, r); err != nil {
			return err
		}
	}
	// Expand control node back ref
	if controlNodes, ok := rData["contrail_control_nodes"]; ok {
		if err = d.interfaceToContrailControlNode(controlNodes, r); err != nil {
			return err
		}
	}
	// Expand webui node back ref
	if webuiNodes, ok := rData["contrail_webui_nodes"]; ok {
		if err = d.interfaceToContrailWebuiNode(webuiNodes, r); err != nil {
			return err
		}
	}
	// Expand analytics node back ref
	if analyticsNodes, ok := rData["contrail_analytics_nodes"]; ok {
		if err = d.interfaceToContrailAnalyticsNode(analyticsNodes, r); err != nil {
			return err
		}
	}
	// Expand analytics database node back ref
	if analyticsDBNodes, ok := rData["contrail_analytics_database_nodes"]; ok {
		if err = d.interfaceToContrailAnalyticsDatabaseNode(analyticsDBNodes, r); err != nil {
			return err
		}
	}
	// Expand analytics alarm node back ref
	if analyticsAlarmNodes, ok := rData["contrail_analytics_alarm_nodes"]; ok {
		if err = d.interfaceToContrailAnalyticsAlarmNode(analyticsAlarmNodes, r); err != nil {
			return err
		}
	}
	// Expand analytics snmp node back ref
	if analyticsSNMPNodes, ok := rData["contrail_analytics_snmp_nodes"]; ok {
		if err = d.interfaceToContrailAnalyticsSNMPNode(analyticsSNMPNodes, r); err != nil {
			return err
		}
	}
	// Expand vcenter fabric manager role node back ref
	if vcFabricManagerNodes, ok := rData["contrail_vcenter_fabric_manager_nodes"]; ok {
		if err = d.interfaceToVcenterFabricManagerNode(vcFabricManagerNodes, r); err != nil {
			return err
		}
	}
	// Expand vouter node back ref
	if vrouterNodes, ok := rData["contrail_vrouter_nodes"]; ok {
		if err = d.interfaceToContrailVrouterNode(vrouterNodes, r); err != nil {
			return err
		}
	}
	if mcGWNodes, ok := rData["contrail_multicloud_gw_nodes"]; ok {
		if err = d.interfaceToContrailMCGWNode(mcGWNodes, r); err != nil {
			return err
		}
	}
	// Expand csn node back ref
	if csnNodes, ok := rData["contrail_service_nodes"]; ok {
		if err = d.interfaceToContrailServiceNode(csnNodes, r); err != nil {
			return err
		}
	}
	// Expand tftp node back ref
	if tftpNodes, ok := rData["contrail_ztp_tftp_nodes"]; ok {
		if err = d.interfaceToContrailZTPTFTPNode(tftpNodes, r); err != nil {
			return err
		}
	}
	// Expand dhcp node back ref
	if dhcpNodes, ok := rData["contrail_ztp_dhcp_nodes"]; ok {
		if err = d.interfaceToContrailZTPDHCPNode(dhcpNodes, r); err != nil {
			return err
		}
	}
	// get all nodes information
	if err = d.updateNodeDetails(r); err != nil {
		return err
	}

	// get all cloud information
	if err = d.updateCloudDetails(r); err != nil {
		return err
	}
	d.DefaultSSHUser, d.DefaultSSHPassword, d.DefaultSSHKey, err = r.getDefaultCredential()
	if err != nil {
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

// GetK8sClusterInfo gets k8s cluster details
func (d *Data) GetK8sClusterInfo() *models.KubernetesCluster {
	if d.getK8sClusterData() != nil {
		return d.getK8sClusterData().ClusterInfo
	}
	return nil
}

func (d *Data) getVCenterClusterData() *VCenterData {

	if len(d.vcenterData) > 0 {
		return d.vcenterData[0]
	}
	return nil
}

// GetVCenterClusterInfo gets VCenter cluster details
func (d *Data) GetVCenterClusterInfo() *models.VCenter {
	if d.getVCenterClusterData() != nil {
		return d.getVCenterClusterData().ClusterInfo
	}
	return nil
}

func (d *Data) getOpenstackClusterData() *OpenstackData {
	// One openstack cluster is the supported topology
	if len(d.openstackClusterData) < 1 {
		return nil
	}
	return d.openstackClusterData[0]
}

// GetOpenstackClusterInfo gets openstack cluster details
func (d *Data) GetOpenstackClusterInfo() *models.OpenstackCluster {
	if cd := d.getOpenstackClusterData(); cd != nil {
		return cd.ClusterInfo
	}
	return nil
}

// KeystoneAdminCredential returns admin credentials from deploy data object.
func (d *Data) KeystoneAdminCredential() (adminUser, adminPassword string) {
	if d.ClusterInfo.Orchestrator != openstack {
		return "", ""
	}
	adminUser, adminPassword = defaultAdminUser, defaultAdminPassword

	for _, kvp := range d.GetOpenstackClusterInfo().GetKollaPasswords().GetKeyValuePair() {
		switch kvp.Key {
		case "keystone_admin_user":
			adminUser = kvp.Value
		case "keystone_admin_password":
			adminPassword = kvp.Value
		}
	}

	return adminUser, adminPassword
}

// GetAllKeypairsInfo gets kepair details
func (d *Data) GetAllKeypairsInfo() []*models.Keypair {
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

// GetAllCredsInfo gets credential details
func (d *Data) GetAllCredsInfo() []*models.Credential {
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

// GetAllNodesInfo gets all node details
func (d *Data) GetAllNodesInfo() []*models.Node {
	nodes := d.NodesInfo
	if d.getOpenstackClusterData() != nil {
		nodes = append(nodes, d.getOpenstackClusterData().nodesInfo...)
	}
	if d.getK8sClusterData() != nil {
		nodes = append(nodes, d.getK8sClusterData().nodesInfo...)
	}
	if d.getAppformixClusterData() != nil {
		nodes = append(nodes, d.getAppformixClusterData().nodesInfo...)
	}

	if d.GetXflowData() != nil {
		nodes = append(nodes, d.GetXflowData().getNodes()...)
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

// GetAppformixMonitoredNodes gets appformix monitored nodes
func (d *Data) GetAppformixMonitoredNodes() []*models.Node {

	// Get all unique contrail+flows nodes.
	nodes := d.NodesInfo
	if d.GetXflowData() != nil {
		nodes = append(nodes, d.GetXflowData().getNodes()...)
	}

	var uniqueNodes []*models.Node
	m := make(map[string]bool)

	for _, node := range nodes {
		if _, ok := m[node.UUID]; !ok {
			m[node.UUID] = true
			uniqueNodes = append(uniqueNodes, node)
		}
	}
	// At this point uniqueNodes has all contrail+flows nodes.

	// For all contrail and flows nodes which donot have appformix role,
	// add appformix barehost role
	// This makes appformix playbook install appformix-manager and monitor the node
	var monitoredNodes []*models.Node
	monitoredNodes = nil
	for _, uniqueNode := range uniqueNodes {
		isFound := false
		if d.getAppformixClusterData() != nil {
			for _, node := range d.getAppformixClusterData().nodesInfo {
				if uniqueNode.UUID == node.UUID {
					isFound = true
					break
				}
			}
		}

		if isFound == false {
			// add barehost role to this node
			// ie add it to AppformixCluser bare-host list
			monitoredNodes = append(monitoredNodes, uniqueNode)
		}
	}

	return monitoredNodes
}

func (d *Data) getConfigNodeIPs() (nodeIPs []string) {
	for _, configNode := range d.ClusterInfo.ContrailConfigNodes {
		for _, nodeRef := range configNode.NodeRefs {
			for _, node := range d.NodesInfo {
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
	for _, configNode := range d.ClusterInfo.ContrailConfigNodes {
		if vip := d.getContrailExternalVip(); vip != "" {
			nodePorts[vip] = make(map[string]int64)
			format.InterfaceToInt64Map(nodePorts[vip])[config] = configNode.APIPublicPort
			return nodePorts
		}
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
	for _, analyticsNode := range d.ClusterInfo.ContrailAnalyticsNodes {
		for _, nodeRef := range analyticsNode.NodeRefs {
			for _, node := range d.NodesInfo {
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
	for _, analyticsNode := range d.ClusterInfo.ContrailAnalyticsNodes {
		if vip := d.getContrailExternalVip(); vip != "" {
			nodePorts[vip] = make(map[string]int64)
			format.InterfaceToInt64Map(nodePorts[vip])[analytics] = analyticsNode.APIPublicPort
			return nodePorts
		}
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
	for _, webuiNode := range d.ClusterInfo.ContrailWebuiNodes {
		for _, nodeRef := range webuiNode.NodeRefs {
			for _, node := range d.NodesInfo {
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
	for _, webuiNode := range d.ClusterInfo.ContrailWebuiNodes {
		if vip := d.getContrailExternalVip(); vip != "" {
			nodePorts[vip] = make(map[string]int64)
			format.InterfaceToInt64Map(nodePorts[vip])[webui] = webuiNode.PublicPort
			return nodePorts
		}
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

// GetAppformixClusterInfo gets appformix cluster details
func (d *Data) GetAppformixClusterInfo() *models.AppformixCluster {
	if d.getAppformixClusterData() != nil {
		return d.getAppformixClusterData().ClusterInfo
	}
	return nil
}

// GetXflowData gets Xflow data
func (d *Data) GetXflowData() *XflowData {
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
	appformixClusterInfo := d.GetAppformixClusterInfo()
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
	appformixClusterInfo := d.GetAppformixClusterInfo()
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

func (x *XflowData) updateClusterDetails(ctx context.Context, uuid string, r *ResourceManager) error {
	resp, err := r.APIServer.GetAppformixFlows(ctx, &services.GetAppformixFlowsRequest{ID: uuid})

	if err != nil {
		return err
	}

	x.ClusterInfo = resp.AppformixFlows

	for i, appformixFlowsNode := range x.ClusterInfo.AppformixFlowsNodes {
		err = x.updateAppformixFlowsNode(ctx, i, appformixFlowsNode.UUID, r)
		if err != nil {
			return err
		}
	}

	return nil
}

func (x *XflowData) updateAppformixFlowsNode(ctx context.Context, i int, uuid string, r *ResourceManager) error {
	resp, err := r.APIServer.GetAppformixFlowsNode(ctx, &services.GetAppformixFlowsNodeRequest{ID: uuid})

	if err != nil {
		return err
	}

	// update the x.ClusterInfo with detailed flow nodes
	x.ClusterInfo.AppformixFlowsNodes[i] = resp.AppformixFlowsNode

	for _, nodeRef := range resp.AppformixFlowsNode.NodeRefs {
		err = x.updateNodeInfo(ctx, nodeRef.UUID, r)
		if err != nil {
			return err
		}
	}

	return nil
}

func (x *XflowData) updateNodeInfo(ctx context.Context, nodeUUID string, r *ResourceManager) error {
	resp, err := r.APIServer.GetNode(ctx, &services.GetNodeRequest{ID: nodeUUID})

	if err != nil {
		return err
	}

	x.NodesInfo[nodeUUID] = resp.Node

	return nil
}
