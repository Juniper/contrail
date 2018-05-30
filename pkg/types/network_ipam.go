package types

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
	"net"
)

func (service *ContrailTypeLogicService) CreateNetworkIpam(
		ctx context.Context,
		request *models.CreateNetworkIpamRequest) (response *models.CreateNetworkIpamResponse, err error) {

	networkIpam := request.GetNetworkIpam()
	err = db.DoInTransaction(
		ctx,
		service.DB.DB(),
		func(ctx context.Context) error {
			ipamSubnets := networkIpam.GetIpamSubnets()
			if (ipamSubnets == nil && networkIpam.IsFlatSubnet()) {
				return common.ErrorBadRequest("Ipam subnets are allowed only with flat-subnet")
			}
			if (networkIpam.IsFlatSubnet()|| ipamSubnets == nil) {
				response, err = service.Next().CreateNetworkIpam(ctx, request)
				return err
			}

			subnets := []*net.IPNet{}
			subnets, err = mergeSubnetIfNoOverlap(subnets, ipamSubnets.GetSubnets())
			if err != nil {
				return err
			}

			for _, ipamSubnet := range ipamSubnets.GetSubnets() {
				err = ipamSubnet.Validate()
				if err != nil {
					return err
				}
			}
			//TODO try to ipam_create_rew in addr_mgmt
			response, err = service.Next().CreateNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

func (service *ContrailTypeLogicService) DeleteNetworkIpam(
		ctx context.Context,
		request *models.DeleteNetworkIpamRequest) (response *models.DeleteNetworkIpamResponse, err error) {

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

			if(networkIpam == nil || networkIpam.IsFlatSubnet() || networkIpam.GetIpamSubnets() == nil) {
				response, err = service.Next().DeleteNetworkIpam(ctx, request)
				return err
			}
			// TODO addr_mgmt ipam delete req
			response, err = service.Next().DeleteNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

func (service *ContrailTypeLogicService) UpdateNetworkIpam(
	ctx context.Context,
	request *models.UpdateNetworkIpamRequest) (response *models.UpdateNetworkIpamResponse, err error) {

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

	networkIpamRes, err := service.DB.GetNetworkIpam(ctx,&models.GetNetworkIpamRequest{ID: id})
	if err != nil {
		return nil, err
	}

	return networkIpamRes.GetNetworkIpam(), err
}
