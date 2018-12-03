package types

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

//CreateAliasIP allocates ip in address manager for alias-ip
func (sv *ContrailTypeLogicService) CreateAliasIP(
	ctx context.Context,
	request *services.CreateAliasIPRequest) (*services.CreateAliasIPResponse, error) {

	var response *services.CreateAliasIPResponse
	aliasIP := request.GetAliasIP()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error
			aliasIPPool, err := sv.getAliasIPAliasPool(ctx, aliasIP)
			if err != nil {
				return errors.Wrap(err, "failed to get alias ip pool")
			}
			network, err := sv.getVirtualNetworkFromAliasIPPool(ctx, aliasIPPool)
			if err != nil {
				return errors.Wrap(err, "failed to get virtual network for alias ip pool")
			}
			if len(aliasIP.AliasIPAddress) != 0 {
				var isAllocated bool
				isAllocated, err = sv.isAllocatedAliasIPAddress(ctx, aliasIP, network)
				if err != nil {
					return errors.Wrap(err, "failed to check if alias ip address is already allocated")
				}
				if isAllocated {
					return grpc.Errorf(codes.AlreadyExists, "ip address %v already in use", aliasIP.AliasIPAddress)
				}
			}
			aipAddr, _, err := sv.AddressManager.AllocateIP(ctx, &ipam.AllocateIPRequest{
				VirtualNetwork: network,
				IPAddress:      aliasIP.AliasIPAddress,
			})
			if err != nil {
				return errors.Wrap(err, "failed to allocate alias ip")
			}
			aliasIP.AliasIPAddress = aipAddr
			response, err = sv.BaseService.CreateAliasIP(ctx, request)
			return err
		})
	return response, err
}

//DeleteAliasIP deletes ip from address manager for alias-ip
func (sv *ContrailTypeLogicService) DeleteAliasIP(
	ctx context.Context,
	request *services.DeleteAliasIPRequest) (*services.DeleteAliasIPResponse, error) {

	var response *services.DeleteAliasIPResponse
	id := request.GetID()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error
			aliasIP, err := sv.getAliasIP(ctx, id)
			if err != nil {
				return errors.Wrap(err, "failed to get alias ip pool")
			}
			aliasIPPool, err := sv.getAliasIPAliasPool(ctx, aliasIP)
			if err != nil {
				return errors.Wrap(err, "failed to get virtual network for alias ip pool")
			}
			network, err := sv.getVirtualNetworkFromAliasIPPool(ctx, aliasIPPool)
			if err != nil {
				return errors.Wrap(err, "failed to get virtual network for alias ip pool")
			}
			deallocateIPParams := &ipam.DeallocateIPRequest{
				VirtualNetwork: network,
				IPAddress:      aliasIP.GetAliasIPAddress(),
			}
			err = sv.AddressManager.DeallocateIP(ctx, deallocateIPParams)
			if err != nil {
				return errors.Wrap(err, "failed to deallocate alias ip")
			}
			response, err = sv.BaseService.DeleteAliasIP(ctx, request)
			return err
		})
	return response, err
}

func (sv *ContrailTypeLogicService) isAllocatedAliasIPAddress(
	ctx context.Context, aliasIP *models.AliasIP, network *models.VirtualNetwork) (bool, error) {
	isIPAllocatedRequest := &ipam.IsIPAllocatedRequest{
		VirtualNetwork: network,
		IPAddress:      aliasIP.AliasIPAddress,
	}

	isAllocated, err := sv.AddressManager.IsIPAllocated(ctx, isIPAllocatedRequest)
	if err != nil {
		return false, err
	}
	return isAllocated, err
}

func (sv *ContrailTypeLogicService) getAliasIPAliasPool(
	ctx context.Context, aliasIP *models.AliasIP) (*models.AliasIPPool, error) {
	aliasIPPool, err := sv.ReadService.GetAliasIPPool(ctx, &services.GetAliasIPPoolRequest{
		ID: aliasIP.GetParentUUID(),
	})
	if err != nil {
		return nil, err
	}
	return aliasIPPool.GetAliasIPPool(), nil
}

func (sv *ContrailTypeLogicService) getVirtualNetworkFromAliasIPPool(
	ctx context.Context, aliasIPPool *models.AliasIPPool) (*models.VirtualNetwork, error) {

	virtualNetworkResponse, err := sv.ReadService.GetVirtualNetwork(ctx,
		&services.GetVirtualNetworkRequest{
			ID: aliasIPPool.GetParentUUID(),
		})
	if err != nil {
		return nil, errutil.ErrorBadRequestf("Missing virtual-network with uuid %s: %v",
			aliasIPPool.GetParentUUID(), err)
	}

	return virtualNetworkResponse.GetVirtualNetwork(), nil
}

func (sv *ContrailTypeLogicService) getAliasIP(ctx context.Context, id string) (*models.AliasIP, error) {
	aliasIPResponse, err := sv.ReadService.GetAliasIP(ctx,
		&services.GetAliasIPRequest{
			ID: id,
		})
	if err != nil {
		return nil, err
	}
	return aliasIPResponse.GetAliasIP(), nil
}
