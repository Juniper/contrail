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

func TestVirtualMachine(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeVirtualMachine()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceInstanceCreateRef []*models.VirtualMachineServiceInstanceRef
	var ServiceInstanceRefModel *models.ServiceInstance

	ServiceInstanceRefUUID := uuid.NewV4().String()
	ServiceInstanceRefUUID1 := uuid.NewV4().String()
	ServiceInstanceRefUUID2 := uuid.NewV4().String()

	ServiceInstanceRefModel = models.MakeServiceInstance()
	ServiceInstanceRefModel.UUID = ServiceInstanceRefUUID
	ServiceInstanceRefModel.FQName = []string{"test", ServiceInstanceRefUUID}
	_, err = db.CreateServiceInstance(ctx, &models.CreateServiceInstanceRequest{
		ServiceInstance: ServiceInstanceRefModel,
	})
	ServiceInstanceRefModel.UUID = ServiceInstanceRefUUID1
	ServiceInstanceRefModel.FQName = []string{"test", ServiceInstanceRefUUID1}
	_, err = db.CreateServiceInstance(ctx, &models.CreateServiceInstanceRequest{
		ServiceInstance: ServiceInstanceRefModel,
	})
	ServiceInstanceRefModel.UUID = ServiceInstanceRefUUID2
	ServiceInstanceRefModel.FQName = []string{"test", ServiceInstanceRefUUID2}
	_, err = db.CreateServiceInstance(ctx, &models.CreateServiceInstanceRequest{
		ServiceInstance: ServiceInstanceRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceInstanceCreateRef = append(ServiceInstanceCreateRef,
		&models.VirtualMachineServiceInstanceRef{UUID: ServiceInstanceRefUUID, To: []string{"test", ServiceInstanceRefUUID}})
	ServiceInstanceCreateRef = append(ServiceInstanceCreateRef,
		&models.VirtualMachineServiceInstanceRef{UUID: ServiceInstanceRefUUID2, To: []string{"test", ServiceInstanceRefUUID2}})
	model.ServiceInstanceRefs = ServiceInstanceCreateRef

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

	_, err = db.CreateVirtualMachine(ctx,
		&models.CreateVirtualMachineRequest{
			VirtualMachine: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListVirtualMachine(ctx, &models.ListVirtualMachineRequest{
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
	if len(response.VirtualMachines) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteVirtualMachine(ctxDemo,
		&models.DeleteVirtualMachineRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateVirtualMachine(ctx,
		&models.CreateVirtualMachineRequest{
			VirtualMachine: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteVirtualMachine(ctx,
		&models.DeleteVirtualMachineRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetVirtualMachine(ctx, &models.GetVirtualMachineRequest{
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
