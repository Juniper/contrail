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

const insertContrailAnalyticsNodeQuery = "insert into `contrail_analytics_node` (`provisioning_state`,`provisioning_log`,`provisioning_progress_stage`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`key_value_pair`,`uuid`,`fq_name`,`provisioning_start_time`,`provisioning_progress`,`display_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailAnalyticsNodeQuery = "update `contrail_analytics_node` set `provisioning_state` = ?,`provisioning_log` = ?,`provisioning_progress_stage` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`key_value_pair` = ?,`uuid` = ?,`fq_name` = ?,`provisioning_start_time` = ?,`provisioning_progress` = ?,`display_name` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteContrailAnalyticsNodeQuery = "delete from `contrail_analytics_node` where uuid = ?"
const listContrailAnalyticsNodeQuery = "select `contrail_analytics_node`.`provisioning_state`,`contrail_analytics_node`.`provisioning_log`,`contrail_analytics_node`.`provisioning_progress_stage`,`contrail_analytics_node`.`enable`,`contrail_analytics_node`.`description`,`contrail_analytics_node`.`created`,`contrail_analytics_node`.`creator`,`contrail_analytics_node`.`user_visible`,`contrail_analytics_node`.`last_modified`,`contrail_analytics_node`.`owner`,`contrail_analytics_node`.`owner_access`,`contrail_analytics_node`.`other_access`,`contrail_analytics_node`.`group`,`contrail_analytics_node`.`group_access`,`contrail_analytics_node`.`key_value_pair`,`contrail_analytics_node`.`uuid`,`contrail_analytics_node`.`fq_name`,`contrail_analytics_node`.`provisioning_start_time`,`contrail_analytics_node`.`provisioning_progress`,`contrail_analytics_node`.`display_name`,`contrail_analytics_node`.`perms2_owner`,`contrail_analytics_node`.`perms2_owner_access`,`contrail_analytics_node`.`global_access`,`contrail_analytics_node`.`share` from `contrail_analytics_node`"
const showContrailAnalyticsNodeQuery = "select `contrail_analytics_node`.`provisioning_state`,`contrail_analytics_node`.`provisioning_log`,`contrail_analytics_node`.`provisioning_progress_stage`,`contrail_analytics_node`.`enable`,`contrail_analytics_node`.`description`,`contrail_analytics_node`.`created`,`contrail_analytics_node`.`creator`,`contrail_analytics_node`.`user_visible`,`contrail_analytics_node`.`last_modified`,`contrail_analytics_node`.`owner`,`contrail_analytics_node`.`owner_access`,`contrail_analytics_node`.`other_access`,`contrail_analytics_node`.`group`,`contrail_analytics_node`.`group_access`,`contrail_analytics_node`.`key_value_pair`,`contrail_analytics_node`.`uuid`,`contrail_analytics_node`.`fq_name`,`contrail_analytics_node`.`provisioning_start_time`,`contrail_analytics_node`.`provisioning_progress`,`contrail_analytics_node`.`display_name`,`contrail_analytics_node`.`perms2_owner`,`contrail_analytics_node`.`perms2_owner_access`,`contrail_analytics_node`.`global_access`,`contrail_analytics_node`.`share` from `contrail_analytics_node` where uuid = ?"

func CreateContrailAnalyticsNode(tx *sql.Tx, model *models.ContrailAnalyticsNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ProvisioningState),
		string(model.ProvisioningLog),
		string(model.ProvisioningProgressStage),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.ProvisioningStartTime),
		int(model.ProvisioningProgress),
		string(model.DisplayName),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share))

	return err
}

func scanContrailAnalyticsNode(rows *sql.Rows) (*models.ContrailAnalyticsNode, error) {
	m := models.MakeContrailAnalyticsNode()

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonPerms2Share string

	if err := rows.Scan(&m.ProvisioningState,
		&m.ProvisioningLog,
		&m.ProvisioningProgressStage,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&jsonFQName,
		&m.ProvisioningStartTime,
		&m.ProvisioningProgress,
		&m.DisplayName,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildContrailAnalyticsNodeWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
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

func ListContrailAnalyticsNode(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ContrailAnalyticsNode, error) {
	result := models.MakeContrailAnalyticsNodeSlice()
	whereQuery, values := buildContrailAnalyticsNodeWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listContrailAnalyticsNodeQuery)
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
		m, _ := scanContrailAnalyticsNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowContrailAnalyticsNode(tx *sql.Tx, uuid string) (*models.ContrailAnalyticsNode, error) {
	rows, err := tx.Query(showContrailAnalyticsNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanContrailAnalyticsNode(rows)
	}
	return nil, nil
}

func UpdateContrailAnalyticsNode(tx *sql.Tx, uuid string, model *models.ContrailAnalyticsNode) error {
	return nil
}

func DeleteContrailAnalyticsNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
