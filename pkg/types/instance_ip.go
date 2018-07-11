package types

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"golang.org/x/net/context"
)

//CreateInstanceIP does pre-check for instance-ip
func (sv *ContrailTypeLogicService) CreateInstanceIP(
	ctx context.Context,
	request *services.CreateInstanceIPRequest) (*services.CreateInstanceIPResponse, error) {

	var response *services.CreateInstanceIPResponse
	instanceIP := request.GetInstanceIP()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			virtualNetworkRefs := instanceIP.GetVirtualNetworkRefs()
			virtualRouterRefs := instanceIP.GetVirtualRouterRefs()
			networkIpamRefs := instanceIP.GetNetworkIpamRefs()

			if len(virtualRouterRefs) > 0 && len(networkIpamRefs) > 0 {
				return common.ErrorBadRequestf("virtual_router_refs and network_ipam_refs are not allowed")
			}

			if len(virtualRouterRefs) > 0 && len(virtualNetworkRefs) > 0 {
				return common.ErrorBadRequestf("virtual_router_refs and virtual_network_refs are not allowed")
			}

			virtualNetwork, err := sv.getVNFromVirtualNetworkRefs(ctx, virtualNetworkRefs)
			if err != nil {
				return err
			}

			if sv.shouldIgnoreAllocation(virtualNetwork) {
				response, err = sv.BaseService.CreateInstanceIP(ctx, request)
				return err
			}

			err = sv.alreadyAllocatedIPGatewayCheck(ctx, virtualNetwork, instanceIP)
			if err != nil {
				return err
			}

			floatingIPAddress, subnetUUID, err := sv.allocateIPAddress(ctx, virtualNetwork, instanceIP)
			if err != nil {
				return err
			}

			instanceIP.InstanceIPAddress = floatingIPAddress
			instanceIP.SubnetUUID = subnetUUID

			response, err = sv.BaseService.CreateInstanceIP(ctx, request)
			return err
		})

	return response, err
}

//UpdateInstanceIP does pre-check for instance-ip
func (sv *ContrailTypeLogicService) UpdateInstanceIP(
	ctx context.Context,
	request *services.UpdateInstanceIPRequest) (*services.UpdateInstanceIPResponse, error) {

	var response *services.UpdateInstanceIPResponse
	requestInstanceIP := request.GetInstanceIP()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {

			databaseInstanceIP, err := sv.getInstanceIP(ctx, requestInstanceIP.GetUUID())
			if err != nil {
				return err
			}

			virtualNetworkRefs := databaseInstanceIP.GetVirtualNetworkRefs()
			if len(virtualNetworkRefs) == 0 {
				var error error
				response, error = sv.BaseService.UpdateInstanceIP(ctx, request)
				if error != nil {
					return err
				}
			}

			virtualNetwork, err := sv.getVNFromVirtualNetworkRefs(ctx, virtualNetworkRefs)
			if err != nil {
				return err
			}

			if sv.shouldIgnoreAllocation(virtualNetwork) {
				var error error
				response, error = sv.BaseService.UpdateInstanceIP(ctx, request)
				if error != nil {
					return error
				}
			}

			requestIPAddress := requestInstanceIP.GetInstanceIPAddress()
			databaseIPAddress := databaseInstanceIP.GetInstanceIPAddress()

			if requestIPAddress != "" && requestIPAddress != databaseIPAddress {
				return common.ErrorBadRequestf("Instance-ip address can not be changed")
			}

			//TODO Gateway IP check

			response, err = sv.BaseService.UpdateInstanceIP(ctx, request)
			return err
		})

	return response, err
}

