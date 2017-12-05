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

const insertSecurityGroupQuery = "insert into `security_group` (`configured_security_group_id`,`display_name`,`key_value_pair`,`policy_rule`,`security_group_id`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSecurityGroupQuery = "update `security_group` set `configured_security_group_id` = ?,`display_name` = ?,`key_value_pair` = ?,`policy_rule` = ?,`security_group_id` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteSecurityGroupQuery = "delete from `security_group` where uuid = ?"
const listSecurityGroupQuery = "select `security_group`.`configured_security_group_id`,`security_group`.`display_name`,`security_group`.`key_value_pair`,`security_group`.`policy_rule`,`security_group`.`security_group_id`,`security_group`.`uuid`,`security_group`.`fq_name`,`security_group`.`creator`,`security_group`.`user_visible`,`security_group`.`last_modified`,`security_group`.`owner_access`,`security_group`.`other_access`,`security_group`.`group`,`security_group`.`group_access`,`security_group`.`owner`,`security_group`.`enable`,`security_group`.`description`,`security_group`.`created`,`security_group`.`global_access`,`security_group`.`share`,`security_group`.`perms2_owner`,`security_group`.`perms2_owner_access` from `security_group`"
const showSecurityGroupQuery = "select `security_group`.`configured_security_group_id`,`security_group`.`display_name`,`security_group`.`key_value_pair`,`security_group`.`policy_rule`,`security_group`.`security_group_id`,`security_group`.`uuid`,`security_group`.`fq_name`,`security_group`.`creator`,`security_group`.`user_visible`,`security_group`.`last_modified`,`security_group`.`owner_access`,`security_group`.`other_access`,`security_group`.`group`,`security_group`.`group_access`,`security_group`.`owner`,`security_group`.`enable`,`security_group`.`description`,`security_group`.`created`,`security_group`.`global_access`,`security_group`.`share`,`security_group`.`perms2_owner`,`security_group`.`perms2_owner_access` from `security_group` where uuid = ?"

func CreateSecurityGroup(tx *sql.Tx, model *models.SecurityGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSecurityGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.ConfiguredSecurityGroupID),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.SecurityGroupEntries.PolicyRule),
		int(model.SecurityGroupID),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess))

	return err
}

func scanSecurityGroup(rows *sql.Rows) (*models.SecurityGroup, error) {
	m := models.MakeSecurityGroup()

	var jsonAnnotationsKeyValuePair string

	var jsonSecurityGroupEntriesPolicyRule string

	var jsonFQName string

	var jsonPerms2Share string

	if err := rows.Scan(&m.ConfiguredSecurityGroupID,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonSecurityGroupEntriesPolicyRule,
		&m.SecurityGroupID,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonSecurityGroupEntriesPolicyRule), &m.SecurityGroupEntries.PolicyRule)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildSecurityGroupWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListSecurityGroup(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.SecurityGroup, error) {
	result := models.MakeSecurityGroupSlice()
	whereQuery, values := buildSecurityGroupWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listSecurityGroupQuery)
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
		m, _ := scanSecurityGroup(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowSecurityGroup(tx *sql.Tx, uuid string) (*models.SecurityGroup, error) {
	rows, err := tx.Query(showSecurityGroupQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanSecurityGroup(rows)
	}
	return nil, nil
}

func UpdateSecurityGroup(tx *sql.Tx, uuid string, model *models.SecurityGroup) error {
	return nil
}

func DeleteSecurityGroup(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteSecurityGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
