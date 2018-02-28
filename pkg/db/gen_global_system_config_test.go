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

func TestGlobalSystemConfig(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "global_system_config")
	// mutexProject := UseTable(db.DB, "global_system_config")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeGlobalSystemConfig()
	model.UUID = "global_system_config_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "global_system_config_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var BGPRoutercreateref []*models.GlobalSystemConfigBGPRouterRef
	var BGPRouterrefModel *models.BGPRouter
	BGPRouterrefModel = models.MakeBGPRouter()
	BGPRouterrefModel.UUID = "global_system_config_bgp_router_ref_uuid"
	BGPRouterrefModel.FQName = []string{"test", "global_system_config_bgp_router_ref_uuid"}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterrefModel,
	})
	BGPRouterrefModel.UUID = "global_system_config_bgp_router_ref_uuid1"
	BGPRouterrefModel.FQName = []string{"test", "global_system_config_bgp_router_ref_uuid1"}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterrefModel,
	})
	BGPRouterrefModel.UUID = "global_system_config_bgp_router_ref_uuid2"
	BGPRouterrefModel.FQName = []string{"test", "global_system_config_bgp_router_ref_uuid2"}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BGPRoutercreateref = append(BGPRoutercreateref, &models.GlobalSystemConfigBGPRouterRef{UUID: "global_system_config_bgp_router_ref_uuid", To: []string{"test", "global_system_config_bgp_router_ref_uuid"}})
	BGPRoutercreateref = append(BGPRoutercreateref, &models.GlobalSystemConfigBGPRouterRef{UUID: "global_system_config_bgp_router_ref_uuid2", To: []string{"test", "global_system_config_bgp_router_ref_uuid2"}})
	model.BGPRouterRefs = BGPRoutercreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "global_system_config_admin_project_uuid"
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
	//    if ".UserDefinedLogStatistics.Statlist" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".UserDefinedLogStatistics.Statlist", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".UserDefinedLogStatistics.Statlist", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    if ".PluginTuning.PluginProperty" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".PluginTuning.PluginProperty", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".PluginTuning.PluginProperty", ".", `{"test": "test"}`)
	//    }
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
	//    common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveTimeWindow", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveLimitAction", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveLimit", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".MacLimitControl.MacLimitAction", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".MacLimitControl.MacLimit", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".MacAgingTime", ".", 1.0)
	//
	//
	//
	//    if ".IPFabricSubnets.Subnet" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".IPFabricSubnets.Subnet", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".IPFabricSubnets.Subnet", ".", `{"test": "test"}`)
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
	//    common.SetValueByPath(updateMap, ".IbgpAutoMesh", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".GracefulRestartParameters.XMPPHelperEnable", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".GracefulRestartParameters.RestartTime", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".GracefulRestartParameters.LongLivedRestartTime", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".GracefulRestartParameters.EndOfRibTimeout", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".GracefulRestartParameters.Enable", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".GracefulRestartParameters.BGPHelperEnable", ".", true)
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
	//    common.SetValueByPath(updateMap, ".ConfigVersion", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BgpaasParameters.PortStart", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BgpaasParameters.PortEnd", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BGPAlwaysCompareMed", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".AutonomousSystem", ".", 1.0)
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
	//    common.SetValueByPath(updateMap, ".AlarmEnable", ".", true)
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "global_system_config_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var BGPRouterref []interface{}
	//    BGPRouterref = append(BGPRouterref, map[string]interface{}{"operation":"delete", "uuid":"global_system_config_bgp_router_ref_uuid", "to": []string{"test", "global_system_config_bgp_router_ref_uuid"}})
	//    BGPRouterref = append(BGPRouterref, map[string]interface{}{"operation":"add", "uuid":"global_system_config_bgp_router_ref_uuid1", "to": []string{"test", "global_system_config_bgp_router_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "BGPRouterRefs", ".", BGPRouterref)
	//
	//
	_, err = db.CreateGlobalSystemConfig(ctx,
		&models.CreateGlobalSystemConfigRequest{
			GlobalSystemConfig: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateGlobalSystemConfig(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_global_system_config_bgp_router` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing BGPRouterRefs delete statement failed")
		}
		_, err = stmt.Exec("global_system_config_dummy_uuid", "global_system_config_bgp_router_ref_uuid")
		_, err = stmt.Exec("global_system_config_dummy_uuid", "global_system_config_bgp_router_ref_uuid1")
		_, err = stmt.Exec("global_system_config_dummy_uuid", "global_system_config_bgp_router_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "BGPRouterRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteBGPRouter(ctx,
		&models.DeleteBGPRouterRequest{
			ID: "global_system_config_bgp_router_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref global_system_config_bgp_router_ref_uuid  failed", err)
	}
	_, err = db.DeleteBGPRouter(ctx,
		&models.DeleteBGPRouterRequest{
			ID: "global_system_config_bgp_router_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref global_system_config_bgp_router_ref_uuid1  failed", err)
	}
	_, err = db.DeleteBGPRouter(
		ctx,
		&models.DeleteBGPRouterRequest{
			ID: "global_system_config_bgp_router_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref global_system_config_bgp_router_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListGlobalSystemConfig(ctx, &models.ListGlobalSystemConfigRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.GlobalSystemConfigs) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteGlobalSystemConfig(ctxDemo,
		&models.DeleteGlobalSystemConfigRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateGlobalSystemConfig(ctx,
		&models.CreateGlobalSystemConfigRequest{
			GlobalSystemConfig: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteGlobalSystemConfig(ctx,
		&models.DeleteGlobalSystemConfigRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListGlobalSystemConfig(ctx, &models.ListGlobalSystemConfigRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.GlobalSystemConfigs) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
