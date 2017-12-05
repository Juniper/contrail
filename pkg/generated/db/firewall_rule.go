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

const insertFirewallRuleQuery = "insert into `firewall_rule` (`protocol`,`start_port`,`end_port`,`src_ports_end_port`,`src_ports_start_port`,`protocol_id`,`match_tags`,`fq_name`,`display_name`,`key_value_pair`,`tags`,`tag_ids`,`virtual_network`,`any`,`address_group`,`ip_prefix`,`ip_prefix_len`,`endpoint_2_tags`,`endpoint_2_tag_ids`,`endpoint_2_virtual_network`,`endpoint_2_any`,`endpoint_2_address_group`,`endpoint_2_subnet_ip_prefix`,`endpoint_2_subnet_ip_prefix_len`,`apply_service`,`gateway_name`,`log`,`alert`,`qos_action`,`assign_routing_instance`,`nh_mode`,`juniper_header`,`nic_assisted_mirroring`,`routing_instance`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`vni`,`encapsulation`,`analyzer_mac_address`,`udp_port`,`nic_assisted_mirroring_vlan`,`analyzer_name`,`analyzer_ip_address`,`simple_action`,`direction`,`tag_type`,`uuid`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFirewallRuleQuery = "update `firewall_rule` set `protocol` = ?,`start_port` = ?,`end_port` = ?,`src_ports_end_port` = ?,`src_ports_start_port` = ?,`protocol_id` = ?,`match_tags` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?,`tags` = ?,`tag_ids` = ?,`virtual_network` = ?,`any` = ?,`address_group` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`endpoint_2_tags` = ?,`endpoint_2_tag_ids` = ?,`endpoint_2_virtual_network` = ?,`endpoint_2_any` = ?,`endpoint_2_address_group` = ?,`endpoint_2_subnet_ip_prefix` = ?,`endpoint_2_subnet_ip_prefix_len` = ?,`apply_service` = ?,`gateway_name` = ?,`log` = ?,`alert` = ?,`qos_action` = ?,`assign_routing_instance` = ?,`nh_mode` = ?,`juniper_header` = ?,`nic_assisted_mirroring` = ?,`routing_instance` = ?,`vtep_dst_ip_address` = ?,`vtep_dst_mac_address` = ?,`vni` = ?,`encapsulation` = ?,`analyzer_mac_address` = ?,`udp_port` = ?,`nic_assisted_mirroring_vlan` = ?,`analyzer_name` = ?,`analyzer_ip_address` = ?,`simple_action` = ?,`direction` = ?,`tag_type` = ?,`uuid` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?;"
const deleteFirewallRuleQuery = "delete from `firewall_rule` where uuid = ?"
const listFirewallRuleQuery = "select `firewall_rule`.`protocol`,`firewall_rule`.`start_port`,`firewall_rule`.`end_port`,`firewall_rule`.`src_ports_end_port`,`firewall_rule`.`src_ports_start_port`,`firewall_rule`.`protocol_id`,`firewall_rule`.`match_tags`,`firewall_rule`.`fq_name`,`firewall_rule`.`display_name`,`firewall_rule`.`key_value_pair`,`firewall_rule`.`tags`,`firewall_rule`.`tag_ids`,`firewall_rule`.`virtual_network`,`firewall_rule`.`any`,`firewall_rule`.`address_group`,`firewall_rule`.`ip_prefix`,`firewall_rule`.`ip_prefix_len`,`firewall_rule`.`endpoint_2_tags`,`firewall_rule`.`endpoint_2_tag_ids`,`firewall_rule`.`endpoint_2_virtual_network`,`firewall_rule`.`endpoint_2_any`,`firewall_rule`.`endpoint_2_address_group`,`firewall_rule`.`endpoint_2_subnet_ip_prefix`,`firewall_rule`.`endpoint_2_subnet_ip_prefix_len`,`firewall_rule`.`apply_service`,`firewall_rule`.`gateway_name`,`firewall_rule`.`log`,`firewall_rule`.`alert`,`firewall_rule`.`qos_action`,`firewall_rule`.`assign_routing_instance`,`firewall_rule`.`nh_mode`,`firewall_rule`.`juniper_header`,`firewall_rule`.`nic_assisted_mirroring`,`firewall_rule`.`routing_instance`,`firewall_rule`.`vtep_dst_ip_address`,`firewall_rule`.`vtep_dst_mac_address`,`firewall_rule`.`vni`,`firewall_rule`.`encapsulation`,`firewall_rule`.`analyzer_mac_address`,`firewall_rule`.`udp_port`,`firewall_rule`.`nic_assisted_mirroring_vlan`,`firewall_rule`.`analyzer_name`,`firewall_rule`.`analyzer_ip_address`,`firewall_rule`.`simple_action`,`firewall_rule`.`direction`,`firewall_rule`.`tag_type`,`firewall_rule`.`uuid`,`firewall_rule`.`creator`,`firewall_rule`.`user_visible`,`firewall_rule`.`last_modified`,`firewall_rule`.`owner`,`firewall_rule`.`owner_access`,`firewall_rule`.`other_access`,`firewall_rule`.`group`,`firewall_rule`.`group_access`,`firewall_rule`.`enable`,`firewall_rule`.`description`,`firewall_rule`.`created`,`firewall_rule`.`share`,`firewall_rule`.`perms2_owner`,`firewall_rule`.`perms2_owner_access`,`firewall_rule`.`global_access` from `firewall_rule`"
const showFirewallRuleQuery = "select `firewall_rule`.`protocol`,`firewall_rule`.`start_port`,`firewall_rule`.`end_port`,`firewall_rule`.`src_ports_end_port`,`firewall_rule`.`src_ports_start_port`,`firewall_rule`.`protocol_id`,`firewall_rule`.`match_tags`,`firewall_rule`.`fq_name`,`firewall_rule`.`display_name`,`firewall_rule`.`key_value_pair`,`firewall_rule`.`tags`,`firewall_rule`.`tag_ids`,`firewall_rule`.`virtual_network`,`firewall_rule`.`any`,`firewall_rule`.`address_group`,`firewall_rule`.`ip_prefix`,`firewall_rule`.`ip_prefix_len`,`firewall_rule`.`endpoint_2_tags`,`firewall_rule`.`endpoint_2_tag_ids`,`firewall_rule`.`endpoint_2_virtual_network`,`firewall_rule`.`endpoint_2_any`,`firewall_rule`.`endpoint_2_address_group`,`firewall_rule`.`endpoint_2_subnet_ip_prefix`,`firewall_rule`.`endpoint_2_subnet_ip_prefix_len`,`firewall_rule`.`apply_service`,`firewall_rule`.`gateway_name`,`firewall_rule`.`log`,`firewall_rule`.`alert`,`firewall_rule`.`qos_action`,`firewall_rule`.`assign_routing_instance`,`firewall_rule`.`nh_mode`,`firewall_rule`.`juniper_header`,`firewall_rule`.`nic_assisted_mirroring`,`firewall_rule`.`routing_instance`,`firewall_rule`.`vtep_dst_ip_address`,`firewall_rule`.`vtep_dst_mac_address`,`firewall_rule`.`vni`,`firewall_rule`.`encapsulation`,`firewall_rule`.`analyzer_mac_address`,`firewall_rule`.`udp_port`,`firewall_rule`.`nic_assisted_mirroring_vlan`,`firewall_rule`.`analyzer_name`,`firewall_rule`.`analyzer_ip_address`,`firewall_rule`.`simple_action`,`firewall_rule`.`direction`,`firewall_rule`.`tag_type`,`firewall_rule`.`uuid`,`firewall_rule`.`creator`,`firewall_rule`.`user_visible`,`firewall_rule`.`last_modified`,`firewall_rule`.`owner`,`firewall_rule`.`owner_access`,`firewall_rule`.`other_access`,`firewall_rule`.`group`,`firewall_rule`.`group_access`,`firewall_rule`.`enable`,`firewall_rule`.`description`,`firewall_rule`.`created`,`firewall_rule`.`share`,`firewall_rule`.`perms2_owner`,`firewall_rule`.`perms2_owner_access`,`firewall_rule`.`global_access` from `firewall_rule` where uuid = ?"

