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

const insertRoutingInstanceQuery = "insert into `routing_instance` (`fq_name`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRoutingInstanceQuery = "update `routing_instance` set `fq_name` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?;"
const deleteRoutingInstanceQuery = "delete from `routing_instance` where uuid = ?"
const listRoutingInstanceQuery = "select `routing_instance`.`fq_name`,`routing_instance`.`last_modified`,`routing_instance`.`owner`,`routing_instance`.`owner_access`,`routing_instance`.`other_access`,`routing_instance`.`group`,`routing_instance`.`group_access`,`routing_instance`.`enable`,`routing_instance`.`description`,`routing_instance`.`created`,`routing_instance`.`creator`,`routing_instance`.`user_visible`,`routing_instance`.`display_name`,`routing_instance`.`key_value_pair`,`routing_instance`.`perms2_owner`,`routing_instance`.`perms2_owner_access`,`routing_instance`.`global_access`,`routing_instance`.`share`,`routing_instance`.`uuid` from `routing_instance`"
const showRoutingInstanceQuery = "select `routing_instance`.`fq_name`,`routing_instance`.`last_modified`,`routing_instance`.`owner`,`routing_instance`.`owner_access`,`routing_instance`.`other_access`,`routing_instance`.`group`,`routing_instance`.`group_access`,`routing_instance`.`enable`,`routing_instance`.`description`,`routing_instance`.`created`,`routing_instance`.`creator`,`routing_instance`.`user_visible`,`routing_instance`.`display_name`,`routing_instance`.`key_value_pair`,`routing_instance`.`perms2_owner`,`routing_instance`.`perms2_owner_access`,`routing_instance`.`global_access`,`routing_instance`.`share`,`routing_instance`.`uuid` from `routing_instance` where uuid = ?"

func CreateRoutingInstance(tx *sql.Tx, model *models.RoutingInstance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRoutingInstanceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID))

	return err
}

func scanRoutingInstance(rows *sql.Rows) (*models.RoutingInstance, error) {
	m := models.MakeRoutingInstance()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
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

func buildRoutingInstanceWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
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

func ListRoutingInstance(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.RoutingInstance, error) {
	result := models.MakeRoutingInstanceSlice()
	whereQuery, values := buildRoutingInstanceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listRoutingInstanceQuery)
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
		m, _ := scanRoutingInstance(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowRoutingInstance(tx *sql.Tx, uuid string) (*models.RoutingInstance, error) {
	rows, err := tx.Query(showRoutingInstanceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRoutingInstance(rows)
	}
	return nil, nil
}

func UpdateRoutingInstance(tx *sql.Tx, uuid string, model *models.RoutingInstance) error {
	return nil
}

func DeleteRoutingInstance(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteRoutingInstanceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
