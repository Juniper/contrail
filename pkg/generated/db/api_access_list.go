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

const insertAPIAccessListQuery = "insert into `api_access_list` (`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`rbac_rule`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAPIAccessListQuery = "update `api_access_list` set `description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`rbac_rule` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteAPIAccessListQuery = "delete from `api_access_list` where uuid = ?"
const listAPIAccessListQuery = "select `api_access_list`.`description`,`api_access_list`.`created`,`api_access_list`.`creator`,`api_access_list`.`user_visible`,`api_access_list`.`last_modified`,`api_access_list`.`owner`,`api_access_list`.`owner_access`,`api_access_list`.`other_access`,`api_access_list`.`group`,`api_access_list`.`group_access`,`api_access_list`.`enable`,`api_access_list`.`display_name`,`api_access_list`.`key_value_pair`,`api_access_list`.`perms2_owner`,`api_access_list`.`perms2_owner_access`,`api_access_list`.`global_access`,`api_access_list`.`share`,`api_access_list`.`rbac_rule`,`api_access_list`.`uuid`,`api_access_list`.`fq_name` from `api_access_list`"
const showAPIAccessListQuery = "select `api_access_list`.`description`,`api_access_list`.`created`,`api_access_list`.`creator`,`api_access_list`.`user_visible`,`api_access_list`.`last_modified`,`api_access_list`.`owner`,`api_access_list`.`owner_access`,`api_access_list`.`other_access`,`api_access_list`.`group`,`api_access_list`.`group_access`,`api_access_list`.`enable`,`api_access_list`.`display_name`,`api_access_list`.`key_value_pair`,`api_access_list`.`perms2_owner`,`api_access_list`.`perms2_owner_access`,`api_access_list`.`global_access`,`api_access_list`.`share`,`api_access_list`.`rbac_rule`,`api_access_list`.`uuid`,`api_access_list`.`fq_name` from `api_access_list` where uuid = ?"

func CreateAPIAccessList(tx *sql.Tx, model *models.APIAccessList) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAPIAccessListQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.APIAccessListEntries.RbacRule),
		string(model.UUID),
		utils.MustJSON(model.FQName))

	return err
}

func scanAPIAccessList(rows *sql.Rows) (*models.APIAccessList, error) {
	m := models.MakeAPIAccessList()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonAPIAccessListEntriesRbacRule string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonAPIAccessListEntriesRbacRule,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAPIAccessListEntriesRbacRule), &m.APIAccessListEntries.RbacRule)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildAPIAccessListWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

func ListAPIAccessList(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.APIAccessList, error) {
	result := models.MakeAPIAccessListSlice()
	whereQuery, values := buildAPIAccessListWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAPIAccessListQuery)
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
		m, _ := scanAPIAccessList(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAPIAccessList(tx *sql.Tx, uuid string) (*models.APIAccessList, error) {
	rows, err := tx.Query(showAPIAccessListQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAPIAccessList(rows)
	}
	return nil, nil
}

func UpdateAPIAccessList(tx *sql.Tx, uuid string, model *models.APIAccessList) error {
	return nil
}

func DeleteAPIAccessList(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAPIAccessListQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
