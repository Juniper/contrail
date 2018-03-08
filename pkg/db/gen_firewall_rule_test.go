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

func TestFirewallRule(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "firewall_rule")
	// mutexProject := UseTable(db.DB, "firewall_rule")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeFirewallRule()
	model.UUID = "firewall_rule_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "firewall_rule_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var AddressGroupcreateref []*models.FirewallRuleAddressGroupRef
	var AddressGrouprefModel *models.AddressGroup
	AddressGrouprefModel = models.MakeAddressGroup()
	AddressGrouprefModel.UUID = "firewall_rule_address_group_ref_uuid"
	AddressGrouprefModel.FQName = []string{"test", "firewall_rule_address_group_ref_uuid"}
	_, err = db.CreateAddressGroup(ctx, &models.CreateAddressGroupRequest{
		AddressGroup: AddressGrouprefModel,
	})
	AddressGrouprefModel.UUID = "firewall_rule_address_group_ref_uuid1"
	AddressGrouprefModel.FQName = []string{"test", "firewall_rule_address_group_ref_uuid1"}
	_, err = db.CreateAddressGroup(ctx, &models.CreateAddressGroupRequest{
		AddressGroup: AddressGrouprefModel,
	})
	AddressGrouprefModel.UUID = "firewall_rule_address_group_ref_uuid2"
	AddressGrouprefModel.FQName = []string{"test", "firewall_rule_address_group_ref_uuid2"}
	_, err = db.CreateAddressGroup(ctx, &models.CreateAddressGroupRequest{
		AddressGroup: AddressGrouprefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	AddressGroupcreateref = append(AddressGroupcreateref, &models.FirewallRuleAddressGroupRef{UUID: "firewall_rule_address_group_ref_uuid", To: []string{"test", "firewall_rule_address_group_ref_uuid"}})
	AddressGroupcreateref = append(AddressGroupcreateref, &models.FirewallRuleAddressGroupRef{UUID: "firewall_rule_address_group_ref_uuid2", To: []string{"test", "firewall_rule_address_group_ref_uuid2"}})
	model.AddressGroupRefs = AddressGroupcreateref

	var SecurityLoggingObjectcreateref []*models.FirewallRuleSecurityLoggingObjectRef
	var SecurityLoggingObjectrefModel *models.SecurityLoggingObject
	SecurityLoggingObjectrefModel = models.MakeSecurityLoggingObject()
	SecurityLoggingObjectrefModel.UUID = "firewall_rule_security_logging_object_ref_uuid"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "firewall_rule_security_logging_object_ref_uuid"}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectrefModel,
	})
	SecurityLoggingObjectrefModel.UUID = "firewall_rule_security_logging_object_ref_uuid1"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "firewall_rule_security_logging_object_ref_uuid1"}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectrefModel,
	})
	SecurityLoggingObjectrefModel.UUID = "firewall_rule_security_logging_object_ref_uuid2"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "firewall_rule_security_logging_object_ref_uuid2"}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.FirewallRuleSecurityLoggingObjectRef{UUID: "firewall_rule_security_logging_object_ref_uuid", To: []string{"test", "firewall_rule_security_logging_object_ref_uuid"}})
	SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.FirewallRuleSecurityLoggingObjectRef{UUID: "firewall_rule_security_logging_object_ref_uuid2", To: []string{"test", "firewall_rule_security_logging_object_ref_uuid2"}})
	model.SecurityLoggingObjectRefs = SecurityLoggingObjectcreateref

	var VirtualNetworkcreateref []*models.FirewallRuleVirtualNetworkRef
	var VirtualNetworkrefModel *models.VirtualNetwork
	VirtualNetworkrefModel = models.MakeVirtualNetwork()
	VirtualNetworkrefModel.UUID = "firewall_rule_virtual_network_ref_uuid"
	VirtualNetworkrefModel.FQName = []string{"test", "firewall_rule_virtual_network_ref_uuid"}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkrefModel,
	})
	VirtualNetworkrefModel.UUID = "firewall_rule_virtual_network_ref_uuid1"
	VirtualNetworkrefModel.FQName = []string{"test", "firewall_rule_virtual_network_ref_uuid1"}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkrefModel,
	})
	VirtualNetworkrefModel.UUID = "firewall_rule_virtual_network_ref_uuid2"
	VirtualNetworkrefModel.FQName = []string{"test", "firewall_rule_virtual_network_ref_uuid2"}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.FirewallRuleVirtualNetworkRef{UUID: "firewall_rule_virtual_network_ref_uuid", To: []string{"test", "firewall_rule_virtual_network_ref_uuid"}})
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.FirewallRuleVirtualNetworkRef{UUID: "firewall_rule_virtual_network_ref_uuid2", To: []string{"test", "firewall_rule_virtual_network_ref_uuid2"}})
	model.VirtualNetworkRefs = VirtualNetworkcreateref

	var ServiceGroupcreateref []*models.FirewallRuleServiceGroupRef
	var ServiceGrouprefModel *models.ServiceGroup
	ServiceGrouprefModel = models.MakeServiceGroup()
	ServiceGrouprefModel.UUID = "firewall_rule_service_group_ref_uuid"
	ServiceGrouprefModel.FQName = []string{"test", "firewall_rule_service_group_ref_uuid"}
	_, err = db.CreateServiceGroup(ctx, &models.CreateServiceGroupRequest{
		ServiceGroup: ServiceGrouprefModel,
	})
	ServiceGrouprefModel.UUID = "firewall_rule_service_group_ref_uuid1"
	ServiceGrouprefModel.FQName = []string{"test", "firewall_rule_service_group_ref_uuid1"}
	_, err = db.CreateServiceGroup(ctx, &models.CreateServiceGroupRequest{
		ServiceGroup: ServiceGrouprefModel,
	})
	ServiceGrouprefModel.UUID = "firewall_rule_service_group_ref_uuid2"
	ServiceGrouprefModel.FQName = []string{"test", "firewall_rule_service_group_ref_uuid2"}
	_, err = db.CreateServiceGroup(ctx, &models.CreateServiceGroupRequest{
		ServiceGroup: ServiceGrouprefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceGroupcreateref = append(ServiceGroupcreateref, &models.FirewallRuleServiceGroupRef{UUID: "firewall_rule_service_group_ref_uuid", To: []string{"test", "firewall_rule_service_group_ref_uuid"}})
	ServiceGroupcreateref = append(ServiceGroupcreateref, &models.FirewallRuleServiceGroupRef{UUID: "firewall_rule_service_group_ref_uuid2", To: []string{"test", "firewall_rule_service_group_ref_uuid2"}})
	model.ServiceGroupRefs = ServiceGroupcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "firewall_rule_admin_project_uuid"
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
	//    common.SetValueByPath(updateMap, ".Service.SRCPorts.StartPort", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Service.SRCPorts.EndPort", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Service.ProtocolID", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Service.Protocol", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Service.DSTPorts.StartPort", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Service.DSTPorts.EndPort", ".", 1.0)
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
	//    if ".MatchTags.TagList" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".MatchTags.TagList", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".MatchTags.TagList", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    if ".MatchTagTypes.TagType" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".MatchTagTypes.TagType", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".MatchTagTypes.TagType", ".", `{"test": "test"}`)
	//    }
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
	//    common.SetValueByPath(updateMap, ".Endpoint2.VirtualNetwork", ".", "test")
	//
	//
	//
	//    if ".Endpoint2.Tags" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Endpoint2.Tags", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Endpoint2.Tags", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    if ".Endpoint2.TagIds" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Endpoint2.TagIds", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Endpoint2.TagIds", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Endpoint2.Subnet.IPPrefixLen", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Endpoint2.Subnet.IPPrefix", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Endpoint2.Any", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Endpoint2.AddressGroup", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Endpoint1.VirtualNetwork", ".", "test")
	//
	//
	//
	//    if ".Endpoint1.Tags" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Endpoint1.Tags", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Endpoint1.Tags", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    if ".Endpoint1.TagIds" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Endpoint1.TagIds", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Endpoint1.TagIds", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Endpoint1.Subnet.IPPrefixLen", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Endpoint1.Subnet.IPPrefix", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Endpoint1.Any", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Endpoint1.AddressGroup", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Direction", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ConfigurationVersion", ".", 1.0)
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
	//
	//    common.SetValueByPath(updateMap, ".ActionList.SimpleAction", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.QosAction", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.UDPPort", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.StaticNHHeader.Vni", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.RoutingInstance", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.NicAssistedMirroringVlan", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.NicAssistedMirroring", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.NHMode", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.JuniperHeader", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.Encapsulation", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.AnalyzerName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.AnalyzerMacAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.MirrorTo.AnalyzerIPAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.Log", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.GatewayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.AssignRoutingInstance", ".", "test")
	//
	//
	//
	//    if ".ActionList.ApplyService" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".ActionList.ApplyService", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".ActionList.ApplyService", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ActionList.Alert", ".", true)
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "firewall_rule_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var SecurityLoggingObjectref []interface{}
	//    SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation":"delete", "uuid":"firewall_rule_security_logging_object_ref_uuid", "to": []string{"test", "firewall_rule_security_logging_object_ref_uuid"}})
	//    SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation":"add", "uuid":"firewall_rule_security_logging_object_ref_uuid1", "to": []string{"test", "firewall_rule_security_logging_object_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "SecurityLoggingObjectRefs", ".", SecurityLoggingObjectref)
	//
	//    var VirtualNetworkref []interface{}
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"delete", "uuid":"firewall_rule_virtual_network_ref_uuid", "to": []string{"test", "firewall_rule_virtual_network_ref_uuid"}})
	//    VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation":"add", "uuid":"firewall_rule_virtual_network_ref_uuid1", "to": []string{"test", "firewall_rule_virtual_network_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualNetworkRefs", ".", VirtualNetworkref)
	//
	//    var ServiceGroupref []interface{}
	//    ServiceGroupref = append(ServiceGroupref, map[string]interface{}{"operation":"delete", "uuid":"firewall_rule_service_group_ref_uuid", "to": []string{"test", "firewall_rule_service_group_ref_uuid"}})
	//    ServiceGroupref = append(ServiceGroupref, map[string]interface{}{"operation":"add", "uuid":"firewall_rule_service_group_ref_uuid1", "to": []string{"test", "firewall_rule_service_group_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ServiceGroupRefs", ".", ServiceGroupref)
	//
	//    var AddressGroupref []interface{}
	//    AddressGroupref = append(AddressGroupref, map[string]interface{}{"operation":"delete", "uuid":"firewall_rule_address_group_ref_uuid", "to": []string{"test", "firewall_rule_address_group_ref_uuid"}})
	//    AddressGroupref = append(AddressGroupref, map[string]interface{}{"operation":"add", "uuid":"firewall_rule_address_group_ref_uuid1", "to": []string{"test", "firewall_rule_address_group_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "AddressGroupRefs", ".", AddressGroupref)
	//
	//
	_, err = db.CreateFirewallRule(ctx,
		&models.CreateFirewallRuleRequest{
			FirewallRule: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateFirewallRule(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_firewall_rule_service_group` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceGroupRefs delete statement failed")
		}
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_service_group_ref_uuid")
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_service_group_ref_uuid1")
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_service_group_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceGroupRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteServiceGroup(ctx,
		&models.DeleteServiceGroupRequest{
			ID: "firewall_rule_service_group_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref firewall_rule_service_group_ref_uuid  failed", err)
	}
	_, err = db.DeleteServiceGroup(ctx,
		&models.DeleteServiceGroupRequest{
			ID: "firewall_rule_service_group_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref firewall_rule_service_group_ref_uuid1  failed", err)
	}
	_, err = db.DeleteServiceGroup(
		ctx,
		&models.DeleteServiceGroupRequest{
			ID: "firewall_rule_service_group_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref firewall_rule_service_group_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_firewall_rule_address_group` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing AddressGroupRefs delete statement failed")
		}
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_address_group_ref_uuid")
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_address_group_ref_uuid1")
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_address_group_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "AddressGroupRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteAddressGroup(ctx,
		&models.DeleteAddressGroupRequest{
			ID: "firewall_rule_address_group_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref firewall_rule_address_group_ref_uuid  failed", err)
	}
	_, err = db.DeleteAddressGroup(ctx,
		&models.DeleteAddressGroupRequest{
			ID: "firewall_rule_address_group_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref firewall_rule_address_group_ref_uuid1  failed", err)
	}
	_, err = db.DeleteAddressGroup(
		ctx,
		&models.DeleteAddressGroupRequest{
			ID: "firewall_rule_address_group_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref firewall_rule_address_group_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_firewall_rule_security_logging_object` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing SecurityLoggingObjectRefs delete statement failed")
		}
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_security_logging_object_ref_uuid")
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_security_logging_object_ref_uuid1")
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_security_logging_object_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteSecurityLoggingObject(ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: "firewall_rule_security_logging_object_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref firewall_rule_security_logging_object_ref_uuid  failed", err)
	}
	_, err = db.DeleteSecurityLoggingObject(ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: "firewall_rule_security_logging_object_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref firewall_rule_security_logging_object_ref_uuid1  failed", err)
	}
	_, err = db.DeleteSecurityLoggingObject(
		ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: "firewall_rule_security_logging_object_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref firewall_rule_security_logging_object_ref_uuid2 failed", err)
	}

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_firewall_rule_virtual_network` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualNetworkRefs delete statement failed")
		}
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_virtual_network_ref_uuid")
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_virtual_network_ref_uuid1")
		_, err = stmt.Exec("firewall_rule_dummy_uuid", "firewall_rule_virtual_network_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteVirtualNetwork(ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: "firewall_rule_virtual_network_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref firewall_rule_virtual_network_ref_uuid  failed", err)
	}
	_, err = db.DeleteVirtualNetwork(ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: "firewall_rule_virtual_network_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref firewall_rule_virtual_network_ref_uuid1  failed", err)
	}
	_, err = db.DeleteVirtualNetwork(
		ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: "firewall_rule_virtual_network_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref firewall_rule_virtual_network_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListFirewallRule(ctx, &models.ListFirewallRuleRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.FirewallRules) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteFirewallRule(ctxDemo,
		&models.DeleteFirewallRuleRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateFirewallRule(ctx,
		&models.CreateFirewallRuleRequest{
			FirewallRule: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteFirewallRule(ctx,
		&models.DeleteFirewallRuleRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListFirewallRule(ctx, &models.ListFirewallRuleRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.FirewallRules) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
