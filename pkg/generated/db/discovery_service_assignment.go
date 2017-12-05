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

const insertDiscoveryServiceAssignmentQuery = "insert into `discovery_service_assignment` (`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDiscoveryServiceAssignmentQuery = "update `discovery_service_assignment` set `key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?;"
const deleteDiscoveryServiceAssignmentQuery = "delete from `discovery_service_assignment` where uuid = ?"
const listDiscoveryServiceAssignmentQuery = "select `discovery_service_assignment`.`key_value_pair`,`discovery_service_assignment`.`owner_access`,`discovery_service_assignment`.`global_access`,`discovery_service_assignment`.`share`,`discovery_service_assignment`.`owner`,`discovery_service_assignment`.`uuid`,`discovery_service_assignment`.`fq_name`,`discovery_service_assignment`.`other_access`,`discovery_service_assignment`.`group`,`discovery_service_assignment`.`group_access`,`discovery_service_assignment`.`permissions_owner`,`discovery_service_assignment`.`permissions_owner_access`,`discovery_service_assignment`.`enable`,`discovery_service_assignment`.`description`,`discovery_service_assignment`.`created`,`discovery_service_assignment`.`creator`,`discovery_service_assignment`.`user_visible`,`discovery_service_assignment`.`last_modified`,`discovery_service_assignment`.`display_name` from `discovery_service_assignment`"
const showDiscoveryServiceAssignmentQuery = "select `discovery_service_assignment`.`key_value_pair`,`discovery_service_assignment`.`owner_access`,`discovery_service_assignment`.`global_access`,`discovery_service_assignment`.`share`,`discovery_service_assignment`.`owner`,`discovery_service_assignment`.`uuid`,`discovery_service_assignment`.`fq_name`,`discovery_service_assignment`.`other_access`,`discovery_service_assignment`.`group`,`discovery_service_assignment`.`group_access`,`discovery_service_assignment`.`permissions_owner`,`discovery_service_assignment`.`permissions_owner_access`,`discovery_service_assignment`.`enable`,`discovery_service_assignment`.`description`,`discovery_service_assignment`.`created`,`discovery_service_assignment`.`creator`,`discovery_service_assignment`.`user_visible`,`discovery_service_assignment`.`last_modified`,`discovery_service_assignment`.`display_name` from `discovery_service_assignment` where uuid = ?"

func CreateDiscoveryServiceAssignment(tx *sql.Tx, model *models.DiscoveryServiceAssignment) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDiscoveryServiceAssignmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
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
		string(model.IDPerms.LastModified),
		string(model.DisplayName))

	return err
}

func scanDiscoveryServiceAssignment(rows *sql.Rows) (*models.DiscoveryServiceAssignment, error) {
	m := models.MakeDiscoveryServiceAssignment()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
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
		&m.IDPerms.LastModified,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildDiscoveryServiceAssignmentWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListDiscoveryServiceAssignment(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.DiscoveryServiceAssignment, error) {
	result := models.MakeDiscoveryServiceAssignmentSlice()
	whereQuery, values := buildDiscoveryServiceAssignmentWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listDiscoveryServiceAssignmentQuery)
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
		m, _ := scanDiscoveryServiceAssignment(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDiscoveryServiceAssignment(tx *sql.Tx, uuid string) (*models.DiscoveryServiceAssignment, error) {
	rows, err := tx.Query(showDiscoveryServiceAssignmentQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDiscoveryServiceAssignment(rows)
	}
	return nil, nil
}

func UpdateDiscoveryServiceAssignment(tx *sql.Tx, uuid string, model *models.DiscoveryServiceAssignment) error {
	return nil
}

func DeleteDiscoveryServiceAssignment(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDiscoveryServiceAssignmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
