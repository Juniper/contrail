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

const insertAccessControlListQuery = "insert into `access_control_list` (`display_name`,`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`access_control_list_hash`,`dynamic`,`acl_rule`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAccessControlListQuery = "update `access_control_list` set `display_name` = ?,`key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`access_control_list_hash` = ?,`dynamic` = ?,`acl_rule` = ?;"
const deleteAccessControlListQuery = "delete from `access_control_list` where uuid = ?"
const listAccessControlListQuery = "select `access_control_list`.`display_name`,`access_control_list`.`key_value_pair`,`access_control_list`.`owner_access`,`access_control_list`.`global_access`,`access_control_list`.`share`,`access_control_list`.`owner`,`access_control_list`.`uuid`,`access_control_list`.`fq_name`,`access_control_list`.`group`,`access_control_list`.`group_access`,`access_control_list`.`permissions_owner`,`access_control_list`.`permissions_owner_access`,`access_control_list`.`other_access`,`access_control_list`.`enable`,`access_control_list`.`description`,`access_control_list`.`created`,`access_control_list`.`creator`,`access_control_list`.`user_visible`,`access_control_list`.`last_modified`,`access_control_list`.`access_control_list_hash`,`access_control_list`.`dynamic`,`access_control_list`.`acl_rule` from `access_control_list`"
const showAccessControlListQuery = "select `access_control_list`.`display_name`,`access_control_list`.`key_value_pair`,`access_control_list`.`owner_access`,`access_control_list`.`global_access`,`access_control_list`.`share`,`access_control_list`.`owner`,`access_control_list`.`uuid`,`access_control_list`.`fq_name`,`access_control_list`.`group`,`access_control_list`.`group_access`,`access_control_list`.`permissions_owner`,`access_control_list`.`permissions_owner_access`,`access_control_list`.`other_access`,`access_control_list`.`enable`,`access_control_list`.`description`,`access_control_list`.`created`,`access_control_list`.`creator`,`access_control_list`.`user_visible`,`access_control_list`.`last_modified`,`access_control_list`.`access_control_list_hash`,`access_control_list`.`dynamic`,`access_control_list`.`acl_rule` from `access_control_list` where uuid = ?"

func CreateAccessControlList(tx *sql.Tx, model *models.AccessControlList) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAccessControlListQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		utils.MustJSON(model.AccessControlListHash),
		bool(model.AccessControlListEntries.Dynamic),
		utils.MustJSON(model.AccessControlListEntries.ACLRule))

	return err
}

func scanAccessControlList(rows *sql.Rows) (*models.AccessControlList, error) {
	m := models.MakeAccessControlList()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAccessControlListHash string

	var jsonAccessControlListEntriesACLRule string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&jsonAccessControlListHash,
		&m.AccessControlListEntries.Dynamic,
		&jsonAccessControlListEntriesACLRule); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAccessControlListHash), &m.AccessControlListHash)

	json.Unmarshal([]byte(jsonAccessControlListEntriesACLRule), &m.AccessControlListEntries.ACLRule)

	return m, nil
}

func buildAccessControlListWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

func ListAccessControlList(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.AccessControlList, error) {
	result := models.MakeAccessControlListSlice()
	whereQuery, values := buildAccessControlListWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAccessControlListQuery)
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
		m, _ := scanAccessControlList(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAccessControlList(tx *sql.Tx, uuid string) (*models.AccessControlList, error) {
	rows, err := tx.Query(showAccessControlListQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAccessControlList(rows)
	}
	return nil, nil
}

func UpdateAccessControlList(tx *sql.Tx, uuid string, model *models.AccessControlList) error {
	return nil
}

func DeleteAccessControlList(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAccessControlListQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
