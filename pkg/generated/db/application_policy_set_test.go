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

func TestApplicationPolicySet(t *testing.T) {
	t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	common.UseTable(db, "metadata")
	common.UseTable(db, "application_policy_set")
	defer func() {
		common.ClearTable(db, "application_policy_set")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeApplicationPolicySet()
	model.UUID = "application_policy_set_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "application_policy_set_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var FirewallPolicycreateref []*models.ApplicationPolicySetFirewallPolicyRef
	var FirewallPolicyrefModel *models.FirewallPolicy
	FirewallPolicyrefModel = models.MakeFirewallPolicy()
	FirewallPolicyrefModel.UUID = "application_policy_set_firewall_policy_ref_uuid"
	FirewallPolicyrefModel.FQName = []string{"test", "application_policy_set_firewall_policy_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFirewallPolicy(ctx, tx, &models.CreateFirewallPolicyRequest{
			FirewallPolicy: FirewallPolicyrefModel,
		})
	})
	FirewallPolicyrefModel.UUID = "application_policy_set_firewall_policy_ref_uuid1"
	FirewallPolicyrefModel.FQName = []string{"test", "application_policy_set_firewall_policy_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFirewallPolicy(ctx, tx, &models.CreateFirewallPolicyRequest{
			FirewallPolicy: FirewallPolicyrefModel,
		})
	})
	FirewallPolicyrefModel.UUID = "application_policy_set_firewall_policy_ref_uuid2"
	FirewallPolicyrefModel.FQName = []string{"test", "application_policy_set_firewall_policy_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFirewallPolicy(ctx, tx, &models.CreateFirewallPolicyRequest{
			FirewallPolicy: FirewallPolicyrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	FirewallPolicycreateref = append(FirewallPolicycreateref, &models.ApplicationPolicySetFirewallPolicyRef{UUID: "application_policy_set_firewall_policy_ref_uuid", To: []string{"test", "application_policy_set_firewall_policy_ref_uuid"}})
	FirewallPolicycreateref = append(FirewallPolicycreateref, &models.ApplicationPolicySetFirewallPolicyRef{UUID: "application_policy_set_firewall_policy_ref_uuid2", To: []string{"test", "application_policy_set_firewall_policy_ref_uuid2"}})
	model.FirewallPolicyRefs = FirewallPolicycreateref

	var GlobalVrouterConfigcreateref []*models.ApplicationPolicySetGlobalVrouterConfigRef
	var GlobalVrouterConfigrefModel *models.GlobalVrouterConfig
	GlobalVrouterConfigrefModel = models.MakeGlobalVrouterConfig()
	GlobalVrouterConfigrefModel.UUID = "application_policy_set_global_vrouter_config_ref_uuid"
	GlobalVrouterConfigrefModel.FQName = []string{"test", "application_policy_set_global_vrouter_config_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateGlobalVrouterConfig(ctx, tx, &models.CreateGlobalVrouterConfigRequest{
			GlobalVrouterConfig: GlobalVrouterConfigrefModel,
		})
	})
	GlobalVrouterConfigrefModel.UUID = "application_policy_set_global_vrouter_config_ref_uuid1"
	GlobalVrouterConfigrefModel.FQName = []string{"test", "application_policy_set_global_vrouter_config_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateGlobalVrouterConfig(ctx, tx, &models.CreateGlobalVrouterConfigRequest{
			GlobalVrouterConfig: GlobalVrouterConfigrefModel,
		})
	})
	GlobalVrouterConfigrefModel.UUID = "application_policy_set_global_vrouter_config_ref_uuid2"
	GlobalVrouterConfigrefModel.FQName = []string{"test", "application_policy_set_global_vrouter_config_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateGlobalVrouterConfig(ctx, tx, &models.CreateGlobalVrouterConfigRequest{
			GlobalVrouterConfig: GlobalVrouterConfigrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	GlobalVrouterConfigcreateref = append(GlobalVrouterConfigcreateref, &models.ApplicationPolicySetGlobalVrouterConfigRef{UUID: "application_policy_set_global_vrouter_config_ref_uuid", To: []string{"test", "application_policy_set_global_vrouter_config_ref_uuid"}})
	GlobalVrouterConfigcreateref = append(GlobalVrouterConfigcreateref, &models.ApplicationPolicySetGlobalVrouterConfigRef{UUID: "application_policy_set_global_vrouter_config_ref_uuid2", To: []string{"test", "application_policy_set_global_vrouter_config_ref_uuid2"}})
	model.GlobalVrouterConfigRefs = GlobalVrouterConfigcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "application_policy_set_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".AllApplications", ".", true)
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "application_policy_set_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var FirewallPolicyref []interface{}
	//    FirewallPolicyref = append(FirewallPolicyref, map[string]interface{}{"operation":"delete", "uuid":"application_policy_set_firewall_policy_ref_uuid", "to": []string{"test", "application_policy_set_firewall_policy_ref_uuid"}})
	//    FirewallPolicyref = append(FirewallPolicyref, map[string]interface{}{"operation":"add", "uuid":"application_policy_set_firewall_policy_ref_uuid1", "to": []string{"test", "application_policy_set_firewall_policy_ref_uuid1"}})
	//
	//    FirewallPolicyAttr := map[string]interface{}{}
	//
	//
	//
	//    common.SetValueByPath(FirewallPolicyAttr, ".Sequence", ".", "test")
	//
	//
	//
	//    FirewallPolicyref = append(FirewallPolicyref, map[string]interface{}{"operation":"update", "uuid":"application_policy_set_firewall_policy_ref_uuid2", "to": []string{"test", "application_policy_set_firewall_policy_ref_uuid2"}, "attr": FirewallPolicyAttr})
	//
	//    common.SetValueByPath(updateMap, "FirewallPolicyRefs", ".", FirewallPolicyref)
	//
	//    var GlobalVrouterConfigref []interface{}
	//    GlobalVrouterConfigref = append(GlobalVrouterConfigref, map[string]interface{}{"operation":"delete", "uuid":"application_policy_set_global_vrouter_config_ref_uuid", "to": []string{"test", "application_policy_set_global_vrouter_config_ref_uuid"}})
	//    GlobalVrouterConfigref = append(GlobalVrouterConfigref, map[string]interface{}{"operation":"add", "uuid":"application_policy_set_global_vrouter_config_ref_uuid1", "to": []string{"test", "application_policy_set_global_vrouter_config_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "GlobalVrouterConfigRefs", ".", GlobalVrouterConfigref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateApplicationPolicySet(ctx, tx,
			&models.CreateApplicationPolicySetRequest{
				ApplicationPolicySet: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateApplicationPolicySet(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_application_policy_set_global_vrouter_config` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing GlobalVrouterConfigRefs delete statement failed")
		}
		_, err = stmt.Exec("application_policy_set_dummy_uuid", "application_policy_set_global_vrouter_config_ref_uuid")
		_, err = stmt.Exec("application_policy_set_dummy_uuid", "application_policy_set_global_vrouter_config_ref_uuid1")
		_, err = stmt.Exec("application_policy_set_dummy_uuid", "application_policy_set_global_vrouter_config_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "GlobalVrouterConfigRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteGlobalVrouterConfig(ctx, tx,
			&models.DeleteGlobalVrouterConfigRequest{
				ID: "application_policy_set_global_vrouter_config_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref application_policy_set_global_vrouter_config_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteGlobalVrouterConfig(ctx, tx,
			&models.DeleteGlobalVrouterConfigRequest{
				ID: "application_policy_set_global_vrouter_config_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref application_policy_set_global_vrouter_config_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteGlobalVrouterConfig(
			ctx,
			tx,
			&models.DeleteGlobalVrouterConfigRequest{
				ID: "application_policy_set_global_vrouter_config_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref application_policy_set_global_vrouter_config_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_application_policy_set_firewall_policy` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing FirewallPolicyRefs delete statement failed")
		}
		_, err = stmt.Exec("application_policy_set_dummy_uuid", "application_policy_set_firewall_policy_ref_uuid")
		_, err = stmt.Exec("application_policy_set_dummy_uuid", "application_policy_set_firewall_policy_ref_uuid1")
		_, err = stmt.Exec("application_policy_set_dummy_uuid", "application_policy_set_firewall_policy_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "FirewallPolicyRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteFirewallPolicy(ctx, tx,
			&models.DeleteFirewallPolicyRequest{
				ID: "application_policy_set_firewall_policy_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref application_policy_set_firewall_policy_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteFirewallPolicy(ctx, tx,
			&models.DeleteFirewallPolicyRequest{
				ID: "application_policy_set_firewall_policy_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref application_policy_set_firewall_policy_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteFirewallPolicy(
			ctx,
			tx,
			&models.DeleteFirewallPolicyRequest{
				ID: "application_policy_set_firewall_policy_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref application_policy_set_firewall_policy_ref_uuid2 failed", err)
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
		response, err := ListApplicationPolicySet(ctx, tx, &models.ListApplicationPolicySetRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.ApplicationPolicySets) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteApplicationPolicySet(ctxDemo, tx,
			&models.DeleteApplicationPolicySetRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteApplicationPolicySet(ctx, tx,
			&models.DeleteApplicationPolicySetRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateApplicationPolicySet(ctx, tx,
			&models.CreateApplicationPolicySetRequest{
				ApplicationPolicySet: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListApplicationPolicySet(ctx, tx, &models.ListApplicationPolicySetRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.ApplicationPolicySets) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
