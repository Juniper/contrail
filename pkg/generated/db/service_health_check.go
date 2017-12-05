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

const insertServiceHealthCheckQuery = "insert into `service_health_check` (`enabled`,`max_retries`,`health_check_type`,`monitor_type`,`timeoutUsecs`,`http_method`,`timeout`,`delay`,`delayUsecs`,`expected_codes`,`url_path`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceHealthCheckQuery = "update `service_health_check` set `enabled` = ?,`max_retries` = ?,`health_check_type` = ?,`monitor_type` = ?,`timeoutUsecs` = ?,`http_method` = ?,`timeout` = ?,`delay` = ?,`delayUsecs` = ?,`expected_codes` = ?,`url_path` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteServiceHealthCheckQuery = "delete from `service_health_check` where uuid = ?"
const listServiceHealthCheckQuery = "select `service_health_check`.`enabled`,`service_health_check`.`max_retries`,`service_health_check`.`health_check_type`,`service_health_check`.`monitor_type`,`service_health_check`.`timeoutUsecs`,`service_health_check`.`http_method`,`service_health_check`.`timeout`,`service_health_check`.`delay`,`service_health_check`.`delayUsecs`,`service_health_check`.`expected_codes`,`service_health_check`.`url_path`,`service_health_check`.`uuid`,`service_health_check`.`fq_name`,`service_health_check`.`creator`,`service_health_check`.`user_visible`,`service_health_check`.`last_modified`,`service_health_check`.`owner_access`,`service_health_check`.`other_access`,`service_health_check`.`group`,`service_health_check`.`group_access`,`service_health_check`.`owner`,`service_health_check`.`enable`,`service_health_check`.`description`,`service_health_check`.`created`,`service_health_check`.`display_name`,`service_health_check`.`key_value_pair`,`service_health_check`.`perms2_owner_access`,`service_health_check`.`global_access`,`service_health_check`.`share`,`service_health_check`.`perms2_owner` from `service_health_check`"
const showServiceHealthCheckQuery = "select `service_health_check`.`enabled`,`service_health_check`.`max_retries`,`service_health_check`.`health_check_type`,`service_health_check`.`monitor_type`,`service_health_check`.`timeoutUsecs`,`service_health_check`.`http_method`,`service_health_check`.`timeout`,`service_health_check`.`delay`,`service_health_check`.`delayUsecs`,`service_health_check`.`expected_codes`,`service_health_check`.`url_path`,`service_health_check`.`uuid`,`service_health_check`.`fq_name`,`service_health_check`.`creator`,`service_health_check`.`user_visible`,`service_health_check`.`last_modified`,`service_health_check`.`owner_access`,`service_health_check`.`other_access`,`service_health_check`.`group`,`service_health_check`.`group_access`,`service_health_check`.`owner`,`service_health_check`.`enable`,`service_health_check`.`description`,`service_health_check`.`created`,`service_health_check`.`display_name`,`service_health_check`.`key_value_pair`,`service_health_check`.`perms2_owner_access`,`service_health_check`.`global_access`,`service_health_check`.`share`,`service_health_check`.`perms2_owner` from `service_health_check` where uuid = ?"

const insertServiceHealthCheckServiceInstanceQuery = "insert into `ref_service_health_check_service_instance` (`from`, `to` ) values (?, ?);"

func CreateServiceHealthCheck(tx *sql.Tx, model *models.ServiceHealthCheck) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceHealthCheckQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.ServiceHealthCheckProperties.Enabled),
		int(model.ServiceHealthCheckProperties.MaxRetries),
		string(model.ServiceHealthCheckProperties.HealthCheckType),
		string(model.ServiceHealthCheckProperties.MonitorType),
		int(model.ServiceHealthCheckProperties.TimeoutUsecs),
		string(model.ServiceHealthCheckProperties.HTTPMethod),
		int(model.ServiceHealthCheckProperties.Timeout),
		int(model.ServiceHealthCheckProperties.Delay),
		int(model.ServiceHealthCheckProperties.DelayUsecs),
		string(model.ServiceHealthCheckProperties.ExpectedCodes),
		string(model.ServiceHealthCheckProperties.URLPath),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))

	stmtServiceInstanceRef, err := tx.Prepare(insertServiceHealthCheckServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanServiceHealthCheck(rows *sql.Rows) (*models.ServiceHealthCheck, error) {
	m := models.MakeServiceHealthCheck()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.ServiceHealthCheckProperties.Enabled,
		&m.ServiceHealthCheckProperties.MaxRetries,
		&m.ServiceHealthCheckProperties.HealthCheckType,
		&m.ServiceHealthCheckProperties.MonitorType,
		&m.ServiceHealthCheckProperties.TimeoutUsecs,
		&m.ServiceHealthCheckProperties.HTTPMethod,
		&m.ServiceHealthCheckProperties.Timeout,
		&m.ServiceHealthCheckProperties.Delay,
		&m.ServiceHealthCheckProperties.DelayUsecs,
		&m.ServiceHealthCheckProperties.ExpectedCodes,
		&m.ServiceHealthCheckProperties.URLPath,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildServiceHealthCheckWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["health_check_type"]; ok {
		results = append(results, "health_check_type = ?")
		values = append(values, value)
	}

	if value, ok := where["monitor_type"]; ok {
		results = append(results, "monitor_type = ?")
		values = append(values, value)
	}

	if value, ok := where["http_method"]; ok {
		results = append(results, "http_method = ?")
		values = append(values, value)
	}

	if value, ok := where["expected_codes"]; ok {
		results = append(results, "expected_codes = ?")
		values = append(values, value)
	}

	if value, ok := where["url_path"]; ok {
		results = append(results, "url_path = ?")
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

func ListServiceHealthCheck(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceHealthCheck, error) {
	result := models.MakeServiceHealthCheckSlice()
	whereQuery, values := buildServiceHealthCheckWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceHealthCheckQuery)
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
		m, _ := scanServiceHealthCheck(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceHealthCheck(tx *sql.Tx, uuid string) (*models.ServiceHealthCheck, error) {
	rows, err := tx.Query(showServiceHealthCheckQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceHealthCheck(rows)
	}
	return nil, nil
}

func UpdateServiceHealthCheck(tx *sql.Tx, uuid string, model *models.ServiceHealthCheck) error {
	return nil
}

func DeleteServiceHealthCheck(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceHealthCheckQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
