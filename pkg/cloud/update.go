package cloud

import (
	"context"
	"fmt"

	tf "github.com/hashicorp/terraform/terraform"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func updateIPDetails(ctx context.Context, cloudID string, data *Data) error {
	tfState, err := readStateFile(GetTFStateFile(cloudID))
	if err != nil {
		return err
	}
	for _, instance := range data.instances {
		err := updateInstanceIP(ctx, instance, tfState)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateInstanceIP(ctx context.Context,
	instance *instanceData, tfState *tf.State) error {

	privateIP, err := getIPFromTFState(tfState,
		fmt.Sprintf("%s.private_ip", instance.info.Hostname))
	if err != nil {
		return err
	}

	if gwRoleExists(instance) {
		portObj, inErr := createPort(ctx, "private", privateIP,
			instance.info, instance.client)
		if inErr != nil {
			return inErr
		}

		inErr = addPortToNode(ctx, portObj, instance.info, instance.client)
		if err != nil {
			return inErr
		}

		publicIP, inErr := getIPFromTFState(tfState,
			fmt.Sprintf("%s.public_ip", instance.info.Hostname))
		if inErr != nil {
			return inErr
		}
		return addIPToNode(ctx, publicIP, instance.info, instance.client)
	}

	return addIPToNode(ctx, privateIP, instance.info, instance.client)
}

func gwRoleExists(instance *instanceData) bool {
	for _, role := range instance.info.CloudInfo.Roles {
		if role == "gateway" {
			return true
		}
	}
	return false
}

func createPort(ctx context.Context, portName string, ip string,
	instance *models.Node, client *client.HTTP) (*models.Port, error) {

	if len(instance.Ports) != 0 {
		for _, p := range instance.Ports {
			if p.Name == portName && p.IPAddress != ip {
				request := new(services.UpdatePortRequest)
				request.Port = p
				request.Port.IPAddress = ip
				portResp, err := client.UpdatePort(ctx, request)
				if err != nil {
					return nil, err
				}
				return portResp.GetPort(), err
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

	portResp, err := client.CreatePort(ctx, request)
	if err != nil {
		return nil, err
	}
	return portResp.GetPort(), err
}

func addPortToNode(ctx context.Context, port *models.Port,
	instance *models.Node, client *client.HTTP) error {

	request := new(services.UpdateNodeRequest)
	request.Node = instance
	request.Node.AddPort(port)
	_, err := client.UpdateNode(ctx, request)
	return err
}

func addIPToNode(ctx context.Context, ip string,
	instance *models.Node, client *client.HTTP) error {

	request := new(services.UpdateNodeRequest)
	request.Node = instance
	request.Node.IPAddress = ip

	_, err := client.UpdateNode(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
