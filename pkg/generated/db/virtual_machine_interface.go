package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualMachineInterfaceQuery = "insert into `virtual_machine_interface` (`key_value_pair`,`dhcp_option`,`virtual_machine_interface_fat_flow_protocols`,`port_security_enabled`,`virtual_machine_interface_device_owner`,`uuid`,`fq_name`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`mac_address`,`virtual_machine_interface_bindings`,`vlan_tag_based_bridge_domain`,`virtual_machine_interface_disable_policy`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`vrf_assign_rule`,`service_interface_type`,`sub_interface_vlan_tag`,`local_preference`,`traffic_direction`,`analyzer_mac_address`,`nic_assisted_mirroring`,`juniper_header`,`nh_mode`,`nic_assisted_mirroring_vlan`,`udp_port`,`analyzer_ip_address`,`analyzer_name`,`encapsulation`,`routing_instance`,`vni`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`route`,`allowed_address_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualMachineInterfaceQuery = "update `virtual_machine_interface` set `key_value_pair` = ?,`dhcp_option` = ?,`virtual_machine_interface_fat_flow_protocols` = ?,`port_security_enabled` = ?,`virtual_machine_interface_device_owner` = ?,`uuid` = ?,`fq_name` = ?,`display_name` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`mac_address` = ?,`virtual_machine_interface_bindings` = ?,`vlan_tag_based_bridge_domain` = ?,`virtual_machine_interface_disable_policy` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`vrf_assign_rule` = ?,`service_interface_type` = ?,`sub_interface_vlan_tag` = ?,`local_preference` = ?,`traffic_direction` = ?,`analyzer_mac_address` = ?,`nic_assisted_mirroring` = ?,`juniper_header` = ?,`nh_mode` = ?,`nic_assisted_mirroring_vlan` = ?,`udp_port` = ?,`analyzer_ip_address` = ?,`analyzer_name` = ?,`encapsulation` = ?,`routing_instance` = ?,`vni` = ?,`vtep_dst_ip_address` = ?,`vtep_dst_mac_address` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`route` = ?,`allowed_address_pair` = ?;"
const deleteVirtualMachineInterfaceQuery = "delete from `virtual_machine_interface` where uuid = ?"

// VirtualMachineInterfaceFields is db columns for VirtualMachineInterface
var VirtualMachineInterfaceFields = []string{
	"key_value_pair",
	"dhcp_option",
	"virtual_machine_interface_fat_flow_protocols",
	"port_security_enabled",
	"virtual_machine_interface_device_owner",
	"uuid",
	"fq_name",
	"display_name",
	"share",
	"owner",
	"owner_access",
	"global_access",
	"mac_address",
	"virtual_machine_interface_bindings",
	"vlan_tag_based_bridge_domain",
	"virtual_machine_interface_disable_policy",
	"user_visible",
	"last_modified",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"permissions_owner",
	"enable",
	"description",
	"created",
	"creator",
	"vrf_assign_rule",
	"service_interface_type",
	"sub_interface_vlan_tag",
	"local_preference",
	"traffic_direction",
	"analyzer_mac_address",
	"nic_assisted_mirroring",
	"juniper_header",
	"nh_mode",
	"nic_assisted_mirroring_vlan",
	"udp_port",
	"analyzer_ip_address",
	"analyzer_name",
	"encapsulation",
	"routing_instance",
	"vni",
	"vtep_dst_ip_address",
	"vtep_dst_mac_address",
	"ip_protocol",
	"source_ip",
	"hashing_configured",
	"source_port",
	"destination_port",
	"destination_ip",
	"route",
	"allowed_address_pair",
}

// VirtualMachineInterfaceRefFields is db reference fields for VirtualMachineInterface
var VirtualMachineInterfaceRefFields = map[string][]string{

	"virtual_machine_interface": {
	// <utils.Schema Value>

	},

	"qos_config": {
	// <utils.Schema Value>

	},

	"service_health_check": {
	// <utils.Schema Value>

	},

	"virtual_network": {
	// <utils.Schema Value>

	},

	"virtual_machine": {
	// <utils.Schema Value>

	},

	"bgp_router": {
	// <utils.Schema Value>

	},

	"service_endpoint": {
	// <utils.Schema Value>

	},

	"interface_route_table": {
	// <utils.Schema Value>

	},

	"routing_instance": {
		// <utils.Schema Value>
		"protocol",
		"ipv6_service_chain_address",
		"direction",
		"mpls_label",
		"vlan_tag",
		"src_mac",
		"service_chain_address",
		"dst_mac",
	},

	"port_tuple": {
	// <utils.Schema Value>

	},

	"physical_interface": {
	// <utils.Schema Value>

	},

	"security_logging_object": {
	// <utils.Schema Value>

	},

	"security_group": {
	// <utils.Schema Value>

	},

	"bridge_domain": {
		// <utils.Schema Value>
		"vlan_tag",
	},
}

