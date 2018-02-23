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

func TestServiceEndpoint(t *testing.T) {
    // t.Parallel()
    db := testDB
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    mutexMetadata := common.UseTable(db, "metadata")
    mutexTable := common.UseTable(db, "service_endpoint")
    mutexProject := common.UseTable(db, "service_endpoint")
    defer func(){
        mutexTable.Unlock()
        mutexMetadata.Unlock()
        if p := recover(); p != nil {
			panic(p)
		}
    }()
    model := models.MakeServiceEndpoint()
    model.UUID = "service_endpoint_dummy_uuid"
    model.FQName = []string{"default", "default-domain", "service_endpoint_dummy"}
    model.Perms2.Owner = "admin"
    var err error

    // Create referred objects
//    
//    var ServiceConnectionModulecreateref []*models.ServiceEndpointServiceConnectionModuleRef
//    var ServiceConnectionModulerefModel *models.ServiceConnectionModule
//    ServiceConnectionModulerefModel = models.MakeServiceConnectionModule()
//	ServiceConnectionModulerefModel.UUID = "service_endpoint_service_connection_module_ref_uuid"
//    ServiceConnectionModulerefModel.FQName = []string{"test", "service_endpoint_service_connection_module_ref_uuid"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateServiceConnectionModule(ctx, tx, &models.CreateServiceConnectionModuleRequest{
//            ServiceConnectionModule: ServiceConnectionModulerefModel,
//        })
//	})
//    ServiceConnectionModulerefModel.UUID = "service_endpoint_service_connection_module_ref_uuid1"
//    ServiceConnectionModulerefModel.FQName = []string{"test", "service_endpoint_service_connection_module_ref_uuid1"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateServiceConnectionModule(ctx, tx, &models.CreateServiceConnectionModuleRequest{
//            ServiceConnectionModule: ServiceConnectionModulerefModel,
//        })
//	})
//    ServiceConnectionModulerefModel.UUID = "service_endpoint_service_connection_module_ref_uuid2"
//    ServiceConnectionModulerefModel.FQName = []string{"test", "service_endpoint_service_connection_module_ref_uuid2"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateServiceConnectionModule(ctx, tx, &models.CreateServiceConnectionModuleRequest{
//            ServiceConnectionModule: ServiceConnectionModulerefModel,
//        })
//	})
//    if err != nil {
//        t.Fatal("ref create failed", err)
//    }
//    ServiceConnectionModulecreateref = append(ServiceConnectionModulecreateref, &models.ServiceEndpointServiceConnectionModuleRef{UUID:"service_endpoint_service_connection_module_ref_uuid", To: []string{"test", "service_endpoint_service_connection_module_ref_uuid"}})
//    ServiceConnectionModulecreateref = append(ServiceConnectionModulecreateref, &models.ServiceEndpointServiceConnectionModuleRef{UUID:"service_endpoint_service_connection_module_ref_uuid2", To: []string{"test", "service_endpoint_service_connection_module_ref_uuid2"}})
//    model.ServiceConnectionModuleRefs = ServiceConnectionModulecreateref
//    
//    var PhysicalRoutercreateref []*models.ServiceEndpointPhysicalRouterRef
//    var PhysicalRouterrefModel *models.PhysicalRouter
//    PhysicalRouterrefModel = models.MakePhysicalRouter()
//	PhysicalRouterrefModel.UUID = "service_endpoint_physical_router_ref_uuid"
//    PhysicalRouterrefModel.FQName = []string{"test", "service_endpoint_physical_router_ref_uuid"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
//            PhysicalRouter: PhysicalRouterrefModel,
//        })
//	})
//    PhysicalRouterrefModel.UUID = "service_endpoint_physical_router_ref_uuid1"
//    PhysicalRouterrefModel.FQName = []string{"test", "service_endpoint_physical_router_ref_uuid1"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
//            PhysicalRouter: PhysicalRouterrefModel,
//        })
//	})
//    PhysicalRouterrefModel.UUID = "service_endpoint_physical_router_ref_uuid2"
//    PhysicalRouterrefModel.FQName = []string{"test", "service_endpoint_physical_router_ref_uuid2"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreatePhysicalRouter(ctx, tx, &models.CreatePhysicalRouterRequest{
//            PhysicalRouter: PhysicalRouterrefModel,
//        })
//	})
//    if err != nil {
//        t.Fatal("ref create failed", err)
//    }
//    PhysicalRoutercreateref = append(PhysicalRoutercreateref, &models.ServiceEndpointPhysicalRouterRef{UUID:"service_endpoint_physical_router_ref_uuid", To: []string{"test", "service_endpoint_physical_router_ref_uuid"}})
//    PhysicalRoutercreateref = append(PhysicalRoutercreateref, &models.ServiceEndpointPhysicalRouterRef{UUID:"service_endpoint_physical_router_ref_uuid2", To: []string{"test", "service_endpoint_physical_router_ref_uuid2"}})
//    model.PhysicalRouterRefs = PhysicalRoutercreateref
//    
//    var ServiceObjectcreateref []*models.ServiceEndpointServiceObjectRef
//    var ServiceObjectrefModel *models.ServiceObject
//    ServiceObjectrefModel = models.MakeServiceObject()
//	ServiceObjectrefModel.UUID = "service_endpoint_service_object_ref_uuid"
//    ServiceObjectrefModel.FQName = []string{"test", "service_endpoint_service_object_ref_uuid"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateServiceObject(ctx, tx, &models.CreateServiceObjectRequest{
//            ServiceObject: ServiceObjectrefModel,
//        })
//	})
//    ServiceObjectrefModel.UUID = "service_endpoint_service_object_ref_uuid1"
//    ServiceObjectrefModel.FQName = []string{"test", "service_endpoint_service_object_ref_uuid1"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateServiceObject(ctx, tx, &models.CreateServiceObjectRequest{
//            ServiceObject: ServiceObjectrefModel,
//        })
//	})
//    ServiceObjectrefModel.UUID = "service_endpoint_service_object_ref_uuid2"
//    ServiceObjectrefModel.FQName = []string{"test", "service_endpoint_service_object_ref_uuid2"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateServiceObject(ctx, tx, &models.CreateServiceObjectRequest{
//            ServiceObject: ServiceObjectrefModel,
//        })
//	})
//    if err != nil {
//        t.Fatal("ref create failed", err)
//    }
//    ServiceObjectcreateref = append(ServiceObjectcreateref, &models.ServiceEndpointServiceObjectRef{UUID:"service_endpoint_service_object_ref_uuid", To: []string{"test", "service_endpoint_service_object_ref_uuid"}})
//    ServiceObjectcreateref = append(ServiceObjectcreateref, &models.ServiceEndpointServiceObjectRef{UUID:"service_endpoint_service_object_ref_uuid2", To: []string{"test", "service_endpoint_service_object_ref_uuid2"}})
//    model.ServiceObjectRefs = ServiceObjectcreateref
//    
//
    //create project to which resource is shared
    projectModel := models.MakeProject()
	projectModel.UUID = "service_endpoint_admin_project_uuid"
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
//    common.SetValueByPath(updateMap, "uuid", ".", "service_endpoint_dummy_uuid")
//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
//
//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
//    
//    var ServiceConnectionModuleref []interface{}
//    ServiceConnectionModuleref = append(ServiceConnectionModuleref, map[string]interface{}{"operation":"delete", "uuid":"service_endpoint_service_connection_module_ref_uuid", "to": []string{"test", "service_endpoint_service_connection_module_ref_uuid"}})
//    ServiceConnectionModuleref = append(ServiceConnectionModuleref, map[string]interface{}{"operation":"add", "uuid":"service_endpoint_service_connection_module_ref_uuid1", "to": []string{"test", "service_endpoint_service_connection_module_ref_uuid1"}})
//    
//    
//    
//    common.SetValueByPath(updateMap, "ServiceConnectionModuleRefs", ".", ServiceConnectionModuleref)
//    
//    var PhysicalRouterref []interface{}
//    PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation":"delete", "uuid":"service_endpoint_physical_router_ref_uuid", "to": []string{"test", "service_endpoint_physical_router_ref_uuid"}})
//    PhysicalRouterref = append(PhysicalRouterref, map[string]interface{}{"operation":"add", "uuid":"service_endpoint_physical_router_ref_uuid1", "to": []string{"test", "service_endpoint_physical_router_ref_uuid1"}})
//    
//    
//    
//    common.SetValueByPath(updateMap, "PhysicalRouterRefs", ".", PhysicalRouterref)
//    
//    var ServiceObjectref []interface{}
//    ServiceObjectref = append(ServiceObjectref, map[string]interface{}{"operation":"delete", "uuid":"service_endpoint_service_object_ref_uuid", "to": []string{"test", "service_endpoint_service_object_ref_uuid"}})
//    ServiceObjectref = append(ServiceObjectref, map[string]interface{}{"operation":"add", "uuid":"service_endpoint_service_object_ref_uuid1", "to": []string{"test", "service_endpoint_service_object_ref_uuid1"}})
//    
//    
//    
//    common.SetValueByPath(updateMap, "ServiceObjectRefs", ".", ServiceObjectref)
//    
//
    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateServiceEndpoint(ctx, tx, 
            &models.CreateServiceEndpointRequest{
                ServiceEndpoint: model,
            })
    })
    if err != nil {
        t.Fatal("create failed", err)
    }

