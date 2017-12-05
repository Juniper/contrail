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

const insertGlobalVrouterConfigQuery = "insert into `global_vrouter_config` (`forwarding_mode`,`vxlan_network_identifier_mode`,`owner`,`owner_access`,`global_access`,`share`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`flow_aging_timeout`,`linklocal_service_entry`,`uuid`,`flow_export_rate`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name`,`encapsulation`,`enable_security_logging`,`fq_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalVrouterConfigQuery = "update `global_vrouter_config` set `forwarding_mode` = ?,`vxlan_network_identifier_mode` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`flow_aging_timeout` = ?,`linklocal_service_entry` = ?,`uuid` = ?,`flow_export_rate` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`encapsulation` = ?,`enable_security_logging` = ?,`fq_name` = ?,`key_value_pair` = ?;"
const deleteGlobalVrouterConfigQuery = "delete from `global_vrouter_config` where uuid = ?"
const listGlobalVrouterConfigQuery = "select `global_vrouter_config`.`forwarding_mode`,`global_vrouter_config`.`vxlan_network_identifier_mode`,`global_vrouter_config`.`owner`,`global_vrouter_config`.`owner_access`,`global_vrouter_config`.`global_access`,`global_vrouter_config`.`share`,`global_vrouter_config`.`ip_protocol`,`global_vrouter_config`.`source_ip`,`global_vrouter_config`.`hashing_configured`,`global_vrouter_config`.`source_port`,`global_vrouter_config`.`destination_port`,`global_vrouter_config`.`destination_ip`,`global_vrouter_config`.`flow_aging_timeout`,`global_vrouter_config`.`linklocal_service_entry`,`global_vrouter_config`.`uuid`,`global_vrouter_config`.`flow_export_rate`,`global_vrouter_config`.`creator`,`global_vrouter_config`.`user_visible`,`global_vrouter_config`.`last_modified`,`global_vrouter_config`.`permissions_owner_access`,`global_vrouter_config`.`other_access`,`global_vrouter_config`.`group`,`global_vrouter_config`.`group_access`,`global_vrouter_config`.`permissions_owner`,`global_vrouter_config`.`enable`,`global_vrouter_config`.`description`,`global_vrouter_config`.`created`,`global_vrouter_config`.`display_name`,`global_vrouter_config`.`encapsulation`,`global_vrouter_config`.`enable_security_logging`,`global_vrouter_config`.`fq_name`,`global_vrouter_config`.`key_value_pair` from `global_vrouter_config`"
const showGlobalVrouterConfigQuery = "select `global_vrouter_config`.`forwarding_mode`,`global_vrouter_config`.`vxlan_network_identifier_mode`,`global_vrouter_config`.`owner`,`global_vrouter_config`.`owner_access`,`global_vrouter_config`.`global_access`,`global_vrouter_config`.`share`,`global_vrouter_config`.`ip_protocol`,`global_vrouter_config`.`source_ip`,`global_vrouter_config`.`hashing_configured`,`global_vrouter_config`.`source_port`,`global_vrouter_config`.`destination_port`,`global_vrouter_config`.`destination_ip`,`global_vrouter_config`.`flow_aging_timeout`,`global_vrouter_config`.`linklocal_service_entry`,`global_vrouter_config`.`uuid`,`global_vrouter_config`.`flow_export_rate`,`global_vrouter_config`.`creator`,`global_vrouter_config`.`user_visible`,`global_vrouter_config`.`last_modified`,`global_vrouter_config`.`permissions_owner_access`,`global_vrouter_config`.`other_access`,`global_vrouter_config`.`group`,`global_vrouter_config`.`group_access`,`global_vrouter_config`.`permissions_owner`,`global_vrouter_config`.`enable`,`global_vrouter_config`.`description`,`global_vrouter_config`.`created`,`global_vrouter_config`.`display_name`,`global_vrouter_config`.`encapsulation`,`global_vrouter_config`.`enable_security_logging`,`global_vrouter_config`.`fq_name`,`global_vrouter_config`.`key_value_pair` from `global_vrouter_config` where uuid = ?"

func CreateGlobalVrouterConfig(tx *sql.Tx, model *models.GlobalVrouterConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalVrouterConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ForwardingMode),
		string(model.VxlanNetworkIdentifierMode),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		utils.MustJSON(model.FlowAgingTimeoutList.FlowAgingTimeout),
		utils.MustJSON(model.LinklocalServices.LinklocalServiceEntry),
		string(model.UUID),
		int(model.FlowExportRate),
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
		string(model.IDPerms.Created),
		string(model.DisplayName),
		utils.MustJSON(model.EncapsulationPriorities.Encapsulation),
		bool(model.EnableSecurityLogging),
		utils.MustJSON(model.FQName),
		utils.MustJSON(model.Annotations.KeyValuePair))

	return err
}

func scanGlobalVrouterConfig(rows *sql.Rows) (*models.GlobalVrouterConfig, error) {
	m := models.MakeGlobalVrouterConfig()

	var jsonPerms2Share string

	var jsonFlowAgingTimeoutListFlowAgingTimeout string

	var jsonLinklocalServicesLinklocalServiceEntry string

	var jsonEncapsulationPrioritiesEncapsulation string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.ForwardingMode,
		&m.VxlanNetworkIdentifierMode,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.EcmpHashingIncludeFields.IPProtocol,
		&m.EcmpHashingIncludeFields.SourceIP,
		&m.EcmpHashingIncludeFields.HashingConfigured,
		&m.EcmpHashingIncludeFields.SourcePort,
		&m.EcmpHashingIncludeFields.DestinationPort,
		&m.EcmpHashingIncludeFields.DestinationIP,
		&jsonFlowAgingTimeoutListFlowAgingTimeout,
		&jsonLinklocalServicesLinklocalServiceEntry,
		&m.UUID,
		&m.FlowExportRate,
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
		&m.IDPerms.Created,
		&m.DisplayName,
		&jsonEncapsulationPrioritiesEncapsulation,
		&m.EnableSecurityLogging,
		&jsonFQName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFlowAgingTimeoutListFlowAgingTimeout), &m.FlowAgingTimeoutList.FlowAgingTimeout)

	json.Unmarshal([]byte(jsonLinklocalServicesLinklocalServiceEntry), &m.LinklocalServices.LinklocalServiceEntry)

	json.Unmarshal([]byte(jsonEncapsulationPrioritiesEncapsulation), &m.EncapsulationPriorities.Encapsulation)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildGlobalVrouterConfigWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["forwarding_mode"]; ok {
		results = append(results, "forwarding_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["vxlan_network_identifier_mode"]; ok {
		results = append(results, "vxlan_network_identifier_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListGlobalVrouterConfig(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.GlobalVrouterConfig, error) {
	result := models.MakeGlobalVrouterConfigSlice()
	whereQuery, values := buildGlobalVrouterConfigWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listGlobalVrouterConfigQuery)
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
		m, _ := scanGlobalVrouterConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowGlobalVrouterConfig(tx *sql.Tx, uuid string) (*models.GlobalVrouterConfig, error) {
	rows, err := tx.Query(showGlobalVrouterConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanGlobalVrouterConfig(rows)
	}
	return nil, nil
}

func UpdateGlobalVrouterConfig(tx *sql.Tx, uuid string, model *models.GlobalVrouterConfig) error {
	return nil
}

func DeleteGlobalVrouterConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteGlobalVrouterConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
