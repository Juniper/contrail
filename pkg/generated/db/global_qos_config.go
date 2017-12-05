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

const insertGlobalQosConfigQuery = "insert into `global_qos_config` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`control`,`analytics`,`dns`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalQosConfigQuery = "update `global_qos_config` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`control` = ?,`analytics` = ?,`dns` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`display_name` = ?;"
const deleteGlobalQosConfigQuery = "delete from `global_qos_config` where uuid = ?"
const listGlobalQosConfigQuery = "select `global_qos_config`.`key_value_pair`,`global_qos_config`.`owner`,`global_qos_config`.`owner_access`,`global_qos_config`.`global_access`,`global_qos_config`.`share`,`global_qos_config`.`uuid`,`global_qos_config`.`fq_name`,`global_qos_config`.`control`,`global_qos_config`.`analytics`,`global_qos_config`.`dns`,`global_qos_config`.`description`,`global_qos_config`.`created`,`global_qos_config`.`creator`,`global_qos_config`.`user_visible`,`global_qos_config`.`last_modified`,`global_qos_config`.`group_access`,`global_qos_config`.`permissions_owner`,`global_qos_config`.`permissions_owner_access`,`global_qos_config`.`other_access`,`global_qos_config`.`group`,`global_qos_config`.`enable`,`global_qos_config`.`display_name` from `global_qos_config`"
const showGlobalQosConfigQuery = "select `global_qos_config`.`key_value_pair`,`global_qos_config`.`owner`,`global_qos_config`.`owner_access`,`global_qos_config`.`global_access`,`global_qos_config`.`share`,`global_qos_config`.`uuid`,`global_qos_config`.`fq_name`,`global_qos_config`.`control`,`global_qos_config`.`analytics`,`global_qos_config`.`dns`,`global_qos_config`.`description`,`global_qos_config`.`created`,`global_qos_config`.`creator`,`global_qos_config`.`user_visible`,`global_qos_config`.`last_modified`,`global_qos_config`.`group_access`,`global_qos_config`.`permissions_owner`,`global_qos_config`.`permissions_owner_access`,`global_qos_config`.`other_access`,`global_qos_config`.`group`,`global_qos_config`.`enable`,`global_qos_config`.`display_name` from `global_qos_config` where uuid = ?"

func CreateGlobalQosConfig(tx *sql.Tx, model *models.GlobalQosConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		int(model.ControlTrafficDSCP.Control),
		int(model.ControlTrafficDSCP.Analytics),
		int(model.ControlTrafficDSCP.DNS),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.DisplayName))

	return err
}

func scanGlobalQosConfig(rows *sql.Rows) (*models.GlobalQosConfig, error) {
	m := models.MakeGlobalQosConfig()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.ControlTrafficDSCP.Control,
		&m.ControlTrafficDSCP.Analytics,
		&m.ControlTrafficDSCP.DNS,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildGlobalQosConfigWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListGlobalQosConfig(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.GlobalQosConfig, error) {
	result := models.MakeGlobalQosConfigSlice()
	whereQuery, values := buildGlobalQosConfigWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listGlobalQosConfigQuery)
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
		m, _ := scanGlobalQosConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowGlobalQosConfig(tx *sql.Tx, uuid string) (*models.GlobalQosConfig, error) {
	rows, err := tx.Query(showGlobalQosConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanGlobalQosConfig(rows)
	}
	return nil, nil
}

func UpdateGlobalQosConfig(tx *sql.Tx, uuid string, model *models.GlobalQosConfig) error {
	return nil
}

func DeleteGlobalQosConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteGlobalQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
