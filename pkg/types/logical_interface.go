package types

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

//CreateLogicalInterface does pre-check for create logical_interface
func (sv *ContrailTypeLogicService) CreateLogicalInterface(
	ctx context.Context,
	request *services.CreateLogicalInterfaceRequest,
) (response *services.CreateLogicalInterfaceResponse, err error) {
	logicalInterface := request.GetLogicalInterface()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var physicalRouter *models.PhysicalRouter
			var physicalInterface *models.PhysicalInterface
			physicalRouter, physicalInterface, err = sv.getLogicalInterfaceParents(ctx, logicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to get parent objects")
			}
			err = validateLogicalInterfaceVLanTag(logicalInterface, physicalRouter)
			if err != nil {
				return errors.Wrapf(err, "failed to check logical interface vlan id")
			}
			err = validateLogicalInterfaceDisplayName(logicalInterface, physicalInterface, physicalRouter)
			if err != nil {
				return errors.Wrapf(err, "failed to validate display name")
			}
			err = sv.validateLogicalInterfaceESI(ctx, logicalInterface, physicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to validate logical interface ESI")
			}
			response, err = sv.BaseService.CreateLogicalInterface(ctx, request)
			return err
		})
	return response, err
}

//UpdateLogicalInterface does pre-check for update logical_interface
func (sv *ContrailTypeLogicService) UpdateLogicalInterface(
	ctx context.Context,
	request *services.UpdateLogicalInterfaceRequest,
) (response *services.UpdateLogicalInterfaceResponse, err error) {
	logicalInterface := request.GetLogicalInterface()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var storedLogicalInterface *models.LogicalInterface
			storedLogicalInterface, err = sv.getLogicalInterface(ctx, logicalInterface.GetUUID())
			if err != nil {
				return errors.Wrapf(err, "failed to get logical interface with uuid %s", logicalInterface.GetUUID())
			}
			var physicalRouter *models.PhysicalRouter
			var physicalInterface *models.PhysicalInterface
			physicalRouter, physicalInterface, err = sv.getLogicalInterfaceParents(ctx, storedLogicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to get parent objects")
			}
			err = validateLogicalInterfaceUpdateDisplayName(logicalInterface, storedLogicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to validate logical interface display name")
			}
			err = validateLogicalInterfaceDisplayName(logicalInterface, physicalInterface, physicalRouter)
			if err != nil {
				return errors.Wrapf(err, "failed to validate logical interface display name")
			}
			err = validateLogicalInterfaceVLanTag(logicalInterface, physicalRouter)
			if err != nil {
				return errors.Wrapf(err, "failed to validate logical interface vlan id")
			}
			err = validateLogicalInterfaceUpdateVLanTag(logicalInterface, storedLogicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to validate logical interface vlan id")
			}
			err = sv.validateLogicalInterfaceESI(ctx, logicalInterface, physicalInterface)
			if err != nil {
				return err
			}
			response, err = sv.BaseService.UpdateLogicalInterface(ctx, request)
			return err
		})
	return response, err
}

func validateLogicalInterfaceVLanTag(
	logicalInterface *models.LogicalInterface,
	physicalRouter *models.PhysicalRouter,
) error {
	if logicalInterface.GetLogicalInterfaceVlanTag() < 0 || logicalInterface.GetLogicalInterfaceVlanTag() > 4094 {
		return errutil.ErrorForbiddenf("invalid logical interface vlan id %d", logicalInterface.GetLogicalInterfaceVlanTag())
	}
	// In case of QFX, check that VLANs 1, 2 and 4094 are not used
	if strings.HasPrefix(strings.ToLower(physicalRouter.GetPhysicalRouterProductName()), "qfx") &&
		logicalInterface.IsReservedQFXVlanTag() {
		return errutil.ErrorBadRequestf("vlan id %d is not allowed on QFX logical interface type %s",
			logicalInterface.GetLogicalInterfaceVlanTag(), logicalInterface.GetLogicalInterfaceType())
	}
	return nil
}

func validateLogicalInterfaceUpdateVLanTag(
	logicalInterface *models.LogicalInterface,
	storedLogicalInterface *models.LogicalInterface,
) error {
	if logicalInterfaceVlanTagExists(logicalInterface) &&
		logicalInterfaceVlanTagExists(storedLogicalInterface) &&
		logicalInterface.GetLogicalInterfaceVlanTag() != storedLogicalInterface.GetLogicalInterfaceVlanTag() {
		return errutil.ErrorForbiddenf("cannot change vlan tag %d", logicalInterface.GetLogicalInterfaceVlanTag())
	}
	return nil
}

