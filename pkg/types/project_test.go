package types

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

func TestCheckVxlanConfig(t *testing.T) {
	tests := []struct {
		name           string
		updateRequest  *services.UpdateProjectRequest
		currentProject *models.Project
		errorCode      codes.Code
	}{
		{
			name: "No vxlan change requested",
			updateRequest: &services.UpdateProjectRequest{
				FieldMask: types.FieldMask{Paths: []string{"Foo", "Bar"}},
				Project:   &models.Project{VxlanRouting: false},
			},
			currentProject: &models.Project{
				VxlanRouting:   true,
				LogicalRouters: []*models.LogicalRouter{{}},
			},
		},
		{
			name: "Vxlan will not change",
			updateRequest: &services.UpdateProjectRequest{
				FieldMask: types.FieldMask{Paths: []string{models.ProjectFieldVxlanRouting}},
				Project:   &models.Project{VxlanRouting: true},
			},
			currentProject: &models.Project{
				VxlanRouting:   true,
				LogicalRouters: []*models.LogicalRouter{{}},
			},
		},
		{
			name: "No vxlan routers attached",
			updateRequest: &services.UpdateProjectRequest{
				FieldMask: types.FieldMask{Paths: []string{models.ProjectFieldVxlanRouting}},
				Project:   &models.Project{VxlanRouting: false},
			},
			currentProject: &models.Project{
				VxlanRouting:   true,
				LogicalRouters: []*models.LogicalRouter{},
			},
		},
		{
			name: "Vxlan routers already attached",
			updateRequest: &services.UpdateProjectRequest{
				FieldMask: types.FieldMask{Paths: []string{models.ProjectFieldVxlanRouting}},
				Project:   &models.Project{VxlanRouting: false},
			},
			currentProject: &models.Project{
				VxlanRouting:   true,
				LogicalRouters: []*models.LogicalRouter{{}},
			},
			errorCode: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sv ContrailTypeLogicService
			err := sv.checkVxlanConfig(tt.currentProject, tt.updateRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestEnsureDefaultApplicationPolicySet(t *testing.T) {
	tests := []struct {
		name      string
		project   models.Project
		initMocks func(*ContrailTypeLogicService)
		fails     bool
	}{
		{
			name: "create returns internal",
			initMocks: func(s *ContrailTypeLogicService) {
				m := s.WriteService.(*servicesmock.MockWriteService) //nolint: errcheck

				m.EXPECT().CreateApplicationPolicySet(
					gomock.Not(gomock.Nil()),
					gomock.Not(gomock.Nil()),
				).Return(nil, errutil.ErrorInternal).Times(1)
			},
			fails: true,
		},
		{
			name: "create returns conflict",
			initMocks: func(s *ContrailTypeLogicService) {
				m := s.WriteService.(*servicesmock.MockWriteService) //nolint: errcheck

				m.EXPECT().CreateApplicationPolicySet(
					gomock.Not(gomock.Nil()),
					gomock.Not(gomock.Nil()),
				).Return(nil, errutil.ErrorConflict).Times(1)
			},
		},
		{
			name: "create returns ApplicationPolicySet object",
			initMocks: func(s *ContrailTypeLogicService) {
				m := s.WriteService.(*servicesmock.MockWriteService) //nolint: errcheck

				m.EXPECT().CreateApplicationPolicySet(
					gomock.Not(gomock.Nil()),
					gomock.Not(gomock.Nil()),
				).Return(
					&services.CreateApplicationPolicySetResponse{
						ApplicationPolicySet: models.MakeApplicationPolicySet(),
					}, nil,
				).Times(1)

				m.EXPECT().UpdateProject(
					gomock.Not(gomock.Nil()),
					gomock.Not(gomock.Nil()),
				).Return(
					&services.UpdateProjectResponse{}, nil,
				).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			service := makeMockedContrailTypeLogicService(mockCtrl)
			if tt.initMocks != nil {
				tt.initMocks(service)
			}

			err := service.ensureDefaultApplicationPolicySet(context.Background(), &tt.project)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
