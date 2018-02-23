package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualMachineInterfaceQuery = "insert into `virtual_machine_interface` (`vrf_assign_rule`,`vlan_tag_based_bridge_domain`,`sub_interface_vlan_tag`,`service_interface_type`,`local_preference`,`traffic_direction`,`udp_port`,`vtep_dst_mac_address`,`vtep_dst_ip_address`,`vni`,`routing_instance`,`nic_assisted_mirroring_vlan`,`nic_assisted_mirroring`,`nh_mode`,`juniper_header`,`encapsulation`,`analyzer_name`,`analyzer_mac_address`,`analyzer_ip_address`,`mac_address`,`route`,`fat_flow_protocol`,`virtual_machine_interface_disable_policy`,`dhcp_option`,`virtual_machine_interface_device_owner`,`key_value_pair`,`allowed_address_pair`,`uuid`,`port_security_enabled`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`source_port`,`source_ip`,`ip_protocol`,`hashing_configured`,`destination_port`,`destination_ip`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteVirtualMachineInterfaceQuery = "delete from `virtual_machine_interface` where uuid = ?"

// VirtualMachineInterfaceFields is db columns for VirtualMachineInterface
var VirtualMachineInterfaceFields = []string{
	"vrf_assign_rule",
	"vlan_tag_based_bridge_domain",
	"sub_interface_vlan_tag",
	"service_interface_type",
	"local_preference",
	"traffic_direction",
	"udp_port",
	"vtep_dst_mac_address",
	"vtep_dst_ip_address",
	"vni",
	"routing_instance",
	"nic_assisted_mirroring_vlan",
	"nic_assisted_mirroring",
	"nh_mode",
	"juniper_header",
	"encapsulation",
	"analyzer_name",
	"analyzer_mac_address",
	"analyzer_ip_address",
	"mac_address",
	"route",
	"fat_flow_protocol",
	"virtual_machine_interface_disable_policy",
	"dhcp_option",
	"virtual_machine_interface_device_owner",
	"key_value_pair",
	"allowed_address_pair",
	"uuid",
	"port_security_enabled",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"user_visible",
	"permissions_owner_access",
	"permissions_owner",
	"other_access",
	"group_access",
	"group",
	"last_modified",
	"enable",
	"description",
	"creator",
	"created",
	"fq_name",
	"source_port",
	"source_ip",
	"ip_protocol",
	"hashing_configured",
	"destination_port",
	"destination_ip",
	"display_name",
	"annotations_key_value_pair",
}

// VirtualMachineInterfaceRefFields is db reference fields for VirtualMachineInterface
var VirtualMachineInterfaceRefFields = map[string][]string{

	"security_logging_object": []string{
		// <schema.Schema Value>

	},

	"interface_route_table": []string{
		// <schema.Schema Value>

	},

	"security_group": []string{
		// <schema.Schema Value>

	},

	"bridge_domain": []string{
		// <schema.Schema Value>
		"vlan_tag",
	},

	"service_endpoint": []string{
		// <schema.Schema Value>

	},

	"virtual_machine_interface": []string{
		// <schema.Schema Value>

	},

	"bgp_router": []string{
		// <schema.Schema Value>

	},

	"port_tuple": []string{
		// <schema.Schema Value>

	},

	"virtual_network": []string{
		// <schema.Schema Value>

	},

	"virtual_machine": []string{
		// <schema.Schema Value>

	},

	"routing_instance": []string{
		// <schema.Schema Value>
		"service_chain_address",
		"dst_mac",
		"protocol",
		"ipv6_service_chain_address",
		"direction",
		"mpls_label",
		"vlan_tag",
		"src_mac",
	},

	"qos_config": []string{
		// <schema.Schema Value>

	},

	"physical_interface": []string{
		// <schema.Schema Value>

	},

	"service_health_check": []string{
		// <schema.Schema Value>

	},
}

// VirtualMachineInterfaceBackRefFields is db back reference fields for VirtualMachineInterface
var VirtualMachineInterfaceBackRefFields = map[string][]string{}

// VirtualMachineInterfaceParentTypes is possible parents for VirtualMachineInterface
var VirtualMachineInterfaceParents = []string{

	"project",

	"virtual_machine",

	"virtual_router",
}

