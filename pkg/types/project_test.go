package types

import (
	"testing"
	"reflect"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestVxlanConfigField(t *testing.T) {
	_, ok := reflect.TypeOf(models.Project{}).FieldByName(vxlanConfigField)
	assert.True(t, ok, "models.Project should have vxlanConfigFiled property")
}

func TestCheckVxlanConfig(t *testing.T) {
	tests := []struct{
		name string
		updateRequest *models.UpdateProjectRequest
		currentProject *models.Project
		errorCode codes.Code
	}{
		{
			name: "No vxlan change requested",
			updateRequest: &models.UpdateProjectRequest{
				FieldMask: types.FieldMask{Paths: []string{"Foo", "Bar"}},
			},
		},
		{
			name: "Vxlan will not change",
			updateRequest: &models.UpdateProjectRequest{
				FieldMask: types.FieldMask{Paths: []string{vxlanConfigField}},
				Project: &models.Project{VxlanRouting: true},
			},
			currentProject:	&models.Project{VxlanRouting: true},
		},
		{
			name: "No vxlan routers attached",
			updateRequest: &models.UpdateProjectRequest{
				FieldMask: types.FieldMask{Paths: []string{vxlanConfigField}},
				Project: &models.Project{VxlanRouting: false},
			},
			currentProject:	&models.Project{VxlanRouting: true},
		},
		{
			name: "Vxlan routers already attached",
			updateRequest: &models.UpdateProjectRequest{
				FieldMask: types.FieldMask{Paths: []string{vxlanConfigField}},
				Project: &models.Project{VxlanRouting: false},
			},
			currentProject:	&models.Project{
				VxlanRouting: true,
				LogicalRouters: []*models.LogicalRouter{{}},
			},
			errorCode: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkVxlanConfig(tt.currentProject, tt.updateRequest)
			if tt.errorCode != 0 {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			}
			})
	}
}

