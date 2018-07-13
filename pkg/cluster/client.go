package cluster

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	uuid "github.com/satori/go.uuid"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func (c *Cluster) createEndpoint(parentUUID, name, publicURL, privateURL string) error {
	endpoint := map[string]string{
		"parent_type":  defaultResource,
		"parent_uuid":  parentUUID,
		"name":         fmt.Sprintf("%s-%s", name, uuid.NewV4().String()),
		"display_name": name,
		"prefix":       name,
		"public_url":   publicURL,
		"private_url":  privateURL,
	}
	endpointData := map[string]map[string]string{"endpoint": endpoint}
	c.log.Infof("Creating endpoint: %s, %s", name, publicURL)
	var endpointResponse map[string]interface{}
	resURI := fmt.Sprintf("%ss", defaultEndpointResPath)
	_, err := c.APIServer.Create(resURI, &endpointData, &endpointResponse)
	return err
}

func (c *Cluster) getEndpoints(parentUUIDs []string) (endpointIDs []string, err error) {
	values := url.Values{
		services.ParentUUIDsKey: parentUUIDs,
		services.ParentTypeKey:  []string{defaultResource},
	}
	var endpointList map[string][]interface{}
	resURI := fmt.Sprintf("%ss?%s", defaultEndpointResPath, values.Encode())
	c.log.Infof("Reading endpoints: %s", resURI)
	_, err = c.APIServer.Read(resURI, &endpointList)
	if err != nil {
		return nil, err
	}
	for _, rawEndpoint := range endpointList[defaultEndpointRes+"s"] {
		endpointID := rawEndpoint.(map[string]interface{})["uuid"].(string)
		endpointIDs = append(endpointIDs, endpointID)
	}
	return endpointIDs, nil
}

func (c *Cluster) deleteEndpoint(endpointUUID string) error {
	var output map[string]interface{}
	resURI := fmt.Sprintf("%s/%s", defaultEndpointResPath, endpointUUID)
	c.log.Infof("Deleting endpoint: %s", resURI)
	_, err := c.APIServer.Delete(resURI, &output)
	return err
}

func (c *Cluster) getResource(resPath string, resID string) (map[string]interface{}, error) {
	var rawResInfo map[string]interface{}
	resURI := fmt.Sprintf("%s/%s", resPath, resID)
	c.log.Infof("Reading: %s", resURI)
	_, err := c.APIServer.Read(resURI, &rawResInfo)
	if err != nil {
		return nil, err
	}
	res := strings.TrimLeft(resPath, "/")
	data, ok := rawResInfo[res].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid resource type")
	}
	return data, nil
}

