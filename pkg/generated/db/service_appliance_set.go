package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceApplianceSetQuery = "insert into `service_appliance_set` (`owner_access`,`global_access`,`share`,`owner`,`key_value_pair`,`service_appliance_ha_mode`,`service_appliance_driver`,`fq_name`,`display_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`annotations_key_value_pair`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceSetQuery = "update `service_appliance_set` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`key_value_pair` = ?,`service_appliance_ha_mode` = ?,`service_appliance_driver` = ?,`fq_name` = ?,`display_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`annotations_key_value_pair` = ?,`uuid` = ?;"
const deleteServiceApplianceSetQuery = "delete from `service_appliance_set` where uuid = ?"

// ServiceApplianceSetFields is db columns for ServiceApplianceSet
var ServiceApplianceSetFields = []string{
	"owner_access",
	"global_access",
	"share",
	"owner",
	"key_value_pair",
	"service_appliance_ha_mode",
	"service_appliance_driver",
	"fq_name",
	"display_name",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"enable",
	"description",
	"annotations_key_value_pair",
	"uuid",
}

// ServiceApplianceSetRefFields is db reference fields for ServiceApplianceSet
var ServiceApplianceSetRefFields = map[string][]string{}

// CreateServiceApplianceSet inserts ServiceApplianceSet to DB
func CreateServiceApplianceSet(tx *sql.Tx, model *models.ServiceApplianceSet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceSetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceApplianceSetQuery,
	}).Debug("create query")
	_, err = stmt.Exec(int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		common.MustJSON(model.ServiceApplianceSetProperties.KeyValuePair),
		string(model.ServiceApplianceHaMode),
		string(model.ServiceApplianceDriver),
		common.MustJSON(model.FQName),
		string(model.DisplayName),
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
		string(model.IDPerms.Description),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanServiceApplianceSet(values map[string]interface{}) (*models.ServiceApplianceSet, error) {
	m := models.MakeServiceApplianceSet()

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

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceApplianceSetProperties.KeyValuePair)

	}

	if value, ok := values["service_appliance_ha_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceApplianceHaMode = castedValue

	}

	if value, ok := values["service_appliance_driver"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceApplianceDriver = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

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

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

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

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["annotations_key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	return m, nil
}

// ListServiceApplianceSet lists ServiceApplianceSet with list spec.
func ListServiceApplianceSet(tx *sql.Tx, spec *common.ListSpec) ([]*models.ServiceApplianceSet, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "service_appliance_set"
	spec.Fields = ServiceApplianceSetFields
	spec.RefFields = ServiceApplianceSetRefFields
	result := models.MakeServiceApplianceSetSlice()
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
		m, err := scanServiceApplianceSet(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowServiceApplianceSet shows ServiceApplianceSet resource
func ShowServiceApplianceSet(tx *sql.Tx, uuid string) (*models.ServiceApplianceSet, error) {
	list, err := ListServiceApplianceSet(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateServiceApplianceSet updates a resource
func UpdateServiceApplianceSet(tx *sql.Tx, uuid string, model *models.ServiceApplianceSet) error {
	//TODO(nati) support update
	return nil
}

// DeleteServiceApplianceSet deletes a resource
func DeleteServiceApplianceSet(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceApplianceSetQuery)
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
