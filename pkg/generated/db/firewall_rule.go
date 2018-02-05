package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertFirewallRuleQuery = "insert into `firewall_rule` (`uuid`,`start_port`,`end_port`,`protocol_id`,`protocol`,`dst_ports_start_port`,`dst_ports_end_port`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`tag_list`,`tag_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`virtual_network`,`tags`,`tag_ids`,`ip_prefix_len`,`ip_prefix`,`any`,`address_group`,`endpoint_1_virtual_network`,`endpoint_1_tags`,`endpoint_1_tag_ids`,`subnet_ip_prefix_len`,`subnet_ip_prefix`,`endpoint_1_any`,`endpoint_1_address_group`,`display_name`,`direction`,`key_value_pair`,`simple_action`,`qos_action`,`udp_port`,`vtep_dst_mac_address`,`vtep_dst_ip_address`,`vni`,`routing_instance`,`nic_assisted_mirroring_vlan`,`nic_assisted_mirroring`,`nh_mode`,`juniper_header`,`encapsulation`,`analyzer_name`,`analyzer_mac_address`,`analyzer_ip_address`,`log`,`gateway_name`,`assign_routing_instance`,`apply_service`,`alert`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteFirewallRuleQuery = "delete from `firewall_rule` where uuid = ?"

// FirewallRuleFields is db columns for FirewallRule
var FirewallRuleFields = []string{
	"uuid",
	"start_port",
	"end_port",
	"protocol_id",
	"protocol",
	"dst_ports_start_port",
	"dst_ports_end_port",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"tag_list",
	"tag_type",
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
	"virtual_network",
	"tags",
	"tag_ids",
	"ip_prefix_len",
	"ip_prefix",
	"any",
	"address_group",
	"endpoint_1_virtual_network",
	"endpoint_1_tags",
	"endpoint_1_tag_ids",
	"subnet_ip_prefix_len",
	"subnet_ip_prefix",
	"endpoint_1_any",
	"endpoint_1_address_group",
	"display_name",
	"direction",
	"key_value_pair",
	"simple_action",
	"qos_action",
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
	"log",
	"gateway_name",
	"assign_routing_instance",
	"apply_service",
	"alert",
}

// FirewallRuleRefFields is db reference fields for FirewallRule
var FirewallRuleRefFields = map[string][]string{

	"virtual_network": {
	// <common.Schema Value>

	},

	"service_group": {
	// <common.Schema Value>

	},

	"address_group": {
	// <common.Schema Value>

	},

	"security_logging_object": {
	// <common.Schema Value>

	},
}

// FirewallRuleBackRefFields is db back reference fields for FirewallRule
var FirewallRuleBackRefFields = map[string][]string{}

// FirewallRuleParentTypes is possible parents for FirewallRule
var FirewallRuleParents = []string{

	"project",

	"policy_management",
}

