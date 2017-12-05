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

const insertNamespaceQuery = "insert into `namespace` (`fq_name`,`ip_prefix`,`ip_prefix_len`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNamespaceQuery = "update `namespace` set `fq_name` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?;"
const deleteNamespaceQuery = "delete from `namespace` where uuid = ?"
const listNamespaceQuery = "select `namespace`.`fq_name`,`namespace`.`ip_prefix`,`namespace`.`ip_prefix_len`,`namespace`.`last_modified`,`namespace`.`group_access`,`namespace`.`owner`,`namespace`.`owner_access`,`namespace`.`other_access`,`namespace`.`group`,`namespace`.`enable`,`namespace`.`description`,`namespace`.`created`,`namespace`.`creator`,`namespace`.`user_visible`,`namespace`.`display_name`,`namespace`.`key_value_pair`,`namespace`.`perms2_owner`,`namespace`.`perms2_owner_access`,`namespace`.`global_access`,`namespace`.`share`,`namespace`.`uuid` from `namespace`"
const showNamespaceQuery = "select `namespace`.`fq_name`,`namespace`.`ip_prefix`,`namespace`.`ip_prefix_len`,`namespace`.`last_modified`,`namespace`.`group_access`,`namespace`.`owner`,`namespace`.`owner_access`,`namespace`.`other_access`,`namespace`.`group`,`namespace`.`enable`,`namespace`.`description`,`namespace`.`created`,`namespace`.`creator`,`namespace`.`user_visible`,`namespace`.`display_name`,`namespace`.`key_value_pair`,`namespace`.`perms2_owner`,`namespace`.`perms2_owner_access`,`namespace`.`global_access`,`namespace`.`share`,`namespace`.`uuid` from `namespace` where uuid = ?"

func CreateNamespace(tx *sql.Tx, model *models.Namespace) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNamespaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		string(model.NamespaceCidr.IPPrefix),
		int(model.NamespaceCidr.IPPrefixLen),
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

func scanNamespace(rows *sql.Rows) (*models.Namespace, error) {
	m := models.MakeNamespace()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.NamespaceCidr.IPPrefix,
		&m.NamespaceCidr.IPPrefixLen,
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

func buildNamespaceWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
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

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListNamespace(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Namespace, error) {
	result := models.MakeNamespaceSlice()
	whereQuery, values := buildNamespaceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listNamespaceQuery)
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
		m, _ := scanNamespace(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNamespace(tx *sql.Tx, uuid string) (*models.Namespace, error) {
	rows, err := tx.Query(showNamespaceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNamespace(rows)
	}
	return nil, nil
}

func UpdateNamespace(tx *sql.Tx, uuid string, model *models.Namespace) error {
	return nil
}

func DeleteNamespace(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNamespaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
