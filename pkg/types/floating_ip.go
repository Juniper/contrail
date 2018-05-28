package types

import (
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// CreateFloatingIP do pre check for floating ip.
// TODO: This needs to be refactored, since this method is too long
func (s *ContrailTypeLogicService) CreateFloatingIP(
	ctx context.Context,
	request *models.CreateFloatingIPRequest) (response *models.CreateFloatingIPResponse, err error) {

	floatingIP := request.GetFloatingIP()
	err = db.DoInTransaction(
		ctx,
		s.DB.DB(),
		func(ctx context.Context) error {
			if floatingIP.IsParentTypeInstanceIp() {
				response, err = s.Next().CreateFloatingIP(ctx, request)
				return err
			}

			var virtualNetwork *models.VirtualNetwork
			virtualNetwork, err = s.getVirtualNetwork(ctx, floatingIP)
			if err != nil {
				return err
			}

			ipAddress := floatingIP.GetFloatingIPAddress()
			isIPAllocatedRequest := &ipam.IsIPAllocatedRequest{
				VirtualNetwork: virtualNetwork,
				IPAddress:      ipAddress,
			}

			if len(ipAddress) > 0 {
				var isAllocated bool
				isAllocated, err = s.AddressManager.IsIPAllocated(ctx, isIPAllocatedRequest)
				if err != nil {
					return err
				}

				if isAllocated {
					return grpc.Errorf(codes.AlreadyExists, "ip address already in use")
				}
			}

			var floatingIPPoolSubnets *models.FloatingIpPoolSubnetType
			floatingIPPoolSubnets, err = s.getFloatingIPPoolSubnets(ctx, floatingIP)
			if err != nil {
				return common.ErrorBadRequest("floating-ip-pool lookup failed with error:" + err.Error())
			}

			var floatingIPAddress string
			var subnetsTried []string

			if floatingIPPoolSubnets == nil {
				// Subnet specification was not found on the floating-ip-pool.
				// Proceed to allocated floating-ip from any of the subnets
				// on the virtual-network.
				floatingIPAddress, _, err = s.AddressManager.AllocateIP(
					ctx, &ipam.AllocateIPRequest{VirtualNetwork: virtualNetwork,
						IPAddress: ipAddress, AllocatorID: floatingIP.GetUUID()})
				if err != nil {
					return err
				}
			} else {
				for _, floatingIPPoolSubnetUUID := range floatingIPPoolSubnets.GetSubnetUUID() {
					// Record the subnets that we try to allocate from.
					subnetsTried = append(subnetsTried, floatingIPPoolSubnetUUID)

					allocateIPParams := &ipam.AllocateIPRequest{
						VirtualNetwork: virtualNetwork,
						IPAddress:      ipAddress,
						SubnetUUID:     floatingIPPoolSubnetUUID,
						AllocatorID:    floatingIP.GetUUID(),
					}

					floatingIPAddress, _, err = s.AddressManager.AllocateIP(ctx, allocateIPParams)
					if addressManagementError, ok := err.(ipam.AddressManagerError); ok &&
						addressManagementError.GetAddressManagementErrorCode() == ipam.SubnetExhausted {
						// This subnet is exhausted. Try next subnet.
						continue
					}

					if err != nil {
						return err
					}
				}

				if floatingIPAddress == "" {
					// Floating-ip could not be allocated from any of the
					// configured subnets. Return ResourceExhausted
					return grpc.Errorf(codes.ResourceExhausted, "subnets tried: %s", strings.Join(subnetsTried, ", "))
				}
			}

			response, err = s.Next().CreateFloatingIP(ctx, request)
			if err != nil {
				return err
			}

			response.GetFloatingIP().FloatingIPAddress = floatingIPAddress

			return nil
		})

	return response, err
}

// DeleteFloatingIP do post actions for delete floating ip.
func (s *ContrailTypeLogicService) DeleteFloatingIP(
	ctx context.Context,
	request *models.DeleteFloatingIPRequest) (response *models.DeleteFloatingIPResponse, err error) {

	id := request.GetID()
	err = db.DoInTransaction(
		ctx,
		s.DB.DB(),
		func(ctx context.Context) error {
			var floatingIP *models.FloatingIP
			floatingIP, err = s.getFloatingIP(ctx, id)
			if err != nil {
				return err
			}

			response, err = s.Next().DeleteFloatingIP(ctx, request)
			if err != nil {
				return err
			}

			if floatingIP.IsParentTypeInstanceIp() {
				return nil
			}

			deallocateIPParams := &ipam.DeallocateIPRequest{
				VirtualNetwork: nil,
				IPAddress:      floatingIP.GetFloatingIPAddress(),
				AllocatorID:    floatingIP.GetUUID(),
			}

			return s.AddressManager.DeallocateIP(ctx, deallocateIPParams)
		})

	return response, err
}

func (s *ContrailTypeLogicService) getVirtualNetwork(
	ctx context.Context, floatingIP *models.FloatingIP) (*models.VirtualNetwork, error) {

	floatingIPPoolRes, err := s.DB.GetFloatingIPPool(ctx, &models.GetFloatingIPPoolRequest{ID: floatingIP.GetParentUUID()})
	if err != nil {
		return nil, err
	}

	virtualNetworkRes, err := s.DB.GetVirtualNetwork(ctx, &models.GetVirtualNetworkRequest{ID: floatingIPPoolRes.GetFloatingIPPool().GetParentUUID()})
	if err != nil {
		return nil, err
	}

	return virtualNetworkRes.GetVirtualNetwork(), nil
}

func (s *ContrailTypeLogicService) getFloatingIP(
	ctx context.Context, id string) (*models.FloatingIP, error) {

	floatingIPRes, err := s.DB.GetFloatingIP(ctx, &models.GetFloatingIPRequest{ID: id})
	if err != nil {
		return nil, err
	}
	return floatingIPRes.GetFloatingIP(), nil
}

// Get any subnets configured on the floating-ip-pool.
// It is acceptable that subnet list may be absent or empty
func (s *ContrailTypeLogicService) getFloatingIPPoolSubnets(
	ctx context.Context,
	floatingIP *models.FloatingIP) (*models.FloatingIpPoolSubnetType, error) {

	floatingIPPoolRes, err := s.DB.GetFloatingIPPool(
		ctx, &models.GetFloatingIPPoolRequest{ID: floatingIP.GetParentUUID()})
	if err != nil {
		return nil, err
	}

	floatingIPPool := floatingIPPoolRes.GetFloatingIPPool()

	return floatingIPPool.GetFloatingIPPoolSubnets(), nil
}
