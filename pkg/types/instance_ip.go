package types

import (
	"context"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
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

			err = validateInstanceIPRefs(instanceIP)
			if err != nil {
				return err
			}

			ipAddress, subnetUUID, err := sv.allocateIPAddressForInstanceIP(ctx, instanceIP)
			if err != nil {
				return err
			}

			instanceIP.InstanceIPAddress = ipAddress
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
				return error
			}

			virtualNetwork, err := sv.getVNFromVirtualNetworkRefs(ctx, virtualNetworkRefs)
			if err != nil {
				return err
			}

			if virtualNetwork.ShouldIgnoreAllocation() {
				var error error
				response, error = sv.BaseService.UpdateInstanceIP(ctx, request)
				return error
			}

			if sv.checkIfIPAddressUpdate(request, requestInstanceIP, databaseInstanceIP) {
				return errutil.ErrorBadRequest("Changing instance-ip-address is not allowed")
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

			if virtualNetwork.ShouldIgnoreAllocation() {
				return nil
			}

			err = sv.deallocIPAddress(ctx, ipAddress, virtualNetwork, nil)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getVNFromVirtualNetworkRefs(
	ctx context.Context, virtualNetworkRefs []*models.InstanceIPVirtualNetworkRef) (*models.VirtualNetwork, error) {

	virtualNetworkResponse, err := sv.ReadService.GetVirtualNetwork(ctx,
		&services.GetVirtualNetworkRequest{
			ID: virtualNetworkRefs[0].GetUUID(),
		})
	if err != nil {
		return nil, err
	}

	return virtualNetworkResponse.GetVirtualNetwork(), nil
}

func validateInstanceIPRefs(instanceIP *models.InstanceIP) error {
	virtualNetworkRefs := instanceIP.GetVirtualNetworkRefs()
	virtualRouterRefs := instanceIP.GetVirtualRouterRefs()
	networkIpamRefs := instanceIP.GetNetworkIpamRefs()

	if len(virtualRouterRefs) > 0 && len(networkIpamRefs) > 0 {
		return errutil.ErrorBadRequest("virtual_router_refs and network_ipam_refs are not allowed")
	}

	if len(virtualRouterRefs) > 0 && len(virtualNetworkRefs) > 0 {
		return errutil.ErrorBadRequest("virtual_router_refs and virtual_network_refs are not allowed")
	}

	return nil
}

func (sv *ContrailTypeLogicService) allocateIPAddressForInstanceIP(
	ctx context.Context, instanceIP *models.InstanceIP,
) (string, string, error) {

	virtualNetworkRefs := instanceIP.GetVirtualNetworkRefs()
	virtualRouterRefs := instanceIP.GetVirtualRouterRefs()

	if len(virtualNetworkRefs) > 0 {
		return sv.allocateIPAddressWithVirtualNetworkRefs(ctx, instanceIP)
	} else if len(virtualRouterRefs) > 0 {
		return sv.allocateIPAddressWithVirtualRouterRefs(ctx, instanceIP)
	}

	return sv.allocateIPAddressWithNetworkIpamRefs(ctx, instanceIP)
}

func (sv *ContrailTypeLogicService) allocateIPAddressWithVirtualNetworkRefs(
	ctx context.Context, instanceIP *models.InstanceIP,
) (string, string, error) {

	virtualNetworkRefs := instanceIP.GetVirtualNetworkRefs()
	virtualNetwork, err := sv.getVNFromVirtualNetworkRefs(ctx, virtualNetworkRefs)
	if err != nil {
		return "", "", err
	}

	if virtualNetwork.ShouldIgnoreAllocation() {
		return instanceIP.InstanceIPAddress, instanceIP.SubnetUUID, nil
	}

	err = sv.alreadyAllocatedIPGatewayCheck(ctx, virtualNetwork, instanceIP)
	if err != nil {
		return "", "", err
	}

	return sv.allocateIPAddress(ctx, virtualNetwork, instanceIP, nil, nil)
}

func (sv *ContrailTypeLogicService) allocateIPAddressWithVirtualRouterRefs(
	ctx context.Context, instanceIP *models.InstanceIP,
) (string, string, error) {

	virtualRouterRefs := instanceIP.GetVirtualRouterRefs()
	virtualRouterNetworkIpamRefs, err := sv.getIpamRefsFromVirtualRouterRefs(ctx, virtualRouterRefs)
	if err != nil {
		return "", "", err
	}

	return sv.allocateIPAddress(ctx, nil, instanceIP, virtualRouterNetworkIpamRefs, nil)
}

func (sv *ContrailTypeLogicService) allocateIPAddressWithNetworkIpamRefs(
	ctx context.Context, instanceIP *models.InstanceIP,
) (string, string, error) {

	instanceIPNetworkIpamRefs := instanceIP.GetNetworkIpamRefs()
	return sv.allocateIPAddress(ctx, nil, instanceIP, nil, instanceIPNetworkIpamRefs)
}

func (sv *ContrailTypeLogicService) getIpamRefsFromVirtualRouterRefs(
	ctx context.Context, virtualRouterRefs []*models.InstanceIPVirtualRouterRef,
) ([]*models.VirtualRouterNetworkIpamRef, error) {

	if len(virtualRouterRefs) > 1 {
		return nil, errutil.ErrorBadRequest("Instance-ip can not refer to multiple vrouters")
	}

	virtualRouterResponse, err := sv.ReadService.GetVirtualRouter(ctx,
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
	if len(ipAddress) == 0 {
		return nil
	}

	isAllocated, err := sv.checkIfRequestedIPAddressIsFree(ctx, virtualNetwork, ipAddress)
	if err != nil {
		return err
	}

	if isAllocated {
		//TODO Gateway IP check
		return nil
	}

	return nil
}

func (sv *ContrailTypeLogicService) allocateIPAddress(
	ctx context.Context,
	virtualNetwork *models.VirtualNetwork,
	instanceIP *models.InstanceIP,
	virtualRouterNetworkIpamRefs []*models.VirtualRouterNetworkIpamRef,
	instanceIPNetworkIpamRefs []*models.InstanceIPNetworkIpamRef,
) (string, string, error) {

	virtualRouterRefs := instanceIP.GetVirtualRouterRefs()
	subnetUUID := instanceIP.GetSubnetUUID()
	ipAddress := instanceIP.GetInstanceIPAddress()
	ipFamily := instanceIP.GetInstanceIPFamily()

	if subnetUUID != "" && len(virtualRouterRefs) > 0 {
		return "", "", errutil.ErrorBadRequest("Subnet uuid based allocation not supported with vrouter")
	}

	if (len(virtualRouterNetworkIpamRefs) > 0 || len(instanceIPNetworkIpamRefs) > 0) && ipAddress != "" {
		return "", "", errutil.ErrorBadRequest("Allocation for requested IP from a network_ipam is not supported")
	}

	allocateIPParams := &ipam.AllocateIPRequest{
		VirtualNetwork:               virtualNetwork,
		IPAddress:                    ipAddress,
		IPFamily:                     ipFamily,
		SubnetUUID:                   subnetUUID,
		VirtualRouterNetworkIpamRefs: virtualRouterNetworkIpamRefs,
		InstanceIPNetworkIpamRefs:    instanceIPNetworkIpamRefs,
	}

	return sv.AddressManager.AllocateIP(ctx, allocateIPParams)
}

func (sv *ContrailTypeLogicService) getInstanceIP(
	ctx context.Context, id string) (*models.InstanceIP, error) {

	instanceIPResponse, err := sv.ReadService.GetInstanceIP(ctx,
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

func (sv *ContrailTypeLogicService) checkIfIPAddressUpdate(request *services.UpdateInstanceIPRequest,
	requestInstanceIP *models.InstanceIP, databaseInstanceIP *models.InstanceIP) bool {
	requestIPAddress := requestInstanceIP.GetInstanceIPAddress()
	databaseIPAddress := databaseInstanceIP.GetInstanceIPAddress()
	fieldMask := request.GetFieldMask()

	if basemodels.FieldMaskContains(&fieldMask, models.InstanceIPFieldInstanceIPAddress) &&
		requestIPAddress != "" && requestIPAddress != databaseIPAddress {
		return true
	}
	return false
}
