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

const insertOpenstackComputeNodeRoleQuery = "insert into `openstack_compute_node_role` (`vrouter_bond_interface`,`display_name`,`provisioning_state`,`provisioning_start_time`,`provisioning_log`,`fq_name`,`provisioning_progress`,`provisioning_progress_stage`,`global_access`,`share`,`owner`,`owner_access`,`default_gateway`,`vrouter_bond_interface_members`,`vrouter_type`,`uuid`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackComputeNodeRoleQuery = "update `openstack_compute_node_role` set `vrouter_bond_interface` = ?,`display_name` = ?,`provisioning_state` = ?,`provisioning_start_time` = ?,`provisioning_log` = ?,`fq_name` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`default_gateway` = ?,`vrouter_bond_interface_members` = ?,`vrouter_type` = ?,`uuid` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`key_value_pair` = ?;"
const deleteOpenstackComputeNodeRoleQuery = "delete from `openstack_compute_node_role` where uuid = ?"
const listOpenstackComputeNodeRoleQuery = "select `openstack_compute_node_role`.`vrouter_bond_interface`,`openstack_compute_node_role`.`display_name`,`openstack_compute_node_role`.`provisioning_state`,`openstack_compute_node_role`.`provisioning_start_time`,`openstack_compute_node_role`.`provisioning_log`,`openstack_compute_node_role`.`fq_name`,`openstack_compute_node_role`.`provisioning_progress`,`openstack_compute_node_role`.`provisioning_progress_stage`,`openstack_compute_node_role`.`global_access`,`openstack_compute_node_role`.`share`,`openstack_compute_node_role`.`owner`,`openstack_compute_node_role`.`owner_access`,`openstack_compute_node_role`.`default_gateway`,`openstack_compute_node_role`.`vrouter_bond_interface_members`,`openstack_compute_node_role`.`vrouter_type`,`openstack_compute_node_role`.`uuid`,`openstack_compute_node_role`.`enable`,`openstack_compute_node_role`.`description`,`openstack_compute_node_role`.`created`,`openstack_compute_node_role`.`creator`,`openstack_compute_node_role`.`user_visible`,`openstack_compute_node_role`.`last_modified`,`openstack_compute_node_role`.`other_access`,`openstack_compute_node_role`.`group`,`openstack_compute_node_role`.`group_access`,`openstack_compute_node_role`.`permissions_owner`,`openstack_compute_node_role`.`permissions_owner_access`,`openstack_compute_node_role`.`key_value_pair` from `openstack_compute_node_role`"
const showOpenstackComputeNodeRoleQuery = "select `openstack_compute_node_role`.`vrouter_bond_interface`,`openstack_compute_node_role`.`display_name`,`openstack_compute_node_role`.`provisioning_state`,`openstack_compute_node_role`.`provisioning_start_time`,`openstack_compute_node_role`.`provisioning_log`,`openstack_compute_node_role`.`fq_name`,`openstack_compute_node_role`.`provisioning_progress`,`openstack_compute_node_role`.`provisioning_progress_stage`,`openstack_compute_node_role`.`global_access`,`openstack_compute_node_role`.`share`,`openstack_compute_node_role`.`owner`,`openstack_compute_node_role`.`owner_access`,`openstack_compute_node_role`.`default_gateway`,`openstack_compute_node_role`.`vrouter_bond_interface_members`,`openstack_compute_node_role`.`vrouter_type`,`openstack_compute_node_role`.`uuid`,`openstack_compute_node_role`.`enable`,`openstack_compute_node_role`.`description`,`openstack_compute_node_role`.`created`,`openstack_compute_node_role`.`creator`,`openstack_compute_node_role`.`user_visible`,`openstack_compute_node_role`.`last_modified`,`openstack_compute_node_role`.`other_access`,`openstack_compute_node_role`.`group`,`openstack_compute_node_role`.`group_access`,`openstack_compute_node_role`.`permissions_owner`,`openstack_compute_node_role`.`permissions_owner_access`,`openstack_compute_node_role`.`key_value_pair` from `openstack_compute_node_role` where uuid = ?"

func CreateOpenstackComputeNodeRole(tx *sql.Tx, model *models.OpenstackComputeNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackComputeNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.VrouterBondInterface),
		string(model.DisplayName),
		string(model.ProvisioningState),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningLog),
		utils.MustJSON(model.FQName),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.DefaultGateway),
		string(model.VrouterBondInterfaceMembers),
		string(model.VrouterType),
		string(model.UUID),
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
		utils.MustJSON(model.Annotations.KeyValuePair))

	return err
}

func scanOpenstackComputeNodeRole(rows *sql.Rows) (*models.OpenstackComputeNodeRole, error) {
	m := models.MakeOpenstackComputeNodeRole()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.VrouterBondInterface,
		&m.DisplayName,
		&m.ProvisioningState,
		&m.ProvisioningStartTime,
		&m.ProvisioningLog,
		&jsonFQName,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.DefaultGateway,
		&m.VrouterBondInterfaceMembers,
		&m.VrouterType,
		&m.UUID,
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
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildOpenstackComputeNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["vrouter_bond_interface"]; ok {
		results = append(results, "vrouter_bond_interface = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["default_gateway"]; ok {
		results = append(results, "default_gateway = ?")
		values = append(values, value)
	}

	if value, ok := where["vrouter_bond_interface_members"]; ok {
		results = append(results, "vrouter_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["vrouter_type"]; ok {
		results = append(results, "vrouter_type = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListOpenstackComputeNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.OpenstackComputeNodeRole, error) {
	result := models.MakeOpenstackComputeNodeRoleSlice()
	whereQuery, values := buildOpenstackComputeNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listOpenstackComputeNodeRoleQuery)
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
		m, _ := scanOpenstackComputeNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowOpenstackComputeNodeRole(tx *sql.Tx, uuid string) (*models.OpenstackComputeNodeRole, error) {
	rows, err := tx.Query(showOpenstackComputeNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanOpenstackComputeNodeRole(rows)
	}
	return nil, nil
}

func UpdateOpenstackComputeNodeRole(tx *sql.Tx, uuid string, model *models.OpenstackComputeNodeRole) error {
	return nil
}

func DeleteOpenstackComputeNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackComputeNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
