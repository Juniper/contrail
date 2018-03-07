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

func TestVirtualRouter(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeVirtualRouter()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var NetworkIpamCreateRef []*models.VirtualRouterNetworkIpamRef
	var NetworkIpamRefModel *models.NetworkIpam

	NetworkIpamRefUUID := uuid.NewV4().String()
	NetworkIpamRefUUID1 := uuid.NewV4().String()
	NetworkIpamRefUUID2 := uuid.NewV4().String()

	NetworkIpamRefModel = models.MakeNetworkIpam()
	NetworkIpamRefModel.UUID = NetworkIpamRefUUID
	NetworkIpamRefModel.FQName = []string{"test", NetworkIpamRefUUID}
	_, err = db.CreateNetworkIpam(ctx, &models.CreateNetworkIpamRequest{
		NetworkIpam: NetworkIpamRefModel,
	})
	NetworkIpamRefModel.UUID = NetworkIpamRefUUID1
	NetworkIpamRefModel.FQName = []string{"test", NetworkIpamRefUUID1}
	_, err = db.CreateNetworkIpam(ctx, &models.CreateNetworkIpamRequest{
		NetworkIpam: NetworkIpamRefModel,
	})
	NetworkIpamRefModel.UUID = NetworkIpamRefUUID2
	NetworkIpamRefModel.FQName = []string{"test", NetworkIpamRefUUID2}
	_, err = db.CreateNetworkIpam(ctx, &models.CreateNetworkIpamRequest{
		NetworkIpam: NetworkIpamRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NetworkIpamCreateRef = append(NetworkIpamCreateRef,
		&models.VirtualRouterNetworkIpamRef{UUID: NetworkIpamRefUUID, To: []string{"test", NetworkIpamRefUUID}})
	NetworkIpamCreateRef = append(NetworkIpamCreateRef,
		&models.VirtualRouterNetworkIpamRef{UUID: NetworkIpamRefUUID2, To: []string{"test", NetworkIpamRefUUID2}})
	model.NetworkIpamRefs = NetworkIpamCreateRef

	var VirtualMachineCreateRef []*models.VirtualRouterVirtualMachineRef
	var VirtualMachineRefModel *models.VirtualMachine

	VirtualMachineRefUUID := uuid.NewV4().String()
	VirtualMachineRefUUID1 := uuid.NewV4().String()
	VirtualMachineRefUUID2 := uuid.NewV4().String()

	VirtualMachineRefModel = models.MakeVirtualMachine()
	VirtualMachineRefModel.UUID = VirtualMachineRefUUID
	VirtualMachineRefModel.FQName = []string{"test", VirtualMachineRefUUID}
	_, err = db.CreateVirtualMachine(ctx, &models.CreateVirtualMachineRequest{
		VirtualMachine: VirtualMachineRefModel,
	})
	VirtualMachineRefModel.UUID = VirtualMachineRefUUID1
	VirtualMachineRefModel.FQName = []string{"test", VirtualMachineRefUUID1}
	_, err = db.CreateVirtualMachine(ctx, &models.CreateVirtualMachineRequest{
		VirtualMachine: VirtualMachineRefModel,
	})
	VirtualMachineRefModel.UUID = VirtualMachineRefUUID2
	VirtualMachineRefModel.FQName = []string{"test", VirtualMachineRefUUID2}
	_, err = db.CreateVirtualMachine(ctx, &models.CreateVirtualMachineRequest{
		VirtualMachine: VirtualMachineRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineCreateRef = append(VirtualMachineCreateRef,
		&models.VirtualRouterVirtualMachineRef{UUID: VirtualMachineRefUUID, To: []string{"test", VirtualMachineRefUUID}})
	VirtualMachineCreateRef = append(VirtualMachineCreateRef,
		&models.VirtualRouterVirtualMachineRef{UUID: VirtualMachineRefUUID2, To: []string{"test", VirtualMachineRefUUID2}})
	model.VirtualMachineRefs = VirtualMachineCreateRef

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

	_, err = db.CreateVirtualRouter(ctx,
		&models.CreateVirtualRouterRequest{
			VirtualRouter: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListVirtualRouter(ctx, &models.ListVirtualRouterRequest{
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
	if len(response.VirtualRouters) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteVirtualRouter(ctxDemo,
		&models.DeleteVirtualRouterRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateVirtualRouter(ctx,
		&models.CreateVirtualRouterRequest{
			VirtualRouter: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteVirtualRouter(ctx,
		&models.DeleteVirtualRouterRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetVirtualRouter(ctx, &models.GetVirtualRouterRequest{
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
