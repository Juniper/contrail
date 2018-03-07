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

func TestServiceAppliance(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeServiceAppliance()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var PhysicalInterfaceCreateRef []*models.ServiceAppliancePhysicalInterfaceRef
	var PhysicalInterfaceRefModel *models.PhysicalInterface

	PhysicalInterfaceRefUUID := uuid.NewV4().String()
	PhysicalInterfaceRefUUID1 := uuid.NewV4().String()
	PhysicalInterfaceRefUUID2 := uuid.NewV4().String()

	PhysicalInterfaceRefModel = models.MakePhysicalInterface()
	PhysicalInterfaceRefModel.UUID = PhysicalInterfaceRefUUID
	PhysicalInterfaceRefModel.FQName = []string{"test", PhysicalInterfaceRefUUID}
	_, err = db.CreatePhysicalInterface(ctx, &models.CreatePhysicalInterfaceRequest{
		PhysicalInterface: PhysicalInterfaceRefModel,
	})
	PhysicalInterfaceRefModel.UUID = PhysicalInterfaceRefUUID1
	PhysicalInterfaceRefModel.FQName = []string{"test", PhysicalInterfaceRefUUID1}
	_, err = db.CreatePhysicalInterface(ctx, &models.CreatePhysicalInterfaceRequest{
		PhysicalInterface: PhysicalInterfaceRefModel,
	})
	PhysicalInterfaceRefModel.UUID = PhysicalInterfaceRefUUID2
	PhysicalInterfaceRefModel.FQName = []string{"test", PhysicalInterfaceRefUUID2}
	_, err = db.CreatePhysicalInterface(ctx, &models.CreatePhysicalInterfaceRequest{
		PhysicalInterface: PhysicalInterfaceRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	PhysicalInterfaceCreateRef = append(PhysicalInterfaceCreateRef,
		&models.ServiceAppliancePhysicalInterfaceRef{UUID: PhysicalInterfaceRefUUID, To: []string{"test", PhysicalInterfaceRefUUID}})
	PhysicalInterfaceCreateRef = append(PhysicalInterfaceCreateRef,
		&models.ServiceAppliancePhysicalInterfaceRef{UUID: PhysicalInterfaceRefUUID2, To: []string{"test", PhysicalInterfaceRefUUID2}})
	model.PhysicalInterfaceRefs = PhysicalInterfaceCreateRef

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

	_, err = db.CreateServiceAppliance(ctx,
		&models.CreateServiceApplianceRequest{
			ServiceAppliance: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListServiceAppliance(ctx, &models.ListServiceApplianceRequest{
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
	if len(response.ServiceAppliances) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteServiceAppliance(ctxDemo,
		&models.DeleteServiceApplianceRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateServiceAppliance(ctx,
		&models.CreateServiceApplianceRequest{
			ServiceAppliance: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteServiceAppliance(ctx,
		&models.DeleteServiceApplianceRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetServiceAppliance(ctx, &models.GetServiceApplianceRequest{
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
