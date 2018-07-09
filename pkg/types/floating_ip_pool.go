package types

import (
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateFloatingIPPool validates if subnets in the floating-ip-pool object exist in
//the virtual-machine. If subnet info is not specified, there is nothing to validate
func (sv *ContrailTypeLogicService) CreateFloatingIPPool(
	ctx context.Context,
	request *services.CreateFloatingIPPoolRequest) (*services.CreateFloatingIPPoolResponse, error) {

	var response *services.CreateFloatingIPPoolResponse
	floatingIPPool := request.GetFloatingIPPool()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			if !floatingIPPool.IsParentTypeVirtualNetwork() || sv.checkMissingSubnetInfo(floatingIPPool) {
				response, err = sv.BaseService.CreateFloatingIPPool(ctx, request)
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

			response, err = sv.BaseService.CreateFloatingIPPool(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) checkMissingSubnetInfo(floatingIPPool *models.FloatingIPPool) bool {

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
		return nil, common.ErrorBadRequestf("Missing virtual-network with uuid %s",
			floatingIPPool.GetParentUUID())
	}

	return virtualNetworkResponse.GetVirtualNetwork(), nil
}

func (sv *ContrailTypeLogicService) validateFloatingIPPoolSubnets(
	floatingIPPool *models.FloatingIPPool, virtualNetwork *models.VirtualNetwork) error {

	for _, floatingIPPoolSubnetUUID := range floatingIPPool.GetFloatingIPPoolSubnets().GetSubnetUUID() {
		subnetFound := false
		for _, ipam := range virtualNetwork.GetNetworkIpamRefs() {
			for _, ipamSubnet := range ipam.GetAttr().GetIpamSubnets() {
				if ipamSubnet.GetSubnetUUID() == floatingIPPoolSubnetUUID {
					subnetFound = true
					break
				}
			}
		}

		if !subnetFound {
			return common.ErrorBadRequestf("Subnet %s was not found in virtual-network %s",
				floatingIPPoolSubnetUUID, virtualNetwork.GetUUID())
		}
	}
	return nil
}
