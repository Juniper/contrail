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

func TestSecurityLoggingObject(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "security_logging_object")
	// mutexProject := UseTable(db.DB, "security_logging_object")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeSecurityLoggingObject()
	model.UUID = "security_logging_object_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "security_logging_object_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var SecurityGroupcreateref []*models.SecurityLoggingObjectSecurityGroupRef
	var SecurityGrouprefModel *models.SecurityGroup
	SecurityGrouprefModel = models.MakeSecurityGroup()
	SecurityGrouprefModel.UUID = "security_logging_object_security_group_ref_uuid"
	SecurityGrouprefModel.FQName = []string{"test", "security_logging_object_security_group_ref_uuid"}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGrouprefModel,
	})
	SecurityGrouprefModel.UUID = "security_logging_object_security_group_ref_uuid1"
	SecurityGrouprefModel.FQName = []string{"test", "security_logging_object_security_group_ref_uuid1"}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGrouprefModel,
	})
	SecurityGrouprefModel.UUID = "security_logging_object_security_group_ref_uuid2"
	SecurityGrouprefModel.FQName = []string{"test", "security_logging_object_security_group_ref_uuid2"}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGrouprefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityGroupcreateref = append(SecurityGroupcreateref, &models.SecurityLoggingObjectSecurityGroupRef{UUID: "security_logging_object_security_group_ref_uuid", To: []string{"test", "security_logging_object_security_group_ref_uuid"}})
	SecurityGroupcreateref = append(SecurityGroupcreateref, &models.SecurityLoggingObjectSecurityGroupRef{UUID: "security_logging_object_security_group_ref_uuid2", To: []string{"test", "security_logging_object_security_group_ref_uuid2"}})
	model.SecurityGroupRefs = SecurityGroupcreateref

	var NetworkPolicycreateref []*models.SecurityLoggingObjectNetworkPolicyRef
	var NetworkPolicyrefModel *models.NetworkPolicy
	NetworkPolicyrefModel = models.MakeNetworkPolicy()
	NetworkPolicyrefModel.UUID = "security_logging_object_network_policy_ref_uuid"
	NetworkPolicyrefModel.FQName = []string{"test", "security_logging_object_network_policy_ref_uuid"}
	_, err = db.CreateNetworkPolicy(ctx, &models.CreateNetworkPolicyRequest{
		NetworkPolicy: NetworkPolicyrefModel,
	})
	NetworkPolicyrefModel.UUID = "security_logging_object_network_policy_ref_uuid1"
	NetworkPolicyrefModel.FQName = []string{"test", "security_logging_object_network_policy_ref_uuid1"}
	_, err = db.CreateNetworkPolicy(ctx, &models.CreateNetworkPolicyRequest{
		NetworkPolicy: NetworkPolicyrefModel,
	})
	NetworkPolicyrefModel.UUID = "security_logging_object_network_policy_ref_uuid2"
	NetworkPolicyrefModel.FQName = []string{"test", "security_logging_object_network_policy_ref_uuid2"}
	_, err = db.CreateNetworkPolicy(ctx, &models.CreateNetworkPolicyRequest{
		NetworkPolicy: NetworkPolicyrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NetworkPolicycreateref = append(NetworkPolicycreateref, &models.SecurityLoggingObjectNetworkPolicyRef{UUID: "security_logging_object_network_policy_ref_uuid", To: []string{"test", "security_logging_object_network_policy_ref_uuid"}})
	NetworkPolicycreateref = append(NetworkPolicycreateref, &models.SecurityLoggingObjectNetworkPolicyRef{UUID: "security_logging_object_network_policy_ref_uuid2", To: []string{"test", "security_logging_object_network_policy_ref_uuid2"}})
	model.NetworkPolicyRefs = NetworkPolicycreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "security_logging_object_admin_project_uuid"
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
	//    if ".SecurityLoggingObjectRules.Rule" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".SecurityLoggingObjectRules.Rule", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".SecurityLoggingObjectRules.Rule", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".SecurityLoggingObjectRate", ".", 1.0)
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "security_logging_object_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var SecurityGroupref []interface{}
	//    SecurityGroupref = append(SecurityGroupref, map[string]interface{}{"operation":"delete", "uuid":"security_logging_object_security_group_ref_uuid", "to": []string{"test", "security_logging_object_security_group_ref_uuid"}})
	//    SecurityGroupref = append(SecurityGroupref, map[string]interface{}{"operation":"add", "uuid":"security_logging_object_security_group_ref_uuid1", "to": []string{"test", "security_logging_object_security_group_ref_uuid1"}})
	//
	//    SecurityGroupAttr := map[string]interface{}{}
	//
	//
	//
	//    common.SetValueByPath(SecurityGroupAttr, ".Rule", ".", map[string]string{"test": "test"})
	//
	//
	//
	//    SecurityGroupref = append(SecurityGroupref, map[string]interface{}{"operation":"update", "uuid":"security_logging_object_security_group_ref_uuid2", "to": []string{"test", "security_logging_object_security_group_ref_uuid2"}, "attr": SecurityGroupAttr})
	//
	//    common.SetValueByPath(updateMap, "SecurityGroupRefs", ".", SecurityGroupref)
	//
	//    var NetworkPolicyref []interface{}
	//    NetworkPolicyref = append(NetworkPolicyref, map[string]interface{}{"operation":"delete", "uuid":"security_logging_object_network_policy_ref_uuid", "to": []string{"test", "security_logging_object_network_policy_ref_uuid"}})
	//    NetworkPolicyref = append(NetworkPolicyref, map[string]interface{}{"operation":"add", "uuid":"security_logging_object_network_policy_ref_uuid1", "to": []string{"test", "security_logging_object_network_policy_ref_uuid1"}})
	//
	//    NetworkPolicyAttr := map[string]interface{}{}
	//
	//
	//
	//    common.SetValueByPath(NetworkPolicyAttr, ".Rule", ".", map[string]string{"test": "test"})
	//
	//
	//
	//    NetworkPolicyref = append(NetworkPolicyref, map[string]interface{}{"operation":"update", "uuid":"security_logging_object_network_policy_ref_uuid2", "to": []string{"test", "security_logging_object_network_policy_ref_uuid2"}, "attr": NetworkPolicyAttr})
	//
	//    common.SetValueByPath(updateMap, "NetworkPolicyRefs", ".", NetworkPolicyref)
	//
	//
	_, err = db.CreateSecurityLoggingObject(ctx,
		&models.CreateSecurityLoggingObjectRequest{
			SecurityLoggingObject: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateSecurityLoggingObject(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_security_logging_object_security_group` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing SecurityGroupRefs delete statement failed")
		}
		_, err = stmt.Exec("security_logging_object_dummy_uuid", "security_logging_object_security_group_ref_uuid")
		_, err = stmt.Exec("security_logging_object_dummy_uuid", "security_logging_object_security_group_ref_uuid1")
		_, err = stmt.Exec("security_logging_object_dummy_uuid", "security_logging_object_security_group_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "SecurityGroupRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteSecurityGroup(ctx,
		&models.DeleteSecurityGroupRequest{
			ID: "security_logging_object_security_group_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref security_logging_object_security_group_ref_uuid  failed", err)
	}
	_, err = db.DeleteSecurityGroup(ctx,
		&models.DeleteSecurityGroupRequest{
			ID: "security_logging_object_security_group_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref security_logging_object_security_group_ref_uuid1  failed", err)
	}
	_, err = db.DeleteSecurityGroup(
		ctx,
		&models.DeleteSecurityGroupRequest{
			ID: "security_logging_object_security_group_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref security_logging_object_security_group_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_security_logging_object_network_policy` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing NetworkPolicyRefs delete statement failed")
		}
		_, err = stmt.Exec("security_logging_object_dummy_uuid", "security_logging_object_network_policy_ref_uuid")
		_, err = stmt.Exec("security_logging_object_dummy_uuid", "security_logging_object_network_policy_ref_uuid1")
		_, err = stmt.Exec("security_logging_object_dummy_uuid", "security_logging_object_network_policy_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "NetworkPolicyRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteNetworkPolicy(ctx,
		&models.DeleteNetworkPolicyRequest{
			ID: "security_logging_object_network_policy_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref security_logging_object_network_policy_ref_uuid  failed", err)
	}
	_, err = db.DeleteNetworkPolicy(ctx,
		&models.DeleteNetworkPolicyRequest{
			ID: "security_logging_object_network_policy_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref security_logging_object_network_policy_ref_uuid1  failed", err)
	}
	_, err = db.DeleteNetworkPolicy(
		ctx,
		&models.DeleteNetworkPolicyRequest{
			ID: "security_logging_object_network_policy_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref security_logging_object_network_policy_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListSecurityLoggingObject(ctx, &models.ListSecurityLoggingObjectRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.SecurityLoggingObjects) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteSecurityLoggingObject(ctxDemo,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateSecurityLoggingObject(ctx,
		&models.CreateSecurityLoggingObjectRequest{
			SecurityLoggingObject: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteSecurityLoggingObject(ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListSecurityLoggingObject(ctx, &models.ListSecurityLoggingObjectRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.SecurityLoggingObjects) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
