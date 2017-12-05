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

const insertAppformixNodeRoleQuery = "insert into `appformix_node_role` (`uuid`,`fq_name`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`provisioning_progress_stage`,`provisioning_log`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAppformixNodeRoleQuery = "update `appformix_node_role` set `uuid` = ?,`fq_name` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`provisioning_progress_stage` = ?,`provisioning_log` = ?;"
const deleteAppformixNodeRoleQuery = "delete from `appformix_node_role` where uuid = ?"
const listAppformixNodeRoleQuery = "select `appformix_node_role`.`uuid`,`appformix_node_role`.`fq_name`,`appformix_node_role`.`provisioning_progress`,`appformix_node_role`.`provisioning_start_time`,`appformix_node_role`.`provisioning_state`,`appformix_node_role`.`created`,`appformix_node_role`.`creator`,`appformix_node_role`.`user_visible`,`appformix_node_role`.`last_modified`,`appformix_node_role`.`owner_access`,`appformix_node_role`.`other_access`,`appformix_node_role`.`group`,`appformix_node_role`.`group_access`,`appformix_node_role`.`owner`,`appformix_node_role`.`enable`,`appformix_node_role`.`description`,`appformix_node_role`.`display_name`,`appformix_node_role`.`key_value_pair`,`appformix_node_role`.`global_access`,`appformix_node_role`.`share`,`appformix_node_role`.`perms2_owner`,`appformix_node_role`.`perms2_owner_access`,`appformix_node_role`.`provisioning_progress_stage`,`appformix_node_role`.`provisioning_log` from `appformix_node_role`"
const showAppformixNodeRoleQuery = "select `appformix_node_role`.`uuid`,`appformix_node_role`.`fq_name`,`appformix_node_role`.`provisioning_progress`,`appformix_node_role`.`provisioning_start_time`,`appformix_node_role`.`provisioning_state`,`appformix_node_role`.`created`,`appformix_node_role`.`creator`,`appformix_node_role`.`user_visible`,`appformix_node_role`.`last_modified`,`appformix_node_role`.`owner_access`,`appformix_node_role`.`other_access`,`appformix_node_role`.`group`,`appformix_node_role`.`group_access`,`appformix_node_role`.`owner`,`appformix_node_role`.`enable`,`appformix_node_role`.`description`,`appformix_node_role`.`display_name`,`appformix_node_role`.`key_value_pair`,`appformix_node_role`.`global_access`,`appformix_node_role`.`share`,`appformix_node_role`.`perms2_owner`,`appformix_node_role`.`perms2_owner_access`,`appformix_node_role`.`provisioning_progress_stage`,`appformix_node_role`.`provisioning_log` from `appformix_node_role` where uuid = ?"

func CreateAppformixNodeRole(tx *sql.Tx, model *models.AppformixNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAppformixNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		string(model.IDPerms.Created),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningLog))

	return err
}

func scanAppformixNodeRole(rows *sql.Rows) (*models.AppformixNodeRole, error) {
	m := models.MakeAppformixNodeRole()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
		&m.ProvisioningProgress,
		&m.ProvisioningStartTime,
		&m.ProvisioningState,
		&m.IDPerms.Created,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.ProvisioningProgressStage,
		&m.ProvisioningLog); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildAppformixNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
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

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListAppformixNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.AppformixNodeRole, error) {
	result := models.MakeAppformixNodeRoleSlice()
	whereQuery, values := buildAppformixNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAppformixNodeRoleQuery)
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
		m, _ := scanAppformixNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAppformixNodeRole(tx *sql.Tx, uuid string) (*models.AppformixNodeRole, error) {
	rows, err := tx.Query(showAppformixNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAppformixNodeRole(rows)
	}
	return nil, nil
}

func UpdateAppformixNodeRole(tx *sql.Tx, uuid string, model *models.AppformixNodeRole) error {
	return nil
}

func DeleteAppformixNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAppformixNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
