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

func TestLoadbalancerListener(t *testing.T) {
    // t.Parallel()
    db := testDB
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    mutexMetadata := common.UseTable(db, "metadata")
    mutexTable := common.UseTable(db, "loadbalancer_listener")
    mutexProject := common.UseTable(db, "loadbalancer_listener")
    defer func(){
        mutexTable.Unlock()
        mutexMetadata.Unlock()
        if p := recover(); p != nil {
			panic(p)
		}
    }()
    model := models.MakeLoadbalancerListener()
    model.UUID = "loadbalancer_listener_dummy_uuid"
    model.FQName = []string{"default", "default-domain", "loadbalancer_listener_dummy"}
    model.Perms2.Owner = "admin"
    var err error

    // Create referred objects
//    
//    var Loadbalancercreateref []*models.LoadbalancerListenerLoadbalancerRef
//    var LoadbalancerrefModel *models.Loadbalancer
//    LoadbalancerrefModel = models.MakeLoadbalancer()
//	LoadbalancerrefModel.UUID = "loadbalancer_listener_loadbalancer_ref_uuid"
//    LoadbalancerrefModel.FQName = []string{"test", "loadbalancer_listener_loadbalancer_ref_uuid"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateLoadbalancer(ctx, tx, &models.CreateLoadbalancerRequest{
//            Loadbalancer: LoadbalancerrefModel,
//        })
//	})
//    LoadbalancerrefModel.UUID = "loadbalancer_listener_loadbalancer_ref_uuid1"
//    LoadbalancerrefModel.FQName = []string{"test", "loadbalancer_listener_loadbalancer_ref_uuid1"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateLoadbalancer(ctx, tx, &models.CreateLoadbalancerRequest{
//            Loadbalancer: LoadbalancerrefModel,
//        })
//	})
//    LoadbalancerrefModel.UUID = "loadbalancer_listener_loadbalancer_ref_uuid2"
//    LoadbalancerrefModel.FQName = []string{"test", "loadbalancer_listener_loadbalancer_ref_uuid2"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateLoadbalancer(ctx, tx, &models.CreateLoadbalancerRequest{
//            Loadbalancer: LoadbalancerrefModel,
//        })
//	})
//    if err != nil {
//        t.Fatal("ref create failed", err)
//    }
//    Loadbalancercreateref = append(Loadbalancercreateref, &models.LoadbalancerListenerLoadbalancerRef{UUID:"loadbalancer_listener_loadbalancer_ref_uuid", To: []string{"test", "loadbalancer_listener_loadbalancer_ref_uuid"}})
//    Loadbalancercreateref = append(Loadbalancercreateref, &models.LoadbalancerListenerLoadbalancerRef{UUID:"loadbalancer_listener_loadbalancer_ref_uuid2", To: []string{"test", "loadbalancer_listener_loadbalancer_ref_uuid2"}})
//    model.LoadbalancerRefs = Loadbalancercreateref
//    
//
    //create project to which resource is shared
    projectModel := models.MakeProject()
	projectModel.UUID = "loadbalancer_listener_admin_project_uuid"
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
//    if ".LoadbalancerListenerProperties.SniContainers" == ".Perms2.Share" {
//        var share []interface{}
//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
//        common.SetValueByPath(updateMap, ".LoadbalancerListenerProperties.SniContainers", ".", share)
//    } else {
//        common.SetValueByPath(updateMap, ".LoadbalancerListenerProperties.SniContainers", ".", `{"test": "test"}`)
//    }
//    
//    
//    
//    common.SetValueByPath(updateMap, ".LoadbalancerListenerProperties.ProtocolPort", ".", 1.0)
//    
//    
//    
//    common.SetValueByPath(updateMap, ".LoadbalancerListenerProperties.Protocol", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".LoadbalancerListenerProperties.DefaultTLSContainer", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".LoadbalancerListenerProperties.ConnectionLimit", ".", 1.0)
//    
//    
//    
//    common.SetValueByPath(updateMap, ".LoadbalancerListenerProperties.AdminState", ".", true)
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
//    common.SetValueByPath(updateMap, "uuid", ".", "loadbalancer_listener_dummy_uuid")
//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
//
//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
//    
//    var Loadbalancerref []interface{}
//    Loadbalancerref = append(Loadbalancerref, map[string]interface{}{"operation":"delete", "uuid":"loadbalancer_listener_loadbalancer_ref_uuid", "to": []string{"test", "loadbalancer_listener_loadbalancer_ref_uuid"}})
//    Loadbalancerref = append(Loadbalancerref, map[string]interface{}{"operation":"add", "uuid":"loadbalancer_listener_loadbalancer_ref_uuid1", "to": []string{"test", "loadbalancer_listener_loadbalancer_ref_uuid1"}})
//    
//    
//    
//    common.SetValueByPath(updateMap, "LoadbalancerRefs", ".", Loadbalancerref)
//    
//
    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateLoadbalancerListener(ctx, tx, 
            &models.CreateLoadbalancerListenerRequest{
                LoadbalancerListener: model,
            })
    })
    if err != nil {
        t.Fatal("create failed", err)
    }

