package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

func TestLogicalRouter(t *testing.T) {
	t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	common.UseTable(db, "metadata")
	common.UseTable(db, "logical_router")
	defer func() {
		common.ClearTable(db, "logical_router")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeLogicalRouter()
	model.UUID = "logical_router_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "logical_router_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var VirtualNetworkcreateref []*models.LogicalRouterVirtualNetworkRef
	var VirtualNetworkrefModel *models.VirtualNetwork
	VirtualNetworkrefModel = models.MakeVirtualNetwork()
	VirtualNetworkrefModel.UUID = "logical_router_virtual_network_ref_uuid"
	VirtualNetworkrefModel.FQName = []string{"test", "logical_router_virtual_network_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(ctx, tx, &models.CreateVirtualNetworkRequest{
			VirtualNetwork: VirtualNetworkrefModel,
		})
	})
	VirtualNetworkrefModel.UUID = "logical_router_virtual_network_ref_uuid1"
	VirtualNetworkrefModel.FQName = []string{"test", "logical_router_virtual_network_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(ctx, tx, &models.CreateVirtualNetworkRequest{
			VirtualNetwork: VirtualNetworkrefModel,
		})
	})
	VirtualNetworkrefModel.UUID = "logical_router_virtual_network_ref_uuid2"
	VirtualNetworkrefModel.FQName = []string{"test", "logical_router_virtual_network_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(ctx, tx, &models.CreateVirtualNetworkRequest{
			VirtualNetwork: VirtualNetworkrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.LogicalRouterVirtualNetworkRef{UUID: "logical_router_virtual_network_ref_uuid", To: []string{"test", "logical_router_virtual_network_ref_uuid"}})
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.LogicalRouterVirtualNetworkRef{UUID: "logical_router_virtual_network_ref_uuid2", To: []string{"test", "logical_router_virtual_network_ref_uuid2"}})
	model.VirtualNetworkRefs = VirtualNetworkcreateref

	var PhysicalRoutercreateref []*models.LogicalRouterPhysicalRouterRef
	var PhysicalRouterrefModel *models.PhysicalRouter
	PhysicalRouterrefModel = models.MakePhysicalRouter()
	PhysicalRouterrefModel.UUID = "logical_router_physical_router_ref_uuid"
	PhysicalRouterrefModel.FQName = []string{"test", "logical_router_physical_router_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
			PhysicalRouter: PhysicalRouterrefModel,
		})
	})
	PhysicalRouterrefModel.UUID = "logical_router_physical_router_ref_uuid1"
	PhysicalRouterrefModel.FQName = []string{"test", "logical_router_physical_router_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
			PhysicalRouter: PhysicalRouterrefModel,
		})
	})
	PhysicalRouterrefModel.UUID = "logical_router_physical_router_ref_uuid2"
	PhysicalRouterrefModel.FQName = []string{"test", "logical_router_physical_router_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
			PhysicalRouter: PhysicalRouterrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	PhysicalRoutercreateref = append(PhysicalRoutercreateref, &models.LogicalRouterPhysicalRouterRef{UUID: "logical_router_physical_router_ref_uuid", To: []string{"test", "logical_router_physical_router_ref_uuid"}})
	PhysicalRoutercreateref = append(PhysicalRoutercreateref, &models.LogicalRouterPhysicalRouterRef{UUID: "logical_router_physical_router_ref_uuid2", To: []string{"test", "logical_router_physical_router_ref_uuid2"}})
	model.PhysicalRouterRefs = PhysicalRoutercreateref

	var BGPVPNcreateref []*models.LogicalRouterBGPVPNRef
	var BGPVPNrefModel *models.BGPVPN
	BGPVPNrefModel = models.MakeBGPVPN()
	BGPVPNrefModel.UUID = "logical_router_bgpvpn_ref_uuid"
	BGPVPNrefModel.FQName = []string{"test", "logical_router_bgpvpn_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateBGPVPN(ctx, tx, &models.CreateBGPVPNRequest{
			BGPVPN: BGPVPNrefModel,
		})
	})
	BGPVPNrefModel.UUID = "logical_router_bgpvpn_ref_uuid1"
	BGPVPNrefModel.FQName = []string{"test", "logical_router_bgpvpn_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateBGPVPN(ctx, tx, &models.CreateBGPVPNRequest{
			BGPVPN: BGPVPNrefModel,
		})
	})
	BGPVPNrefModel.UUID = "logical_router_bgpvpn_ref_uuid2"
	BGPVPNrefModel.FQName = []string{"test", "logical_router_bgpvpn_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateBGPVPN(ctx, tx, &models.CreateBGPVPNRequest{
			BGPVPN: BGPVPNrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BGPVPNcreateref = append(BGPVPNcreateref, &models.LogicalRouterBGPVPNRef{UUID: "logical_router_bgpvpn_ref_uuid", To: []string{"test", "logical_router_bgpvpn_ref_uuid"}})
	BGPVPNcreateref = append(BGPVPNcreateref, &models.LogicalRouterBGPVPNRef{UUID: "logical_router_bgpvpn_ref_uuid2", To: []string{"test", "logical_router_bgpvpn_ref_uuid2"}})
	model.BGPVPNRefs = BGPVPNcreateref

	var RouteTargetcreateref []*models.LogicalRouterRouteTargetRef
	var RouteTargetrefModel *models.RouteTarget
	RouteTargetrefModel = models.MakeRouteTarget()
	RouteTargetrefModel.UUID = "logical_router_route_target_ref_uuid"
	RouteTargetrefModel.FQName = []string{"test", "logical_router_route_target_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateRouteTarget(ctx, tx, &models.CreateRouteTargetRequest{
			RouteTarget: RouteTargetrefModel,
		})
	})
	RouteTargetrefModel.UUID = "logical_router_route_target_ref_uuid1"
	RouteTargetrefModel.FQName = []string{"test", "logical_router_route_target_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateRouteTarget(ctx, tx, &models.CreateRouteTargetRequest{
			RouteTarget: RouteTargetrefModel,
		})
	})
	RouteTargetrefModel.UUID = "logical_router_route_target_ref_uuid2"
	RouteTargetrefModel.FQName = []string{"test", "logical_router_route_target_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateRouteTarget(ctx, tx, &models.CreateRouteTargetRequest{
			RouteTarget: RouteTargetrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	RouteTargetcreateref = append(RouteTargetcreateref, &models.LogicalRouterRouteTargetRef{UUID: "logical_router_route_target_ref_uuid", To: []string{"test", "logical_router_route_target_ref_uuid"}})
	RouteTargetcreateref = append(RouteTargetcreateref, &models.LogicalRouterRouteTargetRef{UUID: "logical_router_route_target_ref_uuid2", To: []string{"test", "logical_router_route_target_ref_uuid2"}})
	model.RouteTargetRefs = RouteTargetcreateref

	var VirtualMachineInterfacecreateref []*models.LogicalRouterVirtualMachineInterfaceRef
	var VirtualMachineInterfacerefModel *models.VirtualMachineInterface
	VirtualMachineInterfacerefModel = models.MakeVirtualMachineInterface()
	VirtualMachineInterfacerefModel.UUID = "logical_router_virtual_machine_interface_ref_uuid"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "logical_router_virtual_machine_interface_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	VirtualMachineInterfacerefModel.UUID = "logical_router_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "logical_router_virtual_machine_interface_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	VirtualMachineInterfacerefModel.UUID = "logical_router_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "logical_router_virtual_machine_interface_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.LogicalRouterVirtualMachineInterfaceRef{UUID: "logical_router_virtual_machine_interface_ref_uuid", To: []string{"test", "logical_router_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.LogicalRouterVirtualMachineInterfaceRef{UUID: "logical_router_virtual_machine_interface_ref_uuid2", To: []string{"test", "logical_router_virtual_machine_interface_ref_uuid2"}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref

	var ServiceInstancecreateref []*models.LogicalRouterServiceInstanceRef
	var ServiceInstancerefModel *models.ServiceInstance
	ServiceInstancerefModel = models.MakeServiceInstance()
	ServiceInstancerefModel.UUID = "logical_router_service_instance_ref_uuid"
	ServiceInstancerefModel.FQName = []string{"test", "logical_router_service_instance_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(ctx, tx, &models.CreateServiceInstanceRequest{
			ServiceInstance: ServiceInstancerefModel,
		})
	})
	ServiceInstancerefModel.UUID = "logical_router_service_instance_ref_uuid1"
	ServiceInstancerefModel.FQName = []string{"test", "logical_router_service_instance_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(ctx, tx, &models.CreateServiceInstanceRequest{
			ServiceInstance: ServiceInstancerefModel,
		})
	})
	ServiceInstancerefModel.UUID = "logical_router_service_instance_ref_uuid2"
	ServiceInstancerefModel.FQName = []string{"test", "logical_router_service_instance_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(ctx, tx, &models.CreateServiceInstanceRequest{
			ServiceInstance: ServiceInstancerefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceInstancecreateref = append(ServiceInstancecreateref, &models.LogicalRouterServiceInstanceRef{UUID: "logical_router_service_instance_ref_uuid", To: []string{"test", "logical_router_service_instance_ref_uuid"}})
	ServiceInstancecreateref = append(ServiceInstancecreateref, &models.LogicalRouterServiceInstanceRef{UUID: "logical_router_service_instance_ref_uuid2", To: []string{"test", "logical_router_service_instance_ref_uuid2"}})
	model.ServiceInstanceRefs = ServiceInstancecreateref

	var RouteTablecreateref []*models.LogicalRouterRouteTableRef
	var RouteTablerefModel *models.RouteTable
	RouteTablerefModel = models.MakeRouteTable()
	RouteTablerefModel.UUID = "logical_router_route_table_ref_uuid"
	RouteTablerefModel.FQName = []string{"test", "logical_router_route_table_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateRouteTable(ctx, tx, &models.CreateRouteTableRequest{
			RouteTable: RouteTablerefModel,
		})
	})
	RouteTablerefModel.UUID = "logical_router_route_table_ref_uuid1"
	RouteTablerefModel.FQName = []string{"test", "logical_router_route_table_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateRouteTable(ctx, tx, &models.CreateRouteTableRequest{
			RouteTable: RouteTablerefModel,
		})
	})
	RouteTablerefModel.UUID = "logical_router_route_table_ref_uuid2"
	RouteTablerefModel.FQName = []string{"test", "logical_router_route_table_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateRouteTable(ctx, tx, &models.CreateRouteTableRequest{
			RouteTable: RouteTablerefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	RouteTablecreateref = append(RouteTablecreateref, &models.LogicalRouterRouteTableRef{UUID: "logical_router_route_table_ref_uuid", To: []string{"test", "logical_router_route_table_ref_uuid"}})
	RouteTablecreateref = append(RouteTablecreateref, &models.LogicalRouterRouteTableRef{UUID: "logical_router_route_table_ref_uuid2", To: []string{"test", "logical_router_route_table_ref_uuid2"}})
	model.RouteTableRefs = RouteTablecreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "logical_router_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(ctx, tx, &models.CreateProjectRequest{
			Project: projectModel,
		})
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//    //populate update map
	//    updateMap := map[string]interface{}{}
	//
	//
	//    common.SetValueByPath(updateMap, ".VxlanNetworkIdentifier", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
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
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    if ".ConfiguredRouteTargetList.RouteTarget" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".ConfiguredRouteTargetList.RouteTarget", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".ConfiguredRouteTargetList.RouteTarget", ".", `{"test": "test"}`)
	//    }
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "logical_router_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var VirtualMachineInterfaceref []interface{}
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"delete", "uuid":"logical_router_virtual_machine_interface_ref_uuid", "to": []string{"test", "logical_router_virtual_machine_interface_ref_uuid"}})
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"add", "uuid":"logical_router_virtual_machine_interface_ref_uuid1", "to": []string{"test", "logical_router_virtual_machine_interface_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)
	//
	//    var ServiceInstanceref []interface{}
	//    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"delete", "uuid":"logical_router_service_instance_ref_uuid", "to": []string{"test", "logical_router_service_instance_ref_uuid"}})
	//    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"add", "uuid":"logical_router_service_instance_ref_uuid1", "to": []string{"test", "logical_router_service_instance_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ServiceInstanceRefs", ".", ServiceInstanceref)
	//
	//    var RouteTableref []interface{}
	//    RouteTableref = append(RouteTableref, map[string]interface{}{"operation":"delete", "uuid":"logical_router_route_table_ref_uuid", "to": []string{"test", "logical_router_route_table_ref_uuid"}})
	//    RouteTableref = append(RouteTableref, map[string]interface{}{"operation":"add", "uuid":"logical_router_route_table_ref_uuid1", "to": []string{"test", "logical_router_route_table_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "RouteTableRefs", ".", RouteTableref)
	//
	//    var VirtualNetworkref []interface{}
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"delete", "uuid":"logical_router_virtual_network_ref_uuid", "to": []string{"test", "logical_router_virtual_network_ref_uuid"}})
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"add", "uuid":"logical_router_virtual_network_ref_uuid1", "to": []string{"test", "logical_router_virtual_network_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualNetworkRefs", ".", VirtualNetworkref)
	//
	//    var PhysicalRouterref []interface{}
	//    PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation":"delete", "uuid":"logical_router_physical_router_ref_uuid", "to": []string{"test", "logical_router_physical_router_ref_uuid"}})
	//    PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation":"add", "uuid":"logical_router_physical_router_ref_uuid1", "to": []string{"test", "logical_router_physical_router_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "PhysicalRouterRefs", ".", PhysicalRouterref)
	//
	//    var BGPVPNref []interface{}
	//    BGPVPNref = append(BGPVPNref, map[string]interface{}{"operation":"delete", "uuid":"logical_router_bgpvpn_ref_uuid", "to": []string{"test", "logical_router_bgpvpn_ref_uuid"}})
	//    BGPVPNref = append(BGPVPNref, map[string]interface{}{"operation":"add", "uuid":"logical_router_bgpvpn_ref_uuid1", "to": []string{"test", "logical_router_bgpvpn_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "BGPVPNRefs", ".", BGPVPNref)
	//
	//    var RouteTargetref []interface{}
	//    RouteTargetref = append(RouteTargetref, map[string]interface{}{"operation":"delete", "uuid":"logical_router_route_target_ref_uuid", "to": []string{"test", "logical_router_route_target_ref_uuid"}})
	//    RouteTargetref = append(RouteTargetref, map[string]interface{}{"operation":"add", "uuid":"logical_router_route_target_ref_uuid1", "to": []string{"test", "logical_router_route_target_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "RouteTargetRefs", ".", RouteTargetref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLogicalRouter(ctx, tx,
			&models.CreateLogicalRouterRequest{
				LogicalRouter: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateLogicalRouter(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_logical_router_service_instance` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceInstanceRefs delete statement failed")
		}
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_service_instance_ref_uuid")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_service_instance_ref_uuid1")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_service_instance_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(ctx, tx,
			&models.DeleteServiceInstanceRequest{
				ID: "logical_router_service_instance_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_service_instance_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(ctx, tx,
			&models.DeleteServiceInstanceRequest{
				ID: "logical_router_service_instance_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_service_instance_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(
			ctx,
			tx,
			&models.DeleteServiceInstanceRequest{
				ID: "logical_router_service_instance_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_service_instance_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_logical_router_route_table` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing RouteTableRefs delete statement failed")
		}
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_route_table_ref_uuid")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_route_table_ref_uuid1")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_route_table_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "RouteTableRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteRouteTable(ctx, tx,
			&models.DeleteRouteTableRequest{
				ID: "logical_router_route_table_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_route_table_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteRouteTable(ctx, tx,
			&models.DeleteRouteTableRequest{
				ID: "logical_router_route_table_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_route_table_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteRouteTable(
			ctx,
			tx,
			&models.DeleteRouteTableRequest{
				ID: "logical_router_route_table_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_route_table_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_logical_router_virtual_network` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualNetworkRefs delete statement failed")
		}
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_virtual_network_ref_uuid")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_virtual_network_ref_uuid1")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_virtual_network_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(ctx, tx,
			&models.DeleteVirtualNetworkRequest{
				ID: "logical_router_virtual_network_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_virtual_network_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(ctx, tx,
			&models.DeleteVirtualNetworkRequest{
				ID: "logical_router_virtual_network_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_virtual_network_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(
			ctx,
			tx,
			&models.DeleteVirtualNetworkRequest{
				ID: "logical_router_virtual_network_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_virtual_network_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_logical_router_physical_router` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing PhysicalRouterRefs delete statement failed")
		}
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_physical_router_ref_uuid")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_physical_router_ref_uuid1")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_physical_router_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(ctx, tx,
			&models.DeletePhysicalRouterRequest{
				ID: "logical_router_physical_router_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_physical_router_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(ctx, tx,
			&models.DeletePhysicalRouterRequest{
				ID: "logical_router_physical_router_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_physical_router_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(
			ctx,
			tx,
			&models.DeletePhysicalRouterRequest{
				ID: "logical_router_physical_router_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_physical_router_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_logical_router_bgpvpn` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing BGPVPNRefs delete statement failed")
		}
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_bgpvpn_ref_uuid")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_bgpvpn_ref_uuid1")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_bgpvpn_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "BGPVPNRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteBGPVPN(ctx, tx,
			&models.DeleteBGPVPNRequest{
				ID: "logical_router_bgpvpn_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_bgpvpn_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteBGPVPN(ctx, tx,
			&models.DeleteBGPVPNRequest{
				ID: "logical_router_bgpvpn_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_bgpvpn_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteBGPVPN(
			ctx,
			tx,
			&models.DeleteBGPVPNRequest{
				ID: "logical_router_bgpvpn_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_bgpvpn_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_logical_router_route_target` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing RouteTargetRefs delete statement failed")
		}
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_route_target_ref_uuid")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_route_target_ref_uuid1")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_route_target_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "RouteTargetRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteRouteTarget(ctx, tx,
			&models.DeleteRouteTargetRequest{
				ID: "logical_router_route_target_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_route_target_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteRouteTarget(ctx, tx,
			&models.DeleteRouteTargetRequest{
				ID: "logical_router_route_target_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_route_target_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteRouteTarget(
			ctx,
			tx,
			&models.DeleteRouteTargetRequest{
				ID: "logical_router_route_target_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_route_target_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_logical_router_virtual_machine_interface` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs delete statement failed")
		}
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_virtual_machine_interface_ref_uuid")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_virtual_machine_interface_ref_uuid1")
		_, err = stmt.Exec("logical_router_dummy_uuid", "logical_router_virtual_machine_interface_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(ctx, tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "logical_router_virtual_machine_interface_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_virtual_machine_interface_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(ctx, tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "logical_router_virtual_machine_interface_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_virtual_machine_interface_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(
			ctx,
			tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "logical_router_virtual_machine_interface_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref logical_router_virtual_machine_interface_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(ctx, tx, &models.DeleteProjectRequest{
			ID: projectModel.UUID})
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListLogicalRouter(ctx, tx, &models.ListLogicalRouterRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.LogicalRouters) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLogicalRouter(ctxDemo, tx,
			&models.DeleteLogicalRouterRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLogicalRouter(ctx, tx,
			&models.DeleteLogicalRouterRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLogicalRouter(ctx, tx,
			&models.CreateLogicalRouterRequest{
				LogicalRouter: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListLogicalRouter(ctx, tx, &models.ListLogicalRouterRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.LogicalRouters) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
