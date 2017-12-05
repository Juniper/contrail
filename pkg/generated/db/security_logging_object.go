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

const insertSecurityLoggingObjectQuery = "insert into `security_logging_object` (`owner_access`,`global_access`,`share`,`owner`,`rule`,`security_logging_object_rate`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSecurityLoggingObjectQuery = "update `security_logging_object` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`rule` = ?,`security_logging_object_rate` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteSecurityLoggingObjectQuery = "delete from `security_logging_object` where uuid = ?"
const listSecurityLoggingObjectQuery = "select `security_logging_object`.`owner_access`,`security_logging_object`.`global_access`,`security_logging_object`.`share`,`security_logging_object`.`owner`,`security_logging_object`.`rule`,`security_logging_object`.`security_logging_object_rate`,`security_logging_object`.`uuid`,`security_logging_object`.`fq_name`,`security_logging_object`.`creator`,`security_logging_object`.`user_visible`,`security_logging_object`.`last_modified`,`security_logging_object`.`permissions_owner`,`security_logging_object`.`permissions_owner_access`,`security_logging_object`.`other_access`,`security_logging_object`.`group`,`security_logging_object`.`group_access`,`security_logging_object`.`enable`,`security_logging_object`.`description`,`security_logging_object`.`created`,`security_logging_object`.`display_name`,`security_logging_object`.`key_value_pair` from `security_logging_object`"
const showSecurityLoggingObjectQuery = "select `security_logging_object`.`owner_access`,`security_logging_object`.`global_access`,`security_logging_object`.`share`,`security_logging_object`.`owner`,`security_logging_object`.`rule`,`security_logging_object`.`security_logging_object_rate`,`security_logging_object`.`uuid`,`security_logging_object`.`fq_name`,`security_logging_object`.`creator`,`security_logging_object`.`user_visible`,`security_logging_object`.`last_modified`,`security_logging_object`.`permissions_owner`,`security_logging_object`.`permissions_owner_access`,`security_logging_object`.`other_access`,`security_logging_object`.`group`,`security_logging_object`.`group_access`,`security_logging_object`.`enable`,`security_logging_object`.`description`,`security_logging_object`.`created`,`security_logging_object`.`display_name`,`security_logging_object`.`key_value_pair` from `security_logging_object` where uuid = ?"

const insertSecurityLoggingObjectSecurityGroupQuery = "insert into `ref_security_logging_object_security_group` (`from`, `to` ,`rule`) values (?, ?,?);"

const insertSecurityLoggingObjectNetworkPolicyQuery = "insert into `ref_security_logging_object_network_policy` (`from`, `to` ) values (?, ?);"

func CreateSecurityLoggingObject(tx *sql.Tx, model *models.SecurityLoggingObject) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSecurityLoggingObjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		utils.MustJSON(model.SecurityLoggingObjectRules.Rule),
		int(model.SecurityLoggingObjectRate),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
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
		utils.MustJSON(model.Annotations.KeyValuePair))

	stmtSecurityGroupRef, err := tx.Prepare(insertSecurityLoggingObjectSecurityGroupQuery)
	if err != nil {
		return err
	}
	defer stmtSecurityGroupRef.Close()
	for _, ref := range model.SecurityGroupRefs {
		_, err = stmtSecurityGroupRef.Exec(model.UUID, ref.UUID, utils.MustJSON(ref.Attr.Rule))
	}

	stmtNetworkPolicyRef, err := tx.Prepare(insertSecurityLoggingObjectNetworkPolicyQuery)
	if err != nil {
		return err
	}
	defer stmtNetworkPolicyRef.Close()
	for _, ref := range model.NetworkPolicyRefs {
		_, err = stmtNetworkPolicyRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanSecurityLoggingObject(rows *sql.Rows) (*models.SecurityLoggingObject, error) {
	m := models.MakeSecurityLoggingObject()

	var jsonPerms2Share string

	var jsonSecurityLoggingObjectRulesRule string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&jsonSecurityLoggingObjectRulesRule,
		&m.SecurityLoggingObjectRate,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Creator,
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
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonSecurityLoggingObjectRulesRule), &m.SecurityLoggingObjectRules.Rule)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildSecurityLoggingObjectWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListSecurityLoggingObject(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.SecurityLoggingObject, error) {
	result := models.MakeSecurityLoggingObjectSlice()
	whereQuery, values := buildSecurityLoggingObjectWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listSecurityLoggingObjectQuery)
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
		m, _ := scanSecurityLoggingObject(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowSecurityLoggingObject(tx *sql.Tx, uuid string) (*models.SecurityLoggingObject, error) {
	rows, err := tx.Query(showSecurityLoggingObjectQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanSecurityLoggingObject(rows)
	}
	return nil, nil
}

func UpdateSecurityLoggingObject(tx *sql.Tx, uuid string, model *models.SecurityLoggingObject) error {
	return nil
}

func DeleteSecurityLoggingObject(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteSecurityLoggingObjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
