package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertContrailAnalyticsDatabaseNodeRoleQuery = "insert into `contrail_analytics_database_node_role` (`uuid`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`provisioning_log`,`provisioning_progress`,`provisioning_start_time`,`fq_name`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`provisioning_progress_stage`,`provisioning_state`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailAnalyticsDatabaseNodeRoleQuery = "update `contrail_analytics_database_node_role` set `uuid` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`fq_name` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`provisioning_progress_stage` = ?,`provisioning_state` = ?;"
const deleteContrailAnalyticsDatabaseNodeRoleQuery = "delete from `contrail_analytics_database_node_role` where uuid = ?"

// ContrailAnalyticsDatabaseNodeRoleFields is db columns for ContrailAnalyticsDatabaseNodeRole
var ContrailAnalyticsDatabaseNodeRoleFields = []string{
	"uuid",
	"key_value_pair",
	"owner",
	"owner_access",
	"global_access",
	"share",
	"provisioning_log",
	"provisioning_progress",
	"provisioning_start_time",
	"fq_name",
	"last_modified",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"permissions_owner",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"display_name",
	"provisioning_progress_stage",
	"provisioning_state",
}

// ContrailAnalyticsDatabaseNodeRoleRefFields is db reference fields for ContrailAnalyticsDatabaseNodeRole
var ContrailAnalyticsDatabaseNodeRoleRefFields = map[string][]string{}

// CreateContrailAnalyticsDatabaseNodeRole inserts ContrailAnalyticsDatabaseNodeRole to DB
func CreateContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, model *models.ContrailAnalyticsDatabaseNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsDatabaseNodeRoleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertContrailAnalyticsDatabaseNodeRoleQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.ProvisioningLog),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		common.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningState))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanContrailAnalyticsDatabaseNodeRole(values map[string]interface{}) (*models.ContrailAnalyticsDatabaseNodeRole, error) {
	m := models.MakeContrailAnalyticsDatabaseNodeRole()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["provisioning_log"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningLog = castedValue

	}

	if value, ok := values["provisioning_progress"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ProvisioningProgress = castedValue

	}

	if value, ok := values["provisioning_start_time"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningStartTime = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	return m, nil
}

// ListContrailAnalyticsDatabaseNodeRole lists ContrailAnalyticsDatabaseNodeRole with list spec.
func ListContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, spec *common.ListSpec) ([]*models.ContrailAnalyticsDatabaseNodeRole, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "contrail_analytics_database_node_role"
	spec.Fields = ContrailAnalyticsDatabaseNodeRoleFields
	spec.RefFields = ContrailAnalyticsDatabaseNodeRoleRefFields
	result := models.MakeContrailAnalyticsDatabaseNodeRoleSlice()
	query, columns, values := common.BuildListQuery(spec)
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
		m, err := scanContrailAnalyticsDatabaseNodeRole(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowContrailAnalyticsDatabaseNodeRole shows ContrailAnalyticsDatabaseNodeRole resource
func ShowContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string) (*models.ContrailAnalyticsDatabaseNodeRole, error) {
	list, err := ListContrailAnalyticsDatabaseNodeRole(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateContrailAnalyticsDatabaseNodeRole updates a resource
func UpdateContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string, model *models.ContrailAnalyticsDatabaseNodeRole) error {
	//TODO(nati) support update
	return nil
}

// DeleteContrailAnalyticsDatabaseNodeRole deletes a resource
func DeleteContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailAnalyticsDatabaseNodeRoleQuery)
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
