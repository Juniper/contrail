package db

import (
        "context" 
        "fmt"
        "time"
        "testing"
        "database/sql"

        "github.com/Juniper/contrail/pkg/common"
        "github.com/Juniper/contrail/pkg/generated/models"
        "github.com/pkg/errors"
        )

//For skip import error.
var _ = errors.New("")

func TestServiceInstance(t *testing.T) {
    // t.Parallel()
    db := testDB
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    mutexMetadata := common.UseTable(db, "metadata")
    mutexTable := common.UseTable(db, "service_instance")
    mutexProject := common.UseTable(db, "service_instance")
    defer func(){
        mutexTable.Unlock()
        mutexMetadata.Unlock()
        if p := recover(); p != nil {
			panic(p)
		}
    }()
    model := models.MakeServiceInstance()
    model.UUID = "service_instance_dummy_uuid"
    model.FQName = []string{"default", "default-domain", "service_instance_dummy"}
    model.Perms2.Owner = "admin"
    var err error

    // Create referred objects
//    
//    var ServiceTemplatecreateref []*models.ServiceInstanceServiceTemplateRef
//    var ServiceTemplaterefModel *models.ServiceTemplate
//    ServiceTemplaterefModel = models.MakeServiceTemplate()
//	ServiceTemplaterefModel.UUID = "service_instance_service_template_ref_uuid"
//    ServiceTemplaterefModel.FQName = []string{"test", "service_instance_service_template_ref_uuid"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateServiceTemplate(ctx, tx, &models.CreateServiceTemplateRequest{
//            ServiceTemplate: ServiceTemplaterefModel,
//        })
//	})
//    ServiceTemplaterefModel.UUID = "service_instance_service_template_ref_uuid1"
//    ServiceTemplaterefModel.FQName = []string{"test", "service_instance_service_template_ref_uuid1"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateServiceTemplate(ctx, tx, &models.CreateServiceTemplateRequest{
//            ServiceTemplate: ServiceTemplaterefModel,
//        })
//	})
//    ServiceTemplaterefModel.UUID = "service_instance_service_template_ref_uuid2"
//    ServiceTemplaterefModel.FQName = []string{"test", "service_instance_service_template_ref_uuid2"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateServiceTemplate(ctx, tx, &models.CreateServiceTemplateRequest{
//            ServiceTemplate: ServiceTemplaterefModel,
//        })
//	})
//    if err != nil {
//        t.Fatal("ref create failed", err)
//    }
//    ServiceTemplatecreateref = append(ServiceTemplatecreateref, &models.ServiceInstanceServiceTemplateRef{UUID:"service_instance_service_template_ref_uuid", To: []string{"test", "service_instance_service_template_ref_uuid"}})
//    ServiceTemplatecreateref = append(ServiceTemplatecreateref, &models.ServiceInstanceServiceTemplateRef{UUID:"service_instance_service_template_ref_uuid2", To: []string{"test", "service_instance_service_template_ref_uuid2"}})
//    model.ServiceTemplateRefs = ServiceTemplatecreateref
//    
//    var InstanceIPcreateref []*models.ServiceInstanceInstanceIPRef
//    var InstanceIPrefModel *models.InstanceIP
//    InstanceIPrefModel = models.MakeInstanceIP()
//	InstanceIPrefModel.UUID = "service_instance_instance_ip_ref_uuid"
//    InstanceIPrefModel.FQName = []string{"test", "service_instance_instance_ip_ref_uuid"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateInstanceIP(ctx, tx, &models.CreateInstanceIPRequest{
//            InstanceIP: InstanceIPrefModel,
//        })
//	})
//    InstanceIPrefModel.UUID = "service_instance_instance_ip_ref_uuid1"
//    InstanceIPrefModel.FQName = []string{"test", "service_instance_instance_ip_ref_uuid1"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateInstanceIP(ctx, tx, &models.CreateInstanceIPRequest{
//            InstanceIP: InstanceIPrefModel,
//        })
//	})
//    InstanceIPrefModel.UUID = "service_instance_instance_ip_ref_uuid2"
//    InstanceIPrefModel.FQName = []string{"test", "service_instance_instance_ip_ref_uuid2"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateInstanceIP(ctx, tx, &models.CreateInstanceIPRequest{
//            InstanceIP: InstanceIPrefModel,
//        })
//	})
//    if err != nil {
//        t.Fatal("ref create failed", err)
//    }
//    InstanceIPcreateref = append(InstanceIPcreateref, &models.ServiceInstanceInstanceIPRef{UUID:"service_instance_instance_ip_ref_uuid", To: []string{"test", "service_instance_instance_ip_ref_uuid"}})
//    InstanceIPcreateref = append(InstanceIPcreateref, &models.ServiceInstanceInstanceIPRef{UUID:"service_instance_instance_ip_ref_uuid2", To: []string{"test", "service_instance_instance_ip_ref_uuid2"}})
//    model.InstanceIPRefs = InstanceIPcreateref
//    
//
    //create project to which resource is shared
    projectModel := models.MakeProject()
	projectModel.UUID = "service_instance_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
    var createShare []*models.ShareType
    createShare = append(createShare, &models.ShareType{Tenant:"default-domain-test:admin-test", TenantAccess:7})
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
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.VirtualRouterID", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.ScaleOut.MaxInstances", ".", 1.0)
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.ScaleOut.AutoScale", ".", true)
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.RightVirtualNetwork", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.RightIPAddress", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.ManagementVirtualNetwork", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.LeftVirtualNetwork", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.LeftIPAddress", ".", "test")
//    
//    
//    
//    if ".ServiceInstanceProperties.InterfaceList" == ".Perms2.Share" {
//        var share []interface{}
//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
//        common.SetValueByPath(updateMap, ".ServiceInstanceProperties.InterfaceList", ".", share)
//    } else {
//        common.SetValueByPath(updateMap, ".ServiceInstanceProperties.InterfaceList", ".", `{"test": "test"}`)
//    }
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.HaMode", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.AvailabilityZone", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".ServiceInstanceProperties.AutoPolicy", ".", true)
//    
//    
//    
//    if ".ServiceInstanceBindings.KeyValuePair" == ".Perms2.Share" {
//        var share []interface{}
//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
//        common.SetValueByPath(updateMap, ".ServiceInstanceBindings.KeyValuePair", ".", share)
//    } else {
//        common.SetValueByPath(updateMap, ".ServiceInstanceBindings.KeyValuePair", ".", `{"test": "test"}`)
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
//    common.SetValueByPath(updateMap, "uuid", ".", "service_instance_dummy_uuid")
//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
//
//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
//    
//    var ServiceTemplateref []interface{}
//    ServiceTemplateref = append(ServiceTemplateref, map[string]interface{}{"operation":"delete", "uuid":"service_instance_service_template_ref_uuid", "to": []string{"test", "service_instance_service_template_ref_uuid"}})
//    ServiceTemplateref = append(ServiceTemplateref, map[string]interface{}{"operation":"add", "uuid":"service_instance_service_template_ref_uuid1", "to": []string{"test", "service_instance_service_template_ref_uuid1"}})
//    
//    
//    
//    common.SetValueByPath(updateMap, "ServiceTemplateRefs", ".", ServiceTemplateref)
//    
//    var InstanceIPref []interface{}
//    InstanceIPref = append(InstanceIPref, map[string]interface{}{"operation":"delete", "uuid":"service_instance_instance_ip_ref_uuid", "to": []string{"test", "service_instance_instance_ip_ref_uuid"}})
//    InstanceIPref = append(InstanceIPref, map[string]interface{}{"operation":"add", "uuid":"service_instance_instance_ip_ref_uuid1", "to": []string{"test", "service_instance_instance_ip_ref_uuid1"}})
//    
//    InstanceIPAttr := map[string]interface{}{}
//    
//    
//    
//    common.SetValueByPath(InstanceIPAttr, ".InterfaceType", ".", "test")
//    
//    
//    
//    InstanceIPref = append(InstanceIPref, map[string]interface{}{"operation":"update", "uuid":"service_instance_instance_ip_ref_uuid2", "to": []string{"test", "service_instance_instance_ip_ref_uuid2"}, "attr": InstanceIPAttr})
//    
//    common.SetValueByPath(updateMap, "InstanceIPRefs", ".", InstanceIPref)
//    
//
    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateServiceInstance(ctx, tx, 
            &models.CreateServiceInstanceRequest{
                ServiceInstance: model,
            })
    })
    if err != nil {
        t.Fatal("create failed", err)
    }

