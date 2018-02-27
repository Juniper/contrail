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

//For skip import error.
var _ = errors.New("")

func TestVirtualRouter(t *testing.T) {
	// t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db, "metadata")
	mutexTable := common.UseTable(db, "virtual_router")
	// mutexProject := common.UseTable(db, "virtual_router")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeVirtualRouter()
	model.UUID = "virtual_router_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "virtual_router_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var NetworkIpamcreateref []*models.VirtualRouterNetworkIpamRef
	var NetworkIpamrefModel *models.NetworkIpam
	NetworkIpamrefModel = models.MakeNetworkIpam()
	NetworkIpamrefModel.UUID = "virtual_router_network_ipam_ref_uuid"
	NetworkIpamrefModel.FQName = []string{"test", "virtual_router_network_ipam_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(ctx, tx, &models.CreateNetworkIpamRequest{
			NetworkIpam: NetworkIpamrefModel,
		})
	})
	NetworkIpamrefModel.UUID = "virtual_router_network_ipam_ref_uuid1"
	NetworkIpamrefModel.FQName = []string{"test", "virtual_router_network_ipam_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(ctx, tx, &models.CreateNetworkIpamRequest{
			NetworkIpam: NetworkIpamrefModel,
		})
	})
	NetworkIpamrefModel.UUID = "virtual_router_network_ipam_ref_uuid2"
	NetworkIpamrefModel.FQName = []string{"test", "virtual_router_network_ipam_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(ctx, tx, &models.CreateNetworkIpamRequest{
			NetworkIpam: NetworkIpamrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NetworkIpamcreateref = append(NetworkIpamcreateref, &models.VirtualRouterNetworkIpamRef{UUID: "virtual_router_network_ipam_ref_uuid", To: []string{"test", "virtual_router_network_ipam_ref_uuid"}})
	NetworkIpamcreateref = append(NetworkIpamcreateref, &models.VirtualRouterNetworkIpamRef{UUID: "virtual_router_network_ipam_ref_uuid2", To: []string{"test", "virtual_router_network_ipam_ref_uuid2"}})
	model.NetworkIpamRefs = NetworkIpamcreateref

	var VirtualMachinecreateref []*models.VirtualRouterVirtualMachineRef
	var VirtualMachinerefModel *models.VirtualMachine
	VirtualMachinerefModel = models.MakeVirtualMachine()
	VirtualMachinerefModel.UUID = "virtual_router_virtual_machine_ref_uuid"
	VirtualMachinerefModel.FQName = []string{"test", "virtual_router_virtual_machine_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachine(ctx, tx, &models.CreateVirtualMachineRequest{
			VirtualMachine: VirtualMachinerefModel,
		})
	})
	VirtualMachinerefModel.UUID = "virtual_router_virtual_machine_ref_uuid1"
	VirtualMachinerefModel.FQName = []string{"test", "virtual_router_virtual_machine_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachine(ctx, tx, &models.CreateVirtualMachineRequest{
			VirtualMachine: VirtualMachinerefModel,
		})
	})
	VirtualMachinerefModel.UUID = "virtual_router_virtual_machine_ref_uuid2"
	VirtualMachinerefModel.FQName = []string{"test", "virtual_router_virtual_machine_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachine(ctx, tx, &models.CreateVirtualMachineRequest{
			VirtualMachine: VirtualMachinerefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachinecreateref = append(VirtualMachinecreateref, &models.VirtualRouterVirtualMachineRef{UUID: "virtual_router_virtual_machine_ref_uuid", To: []string{"test", "virtual_router_virtual_machine_ref_uuid"}})
	VirtualMachinecreateref = append(VirtualMachinecreateref, &models.VirtualRouterVirtualMachineRef{UUID: "virtual_router_virtual_machine_ref_uuid2", To: []string{"test", "virtual_router_virtual_machine_ref_uuid2"}})
	model.VirtualMachineRefs = VirtualMachinecreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "virtual_router_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".VirtualRouterType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualRouterIPAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualRouterDPDKEnabled", ".", true)
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
	//    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "virtual_router_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var VirtualMachineref []interface{}
	//    VirtualMachineref = append(VirtualMachineref, map[string]interface{}{"operation":"delete", "uuid":"virtual_router_virtual_machine_ref_uuid", "to": []string{"test", "virtual_router_virtual_machine_ref_uuid"}})
	//    VirtualMachineref = append(VirtualMachineref, map[string]interface{}{"operation":"add", "uuid":"virtual_router_virtual_machine_ref_uuid1", "to": []string{"test", "virtual_router_virtual_machine_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualMachineRefs", ".", VirtualMachineref)
	//
	//    var NetworkIpamref []interface{}
	//    NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation":"delete", "uuid":"virtual_router_network_ipam_ref_uuid", "to": []string{"test", "virtual_router_network_ipam_ref_uuid"}})
	//    NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation":"add", "uuid":"virtual_router_network_ipam_ref_uuid1", "to": []string{"test", "virtual_router_network_ipam_ref_uuid1"}})
	//
	//    NetworkIpamAttr := map[string]interface{}{}
	//
	//
	//
	//    common.SetValueByPath(NetworkIpamAttr, ".Subnet", ".", map[string]string{"test": "test"})
	//
	//
	//
	//    common.SetValueByPath(NetworkIpamAttr, ".AllocationPools", ".", map[string]string{"test": "test"})
	//
	//
	//
	//    NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation":"update", "uuid":"virtual_router_network_ipam_ref_uuid2", "to": []string{"test", "virtual_router_network_ipam_ref_uuid2"}, "attr": NetworkIpamAttr})
	//
	//    common.SetValueByPath(updateMap, "NetworkIpamRefs", ".", NetworkIpamref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualRouter(ctx, tx,
			&models.CreateVirtualRouterRequest{
				VirtualRouter: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateVirtualRouter(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_router_network_ipam` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing NetworkIpamRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_router_dummy_uuid", "virtual_router_network_ipam_ref_uuid")
		_, err = stmt.Exec("virtual_router_dummy_uuid", "virtual_router_network_ipam_ref_uuid1")
		_, err = stmt.Exec("virtual_router_dummy_uuid", "virtual_router_network_ipam_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "NetworkIpamRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(ctx, tx,
			&models.DeleteNetworkIpamRequest{
				ID: "virtual_router_network_ipam_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref virtual_router_network_ipam_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(ctx, tx,
			&models.DeleteNetworkIpamRequest{
				ID: "virtual_router_network_ipam_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref virtual_router_network_ipam_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(
			ctx,
			tx,
			&models.DeleteNetworkIpamRequest{
				ID: "virtual_router_network_ipam_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref virtual_router_network_ipam_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_router_virtual_machine` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_router_dummy_uuid", "virtual_router_virtual_machine_ref_uuid")
		_, err = stmt.Exec("virtual_router_dummy_uuid", "virtual_router_virtual_machine_ref_uuid1")
		_, err = stmt.Exec("virtual_router_dummy_uuid", "virtual_router_virtual_machine_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachine(ctx, tx,
			&models.DeleteVirtualMachineRequest{
				ID: "virtual_router_virtual_machine_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref virtual_router_virtual_machine_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachine(ctx, tx,
			&models.DeleteVirtualMachineRequest{
				ID: "virtual_router_virtual_machine_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref virtual_router_virtual_machine_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachine(
			ctx,
			tx,
			&models.DeleteVirtualMachineRequest{
				ID: "virtual_router_virtual_machine_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref virtual_router_virtual_machine_ref_uuid2 failed", err)
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
		response, err := ListVirtualRouter(ctx, tx, &models.ListVirtualRouterRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.VirtualRouters) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualRouter(ctxDemo, tx,
			&models.DeleteVirtualRouterRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualRouter(ctx, tx,
			&models.DeleteVirtualRouterRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualRouter(ctx, tx,
			&models.CreateVirtualRouterRequest{
				VirtualRouter: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListVirtualRouter(ctx, tx, &models.ListVirtualRouterRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.VirtualRouters) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
