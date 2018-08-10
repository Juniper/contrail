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

	sgWithID := models.MakeSecurityGroup()
	sgWithID.UUID = sg.UUID
	sgWithID.SecurityGroupID = 8000001

	sgWithID2 := models.MakeSecurityGroup()
	sgWithID2.UUID = sg.UUID
	sgWithID2.SecurityGroupID = 8000002

	runTest(t, "create security group",
		func(t *testing.T, sv *ContrailTypeLogicService) {
			sv.IntPoolAllocator.(*typesmock.MockIntPoolAllocator).
				EXPECT().AllocateInt(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
				Return(int64(1), nil).Times(1)

			sv.Next().(*servicesmock.MockService).
				EXPECT().CreateSecurityGroup(gomock.Not(gomock.Nil()), &services.CreateSecurityGroupRequest{
				SecurityGroup: sgWithID,
			}).Return(&services.CreateSecurityGroupResponse{SecurityGroup: sgWithID}, nil).Times(1)

			ctx := context.Background()
			res, err := sv.CreateSecurityGroup(ctx, &services.CreateSecurityGroupRequest{
				SecurityGroup: sg,
			})

			if assert.NoError(t, err) {
				assert.Equal(t, sgWithID, res.SecurityGroup)
			}
		})

	runTest(t, "fail to create security group with explicit ID",
		func(t *testing.T, sv *ContrailTypeLogicService) {
			ctx := context.Background()
			_, err := sv.CreateSecurityGroup(ctx, &services.CreateSecurityGroupRequest{
				SecurityGroup: sgWithID,
			})

			assert.Error(t, err)
		})

	runTest(t, "update security group with explicit ID equal to the old ID",
		func(t *testing.T, sv *ContrailTypeLogicService) {
			sv.ReadService.(*servicesmock.MockReadService).
				EXPECT().GetSecurityGroup(gomock.Not(gomock.Nil()), &services.GetSecurityGroupRequest{
				ID: sgWithID.UUID,
			}).Return(&services.GetSecurityGroupResponse{SecurityGroup: sgWithID}, nil).Times(1)

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
				SecurityGroup: sgWithID,
				FieldMask: types.FieldMask{
					Paths: []string{
						models.SecurityGroupFieldSecurityGroupID,
					},
				},
			})

			if assert.NoError(t, err) {
				assert.Equal(t, sgWithID, res.SecurityGroup)
			}
		})

	runTest(t, "fail to update security group with explicit ID different than the old ID",
		func(t *testing.T, sv *ContrailTypeLogicService) {
			sv.ReadService.(*servicesmock.MockReadService).
				EXPECT().GetSecurityGroup(gomock.Not(gomock.Nil()), &services.GetSecurityGroupRequest{
				ID: sgWithID.UUID,
			}).Return(&services.GetSecurityGroupResponse{SecurityGroup: sgWithID}, nil).Times(1)

			ctx := context.Background()
			_, err := sv.UpdateSecurityGroup(ctx, &services.UpdateSecurityGroupRequest{
				SecurityGroup: sgWithID2,
				FieldMask: types.FieldMask{
					Paths: []string{
						models.SecurityGroupFieldSecurityGroupID,
					},
				},
			})

			assert.Error(t, err)
		})
}
