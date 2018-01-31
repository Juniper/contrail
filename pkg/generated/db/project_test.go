package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

func TestProject(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "project")
	defer func() {
		common.ClearTable(db, "project")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeProject()
	model.UUID = "project_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "project_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var FloatingIPPoolcreateref []*models.ProjectFloatingIPPoolRef
	var FloatingIPPoolrefModel *models.FloatingIPPool
	FloatingIPPoolrefModel = models.MakeFloatingIPPool()
	FloatingIPPoolrefModel.UUID = "project_floating_ip_pool_ref_uuid"
	FloatingIPPoolrefModel.FQName = []string{"test", "project_floating_ip_pool_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFloatingIPPool(tx, FloatingIPPoolrefModel)
	})
	FloatingIPPoolrefModel.UUID = "project_floating_ip_pool_ref_uuid1"
	FloatingIPPoolrefModel.FQName = []string{"test", "project_floating_ip_pool_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFloatingIPPool(tx, FloatingIPPoolrefModel)
	})
	FloatingIPPoolrefModel.UUID = "project_floating_ip_pool_ref_uuid2"
	FloatingIPPoolrefModel.FQName = []string{"test", "project_floating_ip_pool_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFloatingIPPool(tx, FloatingIPPoolrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	FloatingIPPoolcreateref = append(FloatingIPPoolcreateref, &models.ProjectFloatingIPPoolRef{UUID: "project_floating_ip_pool_ref_uuid", To: []string{"test", "project_floating_ip_pool_ref_uuid"}})
	FloatingIPPoolcreateref = append(FloatingIPPoolcreateref, &models.ProjectFloatingIPPoolRef{UUID: "project_floating_ip_pool_ref_uuid2", To: []string{"test", "project_floating_ip_pool_ref_uuid2"}})
	model.FloatingIPPoolRefs = FloatingIPPoolcreateref

	var AliasIPPoolcreateref []*models.ProjectAliasIPPoolRef
	var AliasIPPoolrefModel *models.AliasIPPool
	AliasIPPoolrefModel = models.MakeAliasIPPool()
	AliasIPPoolrefModel.UUID = "project_alias_ip_pool_ref_uuid"
	AliasIPPoolrefModel.FQName = []string{"test", "project_alias_ip_pool_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateAliasIPPool(tx, AliasIPPoolrefModel)
	})
	AliasIPPoolrefModel.UUID = "project_alias_ip_pool_ref_uuid1"
	AliasIPPoolrefModel.FQName = []string{"test", "project_alias_ip_pool_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateAliasIPPool(tx, AliasIPPoolrefModel)
	})
	AliasIPPoolrefModel.UUID = "project_alias_ip_pool_ref_uuid2"
	AliasIPPoolrefModel.FQName = []string{"test", "project_alias_ip_pool_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateAliasIPPool(tx, AliasIPPoolrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	AliasIPPoolcreateref = append(AliasIPPoolcreateref, &models.ProjectAliasIPPoolRef{UUID: "project_alias_ip_pool_ref_uuid", To: []string{"test", "project_alias_ip_pool_ref_uuid"}})
	AliasIPPoolcreateref = append(AliasIPPoolcreateref, &models.ProjectAliasIPPoolRef{UUID: "project_alias_ip_pool_ref_uuid2", To: []string{"test", "project_alias_ip_pool_ref_uuid2"}})
	model.AliasIPPoolRefs = AliasIPPoolcreateref

	var Namespacecreateref []*models.ProjectNamespaceRef
	var NamespacerefModel *models.Namespace
	NamespacerefModel = models.MakeNamespace()
	NamespacerefModel.UUID = "project_namespace_ref_uuid"
	NamespacerefModel.FQName = []string{"test", "project_namespace_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNamespace(tx, NamespacerefModel)
	})
	NamespacerefModel.UUID = "project_namespace_ref_uuid1"
	NamespacerefModel.FQName = []string{"test", "project_namespace_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNamespace(tx, NamespacerefModel)
	})
	NamespacerefModel.UUID = "project_namespace_ref_uuid2"
	NamespacerefModel.FQName = []string{"test", "project_namespace_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNamespace(tx, NamespacerefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	Namespacecreateref = append(Namespacecreateref, &models.ProjectNamespaceRef{UUID: "project_namespace_ref_uuid", To: []string{"test", "project_namespace_ref_uuid"}})
	Namespacecreateref = append(Namespacecreateref, &models.ProjectNamespaceRef{UUID: "project_namespace_ref_uuid2", To: []string{"test", "project_namespace_ref_uuid2"}})
	model.NamespaceRefs = Namespacecreateref

	var ApplicationPolicySetcreateref []*models.ProjectApplicationPolicySetRef
	var ApplicationPolicySetrefModel *models.ApplicationPolicySet
	ApplicationPolicySetrefModel = models.MakeApplicationPolicySet()
	ApplicationPolicySetrefModel.UUID = "project_application_policy_set_ref_uuid"
	ApplicationPolicySetrefModel.FQName = []string{"test", "project_application_policy_set_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateApplicationPolicySet(tx, ApplicationPolicySetrefModel)
	})
	ApplicationPolicySetrefModel.UUID = "project_application_policy_set_ref_uuid1"
	ApplicationPolicySetrefModel.FQName = []string{"test", "project_application_policy_set_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateApplicationPolicySet(tx, ApplicationPolicySetrefModel)
	})
	ApplicationPolicySetrefModel.UUID = "project_application_policy_set_ref_uuid2"
	ApplicationPolicySetrefModel.FQName = []string{"test", "project_application_policy_set_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateApplicationPolicySet(tx, ApplicationPolicySetrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ApplicationPolicySetcreateref = append(ApplicationPolicySetcreateref, &models.ProjectApplicationPolicySetRef{UUID: "project_application_policy_set_ref_uuid", To: []string{"test", "project_application_policy_set_ref_uuid"}})
	ApplicationPolicySetcreateref = append(ApplicationPolicySetcreateref, &models.ProjectApplicationPolicySetRef{UUID: "project_application_policy_set_ref_uuid2", To: []string{"test", "project_application_policy_set_ref_uuid2"}})
	model.ApplicationPolicySetRefs = ApplicationPolicySetcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "project_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(tx, projectModel)
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//populate update map
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".VxlanRouting", ".", true)

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".Quota.VirtualRouter", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualNetwork", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualMachineInterface", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualIP", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualDNSRecord", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualDNS", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.Subnet", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.ServiceTemplate", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.ServiceInstance", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.SecurityLoggingObject", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.SecurityGroupRule", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.SecurityGroup", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.RouteTable", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.NetworkPolicy", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.NetworkIpam", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.LogicalRouter", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.LoadbalancerPool", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.LoadbalancerMember", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.LoadbalancerHealthmonitor", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.InstanceIP", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.GlobalVrouterConfig", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.FloatingIPPool", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.FloatingIP", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.Defaults", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.BGPRouter", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.AccessControlList", ".", 1.0)

	if ".Perms2.Share" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ParentType", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)

	common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")

	if ".FQName" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".FQName", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	if ".Annotations.KeyValuePair" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".AlarmEnable", ".", true)

	common.SetValueByPath(updateMap, "uuid", ".", "project_dummy_uuid")
	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	// Create Attr values for testing ref update(ADD,UPDATE,DELETE)

	var AliasIPPoolref []interface{}
	AliasIPPoolref = append(AliasIPPoolref, map[string]interface{}{"operation": "delete", "uuid": "project_alias_ip_pool_ref_uuid", "to": []string{"test", "project_alias_ip_pool_ref_uuid"}})
	AliasIPPoolref = append(AliasIPPoolref, map[string]interface{}{"operation": "add", "uuid": "project_alias_ip_pool_ref_uuid1", "to": []string{"test", "project_alias_ip_pool_ref_uuid1"}})

	common.SetValueByPath(updateMap, "AliasIPPoolRefs", ".", AliasIPPoolref)

	var Namespaceref []interface{}
	Namespaceref = append(Namespaceref, map[string]interface{}{"operation": "delete", "uuid": "project_namespace_ref_uuid", "to": []string{"test", "project_namespace_ref_uuid"}})
	Namespaceref = append(Namespaceref, map[string]interface{}{"operation": "add", "uuid": "project_namespace_ref_uuid1", "to": []string{"test", "project_namespace_ref_uuid1"}})

	NamespaceAttr := map[string]interface{}{}

	common.SetValueByPath(NamespaceAttr, ".IPPrefix", ".", "test")

	common.SetValueByPath(NamespaceAttr, ".IPPrefixLen", ".", 1.0)

	Namespaceref = append(Namespaceref, map[string]interface{}{"operation": "update", "uuid": "project_namespace_ref_uuid2", "to": []string{"test", "project_namespace_ref_uuid2"}, "attr": NamespaceAttr})

	common.SetValueByPath(updateMap, "NamespaceRefs", ".", Namespaceref)

	var ApplicationPolicySetref []interface{}
	ApplicationPolicySetref = append(ApplicationPolicySetref, map[string]interface{}{"operation": "delete", "uuid": "project_application_policy_set_ref_uuid", "to": []string{"test", "project_application_policy_set_ref_uuid"}})
	ApplicationPolicySetref = append(ApplicationPolicySetref, map[string]interface{}{"operation": "add", "uuid": "project_application_policy_set_ref_uuid1", "to": []string{"test", "project_application_policy_set_ref_uuid1"}})

	common.SetValueByPath(updateMap, "ApplicationPolicySetRefs", ".", ApplicationPolicySetref)

	var FloatingIPPoolref []interface{}
	FloatingIPPoolref = append(FloatingIPPoolref, map[string]interface{}{"operation": "delete", "uuid": "project_floating_ip_pool_ref_uuid", "to": []string{"test", "project_floating_ip_pool_ref_uuid"}})
	FloatingIPPoolref = append(FloatingIPPoolref, map[string]interface{}{"operation": "add", "uuid": "project_floating_ip_pool_ref_uuid1", "to": []string{"test", "project_floating_ip_pool_ref_uuid1"}})

	common.SetValueByPath(updateMap, "FloatingIPPoolRefs", ".", FloatingIPPoolref)

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateProject(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_project_alias_ip_pool` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing AliasIPPoolRefs delete statement failed")
		}
		_, err = stmt.Exec("project_dummy_uuid", "project_alias_ip_pool_ref_uuid")
		_, err = stmt.Exec("project_dummy_uuid", "project_alias_ip_pool_ref_uuid1")
		_, err = stmt.Exec("project_dummy_uuid", "project_alias_ip_pool_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "AliasIPPoolRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteAliasIPPool(tx, "project_alias_ip_pool_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_alias_ip_pool_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteAliasIPPool(tx, "project_alias_ip_pool_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_alias_ip_pool_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteAliasIPPool(tx, "project_alias_ip_pool_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_alias_ip_pool_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_project_namespace` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing NamespaceRefs delete statement failed")
		}
		_, err = stmt.Exec("project_dummy_uuid", "project_namespace_ref_uuid")
		_, err = stmt.Exec("project_dummy_uuid", "project_namespace_ref_uuid1")
		_, err = stmt.Exec("project_dummy_uuid", "project_namespace_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "NamespaceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNamespace(tx, "project_namespace_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_namespace_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNamespace(tx, "project_namespace_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_namespace_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNamespace(tx, "project_namespace_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_namespace_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_project_application_policy_set` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ApplicationPolicySetRefs delete statement failed")
		}
		_, err = stmt.Exec("project_dummy_uuid", "project_application_policy_set_ref_uuid")
		_, err = stmt.Exec("project_dummy_uuid", "project_application_policy_set_ref_uuid1")
		_, err = stmt.Exec("project_dummy_uuid", "project_application_policy_set_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ApplicationPolicySetRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteApplicationPolicySet(tx, "project_application_policy_set_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_application_policy_set_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteApplicationPolicySet(tx, "project_application_policy_set_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_application_policy_set_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteApplicationPolicySet(tx, "project_application_policy_set_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_application_policy_set_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_project_floating_ip_pool` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing FloatingIPPoolRefs delete statement failed")
		}
		_, err = stmt.Exec("project_dummy_uuid", "project_floating_ip_pool_ref_uuid")
		_, err = stmt.Exec("project_dummy_uuid", "project_floating_ip_pool_ref_uuid1")
		_, err = stmt.Exec("project_dummy_uuid", "project_floating_ip_pool_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "FloatingIPPoolRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteFloatingIPPool(tx, "project_floating_ip_pool_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_floating_ip_pool_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteFloatingIPPool(tx, "project_floating_ip_pool_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_floating_ip_pool_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteFloatingIPPool(tx, "project_floating_ip_pool_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref project_floating_ip_pool_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, projectModel.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListProject(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(tx, model)
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListProject(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