//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
//        return UpdateServiceInstance(tx, model.UUID, updateMap)
//    })
//    if err != nil {
//        t.Fatal("update failed", err)
//    }

//    //Delete ref entries, referred objects
//    
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//        stmt, err := tx.Prepare("delete from `ref_service_instance_service_template` where `from` = ? AND `to` = ?;")
//        if err != nil {
//            return errors.Wrap(err, "preparing ServiceTemplateRefs delete statement failed")
//        }
//        _, err = stmt.Exec( "service_instance_dummy_uuid", "service_instance_service_template_ref_uuid" )
//        _, err = stmt.Exec( "service_instance_dummy_uuid", "service_instance_service_template_ref_uuid1" )
//        _, err = stmt.Exec( "service_instance_dummy_uuid", "service_instance_service_template_ref_uuid2" )
//        if err != nil {
//            return errors.Wrap(err, "ServiceTemplateRefs delete failed")
//        }
//        return nil
//	})
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteServiceTemplate(ctx, tx, 
//            &models.DeleteServiceTemplateRequest{
//                ID: "service_instance_service_template_ref_uuid"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_instance_service_template_ref_uuid  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteServiceTemplate(ctx, tx, 
//            &models.DeleteServiceTemplateRequest{
//                ID: "service_instance_service_template_ref_uuid1"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_instance_service_template_ref_uuid1  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteServiceTemplate(
//            ctx, 
//            tx, 
//            &models.DeleteServiceTemplateRequest{
//                ID: "service_instance_service_template_ref_uuid2",
//            })
//    })
//	if err != nil {
//		t.Fatal("delete ref service_instance_service_template_ref_uuid2 failed", err)
//	}
//    
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//        stmt, err := tx.Prepare("delete from `ref_service_instance_instance_ip` where `from` = ? AND `to` = ?;")
//        if err != nil {
//            return errors.Wrap(err, "preparing InstanceIPRefs delete statement failed")
//        }
//        _, err = stmt.Exec( "service_instance_dummy_uuid", "service_instance_instance_ip_ref_uuid" )
//        _, err = stmt.Exec( "service_instance_dummy_uuid", "service_instance_instance_ip_ref_uuid1" )
//        _, err = stmt.Exec( "service_instance_dummy_uuid", "service_instance_instance_ip_ref_uuid2" )
//        if err != nil {
//            return errors.Wrap(err, "InstanceIPRefs delete failed")
//        }
//        return nil
//	})
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteInstanceIP(ctx, tx, 
//            &models.DeleteInstanceIPRequest{
//                ID: "service_instance_instance_ip_ref_uuid"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_instance_instance_ip_ref_uuid  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteInstanceIP(ctx, tx, 
//            &models.DeleteInstanceIPRequest{
//                ID: "service_instance_instance_ip_ref_uuid1"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_instance_instance_ip_ref_uuid1  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteInstanceIP(
//            ctx, 
//            tx, 
//            &models.DeleteInstanceIPRequest{
//                ID: "service_instance_instance_ip_ref_uuid2",
//            })
//    })
//	if err != nil {
//		t.Fatal("delete ref service_instance_instance_ip_ref_uuid2 failed", err)
//	}
//    

    //Delete the project created for sharing
    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(ctx, tx, &models.DeleteProjectRequest{
            ID: projectModel.UUID})
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        response, err := ListServiceInstance(ctx, tx, &models.ListServiceInstanceRequest{
            Spec: &models.ListSpec{Limit: 1}})
        if err != nil {
            return err
        }
        if len(response.ServiceInstances) != 1 {
            return fmt.Errorf("expected one element")
        }
        return nil
    })
    if err != nil {
        t.Fatal("list failed", err)
    }
    
    ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteServiceInstance(ctxDemo, tx, 
            &models.DeleteServiceInstanceRequest{
                ID: model.UUID}, 
        )
    })
    if err == nil {
        t.Fatal("auth failed")
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteServiceInstance(ctx, tx, 
            &models.DeleteServiceInstanceRequest{
                ID: model.UUID})
    })
    if err != nil {
        t.Fatal("delete failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateServiceInstance(ctx, tx, 
            &models.CreateServiceInstanceRequest{
                ServiceInstance: model})
    })
    if err == nil {
        t.Fatal("Raise Error On Duplicate Create failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        response, err := ListServiceInstance(ctx, tx, &models.ListServiceInstanceRequest{
            Spec: &models.ListSpec{Limit: 1}})
        if err != nil {
            return err
        }
        if len(response.ServiceInstances) != 0 {
            return fmt.Errorf("expected no element")
        }
        return nil
    })
    if err != nil {
        t.Fatal("list failed", err)
    }
    return
}
