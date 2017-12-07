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

const insertContrailControllerNodeRoleQuery = "insert into `contrail_controller_node_role` (`key_value_pair`,`uuid`,`provisioning_log`,`provisioning_progress_stage`,`provisioning_start_time`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`provisioning_progress`,`provisioning_state`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailControllerNodeRoleQuery = "update `contrail_controller_node_role` set `key_value_pair` = ?,`uuid` = ?,`provisioning_log` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`display_name` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`provisioning_progress` = ?,`provisioning_state` = ?;"
const deleteContrailControllerNodeRoleQuery = "delete from `contrail_controller_node_role` where uuid = ?"

// ContrailControllerNodeRoleFields is db columns for ContrailControllerNodeRole
var ContrailControllerNodeRoleFields = []string{
	"key_value_pair",
	"uuid",
	"provisioning_log",
	"provisioning_progress_stage",
	"provisioning_start_time",
	"fq_name",
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
	"display_name",
	"perms2_owner_access",
	"global_access",
	"share",
	"perms2_owner",
	"provisioning_progress",
	"provisioning_state",
}

// ContrailControllerNodeRoleRefFields is db reference fields for ContrailControllerNodeRole
var ContrailControllerNodeRoleRefFields = map[string][]string{}

// CreateContrailControllerNodeRole inserts ContrailControllerNodeRole to DB
func CreateContrailControllerNodeRole(tx *sql.Tx, model *models.ContrailControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailControllerNodeRoleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertContrailControllerNodeRoleQuery,
	}).Debug("create query")
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		string(model.ProvisioningLog),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningStartTime),
		utils.MustJSON(model.FQName),
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
		string(model.DisplayName),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.ProvisioningProgress),
		string(model.ProvisioningState))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanContrailControllerNodeRole(values map[string]interface{}) (*models.ContrailControllerNodeRole, error) {
	m := models.MakeContrailControllerNodeRole()

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["provisioning_log"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningLog = castedValue

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["provisioning_start_time"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningStartTime = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

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

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

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

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["provisioning_progress"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ProvisioningProgress = castedValue

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	return m, nil
}

// ListContrailControllerNodeRole lists ContrailControllerNodeRole with list spec.
func ListContrailControllerNodeRole(tx *sql.Tx, spec *db.ListSpec) ([]*models.ContrailControllerNodeRole, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "contrail_controller_node_role"
	spec.Fields = ContrailControllerNodeRoleFields
	spec.RefFields = ContrailControllerNodeRoleRefFields
	result := models.MakeContrailControllerNodeRoleSlice()
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
		m, err := scanContrailControllerNodeRole(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowContrailControllerNodeRole shows ContrailControllerNodeRole resource
func ShowContrailControllerNodeRole(tx *sql.Tx, uuid string) (*models.ContrailControllerNodeRole, error) {
	list, err := ListContrailControllerNodeRole(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateContrailControllerNodeRole updates a resource
func UpdateContrailControllerNodeRole(tx *sql.Tx, uuid string, model *models.ContrailControllerNodeRole) error {
	//TODO(nati) support update
	return nil
}

// DeleteContrailControllerNodeRole deletes a resource
func DeleteContrailControllerNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailControllerNodeRoleQuery)
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
