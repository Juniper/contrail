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

func TestServiceTemplate(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeServiceTemplate()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceApplianceSetCreateRef []*models.ServiceTemplateServiceApplianceSetRef
	var ServiceApplianceSetRefModel *models.ServiceApplianceSet

	ServiceApplianceSetRefUUID := uuid.NewV4().String()
	ServiceApplianceSetRefUUID1 := uuid.NewV4().String()
	ServiceApplianceSetRefUUID2 := uuid.NewV4().String()

	ServiceApplianceSetRefModel = models.MakeServiceApplianceSet()
	ServiceApplianceSetRefModel.UUID = ServiceApplianceSetRefUUID
	ServiceApplianceSetRefModel.FQName = []string{"test", ServiceApplianceSetRefUUID}
	_, err = db.CreateServiceApplianceSet(ctx, &models.CreateServiceApplianceSetRequest{
		ServiceApplianceSet: ServiceApplianceSetRefModel,
	})
	ServiceApplianceSetRefModel.UUID = ServiceApplianceSetRefUUID1
	ServiceApplianceSetRefModel.FQName = []string{"test", ServiceApplianceSetRefUUID1}
	_, err = db.CreateServiceApplianceSet(ctx, &models.CreateServiceApplianceSetRequest{
		ServiceApplianceSet: ServiceApplianceSetRefModel,
	})
	ServiceApplianceSetRefModel.UUID = ServiceApplianceSetRefUUID2
	ServiceApplianceSetRefModel.FQName = []string{"test", ServiceApplianceSetRefUUID2}
	_, err = db.CreateServiceApplianceSet(ctx, &models.CreateServiceApplianceSetRequest{
		ServiceApplianceSet: ServiceApplianceSetRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceApplianceSetCreateRef = append(ServiceApplianceSetCreateRef,
		&models.ServiceTemplateServiceApplianceSetRef{UUID: ServiceApplianceSetRefUUID, To: []string{"test", ServiceApplianceSetRefUUID}})
	ServiceApplianceSetCreateRef = append(ServiceApplianceSetCreateRef,
		&models.ServiceTemplateServiceApplianceSetRef{UUID: ServiceApplianceSetRefUUID2, To: []string{"test", ServiceApplianceSetRefUUID2}})
	model.ServiceApplianceSetRefs = ServiceApplianceSetCreateRef

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

	_, err = db.CreateServiceTemplate(ctx,
		&models.CreateServiceTemplateRequest{
			ServiceTemplate: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListServiceTemplate(ctx, &models.ListServiceTemplateRequest{
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
	if len(response.ServiceTemplates) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteServiceTemplate(ctxDemo,
		&models.DeleteServiceTemplateRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateServiceTemplate(ctx,
		&models.CreateServiceTemplateRequest{
			ServiceTemplate: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteServiceTemplate(ctx,
		&models.DeleteServiceTemplateRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetServiceTemplate(ctx, &models.GetServiceTemplateRequest{
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
