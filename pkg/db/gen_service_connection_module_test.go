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

func TestServiceConnectionModule(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeServiceConnectionModule()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceObjectCreateRef []*models.ServiceConnectionModuleServiceObjectRef
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
		&models.ServiceConnectionModuleServiceObjectRef{UUID: ServiceObjectRefUUID, To: []string{"test", ServiceObjectRefUUID}})
	ServiceObjectCreateRef = append(ServiceObjectCreateRef,
		&models.ServiceConnectionModuleServiceObjectRef{UUID: ServiceObjectRefUUID2, To: []string{"test", ServiceObjectRefUUID2}})
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

	_, err = db.CreateServiceConnectionModule(ctx,
		&models.CreateServiceConnectionModuleRequest{
			ServiceConnectionModule: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListServiceConnectionModule(ctx, &models.ListServiceConnectionModuleRequest{
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
	if len(response.ServiceConnectionModules) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteServiceConnectionModule(ctxDemo,
		&models.DeleteServiceConnectionModuleRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateServiceConnectionModule(ctx,
		&models.CreateServiceConnectionModuleRequest{
			ServiceConnectionModule: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteServiceConnectionModule(ctx,
		&models.DeleteServiceConnectionModuleRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetServiceConnectionModule(ctx, &models.GetServiceConnectionModuleRequest{
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