//DeleteInstanceIP does post-check for instance-ip
func (sv *ContrailTypeLogicService) DeleteInstanceIP(
	ctx context.Context,
	request *services.DeleteInstanceIPRequest) (*services.DeleteInstanceIPResponse, error) {

	var response *services.DeleteInstanceIPResponse
	id := request.GetID()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {

			instanceIP, err := sv.getInstanceIP(ctx, id)
			if err != nil {
				return err
			}

			response, err = sv.BaseService.DeleteInstanceIP(ctx, request)
			if err != nil {
				return err
			}

			ipAddress := instanceIP.GetInstanceIPAddress()
			if ipAddress == "" {
				return nil
			}

			networkIpamRefs := instanceIP.GetNetworkIpamRefs()
			if len(networkIpamRefs) > 0 {
				err = sv.deallocIPAddress(ctx, ipAddress, nil, networkIpamRefs)
				return err
			}

			virtualNetworkRefs := instanceIP.GetVirtualNetworkRefs()
			if len(virtualNetworkRefs) == 0 {
				return nil
			}

			virtualNetwork, err := sv.getVNFromVirtualNetworkRefs(ctx, virtualNetworkRefs)
			if err != nil {
				return err
			}

			if sv.shouldIgnoreAllocation(virtualNetwork) {
				return nil
			}

			err = sv.deallocIPAddress(ctx, ipAddress, virtualNetwork, nil)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) shouldIgnoreAllocation(
	virtualNetwork *models.VirtualNetwork) bool {

	for _, fqName := range virtualNetwork.GetFQName() {
		if fqName == "ip-fabric" || fqName == "__link_local__" {
			return true
		}
	}

	return false
}

func (sv *ContrailTypeLogicService) getVNFromVirtualNetworkRefs(
	ctx context.Context, virtualNetworkRefs []*models.InstanceIPVirtualNetworkRef) (*models.VirtualNetwork, error) {

	if len(virtualNetworkRefs) == 0 {
		return nil, nil
	}

	virtualNetworkResponse, err := sv.DataService.GetVirtualNetwork(ctx,
		&services.GetVirtualNetworkRequest{
			ID: virtualNetworkRefs[0].GetUUID(),
		})
	if err != nil {
		return nil, err
	}

	return virtualNetworkResponse.GetVirtualNetwork(), nil
}

func (sv *ContrailTypeLogicService) getIpamRefsFromVirtualRouterRefs(
	ctx context.Context, virtualRouterRefs []*models.InstanceIPVirtualRouterRef,
) ([]*models.VirtualRouterNetworkIpamRef, error) {

	switch {
	case len(virtualRouterRefs) == 0:
		return nil, nil
	case len(virtualRouterRefs) > 1:
		return nil, common.ErrorBadRequestf("Instance-ip can not refer to multiple vrouters")
	}

	virtualRouterResponse, err := sv.DataService.GetVirtualRouter(ctx,
		&services.GetVirtualRouterRequest{
			ID: virtualRouterRefs[0].GetUUID(),
		})
	if err != nil {
		return nil, err
	}

	return virtualRouterResponse.GetVirtualRouter().GetNetworkIpamRefs(), nil
}

func (sv *ContrailTypeLogicService) alreadyAllocatedIPGatewayCheck(ctx context.Context,
	virtualNetwork *models.VirtualNetwork, instanceIP *models.InstanceIP) error {

	ipAddress := instanceIP.GetInstanceIPAddress()
	if len(ipAddress) > 0 {
		isAllocated, err := sv.checkIfRequestedIPAddressIsFree(ctx, virtualNetwork, ipAddress)
		if err != nil {
			return err
		}

		if isAllocated {
			//TODO Gateway IP check
			return nil
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) allocateIPAddress(ctx context.Context,
	virtualNetwork *models.VirtualNetwork, instanceIP *models.InstanceIP) (string, string, error) {

	virtualRouterRefs := instanceIP.GetVirtualRouterRefs()
	subnetUUID := instanceIP.GetSubnetUUID()
	ipAddress := instanceIP.GetInstanceIPAddress()
	ipFamily := instanceIP.GetInstanceIPFamily()

	ipamRefs, err := sv.getIpamRefsFromVirtualRouterRefs(ctx, virtualRouterRefs)
	if err != nil {
		return "", "", err
	}

	if subnetUUID != "" && len(virtualRouterRefs) > 0 {
		return "", "", common.ErrorBadRequestf("Subnet uuid based allocation not supported with vrouter")
	}

	if len(ipamRefs) > 0 && ipAddress != "" {
		return "", "", common.ErrorBadRequestf("Allocation for requested ip from a network_ipam is not supported")
	}

	var allocationPools []*models.AllocationPoolType

	for _, ipamRef := range ipamRefs {
		ipamRefAttr := ipamRef.GetAttr()
		if ipamRefAttr != nil {
			allocationPools = append(allocationPools, ipamRefAttr.GetAllocationPools()...)
		}
	}

	allocateIPParams := &ipam.AllocateIPRequest{
		VirtualNetwork:  virtualNetwork,
		IPAddress:       ipAddress,
		IPFamily:        ipFamily,
		SubnetUUID:      subnetUUID,
		IpamRefs:        ipamRefs,
		AllocationPools: allocationPools,
	}

	return sv.AddressManager.AllocateIP(ctx, allocateIPParams)
}

func (sv *ContrailTypeLogicService) getInstanceIP(
	ctx context.Context, id string) (*models.InstanceIP, error) {

	instanceIPResponse, err := sv.DataService.GetInstanceIP(ctx,
		&services.GetInstanceIPRequest{
			ID: id,
		})
	if err != nil {
		return nil, err
	}
	return instanceIPResponse.GetInstanceIP(), nil
}

func (sv *ContrailTypeLogicService) deallocIPAddress(ctx context.Context,
	ipAddress string, virtualNetwork *models.VirtualNetwork, ipamRefs []*models.InstanceIPNetworkIpamRef) error {

	deallocateIPParams := &ipam.DeallocateIPRequest{
		IPAddress:      ipAddress,
		VirtualNetwork: virtualNetwork,
		IpamRefs:       ipamRefs,
	}

	return sv.AddressManager.DeallocateIP(ctx, deallocateIPParams)
}
