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

func TestInstanceIP(t *testing.T) {
	t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	common.UseTable(db, "metadata")
	common.UseTable(db, "instance_ip")
	defer func() {
		common.ClearTable(db, "instance_ip")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeInstanceIP()
	model.UUID = "instance_ip_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "instance_ip_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var PhysicalRoutercreateref []*models.InstanceIPPhysicalRouterRef
	var PhysicalRouterrefModel *models.PhysicalRouter
	PhysicalRouterrefModel = models.MakePhysicalRouter()
	PhysicalRouterrefModel.UUID = "instance_ip_physical_router_ref_uuid"
	PhysicalRouterrefModel.FQName = []string{"test", "instance_ip_physical_router_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
			PhysicalRouter: PhysicalRouterrefModel,
		})
	})
	PhysicalRouterrefModel.UUID = "instance_ip_physical_router_ref_uuid1"
	PhysicalRouterrefModel.FQName = []string{"test", "instance_ip_physical_router_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
			PhysicalRouter: PhysicalRouterrefModel,
		})
	})
	PhysicalRouterrefModel.UUID = "instance_ip_physical_router_ref_uuid2"
	PhysicalRouterrefModel.FQName = []string{"test", "instance_ip_physical_router_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
			PhysicalRouter: PhysicalRouterrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	PhysicalRoutercreateref = append(PhysicalRoutercreateref, &models.InstanceIPPhysicalRouterRef{UUID: "instance_ip_physical_router_ref_uuid", To: []string{"test", "instance_ip_physical_router_ref_uuid"}})
	PhysicalRoutercreateref = append(PhysicalRoutercreateref, &models.InstanceIPPhysicalRouterRef{UUID: "instance_ip_physical_router_ref_uuid2", To: []string{"test", "instance_ip_physical_router_ref_uuid2"}})
	model.PhysicalRouterRefs = PhysicalRoutercreateref

	var VirtualRoutercreateref []*models.InstanceIPVirtualRouterRef
	var VirtualRouterrefModel *models.VirtualRouter
	VirtualRouterrefModel = models.MakeVirtualRouter()
	VirtualRouterrefModel.UUID = "instance_ip_virtual_router_ref_uuid"
	VirtualRouterrefModel.FQName = []string{"test", "instance_ip_virtual_router_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualRouter(ctx, tx, &models.CreateVirtualRouterRequest{
			VirtualRouter: VirtualRouterrefModel,
		})
	})
	VirtualRouterrefModel.UUID = "instance_ip_virtual_router_ref_uuid1"
	VirtualRouterrefModel.FQName = []string{"test", "instance_ip_virtual_router_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualRouter(ctx, tx, &models.CreateVirtualRouterRequest{
			VirtualRouter: VirtualRouterrefModel,
		})
	})
	VirtualRouterrefModel.UUID = "instance_ip_virtual_router_ref_uuid2"
	VirtualRouterrefModel.FQName = []string{"test", "instance_ip_virtual_router_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualRouter(ctx, tx, &models.CreateVirtualRouterRequest{
			VirtualRouter: VirtualRouterrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualRoutercreateref = append(VirtualRoutercreateref, &models.InstanceIPVirtualRouterRef{UUID: "instance_ip_virtual_router_ref_uuid", To: []string{"test", "instance_ip_virtual_router_ref_uuid"}})
	VirtualRoutercreateref = append(VirtualRoutercreateref, &models.InstanceIPVirtualRouterRef{UUID: "instance_ip_virtual_router_ref_uuid2", To: []string{"test", "instance_ip_virtual_router_ref_uuid2"}})
	model.VirtualRouterRefs = VirtualRoutercreateref

	var NetworkIpamcreateref []*models.InstanceIPNetworkIpamRef
	var NetworkIpamrefModel *models.NetworkIpam
	NetworkIpamrefModel = models.MakeNetworkIpam()
	NetworkIpamrefModel.UUID = "instance_ip_network_ipam_ref_uuid"
	NetworkIpamrefModel.FQName = []string{"test", "instance_ip_network_ipam_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(ctx, tx, &models.CreateNetworkIpamRequest{
			NetworkIpam: NetworkIpamrefModel,
		})
	})
	NetworkIpamrefModel.UUID = "instance_ip_network_ipam_ref_uuid1"
	NetworkIpamrefModel.FQName = []string{"test", "instance_ip_network_ipam_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(ctx, tx, &models.CreateNetworkIpamRequest{
			NetworkIpam: NetworkIpamrefModel,
		})
	})
	NetworkIpamrefModel.UUID = "instance_ip_network_ipam_ref_uuid2"
	NetworkIpamrefModel.FQName = []string{"test", "instance_ip_network_ipam_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(ctx, tx, &models.CreateNetworkIpamRequest{
			NetworkIpam: NetworkIpamrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NetworkIpamcreateref = append(NetworkIpamcreateref, &models.InstanceIPNetworkIpamRef{UUID: "instance_ip_network_ipam_ref_uuid", To: []string{"test", "instance_ip_network_ipam_ref_uuid"}})
	NetworkIpamcreateref = append(NetworkIpamcreateref, &models.InstanceIPNetworkIpamRef{UUID: "instance_ip_network_ipam_ref_uuid2", To: []string{"test", "instance_ip_network_ipam_ref_uuid2"}})
	model.NetworkIpamRefs = NetworkIpamcreateref

	var VirtualNetworkcreateref []*models.InstanceIPVirtualNetworkRef
	var VirtualNetworkrefModel *models.VirtualNetwork
	VirtualNetworkrefModel = models.MakeVirtualNetwork()
	VirtualNetworkrefModel.UUID = "instance_ip_virtual_network_ref_uuid"
	VirtualNetworkrefModel.FQName = []string{"test", "instance_ip_virtual_network_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(ctx, tx, &models.CreateVirtualNetworkRequest{
			VirtualNetwork: VirtualNetworkrefModel,
		})
	})
	VirtualNetworkrefModel.UUID = "instance_ip_virtual_network_ref_uuid1"
	VirtualNetworkrefModel.FQName = []string{"test", "instance_ip_virtual_network_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(ctx, tx, &models.CreateVirtualNetworkRequest{
			VirtualNetwork: VirtualNetworkrefModel,
		})
	})
	VirtualNetworkrefModel.UUID = "instance_ip_virtual_network_ref_uuid2"
	VirtualNetworkrefModel.FQName = []string{"test", "instance_ip_virtual_network_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(ctx, tx, &models.CreateVirtualNetworkRequest{
			VirtualNetwork: VirtualNetworkrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.InstanceIPVirtualNetworkRef{UUID: "instance_ip_virtual_network_ref_uuid", To: []string{"test", "instance_ip_virtual_network_ref_uuid"}})
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.InstanceIPVirtualNetworkRef{UUID: "instance_ip_virtual_network_ref_uuid2", To: []string{"test", "instance_ip_virtual_network_ref_uuid2"}})
	model.VirtualNetworkRefs = VirtualNetworkcreateref

	var VirtualMachineInterfacecreateref []*models.InstanceIPVirtualMachineInterfaceRef
	var VirtualMachineInterfacerefModel *models.VirtualMachineInterface
	VirtualMachineInterfacerefModel = models.MakeVirtualMachineInterface()
	VirtualMachineInterfacerefModel.UUID = "instance_ip_virtual_machine_interface_ref_uuid"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "instance_ip_virtual_machine_interface_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	VirtualMachineInterfacerefModel.UUID = "instance_ip_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "instance_ip_virtual_machine_interface_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	VirtualMachineInterfacerefModel.UUID = "instance_ip_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "instance_ip_virtual_machine_interface_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.InstanceIPVirtualMachineInterfaceRef{UUID: "instance_ip_virtual_machine_interface_ref_uuid", To: []string{"test", "instance_ip_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.InstanceIPVirtualMachineInterfaceRef{UUID: "instance_ip_virtual_machine_interface_ref_uuid2", To: []string{"test", "instance_ip_virtual_machine_interface_ref_uuid2"}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "instance_ip_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".SubnetUUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceInstanceIP", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckIP", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".SecondaryIPTrackingIP.IPPrefixLen", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".SecondaryIPTrackingIP.IPPrefix", ".", "test")
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
	//    common.SetValueByPath(updateMap, ".InstanceIPSecondary", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceIPMode", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceIPLocalIP", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceIPFamily", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".InstanceIPAddress", ".", "test")
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
	//    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "instance_ip_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var VirtualMachineInterfaceref []interface{}
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"delete", "uuid":"instance_ip_virtual_machine_interface_ref_uuid", "to": []string{"test", "instance_ip_virtual_machine_interface_ref_uuid"}})
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"add", "uuid":"instance_ip_virtual_machine_interface_ref_uuid1", "to": []string{"test", "instance_ip_virtual_machine_interface_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)
	//
	//    var PhysicalRouterref []interface{}
	//    PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation":"delete", "uuid":"instance_ip_physical_router_ref_uuid", "to": []string{"test", "instance_ip_physical_router_ref_uuid"}})
	//    PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation":"add", "uuid":"instance_ip_physical_router_ref_uuid1", "to": []string{"test", "instance_ip_physical_router_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "PhysicalRouterRefs", ".", PhysicalRouterref)
	//
	//    var VirtualRouterref []interface{}
	//    VirtualRouterref = append(VirtualRouterref, map[string]interface{}{"operation":"delete", "uuid":"instance_ip_virtual_router_ref_uuid", "to": []string{"test", "instance_ip_virtual_router_ref_uuid"}})
	//    VirtualRouterref = append(VirtualRouterref, map[string]interface{}{"operation":"add", "uuid":"instance_ip_virtual_router_ref_uuid1", "to": []string{"test", "instance_ip_virtual_router_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualRouterRefs", ".", VirtualRouterref)
	//
	//    var NetworkIpamref []interface{}
	//    NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation":"delete", "uuid":"instance_ip_network_ipam_ref_uuid", "to": []string{"test", "instance_ip_network_ipam_ref_uuid"}})
	//    NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation":"add", "uuid":"instance_ip_network_ipam_ref_uuid1", "to": []string{"test", "instance_ip_network_ipam_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "NetworkIpamRefs", ".", NetworkIpamref)
	//
	//    var VirtualNetworkref []interface{}
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"delete", "uuid":"instance_ip_virtual_network_ref_uuid", "to": []string{"test", "instance_ip_virtual_network_ref_uuid"}})
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"add", "uuid":"instance_ip_virtual_network_ref_uuid1", "to": []string{"test", "instance_ip_virtual_network_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualNetworkRefs", ".", VirtualNetworkref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateInstanceIP(ctx, tx,
			&models.CreateInstanceIPRequest{
				InstanceIP: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateInstanceIP(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_instance_ip_network_ipam` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing NetworkIpamRefs delete statement failed")
		}
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_network_ipam_ref_uuid")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_network_ipam_ref_uuid1")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_network_ipam_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "NetworkIpamRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(ctx, tx,
			&models.DeleteNetworkIpamRequest{
				ID: "instance_ip_network_ipam_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_network_ipam_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(ctx, tx,
			&models.DeleteNetworkIpamRequest{
				ID: "instance_ip_network_ipam_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_network_ipam_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(
			ctx,
			tx,
			&models.DeleteNetworkIpamRequest{
				ID: "instance_ip_network_ipam_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_network_ipam_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_instance_ip_virtual_network` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualNetworkRefs delete statement failed")
		}
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_virtual_network_ref_uuid")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_virtual_network_ref_uuid1")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_virtual_network_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(ctx, tx,
			&models.DeleteVirtualNetworkRequest{
				ID: "instance_ip_virtual_network_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_network_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(ctx, tx,
			&models.DeleteVirtualNetworkRequest{
				ID: "instance_ip_virtual_network_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_network_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(
			ctx,
			tx,
			&models.DeleteVirtualNetworkRequest{
				ID: "instance_ip_virtual_network_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_network_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_instance_ip_virtual_machine_interface` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs delete statement failed")
		}
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_virtual_machine_interface_ref_uuid")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_virtual_machine_interface_ref_uuid1")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_virtual_machine_interface_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(ctx, tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "instance_ip_virtual_machine_interface_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_machine_interface_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(ctx, tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "instance_ip_virtual_machine_interface_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_machine_interface_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(
			ctx,
			tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "instance_ip_virtual_machine_interface_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_machine_interface_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_instance_ip_physical_router` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing PhysicalRouterRefs delete statement failed")
		}
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_physical_router_ref_uuid")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_physical_router_ref_uuid1")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_physical_router_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(ctx, tx,
			&models.DeletePhysicalRouterRequest{
				ID: "instance_ip_physical_router_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_physical_router_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(ctx, tx,
			&models.DeletePhysicalRouterRequest{
				ID: "instance_ip_physical_router_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_physical_router_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(
			ctx,
			tx,
			&models.DeletePhysicalRouterRequest{
				ID: "instance_ip_physical_router_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_physical_router_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_instance_ip_virtual_router` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualRouterRefs delete statement failed")
		}
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_virtual_router_ref_uuid")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_virtual_router_ref_uuid1")
		_, err = stmt.Exec("instance_ip_dummy_uuid", "instance_ip_virtual_router_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualRouterRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualRouter(ctx, tx,
			&models.DeleteVirtualRouterRequest{
				ID: "instance_ip_virtual_router_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_router_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualRouter(ctx, tx,
			&models.DeleteVirtualRouterRequest{
				ID: "instance_ip_virtual_router_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_router_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualRouter(
			ctx,
			tx,
			&models.DeleteVirtualRouterRequest{
				ID: "instance_ip_virtual_router_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_router_ref_uuid2 failed", err)
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
		response, err := ListInstanceIP(ctx, tx, &models.ListInstanceIPRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.InstanceIPs) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteInstanceIP(ctxDemo, tx,
			&models.DeleteInstanceIPRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteInstanceIP(ctx, tx,
			&models.DeleteInstanceIPRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateInstanceIP(ctx, tx,
			&models.CreateInstanceIPRequest{
				InstanceIP: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListInstanceIP(ctx, tx, &models.ListInstanceIPRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.InstanceIPs) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
