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

const insertNetworkDeviceConfigQuery = "insert into `network_device_config` (`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkDeviceConfigQuery = "update `network_device_config` set `created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteNetworkDeviceConfigQuery = "delete from `network_device_config` where uuid = ?"
const listNetworkDeviceConfigQuery = "select `network_device_config`.`created`,`network_device_config`.`creator`,`network_device_config`.`user_visible`,`network_device_config`.`last_modified`,`network_device_config`.`owner`,`network_device_config`.`owner_access`,`network_device_config`.`other_access`,`network_device_config`.`group`,`network_device_config`.`group_access`,`network_device_config`.`enable`,`network_device_config`.`description`,`network_device_config`.`display_name`,`network_device_config`.`key_value_pair`,`network_device_config`.`share`,`network_device_config`.`perms2_owner`,`network_device_config`.`perms2_owner_access`,`network_device_config`.`global_access`,`network_device_config`.`uuid`,`network_device_config`.`fq_name` from `network_device_config`"
const showNetworkDeviceConfigQuery = "select `network_device_config`.`created`,`network_device_config`.`creator`,`network_device_config`.`user_visible`,`network_device_config`.`last_modified`,`network_device_config`.`owner`,`network_device_config`.`owner_access`,`network_device_config`.`other_access`,`network_device_config`.`group`,`network_device_config`.`group_access`,`network_device_config`.`enable`,`network_device_config`.`description`,`network_device_config`.`display_name`,`network_device_config`.`key_value_pair`,`network_device_config`.`share`,`network_device_config`.`perms2_owner`,`network_device_config`.`perms2_owner_access`,`network_device_config`.`global_access`,`network_device_config`.`uuid`,`network_device_config`.`fq_name` from `network_device_config` where uuid = ?"

const insertNetworkDeviceConfigPhysicalRouterQuery = "insert into `ref_network_device_config_physical_router` (`from`, `to` ) values (?, ?);"

func CreateNetworkDeviceConfig(tx *sql.Tx, model *models.NetworkDeviceConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkDeviceConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Created),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		utils.MustJSON(model.FQName))

	stmtPhysicalRouterRef, err := tx.Prepare(insertNetworkDeviceConfigPhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {
		_, err = stmtPhysicalRouterRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanNetworkDeviceConfig(rows *sql.Rows) (*models.NetworkDeviceConfig, error) {
	m := models.MakeNetworkDeviceConfig()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.Created,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildNetworkDeviceConfigWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListNetworkDeviceConfig(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.NetworkDeviceConfig, error) {
	result := models.MakeNetworkDeviceConfigSlice()
	whereQuery, values := buildNetworkDeviceConfigWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listNetworkDeviceConfigQuery)
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
		m, _ := scanNetworkDeviceConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNetworkDeviceConfig(tx *sql.Tx, uuid string) (*models.NetworkDeviceConfig, error) {
	rows, err := tx.Query(showNetworkDeviceConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNetworkDeviceConfig(rows)
	}
	return nil, nil
}

func UpdateNetworkDeviceConfig(tx *sql.Tx, uuid string, model *models.NetworkDeviceConfig) error {
	return nil
}

func DeleteNetworkDeviceConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNetworkDeviceConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
