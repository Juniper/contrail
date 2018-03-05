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

func TestFirewallPolicy(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "firewall_policy")
	// mutexProject := UseTable(db.DB, "firewall_policy")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeFirewallPolicy()
	model.UUID = "firewall_policy_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "firewall_policy_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var FirewallRulecreateref []*models.FirewallPolicyFirewallRuleRef
	var FirewallRulerefModel *models.FirewallRule
	FirewallRulerefModel = models.MakeFirewallRule()
	FirewallRulerefModel.UUID = "firewall_policy_firewall_rule_ref_uuid"
	FirewallRulerefModel.FQName = []string{"test", "firewall_policy_firewall_rule_ref_uuid"}
	_, err = db.CreateFirewallRule(ctx, &models.CreateFirewallRuleRequest{
		FirewallRule: FirewallRulerefModel,
	})
	FirewallRulerefModel.UUID = "firewall_policy_firewall_rule_ref_uuid1"
	FirewallRulerefModel.FQName = []string{"test", "firewall_policy_firewall_rule_ref_uuid1"}
	_, err = db.CreateFirewallRule(ctx, &models.CreateFirewallRuleRequest{
		FirewallRule: FirewallRulerefModel,
	})
	FirewallRulerefModel.UUID = "firewall_policy_firewall_rule_ref_uuid2"
	FirewallRulerefModel.FQName = []string{"test", "firewall_policy_firewall_rule_ref_uuid2"}
	_, err = db.CreateFirewallRule(ctx, &models.CreateFirewallRuleRequest{
		FirewallRule: FirewallRulerefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	FirewallRulecreateref = append(FirewallRulecreateref, &models.FirewallPolicyFirewallRuleRef{UUID: "firewall_policy_firewall_rule_ref_uuid", To: []string{"test", "firewall_policy_firewall_rule_ref_uuid"}})
	FirewallRulecreateref = append(FirewallRulecreateref, &models.FirewallPolicyFirewallRuleRef{UUID: "firewall_policy_firewall_rule_ref_uuid2", To: []string{"test", "firewall_policy_firewall_rule_ref_uuid2"}})
	model.FirewallRuleRefs = FirewallRulecreateref

	var SecurityLoggingObjectcreateref []*models.FirewallPolicySecurityLoggingObjectRef
	var SecurityLoggingObjectrefModel *models.SecurityLoggingObject
	SecurityLoggingObjectrefModel = models.MakeSecurityLoggingObject()
	SecurityLoggingObjectrefModel.UUID = "firewall_policy_security_logging_object_ref_uuid"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "firewall_policy_security_logging_object_ref_uuid"}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectrefModel,
	})
	SecurityLoggingObjectrefModel.UUID = "firewall_policy_security_logging_object_ref_uuid1"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "firewall_policy_security_logging_object_ref_uuid1"}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectrefModel,
	})
	SecurityLoggingObjectrefModel.UUID = "firewall_policy_security_logging_object_ref_uuid2"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "firewall_policy_security_logging_object_ref_uuid2"}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.FirewallPolicySecurityLoggingObjectRef{UUID: "firewall_policy_security_logging_object_ref_uuid", To: []string{"test", "firewall_policy_security_logging_object_ref_uuid"}})
	SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.FirewallPolicySecurityLoggingObjectRef{UUID: "firewall_policy_security_logging_object_ref_uuid2", To: []string{"test", "firewall_policy_security_logging_object_ref_uuid2"}})
	model.SecurityLoggingObjectRefs = SecurityLoggingObjectcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "firewall_policy_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".ConfigurationVersion", ".", 1.0)
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "firewall_policy_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var FirewallRuleref []interface{}
	//    FirewallRuleref = append(FirewallRuleref, map[string]interface{}{"operation":"delete", "uuid":"firewall_policy_firewall_rule_ref_uuid", "to": []string{"test", "firewall_policy_firewall_rule_ref_uuid"}})
	//    FirewallRuleref = append(FirewallRuleref, map[string]interface{}{"operation":"add", "uuid":"firewall_policy_firewall_rule_ref_uuid1", "to": []string{"test", "firewall_policy_firewall_rule_ref_uuid1"}})
	//
	//    FirewallRuleAttr := map[string]interface{}{}
	//
	//
	//
	//    common.SetValueByPath(FirewallRuleAttr, ".Sequence", ".", "test")
	//
	//
	//
	//    FirewallRuleref = append(FirewallRuleref, map[string]interface{}{"operation":"update", "uuid":"firewall_policy_firewall_rule_ref_uuid2", "to": []string{"test", "firewall_policy_firewall_rule_ref_uuid2"}, "attr": FirewallRuleAttr})
	//
	//    common.SetValueByPath(updateMap, "FirewallRuleRefs", ".", FirewallRuleref)
	//
	//    var SecurityLoggingObjectref []interface{}
	//    SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation":"delete", "uuid":"firewall_policy_security_logging_object_ref_uuid", "to": []string{"test", "firewall_policy_security_logging_object_ref_uuid"}})
	//    SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation":"add", "uuid":"firewall_policy_security_logging_object_ref_uuid1", "to": []string{"test", "firewall_policy_security_logging_object_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "SecurityLoggingObjectRefs", ".", SecurityLoggingObjectref)
	//
	//
	_, err = db.CreateFirewallPolicy(ctx,
		&models.CreateFirewallPolicyRequest{
			FirewallPolicy: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateFirewallPolicy(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_firewall_policy_security_logging_object` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing SecurityLoggingObjectRefs delete statement failed")
		}
		_, err = stmt.Exec("firewall_policy_dummy_uuid", "firewall_policy_security_logging_object_ref_uuid")
		_, err = stmt.Exec("firewall_policy_dummy_uuid", "firewall_policy_security_logging_object_ref_uuid1")
		_, err = stmt.Exec("firewall_policy_dummy_uuid", "firewall_policy_security_logging_object_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteSecurityLoggingObject(ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: "firewall_policy_security_logging_object_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref firewall_policy_security_logging_object_ref_uuid  failed", err)
	}
	_, err = db.DeleteSecurityLoggingObject(ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: "firewall_policy_security_logging_object_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref firewall_policy_security_logging_object_ref_uuid1  failed", err)
	}
	_, err = db.DeleteSecurityLoggingObject(
		ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: "firewall_policy_security_logging_object_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref firewall_policy_security_logging_object_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_firewall_policy_firewall_rule` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing FirewallRuleRefs delete statement failed")
		}
		_, err = stmt.Exec("firewall_policy_dummy_uuid", "firewall_policy_firewall_rule_ref_uuid")
		_, err = stmt.Exec("firewall_policy_dummy_uuid", "firewall_policy_firewall_rule_ref_uuid1")
		_, err = stmt.Exec("firewall_policy_dummy_uuid", "firewall_policy_firewall_rule_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "FirewallRuleRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteFirewallRule(ctx,
		&models.DeleteFirewallRuleRequest{
			ID: "firewall_policy_firewall_rule_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref firewall_policy_firewall_rule_ref_uuid  failed", err)
	}
	_, err = db.DeleteFirewallRule(ctx,
		&models.DeleteFirewallRuleRequest{
			ID: "firewall_policy_firewall_rule_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref firewall_policy_firewall_rule_ref_uuid1  failed", err)
	}
	_, err = db.DeleteFirewallRule(
		ctx,
		&models.DeleteFirewallRuleRequest{
			ID: "firewall_policy_firewall_rule_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref firewall_policy_firewall_rule_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListFirewallPolicy(ctx, &models.ListFirewallPolicyRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.FirewallPolicys) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteFirewallPolicy(ctxDemo,
		&models.DeleteFirewallPolicyRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateFirewallPolicy(ctx,
		&models.CreateFirewallPolicyRequest{
			FirewallPolicy: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteFirewallPolicy(ctx,
		&models.DeleteFirewallPolicyRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListFirewallPolicy(ctx, &models.ListFirewallPolicyRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.FirewallPolicys) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
