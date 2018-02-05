package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

func TestLoadbalancerPool(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "loadbalancer_pool")
	defer func() {
		common.ClearTable(db, "loadbalancer_pool")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeLoadbalancerPool()
	model.UUID = "loadbalancer_pool_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "loadbalancer_pool_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceApplianceSetcreateref []*models.LoadbalancerPoolServiceApplianceSetRef
	var ServiceApplianceSetrefModel *models.ServiceApplianceSet
	ServiceApplianceSetrefModel = models.MakeServiceApplianceSet()
	ServiceApplianceSetrefModel.UUID = "loadbalancer_pool_service_appliance_set_ref_uuid"
	ServiceApplianceSetrefModel.FQName = []string{"test", "loadbalancer_pool_service_appliance_set_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceApplianceSet(tx, ServiceApplianceSetrefModel)
	})
	ServiceApplianceSetrefModel.UUID = "loadbalancer_pool_service_appliance_set_ref_uuid1"
	ServiceApplianceSetrefModel.FQName = []string{"test", "loadbalancer_pool_service_appliance_set_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceApplianceSet(tx, ServiceApplianceSetrefModel)
	})
	ServiceApplianceSetrefModel.UUID = "loadbalancer_pool_service_appliance_set_ref_uuid2"
	ServiceApplianceSetrefModel.FQName = []string{"test", "loadbalancer_pool_service_appliance_set_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceApplianceSet(tx, ServiceApplianceSetrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceApplianceSetcreateref = append(ServiceApplianceSetcreateref, &models.LoadbalancerPoolServiceApplianceSetRef{UUID: "loadbalancer_pool_service_appliance_set_ref_uuid", To: []string{"test", "loadbalancer_pool_service_appliance_set_ref_uuid"}})
	ServiceApplianceSetcreateref = append(ServiceApplianceSetcreateref, &models.LoadbalancerPoolServiceApplianceSetRef{UUID: "loadbalancer_pool_service_appliance_set_ref_uuid2", To: []string{"test", "loadbalancer_pool_service_appliance_set_ref_uuid2"}})
	model.ServiceApplianceSetRefs = ServiceApplianceSetcreateref

	var VirtualMachineInterfacecreateref []*models.LoadbalancerPoolVirtualMachineInterfaceRef
	var VirtualMachineInterfacerefModel *models.VirtualMachineInterface
	VirtualMachineInterfacerefModel = models.MakeVirtualMachineInterface()
	VirtualMachineInterfacerefModel.UUID = "loadbalancer_pool_virtual_machine_interface_ref_uuid"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "loadbalancer_pool_virtual_machine_interface_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(tx, VirtualMachineInterfacerefModel)
	})
	VirtualMachineInterfacerefModel.UUID = "loadbalancer_pool_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "loadbalancer_pool_virtual_machine_interface_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(tx, VirtualMachineInterfacerefModel)
	})
	VirtualMachineInterfacerefModel.UUID = "loadbalancer_pool_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "loadbalancer_pool_virtual_machine_interface_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(tx, VirtualMachineInterfacerefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.LoadbalancerPoolVirtualMachineInterfaceRef{UUID: "loadbalancer_pool_virtual_machine_interface_ref_uuid", To: []string{"test", "loadbalancer_pool_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.LoadbalancerPoolVirtualMachineInterfaceRef{UUID: "loadbalancer_pool_virtual_machine_interface_ref_uuid2", To: []string{"test", "loadbalancer_pool_virtual_machine_interface_ref_uuid2"}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref

	var LoadbalancerListenercreateref []*models.LoadbalancerPoolLoadbalancerListenerRef
	var LoadbalancerListenerrefModel *models.LoadbalancerListener
	LoadbalancerListenerrefModel = models.MakeLoadbalancerListener()
	LoadbalancerListenerrefModel.UUID = "loadbalancer_pool_loadbalancer_listener_ref_uuid"
	LoadbalancerListenerrefModel.FQName = []string{"test", "loadbalancer_pool_loadbalancer_listener_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerListener(tx, LoadbalancerListenerrefModel)
	})
	LoadbalancerListenerrefModel.UUID = "loadbalancer_pool_loadbalancer_listener_ref_uuid1"
	LoadbalancerListenerrefModel.FQName = []string{"test", "loadbalancer_pool_loadbalancer_listener_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerListener(tx, LoadbalancerListenerrefModel)
	})
	LoadbalancerListenerrefModel.UUID = "loadbalancer_pool_loadbalancer_listener_ref_uuid2"
	LoadbalancerListenerrefModel.FQName = []string{"test", "loadbalancer_pool_loadbalancer_listener_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerListener(tx, LoadbalancerListenerrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	LoadbalancerListenercreateref = append(LoadbalancerListenercreateref, &models.LoadbalancerPoolLoadbalancerListenerRef{UUID: "loadbalancer_pool_loadbalancer_listener_ref_uuid", To: []string{"test", "loadbalancer_pool_loadbalancer_listener_ref_uuid"}})
	LoadbalancerListenercreateref = append(LoadbalancerListenercreateref, &models.LoadbalancerPoolLoadbalancerListenerRef{UUID: "loadbalancer_pool_loadbalancer_listener_ref_uuid2", To: []string{"test", "loadbalancer_pool_loadbalancer_listener_ref_uuid2"}})
	model.LoadbalancerListenerRefs = LoadbalancerListenercreateref

	var ServiceInstancecreateref []*models.LoadbalancerPoolServiceInstanceRef
	var ServiceInstancerefModel *models.ServiceInstance
	ServiceInstancerefModel = models.MakeServiceInstance()
	ServiceInstancerefModel.UUID = "loadbalancer_pool_service_instance_ref_uuid"
	ServiceInstancerefModel.FQName = []string{"test", "loadbalancer_pool_service_instance_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(tx, ServiceInstancerefModel)
	})
	ServiceInstancerefModel.UUID = "loadbalancer_pool_service_instance_ref_uuid1"
	ServiceInstancerefModel.FQName = []string{"test", "loadbalancer_pool_service_instance_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(tx, ServiceInstancerefModel)
	})
	ServiceInstancerefModel.UUID = "loadbalancer_pool_service_instance_ref_uuid2"
	ServiceInstancerefModel.FQName = []string{"test", "loadbalancer_pool_service_instance_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(tx, ServiceInstancerefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceInstancecreateref = append(ServiceInstancecreateref, &models.LoadbalancerPoolServiceInstanceRef{UUID: "loadbalancer_pool_service_instance_ref_uuid", To: []string{"test", "loadbalancer_pool_service_instance_ref_uuid"}})
	ServiceInstancecreateref = append(ServiceInstancecreateref, &models.LoadbalancerPoolServiceInstanceRef{UUID: "loadbalancer_pool_service_instance_ref_uuid2", To: []string{"test", "loadbalancer_pool_service_instance_ref_uuid2"}})
	model.ServiceInstanceRefs = ServiceInstancecreateref

	var LoadbalancerHealthmonitorcreateref []*models.LoadbalancerPoolLoadbalancerHealthmonitorRef
	var LoadbalancerHealthmonitorrefModel *models.LoadbalancerHealthmonitor
	LoadbalancerHealthmonitorrefModel = models.MakeLoadbalancerHealthmonitor()
	LoadbalancerHealthmonitorrefModel.UUID = "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid"
	LoadbalancerHealthmonitorrefModel.FQName = []string{"test", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerHealthmonitor(tx, LoadbalancerHealthmonitorrefModel)
	})
	LoadbalancerHealthmonitorrefModel.UUID = "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid1"
	LoadbalancerHealthmonitorrefModel.FQName = []string{"test", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerHealthmonitor(tx, LoadbalancerHealthmonitorrefModel)
	})
	LoadbalancerHealthmonitorrefModel.UUID = "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid2"
	LoadbalancerHealthmonitorrefModel.FQName = []string{"test", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerHealthmonitor(tx, LoadbalancerHealthmonitorrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	LoadbalancerHealthmonitorcreateref = append(LoadbalancerHealthmonitorcreateref, &models.LoadbalancerPoolLoadbalancerHealthmonitorRef{UUID: "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid", To: []string{"test", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid"}})
	LoadbalancerHealthmonitorcreateref = append(LoadbalancerHealthmonitorcreateref, &models.LoadbalancerPoolLoadbalancerHealthmonitorRef{UUID: "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid2", To: []string{"test", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid2"}})
	model.LoadbalancerHealthmonitorRefs = LoadbalancerHealthmonitorcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "loadbalancer_pool_admin_project_uuid"
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

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

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

	common.SetValueByPath(updateMap, ".LoadbalancerPoolProvider", ".", "test")

	common.SetValueByPath(updateMap, ".LoadbalancerPoolProperties.SubnetID", ".", "test")

	common.SetValueByPath(updateMap, ".LoadbalancerPoolProperties.StatusDescription", ".", "test")

	common.SetValueByPath(updateMap, ".LoadbalancerPoolProperties.Status", ".", "test")

	common.SetValueByPath(updateMap, ".LoadbalancerPoolProperties.SessionPersistence", ".", "test")

	common.SetValueByPath(updateMap, ".LoadbalancerPoolProperties.Protocol", ".", "test")

	common.SetValueByPath(updateMap, ".LoadbalancerPoolProperties.PersistenceCookieName", ".", "test")

	common.SetValueByPath(updateMap, ".LoadbalancerPoolProperties.LoadbalancerMethod", ".", "test")

	common.SetValueByPath(updateMap, ".LoadbalancerPoolProperties.AdminState", ".", true)

	if ".LoadbalancerPoolCustomAttributes.KeyValuePair" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".LoadbalancerPoolCustomAttributes.KeyValuePair", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".LoadbalancerPoolCustomAttributes.KeyValuePair", ".", `{"test": "test"}`)
	}

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

	common.SetValueByPath(updateMap, "uuid", ".", "loadbalancer_pool_dummy_uuid")
	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	// Create Attr values for testing ref update(ADD,UPDATE,DELETE)

	var ServiceApplianceSetref []interface{}
	ServiceApplianceSetref = append(ServiceApplianceSetref, map[string]interface{}{"operation": "delete", "uuid": "loadbalancer_pool_service_appliance_set_ref_uuid", "to": []string{"test", "loadbalancer_pool_service_appliance_set_ref_uuid"}})
	ServiceApplianceSetref = append(ServiceApplianceSetref, map[string]interface{}{"operation": "add", "uuid": "loadbalancer_pool_service_appliance_set_ref_uuid1", "to": []string{"test", "loadbalancer_pool_service_appliance_set_ref_uuid1"}})

	common.SetValueByPath(updateMap, "ServiceApplianceSetRefs", ".", ServiceApplianceSetref)

	var VirtualMachineInterfaceref []interface{}
	VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation": "delete", "uuid": "loadbalancer_pool_virtual_machine_interface_ref_uuid", "to": []string{"test", "loadbalancer_pool_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation": "add", "uuid": "loadbalancer_pool_virtual_machine_interface_ref_uuid1", "to": []string{"test", "loadbalancer_pool_virtual_machine_interface_ref_uuid1"}})

	common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)

	var LoadbalancerListenerref []interface{}
	LoadbalancerListenerref = append(LoadbalancerListenerref, map[string]interface{}{"operation": "delete", "uuid": "loadbalancer_pool_loadbalancer_listener_ref_uuid", "to": []string{"test", "loadbalancer_pool_loadbalancer_listener_ref_uuid"}})
	LoadbalancerListenerref = append(LoadbalancerListenerref, map[string]interface{}{"operation": "add", "uuid": "loadbalancer_pool_loadbalancer_listener_ref_uuid1", "to": []string{"test", "loadbalancer_pool_loadbalancer_listener_ref_uuid1"}})

	common.SetValueByPath(updateMap, "LoadbalancerListenerRefs", ".", LoadbalancerListenerref)

	var ServiceInstanceref []interface{}
	ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation": "delete", "uuid": "loadbalancer_pool_service_instance_ref_uuid", "to": []string{"test", "loadbalancer_pool_service_instance_ref_uuid"}})
	ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation": "add", "uuid": "loadbalancer_pool_service_instance_ref_uuid1", "to": []string{"test", "loadbalancer_pool_service_instance_ref_uuid1"}})

	common.SetValueByPath(updateMap, "ServiceInstanceRefs", ".", ServiceInstanceref)

	var LoadbalancerHealthmonitorref []interface{}
	LoadbalancerHealthmonitorref = append(LoadbalancerHealthmonitorref, map[string]interface{}{"operation": "delete", "uuid": "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid", "to": []string{"test", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid"}})
	LoadbalancerHealthmonitorref = append(LoadbalancerHealthmonitorref, map[string]interface{}{"operation": "add", "uuid": "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid1", "to": []string{"test", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid1"}})

	common.SetValueByPath(updateMap, "LoadbalancerHealthmonitorRefs", ".", LoadbalancerHealthmonitorref)

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerPool(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateLoadbalancerPool(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_loadbalancer_pool_loadbalancer_listener` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing LoadbalancerListenerRefs delete statement failed")
		}
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_loadbalancer_listener_ref_uuid")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_loadbalancer_listener_ref_uuid1")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_loadbalancer_listener_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "LoadbalancerListenerRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerListener(tx, "loadbalancer_pool_loadbalancer_listener_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_loadbalancer_listener_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerListener(tx, "loadbalancer_pool_loadbalancer_listener_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_loadbalancer_listener_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerListener(tx, "loadbalancer_pool_loadbalancer_listener_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_loadbalancer_listener_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_loadbalancer_pool_service_instance` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceInstanceRefs delete statement failed")
		}
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_service_instance_ref_uuid")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_service_instance_ref_uuid1")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_service_instance_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(tx, "loadbalancer_pool_service_instance_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_service_instance_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(tx, "loadbalancer_pool_service_instance_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_service_instance_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(tx, "loadbalancer_pool_service_instance_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_service_instance_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_loadbalancer_pool_loadbalancer_healthmonitor` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing LoadbalancerHealthmonitorRefs delete statement failed")
		}
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid1")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "LoadbalancerHealthmonitorRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerHealthmonitor(tx, "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerHealthmonitor(tx, "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerHealthmonitor(tx, "loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_loadbalancer_healthmonitor_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_loadbalancer_pool_service_appliance_set` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceApplianceSetRefs delete statement failed")
		}
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_service_appliance_set_ref_uuid")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_service_appliance_set_ref_uuid1")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_service_appliance_set_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceApplianceSetRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceApplianceSet(tx, "loadbalancer_pool_service_appliance_set_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_service_appliance_set_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceApplianceSet(tx, "loadbalancer_pool_service_appliance_set_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_service_appliance_set_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceApplianceSet(tx, "loadbalancer_pool_service_appliance_set_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_service_appliance_set_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_loadbalancer_pool_virtual_machine_interface` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs delete statement failed")
		}
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_virtual_machine_interface_ref_uuid")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_virtual_machine_interface_ref_uuid1")
		_, err = stmt.Exec("loadbalancer_pool_dummy_uuid", "loadbalancer_pool_virtual_machine_interface_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(tx, "loadbalancer_pool_virtual_machine_interface_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_virtual_machine_interface_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(tx, "loadbalancer_pool_virtual_machine_interface_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_virtual_machine_interface_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualMachineInterface(tx, "loadbalancer_pool_virtual_machine_interface_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref loadbalancer_pool_virtual_machine_interface_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, projectModel.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListLoadbalancerPool(tx, &common.ListSpec{Limit: 1})
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
		return DeleteLoadbalancerPool(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLoadbalancerPool(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLoadbalancerPool(tx, model)
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListLoadbalancerPool(tx, &common.ListSpec{Limit: 1})
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
