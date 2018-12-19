package types

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/errutil"
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
			var physicalInterface *models.PhysicalInterface
			physicalInterface, err = sv.getPhysicalInterfaceFromLogicalInterface(ctx, logicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to get parent physical interface")
			}
			var physicalRouter *models.PhysicalRouter
			physicalRouter, err = sv.getPhysicalRouterFromLogicalInterface(ctx, logicalInterface, physicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to get parent physical router")
			}
			err = validateLogicalInterfaceVLanTag(logicalInterface, physicalRouter)
			if err != nil {
				return errors.Wrapf(err, "failed to check logical interface vlan tag")
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
			storedLogicalInterface, err = sv.getLogicalInterface(ctx, logicalInterface.UUID)
			if err != nil {
				return errors.Wrapf(err, "failed to get logical interface with uuid %s", logicalInterface.UUID)
			}
			var physicalInterface *models.PhysicalInterface
			physicalInterface, err = sv.getPhysicalInterfaceFromLogicalInterface(ctx, storedLogicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to get parent physical interface")
			}
			var physicalRouter *models.PhysicalRouter
			physicalRouter, err = sv.getPhysicalRouterFromLogicalInterface(ctx, storedLogicalInterface, physicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to get parent physical router")
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
				return errors.Wrapf(err, "failed to validate logical interface vlan tag")
			}
			err = validateLogicalInterfaceUpdateVLanTag(logicalInterface, storedLogicalInterface)
			if err != nil {
				return errors.Wrapf(err, "failed to validate logical interface vlan tag")
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
	if logicalInterfaceVlanTagExists(logicalInterface) &&
		logicalInterface.LogicalInterfaceVlanTag < 0 || logicalInterface.LogicalInterfaceVlanTag > 4094 {
		return errutil.ErrorForbiddenf("invalig logical interface vlan tag %d", logicalInterface.LogicalInterfaceVlanTag)
	}
	// In case of QFX, check that VLANs 1, 2 and 4094 are not used
	if strings.HasPrefix(strings.ToLower(physicalRouter.PhysicalRouterProductName), "qfx") &&
		strings.ToLower(logicalInterface.LogicalInterfaceType) == "l2" && (logicalInterface.LogicalInterfaceVlanTag == 1 ||
		logicalInterface.LogicalInterfaceVlanTag == 2 ||
		logicalInterface.LogicalInterfaceVlanTag == 4094) {
		return errutil.ErrorBadRequestf("vlan tag %d is not allowed on QFX logical interface type %s",
			logicalInterface.LogicalInterfaceVlanTag, logicalInterface.LogicalInterfaceType)
	}
	return nil
}

func validateLogicalInterfaceUpdateVLanTag(
	logicalInterface *models.LogicalInterface,
	storedLogicalInterface *models.LogicalInterface,
) error {
	if logicalInterfaceVlanTagExists(logicalInterface) &&
		logicalInterfaceVlanTagExists(storedLogicalInterface) &&
		logicalInterface.LogicalInterfaceVlanTag != storedLogicalInterface.LogicalInterfaceVlanTag {
		return errutil.ErrorForbiddenf("cannot change vlan tag %d", logicalInterface.LogicalInterfaceVlanTag)
	}
	return nil
}

func (sv *ContrailTypeLogicService) getPhysicalInterfaceFromLogicalInterface(
	ctx context.Context,
	logicalInterface *models.LogicalInterface,
) (*models.PhysicalInterface, error) {
	if logicalInterface.ParentType != "physical-interface" {
		// no errors
		return nil, nil
	}
	return sv.getPhysicalInterface(ctx, logicalInterface.ParentUUID)
}

func (sv *ContrailTypeLogicService) getPhysicalRouterFromLogicalInterface(
	ctx context.Context,
	logicalInterface *models.LogicalInterface,
	physicalInterface *models.PhysicalInterface,
) (physicalRouter *models.PhysicalRouter, err error) {
	if logicalInterface.ParentType == "physical-interface" {
		if physicalInterface == nil {
			return nil, errutil.ErrorBadRequestf("failed to get physical router from logical interface with uuid %s",
				logicalInterface.UUID)
		}
		return sv.getPhysicalRouter(ctx, physicalInterface.ParentUUID)
	}
	return sv.getPhysicalRouter(ctx, logicalInterface.ParentUUID)
}

func validateLogicalInterfaceDisplayName(
	logicalInterface *models.LogicalInterface,
	physicalInterface *models.PhysicalInterface,
	physicalRouter *models.PhysicalRouter,
) error {
	for _, pi := range physicalRouter.PhysicalInterfaces {
		if pi.DisplayName == logicalInterface.DisplayName {
			return errutil.ErrorConflictf("display name %s of logical interface already in use", logicalInterface.DisplayName)
		}
	}
	if physicalInterface == nil {
		return nil
	}
	for _, li := range physicalInterface.LogicalInterfaces {
		if li.UUID != logicalInterface.UUID &&
			logicalInterfaceVlanTagExists(li) &&
			li.LogicalInterfaceVlanTag == logicalInterface.LogicalInterfaceVlanTag {
			return errutil.ErrorConflictf("vlan tag of logical interface %d already used"+
				"in another interface %s", logicalInterface.LogicalInterfaceVlanTag, li.UUID)
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) validateLogicalInterfaceESI(
	ctx context.Context,
	logicalInterface *models.LogicalInterface,
	physicalInterface *models.PhysicalInterface,
) error {
	if physicalInterface == nil || len(physicalInterface.EthernetSegmentIdentifier) == 0 {
		return nil
	}
	physicalInterfaceList, err := sv.listPhysicalInterfacesByESI(ctx, physicalInterface.EthernetSegmentIdentifier)
	if err != nil {
		return err
	}
	for _, pi := range physicalInterfaceList {
		for _, li := range pi.LogicalInterfaces {
			if li.UUID == logicalInterface.UUID {
				continue
			}
			otherLogicalInterface, err := sv.getLogicalInterface(ctx, li.UUID)
			if err != nil {
				return err
			}
			if otherLogicalInterface.LogicalInterfaceVlanTag != logicalInterface.LogicalInterfaceVlanTag {
				continue
			}
			if !isEqualVMIRefs(logicalInterface.VirtualMachineInterfaceRefs,
				otherLogicalInterface.VirtualMachineInterfaceRefs) {
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
	if len(logicalInterface.DisplayName) != 0 &&
		logicalInterface.DisplayName != storedLogicalInterface.DisplayName {
		return errutil.ErrorForbiddenf("cannot change display name to %s", logicalInterface.DisplayName)
	}
	return nil
}

func logicalInterfaceVlanTagExists(
	logicalInterface *models.LogicalInterface,
) bool {
	return logicalInterface.LogicalInterfaceVlanTag != 0
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
	return physicalRouterResponse.PhysicalRouter, nil
}
