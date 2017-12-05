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

const insertLoadbalancerHealthmonitorQuery = "insert into `loadbalancer_healthmonitor` (`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`expected_codes`,`max_retries`,`http_method`,`admin_state`,`timeout`,`url_path`,`monitor_type`,`delay`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerHealthmonitorQuery = "update `loadbalancer_healthmonitor` set `key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`expected_codes` = ?,`max_retries` = ?,`http_method` = ?,`admin_state` = ?,`timeout` = ?,`url_path` = ?,`monitor_type` = ?,`delay` = ?;"
const deleteLoadbalancerHealthmonitorQuery = "delete from `loadbalancer_healthmonitor` where uuid = ?"
const listLoadbalancerHealthmonitorQuery = "select `loadbalancer_healthmonitor`.`key_value_pair`,`loadbalancer_healthmonitor`.`owner_access`,`loadbalancer_healthmonitor`.`global_access`,`loadbalancer_healthmonitor`.`share`,`loadbalancer_healthmonitor`.`owner`,`loadbalancer_healthmonitor`.`uuid`,`loadbalancer_healthmonitor`.`fq_name`,`loadbalancer_healthmonitor`.`user_visible`,`loadbalancer_healthmonitor`.`last_modified`,`loadbalancer_healthmonitor`.`permissions_owner`,`loadbalancer_healthmonitor`.`permissions_owner_access`,`loadbalancer_healthmonitor`.`other_access`,`loadbalancer_healthmonitor`.`group`,`loadbalancer_healthmonitor`.`group_access`,`loadbalancer_healthmonitor`.`enable`,`loadbalancer_healthmonitor`.`description`,`loadbalancer_healthmonitor`.`created`,`loadbalancer_healthmonitor`.`creator`,`loadbalancer_healthmonitor`.`display_name`,`loadbalancer_healthmonitor`.`expected_codes`,`loadbalancer_healthmonitor`.`max_retries`,`loadbalancer_healthmonitor`.`http_method`,`loadbalancer_healthmonitor`.`admin_state`,`loadbalancer_healthmonitor`.`timeout`,`loadbalancer_healthmonitor`.`url_path`,`loadbalancer_healthmonitor`.`monitor_type`,`loadbalancer_healthmonitor`.`delay` from `loadbalancer_healthmonitor`"
const showLoadbalancerHealthmonitorQuery = "select `loadbalancer_healthmonitor`.`key_value_pair`,`loadbalancer_healthmonitor`.`owner_access`,`loadbalancer_healthmonitor`.`global_access`,`loadbalancer_healthmonitor`.`share`,`loadbalancer_healthmonitor`.`owner`,`loadbalancer_healthmonitor`.`uuid`,`loadbalancer_healthmonitor`.`fq_name`,`loadbalancer_healthmonitor`.`user_visible`,`loadbalancer_healthmonitor`.`last_modified`,`loadbalancer_healthmonitor`.`permissions_owner`,`loadbalancer_healthmonitor`.`permissions_owner_access`,`loadbalancer_healthmonitor`.`other_access`,`loadbalancer_healthmonitor`.`group`,`loadbalancer_healthmonitor`.`group_access`,`loadbalancer_healthmonitor`.`enable`,`loadbalancer_healthmonitor`.`description`,`loadbalancer_healthmonitor`.`created`,`loadbalancer_healthmonitor`.`creator`,`loadbalancer_healthmonitor`.`display_name`,`loadbalancer_healthmonitor`.`expected_codes`,`loadbalancer_healthmonitor`.`max_retries`,`loadbalancer_healthmonitor`.`http_method`,`loadbalancer_healthmonitor`.`admin_state`,`loadbalancer_healthmonitor`.`timeout`,`loadbalancer_healthmonitor`.`url_path`,`loadbalancer_healthmonitor`.`monitor_type`,`loadbalancer_healthmonitor`.`delay` from `loadbalancer_healthmonitor` where uuid = ?"

func CreateLoadbalancerHealthmonitor(tx *sql.Tx, model *models.LoadbalancerHealthmonitor) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerHealthmonitorQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		utils.MustJSON(model.FQName),
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
		string(model.IDPerms.Creator),
		string(model.DisplayName),
		string(model.LoadbalancerHealthmonitorProperties.ExpectedCodes),
		int(model.LoadbalancerHealthmonitorProperties.MaxRetries),
		string(model.LoadbalancerHealthmonitorProperties.HTTPMethod),
		bool(model.LoadbalancerHealthmonitorProperties.AdminState),
		int(model.LoadbalancerHealthmonitorProperties.Timeout),
		string(model.LoadbalancerHealthmonitorProperties.URLPath),
		string(model.LoadbalancerHealthmonitorProperties.MonitorType),
		int(model.LoadbalancerHealthmonitorProperties.Delay))

	return err
}

func scanLoadbalancerHealthmonitor(rows *sql.Rows) (*models.LoadbalancerHealthmonitor, error) {
	m := models.MakeLoadbalancerHealthmonitor()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Creator,
		&m.DisplayName,
		&m.LoadbalancerHealthmonitorProperties.ExpectedCodes,
		&m.LoadbalancerHealthmonitorProperties.MaxRetries,
		&m.LoadbalancerHealthmonitorProperties.HTTPMethod,
		&m.LoadbalancerHealthmonitorProperties.AdminState,
		&m.LoadbalancerHealthmonitorProperties.Timeout,
		&m.LoadbalancerHealthmonitorProperties.URLPath,
		&m.LoadbalancerHealthmonitorProperties.MonitorType,
		&m.LoadbalancerHealthmonitorProperties.Delay); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildLoadbalancerHealthmonitorWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["expected_codes"]; ok {
		results = append(results, "expected_codes = ?")
		values = append(values, value)
	}

	if value, ok := where["http_method"]; ok {
		results = append(results, "http_method = ?")
		values = append(values, value)
	}

	if value, ok := where["url_path"]; ok {
		results = append(results, "url_path = ?")
		values = append(values, value)
	}

	if value, ok := where["monitor_type"]; ok {
		results = append(results, "monitor_type = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLoadbalancerHealthmonitor(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.LoadbalancerHealthmonitor, error) {
	result := models.MakeLoadbalancerHealthmonitorSlice()
	whereQuery, values := buildLoadbalancerHealthmonitorWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLoadbalancerHealthmonitorQuery)
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
		m, _ := scanLoadbalancerHealthmonitor(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancerHealthmonitor(tx *sql.Tx, uuid string) (*models.LoadbalancerHealthmonitor, error) {
	rows, err := tx.Query(showLoadbalancerHealthmonitorQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancerHealthmonitor(rows)
	}
	return nil, nil
}

func UpdateLoadbalancerHealthmonitor(tx *sql.Tx, uuid string, model *models.LoadbalancerHealthmonitor) error {
	return nil
}

func DeleteLoadbalancerHealthmonitor(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerHealthmonitorQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
