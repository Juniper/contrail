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

const insertVirtualNetworkQuery = "insert into `virtual_network` (`physical_network`,`segmentation_id`,`mac_learning_enabled`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`uuid`,`route_target`,`key_value_pair`,`pbb_etree_enable`,`port_security_enabled`,`display_name`,`mac_limit`,`mac_limit_action`,`router_external`,`flood_unknown_unicast`,`external_ipam`,`multi_policy_service_chains_enabled`,`virtual_network_network_id`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`fq_name`,`is_shared`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`pbb_evpn_enable`,`mac_aging_time`,`export_route_target_list_route_target`,`layer2_control_word`,`vxlan_network_identifier`,`rpf`,`forwarding_mode`,`allow_transit`,`network_id`,`mirror_destination`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`address_allocation_mode`,`import_route_target_list_route_target`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualNetworkQuery = "update `virtual_network` set `physical_network` = ?,`segmentation_id` = ?,`mac_learning_enabled` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`uuid` = ?,`route_target` = ?,`key_value_pair` = ?,`pbb_etree_enable` = ?,`port_security_enabled` = ?,`display_name` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`router_external` = ?,`flood_unknown_unicast` = ?,`external_ipam` = ?,`multi_policy_service_chains_enabled` = ?,`virtual_network_network_id` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`mac_move_time_window` = ?,`fq_name` = ?,`is_shared` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`pbb_evpn_enable` = ?,`mac_aging_time` = ?,`export_route_target_list_route_target` = ?,`layer2_control_word` = ?,`vxlan_network_identifier` = ?,`rpf` = ?,`forwarding_mode` = ?,`allow_transit` = ?,`network_id` = ?,`mirror_destination` = ?,`destination_ip` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`address_allocation_mode` = ?,`import_route_target_list_route_target` = ?;"
const deleteVirtualNetworkQuery = "delete from `virtual_network` where uuid = ?"
const listVirtualNetworkQuery = "select `virtual_network`.`physical_network`,`virtual_network`.`segmentation_id`,`virtual_network`.`mac_learning_enabled`,`virtual_network`.`other_access`,`virtual_network`.`group`,`virtual_network`.`group_access`,`virtual_network`.`owner`,`virtual_network`.`owner_access`,`virtual_network`.`enable`,`virtual_network`.`description`,`virtual_network`.`created`,`virtual_network`.`creator`,`virtual_network`.`user_visible`,`virtual_network`.`last_modified`,`virtual_network`.`uuid`,`virtual_network`.`route_target`,`virtual_network`.`key_value_pair`,`virtual_network`.`pbb_etree_enable`,`virtual_network`.`port_security_enabled`,`virtual_network`.`display_name`,`virtual_network`.`mac_limit`,`virtual_network`.`mac_limit_action`,`virtual_network`.`router_external`,`virtual_network`.`flood_unknown_unicast`,`virtual_network`.`external_ipam`,`virtual_network`.`multi_policy_service_chains_enabled`,`virtual_network`.`virtual_network_network_id`,`virtual_network`.`mac_move_limit`,`virtual_network`.`mac_move_limit_action`,`virtual_network`.`mac_move_time_window`,`virtual_network`.`fq_name`,`virtual_network`.`is_shared`,`virtual_network`.`perms2_owner`,`virtual_network`.`perms2_owner_access`,`virtual_network`.`global_access`,`virtual_network`.`share`,`virtual_network`.`pbb_evpn_enable`,`virtual_network`.`mac_aging_time`,`virtual_network`.`export_route_target_list_route_target`,`virtual_network`.`layer2_control_word`,`virtual_network`.`vxlan_network_identifier`,`virtual_network`.`rpf`,`virtual_network`.`forwarding_mode`,`virtual_network`.`allow_transit`,`virtual_network`.`network_id`,`virtual_network`.`mirror_destination`,`virtual_network`.`destination_ip`,`virtual_network`.`ip_protocol`,`virtual_network`.`source_ip`,`virtual_network`.`hashing_configured`,`virtual_network`.`source_port`,`virtual_network`.`destination_port`,`virtual_network`.`address_allocation_mode`,`virtual_network`.`import_route_target_list_route_target` from `virtual_network`"
const showVirtualNetworkQuery = "select `virtual_network`.`physical_network`,`virtual_network`.`segmentation_id`,`virtual_network`.`mac_learning_enabled`,`virtual_network`.`other_access`,`virtual_network`.`group`,`virtual_network`.`group_access`,`virtual_network`.`owner`,`virtual_network`.`owner_access`,`virtual_network`.`enable`,`virtual_network`.`description`,`virtual_network`.`created`,`virtual_network`.`creator`,`virtual_network`.`user_visible`,`virtual_network`.`last_modified`,`virtual_network`.`uuid`,`virtual_network`.`route_target`,`virtual_network`.`key_value_pair`,`virtual_network`.`pbb_etree_enable`,`virtual_network`.`port_security_enabled`,`virtual_network`.`display_name`,`virtual_network`.`mac_limit`,`virtual_network`.`mac_limit_action`,`virtual_network`.`router_external`,`virtual_network`.`flood_unknown_unicast`,`virtual_network`.`external_ipam`,`virtual_network`.`multi_policy_service_chains_enabled`,`virtual_network`.`virtual_network_network_id`,`virtual_network`.`mac_move_limit`,`virtual_network`.`mac_move_limit_action`,`virtual_network`.`mac_move_time_window`,`virtual_network`.`fq_name`,`virtual_network`.`is_shared`,`virtual_network`.`perms2_owner`,`virtual_network`.`perms2_owner_access`,`virtual_network`.`global_access`,`virtual_network`.`share`,`virtual_network`.`pbb_evpn_enable`,`virtual_network`.`mac_aging_time`,`virtual_network`.`export_route_target_list_route_target`,`virtual_network`.`layer2_control_word`,`virtual_network`.`vxlan_network_identifier`,`virtual_network`.`rpf`,`virtual_network`.`forwarding_mode`,`virtual_network`.`allow_transit`,`virtual_network`.`network_id`,`virtual_network`.`mirror_destination`,`virtual_network`.`destination_ip`,`virtual_network`.`ip_protocol`,`virtual_network`.`source_ip`,`virtual_network`.`hashing_configured`,`virtual_network`.`source_port`,`virtual_network`.`destination_port`,`virtual_network`.`address_allocation_mode`,`virtual_network`.`import_route_target_list_route_target` from `virtual_network` where uuid = ?"

