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

func TestInstanceIP(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeInstanceIP()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var NetworkIpamCreateRef []*models.InstanceIPNetworkIpamRef
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
		&models.InstanceIPNetworkIpamRef{UUID: NetworkIpamRefUUID, To: []string{"test", NetworkIpamRefUUID}})
	NetworkIpamCreateRef = append(NetworkIpamCreateRef,
		&models.InstanceIPNetworkIpamRef{UUID: NetworkIpamRefUUID2, To: []string{"test", NetworkIpamRefUUID2}})
	model.NetworkIpamRefs = NetworkIpamCreateRef

	var VirtualNetworkCreateRef []*models.InstanceIPVirtualNetworkRef
	var VirtualNetworkRefModel *models.VirtualNetwork

	VirtualNetworkRefUUID := uuid.NewV4().String()
	VirtualNetworkRefUUID1 := uuid.NewV4().String()
	VirtualNetworkRefUUID2 := uuid.NewV4().String()

	VirtualNetworkRefModel = models.MakeVirtualNetwork()
	VirtualNetworkRefModel.UUID = VirtualNetworkRefUUID
	VirtualNetworkRefModel.FQName = []string{"test", VirtualNetworkRefUUID}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkRefModel,
	})
	VirtualNetworkRefModel.UUID = VirtualNetworkRefUUID1
	VirtualNetworkRefModel.FQName = []string{"test", VirtualNetworkRefUUID1}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkRefModel,
	})
	VirtualNetworkRefModel.UUID = VirtualNetworkRefUUID2
	VirtualNetworkRefModel.FQName = []string{"test", VirtualNetworkRefUUID2}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualNetworkCreateRef = append(VirtualNetworkCreateRef,
		&models.InstanceIPVirtualNetworkRef{UUID: VirtualNetworkRefUUID, To: []string{"test", VirtualNetworkRefUUID}})
	VirtualNetworkCreateRef = append(VirtualNetworkCreateRef,
		&models.InstanceIPVirtualNetworkRef{UUID: VirtualNetworkRefUUID2, To: []string{"test", VirtualNetworkRefUUID2}})
	model.VirtualNetworkRefs = VirtualNetworkCreateRef

	var VirtualMachineInterfaceCreateRef []*models.InstanceIPVirtualMachineInterfaceRef
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
		&models.InstanceIPVirtualMachineInterfaceRef{UUID: VirtualMachineInterfaceRefUUID, To: []string{"test", VirtualMachineInterfaceRefUUID}})
	VirtualMachineInterfaceCreateRef = append(VirtualMachineInterfaceCreateRef,
		&models.InstanceIPVirtualMachineInterfaceRef{UUID: VirtualMachineInterfaceRefUUID2, To: []string{"test", VirtualMachineInterfaceRefUUID2}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfaceCreateRef

	var PhysicalRouterCreateRef []*models.InstanceIPPhysicalRouterRef
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
		&models.InstanceIPPhysicalRouterRef{UUID: PhysicalRouterRefUUID, To: []string{"test", PhysicalRouterRefUUID}})
	PhysicalRouterCreateRef = append(PhysicalRouterCreateRef,
		&models.InstanceIPPhysicalRouterRef{UUID: PhysicalRouterRefUUID2, To: []string{"test", PhysicalRouterRefUUID2}})
	model.PhysicalRouterRefs = PhysicalRouterCreateRef

	var VirtualRouterCreateRef []*models.InstanceIPVirtualRouterRef
	var VirtualRouterRefModel *models.VirtualRouter

	VirtualRouterRefUUID := uuid.NewV4().String()
	VirtualRouterRefUUID1 := uuid.NewV4().String()
	VirtualRouterRefUUID2 := uuid.NewV4().String()

	VirtualRouterRefModel = models.MakeVirtualRouter()
	VirtualRouterRefModel.UUID = VirtualRouterRefUUID
	VirtualRouterRefModel.FQName = []string{"test", VirtualRouterRefUUID}
	_, err = db.CreateVirtualRouter(ctx, &models.CreateVirtualRouterRequest{
		VirtualRouter: VirtualRouterRefModel,
	})
	VirtualRouterRefModel.UUID = VirtualRouterRefUUID1
	VirtualRouterRefModel.FQName = []string{"test", VirtualRouterRefUUID1}
	_, err = db.CreateVirtualRouter(ctx, &models.CreateVirtualRouterRequest{
		VirtualRouter: VirtualRouterRefModel,
	})
	VirtualRouterRefModel.UUID = VirtualRouterRefUUID2
	VirtualRouterRefModel.FQName = []string{"test", VirtualRouterRefUUID2}
	_, err = db.CreateVirtualRouter(ctx, &models.CreateVirtualRouterRequest{
		VirtualRouter: VirtualRouterRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualRouterCreateRef = append(VirtualRouterCreateRef,
		&models.InstanceIPVirtualRouterRef{UUID: VirtualRouterRefUUID, To: []string{"test", VirtualRouterRefUUID}})
	VirtualRouterCreateRef = append(VirtualRouterCreateRef,
		&models.InstanceIPVirtualRouterRef{UUID: VirtualRouterRefUUID2, To: []string{"test", VirtualRouterRefUUID2}})
	model.VirtualRouterRefs = VirtualRouterCreateRef

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

	_, err = db.CreateInstanceIP(ctx,
		&models.CreateInstanceIPRequest{
			InstanceIP: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListInstanceIP(ctx, &models.ListInstanceIPRequest{
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
	if len(response.InstanceIPs) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteInstanceIP(ctxDemo,
		&models.DeleteInstanceIPRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateInstanceIP(ctx,
		&models.CreateInstanceIPRequest{
			InstanceIP: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteInstanceIP(ctx,
		&models.DeleteInstanceIPRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetInstanceIP(ctx, &models.GetInstanceIPRequest{
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
