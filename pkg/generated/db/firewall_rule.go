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

	"virtual_network": []string{
	// <schema.Schema Value>

	},

	"service_group": []string{
	// <schema.Schema Value>

	},

	"address_group": []string{
	// <schema.Schema Value>

	},

	"security_logging_object": []string{
	// <schema.Schema Value>

	},
}

// FirewallRuleBackRefFields is db back reference fields for FirewallRule
var FirewallRuleBackRefFields = map[string][]string{}

// FirewallRuleParentTypes is possible parents for FirewallRule
var FirewallRuleParents = []string{

	"project",

	"policy_management",
}

const insertFirewallRuleVirtualNetworkQuery = "insert into `ref_firewall_rule_virtual_network` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleServiceGroupQuery = "insert into `ref_firewall_rule_service_group` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleAddressGroupQuery = "insert into `ref_firewall_rule_address_group` (`from`, `to` ) values (?, ?);"

const insertFirewallRuleSecurityLoggingObjectQuery = "insert into `ref_firewall_rule_security_logging_object` (`from`, `to` ) values (?, ?);"

// CreateFirewallRule inserts FirewallRule to DB
func CreateFirewallRule(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateFirewallRuleRequest) error {
	model := request.FirewallRule
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
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		int(model.GetService().GetSRCPorts().GetStartPort()),
		int(model.GetService().GetSRCPorts().GetEndPort()),
		int(model.GetService().GetProtocolID()),
		string(model.GetService().GetProtocol()),
		int(model.GetService().GetDSTPorts().GetStartPort()),
		int(model.GetService().GetDSTPorts().GetEndPort()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		common.MustJSON(model.GetMatchTags().GetTagList()),
		common.MustJSON(model.GetMatchTagTypes().GetTagType()),
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
		string(model.GetEndpoint2().GetVirtualNetwork()),
		common.MustJSON(model.GetEndpoint2().GetTags()),
		common.MustJSON(model.GetEndpoint2().GetTagIds()),
		int(model.GetEndpoint2().GetSubnet().GetIPPrefixLen()),
		string(model.GetEndpoint2().GetSubnet().GetIPPrefix()),
		bool(model.GetEndpoint2().GetAny()),
		string(model.GetEndpoint2().GetAddressGroup()),
		string(model.GetEndpoint1().GetVirtualNetwork()),
		common.MustJSON(model.GetEndpoint1().GetTags()),
		common.MustJSON(model.GetEndpoint1().GetTagIds()),
		int(model.GetEndpoint1().GetSubnet().GetIPPrefixLen()),
		string(model.GetEndpoint1().GetSubnet().GetIPPrefix()),
		bool(model.GetEndpoint1().GetAny()),
		string(model.GetEndpoint1().GetAddressGroup()),
		string(model.GetDisplayName()),
		string(model.GetDirection()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()),
		string(model.GetActionList().GetSimpleAction()),
		string(model.GetActionList().GetQosAction()),
		int(model.GetActionList().GetMirrorTo().GetUDPPort()),
		string(model.GetActionList().GetMirrorTo().GetStaticNHHeader().GetVtepDSTMacAddress()),
		string(model.GetActionList().GetMirrorTo().GetStaticNHHeader().GetVtepDSTIPAddress()),
		int(model.GetActionList().GetMirrorTo().GetStaticNHHeader().GetVni()),
		string(model.GetActionList().GetMirrorTo().GetRoutingInstance()),
		int(model.GetActionList().GetMirrorTo().GetNicAssistedMirroringVlan()),
		bool(model.GetActionList().GetMirrorTo().GetNicAssistedMirroring()),
		string(model.GetActionList().GetMirrorTo().GetNHMode()),
		bool(model.GetActionList().GetMirrorTo().GetJuniperHeader()),
		string(model.GetActionList().GetMirrorTo().GetEncapsulation()),
		string(model.GetActionList().GetMirrorTo().GetAnalyzerName()),
		string(model.GetActionList().GetMirrorTo().GetAnalyzerMacAddress()),
		string(model.GetActionList().GetMirrorTo().GetAnalyzerIPAddress()),
		bool(model.GetActionList().GetLog()),
		string(model.GetActionList().GetGatewayName()),
		string(model.GetActionList().GetAssignRoutingInstance()),
		common.MustJSON(model.GetActionList().GetApplyService()),
		bool(model.GetActionList().GetAlert()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceGroupRef, err := tx.Prepare(insertFirewallRuleServiceGroupQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceGroupRefs create statement failed")
	}
	defer stmtServiceGroupRef.Close()
	for _, ref := range model.ServiceGroupRefs {

		_, err = stmtServiceGroupRef.ExecContext(ctx, model.UUID, ref.UUID)
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

		_, err = stmtAddressGroupRef.ExecContext(ctx, model.UUID, ref.UUID)
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

		_, err = stmtSecurityLoggingObjectRef.ExecContext(ctx, model.UUID, ref.UUID)
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

		_, err = stmtVirtualNetworkRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
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
	err = common.CreateSharing(tx, "firewall_rule", model.UUID, model.GetPerms2().GetShare())
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

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["start_port"]; ok {

		m.Service.SRCPorts.StartPort = schema.InterfaceToInt64(value)

	}

	if value, ok := values["end_port"]; ok {

		m.Service.SRCPorts.EndPort = schema.InterfaceToInt64(value)

	}

	if value, ok := values["protocol_id"]; ok {

		m.Service.ProtocolID = schema.InterfaceToInt64(value)

	}

	if value, ok := values["protocol"]; ok {

		m.Service.Protocol = schema.InterfaceToString(value)

	}

	if value, ok := values["dst_ports_start_port"]; ok {

		m.Service.DSTPorts.StartPort = schema.InterfaceToInt64(value)

	}

	if value, ok := values["dst_ports_end_port"]; ok {

		m.Service.DSTPorts.EndPort = schema.InterfaceToInt64(value)

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

	if value, ok := values["tag_list"]; ok {

		json.Unmarshal(value.([]byte), &m.MatchTags.TagList)

	}

	if value, ok := values["tag_type"]; ok {

		json.Unmarshal(value.([]byte), &m.MatchTagTypes.TagType)

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

	if value, ok := values["virtual_network"]; ok {

		m.Endpoint2.VirtualNetwork = schema.InterfaceToString(value)

	}

	if value, ok := values["tags"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint2.Tags)

	}

	if value, ok := values["tag_ids"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint2.TagIds)

	}

	if value, ok := values["ip_prefix_len"]; ok {

		m.Endpoint2.Subnet.IPPrefixLen = schema.InterfaceToInt64(value)

	}

	if value, ok := values["ip_prefix"]; ok {

		m.Endpoint2.Subnet.IPPrefix = schema.InterfaceToString(value)

	}

	if value, ok := values["any"]; ok {

		m.Endpoint2.Any = schema.InterfaceToBool(value)

	}

	if value, ok := values["address_group"]; ok {

		m.Endpoint2.AddressGroup = schema.InterfaceToString(value)

	}

	if value, ok := values["endpoint_1_virtual_network"]; ok {

		m.Endpoint1.VirtualNetwork = schema.InterfaceToString(value)

	}

	if value, ok := values["endpoint_1_tags"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint1.Tags)

	}

	if value, ok := values["endpoint_1_tag_ids"]; ok {

		json.Unmarshal(value.([]byte), &m.Endpoint1.TagIds)

	}

	if value, ok := values["subnet_ip_prefix_len"]; ok {

		m.Endpoint1.Subnet.IPPrefixLen = schema.InterfaceToInt64(value)

	}

	if value, ok := values["subnet_ip_prefix"]; ok {

		m.Endpoint1.Subnet.IPPrefix = schema.InterfaceToString(value)

	}

	if value, ok := values["endpoint_1_any"]; ok {

		m.Endpoint1.Any = schema.InterfaceToBool(value)

	}

	if value, ok := values["endpoint_1_address_group"]; ok {

		m.Endpoint1.AddressGroup = schema.InterfaceToString(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["direction"]; ok {

		m.Direction = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["simple_action"]; ok {

		m.ActionList.SimpleAction = schema.InterfaceToString(value)

	}

	if value, ok := values["qos_action"]; ok {

		m.ActionList.QosAction = schema.InterfaceToString(value)

	}

	if value, ok := values["udp_port"]; ok {

		m.ActionList.MirrorTo.UDPPort = schema.InterfaceToInt64(value)

	}

	if value, ok := values["vtep_dst_mac_address"]; ok {

		m.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["vtep_dst_ip_address"]; ok {

		m.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["vni"]; ok {

		m.ActionList.MirrorTo.StaticNHHeader.Vni = schema.InterfaceToInt64(value)

	}

	if value, ok := values["routing_instance"]; ok {

		m.ActionList.MirrorTo.RoutingInstance = schema.InterfaceToString(value)

	}

	if value, ok := values["nic_assisted_mirroring_vlan"]; ok {

		m.ActionList.MirrorTo.NicAssistedMirroringVlan = schema.InterfaceToInt64(value)

	}

	if value, ok := values["nic_assisted_mirroring"]; ok {

		m.ActionList.MirrorTo.NicAssistedMirroring = schema.InterfaceToBool(value)

	}

	if value, ok := values["nh_mode"]; ok {

		m.ActionList.MirrorTo.NHMode = schema.InterfaceToString(value)

	}

	if value, ok := values["juniper_header"]; ok {

		m.ActionList.MirrorTo.JuniperHeader = schema.InterfaceToBool(value)

	}

	if value, ok := values["encapsulation"]; ok {

		m.ActionList.MirrorTo.Encapsulation = schema.InterfaceToString(value)

	}

	if value, ok := values["analyzer_name"]; ok {

		m.ActionList.MirrorTo.AnalyzerName = schema.InterfaceToString(value)

	}

	if value, ok := values["analyzer_mac_address"]; ok {

		m.ActionList.MirrorTo.AnalyzerMacAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["analyzer_ip_address"]; ok {

		m.ActionList.MirrorTo.AnalyzerIPAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["log"]; ok {

		m.ActionList.Log = schema.InterfaceToBool(value)

	}

	if value, ok := values["gateway_name"]; ok {

		m.ActionList.GatewayName = schema.InterfaceToString(value)

	}

	if value, ok := values["assign_routing_instance"]; ok {

		m.ActionList.AssignRoutingInstance = schema.InterfaceToString(value)

	}

	if value, ok := values["apply_service"]; ok {

		json.Unmarshal(value.([]byte), &m.ActionList.ApplyService)

	}

	if value, ok := values["alert"]; ok {

		m.ActionList.Alert = schema.InterfaceToBool(value)

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
			referenceModel := &models.FirewallRuleSecurityLoggingObjectRef{}
			referenceModel.UUID = uuid
			m.SecurityLoggingObjectRefs = append(m.SecurityLoggingObjectRefs, referenceModel)

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
			referenceModel := &models.FirewallRuleVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_group"]; ok {
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
			referenceModel := &models.FirewallRuleServiceGroupRef{}
			referenceModel.UUID = uuid
			m.ServiceGroupRefs = append(m.ServiceGroupRefs, referenceModel)

		}
	}

	if value, ok := values["ref_address_group"]; ok {
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
			referenceModel := &models.FirewallRuleAddressGroupRef{}
			referenceModel.UUID = uuid
			m.AddressGroupRefs = append(m.AddressGroupRefs, referenceModel)

		}
	}

	return m, nil
}

// ListFirewallRule lists FirewallRule with list spec.
func ListFirewallRule(ctx context.Context, tx *sql.Tx, request *models.ListFirewallRuleRequest) (response *models.ListFirewallRuleResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "firewall_rule"
	qb.Fields = FirewallRuleFields
	qb.RefFields = FirewallRuleRefFields
	qb.BackRefFields = FirewallRuleBackRefFields
	result := []*models.FirewallRule{}

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
		m, err := scanFirewallRule(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListFirewallRuleResponse{
		FirewallRules: result,
	}
	return response, nil
}

// UpdateFirewallRule updates a resource
func UpdateFirewallRule(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateFirewallRuleRequest,
) error {
	//TODO
	return nil
}

// DeleteFirewallRule deletes a resource
func DeleteFirewallRule(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteFirewallRuleRequest) error {
	deleteQuery := deleteFirewallRuleQuery
	selectQuery := "select count(uuid) from firewall_rule where uuid = ?"
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
