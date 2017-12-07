package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualMachineInterfaceQuery = "insert into `virtual_machine_interface` (`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`dhcp_option`,`virtual_machine_interface_device_owner`,`fq_name`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`allowed_address_pair`,`virtual_machine_interface_fat_flow_protocols`,`vrf_assign_rule`,`service_interface_type`,`sub_interface_vlan_tag`,`local_preference`,`traffic_direction`,`analyzer_ip_address`,`nh_mode`,`juniper_header`,`encapsulation`,`nic_assisted_mirroring`,`vni`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`analyzer_name`,`udp_port`,`nic_assisted_mirroring_vlan`,`routing_instance`,`analyzer_mac_address`,`uuid`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`display_name`,`route`,`mac_address`,`virtual_machine_interface_bindings`,`virtual_machine_interface_disable_policy`,`vlan_tag_based_bridge_domain`,`port_security_enabled`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualMachineInterfaceQuery = "update `virtual_machine_interface` set `ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`dhcp_option` = ?,`virtual_machine_interface_device_owner` = ?,`fq_name` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`allowed_address_pair` = ?,`virtual_machine_interface_fat_flow_protocols` = ?,`vrf_assign_rule` = ?,`service_interface_type` = ?,`sub_interface_vlan_tag` = ?,`local_preference` = ?,`traffic_direction` = ?,`analyzer_ip_address` = ?,`nh_mode` = ?,`juniper_header` = ?,`encapsulation` = ?,`nic_assisted_mirroring` = ?,`vni` = ?,`vtep_dst_ip_address` = ?,`vtep_dst_mac_address` = ?,`analyzer_name` = ?,`udp_port` = ?,`nic_assisted_mirroring_vlan` = ?,`routing_instance` = ?,`analyzer_mac_address` = ?,`uuid` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`display_name` = ?,`route` = ?,`mac_address` = ?,`virtual_machine_interface_bindings` = ?,`virtual_machine_interface_disable_policy` = ?,`vlan_tag_based_bridge_domain` = ?,`port_security_enabled` = ?;"
const deleteVirtualMachineInterfaceQuery = "delete from `virtual_machine_interface` where uuid = ?"

// VirtualMachineInterfaceFields is db columns for VirtualMachineInterface
var VirtualMachineInterfaceFields = []string{
	"ip_protocol",
	"source_ip",
	"hashing_configured",
	"source_port",
	"destination_port",
	"destination_ip",
	"dhcp_option",
	"virtual_machine_interface_device_owner",
	"fq_name",
	"other_access",
	"group",
	"group_access",
	"owner",
	"owner_access",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"allowed_address_pair",
	"virtual_machine_interface_fat_flow_protocols",
	"vrf_assign_rule",
	"service_interface_type",
	"sub_interface_vlan_tag",
	"local_preference",
	"traffic_direction",
	"analyzer_ip_address",
	"nh_mode",
	"juniper_header",
	"encapsulation",
	"nic_assisted_mirroring",
	"vni",
	"vtep_dst_ip_address",
	"vtep_dst_mac_address",
	"analyzer_name",
	"udp_port",
	"nic_assisted_mirroring_vlan",
	"routing_instance",
	"analyzer_mac_address",
	"uuid",
	"key_value_pair",
	"global_access",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"display_name",
	"route",
	"mac_address",
	"virtual_machine_interface_bindings",
	"virtual_machine_interface_disable_policy",
	"vlan_tag_based_bridge_domain",
	"port_security_enabled",
}

// VirtualMachineInterfaceRefFields is db reference fields for VirtualMachineInterface
var VirtualMachineInterfaceRefFields = map[string][]string{

	"bgp_router": {
	// <common.Schema Value>

	},

	"interface_route_table": {
	// <common.Schema Value>

	},

	"qos_config": {
	// <common.Schema Value>

	},

	"physical_interface": {
	// <common.Schema Value>

	},

	"bridge_domain": {
		// <common.Schema Value>
		"vlan_tag",
	},

	"routing_instance": {
		// <common.Schema Value>
		"vlan_tag",
		"src_mac",
		"service_chain_address",
		"dst_mac",
		"protocol",
		"ipv6_service_chain_address",
		"direction",
		"mpls_label",
	},

	"service_health_check": {
	// <common.Schema Value>

	},

	"security_group": {
	// <common.Schema Value>

	},

	"security_logging_object": {
	// <common.Schema Value>

	},

	"port_tuple": {
	// <common.Schema Value>

	},

	"virtual_network": {
	// <common.Schema Value>

	},

	"virtual_machine_interface": {
	// <common.Schema Value>

	},

	"virtual_machine": {
	// <common.Schema Value>

	},

	"service_endpoint": {
	// <common.Schema Value>

	},
}

const insertVirtualMachineInterfaceQosConfigQuery = "insert into `ref_virtual_machine_interface_qos_config` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfacePhysicalInterfaceQuery = "insert into `ref_virtual_machine_interface_physical_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceBridgeDomainQuery = "insert into `ref_virtual_machine_interface_bridge_domain` (`from`, `to` ,`vlan_tag`) values (?, ?,?);"

