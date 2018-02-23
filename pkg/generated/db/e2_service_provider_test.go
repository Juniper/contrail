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

func TestE2ServiceProvider(t *testing.T) {
	// t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db, "metadata")
	mutexTable := common.UseTable(db, "e2_service_provider")
	// mutexProject := common.UseTable(db, "e2_service_provider")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeE2ServiceProvider()
	model.UUID = "e2_service_provider_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "e2_service_provider_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var PhysicalRoutercreateref []*models.E2ServiceProviderPhysicalRouterRef
	var PhysicalRouterrefModel *models.PhysicalRouter
	PhysicalRouterrefModel = models.MakePhysicalRouter()
	PhysicalRouterrefModel.UUID = "e2_service_provider_physical_router_ref_uuid"
	PhysicalRouterrefModel.FQName = []string{"test", "e2_service_provider_physical_router_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
			PhysicalRouter: PhysicalRouterrefModel,
		})
	})
	PhysicalRouterrefModel.UUID = "e2_service_provider_physical_router_ref_uuid1"
	PhysicalRouterrefModel.FQName = []string{"test", "e2_service_provider_physical_router_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
			PhysicalRouter: PhysicalRouterrefModel,
		})
	})
	PhysicalRouterrefModel.UUID = "e2_service_provider_physical_router_ref_uuid2"
	PhysicalRouterrefModel.FQName = []string{"test", "e2_service_provider_physical_router_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
			PhysicalRouter: PhysicalRouterrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	PhysicalRoutercreateref = append(PhysicalRoutercreateref, &models.E2ServiceProviderPhysicalRouterRef{UUID: "e2_service_provider_physical_router_ref_uuid", To: []string{"test", "e2_service_provider_physical_router_ref_uuid"}})
	PhysicalRoutercreateref = append(PhysicalRoutercreateref, &models.E2ServiceProviderPhysicalRouterRef{UUID: "e2_service_provider_physical_router_ref_uuid2", To: []string{"test", "e2_service_provider_physical_router_ref_uuid2"}})
	model.PhysicalRouterRefs = PhysicalRoutercreateref

	var PeeringPolicycreateref []*models.E2ServiceProviderPeeringPolicyRef
	var PeeringPolicyrefModel *models.PeeringPolicy
	PeeringPolicyrefModel = models.MakePeeringPolicy()
	PeeringPolicyrefModel.UUID = "e2_service_provider_peering_policy_ref_uuid"
	PeeringPolicyrefModel.FQName = []string{"test", "e2_service_provider_peering_policy_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePeeringPolicy(ctx, tx, &models.CreatePeeringPolicyRequest{
			PeeringPolicy: PeeringPolicyrefModel,
		})
	})
	PeeringPolicyrefModel.UUID = "e2_service_provider_peering_policy_ref_uuid1"
	PeeringPolicyrefModel.FQName = []string{"test", "e2_service_provider_peering_policy_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePeeringPolicy(ctx, tx, &models.CreatePeeringPolicyRequest{
			PeeringPolicy: PeeringPolicyrefModel,
		})
	})
	PeeringPolicyrefModel.UUID = "e2_service_provider_peering_policy_ref_uuid2"
	PeeringPolicyrefModel.FQName = []string{"test", "e2_service_provider_peering_policy_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePeeringPolicy(ctx, tx, &models.CreatePeeringPolicyRequest{
			PeeringPolicy: PeeringPolicyrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	PeeringPolicycreateref = append(PeeringPolicycreateref, &models.E2ServiceProviderPeeringPolicyRef{UUID: "e2_service_provider_peering_policy_ref_uuid", To: []string{"test", "e2_service_provider_peering_policy_ref_uuid"}})
	PeeringPolicycreateref = append(PeeringPolicycreateref, &models.E2ServiceProviderPeeringPolicyRef{UUID: "e2_service_provider_peering_policy_ref_uuid2", To: []string{"test", "e2_service_provider_peering_policy_ref_uuid2"}})
	model.PeeringPolicyRefs = PeeringPolicycreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "e2_service_provider_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".E2ServiceProviderPromiscuous", ".", true)
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "e2_service_provider_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var PhysicalRouterref []interface{}
	//    PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation":"delete", "uuid":"e2_service_provider_physical_router_ref_uuid", "to": []string{"test", "e2_service_provider_physical_router_ref_uuid"}})
	//    PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation":"add", "uuid":"e2_service_provider_physical_router_ref_uuid1", "to": []string{"test", "e2_service_provider_physical_router_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "PhysicalRouterRefs", ".", PhysicalRouterref)
	//
	//    var PeeringPolicyref []interface{}
	//    PeeringPolicyref = append(PeeringPolicyref, map[string]interface{}{"operation":"delete", "uuid":"e2_service_provider_peering_policy_ref_uuid", "to": []string{"test", "e2_service_provider_peering_policy_ref_uuid"}})
	//    PeeringPolicyref = append(PeeringPolicyref, map[string]interface{}{"operation":"add", "uuid":"e2_service_provider_peering_policy_ref_uuid1", "to": []string{"test", "e2_service_provider_peering_policy_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "PeeringPolicyRefs", ".", PeeringPolicyref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateE2ServiceProvider(ctx, tx,
			&models.CreateE2ServiceProviderRequest{
				E2ServiceProvider: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateE2ServiceProvider(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_e2_service_provider_physical_router` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing PhysicalRouterRefs delete statement failed")
		}
		_, err = stmt.Exec("e2_service_provider_dummy_uuid", "e2_service_provider_physical_router_ref_uuid")
		_, err = stmt.Exec("e2_service_provider_dummy_uuid", "e2_service_provider_physical_router_ref_uuid1")
		_, err = stmt.Exec("e2_service_provider_dummy_uuid", "e2_service_provider_physical_router_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(ctx, tx,
			&models.DeletePhysicalRouterRequest{
				ID: "e2_service_provider_physical_router_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref e2_service_provider_physical_router_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(ctx, tx,
			&models.DeletePhysicalRouterRequest{
				ID: "e2_service_provider_physical_router_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref e2_service_provider_physical_router_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(
			ctx,
			tx,
			&models.DeletePhysicalRouterRequest{
				ID: "e2_service_provider_physical_router_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref e2_service_provider_physical_router_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_e2_service_provider_peering_policy` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing PeeringPolicyRefs delete statement failed")
		}
		_, err = stmt.Exec("e2_service_provider_dummy_uuid", "e2_service_provider_peering_policy_ref_uuid")
		_, err = stmt.Exec("e2_service_provider_dummy_uuid", "e2_service_provider_peering_policy_ref_uuid1")
		_, err = stmt.Exec("e2_service_provider_dummy_uuid", "e2_service_provider_peering_policy_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "PeeringPolicyRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePeeringPolicy(ctx, tx,
			&models.DeletePeeringPolicyRequest{
				ID: "e2_service_provider_peering_policy_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref e2_service_provider_peering_policy_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePeeringPolicy(ctx, tx,
			&models.DeletePeeringPolicyRequest{
				ID: "e2_service_provider_peering_policy_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref e2_service_provider_peering_policy_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePeeringPolicy(
			ctx,
			tx,
			&models.DeletePeeringPolicyRequest{
				ID: "e2_service_provider_peering_policy_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref e2_service_provider_peering_policy_ref_uuid2 failed", err)
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
		response, err := ListE2ServiceProvider(ctx, tx, &models.ListE2ServiceProviderRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.E2ServiceProviders) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteE2ServiceProvider(ctxDemo, tx,
			&models.DeleteE2ServiceProviderRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteE2ServiceProvider(ctx, tx,
			&models.DeleteE2ServiceProviderRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateE2ServiceProvider(ctx, tx,
			&models.CreateE2ServiceProviderRequest{
				E2ServiceProvider: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListE2ServiceProvider(ctx, tx, &models.ListE2ServiceProviderRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.E2ServiceProviders) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
