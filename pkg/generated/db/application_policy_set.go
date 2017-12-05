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

const insertApplicationPolicySetQuery = "insert into `application_policy_set` (`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`all_applications`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateApplicationPolicySetQuery = "update `application_policy_set` set `global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`all_applications` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteApplicationPolicySetQuery = "delete from `application_policy_set` where uuid = ?"
const listApplicationPolicySetQuery = "select `application_policy_set`.`global_access`,`application_policy_set`.`share`,`application_policy_set`.`owner`,`application_policy_set`.`owner_access`,`application_policy_set`.`uuid`,`application_policy_set`.`fq_name`,`application_policy_set`.`all_applications`,`application_policy_set`.`permissions_owner_access`,`application_policy_set`.`other_access`,`application_policy_set`.`group`,`application_policy_set`.`group_access`,`application_policy_set`.`permissions_owner`,`application_policy_set`.`enable`,`application_policy_set`.`description`,`application_policy_set`.`created`,`application_policy_set`.`creator`,`application_policy_set`.`user_visible`,`application_policy_set`.`last_modified`,`application_policy_set`.`display_name`,`application_policy_set`.`key_value_pair` from `application_policy_set`"
const showApplicationPolicySetQuery = "select `application_policy_set`.`global_access`,`application_policy_set`.`share`,`application_policy_set`.`owner`,`application_policy_set`.`owner_access`,`application_policy_set`.`uuid`,`application_policy_set`.`fq_name`,`application_policy_set`.`all_applications`,`application_policy_set`.`permissions_owner_access`,`application_policy_set`.`other_access`,`application_policy_set`.`group`,`application_policy_set`.`group_access`,`application_policy_set`.`permissions_owner`,`application_policy_set`.`enable`,`application_policy_set`.`description`,`application_policy_set`.`created`,`application_policy_set`.`creator`,`application_policy_set`.`user_visible`,`application_policy_set`.`last_modified`,`application_policy_set`.`display_name`,`application_policy_set`.`key_value_pair` from `application_policy_set` where uuid = ?"

const insertApplicationPolicySetFirewallPolicyQuery = "insert into `ref_application_policy_set_firewall_policy` (`from`, `to` ,`sequence`) values (?, ?,?);"

const insertApplicationPolicySetGlobalVrouterConfigQuery = "insert into `ref_application_policy_set_global_vrouter_config` (`from`, `to` ) values (?, ?);"

func CreateApplicationPolicySet(tx *sql.Tx, model *models.ApplicationPolicySet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertApplicationPolicySetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		bool(model.AllApplications),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair))

	stmtFirewallPolicyRef, err := tx.Prepare(insertApplicationPolicySetFirewallPolicyQuery)
	if err != nil {
		return err
	}
	defer stmtFirewallPolicyRef.Close()
	for _, ref := range model.FirewallPolicyRefs {
		_, err = stmtFirewallPolicyRef.Exec(model.UUID, ref.UUID, string(ref.Attr.Sequence))
	}

	stmtGlobalVrouterConfigRef, err := tx.Prepare(insertApplicationPolicySetGlobalVrouterConfigQuery)
	if err != nil {
		return err
	}
	defer stmtGlobalVrouterConfigRef.Close()
	for _, ref := range model.GlobalVrouterConfigRefs {
		_, err = stmtGlobalVrouterConfigRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanApplicationPolicySet(rows *sql.Rows) (*models.ApplicationPolicySet, error) {
	m := models.MakeApplicationPolicySet()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&jsonFQName,
		&m.AllApplications,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildApplicationPolicySetWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

func ListApplicationPolicySet(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ApplicationPolicySet, error) {
	result := models.MakeApplicationPolicySetSlice()
	whereQuery, values := buildApplicationPolicySetWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listApplicationPolicySetQuery)
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
		m, _ := scanApplicationPolicySet(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowApplicationPolicySet(tx *sql.Tx, uuid string) (*models.ApplicationPolicySet, error) {
	rows, err := tx.Query(showApplicationPolicySetQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanApplicationPolicySet(rows)
	}
	return nil, nil
}

func UpdateApplicationPolicySet(tx *sql.Tx, uuid string, model *models.ApplicationPolicySet) error {
	return nil
}

func DeleteApplicationPolicySet(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteApplicationPolicySetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
