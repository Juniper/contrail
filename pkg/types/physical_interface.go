package types

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

//CreatePhysicalInterface does pre-check for create physical_interface
func (sv *ContrailTypeLogicService) CreatePhysicalInterface(
	ctx context.Context,
	request *services.CreatePhysicalInterfaceRequest,
) (response *services.CreatePhysicalInterfaceResponse, err error) {
	physicalInterface := request.GetPhysicalInterface()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			err = sv.validatePhysicalInterfaceDisplayName(ctx, physicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to validate display name")
			}
			err = physicalInterface.ValidateESIFormat()
			if err != nil {
				return errors.Wrapf(err, "failed to validate ESI format")
			}
			response, err = sv.BaseService.CreatePhysicalInterface(ctx, request)
			return err
		})
	return response, err
}

//UpdatePhysicalInterface does pre-check for update physical_interface
func (sv *ContrailTypeLogicService) UpdatePhysicalInterface(
	ctx context.Context,
	request *services.UpdatePhysicalInterfaceRequest,
) (response *services.UpdatePhysicalInterfaceResponse, err error) {
	physicalInterface := request.GetPhysicalInterface()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var storedPhysicalInterface *models.PhysicalInterface
			storedPhysicalInterface, err = sv.getPhysicalInterface(ctx, physicalInterface.UUID)
			if err != nil {
				return errors.Wrapf(err, "failed to get physical interface router")
			}
			err = validatePhysicalInterfaceUpdateDisplayName(physicalInterface, storedPhysicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to change display name")
			}
			err = physicalInterface.ValidateESIFormat()
			if err != nil {
				return errors.Wrapf(err, "failed to validate ESI name")
			}
			err = sv.validatePhysicalInterfaceESI(ctx, physicalInterface, storedPhysicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to validate ESI")
			}
			response, err = sv.BaseService.UpdatePhysicalInterface(ctx, request)
			return err
		})
	return response, err
}

func (sv *ContrailTypeLogicService) validatePhysicalInterfaceDisplayName(
	ctx context.Context,
	physicalInterface *models.PhysicalInterface,
) error {
	physicalRouterResponse, err := sv.ReadService.GetPhysicalRouter(ctx, &services.GetPhysicalRouterRequest{
		ID: physicalInterface.GetParentUUID(),
	})
	if err != nil {
		return errutil.ErrorInternalf("failed to get physical interface parent router with uuid %s",
			physicalInterface.GetParentUUID())
	}
	// the display name of the physical interface must be unique
	for _, pi := range physicalRouterResponse.GetPhysicalRouter().PhysicalInterfaces {
		if pi.DisplayName == physicalInterface.DisplayName {
			return errutil.ErrorConflictf("display name %s of physical interface already in use",
				physicalInterface.DisplayName)
		}
	}
	return nil
}

func validatePhysicalInterfaceUpdateDisplayName(
	physicalInterface *models.PhysicalInterface,
	storedPhysicalInterface *models.PhysicalInterface,
) error {
	if len(physicalInterface.DisplayName) != 0 &&
		physicalInterface.DisplayName != storedPhysicalInterface.DisplayName {
		return errutil.ErrorBadRequestf("cannot change display name from %s to %s",
			storedPhysicalInterface.DisplayName, physicalInterface.DisplayName)
	}
	return nil
}

func (sv *ContrailTypeLogicService) validatePhysicalInterfaceESI(
	ctx context.Context,
	physicalInterface *models.PhysicalInterface,
	storedPhysicalInterface *models.PhysicalInterface,
) error {
	if len(physicalInterface.EthernetSegmentIdentifier) == 0 {
		return nil
	}
	norm, err := sv.readPhysicalInterfaceESIStored(ctx, storedPhysicalInterface)
	if err != nil {
		return err
	}
	return sv.validatePhysicalInterfaceESIEqual(ctx, norm, physicalInterface)
}

