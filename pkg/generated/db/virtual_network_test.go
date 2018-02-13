package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

func TestVirtualNetwork(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "virtual_network")
	defer func() {
		common.ClearTable(db, "virtual_network")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeVirtualNetwork()
	model.UUID = "virtual_network_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "virtual_network_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var RouteTablecreateref []*models.VirtualNetworkRouteTableRef
	var RouteTablerefModel *models.RouteTable
	RouteTablerefModel = models.MakeRouteTable()
	RouteTablerefModel.UUID = "virtual_network_route_table_ref_uuid"
	RouteTablerefModel.FQName = []string{"test", "virtual_network_route_table_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateRouteTable(tx, RouteTablerefModel)
	})
	RouteTablerefModel.UUID = "virtual_network_route_table_ref_uuid1"
	RouteTablerefModel.FQName = []string{"test", "virtual_network_route_table_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateRouteTable(tx, RouteTablerefModel)
	})
	RouteTablerefModel.UUID = "virtual_network_route_table_ref_uuid2"
	RouteTablerefModel.FQName = []string{"test", "virtual_network_route_table_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateRouteTable(tx, RouteTablerefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	RouteTablecreateref = append(RouteTablecreateref, &models.VirtualNetworkRouteTableRef{UUID: "virtual_network_route_table_ref_uuid", To: []string{"test", "virtual_network_route_table_ref_uuid"}})
	RouteTablecreateref = append(RouteTablecreateref, &models.VirtualNetworkRouteTableRef{UUID: "virtual_network_route_table_ref_uuid2", To: []string{"test", "virtual_network_route_table_ref_uuid2"}})
	model.RouteTableRefs = RouteTablecreateref

	var VirtualNetworkcreateref []*models.VirtualNetworkVirtualNetworkRef
	var VirtualNetworkrefModel *models.VirtualNetwork
	VirtualNetworkrefModel = models.MakeVirtualNetwork()
	VirtualNetworkrefModel.UUID = "virtual_network_virtual_network_ref_uuid"
	VirtualNetworkrefModel.FQName = []string{"test", "virtual_network_virtual_network_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(tx, VirtualNetworkrefModel)
	})
	VirtualNetworkrefModel.UUID = "virtual_network_virtual_network_ref_uuid1"
	VirtualNetworkrefModel.FQName = []string{"test", "virtual_network_virtual_network_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(tx, VirtualNetworkrefModel)
	})
	VirtualNetworkrefModel.UUID = "virtual_network_virtual_network_ref_uuid2"
	VirtualNetworkrefModel.FQName = []string{"test", "virtual_network_virtual_network_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(tx, VirtualNetworkrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.VirtualNetworkVirtualNetworkRef{UUID: "virtual_network_virtual_network_ref_uuid", To: []string{"test", "virtual_network_virtual_network_ref_uuid"}})
	VirtualNetworkcreateref = append(VirtualNetworkcreateref, &models.VirtualNetworkVirtualNetworkRef{UUID: "virtual_network_virtual_network_ref_uuid2", To: []string{"test", "virtual_network_virtual_network_ref_uuid2"}})
	model.VirtualNetworkRefs = VirtualNetworkcreateref

	var BGPVPNcreateref []*models.VirtualNetworkBGPVPNRef
	var BGPVPNrefModel *models.BGPVPN
	BGPVPNrefModel = models.MakeBGPVPN()
	BGPVPNrefModel.UUID = "virtual_network_bgpvpn_ref_uuid"
	BGPVPNrefModel.FQName = []string{"test", "virtual_network_bgpvpn_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateBGPVPN(tx, BGPVPNrefModel)
	})
	BGPVPNrefModel.UUID = "virtual_network_bgpvpn_ref_uuid1"
	BGPVPNrefModel.FQName = []string{"test", "virtual_network_bgpvpn_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateBGPVPN(tx, BGPVPNrefModel)
	})
	BGPVPNrefModel.UUID = "virtual_network_bgpvpn_ref_uuid2"
	BGPVPNrefModel.FQName = []string{"test", "virtual_network_bgpvpn_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateBGPVPN(tx, BGPVPNrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BGPVPNcreateref = append(BGPVPNcreateref, &models.VirtualNetworkBGPVPNRef{UUID: "virtual_network_bgpvpn_ref_uuid", To: []string{"test", "virtual_network_bgpvpn_ref_uuid"}})
	BGPVPNcreateref = append(BGPVPNcreateref, &models.VirtualNetworkBGPVPNRef{UUID: "virtual_network_bgpvpn_ref_uuid2", To: []string{"test", "virtual_network_bgpvpn_ref_uuid2"}})
	model.BGPVPNRefs = BGPVPNcreateref

	var NetworkIpamcreateref []*models.VirtualNetworkNetworkIpamRef
	var NetworkIpamrefModel *models.NetworkIpam
	NetworkIpamrefModel = models.MakeNetworkIpam()
	NetworkIpamrefModel.UUID = "virtual_network_network_ipam_ref_uuid"
	NetworkIpamrefModel.FQName = []string{"test", "virtual_network_network_ipam_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(tx, NetworkIpamrefModel)
	})
	NetworkIpamrefModel.UUID = "virtual_network_network_ipam_ref_uuid1"
	NetworkIpamrefModel.FQName = []string{"test", "virtual_network_network_ipam_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(tx, NetworkIpamrefModel)
	})
	NetworkIpamrefModel.UUID = "virtual_network_network_ipam_ref_uuid2"
	NetworkIpamrefModel.FQName = []string{"test", "virtual_network_network_ipam_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(tx, NetworkIpamrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NetworkIpamcreateref = append(NetworkIpamcreateref, &models.VirtualNetworkNetworkIpamRef{UUID: "virtual_network_network_ipam_ref_uuid", To: []string{"test", "virtual_network_network_ipam_ref_uuid"}})
	NetworkIpamcreateref = append(NetworkIpamcreateref, &models.VirtualNetworkNetworkIpamRef{UUID: "virtual_network_network_ipam_ref_uuid2", To: []string{"test", "virtual_network_network_ipam_ref_uuid2"}})
	model.NetworkIpamRefs = NetworkIpamcreateref

	var SecurityLoggingObjectcreateref []*models.VirtualNetworkSecurityLoggingObjectRef
	var SecurityLoggingObjectrefModel *models.SecurityLoggingObject
	SecurityLoggingObjectrefModel = models.MakeSecurityLoggingObject()
	SecurityLoggingObjectrefModel.UUID = "virtual_network_security_logging_object_ref_uuid"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "virtual_network_security_logging_object_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateSecurityLoggingObject(tx, SecurityLoggingObjectrefModel)
	})
	SecurityLoggingObjectrefModel.UUID = "virtual_network_security_logging_object_ref_uuid1"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "virtual_network_security_logging_object_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateSecurityLoggingObject(tx, SecurityLoggingObjectrefModel)
	})
	SecurityLoggingObjectrefModel.UUID = "virtual_network_security_logging_object_ref_uuid2"
	SecurityLoggingObjectrefModel.FQName = []string{"test", "virtual_network_security_logging_object_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateSecurityLoggingObject(tx, SecurityLoggingObjectrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.VirtualNetworkSecurityLoggingObjectRef{UUID: "virtual_network_security_logging_object_ref_uuid", To: []string{"test", "virtual_network_security_logging_object_ref_uuid"}})
	SecurityLoggingObjectcreateref = append(SecurityLoggingObjectcreateref, &models.VirtualNetworkSecurityLoggingObjectRef{UUID: "virtual_network_security_logging_object_ref_uuid2", To: []string{"test", "virtual_network_security_logging_object_ref_uuid2"}})
	model.SecurityLoggingObjectRefs = SecurityLoggingObjectcreateref

	var NetworkPolicycreateref []*models.VirtualNetworkNetworkPolicyRef
	var NetworkPolicyrefModel *models.NetworkPolicy
	NetworkPolicyrefModel = models.MakeNetworkPolicy()
	NetworkPolicyrefModel.UUID = "virtual_network_network_policy_ref_uuid"
	NetworkPolicyrefModel.FQName = []string{"test", "virtual_network_network_policy_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkPolicy(tx, NetworkPolicyrefModel)
	})
	NetworkPolicyrefModel.UUID = "virtual_network_network_policy_ref_uuid1"
	NetworkPolicyrefModel.FQName = []string{"test", "virtual_network_network_policy_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkPolicy(tx, NetworkPolicyrefModel)
	})
	NetworkPolicyrefModel.UUID = "virtual_network_network_policy_ref_uuid2"
	NetworkPolicyrefModel.FQName = []string{"test", "virtual_network_network_policy_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkPolicy(tx, NetworkPolicyrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NetworkPolicycreateref = append(NetworkPolicycreateref, &models.VirtualNetworkNetworkPolicyRef{UUID: "virtual_network_network_policy_ref_uuid", To: []string{"test", "virtual_network_network_policy_ref_uuid"}})
	NetworkPolicycreateref = append(NetworkPolicycreateref, &models.VirtualNetworkNetworkPolicyRef{UUID: "virtual_network_network_policy_ref_uuid2", To: []string{"test", "virtual_network_network_policy_ref_uuid2"}})
	model.NetworkPolicyRefs = NetworkPolicycreateref

	var QosConfigcreateref []*models.VirtualNetworkQosConfigRef
	var QosConfigrefModel *models.QosConfig
	QosConfigrefModel = models.MakeQosConfig()
	QosConfigrefModel.UUID = "virtual_network_qos_config_ref_uuid"
	QosConfigrefModel.FQName = []string{"test", "virtual_network_qos_config_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateQosConfig(tx, QosConfigrefModel)
	})
	QosConfigrefModel.UUID = "virtual_network_qos_config_ref_uuid1"
	QosConfigrefModel.FQName = []string{"test", "virtual_network_qos_config_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateQosConfig(tx, QosConfigrefModel)
	})
	QosConfigrefModel.UUID = "virtual_network_qos_config_ref_uuid2"
	QosConfigrefModel.FQName = []string{"test", "virtual_network_qos_config_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateQosConfig(tx, QosConfigrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	QosConfigcreateref = append(QosConfigcreateref, &models.VirtualNetworkQosConfigRef{UUID: "virtual_network_qos_config_ref_uuid", To: []string{"test", "virtual_network_qos_config_ref_uuid"}})
	QosConfigcreateref = append(QosConfigcreateref, &models.VirtualNetworkQosConfigRef{UUID: "virtual_network_qos_config_ref_uuid2", To: []string{"test", "virtual_network_qos_config_ref_uuid2"}})
	model.QosConfigRefs = QosConfigcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "virtual_network_admin_project_uuid"
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

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.VxlanNetworkIdentifier", ".", 1.0)

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.RPF", ".", "test")

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.NetworkID", ".", 1.0)

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.MirrorDestination", ".", true)

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.ForwardingMode", ".", "test")

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.AllowTransit", ".", true)

	common.SetValueByPath(updateMap, ".VirtualNetworkNetworkID", ".", 1.0)

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".RouterExternal", ".", true)

	if ".RouteTargetList.RouteTarget" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".RouteTargetList.RouteTarget", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".RouteTargetList.RouteTarget", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".ProviderProperties.SegmentationID", ".", 1.0)

	common.SetValueByPath(updateMap, ".ProviderProperties.PhysicalNetwork", ".", "test")

	common.SetValueByPath(updateMap, ".PortSecurityEnabled", ".", true)

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

	common.SetValueByPath(updateMap, ".PBBEvpnEnable", ".", true)

	common.SetValueByPath(updateMap, ".PBBEtreeEnable", ".", true)

	common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ParentType", ".", "test")

	common.SetValueByPath(updateMap, ".MultiPolicyServiceChainsEnabled", ".", true)

	common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveTimeWindow", ".", 1.0)

	common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveLimitAction", ".", "test")

	common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveLimit", ".", 1.0)

	common.SetValueByPath(updateMap, ".MacLimitControl.MacLimitAction", ".", "test")

	common.SetValueByPath(updateMap, ".MacLimitControl.MacLimit", ".", 1.0)

	common.SetValueByPath(updateMap, ".MacLearningEnabled", ".", true)

	common.SetValueByPath(updateMap, ".MacAgingTime", ".", 1.0)

	common.SetValueByPath(updateMap, ".Layer2ControlWord", ".", true)

	common.SetValueByPath(updateMap, ".IsShared", ".", true)

	if ".ImportRouteTargetList.RouteTarget" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".ImportRouteTargetList.RouteTarget", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".ImportRouteTargetList.RouteTarget", ".", `{"test": "test"}`)
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

	common.SetValueByPath(updateMap, ".FloodUnknownUnicast", ".", true)

	common.SetValueByPath(updateMap, ".ExternalIpam", ".", true)

	if ".ExportRouteTargetList.RouteTarget" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".ExportRouteTargetList.RouteTarget", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".ExportRouteTargetList.RouteTarget", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.SourcePort", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.SourceIP", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.IPProtocol", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.HashingConfigured", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.DestinationPort", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.DestinationIP", ".", true)

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	if ".Annotations.KeyValuePair" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".AddressAllocationMode", ".", "test")

	common.SetValueByPath(updateMap, "uuid", ".", "virtual_network_dummy_uuid")
	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	// Create Attr values for testing ref update(ADD,UPDATE,DELETE)

	var NetworkPolicyref []interface{}
	NetworkPolicyref = append(NetworkPolicyref, map[string]interface{}{"operation": "delete", "uuid": "virtual_network_network_policy_ref_uuid", "to": []string{"test", "virtual_network_network_policy_ref_uuid"}})
	NetworkPolicyref = append(NetworkPolicyref, map[string]interface{}{"operation": "add", "uuid": "virtual_network_network_policy_ref_uuid1", "to": []string{"test", "virtual_network_network_policy_ref_uuid1"}})

	NetworkPolicyAttr := map[string]interface{}{}

	common.SetValueByPath(NetworkPolicyAttr, ".Sequence.Major", ".", 1.0)

	common.SetValueByPath(NetworkPolicyAttr, ".Sequence.Minor", ".", 1.0)

	common.SetValueByPath(NetworkPolicyAttr, ".Timer.EndTime", ".", "test")

	common.SetValueByPath(NetworkPolicyAttr, ".Timer.StartTime", ".", "test")

	common.SetValueByPath(NetworkPolicyAttr, ".Timer.OffInterval", ".", "test")

	common.SetValueByPath(NetworkPolicyAttr, ".Timer.OnInterval", ".", "test")

	NetworkPolicyref = append(NetworkPolicyref, map[string]interface{}{"operation": "update", "uuid": "virtual_network_network_policy_ref_uuid2", "to": []string{"test", "virtual_network_network_policy_ref_uuid2"}, "attr": NetworkPolicyAttr})

	common.SetValueByPath(updateMap, "NetworkPolicyRefs", ".", NetworkPolicyref)

	var QosConfigref []interface{}
	QosConfigref = append(QosConfigref, map[string]interface{}{"operation": "delete", "uuid": "virtual_network_qos_config_ref_uuid", "to": []string{"test", "virtual_network_qos_config_ref_uuid"}})
	QosConfigref = append(QosConfigref, map[string]interface{}{"operation": "add", "uuid": "virtual_network_qos_config_ref_uuid1", "to": []string{"test", "virtual_network_qos_config_ref_uuid1"}})

	common.SetValueByPath(updateMap, "QosConfigRefs", ".", QosConfigref)

	var RouteTableref []interface{}
	RouteTableref = append(RouteTableref, map[string]interface{}{"operation": "delete", "uuid": "virtual_network_route_table_ref_uuid", "to": []string{"test", "virtual_network_route_table_ref_uuid"}})
	RouteTableref = append(RouteTableref, map[string]interface{}{"operation": "add", "uuid": "virtual_network_route_table_ref_uuid1", "to": []string{"test", "virtual_network_route_table_ref_uuid1"}})

	common.SetValueByPath(updateMap, "RouteTableRefs", ".", RouteTableref)

	var VirtualNetworkref []interface{}
	VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation": "delete", "uuid": "virtual_network_virtual_network_ref_uuid", "to": []string{"test", "virtual_network_virtual_network_ref_uuid"}})
	VirtualNetworkref = append(VirtualNetworkref, map[string]interface{}{"operation": "add", "uuid": "virtual_network_virtual_network_ref_uuid1", "to": []string{"test", "virtual_network_virtual_network_ref_uuid1"}})

	common.SetValueByPath(updateMap, "VirtualNetworkRefs", ".", VirtualNetworkref)

	var BGPVPNref []interface{}
	BGPVPNref = append(BGPVPNref, map[string]interface{}{"operation": "delete", "uuid": "virtual_network_bgpvpn_ref_uuid", "to": []string{"test", "virtual_network_bgpvpn_ref_uuid"}})
	BGPVPNref = append(BGPVPNref, map[string]interface{}{"operation": "add", "uuid": "virtual_network_bgpvpn_ref_uuid1", "to": []string{"test", "virtual_network_bgpvpn_ref_uuid1"}})

	common.SetValueByPath(updateMap, "BGPVPNRefs", ".", BGPVPNref)

	var NetworkIpamref []interface{}
	NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation": "delete", "uuid": "virtual_network_network_ipam_ref_uuid", "to": []string{"test", "virtual_network_network_ipam_ref_uuid"}})
	NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation": "add", "uuid": "virtual_network_network_ipam_ref_uuid1", "to": []string{"test", "virtual_network_network_ipam_ref_uuid1"}})

	NetworkIpamAttr := map[string]interface{}{}

	common.SetValueByPath(NetworkIpamAttr, ".IpamSubnets", ".", map[string]string{"test": "test"})

	common.SetValueByPath(NetworkIpamAttr, ".HostRoutes.Route", ".", map[string]string{"test": "test"})

	NetworkIpamref = append(NetworkIpamref, map[string]interface{}{"operation": "update", "uuid": "virtual_network_network_ipam_ref_uuid2", "to": []string{"test", "virtual_network_network_ipam_ref_uuid2"}, "attr": NetworkIpamAttr})

	common.SetValueByPath(updateMap, "NetworkIpamRefs", ".", NetworkIpamref)

	var SecurityLoggingObjectref []interface{}
	SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation": "delete", "uuid": "virtual_network_security_logging_object_ref_uuid", "to": []string{"test", "virtual_network_security_logging_object_ref_uuid"}})
	SecurityLoggingObjectref = append(SecurityLoggingObjectref, map[string]interface{}{"operation": "add", "uuid": "virtual_network_security_logging_object_ref_uuid1", "to": []string{"test", "virtual_network_security_logging_object_ref_uuid1"}})

	common.SetValueByPath(updateMap, "SecurityLoggingObjectRefs", ".", SecurityLoggingObjectref)

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateVirtualNetwork(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_network_route_table` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing RouteTableRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_route_table_ref_uuid")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_route_table_ref_uuid1")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_route_table_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "RouteTableRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteRouteTable(tx, "virtual_network_route_table_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_route_table_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteRouteTable(tx, "virtual_network_route_table_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_route_table_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteRouteTable(tx, "virtual_network_route_table_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_route_table_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_network_virtual_network` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualNetworkRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_virtual_network_ref_uuid")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_virtual_network_ref_uuid1")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_virtual_network_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(tx, "virtual_network_virtual_network_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_virtual_network_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(tx, "virtual_network_virtual_network_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_virtual_network_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(tx, "virtual_network_virtual_network_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_virtual_network_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_network_bgpvpn` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing BGPVPNRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_bgpvpn_ref_uuid")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_bgpvpn_ref_uuid1")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_bgpvpn_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "BGPVPNRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteBGPVPN(tx, "virtual_network_bgpvpn_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_bgpvpn_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteBGPVPN(tx, "virtual_network_bgpvpn_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_bgpvpn_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteBGPVPN(tx, "virtual_network_bgpvpn_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_bgpvpn_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_network_network_ipam` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing NetworkIpamRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_network_ipam_ref_uuid")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_network_ipam_ref_uuid1")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_network_ipam_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "NetworkIpamRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(tx, "virtual_network_network_ipam_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_network_ipam_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(tx, "virtual_network_network_ipam_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_network_ipam_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(tx, "virtual_network_network_ipam_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_network_ipam_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_network_security_logging_object` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing SecurityLoggingObjectRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_security_logging_object_ref_uuid")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_security_logging_object_ref_uuid1")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_security_logging_object_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteSecurityLoggingObject(tx, "virtual_network_security_logging_object_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_security_logging_object_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteSecurityLoggingObject(tx, "virtual_network_security_logging_object_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_security_logging_object_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteSecurityLoggingObject(tx, "virtual_network_security_logging_object_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_security_logging_object_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_network_network_policy` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing NetworkPolicyRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_network_policy_ref_uuid")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_network_policy_ref_uuid1")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_network_policy_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "NetworkPolicyRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkPolicy(tx, "virtual_network_network_policy_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_network_policy_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkPolicy(tx, "virtual_network_network_policy_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_network_policy_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkPolicy(tx, "virtual_network_network_policy_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_network_policy_ref_uuid2 failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_virtual_network_qos_config` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing QosConfigRefs delete statement failed")
		}
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_qos_config_ref_uuid")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_qos_config_ref_uuid1")
		_, err = stmt.Exec("virtual_network_dummy_uuid", "virtual_network_qos_config_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "QosConfigRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteQosConfig(tx, "virtual_network_qos_config_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_qos_config_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteQosConfig(tx, "virtual_network_qos_config_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_qos_config_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteQosConfig(tx, "virtual_network_qos_config_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref virtual_network_qos_config_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, projectModel.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListVirtualNetwork(tx, &common.ListSpec{Limit: 1})
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
		return DeleteVirtualNetwork(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(tx, model)
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListVirtualNetwork(tx, &common.ListSpec{Limit: 1})
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
