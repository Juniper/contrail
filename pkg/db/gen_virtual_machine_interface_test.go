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

func TestVirtualMachineInterface(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeVirtualMachineInterface()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var RoutingInstanceCreateRef []*models.VirtualMachineInterfaceRoutingInstanceRef
	var RoutingInstanceRefModel *models.RoutingInstance

	RoutingInstanceRefUUID := uuid.NewV4().String()
	RoutingInstanceRefUUID1 := uuid.NewV4().String()
	RoutingInstanceRefUUID2 := uuid.NewV4().String()

	RoutingInstanceRefModel = models.MakeRoutingInstance()
	RoutingInstanceRefModel.UUID = RoutingInstanceRefUUID
	RoutingInstanceRefModel.FQName = []string{"test", RoutingInstanceRefUUID}
	_, err = db.CreateRoutingInstance(ctx, &models.CreateRoutingInstanceRequest{
		RoutingInstance: RoutingInstanceRefModel,
	})
	RoutingInstanceRefModel.UUID = RoutingInstanceRefUUID1
	RoutingInstanceRefModel.FQName = []string{"test", RoutingInstanceRefUUID1}
	_, err = db.CreateRoutingInstance(ctx, &models.CreateRoutingInstanceRequest{
		RoutingInstance: RoutingInstanceRefModel,
	})
	RoutingInstanceRefModel.UUID = RoutingInstanceRefUUID2
	RoutingInstanceRefModel.FQName = []string{"test", RoutingInstanceRefUUID2}
	_, err = db.CreateRoutingInstance(ctx, &models.CreateRoutingInstanceRequest{
		RoutingInstance: RoutingInstanceRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	RoutingInstanceCreateRef = append(RoutingInstanceCreateRef,
		&models.VirtualMachineInterfaceRoutingInstanceRef{UUID: RoutingInstanceRefUUID, To: []string{"test", RoutingInstanceRefUUID}})
	RoutingInstanceCreateRef = append(RoutingInstanceCreateRef,
		&models.VirtualMachineInterfaceRoutingInstanceRef{UUID: RoutingInstanceRefUUID2, To: []string{"test", RoutingInstanceRefUUID2}})
	model.RoutingInstanceRefs = RoutingInstanceCreateRef

	var PortTupleCreateRef []*models.VirtualMachineInterfacePortTupleRef
	var PortTupleRefModel *models.PortTuple

	PortTupleRefUUID := uuid.NewV4().String()
	PortTupleRefUUID1 := uuid.NewV4().String()
	PortTupleRefUUID2 := uuid.NewV4().String()

	PortTupleRefModel = models.MakePortTuple()
	PortTupleRefModel.UUID = PortTupleRefUUID
	PortTupleRefModel.FQName = []string{"test", PortTupleRefUUID}
	_, err = db.CreatePortTuple(ctx, &models.CreatePortTupleRequest{
		PortTuple: PortTupleRefModel,
	})
	PortTupleRefModel.UUID = PortTupleRefUUID1
	PortTupleRefModel.FQName = []string{"test", PortTupleRefUUID1}
	_, err = db.CreatePortTuple(ctx, &models.CreatePortTupleRequest{
		PortTuple: PortTupleRefModel,
	})
	PortTupleRefModel.UUID = PortTupleRefUUID2
	PortTupleRefModel.FQName = []string{"test", PortTupleRefUUID2}
	_, err = db.CreatePortTuple(ctx, &models.CreatePortTupleRequest{
		PortTuple: PortTupleRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	PortTupleCreateRef = append(PortTupleCreateRef,
		&models.VirtualMachineInterfacePortTupleRef{UUID: PortTupleRefUUID, To: []string{"test", PortTupleRefUUID}})
	PortTupleCreateRef = append(PortTupleCreateRef,
		&models.VirtualMachineInterfacePortTupleRef{UUID: PortTupleRefUUID2, To: []string{"test", PortTupleRefUUID2}})
	model.PortTupleRefs = PortTupleCreateRef

	var PhysicalInterfaceCreateRef []*models.VirtualMachineInterfacePhysicalInterfaceRef
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
		&models.VirtualMachineInterfacePhysicalInterfaceRef{UUID: PhysicalInterfaceRefUUID, To: []string{"test", PhysicalInterfaceRefUUID}})
	PhysicalInterfaceCreateRef = append(PhysicalInterfaceCreateRef,
		&models.VirtualMachineInterfacePhysicalInterfaceRef{UUID: PhysicalInterfaceRefUUID2, To: []string{"test", PhysicalInterfaceRefUUID2}})
	model.PhysicalInterfaceRefs = PhysicalInterfaceCreateRef

	var ServiceHealthCheckCreateRef []*models.VirtualMachineInterfaceServiceHealthCheckRef
	var ServiceHealthCheckRefModel *models.ServiceHealthCheck

	ServiceHealthCheckRefUUID := uuid.NewV4().String()
	ServiceHealthCheckRefUUID1 := uuid.NewV4().String()
	ServiceHealthCheckRefUUID2 := uuid.NewV4().String()

	ServiceHealthCheckRefModel = models.MakeServiceHealthCheck()
	ServiceHealthCheckRefModel.UUID = ServiceHealthCheckRefUUID
	ServiceHealthCheckRefModel.FQName = []string{"test", ServiceHealthCheckRefUUID}
	_, err = db.CreateServiceHealthCheck(ctx, &models.CreateServiceHealthCheckRequest{
		ServiceHealthCheck: ServiceHealthCheckRefModel,
	})
	ServiceHealthCheckRefModel.UUID = ServiceHealthCheckRefUUID1
	ServiceHealthCheckRefModel.FQName = []string{"test", ServiceHealthCheckRefUUID1}
	_, err = db.CreateServiceHealthCheck(ctx, &models.CreateServiceHealthCheckRequest{
		ServiceHealthCheck: ServiceHealthCheckRefModel,
	})
	ServiceHealthCheckRefModel.UUID = ServiceHealthCheckRefUUID2
	ServiceHealthCheckRefModel.FQName = []string{"test", ServiceHealthCheckRefUUID2}
	_, err = db.CreateServiceHealthCheck(ctx, &models.CreateServiceHealthCheckRequest{
		ServiceHealthCheck: ServiceHealthCheckRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceHealthCheckCreateRef = append(ServiceHealthCheckCreateRef,
		&models.VirtualMachineInterfaceServiceHealthCheckRef{UUID: ServiceHealthCheckRefUUID, To: []string{"test", ServiceHealthCheckRefUUID}})
	ServiceHealthCheckCreateRef = append(ServiceHealthCheckCreateRef,
		&models.VirtualMachineInterfaceServiceHealthCheckRef{UUID: ServiceHealthCheckRefUUID2, To: []string{"test", ServiceHealthCheckRefUUID2}})
	model.ServiceHealthCheckRefs = ServiceHealthCheckCreateRef

	var VirtualNetworkCreateRef []*models.VirtualMachineInterfaceVirtualNetworkRef
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
		&models.VirtualMachineInterfaceVirtualNetworkRef{UUID: VirtualNetworkRefUUID, To: []string{"test", VirtualNetworkRefUUID}})
	VirtualNetworkCreateRef = append(VirtualNetworkCreateRef,
		&models.VirtualMachineInterfaceVirtualNetworkRef{UUID: VirtualNetworkRefUUID2, To: []string{"test", VirtualNetworkRefUUID2}})
	model.VirtualNetworkRefs = VirtualNetworkCreateRef

	var InterfaceRouteTableCreateRef []*models.VirtualMachineInterfaceInterfaceRouteTableRef
	var InterfaceRouteTableRefModel *models.InterfaceRouteTable

	InterfaceRouteTableRefUUID := uuid.NewV4().String()
	InterfaceRouteTableRefUUID1 := uuid.NewV4().String()
	InterfaceRouteTableRefUUID2 := uuid.NewV4().String()

	InterfaceRouteTableRefModel = models.MakeInterfaceRouteTable()
	InterfaceRouteTableRefModel.UUID = InterfaceRouteTableRefUUID
	InterfaceRouteTableRefModel.FQName = []string{"test", InterfaceRouteTableRefUUID}
	_, err = db.CreateInterfaceRouteTable(ctx, &models.CreateInterfaceRouteTableRequest{
		InterfaceRouteTable: InterfaceRouteTableRefModel,
	})
	InterfaceRouteTableRefModel.UUID = InterfaceRouteTableRefUUID1
	InterfaceRouteTableRefModel.FQName = []string{"test", InterfaceRouteTableRefUUID1}
	_, err = db.CreateInterfaceRouteTable(ctx, &models.CreateInterfaceRouteTableRequest{
		InterfaceRouteTable: InterfaceRouteTableRefModel,
	})
	InterfaceRouteTableRefModel.UUID = InterfaceRouteTableRefUUID2
	InterfaceRouteTableRefModel.FQName = []string{"test", InterfaceRouteTableRefUUID2}
	_, err = db.CreateInterfaceRouteTable(ctx, &models.CreateInterfaceRouteTableRequest{
		InterfaceRouteTable: InterfaceRouteTableRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	InterfaceRouteTableCreateRef = append(InterfaceRouteTableCreateRef,
		&models.VirtualMachineInterfaceInterfaceRouteTableRef{UUID: InterfaceRouteTableRefUUID, To: []string{"test", InterfaceRouteTableRefUUID}})
	InterfaceRouteTableCreateRef = append(InterfaceRouteTableCreateRef,
		&models.VirtualMachineInterfaceInterfaceRouteTableRef{UUID: InterfaceRouteTableRefUUID2, To: []string{"test", InterfaceRouteTableRefUUID2}})
	model.InterfaceRouteTableRefs = InterfaceRouteTableCreateRef

	var BGPRouterCreateRef []*models.VirtualMachineInterfaceBGPRouterRef
	var BGPRouterRefModel *models.BGPRouter

	BGPRouterRefUUID := uuid.NewV4().String()
	BGPRouterRefUUID1 := uuid.NewV4().String()
	BGPRouterRefUUID2 := uuid.NewV4().String()

	BGPRouterRefModel = models.MakeBGPRouter()
	BGPRouterRefModel.UUID = BGPRouterRefUUID
	BGPRouterRefModel.FQName = []string{"test", BGPRouterRefUUID}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterRefModel,
	})
	BGPRouterRefModel.UUID = BGPRouterRefUUID1
	BGPRouterRefModel.FQName = []string{"test", BGPRouterRefUUID1}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterRefModel,
	})
	BGPRouterRefModel.UUID = BGPRouterRefUUID2
	BGPRouterRefModel.FQName = []string{"test", BGPRouterRefUUID2}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BGPRouterCreateRef = append(BGPRouterCreateRef,
		&models.VirtualMachineInterfaceBGPRouterRef{UUID: BGPRouterRefUUID, To: []string{"test", BGPRouterRefUUID}})
	BGPRouterCreateRef = append(BGPRouterCreateRef,
		&models.VirtualMachineInterfaceBGPRouterRef{UUID: BGPRouterRefUUID2, To: []string{"test", BGPRouterRefUUID2}})
	model.BGPRouterRefs = BGPRouterCreateRef

	var SecurityLoggingObjectCreateRef []*models.VirtualMachineInterfaceSecurityLoggingObjectRef
	var SecurityLoggingObjectRefModel *models.SecurityLoggingObject

	SecurityLoggingObjectRefUUID := uuid.NewV4().String()
	SecurityLoggingObjectRefUUID1 := uuid.NewV4().String()
	SecurityLoggingObjectRefUUID2 := uuid.NewV4().String()

	SecurityLoggingObjectRefModel = models.MakeSecurityLoggingObject()
	SecurityLoggingObjectRefModel.UUID = SecurityLoggingObjectRefUUID
	SecurityLoggingObjectRefModel.FQName = []string{"test", SecurityLoggingObjectRefUUID}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectRefModel,
	})
	SecurityLoggingObjectRefModel.UUID = SecurityLoggingObjectRefUUID1
	SecurityLoggingObjectRefModel.FQName = []string{"test", SecurityLoggingObjectRefUUID1}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectRefModel,
	})
	SecurityLoggingObjectRefModel.UUID = SecurityLoggingObjectRefUUID2
	SecurityLoggingObjectRefModel.FQName = []string{"test", SecurityLoggingObjectRefUUID2}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityLoggingObjectCreateRef = append(SecurityLoggingObjectCreateRef,
		&models.VirtualMachineInterfaceSecurityLoggingObjectRef{UUID: SecurityLoggingObjectRefUUID, To: []string{"test", SecurityLoggingObjectRefUUID}})
	SecurityLoggingObjectCreateRef = append(SecurityLoggingObjectCreateRef,
		&models.VirtualMachineInterfaceSecurityLoggingObjectRef{UUID: SecurityLoggingObjectRefUUID2, To: []string{"test", SecurityLoggingObjectRefUUID2}})
	model.SecurityLoggingObjectRefs = SecurityLoggingObjectCreateRef

	var QosConfigCreateRef []*models.VirtualMachineInterfaceQosConfigRef
	var QosConfigRefModel *models.QosConfig

	QosConfigRefUUID := uuid.NewV4().String()
	QosConfigRefUUID1 := uuid.NewV4().String()
	QosConfigRefUUID2 := uuid.NewV4().String()

	QosConfigRefModel = models.MakeQosConfig()
	QosConfigRefModel.UUID = QosConfigRefUUID
	QosConfigRefModel.FQName = []string{"test", QosConfigRefUUID}
	_, err = db.CreateQosConfig(ctx, &models.CreateQosConfigRequest{
		QosConfig: QosConfigRefModel,
	})
	QosConfigRefModel.UUID = QosConfigRefUUID1
	QosConfigRefModel.FQName = []string{"test", QosConfigRefUUID1}
	_, err = db.CreateQosConfig(ctx, &models.CreateQosConfigRequest{
		QosConfig: QosConfigRefModel,
	})
	QosConfigRefModel.UUID = QosConfigRefUUID2
	QosConfigRefModel.FQName = []string{"test", QosConfigRefUUID2}
	_, err = db.CreateQosConfig(ctx, &models.CreateQosConfigRequest{
		QosConfig: QosConfigRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	QosConfigCreateRef = append(QosConfigCreateRef,
		&models.VirtualMachineInterfaceQosConfigRef{UUID: QosConfigRefUUID, To: []string{"test", QosConfigRefUUID}})
	QosConfigCreateRef = append(QosConfigCreateRef,
		&models.VirtualMachineInterfaceQosConfigRef{UUID: QosConfigRefUUID2, To: []string{"test", QosConfigRefUUID2}})
	model.QosConfigRefs = QosConfigCreateRef

	var SecurityGroupCreateRef []*models.VirtualMachineInterfaceSecurityGroupRef
	var SecurityGroupRefModel *models.SecurityGroup

	SecurityGroupRefUUID := uuid.NewV4().String()
	SecurityGroupRefUUID1 := uuid.NewV4().String()
	SecurityGroupRefUUID2 := uuid.NewV4().String()

	SecurityGroupRefModel = models.MakeSecurityGroup()
	SecurityGroupRefModel.UUID = SecurityGroupRefUUID
	SecurityGroupRefModel.FQName = []string{"test", SecurityGroupRefUUID}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGroupRefModel,
	})
	SecurityGroupRefModel.UUID = SecurityGroupRefUUID1
	SecurityGroupRefModel.FQName = []string{"test", SecurityGroupRefUUID1}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGroupRefModel,
	})
	SecurityGroupRefModel.UUID = SecurityGroupRefUUID2
	SecurityGroupRefModel.FQName = []string{"test", SecurityGroupRefUUID2}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGroupRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityGroupCreateRef = append(SecurityGroupCreateRef,
		&models.VirtualMachineInterfaceSecurityGroupRef{UUID: SecurityGroupRefUUID, To: []string{"test", SecurityGroupRefUUID}})
	SecurityGroupCreateRef = append(SecurityGroupCreateRef,
		&models.VirtualMachineInterfaceSecurityGroupRef{UUID: SecurityGroupRefUUID2, To: []string{"test", SecurityGroupRefUUID2}})
	model.SecurityGroupRefs = SecurityGroupCreateRef

	var ServiceEndpointCreateRef []*models.VirtualMachineInterfaceServiceEndpointRef
	var ServiceEndpointRefModel *models.ServiceEndpoint

	ServiceEndpointRefUUID := uuid.NewV4().String()
	ServiceEndpointRefUUID1 := uuid.NewV4().String()
	ServiceEndpointRefUUID2 := uuid.NewV4().String()

	ServiceEndpointRefModel = models.MakeServiceEndpoint()
	ServiceEndpointRefModel.UUID = ServiceEndpointRefUUID
	ServiceEndpointRefModel.FQName = []string{"test", ServiceEndpointRefUUID}
	_, err = db.CreateServiceEndpoint(ctx, &models.CreateServiceEndpointRequest{
		ServiceEndpoint: ServiceEndpointRefModel,
	})
	ServiceEndpointRefModel.UUID = ServiceEndpointRefUUID1
	ServiceEndpointRefModel.FQName = []string{"test", ServiceEndpointRefUUID1}
	_, err = db.CreateServiceEndpoint(ctx, &models.CreateServiceEndpointRequest{
		ServiceEndpoint: ServiceEndpointRefModel,
	})
	ServiceEndpointRefModel.UUID = ServiceEndpointRefUUID2
	ServiceEndpointRefModel.FQName = []string{"test", ServiceEndpointRefUUID2}
	_, err = db.CreateServiceEndpoint(ctx, &models.CreateServiceEndpointRequest{
		ServiceEndpoint: ServiceEndpointRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceEndpointCreateRef = append(ServiceEndpointCreateRef,
		&models.VirtualMachineInterfaceServiceEndpointRef{UUID: ServiceEndpointRefUUID, To: []string{"test", ServiceEndpointRefUUID}})
	ServiceEndpointCreateRef = append(ServiceEndpointCreateRef,
		&models.VirtualMachineInterfaceServiceEndpointRef{UUID: ServiceEndpointRefUUID2, To: []string{"test", ServiceEndpointRefUUID2}})
	model.ServiceEndpointRefs = ServiceEndpointCreateRef

	var VirtualMachineInterfaceCreateRef []*models.VirtualMachineInterfaceVirtualMachineInterfaceRef
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
		&models.VirtualMachineInterfaceVirtualMachineInterfaceRef{UUID: VirtualMachineInterfaceRefUUID, To: []string{"test", VirtualMachineInterfaceRefUUID}})
	VirtualMachineInterfaceCreateRef = append(VirtualMachineInterfaceCreateRef,
		&models.VirtualMachineInterfaceVirtualMachineInterfaceRef{UUID: VirtualMachineInterfaceRefUUID2, To: []string{"test", VirtualMachineInterfaceRefUUID2}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfaceCreateRef

	var BridgeDomainCreateRef []*models.VirtualMachineInterfaceBridgeDomainRef
	var BridgeDomainRefModel *models.BridgeDomain

	BridgeDomainRefUUID := uuid.NewV4().String()
	BridgeDomainRefUUID1 := uuid.NewV4().String()
	BridgeDomainRefUUID2 := uuid.NewV4().String()

	BridgeDomainRefModel = models.MakeBridgeDomain()
	BridgeDomainRefModel.UUID = BridgeDomainRefUUID
	BridgeDomainRefModel.FQName = []string{"test", BridgeDomainRefUUID}
	_, err = db.CreateBridgeDomain(ctx, &models.CreateBridgeDomainRequest{
		BridgeDomain: BridgeDomainRefModel,
	})
	BridgeDomainRefModel.UUID = BridgeDomainRefUUID1
	BridgeDomainRefModel.FQName = []string{"test", BridgeDomainRefUUID1}
	_, err = db.CreateBridgeDomain(ctx, &models.CreateBridgeDomainRequest{
		BridgeDomain: BridgeDomainRefModel,
	})
	BridgeDomainRefModel.UUID = BridgeDomainRefUUID2
	BridgeDomainRefModel.FQName = []string{"test", BridgeDomainRefUUID2}
	_, err = db.CreateBridgeDomain(ctx, &models.CreateBridgeDomainRequest{
		BridgeDomain: BridgeDomainRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BridgeDomainCreateRef = append(BridgeDomainCreateRef,
		&models.VirtualMachineInterfaceBridgeDomainRef{UUID: BridgeDomainRefUUID, To: []string{"test", BridgeDomainRefUUID}})
	BridgeDomainCreateRef = append(BridgeDomainCreateRef,
		&models.VirtualMachineInterfaceBridgeDomainRef{UUID: BridgeDomainRefUUID2, To: []string{"test", BridgeDomainRefUUID2}})
	model.BridgeDomainRefs = BridgeDomainCreateRef

	var VirtualMachineCreateRef []*models.VirtualMachineInterfaceVirtualMachineRef
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
		&models.VirtualMachineInterfaceVirtualMachineRef{UUID: VirtualMachineRefUUID, To: []string{"test", VirtualMachineRefUUID}})
	VirtualMachineCreateRef = append(VirtualMachineCreateRef,
		&models.VirtualMachineInterfaceVirtualMachineRef{UUID: VirtualMachineRefUUID2, To: []string{"test", VirtualMachineRefUUID2}})
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

	_, err = db.CreateVirtualMachineInterface(ctx,
		&models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListVirtualMachineInterface(ctx, &models.ListVirtualMachineInterfaceRequest{
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
	if len(response.VirtualMachineInterfaces) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteVirtualMachineInterface(ctxDemo,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateVirtualMachineInterface(ctx,
		&models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteVirtualMachineInterface(ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetVirtualMachineInterface(ctx, &models.GetVirtualMachineInterfaceRequest{
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
