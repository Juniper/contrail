package types

import (
	"context"

	protobuf "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateVirtualRouter virtual-router create specific logic.
func (sv *ContrailTypeLogicService) CreateVirtualRouter(
	ctx context.Context, request *services.CreateVirtualRouterRequest,
) (*services.CreateVirtualRouterResponse, error) {
	var response *services.CreateVirtualRouterResponse
	vr := request.VirtualRouter
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			err := sv.validateVrouterAllocationPools(ctx, vr)
			if err != nil {
				return errors.Wrapf(err, "virtual-router %s allocation pools validations failed",
					vr.GetUUID())
			}
			response, err = sv.BaseService.CreateVirtualRouter(ctx, request)
			return err
		})

	return response, err
}

// UpdateVirtualRouter virtual-router update specific logic.
func (sv *ContrailTypeLogicService) UpdateVirtualRouter(
	ctx context.Context, request *services.UpdateVirtualRouterRequest,
) (*services.UpdateVirtualRouterResponse, error) {
	var response *services.UpdateVirtualRouterResponse
	vr := request.GetVirtualRouter()
	fm := request.GetFieldMask()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			dbVrRes, err := sv.ReadService.GetVirtualRouter(ctx, &services.GetVirtualRouterRequest{
				ID: vr.GetUUID(),
			})
			if err != nil {
				return err
			}

			if err = sv.vrouterCheckAllocationPoolsDelete(ctx, vr, &fm, dbVrRes.GetVirtualRouter()); err != nil {
				return errors.Wrapf(err, "couldn't delete allocation pools from  virtual-router %s",
					vr.GetUUID())
			}

			if err = sv.validateVrouterAllocationPools(ctx, vr); err != nil {
				return errors.Wrapf(err, "couldn't add allocation pools to virtual-router %s",
					vr.GetUUID())
			}

			response, err = sv.BaseService.UpdateVirtualRouter(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) vrouterCheckAllocationPoolsDelete(
	ctx context.Context,
	vr *models.VirtualRouter,
	fm *protobuf.FieldMask,
	dbVr *models.VirtualRouter,
) error {
	if !basemodels.FieldMaskContains(fm, models.VirtualRouterFieldNetworkIpamRefs) {
		return nil
	}

	ipamRefs := vr.GetNetworkIpamRefs()
	if len(ipamRefs) == 0 {
		if len(dbVr.GetInstanceIPBackRefs()) != 0 {
			return errutil.ErrorBadRequestf("cannot delete allocation pool, ip address in use")
		}
	}

	existingVrPools := dbVr.GetAllocationPools()
	if len(existingVrPools) == 0 {
		return nil
	}

	toDeletePools := models.AllocationPoolsSubtract(existingVrPools, vr.GetAllocationPools())
	if len(toDeletePools) == 0 {
		return nil
	}

	return sv.vrouterCanDeleteAllocationPools(ctx, toDeletePools, dbVr)
}

// check if any ip address from given alloc_pool sets is used in
// in given virtual router, for instance_ip
func (sv *ContrailTypeLogicService) vrouterCanDeleteAllocationPools(
	_ context.Context,
	poolSet []*models.AllocationPoolType,
	dbVr *models.VirtualRouter,
) error {
	for _, iip := range dbVr.GetInstanceIPBackRefs() {
		if iip.GetInstanceIPAddress() == "" {
			continue
		}

		for _, allocPool := range poolSet {
			contains, err := allocPool.ContainsIPAddress(iip.GetInstanceIPAddress())
			if err != nil {
				return errutil.ErrorBadRequestf("cannot delete allocation pool: %v", err)
			}

			if contains {
				return errutil.ErrorBadRequestf("cannot delete allocation pool, %s in use", iip.GetInstanceIPAddress())
			}
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) validateVrouterAllocationPools(
	ctx context.Context,
	vr *models.VirtualRouter,
) error {
	if len(vr.GetNetworkIpamRefs()) == 0 {
		return nil
	}

	networkIpamsRes, err := sv.ReadService.ListNetworkIpam(ctx, &services.ListNetworkIpamRequest{
		Spec: &baseservices.ListSpec{
			ObjectUUIDs: vr.GetNetworkIpamRefUUIDs(),
			Detail:      true,
		},
	})

	if err != nil {
		return errutil.ErrorBadRequestf("error in dbe_list: %v", err)
	}

	refUUIDToNetworkIpamRefMap := vr.GetRefUUIDToNetworkIpamRefMap()
	for _, netIpam := range networkIpamsRes.GetNetworkIpams() {
		if netIpam.GetIpamSubnetMethod() != models.FlatSubnet {
			return errutil.ErrorBadRequestf(
				"only flat-subnet ipam can be attached to vrouter: NetworkIpam %s has subnet method %s",
				netIpam.GetUUID(),
				netIpam.GetIpamSubnetMethod(),
			)
		}
		netIpamRef := refUUIDToNetworkIpamRefMap[netIpam.GetUUID()]

		err = netIpamRef.Validate()
		if err != nil {
			return errutil.ErrorBadRequestf("failed to validate ref to network-ipam %s: %v", netIpamRef.GetUUID(), err)
		}

		err = netIpamRef.ValidateLinkToIpam(netIpam)
		if err != nil {
			return errutil.ErrorBadRequestf("failed to validate ref to network-ipam %s: %v", netIpamRef.GetUUID(), err)
		}

		if len(netIpam.GetVirtualRouterBackRefs()) == 0 {
			continue
		}

		// TODO: handle ipam back-refs to virtual routers
		//       validate if allocation-pools are not used in other vrouters
	}

	return nil
}
