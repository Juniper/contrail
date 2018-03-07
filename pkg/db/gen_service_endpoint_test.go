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

func TestServiceEndpoint(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeServiceEndpoint()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceConnectionModuleCreateRef []*models.ServiceEndpointServiceConnectionModuleRef
	var ServiceConnectionModuleRefModel *models.ServiceConnectionModule

	ServiceConnectionModuleRefUUID := uuid.NewV4().String()
	ServiceConnectionModuleRefUUID1 := uuid.NewV4().String()
	ServiceConnectionModuleRefUUID2 := uuid.NewV4().String()

	ServiceConnectionModuleRefModel = models.MakeServiceConnectionModule()
	ServiceConnectionModuleRefModel.UUID = ServiceConnectionModuleRefUUID
	ServiceConnectionModuleRefModel.FQName = []string{"test", ServiceConnectionModuleRefUUID}
	_, err = db.CreateServiceConnectionModule(ctx, &models.CreateServiceConnectionModuleRequest{
		ServiceConnectionModule: ServiceConnectionModuleRefModel,
	})
	ServiceConnectionModuleRefModel.UUID = ServiceConnectionModuleRefUUID1
	ServiceConnectionModuleRefModel.FQName = []string{"test", ServiceConnectionModuleRefUUID1}
	_, err = db.CreateServiceConnectionModule(ctx, &models.CreateServiceConnectionModuleRequest{
		ServiceConnectionModule: ServiceConnectionModuleRefModel,
	})
	ServiceConnectionModuleRefModel.UUID = ServiceConnectionModuleRefUUID2
	ServiceConnectionModuleRefModel.FQName = []string{"test", ServiceConnectionModuleRefUUID2}
	_, err = db.CreateServiceConnectionModule(ctx, &models.CreateServiceConnectionModuleRequest{
		ServiceConnectionModule: ServiceConnectionModuleRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceConnectionModuleCreateRef = append(ServiceConnectionModuleCreateRef,
		&models.ServiceEndpointServiceConnectionModuleRef{UUID: ServiceConnectionModuleRefUUID, To: []string{"test", ServiceConnectionModuleRefUUID}})
	ServiceConnectionModuleCreateRef = append(ServiceConnectionModuleCreateRef,
		&models.ServiceEndpointServiceConnectionModuleRef{UUID: ServiceConnectionModuleRefUUID2, To: []string{"test", ServiceConnectionModuleRefUUID2}})
	model.ServiceConnectionModuleRefs = ServiceConnectionModuleCreateRef

	var PhysicalRouterCreateRef []*models.ServiceEndpointPhysicalRouterRef
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
		&models.ServiceEndpointPhysicalRouterRef{UUID: PhysicalRouterRefUUID, To: []string{"test", PhysicalRouterRefUUID}})
	PhysicalRouterCreateRef = append(PhysicalRouterCreateRef,
		&models.ServiceEndpointPhysicalRouterRef{UUID: PhysicalRouterRefUUID2, To: []string{"test", PhysicalRouterRefUUID2}})
	model.PhysicalRouterRefs = PhysicalRouterCreateRef

	var ServiceObjectCreateRef []*models.ServiceEndpointServiceObjectRef
	var ServiceObjectRefModel *models.ServiceObject

	ServiceObjectRefUUID := uuid.NewV4().String()
	ServiceObjectRefUUID1 := uuid.NewV4().String()
	ServiceObjectRefUUID2 := uuid.NewV4().String()

	ServiceObjectRefModel = models.MakeServiceObject()
	ServiceObjectRefModel.UUID = ServiceObjectRefUUID
	ServiceObjectRefModel.FQName = []string{"test", ServiceObjectRefUUID}
	_, err = db.CreateServiceObject(ctx, &models.CreateServiceObjectRequest{
		ServiceObject: ServiceObjectRefModel,
	})
	ServiceObjectRefModel.UUID = ServiceObjectRefUUID1
	ServiceObjectRefModel.FQName = []string{"test", ServiceObjectRefUUID1}
	_, err = db.CreateServiceObject(ctx, &models.CreateServiceObjectRequest{
		ServiceObject: ServiceObjectRefModel,
	})
	ServiceObjectRefModel.UUID = ServiceObjectRefUUID2
	ServiceObjectRefModel.FQName = []string{"test", ServiceObjectRefUUID2}
	_, err = db.CreateServiceObject(ctx, &models.CreateServiceObjectRequest{
		ServiceObject: ServiceObjectRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceObjectCreateRef = append(ServiceObjectCreateRef,
		&models.ServiceEndpointServiceObjectRef{UUID: ServiceObjectRefUUID, To: []string{"test", ServiceObjectRefUUID}})
	ServiceObjectCreateRef = append(ServiceObjectCreateRef,
		&models.ServiceEndpointServiceObjectRef{UUID: ServiceObjectRefUUID2, To: []string{"test", ServiceObjectRefUUID2}})
	model.ServiceObjectRefs = ServiceObjectCreateRef

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

	_, err = db.CreateServiceEndpoint(ctx,
		&models.CreateServiceEndpointRequest{
			ServiceEndpoint: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListServiceEndpoint(ctx, &models.ListServiceEndpointRequest{
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
	if len(response.ServiceEndpoints) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteServiceEndpoint(ctxDemo,
		&models.DeleteServiceEndpointRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateServiceEndpoint(ctx,
		&models.CreateServiceEndpointRequest{
			ServiceEndpoint: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteServiceEndpoint(ctx,
		&models.DeleteServiceEndpointRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetServiceEndpoint(ctx, &models.GetServiceEndpointRequest{
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
