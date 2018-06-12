package cluster

import (
	"github.com/Juniper/contrail/pkg/models"
)

// OpenstackData is the representation of openstack cluster details.
type OpenstackData struct {
	clusterInfo *models.OpenstackCluster
	nodesInfo   []*models.Node
}

// KubernetesData is the representation of kubernetes cluster details.
type KubernetesData struct {
	clusterInfo *models.KubernetesCluster
	nodesInfo   []*models.Node
}

// Data is the representation of cluster details.
type Data struct {
	clusterInfo           *models.ContrailCluster
	nodesInfo             []*models.Node
	openstackClusterData  []*OpenstackData
	kubernetesClusterData []*KubernetesData
	// TODO (ijohnson): Add gce/aws/kvm info
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
