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

func TestVirtualIP(t *testing.T) {
	t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	common.UseTable(db, "metadata")
	common.UseTable(db, "virtual_ip")
	defer func() {
		common.ClearTable(db, "virtual_ip")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeVirtualIP()
	model.UUID = "virtual_ip_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "virtual_ip_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var LoadbalancerPoolcreateref []*models.VirtualIPLoadbalancerPoolRef
	var LoadbalancerPoolrefModel *models.LoadbalancerPool
	LoadbalancerPoolrefModel = models.MakeLoadbalancerPool()
	LoadbalancerPoolrefModel.UUID = "virtual_ip_loadbalancer_pool_ref_uuid"
	LoadbalancerPoolrefModel.FQName = []string{"test", "virtual_ip_loadbalancer_pool_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerPool(ctx, tx, &models.CreateLoadbalancerPoolRequest{
			LoadbalancerPool: LoadbalancerPoolrefModel,
		})
	})
	LoadbalancerPoolrefModel.UUID = "virtual_ip_loadbalancer_pool_ref_uuid1"
	LoadbalancerPoolrefModel.FQName = []string{"test", "virtual_ip_loadbalancer_pool_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerPool(ctx, tx, &models.CreateLoadbalancerPoolRequest{
			LoadbalancerPool: LoadbalancerPoolrefModel,
		})
	})
	LoadbalancerPoolrefModel.UUID = "virtual_ip_loadbalancer_pool_ref_uuid2"
	LoadbalancerPoolrefModel.FQName = []string{"test", "virtual_ip_loadbalancer_pool_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerPool(ctx, tx, &models.CreateLoadbalancerPoolRequest{
			LoadbalancerPool: LoadbalancerPoolrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	LoadbalancerPoolcreateref = append(LoadbalancerPoolcreateref, &models.VirtualIPLoadbalancerPoolRef{UUID: "virtual_ip_loadbalancer_pool_ref_uuid", To: []string{"test", "virtual_ip_loadbalancer_pool_ref_uuid"}})
	LoadbalancerPoolcreateref = append(LoadbalancerPoolcreateref, &models.VirtualIPLoadbalancerPoolRef{UUID: "virtual_ip_loadbalancer_pool_ref_uuid2", To: []string{"test", "virtual_ip_loadbalancer_pool_ref_uuid2"}})
	model.LoadbalancerPoolRefs = LoadbalancerPoolcreateref

	var VirtualMachineInterfacecreateref []*models.VirtualIPVirtualMachineInterfaceRef
	var VirtualMachineInterfacerefModel *models.VirtualMachineInterface
	VirtualMachineInterfacerefModel = models.MakeVirtualMachineInterface()
	VirtualMachineInterfacerefModel.UUID = "virtual_ip_virtual_machine_interface_ref_uuid"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "virtual_ip_virtual_machine_interface_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	VirtualMachineInterfacerefModel.UUID = "virtual_ip_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "virtual_ip_virtual_machine_interface_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	VirtualMachineInterfacerefModel.UUID = "virtual_ip_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "virtual_ip_virtual_machine_interface_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.VirtualIPVirtualMachineInterfaceRef{UUID: "virtual_ip_virtual_machine_interface_ref_uuid", To: []string{"test", "virtual_ip_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.VirtualIPVirtualMachineInterfaceRef{UUID: "virtual_ip_virtual_machine_interface_ref_uuid2", To: []string{"test", "virtual_ip_virtual_machine_interface_ref_uuid2"}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "virtual_ip_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.SubnetID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.StatusDescription", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.Status", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.ProtocolPort", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.Protocol", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.PersistenceType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.PersistenceCookieName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.ConnectionLimit", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.AdminState", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".VirtualIPProperties.Address", ".", "test")
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "virtual_ip_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var VirtualMachineInterfaceref []interface{}
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"delete", "uuid":"virtual_ip_virtual_machine_interface_ref_uuid", "to": []string{"test", "virtual_ip_virtual_machine_interface_ref_uuid"}})
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"add", "uuid":"virtual_ip_virtual_machine_interface_ref_uuid1", "to": []string{"test", "virtual_ip_virtual_machine_interface_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)
	//
	//    var LoadbalancerPoolref []interface{}
	//    LoadbalancerPoolref = append(LoadbalancerPoolref, map[string]interface{}{"operation":"delete", "uuid":"virtual_ip_loadbalancer_pool_ref_uuid", "to": []string{"test", "virtual_ip_loadbalancer_pool_ref_uuid"}})
	//    LoadbalancerPoolref = append(LoadbalancerPoolref, map[string]interface{}{"operation":"add", "uuid":"virtual_ip_loadbalancer_pool_ref_uuid1", "to": []string{"test", "virtual_ip_loadbalancer_pool_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "LoadbalancerPoolRefs", ".", LoadbalancerPoolref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualIP(ctx, tx,
			&models.CreateVirtualIPRequest{
				VirtualIP: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateVirtualIP(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_ip_virtual_machine_interface` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_ip_dummy_uuid", "virtual_ip_virtual_machine_interface_ref_uuid")
		_, err = stmt.Exec("virtual_ip_dummy_uuid", "virtual_ip_virtual_machine_interface_ref_uuid1")
		_, err = stmt.Exec("virtual_ip_dummy_uuid", "virtual_ip_virtual_machine_interface_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(ctx, tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "virtual_ip_virtual_machine_interface_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref virtual_ip_virtual_machine_interface_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(ctx, tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "virtual_ip_virtual_machine_interface_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref virtual_ip_virtual_machine_interface_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(
			ctx,
			tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "virtual_ip_virtual_machine_interface_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref virtual_ip_virtual_machine_interface_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_ip_loadbalancer_pool` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing LoadbalancerPoolRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_ip_dummy_uuid", "virtual_ip_loadbalancer_pool_ref_uuid")
		_, err = stmt.Exec("virtual_ip_dummy_uuid", "virtual_ip_loadbalancer_pool_ref_uuid1")
		_, err = stmt.Exec("virtual_ip_dummy_uuid", "virtual_ip_loadbalancer_pool_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "LoadbalancerPoolRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerPool(ctx, tx,
			&models.DeleteLoadbalancerPoolRequest{
				ID: "virtual_ip_loadbalancer_pool_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref virtual_ip_loadbalancer_pool_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerPool(ctx, tx,
			&models.DeleteLoadbalancerPoolRequest{
				ID: "virtual_ip_loadbalancer_pool_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref virtual_ip_loadbalancer_pool_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerPool(
			ctx,
			tx,
			&models.DeleteLoadbalancerPoolRequest{
				ID: "virtual_ip_loadbalancer_pool_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref virtual_ip_loadbalancer_pool_ref_uuid2 failed", err)
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
		response, err := ListVirtualIP(ctx, tx, &models.ListVirtualIPRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.VirtualIPs) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualIP(ctxDemo, tx,
			&models.DeleteVirtualIPRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualIP(ctx, tx,
			&models.DeleteVirtualIPRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualIP(ctx, tx,
			&models.CreateVirtualIPRequest{
				VirtualIP: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListVirtualIP(ctx, tx, &models.ListVirtualIPRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.VirtualIPs) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
