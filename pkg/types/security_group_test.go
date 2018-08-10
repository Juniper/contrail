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

func TestCreateSecurityGroup(t *testing.T) {
	tests := []struct {
		name         string
		requestedSG  *models.SecurityGroup
		expectedSG   *models.SecurityGroup
		allocatedInt int64
		fails        bool
	}{
		{
			name:         "create security group",
			requestedSG:  &models.SecurityGroup{UUID: "sg_uuid"},
			expectedSG:   &models.SecurityGroup{UUID: "sg_uuid", SecurityGroupID: 8000001},
			allocatedInt: 1,
		},
		{
			name:        "fail to create security group with explicit ID",
			requestedSG: &models.SecurityGroup{UUID: "sg_uuid", SecurityGroupID: 8000001},
			fails:       true,
		},
	}

	for _, tt := range tests {
		runTest(t, tt.name, func(t *testing.T, sv *ContrailTypeLogicService) {
			allocateCall := sv.IntPoolAllocator.(*typesmock.MockIntPoolAllocator).
				EXPECT().AllocateInt(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
				Return(tt.allocatedInt, nil)
			createCall := sv.Next().(*servicesmock.MockService).
				EXPECT().CreateSecurityGroup(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
				DoAndReturn(func(_ context.Context, request *services.CreateSecurityGroupRequest) (
					*services.CreateSecurityGroupResponse, error) {
					return &services.CreateSecurityGroupResponse{
						SecurityGroup: request.SecurityGroup,
					}, nil
				})

			if tt.fails {
				allocateCall.MaxTimes(1)
				createCall.MaxTimes(1)
			} else {
				allocateCall.Times(1)
				createCall.Times(1)
			}

			ctx := context.Background()
			res, err := sv.CreateSecurityGroup(ctx, &services.CreateSecurityGroupRequest{
				SecurityGroup: tt.requestedSG,
			})

			if tt.fails {
				assert.Error(t, err)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tt.expectedSG, res.SecurityGroup)
			}
		})
	}
}

func TestUpdateSecurityGroup(t *testing.T) {
	tests := []struct {
		name           string
		existingSG     *models.SecurityGroup
		requestedSG    *models.SecurityGroup
		expectedSG     *models.SecurityGroup
		fieldMaskPaths []string
		fails          bool
	}{
		{
			name:           "update security group",
			existingSG:     &models.SecurityGroup{UUID: "sg_uuid", SecurityGroupID: 8000001},
			requestedSG:    &models.SecurityGroup{UUID: "sg_uuid"},
			expectedSG:     &models.SecurityGroup{UUID: "sg_uuid", SecurityGroupID: 8000001},
			fieldMaskPaths: []string{},
		},
		{
			name:           "update security group with explicit ID equal to the old ID",
			existingSG:     &models.SecurityGroup{UUID: "sg_uuid", SecurityGroupID: 8000001},
			requestedSG:    &models.SecurityGroup{UUID: "sg_uuid", SecurityGroupID: 8000001},
			expectedSG:     &models.SecurityGroup{UUID: "sg_uuid", SecurityGroupID: 8000001},
			fieldMaskPaths: []string{models.SecurityGroupFieldSecurityGroupID},
		},
		{
			name:           "fail to update security group with explicit ID different than the old ID",
			existingSG:     &models.SecurityGroup{UUID: "sg_uuid", SecurityGroupID: 8000001},
			requestedSG:    &models.SecurityGroup{UUID: "sg_uuid", SecurityGroupID: 8000002},
			fieldMaskPaths: []string{models.SecurityGroupFieldSecurityGroupID},
			fails:          true,
		},
	}

	for _, tt := range tests {
		runTest(t, tt.name, func(t *testing.T, sv *ContrailTypeLogicService) {
			sv.ReadService.(*servicesmock.MockReadService).
				EXPECT().GetSecurityGroup(gomock.Not(gomock.Nil()), &services.GetSecurityGroupRequest{
				ID: tt.requestedSG.UUID,
			}).Return(&services.GetSecurityGroupResponse{SecurityGroup: tt.existingSG}, nil).Times(1)

			updateCall := sv.Next().(*servicesmock.MockService).
				EXPECT().UpdateSecurityGroup(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
				DoAndReturn(func(_ context.Context, request *services.UpdateSecurityGroupRequest) (
					*services.UpdateSecurityGroupResponse, error) {
					return &services.UpdateSecurityGroupResponse{
						SecurityGroup: request.SecurityGroup,
					}, nil
				})

			if tt.fails {
				updateCall.MaxTimes(1)
			} else {
				updateCall.Times(1)
			}

			ctx := context.Background()
			res, err := sv.UpdateSecurityGroup(ctx, &services.UpdateSecurityGroupRequest{
				SecurityGroup: tt.requestedSG,
				FieldMask: types.FieldMask{
					Paths: tt.fieldMaskPaths,
				},
			})

			if tt.fails {
				assert.Error(t, err)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tt.expectedSG, res.SecurityGroup)
			}
		})
	}
}