const insertFirewallRuleSecurityLoggingObjectQuery = "insert into `ref_firewall_rule_security_logging_object` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleVirtualNetworkQuery = "insert into `ref_firewall_rule_virtual_network` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleServiceGroupQuery = "insert into `ref_firewall_rule_service_group` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleAddressGroupQuery = "insert into `ref_firewall_rule_address_group` (`from`, `to` ) values (?, ?);"

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
	_, err = stmt.Exec(string(model.UUID),
		int(model.Service.SRCPorts.StartPort),
		int(model.Service.SRCPorts.EndPort),
		int(model.Service.ProtocolID),
		string(model.Service.Protocol),
		int(model.Service.DSTPorts.StartPort),
		int(model.Service.DSTPorts.EndPort),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		common.MustJSON(model.MatchTags.TagList),
		common.MustJSON(model.MatchTagTypes.TagType),
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
		string(model.Endpoint2.VirtualNetwork),
		common.MustJSON(model.Endpoint2.Tags),
		common.MustJSON(model.Endpoint2.TagIds),
		int(model.Endpoint2.Subnet.IPPrefixLen),
		string(model.Endpoint2.Subnet.IPPrefix),
		bool(model.Endpoint2.Any),
		string(model.Endpoint2.AddressGroup),
		string(model.Endpoint1.VirtualNetwork),
		common.MustJSON(model.Endpoint1.Tags),
		common.MustJSON(model.Endpoint1.TagIds),
		int(model.Endpoint1.Subnet.IPPrefixLen),
		string(model.Endpoint1.Subnet.IPPrefix),
		bool(model.Endpoint1.Any),
		string(model.Endpoint1.AddressGroup),
		string(model.DisplayName),
		string(model.Direction),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.ActionList.SimpleAction),
		string(model.ActionList.QosAction),
		int(model.ActionList.MirrorTo.UDPPort),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		int(model.ActionList.MirrorTo.StaticNHHeader.Vni),
		string(model.ActionList.MirrorTo.RoutingInstance),
		int(model.ActionList.MirrorTo.NicAssistedMirroringVlan),
		bool(model.ActionList.MirrorTo.NicAssistedMirroring),
		string(model.ActionList.MirrorTo.NHMode),
		bool(model.ActionList.MirrorTo.JuniperHeader),
		string(model.ActionList.MirrorTo.Encapsulation),
		string(model.ActionList.MirrorTo.AnalyzerName),
		string(model.ActionList.MirrorTo.AnalyzerMacAddress),
		string(model.ActionList.MirrorTo.AnalyzerIPAddress),
		bool(model.ActionList.Log),
		string(model.ActionList.GatewayName),
		string(model.ActionList.AssignRoutingInstance),
		common.MustJSON(model.ActionList.ApplyService),
		bool(model.ActionList.Alert))
	if err != nil {
		return errors.Wrap(err, "create failed")
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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "firewall_rule",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "firewall_rule", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanFirewallRule(values map[string]interface{}) (*models.FirewallRule, error) {
	m := models.MakeFirewallRule()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["start_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Service.SRCPorts.StartPort = models.L4PortType(castedValue)

	}

	if value, ok := values["end_port"]; ok {

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

	if value, ok := values["dst_ports_start_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Service.DSTPorts.StartPort = models.L4PortType(castedValue)

	}

	if value, ok := values["dst_ports_end_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Service.DSTPorts.EndPort = models.L4PortType(castedValue)

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

	if value, ok := values["tag_list"]; ok {

		json.Unmarshal(value.([]byte), &m.MatchTags.TagList)

	}

	if value, ok := values["tag_type"]; ok {

		json.Unmarshal(value.([]byte), &m.MatchTagTypes.TagType)

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

	if value, ok := values["virtual_network"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint2.VirtualNetwork = castedValue

	}

	if value, ok := values["tags"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint2.Tags)

	}

	if value, ok := values["tag_ids"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint2.TagIds)

	}

	if value, ok := values["ip_prefix_len"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Endpoint2.Subnet.IPPrefixLen = castedValue

	}

	if value, ok := values["ip_prefix"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint2.Subnet.IPPrefix = castedValue

	}

	if value, ok := values["any"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.Endpoint2.Any = castedValue

	}

	if value, ok := values["address_group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint2.AddressGroup = castedValue

	}

	if value, ok := values["endpoint_1_virtual_network"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint1.VirtualNetwork = castedValue

	}

	if value, ok := values["endpoint_1_tags"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint1.Tags)

	}

	if value, ok := values["endpoint_1_tag_ids"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint1.TagIds)

	}

	if value, ok := values["subnet_ip_prefix_len"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Endpoint1.Subnet.IPPrefixLen = castedValue

	}

	if value, ok := values["subnet_ip_prefix"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint1.Subnet.IPPrefix = castedValue

	}

	if value, ok := values["endpoint_1_any"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.Endpoint1.Any = castedValue

	}

	if value, ok := values["endpoint_1_address_group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Endpoint1.AddressGroup = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["direction"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Direction = models.FirewallRuleDirectionType(castedValue)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["simple_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.SimpleAction = models.SimpleActionType(castedValue)

	}

	if value, ok := values["qos_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.QosAction = castedValue

	}

	if value, ok := values["udp_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ActionList.MirrorTo.UDPPort = castedValue

	}

	if value, ok := values["vtep_dst_mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress = castedValue

	}

	if value, ok := values["vtep_dst_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress = castedValue

	}

	if value, ok := values["vni"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ActionList.MirrorTo.StaticNHHeader.Vni = models.VxlanNetworkIdentifierType(castedValue)

	}

	if value, ok := values["routing_instance"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.RoutingInstance = castedValue

	}

	if value, ok := values["nic_assisted_mirroring_vlan"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ActionList.MirrorTo.NicAssistedMirroringVlan = models.VlanIdType(castedValue)

	}

	if value, ok := values["nic_assisted_mirroring"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ActionList.MirrorTo.NicAssistedMirroring = castedValue

	}

	if value, ok := values["nh_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.NHMode = models.NHModeType(castedValue)

	}

	if value, ok := values["juniper_header"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ActionList.MirrorTo.JuniperHeader = castedValue

	}

	if value, ok := values["encapsulation"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.Encapsulation = castedValue

	}

	if value, ok := values["analyzer_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.AnalyzerName = castedValue

	}

	if value, ok := values["analyzer_mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.AnalyzerMacAddress = castedValue

	}

	if value, ok := values["analyzer_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.MirrorTo.AnalyzerIPAddress = castedValue

	}

	if value, ok := values["log"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ActionList.Log = castedValue

	}

	if value, ok := values["gateway_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.GatewayName = castedValue

	}

	if value, ok := values["assign_routing_instance"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ActionList.AssignRoutingInstance = castedValue

	}

	if value, ok := values["apply_service"]; ok {

		json.Unmarshal(value.([]byte), &m.ActionList.ApplyService)

	}

	if value, ok := values["alert"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ActionList.Alert = castedValue

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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.FirewallRuleAddressGroupRef{}
			referenceModel.UUID = uuid
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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.FirewallRuleSecurityLoggingObjectRef{}
			referenceModel.UUID = uuid
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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.FirewallRuleVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.FirewallRuleServiceGroupRef{}
			referenceModel.UUID = uuid
			m.ServiceGroupRefs = append(m.ServiceGroupRefs, referenceModel)

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
	spec.BackRefFields = FirewallRuleBackRefFields
	result := models.MakeFirewallRuleSlice()

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
		m, err := scanFirewallRule(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateFirewallRule updates a resource
func UpdateFirewallRule(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	// Prepare statement for updating data
	var updateFirewallRuleQuery = "update `firewall_rule` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateFirewallRuleQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Service.SRCPorts.StartPort", "."); ok {
		updateFirewallRuleQuery += "`start_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Service.SRCPorts.EndPort", "."); ok {
		updateFirewallRuleQuery += "`end_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Service.ProtocolID", "."); ok {
		updateFirewallRuleQuery += "`protocol_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Service.Protocol", "."); ok {
		updateFirewallRuleQuery += "`protocol` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Service.DSTPorts.StartPort", "."); ok {
		updateFirewallRuleQuery += "`dst_ports_start_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Service.DSTPorts.EndPort", "."); ok {
		updateFirewallRuleQuery += "`dst_ports_end_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateFirewallRuleQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateFirewallRuleQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateFirewallRuleQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateFirewallRuleQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateFirewallRuleQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateFirewallRuleQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MatchTags.TagList", "."); ok {
		updateFirewallRuleQuery += "`tag_list` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MatchTagTypes.TagType", "."); ok {
		updateFirewallRuleQuery += "`tag_type` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateFirewallRuleQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateFirewallRuleQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateFirewallRuleQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateFirewallRuleQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateFirewallRuleQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateFirewallRuleQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateFirewallRuleQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateFirewallRuleQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateFirewallRuleQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateFirewallRuleQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateFirewallRuleQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateFirewallRuleQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint2.VirtualNetwork", "."); ok {
		updateFirewallRuleQuery += "`virtual_network` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint2.Tags", "."); ok {
		updateFirewallRuleQuery += "`tags` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint2.TagIds", "."); ok {
		updateFirewallRuleQuery += "`tag_ids` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint2.Subnet.IPPrefixLen", "."); ok {
		updateFirewallRuleQuery += "`ip_prefix_len` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint2.Subnet.IPPrefix", "."); ok {
		updateFirewallRuleQuery += "`ip_prefix` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint2.Any", "."); ok {
		updateFirewallRuleQuery += "`any` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint2.AddressGroup", "."); ok {
		updateFirewallRuleQuery += "`address_group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint1.VirtualNetwork", "."); ok {
		updateFirewallRuleQuery += "`endpoint_1_virtual_network` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint1.Tags", "."); ok {
		updateFirewallRuleQuery += "`endpoint_1_tags` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint1.TagIds", "."); ok {
		updateFirewallRuleQuery += "`endpoint_1_tag_ids` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint1.Subnet.IPPrefixLen", "."); ok {
		updateFirewallRuleQuery += "`subnet_ip_prefix_len` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint1.Subnet.IPPrefix", "."); ok {
		updateFirewallRuleQuery += "`subnet_ip_prefix` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint1.Any", "."); ok {
		updateFirewallRuleQuery += "`endpoint_1_any` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Endpoint1.AddressGroup", "."); ok {
		updateFirewallRuleQuery += "`endpoint_1_address_group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateFirewallRuleQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Direction", "."); ok {
		updateFirewallRuleQuery += "`direction` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateFirewallRuleQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.SimpleAction", "."); ok {
		updateFirewallRuleQuery += "`simple_action` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.QosAction", "."); ok {
		updateFirewallRuleQuery += "`qos_action` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.UDPPort", "."); ok {
		updateFirewallRuleQuery += "`udp_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress", "."); ok {
		updateFirewallRuleQuery += "`vtep_dst_mac_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress", "."); ok {
		updateFirewallRuleQuery += "`vtep_dst_ip_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.StaticNHHeader.Vni", "."); ok {
		updateFirewallRuleQuery += "`vni` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.RoutingInstance", "."); ok {
		updateFirewallRuleQuery += "`routing_instance` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.NicAssistedMirroringVlan", "."); ok {
		updateFirewallRuleQuery += "`nic_assisted_mirroring_vlan` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.NicAssistedMirroring", "."); ok {
		updateFirewallRuleQuery += "`nic_assisted_mirroring` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.NHMode", "."); ok {
		updateFirewallRuleQuery += "`nh_mode` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.JuniperHeader", "."); ok {
		updateFirewallRuleQuery += "`juniper_header` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.Encapsulation", "."); ok {
		updateFirewallRuleQuery += "`encapsulation` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.AnalyzerName", "."); ok {
		updateFirewallRuleQuery += "`analyzer_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.AnalyzerMacAddress", "."); ok {
		updateFirewallRuleQuery += "`analyzer_mac_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.MirrorTo.AnalyzerIPAddress", "."); ok {
		updateFirewallRuleQuery += "`analyzer_ip_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.Log", "."); ok {
		updateFirewallRuleQuery += "`log` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.GatewayName", "."); ok {
		updateFirewallRuleQuery += "`gateway_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.AssignRoutingInstance", "."); ok {
		updateFirewallRuleQuery += "`assign_routing_instance` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.ApplyService", "."); ok {
		updateFirewallRuleQuery += "`apply_service` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateFirewallRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ActionList.Alert", "."); ok {
		updateFirewallRuleQuery += "`alert` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateFirewallRuleQuery += ","
	}

	updateFirewallRuleQuery =
		updateFirewallRuleQuery[:len(updateFirewallRuleQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateFirewallRuleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateFirewallRuleQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	if value, ok := common.GetValueByPath(model, "ServiceGroupRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_firewall_rule_service_group` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_firewall_rule_service_group` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref ServiceGroupRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_firewall_rule_service_group` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing ServiceGroupRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "ServiceGroupRefs update failed")
			}
		}
	}

	if value, ok := common.GetValueByPath(model, "AddressGroupRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_firewall_rule_address_group` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_firewall_rule_address_group` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref AddressGroupRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_firewall_rule_address_group` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing AddressGroupRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "AddressGroupRefs update failed")
			}
		}
	}

	if value, ok := common.GetValueByPath(model, "SecurityLoggingObjectRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_firewall_rule_security_logging_object` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_firewall_rule_security_logging_object` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref SecurityLoggingObjectRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_firewall_rule_security_logging_object` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing SecurityLoggingObjectRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "SecurityLoggingObjectRefs update failed")
			}
		}
	}

	if value, ok := common.GetValueByPath(model, "VirtualNetworkRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_firewall_rule_virtual_network` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_firewall_rule_virtual_network` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref VirtualNetworkRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_firewall_rule_virtual_network` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing VirtualNetworkRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "VirtualNetworkRefs update failed")
			}
		}
	}

	share, ok := common.GetValueByPath(model, ".Perms2.Share", ".")
	if ok {
		err = common.UpdateSharing(tx, "firewall_rule", string(uuid), share.([]interface{}))
		if err != nil {
			return err
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteFirewallRule deletes a resource
func DeleteFirewallRule(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteFirewallRuleQuery
	selectQuery := "select count(uuid) from firewall_rule where uuid = ?"
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
