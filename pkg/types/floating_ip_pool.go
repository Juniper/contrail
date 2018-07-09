package types

import (
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateFloatingIPPool validates data before creating floating-ip-pool
func (sv *ContrailTypeLogicService) CreateFloatingIPPool(
	ctx context.Context,
	request *services.CreateFloatingIPPoolRequest) (*services.CreateFloatingIPPoolResponse, error) {

	var response *services.CreateFloatingIPPoolResponse
	floatingIPPool := request.GetFloatingIPPool()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error
			if sv.checkIfSubnetInfoIsNotSpecified(floatingIPPool) {
				response, err = sv.Next().CreateFloatingIPPool(ctx, request)
				return err
			}

			virtualNetwork, err := sv.getVirtualNetworkFromFloatingIPPool(ctx, floatingIPPool)
			if err != nil {
				return err
			}

			err = sv.validateFloatingIPPoolSubnets(floatingIPPool, virtualNetwork)
			if err != nil {
				return err
			}

			response, err = sv.Next().CreateFloatingIPPool(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) checkIfSubnetInfoIsNotSpecified(floatingIPPool *models.FloatingIPPool) bool {

	if !floatingIPPool.IsParentTypeVirtualNetwork() {
		return true
	}

	floatingIPPoolSubnets := floatingIPPool.GetFloatingIPPoolSubnets()
	return floatingIPPoolSubnets == nil || len(floatingIPPoolSubnets.GetSubnetUUID()) == 0
}

func (sv *ContrailTypeLogicService) getVirtualNetworkFromFloatingIPPool(
	ctx context.Context, floatingIPPool *models.FloatingIPPool) (*models.VirtualNetwork, error) {

	virtualNetworkResponse, err := sv.DataService.GetVirtualNetwork(ctx,
		&services.GetVirtualNetworkRequest{
			ID: floatingIPPool.GetParentUUID(),
		})
	if err != nil {
		return nil, err
	}

	return virtualNetworkResponse.GetVirtualNetwork(), nil
}

func (sv *ContrailTypeLogicService) validateFloatingIPPoolSubnets(
	floatingIPPool *models.FloatingIPPool, virtualNetwork *models.VirtualNetwork) error {

	for _, floatingIPPoolSubnetUUID := range floatingIPPool.GetFloatingIPPoolSubnets().GetSubnetUUID() {
		subnetFound := false
		for _, ipam := range virtualNetwork.GetNetworkIpamRefs() {
			if ipam.GetAttr() == nil {
				continue
			}

			for _, ipamSubnet := range ipam.GetAttr().GetIpamSubnets() {
				if ipamSubnet.GetSubnetUUID() == floatingIPPoolSubnetUUID {
					subnetFound = true
					break
				}
			}
		}

		if !subnetFound {
			return common.ErrorBadRequest("Subnet " + floatingIPPoolSubnetUUID +
				" was not found in virtual-network " + virtualNetwork.GetUUID())
		}
	}
	return nil
}