//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
//        return UpdateServiceEndpoint(tx, model.UUID, updateMap)
//    })
//    if err != nil {
//        t.Fatal("update failed", err)
//    }

//    //Delete ref entries, referred objects
//    
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//        stmt, err := tx.Prepare("delete from `ref_service_endpoint_service_object` where `from` = ? AND `to` = ?;")
//        if err != nil {
//            return errors.Wrap(err, "preparing ServiceObjectRefs delete statement failed")
//        }
//        _, err = stmt.Exec( "service_endpoint_dummy_uuid", "service_endpoint_service_object_ref_uuid" )
//        _, err = stmt.Exec( "service_endpoint_dummy_uuid", "service_endpoint_service_object_ref_uuid1" )
//        _, err = stmt.Exec( "service_endpoint_dummy_uuid", "service_endpoint_service_object_ref_uuid2" )
//        if err != nil {
//            return errors.Wrap(err, "ServiceObjectRefs delete failed")
//        }
//        return nil
//	})
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteServiceObject(ctx, tx, 
//            &models.DeleteServiceObjectRequest{
//                ID: "service_endpoint_service_object_ref_uuid"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_endpoint_service_object_ref_uuid  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteServiceObject(ctx, tx, 
//            &models.DeleteServiceObjectRequest{
//                ID: "service_endpoint_service_object_ref_uuid1"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_endpoint_service_object_ref_uuid1  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteServiceObject(
//            ctx, 
//            tx, 
//            &models.DeleteServiceObjectRequest{
//                ID: "service_endpoint_service_object_ref_uuid2",
//            })
//    })
//	if err != nil {
//		t.Fatal("delete ref service_endpoint_service_object_ref_uuid2 failed", err)
//	}
//    
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//        stmt, err := tx.Prepare("delete from `ref_service_endpoint_service_connection_module` where `from` = ? AND `to` = ?;")
//        if err != nil {
//            return errors.Wrap(err, "preparing ServiceConnectionModuleRefs delete statement failed")
//        }
//        _, err = stmt.Exec( "service_endpoint_dummy_uuid", "service_endpoint_service_connection_module_ref_uuid" )
//        _, err = stmt.Exec( "service_endpoint_dummy_uuid", "service_endpoint_service_connection_module_ref_uuid1" )
//        _, err = stmt.Exec( "service_endpoint_dummy_uuid", "service_endpoint_service_connection_module_ref_uuid2" )
//        if err != nil {
//            return errors.Wrap(err, "ServiceConnectionModuleRefs delete failed")
//        }
//        return nil
//	})
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteServiceConnectionModule(ctx, tx, 
//            &models.DeleteServiceConnectionModuleRequest{
//                ID: "service_endpoint_service_connection_module_ref_uuid"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_endpoint_service_connection_module_ref_uuid  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteServiceConnectionModule(ctx, tx, 
//            &models.DeleteServiceConnectionModuleRequest{
//                ID: "service_endpoint_service_connection_module_ref_uuid1"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_endpoint_service_connection_module_ref_uuid1  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteServiceConnectionModule(
//            ctx, 
//            tx, 
//            &models.DeleteServiceConnectionModuleRequest{
//                ID: "service_endpoint_service_connection_module_ref_uuid2",
//            })
//    })
//	if err != nil {
//		t.Fatal("delete ref service_endpoint_service_connection_module_ref_uuid2 failed", err)
//	}
//    
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//        stmt, err := tx.Prepare("delete from `ref_service_endpoint_physical_router` where `from` = ? AND `to` = ?;")
//        if err != nil {
//            return errors.Wrap(err, "preparing PhysicalRouterRefs delete statement failed")
//        }
//        _, err = stmt.Exec( "service_endpoint_dummy_uuid", "service_endpoint_physical_router_ref_uuid" )
//        _, err = stmt.Exec( "service_endpoint_dummy_uuid", "service_endpoint_physical_router_ref_uuid1" )
//        _, err = stmt.Exec( "service_endpoint_dummy_uuid", "service_endpoint_physical_router_ref_uuid2" )
//        if err != nil {
//            return errors.Wrap(err, "PhysicalRouterRefs delete failed")
//        }
//        return nil
//	})
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeletePhysicalRouter(ctx, tx, 
//            &models.DeletePhysicalRouterRequest{
//                ID: "service_endpoint_physical_router_ref_uuid"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_endpoint_physical_router_ref_uuid  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeletePhysicalRouter(ctx, tx, 
//            &models.DeletePhysicalRouterRequest{
//                ID: "service_endpoint_physical_router_ref_uuid1"})
//    })
//	if err != nil {
//		t.Fatal("delete ref service_endpoint_physical_router_ref_uuid1  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeletePhysicalRouter(
//            ctx, 
//            tx, 
//            &models.DeletePhysicalRouterRequest{
//                ID: "service_endpoint_physical_router_ref_uuid2",
//            })
//    })
//	if err != nil {
//		t.Fatal("delete ref service_endpoint_physical_router_ref_uuid2 failed", err)
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
        response, err := ListServiceEndpoint(ctx, tx, &models.ListServiceEndpointRequest{
            Spec: &models.ListSpec{Limit: 1}})
        if err != nil {
            return err
        }
        if len(response.ServiceEndpoints) != 1 {
            return fmt.Errorf("expected one element")
        }
        return nil
    })
    if err != nil {
        t.Fatal("list failed", err)
    }
    
    ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteServiceEndpoint(ctxDemo, tx, 
            &models.DeleteServiceEndpointRequest{
                ID: model.UUID}, 
        )
    })
    if err == nil {
        t.Fatal("auth failed")
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteServiceEndpoint(ctx, tx, 
            &models.DeleteServiceEndpointRequest{
                ID: model.UUID})
    })
    if err != nil {
        t.Fatal("delete failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateServiceEndpoint(ctx, tx, 
            &models.CreateServiceEndpointRequest{
                ServiceEndpoint: model})
    })
    if err == nil {
        t.Fatal("Raise Error On Duplicate Create failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        response, err := ListServiceEndpoint(ctx, tx, &models.ListServiceEndpointRequest{
            Spec: &models.ListSpec{Limit: 1}})
        if err != nil {
            return err
        }
        if len(response.ServiceEndpoints) != 0 {
            return fmt.Errorf("expected no element")
        }
        return nil
    })
    if err != nil {
        t.Fatal("list failed", err)
    }
    return
}
