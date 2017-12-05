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

const insertPhysicalRouterQuery = "insert into `physical_router` (`retries`,`v3_privacy_password`,`v3_authentication_password`,`v3_context_engine_id`,`local_port`,`v3_security_level`,`v3_authentication_protocol`,`timeout`,`v3_security_name`,`v3_security_engine_id`,`v3_context`,`version`,`v3_privacy_protocol`,`v3_engine_boots`,`v3_engine_id`,`v2_community`,`v3_engine_time`,`physical_router_loopback_ip`,`physical_router_dataplane_ip`,`owner`,`owner_access`,`global_access`,`share`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`physical_router_management_ip`,`physical_router_vendor_name`,`physical_router_product_name`,`physical_router_snmp`,`uuid`,`physical_router_role`,`resource`,`server_port`,`server_ip`,`key_value_pair`,`fq_name`,`username`,`password`,`physical_router_vnc_managed`,`physical_router_lldp`,`physical_router_image_uri`,`service_port`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePhysicalRouterQuery = "update `physical_router` set `retries` = ?,`v3_privacy_password` = ?,`v3_authentication_password` = ?,`v3_context_engine_id` = ?,`local_port` = ?,`v3_security_level` = ?,`v3_authentication_protocol` = ?,`timeout` = ?,`v3_security_name` = ?,`v3_security_engine_id` = ?,`v3_context` = ?,`version` = ?,`v3_privacy_protocol` = ?,`v3_engine_boots` = ?,`v3_engine_id` = ?,`v2_community` = ?,`v3_engine_time` = ?,`physical_router_loopback_ip` = ?,`physical_router_dataplane_ip` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`physical_router_management_ip` = ?,`physical_router_vendor_name` = ?,`physical_router_product_name` = ?,`physical_router_snmp` = ?,`uuid` = ?,`physical_router_role` = ?,`resource` = ?,`server_port` = ?,`server_ip` = ?,`key_value_pair` = ?,`fq_name` = ?,`username` = ?,`password` = ?,`physical_router_vnc_managed` = ?,`physical_router_lldp` = ?,`physical_router_image_uri` = ?,`service_port` = ?,`display_name` = ?;"
const deletePhysicalRouterQuery = "delete from `physical_router` where uuid = ?"
const listPhysicalRouterQuery = "select `physical_router`.`retries`,`physical_router`.`v3_privacy_password`,`physical_router`.`v3_authentication_password`,`physical_router`.`v3_context_engine_id`,`physical_router`.`local_port`,`physical_router`.`v3_security_level`,`physical_router`.`v3_authentication_protocol`,`physical_router`.`timeout`,`physical_router`.`v3_security_name`,`physical_router`.`v3_security_engine_id`,`physical_router`.`v3_context`,`physical_router`.`version`,`physical_router`.`v3_privacy_protocol`,`physical_router`.`v3_engine_boots`,`physical_router`.`v3_engine_id`,`physical_router`.`v2_community`,`physical_router`.`v3_engine_time`,`physical_router`.`physical_router_loopback_ip`,`physical_router`.`physical_router_dataplane_ip`,`physical_router`.`owner`,`physical_router`.`owner_access`,`physical_router`.`global_access`,`physical_router`.`share`,`physical_router`.`created`,`physical_router`.`creator`,`physical_router`.`user_visible`,`physical_router`.`last_modified`,`physical_router`.`permissions_owner_access`,`physical_router`.`other_access`,`physical_router`.`group`,`physical_router`.`group_access`,`physical_router`.`permissions_owner`,`physical_router`.`enable`,`physical_router`.`description`,`physical_router`.`physical_router_management_ip`,`physical_router`.`physical_router_vendor_name`,`physical_router`.`physical_router_product_name`,`physical_router`.`physical_router_snmp`,`physical_router`.`uuid`,`physical_router`.`physical_router_role`,`physical_router`.`resource`,`physical_router`.`server_port`,`physical_router`.`server_ip`,`physical_router`.`key_value_pair`,`physical_router`.`fq_name`,`physical_router`.`username`,`physical_router`.`password`,`physical_router`.`physical_router_vnc_managed`,`physical_router`.`physical_router_lldp`,`physical_router`.`physical_router_image_uri`,`physical_router`.`service_port`,`physical_router`.`display_name` from `physical_router`"
const showPhysicalRouterQuery = "select `physical_router`.`retries`,`physical_router`.`v3_privacy_password`,`physical_router`.`v3_authentication_password`,`physical_router`.`v3_context_engine_id`,`physical_router`.`local_port`,`physical_router`.`v3_security_level`,`physical_router`.`v3_authentication_protocol`,`physical_router`.`timeout`,`physical_router`.`v3_security_name`,`physical_router`.`v3_security_engine_id`,`physical_router`.`v3_context`,`physical_router`.`version`,`physical_router`.`v3_privacy_protocol`,`physical_router`.`v3_engine_boots`,`physical_router`.`v3_engine_id`,`physical_router`.`v2_community`,`physical_router`.`v3_engine_time`,`physical_router`.`physical_router_loopback_ip`,`physical_router`.`physical_router_dataplane_ip`,`physical_router`.`owner`,`physical_router`.`owner_access`,`physical_router`.`global_access`,`physical_router`.`share`,`physical_router`.`created`,`physical_router`.`creator`,`physical_router`.`user_visible`,`physical_router`.`last_modified`,`physical_router`.`permissions_owner_access`,`physical_router`.`other_access`,`physical_router`.`group`,`physical_router`.`group_access`,`physical_router`.`permissions_owner`,`physical_router`.`enable`,`physical_router`.`description`,`physical_router`.`physical_router_management_ip`,`physical_router`.`physical_router_vendor_name`,`physical_router`.`physical_router_product_name`,`physical_router`.`physical_router_snmp`,`physical_router`.`uuid`,`physical_router`.`physical_router_role`,`physical_router`.`resource`,`physical_router`.`server_port`,`physical_router`.`server_ip`,`physical_router`.`key_value_pair`,`physical_router`.`fq_name`,`physical_router`.`username`,`physical_router`.`password`,`physical_router`.`physical_router_vnc_managed`,`physical_router`.`physical_router_lldp`,`physical_router`.`physical_router_image_uri`,`physical_router`.`service_port`,`physical_router`.`display_name` from `physical_router` where uuid = ?"

