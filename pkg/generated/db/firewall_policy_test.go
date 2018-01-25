package db

import ("fmt"
        "testing"
        "database/sql"

        "github.com/Juniper/contrail/pkg/common"
        "github.com/Juniper/contrail/pkg/generated/models"
        "github.com/pkg/errors"
        )

func TestFirewallPolicy(t *testing.T) {
    t.Parallel()
    db := testDB
    common.UseTable(db, "metadata")
    common.UseTable(db, "firewall_policy")
    defer func(){
        common.ClearTable(db, "firewall_policy")
        common.ClearTable(db, "metadata")
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
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFirewallRule(tx, FirewallRulerefModel)
	})
    FirewallRulerefModel.UUID = "firewall_policy_firewall_rule_ref_uuid1"
    FirewallRulerefModel.FQName = []string{"test", "firewall_policy_firewall_rule_ref_uuid1"}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFirewallRule(tx, FirewallRulerefModel)
	})
    FirewallRulerefModel.UUID = "firewall_policy_firewall_rule_ref_uuid2"
    FirewallRulerefModel.FQName = []string{"test", "firewall_policy_firewall_rule_ref_uuid2"}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFirewallRule(tx, FirewallRulerefModel)
	})
    if err != nil {
        t.Fatal("ref create failed", err)
    }
    FirewallRulecreateref = append(FirewallRulecreateref, &models.FirewallPolicyFirewallRuleRef{UUID:"firewall_policy_firewall_rule_ref_uuid", To: []string{"test", "firewall_policy_firewall_rule_ref_uuid"}})
    FirewallRulecreateref = append(FirewallRulecreateref, &models.FirewallPolicyFirewallRuleRef{UUID:"firewall_policy_firewall_rule_ref_uuid2", To: []string{"test", "firewall_policy_firewall_rule_ref_uuid2"}})
    model.FirewallRuleRefs = FirewallRulecreateref
    
    var SecurityLoggingObjectcreateref []*models.FirewallPolicySecurityLoggingObjectRef
    var SecurityLoggingObjectrefModel *models.SecurityLoggingObject
    SecurityLoggingObjectrefModel = models.MakeSecurityLoggingObject()
	SecurityLoggingObjectrefModel.UUID = "firewall_policy_security_logging_object_ref_uuid"
    SecurityLoggingObjectrefModel.FQName = []string{"test", "firewall_policy_security_logging_object_ref_uuid"}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateSecurityLoggingObject(tx, SecurityLoggingObjectrefModel)
	})
    SecurityLoggingObjectrefModel.UUID = "firewall_policy_security_logging_object_ref_uuid1"
    SecurityLoggingObjectrefModel.FQName = []string{"test", "firewall_policy_security_logging_object_ref_uuid1"}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateSecurityLoggingObject(tx, SecurityLoggingObjectrefModel)
	})
    SecurityLoggingObjectrefModel.UUID = "firewall_policy_security_logging_object_ref_uuid2"
    SecurityLoggingObjectrefModel.FQName = []string{"test", "firewall_policy_security_logging_object_ref_uuid2"}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateSecurityLoggingObject(tx, SecurityLoggingObjectrefModel)
	})
    if err != nil {
        t.Fatal("ref create failed", err)
    }
    SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.FirewallPolicySecurityLoggingObjectRef{UUID:"firewall_policy_security_logging_object_ref_uuid", To: []string{"test", "firewall_policy_security_logging_object_ref_uuid"}})
    SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.FirewallPolicySecurityLoggingObjectRef{UUID:"firewall_policy_security_logging_object_ref_uuid2", To: []string{"test", "firewall_policy_security_logging_object_ref_uuid2"}})
    model.SecurityLoggingObjectRefs = SecurityLoggingObjectcreateref
    

    //create project to which resource is shared
    projectModel := models.MakeProject()
	projectModel.UUID = "firewall_policy_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
    var createShare []*models.ShareType
    createShare = append(createShare, &models.ShareType{Tenant:"default-domain-test:admin-test", TenantAccess:7})
    model.Perms2.Share = createShare
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(tx, projectModel)
	})
    if err != nil {
        t.Fatal("project create failed", err)
    }

    //populate update map
    updateMap := map[string]interface{}{}
    
    
    common.SetValueByPath(updateMap, ".UUID", ".", "test")
    
    
    
    if ".Perms2.Share" == ".Perms2.Share" {
        var share []interface{}
        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
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
        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
        common.SetValueByPath(updateMap, ".FQName", ".", share)
    } else {
        common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
    }
    
    
    
    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
    
    
    
    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
        var share []interface{}
        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
    } else {
        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
    }
    
    
    common.SetValueByPath(updateMap, "uuid", ".", "firewall_policy_dummy_uuid")
    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
    
    var FirewallRuleref []interface{}
    FirewallRuleref = append(FirewallRuleref, map[string]interface{}{"operation":"delete", "uuid":"firewall_policy_firewall_rule_ref_uuid", "to": []string{"test", "firewall_policy_firewall_rule_ref_uuid"}})
    FirewallRuleref = append(FirewallRuleref, map[string]interface{}{"operation":"add", "uuid":"firewall_policy_firewall_rule_ref_uuid1", "to": []string{"test", "firewall_policy_firewall_rule_ref_uuid1"}})
    
    FirewallRuleAttr := map[string]interface{}{}
    
    
    
    common.SetValueByPath(FirewallRuleAttr, ".Sequence", ".", "test")
    
    
    
    FirewallRuleref = append(FirewallRuleref, map[string]interface{}{"operation":"update", "uuid":"firewall_policy_firewall_rule_ref_uuid2", "to": []string{"test", "firewall_policy_firewall_rule_ref_uuid2"}, "attr": FirewallRuleAttr})
    
    common.SetValueByPath(updateMap, "FirewallRuleRefs", ".", FirewallRuleref)
    
    var SecurityLoggingObjectref []interface{}
    SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation":"delete", "uuid":"firewall_policy_security_logging_object_ref_uuid", "to": []string{"test", "firewall_policy_security_logging_object_ref_uuid"}})
    SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation":"add", "uuid":"firewall_policy_security_logging_object_ref_uuid1", "to": []string{"test", "firewall_policy_security_logging_object_ref_uuid1"}})
    
    
    
    common.SetValueByPath(updateMap, "SecurityLoggingObjectRefs", ".", SecurityLoggingObjectref)
    

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateFirewallPolicy(tx, model)
    })
    if err != nil {
        t.Fatal("create failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return UpdateFirewallPolicy(tx, model.UUID, updateMap)
    })
    if err != nil {
        t.Fatal("update failed", err)
    }

    //Delete ref entries, referred objects
    
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
        stmt, err := tx.Prepare("delete from `ref_firewall_policy_firewall_rule` where `from` = ? AND `to` = ?;")
        if err != nil {
            return errors.Wrap(err, "preparing FirewallRuleRefs delete statement failed")
        }
        _, err = stmt.Exec( "firewall_policy_dummy_uuid", "firewall_policy_firewall_rule_ref_uuid" )
        _, err = stmt.Exec( "firewall_policy_dummy_uuid", "firewall_policy_firewall_rule_ref_uuid1" )
        _, err = stmt.Exec( "firewall_policy_dummy_uuid", "firewall_policy_firewall_rule_ref_uuid2" )
        if err != nil {
            return errors.Wrap(err, "FirewallRuleRefs delete failed")
        }
        return nil
	})
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
    	return DeleteFirewallRule(tx, "firewall_policy_firewall_rule_ref_uuid", nil)
    })
	if err != nil {
		t.Fatal("delete ref firewall_policy_firewall_rule_ref_uuid  failed", err)
	}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
    	return DeleteFirewallRule(tx, "firewall_policy_firewall_rule_ref_uuid1", nil)
    })
	if err != nil {
		t.Fatal("delete ref firewall_policy_firewall_rule_ref_uuid1  failed", err)
	}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
    	return DeleteFirewallRule(tx, "firewall_policy_firewall_rule_ref_uuid2", nil)
    })
	if err != nil {
		t.Fatal("delete ref firewall_policy_firewall_rule_ref_uuid2 failed", err)
	}
    
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
        stmt, err := tx.Prepare("delete from `ref_firewall_policy_security_logging_object` where `from` = ? AND `to` = ?;")
        if err != nil {
            return errors.Wrap(err, "preparing SecurityLoggingObjectRefs delete statement failed")
        }
        _, err = stmt.Exec( "firewall_policy_dummy_uuid", "firewall_policy_security_logging_object_ref_uuid" )
        _, err = stmt.Exec( "firewall_policy_dummy_uuid", "firewall_policy_security_logging_object_ref_uuid1" )
        _, err = stmt.Exec( "firewall_policy_dummy_uuid", "firewall_policy_security_logging_object_ref_uuid2" )
        if err != nil {
            return errors.Wrap(err, "SecurityLoggingObjectRefs delete failed")
        }
        return nil
	})
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
    	return DeleteSecurityLoggingObject(tx, "firewall_policy_security_logging_object_ref_uuid", nil)
    })
	if err != nil {
		t.Fatal("delete ref firewall_policy_security_logging_object_ref_uuid  failed", err)
	}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
    	return DeleteSecurityLoggingObject(tx, "firewall_policy_security_logging_object_ref_uuid1", nil)
    })
	if err != nil {
		t.Fatal("delete ref firewall_policy_security_logging_object_ref_uuid1  failed", err)
	}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
    	return DeleteSecurityLoggingObject(tx, "firewall_policy_security_logging_object_ref_uuid2", nil)
    })
	if err != nil {
		t.Fatal("delete ref firewall_policy_security_logging_object_ref_uuid2 failed", err)
	}
    

    //Delete the project created for sharing
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, projectModel.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        models, err := ListFirewallPolicy(tx, &common.ListSpec{Limit: 1})
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

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteFirewallPolicy(tx, model.UUID, 
            common.NewAuthContext("default", "demo", "demo", []string{}), 
        )
    })
    if err == nil {
        t.Fatal("auth failed")
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteFirewallPolicy(tx, model.UUID, nil)
    })
    if err != nil {
        t.Fatal("delete failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateFirewallPolicy(tx, model)
    })
    if err == nil {
        t.Fatal("Raise Error On Duplicate Create failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        models, err := ListFirewallPolicy(tx, &common.ListSpec{Limit: 1})
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
