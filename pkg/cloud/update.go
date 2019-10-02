package cloud

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

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

func updateInstanceIP(ctx context.Context, instance *instanceData, tfState *tf.State) error {
	if gwRoleExists(instance) {
		publicIPs := []string{}
		for i := int64(1); i <= instance.info.CloudInfo.Count; i++ {
			publicIP, err := getIPFromTFState(tfState, fmt.Sprintf("%s%v.public_ip", instance.info.Hostname, i))
			if err != nil {
				return err
			}
			publicIPs = append(publicIPs, publicIP)
		}

		if err := setIPAddresses(ctx, publicIPs, "public", instance.info, instance.client); err != nil {
			return err
		}
	}

	privateIPs := []string{}
	for i := int64(1); i <= instance.info.CloudInfo.Count; i++ {
		privateIP, err := getIPFromTFState(tfState, fmt.Sprintf("%s%v.private_ip", instance.info.Hostname, i))
		if err != nil {
			return err
		}
		privateIPs = append(privateIPs, privateIP)
	}

	return setIPAddresses(ctx, privateIPs, "private", instance.info, instance.client)
}

func gwRoleExists(instance *instanceData) bool {
	for _, role := range instance.info.CloudInfo.Roles {
		if role == "gateway" {
			return true
		}
	}
	return false
}

func setIPAddresses(ctx context.Context, ips []string, ipType string, instance *models.Node, cli *client.HTTP) error {
	request := &services.UpdateNodeRequest{Node: instance}
	switch ipType {
	case "public":
		request.Node.CloudInfo.PublicIPAddresses = ips
	case "private":
		request.Node.CloudInfo.PrivIPAddresses = ips
	default:
		return errors.Errorf("unknown type of ip address: %v", ipType)
	}

	_, err := cli.UpdateNode(ctx, request)
	return err
}