const insertVirtualNetworkNetworkIpamQuery = "insert into `ref_virtual_network_network_ipam` (`from`, `to` ,`ipam_subnets`,`route`) values (?, ?,?,?);"

const insertVirtualNetworkSecurityLoggingObjectQuery = "insert into `ref_virtual_network_security_logging_object` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkNetworkPolicyQuery = "insert into `ref_virtual_network_network_policy` (`from`, `to` ,`end_time`,`start_time`,`off_interval`,`on_interval`,`minor`,`major`) values (?, ?,?,?,?,?,?,?);"

const insertVirtualNetworkQosConfigQuery = "insert into `ref_virtual_network_qos_config` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkRouteTableQuery = "insert into `ref_virtual_network_route_table` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkVirtualNetworkQuery = "insert into `ref_virtual_network_virtual_network` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkBGPVPNQuery = "insert into `ref_virtual_network_bgpvpn` (`from`, `to` ) values (?, ?);"

func CreateVirtualNetwork(tx *sql.Tx, model *models.VirtualNetwork) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ProviderProperties.PhysicalNetwork),
		int(model.ProviderProperties.SegmentationID),
		bool(model.MacLearningEnabled),
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
		string(model.UUID),
		utils.MustJSON(model.RouteTargetList.RouteTarget),
		utils.MustJSON(model.Annotations.KeyValuePair),
		bool(model.PBBEtreeEnable),
		bool(model.PortSecurityEnabled),
		string(model.DisplayName),
		int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		bool(model.RouterExternal),
		bool(model.FloodUnknownUnicast),
		bool(model.ExternalIpam),
		bool(model.MultiPolicyServiceChainsEnabled),
		int(model.VirtualNetworkNetworkID),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.MacMoveControl.MacMoveTimeWindow),
		utils.MustJSON(model.FQName),
		bool(model.IsShared),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		bool(model.PBBEvpnEnable),
		int(model.MacAgingTime),
		utils.MustJSON(model.ExportRouteTargetList.RouteTarget),
		bool(model.Layer2ControlWord),
		int(model.VirtualNetworkProperties.VxlanNetworkIdentifier),
		string(model.VirtualNetworkProperties.RPF),
		string(model.VirtualNetworkProperties.ForwardingMode),
		bool(model.VirtualNetworkProperties.AllowTransit),
		int(model.VirtualNetworkProperties.NetworkID),
		bool(model.VirtualNetworkProperties.MirrorDestination),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		string(model.AddressAllocationMode),
		utils.MustJSON(model.ImportRouteTargetList.RouteTarget))

	stmtBGPVPNRef, err := tx.Prepare(insertVirtualNetworkBGPVPNQuery)
	if err != nil {
		return err
	}
	defer stmtBGPVPNRef.Close()
	for _, ref := range model.BGPVPNRefs {
		_, err = stmtBGPVPNRef.Exec(model.UUID, ref.UUID)
	}

	stmtNetworkIpamRef, err := tx.Prepare(insertVirtualNetworkNetworkIpamQuery)
	if err != nil {
		return err
	}
	defer stmtNetworkIpamRef.Close()
	for _, ref := range model.NetworkIpamRefs {
		_, err = stmtNetworkIpamRef.Exec(model.UUID, ref.UUID, utils.MustJSON(ref.Attr.IpamSubnets),
			utils.MustJSON(ref.Attr.HostRoutes.Route))
	}

	stmtSecurityLoggingObjectRef, err := tx.Prepare(insertVirtualNetworkSecurityLoggingObjectQuery)
	if err != nil {
		return err
	}
	defer stmtSecurityLoggingObjectRef.Close()
	for _, ref := range model.SecurityLoggingObjectRefs {
		_, err = stmtSecurityLoggingObjectRef.Exec(model.UUID, ref.UUID)
	}

	stmtNetworkPolicyRef, err := tx.Prepare(insertVirtualNetworkNetworkPolicyQuery)
	if err != nil {
		return err
	}
	defer stmtNetworkPolicyRef.Close()
	for _, ref := range model.NetworkPolicyRefs {
		_, err = stmtNetworkPolicyRef.Exec(model.UUID, ref.UUID, string(ref.Attr.Timer.EndTime),
			string(ref.Attr.Timer.StartTime),
			string(ref.Attr.Timer.OffInterval),
			string(ref.Attr.Timer.OnInterval),
			int(ref.Attr.Sequence.Minor),
			int(ref.Attr.Sequence.Major))
	}

	stmtQosConfigRef, err := tx.Prepare(insertVirtualNetworkQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmtQosConfigRef.Close()
	for _, ref := range model.QosConfigRefs {
		_, err = stmtQosConfigRef.Exec(model.UUID, ref.UUID)
	}

	stmtRouteTableRef, err := tx.Prepare(insertVirtualNetworkRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmtRouteTableRef.Close()
	for _, ref := range model.RouteTableRefs {
		_, err = stmtRouteTableRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertVirtualNetworkVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {
		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanVirtualNetwork(rows *sql.Rows) (*models.VirtualNetwork, error) {
	m := models.MakeVirtualNetwork()

	var jsonRouteTargetListRouteTarget string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonPerms2Share string

	var jsonExportRouteTargetListRouteTarget string

	var jsonImportRouteTargetListRouteTarget string

	if err := rows.Scan(&m.ProviderProperties.PhysicalNetwork,
		&m.ProviderProperties.SegmentationID,
		&m.MacLearningEnabled,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.UUID,
		&jsonRouteTargetListRouteTarget,
		&jsonAnnotationsKeyValuePair,
		&m.PBBEtreeEnable,
		&m.PortSecurityEnabled,
		&m.DisplayName,
		&m.MacLimitControl.MacLimit,
		&m.MacLimitControl.MacLimitAction,
		&m.RouterExternal,
		&m.FloodUnknownUnicast,
		&m.ExternalIpam,
		&m.MultiPolicyServiceChainsEnabled,
		&m.VirtualNetworkNetworkID,
		&m.MacMoveControl.MacMoveLimit,
		&m.MacMoveControl.MacMoveLimitAction,
		&m.MacMoveControl.MacMoveTimeWindow,
		&jsonFQName,
		&m.IsShared,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.PBBEvpnEnable,
		&m.MacAgingTime,
		&jsonExportRouteTargetListRouteTarget,
		&m.Layer2ControlWord,
		&m.VirtualNetworkProperties.VxlanNetworkIdentifier,
		&m.VirtualNetworkProperties.RPF,
		&m.VirtualNetworkProperties.ForwardingMode,
		&m.VirtualNetworkProperties.AllowTransit,
		&m.VirtualNetworkProperties.NetworkID,
		&m.VirtualNetworkProperties.MirrorDestination,
		&m.EcmpHashingIncludeFields.DestinationIP,
		&m.EcmpHashingIncludeFields.IPProtocol,
		&m.EcmpHashingIncludeFields.SourceIP,
		&m.EcmpHashingIncludeFields.HashingConfigured,
		&m.EcmpHashingIncludeFields.SourcePort,
		&m.EcmpHashingIncludeFields.DestinationPort,
		&m.AddressAllocationMode,
		&jsonImportRouteTargetListRouteTarget); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonRouteTargetListRouteTarget), &m.RouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonExportRouteTargetListRouteTarget), &m.ExportRouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonImportRouteTargetListRouteTarget), &m.ImportRouteTargetList.RouteTarget)

	return m, nil
}

func buildVirtualNetworkWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["physical_network"]; ok {
		results = append(results, "physical_network = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["mac_limit_action"]; ok {
		results = append(results, "mac_limit_action = ?")
		values = append(values, value)
	}

	if value, ok := where["mac_move_limit_action"]; ok {
		results = append(results, "mac_move_limit_action = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["rpf"]; ok {
		results = append(results, "rpf = ?")
		values = append(values, value)
	}

	if value, ok := where["forwarding_mode"]; ok {
		results = append(results, "forwarding_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["address_allocation_mode"]; ok {
		results = append(results, "address_allocation_mode = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualNetwork(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualNetwork, error) {
	result := models.MakeVirtualNetworkSlice()
	whereQuery, values := buildVirtualNetworkWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualNetworkQuery)
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
		m, _ := scanVirtualNetwork(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualNetwork(tx *sql.Tx, uuid string) (*models.VirtualNetwork, error) {
	rows, err := tx.Query(showVirtualNetworkQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualNetwork(rows)
	}
	return nil, nil
}

func UpdateVirtualNetwork(tx *sql.Tx, uuid string, model *models.VirtualNetwork) error {
	return nil
}

func DeleteVirtualNetwork(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
