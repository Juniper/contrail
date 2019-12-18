package types

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

// CreateFloatingIP checks parent type and if parent type isn't instance-ip then this method tries to
// allocate IP using AddressManager in subnets from floating-ip-pool(parent).
func (sv *ContrailTypeLogicService) CreateFloatingIP(
	ctx context.Context,
	request *services.CreateFloatingIPRequest) (*services.CreateFloatingIPResponse, error) {

	var response *services.CreateFloatingIPResponse
	floatingIP := request.GetFloatingIP()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error
			if sv.checkIfParentTypeIsInstanceIP(floatingIP) {
				response, err = sv.BaseService.CreateFloatingIP(ctx, request)
				return err
			}

			virtualNetwork, err := sv.getVirtualNetwork(ctx, floatingIP)
			if err != nil {
				return err
			}

			ipAddress := floatingIP.GetFloatingIPAddress()
			if len(ipAddress) > 0 {
				var isAllocated bool
				isAllocated, err = sv.checkIfRequestedIPAddressIsFree(ctx, virtualNetwork, ipAddress)
				if err != nil {
					return err
				}
				if isAllocated {
					return grpc.Errorf(codes.AlreadyExists, "Ip address %v already in use", ipAddress)
				}
			}

			floatingIPAddress, err := sv.tryToAllocateIPAddress(ctx, virtualNetwork, floatingIP)
			if err != nil {
				return err
			}
			floatingIP.FloatingIPAddress = floatingIPAddress

			response, err = sv.BaseService.CreateFloatingIP(ctx, request)
			return err
		})

	return response, err
}

// DeleteFloatingIP checks parent type and if it isn't instance-ip then this method tries to
// deallocate IP using AddressManager
func (sv *ContrailTypeLogicService) DeleteFloatingIP(
	ctx context.Context,
	request *services.DeleteFloatingIPRequest) (*services.DeleteFloatingIPResponse, error) {

	var response *services.DeleteFloatingIPResponse

	id := request.GetID()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			floatingIP, err := sv.getFloatingIP(ctx, id)
			if err != nil {
				return err
			}

			response, err = sv.BaseService.DeleteFloatingIP(ctx, request)
			if err != nil {
				return err
			}

			if sv.checkIfParentTypeIsInstanceIP(floatingIP) {
				return nil
			}

			virtualNetwork, err := sv.getVirtualNetwork(ctx, floatingIP)
			if err != nil {
				return err
			}

			return sv.deallocateIPAddress(ctx, virtualNetwork, floatingIP)
		})

	return response, err
}

func (sv *ContrailTypeLogicService) deallocateIPAddress(ctx context.Context, virtualNetwork *models.VirtualNetwork,
	floatingIP *models.FloatingIP) error {

	deallocateIPParams := &ipam.DeallocateIPRequest{
		VirtualNetwork: virtualNetwork,
		IPAddress:      floatingIP.GetFloatingIPAddress(),
	}

	return sv.AddressManager.DeallocateIP(ctx, deallocateIPParams)
}

func (sv *ContrailTypeLogicService) checkIfParentTypeIsInstanceIP(floatingIP *models.FloatingIP) bool {
	return floatingIP.IsParentTypeInstanceIP()
}

func (sv *ContrailTypeLogicService) checkIfRequestedIPAddressIsFree(ctx context.Context,
	virtualNetwork *models.VirtualNetwork, ipAddress string) (bool, error) {

	isIPAllocatedRequest := &ipam.IsIPAllocatedRequest{
		VirtualNetwork: virtualNetwork,
		IPAddress:      ipAddress,
	}

	isAllocated, err := sv.AddressManager.IsIPAllocated(ctx, isIPAllocatedRequest)
	if err != nil {
		return false, err
	}
	return isAllocated, err
}

