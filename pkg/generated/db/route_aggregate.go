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

const insertRouteAggregateQuery = "insert into `route_aggregate` (`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRouteAggregateQuery = "update `route_aggregate` set `display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?;"
const deleteRouteAggregateQuery = "delete from `route_aggregate` where uuid = ?"
const listRouteAggregateQuery = "select `route_aggregate`.`display_name`,`route_aggregate`.`key_value_pair`,`route_aggregate`.`owner`,`route_aggregate`.`owner_access`,`route_aggregate`.`global_access`,`route_aggregate`.`share`,`route_aggregate`.`uuid`,`route_aggregate`.`fq_name`,`route_aggregate`.`description`,`route_aggregate`.`created`,`route_aggregate`.`creator`,`route_aggregate`.`user_visible`,`route_aggregate`.`last_modified`,`route_aggregate`.`permissions_owner`,`route_aggregate`.`permissions_owner_access`,`route_aggregate`.`other_access`,`route_aggregate`.`group`,`route_aggregate`.`group_access`,`route_aggregate`.`enable` from `route_aggregate`"
const showRouteAggregateQuery = "select `route_aggregate`.`display_name`,`route_aggregate`.`key_value_pair`,`route_aggregate`.`owner`,`route_aggregate`.`owner_access`,`route_aggregate`.`global_access`,`route_aggregate`.`share`,`route_aggregate`.`uuid`,`route_aggregate`.`fq_name`,`route_aggregate`.`description`,`route_aggregate`.`created`,`route_aggregate`.`creator`,`route_aggregate`.`user_visible`,`route_aggregate`.`last_modified`,`route_aggregate`.`permissions_owner`,`route_aggregate`.`permissions_owner_access`,`route_aggregate`.`other_access`,`route_aggregate`.`group`,`route_aggregate`.`group_access`,`route_aggregate`.`enable` from `route_aggregate` where uuid = ?"

const insertRouteAggregateServiceInstanceQuery = "insert into `ref_route_aggregate_service_instance` (`from`, `to` ) values (?, ?);"

func CreateRouteAggregate(tx *sql.Tx, model *models.RouteAggregate) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteAggregateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		utils.MustJSON(model.FQName),
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
		bool(model.IDPerms.Enable))

	stmtServiceInstanceRef, err := tx.Prepare(insertRouteAggregateServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanRouteAggregate(rows *sql.Rows) (*models.RouteAggregate, error) {
	m := models.MakeRouteAggregate()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Enable); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildRouteAggregateWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListRouteAggregate(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.RouteAggregate, error) {
	result := models.MakeRouteAggregateSlice()
	whereQuery, values := buildRouteAggregateWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listRouteAggregateQuery)
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
		m, _ := scanRouteAggregate(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowRouteAggregate(tx *sql.Tx, uuid string) (*models.RouteAggregate, error) {
	rows, err := tx.Query(showRouteAggregateQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRouteAggregate(rows)
	}
	return nil, nil
}

func UpdateRouteAggregate(tx *sql.Tx, uuid string, model *models.RouteAggregate) error {
	return nil
}

func DeleteRouteAggregate(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteRouteAggregateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
