package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertFirewallRuleQuery = "insert into `firewall_rule` (`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`display_name`,`apply_service`,`gateway_name`,`log`,`alert`,`qos_action`,`assign_routing_instance`,`analyzer_name`,`nh_mode`,`nic_assisted_mirroring_vlan`,`vni`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`analyzer_ip_address`,`analyzer_mac_address`,`udp_port`,`encapsulation`,`routing_instance`,`nic_assisted_mirroring`,`juniper_header`,`simple_action`,`tag_type`,`fq_name`,`direction`,`match_tags`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`key_value_pair`,`ip_prefix_len`,`ip_prefix`,`tags`,`tag_ids`,`virtual_network`,`any`,`address_group`,`endpoint_2_subnet_ip_prefix`,`endpoint_2_subnet_ip_prefix_len`,`endpoint_2_tags`,`endpoint_2_tag_ids`,`endpoint_2_virtual_network`,`endpoint_2_any`,`endpoint_2_address_group`,`end_port`,`start_port`,`src_ports_start_port`,`src_ports_end_port`,`protocol_id`,`protocol`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFirewallRuleQuery = "update `firewall_rule` set `enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`display_name` = ?,`apply_service` = ?,`gateway_name` = ?,`log` = ?,`alert` = ?,`qos_action` = ?,`assign_routing_instance` = ?,`analyzer_name` = ?,`nh_mode` = ?,`nic_assisted_mirroring_vlan` = ?,`vni` = ?,`vtep_dst_ip_address` = ?,`vtep_dst_mac_address` = ?,`analyzer_ip_address` = ?,`analyzer_mac_address` = ?,`udp_port` = ?,`encapsulation` = ?,`routing_instance` = ?,`nic_assisted_mirroring` = ?,`juniper_header` = ?,`simple_action` = ?,`tag_type` = ?,`fq_name` = ?,`direction` = ?,`match_tags` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`key_value_pair` = ?,`ip_prefix_len` = ?,`ip_prefix` = ?,`tags` = ?,`tag_ids` = ?,`virtual_network` = ?,`any` = ?,`address_group` = ?,`endpoint_2_subnet_ip_prefix` = ?,`endpoint_2_subnet_ip_prefix_len` = ?,`endpoint_2_tags` = ?,`endpoint_2_tag_ids` = ?,`endpoint_2_virtual_network` = ?,`endpoint_2_any` = ?,`endpoint_2_address_group` = ?,`end_port` = ?,`start_port` = ?,`src_ports_start_port` = ?,`src_ports_end_port` = ?,`protocol_id` = ?,`protocol` = ?;"
const deleteFirewallRuleQuery = "delete from `firewall_rule` where uuid = ?"

// FirewallRuleFields is db columns for FirewallRule
var FirewallRuleFields = []string{
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"owner",
	"display_name",
	"apply_service",
	"gateway_name",
	"log",
	"alert",
	"qos_action",
	"assign_routing_instance",
	"analyzer_name",
	"nh_mode",
	"nic_assisted_mirroring_vlan",
	"vni",
	"vtep_dst_ip_address",
	"vtep_dst_mac_address",
	"analyzer_ip_address",
	"analyzer_mac_address",
	"udp_port",
	"encapsulation",
	"routing_instance",
	"nic_assisted_mirroring",
	"juniper_header",
	"simple_action",
	"tag_type",
	"fq_name",
	"direction",
	"match_tags",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"uuid",
	"key_value_pair",
	"ip_prefix_len",
	"ip_prefix",
	"tags",
	"tag_ids",
	"virtual_network",
	"any",
	"address_group",
	"endpoint_2_subnet_ip_prefix",
	"endpoint_2_subnet_ip_prefix_len",
	"endpoint_2_tags",
	"endpoint_2_tag_ids",
	"endpoint_2_virtual_network",
	"endpoint_2_any",
	"endpoint_2_address_group",
	"end_port",
	"start_port",
	"src_ports_start_port",
	"src_ports_end_port",
	"protocol_id",
	"protocol",
}

// FirewallRuleRefFields is db reference fields for FirewallRule
var FirewallRuleRefFields = map[string][]string{

	"service_group": {
	// <common.Schema Value>

	},

	"address_group": {
	// <common.Schema Value>

	},

	"security_logging_object": {
	// <common.Schema Value>

	},

	"virtual_network": {
	// <common.Schema Value>

	},
}

const insertFirewallRuleServiceGroupQuery = "insert into `ref_firewall_rule_service_group` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleAddressGroupQuery = "insert into `ref_firewall_rule_address_group` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleSecurityLoggingObjectQuery = "insert into `ref_firewall_rule_security_logging_object` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleVirtualNetworkQuery = "insert into `ref_firewall_rule_virtual_network` (`from`, `to` ) values (?, ?);"

// CreateFirewallRule inserts FirewallRule to DB
func CreateFirewallRule(tx *sql.Tx, model *models.FirewallRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFirewallRuleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertFirewallRuleQuery,
	}).Debug("create query")
	_, err = stmt.Exec(bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		string(model.DisplayName),
		common.MustJSON(model.ActionList.ApplyService),
		string(model.ActionList.GatewayName),
		bool(model.ActionList.Log),
		bool(model.ActionList.Alert),
		string(model.ActionList.QosAction),
		string(model.ActionList.AssignRoutingInstance),
		string(model.ActionList.MirrorTo.AnalyzerName),
		string(model.ActionList.MirrorTo.NHMode),
		int(model.ActionList.MirrorTo.NicAssistedMirroringVlan),
		int(model.ActionList.MirrorTo.StaticNHHeader.Vni),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		string(model.ActionList.MirrorTo.AnalyzerIPAddress),
		string(model.ActionList.MirrorTo.AnalyzerMacAddress),
		int(model.ActionList.MirrorTo.UDPPort),
		string(model.ActionList.MirrorTo.Encapsulation),
		string(model.ActionList.MirrorTo.RoutingInstance),
		bool(model.ActionList.MirrorTo.NicAssistedMirroring),
		bool(model.ActionList.MirrorTo.JuniperHeader),
		string(model.ActionList.SimpleAction),
		common.MustJSON(model.MatchTagTypes.TagType),
		common.MustJSON(model.FQName),
		string(model.Direction),
		common.MustJSON(model.MatchTags),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		common.MustJSON(model.Annotations.KeyValuePair),
		int(model.Endpoint1.Subnet.IPPrefixLen),
		string(model.Endpoint1.Subnet.IPPrefix),
		common.MustJSON(model.Endpoint1.Tags),
		common.MustJSON(model.Endpoint1.TagIds),
		string(model.Endpoint1.VirtualNetwork),
		bool(model.Endpoint1.Any),
		string(model.Endpoint1.AddressGroup),
		string(model.Endpoint2.Subnet.IPPrefix),
		int(model.Endpoint2.Subnet.IPPrefixLen),
		common.MustJSON(model.Endpoint2.Tags),
		common.MustJSON(model.Endpoint2.TagIds),
		string(model.Endpoint2.VirtualNetwork),
		bool(model.Endpoint2.Any),
		string(model.Endpoint2.AddressGroup),
		int(model.Service.DSTPorts.EndPort),
		int(model.Service.DSTPorts.StartPort),
		int(model.Service.SRCPorts.StartPort),
		int(model.Service.SRCPorts.EndPort),
		int(model.Service.ProtocolID),
		string(model.Service.Protocol))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceGroupRef, err := tx.Prepare(insertFirewallRuleServiceGroupQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceGroupRefs create statement failed")
	}
	defer stmtServiceGroupRef.Close()
	for _, ref := range model.ServiceGroupRefs {

		_, err = stmtServiceGroupRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceGroupRefs create failed")
		}
	}

	stmtAddressGroupRef, err := tx.Prepare(insertFirewallRuleAddressGroupQuery)
	if err != nil {
		return errors.Wrap(err, "preparing AddressGroupRefs create statement failed")
	}
	defer stmtAddressGroupRef.Close()
	for _, ref := range model.AddressGroupRefs {

		_, err = stmtAddressGroupRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "AddressGroupRefs create failed")
		}
	}

	stmtSecurityLoggingObjectRef, err := tx.Prepare(insertFirewallRuleSecurityLoggingObjectQuery)
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

	stmtVirtualNetworkRef, err := tx.Prepare(insertFirewallRuleVirtualNetworkQuery)
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

func scanFirewallRule(values map[string]interface{}) (*models.FirewallRule, error) {
	m := models.MakeFirewallRule()

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

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

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

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["apply_service"]; ok {

		json.Unmarshal(value.([]byte), &m.ActionList.ApplyService)

	}

	if value, ok := values["gateway_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.GatewayName = castedValue

	}

	if value, ok := values["log"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ActionList.Log = castedValue

	}

	if value, ok := values["alert"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ActionList.Alert = castedValue

	}

	if value, ok := values["qos_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.QosAction = castedValue

	}

	if value, ok := values["assign_routing_instance"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.AssignRoutingInstance = castedValue

	}

	if value, ok := values["analyzer_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.AnalyzerName = castedValue

	}

	if value, ok := values["nh_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.NHMode = models.NHModeType(castedValue)

	}

	if value, ok := values["nic_assisted_mirroring_vlan"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ActionList.MirrorTo.NicAssistedMirroringVlan = models.VlanIdType(castedValue)

	}

	if value, ok := values["vni"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ActionList.MirrorTo.StaticNHHeader.Vni = models.VxlanNetworkIdentifierType(castedValue)

	}

	if value, ok := values["vtep_dst_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress = castedValue

	}

	if value, ok := values["vtep_dst_mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress = castedValue

	}

	if value, ok := values["analyzer_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.AnalyzerIPAddress = castedValue

	}

	if value, ok := values["analyzer_mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.AnalyzerMacAddress = castedValue

	}

	if value, ok := values["udp_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ActionList.MirrorTo.UDPPort = castedValue

	}

	if value, ok := values["encapsulation"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.Encapsulation = castedValue

	}

	if value, ok := values["routing_instance"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.RoutingInstance = castedValue

	}

	if value, ok := values["nic_assisted_mirroring"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ActionList.MirrorTo.NicAssistedMirroring = castedValue

	}

	if value, ok := values["juniper_header"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ActionList.MirrorTo.JuniperHeader = castedValue

	}

	if value, ok := values["simple_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.SimpleAction = models.SimpleActionType(castedValue)

	}

	if value, ok := values["tag_type"]; ok {

		json.Unmarshal(value.([]byte), &m.MatchTagTypes.TagType)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["direction"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Direction = models.FirewallRuleDirectionType(castedValue)

	}

	if value, ok := values["match_tags"]; ok {

		json.Unmarshal(value.([]byte), &m.MatchTags)

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

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ip_prefix_len"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Endpoint1.Subnet.IPPrefixLen = castedValue

	}

	if value, ok := values["ip_prefix"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint1.Subnet.IPPrefix = castedValue

	}

	if value, ok := values["tags"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint1.Tags)

	}

	if value, ok := values["tag_ids"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint1.TagIds)

	}

	if value, ok := values["virtual_network"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint1.VirtualNetwork = castedValue

	}

	if value, ok := values["any"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.Endpoint1.Any = castedValue

	}

	if value, ok := values["address_group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint1.AddressGroup = castedValue

	}

	if value, ok := values["endpoint_2_subnet_ip_prefix"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint2.Subnet.IPPrefix = castedValue

	}

	if value, ok := values["endpoint_2_subnet_ip_prefix_len"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Endpoint2.Subnet.IPPrefixLen = castedValue

	}

	if value, ok := values["endpoint_2_tags"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint2.Tags)

	}

	if value, ok := values["endpoint_2_tag_ids"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint2.TagIds)

	}

	if value, ok := values["endpoint_2_virtual_network"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint2.VirtualNetwork = castedValue

	}

	if value, ok := values["endpoint_2_any"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.Endpoint2.Any = castedValue

	}

	if value, ok := values["endpoint_2_address_group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint2.AddressGroup = castedValue

	}

	if value, ok := values["end_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Service.DSTPorts.EndPort = models.L4PortType(castedValue)

	}

	if value, ok := values["start_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Service.DSTPorts.StartPort = models.L4PortType(castedValue)

	}

	if value, ok := values["src_ports_start_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Service.SRCPorts.StartPort = models.L4PortType(castedValue)

	}

	if value, ok := values["src_ports_end_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Service.SRCPorts.EndPort = models.L4PortType(castedValue)

	}

	if value, ok := values["protocol_id"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Service.ProtocolID = castedValue

	}

	if value, ok := values["protocol"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Service.Protocol = castedValue

	}

	if value, ok := values["ref_service_group"]; ok {
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
			referenceModel := &models.FirewallRuleServiceGroupRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceGroupRefs = append(m.ServiceGroupRefs, referenceModel)

		}
	}

	if value, ok := values["ref_address_group"]; ok {
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
			referenceModel := &models.FirewallRuleAddressGroupRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.AddressGroupRefs = append(m.AddressGroupRefs, referenceModel)

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
			referenceModel := &models.FirewallRuleSecurityLoggingObjectRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.SecurityLoggingObjectRefs = append(m.SecurityLoggingObjectRefs, referenceModel)

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
			referenceModel := &models.FirewallRuleVirtualNetworkRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	return m, nil
}

// ListFirewallRule lists FirewallRule with list spec.
func ListFirewallRule(tx *sql.Tx, spec *common.ListSpec) ([]*models.FirewallRule, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "firewall_rule"
	spec.Fields = FirewallRuleFields
	spec.RefFields = FirewallRuleRefFields
	result := models.MakeFirewallRuleSlice()
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
		m, err := scanFirewallRule(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowFirewallRule shows FirewallRule resource
func ShowFirewallRule(tx *sql.Tx, uuid string) (*models.FirewallRule, error) {
	list, err := ListFirewallRule(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateFirewallRule updates a resource
func UpdateFirewallRule(tx *sql.Tx, uuid string, model *models.FirewallRule) error {
	//TODO(nati) support update
	return nil
}

// DeleteFirewallRule deletes a resource
func DeleteFirewallRule(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFirewallRuleQuery)
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
