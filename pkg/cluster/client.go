package cluster

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
)

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

func (c *Cluster) interfaceToVrouterNode(vrouterNodes interface{}) ([]*models.ContrailVrouterNode, error) {
	var vrouterNodesData []*models.ContrailVrouterNode
	for _, vrouterNode := range vrouterNodes.([]interface{}) {
		vrouterNodeInfo := models.InterfaceToContrailVrouterNode(vrouterNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		vrouterNodeData, err := c.getResource(defaultVrouterNodeResPath, vrouterNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of config node")
		}
		vrouterNodeInfo = models.InterfaceToContrailVrouterNode(vrouterNodeData)
		// Expand node refs
		if nodeRefs, ok := vrouterNodeData["node_refs"]; ok {
			var nodesData []*models.ContrailVrouterNodeNodeRef
			for _, nodeRef := range nodeRefs.([]interface{}) {
				nodeRefMap, ok := nodeRef.(map[string]interface{})
				if !ok {
					return nil, errors.New("unable to get node refs")
				}
				nodeInfo := &models.ContrailVrouterNodeNodeRef{}
				if uuid, ok := nodeRefMap["uuid"]; ok {
					nodeInfo.UUID = common.InterfaceToString(uuid)
				}
				if to, ok := nodeRefMap["to"]; ok {
					nodeInfo.To = common.InterfaceToStringList(to)
				}
				nodesData = append(nodesData, nodeInfo)
			}
			vrouterNodeInfo.NodeRefs = nodesData
		}
		vrouterNodesData = append(vrouterNodesData, vrouterNodeInfo)
	}
	return vrouterNodesData, nil
}
func (c *Cluster) interfaceToAnalyticsDBNode(
	analyticsDBNodes interface{}) ([]*models.ContrailAnalyticsDatabaseNode, error) {
	var analyticsDBNodesData []*models.ContrailAnalyticsDatabaseNode
	for _, analyticsDBNode := range analyticsDBNodes.([]interface{}) {
		analyticsDBNodeInfo := models.InterfaceToContrailAnalyticsDatabaseNode(analyticsDBNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		analyticsDBNodeData, err := c.getResource(defaultAnalyticsDBNodeResPath, analyticsDBNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of config node")
		}
		analyticsDBNodeInfo = models.InterfaceToContrailAnalyticsDatabaseNode(analyticsDBNodeData)
		// Expand node refs
		if nodeRefs, ok := analyticsDBNodeData["node_refs"]; ok {
			var nodesData []*models.ContrailAnalyticsDatabaseNodeNodeRef
			for _, nodeRef := range nodeRefs.([]interface{}) {
				nodeRefMap, ok := nodeRef.(map[string]interface{})
				if !ok {
					return nil, errors.New("unable to get node refs")
				}
				nodeInfo := &models.ContrailAnalyticsDatabaseNodeNodeRef{}
				if uuid, ok := nodeRefMap["uuid"]; ok {
					nodeInfo.UUID = common.InterfaceToString(uuid)
				}
				if to, ok := nodeRefMap["to"]; ok {
					nodeInfo.To = common.InterfaceToStringList(to)
				}
				nodesData = append(nodesData, nodeInfo)
			}
			analyticsDBNodeInfo.NodeRefs = nodesData
		}
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
			return nil, errors.New("unable to get information of config node")
		}
		analyticsNodeInfo = models.InterfaceToContrailAnalyticsNode(analyticsNodeData)
		// Expand node refs
		if nodeRefs, ok := analyticsNodeData["node_refs"]; ok {
			var nodesData []*models.ContrailAnalyticsNodeNodeRef
			for _, nodeRef := range nodeRefs.([]interface{}) {
				nodeRefMap, ok := nodeRef.(map[string]interface{})
				if !ok {
					return nil, errors.New("unable to get node refs")
				}
				nodeInfo := &models.ContrailAnalyticsNodeNodeRef{}
				if uuid, ok := nodeRefMap["uuid"]; ok {
					nodeInfo.UUID = common.InterfaceToString(uuid)
				}
				if to, ok := nodeRefMap["to"]; ok {
					nodeInfo.To = common.InterfaceToStringList(to)
				}
				nodesData = append(nodesData, nodeInfo)
			}
			analyticsNodeInfo.NodeRefs = nodesData
		}
		analyticsNodesData = append(analyticsNodesData, analyticsNodeInfo)
	}
	return analyticsNodesData, nil
}

/*
func (c *Cluster) interfaceToWebuiNode(webuiNodes interface{}) ([]*models.ContrailWebuiNode, error) {
	var webuiNodesData []*models.ContrailWebuiNode
	for _, webuiNode := range webuiNodes.([]interface{}) {
		webuiNodeInfo := models.InterfaceToContrailWebuiNode(webuiNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		webuiNodeData, err := c.getResource(defaultWebuiNodeResPath, webuiNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of config node")
		}
		webuiNodeInfo = models.InterfaceToContrailWebuiNode(webuiNodeData)
		// Expand node refs
		if nodeRefs, ok := webuiNodeData["node_refs"]; ok {
			var nodesData []*models.ContrailWebuiNodeNodeRef
			for _, nodeRef := range nodeRefs.([]interface{}) {
				nodeRefMap, ok := nodeRef.(map[string]interface{})
				if !ok {
					return nil, errors.New("unable to get node refs")
				}
				nodeInfo := &models.ContrailWebuiNodeNodeRef{}
				if uuid, ok := nodeRefMap["uuid"]; ok {
					nodeInfo.UUID = schema.InterfaceToString(uuid)
				}
				if to, ok := nodeRefMap["to"]; ok {
					nodeInfo.To = schema.InterfaceToStringList(to)
				}
				nodesData = append(nodesData, nodeInfo)
			}
			webuiNodeInfo.NodeRefs = nodesData
		}
		webuiNodesData = append(webuiNodesData, webuiNodeInfo)
	}
	return webuiNodesData, nil
}
*/

func (c *Cluster) interfaceToControlNode(controlNodes interface{}) ([]*models.ContrailControlNode, error) {
	var controlNodesData []*models.ContrailControlNode
	for _, controlNode := range controlNodes.([]interface{}) {
		controlNodeInfo := models.InterfaceToContrailControlNode(controlNode.(map[string]interface{}))
		// Read contrail role node to get the node refs information
		controlNodeData, err := c.getResource(defaultControlNodeResPath, controlNodeInfo.UUID)
		if err != nil {
			return nil, errors.New("unable to get information of config node")
		}
		controlNodeInfo = models.InterfaceToContrailControlNode(controlNodeData)
		// Expand node refs
		if nodeRefs, ok := controlNodeData["node_refs"]; ok {
			var nodesData []*models.ContrailControlNodeNodeRef
			for _, nodeRef := range nodeRefs.([]interface{}) {
				nodeRefMap, ok := nodeRef.(map[string]interface{})
				if !ok {
					return nil, errors.New("unable to get node refs")
				}
				nodeInfo := &models.ContrailControlNodeNodeRef{}
				if uuid, ok := nodeRefMap["uuid"]; ok {
					nodeInfo.UUID = common.InterfaceToString(uuid)
				}
				if to, ok := nodeRefMap["to"]; ok {
					nodeInfo.To = common.InterfaceToStringList(to)
				}
				nodesData = append(nodesData, nodeInfo)
			}
			controlNodeInfo.NodeRefs = nodesData
		}
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
			return nil, errors.New("unable to get information of config node")
		}
		configDBNodeInfo = models.InterfaceToContrailConfigDatabaseNode(configDBNodeData)
		// Expand node refs
		if nodeRefs, ok := configDBNodeData["node_refs"]; ok {
			var nodesData []*models.ContrailConfigDatabaseNodeNodeRef
			for _, nodeRef := range nodeRefs.([]interface{}) {
				nodeRefMap, ok := nodeRef.(map[string]interface{})
				if !ok {
					return nil, errors.New("unable to get node refs")
				}
				nodeInfo := &models.ContrailConfigDatabaseNodeNodeRef{}
				if uuid, ok := nodeRefMap["uuid"]; ok {
					nodeInfo.UUID = common.InterfaceToString(uuid)
				}
				if to, ok := nodeRefMap["to"]; ok {
					nodeInfo.To = common.InterfaceToStringList(to)
				}
				nodesData = append(nodesData, nodeInfo)
			}
			configDBNodeInfo.NodeRefs = nodesData
		}
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
		// Expand node refs
		if nodeRefs, ok := configNodeData["node_refs"]; ok {
			var nodesData []*models.ContrailConfigNodeNodeRef
			for _, nodeRef := range nodeRefs.([]interface{}) {
				nodeRefMap, ok := nodeRef.(map[string]interface{})
				if !ok {
					return nil, errors.New("unable to get node refs")
				}
				nodeInfo := &models.ContrailConfigNodeNodeRef{}
				if uuid, ok := nodeRefMap["uuid"]; ok {
					nodeInfo.UUID = common.InterfaceToString(uuid)
				}
				if to, ok := nodeRefMap["to"]; ok {
					nodeInfo.To = common.InterfaceToStringList(to)
				}
				nodesData = append(nodesData, nodeInfo)
			}
			configNodeInfo.NodeRefs = nodesData
		}
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
	/*
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
	*/
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
	/*
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
	*/
	return nodesInfo, nil
}

func (c *Cluster) getClusterDetails(clusterID string) (*Data, error) {
	rData, err := c.getResource(defaultResourcePath, clusterID)
	if err != nil {
		return nil, errors.New("unable to gather cluster information")
	}
	clusterInfo := models.InterfaceToContrailCluster(rData)

	// Expand config node back ref
	if configNodes, ok := rData["contrail_config_nodes"]; ok {
		configNodesInfo, err := c.interfaceToConfigNode(configNodes)
		if err != nil {
			return nil, err
		}
		clusterInfo.ContrailConfigNodes = configNodesInfo
	}
	// Expand config database node back ref
	if configDBNodes, ok := rData["contrail_config_database_nodes"]; ok {
		configDBNodesInfo, err := c.interfaceToConfigDBNode(configDBNodes)
		if err != nil {
			return nil, err
		}
		clusterInfo.ContrailConfigDatabaseNodes = configDBNodesInfo
	}
	// Expand control node back ref
	if controlNodes, ok := rData["contrail_control_nodes"]; ok {
		controlNodesInfo, err := c.interfaceToControlNode(controlNodes)
		if err != nil {
			return nil, err
		}
		clusterInfo.ContrailControlNodes = controlNodesInfo
	}
	/*
		// Expand webui node back ref
		if webuiNodes, ok := rData["contrail_webui_nodes"]; ok {
			webuiNodesInfo, err := c.interfaceToWebuiNode(webuiNodes)
			if err != nil {
				return nil, err
			}
			clusterInfo.ContrailWebuiNodes = webuiNodesInfo
		}
	*/
	// Expand analytics node back ref
	if analyticsNodes, ok := rData["contrail_analytics_nodes"]; ok {
		analyticsNodesInfo, err := c.interfaceToAnalyticsNode(analyticsNodes)
		if err != nil {
			return nil, err
		}
		clusterInfo.ContrailAnalyticsNodes = analyticsNodesInfo
	}
	// Expand analytics database node back ref
	if analyticsDBNodes, ok := rData["contrail_analytics_database_nodes"]; ok {
		analyticsDBNodesInfo, err := c.interfaceToAnalyticsDBNode(analyticsDBNodes)
		if err != nil {
			return nil, err
		}
		clusterInfo.ContrailAnalyticsDatabaseNodes = analyticsDBNodesInfo
	}
	// Expand vouter node back ref
	if vrouterNodes, ok := rData["contrail_vrouter_nodes"]; ok {
		vrouterNodesInfo, err := c.interfaceToVrouterNode(vrouterNodes)
		if err != nil {
			return nil, err
		}
		clusterInfo.ContrailVrouterNodes = vrouterNodesInfo
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

	return clusterData, nil
}
