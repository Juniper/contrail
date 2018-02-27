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

//For skip import error.
var _ = errors.New("")

func TestLoadbalancer(t *testing.T) {
	// t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db, "metadata")
	mutexTable := common.UseTable(db, "loadbalancer")
	// mutexProject := common.UseTable(db, "loadbalancer")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeLoadbalancer()
	model.UUID = "loadbalancer_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "loadbalancer_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceApplianceSetcreateref []*models.LoadbalancerServiceApplianceSetRef
	var ServiceApplianceSetrefModel *models.ServiceApplianceSet
	ServiceApplianceSetrefModel = models.MakeServiceApplianceSet()
	ServiceApplianceSetrefModel.UUID = "loadbalancer_service_appliance_set_ref_uuid"
	ServiceApplianceSetrefModel.FQName = []string{"test", "loadbalancer_service_appliance_set_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceApplianceSet(ctx, tx, &models.CreateServiceApplianceSetRequest{
			ServiceApplianceSet: ServiceApplianceSetrefModel,
		})
	})
	ServiceApplianceSetrefModel.UUID = "loadbalancer_service_appliance_set_ref_uuid1"
	ServiceApplianceSetrefModel.FQName = []string{"test", "loadbalancer_service_appliance_set_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceApplianceSet(ctx, tx, &models.CreateServiceApplianceSetRequest{
			ServiceApplianceSet: ServiceApplianceSetrefModel,
		})
	})
	ServiceApplianceSetrefModel.UUID = "loadbalancer_service_appliance_set_ref_uuid2"
	ServiceApplianceSetrefModel.FQName = []string{"test", "loadbalancer_service_appliance_set_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceApplianceSet(ctx, tx, &models.CreateServiceApplianceSetRequest{
			ServiceApplianceSet: ServiceApplianceSetrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceApplianceSetcreateref = append(ServiceApplianceSetcreateref, &models.LoadbalancerServiceApplianceSetRef{UUID: "loadbalancer_service_appliance_set_ref_uuid", To: []string{"test", "loadbalancer_service_appliance_set_ref_uuid"}})
	ServiceApplianceSetcreateref = append(ServiceApplianceSetcreateref, &models.LoadbalancerServiceApplianceSetRef{UUID: "loadbalancer_service_appliance_set_ref_uuid2", To: []string{"test", "loadbalancer_service_appliance_set_ref_uuid2"}})
	model.ServiceApplianceSetRefs = ServiceApplianceSetcreateref

	var VirtualMachineInterfacecreateref []*models.LoadbalancerVirtualMachineInterfaceRef
	var VirtualMachineInterfacerefModel *models.VirtualMachineInterface
	VirtualMachineInterfacerefModel = models.MakeVirtualMachineInterface()
	VirtualMachineInterfacerefModel.UUID = "loadbalancer_virtual_machine_interface_ref_uuid"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "loadbalancer_virtual_machine_interface_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	VirtualMachineInterfacerefModel.UUID = "loadbalancer_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "loadbalancer_virtual_machine_interface_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	VirtualMachineInterfacerefModel.UUID = "loadbalancer_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "loadbalancer_virtual_machine_interface_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(ctx, tx, &models.CreateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: VirtualMachineInterfacerefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.LoadbalancerVirtualMachineInterfaceRef{UUID: "loadbalancer_virtual_machine_interface_ref_uuid", To: []string{"test", "loadbalancer_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.LoadbalancerVirtualMachineInterfaceRef{UUID: "loadbalancer_virtual_machine_interface_ref_uuid2", To: []string{"test", "loadbalancer_virtual_machine_interface_ref_uuid2"}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref

	var ServiceInstancecreateref []*models.LoadbalancerServiceInstanceRef
	var ServiceInstancerefModel *models.ServiceInstance
	ServiceInstancerefModel = models.MakeServiceInstance()
	ServiceInstancerefModel.UUID = "loadbalancer_service_instance_ref_uuid"
	ServiceInstancerefModel.FQName = []string{"test", "loadbalancer_service_instance_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(ctx, tx, &models.CreateServiceInstanceRequest{
			ServiceInstance: ServiceInstancerefModel,
		})
	})
	ServiceInstancerefModel.UUID = "loadbalancer_service_instance_ref_uuid1"
	ServiceInstancerefModel.FQName = []string{"test", "loadbalancer_service_instance_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(ctx, tx, &models.CreateServiceInstanceRequest{
			ServiceInstance: ServiceInstancerefModel,
		})
	})
	ServiceInstancerefModel.UUID = "loadbalancer_service_instance_ref_uuid2"
	ServiceInstancerefModel.FQName = []string{"test", "loadbalancer_service_instance_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(ctx, tx, &models.CreateServiceInstanceRequest{
			ServiceInstance: ServiceInstancerefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceInstancecreateref = append(ServiceInstancecreateref, &models.LoadbalancerServiceInstanceRef{UUID: "loadbalancer_service_instance_ref_uuid", To: []string{"test", "loadbalancer_service_instance_ref_uuid"}})
	ServiceInstancecreateref = append(ServiceInstancecreateref, &models.LoadbalancerServiceInstanceRef{UUID: "loadbalancer_service_instance_ref_uuid2", To: []string{"test", "loadbalancer_service_instance_ref_uuid2"}})
	model.ServiceInstanceRefs = ServiceInstancecreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "loadbalancer_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".LoadbalancerProvider", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".LoadbalancerProperties.VipSubnetID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".LoadbalancerProperties.VipAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".LoadbalancerProperties.Status", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".LoadbalancerProperties.ProvisioningStatus", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".LoadbalancerProperties.OperatingStatus", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".LoadbalancerProperties.AdminState", ".", true)
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
	//    common.SetValueByPath(updateMap, "uuid", ".", "loadbalancer_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var VirtualMachineInterfaceref []interface{}
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"delete", "uuid":"loadbalancer_virtual_machine_interface_ref_uuid", "to": []string{"test", "loadbalancer_virtual_machine_interface_ref_uuid"}})
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"add", "uuid":"loadbalancer_virtual_machine_interface_ref_uuid1", "to": []string{"test", "loadbalancer_virtual_machine_interface_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)
	//
	//    var ServiceInstanceref []interface{}
	//    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"delete", "uuid":"loadbalancer_service_instance_ref_uuid", "to": []string{"test", "loadbalancer_service_instance_ref_uuid"}})
	//    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"add", "uuid":"loadbalancer_service_instance_ref_uuid1", "to": []string{"test", "loadbalancer_service_instance_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ServiceInstanceRefs", ".", ServiceInstanceref)
	//
	//    var ServiceApplianceSetref []interface{}
	//    ServiceApplianceSetref = append(ServiceApplianceSetref, map[string]interface{}{"operation":"delete", "uuid":"loadbalancer_service_appliance_set_ref_uuid", "to": []string{"test", "loadbalancer_service_appliance_set_ref_uuid"}})
	//    ServiceApplianceSetref = append(ServiceApplianceSetref, map[string]interface{}{"operation":"add", "uuid":"loadbalancer_service_appliance_set_ref_uuid1", "to": []string{"test", "loadbalancer_service_appliance_set_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ServiceApplianceSetRefs", ".", ServiceApplianceSetref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancer(ctx, tx,
			&models.CreateLoadbalancerRequest{
				Loadbalancer: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateLoadbalancer(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_loadbalancer_virtual_machine_interface` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs delete statement failed")
		}
		_, err = stmt.Exec("loadbalancer_dummy_uuid", "loadbalancer_virtual_machine_interface_ref_uuid")
		_, err = stmt.Exec("loadbalancer_dummy_uuid", "loadbalancer_virtual_machine_interface_ref_uuid1")
		_, err = stmt.Exec("loadbalancer_dummy_uuid", "loadbalancer_virtual_machine_interface_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(ctx, tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "loadbalancer_virtual_machine_interface_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_virtual_machine_interface_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(ctx, tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "loadbalancer_virtual_machine_interface_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_virtual_machine_interface_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(
			ctx,
			tx,
			&models.DeleteVirtualMachineInterfaceRequest{
				ID: "loadbalancer_virtual_machine_interface_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_virtual_machine_interface_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_loadbalancer_service_instance` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceInstanceRefs delete statement failed")
		}
		_, err = stmt.Exec("loadbalancer_dummy_uuid", "loadbalancer_service_instance_ref_uuid")
		_, err = stmt.Exec("loadbalancer_dummy_uuid", "loadbalancer_service_instance_ref_uuid1")
		_, err = stmt.Exec("loadbalancer_dummy_uuid", "loadbalancer_service_instance_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(ctx, tx,
			&models.DeleteServiceInstanceRequest{
				ID: "loadbalancer_service_instance_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_instance_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(ctx, tx,
			&models.DeleteServiceInstanceRequest{
				ID: "loadbalancer_service_instance_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_instance_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(
			ctx,
			tx,
			&models.DeleteServiceInstanceRequest{
				ID: "loadbalancer_service_instance_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_instance_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_loadbalancer_service_appliance_set` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceApplianceSetRefs delete statement failed")
		}
		_, err = stmt.Exec("loadbalancer_dummy_uuid", "loadbalancer_service_appliance_set_ref_uuid")
		_, err = stmt.Exec("loadbalancer_dummy_uuid", "loadbalancer_service_appliance_set_ref_uuid1")
		_, err = stmt.Exec("loadbalancer_dummy_uuid", "loadbalancer_service_appliance_set_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceApplianceSetRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceApplianceSet(ctx, tx,
			&models.DeleteServiceApplianceSetRequest{
				ID: "loadbalancer_service_appliance_set_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_appliance_set_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceApplianceSet(ctx, tx,
			&models.DeleteServiceApplianceSetRequest{
				ID: "loadbalancer_service_appliance_set_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_appliance_set_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceApplianceSet(
			ctx,
			tx,
			&models.DeleteServiceApplianceSetRequest{
				ID: "loadbalancer_service_appliance_set_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_appliance_set_ref_uuid2 failed", err)
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
		response, err := ListLoadbalancer(ctx, tx, &models.ListLoadbalancerRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.Loadbalancers) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancer(ctxDemo, tx,
			&models.DeleteLoadbalancerRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancer(ctx, tx,
			&models.DeleteLoadbalancerRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancer(ctx, tx,
			&models.CreateLoadbalancerRequest{
				Loadbalancer: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListLoadbalancer(ctx, tx, &models.ListLoadbalancerRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.Loadbalancers) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
