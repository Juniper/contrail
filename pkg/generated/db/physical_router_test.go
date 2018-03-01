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

func TestPhysicalRouter(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB: testDB,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "physical_router")
	// mutexProject := common.UseTable(db.DB, "physical_router")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakePhysicalRouter()
	model.UUID = "physical_router_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "physical_router_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var BGPRoutercreateref []*models.PhysicalRouterBGPRouterRef
	var BGPRouterrefModel *models.BGPRouter
	BGPRouterrefModel = models.MakeBGPRouter()
	BGPRouterrefModel.UUID = "physical_router_bgp_router_ref_uuid"
	BGPRouterrefModel.FQName = []string{"test", "physical_router_bgp_router_ref_uuid"}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterrefModel,
	})
	BGPRouterrefModel.UUID = "physical_router_bgp_router_ref_uuid1"
	BGPRouterrefModel.FQName = []string{"test", "physical_router_bgp_router_ref_uuid1"}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterrefModel,
	})
	BGPRouterrefModel.UUID = "physical_router_bgp_router_ref_uuid2"
	BGPRouterrefModel.FQName = []string{"test", "physical_router_bgp_router_ref_uuid2"}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BGPRoutercreateref = append(BGPRoutercreateref, &models.PhysicalRouterBGPRouterRef{UUID: "physical_router_bgp_router_ref_uuid", To: []string{"test", "physical_router_bgp_router_ref_uuid"}})
	BGPRoutercreateref = append(BGPRoutercreateref, &models.PhysicalRouterBGPRouterRef{UUID: "physical_router_bgp_router_ref_uuid2", To: []string{"test", "physical_router_bgp_router_ref_uuid2"}})
	model.BGPRouterRefs = BGPRoutercreateref

	var VirtualRoutercreateref []*models.PhysicalRouterVirtualRouterRef
	var VirtualRouterrefModel *models.VirtualRouter
	VirtualRouterrefModel = models.MakeVirtualRouter()
	VirtualRouterrefModel.UUID = "physical_router_virtual_router_ref_uuid"
	VirtualRouterrefModel.FQName = []string{"test", "physical_router_virtual_router_ref_uuid"}
	_, err = db.CreateVirtualRouter(ctx, &models.CreateVirtualRouterRequest{
		VirtualRouter: VirtualRouterrefModel,
	})
	VirtualRouterrefModel.UUID = "physical_router_virtual_router_ref_uuid1"
	VirtualRouterrefModel.FQName = []string{"test", "physical_router_virtual_router_ref_uuid1"}
	_, err = db.CreateVirtualRouter(ctx, &models.CreateVirtualRouterRequest{
		VirtualRouter: VirtualRouterrefModel,
	})
	VirtualRouterrefModel.UUID = "physical_router_virtual_router_ref_uuid2"
	VirtualRouterrefModel.FQName = []string{"test", "physical_router_virtual_router_ref_uuid2"}
	_, err = db.CreateVirtualRouter(ctx, &models.CreateVirtualRouterRequest{
		VirtualRouter: VirtualRouterrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualRoutercreateref = append(VirtualRoutercreateref, &models.PhysicalRouterVirtualRouterRef{UUID: "physical_router_virtual_router_ref_uuid", To: []string{"test", "physical_router_virtual_router_ref_uuid"}})
	VirtualRoutercreateref = append(VirtualRoutercreateref, &models.PhysicalRouterVirtualRouterRef{UUID: "physical_router_virtual_router_ref_uuid2", To: []string{"test", "physical_router_virtual_router_ref_uuid2"}})
	model.VirtualRouterRefs = VirtualRoutercreateref

	var VirtualNetworkcreateref []*models.PhysicalRouterVirtualNetworkRef
	var VirtualNetworkrefModel *models.VirtualNetwork
	VirtualNetworkrefModel = models.MakeVirtualNetwork()
	VirtualNetworkrefModel.UUID = "physical_router_virtual_network_ref_uuid"
	VirtualNetworkrefModel.FQName = []string{"test", "physical_router_virtual_network_ref_uuid"}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkrefModel,
	})
	VirtualNetworkrefModel.UUID = "physical_router_virtual_network_ref_uuid1"
	VirtualNetworkrefModel.FQName = []string{"test", "physical_router_virtual_network_ref_uuid1"}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkrefModel,
	})
	VirtualNetworkrefModel.UUID = "physical_router_virtual_network_ref_uuid2"
	VirtualNetworkrefModel.FQName = []string{"test", "physical_router_virtual_network_ref_uuid2"}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.PhysicalRouterVirtualNetworkRef{UUID: "physical_router_virtual_network_ref_uuid", To: []string{"test", "physical_router_virtual_network_ref_uuid"}})
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.PhysicalRouterVirtualNetworkRef{UUID: "physical_router_virtual_network_ref_uuid2", To: []string{"test", "physical_router_virtual_network_ref_uuid2"}})
	model.VirtualNetworkRefs = VirtualNetworkcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "physical_router_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".TelemetryInfo.ServerPort", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".TelemetryInfo.ServerIP", ".", "test")
	//
	//
	//
	//    if ".TelemetryInfo.Resource" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".TelemetryInfo.Resource", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".TelemetryInfo.Resource", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterVNCManaged", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterVendorName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterUserCredentials.Username", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterUserCredentials.Password", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.Version", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3SecurityName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3SecurityLevel", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3SecurityEngineID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3PrivacyProtocol", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3PrivacyPassword", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3EngineTime", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3EngineID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3EngineBoots", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3ContextEngineID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3Context", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3AuthenticationProtocol", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3AuthenticationPassword", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V2Community", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.Timeout", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.Retries", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.LocalPort", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterSNMP", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterRole", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterProductName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterManagementIP", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterLoopbackIP", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterLLDP", ".", true)
	//
	//
	//
	//    if ".PhysicalRouterJunosServicePorts.ServicePort" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".PhysicalRouterJunosServicePorts.ServicePort", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".PhysicalRouterJunosServicePorts.ServicePort", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterImageURI", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PhysicalRouterDataplaneIP", ".", "test")
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "physical_router_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var VirtualNetworkref []interface{}
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"delete", "uuid":"physical_router_virtual_network_ref_uuid", "to": []string{"test", "physical_router_virtual_network_ref_uuid"}})
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"add", "uuid":"physical_router_virtual_network_ref_uuid1", "to": []string{"test", "physical_router_virtual_network_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualNetworkRefs", ".", VirtualNetworkref)
	//
	//    var BGPRouterref []interface{}
	//    BGPRouterref = append(BGPRouterref, map[string]interface{}{"operation":"delete", "uuid":"physical_router_bgp_router_ref_uuid", "to": []string{"test", "physical_router_bgp_router_ref_uuid"}})
	//    BGPRouterref = append(BGPRouterref, map[string]interface{}{"operation":"add", "uuid":"physical_router_bgp_router_ref_uuid1", "to": []string{"test", "physical_router_bgp_router_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "BGPRouterRefs", ".", BGPRouterref)
	//
	//    var VirtualRouterref []interface{}
	//    VirtualRouterref = append(VirtualRouterref, map[string]interface{}{"operation":"delete", "uuid":"physical_router_virtual_router_ref_uuid", "to": []string{"test", "physical_router_virtual_router_ref_uuid"}})
	//    VirtualRouterref = append(VirtualRouterref, map[string]interface{}{"operation":"add", "uuid":"physical_router_virtual_router_ref_uuid1", "to": []string{"test", "physical_router_virtual_router_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualRouterRefs", ".", VirtualRouterref)
	//
	//
	_, err = db.CreatePhysicalRouter(ctx,
		&models.CreatePhysicalRouterRequest{
			PhysicalRouter: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdatePhysicalRouter(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := common.GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_physical_router_bgp_router` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing BGPRouterRefs delete statement failed")
		}
		_, err = stmt.Exec("physical_router_dummy_uuid", "physical_router_bgp_router_ref_uuid")
		_, err = stmt.Exec("physical_router_dummy_uuid", "physical_router_bgp_router_ref_uuid1")
		_, err = stmt.Exec("physical_router_dummy_uuid", "physical_router_bgp_router_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "BGPRouterRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteBGPRouter(ctx,
		&models.DeleteBGPRouterRequest{
			ID: "physical_router_bgp_router_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref physical_router_bgp_router_ref_uuid  failed", err)
	}
	_, err = db.DeleteBGPRouter(ctx,
		&models.DeleteBGPRouterRequest{
			ID: "physical_router_bgp_router_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref physical_router_bgp_router_ref_uuid1  failed", err)
	}
	_, err = db.DeleteBGPRouter(
		ctx,
		&models.DeleteBGPRouterRequest{
			ID: "physical_router_bgp_router_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref physical_router_bgp_router_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := common.GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_physical_router_virtual_router` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualRouterRefs delete statement failed")
		}
		_, err = stmt.Exec("physical_router_dummy_uuid", "physical_router_virtual_router_ref_uuid")
		_, err = stmt.Exec("physical_router_dummy_uuid", "physical_router_virtual_router_ref_uuid1")
		_, err = stmt.Exec("physical_router_dummy_uuid", "physical_router_virtual_router_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualRouterRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteVirtualRouter(ctx,
		&models.DeleteVirtualRouterRequest{
			ID: "physical_router_virtual_router_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref physical_router_virtual_router_ref_uuid  failed", err)
	}
	_, err = db.DeleteVirtualRouter(ctx,
		&models.DeleteVirtualRouterRequest{
			ID: "physical_router_virtual_router_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref physical_router_virtual_router_ref_uuid1  failed", err)
	}
	_, err = db.DeleteVirtualRouter(
		ctx,
		&models.DeleteVirtualRouterRequest{
			ID: "physical_router_virtual_router_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref physical_router_virtual_router_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := common.GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_physical_router_virtual_network` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualNetworkRefs delete statement failed")
		}
		_, err = stmt.Exec("physical_router_dummy_uuid", "physical_router_virtual_network_ref_uuid")
		_, err = stmt.Exec("physical_router_dummy_uuid", "physical_router_virtual_network_ref_uuid1")
		_, err = stmt.Exec("physical_router_dummy_uuid", "physical_router_virtual_network_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteVirtualNetwork(ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: "physical_router_virtual_network_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref physical_router_virtual_network_ref_uuid  failed", err)
	}
	_, err = db.DeleteVirtualNetwork(ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: "physical_router_virtual_network_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref physical_router_virtual_network_ref_uuid1  failed", err)
	}
	_, err = db.DeleteVirtualNetwork(
		ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: "physical_router_virtual_network_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref physical_router_virtual_network_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListPhysicalRouter(ctx, &models.ListPhysicalRouterRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.PhysicalRouters) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeletePhysicalRouter(ctxDemo,
		&models.DeletePhysicalRouterRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreatePhysicalRouter(ctx,
		&models.CreatePhysicalRouterRequest{
			PhysicalRouter: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeletePhysicalRouter(ctx,
		&models.DeletePhysicalRouterRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListPhysicalRouter(ctx, &models.ListPhysicalRouterRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.PhysicalRouters) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
