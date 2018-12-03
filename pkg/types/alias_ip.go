package types

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

//CreateAliasIP allocates ip in address manager for alias-ip
func (sv *ContrailTypeLogicService) CreateAliasIP(
	ctx context.Context,
	request *services.CreateAliasIPRequest,
) (response *services.CreateAliasIPResponse, err error) {
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			aliasIP := request.GetAliasIP()
			aipAddr, err := sv.allocateAliasIPAddress(ctx, aliasIP)
			if err != nil {
				return errors.Wrapf(err, "failed to allocate alias ip")
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
	request *services.DeleteAliasIPRequest,
) (response *services.DeleteAliasIPResponse, err error) {
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			aliasIP, err := sv.getAliasIP(ctx, request.GetID())
			if err != nil {
				return errors.Wrapf(err, "failed to get alias ip")
			}
			err = sv.deallocateAliasIPAddress(ctx, aliasIP)
			if err != nil {
				return errors.Wrapf(err, "failed to deallocate alias ip")
			}
			response, err = sv.BaseService.DeleteAliasIP(ctx, request)
			return err
		})
	return response, err
}

func (sv *ContrailTypeLogicService) allocateAliasIPAddress(
	ctx context.Context,
	aliasIP *models.AliasIP,
) (string, error) {
	network, err := sv.getVirtualNetworkFromAliasIP(ctx, aliasIP)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get virtual network for alias ip")
	}
	if len(aliasIP.AliasIPAddress) != 0 {
		var isAllocated bool
		isIPAllocatedRequest := &ipam.IsIPAllocatedRequest{
			VirtualNetwork: network,
			IPAddress:      aliasIP.AliasIPAddress,
		}
		isAllocated, err = sv.AddressManager.IsIPAllocated(ctx, isIPAllocatedRequest)
		if err != nil {
			return "", errors.Wrapf(err, "failed to check if alias ip address is already allocated")
		}
		if isAllocated {
			return "", errors.Wrapf(
				grpc.Errorf(codes.AlreadyExists, "ip address %v already in use", aliasIP.AliasIPAddress), "test")
		}
	}
	aipAddr, _, err := sv.AddressManager.AllocateIP(ctx, &ipam.AllocateIPRequest{
		VirtualNetwork: network,
		IPAddress:      aliasIP.AliasIPAddress,
	})
	return aipAddr, err
}

func (sv *ContrailTypeLogicService) deallocateAliasIPAddress(
	ctx context.Context,
	aliasIP *models.AliasIP,
) error {
	network, err := sv.getVirtualNetworkFromAliasIP(ctx, aliasIP)
	if err != nil {
		return errors.Wrapf(err, "failed to get virtual network for alias ip")
	}
	return sv.AddressManager.DeallocateIP(ctx, &ipam.DeallocateIPRequest{
		VirtualNetwork: network,
		IPAddress:      aliasIP.AliasIPAddress,
	})
}

func (sv *ContrailTypeLogicService) getVirtualNetworkFromAliasIP(
	ctx context.Context,
	aliasIP *models.AliasIP,
) (*models.VirtualNetwork, error) {
	aliasIPPoolResponse, err := sv.ReadService.GetAliasIPPool(ctx, &services.GetAliasIPPoolRequest{
		ID: aliasIP.GetParentUUID(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get alias ip pool with uuid %s",
			aliasIP.GetParentUUID())
	}
	virtualNetworkResponse, err := sv.ReadService.GetVirtualNetwork(ctx,
		&services.GetVirtualNetworkRequest{
			ID: aliasIPPoolResponse.GetAliasIPPool().GetParentUUID(),
		})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get virtual-network with uuid %s",
			aliasIPPoolResponse.GetAliasIPPool().GetParentUUID())
	}
	return virtualNetworkResponse.GetVirtualNetwork(), nil
}

func (sv *ContrailTypeLogicService) getAliasIP(
	ctx context.Context,
	id string,
) (*models.AliasIP, error) {
	aliasIPResponse, err := sv.ReadService.GetAliasIP(ctx,
		&services.GetAliasIPRequest{
			ID: id,
		})
	if err != nil {
		return nil, err
	}
	return aliasIPResponse.GetAliasIP(), nil
}
