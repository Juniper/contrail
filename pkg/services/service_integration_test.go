package services_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestCreateObjectWithoutPassingFqNameParentUuidAndTypeWhenHasConfigRootAsPossibleParent(t *testing.T) {
	s := integration.NewRunningAPIServer(t, "../../..", db.DriverPostgreSQL)
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())

	pmUUID := "PolicyManagementTest"
	dUUID := "DomainTest"
	aclUUID := "AccessControlListTest"

	// Test for object with multiple parents including ConfigRoot
	_, err := hc.CreatePolicyManagement(
		context.Background(),
		&services.CreatePolicyManagementRequest{
			PolicyManagement: &models.PolicyManagement{
				UUID: pmUUID,
				Name: "test",
			},
		},
	)
	assert.NoError(t, err)

	_, err = hc.DeletePolicyManagement(
		context.Background(),
		&services.DeletePolicyManagementRequest{
			ID: pmUUID,
		},
	)
	assert.NoError(t, err)

	// Test for object with ConfigRoot as the only possible parent
	_, err = hc.CreateDomain(
		context.Background(),
		&services.CreateDomainRequest{
			Domain: &models.Domain{
				UUID: dUUID,
				Name: "test",
			},
		},
	)
	assert.NoError(t, err)

	_, err = hc.DeleteDomain(
		context.Background(),
		&services.DeleteDomainRequest{
			ID: dUUID,
		},
	)
	assert.NoError(t, err)

	// Test for object with multiple parents without ConfigRoot
	_, err = hc.CreateAccessControlList(
		context.Background(),
		&services.CreateAccessControlListRequest{
			AccessControlList: &models.AccessControlList{
				UUID: aclUUID,
				Name: "test",
			},
		},
	)
	assert.Error(t, err)
}