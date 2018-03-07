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

func TestServiceInstance(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeServiceInstance()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceTemplateCreateRef []*models.ServiceInstanceServiceTemplateRef
	var ServiceTemplateRefModel *models.ServiceTemplate

	ServiceTemplateRefUUID := uuid.NewV4().String()
	ServiceTemplateRefUUID1 := uuid.NewV4().String()
	ServiceTemplateRefUUID2 := uuid.NewV4().String()

	ServiceTemplateRefModel = models.MakeServiceTemplate()
	ServiceTemplateRefModel.UUID = ServiceTemplateRefUUID
	ServiceTemplateRefModel.FQName = []string{"test", ServiceTemplateRefUUID}
	_, err = db.CreateServiceTemplate(ctx, &models.CreateServiceTemplateRequest{
		ServiceTemplate: ServiceTemplateRefModel,
	})
	ServiceTemplateRefModel.UUID = ServiceTemplateRefUUID1
	ServiceTemplateRefModel.FQName = []string{"test", ServiceTemplateRefUUID1}
	_, err = db.CreateServiceTemplate(ctx, &models.CreateServiceTemplateRequest{
		ServiceTemplate: ServiceTemplateRefModel,
	})
	ServiceTemplateRefModel.UUID = ServiceTemplateRefUUID2
	ServiceTemplateRefModel.FQName = []string{"test", ServiceTemplateRefUUID2}
	_, err = db.CreateServiceTemplate(ctx, &models.CreateServiceTemplateRequest{
		ServiceTemplate: ServiceTemplateRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceTemplateCreateRef = append(ServiceTemplateCreateRef,
		&models.ServiceInstanceServiceTemplateRef{UUID: ServiceTemplateRefUUID, To: []string{"test", ServiceTemplateRefUUID}})
	ServiceTemplateCreateRef = append(ServiceTemplateCreateRef,
		&models.ServiceInstanceServiceTemplateRef{UUID: ServiceTemplateRefUUID2, To: []string{"test", ServiceTemplateRefUUID2}})
	model.ServiceTemplateRefs = ServiceTemplateCreateRef

	var InstanceIPCreateRef []*models.ServiceInstanceInstanceIPRef
	var InstanceIPRefModel *models.InstanceIP

	InstanceIPRefUUID := uuid.NewV4().String()
	InstanceIPRefUUID1 := uuid.NewV4().String()
	InstanceIPRefUUID2 := uuid.NewV4().String()

	InstanceIPRefModel = models.MakeInstanceIP()
	InstanceIPRefModel.UUID = InstanceIPRefUUID
	InstanceIPRefModel.FQName = []string{"test", InstanceIPRefUUID}
	_, err = db.CreateInstanceIP(ctx, &models.CreateInstanceIPRequest{
		InstanceIP: InstanceIPRefModel,
	})
	InstanceIPRefModel.UUID = InstanceIPRefUUID1
	InstanceIPRefModel.FQName = []string{"test", InstanceIPRefUUID1}
	_, err = db.CreateInstanceIP(ctx, &models.CreateInstanceIPRequest{
		InstanceIP: InstanceIPRefModel,
	})
	InstanceIPRefModel.UUID = InstanceIPRefUUID2
	InstanceIPRefModel.FQName = []string{"test", InstanceIPRefUUID2}
	_, err = db.CreateInstanceIP(ctx, &models.CreateInstanceIPRequest{
		InstanceIP: InstanceIPRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	InstanceIPCreateRef = append(InstanceIPCreateRef,
		&models.ServiceInstanceInstanceIPRef{UUID: InstanceIPRefUUID, To: []string{"test", InstanceIPRefUUID}})
	InstanceIPCreateRef = append(InstanceIPCreateRef,
		&models.ServiceInstanceInstanceIPRef{UUID: InstanceIPRefUUID2, To: []string{"test", InstanceIPRefUUID2}})
	model.InstanceIPRefs = InstanceIPCreateRef

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

	_, err = db.CreateServiceInstance(ctx,
		&models.CreateServiceInstanceRequest{
			ServiceInstance: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListServiceInstance(ctx, &models.ListServiceInstanceRequest{
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
	if len(response.ServiceInstances) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteServiceInstance(ctxDemo,
		&models.DeleteServiceInstanceRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateServiceInstance(ctx,
		&models.CreateServiceInstanceRequest{
			ServiceInstance: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteServiceInstance(ctx,
		&models.DeleteServiceInstanceRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetServiceInstance(ctx, &models.GetServiceInstanceRequest{
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
