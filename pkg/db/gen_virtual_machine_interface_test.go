// nolint
package db

import (
	"context"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestVirtualMachineInterface(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "virtual_machine_interface")
	// mutexProject := UseTable(db.DB, "virtual_machine_interface")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeVirtualMachineInterface()
	model.UUID = "virtual_machine_interface_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "virtual_machine_interface_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var VirtualMachineInterfacecreateref []*models.VirtualMachineInterfaceVirtualMachineInterfaceRef
	var VirtualMachineInterfacerefModel *models.VirtualMachineInterface
	VirtualMachineInterfacerefModel = models.MakeVirtualMachineInterface()
	VirtualMachineInterfacerefModel.UUID = "virtual_machine_interface_virtual_machine_interface_ref_uuid"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "virtual_machine_interface_virtual_machine_interface_ref_uuid"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	VirtualMachineInterfacerefModel.UUID = "virtual_machine_interface_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "virtual_machine_interface_virtual_machine_interface_ref_uuid1"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	VirtualMachineInterfacerefModel.UUID = "virtual_machine_interface_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "virtual_machine_interface_virtual_machine_interface_ref_uuid2"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.VirtualMachineInterfaceVirtualMachineInterfaceRef{UUID: "virtual_machine_interface_virtual_machine_interface_ref_uuid", To: []string{"test", "virtual_machine_interface_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.VirtualMachineInterfaceVirtualMachineInterfaceRef{UUID: "virtual_machine_interface_virtual_machine_interface_ref_uuid2", To: []string{"test", "virtual_machine_interface_virtual_machine_interface_ref_uuid2"}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref

	var InterfaceRouteTablecreateref []*models.VirtualMachineInterfaceInterfaceRouteTableRef
	var InterfaceRouteTablerefModel *models.InterfaceRouteTable
	InterfaceRouteTablerefModel = models.MakeInterfaceRouteTable()
	InterfaceRouteTablerefModel.UUID = "virtual_machine_interface_interface_route_table_ref_uuid"
	InterfaceRouteTablerefModel.FQName = []string{"test", "virtual_machine_interface_interface_route_table_ref_uuid"}
	_, err = db.CreateInterfaceRouteTable(ctx, &models.CreateInterfaceRouteTableRequest{
		InterfaceRouteTable: InterfaceRouteTablerefModel,
	})
	InterfaceRouteTablerefModel.UUID = "virtual_machine_interface_interface_route_table_ref_uuid1"
	InterfaceRouteTablerefModel.FQName = []string{"test", "virtual_machine_interface_interface_route_table_ref_uuid1"}
	_, err = db.CreateInterfaceRouteTable(ctx, &models.CreateInterfaceRouteTableRequest{
		InterfaceRouteTable: InterfaceRouteTablerefModel,
	})
	InterfaceRouteTablerefModel.UUID = "virtual_machine_interface_interface_route_table_ref_uuid2"
	InterfaceRouteTablerefModel.FQName = []string{"test", "virtual_machine_interface_interface_route_table_ref_uuid2"}
	_, err = db.CreateInterfaceRouteTable(ctx, &models.CreateInterfaceRouteTableRequest{
		InterfaceRouteTable: InterfaceRouteTablerefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	InterfaceRouteTablecreateref = append(InterfaceRouteTablecreateref, &models.VirtualMachineInterfaceInterfaceRouteTableRef{UUID: "virtual_machine_interface_interface_route_table_ref_uuid", To: []string{"test", "virtual_machine_interface_interface_route_table_ref_uuid"}})
	InterfaceRouteTablecreateref = append(InterfaceRouteTablecreateref, &models.VirtualMachineInterfaceInterfaceRouteTableRef{UUID: "virtual_machine_interface_interface_route_table_ref_uuid2", To: []string{"test", "virtual_machine_interface_interface_route_table_ref_uuid2"}})
	model.InterfaceRouteTableRefs = InterfaceRouteTablecreateref

	var ServiceHealthCheckcreateref []*models.VirtualMachineInterfaceServiceHealthCheckRef
	var ServiceHealthCheckrefModel *models.ServiceHealthCheck
	ServiceHealthCheckrefModel = models.MakeServiceHealthCheck()
	ServiceHealthCheckrefModel.UUID = "virtual_machine_interface_service_health_check_ref_uuid"
	ServiceHealthCheckrefModel.FQName = []string{"test", "virtual_machine_interface_service_health_check_ref_uuid"}
	_, err = db.CreateServiceHealthCheck(ctx, &models.CreateServiceHealthCheckRequest{
		ServiceHealthCheck: ServiceHealthCheckrefModel,
	})
	ServiceHealthCheckrefModel.UUID = "virtual_machine_interface_service_health_check_ref_uuid1"
	ServiceHealthCheckrefModel.FQName = []string{"test", "virtual_machine_interface_service_health_check_ref_uuid1"}
	_, err = db.CreateServiceHealthCheck(ctx, &models.CreateServiceHealthCheckRequest{
		ServiceHealthCheck: ServiceHealthCheckrefModel,
	})
	ServiceHealthCheckrefModel.UUID = "virtual_machine_interface_service_health_check_ref_uuid2"
	ServiceHealthCheckrefModel.FQName = []string{"test", "virtual_machine_interface_service_health_check_ref_uuid2"}
	_, err = db.CreateServiceHealthCheck(ctx, &models.CreateServiceHealthCheckRequest{
		ServiceHealthCheck: ServiceHealthCheckrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceHealthCheckcreateref = append(ServiceHealthCheckcreateref, &models.VirtualMachineInterfaceServiceHealthCheckRef{UUID: "virtual_machine_interface_service_health_check_ref_uuid", To: []string{"test", "virtual_machine_interface_service_health_check_ref_uuid"}})
	ServiceHealthCheckcreateref = append(ServiceHealthCheckcreateref, &models.VirtualMachineInterfaceServiceHealthCheckRef{UUID: "virtual_machine_interface_service_health_check_ref_uuid2", To: []string{"test", "virtual_machine_interface_service_health_check_ref_uuid2"}})
	model.ServiceHealthCheckRefs = ServiceHealthCheckcreateref

	var VirtualNetworkcreateref []*models.VirtualMachineInterfaceVirtualNetworkRef
	var VirtualNetworkrefModel *models.VirtualNetwork
	VirtualNetworkrefModel = models.MakeVirtualNetwork()
	VirtualNetworkrefModel.UUID = "virtual_machine_interface_virtual_network_ref_uuid"
	VirtualNetworkrefModel.FQName = []string{"test", "virtual_machine_interface_virtual_network_ref_uuid"}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkrefModel,
	})
	VirtualNetworkrefModel.UUID = "virtual_machine_interface_virtual_network_ref_uuid1"
	VirtualNetworkrefModel.FQName = []string{"test", "virtual_machine_interface_virtual_network_ref_uuid1"}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkrefModel,
	})
	VirtualNetworkrefModel.UUID = "virtual_machine_interface_virtual_network_ref_uuid2"
	VirtualNetworkrefModel.FQName = []string{"test", "virtual_machine_interface_virtual_network_ref_uuid2"}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.VirtualMachineInterfaceVirtualNetworkRef{UUID: "virtual_machine_interface_virtual_network_ref_uuid", To: []string{"test", "virtual_machine_interface_virtual_network_ref_uuid"}})
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.VirtualMachineInterfaceVirtualNetworkRef{UUID: "virtual_machine_interface_virtual_network_ref_uuid2", To: []string{"test", "virtual_machine_interface_virtual_network_ref_uuid2"}})
	model.VirtualNetworkRefs = VirtualNetworkcreateref

	var RoutingInstancecreateref []*models.VirtualMachineInterfaceRoutingInstanceRef
	var RoutingInstancerefModel *models.RoutingInstance
	RoutingInstancerefModel = models.MakeRoutingInstance()
	RoutingInstancerefModel.UUID = "virtual_machine_interface_routing_instance_ref_uuid"
	RoutingInstancerefModel.FQName = []string{"test", "virtual_machine_interface_routing_instance_ref_uuid"}
	_, err = db.CreateRoutingInstance(ctx, &models.CreateRoutingInstanceRequest{
		RoutingInstance: RoutingInstancerefModel,
	})
	RoutingInstancerefModel.UUID = "virtual_machine_interface_routing_instance_ref_uuid1"
	RoutingInstancerefModel.FQName = []string{"test", "virtual_machine_interface_routing_instance_ref_uuid1"}
	_, err = db.CreateRoutingInstance(ctx, &models.CreateRoutingInstanceRequest{
		RoutingInstance: RoutingInstancerefModel,
	})
	RoutingInstancerefModel.UUID = "virtual_machine_interface_routing_instance_ref_uuid2"
	RoutingInstancerefModel.FQName = []string{"test", "virtual_machine_interface_routing_instance_ref_uuid2"}
	_, err = db.CreateRoutingInstance(ctx, &models.CreateRoutingInstanceRequest{
		RoutingInstance: RoutingInstancerefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	RoutingInstancecreateref = append(RoutingInstancecreateref, &models.VirtualMachineInterfaceRoutingInstanceRef{UUID: "virtual_machine_interface_routing_instance_ref_uuid", To: []string{"test", "virtual_machine_interface_routing_instance_ref_uuid"}})
	RoutingInstancecreateref = append(RoutingInstancecreateref, &models.VirtualMachineInterfaceRoutingInstanceRef{UUID: "virtual_machine_interface_routing_instance_ref_uuid2", To: []string{"test", "virtual_machine_interface_routing_instance_ref_uuid2"}})
	model.RoutingInstanceRefs = RoutingInstancecreateref

	var BridgeDomaincreateref []*models.VirtualMachineInterfaceBridgeDomainRef
	var BridgeDomainrefModel *models.BridgeDomain
	BridgeDomainrefModel = models.MakeBridgeDomain()
	BridgeDomainrefModel.UUID = "virtual_machine_interface_bridge_domain_ref_uuid"
	BridgeDomainrefModel.FQName = []string{"test", "virtual_machine_interface_bridge_domain_ref_uuid"}
	_, err = db.CreateBridgeDomain(ctx, &models.CreateBridgeDomainRequest{
		BridgeDomain: BridgeDomainrefModel,
	})
	BridgeDomainrefModel.UUID = "virtual_machine_interface_bridge_domain_ref_uuid1"
	BridgeDomainrefModel.FQName = []string{"test", "virtual_machine_interface_bridge_domain_ref_uuid1"}
	_, err = db.CreateBridgeDomain(ctx, &models.CreateBridgeDomainRequest{
		BridgeDomain: BridgeDomainrefModel,
	})
	BridgeDomainrefModel.UUID = "virtual_machine_interface_bridge_domain_ref_uuid2"
	BridgeDomainrefModel.FQName = []string{"test", "virtual_machine_interface_bridge_domain_ref_uuid2"}
	_, err = db.CreateBridgeDomain(ctx, &models.CreateBridgeDomainRequest{
		BridgeDomain: BridgeDomainrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BridgeDomaincreateref = append(BridgeDomaincreateref, &models.VirtualMachineInterfaceBridgeDomainRef{UUID: "virtual_machine_interface_bridge_domain_ref_uuid", To: []string{"test", "virtual_machine_interface_bridge_domain_ref_uuid"}})
	BridgeDomaincreateref = append(BridgeDomaincreateref, &models.VirtualMachineInterfaceBridgeDomainRef{UUID: "virtual_machine_interface_bridge_domain_ref_uuid2", To: []string{"test", "virtual_machine_interface_bridge_domain_ref_uuid2"}})
	model.BridgeDomainRefs = BridgeDomaincreateref

	var BGPRoutercreateref []*models.VirtualMachineInterfaceBGPRouterRef
	var BGPRouterrefModel *models.BGPRouter
	BGPRouterrefModel = models.MakeBGPRouter()
	BGPRouterrefModel.UUID = "virtual_machine_interface_bgp_router_ref_uuid"
	BGPRouterrefModel.FQName = []string{"test", "virtual_machine_interface_bgp_router_ref_uuid"}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterrefModel,
	})
	BGPRouterrefModel.UUID = "virtual_machine_interface_bgp_router_ref_uuid1"
	BGPRouterrefModel.FQName = []string{"test", "virtual_machine_interface_bgp_router_ref_uuid1"}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterrefModel,
	})
	BGPRouterrefModel.UUID = "virtual_machine_interface_bgp_router_ref_uuid2"
	BGPRouterrefModel.FQName = []string{"test", "virtual_machine_interface_bgp_router_ref_uuid2"}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BGPRoutercreateref = append(BGPRoutercreateref, &models.VirtualMachineInterfaceBGPRouterRef{UUID: "virtual_machine_interface_bgp_router_ref_uuid", To: []string{"test", "virtual_machine_interface_bgp_router_ref_uuid"}})
	BGPRoutercreateref = append(BGPRoutercreateref, &models.VirtualMachineInterfaceBGPRouterRef{UUID: "virtual_machine_interface_bgp_router_ref_uuid2", To: []string{"test", "virtual_machine_interface_bgp_router_ref_uuid2"}})
	model.BGPRouterRefs = BGPRoutercreateref

	var SecurityLoggingObjectcreateref []*models.VirtualMachineInterfaceSecurityLoggingObjectRef
	var SecurityLoggingObjectrefModel *models.SecurityLoggingObject
	SecurityLoggingObjectrefModel = models.MakeSecurityLoggingObject()
	SecurityLoggingObjectrefModel.UUID = "virtual_machine_interface_security_logging_object_ref_uuid"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "virtual_machine_interface_security_logging_object_ref_uuid"}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectrefModel,
	})
	SecurityLoggingObjectrefModel.UUID = "virtual_machine_interface_security_logging_object_ref_uuid1"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "virtual_machine_interface_security_logging_object_ref_uuid1"}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectrefModel,
	})
	SecurityLoggingObjectrefModel.UUID = "virtual_machine_interface_security_logging_object_ref_uuid2"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "virtual_machine_interface_security_logging_object_ref_uuid2"}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.VirtualMachineInterfaceSecurityLoggingObjectRef{UUID: "virtual_machine_interface_security_logging_object_ref_uuid", To: []string{"test", "virtual_machine_interface_security_logging_object_ref_uuid"}})
	SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.VirtualMachineInterfaceSecurityLoggingObjectRef{UUID: "virtual_machine_interface_security_logging_object_ref_uuid2", To: []string{"test", "virtual_machine_interface_security_logging_object_ref_uuid2"}})
	model.SecurityLoggingObjectRefs = SecurityLoggingObjectcreateref

	var QosConfigcreateref []*models.VirtualMachineInterfaceQosConfigRef
	var QosConfigrefModel *models.QosConfig
	QosConfigrefModel = models.MakeQosConfig()
	QosConfigrefModel.UUID = "virtual_machine_interface_qos_config_ref_uuid"
	QosConfigrefModel.FQName = []string{"test", "virtual_machine_interface_qos_config_ref_uuid"}
	_, err = db.CreateQosConfig(ctx, &models.CreateQosConfigRequest{
		QosConfig: QosConfigrefModel,
	})
	QosConfigrefModel.UUID = "virtual_machine_interface_qos_config_ref_uuid1"
	QosConfigrefModel.FQName = []string{"test", "virtual_machine_interface_qos_config_ref_uuid1"}
	_, err = db.CreateQosConfig(ctx, &models.CreateQosConfigRequest{
		QosConfig: QosConfigrefModel,
	})
	QosConfigrefModel.UUID = "virtual_machine_interface_qos_config_ref_uuid2"
	QosConfigrefModel.FQName = []string{"test", "virtual_machine_interface_qos_config_ref_uuid2"}
	_, err = db.CreateQosConfig(ctx, &models.CreateQosConfigRequest{
		QosConfig: QosConfigrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	QosConfigcreateref = append(QosConfigcreateref, &models.VirtualMachineInterfaceQosConfigRef{UUID: "virtual_machine_interface_qos_config_ref_uuid", To: []string{"test", "virtual_machine_interface_qos_config_ref_uuid"}})
	QosConfigcreateref = append(QosConfigcreateref, &models.VirtualMachineInterfaceQosConfigRef{UUID: "virtual_machine_interface_qos_config_ref_uuid2", To: []string{"test", "virtual_machine_interface_qos_config_ref_uuid2"}})
	model.QosConfigRefs = QosConfigcreateref

	var PortTuplecreateref []*models.VirtualMachineInterfacePortTupleRef
	var PortTuplerefModel *models.PortTuple
	PortTuplerefModel = models.MakePortTuple()
	PortTuplerefModel.UUID = "virtual_machine_interface_port_tuple_ref_uuid"
	PortTuplerefModel.FQName = []string{"test", "virtual_machine_interface_port_tuple_ref_uuid"}
	_, err = db.CreatePortTuple(ctx, &models.CreatePortTupleRequest{
		PortTuple: PortTuplerefModel,
	})
	PortTuplerefModel.UUID = "virtual_machine_interface_port_tuple_ref_uuid1"
	PortTuplerefModel.FQName = []string{"test", "virtual_machine_interface_port_tuple_ref_uuid1"}
	_, err = db.CreatePortTuple(ctx, &models.CreatePortTupleRequest{
		PortTuple: PortTuplerefModel,
	})
	PortTuplerefModel.UUID = "virtual_machine_interface_port_tuple_ref_uuid2"
	PortTuplerefModel.FQName = []string{"test", "virtual_machine_interface_port_tuple_ref_uuid2"}
	_, err = db.CreatePortTuple(ctx, &models.CreatePortTupleRequest{
		PortTuple: PortTuplerefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	PortTuplecreateref = append(PortTuplecreateref, &models.VirtualMachineInterfacePortTupleRef{UUID: "virtual_machine_interface_port_tuple_ref_uuid", To: []string{"test", "virtual_machine_interface_port_tuple_ref_uuid"}})
	PortTuplecreateref = append(PortTuplecreateref, &models.VirtualMachineInterfacePortTupleRef{UUID: "virtual_machine_interface_port_tuple_ref_uuid2", To: []string{"test", "virtual_machine_interface_port_tuple_ref_uuid2"}})
	model.PortTupleRefs = PortTuplecreateref

	var PhysicalInterfacecreateref []*models.VirtualMachineInterfacePhysicalInterfaceRef
	var PhysicalInterfacerefModel *models.PhysicalInterface
	PhysicalInterfacerefModel = models.MakePhysicalInterface()
	PhysicalInterfacerefModel.UUID = "virtual_machine_interface_physical_interface_ref_uuid"
	PhysicalInterfacerefModel.FQName = []string{"test", "virtual_machine_interface_physical_interface_ref_uuid"}
	_, err = db.CreatePhysicalInterface(ctx, &models.CreatePhysicalInterfaceRequest{
		PhysicalInterface: PhysicalInterfacerefModel,
	})
	PhysicalInterfacerefModel.UUID = "virtual_machine_interface_physical_interface_ref_uuid1"
	PhysicalInterfacerefModel.FQName = []string{"test", "virtual_machine_interface_physical_interface_ref_uuid1"}
	_, err = db.CreatePhysicalInterface(ctx, &models.CreatePhysicalInterfaceRequest{
		PhysicalInterface: PhysicalInterfacerefModel,
	})
	PhysicalInterfacerefModel.UUID = "virtual_machine_interface_physical_interface_ref_uuid2"
	PhysicalInterfacerefModel.FQName = []string{"test", "virtual_machine_interface_physical_interface_ref_uuid2"}
	_, err = db.CreatePhysicalInterface(ctx, &models.CreatePhysicalInterfaceRequest{
		PhysicalInterface: PhysicalInterfacerefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	PhysicalInterfacecreateref = append(PhysicalInterfacecreateref, &models.VirtualMachineInterfacePhysicalInterfaceRef{UUID: "virtual_machine_interface_physical_interface_ref_uuid", To: []string{"test", "virtual_machine_interface_physical_interface_ref_uuid"}})
	PhysicalInterfacecreateref = append(PhysicalInterfacecreateref, &models.VirtualMachineInterfacePhysicalInterfaceRef{UUID: "virtual_machine_interface_physical_interface_ref_uuid2", To: []string{"test", "virtual_machine_interface_physical_interface_ref_uuid2"}})
	model.PhysicalInterfaceRefs = PhysicalInterfacecreateref

	var VirtualMachinecreateref []*models.VirtualMachineInterfaceVirtualMachineRef
	var VirtualMachinerefModel *models.VirtualMachine
	VirtualMachinerefModel = models.MakeVirtualMachine()
	VirtualMachinerefModel.UUID = "virtual_machine_interface_virtual_machine_ref_uuid"
	VirtualMachinerefModel.FQName = []string{"test", "virtual_machine_interface_virtual_machine_ref_uuid"}
	_, err = db.CreateVirtualMachine(ctx, &models.CreateVirtualMachineRequest{
		VirtualMachine: VirtualMachinerefModel,
	})
	VirtualMachinerefModel.UUID = "virtual_machine_interface_virtual_machine_ref_uuid1"
	VirtualMachinerefModel.FQName = []string{"test", "virtual_machine_interface_virtual_machine_ref_uuid1"}
	_, err = db.CreateVirtualMachine(ctx, &models.CreateVirtualMachineRequest{
		VirtualMachine: VirtualMachinerefModel,
	})
	VirtualMachinerefModel.UUID = "virtual_machine_interface_virtual_machine_ref_uuid2"
	VirtualMachinerefModel.FQName = []string{"test", "virtual_machine_interface_virtual_machine_ref_uuid2"}
	_, err = db.CreateVirtualMachine(ctx, &models.CreateVirtualMachineRequest{
		VirtualMachine: VirtualMachinerefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachinecreateref = append(VirtualMachinecreateref, &models.VirtualMachineInterfaceVirtualMachineRef{UUID: "virtual_machine_interface_virtual_machine_ref_uuid", To: []string{"test", "virtual_machine_interface_virtual_machine_ref_uuid"}})
	VirtualMachinecreateref = append(VirtualMachinecreateref, &models.VirtualMachineInterfaceVirtualMachineRef{UUID: "virtual_machine_interface_virtual_machine_ref_uuid2", To: []string{"test", "virtual_machine_interface_virtual_machine_ref_uuid2"}})
	model.VirtualMachineRefs = VirtualMachinecreateref

	var SecurityGroupcreateref []*models.VirtualMachineInterfaceSecurityGroupRef
	var SecurityGrouprefModel *models.SecurityGroup
	SecurityGrouprefModel = models.MakeSecurityGroup()
	SecurityGrouprefModel.UUID = "virtual_machine_interface_security_group_ref_uuid"
	SecurityGrouprefModel.FQName = []string{"test", "virtual_machine_interface_security_group_ref_uuid"}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGrouprefModel,
	})
	SecurityGrouprefModel.UUID = "virtual_machine_interface_security_group_ref_uuid1"
	SecurityGrouprefModel.FQName = []string{"test", "virtual_machine_interface_security_group_ref_uuid1"}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGrouprefModel,
	})
	SecurityGrouprefModel.UUID = "virtual_machine_interface_security_group_ref_uuid2"
	SecurityGrouprefModel.FQName = []string{"test", "virtual_machine_interface_security_group_ref_uuid2"}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGrouprefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityGroupcreateref = append(SecurityGroupcreateref, &models.VirtualMachineInterfaceSecurityGroupRef{UUID: "virtual_machine_interface_security_group_ref_uuid", To: []string{"test", "virtual_machine_interface_security_group_ref_uuid"}})
	SecurityGroupcreateref = append(SecurityGroupcreateref, &models.VirtualMachineInterfaceSecurityGroupRef{UUID: "virtual_machine_interface_security_group_ref_uuid2", To: []string{"test", "virtual_machine_interface_security_group_ref_uuid2"}})
	model.SecurityGroupRefs = SecurityGroupcreateref

	var ServiceEndpointcreateref []*models.VirtualMachineInterfaceServiceEndpointRef
	var ServiceEndpointrefModel *models.ServiceEndpoint
	ServiceEndpointrefModel = models.MakeServiceEndpoint()
	ServiceEndpointrefModel.UUID = "virtual_machine_interface_service_endpoint_ref_uuid"
	ServiceEndpointrefModel.FQName = []string{"test", "virtual_machine_interface_service_endpoint_ref_uuid"}
	_, err = db.CreateServiceEndpoint(ctx, &models.CreateServiceEndpointRequest{
		ServiceEndpoint: ServiceEndpointrefModel,
	})
	ServiceEndpointrefModel.UUID = "virtual_machine_interface_service_endpoint_ref_uuid1"
	ServiceEndpointrefModel.FQName = []string{"test", "virtual_machine_interface_service_endpoint_ref_uuid1"}
	_, err = db.CreateServiceEndpoint(ctx, &models.CreateServiceEndpointRequest{
		ServiceEndpoint: ServiceEndpointrefModel,
	})
	ServiceEndpointrefModel.UUID = "virtual_machine_interface_service_endpoint_ref_uuid2"
	ServiceEndpointrefModel.FQName = []string{"test", "virtual_machine_interface_service_endpoint_ref_uuid2"}
	_, err = db.CreateServiceEndpoint(ctx, &models.CreateServiceEndpointRequest{
		ServiceEndpoint: ServiceEndpointrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceEndpointcreateref = append(ServiceEndpointcreateref, &models.VirtualMachineInterfaceServiceEndpointRef{UUID: "virtual_machine_interface_service_endpoint_ref_uuid", To: []string{"test", "virtual_machine_interface_service_endpoint_ref_uuid"}})
	ServiceEndpointcreateref = append(ServiceEndpointcreateref, &models.VirtualMachineInterfaceServiceEndpointRef{UUID: "virtual_machine_interface_service_endpoint_ref_uuid2", To: []string{"test", "virtual_machine_interface_service_endpoint_ref_uuid2"}})
	model.ServiceEndpointRefs = ServiceEndpointcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "virtual_machine_interface_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare

	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: projectModel,
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//    //populate update map
	//    updateMap := map[string]interface{}{}
	//
	//
	//    if ".VRFAssignTable.VRFAssignRule" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".VRFAssignTable.VRFAssignRule", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".VRFAssignTable.VRFAssignRule", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VlanTagBasedBridgeDomain", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.SubInterfaceVlanTag", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.ServiceInterfaceType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.LocalPreference", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress", ".", "test")
	//
	//
	//
	//    if ".VirtualMachineInterfaceMacAddresses.MacAddress" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceMacAddresses.MacAddress", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceMacAddresses.MacAddress", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    if ".VirtualMachineInterfaceHostRoutes.Route" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceHostRoutes.Route", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceHostRoutes.Route", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    if ".VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceDisablePolicy", ".", true)
	//
	//
	//
	//    if ".VirtualMachineInterfaceDHCPOptionList.DHCPOption" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceDHCPOptionList.DHCPOption", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceDHCPOptionList.DHCPOption", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualMachineInterfaceDeviceOwner", ".", "test")
	//
	//
	//
	//    if ".VirtualMachineInterfaceBindings.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceBindings.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceBindings.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    if ".VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PortSecurityEnabled", ".", true)
	//
	//
	//
	//    if ".Perms2.Share" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")
	//
	//
	//
	//    if ".FQName" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".FQName", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.SourcePort", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.SourceIP", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.IPProtocol", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.HashingConfigured", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.DestinationPort", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.DestinationIP", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ConfigurationVersion", ".", 1.0)
	//
	//
	//
	//    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "virtual_machine_interface_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var QosConfigref []interface{}
	//    QosConfigref = append(QosConfigref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_qos_config_ref_uuid", "to": []string{"test", "virtual_machine_interface_qos_config_ref_uuid"}})
	//    QosConfigref = append(QosConfigref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_qos_config_ref_uuid1", "to": []string{"test", "virtual_machine_interface_qos_config_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "QosConfigRefs", ".", QosConfigref)
	//
	//    var PortTupleref []interface{}
	//    PortTupleref = append(PortTupleref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_port_tuple_ref_uuid", "to": []string{"test", "virtual_machine_interface_port_tuple_ref_uuid"}})
	//    PortTupleref = append(PortTupleref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_port_tuple_ref_uuid1", "to": []string{"test", "virtual_machine_interface_port_tuple_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "PortTupleRefs", ".", PortTupleref)
	//
	//    var PhysicalInterfaceref []interface{}
	//    PhysicalInterfaceref = append(PhysicalInterfaceref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_physical_interface_ref_uuid", "to": []string{"test", "virtual_machine_interface_physical_interface_ref_uuid"}})
	//    PhysicalInterfaceref = append(PhysicalInterfaceref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_physical_interface_ref_uuid1", "to": []string{"test", "virtual_machine_interface_physical_interface_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "PhysicalInterfaceRefs", ".", PhysicalInterfaceref)
	//
	//    var BGPRouterref []interface{}
	//    BGPRouterref = append(BGPRouterref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_bgp_router_ref_uuid", "to": []string{"test", "virtual_machine_interface_bgp_router_ref_uuid"}})
	//    BGPRouterref = append(BGPRouterref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_bgp_router_ref_uuid1", "to": []string{"test", "virtual_machine_interface_bgp_router_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "BGPRouterRefs", ".", BGPRouterref)
	//
	//    var SecurityLoggingObjectref []interface{}
	//    SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_security_logging_object_ref_uuid", "to": []string{"test", "virtual_machine_interface_security_logging_object_ref_uuid"}})
	//    SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_security_logging_object_ref_uuid1", "to": []string{"test", "virtual_machine_interface_security_logging_object_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "SecurityLoggingObjectRefs", ".", SecurityLoggingObjectref)
	//
	//    var ServiceEndpointref []interface{}
	//    ServiceEndpointref = append(ServiceEndpointref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_service_endpoint_ref_uuid", "to": []string{"test", "virtual_machine_interface_service_endpoint_ref_uuid"}})
	//    ServiceEndpointref = append(ServiceEndpointref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_service_endpoint_ref_uuid1", "to": []string{"test", "virtual_machine_interface_service_endpoint_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ServiceEndpointRefs", ".", ServiceEndpointref)
	//
	//    var VirtualMachineref []interface{}
	//    VirtualMachineref = append(VirtualMachineref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_virtual_machine_ref_uuid", "to": []string{"test", "virtual_machine_interface_virtual_machine_ref_uuid"}})
	//    VirtualMachineref = append(VirtualMachineref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_virtual_machine_ref_uuid1", "to": []string{"test", "virtual_machine_interface_virtual_machine_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualMachineRefs", ".", VirtualMachineref)
	//
	//    var SecurityGroupref []interface{}
	//    SecurityGroupref = append(SecurityGroupref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_security_group_ref_uuid", "to": []string{"test", "virtual_machine_interface_security_group_ref_uuid"}})
	//    SecurityGroupref = append(SecurityGroupref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_security_group_ref_uuid1", "to": []string{"test", "virtual_machine_interface_security_group_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "SecurityGroupRefs", ".", SecurityGroupref)
	//
	//    var ServiceHealthCheckref []interface{}
	//    ServiceHealthCheckref = append(ServiceHealthCheckref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_service_health_check_ref_uuid", "to": []string{"test", "virtual_machine_interface_service_health_check_ref_uuid"}})
	//    ServiceHealthCheckref = append(ServiceHealthCheckref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_service_health_check_ref_uuid1", "to": []string{"test", "virtual_machine_interface_service_health_check_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ServiceHealthCheckRefs", ".", ServiceHealthCheckref)
	//
	//    var VirtualNetworkref []interface{}
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_virtual_network_ref_uuid", "to": []string{"test", "virtual_machine_interface_virtual_network_ref_uuid"}})
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_virtual_network_ref_uuid1", "to": []string{"test", "virtual_machine_interface_virtual_network_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualNetworkRefs", ".", VirtualNetworkref)
	//
	//    var VirtualMachineInterfaceref []interface{}
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_virtual_machine_interface_ref_uuid", "to": []string{"test", "virtual_machine_interface_virtual_machine_interface_ref_uuid"}})
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_virtual_machine_interface_ref_uuid1", "to": []string{"test", "virtual_machine_interface_virtual_machine_interface_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)
	//
	//    var InterfaceRouteTableref []interface{}
	//    InterfaceRouteTableref = append(InterfaceRouteTableref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_interface_route_table_ref_uuid", "to": []string{"test", "virtual_machine_interface_interface_route_table_ref_uuid"}})
	//    InterfaceRouteTableref = append(InterfaceRouteTableref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_interface_route_table_ref_uuid1", "to": []string{"test", "virtual_machine_interface_interface_route_table_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "InterfaceRouteTableRefs", ".", InterfaceRouteTableref)
	//
	//    var RoutingInstanceref []interface{}
	//    RoutingInstanceref = append(RoutingInstanceref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_routing_instance_ref_uuid", "to": []string{"test", "virtual_machine_interface_routing_instance_ref_uuid"}})
	//    RoutingInstanceref = append(RoutingInstanceref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_routing_instance_ref_uuid1", "to": []string{"test", "virtual_machine_interface_routing_instance_ref_uuid1"}})
	//
	//    RoutingInstanceAttr := map[string]interface{}{}
	//
	//
	//
	//    common.SetValueByPath(RoutingInstanceAttr, ".Ipv6ServiceChainAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(RoutingInstanceAttr, ".Direction", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(RoutingInstanceAttr, ".MPLSLabel", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(RoutingInstanceAttr, ".VlanTag", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(RoutingInstanceAttr, ".SRCMac", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(RoutingInstanceAttr, ".ServiceChainAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(RoutingInstanceAttr, ".DSTMac", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(RoutingInstanceAttr, ".Protocol", ".", "test")
	//
	//
	//
	//    RoutingInstanceref = append(RoutingInstanceref, map[string]interface{}{"operation":"update", "uuid":"virtual_machine_interface_routing_instance_ref_uuid2", "to": []string{"test", "virtual_machine_interface_routing_instance_ref_uuid2"}, "attr": RoutingInstanceAttr})
	//
	//    common.SetValueByPath(updateMap, "RoutingInstanceRefs", ".", RoutingInstanceref)
	//
	//    var BridgeDomainref []interface{}
	//    BridgeDomainref = append(BridgeDomainref, map[string]interface{}{"operation":"delete", "uuid":"virtual_machine_interface_bridge_domain_ref_uuid", "to": []string{"test", "virtual_machine_interface_bridge_domain_ref_uuid"}})
	//    BridgeDomainref = append(BridgeDomainref, map[string]interface{}{"operation":"add", "uuid":"virtual_machine_interface_bridge_domain_ref_uuid1", "to": []string{"test", "virtual_machine_interface_bridge_domain_ref_uuid1"}})
	//
	//    BridgeDomainAttr := map[string]interface{}{}
	//
	//
	//
	//    common.SetValueByPath(BridgeDomainAttr, ".VlanTag", ".", 1.0)
	//
	//
	//
	//    BridgeDomainref = append(BridgeDomainref, map[string]interface{}{"operation":"update", "uuid":"virtual_machine_interface_bridge_domain_ref_uuid2", "to": []string{"test", "virtual_machine_interface_bridge_domain_ref_uuid2"}, "attr": BridgeDomainAttr})
	//
	//    common.SetValueByPath(updateMap, "BridgeDomainRefs", ".", BridgeDomainref)
	//
	//
	_, err = db.CreateVirtualMachineInterface(ctx,
		&models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateVirtualMachineInterface(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_routing_instance` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing RoutingInstanceRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_routing_instance_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_routing_instance_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_routing_instance_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "RoutingInstanceRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteRoutingInstance(ctx,
		&models.DeleteRoutingInstanceRequest{
			ID: "virtual_machine_interface_routing_instance_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_routing_instance_ref_uuid  failed", err)
	}
	_, err = db.DeleteRoutingInstance(ctx,
		&models.DeleteRoutingInstanceRequest{
			ID: "virtual_machine_interface_routing_instance_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_routing_instance_ref_uuid1  failed", err)
	}
	_, err = db.DeleteRoutingInstance(
		ctx,
		&models.DeleteRoutingInstanceRequest{
			ID: "virtual_machine_interface_routing_instance_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_routing_instance_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_bridge_domain` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing BridgeDomainRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_bridge_domain_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_bridge_domain_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_bridge_domain_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "BridgeDomainRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteBridgeDomain(ctx,
		&models.DeleteBridgeDomainRequest{
			ID: "virtual_machine_interface_bridge_domain_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_bridge_domain_ref_uuid  failed", err)
	}
	_, err = db.DeleteBridgeDomain(ctx,
		&models.DeleteBridgeDomainRequest{
			ID: "virtual_machine_interface_bridge_domain_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_bridge_domain_ref_uuid1  failed", err)
	}
	_, err = db.DeleteBridgeDomain(
		ctx,
		&models.DeleteBridgeDomainRequest{
			ID: "virtual_machine_interface_bridge_domain_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_bridge_domain_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_bgp_router` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing BGPRouterRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_bgp_router_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_bgp_router_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_bgp_router_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "BGPRouterRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteBGPRouter(ctx,
		&models.DeleteBGPRouterRequest{
			ID: "virtual_machine_interface_bgp_router_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_bgp_router_ref_uuid  failed", err)
	}
	_, err = db.DeleteBGPRouter(ctx,
		&models.DeleteBGPRouterRequest{
			ID: "virtual_machine_interface_bgp_router_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_bgp_router_ref_uuid1  failed", err)
	}
	_, err = db.DeleteBGPRouter(
		ctx,
		&models.DeleteBGPRouterRequest{
			ID: "virtual_machine_interface_bgp_router_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_bgp_router_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_security_logging_object` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing SecurityLoggingObjectRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_security_logging_object_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_security_logging_object_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_security_logging_object_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteSecurityLoggingObject(ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: "virtual_machine_interface_security_logging_object_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_security_logging_object_ref_uuid  failed", err)
	}
	_, err = db.DeleteSecurityLoggingObject(ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: "virtual_machine_interface_security_logging_object_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_security_logging_object_ref_uuid1  failed", err)
	}
	_, err = db.DeleteSecurityLoggingObject(
		ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: "virtual_machine_interface_security_logging_object_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_security_logging_object_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_qos_config` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing QosConfigRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_qos_config_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_qos_config_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_qos_config_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "QosConfigRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteQosConfig(ctx,
		&models.DeleteQosConfigRequest{
			ID: "virtual_machine_interface_qos_config_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_qos_config_ref_uuid  failed", err)
	}
	_, err = db.DeleteQosConfig(ctx,
		&models.DeleteQosConfigRequest{
			ID: "virtual_machine_interface_qos_config_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_qos_config_ref_uuid1  failed", err)
	}
	_, err = db.DeleteQosConfig(
		ctx,
		&models.DeleteQosConfigRequest{
			ID: "virtual_machine_interface_qos_config_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_qos_config_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_port_tuple` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing PortTupleRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_port_tuple_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_port_tuple_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_port_tuple_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "PortTupleRefs delete failed")
		}
		return nil
	})
	_, err = db.DeletePortTuple(ctx,
		&models.DeletePortTupleRequest{
			ID: "virtual_machine_interface_port_tuple_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_port_tuple_ref_uuid  failed", err)
	}
	_, err = db.DeletePortTuple(ctx,
		&models.DeletePortTupleRequest{
			ID: "virtual_machine_interface_port_tuple_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_port_tuple_ref_uuid1  failed", err)
	}
	_, err = db.DeletePortTuple(
		ctx,
		&models.DeletePortTupleRequest{
			ID: "virtual_machine_interface_port_tuple_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_port_tuple_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_physical_interface` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing PhysicalInterfaceRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_physical_interface_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_physical_interface_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_physical_interface_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "PhysicalInterfaceRefs delete failed")
		}
		return nil
	})
	_, err = db.DeletePhysicalInterface(ctx,
		&models.DeletePhysicalInterfaceRequest{
			ID: "virtual_machine_interface_physical_interface_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_physical_interface_ref_uuid  failed", err)
	}
	_, err = db.DeletePhysicalInterface(ctx,
		&models.DeletePhysicalInterfaceRequest{
			ID: "virtual_machine_interface_physical_interface_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_physical_interface_ref_uuid1  failed", err)
	}
	_, err = db.DeletePhysicalInterface(
		ctx,
		&models.DeletePhysicalInterfaceRequest{
			ID: "virtual_machine_interface_physical_interface_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_physical_interface_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_virtual_machine` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_virtual_machine_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_virtual_machine_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_virtual_machine_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteVirtualMachine(ctx,
		&models.DeleteVirtualMachineRequest{
			ID: "virtual_machine_interface_virtual_machine_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_virtual_machine_ref_uuid  failed", err)
	}
	_, err = db.DeleteVirtualMachine(ctx,
		&models.DeleteVirtualMachineRequest{
			ID: "virtual_machine_interface_virtual_machine_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_virtual_machine_ref_uuid1  failed", err)
	}
	_, err = db.DeleteVirtualMachine(
		ctx,
		&models.DeleteVirtualMachineRequest{
			ID: "virtual_machine_interface_virtual_machine_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_virtual_machine_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_security_group` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing SecurityGroupRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_security_group_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_security_group_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_security_group_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "SecurityGroupRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteSecurityGroup(ctx,
		&models.DeleteSecurityGroupRequest{
			ID: "virtual_machine_interface_security_group_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_security_group_ref_uuid  failed", err)
	}
	_, err = db.DeleteSecurityGroup(ctx,
		&models.DeleteSecurityGroupRequest{
			ID: "virtual_machine_interface_security_group_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_security_group_ref_uuid1  failed", err)
	}
	_, err = db.DeleteSecurityGroup(
		ctx,
		&models.DeleteSecurityGroupRequest{
			ID: "virtual_machine_interface_security_group_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_security_group_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_service_endpoint` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceEndpointRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_service_endpoint_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_service_endpoint_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_service_endpoint_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceEndpointRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteServiceEndpoint(ctx,
		&models.DeleteServiceEndpointRequest{
			ID: "virtual_machine_interface_service_endpoint_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_service_endpoint_ref_uuid  failed", err)
	}
	_, err = db.DeleteServiceEndpoint(ctx,
		&models.DeleteServiceEndpointRequest{
			ID: "virtual_machine_interface_service_endpoint_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_service_endpoint_ref_uuid1  failed", err)
	}
	_, err = db.DeleteServiceEndpoint(
		ctx,
		&models.DeleteServiceEndpointRequest{
			ID: "virtual_machine_interface_service_endpoint_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_service_endpoint_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_virtual_machine_interface` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_virtual_machine_interface_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_virtual_machine_interface_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_virtual_machine_interface_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteVirtualMachineInterface(ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "virtual_machine_interface_virtual_machine_interface_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_virtual_machine_interface_ref_uuid  failed", err)
	}
	_, err = db.DeleteVirtualMachineInterface(ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "virtual_machine_interface_virtual_machine_interface_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_virtual_machine_interface_ref_uuid1  failed", err)
	}
	_, err = db.DeleteVirtualMachineInterface(
		ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "virtual_machine_interface_virtual_machine_interface_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_virtual_machine_interface_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_interface_route_table` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing InterfaceRouteTableRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_interface_route_table_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_interface_route_table_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_interface_route_table_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "InterfaceRouteTableRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteInterfaceRouteTable(ctx,
		&models.DeleteInterfaceRouteTableRequest{
			ID: "virtual_machine_interface_interface_route_table_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_interface_route_table_ref_uuid  failed", err)
	}
	_, err = db.DeleteInterfaceRouteTable(ctx,
		&models.DeleteInterfaceRouteTableRequest{
			ID: "virtual_machine_interface_interface_route_table_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_interface_route_table_ref_uuid1  failed", err)
	}
	_, err = db.DeleteInterfaceRouteTable(
		ctx,
		&models.DeleteInterfaceRouteTableRequest{
			ID: "virtual_machine_interface_interface_route_table_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_interface_route_table_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_service_health_check` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceHealthCheckRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_service_health_check_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_service_health_check_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_service_health_check_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceHealthCheckRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteServiceHealthCheck(ctx,
		&models.DeleteServiceHealthCheckRequest{
			ID: "virtual_machine_interface_service_health_check_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_service_health_check_ref_uuid  failed", err)
	}
	_, err = db.DeleteServiceHealthCheck(ctx,
		&models.DeleteServiceHealthCheckRequest{
			ID: "virtual_machine_interface_service_health_check_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_service_health_check_ref_uuid1  failed", err)
	}
	_, err = db.DeleteServiceHealthCheck(
		ctx,
		&models.DeleteServiceHealthCheckRequest{
			ID: "virtual_machine_interface_service_health_check_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_service_health_check_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_virtual_machine_interface_virtual_network` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualNetworkRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_virtual_network_ref_uuid")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_virtual_network_ref_uuid1")
		_, err = stmt.Exec("virtual_machine_interface_dummy_uuid", "virtual_machine_interface_virtual_network_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteVirtualNetwork(ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: "virtual_machine_interface_virtual_network_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_virtual_network_ref_uuid  failed", err)
	}
	_, err = db.DeleteVirtualNetwork(ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: "virtual_machine_interface_virtual_network_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_virtual_network_ref_uuid1  failed", err)
	}
	_, err = db.DeleteVirtualNetwork(
		ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: "virtual_machine_interface_virtual_network_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref virtual_machine_interface_virtual_network_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListVirtualMachineInterface(ctx, &models.ListVirtualMachineInterfaceRequest{
		Spec: &models.ListSpec{Limit: 1}})
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

	response, err = db.ListVirtualMachineInterface(ctx, &models.ListVirtualMachineInterfaceRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.VirtualMachineInterfaces) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
