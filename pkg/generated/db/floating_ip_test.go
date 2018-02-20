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

func TestFloatingIP(t *testing.T) {
    // t.Parallel()
    db := testDB
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    mutexMetadata := common.UseTable(db, "metadata")
    mutexTable := common.UseTable(db, "floating_ip")
    mutexProject := common.UseTable(db, "floating_ip")
    defer func(){
        mutexTable.Unlock()
        mutexMetadata.Unlock()
        if p := recover(); p != nil {
			panic(p)
		}
    }()
    model := models.MakeFloatingIP()
    model.UUID = "floating_ip_dummy_uuid"
    model.FQName = []string{"default", "default-domain", "floating_ip_dummy"}
    model.Perms2.Owner = "admin"
    var err error

    // Create referred objects
//    
//    var Projectcreateref []*models.FloatingIPProjectRef
//    var ProjectrefModel *models.Project
//    ProjectrefModel = models.MakeProject()
//	ProjectrefModel.UUID = "floating_ip_project_ref_uuid"
//    ProjectrefModel.FQName = []string{"test", "floating_ip_project_ref_uuid"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateProject(ctx, tx, &models.CreateProjectRequest{
//            Project: ProjectrefModel,
//        })
//	})
//    ProjectrefModel.UUID = "floating_ip_project_ref_uuid1"
//    ProjectrefModel.FQName = []string{"test", "floating_ip_project_ref_uuid1"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateProject(ctx, tx, &models.CreateProjectRequest{
//            Project: ProjectrefModel,
//        })
//	})
//    ProjectrefModel.UUID = "floating_ip_project_ref_uuid2"
//    ProjectrefModel.FQName = []string{"test", "floating_ip_project_ref_uuid2"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateProject(ctx, tx, &models.CreateProjectRequest{
//            Project: ProjectrefModel,
//        })
//	})
//    if err != nil {
//        t.Fatal("ref create failed", err)
//    }
//    Projectcreateref = append(Projectcreateref, &models.FloatingIPProjectRef{UUID:"floating_ip_project_ref_uuid", To: []string{"test", "floating_ip_project_ref_uuid"}})
//    Projectcreateref = append(Projectcreateref, &models.FloatingIPProjectRef{UUID:"floating_ip_project_ref_uuid2", To: []string{"test", "floating_ip_project_ref_uuid2"}})
//    model.ProjectRefs = Projectcreateref
//    
//    var VirtualMachineInterfacecreateref []*models.FloatingIPVirtualMachineInterfaceRef
//    var VirtualMachineInterfacerefModel *models.VirtualMachineInterface
//    VirtualMachineInterfacerefModel = models.MakeVirtualMachineInterface()
//	VirtualMachineInterfacerefModel.UUID = "floating_ip_virtual_machine_interface_ref_uuid"
//    VirtualMachineInterfacerefModel.FQName = []string{"test", "floating_ip_virtual_machine_interface_ref_uuid"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
//            VirtualMachineInterface: VirtualMachineInterfacerefModel,
//        })
//	})
//    VirtualMachineInterfacerefModel.UUID = "floating_ip_virtual_machine_interface_ref_uuid1"
//    VirtualMachineInterfacerefModel.FQName = []string{"test", "floating_ip_virtual_machine_interface_ref_uuid1"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
//            VirtualMachineInterface: VirtualMachineInterfacerefModel,
//        })
//	})
//    VirtualMachineInterfacerefModel.UUID = "floating_ip_virtual_machine_interface_ref_uuid2"
//    VirtualMachineInterfacerefModel.FQName = []string{"test", "floating_ip_virtual_machine_interface_ref_uuid2"}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
//            VirtualMachineInterface: VirtualMachineInterfacerefModel,
//        })
//	})
//    if err != nil {
//        t.Fatal("ref create failed", err)
//    }
//    VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.FloatingIPVirtualMachineInterfaceRef{UUID:"floating_ip_virtual_machine_interface_ref_uuid", To: []string{"test", "floating_ip_virtual_machine_interface_ref_uuid"}})
//    VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.FloatingIPVirtualMachineInterfaceRef{UUID:"floating_ip_virtual_machine_interface_ref_uuid2", To: []string{"test", "floating_ip_virtual_machine_interface_ref_uuid2"}})
//    model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref
//    
//
    //create project to which resource is shared
    projectModel := models.MakeProject()
	projectModel.UUID = "floating_ip_admin_project_uuid"
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
//    common.SetValueByPath(updateMap, ".FloatingIPTrafficDirection", ".", "test")
//    
//    
//    
//    if ".FloatingIPPortMappings.PortMappings" == ".Perms2.Share" {
//        var share []interface{}
//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
//        common.SetValueByPath(updateMap, ".FloatingIPPortMappings.PortMappings", ".", share)
//    } else {
//        common.SetValueByPath(updateMap, ".FloatingIPPortMappings.PortMappings", ".", `{"test": "test"}`)
//    }
//    
//    
//    
//    common.SetValueByPath(updateMap, ".FloatingIPPortMappingsEnable", ".", true)
//    
//    
//    
//    common.SetValueByPath(updateMap, ".FloatingIPIsVirtualIP", ".", true)
//    
//    
//    
//    common.SetValueByPath(updateMap, ".FloatingIPFixedIPAddress", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".FloatingIPAddressFamily", ".", "test")
//    
//    
//    
//    common.SetValueByPath(updateMap, ".FloatingIPAddress", ".", "test")
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
//    common.SetValueByPath(updateMap, "uuid", ".", "floating_ip_dummy_uuid")
//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
//
//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
//    
//    var VirtualMachineInterfaceref []interface{}
//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"delete", "uuid":"floating_ip_virtual_machine_interface_ref_uuid", "to": []string{"test", "floating_ip_virtual_machine_interface_ref_uuid"}})
//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"add", "uuid":"floating_ip_virtual_machine_interface_ref_uuid1", "to": []string{"test", "floating_ip_virtual_machine_interface_ref_uuid1"}})
//    
//    
//    
//    common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)
//    
//    var Projectref []interface{}
//    Projectref = append(Projectref, map[string]interface{}{"operation":"delete", "uuid":"floating_ip_project_ref_uuid", "to": []string{"test", "floating_ip_project_ref_uuid"}})
//    Projectref = append(Projectref, map[string]interface{}{"operation":"add", "uuid":"floating_ip_project_ref_uuid1", "to": []string{"test", "floating_ip_project_ref_uuid1"}})
//    
//    
//    
//    common.SetValueByPath(updateMap, "ProjectRefs", ".", Projectref)
//    
//
    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateFloatingIP(ctx, tx, 
            &models.CreateFloatingIPRequest{
                FloatingIP: model,
            })
    })
    if err != nil {
        t.Fatal("create failed", err)
    }

