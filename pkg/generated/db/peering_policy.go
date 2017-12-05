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

const insertPeeringPolicyQuery = "insert into `peering_policy` (`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`peering_service`,`uuid`,`fq_name`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePeeringPolicyQuery = "update `peering_policy` set `display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`peering_service` = ?,`uuid` = ?,`fq_name` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?;"
const deletePeeringPolicyQuery = "delete from `peering_policy` where uuid = ?"
const listPeeringPolicyQuery = "select `peering_policy`.`display_name`,`peering_policy`.`key_value_pair`,`peering_policy`.`owner`,`peering_policy`.`owner_access`,`peering_policy`.`global_access`,`peering_policy`.`share`,`peering_policy`.`peering_service`,`peering_policy`.`uuid`,`peering_policy`.`fq_name`,`peering_policy`.`other_access`,`peering_policy`.`group`,`peering_policy`.`group_access`,`peering_policy`.`permissions_owner`,`peering_policy`.`permissions_owner_access`,`peering_policy`.`enable`,`peering_policy`.`description`,`peering_policy`.`created`,`peering_policy`.`creator`,`peering_policy`.`user_visible`,`peering_policy`.`last_modified` from `peering_policy`"
const showPeeringPolicyQuery = "select `peering_policy`.`display_name`,`peering_policy`.`key_value_pair`,`peering_policy`.`owner`,`peering_policy`.`owner_access`,`peering_policy`.`global_access`,`peering_policy`.`share`,`peering_policy`.`peering_service`,`peering_policy`.`uuid`,`peering_policy`.`fq_name`,`peering_policy`.`other_access`,`peering_policy`.`group`,`peering_policy`.`group_access`,`peering_policy`.`permissions_owner`,`peering_policy`.`permissions_owner_access`,`peering_policy`.`enable`,`peering_policy`.`description`,`peering_policy`.`created`,`peering_policy`.`creator`,`peering_policy`.`user_visible`,`peering_policy`.`last_modified` from `peering_policy` where uuid = ?"

func CreatePeeringPolicy(tx *sql.Tx, model *models.PeeringPolicy) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPeeringPolicyQuery)
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
		string(model.PeeringService),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified))

	return err
}

func scanPeeringPolicy(rows *sql.Rows) (*models.PeeringPolicy, error) {
	m := models.MakePeeringPolicy()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.PeeringService,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildPeeringPolicyWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["peering_service"]; ok {
		results = append(results, "peering_service = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListPeeringPolicy(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.PeeringPolicy, error) {
	result := models.MakePeeringPolicySlice()
	whereQuery, values := buildPeeringPolicyWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listPeeringPolicyQuery)
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
		m, _ := scanPeeringPolicy(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowPeeringPolicy(tx *sql.Tx, uuid string) (*models.PeeringPolicy, error) {
	rows, err := tx.Query(showPeeringPolicyQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanPeeringPolicy(rows)
	}
	return nil, nil
}

func UpdatePeeringPolicy(tx *sql.Tx, uuid string, model *models.PeeringPolicy) error {
	return nil
}

func DeletePeeringPolicy(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deletePeeringPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