const insertVirtualMachineInterfaceBGPRouterQuery = "insert into `ref_virtual_machine_interface_bgp_router` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceInterfaceRouteTableQuery = "insert into `ref_virtual_machine_interface_interface_route_table` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceSecurityGroupQuery = "insert into `ref_virtual_machine_interface_security_group` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceRoutingInstanceQuery = "insert into `ref_virtual_machine_interface_routing_instance` (`from`, `to` ,`vlan_tag`,`src_mac`,`service_chain_address`,`dst_mac`,`protocol`,`ipv6_service_chain_address`,`direction`,`mpls_label`) values (?, ?,?,?,?,?,?,?,?,?);"

const insertVirtualMachineInterfaceServiceHealthCheckQuery = "insert into `ref_virtual_machine_interface_service_health_check` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualNetworkQuery = "insert into `ref_virtual_machine_interface_virtual_network` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceSecurityLoggingObjectQuery = "insert into `ref_virtual_machine_interface_security_logging_object` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfacePortTupleQuery = "insert into `ref_virtual_machine_interface_port_tuple` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceServiceEndpointQuery = "insert into `ref_virtual_machine_interface_service_endpoint` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualMachineInterfaceQuery = "insert into `ref_virtual_machine_interface_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualMachineQuery = "insert into `ref_virtual_machine_interface_virtual_machine` (`from`, `to` ) values (?, ?);"

// CreateVirtualMachineInterface inserts VirtualMachineInterface to DB
func CreateVirtualMachineInterface(tx *sql.Tx, model *models.VirtualMachineInterface) error {
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
	_, err = stmt.Exec(bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		common.MustJSON(model.VirtualMachineInterfaceDHCPOptionList.DHCPOption),
		string(model.VirtualMachineInterfaceDeviceOwner),
		common.MustJSON(model.FQName),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		common.MustJSON(model.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair),
		common.MustJSON(model.VirtualMachineInterfaceFatFlowProtocols),
		common.MustJSON(model.VRFAssignTable.VRFAssignRule),
		string(model.VirtualMachineInterfaceProperties.ServiceInterfaceType),
		int(model.VirtualMachineInterfaceProperties.SubInterfaceVlanTag),
		int(model.VirtualMachineInterfaceProperties.LocalPreference),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress),
		string(model.UUID),
		common.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.DisplayName),
		common.MustJSON(model.VirtualMachineInterfaceHostRoutes.Route),
		common.MustJSON(model.VirtualMachineInterfaceMacAddresses.MacAddress),
		common.MustJSON(model.VirtualMachineInterfaceBindings),
		bool(model.VirtualMachineInterfaceDisablePolicy),
		bool(model.VlanTagBasedBridgeDomain),
		bool(model.PortSecurityEnabled))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertVirtualMachineInterfaceVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	stmtVirtualMachineRef, err := tx.Prepare(insertVirtualMachineInterfaceVirtualMachineQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineRefs create statement failed")
	}
	defer stmtVirtualMachineRef.Close()
	for _, ref := range model.VirtualMachineRefs {

		_, err = stmtVirtualMachineRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineRefs create failed")
		}
	}

	stmtServiceEndpointRef, err := tx.Prepare(insertVirtualMachineInterfaceServiceEndpointQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceEndpointRefs create statement failed")
	}
	defer stmtServiceEndpointRef.Close()
	for _, ref := range model.ServiceEndpointRefs {

		_, err = stmtServiceEndpointRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceEndpointRefs create failed")
		}
	}

	stmtBGPRouterRef, err := tx.Prepare(insertVirtualMachineInterfaceBGPRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing BGPRouterRefs create statement failed")
	}
	defer stmtBGPRouterRef.Close()
	for _, ref := range model.BGPRouterRefs {

		_, err = stmtBGPRouterRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "BGPRouterRefs create failed")
		}
	}

	stmtInterfaceRouteTableRef, err := tx.Prepare(insertVirtualMachineInterfaceInterfaceRouteTableQuery)
	if err != nil {
		return errors.Wrap(err, "preparing InterfaceRouteTableRefs create statement failed")
	}
	defer stmtInterfaceRouteTableRef.Close()
	for _, ref := range model.InterfaceRouteTableRefs {

		_, err = stmtInterfaceRouteTableRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "InterfaceRouteTableRefs create failed")
		}
	}

	stmtQosConfigRef, err := tx.Prepare(insertVirtualMachineInterfaceQosConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing QosConfigRefs create statement failed")
	}
	defer stmtQosConfigRef.Close()
	for _, ref := range model.QosConfigRefs {

		_, err = stmtQosConfigRef.Exec(model.UUID, ref.UUID)
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

		_, err = stmtPhysicalInterfaceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalInterfaceRefs create failed")
		}
	}

	stmtBridgeDomainRef, err := tx.Prepare(insertVirtualMachineInterfaceBridgeDomainQuery)
	if err != nil {
		return errors.Wrap(err, "preparing BridgeDomainRefs create statement failed")
	}
	defer stmtBridgeDomainRef.Close()
	for _, ref := range model.BridgeDomainRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeBridgeDomainMembershipType()
		}

		_, err = stmtBridgeDomainRef.Exec(model.UUID, ref.UUID, int(ref.Attr.VlanTag))
		if err != nil {
			return errors.Wrap(err, "BridgeDomainRefs create failed")
		}
	}

	stmtRoutingInstanceRef, err := tx.Prepare(insertVirtualMachineInterfaceRoutingInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing RoutingInstanceRefs create statement failed")
	}
	defer stmtRoutingInstanceRef.Close()
	for _, ref := range model.RoutingInstanceRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakePolicyBasedForwardingRuleType()
		}

		_, err = stmtRoutingInstanceRef.Exec(model.UUID, ref.UUID, int(ref.Attr.VlanTag),
			string(ref.Attr.SRCMac),
			string(ref.Attr.ServiceChainAddress),
			string(ref.Attr.DSTMac),
			string(ref.Attr.Protocol),
			string(ref.Attr.Ipv6ServiceChainAddress),
			string(ref.Attr.Direction),
			int(ref.Attr.MPLSLabel))
		if err != nil {
			return errors.Wrap(err, "RoutingInstanceRefs create failed")
		}
	}

	stmtServiceHealthCheckRef, err := tx.Prepare(insertVirtualMachineInterfaceServiceHealthCheckQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceHealthCheckRefs create statement failed")
	}
	defer stmtServiceHealthCheckRef.Close()
	for _, ref := range model.ServiceHealthCheckRefs {

		_, err = stmtServiceHealthCheckRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceHealthCheckRefs create failed")
		}
	}

	stmtSecurityGroupRef, err := tx.Prepare(insertVirtualMachineInterfaceSecurityGroupQuery)
	if err != nil {
		return errors.Wrap(err, "preparing SecurityGroupRefs create statement failed")
	}
	defer stmtSecurityGroupRef.Close()
	for _, ref := range model.SecurityGroupRefs {

		_, err = stmtSecurityGroupRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "SecurityGroupRefs create failed")
		}
	}

	stmtSecurityLoggingObjectRef, err := tx.Prepare(insertVirtualMachineInterfaceSecurityLoggingObjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing SecurityLoggingObjectRefs create statement failed")
	}
	defer stmtSecurityLoggingObjectRef.Close()
	for _, ref := range model.SecurityLoggingObjectRefs {

		_, err = stmtSecurityLoggingObjectRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs create failed")
		}
	}

	stmtPortTupleRef, err := tx.Prepare(insertVirtualMachineInterfacePortTupleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PortTupleRefs create statement failed")
	}
	defer stmtPortTupleRef.Close()
	for _, ref := range model.PortTupleRefs {

		_, err = stmtPortTupleRef.Exec(model.UUID, ref.UUID)
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

		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanVirtualMachineInterface(values map[string]interface{}) (*models.VirtualMachineInterface, error) {
	m := models.MakeVirtualMachineInterface()

	if value, ok := values["ip_protocol"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.IPProtocol = castedValue

	}

	if value, ok := values["source_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourceIP = castedValue

	}

	if value, ok := values["hashing_configured"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.HashingConfigured = castedValue

	}

	if value, ok := values["source_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourcePort = castedValue

	}

	if value, ok := values["destination_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationPort = castedValue

	}

	if value, ok := values["destination_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationIP = castedValue

	}

	if value, ok := values["dhcp_option"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

	}

	if value, ok := values["virtual_machine_interface_device_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceDeviceOwner = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["allowed_address_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

	}

	if value, ok := values["virtual_machine_interface_fat_flow_protocols"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceFatFlowProtocols)

	}

	if value, ok := values["vrf_assign_rule"]; ok {

		json.Unmarshal(value.([]byte), &m.VRFAssignTable.VRFAssignRule)

	}

	if value, ok := values["service_interface_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.ServiceInterfaceType = models.ServiceInterfaceType(castedValue)

	}

	if value, ok := values["sub_interface_vlan_tag"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.SubInterfaceVlanTag = castedValue

	}

	if value, ok := values["local_preference"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.LocalPreference = castedValue

	}

	if value, ok := values["traffic_direction"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection = models.TrafficDirectionType(castedValue)

	}

	if value, ok := values["analyzer_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress = castedValue

	}

	if value, ok := values["nh_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode = models.NHModeType(castedValue)

	}

	if value, ok := values["juniper_header"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader = castedValue

	}

	if value, ok := values["encapsulation"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation = castedValue

	}

	if value, ok := values["nic_assisted_mirroring"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring = castedValue

	}

	if value, ok := values["vni"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni = models.VxlanNetworkIdentifierType(castedValue)

	}

	if value, ok := values["vtep_dst_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress = castedValue

	}

	if value, ok := values["vtep_dst_mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress = castedValue

	}

	if value, ok := values["analyzer_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName = castedValue

	}

	if value, ok := values["udp_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort = castedValue

	}

	if value, ok := values["nic_assisted_mirroring_vlan"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan = models.VlanIdType(castedValue)

	}

	if value, ok := values["routing_instance"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance = castedValue

	}

	if value, ok := values["analyzer_mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["route"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceHostRoutes.Route)

	}

	if value, ok := values["mac_address"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceMacAddresses.MacAddress)

	}

	if value, ok := values["virtual_machine_interface_bindings"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceBindings)

	}

	if value, ok := values["virtual_machine_interface_disable_policy"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualMachineInterfaceDisablePolicy = castedValue

	}

	if value, ok := values["vlan_tag_based_bridge_domain"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VlanTagBasedBridgeDomain = castedValue

	}

	if value, ok := values["port_security_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PortSecurityEnabled = castedValue

	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualMachineInterfaceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualMachineRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualMachineRefs = append(m.VirtualMachineRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_endpoint"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceServiceEndpointRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceEndpointRefs = append(m.ServiceEndpointRefs, referenceModel)

		}
	}

	if value, ok := values["ref_physical_interface"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfacePhysicalInterfaceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.PhysicalInterfaceRefs = append(m.PhysicalInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bridge_domain"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceBridgeDomainRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.BridgeDomainRefs = append(m.BridgeDomainRefs, referenceModel)

			attr := models.MakeBridgeDomainMembershipType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_bgp_router"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceBGPRouterRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_interface_route_table"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceInterfaceRouteTableRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.InterfaceRouteTableRefs = append(m.InterfaceRouteTableRefs, referenceModel)

		}
	}

	if value, ok := values["ref_qos_config"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceQosConfigRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.QosConfigRefs = append(m.QosConfigRefs, referenceModel)

		}
	}

	if value, ok := values["ref_routing_instance"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceRoutingInstanceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.RoutingInstanceRefs = append(m.RoutingInstanceRefs, referenceModel)

			attr := models.MakePolicyBasedForwardingRuleType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_service_health_check"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceServiceHealthCheckRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceHealthCheckRefs = append(m.ServiceHealthCheckRefs, referenceModel)

		}
	}

	if value, ok := values["ref_security_group"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceSecurityGroupRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.SecurityGroupRefs = append(m.SecurityGroupRefs, referenceModel)

		}
	}

	if value, ok := values["ref_security_logging_object"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceSecurityLoggingObjectRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.SecurityLoggingObjectRefs = append(m.SecurityLoggingObjectRefs, referenceModel)

		}
	}

	if value, ok := values["ref_port_tuple"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfacePortTupleRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.PortTupleRefs = append(m.PortTupleRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_network"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualNetworkRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVirtualMachineInterface lists VirtualMachineInterface with list spec.
func ListVirtualMachineInterface(tx *sql.Tx, spec *common.ListSpec) ([]*models.VirtualMachineInterface, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_machine_interface"
	spec.Fields = VirtualMachineInterfaceFields
	spec.RefFields = VirtualMachineInterfaceRefFields
	result := models.MakeVirtualMachineInterfaceSlice()
	query, columns, values := common.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
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
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanVirtualMachineInterface(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowVirtualMachineInterface shows VirtualMachineInterface resource
func ShowVirtualMachineInterface(tx *sql.Tx, uuid string) (*models.VirtualMachineInterface, error) {
	list, err := ListVirtualMachineInterface(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateVirtualMachineInterface updates a resource
func UpdateVirtualMachineInterface(tx *sql.Tx, uuid string, model *models.VirtualMachineInterface) error {
	//TODO(nati) support update
	return nil
}

// DeleteVirtualMachineInterface deletes a resource
func DeleteVirtualMachineInterface(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
