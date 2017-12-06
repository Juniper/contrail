package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertOpenstackStorageNodeRoleQuery = "insert into `openstack_storage_node_role` (`storage_backend_bond_interface_members`,`uuid`,`provisioning_progress_stage`,`provisioning_start_time`,`provisioning_state`,`provisioning_progress`,`fq_name`,`display_name`,`key_value_pair`,`provisioning_log`,`journal_drives`,`osd_drives`,`storage_access_bond_interface_members`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackStorageNodeRoleQuery = "update `openstack_storage_node_role` set `storage_backend_bond_interface_members` = ?,`uuid` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`provisioning_progress` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?,`provisioning_log` = ?,`journal_drives` = ?,`osd_drives` = ?,`storage_access_bond_interface_members` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteOpenstackStorageNodeRoleQuery = "delete from `openstack_storage_node_role` where uuid = ?"

// OpenstackStorageNodeRoleFields is db columns for OpenstackStorageNodeRole
var OpenstackStorageNodeRoleFields = []string{
	"storage_backend_bond_interface_members",
	"uuid",
	"provisioning_progress_stage",
	"provisioning_start_time",
	"provisioning_state",
	"provisioning_progress",
	"fq_name",
	"display_name",
	"key_value_pair",
	"provisioning_log",
	"journal_drives",
	"osd_drives",
	"storage_access_bond_interface_members",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"enable",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
}

// OpenstackStorageNodeRoleRefFields is db reference fields for OpenstackStorageNodeRole
var OpenstackStorageNodeRoleRefFields = map[string][]string{}

// CreateOpenstackStorageNodeRole inserts OpenstackStorageNodeRole to DB
func CreateOpenstackStorageNodeRole(tx *sql.Tx, model *models.OpenstackStorageNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackStorageNodeRoleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertOpenstackStorageNodeRoleQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.StorageBackendBondInterfaceMembers),
		string(model.UUID),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		int(model.ProvisioningProgress),
		utils.MustJSON(model.FQName),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningLog),
		string(model.JournalDrives),
		string(model.OsdDrives),
		string(model.StorageAccessBondInterfaceMembers),
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
		bool(model.IDPerms.Enable),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanOpenstackStorageNodeRole(values map[string]interface{}) (*models.OpenstackStorageNodeRole, error) {
	m := models.MakeOpenstackStorageNodeRole()

	if value, ok := values["storage_backend_bond_interface_members"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.StorageBackendBondInterfaceMembers = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["provisioning_start_time"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningStartTime = castedValue

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	if value, ok := values["provisioning_progress"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ProvisioningProgress = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["provisioning_log"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningLog = castedValue

	}

	if value, ok := values["journal_drives"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.JournalDrives = castedValue

	}

	if value, ok := values["osd_drives"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.OsdDrives = castedValue

	}

	if value, ok := values["storage_access_bond_interface_members"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.StorageAccessBondInterfaceMembers = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	return m, nil
}

// ListOpenstackStorageNodeRole lists OpenstackStorageNodeRole with list spec.
func ListOpenstackStorageNodeRole(tx *sql.Tx, spec *db.ListSpec) ([]*models.OpenstackStorageNodeRole, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "openstack_storage_node_role"
	spec.Fields = OpenstackStorageNodeRoleFields
	spec.RefFields = OpenstackStorageNodeRoleRefFields
	result := models.MakeOpenstackStorageNodeRoleSlice()
	query, columns, values := db.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}
	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanOpenstackStorageNodeRole(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowOpenstackStorageNodeRole shows OpenstackStorageNodeRole resource
func ShowOpenstackStorageNodeRole(tx *sql.Tx, uuid string) (*models.OpenstackStorageNodeRole, error) {
	list, err := ListOpenstackStorageNodeRole(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateOpenstackStorageNodeRole updates a resource
func UpdateOpenstackStorageNodeRole(tx *sql.Tx, uuid string, model *models.OpenstackStorageNodeRole) error {
	//TODO(nati) support update
	return nil
}

// DeleteOpenstackStorageNodeRole deletes a resource
func DeleteOpenstackStorageNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackStorageNodeRoleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
