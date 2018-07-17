package types

import (
	"context"
	"net"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

func isSubnetOverlap(subnet1, subnet2 *net.IPNet) bool {
	return subnet1.Contains(subnet2.IP) || subnet2.Contains(subnet1.IP)
}

func appendSubnetIfNoOverlap(subnets []*net.IPNet, subnet *net.IPNet) ([]*net.IPNet, error) {
	for _, existingSubnet := range subnets {
		if isSubnetOverlap(subnet, existingSubnet) {
			return nil, errors.Errorf(
				"overlapping addresses: %s, %s", subnet.String(), existingSubnet.String())
		}
	}
	return append(subnets, subnet), nil
}

func checkSubnetsOverlap(ipamSubnets []*models.IpamSubnetType) error {
	subnets := []*net.IPNet{}
	for _, subnet := range ipamSubnets {
		n, err := subnet.Subnet.Net()
		if err != nil {
			return err
		}
		subnets, err = appendSubnetIfNoOverlap(subnets, n)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateNetworkIpam do pre check for network ipam
func (sv *ContrailTypeLogicService) CreateNetworkIpam(
	ctx context.Context,
	request *services.CreateNetworkIpamRequest) (response *services.CreateNetworkIpamResponse, err error) {

	networkIpam := request.GetNetworkIpam()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			ipamSubnets := networkIpam.GetIpamSubnets()
			if ipamSubnets == nil {
				response, err = sv.BaseService.CreateNetworkIpam(ctx, request)
				return err
			}

			if !networkIpam.IsFlatSubnet() {
				return common.ErrorBadRequest("Ipam subnets are allowed only with flat-subnet")
			}

			err = checkSubnetsOverlap(ipamSubnets.GetSubnets())
			if err != nil {
				return common.ErrorBadRequest(err.Error())
			}

			for _, ipamSubnet := range ipamSubnets.GetSubnets() {
				err = ipamSubnet.CheckIfSubnetParamsAreValid()
				if err != nil {
					return err
				}
			}

			ipamUUID := networkIpam.GetUUID()
			for _, ipamSubnet := range ipamSubnets.GetSubnets() {
				subnetUUID, cErr := sv.createIpamSubnet(ctx, ipamSubnet, ipamUUID)
				if cErr != nil {
					return cErr
				}
				ipamSubnet.SubnetUUID = subnetUUID
			}

			response, err = sv.BaseService.CreateNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

// DeleteNetworkIpam do pre check for network ipam deletion
func (sv *ContrailTypeLogicService) DeleteNetworkIpam(
	ctx context.Context,
	request *services.DeleteNetworkIpamRequest) (response *services.DeleteNetworkIpamResponse, err error) {

	id := request.GetID()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var networkIpam *models.NetworkIpam
			networkIpam, err = sv.getNetworkIpam(ctx, id)
			if err != nil {
				return err
			}

			ipamSubnets := networkIpam.GetIpamSubnets()
			if networkIpam != nil && networkIpam.IsFlatSubnet() && ipamSubnets != nil {
				for _, ipamSubnet := range ipamSubnets.GetSubnets() {
					err = sv.deleteIpamSubnet(ctx, ipamSubnet.GetSubnetUUID())
					if err != nil {
						return err
					}
				}
			}

			response, err = sv.BaseService.DeleteNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

// UpdateNetworkIpam do pre check for network ipam update
func (sv *ContrailTypeLogicService) UpdateNetworkIpam(
	ctx context.Context,
	request *services.UpdateNetworkIpamRequest) (response *services.UpdateNetworkIpamResponse, err error) {

	//networkIpam := request.GetNetworkIpam()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			//TODO ipam_mgmt check
			//TODO check new subnet method
			//TODO check flat-subnet
			//TODO rest of update checks
			response, err = sv.BaseService.UpdateNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

func (sv *ContrailTypeLogicService) getNetworkIpam(
	ctx context.Context,
	id string) (*models.NetworkIpam, error) {

	networkIpamRes, err := sv.DataService.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{ID: id})
	if err != nil {
		return nil, err
	}
	return networkIpamRes.GetNetworkIpam(), err
}

func (sv *ContrailTypeLogicService) createIpamSubnet(
	ctx context.Context,
	ipamSubnet *models.IpamSubnetType,
	ipamUUID string) (subnetUUID string, err error) {
	createIpamSubnetParams := &ipam.CreateIpamSubnetRequest{
		IpamSubnet:      ipamSubnet,
		NetworkIpamUUID: ipamUUID,
	}
	return sv.AddressManager.CreateIpamSubnet(ctx, createIpamSubnetParams)
}

func (sv *ContrailTypeLogicService) deleteIpamSubnet(
	ctx context.Context,
	subnetUUID string) error {
	deleteIpamSubnetParams := &ipam.DeleteIpamSubnetRequest{
		SubnetUUID: subnetUUID,
	}
	return sv.AddressManager.DeleteIpamSubnet(ctx, deleteIpamSubnetParams)
}
