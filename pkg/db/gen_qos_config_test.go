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
func TestQosConfig(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB: testDB,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "qos_config")
	// mutexProject := UseTable(db.DB, "qos_config")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeQosConfig()
	model.UUID = "qos_config_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "qos_config_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var GlobalSystemConfigcreateref []*models.QosConfigGlobalSystemConfigRef
	var GlobalSystemConfigrefModel *models.GlobalSystemConfig
	GlobalSystemConfigrefModel = models.MakeGlobalSystemConfig()
	GlobalSystemConfigrefModel.UUID = "qos_config_global_system_config_ref_uuid"
	GlobalSystemConfigrefModel.FQName = []string{"test", "qos_config_global_system_config_ref_uuid"}
	_, err = db.CreateGlobalSystemConfig(ctx, &models.CreateGlobalSystemConfigRequest{
		GlobalSystemConfig: GlobalSystemConfigrefModel,
	})
	GlobalSystemConfigrefModel.UUID = "qos_config_global_system_config_ref_uuid1"
	GlobalSystemConfigrefModel.FQName = []string{"test", "qos_config_global_system_config_ref_uuid1"}
	_, err = db.CreateGlobalSystemConfig(ctx, &models.CreateGlobalSystemConfigRequest{
		GlobalSystemConfig: GlobalSystemConfigrefModel,
	})
	GlobalSystemConfigrefModel.UUID = "qos_config_global_system_config_ref_uuid2"
	GlobalSystemConfigrefModel.FQName = []string{"test", "qos_config_global_system_config_ref_uuid2"}
	_, err = db.CreateGlobalSystemConfig(ctx, &models.CreateGlobalSystemConfigRequest{
		GlobalSystemConfig: GlobalSystemConfigrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	GlobalSystemConfigcreateref = append(GlobalSystemConfigcreateref, &models.QosConfigGlobalSystemConfigRef{UUID: "qos_config_global_system_config_ref_uuid", To: []string{"test", "qos_config_global_system_config_ref_uuid"}})
	GlobalSystemConfigcreateref = append(GlobalSystemConfigcreateref, &models.QosConfigGlobalSystemConfigRef{UUID: "qos_config_global_system_config_ref_uuid2", To: []string{"test", "qos_config_global_system_config_ref_uuid2"}})
	model.GlobalSystemConfigRefs = GlobalSystemConfigcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "qos_config_admin_project_uuid"
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
	//    if ".VlanPriorityEntries.QosIDForwardingClassPair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".VlanPriorityEntries.QosIDForwardingClassPair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".VlanPriorityEntries.QosIDForwardingClassPair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".QosConfigType", ".", "test")
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
	//    if ".MPLSExpEntries.QosIDForwardingClassPair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".MPLSExpEntries.QosIDForwardingClassPair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".MPLSExpEntries.QosIDForwardingClassPair", ".", `{"test": "test"}`)
	//    }
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
	//    if ".DSCPEntries.QosIDForwardingClassPair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".DSCPEntries.QosIDForwardingClassPair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".DSCPEntries.QosIDForwardingClassPair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DefaultForwardingClassID", ".", 1.0)
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "qos_config_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var GlobalSystemConfigref []interface{}
	//    GlobalSystemConfigref = append(GlobalSystemConfigref, map[string]interface{}{"operation":"delete", "uuid":"qos_config_global_system_config_ref_uuid", "to": []string{"test", "qos_config_global_system_config_ref_uuid"}})
	//    GlobalSystemConfigref = append(GlobalSystemConfigref, map[string]interface{}{"operation":"add", "uuid":"qos_config_global_system_config_ref_uuid1", "to": []string{"test", "qos_config_global_system_config_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "GlobalSystemConfigRefs", ".", GlobalSystemConfigref)
	//
	//
	_, err = db.CreateQosConfig(ctx,
		&models.CreateQosConfigRequest{
			QosConfig: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateQosConfig(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_qos_config_global_system_config` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing GlobalSystemConfigRefs delete statement failed")
		}
		_, err = stmt.Exec("qos_config_dummy_uuid", "qos_config_global_system_config_ref_uuid")
		_, err = stmt.Exec("qos_config_dummy_uuid", "qos_config_global_system_config_ref_uuid1")
		_, err = stmt.Exec("qos_config_dummy_uuid", "qos_config_global_system_config_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "GlobalSystemConfigRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteGlobalSystemConfig(ctx,
		&models.DeleteGlobalSystemConfigRequest{
			ID: "qos_config_global_system_config_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref qos_config_global_system_config_ref_uuid  failed", err)
	}
	_, err = db.DeleteGlobalSystemConfig(ctx,
		&models.DeleteGlobalSystemConfigRequest{
			ID: "qos_config_global_system_config_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref qos_config_global_system_config_ref_uuid1  failed", err)
	}
	_, err = db.DeleteGlobalSystemConfig(
		ctx,
		&models.DeleteGlobalSystemConfigRequest{
			ID: "qos_config_global_system_config_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref qos_config_global_system_config_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListQosConfig(ctx, &models.ListQosConfigRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.QosConfigs) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteQosConfig(ctxDemo,
		&models.DeleteQosConfigRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateQosConfig(ctx,
		&models.CreateQosConfigRequest{
			QosConfig: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteQosConfig(ctx,
		&models.DeleteQosConfigRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListQosConfig(ctx, &models.ListQosConfigRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.QosConfigs) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
