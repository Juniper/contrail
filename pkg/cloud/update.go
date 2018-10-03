package cloud

import (
	"context"
	"fmt"

	tf "github.com/hashicorp/terraform/terraform"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func (c *Cloud) updateIPDetails(ctx context.Context, data *Data) error {
	tfState, err := readStateFile(GetTFStateFile(c.config.CloudID))
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

func updateInstanceIP(ctx context.Context, instance *instanceData, tfState *tf.State) error {

	gwExists := gwRoleExists(instance)

	privateIP, err := getIPFromTFState(tfState,
		fmt.Sprintf("private_ip.%s", instance.info.Name))

	if gwExists {
		portObj, err := createPort(ctx, "private", privateIP, instance.client)
		err = addPortToNode(ctx, portObj, instance.info, instance.client)
		if err != nil {
			return err
		}
		return nil

		publicIP, err := getIPFromTFState(tfState, fmt.Sprintf("public_ip.%s", instance.info.Name))
		err = addIPToNode(ctx, publicIP, instance.info, instance.client)
		if err != nil {
			return err
		}
	}

	err = addIPToNode(ctx, privateIP, instance.info, instance.client)
	if err != nil {
		return err
	}
	return nil
}

func gwRoleExists(instance *instanceData) bool {
	for _, role := range instance.info.CloudInfo.Roles {
		if role == "gateway" {
			return true
		}
	}
	return false
}

func createPort(ctx context.Context, portName string, ip string, client *client.HTTP) (*models.Port, error) {

	request := new(services.CreatePortRequest)
	request.Port.Name = portName
	request.Port.IPAddress = ip

	portResp, err := client.CreatePort(ctx, request)
	if err != nil {
		return nil, err
	}
	return portResp.GetPort(), nil
}

func addPortToNode(ctx context.Context, port *models.Port, instance *models.Node, client *client.HTTP) error {

	request := new(services.UpdateNodeRequest)
	request.Node = instance
	request.Node.AddPort(port)

	_, err := client.UpdateNode(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func addIPToNode(ctx context.Context, ip string, instance *models.Node, client *client.HTTP) error {

	request := new(services.UpdateNodeRequest)
	request.Node = instance
	request.Node.IPAddress = ip

	_, err := client.UpdateNode(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
