package db

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"strings"
)

const insertVirtualMachineInterfaceQuery = "insert into `virtual_machine_interface` (`route`,`dhcp_option`,`virtual_machine_interface_device_owner`,`port_security_enabled`,`traffic_direction`,`analyzer_name`,`nic_assisted_mirroring_vlan`,`nic_assisted_mirroring`,`analyzer_ip_address`,`encapsulation`,`vni`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`juniper_header`,`udp_port`,`routing_instance`,`nh_mode`,`analyzer_mac_address`,`service_interface_type`,`sub_interface_vlan_tag`,`local_preference`,`key_value_pair`,`mac_address`,`virtual_machine_interface_bindings`,`virtual_machine_interface_disable_policy`,`display_name`,`virtual_machine_interface_fat_flow_protocols`,`vrf_assign_rule`,`fq_name`,`uuid`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`allowed_address_pair`,`vlan_tag_based_bridge_domain`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualMachineInterfaceQuery = "update `virtual_machine_interface` set `route` = ?,`dhcp_option` = ?,`virtual_machine_interface_device_owner` = ?,`port_security_enabled` = ?,`traffic_direction` = ?,`analyzer_name` = ?,`nic_assisted_mirroring_vlan` = ?,`nic_assisted_mirroring` = ?,`analyzer_ip_address` = ?,`encapsulation` = ?,`vni` = ?,`vtep_dst_ip_address` = ?,`vtep_dst_mac_address` = ?,`juniper_header` = ?,`udp_port` = ?,`routing_instance` = ?,`nh_mode` = ?,`analyzer_mac_address` = ?,`service_interface_type` = ?,`sub_interface_vlan_tag` = ?,`local_preference` = ?,`key_value_pair` = ?,`mac_address` = ?,`virtual_machine_interface_bindings` = ?,`virtual_machine_interface_disable_policy` = ?,`display_name` = ?,`virtual_machine_interface_fat_flow_protocols` = ?,`vrf_assign_rule` = ?,`fq_name` = ?,`uuid` = ?,`destination_ip` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`allowed_address_pair` = ?,`vlan_tag_based_bridge_domain` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteVirtualMachineInterfaceQuery = "delete from `virtual_machine_interface` where uuid = ?"
const listVirtualMachineInterfaceQuery = "select `virtual_machine_interface`.`route`,`virtual_machine_interface`.`dhcp_option`,`virtual_machine_interface`.`virtual_machine_interface_device_owner`,`virtual_machine_interface`.`port_security_enabled`,`virtual_machine_interface`.`traffic_direction`,`virtual_machine_interface`.`analyzer_name`,`virtual_machine_interface`.`nic_assisted_mirroring_vlan`,`virtual_machine_interface`.`nic_assisted_mirroring`,`virtual_machine_interface`.`analyzer_ip_address`,`virtual_machine_interface`.`encapsulation`,`virtual_machine_interface`.`vni`,`virtual_machine_interface`.`vtep_dst_ip_address`,`virtual_machine_interface`.`vtep_dst_mac_address`,`virtual_machine_interface`.`juniper_header`,`virtual_machine_interface`.`udp_port`,`virtual_machine_interface`.`routing_instance`,`virtual_machine_interface`.`nh_mode`,`virtual_machine_interface`.`analyzer_mac_address`,`virtual_machine_interface`.`service_interface_type`,`virtual_machine_interface`.`sub_interface_vlan_tag`,`virtual_machine_interface`.`local_preference`,`virtual_machine_interface`.`key_value_pair`,`virtual_machine_interface`.`mac_address`,`virtual_machine_interface`.`virtual_machine_interface_bindings`,`virtual_machine_interface`.`virtual_machine_interface_disable_policy`,`virtual_machine_interface`.`display_name`,`virtual_machine_interface`.`virtual_machine_interface_fat_flow_protocols`,`virtual_machine_interface`.`vrf_assign_rule`,`virtual_machine_interface`.`fq_name`,`virtual_machine_interface`.`uuid`,`virtual_machine_interface`.`destination_ip`,`virtual_machine_interface`.`ip_protocol`,`virtual_machine_interface`.`source_ip`,`virtual_machine_interface`.`hashing_configured`,`virtual_machine_interface`.`source_port`,`virtual_machine_interface`.`destination_port`,`virtual_machine_interface`.`allowed_address_pair`,`virtual_machine_interface`.`vlan_tag_based_bridge_domain`,`virtual_machine_interface`.`created`,`virtual_machine_interface`.`creator`,`virtual_machine_interface`.`user_visible`,`virtual_machine_interface`.`last_modified`,`virtual_machine_interface`.`owner`,`virtual_machine_interface`.`owner_access`,`virtual_machine_interface`.`other_access`,`virtual_machine_interface`.`group`,`virtual_machine_interface`.`group_access`,`virtual_machine_interface`.`enable`,`virtual_machine_interface`.`description`,`virtual_machine_interface`.`global_access`,`virtual_machine_interface`.`share`,`virtual_machine_interface`.`perms2_owner`,`virtual_machine_interface`.`perms2_owner_access` from `virtual_machine_interface`"
const showVirtualMachineInterfaceQuery = "select `virtual_machine_interface`.`route`,`virtual_machine_interface`.`dhcp_option`,`virtual_machine_interface`.`virtual_machine_interface_device_owner`,`virtual_machine_interface`.`port_security_enabled`,`virtual_machine_interface`.`traffic_direction`,`virtual_machine_interface`.`analyzer_name`,`virtual_machine_interface`.`nic_assisted_mirroring_vlan`,`virtual_machine_interface`.`nic_assisted_mirroring`,`virtual_machine_interface`.`analyzer_ip_address`,`virtual_machine_interface`.`encapsulation`,`virtual_machine_interface`.`vni`,`virtual_machine_interface`.`vtep_dst_ip_address`,`virtual_machine_interface`.`vtep_dst_mac_address`,`virtual_machine_interface`.`juniper_header`,`virtual_machine_interface`.`udp_port`,`virtual_machine_interface`.`routing_instance`,`virtual_machine_interface`.`nh_mode`,`virtual_machine_interface`.`analyzer_mac_address`,`virtual_machine_interface`.`service_interface_type`,`virtual_machine_interface`.`sub_interface_vlan_tag`,`virtual_machine_interface`.`local_preference`,`virtual_machine_interface`.`key_value_pair`,`virtual_machine_interface`.`mac_address`,`virtual_machine_interface`.`virtual_machine_interface_bindings`,`virtual_machine_interface`.`virtual_machine_interface_disable_policy`,`virtual_machine_interface`.`display_name`,`virtual_machine_interface`.`virtual_machine_interface_fat_flow_protocols`,`virtual_machine_interface`.`vrf_assign_rule`,`virtual_machine_interface`.`fq_name`,`virtual_machine_interface`.`uuid`,`virtual_machine_interface`.`destination_ip`,`virtual_machine_interface`.`ip_protocol`,`virtual_machine_interface`.`source_ip`,`virtual_machine_interface`.`hashing_configured`,`virtual_machine_interface`.`source_port`,`virtual_machine_interface`.`destination_port`,`virtual_machine_interface`.`allowed_address_pair`,`virtual_machine_interface`.`vlan_tag_based_bridge_domain`,`virtual_machine_interface`.`created`,`virtual_machine_interface`.`creator`,`virtual_machine_interface`.`user_visible`,`virtual_machine_interface`.`last_modified`,`virtual_machine_interface`.`owner`,`virtual_machine_interface`.`owner_access`,`virtual_machine_interface`.`other_access`,`virtual_machine_interface`.`group`,`virtual_machine_interface`.`group_access`,`virtual_machine_interface`.`enable`,`virtual_machine_interface`.`description`,`virtual_machine_interface`.`global_access`,`virtual_machine_interface`.`share`,`virtual_machine_interface`.`perms2_owner`,`virtual_machine_interface`.`perms2_owner_access` from `virtual_machine_interface` where uuid = ?"

