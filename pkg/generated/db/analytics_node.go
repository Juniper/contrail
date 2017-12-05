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

const insertAnalyticsNodeQuery = "insert into `analytics_node` (`key_value_pair`,`analytics_node_ip_address`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAnalyticsNodeQuery = "update `analytics_node` set `key_value_pair` = ?,`analytics_node_ip_address` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?;"
const deleteAnalyticsNodeQuery = "delete from `analytics_node` where uuid = ?"
const listAnalyticsNodeQuery = "select `analytics_node`.`key_value_pair`,`analytics_node`.`analytics_node_ip_address`,`analytics_node`.`owner`,`analytics_node`.`owner_access`,`analytics_node`.`global_access`,`analytics_node`.`share`,`analytics_node`.`uuid`,`analytics_node`.`fq_name`,`analytics_node`.`enable`,`analytics_node`.`description`,`analytics_node`.`created`,`analytics_node`.`creator`,`analytics_node`.`user_visible`,`analytics_node`.`last_modified`,`analytics_node`.`permissions_owner`,`analytics_node`.`permissions_owner_access`,`analytics_node`.`other_access`,`analytics_node`.`group`,`analytics_node`.`group_access`,`analytics_node`.`display_name` from `analytics_node`"
const showAnalyticsNodeQuery = "select `analytics_node`.`key_value_pair`,`analytics_node`.`analytics_node_ip_address`,`analytics_node`.`owner`,`analytics_node`.`owner_access`,`analytics_node`.`global_access`,`analytics_node`.`share`,`analytics_node`.`uuid`,`analytics_node`.`fq_name`,`analytics_node`.`enable`,`analytics_node`.`description`,`analytics_node`.`created`,`analytics_node`.`creator`,`analytics_node`.`user_visible`,`analytics_node`.`last_modified`,`analytics_node`.`permissions_owner`,`analytics_node`.`permissions_owner_access`,`analytics_node`.`other_access`,`analytics_node`.`group`,`analytics_node`.`group_access`,`analytics_node`.`display_name` from `analytics_node` where uuid = ?"

func CreateAnalyticsNode(tx *sql.Tx, model *models.AnalyticsNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.AnalyticsNodeIPAddress),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.DisplayName))

	return err
}

func scanAnalyticsNode(rows *sql.Rows) (*models.AnalyticsNode, error) {
	m := models.MakeAnalyticsNode()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.AnalyticsNodeIPAddress,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildAnalyticsNodeWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["analytics_node_ip_address"]; ok {
		results = append(results, "analytics_node_ip_address = ?")
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

func ListAnalyticsNode(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.AnalyticsNode, error) {
	result := models.MakeAnalyticsNodeSlice()
	whereQuery, values := buildAnalyticsNodeWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAnalyticsNodeQuery)
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
		m, _ := scanAnalyticsNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAnalyticsNode(tx *sql.Tx, uuid string) (*models.AnalyticsNode, error) {
	rows, err := tx.Query(showAnalyticsNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAnalyticsNode(rows)
	}
	return nil, nil
}

func UpdateAnalyticsNode(tx *sql.Tx, uuid string, model *models.AnalyticsNode) error {
	return nil
}

func DeleteAnalyticsNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
