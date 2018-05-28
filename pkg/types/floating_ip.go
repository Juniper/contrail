package types

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
)

//CreateFloatingIP do pre check for floating ip.
func (service *ContrailTypeLogicService) CreateFloatingIP(
	ctx context.Context,
	request *models.CreateFloatingIPRequest) (response *models.CreateFloatingIPResponse, err error) {

	floatingIP := request.GetFloatingIP()
	err = db.DoInTransaction(
		ctx,
		service.DB.DB(),
		func(ctx context.Context) error {

			if floatingIP.GetParentType() == "instance-ip" {
				response, err = service.Next().CreateFloatingIP(ctx, request)
				return err
			}
			//TODO: Implement addr mgmt
			//TODO: Check if requested ip address is already in use.

			floatingIPPoolSubnets, err := service.getFloatingIPPoolSubnets(ctx, floatingIP)
			if err != nil {
				return common.ErrorBadRequest("Floating-ip-pool lookup failed with error:" + err.Error())
			}

			if floatingIPPoolSubnets == nil {
				//TODO: Subnet specification was not found on the floating-ip-pool.
				//		Proceed to allocated floating-ip from any of the subnets
				//		on the virtual-network.

				//TODO: allocate ip using addr mgmt
				response, err = service.Next().CreateFloatingIP(ctx, request)
				return err
			}
			//TODO: Iterate through configured subnets on floating-ip-pool.
			//		We will try to allocate floating-ip by iterating through
			//		the list of configured subnets.

			//TODO: try to allocate ip using addr mgmt

			response, err = service.Next().CreateFloatingIP(ctx, request)
			return nil
		})

	return response, err
}

// DeleteFloatingIP do post actions for delete floating ip.
func (service *ContrailTypeLogicService) DeleteFloatingIP(
	ctx context.Context,
	request *models.DeleteFloatingIPRequest) (response *models.DeleteFloatingIPResponse, err error) {

	id := request.GetID()
	err = db.DoInTransaction(
		ctx,
		service.DB.DB(),
		func(ctx context.Context) error {

			response, err = service.Next().DeleteFloatingIP(ctx, request)
			if err != nil {
				return err
			}

			floatingIP, err := service.getFloatingIP(ctx, id)
			if err != nil {
				return err
			}

			if floatingIP.GetParentType() == "instance-ip" {
				return nil
			}

			// TODO: free ip using addr mgmt
			return nil
		})

	return response, err
}

func (service *ContrailTypeLogicService) getFloatingIP(
	ctx context.Context, id string) (*models.FloatingIP, error) {

	floatingIPRes, err := service.DB.GetFloatingIP(ctx, &models.GetFloatingIPRequest{ID: id})
	if err != nil {
		return nil, err
	}
	return floatingIPRes.GetFloatingIP(), nil
}

// Get any subnets configured on the floating-ip-pool.
// It is acceptable that subnet list may be absent or empty
func (service *ContrailTypeLogicService) getFloatingIPPoolSubnets(
	ctx context.Context,
	floatingIP *models.FloatingIP) (*models.FloatingIpPoolSubnetType, error) {

	floatingIPPoolRes, err := service.DB.GetFloatingIPPool(
		ctx, &models.GetFloatingIPPoolRequest{ID: floatingIP.GetParentUUID()})
	if err != nil {
		return nil, err
	}

	floatingIPPool := floatingIPPoolRes.GetFloatingIPPool()

	return floatingIPPool.GetFloatingIPPoolSubnets(), nil
}
