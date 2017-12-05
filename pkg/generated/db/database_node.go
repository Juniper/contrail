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

const insertDatabaseNodeQuery = "insert into `database_node` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`database_node_ip_address`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDatabaseNodeQuery = "update `database_node` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`database_node_ip_address` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`display_name` = ?;"
const deleteDatabaseNodeQuery = "delete from `database_node` where uuid = ?"
const listDatabaseNodeQuery = "select `database_node`.`key_value_pair`,`database_node`.`owner`,`database_node`.`owner_access`,`database_node`.`global_access`,`database_node`.`share`,`database_node`.`database_node_ip_address`,`database_node`.`uuid`,`database_node`.`fq_name`,`database_node`.`description`,`database_node`.`created`,`database_node`.`creator`,`database_node`.`user_visible`,`database_node`.`last_modified`,`database_node`.`other_access`,`database_node`.`group`,`database_node`.`group_access`,`database_node`.`permissions_owner`,`database_node`.`permissions_owner_access`,`database_node`.`enable`,`database_node`.`display_name` from `database_node`"
const showDatabaseNodeQuery = "select `database_node`.`key_value_pair`,`database_node`.`owner`,`database_node`.`owner_access`,`database_node`.`global_access`,`database_node`.`share`,`database_node`.`database_node_ip_address`,`database_node`.`uuid`,`database_node`.`fq_name`,`database_node`.`description`,`database_node`.`created`,`database_node`.`creator`,`database_node`.`user_visible`,`database_node`.`last_modified`,`database_node`.`other_access`,`database_node`.`group`,`database_node`.`group_access`,`database_node`.`permissions_owner`,`database_node`.`permissions_owner_access`,`database_node`.`enable`,`database_node`.`display_name` from `database_node` where uuid = ?"

func CreateDatabaseNode(tx *sql.Tx, model *models.DatabaseNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDatabaseNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.DatabaseNodeIPAddress),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName))

	return err
}

func scanDatabaseNode(rows *sql.Rows) (*models.DatabaseNode, error) {
	m := models.MakeDatabaseNode()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.DatabaseNodeIPAddress,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildDatabaseNodeWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["database_node_ip_address"]; ok {
		results = append(results, "database_node_ip_address = ?")
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

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListDatabaseNode(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.DatabaseNode, error) {
	result := models.MakeDatabaseNodeSlice()
	whereQuery, values := buildDatabaseNodeWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listDatabaseNodeQuery)
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
		m, _ := scanDatabaseNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDatabaseNode(tx *sql.Tx, uuid string) (*models.DatabaseNode, error) {
	rows, err := tx.Query(showDatabaseNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDatabaseNode(rows)
	}
	return nil, nil
}

func UpdateDatabaseNode(tx *sql.Tx, uuid string, model *models.DatabaseNode) error {
	return nil
}

func DeleteDatabaseNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDatabaseNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
