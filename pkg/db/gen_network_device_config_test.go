// nolint
package db

import (
	"context"
	"github.com/satori/go.uuid"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestNetworkDeviceConfig(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeNetworkDeviceConfig()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var PhysicalRouterCreateRef []*models.NetworkDeviceConfigPhysicalRouterRef
	var PhysicalRouterRefModel *models.PhysicalRouter

	PhysicalRouterRefUUID := uuid.NewV4().String()
	PhysicalRouterRefUUID1 := uuid.NewV4().String()
	PhysicalRouterRefUUID2 := uuid.NewV4().String()

	PhysicalRouterRefModel = models.MakePhysicalRouter()
	PhysicalRouterRefModel.UUID = PhysicalRouterRefUUID
	PhysicalRouterRefModel.FQName = []string{"test", PhysicalRouterRefUUID}
	_, err = db.CreatePhysicalRouter(ctx, &models.CreatePhysicalRouterRequest{
		PhysicalRouter: PhysicalRouterRefModel,
	})
	PhysicalRouterRefModel.UUID = PhysicalRouterRefUUID1
	PhysicalRouterRefModel.FQName = []string{"test", PhysicalRouterRefUUID1}
	_, err = db.CreatePhysicalRouter(ctx, &models.CreatePhysicalRouterRequest{
		PhysicalRouter: PhysicalRouterRefModel,
	})
	PhysicalRouterRefModel.UUID = PhysicalRouterRefUUID2
	PhysicalRouterRefModel.FQName = []string{"test", PhysicalRouterRefUUID2}
	_, err = db.CreatePhysicalRouter(ctx, &models.CreatePhysicalRouterRequest{
		PhysicalRouter: PhysicalRouterRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	PhysicalRouterCreateRef = append(PhysicalRouterCreateRef,
		&models.NetworkDeviceConfigPhysicalRouterRef{UUID: PhysicalRouterRefUUID, To: []string{"test", PhysicalRouterRefUUID}})
	PhysicalRouterCreateRef = append(PhysicalRouterCreateRef,
		&models.NetworkDeviceConfigPhysicalRouterRef{UUID: PhysicalRouterRefUUID2, To: []string{"test", PhysicalRouterRefUUID2}})
	model.PhysicalRouterRefs = PhysicalRouterCreateRef

	//create project to which resource is shared
	projectModel := models.MakeProject()

	projectModel.UUID = uuid.NewV4().String()
	projectModel.FQName = []string{"default-domain-test", projectModel.UUID}
	projectModel.Perms2.Owner = "admin"

	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:" + projectModel.UUID, TenantAccess: 7})
	model.Perms2.Share = createShare

	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: projectModel,
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	_, err = db.CreateNetworkDeviceConfig(ctx,
		&models.CreateNetworkDeviceConfigRequest{
			NetworkDeviceConfig: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListNetworkDeviceConfig(ctx, &models.ListNetworkDeviceConfigRequest{
		Spec: &models.ListSpec{Limit: 1,
			Filters: []*models.Filter{
				&models.Filter{
					Key:    "uuid",
					Values: []string{model.UUID},
				},
			},
		}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.NetworkDeviceConfigs) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteNetworkDeviceConfig(ctxDemo,
		&models.DeleteNetworkDeviceConfigRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateNetworkDeviceConfig(ctx,
		&models.CreateNetworkDeviceConfigRequest{
			NetworkDeviceConfig: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteNetworkDeviceConfig(ctx,
		&models.DeleteNetworkDeviceConfigRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetNetworkDeviceConfig(ctx, &models.GetNetworkDeviceConfigRequest{
		ID: model.UUID})
	if err == nil {
		t.Fatal("expected not found error")
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}
	return
}