const insertVirtualMachineInterfaceVirtualMachineQuery = "insert into `ref_virtual_machine_interface_virtual_machine` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceInterfaceRouteTableQuery = "insert into `ref_virtual_machine_interface_interface_route_table` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceRoutingInstanceQuery = "insert into `ref_virtual_machine_interface_routing_instance` (`from`, `to` ,`direction`,`mpls_label`,`vlan_tag`,`src_mac`,`service_chain_address`,`dst_mac`,`protocol`,`ipv6_service_chain_address`) values (?, ?,?,?,?,?,?,?,?,?);"

const insertVirtualMachineInterfaceBridgeDomainQuery = "insert into `ref_virtual_machine_interface_bridge_domain` (`from`, `to` ,`vlan_tag`) values (?, ?,?);"

const insertVirtualMachineInterfaceVirtualMachineInterfaceQuery = "insert into `ref_virtual_machine_interface_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfacePortTupleQuery = "insert into `ref_virtual_machine_interface_port_tuple` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceServiceHealthCheckQuery = "insert into `ref_virtual_machine_interface_service_health_check` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceSecurityLoggingObjectQuery = "insert into `ref_virtual_machine_interface_security_logging_object` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceQosConfigQuery = "insert into `ref_virtual_machine_interface_qos_config` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfacePhysicalInterfaceQuery = "insert into `ref_virtual_machine_interface_physical_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceSecurityGroupQuery = "insert into `ref_virtual_machine_interface_security_group` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceVirtualNetworkQuery = "insert into `ref_virtual_machine_interface_virtual_network` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceServiceEndpointQuery = "insert into `ref_virtual_machine_interface_service_endpoint` (`from`, `to` ) values (?, ?);"

const insertVirtualMachineInterfaceBGPRouterQuery = "insert into `ref_virtual_machine_interface_bgp_router` (`from`, `to` ) values (?, ?);"

func CreateVirtualMachineInterface(tx *sql.Tx, model *models.VirtualMachineInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.VirtualMachineInterfaceHostRoutes.Route),
		utils.MustJSON(model.VirtualMachineInterfaceDHCPOptionList.DHCPOption),
		string(model.VirtualMachineInterfaceDeviceOwner),
		bool(model.PortSecurityEnabled),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress),
		string(model.VirtualMachineInterfaceProperties.ServiceInterfaceType),
		int(model.VirtualMachineInterfaceProperties.SubInterfaceVlanTag),
		int(model.VirtualMachineInterfaceProperties.LocalPreference),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.VirtualMachineInterfaceMacAddresses.MacAddress),
		utils.MustJSON(model.VirtualMachineInterfaceBindings),
		bool(model.VirtualMachineInterfaceDisablePolicy),
		string(model.DisplayName),
		utils.MustJSON(model.VirtualMachineInterfaceFatFlowProtocols),
		utils.MustJSON(model.VRFAssignTable.VRFAssignRule),
		utils.MustJSON(model.FQName),
		string(model.UUID),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		utils.MustJSON(model.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair),
		bool(model.VlanTagBasedBridgeDomain),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess))

	stmtSecurityLoggingObjectRef, err := tx.Prepare(insertVirtualMachineInterfaceSecurityLoggingObjectQuery)
	if err != nil {
		return err
	}
	defer stmtSecurityLoggingObjectRef.Close()
	for _, ref := range model.SecurityLoggingObjectRefs {
		_, err = stmtSecurityLoggingObjectRef.Exec(model.UUID, ref.UUID)
	}

	stmtQosConfigRef, err := tx.Prepare(insertVirtualMachineInterfaceQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmtQosConfigRef.Close()
	for _, ref := range model.QosConfigRefs {
		_, err = stmtQosConfigRef.Exec(model.UUID, ref.UUID)
	}

	stmtPhysicalInterfaceRef, err := tx.Prepare(insertVirtualMachineInterfacePhysicalInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtPhysicalInterfaceRef.Close()
	for _, ref := range model.PhysicalInterfaceRefs {
		_, err = stmtPhysicalInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	stmtSecurityGroupRef, err := tx.Prepare(insertVirtualMachineInterfaceSecurityGroupQuery)
	if err != nil {
		return err
	}
	defer stmtSecurityGroupRef.Close()
	for _, ref := range model.SecurityGroupRefs {
		_, err = stmtSecurityGroupRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertVirtualMachineInterfaceVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {
		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
	}

	stmtServiceEndpointRef, err := tx.Prepare(insertVirtualMachineInterfaceServiceEndpointQuery)
	if err != nil {
		return err
	}
	defer stmtServiceEndpointRef.Close()
	for _, ref := range model.ServiceEndpointRefs {
		_, err = stmtServiceEndpointRef.Exec(model.UUID, ref.UUID)
	}

	stmtBGPRouterRef, err := tx.Prepare(insertVirtualMachineInterfaceBGPRouterQuery)
	if err != nil {
		return err
	}
	defer stmtBGPRouterRef.Close()
	for _, ref := range model.BGPRouterRefs {
		_, err = stmtBGPRouterRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualMachineRef, err := tx.Prepare(insertVirtualMachineInterfaceVirtualMachineQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineRef.Close()
	for _, ref := range model.VirtualMachineRefs {
		_, err = stmtVirtualMachineRef.Exec(model.UUID, ref.UUID)
	}

	stmtInterfaceRouteTableRef, err := tx.Prepare(insertVirtualMachineInterfaceInterfaceRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmtInterfaceRouteTableRef.Close()
	for _, ref := range model.InterfaceRouteTableRefs {
		_, err = stmtInterfaceRouteTableRef.Exec(model.UUID, ref.UUID)
	}

	stmtRoutingInstanceRef, err := tx.Prepare(insertVirtualMachineInterfaceRoutingInstanceQuery)
	if err != nil {
		return err
	}
	defer stmtRoutingInstanceRef.Close()
	for _, ref := range model.RoutingInstanceRefs {
		_, err = stmtRoutingInstanceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.Direction),
			int(ref.Attr.MPLSLabel),
			int(ref.Attr.VlanTag),
			string(ref.Attr.SRCMac),
			string(ref.Attr.ServiceChainAddress),
			string(ref.Attr.DSTMac),
			string(ref.Attr.Protocol),
			string(ref.Attr.Ipv6ServiceChainAddress))
	}

	stmtBridgeDomainRef, err := tx.Prepare(insertVirtualMachineInterfaceBridgeDomainQuery)
	if err != nil {
		return err
	}
	defer stmtBridgeDomainRef.Close()
	for _, ref := range model.BridgeDomainRefs {
		_, err = stmtBridgeDomainRef.Exec(model.UUID, ref.UUID, int(ref.Attr.VlanTag))
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertVirtualMachineInterfaceVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	stmtPortTupleRef, err := tx.Prepare(insertVirtualMachineInterfacePortTupleQuery)
	if err != nil {
		return err
	}
	defer stmtPortTupleRef.Close()
	for _, ref := range model.PortTupleRefs {
		_, err = stmtPortTupleRef.Exec(model.UUID, ref.UUID)
	}

	stmtServiceHealthCheckRef, err := tx.Prepare(insertVirtualMachineInterfaceServiceHealthCheckQuery)
	if err != nil {
		return err
	}
	defer stmtServiceHealthCheckRef.Close()
	for _, ref := range model.ServiceHealthCheckRefs {
		_, err = stmtServiceHealthCheckRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanVirtualMachineInterface(rows *sql.Rows) (*models.VirtualMachineInterface, error) {
	m := models.MakeVirtualMachineInterface()

	var jsonVirtualMachineInterfaceHostRoutesRoute string

	var jsonVirtualMachineInterfaceDHCPOptionListDHCPOption string

	var jsonAnnotationsKeyValuePair string

	var jsonVirtualMachineInterfaceMacAddressesMacAddress string

	var jsonVirtualMachineInterfaceBindings string

	var jsonVirtualMachineInterfaceFatFlowProtocols string

	var jsonVRFAssignTableVRFAssignRule string

	var jsonFQName string

	var jsonVirtualMachineInterfaceAllowedAddressPairsAllowedAddressPair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonVirtualMachineInterfaceHostRoutesRoute,
		&jsonVirtualMachineInterfaceDHCPOptionListDHCPOption,
		&m.VirtualMachineInterfaceDeviceOwner,
		&m.PortSecurityEnabled,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress,
		&m.VirtualMachineInterfaceProperties.ServiceInterfaceType,
		&m.VirtualMachineInterfaceProperties.SubInterfaceVlanTag,
		&m.VirtualMachineInterfaceProperties.LocalPreference,
		&jsonAnnotationsKeyValuePair,
		&jsonVirtualMachineInterfaceMacAddressesMacAddress,
		&jsonVirtualMachineInterfaceBindings,
		&m.VirtualMachineInterfaceDisablePolicy,
		&m.DisplayName,
		&jsonVirtualMachineInterfaceFatFlowProtocols,
		&jsonVRFAssignTableVRFAssignRule,
		&jsonFQName,
		&m.UUID,
		&m.EcmpHashingIncludeFields.DestinationIP,
		&m.EcmpHashingIncludeFields.IPProtocol,
		&m.EcmpHashingIncludeFields.SourceIP,
		&m.EcmpHashingIncludeFields.HashingConfigured,
		&m.EcmpHashingIncludeFields.SourcePort,
		&m.EcmpHashingIncludeFields.DestinationPort,
		&jsonVirtualMachineInterfaceAllowedAddressPairsAllowedAddressPair,
		&m.VlanTagBasedBridgeDomain,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceHostRoutesRoute), &m.VirtualMachineInterfaceHostRoutes.Route)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceDHCPOptionListDHCPOption), &m.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceMacAddressesMacAddress), &m.VirtualMachineInterfaceMacAddresses.MacAddress)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceBindings), &m.VirtualMachineInterfaceBindings)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceFatFlowProtocols), &m.VirtualMachineInterfaceFatFlowProtocols)

	json.Unmarshal([]byte(jsonVRFAssignTableVRFAssignRule), &m.VRFAssignTable.VRFAssignRule)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceAllowedAddressPairsAllowedAddressPair), &m.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildVirtualMachineInterfaceWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["virtual_machine_interface_device_owner"]; ok {
		results = append(results, "virtual_machine_interface_device_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["traffic_direction"]; ok {
		results = append(results, "traffic_direction = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_name"]; ok {
		results = append(results, "analyzer_name = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_ip_address"]; ok {
		results = append(results, "analyzer_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["encapsulation"]; ok {
		results = append(results, "encapsulation = ?")
		values = append(values, value)
	}

	if value, ok := where["vtep_dst_ip_address"]; ok {
		results = append(results, "vtep_dst_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["vtep_dst_mac_address"]; ok {
		results = append(results, "vtep_dst_mac_address = ?")
		values = append(values, value)
	}

	if value, ok := where["routing_instance"]; ok {
		results = append(results, "routing_instance = ?")
		values = append(values, value)
	}

	if value, ok := where["nh_mode"]; ok {
		results = append(results, "nh_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_mac_address"]; ok {
		results = append(results, "analyzer_mac_address = ?")
		values = append(values, value)
	}

	if value, ok := where["service_interface_type"]; ok {
		results = append(results, "service_interface_type = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualMachineInterface(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualMachineInterface, error) {
	result := models.MakeVirtualMachineInterfaceSlice()
	whereQuery, values := buildVirtualMachineInterfaceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualMachineInterfaceQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanVirtualMachineInterface(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualMachineInterface(tx *sql.Tx, uuid string) (*models.VirtualMachineInterface, error) {
	rows, err := tx.Query(showVirtualMachineInterfaceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualMachineInterface(rows)
	}
	return nil, nil
}

func UpdateVirtualMachineInterface(tx *sql.Tx, uuid string, model *models.VirtualMachineInterface) error {
	return nil
}

func DeleteVirtualMachineInterface(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
