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

func TestLogicalRouter(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeLogicalRouter()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var VirtualMachineInterfaceCreateRef []*models.LogicalRouterVirtualMachineInterfaceRef
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
		&models.LogicalRouterVirtualMachineInterfaceRef{UUID: VirtualMachineInterfaceRefUUID, To: []string{"test", VirtualMachineInterfaceRefUUID}})
	VirtualMachineInterfaceCreateRef = append(VirtualMachineInterfaceCreateRef,
		&models.LogicalRouterVirtualMachineInterfaceRef{UUID: VirtualMachineInterfaceRefUUID2, To: []string{"test", VirtualMachineInterfaceRefUUID2}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfaceCreateRef

	var ServiceInstanceCreateRef []*models.LogicalRouterServiceInstanceRef
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
		&models.LogicalRouterServiceInstanceRef{UUID: ServiceInstanceRefUUID, To: []string{"test", ServiceInstanceRefUUID}})
	ServiceInstanceCreateRef = append(ServiceInstanceCreateRef,
		&models.LogicalRouterServiceInstanceRef{UUID: ServiceInstanceRefUUID2, To: []string{"test", ServiceInstanceRefUUID2}})
	model.ServiceInstanceRefs = ServiceInstanceCreateRef

	var RouteTableCreateRef []*models.LogicalRouterRouteTableRef
	var RouteTableRefModel *models.RouteTable

	RouteTableRefUUID := uuid.NewV4().String()
	RouteTableRefUUID1 := uuid.NewV4().String()
	RouteTableRefUUID2 := uuid.NewV4().String()

	RouteTableRefModel = models.MakeRouteTable()
	RouteTableRefModel.UUID = RouteTableRefUUID
	RouteTableRefModel.FQName = []string{"test", RouteTableRefUUID}
	_, err = db.CreateRouteTable(ctx, &models.CreateRouteTableRequest{
		RouteTable: RouteTableRefModel,
	})
	RouteTableRefModel.UUID = RouteTableRefUUID1
	RouteTableRefModel.FQName = []string{"test", RouteTableRefUUID1}
	_, err = db.CreateRouteTable(ctx, &models.CreateRouteTableRequest{
		RouteTable: RouteTableRefModel,
	})
	RouteTableRefModel.UUID = RouteTableRefUUID2
	RouteTableRefModel.FQName = []string{"test", RouteTableRefUUID2}
	_, err = db.CreateRouteTable(ctx, &models.CreateRouteTableRequest{
		RouteTable: RouteTableRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	RouteTableCreateRef = append(RouteTableCreateRef,
		&models.LogicalRouterRouteTableRef{UUID: RouteTableRefUUID, To: []string{"test", RouteTableRefUUID}})
	RouteTableCreateRef = append(RouteTableCreateRef,
		&models.LogicalRouterRouteTableRef{UUID: RouteTableRefUUID2, To: []string{"test", RouteTableRefUUID2}})
	model.RouteTableRefs = RouteTableCreateRef

	var VirtualNetworkCreateRef []*models.LogicalRouterVirtualNetworkRef
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
		&models.LogicalRouterVirtualNetworkRef{UUID: VirtualNetworkRefUUID, To: []string{"test", VirtualNetworkRefUUID}})
	VirtualNetworkCreateRef = append(VirtualNetworkCreateRef,
		&models.LogicalRouterVirtualNetworkRef{UUID: VirtualNetworkRefUUID2, To: []string{"test", VirtualNetworkRefUUID2}})
	model.VirtualNetworkRefs = VirtualNetworkCreateRef

	var PhysicalRouterCreateRef []*models.LogicalRouterPhysicalRouterRef
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
		&models.LogicalRouterPhysicalRouterRef{UUID: PhysicalRouterRefUUID, To: []string{"test", PhysicalRouterRefUUID}})
	PhysicalRouterCreateRef = append(PhysicalRouterCreateRef,
		&models.LogicalRouterPhysicalRouterRef{UUID: PhysicalRouterRefUUID2, To: []string{"test", PhysicalRouterRefUUID2}})
	model.PhysicalRouterRefs = PhysicalRouterCreateRef

	var BGPVPNCreateRef []*models.LogicalRouterBGPVPNRef
	var BGPVPNRefModel *models.BGPVPN

	BGPVPNRefUUID := uuid.NewV4().String()
	BGPVPNRefUUID1 := uuid.NewV4().String()
	BGPVPNRefUUID2 := uuid.NewV4().String()

	BGPVPNRefModel = models.MakeBGPVPN()
	BGPVPNRefModel.UUID = BGPVPNRefUUID
	BGPVPNRefModel.FQName = []string{"test", BGPVPNRefUUID}
	_, err = db.CreateBGPVPN(ctx, &models.CreateBGPVPNRequest{
		BGPVPN: BGPVPNRefModel,
	})
	BGPVPNRefModel.UUID = BGPVPNRefUUID1
	BGPVPNRefModel.FQName = []string{"test", BGPVPNRefUUID1}
	_, err = db.CreateBGPVPN(ctx, &models.CreateBGPVPNRequest{
		BGPVPN: BGPVPNRefModel,
	})
	BGPVPNRefModel.UUID = BGPVPNRefUUID2
	BGPVPNRefModel.FQName = []string{"test", BGPVPNRefUUID2}
	_, err = db.CreateBGPVPN(ctx, &models.CreateBGPVPNRequest{
		BGPVPN: BGPVPNRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BGPVPNCreateRef = append(BGPVPNCreateRef,
		&models.LogicalRouterBGPVPNRef{UUID: BGPVPNRefUUID, To: []string{"test", BGPVPNRefUUID}})
	BGPVPNCreateRef = append(BGPVPNCreateRef,
		&models.LogicalRouterBGPVPNRef{UUID: BGPVPNRefUUID2, To: []string{"test", BGPVPNRefUUID2}})
	model.BGPVPNRefs = BGPVPNCreateRef

	var RouteTargetCreateRef []*models.LogicalRouterRouteTargetRef
	var RouteTargetRefModel *models.RouteTarget

	RouteTargetRefUUID := uuid.NewV4().String()
	RouteTargetRefUUID1 := uuid.NewV4().String()
	RouteTargetRefUUID2 := uuid.NewV4().String()

	RouteTargetRefModel = models.MakeRouteTarget()
	RouteTargetRefModel.UUID = RouteTargetRefUUID
	RouteTargetRefModel.FQName = []string{"test", RouteTargetRefUUID}
	_, err = db.CreateRouteTarget(ctx, &models.CreateRouteTargetRequest{
		RouteTarget: RouteTargetRefModel,
	})
	RouteTargetRefModel.UUID = RouteTargetRefUUID1
	RouteTargetRefModel.FQName = []string{"test", RouteTargetRefUUID1}
	_, err = db.CreateRouteTarget(ctx, &models.CreateRouteTargetRequest{
		RouteTarget: RouteTargetRefModel,
	})
	RouteTargetRefModel.UUID = RouteTargetRefUUID2
	RouteTargetRefModel.FQName = []string{"test", RouteTargetRefUUID2}
	_, err = db.CreateRouteTarget(ctx, &models.CreateRouteTargetRequest{
		RouteTarget: RouteTargetRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	RouteTargetCreateRef = append(RouteTargetCreateRef,
		&models.LogicalRouterRouteTargetRef{UUID: RouteTargetRefUUID, To: []string{"test", RouteTargetRefUUID}})
	RouteTargetCreateRef = append(RouteTargetCreateRef,
		&models.LogicalRouterRouteTargetRef{UUID: RouteTargetRefUUID2, To: []string{"test", RouteTargetRefUUID2}})
	model.RouteTargetRefs = RouteTargetCreateRef

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

	_, err = db.CreateLogicalRouter(ctx,
		&models.CreateLogicalRouterRequest{
			LogicalRouter: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListLogicalRouter(ctx, &models.ListLogicalRouterRequest{
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
	if len(response.LogicalRouters) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteLogicalRouter(ctxDemo,
		&models.DeleteLogicalRouterRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateLogicalRouter(ctx,
		&models.CreateLogicalRouterRequest{
			LogicalRouter: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteLogicalRouter(ctx,
		&models.DeleteLogicalRouterRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetLogicalRouter(ctx, &models.GetLogicalRouterRequest{
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
