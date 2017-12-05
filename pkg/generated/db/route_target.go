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

const insertRouteTargetQuery = "insert into `route_target` (`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRouteTargetQuery = "update `route_target` set `uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteRouteTargetQuery = "delete from `route_target` where uuid = ?"
const listRouteTargetQuery = "select `route_target`.`uuid`,`route_target`.`fq_name`,`route_target`.`user_visible`,`route_target`.`last_modified`,`route_target`.`group_access`,`route_target`.`owner`,`route_target`.`owner_access`,`route_target`.`other_access`,`route_target`.`group`,`route_target`.`enable`,`route_target`.`description`,`route_target`.`created`,`route_target`.`creator`,`route_target`.`display_name`,`route_target`.`key_value_pair`,`route_target`.`perms2_owner`,`route_target`.`perms2_owner_access`,`route_target`.`global_access`,`route_target`.`share` from `route_target`"
const showRouteTargetQuery = "select `route_target`.`uuid`,`route_target`.`fq_name`,`route_target`.`user_visible`,`route_target`.`last_modified`,`route_target`.`group_access`,`route_target`.`owner`,`route_target`.`owner_access`,`route_target`.`other_access`,`route_target`.`group`,`route_target`.`enable`,`route_target`.`description`,`route_target`.`created`,`route_target`.`creator`,`route_target`.`display_name`,`route_target`.`key_value_pair`,`route_target`.`perms2_owner`,`route_target`.`perms2_owner_access`,`route_target`.`global_access`,`route_target`.`share` from `route_target` where uuid = ?"

func CreateRouteTarget(tx *sql.Tx, model *models.RouteTarget) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteTargetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share))

	return err
}

func scanRouteTarget(rows *sql.Rows) (*models.RouteTarget, error) {
	m := models.MakeRouteTarget()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildRouteTargetWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListRouteTarget(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.RouteTarget, error) {
	result := models.MakeRouteTargetSlice()
	whereQuery, values := buildRouteTargetWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listRouteTargetQuery)
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
		m, _ := scanRouteTarget(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowRouteTarget(tx *sql.Tx, uuid string) (*models.RouteTarget, error) {
	rows, err := tx.Query(showRouteTargetQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRouteTarget(rows)
	}
	return nil, nil
}

func UpdateRouteTarget(tx *sql.Tx, uuid string, model *models.RouteTarget) error {
	return nil
}

func DeleteRouteTarget(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteRouteTargetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
