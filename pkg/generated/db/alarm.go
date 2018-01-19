package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertAlarmQuery = "insert into `alarm` (`uve_key`,`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`,`alarm_severity`,`or_list`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteAlarmQuery = "delete from `alarm` where uuid = ?"

// AlarmFields is db columns for Alarm
var AlarmFields = []string{
	"uve_key",
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"user_visible",
	"permissions_owner_access",
	"permissions_owner",
	"other_access",
	"group_access",
	"group",
	"last_modified",
	"enable",
	"description",
	"creator",
	"created",
	"fq_name",
	"display_name",
	"key_value_pair",
	"alarm_severity",
	"or_list",
}

// AlarmRefFields is db reference fields for Alarm
var AlarmRefFields = map[string][]string{}

// AlarmBackRefFields is db back reference fields for Alarm
var AlarmBackRefFields = map[string][]string{}

// AlarmParentTypes is possible parents for Alarm
var AlarmParents = []string{

	"global_system_config",

	"project",
}

// CreateAlarm inserts Alarm to DB
func CreateAlarm(tx *sql.Tx, model *models.Alarm) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAlarmQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertAlarmQuery,
	}).Debug("create query")
	_, err = stmt.Exec(common.MustJSON(model.UveKeys.UveKey),
		string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		common.MustJSON(model.FQName),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		int(model.AlarmSeverity),
		common.MustJSON(model.AlarmRules.OrList))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "alarm",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "alarm", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanAlarm(values map[string]interface{}) (*models.Alarm, error) {
	m := models.MakeAlarm()

	if value, ok := values["uve_key"]; ok {

		json.Unmarshal(value.([]byte), &m.UveKeys.UveKey)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["alarm_severity"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.AlarmSeverity = models.AlarmSeverity(castedValue)

	}

	if value, ok := values["or_list"]; ok {

		json.Unmarshal(value.([]byte), &m.AlarmRules.OrList)

	}

	return m, nil
}

// ListAlarm lists Alarm with list spec.
func ListAlarm(tx *sql.Tx, spec *common.ListSpec) ([]*models.Alarm, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "alarm"
	spec.Fields = AlarmFields
	spec.RefFields = AlarmRefFields
	spec.BackRefFields = AlarmBackRefFields
	result := models.MakeAlarmSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
	}

	query := spec.BuildQuery()
	columns := spec.Columns
	values := spec.Values
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
		m, err := scanAlarm(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateAlarm updates a resource
func UpdateAlarm(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateAlarmQuery = "update `alarm` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UveKeys.UveKey", "."); ok {
		updateAlarmQuery += "`uve_key` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateAlarmQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateAlarmQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateAlarmQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateAlarmQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateAlarmQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateAlarmQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateAlarmQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateAlarmQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateAlarmQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateAlarmQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateAlarmQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateAlarmQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateAlarmQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateAlarmQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateAlarmQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateAlarmQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateAlarmQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateAlarmQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateAlarmQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateAlarmQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateAlarmQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".AlarmSeverity", "."); ok {
		updateAlarmQuery += "`alarm_severity` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateAlarmQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".AlarmRules.OrList", "."); ok {
		updateAlarmQuery += "`or_list` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateAlarmQuery += ","
	}

	updateAlarmQuery =
		updateAlarmQuery[:len(updateAlarmQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateAlarmQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateAlarmQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteAlarm deletes a resource
func DeleteAlarm(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteAlarmQuery
	selectQuery := "select count(uuid) from alarm where uuid = ?"
	var err error
	var count int

	if auth.IsAdmin() {
		row := tx.QueryRow(selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRow(selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = common.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}
