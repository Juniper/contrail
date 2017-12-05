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

const insertRouteTableQuery = "insert into `route_table` (`route`,`fq_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRouteTableQuery = "update `route_table` set `route` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?;"
const deleteRouteTableQuery = "delete from `route_table` where uuid = ?"
const listRouteTableQuery = "select `route_table`.`route`,`route_table`.`fq_name`,`route_table`.`user_visible`,`route_table`.`last_modified`,`route_table`.`owner`,`route_table`.`owner_access`,`route_table`.`other_access`,`route_table`.`group`,`route_table`.`group_access`,`route_table`.`enable`,`route_table`.`description`,`route_table`.`created`,`route_table`.`creator`,`route_table`.`display_name`,`route_table`.`key_value_pair`,`route_table`.`perms2_owner_access`,`route_table`.`global_access`,`route_table`.`share`,`route_table`.`perms2_owner`,`route_table`.`uuid` from `route_table`"
const showRouteTableQuery = "select `route_table`.`route`,`route_table`.`fq_name`,`route_table`.`user_visible`,`route_table`.`last_modified`,`route_table`.`owner`,`route_table`.`owner_access`,`route_table`.`other_access`,`route_table`.`group`,`route_table`.`group_access`,`route_table`.`enable`,`route_table`.`description`,`route_table`.`created`,`route_table`.`creator`,`route_table`.`display_name`,`route_table`.`key_value_pair`,`route_table`.`perms2_owner_access`,`route_table`.`global_access`,`route_table`.`share`,`route_table`.`perms2_owner`,`route_table`.`uuid` from `route_table` where uuid = ?"

func CreateRouteTable(tx *sql.Tx, model *models.RouteTable) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Routes.Route),
		utils.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID))

	return err
}

func scanRouteTable(rows *sql.Rows) (*models.RouteTable, error) {
	m := models.MakeRouteTable()

	var jsonRoutesRoute string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonRoutesRoute,
		&jsonFQName,
		&m.IDPerms.UserVisible,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonRoutesRoute), &m.Routes.Route)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildRouteTableWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

func ListRouteTable(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.RouteTable, error) {
	result := models.MakeRouteTableSlice()
	whereQuery, values := buildRouteTableWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listRouteTableQuery)
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
		m, _ := scanRouteTable(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowRouteTable(tx *sql.Tx, uuid string) (*models.RouteTable, error) {
	rows, err := tx.Query(showRouteTableQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRouteTable(rows)
	}
	return nil, nil
}

func UpdateRouteTable(tx *sql.Tx, uuid string, model *models.RouteTable) error {
	return nil
}

func DeleteRouteTable(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