func (sv *ContrailTypeLogicService) getLogicalInterfaceParents(
	ctx context.Context,
	logicalInterface *models.LogicalInterface,
) (physicalRouter *models.PhysicalRouter, physicalInterface *models.PhysicalInterface, err error) {
	physicalRouterUUID := logicalInterface.GetParentUUID()
	if logicalInterface.ParentType == "physical-interface" {
		physicalInterface, err = sv.getPhysicalInterface(ctx, logicalInterface.GetParentUUID())
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to get logical interface physical interface with uuid %s",
				logicalInterface.GetParentUUID())
		}
		physicalRouterUUID = physicalInterface.GetParentUUID()
	}
	physicalRouter, err = sv.getPhysicalRouter(ctx, physicalRouterUUID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to get logical interface physical router width uuid %s",
			physicalRouterUUID)
	}
	return physicalRouter, physicalInterface, nil
}

func validateLogicalInterfaceDisplayName(
	logicalInterface *models.LogicalInterface,
	physicalInterface *models.PhysicalInterface,
	physicalRouter *models.PhysicalRouter,
) error {
	for _, pi := range physicalRouter.GetPhysicalInterfaces() {
		if pi.GetDisplayName() == logicalInterface.GetDisplayName() {
			return errutil.ErrorConflictf("display name %s of logical interface already in use",
				logicalInterface.GetDisplayName())
		}
	}
	if physicalInterface == nil {
		return nil
	}
	for _, li := range physicalInterface.GetLogicalInterfaces() {
		if li.GetUUID() != logicalInterface.GetUUID() &&
			logicalInterfaceVlanTagExists(li) &&
			li.GetLogicalInterfaceVlanTag() == logicalInterface.GetLogicalInterfaceVlanTag() {
			return errutil.ErrorConflictf("vlan id of logical interface %d already used"+
				"in another interface %s", logicalInterface.GetLogicalInterfaceVlanTag(), li.GetUUID())
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) validateLogicalInterfaceESI(
	ctx context.Context,
	logicalInterface *models.LogicalInterface,
	physicalInterface *models.PhysicalInterface,
) error {
	if physicalInterface == nil || len(physicalInterface.GetEthernetSegmentIdentifier()) == 0 {
		return nil
	}
	physicalInterfaceList, err := sv.listPhysicalInterfacesByESI(ctx, physicalInterface.GetEthernetSegmentIdentifier())
	if err != nil {
		return err
	}
	for _, pi := range physicalInterfaceList {
		for _, li := range pi.GetLogicalInterfaces() {
			if li.GetUUID() == logicalInterface.GetUUID() {
				continue
			}
			otherLogicalInterface, err := sv.getLogicalInterface(ctx, li.GetUUID())
			if err != nil {
				return err
			}
			if otherLogicalInterface.GetLogicalInterfaceVlanTag() != logicalInterface.GetLogicalInterfaceVlanTag() {
				continue
			}
			if !isEqualVMIRefs(logicalInterface.GetVirtualMachineInterfaceRefs(),
				otherLogicalInterface.GetVirtualMachineInterfaceRefs()) {
				return errutil.ErrorForbiddenf(
					"LI should refer to the same set of VMIs as peer LIs belonging to the same ESI")
			}
		}
	}
	return nil
}

func validateLogicalInterfaceUpdateDisplayName(
	logicalInterface *models.LogicalInterface,
	storedLogicalInterface *models.LogicalInterface,
) error {
	if len(logicalInterface.GetDisplayName()) > 0 &&
		logicalInterface.GetDisplayName() != storedLogicalInterface.GetDisplayName() {
		return errutil.ErrorForbiddenf("cannot change display name to %s", logicalInterface.GetDisplayName())
	}
	return nil
}

func logicalInterfaceVlanTagExists(
	logicalInterface *models.LogicalInterface,
) bool {
	return logicalInterface.GetLogicalInterfaceVlanTag() != 0
}

func (sv *ContrailTypeLogicService) getPhysicalRouter(
	ctx context.Context,
	uuid string,
) (*models.PhysicalRouter, error) {
	physicalRouterResponse, err := sv.ReadService.GetPhysicalRouter(ctx, &services.GetPhysicalRouterRequest{
		ID: uuid,
	})
	if err != nil {
		return nil, err
	}
	return physicalRouterResponse.GetPhysicalRouter(), nil
}
