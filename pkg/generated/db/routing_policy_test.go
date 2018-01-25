package db

import ("fmt"
        "testing"
        "database/sql"

        "github.com/Juniper/contrail/pkg/common"
        "github.com/Juniper/contrail/pkg/generated/models"
        "github.com/pkg/errors"
        )

func TestRoutingPolicy(t *testing.T) {
    t.Parallel()
    db := testDB
    common.UseTable(db, "metadata")
    common.UseTable(db, "routing_policy")
    defer func(){
        common.ClearTable(db, "routing_policy")
        common.ClearTable(db, "metadata")
        if p := recover(); p != nil {
			panic(p)
		}
    }()
    model := models.MakeRoutingPolicy()
    model.UUID = "routing_policy_dummy_uuid"
    model.FQName = []string{"default", "default-domain", "routing_policy_dummy"}
    model.Perms2.Owner = "admin"
    var err error

    // Create referred objects
    
    var ServiceInstancecreateref []*models.RoutingPolicyServiceInstanceRef
    var ServiceInstancerefModel *models.ServiceInstance
    ServiceInstancerefModel = models.MakeServiceInstance()
	ServiceInstancerefModel.UUID = "routing_policy_service_instance_ref_uuid"
    ServiceInstancerefModel.FQName = []string{"test", "routing_policy_service_instance_ref_uuid"}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(tx, ServiceInstancerefModel)
	})
    ServiceInstancerefModel.UUID = "routing_policy_service_instance_ref_uuid1"
    ServiceInstancerefModel.FQName = []string{"test", "routing_policy_service_instance_ref_uuid1"}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(tx, ServiceInstancerefModel)
	})
    ServiceInstancerefModel.UUID = "routing_policy_service_instance_ref_uuid2"
    ServiceInstancerefModel.FQName = []string{"test", "routing_policy_service_instance_ref_uuid2"}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(tx, ServiceInstancerefModel)
	})
    if err != nil {
        t.Fatal("ref create failed", err)
    }
    ServiceInstancecreateref = append(ServiceInstancecreateref, &models.RoutingPolicyServiceInstanceRef{UUID:"routing_policy_service_instance_ref_uuid", To: []string{"test", "routing_policy_service_instance_ref_uuid"}})
    ServiceInstancecreateref = append(ServiceInstancecreateref, &models.RoutingPolicyServiceInstanceRef{UUID:"routing_policy_service_instance_ref_uuid2", To: []string{"test", "routing_policy_service_instance_ref_uuid2"}})
    model.ServiceInstanceRefs = ServiceInstancecreateref
    

    //create project to which resource is shared
    projectModel := models.MakeProject()
	projectModel.UUID = "routing_policy_admin_project_uuid"
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
    
    
    common.SetValueByPath(updateMap, "uuid", ".", "routing_policy_dummy_uuid")
    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
    
    var ServiceInstanceref []interface{}
    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"delete", "uuid":"routing_policy_service_instance_ref_uuid", "to": []string{"test", "routing_policy_service_instance_ref_uuid"}})
    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"add", "uuid":"routing_policy_service_instance_ref_uuid1", "to": []string{"test", "routing_policy_service_instance_ref_uuid1"}})
    
    ServiceInstanceAttr := map[string]interface{}{}
    
    
    
    common.SetValueByPath(ServiceInstanceAttr, ".RightSequence", ".", "test")
    
    
    
    common.SetValueByPath(ServiceInstanceAttr, ".LeftSequence", ".", "test")
    
    
    
    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"update", "uuid":"routing_policy_service_instance_ref_uuid2", "to": []string{"test", "routing_policy_service_instance_ref_uuid2"}, "attr": ServiceInstanceAttr})
    
    common.SetValueByPath(updateMap, "ServiceInstanceRefs", ".", ServiceInstanceref)
    

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateRoutingPolicy(tx, model)
    })
    if err != nil {
        t.Fatal("create failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return UpdateRoutingPolicy(tx, model.UUID, updateMap)
    })
    if err != nil {
        t.Fatal("update failed", err)
    }

    //Delete ref entries, referred objects
    
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
        stmt, err := tx.Prepare("delete from `ref_routing_policy_service_instance` where `from` = ? AND `to` = ?;")
        if err != nil {
            return errors.Wrap(err, "preparing ServiceInstanceRefs delete statement failed")
        }
        _, err = stmt.Exec( "routing_policy_dummy_uuid", "routing_policy_service_instance_ref_uuid" )
        _, err = stmt.Exec( "routing_policy_dummy_uuid", "routing_policy_service_instance_ref_uuid1" )
        _, err = stmt.Exec( "routing_policy_dummy_uuid", "routing_policy_service_instance_ref_uuid2" )
        if err != nil {
            return errors.Wrap(err, "ServiceInstanceRefs delete failed")
        }
        return nil
	})
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
    	return DeleteServiceInstance(tx, "routing_policy_service_instance_ref_uuid", nil)
    })
	if err != nil {
		t.Fatal("delete ref routing_policy_service_instance_ref_uuid  failed", err)
	}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
    	return DeleteServiceInstance(tx, "routing_policy_service_instance_ref_uuid1", nil)
    })
	if err != nil {
		t.Fatal("delete ref routing_policy_service_instance_ref_uuid1  failed", err)
	}
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
    	return DeleteServiceInstance(tx, "routing_policy_service_instance_ref_uuid2", nil)
    })
	if err != nil {
		t.Fatal("delete ref routing_policy_service_instance_ref_uuid2 failed", err)
	}
    

    //Delete the project created for sharing
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, projectModel.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        models, err := ListRoutingPolicy(tx, &common.ListSpec{Limit: 1})
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
        return DeleteRoutingPolicy(tx, model.UUID, 
            common.NewAuthContext("default", "demo", "demo", []string{}), 
        )
    })
    if err == nil {
        t.Fatal("auth failed")
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteRoutingPolicy(tx, model.UUID, nil)
    })
    if err != nil {
        t.Fatal("delete failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateRoutingPolicy(tx, model)
    })
    if err == nil {
        t.Fatal("Raise Error On Duplicate Create failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        models, err := ListRoutingPolicy(tx, &common.ListSpec{Limit: 1})
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
