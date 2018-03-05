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

func TestServer(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "server")
	// mutexProject := UseTable(db.DB, "server")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeServer()
	model.UUID = "server_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "server_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "server_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".UserID", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Updated", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".TenantID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Status", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Progress", ".", 1.0)
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
	//    common.SetValueByPath(updateMap, ".Name", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Locked", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Image.Links.Type", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Image.Links.Rel", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Image.Links.Href", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Image.ID", ".", "test")
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
	//    common.SetValueByPath(updateMap, ".ID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".HostStatus", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".HostId", ".", "test")
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
	//    common.SetValueByPath(updateMap, ".Flavor.Links.Type", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Flavor.Links.Rel", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Flavor.Links.Href", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Flavor.ID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Created", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ConfigurationVersion", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ConfigDrive", ".", true)
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
	//    common.SetValueByPath(updateMap, ".Addresses.Addr", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".AccessIPv6", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".AccessIPv4", ".", "test")
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "server_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//
	_, err = db.CreateServer(ctx,
		&models.CreateServerRequest{
			Server: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateServer(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListServer(ctx, &models.ListServerRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.Servers) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteServer(ctxDemo,
		&models.DeleteServerRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateServer(ctx,
		&models.CreateServerRequest{
			Server: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteServer(ctx,
		&models.DeleteServerRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListServer(ctx, &models.ListServerRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.Servers) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
