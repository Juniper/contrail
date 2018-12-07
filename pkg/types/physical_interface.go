package types

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/services/baseservices"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/Juniper/contrail/pkg/models"

	//"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	//"github.com/Juniper/contrail/pkg/types/ipam"
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
			physicalRouter, err := sv.getPhysicalRouterFromPhysicalInterface(ctx, physicalInterface)
			if err != nil {
				return err
			}
			err = sv.validatePhysicalInterfaceName(ctx, physicalInterface, physicalRouter)
			if err != nil {
				return grpc.Errorf(codes.AlreadyExists, "failed to validate interface name: %v", err)
			}
			if len(physicalInterface.EthernetSegmentIdentifier) != 0 {
				err = validatePhysicalInterfaceESIFormat(physicalInterface)
				if err != nil {
					return errutil.ErrorBadRequestf("failed to test ESI: %v", err)
				}
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
			storedPhysicalInterface, err := sv.getPhysicalInterface(ctx, physicalInterface.UUID)
			if err != nil {
				return err
			}
			if physicalInterface.DisplayName != storedPhysicalInterface.DisplayName {
				return errors.New("cannot change display name")
			}
			if len(physicalInterface.EthernetSegmentIdentifier) != 0 {
				err = validatePhysicalInterfaceESIFormat(physicalInterface)
				if err != nil {
					return err
				}
				err = sv.validatePhysicalInterfaceESI(ctx, physicalInterface)
				if err != nil {
					return err
				}
			}
			responce, err = sv.BaseService.UpdatePhysicalInterface(ctx, request)
			return err
		})
	return responce, err
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

func (sv *ContrailTypeLogicService) getPhysicalRouterFromPhysicalInterface(
	ctx context.Context,
	physicalInterface *models.PhysicalInterface,
) (*models.PhysicalRouter, error) {
	physicalRouterResponse, err := sv.ReadService.GetPhysicalRouter(ctx, &services.GetPhysicalRouterRequest{
		ID: physicalInterface.GetParentUUID(),
	})
	if err != nil {
		return nil, err
	}
	return physicalRouterResponse.GetPhysicalRouter(), nil
}

func (sv *ContrailTypeLogicService) listPhysicalInterfacesByESI(
	ctx context.Context,
	esi string,
) ([]*models.PhysicalInterface, error) {
	response, err := sv.ReadService.ListPhysicalInterface(ctx, &services.ListPhysicalInterfaceRequest{
		Spec: &baseservices.ListSpec{
			Detail: true,
			Filters: []*baseservices.Filter{
				{
					Key:    models.PhysicalInterfaceFieldEthernetSegmentIdentifier,
					Values: []string{esi},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return response.GetPhysicalInterfaces(), nil
}

/*
func (sv *ContrailTypeLogicService) listLogicalInterfacesByPhysicalInterface(
	ctx context.Context,
	physicalInterface *models.PhysicalInterface,
) ([]*models.LogicalInterface, error) {
	response, err := sv.ReadService.ListLogicalInterface(ctx, &services.ListLogicalInterfaceRequest{
		Spec: &baseservices.ListSpec{
			Detail: true,
			Filters: []*baseservices.Filter{
								{
								Key:    models.LogicalInterfaceFieldParentType
								Values: []string{"physical-interface"},
							},
				{
					Key:    models.LogicalInterfaceFieldParentUUID,
					Values: []string{physicalInterface.UUID},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return response.GetLogicalInterfaces(), nil
}
*/

func (sv *ContrailTypeLogicService) validatePhysicalInterfaceName(
	ctx context.Context,
	physicalInterface *models.PhysicalInterface,
	physicalRouter *models.PhysicalRouter,
) error {
	for _, pi := range physicalRouter.PhysicalInterfaces {
		pInterface, err := sv.getPhysicalInterface(ctx, pi.UUID)
		if err != nil {
			return err
		}
		// the display name of the physical interface must be unique
		if pInterface.DisplayName == physicalInterface.DisplayName {
			return grpc.Errorf(codes.AlreadyExists, "display name %s of physical interface already in use", physicalInterface.DisplayName)
		}
		// check vlan tags on the physical interface
		for _, tag := range pInterface.TagRefs {
			err = validatePhysicalInterfaceTags(physicalInterface, tag.UUID)
			if err != nil {
				return err
			}
		}
	}
	for _, li := range physicalRouter.LogicalInterfaces {
		//fmt.Println("Step OK")
		lInterface, err := sv.getLogicalInterface(ctx, li.UUID)
		if err != nil {
			return err
		}
		// check vlan tags on the logical interface
		for _, tag := range lInterface.TagRefs {
			err = validatePhysicalInterfaceTags(physicalInterface, tag.UUID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validatePhysicalInterfaceTags(
	physicalInterface *models.PhysicalInterface,
	uuid string,
) error {
	for _, tag := range physicalInterface.TagRefs {
		//fmt.Println("+++", tag.UUID, uuid)
		if tag.UUID == uuid {
			return errors.New("Vlan tag " + tag.UUID + " already used in another interface ")
		}
	}
	return nil
}

func validatePhysicalInterfaceESIFormat(
	physicalInterface *models.PhysicalInterface,
) error {
	re := regexp.MustCompile("^([0-9A-Fa-f]{2}[:]){9}[0-9A-Fa-f]{2}")
	ok := re.MatchString(physicalInterface.EthernetSegmentIdentifier)
	if !ok {
		return errors.New("invalid esi string format")
	}
	return nil
}

func (sv *ContrailTypeLogicService) validatePhysicalInterfaceESI(
	ctx context.Context,
	physicalInterface *models.PhysicalInterface,
) error {
	norm := map[int64][]*models.LogicalInterfaceVirtualMachineInterfaceRef{}
	// logical interfaces associated with physical interface
	//fmt.Println("Etalon start")
	for _, li := range physicalInterface.LogicalInterfaces {
		fmt.Println("Etalon step")
		lInterface, err := sv.getLogicalInterface(ctx, li.UUID)
		if err != nil {
			return err
		}
		norm[lInterface.LogicalInterfaceVlanTag] = lInterface.VirtualMachineInterfaceRefs
		//fmt.Println("Etalon", lInterface.LogicalInterfaceVlanTag, lInterface.VirtualMachineInterfaceRefs)
	}

	// physical interfaces width same ESI
	physicalInterfaceList, err := sv.listPhysicalInterfacesByESI(ctx, physicalInterface.EthernetSegmentIdentifier)
	if err != nil {
		return err
	}
	//fmt.Println("ESI", physicalInterface.EthernetSegmentIdentifier)
	for _, pi := range physicalInterfaceList {
		//fmt.Println("UUID", pi.UUID)
		for _, li := range pi.LogicalInterfaces {
			// verify logical interface VMIs
			lInterface, err := sv.getLogicalInterface(ctx, li.UUID)
			if err != nil {
				return err
			}
			//fmt.Println("Compare", lInterface.LogicalInterfaceVlanTag, lInterface.VirtualMachineInterfaceRefs)
			if !isEqualVMIRefs(norm[lInterface.LogicalInterfaceVlanTag], lInterface.VirtualMachineInterfaceRefs) {
				return errors.New("LI associated with the PI should have the same VMIs as LIs (associated with the PIs) of the same ESI family")
			}
		}
	}
	return nil
}

func isEqualVMIRefs(left, right []*models.LogicalInterfaceVirtualMachineInterfaceRef) bool {
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
