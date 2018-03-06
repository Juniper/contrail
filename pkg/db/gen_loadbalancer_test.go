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

func TestLoadbalancer(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "loadbalancer")
	// mutexProject := UseTable(db.DB, "loadbalancer")
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
	_, err = db.CreateServiceApplianceSet(ctx, &models.CreateServiceApplianceSetRequest{
		ServiceApplianceSet: ServiceApplianceSetrefModel,
	})
	ServiceApplianceSetrefModel.UUID = "loadbalancer_service_appliance_set_ref_uuid1"
	ServiceApplianceSetrefModel.FQName = []string{"test", "loadbalancer_service_appliance_set_ref_uuid1"}
	_, err = db.CreateServiceApplianceSet(ctx, &models.CreateServiceApplianceSetRequest{
		ServiceApplianceSet: ServiceApplianceSetrefModel,
	})
	ServiceApplianceSetrefModel.UUID = "loadbalancer_service_appliance_set_ref_uuid2"
	ServiceApplianceSetrefModel.FQName = []string{"test", "loadbalancer_service_appliance_set_ref_uuid2"}
	_, err = db.CreateServiceApplianceSet(ctx, &models.CreateServiceApplianceSetRequest{
		ServiceApplianceSet: ServiceApplianceSetrefModel,
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
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	VirtualMachineInterfacerefModel.UUID = "loadbalancer_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "loadbalancer_virtual_machine_interface_ref_uuid1"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	VirtualMachineInterfacerefModel.UUID = "loadbalancer_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "loadbalancer_virtual_machine_interface_ref_uuid2"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
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
	_, err = db.CreateServiceInstance(ctx, &models.CreateServiceInstanceRequest{
		ServiceInstance: ServiceInstancerefModel,
	})
	ServiceInstancerefModel.UUID = "loadbalancer_service_instance_ref_uuid1"
	ServiceInstancerefModel.FQName = []string{"test", "loadbalancer_service_instance_ref_uuid1"}
	_, err = db.CreateServiceInstance(ctx, &models.CreateServiceInstanceRequest{
		ServiceInstance: ServiceInstancerefModel,
	})
	ServiceInstancerefModel.UUID = "loadbalancer_service_instance_ref_uuid2"
	ServiceInstancerefModel.FQName = []string{"test", "loadbalancer_service_instance_ref_uuid2"}
	_, err = db.CreateServiceInstance(ctx, &models.CreateServiceInstanceRequest{
		ServiceInstance: ServiceInstancerefModel,
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
	//    var ServiceApplianceSetref []interface{}
	//    ServiceApplianceSetref = append(ServiceApplianceSetref, map[string]interface{}{"operation":"delete", "uuid":"loadbalancer_service_appliance_set_ref_uuid", "to": []string{"test", "loadbalancer_service_appliance_set_ref_uuid"}})
	//    ServiceApplianceSetref = append(ServiceApplianceSetref, map[string]interface{}{"operation":"add", "uuid":"loadbalancer_service_appliance_set_ref_uuid1", "to": []string{"test", "loadbalancer_service_appliance_set_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ServiceApplianceSetRefs", ".", ServiceApplianceSetref)
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
	//
	_, err = db.CreateLoadbalancer(ctx,
		&models.CreateLoadbalancerRequest{
			Loadbalancer: model,
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

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
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
	_, err = db.DeleteServiceApplianceSet(ctx,
		&models.DeleteServiceApplianceSetRequest{
			ID: "loadbalancer_service_appliance_set_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_appliance_set_ref_uuid  failed", err)
	}
	_, err = db.DeleteServiceApplianceSet(ctx,
		&models.DeleteServiceApplianceSetRequest{
			ID: "loadbalancer_service_appliance_set_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_appliance_set_ref_uuid1  failed", err)
	}
	_, err = db.DeleteServiceApplianceSet(
		ctx,
		&models.DeleteServiceApplianceSetRequest{
			ID: "loadbalancer_service_appliance_set_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_appliance_set_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
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
	_, err = db.DeleteVirtualMachineInterface(ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "loadbalancer_virtual_machine_interface_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref loadbalancer_virtual_machine_interface_ref_uuid  failed", err)
	}
	_, err = db.DeleteVirtualMachineInterface(ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "loadbalancer_virtual_machine_interface_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref loadbalancer_virtual_machine_interface_ref_uuid1  failed", err)
	}
	_, err = db.DeleteVirtualMachineInterface(
		ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "loadbalancer_virtual_machine_interface_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref loadbalancer_virtual_machine_interface_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
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
	_, err = db.DeleteServiceInstance(ctx,
		&models.DeleteServiceInstanceRequest{
			ID: "loadbalancer_service_instance_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_instance_ref_uuid  failed", err)
	}
	_, err = db.DeleteServiceInstance(ctx,
		&models.DeleteServiceInstanceRequest{
			ID: "loadbalancer_service_instance_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_instance_ref_uuid1  failed", err)
	}
	_, err = db.DeleteServiceInstance(
		ctx,
		&models.DeleteServiceInstanceRequest{
			ID: "loadbalancer_service_instance_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref loadbalancer_service_instance_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListLoadbalancer(ctx, &models.ListLoadbalancerRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.Loadbalancers) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteLoadbalancer(ctxDemo,
		&models.DeleteLoadbalancerRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateLoadbalancer(ctx,
		&models.CreateLoadbalancerRequest{
			Loadbalancer: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteLoadbalancer(ctx,
		&models.DeleteLoadbalancerRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListLoadbalancer(ctx, &models.ListLoadbalancerRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.Loadbalancers) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