const insertFirewallRuleVirtualNetworkQuery = "insert into `ref_firewall_rule_virtual_network` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleServiceGroupQuery = "insert into `ref_firewall_rule_service_group` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleAddressGroupQuery = "insert into `ref_firewall_rule_address_group` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleSecurityLoggingObjectQuery = "insert into `ref_firewall_rule_security_logging_object` (`from`, `to` ) values (?, ?);"

func CreateFirewallRule(tx *sql.Tx, model *models.FirewallRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFirewallRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Service.Protocol),
		int(model.Service.DSTPorts.StartPort),
		int(model.Service.DSTPorts.EndPort),
		int(model.Service.SRCPorts.EndPort),
		int(model.Service.SRCPorts.StartPort),
		int(model.Service.ProtocolID),
		utils.MustJSON(model.MatchTags),
		utils.MustJSON(model.FQName),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Endpoint1.Tags),
		utils.MustJSON(model.Endpoint1.TagIds),
		string(model.Endpoint1.VirtualNetwork),
		bool(model.Endpoint1.Any),
		string(model.Endpoint1.AddressGroup),
		string(model.Endpoint1.Subnet.IPPrefix),
		int(model.Endpoint1.Subnet.IPPrefixLen),
		utils.MustJSON(model.Endpoint2.Tags),
		utils.MustJSON(model.Endpoint2.TagIds),
		string(model.Endpoint2.VirtualNetwork),
		bool(model.Endpoint2.Any),
		string(model.Endpoint2.AddressGroup),
		string(model.Endpoint2.Subnet.IPPrefix),
		int(model.Endpoint2.Subnet.IPPrefixLen),
		utils.MustJSON(model.ActionList.ApplyService),
		string(model.ActionList.GatewayName),
		bool(model.ActionList.Log),
		bool(model.ActionList.Alert),
		string(model.ActionList.QosAction),
		string(model.ActionList.AssignRoutingInstance),
		string(model.ActionList.MirrorTo.NHMode),
		bool(model.ActionList.MirrorTo.JuniperHeader),
		bool(model.ActionList.MirrorTo.NicAssistedMirroring),
		string(model.ActionList.MirrorTo.RoutingInstance),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		int(model.ActionList.MirrorTo.StaticNHHeader.Vni),
		string(model.ActionList.MirrorTo.Encapsulation),
		string(model.ActionList.MirrorTo.AnalyzerMacAddress),
		int(model.ActionList.MirrorTo.UDPPort),
		int(model.ActionList.MirrorTo.NicAssistedMirroringVlan),
		string(model.ActionList.MirrorTo.AnalyzerName),
		string(model.ActionList.MirrorTo.AnalyzerIPAddress),
		string(model.ActionList.SimpleAction),
		string(model.Direction),
		utils.MustJSON(model.MatchTagTypes.TagType),
		string(model.UUID),
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
		string(model.IDPerms.Created),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess))

	stmtServiceGroupRef, err := tx.Prepare(insertFirewallRuleServiceGroupQuery)
	if err != nil {
		return err
	}
	defer stmtServiceGroupRef.Close()
	for _, ref := range model.ServiceGroupRefs {
		_, err = stmtServiceGroupRef.Exec(model.UUID, ref.UUID)
	}

	stmtAddressGroupRef, err := tx.Prepare(insertFirewallRuleAddressGroupQuery)
	if err != nil {
		return err
	}
	defer stmtAddressGroupRef.Close()
	for _, ref := range model.AddressGroupRefs {
		_, err = stmtAddressGroupRef.Exec(model.UUID, ref.UUID)
	}

	stmtSecurityLoggingObjectRef, err := tx.Prepare(insertFirewallRuleSecurityLoggingObjectQuery)
	if err != nil {
		return err
	}
	defer stmtSecurityLoggingObjectRef.Close()
	for _, ref := range model.SecurityLoggingObjectRefs {
		_, err = stmtSecurityLoggingObjectRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertFirewallRuleVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {
		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanFirewallRule(rows *sql.Rows) (*models.FirewallRule, error) {
	m := models.MakeFirewallRule()

	var jsonMatchTags string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonEndpoint1Tags string

	var jsonEndpoint1TagIds string

	var jsonEndpoint2Tags string

	var jsonEndpoint2TagIds string

	var jsonActionListApplyService string

	var jsonMatchTagTypesTagType string

	var jsonPerms2Share string

	if err := rows.Scan(&m.Service.Protocol,
		&m.Service.DSTPorts.StartPort,
		&m.Service.DSTPorts.EndPort,
		&m.Service.SRCPorts.EndPort,
		&m.Service.SRCPorts.StartPort,
		&m.Service.ProtocolID,
		&jsonMatchTags,
		&jsonFQName,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonEndpoint1Tags,
		&jsonEndpoint1TagIds,
		&m.Endpoint1.VirtualNetwork,
		&m.Endpoint1.Any,
		&m.Endpoint1.AddressGroup,
		&m.Endpoint1.Subnet.IPPrefix,
		&m.Endpoint1.Subnet.IPPrefixLen,
		&jsonEndpoint2Tags,
		&jsonEndpoint2TagIds,
		&m.Endpoint2.VirtualNetwork,
		&m.Endpoint2.Any,
		&m.Endpoint2.AddressGroup,
		&m.Endpoint2.Subnet.IPPrefix,
		&m.Endpoint2.Subnet.IPPrefixLen,
		&jsonActionListApplyService,
		&m.ActionList.GatewayName,
		&m.ActionList.Log,
		&m.ActionList.Alert,
		&m.ActionList.QosAction,
		&m.ActionList.AssignRoutingInstance,
		&m.ActionList.MirrorTo.NHMode,
		&m.ActionList.MirrorTo.JuniperHeader,
		&m.ActionList.MirrorTo.NicAssistedMirroring,
		&m.ActionList.MirrorTo.RoutingInstance,
		&m.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
		&m.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
		&m.ActionList.MirrorTo.StaticNHHeader.Vni,
		&m.ActionList.MirrorTo.Encapsulation,
		&m.ActionList.MirrorTo.AnalyzerMacAddress,
		&m.ActionList.MirrorTo.UDPPort,
		&m.ActionList.MirrorTo.NicAssistedMirroringVlan,
		&m.ActionList.MirrorTo.AnalyzerName,
		&m.ActionList.MirrorTo.AnalyzerIPAddress,
		&m.ActionList.SimpleAction,
		&m.Direction,
		&jsonMatchTagTypesTagType,
		&m.UUID,
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
		&m.IDPerms.Created,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonMatchTags), &m.MatchTags)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonEndpoint1Tags), &m.Endpoint1.Tags)

	json.Unmarshal([]byte(jsonEndpoint1TagIds), &m.Endpoint1.TagIds)

	json.Unmarshal([]byte(jsonEndpoint2Tags), &m.Endpoint2.Tags)

	json.Unmarshal([]byte(jsonEndpoint2TagIds), &m.Endpoint2.TagIds)

	json.Unmarshal([]byte(jsonActionListApplyService), &m.ActionList.ApplyService)

	json.Unmarshal([]byte(jsonMatchTagTypesTagType), &m.MatchTagTypes.TagType)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildFirewallRuleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["protocol"]; ok {
		results = append(results, "protocol = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["virtual_network"]; ok {
		results = append(results, "virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["address_group"]; ok {
		results = append(results, "address_group = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
		values = append(values, value)
	}

	if value, ok := where["endpoint_2_virtual_network"]; ok {
		results = append(results, "endpoint_2_virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["endpoint_2_address_group"]; ok {
		results = append(results, "endpoint_2_address_group = ?")
		values = append(values, value)
	}

	if value, ok := where["endpoint_2_subnet_ip_prefix"]; ok {
		results = append(results, "endpoint_2_subnet_ip_prefix = ?")
		values = append(values, value)
	}

	if value, ok := where["gateway_name"]; ok {
		results = append(results, "gateway_name = ?")
		values = append(values, value)
	}

	if value, ok := where["qos_action"]; ok {
		results = append(results, "qos_action = ?")
		values = append(values, value)
	}

	if value, ok := where["assign_routing_instance"]; ok {
		results = append(results, "assign_routing_instance = ?")
		values = append(values, value)
	}

	if value, ok := where["nh_mode"]; ok {
		results = append(results, "nh_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["routing_instance"]; ok {
		results = append(results, "routing_instance = ?")
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

	if value, ok := where["encapsulation"]; ok {
		results = append(results, "encapsulation = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_mac_address"]; ok {
		results = append(results, "analyzer_mac_address = ?")
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

	if value, ok := where["simple_action"]; ok {
		results = append(results, "simple_action = ?")
		values = append(values, value)
	}

	if value, ok := where["direction"]; ok {
		results = append(results, "direction = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListFirewallRule(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.FirewallRule, error) {
	result := models.MakeFirewallRuleSlice()
	whereQuery, values := buildFirewallRuleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listFirewallRuleQuery)
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
		m, _ := scanFirewallRule(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowFirewallRule(tx *sql.Tx, uuid string) (*models.FirewallRule, error) {
	rows, err := tx.Query(showFirewallRuleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanFirewallRule(rows)
	}
	return nil, nil
}

func UpdateFirewallRule(tx *sql.Tx, uuid string, model *models.FirewallRule) error {
	return nil
}

func DeleteFirewallRule(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFirewallRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
