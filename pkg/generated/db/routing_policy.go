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

const insertRoutingPolicyQuery = "insert into `routing_policy` (`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRoutingPolicyQuery = "update `routing_policy` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteRoutingPolicyQuery = "delete from `routing_policy` where uuid = ?"
const listRoutingPolicyQuery = "select `routing_policy`.`owner`,`routing_policy`.`owner_access`,`routing_policy`.`global_access`,`routing_policy`.`share`,`routing_policy`.`uuid`,`routing_policy`.`fq_name`,`routing_policy`.`enable`,`routing_policy`.`description`,`routing_policy`.`created`,`routing_policy`.`creator`,`routing_policy`.`user_visible`,`routing_policy`.`last_modified`,`routing_policy`.`permissions_owner`,`routing_policy`.`permissions_owner_access`,`routing_policy`.`other_access`,`routing_policy`.`group`,`routing_policy`.`group_access`,`routing_policy`.`display_name`,`routing_policy`.`key_value_pair` from `routing_policy`"
const showRoutingPolicyQuery = "select `routing_policy`.`owner`,`routing_policy`.`owner_access`,`routing_policy`.`global_access`,`routing_policy`.`share`,`routing_policy`.`uuid`,`routing_policy`.`fq_name`,`routing_policy`.`enable`,`routing_policy`.`description`,`routing_policy`.`created`,`routing_policy`.`creator`,`routing_policy`.`user_visible`,`routing_policy`.`last_modified`,`routing_policy`.`permissions_owner`,`routing_policy`.`permissions_owner_access`,`routing_policy`.`other_access`,`routing_policy`.`group`,`routing_policy`.`group_access`,`routing_policy`.`display_name`,`routing_policy`.`key_value_pair` from `routing_policy` where uuid = ?"

const insertRoutingPolicyServiceInstanceQuery = "insert into `ref_routing_policy_service_instance` (`from`, `to` ,`right_sequence`,`left_sequence`) values (?, ?,?,?);"

func CreateRoutingPolicy(tx *sql.Tx, model *models.RoutingPolicy) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRoutingPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Perms2.Owner),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair))

	stmtServiceInstanceRef, err := tx.Prepare(insertRoutingPolicyServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.RightSequence),
			string(ref.Attr.LeftSequence))
	}

	return err
}

func scanRoutingPolicy(rows *sql.Rows) (*models.RoutingPolicy, error) {
	m := models.MakeRoutingPolicy()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildRoutingPolicyWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

func ListRoutingPolicy(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.RoutingPolicy, error) {
	result := models.MakeRoutingPolicySlice()
	whereQuery, values := buildRoutingPolicyWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listRoutingPolicyQuery)
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
		m, _ := scanRoutingPolicy(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowRoutingPolicy(tx *sql.Tx, uuid string) (*models.RoutingPolicy, error) {
	rows, err := tx.Query(showRoutingPolicyQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRoutingPolicy(rows)
	}
	return nil, nil
}

func UpdateRoutingPolicy(tx *sql.Tx, uuid string, model *models.RoutingPolicy) error {
	return nil
}

func DeleteRoutingPolicy(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteRoutingPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