const insertVirtualMachineInterfaceSecurityLoggingObjectQuery = "insert into `ref_virtual_machine_interface_security_logging_object` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceSecurityGroupQuery = "insert into `ref_virtual_machine_interface_security_group` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceBridgeDomainQuery = "insert into `ref_virtual_machine_interface_bridge_domain` (`from`, `to` ,`vlan_tag`) values (?, ?,?);"

const insertVirtualMachineInterfaceVirtualMachineInterfaceQuery = "insert into `ref_virtual_machine_interface_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceQosConfigQuery = "insert into `ref_virtual_machine_interface_qos_config` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceServiceHealthCheckQuery = "insert into `ref_virtual_machine_interface_service_health_check` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualNetworkQuery = "insert into `ref_virtual_machine_interface_virtual_network` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualMachineQuery = "insert into `ref_virtual_machine_interface_virtual_machine` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceBGPRouterQuery = "insert into `ref_virtual_machine_interface_bgp_router` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceServiceEndpointQuery = "insert into `ref_virtual_machine_interface_service_endpoint` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceInterfaceRouteTableQuery = "insert into `ref_virtual_machine_interface_interface_route_table` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceRoutingInstanceQuery = "insert into `ref_virtual_machine_interface_routing_instance` (`from`, `to` ,`protocol`,`ipv6_service_chain_address`,`direction`,`mpls_label`,`vlan_tag`,`src_mac`,`service_chain_address`,`dst_mac`) values (?, ?,?,?,?,?,?,?,?,?);"