func (sv *ContrailTypeLogicService) readPhysicalInterfaceESIStored(
	ctx context.Context,
	storedPhysicalInterface *models.PhysicalInterface,
) (map[int64][]*models.LogicalInterfaceVirtualMachineInterfaceRef, error) {
	norm := map[int64][]*models.LogicalInterfaceVirtualMachineInterfaceRef{}
	// logical interfaces associated with the physical interface
	for _, li := range storedPhysicalInterface.LogicalInterfaces {
		logicalInterface, err := sv.getLogicalInterface(ctx, li.UUID)
		if err != nil {
			return nil, err
		}
		if len(logicalInterface.VirtualMachineInterfaceRefs) != 0 {
			norm[logicalInterface.LogicalInterfaceVlanTag] = logicalInterface.VirtualMachineInterfaceRefs
		}
	}
	return norm, nil
}

func (sv *ContrailTypeLogicService) validatePhysicalInterfaceESIEqual(
	ctx context.Context,
	norm map[int64][]*models.LogicalInterfaceVirtualMachineInterfaceRef,
	physicalInterface *models.PhysicalInterface,
) error {
	// physical interfaces width same ESI
	physicalInterfaceList, err := sv.listPhysicalInterfacesByESI(ctx, physicalInterface.EthernetSegmentIdentifier)
	if err != nil {
		return err
	}
	for _, pi := range physicalInterfaceList {
		for _, li := range pi.LogicalInterfaces {
			// verify logical interface VMIs
			logicalInterface, err := sv.getLogicalInterface(ctx, li.UUID)
			if err != nil {
				return err
			}
			if len(logicalInterface.VirtualMachineInterfaceRefs) == 0 {
				continue
			}
			if !isEqualVMIRefs(norm[logicalInterface.LogicalInterfaceVlanTag],
				logicalInterface.VirtualMachineInterfaceRefs) {
				// return 403 code, implemented "as is"
				// https://github.com/Juniper/contrail-controller/blob/0e55b227581a7ab1f705734a5bd3e4360ad9a9e5/
				// src/config/api-server/vnc_cfg_api_server/vnc_cfg_types.py#L5539
				return errutil.ErrorForbiddenf("LI associated with the PI should have the same VMIs" +
					" as LIs (associated with the PIs) of the same ESI family")
			}
		}
	}
	return nil
}

func isEqualVMIRefs(
	left, right []*models.LogicalInterfaceVirtualMachineInterfaceRef,
) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i].UUID != right[i].UUID {
			return false
		}
	}
	return true
}

func (sv *ContrailTypeLogicService) getPhysicalInterface(
	ctx context.Context,
	uuid string,
) (*models.PhysicalInterface, error) {
	physicalInterfaceResponse, err := sv.ReadService.GetPhysicalInterface(ctx, &services.GetPhysicalInterfaceRequest{
		ID: uuid,
	})
	if err != nil {
		return nil, err
	}
	return physicalInterfaceResponse.GetPhysicalInterface(), nil
}

func (sv *ContrailTypeLogicService) getLogicalInterface(
	ctx context.Context,
	uuid string,
) (*models.LogicalInterface, error) {
	logicalInterfaceResponse, err := sv.ReadService.GetLogicalInterface(ctx, &services.GetLogicalInterfaceRequest{
		ID: uuid,
	})
	if err != nil {
		return nil, err
	}
	return logicalInterfaceResponse.GetLogicalInterface(), nil
}

func (sv *ContrailTypeLogicService) listPhysicalInterfacesByESI(
	ctx context.Context,
	esi string,
) ([]*models.PhysicalInterface, error) {
	response, err := sv.ReadService.ListPhysicalInterface(ctx, &services.ListPhysicalInterfaceRequest{
		Spec: &baseservices.ListSpec{
			Detail: true,
			Filters: []*baseservices.Filter{{
				Key:    models.PhysicalInterfaceFieldEthernetSegmentIdentifier,
				Values: []string{esi},
			}},
		},
	})
	if err != nil {
		return nil, err
	}
	return response.GetPhysicalInterfaces(), nil
}
