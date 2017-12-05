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

const insertFirewallPolicyQuery = "insert into `firewall_policy` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFirewallPolicyQuery = "update `firewall_policy` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`display_name` = ?;"
const deleteFirewallPolicyQuery = "delete from `firewall_policy` where uuid = ?"
const listFirewallPolicyQuery = "select `firewall_policy`.`key_value_pair`,`firewall_policy`.`owner`,`firewall_policy`.`owner_access`,`firewall_policy`.`global_access`,`firewall_policy`.`share`,`firewall_policy`.`uuid`,`firewall_policy`.`fq_name`,`firewall_policy`.`enable`,`firewall_policy`.`description`,`firewall_policy`.`created`,`firewall_policy`.`creator`,`firewall_policy`.`user_visible`,`firewall_policy`.`last_modified`,`firewall_policy`.`other_access`,`firewall_policy`.`group`,`firewall_policy`.`group_access`,`firewall_policy`.`permissions_owner`,`firewall_policy`.`permissions_owner_access`,`firewall_policy`.`display_name` from `firewall_policy`"
const showFirewallPolicyQuery = "select `firewall_policy`.`key_value_pair`,`firewall_policy`.`owner`,`firewall_policy`.`owner_access`,`firewall_policy`.`global_access`,`firewall_policy`.`share`,`firewall_policy`.`uuid`,`firewall_policy`.`fq_name`,`firewall_policy`.`enable`,`firewall_policy`.`description`,`firewall_policy`.`created`,`firewall_policy`.`creator`,`firewall_policy`.`user_visible`,`firewall_policy`.`last_modified`,`firewall_policy`.`other_access`,`firewall_policy`.`group`,`firewall_policy`.`group_access`,`firewall_policy`.`permissions_owner`,`firewall_policy`.`permissions_owner_access`,`firewall_policy`.`display_name` from `firewall_policy` where uuid = ?"

const insertFirewallPolicySecurityLoggingObjectQuery = "insert into `ref_firewall_policy_security_logging_object` (`from`, `to` ) values (?, ?);"

const insertFirewallPolicyFirewallRuleQuery = "insert into `ref_firewall_policy_firewall_rule` (`from`, `to` ) values (?, ?);"

func CreateFirewallPolicy(tx *sql.Tx, model *models.FirewallPolicy) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFirewallPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
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
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.DisplayName))

	stmtFirewallRuleRef, err := tx.Prepare(insertFirewallPolicyFirewallRuleQuery)
	if err != nil {
		return err
	}
	defer stmtFirewallRuleRef.Close()
	for _, ref := range model.FirewallRuleRefs {
		_, err = stmtFirewallRuleRef.Exec(model.UUID, ref.UUID)
	}

	stmtSecurityLoggingObjectRef, err := tx.Prepare(insertFirewallPolicySecurityLoggingObjectQuery)
	if err != nil {
		return err
	}
	defer stmtSecurityLoggingObjectRef.Close()
	for _, ref := range model.SecurityLoggingObjectRefs {
		_, err = stmtSecurityLoggingObjectRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanFirewallPolicy(rows *sql.Rows) (*models.FirewallPolicy, error) {
	m := models.MakeFirewallPolicy()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
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
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildFirewallPolicyWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListFirewallPolicy(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.FirewallPolicy, error) {
	result := models.MakeFirewallPolicySlice()
	whereQuery, values := buildFirewallPolicyWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listFirewallPolicyQuery)
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
		m, _ := scanFirewallPolicy(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowFirewallPolicy(tx *sql.Tx, uuid string) (*models.FirewallPolicy, error) {
	rows, err := tx.Query(showFirewallPolicyQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanFirewallPolicy(rows)
	}
	return nil, nil
}

func UpdateFirewallPolicy(tx *sql.Tx, uuid string, model *models.FirewallPolicy) error {
	return nil
}

func DeleteFirewallPolicy(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFirewallPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
