package types

import (
	"context"
	"regexp"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

//CreatePhysicalInterface does pre-check for create physical_interface
func (sv *ContrailTypeLogicService) CreatePhysicalInterface(
	ctx context.Context,
	request *services.CreatePhysicalInterfaceRequest,
) (responce *services.CreatePhysicalInterfaceResponse, err error) {
	physicalInterface := request.GetPhysicalInterface()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			err = sv.validatePhysicalInterfaceDisplayName(ctx, physicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to validate display name")
			}
			err = validatePhysicalInterfaceESIFormat(physicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to test ESI format")
			}
			responce, err = sv.BaseService.CreatePhysicalInterface(ctx, request)
			return err
		})
	return responce, err
}

//UpdatePhysicalInterface does pre-check for update physical_interface
func (sv *ContrailTypeLogicService) UpdatePhysicalInterface(
	ctx context.Context,
	request *services.UpdatePhysicalInterfaceRequest,
) (responce *services.UpdatePhysicalInterfaceResponse, err error) {
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
			err = validatePhysicalInterfaceESIFormat(physicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to validate ESI name")
			}
			err = sv.validatePhysicalInterfaceESI(ctx, physicalInterface, storedPhysicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to validate ESI")
			}
			responce, err = sv.BaseService.UpdatePhysicalInterface(ctx, request)
			return err
		})
	return responce, err
}

func (sv *ContrailTypeLogicService) validatePhysicalInterfaceDisplayName(
	ctx context.Context,
	physicalInterface *models.PhysicalInterface,
) error {
	physicalRouterResponse, err := sv.ReadService.GetPhysicalRouter(ctx, &services.GetPhysicalRouterRequest{
		ID: physicalInterface.GetParentUUID(),
	})
	if err != nil {
		return errors.Wrapf(err, "failed to get physical interface router")
	}
	// the display name of the physical interface must be unique
	for _, pi := range physicalRouterResponse.GetPhysicalRouter().PhysicalInterfaces {
		if pi.DisplayName == physicalInterface.DisplayName {
			return errutil.ErrorConflictf("display name %s of physical interface already in use", physicalInterface.DisplayName)
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

var regexpESIFormat = regexp.MustCompile("^([0-9A-Fa-f]{2}[:]){9}[0-9A-Fa-f]{2}")

func validatePhysicalInterfaceESIFormat(
	physicalInterface *models.PhysicalInterface,
) error {
	if len(physicalInterface.EthernetSegmentIdentifier) != 0 &&
		!regexpESIFormat.MatchString(physicalInterface.EthernetSegmentIdentifier) {
		return errutil.ErrorBadRequestf("invalid esi string format %s", physicalInterface.EthernetSegmentIdentifier)
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
	norm := map[int64][]*models.LogicalInterfaceVirtualMachineInterfaceRef{}
	// logical interfaces associated with the physical interface
	for _, li := range storedPhysicalInterface.LogicalInterfaces {
		lInterface, err := sv.getLogicalInterface(ctx, li.UUID)
		if err != nil {
			return err
		}
		if len(lInterface.VirtualMachineInterfaceRefs) != 0 {
			norm[lInterface.LogicalInterfaceVlanTag] = lInterface.VirtualMachineInterfaceRefs
		}
	}
	// physical interfaces width same ESI
	physicalInterfaceList, err := sv.listPhysicalInterfacesByESI(ctx, physicalInterface.EthernetSegmentIdentifier)
	if err != nil {
		return err
	}
	for _, pi := range physicalInterfaceList {
		for _, li := range pi.LogicalInterfaces {
			// verify logical interface VMIs
			lInterface, err := sv.getLogicalInterface(ctx, li.UUID)
			if err != nil {
				return err
			}
			if len(lInterface.VirtualMachineInterfaceRefs) != 0 &&
				!isEqualVMIRefs(norm[lInterface.LogicalInterfaceVlanTag], lInterface.VirtualMachineInterfaceRefs) {
				return errutil.ErrorConflictf(
					"LI associated with the PI should have the same VMIs" +
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
