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

func TestContrailCluster(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "contrail_cluster")
	// mutexProject := UseTable(db.DB, "contrail_cluster")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeContrailCluster()
	model.UUID = "contrail_cluster_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "contrail_cluster_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "contrail_cluster_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".SudoPassword", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".StatisticsTTL", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".RabbitmqPort", ".", "test")
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
	//    common.SetValueByPath(updateMap, ".ProvisionerType", ".", "test")
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
	//    common.SetValueByPath(updateMap, ".Orchestrator", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".OpenstackInternalVipInterface", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".OpenstackInternalVip", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".OpenstackExternalVipInterface", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".OpenstackExternalVip", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".OpenstackEnableHaproxy", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".NTPServer", ".", "test")
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
	//    common.SetValueByPath(updateMap, ".FlowTTL", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DefaultVrouterBondInterfaceMembers", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DefaultVrouterBondInterface", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DefaultGateway", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ContrailVersion", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ContainerRegistry", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ConfigurationVersion", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ConfigAuditTTL", ".", "test")
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "contrail_cluster_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//
	_, err = db.CreateContrailCluster(ctx,
		&models.CreateContrailClusterRequest{
			ContrailCluster: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateContrailCluster(tx, model.UUID, updateMap)
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

	response, err := db.ListContrailCluster(ctx, &models.ListContrailClusterRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.ContrailClusters) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteContrailCluster(ctxDemo,
		&models.DeleteContrailClusterRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateContrailCluster(ctx,
		&models.CreateContrailClusterRequest{
			ContrailCluster: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteContrailCluster(ctx,
		&models.DeleteContrailClusterRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListContrailCluster(ctx, &models.ListContrailClusterRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.ContrailClusters) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
