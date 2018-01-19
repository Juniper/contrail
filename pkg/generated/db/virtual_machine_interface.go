package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
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

	"qos_config": {
	// <common.Schema Value>

	},

	"virtual_network": {
	// <common.Schema Value>

	},

	"service_endpoint": {
	// <common.Schema Value>

	},

	"bgp_router": {
	// <common.Schema Value>

	},

	"security_logging_object": {
	// <common.Schema Value>

	},

	"interface_route_table": {
	// <common.Schema Value>

	},

	"port_tuple": {
	// <common.Schema Value>

	},

	"virtual_machine_interface": {
	// <common.Schema Value>

	},

	"virtual_machine": {
	// <common.Schema Value>

	},

	"routing_instance": {
		// <common.Schema Value>
		"dst_mac",
		"protocol",
		"ipv6_service_chain_address",
		"direction",
		"mpls_label",
		"vlan_tag",
		"src_mac",
		"service_chain_address",
	},

	"physical_interface": {
	// <common.Schema Value>

	},

	"service_health_check": {
	// <common.Schema Value>

	},

	"security_group": {
	// <common.Schema Value>

	},

	"bridge_domain": {
		// <common.Schema Value>
		"vlan_tag",
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

const insertVirtualMachineInterfacePortTupleQuery = "insert into `ref_virtual_machine_interface_port_tuple` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceServiceEndpointQuery = "insert into `ref_virtual_machine_interface_service_endpoint` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceBGPRouterQuery = "insert into `ref_virtual_machine_interface_bgp_router` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceSecurityLoggingObjectQuery = "insert into `ref_virtual_machine_interface_security_logging_object` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceInterfaceRouteTableQuery = "insert into `ref_virtual_machine_interface_interface_route_table` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfacePhysicalInterfaceQuery = "insert into `ref_virtual_machine_interface_physical_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualMachineInterfaceQuery = "insert into `ref_virtual_machine_interface_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualMachineQuery = "insert into `ref_virtual_machine_interface_virtual_machine` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceRoutingInstanceQuery = "insert into `ref_virtual_machine_interface_routing_instance` (`from`, `to` ,`dst_mac`,`protocol`,`ipv6_service_chain_address`,`direction`,`mpls_label`,`vlan_tag`,`src_mac`,`service_chain_address`) values (?, ?,?,?,?,?,?,?,?,?);"

const insertVirtualMachineInterfaceServiceHealthCheckQuery = "insert into `ref_virtual_machine_interface_service_health_check` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceSecurityGroupQuery = "insert into `ref_virtual_machine_interface_security_group` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceBridgeDomainQuery = "insert into `ref_virtual_machine_interface_bridge_domain` (`from`, `to` ,`vlan_tag`) values (?, ?,?);"

const insertVirtualMachineInterfaceQosConfigQuery = "insert into `ref_virtual_machine_interface_qos_config` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualNetworkQuery = "insert into `ref_virtual_machine_interface_virtual_network` (`from`, `to` ) values (?, ?);"

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
	_, err = stmt.Exec(common.MustJSON(model.VRFAssignTable.VRFAssignRule),
		bool(model.VlanTagBasedBridgeDomain),
		int(model.VirtualMachineInterfaceProperties.SubInterfaceVlanTag),
		string(model.VirtualMachineInterfaceProperties.ServiceInterfaceType),
		int(model.VirtualMachineInterfaceProperties.LocalPreference),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress),
		common.MustJSON(model.VirtualMachineInterfaceMacAddresses.MacAddress),
		common.MustJSON(model.VirtualMachineInterfaceHostRoutes.Route),
		common.MustJSON(model.VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol),
		bool(model.VirtualMachineInterfaceDisablePolicy),
		common.MustJSON(model.VirtualMachineInterfaceDHCPOptionList.DHCPOption),
		string(model.VirtualMachineInterfaceDeviceOwner),
		common.MustJSON(model.VirtualMachineInterfaceBindings.KeyValuePair),
		common.MustJSON(model.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair),
		string(model.UUID),
		bool(model.PortSecurityEnabled),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		common.MustJSON(model.FQName),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
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

	stmtRoutingInstanceRef, err := tx.Prepare(insertVirtualMachineInterfaceRoutingInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing RoutingInstanceRefs create statement failed")
	}
	defer stmtRoutingInstanceRef.Close()
	for _, ref := range model.RoutingInstanceRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakePolicyBasedForwardingRuleType()
		}

		_, err = stmtRoutingInstanceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.DSTMac),
			string(ref.Attr.Protocol),
			string(ref.Attr.Ipv6ServiceChainAddress),
			string(ref.Attr.Direction),
			int(ref.Attr.MPLSLabel),
			int(ref.Attr.VlanTag),
			string(ref.Attr.SRCMac),
			string(ref.Attr.ServiceChainAddress))
		if err != nil {
			return errors.Wrap(err, "RoutingInstanceRefs create failed")
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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "virtual_machine_interface",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "virtual_machine_interface", model.UUID, model.Perms2.Share)
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

		castedValue := common.InterfaceToBool(value)

		m.VlanTagBasedBridgeDomain = castedValue

	}

	if value, ok := values["sub_interface_vlan_tag"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.SubInterfaceVlanTag = castedValue

	}

	if value, ok := values["service_interface_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.ServiceInterfaceType = models.ServiceInterfaceType(castedValue)

	}

	if value, ok := values["local_preference"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.LocalPreference = castedValue

	}

	if value, ok := values["traffic_direction"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection = models.TrafficDirectionType(castedValue)

	}

	if value, ok := values["udp_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort = castedValue

	}

	if value, ok := values["vtep_dst_mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress = castedValue

	}

	if value, ok := values["vtep_dst_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress = castedValue

	}

	if value, ok := values["vni"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni = models.VxlanNetworkIdentifierType(castedValue)

	}

	if value, ok := values["routing_instance"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance = castedValue

	}

	if value, ok := values["nic_assisted_mirroring_vlan"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan = models.VlanIdType(castedValue)

	}

	if value, ok := values["nic_assisted_mirroring"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring = castedValue

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

	if value, ok := values["analyzer_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName = castedValue

	}

	if value, ok := values["analyzer_mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress = castedValue

	}

	if value, ok := values["analyzer_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress = castedValue

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

		castedValue := common.InterfaceToBool(value)

		m.VirtualMachineInterfaceDisablePolicy = castedValue

	}

	if value, ok := values["dhcp_option"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

	}

	if value, ok := values["virtual_machine_interface_device_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualMachineInterfaceDeviceOwner = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceBindings.KeyValuePair)

	}

	if value, ok := values["allowed_address_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["port_security_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PortSecurityEnabled = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["source_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourcePort = castedValue

	}

	if value, ok := values["source_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourceIP = castedValue

	}

	if value, ok := values["ip_protocol"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.IPProtocol = castedValue

	}

	if value, ok := values["hashing_configured"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.HashingConfigured = castedValue

	}

	if value, ok := values["destination_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationPort = castedValue

	}

	if value, ok := values["destination_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationIP = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["annotations_key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceBGPRouterRef{}
			referenceModel.UUID = uuid
			m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)

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
			uuid := common.InterfaceToString(referenceMap["to"])
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
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceInterfaceRouteTableRef{}
			referenceModel.UUID = uuid
			m.InterfaceRouteTableRefs = append(m.InterfaceRouteTableRefs, referenceModel)

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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfacePortTupleRef{}
			referenceModel.UUID = uuid
			m.PortTupleRefs = append(m.PortTupleRefs, referenceModel)

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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceServiceEndpointRef{}
			referenceModel.UUID = uuid
			m.ServiceEndpointRefs = append(m.ServiceEndpointRefs, referenceModel)

		}
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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
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
			uuid := common.InterfaceToString(referenceMap["to"])
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
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
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

	if value, ok := values["ref_physical_interface"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
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
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceServiceHealthCheckRef{}
			referenceModel.UUID = uuid
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
			uuid := common.InterfaceToString(referenceMap["to"])
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
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
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

	if value, ok := values["ref_qos_config"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceQosConfigRef{}
			referenceModel.UUID = uuid
			m.QosConfigRefs = append(m.QosConfigRefs, referenceModel)

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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualNetworkRef{}
			referenceModel.UUID = uuid
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
	spec.BackRefFields = VirtualMachineInterfaceBackRefFields
	result := models.MakeVirtualMachineInterfaceSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
	}

	query := spec.BuildQuery()
	columns := spec.Columns
	values := spec.Values
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
		m, err := scanVirtualMachineInterface(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateVirtualMachineInterface updates a resource
func UpdateVirtualMachineInterface(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateVirtualMachineInterfaceQuery = "update `virtual_machine_interface` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".VRFAssignTable.VRFAssignRule", "."); ok {
		updateVirtualMachineInterfaceQuery += "`vrf_assign_rule` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VlanTagBasedBridgeDomain", "."); ok {
		updateVirtualMachineInterfaceQuery += "`vlan_tag_based_bridge_domain` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.SubInterfaceVlanTag", "."); ok {
		updateVirtualMachineInterfaceQuery += "`sub_interface_vlan_tag` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.ServiceInterfaceType", "."); ok {
		updateVirtualMachineInterfaceQuery += "`service_interface_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.LocalPreference", "."); ok {
		updateVirtualMachineInterfaceQuery += "`local_preference` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection", "."); ok {
		updateVirtualMachineInterfaceQuery += "`traffic_direction` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort", "."); ok {
		updateVirtualMachineInterfaceQuery += "`udp_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress", "."); ok {
		updateVirtualMachineInterfaceQuery += "`vtep_dst_mac_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress", "."); ok {
		updateVirtualMachineInterfaceQuery += "`vtep_dst_ip_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni", "."); ok {
		updateVirtualMachineInterfaceQuery += "`vni` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance", "."); ok {
		updateVirtualMachineInterfaceQuery += "`routing_instance` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan", "."); ok {
		updateVirtualMachineInterfaceQuery += "`nic_assisted_mirroring_vlan` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring", "."); ok {
		updateVirtualMachineInterfaceQuery += "`nic_assisted_mirroring` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode", "."); ok {
		updateVirtualMachineInterfaceQuery += "`nh_mode` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader", "."); ok {
		updateVirtualMachineInterfaceQuery += "`juniper_header` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation", "."); ok {
		updateVirtualMachineInterfaceQuery += "`encapsulation` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName", "."); ok {
		updateVirtualMachineInterfaceQuery += "`analyzer_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress", "."); ok {
		updateVirtualMachineInterfaceQuery += "`analyzer_mac_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress", "."); ok {
		updateVirtualMachineInterfaceQuery += "`analyzer_ip_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceMacAddresses.MacAddress", "."); ok {
		updateVirtualMachineInterfaceQuery += "`mac_address` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceHostRoutes.Route", "."); ok {
		updateVirtualMachineInterfaceQuery += "`route` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol", "."); ok {
		updateVirtualMachineInterfaceQuery += "`fat_flow_protocol` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceDisablePolicy", "."); ok {
		updateVirtualMachineInterfaceQuery += "`virtual_machine_interface_disable_policy` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceDHCPOptionList.DHCPOption", "."); ok {
		updateVirtualMachineInterfaceQuery += "`dhcp_option` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceDeviceOwner", "."); ok {
		updateVirtualMachineInterfaceQuery += "`virtual_machine_interface_device_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceBindings.KeyValuePair", "."); ok {
		updateVirtualMachineInterfaceQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair", "."); ok {
		updateVirtualMachineInterfaceQuery += "`allowed_address_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateVirtualMachineInterfaceQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PortSecurityEnabled", "."); ok {
		updateVirtualMachineInterfaceQuery += "`port_security_enabled` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateVirtualMachineInterfaceQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateVirtualMachineInterfaceQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateVirtualMachineInterfaceQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateVirtualMachineInterfaceQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateVirtualMachineInterfaceQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateVirtualMachineInterfaceQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateVirtualMachineInterfaceQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateVirtualMachineInterfaceQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateVirtualMachineInterfaceQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateVirtualMachineInterfaceQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateVirtualMachineInterfaceQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateVirtualMachineInterfaceQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateVirtualMachineInterfaceQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateVirtualMachineInterfaceQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateVirtualMachineInterfaceQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateVirtualMachineInterfaceQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateVirtualMachineInterfaceQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateVirtualMachineInterfaceQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.SourcePort", "."); ok {
		updateVirtualMachineInterfaceQuery += "`source_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.SourceIP", "."); ok {
		updateVirtualMachineInterfaceQuery += "`source_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.IPProtocol", "."); ok {
		updateVirtualMachineInterfaceQuery += "`ip_protocol` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.HashingConfigured", "."); ok {
		updateVirtualMachineInterfaceQuery += "`hashing_configured` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.DestinationPort", "."); ok {
		updateVirtualMachineInterfaceQuery += "`destination_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.DestinationIP", "."); ok {
		updateVirtualMachineInterfaceQuery += "`destination_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateVirtualMachineInterfaceQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateVirtualMachineInterfaceQuery += "`annotations_key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineInterfaceQuery += ","
	}

	updateVirtualMachineInterfaceQuery =
		updateVirtualMachineInterfaceQuery[:len(updateVirtualMachineInterfaceQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateVirtualMachineInterfaceQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteVirtualMachineInterface deletes a resource
func DeleteVirtualMachineInterface(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteVirtualMachineInterfaceQuery
	selectQuery := "select count(uuid) from virtual_machine_interface where uuid = ?"
	var err error
	var count int

	if auth.IsAdmin() {
		row := tx.QueryRow(selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRow(selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid, auth.ProjectID())
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