const insertPhysicalRouterVirtualRouterQuery = "insert into `ref_physical_router_virtual_router` (`from`, `to` ) values (?, ?);"

const insertPhysicalRouterVirtualNetworkQuery = "insert into `ref_physical_router_virtual_network` (`from`, `to` ) values (?, ?);"

const insertPhysicalRouterBGPRouterQuery = "insert into `ref_physical_router_bgp_router` (`from`, `to` ) values (?, ?);"

func CreatePhysicalRouter(tx *sql.Tx, model *models.PhysicalRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.PhysicalRouterSNMPCredentials.Retries),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyPassword),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationPassword),
		string(model.PhysicalRouterSNMPCredentials.V3ContextEngineID),
		int(model.PhysicalRouterSNMPCredentials.LocalPort),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityLevel),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol),
		int(model.PhysicalRouterSNMPCredentials.Timeout),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityName),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityEngineID),
		string(model.PhysicalRouterSNMPCredentials.V3Context),
		int(model.PhysicalRouterSNMPCredentials.Version),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyProtocol),
		int(model.PhysicalRouterSNMPCredentials.V3EngineBoots),
		string(model.PhysicalRouterSNMPCredentials.V3EngineID),
		string(model.PhysicalRouterSNMPCredentials.V2Community),
		int(model.PhysicalRouterSNMPCredentials.V3EngineTime),
		string(model.PhysicalRouterLoopbackIP),
		string(model.PhysicalRouterDataplaneIP),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.PhysicalRouterManagementIP),
		string(model.PhysicalRouterVendorName),
		string(model.PhysicalRouterProductName),
		bool(model.PhysicalRouterSNMP),
		string(model.UUID),
		string(model.PhysicalRouterRole),
		utils.MustJSON(model.TelemetryInfo.Resource),
		int(model.TelemetryInfo.ServerPort),
		string(model.TelemetryInfo.ServerIP),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.FQName),
		string(model.PhysicalRouterUserCredentials.Username),
		string(model.PhysicalRouterUserCredentials.Password),
		bool(model.PhysicalRouterVNCManaged),
		bool(model.PhysicalRouterLLDP),
		string(model.PhysicalRouterImageURI),
		utils.MustJSON(model.PhysicalRouterJunosServicePorts.ServicePort),
		string(model.DisplayName))

	stmtVirtualNetworkRef, err := tx.Prepare(insertPhysicalRouterVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {
		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
	}

	stmtBGPRouterRef, err := tx.Prepare(insertPhysicalRouterBGPRouterQuery)
	if err != nil {
		return err
	}
	defer stmtBGPRouterRef.Close()
	for _, ref := range model.BGPRouterRefs {
		_, err = stmtBGPRouterRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualRouterRef, err := tx.Prepare(insertPhysicalRouterVirtualRouterQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualRouterRef.Close()
	for _, ref := range model.VirtualRouterRefs {
		_, err = stmtVirtualRouterRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanPhysicalRouter(rows *sql.Rows) (*models.PhysicalRouter, error) {
	m := models.MakePhysicalRouter()

	var jsonPerms2Share string

	var jsonTelemetryInfoResource string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonPhysicalRouterJunosServicePortsServicePort string

	if err := rows.Scan(&m.PhysicalRouterSNMPCredentials.Retries,
		&m.PhysicalRouterSNMPCredentials.V3PrivacyPassword,
		&m.PhysicalRouterSNMPCredentials.V3AuthenticationPassword,
		&m.PhysicalRouterSNMPCredentials.V3ContextEngineID,
		&m.PhysicalRouterSNMPCredentials.LocalPort,
		&m.PhysicalRouterSNMPCredentials.V3SecurityLevel,
		&m.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol,
		&m.PhysicalRouterSNMPCredentials.Timeout,
		&m.PhysicalRouterSNMPCredentials.V3SecurityName,
		&m.PhysicalRouterSNMPCredentials.V3SecurityEngineID,
		&m.PhysicalRouterSNMPCredentials.V3Context,
		&m.PhysicalRouterSNMPCredentials.Version,
		&m.PhysicalRouterSNMPCredentials.V3PrivacyProtocol,
		&m.PhysicalRouterSNMPCredentials.V3EngineBoots,
		&m.PhysicalRouterSNMPCredentials.V3EngineID,
		&m.PhysicalRouterSNMPCredentials.V2Community,
		&m.PhysicalRouterSNMPCredentials.V3EngineTime,
		&m.PhysicalRouterLoopbackIP,
		&m.PhysicalRouterDataplaneIP,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.PhysicalRouterManagementIP,
		&m.PhysicalRouterVendorName,
		&m.PhysicalRouterProductName,
		&m.PhysicalRouterSNMP,
		&m.UUID,
		&m.PhysicalRouterRole,
		&jsonTelemetryInfoResource,
		&m.TelemetryInfo.ServerPort,
		&m.TelemetryInfo.ServerIP,
		&jsonAnnotationsKeyValuePair,
		&jsonFQName,
		&m.PhysicalRouterUserCredentials.Username,
		&m.PhysicalRouterUserCredentials.Password,
		&m.PhysicalRouterVNCManaged,
		&m.PhysicalRouterLLDP,
		&m.PhysicalRouterImageURI,
		&jsonPhysicalRouterJunosServicePortsServicePort,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonTelemetryInfoResource), &m.TelemetryInfo.Resource)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPhysicalRouterJunosServicePortsServicePort), &m.PhysicalRouterJunosServicePorts.ServicePort)

	return m, nil
}

func buildPhysicalRouterWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["v3_privacy_password"]; ok {
		results = append(results, "v3_privacy_password = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_authentication_password"]; ok {
		results = append(results, "v3_authentication_password = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_context_engine_id"]; ok {
		results = append(results, "v3_context_engine_id = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_security_level"]; ok {
		results = append(results, "v3_security_level = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_authentication_protocol"]; ok {
		results = append(results, "v3_authentication_protocol = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_security_name"]; ok {
		results = append(results, "v3_security_name = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_security_engine_id"]; ok {
		results = append(results, "v3_security_engine_id = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_context"]; ok {
		results = append(results, "v3_context = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_privacy_protocol"]; ok {
		results = append(results, "v3_privacy_protocol = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_engine_id"]; ok {
		results = append(results, "v3_engine_id = ?")
		values = append(values, value)
	}

	if value, ok := where["v2_community"]; ok {
		results = append(results, "v2_community = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_loopback_ip"]; ok {
		results = append(results, "physical_router_loopback_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_dataplane_ip"]; ok {
		results = append(results, "physical_router_dataplane_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_management_ip"]; ok {
		results = append(results, "physical_router_management_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_vendor_name"]; ok {
		results = append(results, "physical_router_vendor_name = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_product_name"]; ok {
		results = append(results, "physical_router_product_name = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_role"]; ok {
		results = append(results, "physical_router_role = ?")
		values = append(values, value)
	}

	if value, ok := where["server_ip"]; ok {
		results = append(results, "server_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["username"]; ok {
		results = append(results, "username = ?")
		values = append(values, value)
	}

	if value, ok := where["password"]; ok {
		results = append(results, "password = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_image_uri"]; ok {
		results = append(results, "physical_router_image_uri = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListPhysicalRouter(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.PhysicalRouter, error) {
	result := models.MakePhysicalRouterSlice()
	whereQuery, values := buildPhysicalRouterWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listPhysicalRouterQuery)
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
		m, _ := scanPhysicalRouter(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowPhysicalRouter(tx *sql.Tx, uuid string) (*models.PhysicalRouter, error) {
	rows, err := tx.Query(showPhysicalRouterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanPhysicalRouter(rows)
	}
	return nil, nil
}

func UpdatePhysicalRouter(tx *sql.Tx, uuid string, model *models.PhysicalRouter) error {
	return nil
}

func DeletePhysicalRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deletePhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
