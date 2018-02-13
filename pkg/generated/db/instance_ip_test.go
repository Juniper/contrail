package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

func TestInstanceIP(t *testing.T) {
	t.Parallel()
	db := testDB
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

	var NetworkIpamcreateref []*models.InstanceIPNetworkIpamRef
	var NetworkIpamrefModel *models.NetworkIpam
	NetworkIpamrefModel = models.MakeNetworkIpam()
	NetworkIpamrefModel.UUID = "instance_ip_network_ipam_ref_uuid"
	NetworkIpamrefModel.FQName = []string{"test", "instance_ip_network_ipam_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(tx, NetworkIpamrefModel)
	})
	NetworkIpamrefModel.UUID = "instance_ip_network_ipam_ref_uuid1"
	NetworkIpamrefModel.FQName = []string{"test", "instance_ip_network_ipam_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(tx, NetworkIpamrefModel)
	})
	NetworkIpamrefModel.UUID = "instance_ip_network_ipam_ref_uuid2"
	NetworkIpamrefModel.FQName = []string{"test", "instance_ip_network_ipam_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(tx, NetworkIpamrefModel)
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
		return CreateVirtualNetwork(tx, VirtualNetworkrefModel)
	})
	VirtualNetworkrefModel.UUID = "instance_ip_virtual_network_ref_uuid1"
	VirtualNetworkrefModel.FQName = []string{"test", "instance_ip_virtual_network_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(tx, VirtualNetworkrefModel)
	})
	VirtualNetworkrefModel.UUID = "instance_ip_virtual_network_ref_uuid2"
	VirtualNetworkrefModel.FQName = []string{"test", "instance_ip_virtual_network_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(tx, VirtualNetworkrefModel)
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
		return CreateVirtualMachineInterface(tx, VirtualMachineInterfacerefModel)
	})
	VirtualMachineInterfacerefModel.UUID = "instance_ip_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "instance_ip_virtual_machine_interface_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(tx, VirtualMachineInterfacerefModel)
	})
	VirtualMachineInterfacerefModel.UUID = "instance_ip_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "instance_ip_virtual_machine_interface_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(tx, VirtualMachineInterfacerefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.InstanceIPVirtualMachineInterfaceRef{UUID: "instance_ip_virtual_machine_interface_ref_uuid", To: []string{"test", "instance_ip_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.InstanceIPVirtualMachineInterfaceRef{UUID: "instance_ip_virtual_machine_interface_ref_uuid2", To: []string{"test", "instance_ip_virtual_machine_interface_ref_uuid2"}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref

	var PhysicalRoutercreateref []*models.InstanceIPPhysicalRouterRef
	var PhysicalRouterrefModel *models.PhysicalRouter
	PhysicalRouterrefModel = models.MakePhysicalRouter()
	PhysicalRouterrefModel.UUID = "instance_ip_physical_router_ref_uuid"
	PhysicalRouterrefModel.FQName = []string{"test", "instance_ip_physical_router_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(tx, PhysicalRouterrefModel)
	})
	PhysicalRouterrefModel.UUID = "instance_ip_physical_router_ref_uuid1"
	PhysicalRouterrefModel.FQName = []string{"test", "instance_ip_physical_router_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(tx, PhysicalRouterrefModel)
	})
	PhysicalRouterrefModel.UUID = "instance_ip_physical_router_ref_uuid2"
	PhysicalRouterrefModel.FQName = []string{"test", "instance_ip_physical_router_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(tx, PhysicalRouterrefModel)
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
		return CreateVirtualRouter(tx, VirtualRouterrefModel)
	})
	VirtualRouterrefModel.UUID = "instance_ip_virtual_router_ref_uuid1"
	VirtualRouterrefModel.FQName = []string{"test", "instance_ip_virtual_router_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualRouter(tx, VirtualRouterrefModel)
	})
	VirtualRouterrefModel.UUID = "instance_ip_virtual_router_ref_uuid2"
	VirtualRouterrefModel.FQName = []string{"test", "instance_ip_virtual_router_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualRouter(tx, VirtualRouterrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualRoutercreateref = append(VirtualRoutercreateref, &models.InstanceIPVirtualRouterRef{UUID: "instance_ip_virtual_router_ref_uuid", To: []string{"test", "instance_ip_virtual_router_ref_uuid"}})
	VirtualRoutercreateref = append(VirtualRoutercreateref, &models.InstanceIPVirtualRouterRef{UUID: "instance_ip_virtual_router_ref_uuid2", To: []string{"test", "instance_ip_virtual_router_ref_uuid2"}})
	model.VirtualRouterRefs = VirtualRoutercreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "instance_ip_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(tx, projectModel)
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//populate update map
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".SubnetUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ServiceInstanceIP", ".", true)

	common.SetValueByPath(updateMap, ".ServiceHealthCheckIP", ".", true)

	common.SetValueByPath(updateMap, ".SecondaryIPTrackingIP.IPPrefixLen", ".", 1.0)

	common.SetValueByPath(updateMap, ".SecondaryIPTrackingIP.IPPrefix", ".", "test")

	if ".Perms2.Share" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ParentType", ".", "test")

	common.SetValueByPath(updateMap, ".InstanceIPSecondary", ".", true)

	common.SetValueByPath(updateMap, ".InstanceIPMode", ".", "test")

	common.SetValueByPath(updateMap, ".InstanceIPLocalIP", ".", true)

	common.SetValueByPath(updateMap, ".InstanceIPFamily", ".", "test")

	common.SetValueByPath(updateMap, ".InstanceIPAddress", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)

	common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")

	if ".FQName" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".FQName", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	if ".Annotations.KeyValuePair" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, "uuid", ".", "instance_ip_dummy_uuid")
	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	// Create Attr values for testing ref update(ADD,UPDATE,DELETE)

	var VirtualNetworkref []interface{}
	VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation": "delete", "uuid": "instance_ip_virtual_network_ref_uuid", "to": []string{"test", "instance_ip_virtual_network_ref_uuid"}})
	VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation": "add", "uuid": "instance_ip_virtual_network_ref_uuid1", "to": []string{"test", "instance_ip_virtual_network_ref_uuid1"}})

	common.SetValueByPath(updateMap, "VirtualNetworkRefs", ".", VirtualNetworkref)

	var VirtualMachineInterfaceref []interface{}
	VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation": "delete", "uuid": "instance_ip_virtual_machine_interface_ref_uuid", "to": []string{"test", "instance_ip_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation": "add", "uuid": "instance_ip_virtual_machine_interface_ref_uuid1", "to": []string{"test", "instance_ip_virtual_machine_interface_ref_uuid1"}})

	common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)

	var PhysicalRouterref []interface{}
	PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation": "delete", "uuid": "instance_ip_physical_router_ref_uuid", "to": []string{"test", "instance_ip_physical_router_ref_uuid"}})
	PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation": "add", "uuid": "instance_ip_physical_router_ref_uuid1", "to": []string{"test", "instance_ip_physical_router_ref_uuid1"}})

	common.SetValueByPath(updateMap, "PhysicalRouterRefs", ".", PhysicalRouterref)

	var VirtualRouterref []interface{}
	VirtualRouterref = append(VirtualRouterref, map[string]interface{}{"operation": "delete", "uuid": "instance_ip_virtual_router_ref_uuid", "to": []string{"test", "instance_ip_virtual_router_ref_uuid"}})
	VirtualRouterref = append(VirtualRouterref, map[string]interface{}{"operation": "add", "uuid": "instance_ip_virtual_router_ref_uuid1", "to": []string{"test", "instance_ip_virtual_router_ref_uuid1"}})

	common.SetValueByPath(updateMap, "VirtualRouterRefs", ".", VirtualRouterref)

	var NetworkIpamref []interface{}
	NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation": "delete", "uuid": "instance_ip_network_ipam_ref_uuid", "to": []string{"test", "instance_ip_network_ipam_ref_uuid"}})
	NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation": "add", "uuid": "instance_ip_network_ipam_ref_uuid1", "to": []string{"test", "instance_ip_network_ipam_ref_uuid1"}})

	common.SetValueByPath(updateMap, "NetworkIpamRefs", ".", NetworkIpamref)

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateInstanceIP(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateInstanceIP(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	//Delete ref entries, referred objects

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
		return DeleteVirtualMachineInterface(tx, "instance_ip_virtual_machine_interface_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_machine_interface_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(tx, "instance_ip_virtual_machine_interface_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_machine_interface_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(tx, "instance_ip_virtual_machine_interface_ref_uuid2", nil)
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
		return DeletePhysicalRouter(tx, "instance_ip_physical_router_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_physical_router_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(tx, "instance_ip_physical_router_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_physical_router_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(tx, "instance_ip_physical_router_ref_uuid2", nil)
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
		return DeleteVirtualRouter(tx, "instance_ip_virtual_router_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_router_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualRouter(tx, "instance_ip_virtual_router_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_router_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualRouter(tx, "instance_ip_virtual_router_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_router_ref_uuid2 failed", err)
	}

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
		return DeleteNetworkIpam(tx, "instance_ip_network_ipam_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_network_ipam_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(tx, "instance_ip_network_ipam_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_network_ipam_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(tx, "instance_ip_network_ipam_ref_uuid2", nil)
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
		return DeleteVirtualNetwork(tx, "instance_ip_virtual_network_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_network_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(tx, "instance_ip_virtual_network_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_network_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(tx, "instance_ip_virtual_network_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref instance_ip_virtual_network_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, projectModel.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListInstanceIP(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteInstanceIP(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteInstanceIP(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateInstanceIP(tx, model)
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListInstanceIP(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
