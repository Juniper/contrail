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

const insertAddressGroupQuery = "insert into `address_group` (`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`address_group_prefix`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAddressGroupQuery = "update `address_group` set `creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`address_group_prefix` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteAddressGroupQuery = "delete from `address_group` where uuid = ?"
const listAddressGroupQuery = "select `address_group`.`creator`,`address_group`.`user_visible`,`address_group`.`last_modified`,`address_group`.`owner`,`address_group`.`owner_access`,`address_group`.`other_access`,`address_group`.`group`,`address_group`.`group_access`,`address_group`.`enable`,`address_group`.`description`,`address_group`.`created`,`address_group`.`display_name`,`address_group`.`key_value_pair`,`address_group`.`perms2_owner`,`address_group`.`perms2_owner_access`,`address_group`.`global_access`,`address_group`.`share`,`address_group`.`address_group_prefix`,`address_group`.`uuid`,`address_group`.`fq_name` from `address_group`"
const showAddressGroupQuery = "select `address_group`.`creator`,`address_group`.`user_visible`,`address_group`.`last_modified`,`address_group`.`owner`,`address_group`.`owner_access`,`address_group`.`other_access`,`address_group`.`group`,`address_group`.`group_access`,`address_group`.`enable`,`address_group`.`description`,`address_group`.`created`,`address_group`.`display_name`,`address_group`.`key_value_pair`,`address_group`.`perms2_owner`,`address_group`.`perms2_owner_access`,`address_group`.`global_access`,`address_group`.`share`,`address_group`.`address_group_prefix`,`address_group`.`uuid`,`address_group`.`fq_name` from `address_group` where uuid = ?"

func CreateAddressGroup(tx *sql.Tx, model *models.AddressGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAddressGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Creator),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.AddressGroupPrefix),
		string(model.UUID),
		utils.MustJSON(model.FQName))

	return err
}

func scanAddressGroup(rows *sql.Rows) (*models.AddressGroup, error) {
	m := models.MakeAddressGroup()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonAddressGroupPrefix string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.Creator,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonAddressGroupPrefix,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAddressGroupPrefix), &m.AddressGroupPrefix)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildAddressGroupWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
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

func ListAddressGroup(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.AddressGroup, error) {
	result := models.MakeAddressGroupSlice()
	whereQuery, values := buildAddressGroupWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAddressGroupQuery)
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
		m, _ := scanAddressGroup(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAddressGroup(tx *sql.Tx, uuid string) (*models.AddressGroup, error) {
	rows, err := tx.Query(showAddressGroupQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAddressGroup(rows)
	}
	return nil, nil
}

func UpdateAddressGroup(tx *sql.Tx, uuid string, model *models.AddressGroup) error {
	return nil
}

func DeleteAddressGroup(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAddressGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
