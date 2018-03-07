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

func TestLoadbalancerPool(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeLoadbalancerPool()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceApplianceSetCreateRef []*models.LoadbalancerPoolServiceApplianceSetRef
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
		&models.LoadbalancerPoolServiceApplianceSetRef{UUID: ServiceApplianceSetRefUUID, To: []string{"test", ServiceApplianceSetRefUUID}})
	ServiceApplianceSetCreateRef = append(ServiceApplianceSetCreateRef,
		&models.LoadbalancerPoolServiceApplianceSetRef{UUID: ServiceApplianceSetRefUUID2, To: []string{"test", ServiceApplianceSetRefUUID2}})
	model.ServiceApplianceSetRefs = ServiceApplianceSetCreateRef

	var VirtualMachineInterfaceCreateRef []*models.LoadbalancerPoolVirtualMachineInterfaceRef
	var VirtualMachineInterfaceRefModel *models.VirtualMachineInterface

	VirtualMachineInterfaceRefUUID := uuid.NewV4().String()
	VirtualMachineInterfaceRefUUID1 := uuid.NewV4().String()
	VirtualMachineInterfaceRefUUID2 := uuid.NewV4().String()

	VirtualMachineInterfaceRefModel = models.MakeVirtualMachineInterface()
	VirtualMachineInterfaceRefModel.UUID = VirtualMachineInterfaceRefUUID
	VirtualMachineInterfaceRefModel.FQName = []string{"test", VirtualMachineInterfaceRefUUID}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfaceRefModel,
	})
	VirtualMachineInterfaceRefModel.UUID = VirtualMachineInterfaceRefUUID1
	VirtualMachineInterfaceRefModel.FQName = []string{"test", VirtualMachineInterfaceRefUUID1}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfaceRefModel,
	})
	VirtualMachineInterfaceRefModel.UUID = VirtualMachineInterfaceRefUUID2
	VirtualMachineInterfaceRefModel.FQName = []string{"test", VirtualMachineInterfaceRefUUID2}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfaceRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfaceCreateRef = append(VirtualMachineInterfaceCreateRef,
		&models.LoadbalancerPoolVirtualMachineInterfaceRef{UUID: VirtualMachineInterfaceRefUUID, To: []string{"test", VirtualMachineInterfaceRefUUID}})
	VirtualMachineInterfaceCreateRef = append(VirtualMachineInterfaceCreateRef,
		&models.LoadbalancerPoolVirtualMachineInterfaceRef{UUID: VirtualMachineInterfaceRefUUID2, To: []string{"test", VirtualMachineInterfaceRefUUID2}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfaceCreateRef

	var LoadbalancerListenerCreateRef []*models.LoadbalancerPoolLoadbalancerListenerRef
	var LoadbalancerListenerRefModel *models.LoadbalancerListener

	LoadbalancerListenerRefUUID := uuid.NewV4().String()
	LoadbalancerListenerRefUUID1 := uuid.NewV4().String()
	LoadbalancerListenerRefUUID2 := uuid.NewV4().String()

	LoadbalancerListenerRefModel = models.MakeLoadbalancerListener()
	LoadbalancerListenerRefModel.UUID = LoadbalancerListenerRefUUID
	LoadbalancerListenerRefModel.FQName = []string{"test", LoadbalancerListenerRefUUID}
	_, err = db.CreateLoadbalancerListener(ctx, &models.CreateLoadbalancerListenerRequest{
		LoadbalancerListener: LoadbalancerListenerRefModel,
	})
	LoadbalancerListenerRefModel.UUID = LoadbalancerListenerRefUUID1
	LoadbalancerListenerRefModel.FQName = []string{"test", LoadbalancerListenerRefUUID1}
	_, err = db.CreateLoadbalancerListener(ctx, &models.CreateLoadbalancerListenerRequest{
		LoadbalancerListener: LoadbalancerListenerRefModel,
	})
	LoadbalancerListenerRefModel.UUID = LoadbalancerListenerRefUUID2
	LoadbalancerListenerRefModel.FQName = []string{"test", LoadbalancerListenerRefUUID2}
	_, err = db.CreateLoadbalancerListener(ctx, &models.CreateLoadbalancerListenerRequest{
		LoadbalancerListener: LoadbalancerListenerRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	LoadbalancerListenerCreateRef = append(LoadbalancerListenerCreateRef,
		&models.LoadbalancerPoolLoadbalancerListenerRef{UUID: LoadbalancerListenerRefUUID, To: []string{"test", LoadbalancerListenerRefUUID}})
	LoadbalancerListenerCreateRef = append(LoadbalancerListenerCreateRef,
		&models.LoadbalancerPoolLoadbalancerListenerRef{UUID: LoadbalancerListenerRefUUID2, To: []string{"test", LoadbalancerListenerRefUUID2}})
	model.LoadbalancerListenerRefs = LoadbalancerListenerCreateRef

	var ServiceInstanceCreateRef []*models.LoadbalancerPoolServiceInstanceRef
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
		&models.LoadbalancerPoolServiceInstanceRef{UUID: ServiceInstanceRefUUID, To: []string{"test", ServiceInstanceRefUUID}})
	ServiceInstanceCreateRef = append(ServiceInstanceCreateRef,
		&models.LoadbalancerPoolServiceInstanceRef{UUID: ServiceInstanceRefUUID2, To: []string{"test", ServiceInstanceRefUUID2}})
	model.ServiceInstanceRefs = ServiceInstanceCreateRef

	var LoadbalancerHealthmonitorCreateRef []*models.LoadbalancerPoolLoadbalancerHealthmonitorRef
	var LoadbalancerHealthmonitorRefModel *models.LoadbalancerHealthmonitor

	LoadbalancerHealthmonitorRefUUID := uuid.NewV4().String()
	LoadbalancerHealthmonitorRefUUID1 := uuid.NewV4().String()
	LoadbalancerHealthmonitorRefUUID2 := uuid.NewV4().String()

	LoadbalancerHealthmonitorRefModel = models.MakeLoadbalancerHealthmonitor()
	LoadbalancerHealthmonitorRefModel.UUID = LoadbalancerHealthmonitorRefUUID
	LoadbalancerHealthmonitorRefModel.FQName = []string{"test", LoadbalancerHealthmonitorRefUUID}
	_, err = db.CreateLoadbalancerHealthmonitor(ctx, &models.CreateLoadbalancerHealthmonitorRequest{
		LoadbalancerHealthmonitor: LoadbalancerHealthmonitorRefModel,
	})
	LoadbalancerHealthmonitorRefModel.UUID = LoadbalancerHealthmonitorRefUUID1
	LoadbalancerHealthmonitorRefModel.FQName = []string{"test", LoadbalancerHealthmonitorRefUUID1}
	_, err = db.CreateLoadbalancerHealthmonitor(ctx, &models.CreateLoadbalancerHealthmonitorRequest{
		LoadbalancerHealthmonitor: LoadbalancerHealthmonitorRefModel,
	})
	LoadbalancerHealthmonitorRefModel.UUID = LoadbalancerHealthmonitorRefUUID2
	LoadbalancerHealthmonitorRefModel.FQName = []string{"test", LoadbalancerHealthmonitorRefUUID2}
	_, err = db.CreateLoadbalancerHealthmonitor(ctx, &models.CreateLoadbalancerHealthmonitorRequest{
		LoadbalancerHealthmonitor: LoadbalancerHealthmonitorRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	LoadbalancerHealthmonitorCreateRef = append(LoadbalancerHealthmonitorCreateRef,
		&models.LoadbalancerPoolLoadbalancerHealthmonitorRef{UUID: LoadbalancerHealthmonitorRefUUID, To: []string{"test", LoadbalancerHealthmonitorRefUUID}})
	LoadbalancerHealthmonitorCreateRef = append(LoadbalancerHealthmonitorCreateRef,
		&models.LoadbalancerPoolLoadbalancerHealthmonitorRef{UUID: LoadbalancerHealthmonitorRefUUID2, To: []string{"test", LoadbalancerHealthmonitorRefUUID2}})
	model.LoadbalancerHealthmonitorRefs = LoadbalancerHealthmonitorCreateRef

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

	_, err = db.CreateLoadbalancerPool(ctx,
		&models.CreateLoadbalancerPoolRequest{
			LoadbalancerPool: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListLoadbalancerPool(ctx, &models.ListLoadbalancerPoolRequest{
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
	if len(response.LoadbalancerPools) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteLoadbalancerPool(ctxDemo,
		&models.DeleteLoadbalancerPoolRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateLoadbalancerPool(ctx,
		&models.CreateLoadbalancerPoolRequest{
			LoadbalancerPool: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteLoadbalancerPool(ctx,
		&models.DeleteLoadbalancerPoolRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetLoadbalancerPool(ctx, &models.GetLoadbalancerPoolRequest{
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
