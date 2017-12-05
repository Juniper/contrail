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

const insertNetworkPolicyQuery = "insert into `network_policy` (`uuid`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`policy_rule`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkPolicyQuery = "update `network_policy` set `uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`policy_rule` = ?;"
const deleteNetworkPolicyQuery = "delete from `network_policy` where uuid = ?"
const listNetworkPolicyQuery = "select `network_policy`.`uuid`,`network_policy`.`fq_name`,`network_policy`.`user_visible`,`network_policy`.`last_modified`,`network_policy`.`other_access`,`network_policy`.`group`,`network_policy`.`group_access`,`network_policy`.`owner`,`network_policy`.`owner_access`,`network_policy`.`enable`,`network_policy`.`description`,`network_policy`.`created`,`network_policy`.`creator`,`network_policy`.`display_name`,`network_policy`.`key_value_pair`,`network_policy`.`global_access`,`network_policy`.`share`,`network_policy`.`perms2_owner`,`network_policy`.`perms2_owner_access`,`network_policy`.`policy_rule` from `network_policy`"
const showNetworkPolicyQuery = "select `network_policy`.`uuid`,`network_policy`.`fq_name`,`network_policy`.`user_visible`,`network_policy`.`last_modified`,`network_policy`.`other_access`,`network_policy`.`group`,`network_policy`.`group_access`,`network_policy`.`owner`,`network_policy`.`owner_access`,`network_policy`.`enable`,`network_policy`.`description`,`network_policy`.`created`,`network_policy`.`creator`,`network_policy`.`display_name`,`network_policy`.`key_value_pair`,`network_policy`.`global_access`,`network_policy`.`share`,`network_policy`.`perms2_owner`,`network_policy`.`perms2_owner_access`,`network_policy`.`policy_rule` from `network_policy` where uuid = ?"

func CreateNetworkPolicy(tx *sql.Tx, model *models.NetworkPolicy) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		utils.MustJSON(model.NetworkPolicyEntries.PolicyRule))

	return err
}

func scanNetworkPolicy(rows *sql.Rows) (*models.NetworkPolicy, error) {
	m := models.MakeNetworkPolicy()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonNetworkPolicyEntriesPolicyRule string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonNetworkPolicyEntriesPolicyRule); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonNetworkPolicyEntriesPolicyRule), &m.NetworkPolicyEntries.PolicyRule)

	return m, nil
}

func buildNetworkPolicyWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

func ListNetworkPolicy(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.NetworkPolicy, error) {
	result := models.MakeNetworkPolicySlice()
	whereQuery, values := buildNetworkPolicyWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listNetworkPolicyQuery)
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
		m, _ := scanNetworkPolicy(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNetworkPolicy(tx *sql.Tx, uuid string) (*models.NetworkPolicy, error) {
	rows, err := tx.Query(showNetworkPolicyQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNetworkPolicy(rows)
	}
	return nil, nil
}

func UpdateNetworkPolicy(tx *sql.Tx, uuid string, model *models.NetworkPolicy) error {
	return nil
}

func DeleteNetworkPolicy(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNetworkPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
