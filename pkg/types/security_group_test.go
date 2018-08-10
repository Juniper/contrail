package types

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
)

func TestSecurityGroupLogicValidation(t *testing.T) {

	sg := models.MakeSecurityGroup()
	sg.UUID = "sg_uuid"

	sg_with_id := models.MakeSecurityGroup()
	sg_with_id.UUID = sg.UUID
	sg_with_id.SecurityGroupID = 8000001

	sg_with_id2 := models.MakeSecurityGroup()
	sg_with_id2.UUID = sg.UUID
	sg_with_id2.SecurityGroupID = 8000002

	t.Run("create security group", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sv := makeMockedContrailTypeLogicService(ctrl)

		sv.IntPoolAllocator.(*typesmock.MockIntPoolAllocator).
			EXPECT().AllocateInt(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
			Return(int64(1), nil).Times(1)

		sv.Next().(*servicesmock.MockService).
			EXPECT().CreateSecurityGroup(gomock.Not(gomock.Nil()), &services.CreateSecurityGroupRequest{
			SecurityGroup: sg_with_id,
		}).Return(&services.CreateSecurityGroupResponse{SecurityGroup: sg_with_id}, nil).Times(1)

		ctx := context.Background()
		res, err := sv.CreateSecurityGroup(ctx, &services.CreateSecurityGroupRequest{
			SecurityGroup: sg,
		})

		if assert.NoError(t, err) {
			assert.Equal(t, sg_with_id, res.SecurityGroup)
		}
	})

	t.Run("fail to create security group with explicit ID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sv := makeMockedContrailTypeLogicService(ctrl)

		ctx := context.Background()
		_, err := sv.CreateSecurityGroup(ctx, &services.CreateSecurityGroupRequest{
			SecurityGroup: sg_with_id,
		})

		assert.Error(t, err)
	})

	t.Run("update security group with explicit ID equal to the old ID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sv := makeMockedContrailTypeLogicService(ctrl)

		sv.ReadService.(*servicesmock.MockReadService).
			EXPECT().GetSecurityGroup(gomock.Not(gomock.Nil()), &services.GetSecurityGroupRequest{
			ID: sg_with_id.UUID,
		}).Return(&services.GetSecurityGroupResponse{SecurityGroup: sg_with_id}, nil).Times(1)

		sv.Next().(*servicesmock.MockService).
			EXPECT().UpdateSecurityGroup(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
			DoAndReturn(func(_ context.Context, request *services.UpdateSecurityGroupRequest) (
				*services.UpdateSecurityGroupResponse, error) {
				return &services.UpdateSecurityGroupResponse{
					SecurityGroup: request.SecurityGroup,
				}, nil
			}).Times(1)

		ctx := context.Background()
		res, err := sv.UpdateSecurityGroup(ctx, &services.UpdateSecurityGroupRequest{
			SecurityGroup: sg_with_id,
			FieldMask: types.FieldMask{
				Paths: []string{
					models.SecurityGroupFieldSecurityGroupID,
				},
			},
		})

		if assert.NoError(t, err) {
			assert.Equal(t, sg_with_id, res.SecurityGroup)
		}
	})

	t.Run("fail to update security group with explicit ID different than the old ID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sv := makeMockedContrailTypeLogicService(ctrl)

		sv.ReadService.(*servicesmock.MockReadService).
			EXPECT().GetSecurityGroup(gomock.Not(gomock.Nil()), &services.GetSecurityGroupRequest{
			ID: sg_with_id.UUID,
		}).Return(&services.GetSecurityGroupResponse{SecurityGroup: sg_with_id}, nil).Times(1)

		ctx := context.Background()
		_, err := sv.UpdateSecurityGroup(ctx, &services.UpdateSecurityGroupRequest{
			SecurityGroup: sg_with_id2,
			FieldMask: types.FieldMask{
				Paths: []string{
					models.SecurityGroupFieldSecurityGroupID,
				},
			},
		})

		assert.Error(t, err)
	})
}
