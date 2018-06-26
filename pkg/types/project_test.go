package types

import (
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
				FieldMask: types.FieldMask{Paths: []string{models.ProjectPropertyIDVxlanRouting}},
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
				FieldMask: types.FieldMask{Paths: []string{models.ProjectPropertyIDVxlanRouting}},
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
				FieldMask: types.FieldMask{Paths: []string{models.ProjectPropertyIDVxlanRouting}},
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
