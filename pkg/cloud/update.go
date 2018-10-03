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
		fmt.Sprintf("private_ip.%s", instance.info.Name))
	if err != nil {
		return err
	}

	if gwRoleExists(instance) {
		portObj, inErr := createPort("private", privateIP, instance.client)
		if inErr != nil {
			return inErr
		}
		inErr = addPortToNode(portObj, instance.info, instance.client)
		if err != nil {
			return inErr
		}

		publicIP, inErr := getIPFromTFState(tfState,
			fmt.Sprintf("public_ip.%s", instance.info.Name))
		if inErr != nil {
			return inErr
		}
		inErr = addIPToNode(publicIP, instance.info, instance.client)
		if err != nil {
			return inErr
		}
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
	client *client.HTTP) (*models.Port, error) {

	request := new(services.CreatePortRequest)
	request.Port.Name = portName
	request.Port.IPAddress = ip

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
