package types

import (
	"context"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateFloatingIPPool validates if subnets in the floating-ip-pool object exist in
//the virtual-network. If subnet info is not specified, there is nothing to validate
func (sv *ContrailTypeLogicService) CreateFloatingIPPool(
	ctx context.Context,
	request *services.CreateFloatingIPPoolRequest) (*services.CreateFloatingIPPoolResponse, error) {

	var response *services.CreateFloatingIPPoolResponse
	floatingIPPool := request.GetFloatingIPPool()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error
			var virtualNetwork *models.VirtualNetwork

			if floatingIPPool.IsParentTypeVirtualNetwork() && floatingIPPool.HasSubnets() {
				virtualNetwork, err = sv.getVirtualNetworkFromFloatingIPPool(ctx, floatingIPPool)
				if err != nil {
					return err
				}

				err = floatingIPPool.CheckAreSubnetsInVirtualNetworkSubnets(virtualNetwork)
				if err != nil {
					return err
				}
			}

			response, err = sv.BaseService.CreateFloatingIPPool(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getVirtualNetworkFromFloatingIPPool(
	ctx context.Context, floatingIPPool *models.FloatingIPPool) (*models.VirtualNetwork, error) {

	virtualNetworkResponse, err := sv.ReadService.GetVirtualNetwork(ctx,
		&services.GetVirtualNetworkRequest{
			ID: floatingIPPool.GetParentUUID(),
		})
	if err != nil {
		return nil, errutil.ErrorBadRequestf("Missing virtual-network with uuid %s: %v",
			floatingIPPool.GetParentUUID(), err)
	}

	return virtualNetworkResponse.GetVirtualNetwork(), nil
}
