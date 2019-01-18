package cloud

import (
	"fmt"

	tf "github.com/hashicorp/terraform/terraform"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func (c *Cloud) updateIPDetails(data *Data) error {
	tfState, err := readStateFile(GetTFStateFile(c.config.CloudID))
	if err != nil {
		return err
	}
	for _, instance := range data.instances {
		err := updateInstanceIP(instance, tfState)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateInstanceIP(instance *instanceData, tfState *tf.State) error {

	privateIP, err := getIPFromTFState(tfState,
		fmt.Sprintf("%s.private_ip", instance.info.Hostname))
	if err != nil {
		return err
	}

	if gwRoleExists(instance) {
		portObj, inErr := createPort("private", privateIP,
			instance.info, instance.client)
		if inErr != nil {
			return inErr
		}

		inErr = addPortToNode(portObj, instance.info, instance.client)
		if inErr != nil {
			return inErr
		}

		publicIP, inErr := getIPFromTFState(tfState,
			fmt.Sprintf("%s.public_ip", instance.info.Hostname))
		if inErr != nil {
			return inErr
		}
		return addIPToNode(publicIP, instance.info, instance.client)
	}

	return addIPToNode(privateIP, instance.info, instance.client)
}

func gwRoleExists(instance *instanceData) bool {
	for _, role := range instance.info.CloudInfo.Roles {
		if role == "gateway" {
			return true
		}
	}
	return false
}

func createPort(portName string, ip string,
	instance *models.Node, client *client.HTTP) (*models.Port, error) {

	if len(instance.Ports) != 0 {
		for _, p := range instance.Ports {
			if p.Name == portName && p.IPAddress != ip {
				request := new(services.UpdatePortRequest)
				request.Port = p
				request.Port.IPAddress = ip

				response := new(services.UpdatePortResponse)
				_, err := client.Update("/port/"+request.Port.UUID, request, response)
				if err != nil {
					return nil, err
				}
				return response.GetPort(), err
			} else if p.Name == portName && p.IPAddress == ip {
				return p, nil
			}
		}
	}

	port := new(models.Port)
	port.Name = portName
	port.ParentType = "node"
	port.ParentUUID = instance.UUID
	port.IPAddress = ip

	request := new(services.CreatePortRequest)
	request.Port = port

	response := new(services.CreatePortResponse)

	_, err := client.Create("/ports", request, response)
	if err != nil {
		return nil, err
	}
	return response.GetPort(), nil
}

func addPortToNode(port *models.Port,
	instance *models.Node, client *client.HTTP) error {

	request := new(services.UpdateNodeRequest)
	request.Node = instance
	request.Node.Ports = append(request.Node.Ports, port)

	_, err := client.Update("/node/"+request.Node.UUID,
		request, new(services.UpdateNodeResponse))
	return err

}

func addIPToNode(ip string,
	instance *models.Node, client *client.HTTP) error {

	request := new(services.UpdateNodeRequest)
	request.Node = instance
	request.Node.IPAddress = ip

	_, err := client.Update("/node/"+request.Node.UUID,
		request, new(services.UpdateNodeResponse))
	return err
}
