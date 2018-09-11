package types

import (
	"context"

	protobuf "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

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
			return common.ErrorBadRequestf("cannot delete allocation pool, ip address in use")
		}
	}

	existingVrPools := vr.GetAllocationPools()
	if len(existingVrPools) == 0 {
		return nil
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
		},
	})

	if err != nil {
		return common.ErrorBadRequestf("error in dbe_list: %v", err)
	}

	refUUIDToNetworkIpamRefMap := vr.GetRefUUIDToNetworkIpamRefMap()
	for _, netIpam := range networkIpamsRes.GetNetworkIpams() {
		if netIpam.GetIpamSubnetMethod() != models.FlatSubnet {
			return errors.Errorf(
				"only flat-subnet ipam can be attached to vrouter: NetworkIpam %s has subnet method %s",
				netIpam.GetUUID(), netIpam.GetIpamSubnetMethod())
		}
		netIpamRef := refUUIDToNetworkIpamRefMap[netIpam.GetUUID()]

		// read data on the link between vrouter and ipam
		// if alloc pool exists, then make sure that alloc-pools are
		// configured in ipam subnet with a flag indicating
		// vrouter specific allocation pool
		vrIpamData := netIpamRef.GetAttr()
		if len(vrIpamData.GetAllocationPools()) == 0 {
			return common.ErrorBadRequestf("no allocation-pools for this vrouter")
		}

		vrSubnets := vrIpamData.GetSubnet()
		for _, vrSubnet := range vrSubnets {
			if err := vrSubnet.Validate(); err != nil {
				return common.ErrorBadRequestf("couldn't validate vrouter subnet: %v", err)
			}
		}

		vrAllocPools := netIpamRef.GetVrouterSpecificAllocationPools()
		for _, vrAllocPool := range vrAllocPools {
			if !netIpam.ContainsAllocationPool(vrAllocPool) {
				return common.ErrorBadRequestf("vrouter allocation-pool start:%s, end %s not in ipam %s",
					vrAllocPool.GetStart(), vrAllocPool.GetEnd(), netIpam.GetUUID())
			}
		}

		if len(netIpam.GetVirtualRouterBackRefs()) == 0 {
			continue
		}

		// TODO: handle ipam back-refs to virtual routers
		//       validate if allocation-pools are not used in other vrouters

	}

	//TODO: Validate vrouter allocation pools

	return nil
}

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

			dbVr := dbVrRes.GetVirtualRouter()
			if err := sv.vrouterCheckAllocationPoolsDelete(ctx, vr, &fm, dbVr); err != nil {
				return errors.Wrapf(err, "couldn't delete allocation pools from  virtual-router %s",
					vr.GetUUID())
			}

			if err := sv.validateVrouterAllocationPools(ctx, vr); err != nil {
				return errors.Wrapf(err, "couldn't add allocation pools validations to virtual-router %s",
					vr.GetUUID())
			}

			response, err = sv.BaseService.UpdateVirtualRouter(ctx, request)
			return err
		})

	return response, err
}

// DeleteVirtualRouter virtual-router delete specific logic.
func (sv *ContrailTypeLogicService) DeleteVirtualRouter(
	ctx context.Context, request *services.DeleteVirtualRouterRequest,
) (*services.DeleteVirtualRouterResponse, error) {
	return sv.BaseService.DeleteVirtualRouter(ctx, request)
}