//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
//        return UpdateFloatingIP(tx, model.UUID, updateMap)
//    })
//    if err != nil {
//        t.Fatal("update failed", err)
//    }

//    //Delete ref entries, referred objects
//    
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//        stmt, err := tx.Prepare("delete from `ref_floating_ip_virtual_machine_interface` where `from` = ? AND `to` = ?;")
//        if err != nil {
//            return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs delete statement failed")
//        }
//        _, err = stmt.Exec( "floating_ip_dummy_uuid", "floating_ip_virtual_machine_interface_ref_uuid" )
//        _, err = stmt.Exec( "floating_ip_dummy_uuid", "floating_ip_virtual_machine_interface_ref_uuid1" )
//        _, err = stmt.Exec( "floating_ip_dummy_uuid", "floating_ip_virtual_machine_interface_ref_uuid2" )
//        if err != nil {
//            return errors.Wrap(err, "VirtualMachineInterfaceRefs delete failed")
//        }
//        return nil
//	})
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteVirtualMachineInterface(ctx, tx, 
//            &models.DeleteVirtualMachineInterfaceRequest{
//                ID: "floating_ip_virtual_machine_interface_ref_uuid"})
//    })
//	if err != nil {
//		t.Fatal("delete ref floating_ip_virtual_machine_interface_ref_uuid  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteVirtualMachineInterface(ctx, tx, 
//            &models.DeleteVirtualMachineInterfaceRequest{
//                ID: "floating_ip_virtual_machine_interface_ref_uuid1"})
//    })
//	if err != nil {
//		t.Fatal("delete ref floating_ip_virtual_machine_interface_ref_uuid1  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteVirtualMachineInterface(
//            ctx, 
//            tx, 
//            &models.DeleteVirtualMachineInterfaceRequest{
//                ID: "floating_ip_virtual_machine_interface_ref_uuid2",
//            })
//    })
//	if err != nil {
//		t.Fatal("delete ref floating_ip_virtual_machine_interface_ref_uuid2 failed", err)
//	}
//    
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//        stmt, err := tx.Prepare("delete from `ref_floating_ip_project` where `from` = ? AND `to` = ?;")
//        if err != nil {
//            return errors.Wrap(err, "preparing ProjectRefs delete statement failed")
//        }
//        _, err = stmt.Exec( "floating_ip_dummy_uuid", "floating_ip_project_ref_uuid" )
//        _, err = stmt.Exec( "floating_ip_dummy_uuid", "floating_ip_project_ref_uuid1" )
//        _, err = stmt.Exec( "floating_ip_dummy_uuid", "floating_ip_project_ref_uuid2" )
//        if err != nil {
//            return errors.Wrap(err, "ProjectRefs delete failed")
//        }
//        return nil
//	})
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteProject(ctx, tx, 
//            &models.DeleteProjectRequest{
//                ID: "floating_ip_project_ref_uuid"})
//    })
//	if err != nil {
//		t.Fatal("delete ref floating_ip_project_ref_uuid  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteProject(ctx, tx, 
//            &models.DeleteProjectRequest{
//                ID: "floating_ip_project_ref_uuid1"})
//    })
//	if err != nil {
//		t.Fatal("delete ref floating_ip_project_ref_uuid1  failed", err)
//	}
//    err = common.DoInTransaction(db, func(tx *sql.Tx) error {
//    	return DeleteProject(
//            ctx, 
//            tx, 
//            &models.DeleteProjectRequest{
//                ID: "floating_ip_project_ref_uuid2",
//            })
//    })
//	if err != nil {
//		t.Fatal("delete ref floating_ip_project_ref_uuid2 failed", err)
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
        response, err := ListFloatingIP(ctx, tx, &models.ListFloatingIPRequest{
            Spec: &models.ListSpec{Limit: 1}})
        if err != nil {
            return err
        }
        if len(response.FloatingIPs) != 1 {
            return fmt.Errorf("expected one element")
        }
        return nil
    })
    if err != nil {
        t.Fatal("list failed", err)
    }
    
    ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteFloatingIP(ctxDemo, tx, 
            &models.DeleteFloatingIPRequest{
                ID: model.UUID}, 
        )
    })
    if err == nil {
        t.Fatal("auth failed")
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return DeleteFloatingIP(ctx, tx, 
            &models.DeleteFloatingIPRequest{
                ID: model.UUID})
    })
    if err != nil {
        t.Fatal("delete failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        return CreateFloatingIP(ctx, tx, 
            &models.CreateFloatingIPRequest{
                FloatingIP: model})
    })
    if err == nil {
        t.Fatal("Raise Error On Duplicate Create failed", err)
    }

    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
        response, err := ListFloatingIP(ctx, tx, &models.ListFloatingIPRequest{
            Spec: &models.ListSpec{Limit: 1}})
        if err != nil {
            return err
        }
        if len(response.FloatingIPs) != 0 {
            return fmt.Errorf("expected no element")
        }
        return nil
    })
    if err != nil {
        t.Fatal("list failed", err)
    }
    return
}
