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

const insertContrailControllerNodeRoleQuery = "insert into `contrail_controller_node_role` (`fq_name`,`key_value_pair`,`provisioning_log`,`provisioning_progress`,`provisioning_state`,`provisioning_progress_stage`,`provisioning_start_time`,`uuid`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailControllerNodeRoleQuery = "update `contrail_controller_node_role` set `fq_name` = ?,`key_value_pair` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`provisioning_state` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`uuid` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteContrailControllerNodeRoleQuery = "delete from `contrail_controller_node_role` where uuid = ?"
const listContrailControllerNodeRoleQuery = "select `contrail_controller_node_role`.`fq_name`,`contrail_controller_node_role`.`key_value_pair`,`contrail_controller_node_role`.`provisioning_log`,`contrail_controller_node_role`.`provisioning_progress`,`contrail_controller_node_role`.`provisioning_state`,`contrail_controller_node_role`.`provisioning_progress_stage`,`contrail_controller_node_role`.`provisioning_start_time`,`contrail_controller_node_role`.`uuid`,`contrail_controller_node_role`.`creator`,`contrail_controller_node_role`.`user_visible`,`contrail_controller_node_role`.`last_modified`,`contrail_controller_node_role`.`owner`,`contrail_controller_node_role`.`owner_access`,`contrail_controller_node_role`.`other_access`,`contrail_controller_node_role`.`group`,`contrail_controller_node_role`.`group_access`,`contrail_controller_node_role`.`enable`,`contrail_controller_node_role`.`description`,`contrail_controller_node_role`.`created`,`contrail_controller_node_role`.`display_name`,`contrail_controller_node_role`.`global_access`,`contrail_controller_node_role`.`share`,`contrail_controller_node_role`.`perms2_owner`,`contrail_controller_node_role`.`perms2_owner_access` from `contrail_controller_node_role`"
const showContrailControllerNodeRoleQuery = "select `contrail_controller_node_role`.`fq_name`,`contrail_controller_node_role`.`key_value_pair`,`contrail_controller_node_role`.`provisioning_log`,`contrail_controller_node_role`.`provisioning_progress`,`contrail_controller_node_role`.`provisioning_state`,`contrail_controller_node_role`.`provisioning_progress_stage`,`contrail_controller_node_role`.`provisioning_start_time`,`contrail_controller_node_role`.`uuid`,`contrail_controller_node_role`.`creator`,`contrail_controller_node_role`.`user_visible`,`contrail_controller_node_role`.`last_modified`,`contrail_controller_node_role`.`owner`,`contrail_controller_node_role`.`owner_access`,`contrail_controller_node_role`.`other_access`,`contrail_controller_node_role`.`group`,`contrail_controller_node_role`.`group_access`,`contrail_controller_node_role`.`enable`,`contrail_controller_node_role`.`description`,`contrail_controller_node_role`.`created`,`contrail_controller_node_role`.`display_name`,`contrail_controller_node_role`.`global_access`,`contrail_controller_node_role`.`share`,`contrail_controller_node_role`.`perms2_owner`,`contrail_controller_node_role`.`perms2_owner_access` from `contrail_controller_node_role` where uuid = ?"

func CreateContrailControllerNodeRole(tx *sql.Tx, model *models.ContrailControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningLog),
		int(model.ProvisioningProgress),
		string(model.ProvisioningState),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningStartTime),
		string(model.UUID),
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
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess))

	return err
}

func scanContrailControllerNodeRole(rows *sql.Rows) (*models.ContrailControllerNodeRole, error) {
	m := models.MakeContrailControllerNodeRole()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&jsonAnnotationsKeyValuePair,
		&m.ProvisioningLog,
		&m.ProvisioningProgress,
		&m.ProvisioningState,
		&m.ProvisioningProgressStage,
		&m.ProvisioningStartTime,
		&m.UUID,
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
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildContrailControllerNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListContrailControllerNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ContrailControllerNodeRole, error) {
	result := models.MakeContrailControllerNodeRoleSlice()
	whereQuery, values := buildContrailControllerNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listContrailControllerNodeRoleQuery)
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
		m, _ := scanContrailControllerNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowContrailControllerNodeRole(tx *sql.Tx, uuid string) (*models.ContrailControllerNodeRole, error) {
	rows, err := tx.Query(showContrailControllerNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanContrailControllerNodeRole(rows)
	}
	return nil, nil
}

func UpdateContrailControllerNodeRole(tx *sql.Tx, uuid string, model *models.ContrailControllerNodeRole) error {
	return nil
}

func DeleteContrailControllerNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
