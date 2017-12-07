package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVPNGroupQuery = "insert into `vpn_group` (`key_value_pair`,`provisioning_progress_stage`,`provisioning_state`,`fq_name`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`uuid`,`display_name`,`provisioning_log`,`provisioning_progress`,`provisioning_start_time`,`type`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVPNGroupQuery = "update `vpn_group` set `key_value_pair` = ?,`provisioning_progress_stage` = ?,`provisioning_state` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`uuid` = ?,`display_name` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`type` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteVPNGroupQuery = "delete from `vpn_group` where uuid = ?"

// VPNGroupFields is db columns for VPNGroup
var VPNGroupFields = []string{
	"key_value_pair",
	"provisioning_progress_stage",
	"provisioning_state",
	"fq_name",
	"user_visible",
	"last_modified",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"owner",
	"enable",
	"description",
	"created",
	"creator",
	"uuid",
	"display_name",
	"provisioning_log",
	"provisioning_progress",
	"provisioning_start_time",
	"type",
	"perms2_owner_access",
	"global_access",
	"share",
	"perms2_owner",
}

// VPNGroupRefFields is db reference fields for VPNGroup
var VPNGroupRefFields = map[string][]string{

	"location": {
	// <common.Schema Value>

	},
}

const insertVPNGroupLocationQuery = "insert into `ref_vpn_group_location` (`from`, `to` ) values (?, ?);"

// CreateVPNGroup inserts VPNGroup to DB
func CreateVPNGroup(tx *sql.Tx, model *models.VPNGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVPNGroupQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVPNGroupQuery,
	}).Debug("create query")
	_, err = stmt.Exec(common.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningState),
		common.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
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
		string(model.UUID),
		string(model.DisplayName),
		string(model.ProvisioningLog),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		string(model.Type),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtLocationRef, err := tx.Prepare(insertVPNGroupLocationQuery)
	if err != nil {
		return errors.Wrap(err, "preparing LocationRefs create statement failed")
	}
	defer stmtLocationRef.Close()
	for _, ref := range model.LocationRefs {

		_, err = stmtLocationRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LocationRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanVPNGroup(values map[string]interface{}) (*models.VPNGroup, error) {
	m := models.MakeVPNGroup()

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["owner_access"]; ok {

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

	if value, ok := values["owner"]; ok {

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

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

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

	if value, ok := values["type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Type = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

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

	if value, ok := values["perms2_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["ref_location"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VPNGroupLocationRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.LocationRefs = append(m.LocationRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVPNGroup lists VPNGroup with list spec.
func ListVPNGroup(tx *sql.Tx, spec *common.ListSpec) ([]*models.VPNGroup, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "vpn_group"
	spec.Fields = VPNGroupFields
	spec.RefFields = VPNGroupRefFields
	result := models.MakeVPNGroupSlice()
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
		m, err := scanVPNGroup(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowVPNGroup shows VPNGroup resource
func ShowVPNGroup(tx *sql.Tx, uuid string) (*models.VPNGroup, error) {
	list, err := ListVPNGroup(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateVPNGroup updates a resource
func UpdateVPNGroup(tx *sql.Tx, uuid string, model *models.VPNGroup) error {
	//TODO(nati) support update
	return nil
}

// DeleteVPNGroup deletes a resource
func DeleteVPNGroup(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVPNGroupQuery)
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
