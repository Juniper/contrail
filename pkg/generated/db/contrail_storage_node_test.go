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

func TestContrailStorageNode(t *testing.T) {
	// t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db, "metadata")
	mutexTable := common.UseTable(db, "contrail_storage_node")
	// mutexProject := common.UseTable(db, "contrail_storage_node")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeContrailStorageNode()
	model.UUID = "contrail_storage_node_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "contrail_storage_node_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var Nodecreateref []*models.ContrailStorageNodeNodeRef
	var NoderefModel *models.Node
	NoderefModel = models.MakeNode()
	NoderefModel.UUID = "contrail_storage_node_node_ref_uuid"
	NoderefModel.FQName = []string{"test", "contrail_storage_node_node_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNode(ctx, tx, &models.CreateNodeRequest{
			Node: NoderefModel,
		})
	})
	NoderefModel.UUID = "contrail_storage_node_node_ref_uuid1"
	NoderefModel.FQName = []string{"test", "contrail_storage_node_node_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNode(ctx, tx, &models.CreateNodeRequest{
			Node: NoderefModel,
		})
	})
	NoderefModel.UUID = "contrail_storage_node_node_ref_uuid2"
	NoderefModel.FQName = []string{"test", "contrail_storage_node_node_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNode(ctx, tx, &models.CreateNodeRequest{
			Node: NoderefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	Nodecreateref = append(Nodecreateref, &models.ContrailStorageNodeNodeRef{UUID: "contrail_storage_node_node_ref_uuid", To: []string{"test", "contrail_storage_node_node_ref_uuid"}})
	Nodecreateref = append(Nodecreateref, &models.ContrailStorageNodeNodeRef{UUID: "contrail_storage_node_node_ref_uuid2", To: []string{"test", "contrail_storage_node_node_ref_uuid2"}})
	model.NodeRefs = Nodecreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "contrail_storage_node_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".StorageBackendBondInterfaceMembers", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".StorageAccessBondInterfaceMembers", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ProvisioningState", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ProvisioningStartTime", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ProvisioningProgressStage", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ProvisioningProgress", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ProvisioningLog", ".", "test")
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
	//    common.SetValueByPath(updateMap, ".OsdDrives", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".JournalDrives", ".", "test")
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "contrail_storage_node_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var Noderef []interface{}
	//    Noderef = append(Noderef, map[string]interface{}{"operation":"delete", "uuid":"contrail_storage_node_node_ref_uuid", "to": []string{"test", "contrail_storage_node_node_ref_uuid"}})
	//    Noderef = append(Noderef, map[string]interface{}{"operation":"add", "uuid":"contrail_storage_node_node_ref_uuid1", "to": []string{"test", "contrail_storage_node_node_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "NodeRefs", ".", Noderef)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateContrailStorageNode(ctx, tx,
			&models.CreateContrailStorageNodeRequest{
				ContrailStorageNode: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateContrailStorageNode(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_contrail_storage_node_node` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing NodeRefs delete statement failed")
		}
		_, err = stmt.Exec("contrail_storage_node_dummy_uuid", "contrail_storage_node_node_ref_uuid")
		_, err = stmt.Exec("contrail_storage_node_dummy_uuid", "contrail_storage_node_node_ref_uuid1")
		_, err = stmt.Exec("contrail_storage_node_dummy_uuid", "contrail_storage_node_node_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "NodeRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNode(ctx, tx,
			&models.DeleteNodeRequest{
				ID: "contrail_storage_node_node_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref contrail_storage_node_node_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNode(ctx, tx,
			&models.DeleteNodeRequest{
				ID: "contrail_storage_node_node_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref contrail_storage_node_node_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNode(
			ctx,
			tx,
			&models.DeleteNodeRequest{
				ID: "contrail_storage_node_node_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref contrail_storage_node_node_ref_uuid2 failed", err)
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
		response, err := ListContrailStorageNode(ctx, tx, &models.ListContrailStorageNodeRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.ContrailStorageNodes) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteContrailStorageNode(ctxDemo, tx,
			&models.DeleteContrailStorageNodeRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteContrailStorageNode(ctx, tx,
			&models.DeleteContrailStorageNodeRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateContrailStorageNode(ctx, tx,
			&models.CreateContrailStorageNodeRequest{
				ContrailStorageNode: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListContrailStorageNode(ctx, tx, &models.ListContrailStorageNodeRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.ContrailStorageNodes) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
