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

const insertDsaRuleQuery = "insert into `dsa_rule` (`uuid`,`fq_name`,`subscriber`,`ep_version`,`ep_id`,`ep_type`,`ip_prefix`,`ip_prefix_len`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDsaRuleQuery = "update `dsa_rule` set `uuid` = ?,`fq_name` = ?,`subscriber` = ?,`ep_version` = ?,`ep_id` = ?,`ep_type` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteDsaRuleQuery = "delete from `dsa_rule` where uuid = ?"
const listDsaRuleQuery = "select `dsa_rule`.`uuid`,`dsa_rule`.`fq_name`,`dsa_rule`.`subscriber`,`dsa_rule`.`ep_version`,`dsa_rule`.`ep_id`,`dsa_rule`.`ep_type`,`dsa_rule`.`ip_prefix`,`dsa_rule`.`ip_prefix_len`,`dsa_rule`.`creator`,`dsa_rule`.`user_visible`,`dsa_rule`.`last_modified`,`dsa_rule`.`group`,`dsa_rule`.`group_access`,`dsa_rule`.`owner`,`dsa_rule`.`owner_access`,`dsa_rule`.`other_access`,`dsa_rule`.`enable`,`dsa_rule`.`description`,`dsa_rule`.`created`,`dsa_rule`.`display_name`,`dsa_rule`.`key_value_pair`,`dsa_rule`.`global_access`,`dsa_rule`.`share`,`dsa_rule`.`perms2_owner`,`dsa_rule`.`perms2_owner_access` from `dsa_rule`"
const showDsaRuleQuery = "select `dsa_rule`.`uuid`,`dsa_rule`.`fq_name`,`dsa_rule`.`subscriber`,`dsa_rule`.`ep_version`,`dsa_rule`.`ep_id`,`dsa_rule`.`ep_type`,`dsa_rule`.`ip_prefix`,`dsa_rule`.`ip_prefix_len`,`dsa_rule`.`creator`,`dsa_rule`.`user_visible`,`dsa_rule`.`last_modified`,`dsa_rule`.`group`,`dsa_rule`.`group_access`,`dsa_rule`.`owner`,`dsa_rule`.`owner_access`,`dsa_rule`.`other_access`,`dsa_rule`.`enable`,`dsa_rule`.`description`,`dsa_rule`.`created`,`dsa_rule`.`display_name`,`dsa_rule`.`key_value_pair`,`dsa_rule`.`global_access`,`dsa_rule`.`share`,`dsa_rule`.`perms2_owner`,`dsa_rule`.`perms2_owner_access` from `dsa_rule` where uuid = ?"

func CreateDsaRule(tx *sql.Tx, model *models.DsaRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDsaRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
		utils.MustJSON(model.DsaRuleEntry.Subscriber),
		string(model.DsaRuleEntry.Publisher.EpVersion),
		string(model.DsaRuleEntry.Publisher.EpID),
		string(model.DsaRuleEntry.Publisher.EpType),
		string(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefix),
		int(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess))

	return err
}

func scanDsaRule(rows *sql.Rows) (*models.DsaRule, error) {
	m := models.MakeDsaRule()

	var jsonFQName string

	var jsonDsaRuleEntrySubscriber string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
		&jsonDsaRuleEntrySubscriber,
		&m.DsaRuleEntry.Publisher.EpVersion,
		&m.DsaRuleEntry.Publisher.EpID,
		&m.DsaRuleEntry.Publisher.EpType,
		&m.DsaRuleEntry.Publisher.EpPrefix.IPPrefix,
		&m.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonDsaRuleEntrySubscriber), &m.DsaRuleEntry.Subscriber)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildDsaRuleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["ep_version"]; ok {
		results = append(results, "ep_version = ?")
		values = append(values, value)
	}

	if value, ok := where["ep_id"]; ok {
		results = append(results, "ep_id = ?")
		values = append(values, value)
	}

	if value, ok := where["ep_type"]; ok {
		results = append(results, "ep_type = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
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

func ListDsaRule(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.DsaRule, error) {
	result := models.MakeDsaRuleSlice()
	whereQuery, values := buildDsaRuleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listDsaRuleQuery)
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
		m, _ := scanDsaRule(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDsaRule(tx *sql.Tx, uuid string) (*models.DsaRule, error) {
	rows, err := tx.Query(showDsaRuleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDsaRule(rows)
	}
	return nil, nil
}

func UpdateDsaRule(tx *sql.Tx, uuid string, model *models.DsaRule) error {
	return nil
}

func DeleteDsaRule(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDsaRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
