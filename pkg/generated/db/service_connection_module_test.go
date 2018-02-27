package db

import (
	"context"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestServiceConnectionModule(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB: testDB,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "service_connection_module")
	// mutexProject := common.UseTable(db.DB, "service_connection_module")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeServiceConnectionModule()
	model.UUID = "service_connection_module_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "service_connection_module_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceObjectcreateref []*models.ServiceConnectionModuleServiceObjectRef
	var ServiceObjectrefModel *models.ServiceObject
	ServiceObjectrefModel = models.MakeServiceObject()
	ServiceObjectrefModel.UUID = "service_connection_module_service_object_ref_uuid"
	ServiceObjectrefModel.FQName = []string{"test", "service_connection_module_service_object_ref_uuid"}
	_, err = db.CreateServiceObject(ctx, &models.CreateServiceObjectRequest{
		ServiceObject: ServiceObjectrefModel,
	})
	ServiceObjectrefModel.UUID = "service_connection_module_service_object_ref_uuid1"
	ServiceObjectrefModel.FQName = []string{"test", "service_connection_module_service_object_ref_uuid1"}
	_, err = db.CreateServiceObject(ctx, &models.CreateServiceObjectRequest{
		ServiceObject: ServiceObjectrefModel,
	})
	ServiceObjectrefModel.UUID = "service_connection_module_service_object_ref_uuid2"
	ServiceObjectrefModel.FQName = []string{"test", "service_connection_module_service_object_ref_uuid2"}
	_, err = db.CreateServiceObject(ctx, &models.CreateServiceObjectRequest{
		ServiceObject: ServiceObjectrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceObjectcreateref = append(ServiceObjectcreateref, &models.ServiceConnectionModuleServiceObjectRef{UUID: "service_connection_module_service_object_ref_uuid", To: []string{"test", "service_connection_module_service_object_ref_uuid"}})
	ServiceObjectcreateref = append(ServiceObjectcreateref, &models.ServiceConnectionModuleServiceObjectRef{UUID: "service_connection_module_service_object_ref_uuid2", To: []string{"test", "service_connection_module_service_object_ref_uuid2"}})
	model.ServiceObjectRefs = ServiceObjectcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "service_connection_module_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".ServiceType", ".", "test")
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
	//    common.SetValueByPath(updateMap, ".E2Service", ".", "test")
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "service_connection_module_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var ServiceObjectref []interface{}
	//    ServiceObjectref = append(ServiceObjectref, map[string]interface{}{"operation":"delete", "uuid":"service_connection_module_service_object_ref_uuid", "to": []string{"test", "service_connection_module_service_object_ref_uuid"}})
	//    ServiceObjectref = append(ServiceObjectref, map[string]interface{}{"operation":"add", "uuid":"service_connection_module_service_object_ref_uuid1", "to": []string{"test", "service_connection_module_service_object_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ServiceObjectRefs", ".", ServiceObjectref)
	//
	//
	_, err = db.CreateServiceConnectionModule(ctx,
		&models.CreateServiceConnectionModuleRequest{
			ServiceConnectionModule: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateServiceConnectionModule(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := common.GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_service_connection_module_service_object` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceObjectRefs delete statement failed")
		}
		_, err = stmt.Exec("service_connection_module_dummy_uuid", "service_connection_module_service_object_ref_uuid")
		_, err = stmt.Exec("service_connection_module_dummy_uuid", "service_connection_module_service_object_ref_uuid1")
		_, err = stmt.Exec("service_connection_module_dummy_uuid", "service_connection_module_service_object_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceObjectRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteServiceObject(ctx,
		&models.DeleteServiceObjectRequest{
			ID: "service_connection_module_service_object_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref service_connection_module_service_object_ref_uuid  failed", err)
	}
	_, err = db.DeleteServiceObject(ctx,
		&models.DeleteServiceObjectRequest{
			ID: "service_connection_module_service_object_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref service_connection_module_service_object_ref_uuid1  failed", err)
	}
	_, err = db.DeleteServiceObject(
		ctx,
		&models.DeleteServiceObjectRequest{
			ID: "service_connection_module_service_object_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref service_connection_module_service_object_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListServiceConnectionModule(ctx, &models.ListServiceConnectionModuleRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.ServiceConnectionModules) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteServiceConnectionModule(ctxDemo,
		&models.DeleteServiceConnectionModuleRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateServiceConnectionModule(ctx,
		&models.CreateServiceConnectionModuleRequest{
			ServiceConnectionModule: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteServiceConnectionModule(ctx,
		&models.DeleteServiceConnectionModuleRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListServiceConnectionModule(ctx, &models.ListServiceConnectionModuleRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.ServiceConnectionModules) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
