package undercloud

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// BaremetalServer is the ironic introspected server
type BaremetalServer struct {
	IMPMIAddress string        `json:"ipmi_address"`
	NumaTopology *NumaTopology `json:"numa_topology"`
}

// NumaTopology is the Nic topology of the server
type NumaTopology struct {
	Nics []*Nic `json:"nics"`
}

// Nic is the Nic in the server
type Nic struct {
	Name string `json:"name"`
}

func (c *contrailCloudDeployer) getPortsOfNode(
	nodeID string) (ports []string, err error) {

	request := &services.ListPortRequest{
		Spec: &baseservices.ListSpec{
			Fields: []string{"name"},
			Filters: []*baseservices.Filter{
				&baseservices.Filter{
					Key:    "parent_uuid",
					Values: []string{nodeID},
				},
			},
		},
	}
	resp, err := c.undercloud.APIServer.ListPort(context.Background(), request)
	if err != nil {
		c.Log.Errorf("unable to get ports with parent node: %s in inventory", nodeID)
		return nil, err
	}
	for _, p := range resp.GetPorts() {
		ports = append(ports, p.Name)
	}
	return ports, nil
}

func (c *contrailCloudDeployer) createPort(nic *Nic, nodeID string) error {
	port := &models.Port{
		Name:       nic.Name,
		ParentUUID: nodeID,
	}
	request := &services.CreatePortRequest{Port: port}
	_, err := c.undercloud.APIServer.CreatePort(context.Background(), request)
	return err
}

func (c *contrailCloudDeployer) updateNodeWithIntrospectData(
	UUID string, bms *BaremetalServer) error {
	node := &models.Node{
		UUID: UUID,
		Type: "private",
	}
	request := &services.UpdateNodeRequest{Node: node}
	if _, err := c.undercloud.APIServer.UpdateNode(context.Background(), request); err != nil {
		c.Log.Errorf("update of node: %s failed", UUID)
		return err
	}
	ports, err := c.getPortsOfNode(UUID)
	if err != nil {
		return err
	}
	portMap := map[string]bool{}
	for _, port := range ports {
		portMap[port] = true
	}
	for _, nic := range bms.NumaTopology.Nics {
		if _, ok := portMap[nic.Name]; !ok {
			if err = c.createPort(nic, UUID); err != nil {
				c.Log.Errorf("creating port: %s for parent node: %s failed", nic.Name, UUID)
				return err
			}
		}
	}
	return err
}

func (c *contrailCloudDeployer) getNodeByIPMIAddress(
	IMPMIAddress string) (nodes []*models.Node, err error) {

	request := &services.ListNodeRequest{
		Spec: &baseservices.ListSpec{
			Fields: []string{"uuid"},
			Filters: []*baseservices.Filter{
				&baseservices.Filter{
					Key:    "ipmi_address",
					Values: []string{IMPMIAddress},
				},
			},
		},
	}
	resp, err := c.undercloud.APIServer.ListNode(context.Background(), request)
	if err != nil {
		c.Log.Errorf("unable to find node with ipmi: %s in inventory", IMPMIAddress)
		return nil, err
	}
	return resp.GetNodes(), nil
}

func (c *contrailCloudDeployer) updateNodes(bms *BaremetalServer) error {

	if bms.NumaTopology == nil {
		return nil
	}
	nodes, err := c.getNodeByIPMIAddress(bms.IMPMIAddress)
	if err != nil {
		return err
	}
	for i, node := range nodes {
		if i > 0 {
			c.Log.Warningf("duplicate node: %s found with ipmi: %s",
				node.UUID, bms.IMPMIAddress)
		}
		if err = c.updateNodeWithIntrospectData(node.UUID, bms); err != nil {
			return err
		}
	}
	return nil
}

func (c *contrailCloudDeployer) introspectAndUpdateNodes() error {
	introspectDir := c.getIntrospectionDir()
	_, err := os.Stat(introspectDir)
	if err != nil {
		c.Log.Errorf("introspect dir: %s not found", introspectDir)
		return err
	}

	files, err := ioutil.ReadDir(introspectDir)
	if err != nil {
		return err
	}

	for _, aFile := range files {
		fileName := filepath.Join(introspectDir, aFile.Name())
		var introspectFile *os.File
		introspectFile, err = os.Open(fileName)
		if err != nil {
			c.Log.Errorf("unable to read introspect file: %s", fileName)
			return err
		}
		defer introspectFile.Close() //nolint: errcheck

		byteValue, err := ioutil.ReadAll(introspectFile)
		if err != nil {
			c.Log.Errorf("unable to read introspect file: %s", fileName)
			return err
		}
		var baremetalServer *BaremetalServer
		err = json.Unmarshal(byteValue, &baremetalServer)
		if err != nil {
			return err
		}
		err = c.updateNodes(baremetalServer)
		if err != nil {
			return err
		}
	}
	return nil
}