const insertVirtualMachineInterfacePortTupleQuery = "insert into `ref_virtual_machine_interface_port_tuple` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfacePhysicalInterfaceQuery = "insert into `ref_virtual_machine_interface_physical_interface` (`from`, `to` ) values (?, ?);"

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
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.VirtualMachineInterfaceDHCPOptionList.DHCPOption),
		utils.MustJSON(model.VirtualMachineInterfaceFatFlowProtocols),
		bool(model.PortSecurityEnabled),
		string(model.VirtualMachineInterfaceDeviceOwner),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.DisplayName),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.VirtualMachineInterfaceMacAddresses.MacAddress),
		utils.MustJSON(model.VirtualMachineInterfaceBindings),
		bool(model.VlanTagBasedBridgeDomain),
		bool(model.VirtualMachineInterfaceDisablePolicy),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		utils.MustJSON(model.VRFAssignTable.VRFAssignRule),
		string(model.VirtualMachineInterfaceProperties.ServiceInterfaceType),
		int(model.VirtualMachineInterfaceProperties.SubInterfaceVlanTag),
		int(model.VirtualMachineInterfaceProperties.LocalPreference),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		utils.MustJSON(model.VirtualMachineInterfaceHostRoutes.Route),
		utils.MustJSON(model.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair))
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

	stmtRoutingInstanceRef, err := tx.Prepare(insertVirtualMachineInterfaceRoutingInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing RoutingInstanceRefs create statement failed")
	}
	defer stmtRoutingInstanceRef.Close()
	for _, ref := range model.RoutingInstanceRefs {
		_, err = stmtRoutingInstanceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.Protocol),
			string(ref.Attr.Ipv6ServiceChainAddress),
			string(ref.Attr.Direction),
			int(ref.Attr.MPLSLabel),
			int(ref.Attr.VlanTag),
			string(ref.Attr.SRCMac),
			string(ref.Attr.ServiceChainAddress),
			string(ref.Attr.DSTMac))
		if err != nil {
			return errors.Wrap(err, "RoutingInstanceRefs create failed")
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

	stmtBridgeDomainRef, err := tx.Prepare(insertVirtualMachineInterfaceBridgeDomainQuery)
	if err != nil {
		return errors.Wrap(err, "preparing BridgeDomainRefs create statement failed")
	}
	defer stmtBridgeDomainRef.Close()
	for _, ref := range model.BridgeDomainRefs {
		_, err = stmtBridgeDomainRef.Exec(model.UUID, ref.UUID, int(ref.Attr.VlanTag))
		if err != nil {
			return errors.Wrap(err, "BridgeDomainRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanVirtualMachineInterface(values map[string]interface{}) (*models.VirtualMachineInterface, error) {
	m := models.MakeVirtualMachineInterface()

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["dhcp_option"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

	}

	if value, ok := values["virtual_machine_interface_fat_flow_protocols"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceFatFlowProtocols)

	}

	if value, ok := values["port_security_enabled"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.PortSecurityEnabled = castedValue

	}

	if value, ok := values["virtual_machine_interface_device_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceDeviceOwner = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["mac_address"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceMacAddresses.MacAddress)

	}

	if value, ok := values["virtual_machine_interface_bindings"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceBindings)

	}

	if value, ok := values["vlan_tag_based_bridge_domain"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.VlanTagBasedBridgeDomain = castedValue

	}

	if value, ok := values["virtual_machine_interface_disable_policy"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.VirtualMachineInterfaceDisablePolicy = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["vrf_assign_rule"]; ok {

		json.Unmarshal(value.([]byte), &m.VRFAssignTable.VRFAssignRule)

	}

	if value, ok := values["service_interface_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.ServiceInterfaceType = models.ServiceInterfaceType(castedValue)

	}

	if value, ok := values["sub_interface_vlan_tag"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.SubInterfaceVlanTag = castedValue

	}

	if value, ok := values["local_preference"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.LocalPreference = castedValue

	}

	if value, ok := values["traffic_direction"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection = models.TrafficDirectionType(castedValue)

	}

	if value, ok := values["analyzer_mac_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress = castedValue

	}

	if value, ok := values["nic_assisted_mirroring"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring = castedValue

	}

	if value, ok := values["juniper_header"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader = castedValue

	}

	if value, ok := values["nh_mode"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode = models.NHModeType(castedValue)

	}

	if value, ok := values["nic_assisted_mirroring_vlan"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan = models.VlanIdType(castedValue)

	}

	if value, ok := values["udp_port"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort = castedValue

	}

	if value, ok := values["analyzer_ip_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress = castedValue

	}

	if value, ok := values["analyzer_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName = castedValue

	}

	if value, ok := values["encapsulation"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation = castedValue

	}

	if value, ok := values["routing_instance"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance = castedValue

	}

	if value, ok := values["vni"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni = models.VxlanNetworkIdentifierType(castedValue)

	}

	if value, ok := values["vtep_dst_ip_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress = castedValue

	}

	if value, ok := values["vtep_dst_mac_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress = castedValue

	}

	if value, ok := values["ip_protocol"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.IPProtocol = castedValue

	}

	if value, ok := values["source_ip"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourceIP = castedValue

	}

	if value, ok := values["hashing_configured"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.HashingConfigured = castedValue

	}

	if value, ok := values["source_port"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourcePort = castedValue

	}

	if value, ok := values["destination_port"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationPort = castedValue

	}

	if value, ok := values["destination_ip"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationIP = castedValue

	}

	if value, ok := values["route"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceHostRoutes.Route)

	}

	if value, ok := values["allowed_address_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

	}

	if value, ok := values["ref_routing_instance"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceRoutingInstanceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.RoutingInstanceRefs = append(m.RoutingInstanceRefs, referenceModel)

			attr := models.MakePolicyBasedForwardingRuleType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_port_tuple"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfacePortTupleRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.PortTupleRefs = append(m.PortTupleRefs, referenceModel)

		}
	}

	if value, ok := values["ref_physical_interface"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfacePhysicalInterfaceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.PhysicalInterfaceRefs = append(m.PhysicalInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_interface_route_table"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceInterfaceRouteTableRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.InterfaceRouteTableRefs = append(m.InterfaceRouteTableRefs, referenceModel)

		}
	}

	if value, ok := values["ref_security_group"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceSecurityGroupRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.SecurityGroupRefs = append(m.SecurityGroupRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bridge_domain"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceBridgeDomainRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.BridgeDomainRefs = append(m.BridgeDomainRefs, referenceModel)

			attr := models.MakeBridgeDomainMembershipType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_security_logging_object"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceSecurityLoggingObjectRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.SecurityLoggingObjectRefs = append(m.SecurityLoggingObjectRefs, referenceModel)

		}
	}

	if value, ok := values["ref_qos_config"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceQosConfigRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.QosConfigRefs = append(m.QosConfigRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_health_check"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceServiceHealthCheckRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.ServiceHealthCheckRefs = append(m.ServiceHealthCheckRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_network"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceVirtualNetworkRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceVirtualMachineInterfaceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bgp_router"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceBGPRouterRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_endpoint"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceServiceEndpointRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.ServiceEndpointRefs = append(m.ServiceEndpointRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualMachineInterfaceVirtualMachineRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualMachineRefs = append(m.VirtualMachineRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVirtualMachineInterface lists VirtualMachineInterface with list spec.
func ListVirtualMachineInterface(tx *sql.Tx, spec *db.ListSpec) ([]*models.VirtualMachineInterface, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_machine_interface"
	spec.Fields = VirtualMachineInterfaceFields
	spec.RefFields = VirtualMachineInterfaceRefFields
	result := models.MakeVirtualMachineInterfaceSlice()
	query, columns, values := db.BuildListQuery(spec)
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
	list, err := ListVirtualMachineInterface(tx, &db.ListSpec{
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
