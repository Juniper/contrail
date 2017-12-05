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

const insertContrailAnalyticsDatabaseNodeRoleQuery = "insert into `contrail_analytics_database_node_role` (`provisioning_start_time`,`fq_name`,`display_name`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`provisioning_progress`,`provisioning_progress_stage`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`key_value_pair`,`provisioning_state`,`provisioning_log`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailAnalyticsDatabaseNodeRoleQuery = "update `contrail_analytics_database_node_role` set `provisioning_start_time` = ?,`fq_name` = ?,`display_name` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`key_value_pair` = ?,`provisioning_state` = ?,`provisioning_log` = ?;"
const deleteContrailAnalyticsDatabaseNodeRoleQuery = "delete from `contrail_analytics_database_node_role` where uuid = ?"
const listContrailAnalyticsDatabaseNodeRoleQuery = "select `contrail_analytics_database_node_role`.`provisioning_start_time`,`contrail_analytics_database_node_role`.`fq_name`,`contrail_analytics_database_node_role`.`display_name`,`contrail_analytics_database_node_role`.`global_access`,`contrail_analytics_database_node_role`.`share`,`contrail_analytics_database_node_role`.`owner`,`contrail_analytics_database_node_role`.`owner_access`,`contrail_analytics_database_node_role`.`uuid`,`contrail_analytics_database_node_role`.`provisioning_progress`,`contrail_analytics_database_node_role`.`provisioning_progress_stage`,`contrail_analytics_database_node_role`.`created`,`contrail_analytics_database_node_role`.`creator`,`contrail_analytics_database_node_role`.`user_visible`,`contrail_analytics_database_node_role`.`last_modified`,`contrail_analytics_database_node_role`.`other_access`,`contrail_analytics_database_node_role`.`group`,`contrail_analytics_database_node_role`.`group_access`,`contrail_analytics_database_node_role`.`permissions_owner`,`contrail_analytics_database_node_role`.`permissions_owner_access`,`contrail_analytics_database_node_role`.`enable`,`contrail_analytics_database_node_role`.`description`,`contrail_analytics_database_node_role`.`key_value_pair`,`contrail_analytics_database_node_role`.`provisioning_state`,`contrail_analytics_database_node_role`.`provisioning_log` from `contrail_analytics_database_node_role`"
const showContrailAnalyticsDatabaseNodeRoleQuery = "select `contrail_analytics_database_node_role`.`provisioning_start_time`,`contrail_analytics_database_node_role`.`fq_name`,`contrail_analytics_database_node_role`.`display_name`,`contrail_analytics_database_node_role`.`global_access`,`contrail_analytics_database_node_role`.`share`,`contrail_analytics_database_node_role`.`owner`,`contrail_analytics_database_node_role`.`owner_access`,`contrail_analytics_database_node_role`.`uuid`,`contrail_analytics_database_node_role`.`provisioning_progress`,`contrail_analytics_database_node_role`.`provisioning_progress_stage`,`contrail_analytics_database_node_role`.`created`,`contrail_analytics_database_node_role`.`creator`,`contrail_analytics_database_node_role`.`user_visible`,`contrail_analytics_database_node_role`.`last_modified`,`contrail_analytics_database_node_role`.`other_access`,`contrail_analytics_database_node_role`.`group`,`contrail_analytics_database_node_role`.`group_access`,`contrail_analytics_database_node_role`.`permissions_owner`,`contrail_analytics_database_node_role`.`permissions_owner_access`,`contrail_analytics_database_node_role`.`enable`,`contrail_analytics_database_node_role`.`description`,`contrail_analytics_database_node_role`.`key_value_pair`,`contrail_analytics_database_node_role`.`provisioning_state`,`contrail_analytics_database_node_role`.`provisioning_log` from `contrail_analytics_database_node_role` where uuid = ?"

func CreateContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, model *models.ContrailAnalyticsDatabaseNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsDatabaseNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ProvisioningStartTime),
		utils.MustJSON(model.FQName),
		string(model.DisplayName),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningState),
		string(model.ProvisioningLog))

	return err
}

func scanContrailAnalyticsDatabaseNodeRole(rows *sql.Rows) (*models.ContrailAnalyticsDatabaseNodeRole, error) {
	m := models.MakeContrailAnalyticsDatabaseNodeRole()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.ProvisioningStartTime,
		&jsonFQName,
		&m.DisplayName,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&jsonAnnotationsKeyValuePair,
		&m.ProvisioningState,
		&m.ProvisioningLog); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildContrailAnalyticsDatabaseNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
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

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ContrailAnalyticsDatabaseNodeRole, error) {
	result := models.MakeContrailAnalyticsDatabaseNodeRoleSlice()
	whereQuery, values := buildContrailAnalyticsDatabaseNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listContrailAnalyticsDatabaseNodeRoleQuery)
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
		m, _ := scanContrailAnalyticsDatabaseNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string) (*models.ContrailAnalyticsDatabaseNodeRole, error) {
	rows, err := tx.Query(showContrailAnalyticsDatabaseNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanContrailAnalyticsDatabaseNodeRole(rows)
	}
	return nil, nil
}

func UpdateContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string, model *models.ContrailAnalyticsDatabaseNodeRole) error {
	return nil
}

func DeleteContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailAnalyticsDatabaseNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
