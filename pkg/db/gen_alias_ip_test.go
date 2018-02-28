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

// nolint
func TestAliasIP(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB: testDB,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "alias_ip")
	// mutexProject := UseTable(db.DB, "alias_ip")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeAliasIP()
	model.UUID = "alias_ip_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "alias_ip_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var Projectcreateref []*models.AliasIPProjectRef
	var ProjectrefModel *models.Project
	ProjectrefModel = models.MakeProject()
	ProjectrefModel.UUID = "alias_ip_project_ref_uuid"
	ProjectrefModel.FQName = []string{"test", "alias_ip_project_ref_uuid"}
	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: ProjectrefModel,
	})
	ProjectrefModel.UUID = "alias_ip_project_ref_uuid1"
	ProjectrefModel.FQName = []string{"test", "alias_ip_project_ref_uuid1"}
	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: ProjectrefModel,
	})
	ProjectrefModel.UUID = "alias_ip_project_ref_uuid2"
	ProjectrefModel.FQName = []string{"test", "alias_ip_project_ref_uuid2"}
	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: ProjectrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	Projectcreateref = append(Projectcreateref, &models.AliasIPProjectRef{UUID: "alias_ip_project_ref_uuid", To: []string{"test", "alias_ip_project_ref_uuid"}})
	Projectcreateref = append(Projectcreateref, &models.AliasIPProjectRef{UUID: "alias_ip_project_ref_uuid2", To: []string{"test", "alias_ip_project_ref_uuid2"}})
	model.ProjectRefs = Projectcreateref

	var VirtualMachineInterfacecreateref []*models.AliasIPVirtualMachineInterfaceRef
	var VirtualMachineInterfacerefModel *models.VirtualMachineInterface
	VirtualMachineInterfacerefModel = models.MakeVirtualMachineInterface()
	VirtualMachineInterfacerefModel.UUID = "alias_ip_virtual_machine_interface_ref_uuid"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "alias_ip_virtual_machine_interface_ref_uuid"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	VirtualMachineInterfacerefModel.UUID = "alias_ip_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "alias_ip_virtual_machine_interface_ref_uuid1"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	VirtualMachineInterfacerefModel.UUID = "alias_ip_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "alias_ip_virtual_machine_interface_ref_uuid2"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.AliasIPVirtualMachineInterfaceRef{UUID: "alias_ip_virtual_machine_interface_ref_uuid", To: []string{"test", "alias_ip_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.AliasIPVirtualMachineInterfaceRef{UUID: "alias_ip_virtual_machine_interface_ref_uuid2", To: []string{"test", "alias_ip_virtual_machine_interface_ref_uuid2"}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "alias_ip_admin_project_uuid"
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
	//
	//    common.SetValueByPath(updateMap, ".AliasIPAddressFamily", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".AliasIPAddress", ".", "test")
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "alias_ip_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var Projectref []interface{}
	//    Projectref = append(Projectref, map[string]interface{}{"operation":"delete", "uuid":"alias_ip_project_ref_uuid", "to": []string{"test", "alias_ip_project_ref_uuid"}})
	//    Projectref = append(Projectref, map[string]interface{}{"operation":"add", "uuid":"alias_ip_project_ref_uuid1", "to": []string{"test", "alias_ip_project_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ProjectRefs", ".", Projectref)
	//
	//    var VirtualMachineInterfaceref []interface{}
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"delete", "uuid":"alias_ip_virtual_machine_interface_ref_uuid", "to": []string{"test", "alias_ip_virtual_machine_interface_ref_uuid"}})
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"add", "uuid":"alias_ip_virtual_machine_interface_ref_uuid1", "to": []string{"test", "alias_ip_virtual_machine_interface_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)
	//
	//
	_, err = db.CreateAliasIP(ctx,
		&models.CreateAliasIPRequest{
			AliasIP: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateAliasIP(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_alias_ip_project` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ProjectRefs delete statement failed")
		}
		_, err = stmt.Exec("alias_ip_dummy_uuid", "alias_ip_project_ref_uuid")
		_, err = stmt.Exec("alias_ip_dummy_uuid", "alias_ip_project_ref_uuid1")
		_, err = stmt.Exec("alias_ip_dummy_uuid", "alias_ip_project_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ProjectRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteProject(ctx,
		&models.DeleteProjectRequest{
			ID: "alias_ip_project_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref alias_ip_project_ref_uuid  failed", err)
	}
	_, err = db.DeleteProject(ctx,
		&models.DeleteProjectRequest{
			ID: "alias_ip_project_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref alias_ip_project_ref_uuid1  failed", err)
	}
	_, err = db.DeleteProject(
		ctx,
		&models.DeleteProjectRequest{
			ID: "alias_ip_project_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref alias_ip_project_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_alias_ip_virtual_machine_interface` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs delete statement failed")
		}
		_, err = stmt.Exec("alias_ip_dummy_uuid", "alias_ip_virtual_machine_interface_ref_uuid")
		_, err = stmt.Exec("alias_ip_dummy_uuid", "alias_ip_virtual_machine_interface_ref_uuid1")
		_, err = stmt.Exec("alias_ip_dummy_uuid", "alias_ip_virtual_machine_interface_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteVirtualMachineInterface(ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "alias_ip_virtual_machine_interface_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref alias_ip_virtual_machine_interface_ref_uuid  failed", err)
	}
	_, err = db.DeleteVirtualMachineInterface(ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "alias_ip_virtual_machine_interface_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref alias_ip_virtual_machine_interface_ref_uuid1  failed", err)
	}
	_, err = db.DeleteVirtualMachineInterface(
		ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "alias_ip_virtual_machine_interface_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref alias_ip_virtual_machine_interface_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListAliasIP(ctx, &models.ListAliasIPRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.AliasIPs) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteAliasIP(ctxDemo,
		&models.DeleteAliasIPRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateAliasIP(ctx,
		&models.CreateAliasIPRequest{
			AliasIP: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteAliasIP(ctx,
		&models.DeleteAliasIPRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListAliasIP(ctx, &models.ListAliasIPRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.AliasIPs) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
