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

const insertControllerNodeRoleQuery = "insert into `controller_node_role` (`internalapi_bond_interface_members`,`performance_drives`,`provisioning_log`,`provisioning_progress`,`provisioning_progress_stage`,`capacity_drives`,`storage_management_bond_interface_members`,`fq_name`,`uuid`,`provisioning_start_time`,`provisioning_state`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateControllerNodeRoleQuery = "update `controller_node_role` set `internalapi_bond_interface_members` = ?,`performance_drives` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`capacity_drives` = ?,`storage_management_bond_interface_members` = ?,`fq_name` = ?,`uuid` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteControllerNodeRoleQuery = "delete from `controller_node_role` where uuid = ?"
const listControllerNodeRoleQuery = "select `controller_node_role`.`internalapi_bond_interface_members`,`controller_node_role`.`performance_drives`,`controller_node_role`.`provisioning_log`,`controller_node_role`.`provisioning_progress`,`controller_node_role`.`provisioning_progress_stage`,`controller_node_role`.`capacity_drives`,`controller_node_role`.`storage_management_bond_interface_members`,`controller_node_role`.`fq_name`,`controller_node_role`.`uuid`,`controller_node_role`.`provisioning_start_time`,`controller_node_role`.`provisioning_state`,`controller_node_role`.`enable`,`controller_node_role`.`description`,`controller_node_role`.`created`,`controller_node_role`.`creator`,`controller_node_role`.`user_visible`,`controller_node_role`.`last_modified`,`controller_node_role`.`owner`,`controller_node_role`.`owner_access`,`controller_node_role`.`other_access`,`controller_node_role`.`group`,`controller_node_role`.`group_access`,`controller_node_role`.`display_name`,`controller_node_role`.`key_value_pair`,`controller_node_role`.`perms2_owner_access`,`controller_node_role`.`global_access`,`controller_node_role`.`share`,`controller_node_role`.`perms2_owner` from `controller_node_role`"
const showControllerNodeRoleQuery = "select `controller_node_role`.`internalapi_bond_interface_members`,`controller_node_role`.`performance_drives`,`controller_node_role`.`provisioning_log`,`controller_node_role`.`provisioning_progress`,`controller_node_role`.`provisioning_progress_stage`,`controller_node_role`.`capacity_drives`,`controller_node_role`.`storage_management_bond_interface_members`,`controller_node_role`.`fq_name`,`controller_node_role`.`uuid`,`controller_node_role`.`provisioning_start_time`,`controller_node_role`.`provisioning_state`,`controller_node_role`.`enable`,`controller_node_role`.`description`,`controller_node_role`.`created`,`controller_node_role`.`creator`,`controller_node_role`.`user_visible`,`controller_node_role`.`last_modified`,`controller_node_role`.`owner`,`controller_node_role`.`owner_access`,`controller_node_role`.`other_access`,`controller_node_role`.`group`,`controller_node_role`.`group_access`,`controller_node_role`.`display_name`,`controller_node_role`.`key_value_pair`,`controller_node_role`.`perms2_owner_access`,`controller_node_role`.`global_access`,`controller_node_role`.`share`,`controller_node_role`.`perms2_owner` from `controller_node_role` where uuid = ?"

func CreateControllerNodeRole(tx *sql.Tx, model *models.ControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.InternalapiBondInterfaceMembers),
		string(model.PerformanceDrives),
		string(model.ProvisioningLog),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.CapacityDrives),
		string(model.StorageManagementBondInterfaceMembers),
		utils.MustJSON(model.FQName),
		string(model.UUID),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))

	return err
}

func scanControllerNodeRole(rows *sql.Rows) (*models.ControllerNodeRole, error) {
	m := models.MakeControllerNodeRole()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.InternalapiBondInterfaceMembers,
		&m.PerformanceDrives,
		&m.ProvisioningLog,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.CapacityDrives,
		&m.StorageManagementBondInterfaceMembers,
		&jsonFQName,
		&m.UUID,
		&m.ProvisioningStartTime,
		&m.ProvisioningState,
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

func buildControllerNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["internalapi_bond_interface_members"]; ok {
		results = append(results, "internalapi_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["performance_drives"]; ok {
		results = append(results, "performance_drives = ?")
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

	if value, ok := where["capacity_drives"]; ok {
		results = append(results, "capacity_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["storage_management_bond_interface_members"]; ok {
		results = append(results, "storage_management_bond_interface_members = ?")
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

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
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

func ListControllerNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ControllerNodeRole, error) {
	result := models.MakeControllerNodeRoleSlice()
	whereQuery, values := buildControllerNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listControllerNodeRoleQuery)
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
		m, _ := scanControllerNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowControllerNodeRole(tx *sql.Tx, uuid string) (*models.ControllerNodeRole, error) {
	rows, err := tx.Query(showControllerNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanControllerNodeRole(rows)
	}
	return nil, nil
}

func UpdateControllerNodeRole(tx *sql.Tx, uuid string, model *models.ControllerNodeRole) error {
	return nil
}

func DeleteControllerNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