func (sv *ContrailTypeLogicService) tryToAllocateIPAddress(ctx context.Context,
	virtualNetwork *models.VirtualNetwork, floatingIP *models.FloatingIP) (string, error) {

	floatingIPPoolSubnets, err := sv.getFloatingIPPoolSubnets(ctx, floatingIP)
	if err != nil {
		return "", errutil.ErrorBadRequest("floating-ip-pool lookup failed with error:" + err.Error())
	}

	var floatingIPAddress string
	if floatingIPPoolSubnets == nil || len(floatingIPPoolSubnets.GetSubnetUUID()) == 0 {
		// Subnet specification was not found on the floating-ip-pool.
		// Proceed to allocated floating-ip from any of the subnets
		// on the virtual-network.
		floatingIPAddress, _, err = sv.AddressManager.AllocateIP(
			ctx, &ipam.AllocateIPRequest{
				VirtualNetwork: virtualNetwork,
				IPAddress:      floatingIP.GetFloatingIPAddress(),
			})
		if err != nil {
			return "", err
		}
		return floatingIPAddress, nil
	}
	var subnetsTried []string
	for _, floatingIPPoolSubnetUUID := range floatingIPPoolSubnets.GetSubnetUUID() {
		// Record the subnets that we try to allocate from.
		subnetsTried = append(subnetsTried, floatingIPPoolSubnetUUID)

		allocateIPParams := &ipam.AllocateIPRequest{
			VirtualNetwork: virtualNetwork,
			IPAddress:      floatingIP.GetFloatingIPAddress(),
			SubnetUUID:     floatingIPPoolSubnetUUID,
		}

		floatingIPAddress, _, err = sv.AddressManager.AllocateIP(ctx, allocateIPParams)
		if _, ok := err.(ipam.ErrSubnetExhausted); ok {
			// This subnet is exhausted. Try next subnet.
			continue
		}

		if err != nil {
			return "", err
		}
	}

	if floatingIPAddress == "" {
		// Floating-ip could not be allocated from any of the
		// configured subnets. Return ResourceExhausted
		return "", grpc.Errorf(codes.ResourceExhausted, "subnets tried: %s", strings.Join(subnetsTried, ", "))
	}
	return floatingIPAddress, nil
}

// Get virtual network associated with floatingIP and floatingIPPool
func (sv *ContrailTypeLogicService) getVirtualNetwork(
	ctx context.Context, floatingIP *models.FloatingIP) (*models.VirtualNetwork, error) {

	floatingIPPoolResponse, err := sv.ReadService.GetFloatingIPPool(ctx,
		&services.GetFloatingIPPoolRequest{
			ID: floatingIP.GetParentUUID(),
		})
	if err != nil {
		return nil, err
	}

	virtualNetworkResponse, err := sv.ReadService.GetVirtualNetwork(ctx,
		&services.GetVirtualNetworkRequest{
			ID: floatingIPPoolResponse.GetFloatingIPPool().GetParentUUID(),
		})
	if err != nil {
		return nil, err
	}

	return virtualNetworkResponse.GetVirtualNetwork(), nil
}

func (sv *ContrailTypeLogicService) getFloatingIP(
	ctx context.Context, id string) (*models.FloatingIP, error) {

	floatingIPResponse, err := sv.ReadService.GetFloatingIP(ctx,
		&services.GetFloatingIPRequest{
			ID: id,
		})
	if err != nil {
		return nil, err
	}
	return floatingIPResponse.GetFloatingIP(), nil
}

// Get all subnets configured on the floating-ip-pool.
// It is acceptable that list of subnets may be absent or empty
func (sv *ContrailTypeLogicService) getFloatingIPPoolSubnets(
	ctx context.Context,
	floatingIP *models.FloatingIP) (*models.FloatingIpPoolSubnetType, error) {

	floatingIPPoolResponse, err := sv.ReadService.GetFloatingIPPool(ctx,
		&services.GetFloatingIPPoolRequest{
			ID: floatingIP.GetParentUUID(),
		})
	if err != nil {
		return nil, err
	}

	floatingIPPool := floatingIPPoolResponse.GetFloatingIPPool()
	return floatingIPPool.GetFloatingIPPoolSubnets(), nil
}