func (c *Cluster) interfaceToKubernetesNode(kubernetesNodes interface{}) ([]*models.KubernetesNode, error) {
	var kubernetesNodesData []*models.KubernetesNode
	for _, kubernetesNode := range kubernetesNodes.([]interface{}) {
		kubernetesNodeInfo := models.InterfaceToKubernetesNode(kubernetesNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		kubernetesNodeData, err := c.getResource(defaultKubernetesNodeResPath, kubernetesNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of kubernetes node")
		}
		kubernetesNodeInfo = models.InterfaceToKubernetesNode(kubernetesNodeData)
		kubernetesNodesData = append(kubernetesNodesData, kubernetesNodeInfo)
	}
	return kubernetesNodesData, nil
}

func (c *Cluster) interfaceToKubernetesMasterNode(
	kubernetesMasterNodes interface{}) ([]*models.KubernetesMasterNode, error) {
	var kubernetesMasterNodesData []*models.KubernetesMasterNode
	for _, kubernetesMasterNode := range kubernetesMasterNodes.([]interface{}) {
		kubernetesMasterNodeInfo := models.InterfaceToKubernetesMasterNode(kubernetesMasterNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		kubernetesMasterNodeData, err := c.getResource(defaultKubernetesMasterNodeResPath, kubernetesMasterNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of kubernetes master node")
		}
		kubernetesMasterNodeInfo = models.InterfaceToKubernetesMasterNode(kubernetesMasterNodeData)
		kubernetesMasterNodesData = append(kubernetesMasterNodesData, kubernetesMasterNodeInfo)
	}
	return kubernetesMasterNodesData, nil
}

func (c *Cluster) interfaceToOpenstackControlNode(
	openstackControlNodes interface{}) ([]*models.OpenstackControlNode, error) {
	var openstackControlNodesData []*models.OpenstackControlNode
	for _, openstackControlNode := range openstackControlNodes.([]interface{}) {
		openstackControlNodeInfo := models.InterfaceToOpenstackControlNode(openstackControlNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		openstackControlNodeData, err := c.getResource(defaultOpenstackControlNodeResPath, openstackControlNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of openstack control node")
		}
		openstackControlNodeInfo = models.InterfaceToOpenstackControlNode(openstackControlNodeData)
		openstackControlNodesData = append(openstackControlNodesData, openstackControlNodeInfo)
	}
	return openstackControlNodesData, nil
}

func (c *Cluster) interfaceToOpenstackMonitoringNode(
	openstackMonitoringNodes interface{}) ([]*models.OpenstackMonitoringNode, error) {
	var openstackMonitoringNodesData []*models.OpenstackMonitoringNode
	for _, openstackMonitoringNode := range openstackMonitoringNodes.([]interface{}) {
		openstackMonitoringNodeInfo := models.InterfaceToOpenstackMonitoringNode(
			openstackMonitoringNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		openstackMonitoringNodeData, err := c.getResource(
			defaultOpenstackMonitoringNodeResPath, openstackMonitoringNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of openstack monitoring node")
		}
		openstackMonitoringNodeInfo = models.InterfaceToOpenstackMonitoringNode(openstackMonitoringNodeData)
		openstackMonitoringNodesData = append(openstackMonitoringNodesData, openstackMonitoringNodeInfo)
	}
	return openstackMonitoringNodesData, nil
}

func (c *Cluster) interfaceToOpenstackNetworkNode(
	openstackNetworkNodes interface{}) ([]*models.OpenstackNetworkNode, error) {
	var openstackNetworkNodesData []*models.OpenstackNetworkNode
	for _, openstackNetworkNode := range openstackNetworkNodes.([]interface{}) {
		openstackNetworkNodeInfo := models.InterfaceToOpenstackNetworkNode(openstackNetworkNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		openstackNetworkNodeData, err := c.getResource(defaultOpenstackNetworkNodeResPath, openstackNetworkNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of openstack_network node")
		}
		openstackNetworkNodeInfo = models.InterfaceToOpenstackNetworkNode(openstackNetworkNodeData)
		openstackNetworkNodesData = append(
			openstackNetworkNodesData, openstackNetworkNodeInfo)
	}
	return openstackNetworkNodesData, nil
}

func (c *Cluster) interfaceToOpenstackStorageNode(
	openstackStorageNodes interface{}) ([]*models.OpenstackStorageNode, error) {
	var openstackStorageNodesData []*models.OpenstackStorageNode
	for _, openstackStorageNode := range openstackStorageNodes.([]interface{}) {
		openstackStorageNodeInfo := models.InterfaceToOpenstackStorageNode(
			openstackStorageNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		openstackStorageNodeData, err := c.getResource(
			defaultOpenstackStorageNodeResPath, openstackStorageNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of openstack storage node")
		}
		openstackStorageNodeInfo = models.InterfaceToOpenstackStorageNode(openstackStorageNodeData)
		openstackStorageNodesData = append(
			openstackStorageNodesData, openstackStorageNodeInfo)
	}
	return openstackStorageNodesData, nil
}

func (c *Cluster) interfaceToOpenstackComputeNode(
	openstackComputeNodes interface{}) ([]*models.OpenstackComputeNode, error) {
	var openstackComputeNodesData []*models.OpenstackComputeNode
	for _, openstackComputeNode := range openstackComputeNodes.([]interface{}) {
		openstackComputeNodeInfo := models.InterfaceToOpenstackComputeNode(
			openstackComputeNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		openstackComputeNodeData, err := c.getResource(
			defaultOpenstackComputeNodeResPath, openstackComputeNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of openstack_compute")
		}
		openstackComputeNodeInfo = models.InterfaceToOpenstackComputeNode(openstackComputeNodeData)
		openstackComputeNodesData = append(openstackComputeNodesData, openstackComputeNodeInfo)
	}
	return openstackComputeNodesData, nil
}

func (c *Cluster) interfaceToVrouterNode(vrouterNodes interface{}) ([]*models.ContrailVrouterNode, error) {
	var vrouterNodesData []*models.ContrailVrouterNode
	for _, vrouterNode := range vrouterNodes.([]interface{}) {
		vrouterNodeInfo := models.InterfaceToContrailVrouterNode(vrouterNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		vrouterNodeData, err := c.getResource(defaultVrouterNodeResPath, vrouterNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of vrouter node")
		}
		vrouterNodeInfo = models.InterfaceToContrailVrouterNode(vrouterNodeData)
		vrouterNodesData = append(vrouterNodesData, vrouterNodeInfo)
	}
	return vrouterNodesData, nil
}

func (c *Cluster) interfaceToServiceNode(serviceNodes interface{}) ([]*models.ContrailServiceNode, error) {
	var serviceNodesData []*models.ContrailServiceNode
	for _, serviceNode := range serviceNodes.([]interface{}) {
		serviceNodeInfo := models.InterfaceToContrailServiceNode(serviceNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		serviceNodeData, err := c.getResource(defaultServiceNodeResPath, serviceNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of service node")
		}
		serviceNodeInfo = models.InterfaceToContrailServiceNode(serviceNodeData)
		serviceNodesData = append(serviceNodesData, serviceNodeInfo)
	}
	return serviceNodesData, nil
}

func (c *Cluster) interfaceToAnalyticsDBNode(
	analyticsDBNodes interface{}) ([]*models.ContrailAnalyticsDatabaseNode, error) {
	var analyticsDBNodesData []*models.ContrailAnalyticsDatabaseNode
	for _, analyticsDBNode := range analyticsDBNodes.([]interface{}) {
		analyticsDBNodeInfo := models.InterfaceToContrailAnalyticsDatabaseNode(analyticsDBNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		analyticsDBNodeData, err := c.getResource(defaultAnalyticsDBNodeResPath, analyticsDBNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of analytics db node")
		}
		analyticsDBNodeInfo = models.InterfaceToContrailAnalyticsDatabaseNode(analyticsDBNodeData)
		analyticsDBNodesData = append(analyticsDBNodesData, analyticsDBNodeInfo)
	}
	return analyticsDBNodesData, nil
}
func (c *Cluster) interfaceToAnalyticsNode(analyticsNodes interface{}) ([]*models.ContrailAnalyticsNode, error) {
	var analyticsNodesData []*models.ContrailAnalyticsNode
	for _, analyticsNode := range analyticsNodes.([]interface{}) {
		analyticsNodeInfo := models.InterfaceToContrailAnalyticsNode(analyticsNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		analyticsNodeData, err := c.getResource(defaultAnalyticsNodeResPath, analyticsNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of analytics node")
		}
		analyticsNodeInfo = models.InterfaceToContrailAnalyticsNode(analyticsNodeData)
		analyticsNodesData = append(analyticsNodesData, analyticsNodeInfo)
	}
	return analyticsNodesData, nil
}

func (c *Cluster) interfaceToWebuiNode(webuiNodes interface{}) ([]*models.ContrailWebuiNode, error) {
	var webuiNodesData []*models.ContrailWebuiNode
	for _, webuiNode := range webuiNodes.([]interface{}) {
		webuiNodeInfo := models.InterfaceToContrailWebuiNode(webuiNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		webuiNodeData, err := c.getResource(defaultWebuiNodeResPath, webuiNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of webui node")
		}
		webuiNodeInfo = models.InterfaceToContrailWebuiNode(webuiNodeData)
		webuiNodesData = append(webuiNodesData, webuiNodeInfo)
	}
	return webuiNodesData, nil
}

func (c *Cluster) interfaceToControlNode(controlNodes interface{}) ([]*models.ContrailControlNode, error) {
	var controlNodesData []*models.ContrailControlNode
	for _, controlNode := range controlNodes.([]interface{}) {
		controlNodeInfo := models.InterfaceToContrailControlNode(controlNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		controlNodeData, err := c.getResource(defaultControlNodeResPath, controlNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of control node")
		}
		controlNodeInfo = models.InterfaceToContrailControlNode(controlNodeData)
		controlNodesData = append(controlNodesData, controlNodeInfo)
	}
	return controlNodesData, nil
}

func (c *Cluster) interfaceToConfigDBNode(configDBNodes interface{}) ([]*models.ContrailConfigDatabaseNode, error) {
	var configDBNodesData []*models.ContrailConfigDatabaseNode
	for _, configDBNode := range configDBNodes.([]interface{}) {
		configDBNodeInfo := models.InterfaceToContrailConfigDatabaseNode(configDBNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		configDBNodeData, err := c.getResource(defaultConfigDBNodeResPath, configDBNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of config db node")
		}
		configDBNodeInfo = models.InterfaceToContrailConfigDatabaseNode(configDBNodeData)
		configDBNodesData = append(configDBNodesData, configDBNodeInfo)
	}
	return configDBNodesData, nil
}

func (c *Cluster) interfaceToConfigNode(configNodes interface{}) ([]*models.ContrailConfigNode, error) {
	var configNodesData []*models.ContrailConfigNode
	for _, configNode := range configNodes.([]interface{}) {
		configNodeInfo := models.InterfaceToContrailConfigNode(configNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		configNodeData, err := c.getResource(defaultConfigNodeResPath, configNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of config node")
		}
		configNodeInfo = models.InterfaceToContrailConfigNode(configNodeData)
		configNodesData = append(configNodesData, configNodeInfo)
	}
	return configNodesData, nil
}

func (c *Cluster) getNode(nodeID string, m map[string]bool) (*models.Node, error) {
	if _, ok := m[nodeID]; !ok {
		m[nodeID] = true
		n, err := c.getResource(defaultNodeResPath, nodeID)
		if err != nil {
			return nil, err
		}
		ni := models.InterfaceToNode(n)
		return ni, nil
	}
	return nil, nil
}

// nolint: gocyclo
func (c *Cluster) getNodeDetails(clusterInfo *models.ContrailCluster) ([]*models.Node, error) {
	var nodesInfo []*models.Node
	m := make(map[string]bool)
	for _, node := range clusterInfo.ContrailConfigNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.ContrailConfigDatabaseNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.ContrailControlNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.ContrailWebuiNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.ContrailAnalyticsNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.ContrailAnalyticsDatabaseNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.ContrailVrouterNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.ContrailServiceNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	return nodesInfo, nil
}

func (c *Cluster) getK8sNodeDetails(clusterInfo *models.KubernetesCluster) ([]*models.Node, error) {
	var nodesInfo []*models.Node
	m := make(map[string]bool)
	for _, node := range clusterInfo.KubernetesMasterNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.KubernetesNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	return nodesInfo, nil
}

// nolint: gocyclo
func (c *Cluster) getOpenstackNodeDetails(clusterInfo *models.OpenstackCluster) ([]*models.Node, error) {
	var nodesInfo []*models.Node
	m := make(map[string]bool)
	for _, node := range clusterInfo.OpenstackControlNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.OpenstackNetworkNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.OpenstackStorageNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.OpenstackMonitoringNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	for _, node := range clusterInfo.OpenstackComputeNodes {
		for _, nodeRef := range node.NodeRefs {
			n, err := c.getNode(nodeRef.UUID, m)
			if err != nil {
				return nil, err
			}
			if n != nil {
				nodesInfo = append(nodesInfo, n)
			}
		}
	}
	return nodesInfo, nil
}

// nolint: gocyclo
func (c *Cluster) getClusterDetails(clusterID string) (*Data, error) {
	rData, err := c.getResource(defaultResourcePath, clusterID)
	if err != nil {
		return nil, errors.New("unable to gather cluster information")
	}
	clusterInfo := models.InterfaceToContrailCluster(rData)

	// Expand config node back ref
	if configNodes, ok := rData["contrail_config_nodes"]; ok {
		clusterInfo.ContrailConfigNodes, err = c.interfaceToConfigNode(configNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand config database node back ref
	if configDBNodes, ok := rData["contrail_config_database_nodes"]; ok {
		clusterInfo.ContrailConfigDatabaseNodes, err = c.interfaceToConfigDBNode(configDBNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand control node back ref
	if controlNodes, ok := rData["contrail_control_nodes"]; ok {
		clusterInfo.ContrailControlNodes, err = c.interfaceToControlNode(controlNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand webui node back ref
	if webuiNodes, ok := rData["contrail_webui_nodes"]; ok {
		clusterInfo.ContrailWebuiNodes, err = c.interfaceToWebuiNode(webuiNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand analytics node back ref
	if analyticsNodes, ok := rData["contrail_analytics_nodes"]; ok {
		clusterInfo.ContrailAnalyticsNodes, err = c.interfaceToAnalyticsNode(analyticsNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand analytics database node back ref
	if analyticsDBNodes, ok := rData["contrail_analytics_database_nodes"]; ok {
		clusterInfo.ContrailAnalyticsDatabaseNodes, err = c.interfaceToAnalyticsDBNode(analyticsDBNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand vouter node back ref
	if vrouterNodes, ok := rData["contrail_vrouter_nodes"]; ok {
		clusterInfo.ContrailVrouterNodes, err = c.interfaceToVrouterNode(vrouterNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand csn node back ref
	if csnNodes, ok := rData["contrail_service_nodes"]; ok {
		clusterInfo.ContrailServiceNodes, err = c.interfaceToServiceNode(csnNodes)
		if err != nil {
			return nil, err
		}
	}
	// get all nodes information
	nodesInfo, err := c.getNodeDetails(clusterInfo)
	if err != nil {
		return nil, err
	}
	clusterData := &Data{
		clusterInfo: clusterInfo,
		nodesInfo:   nodesInfo,
	}

	// get all referred openstack cluster information
	var openstackClusterData []*OpenstackData
	for _, openstackClusterRef := range clusterInfo.OpenstackClusterRefs {
		openstackData, err := c.getOpenstackClusterDetails(openstackClusterRef.UUID)
		if err != nil {
			return nil, err
		}
		openstackClusterData = append(openstackClusterData, openstackData)
	}
	clusterData.openstackClusterData = openstackClusterData

	// get all referred kubernetes cluster information
	var kubernetesClusterData []*KubernetesData
	for _, kubernetesClusterRef := range clusterInfo.KubernetesClusterRefs {
		k8sData, err := c.getKubernetesClusterDetails(kubernetesClusterRef.UUID)
		if err != nil {
			return nil, err
		}
		kubernetesClusterData = append(kubernetesClusterData, k8sData)
	}
	clusterData.kubernetesClusterData = kubernetesClusterData

	return clusterData, nil
}

// nolint: gocyclo
func (c *Cluster) getOpenstackClusterDetails(clusterID string) (*OpenstackData, error) {
	rData, err := c.getResource(defaultOpenstackResourcePath, clusterID)
	if err != nil {
		return nil, errors.New("unable to gather openstack cluster information")
	}
	clusterInfo := models.InterfaceToOpenstackCluster(rData)

	// Expand openstack_compute back ref
	if openstackComputeNodes, ok := rData["openstack_compute_nodes"]; ok {
		clusterInfo.OpenstackComputeNodes, err = c.interfaceToOpenstackComputeNode(openstackComputeNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand openstack_storage node back ref
	if openstackStorageNodes, ok := rData["openstack_storage_nodes"]; ok {
		clusterInfo.OpenstackStorageNodes, err = c.interfaceToOpenstackStorageNode(openstackStorageNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand openstack_network node back ref
	if openstackNetworkNodes, ok := rData["openstack_network_nodes"]; ok {
		clusterInfo.OpenstackNetworkNodes, err = c.interfaceToOpenstackNetworkNode(openstackNetworkNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand openstack_monitoring node back ref
	if openstackMonitoringNodes, ok := rData["openstack_monitoring_nodes"]; ok {
		clusterInfo.OpenstackMonitoringNodes, err = c.interfaceToOpenstackMonitoringNode(openstackMonitoringNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand openstack_control node back ref
	if openstackControlNodes, ok := rData["openstack_control_nodes"]; ok {
		clusterInfo.OpenstackControlNodes, err = c.interfaceToOpenstackControlNode(openstackControlNodes)
		if err != nil {
			return nil, err
		}
	}
	// get all nodes information
	nodesInfo, err := c.getOpenstackNodeDetails(clusterInfo)
	if err != nil {
		return nil, err
	}
	clusterData := &OpenstackData{
		clusterInfo: clusterInfo,
		nodesInfo:   nodesInfo,
	}

	return clusterData, nil
}

func (c *Cluster) getKubernetesClusterDetails(clusterID string) (*KubernetesData, error) {
	rData, err := c.getResource(defaultK8sResourcePath, clusterID)
	if err != nil {
		return nil, errors.New("unable to gather k8s cluster information")
	}
	clusterInfo := models.InterfaceToKubernetesCluster(rData)

	// Expand kubernetes_master back ref
	if kubernetesMasterNodes, ok := rData["kubernetes_master_nodes"]; ok {
		clusterInfo.KubernetesMasterNodes, err = c.interfaceToKubernetesMasterNode(kubernetesMasterNodes)
		if err != nil {
			return nil, err
		}
	}
	// Expand kubernetes node back ref
	if kubernetesNodes, ok := rData["kubernetes_nodes"]; ok {
		clusterInfo.KubernetesNodes, err = c.interfaceToKubernetesNode(kubernetesNodes)
		if err != nil {
			return nil, err
		}
	}

	// get all nodes information
	nodesInfo, err := c.getK8sNodeDetails(clusterInfo)
	if err != nil {
		return nil, err
	}
	clusterData := &KubernetesData{
		clusterInfo: clusterInfo,
		nodesInfo:   nodesInfo,
	}

	return clusterData, nil
}
