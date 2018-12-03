package cluster

import (
	"context"

	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/models"
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
}

// KubernetesData is the representation of kubernetes cluster details.
type KubernetesData struct {
	clusterInfo  *models.KubernetesCluster
	nodesInfo    []*models.Node
	keypairsInfo []*models.Keypair
	credsInfo    []*models.Credential
}

// Data is the representation of cluster details.
type Data struct {
	clusterInfo           *models.ContrailCluster
	nodesInfo             []*models.Node
	keypairsInfo          []*models.Keypair
	credsInfo             []*models.Credential
	cloudInfo             []*models.Cloud
	openstackClusterData  []*OpenstackData
	kubernetesClusterData []*KubernetesData
	// TODO (ijohnson): Add gce/aws/kvm info
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

func (k *KubernetesData) interfaceToKubernetesNode(
	kubernetesNodes interface{}, c *Cluster) error {
	for _, kubernetesNode := range kubernetesNodes.([]interface{}) {
		kubernetesNodeInfo := models.InterfaceToKubernetesNode(
			kubernetesNode.(map[string]interface{}))
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

func (k *KubernetesData) interfaceToKubernetesMasterNode(
	kubernetesMasterNodes interface{}, c *Cluster) error {
	for _, kubernetesMasterNode := range kubernetesMasterNodes.([]interface{}) {
		kubernetesMasterNodeInfo := models.InterfaceToKubernetesMasterNode(
			kubernetesMasterNode.(map[string]interface{}))
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
	kubernetesKubemanagerNodes interface{}, c *Cluster) error {
	for _, kubernetesKubemanagerNode := range kubernetesKubemanagerNodes.([]interface{}) {
		kubernetesKubemanagerNodeInfo := models.InterfaceToKubernetesKubemanagerNode(
			kubernetesKubemanagerNode.(map[string]interface{}))
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

func (o *OpenstackData) interfaceToOpenstackControlNode(
	openstackControlNodes interface{}, c *Cluster) error {
	for _, openstackControlNode := range openstackControlNodes.([]interface{}) {
		openstackControlNodeInfo := models.InterfaceToOpenstackControlNode(
			openstackControlNode.(map[string]interface{}))
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

func (o *OpenstackData) interfaceToOpenstackMonitoringNode(
	openstackMonitoringNodes interface{}, c *Cluster) error {
	for _, openstackMonitoringNode := range openstackMonitoringNodes.([]interface{}) {
		openstackMonitoringNodeInfo := models.InterfaceToOpenstackMonitoringNode(
			openstackMonitoringNode.(map[string]interface{}))
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

func (o *OpenstackData) interfaceToOpenstackNetworkNode(
	openstackNetworkNodes interface{}, c *Cluster) error {
	for _, openstackNetworkNode := range openstackNetworkNodes.([]interface{}) {
		openstackNetworkNodeInfo := models.InterfaceToOpenstackNetworkNode(
			openstackNetworkNode.(map[string]interface{}))
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

func (o *OpenstackData) interfaceToOpenstackStorageNode(
	openstackStorageNodes interface{}, c *Cluster) error {
	for _, openstackStorageNode := range openstackStorageNodes.([]interface{}) {
		openstackStorageNodeInfo := models.InterfaceToOpenstackStorageNode(
			openstackStorageNode.(map[string]interface{}))
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

func (o *OpenstackData) interfaceToOpenstackComputeNode(
	openstackComputeNodes interface{}, c *Cluster) error {
	for _, openstackComputeNode := range openstackComputeNodes.([]interface{}) {
		openstackComputeNodeInfo := models.InterfaceToOpenstackComputeNode(
			openstackComputeNode.(map[string]interface{}))
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

func (d *Data) addKeypair(keypair *models.Keypair) {
	d.keypairsInfo = append(d.keypairsInfo, keypair)
}

func (d *Data) addCredential(cred *models.Credential) {
	d.credsInfo = append(d.credsInfo, cred)
}

func (d *Data) addNode(node *models.Node) {
	d.nodesInfo = append(d.nodesInfo, node)
}

func (d *Data) interfaceToContrailVrouterNode(
	contrailVrouterNodes interface{}, c *Cluster) error {
	for _, contrailVrouterNode := range contrailVrouterNodes.([]interface{}) {
		contrailVrouterNodeInfo := models.InterfaceToContrailVrouterNode(
			contrailVrouterNode.(map[string]interface{}))
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

func (d *Data) interfaceToContrailZTPTFTPNode(
	contrailZTPTFTPNodes interface{}, c *Cluster) error {
	for _, contrailZTPTFTPNode := range contrailZTPTFTPNodes.([]interface{}) {
		contrailZTPTFTPNodeInfo := models.InterfaceToContrailZTPTFTPNode(
			contrailZTPTFTPNode.(map[string]interface{}))
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

func (d *Data) interfaceToContrailZTPDHCPNode(
	contrailZTPDHCPNodes interface{}, c *Cluster) error {
	for _, contrailZTPDHCPNode := range contrailZTPDHCPNodes.([]interface{}) {
		contrailZTPDHCPNodeInfo := models.InterfaceToContrailZTPDHCPNode(
			contrailZTPDHCPNode.(map[string]interface{}))
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

func (d *Data) interfaceToContrailServiceNode(
	contrailServiceNodes interface{}, c *Cluster) error {
	for _, contrailServiceNode := range contrailServiceNodes.([]interface{}) {
		contrailServiceNodeInfo := models.InterfaceToContrailServiceNode(
			contrailServiceNode.(map[string]interface{}))
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

func (d *Data) interfaceToContrailAnalyticsDatabaseNode(
	contrailAnalyticsDatabaseNodes interface{}, c *Cluster) error {
	for _, contrailAnalyticsDatabaseNode := range contrailAnalyticsDatabaseNodes.([]interface{}) {
		contrailAnalyticsDatabaseNodeInfo := models.InterfaceToContrailAnalyticsDatabaseNode(
			contrailAnalyticsDatabaseNode.(map[string]interface{}))
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

func (d *Data) interfaceToContrailAnalyticsNode(
	contrailAnalyticsNodes interface{}, c *Cluster) error {
	for _, contrailAnalyticsNode := range contrailAnalyticsNodes.([]interface{}) {
		contrailAnalyticsNodeInfo := models.InterfaceToContrailAnalyticsNode(
			contrailAnalyticsNode.(map[string]interface{}))
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

func (d *Data) interfaceToContrailWebuiNode(
	contrailWebuiNodes interface{}, c *Cluster) error {
	for _, contrailWebuiNode := range contrailWebuiNodes.([]interface{}) {
		contrailWebuiNodeInfo := models.InterfaceToContrailWebuiNode(
			contrailWebuiNode.(map[string]interface{}))
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

func (d *Data) interfaceToContrailControlNode(
	contrailControlNodes interface{}, c *Cluster) error {
	for _, contrailControlNode := range contrailControlNodes.([]interface{}) {
		contrailControlNodeInfo := models.InterfaceToContrailControlNode(
			contrailControlNode.(map[string]interface{}))
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

func (d *Data) interfaceToContrailConfigDatabaseNode(
	contrailConfigDatabaseNodes interface{}, c *Cluster) error {
	for _, contrailConfigDatabaseNode := range contrailConfigDatabaseNodes.([]interface{}) {
		contrailConfigDatabaseNodeInfo := models.InterfaceToContrailConfigDatabaseNode(
			contrailConfigDatabaseNode.(map[string]interface{}))
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

func (d *Data) interfaceToContrailConfigNode(
	contrailConfigNodes interface{}, c *Cluster) error {
	for _, contrailConfigNode := range contrailConfigNodes.([]interface{}) {
		contrailConfigNodeInfo := models.InterfaceToContrailConfigNode(

			contrailConfigNode.(map[string]interface{}))
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

func (d *Data) getCloudRefs() ([]*models.Cloud, error) {
	return nil, nil
}