const insertVirtualMachineInterfaceBGPRouterQuery = "insert into `ref_virtual_machine_interface_bgp_router` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfacePortTupleQuery = "insert into `ref_virtual_machine_interface_port_tuple` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualNetworkQuery = "insert into `ref_virtual_machine_interface_virtual_network` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceServiceEndpointQuery = "insert into `ref_virtual_machine_interface_service_endpoint` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualMachineInterfaceQuery = "insert into `ref_virtual_machine_interface_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceRoutingInstanceQuery = "insert into `ref_virtual_machine_interface_routing_instance` (`from`, `to` ,`service_chain_address`,`dst_mac`,`protocol`,`ipv6_service_chain_address`,`direction`,`mpls_label`,`vlan_tag`,`src_mac`) values (?, ?,?,?,?,?,?,?,?,?);"

const insertVirtualMachineInterfaceQosConfigQuery = "insert into `ref_virtual_machine_interface_qos_config` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualMachineQuery = "insert into `ref_virtual_machine_interface_virtual_machine` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceServiceHealthCheckQuery = "insert into `ref_virtual_machine_interface_service_health_check` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfacePhysicalInterfaceQuery = "insert into `ref_virtual_machine_interface_physical_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceInterfaceRouteTableQuery = "insert into `ref_virtual_machine_interface_interface_route_table` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceSecurityGroupQuery = "insert into `ref_virtual_machine_interface_security_group` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceBridgeDomainQuery = "insert into `ref_virtual_machine_interface_bridge_domain` (`from`, `to` ,`vlan_tag`) values (?, ?,?);"

const insertVirtualMachineInterfaceSecurityLoggingObjectQuery = "insert into `ref_virtual_machine_interface_security_logging_object` (`from`, `to` ) values (?, ?);"

// CreateVirtualMachineInterface inserts VirtualMachineInterface to DB
func CreateVirtualMachineInterface(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateVirtualMachineInterfaceRequest) error {
	model := request.VirtualMachineInterface
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVirtualMachineInterfaceQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, common.MustJSON(model.GetVRFAssignTable().GetVRFAssignRule()),
		bool(model.GetVlanTagBasedBridgeDomain()),
		int(model.GetVirtualMachineInterfaceProperties().GetSubInterfaceVlanTag()),
		string(model.GetVirtualMachineInterfaceProperties().GetServiceInterfaceType()),
		int(model.GetVirtualMachineInterfaceProperties().GetLocalPreference()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetTrafficDirection()),
		int(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetUDPPort()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetStaticNHHeader().GetVtepDSTMacAddress()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetStaticNHHeader().GetVtepDSTIPAddress()),
		int(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetStaticNHHeader().GetVni()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetRoutingInstance()),
		int(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetNicAssistedMirroringVlan()),
		bool(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetNicAssistedMirroring()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetNHMode()),
		bool(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetJuniperHeader()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetEncapsulation()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetAnalyzerName()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetAnalyzerMacAddress()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetAnalyzerIPAddress()),
		common.MustJSON(model.GetVirtualMachineInterfaceMacAddresses().GetMacAddress()),
		common.MustJSON(model.GetVirtualMachineInterfaceHostRoutes().GetRoute()),
		common.MustJSON(model.GetVirtualMachineInterfaceFatFlowProtocols().GetFatFlowProtocol()),
		bool(model.GetVirtualMachineInterfaceDisablePolicy()),
		common.MustJSON(model.GetVirtualMachineInterfaceDHCPOptionList().GetDHCPOption()),
		string(model.GetVirtualMachineInterfaceDeviceOwner()),
		common.MustJSON(model.GetVirtualMachineInterfaceBindings().GetKeyValuePair()),
		common.MustJSON(model.GetVirtualMachineInterfaceAllowedAddressPairs().GetAllowedAddressPair()),
		string(model.GetUUID()),
		bool(model.GetPortSecurityEnabled()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetIDPerms().GetUserVisible()),
		int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
		string(model.GetIDPerms().GetPermissions().GetOwner()),
		int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
		int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
		string(model.GetIDPerms().GetPermissions().GetGroup()),
		string(model.GetIDPerms().GetLastModified()),
		bool(model.GetIDPerms().GetEnable()),
		string(model.GetIDPerms().GetDescription()),
		string(model.GetIDPerms().GetCreator()),
		string(model.GetIDPerms().GetCreated()),
		common.MustJSON(model.GetFQName()),
		bool(model.GetEcmpHashingIncludeFields().GetSourcePort()),
		bool(model.GetEcmpHashingIncludeFields().GetSourceIP()),
		bool(model.GetEcmpHashingIncludeFields().GetIPProtocol()),
		bool(model.GetEcmpHashingIncludeFields().GetHashingConfigured()),
		bool(model.GetEcmpHashingIncludeFields().GetDestinationPort()),
		bool(model.GetEcmpHashingIncludeFields().GetDestinationIP()),
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtSecurityLoggingObjectRef, err := tx.Prepare(insertVirtualMachineInterfaceSecurityLoggingObjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing SecurityLoggingObjectRefs create statement failed")
	}
	defer stmtSecurityLoggingObjectRef.Close()
	for _, ref := range model.SecurityLoggingObjectRefs {

		_, err = stmtSecurityLoggingObjectRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs create failed")
		}
	}

	stmtInterfaceRouteTableRef, err := tx.Prepare(insertVirtualMachineInterfaceInterfaceRouteTableQuery)
	if err != nil {
		return errors.Wrap(err, "preparing InterfaceRouteTableRefs create statement failed")
	}
	defer stmtInterfaceRouteTableRef.Close()
	for _, ref := range model.InterfaceRouteTableRefs {

		_, err = stmtInterfaceRouteTableRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "InterfaceRouteTableRefs create failed")
		}
	}

	stmtSecurityGroupRef, err := tx.Prepare(insertVirtualMachineInterfaceSecurityGroupQuery)
	if err != nil {
		return errors.Wrap(err, "preparing SecurityGroupRefs create statement failed")
	}
	defer stmtSecurityGroupRef.Close()
	for _, ref := range model.SecurityGroupRefs {

		_, err = stmtSecurityGroupRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "SecurityGroupRefs create failed")
		}
	}

	stmtBridgeDomainRef, err := tx.Prepare(insertVirtualMachineInterfaceBridgeDomainQuery)
	if err != nil {
		return errors.Wrap(err, "preparing BridgeDomainRefs create statement failed")
	}
	defer stmtBridgeDomainRef.Close()
	for _, ref := range model.BridgeDomainRefs {

		if ref.Attr == nil {
			ref.Attr = &models.BridgeDomainMembershipType{}
		}

		_, err = stmtBridgeDomainRef.ExecContext(ctx, model.UUID, ref.UUID, int(ref.Attr.GetVlanTag()))
		if err != nil {
			return errors.Wrap(err, "BridgeDomainRefs create failed")
		}
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertVirtualMachineInterfaceVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = stmtVirtualMachineInterfaceRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	stmtBGPRouterRef, err := tx.Prepare(insertVirtualMachineInterfaceBGPRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing BGPRouterRefs create statement failed")
	}
	defer stmtBGPRouterRef.Close()
	for _, ref := range model.BGPRouterRefs {

		_, err = stmtBGPRouterRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "BGPRouterRefs create failed")
		}
	}

	stmtPortTupleRef, err := tx.Prepare(insertVirtualMachineInterfacePortTupleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PortTupleRefs create statement failed")
	}
	defer stmtPortTupleRef.Close()
	for _, ref := range model.PortTupleRefs {

		_, err = stmtPortTupleRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PortTupleRefs create failed")
		}
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertVirtualMachineInterfaceVirtualNetworkQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualNetworkRefs create statement failed")
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {

		_, err = stmtVirtualNetworkRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
		}
	}

	stmtServiceEndpointRef, err := tx.Prepare(insertVirtualMachineInterfaceServiceEndpointQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceEndpointRefs create statement failed")
	}
	defer stmtServiceEndpointRef.Close()
	for _, ref := range model.ServiceEndpointRefs {

		_, err = stmtServiceEndpointRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceEndpointRefs create failed")
		}
	}

	stmtVirtualMachineRef, err := tx.Prepare(insertVirtualMachineInterfaceVirtualMachineQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineRefs create statement failed")
	}
	defer stmtVirtualMachineRef.Close()
	for _, ref := range model.VirtualMachineRefs {

		_, err = stmtVirtualMachineRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineRefs create failed")
		}
	}

	stmtRoutingInstanceRef, err := tx.Prepare(insertVirtualMachineInterfaceRoutingInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing RoutingInstanceRefs create statement failed")
	}
	defer stmtRoutingInstanceRef.Close()
	for _, ref := range model.RoutingInstanceRefs {

		if ref.Attr == nil {
			ref.Attr = &models.PolicyBasedForwardingRuleType{}
		}

		_, err = stmtRoutingInstanceRef.ExecContext(ctx, model.UUID, ref.UUID, string(ref.Attr.GetServiceChainAddress()),
			string(ref.Attr.GetDSTMac()),
			string(ref.Attr.GetProtocol()),
			string(ref.Attr.GetIpv6ServiceChainAddress()),
			string(ref.Attr.GetDirection()),
			int(ref.Attr.GetMPLSLabel()),
			int(ref.Attr.GetVlanTag()),
			string(ref.Attr.GetSRCMac()))
		if err != nil {
			return errors.Wrap(err, "RoutingInstanceRefs create failed")
		}
	}

	stmtQosConfigRef, err := tx.Prepare(insertVirtualMachineInterfaceQosConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing QosConfigRefs create statement failed")
	}
	defer stmtQosConfigRef.Close()
	for _, ref := range model.QosConfigRefs {

		_, err = stmtQosConfigRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "QosConfigRefs create failed")
		}
	}

	stmtPhysicalInterfaceRef, err := tx.Prepare(insertVirtualMachineInterfacePhysicalInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PhysicalInterfaceRefs create statement failed")
	}
	defer stmtPhysicalInterfaceRef.Close()
	for _, ref := range model.PhysicalInterfaceRefs {

		_, err = stmtPhysicalInterfaceRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalInterfaceRefs create failed")
		}
	}

	stmtServiceHealthCheckRef, err := tx.Prepare(insertVirtualMachineInterfaceServiceHealthCheckQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceHealthCheckRefs create statement failed")
	}
	defer stmtServiceHealthCheckRef.Close()
	for _, ref := range model.ServiceHealthCheckRefs {

		_, err = stmtServiceHealthCheckRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceHealthCheckRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "virtual_machine_interface",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "virtual_machine_interface", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanVirtualMachineInterface(values map[string]interface{}) (*models.VirtualMachineInterface, error) {
	m := models.MakeVirtualMachineInterface()

	if value, ok := values["vrf_assign_rule"]; ok {

		json.Unmarshal(value.([]byte), &m.VRFAssignTable.VRFAssignRule)

	}

	if value, ok := values["vlan_tag_based_bridge_domain"]; ok {

		m.VlanTagBasedBridgeDomain = schema.InterfaceToBool(value)

	}

	if value, ok := values["sub_interface_vlan_tag"]; ok {

		m.VirtualMachineInterfaceProperties.SubInterfaceVlanTag = schema.InterfaceToInt64(value)

	}

	if value, ok := values["service_interface_type"]; ok {

		m.VirtualMachineInterfaceProperties.ServiceInterfaceType = schema.InterfaceToString(value)

	}

	if value, ok := values["local_preference"]; ok {

		m.VirtualMachineInterfaceProperties.LocalPreference = schema.InterfaceToInt64(value)

	}

	if value, ok := values["traffic_direction"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection = schema.InterfaceToString(value)

	}

	if value, ok := values["udp_port"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort = schema.InterfaceToInt64(value)

	}

	if value, ok := values["vtep_dst_mac_address"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["vtep_dst_ip_address"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["vni"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni = schema.InterfaceToInt64(value)

	}

	if value, ok := values["routing_instance"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance = schema.InterfaceToString(value)

	}

	if value, ok := values["nic_assisted_mirroring_vlan"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan = schema.InterfaceToInt64(value)

	}

	if value, ok := values["nic_assisted_mirroring"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring = schema.InterfaceToBool(value)

	}

	if value, ok := values["nh_mode"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode = schema.InterfaceToString(value)

	}

	if value, ok := values["juniper_header"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader = schema.InterfaceToBool(value)

	}

	if value, ok := values["encapsulation"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation = schema.InterfaceToString(value)

	}

	if value, ok := values["analyzer_name"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName = schema.InterfaceToString(value)

	}

	if value, ok := values["analyzer_mac_address"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["analyzer_ip_address"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["mac_address"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceMacAddresses.MacAddress)

	}

	if value, ok := values["route"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceHostRoutes.Route)

	}

	if value, ok := values["fat_flow_protocol"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol)

	}

	if value, ok := values["virtual_machine_interface_disable_policy"]; ok {

		m.VirtualMachineInterfaceDisablePolicy = schema.InterfaceToBool(value)

	}

	if value, ok := values["dhcp_option"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

	}

	if value, ok := values["virtual_machine_interface_device_owner"]; ok {

		m.VirtualMachineInterfaceDeviceOwner = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceBindings.KeyValuePair)

	}

	if value, ok := values["allowed_address_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["port_security_enabled"]; ok {

		m.PortSecurityEnabled = schema.InterfaceToBool(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = schema.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = schema.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = schema.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = schema.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = schema.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = schema.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = schema.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = schema.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = schema.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["source_port"]; ok {

		m.EcmpHashingIncludeFields.SourcePort = schema.InterfaceToBool(value)

	}

	if value, ok := values["source_ip"]; ok {

		m.EcmpHashingIncludeFields.SourceIP = schema.InterfaceToBool(value)

	}

	if value, ok := values["ip_protocol"]; ok {

		m.EcmpHashingIncludeFields.IPProtocol = schema.InterfaceToBool(value)

	}

	if value, ok := values["hashing_configured"]; ok {

		m.EcmpHashingIncludeFields.HashingConfigured = schema.InterfaceToBool(value)

	}

	if value, ok := values["destination_port"]; ok {

		m.EcmpHashingIncludeFields.DestinationPort = schema.InterfaceToBool(value)

	}

	if value, ok := values["destination_ip"]; ok {

		m.EcmpHashingIncludeFields.DestinationIP = schema.InterfaceToBool(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["annotations_key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_security_logging_object"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceSecurityLoggingObjectRef{}
			referenceModel.UUID = uuid
			m.SecurityLoggingObjectRefs = append(m.SecurityLoggingObjectRefs, referenceModel)

		}
	}

	if value, ok := values["ref_interface_route_table"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceInterfaceRouteTableRef{}
			referenceModel.UUID = uuid
			m.InterfaceRouteTableRefs = append(m.InterfaceRouteTableRefs, referenceModel)

		}
	}

	if value, ok := values["ref_security_group"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceSecurityGroupRef{}
			referenceModel.UUID = uuid
			m.SecurityGroupRefs = append(m.SecurityGroupRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bridge_domain"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceBridgeDomainRef{}
			referenceModel.UUID = uuid
			m.BridgeDomainRefs = append(m.BridgeDomainRefs, referenceModel)

			attr := models.MakeBridgeDomainMembershipType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bgp_router"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceBGPRouterRef{}
			referenceModel.UUID = uuid
			m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_port_tuple"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfacePortTupleRef{}
			referenceModel.UUID = uuid
			m.PortTupleRefs = append(m.PortTupleRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_network"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_endpoint"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceServiceEndpointRef{}
			referenceModel.UUID = uuid
			m.ServiceEndpointRefs = append(m.ServiceEndpointRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualMachineRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineRefs = append(m.VirtualMachineRefs, referenceModel)

		}
	}

	if value, ok := values["ref_routing_instance"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceRoutingInstanceRef{}
			referenceModel.UUID = uuid
			m.RoutingInstanceRefs = append(m.RoutingInstanceRefs, referenceModel)

			attr := models.MakePolicyBasedForwardingRuleType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_qos_config"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceQosConfigRef{}
			referenceModel.UUID = uuid
			m.QosConfigRefs = append(m.QosConfigRefs, referenceModel)

		}
	}

	if value, ok := values["ref_physical_interface"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfacePhysicalInterfaceRef{}
			referenceModel.UUID = uuid
			m.PhysicalInterfaceRefs = append(m.PhysicalInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_health_check"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceServiceHealthCheckRef{}
			referenceModel.UUID = uuid
			m.ServiceHealthCheckRefs = append(m.ServiceHealthCheckRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVirtualMachineInterface lists VirtualMachineInterface with list spec.
func ListVirtualMachineInterface(ctx context.Context, tx *sql.Tx, request *models.ListVirtualMachineInterfaceRequest) (response *models.ListVirtualMachineInterfaceResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "virtual_machine_interface"
	qb.Fields = VirtualMachineInterfaceFields
	qb.RefFields = VirtualMachineInterfaceRefFields
	qb.BackRefFields = VirtualMachineInterfaceBackRefFields
	result := []*models.VirtualMachineInterface{}

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = common.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}

	query := qb.BuildQuery()
	columns := qb.Columns
	values := qb.Values
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}

	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		m, err := scanVirtualMachineInterface(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListVirtualMachineInterfaceResponse{
		VirtualMachineInterfaces: result,
	}
	return response, nil
}

// UpdateVirtualMachineInterface updates a resource
func UpdateVirtualMachineInterface(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateVirtualMachineInterfaceRequest,
) error {
	//TODO
	return nil
}

// DeleteVirtualMachineInterface deletes a resource
func DeleteVirtualMachineInterface(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteVirtualMachineInterfaceRequest) error {
	deleteQuery := deleteVirtualMachineInterfaceQuery
	selectQuery := "select count(uuid) from virtual_machine_interface where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	auth := common.GetAuthCTX(ctx)
	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = common.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}