//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
//        return UpdateLoadbalancerListener(tx, model.UUID, updateMap)
//    })
//    if err != nil {
//        t.Fatal("update failed", err)
//    }

//    //Delete ref entries, referred objects
//    
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//        stmt, err := tx.Prepare("delete from `ref_loadbalancer_listener_loadbalancer` where `from` = ? AND `to` = ?;")
//        if err != nil {
//            return errors.Wrap(err, "preparing LoadbalancerRefs delete statement failed")
//        }
//        _, err = stmt.Exec( "loadbalancer_listener_dummy_uuid", "loadbalancer_listener_loadbalancer_ref_uuid" )
//        _, err = stmt.Exec( "loadbalancer_listener_dummy_uuid", "loadbalancer_listener_loadbalancer_ref_uuid1" )
//        _, err = stmt.Exec( "loadbalancer_listener_dummy_uuid", "loadbalancer_listener_loadbalancer_ref_uuid2" )
//        if err != nil {
//            return errors.Wrap(err, "LoadbalancerRefs delete failed")
//        }
//        return nil
//	})
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteLoadbalancer(ctx, tx, 
//            &models.DeleteLoadbalancerRequest{
//                ID: "loadbalancer_listener_loadbalancer_ref_uuid"})
//    })
//	if err != nil {
//		t.Fatal("delete ref loadbalancer_listener_loadbalancer_ref_uuid  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteLoadbalancer(ctx, tx, 
//            &models.DeleteLoadbalancerRequest{
//                ID: "loadbalancer_listener_loadbalancer_ref_uuid1"})
//    })
//	if err != nil {
//		t.Fatal("delete ref loadbalancer_listener_loadbalancer_ref_uuid1  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteLoadbalancer(
//            ctx, 
//            tx, 
//            &models.DeleteLoadbalancerRequest{
//                ID: "loadbalancer_listener_loadbalancer_ref_uuid2",
//            })
//    })
//	if err != nil {
//		t.Fatal("delete ref loadbalancer_listener_loadbalancer_ref_uuid2 failed", err)
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
        response, err := ListLoadbalancerListener(ctx, tx, &models.ListLoadbalancerListenerRequest{
            Spec: &models.ListSpec{Limit: 1}})
        if err != nil {
            return err
        }
        if len(response.LoadbalancerListeners) != 1 {
            return fmt.Errorf("expected one element")
        }
        return nil
    })
    if err != nil {
        t.Fatal("list failed", err)
    }
    
    ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteLoadbalancerListener(ctxDemo, tx, 
            &models.DeleteLoadbalancerListenerRequest{
                ID: model.UUID}, 
        )
    })
    if err == nil {
        t.Fatal("auth failed")
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteLoadbalancerListener(ctx, tx, 
            &models.DeleteLoadbalancerListenerRequest{
                ID: model.UUID})
    })
    if err != nil {
        t.Fatal("delete failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateLoadbalancerListener(ctx, tx, 
            &models.CreateLoadbalancerListenerRequest{
                LoadbalancerListener: model})
    })
    if err == nil {
        t.Fatal("Raise Error On Duplicate Create failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        response, err := ListLoadbalancerListener(ctx, tx, &models.ListLoadbalancerListenerRequest{
            Spec: &models.ListSpec{Limit: 1}})
        if err != nil {
            return err
        }
        if len(response.LoadbalancerListeners) != 0 {
            return fmt.Errorf("expected no element")
        }
        return nil
    })
    if err != nil {
        t.Fatal("list failed", err)
    }
    return
}
