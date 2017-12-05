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

const insertOpenstackStorageNodeRoleQuery = "insert into `openstack_storage_node_role` (`key_value_pair`,`provisioning_log`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_start_time`,`provisioning_state`,`journal_drives`,`storage_backend_bond_interface_members`,`uuid`,`fq_name`,`osd_drives`,`storage_access_bond_interface_members`,`display_name`,`global_access`,`share`,`owner`,`owner_access`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackStorageNodeRoleQuery = "update `openstack_storage_node_role` set `key_value_pair` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`journal_drives` = ?,`storage_backend_bond_interface_members` = ?,`uuid` = ?,`fq_name` = ?,`osd_drives` = ?,`storage_access_bond_interface_members` = ?,`display_name` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?;"
const deleteOpenstackStorageNodeRoleQuery = "delete from `openstack_storage_node_role` where uuid = ?"
const listOpenstackStorageNodeRoleQuery = "select `openstack_storage_node_role`.`key_value_pair`,`openstack_storage_node_role`.`provisioning_log`,`openstack_storage_node_role`.`provisioning_progress`,`openstack_storage_node_role`.`provisioning_progress_stage`,`openstack_storage_node_role`.`provisioning_start_time`,`openstack_storage_node_role`.`provisioning_state`,`openstack_storage_node_role`.`journal_drives`,`openstack_storage_node_role`.`storage_backend_bond_interface_members`,`openstack_storage_node_role`.`uuid`,`openstack_storage_node_role`.`fq_name`,`openstack_storage_node_role`.`osd_drives`,`openstack_storage_node_role`.`storage_access_bond_interface_members`,`openstack_storage_node_role`.`display_name`,`openstack_storage_node_role`.`global_access`,`openstack_storage_node_role`.`share`,`openstack_storage_node_role`.`owner`,`openstack_storage_node_role`.`owner_access`,`openstack_storage_node_role`.`user_visible`,`openstack_storage_node_role`.`last_modified`,`openstack_storage_node_role`.`other_access`,`openstack_storage_node_role`.`group`,`openstack_storage_node_role`.`group_access`,`openstack_storage_node_role`.`permissions_owner`,`openstack_storage_node_role`.`permissions_owner_access`,`openstack_storage_node_role`.`enable`,`openstack_storage_node_role`.`description`,`openstack_storage_node_role`.`created`,`openstack_storage_node_role`.`creator` from `openstack_storage_node_role`"
const showOpenstackStorageNodeRoleQuery = "select `openstack_storage_node_role`.`key_value_pair`,`openstack_storage_node_role`.`provisioning_log`,`openstack_storage_node_role`.`provisioning_progress`,`openstack_storage_node_role`.`provisioning_progress_stage`,`openstack_storage_node_role`.`provisioning_start_time`,`openstack_storage_node_role`.`provisioning_state`,`openstack_storage_node_role`.`journal_drives`,`openstack_storage_node_role`.`storage_backend_bond_interface_members`,`openstack_storage_node_role`.`uuid`,`openstack_storage_node_role`.`fq_name`,`openstack_storage_node_role`.`osd_drives`,`openstack_storage_node_role`.`storage_access_bond_interface_members`,`openstack_storage_node_role`.`display_name`,`openstack_storage_node_role`.`global_access`,`openstack_storage_node_role`.`share`,`openstack_storage_node_role`.`owner`,`openstack_storage_node_role`.`owner_access`,`openstack_storage_node_role`.`user_visible`,`openstack_storage_node_role`.`last_modified`,`openstack_storage_node_role`.`other_access`,`openstack_storage_node_role`.`group`,`openstack_storage_node_role`.`group_access`,`openstack_storage_node_role`.`permissions_owner`,`openstack_storage_node_role`.`permissions_owner_access`,`openstack_storage_node_role`.`enable`,`openstack_storage_node_role`.`description`,`openstack_storage_node_role`.`created`,`openstack_storage_node_role`.`creator` from `openstack_storage_node_role` where uuid = ?"

func CreateOpenstackStorageNodeRole(tx *sql.Tx, model *models.OpenstackStorageNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackStorageNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningLog),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		string(model.JournalDrives),
		string(model.StorageBackendBondInterfaceMembers),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.OsdDrives),
		string(model.StorageAccessBondInterfaceMembers),
		string(model.DisplayName),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator))

	return err
}

func scanOpenstackStorageNodeRole(rows *sql.Rows) (*models.OpenstackStorageNodeRole, error) {
	m := models.MakeOpenstackStorageNodeRole()

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.ProvisioningLog,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.ProvisioningStartTime,
		&m.ProvisioningState,
		&m.JournalDrives,
		&m.StorageBackendBondInterfaceMembers,
		&m.UUID,
		&jsonFQName,
		&m.OsdDrives,
		&m.StorageAccessBondInterfaceMembers,
		&m.DisplayName,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildOpenstackStorageNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
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

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["journal_drives"]; ok {
		results = append(results, "journal_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["storage_backend_bond_interface_members"]; ok {
		results = append(results, "storage_backend_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["osd_drives"]; ok {
		results = append(results, "osd_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["storage_access_bond_interface_members"]; ok {
		results = append(results, "storage_access_bond_interface_members = ?")
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

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListOpenstackStorageNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.OpenstackStorageNodeRole, error) {
	result := models.MakeOpenstackStorageNodeRoleSlice()
	whereQuery, values := buildOpenstackStorageNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listOpenstackStorageNodeRoleQuery)
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
		m, _ := scanOpenstackStorageNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowOpenstackStorageNodeRole(tx *sql.Tx, uuid string) (*models.OpenstackStorageNodeRole, error) {
	rows, err := tx.Query(showOpenstackStorageNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanOpenstackStorageNodeRole(rows)
	}
	return nil, nil
}

func UpdateOpenstackStorageNodeRole(tx *sql.Tx, uuid string, model *models.OpenstackStorageNodeRole) error {
	return nil
}

func DeleteOpenstackStorageNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackStorageNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
