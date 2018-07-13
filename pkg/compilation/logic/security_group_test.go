package logic

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateSecurityGroupCreatesACLs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAPIService := servicesmock.NewMockService(mockCtrl)
	service := NewService(mockAPIService)

	expectedIngressACL := &models.AccessControlList{
		Name: "ingress-access-control-list",
	}
	expectedEgressACL := &models.AccessControlList{
		Name: "egress-access-control-list",
	}

	expectCreateACL(mockAPIService, expectedIngressACL)
	expectCreateACL(mockAPIService, expectedEgressACL)

	_, err := service.CreateSecurityGroup(context.Background(), &services.CreateSecurityGroupRequest{
		SecurityGroup: &models.SecurityGroup{},
	})
	assert.NoError(t, err)
}

func expectCreateACL(mockAPIService *servicesmock.MockService, expectedACL *models.AccessControlList) {
	mockAPIService.EXPECT().CreateAccessControlList(notNil(), &services.CreateAccessControlListRequest{
		expectedACL,
	}).Return(&services.CreateAccessControlListResponse{expectedACL}, nil)
}

func notNil() gomock.Matcher {
	return gomock.Not(gomock.Nil())
}
