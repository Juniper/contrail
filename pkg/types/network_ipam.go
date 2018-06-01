package types

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"golang.org/x/net/context"
	"net"
)

func isSubnetOverlap(subnet1, subnet2 *net.IPNet) bool {
	return subnet1.Contains(subnet2.IP) || subnet2.Contains(subnet1.IP)
}

func appendSubnetIfNoOverlap(subnets []*net.IPNet, subnet *net.IPNet) ([]*net.IPNet, error) {
	for _, existingSubnet := range subnets {
		if isSubnetOverlap(subnet, existingSubnet) {
			return nil, common.ErrorBadRequest(
				"Overlapping addresses: " + subnet.String() + "," + existingSubnet.String())
		}
	}
	return append(subnets, subnet), nil
}

func mergeSubnetIfNoOverlap(subnets []*net.IPNet, ipamSubnets []*models.IpamSubnetType) ([]*net.IPNet, error) {
	for _, subnet := range ipamSubnets {
		n, err := subnet.Subnet.Net()
		if err != nil {
			return nil, err
		}
		subnets, err = appendSubnetIfNoOverlap(subnets, n)
		if err != nil {
			return nil, err
		}
	}
	return subnets, nil
}

// CreateNetworkIpam do pre check for network ipam
func (service *ContrailTypeLogicService) CreateNetworkIpam(
	ctx context.Context,
	request *services.CreateNetworkIpamRequest) (response *services.CreateNetworkIpamResponse, err error) {

	networkIpam := request.GetNetworkIpam()
	err = db.DoInTransaction(
		ctx,
		service.DB.DB(),
		func(ctx context.Context) error {
			ipamSubnets := networkIpam.GetIpamSubnets()
			if ipamSubnets == nil {
				response, err = service.Next().CreateNetworkIpam(ctx, request)
				return err
			}

			if !networkIpam.IsFlatSubnet() {
				return common.ErrorBadRequest("Ipam subnets are allowed only with flat-subnet")
			}

			subnets := []*net.IPNet{}
			subnets, err = mergeSubnetIfNoOverlap(subnets, ipamSubnets.GetSubnets())
			if err != nil {
				return err
			}

			for _, ipamSubnet := range ipamSubnets.GetSubnets() {
				err = ipamSubnet.CheckIfSubnetParamsAreValid()
				if err != nil {
					return err
				}
			}

			ipamUUID := networkIpam.GetUUID()

			for _, ipamSubnet := range ipamSubnets.GetSubnets() {
				subnetUUID, err := service.createIpamSubnet(ctx, ipamSubnet, ipamUUID)
				if err != nil {
					return err
				}
				ipamSubnet.SubnetUUID = subnetUUID
			}

			response, err = service.Next().CreateNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

// DeleteNetworkIpam do pre check for network ipam deletion
func (service *ContrailTypeLogicService) DeleteNetworkIpam(
	ctx context.Context,
	request *services.DeleteNetworkIpamRequest) (response *services.DeleteNetworkIpamResponse, err error) {

	id := request.GetID()
	err = db.DoInTransaction(
		ctx,
		service.DB.DB(),
		func(ctx context.Context) error {
			var networkIpam *models.NetworkIpam
			networkIpam, err = service.getNetworkIpam(ctx, id)
			if err != nil {
				return err
			}

			ipamSubnets := networkIpam.GetIpamSubnets()
			if networkIpam == nil || !networkIpam.IsFlatSubnet() || ipamSubnets == nil {
				response, err = service.Next().DeleteNetworkIpam(ctx, request)
				return err
			}

			for _, ipamSubnet := range ipamSubnets.GetSubnets() {
				err := service.deleteIpamSubnet(ctx, ipamSubnet.GetSubnetUUID())
				if err != nil {
					return err
				}
			}

			response, err = service.Next().DeleteNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

// UpdateNetworkIpam do pre check for network ipam update
func (service *ContrailTypeLogicService) UpdateNetworkIpam(
	ctx context.Context,
	request *services.UpdateNetworkIpamRequest) (response *services.UpdateNetworkIpamResponse, err error) {

	//networkIpam := request.GetNetworkIpam()
	err = db.DoInTransaction(
		ctx,
		service.DB.DB(),
		func(ctx context.Context) error {
			//TODO ipam_mgmt check
			//TODO check new subnet method
			//TODO check flat-subnet
			//TODO rest of update checks
			response, err = service.Next().UpdateNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

func (service *ContrailTypeLogicService) getNetworkIpam(
	ctx context.Context,
	id string) (*models.NetworkIpam, error) {

	networkIpamRes, err := service.DB.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{ID: id})
	if err != nil {
		return nil, err
	}
	return networkIpamRes.GetNetworkIpam(), err
}

func (service *ContrailTypeLogicService) createIpamSubnet(
	ctx context.Context,
	ipamSubnet *models.IpamSubnetType,
	ipamUUID string) (string, error) {
	createIpamSubnetParams := &ipam.CreateIpamSubnetRequest{
		IpamSubnet:      ipamSubnet,
		NetworkIpamUUID: ipamUUID,
	}
	return service.AddressManager.CreateIpamSubnet(ctx, createIpamSubnetParams)
}

func (service *ContrailTypeLogicService) deleteIpamSubnet(
	ctx context.Context,
	subnetUUID string) error {
	deleteIpamSubnetParams := &ipam.DeleteIpamSubnetRequest{
		SubnetUUID: subnetUUID,
	}
	return service.AddressManager.DeleteIpamSubnet(ctx, deleteIpamSubnetParams)
}
