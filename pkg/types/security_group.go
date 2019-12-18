package types

import (
	"context"

	protobuf "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateSecurityGroup performs type specific validation and setup for creating security groups.
func (sv *ContrailTypeLogicService) CreateSecurityGroup(
	ctx context.Context, request *services.CreateSecurityGroupRequest) (
	response *services.CreateSecurityGroupResponse, err error) {

	sg := request.SecurityGroup

	if err = sv.disallowManualSecurityGroupID(nil, sg, nil); err != nil {
		return nil, err
	}

	err = sv.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		err = sg.GetSecurityGroupEntries().CheckSecurityGroupRules()
		if err != nil {
			return errors.Wrapf(err, "failed to check Policy Rules")
		}
		sg.SecurityGroupEntries.FillRuleUUIDs()
		// TODO: handle configured security group ID
		if sg.SecurityGroupID, err = sv.allocateSecurityGroupID(ctx); err != nil {
			return err
		}

		if response, err = sv.BaseService.CreateSecurityGroup(ctx, request); err != nil {
			return err
		}

		return nil
	})

	return response, err
}

// UpdateSecurityGroup performs type specific validation and setup for updating security groups.
func (sv *ContrailTypeLogicService) UpdateSecurityGroup(
	ctx context.Context, request *services.UpdateSecurityGroupRequest) (
	response *services.UpdateSecurityGroupResponse, err error) {

	sg := request.SecurityGroup

	err = sv.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		var getResponse *services.GetSecurityGroupResponse
		if getResponse, err = sv.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
			ID: sg.UUID,
		}); err != nil {
			return err
		}
		current := getResponse.GetSecurityGroup()

		if err = sv.disallowManualSecurityGroupID(current, sg, &request.FieldMask); err != nil {
			return err
		}

		err = sg.GetSecurityGroupEntries().CheckSecurityGroupRules()
		if err != nil {
			return errors.Wrapf(err, "failed to check Policy Rules")
		}
		sg.SecurityGroupEntries.FillRuleUUIDs()

		response, err = sv.BaseService.UpdateSecurityGroup(ctx, request)
		return err
	})

	return response, err
}

// DeleteSecurityGroup performs type specific validation and teardown for deleting security groups.
func (sv *ContrailTypeLogicService) DeleteSecurityGroup(
	ctx context.Context, request *services.DeleteSecurityGroupRequest) (
	response *services.DeleteSecurityGroupResponse, err error) {

	uuid := request.ID

	err = sv.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		var getResponse *services.GetSecurityGroupResponse
		if getResponse, err = sv.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
			ID: uuid,
		}); err != nil {
			return err
		}

		sg := getResponse.GetSecurityGroup()

		// TODO: handle configured security group ID
		if err = sv.deallocateSecurityGroupID(ctx, sg.SecurityGroupID); err != nil {
			return err
		}

		response, err = sv.BaseService.DeleteSecurityGroup(ctx, request)
		return err
	})

	return response, err
}

func (sv *ContrailTypeLogicService) allocateSecurityGroupID(ctx context.Context) (int64, error) {
	id, err := sv.IntPoolAllocator.AllocateInt(ctx, SecurityGroupIDPoolKey, db.EmptyIntOwner)
	if err != nil {
		return 0, err
	}
	return id + models.SecurityGroupIdTypeMinimum, nil
}

func (sv *ContrailTypeLogicService) deallocateSecurityGroupID(ctx context.Context, id int64) error {
	return sv.IntPoolAllocator.DeallocateInt(ctx, SecurityGroupIDPoolKey, id-models.SecurityGroupIdTypeMinimum)
}

func (sv *ContrailTypeLogicService) disallowManualSecurityGroupID(
	current *models.SecurityGroup, requested *models.SecurityGroup, fieldMask *protobuf.FieldMask) error {
	if current == nil {
		if requested.SecurityGroupID == 0 {
			return nil
		}

		return errutil.ErrorForbidden("cannot set the security group ID, it's allocated by the server")
	}

	if !basemodels.FieldMaskContains(fieldMask, models.SecurityGroupFieldSecurityGroupID) {
		return nil
	}

	if current.SecurityGroupID == requested.SecurityGroupID {
		return nil
	}

	return errutil.ErrorForbidden("cannot update the security group ID, it's allocated by the server")
}
