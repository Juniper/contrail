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

func TestCreateObjectWithConfigRootAsOneOfManyParents(t *testing.T) {
	s := integration.NewRunningAPIServer(t, "../../..", db.DriverPostgreSQL)
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())

	pmUUID := "test_test"

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
}

func TestCreateObjectWithConfigRootAsSingleParent(t *testing.T) {
	s := integration.NewRunningAPIServer(t, "../../..", db.DriverPostgreSQL)
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())

	dUUID := "test_test"

	_, err := hc.CreateDomain(
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
}

func TestCreateObjectWithMultipleParentsButnoConfigRoot(t *testing.T) {
	s := integration.NewRunningAPIServer(t, "../../..", db.DriverPostgreSQL)
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())

	aclUUID := "test_test"

	_, err := hc.CreateAccessControlList(
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