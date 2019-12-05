package cloud

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func updateIPDetails(ctx context.Context, instances []*instanceData, tfState terraformState) error {
	for _, instance := range instances {
		if err := updateIP(ctx, instance, tfState); err != nil {
			return err
		}
	}
	return nil
}

func updateIP(ctx context.Context,
	instance *instanceData, tfState terraformState) error {
	if isGateway(instance) {
		return updateGatewayInstanceIP(ctx, instance, tfState)
	}
	return updateInstanceIP(ctx, instance, tfState)
}

func isGateway(instance *instanceData) bool {
	for _, role := range instance.info.CloudInfo.Roles {
		if role == "gateway" {
			return true
		}
	}
	return false
}

func updateInstanceIP(ctx context.Context,
	instance *instanceData, tfState terraformState) error {
	privateIP, err := tfState.GetPrivateIp(instance.info.Hostname)
	if err != nil {
		return err
	}
	return addIPToNode(ctx, privateIP, instance.info, instance.client)
}

func updateGatewayInstanceIP(ctx context.Context,
	instance *instanceData, tfState terraformState) error {

	privateIP, err := tfState.GetPrivateIp(instance.info.Hostname)
	if err != nil {
		return err
	}

	portObj, err := createPort(ctx, "private", privateIP,
		instance.info, instance.client)
	if err != nil {
		return err
	}

	if err = addPortToNode(ctx, portObj, instance.info, instance.client); err != nil {
		return err
	}

	publicIP, err := tfState.GetPublicIp(instance.info.Hostname)
	if err != nil {
		return err
	}
	return addIPToNode(ctx, publicIP, instance.info, instance.client)
}

func createPort(ctx context.Context, portName string, ip string,
	instance *models.Node, client services.Service) (*models.Port, error) {

	if len(instance.Ports) != 0 {
		for _, p := range instance.Ports {
			if p.Name == portName && p.IPAddress != ip {
				request := &services.UpdatePortRequest{Port: p}
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

	request := &services.CreatePortRequest{
		Port: &models.Port{
			Name:       portName,
			ParentType: "node",
			ParentUUID: instance.UUID,
			IPAddress:  ip,
		}}

	portResp, err := client.CreatePort(ctx, request)
	return portResp.GetPort(), err
}

func addPortToNode(ctx context.Context, port *models.Port,
	instance *models.Node, client services.Service) error {

	request := &services.UpdateNodeRequest{Node: instance}
	request.Node.AddPort(port)

	_, err := client.UpdateNode(ctx, request)
	return err
}

func addIPToNode(ctx context.Context, ip string,
	instance *models.Node, client services.Service) error {

	request := &services.UpdateNodeRequest{Node: instance}
	request.Node.IPAddress = ip

	_, err := client.UpdateNode(ctx, request)
	return err
}
