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

const insertConfigNodeQuery = "insert into `config_node` (`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`config_node_ip_address`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateConfigNodeQuery = "update `config_node` set `fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`config_node_ip_address` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?;"
const deleteConfigNodeQuery = "delete from `config_node` where uuid = ?"
const listConfigNodeQuery = "select `config_node`.`fq_name`,`config_node`.`description`,`config_node`.`created`,`config_node`.`creator`,`config_node`.`user_visible`,`config_node`.`last_modified`,`config_node`.`owner_access`,`config_node`.`other_access`,`config_node`.`group`,`config_node`.`group_access`,`config_node`.`owner`,`config_node`.`enable`,`config_node`.`config_node_ip_address`,`config_node`.`display_name`,`config_node`.`key_value_pair`,`config_node`.`perms2_owner`,`config_node`.`perms2_owner_access`,`config_node`.`global_access`,`config_node`.`share`,`config_node`.`uuid` from `config_node`"
const showConfigNodeQuery = "select `config_node`.`fq_name`,`config_node`.`description`,`config_node`.`created`,`config_node`.`creator`,`config_node`.`user_visible`,`config_node`.`last_modified`,`config_node`.`owner_access`,`config_node`.`other_access`,`config_node`.`group`,`config_node`.`group_access`,`config_node`.`owner`,`config_node`.`enable`,`config_node`.`config_node_ip_address`,`config_node`.`display_name`,`config_node`.`key_value_pair`,`config_node`.`perms2_owner`,`config_node`.`perms2_owner_access`,`config_node`.`global_access`,`config_node`.`share`,`config_node`.`uuid` from `config_node` where uuid = ?"

func CreateConfigNode(tx *sql.Tx, model *models.ConfigNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertConfigNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
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
		bool(model.IDPerms.Enable),
		string(model.ConfigNodeIPAddress),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID))

	return err
}

func scanConfigNode(rows *sql.Rows) (*models.ConfigNode, error) {
	m := models.MakeConfigNode()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.IDPerms.Description,
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
		&m.ConfigNodeIPAddress,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildConfigNodeWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["config_node_ip_address"]; ok {
		results = append(results, "config_node_ip_address = ?")
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

func ListConfigNode(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ConfigNode, error) {
	result := models.MakeConfigNodeSlice()
	whereQuery, values := buildConfigNodeWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listConfigNodeQuery)
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
		m, _ := scanConfigNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowConfigNode(tx *sql.Tx, uuid string) (*models.ConfigNode, error) {
	rows, err := tx.Query(showConfigNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanConfigNode(rows)
	}
	return nil, nil
}

func UpdateConfigNode(tx *sql.Tx, uuid string, model *models.ConfigNode) error {
	return nil
}

func DeleteConfigNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteConfigNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
