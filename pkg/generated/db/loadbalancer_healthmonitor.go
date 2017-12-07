package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLoadbalancerHealthmonitorQuery = "insert into `loadbalancer_healthmonitor` (`http_method`,`admin_state`,`timeout`,`url_path`,`monitor_type`,`delay`,`expected_codes`,`max_retries`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerHealthmonitorQuery = "update `loadbalancer_healthmonitor` set `http_method` = ?,`admin_state` = ?,`timeout` = ?,`url_path` = ?,`monitor_type` = ?,`delay` = ?,`expected_codes` = ?,`max_retries` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteLoadbalancerHealthmonitorQuery = "delete from `loadbalancer_healthmonitor` where uuid = ?"

// LoadbalancerHealthmonitorFields is db columns for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorFields = []string{
	"http_method",
	"admin_state",
	"timeout",
	"url_path",
	"monitor_type",
	"delay",
	"expected_codes",
	"max_retries",
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
	"key_value_pair",
	"global_access",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"uuid",
	"fq_name",
}

// LoadbalancerHealthmonitorRefFields is db reference fields for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorRefFields = map[string][]string{}

// CreateLoadbalancerHealthmonitor inserts LoadbalancerHealthmonitor to DB
func CreateLoadbalancerHealthmonitor(tx *sql.Tx, model *models.LoadbalancerHealthmonitor) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerHealthmonitorQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLoadbalancerHealthmonitorQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.LoadbalancerHealthmonitorProperties.HTTPMethod),
		bool(model.LoadbalancerHealthmonitorProperties.AdminState),
		int(model.LoadbalancerHealthmonitorProperties.Timeout),
		string(model.LoadbalancerHealthmonitorProperties.URLPath),
		string(model.LoadbalancerHealthmonitorProperties.MonitorType),
		int(model.LoadbalancerHealthmonitorProperties.Delay),
		string(model.LoadbalancerHealthmonitorProperties.ExpectedCodes),
		int(model.LoadbalancerHealthmonitorProperties.MaxRetries),
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
		common.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		common.MustJSON(model.FQName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanLoadbalancerHealthmonitor(values map[string]interface{}) (*models.LoadbalancerHealthmonitor, error) {
	m := models.MakeLoadbalancerHealthmonitor()

	if value, ok := values["http_method"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerHealthmonitorProperties.HTTPMethod = castedValue

	}

	if value, ok := values["admin_state"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.LoadbalancerHealthmonitorProperties.AdminState = castedValue

	}

	if value, ok := values["timeout"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.LoadbalancerHealthmonitorProperties.Timeout = castedValue

	}

	if value, ok := values["url_path"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerHealthmonitorProperties.URLPath = castedValue

	}

	if value, ok := values["monitor_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerHealthmonitorProperties.MonitorType = models.HealthmonitorType(castedValue)

	}

	if value, ok := values["delay"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.LoadbalancerHealthmonitorProperties.Delay = castedValue

	}

	if value, ok := values["expected_codes"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerHealthmonitorProperties.ExpectedCodes = castedValue

	}

	if value, ok := values["max_retries"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.LoadbalancerHealthmonitorProperties.MaxRetries = castedValue

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

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

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

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	return m, nil
}

// ListLoadbalancerHealthmonitor lists LoadbalancerHealthmonitor with list spec.
func ListLoadbalancerHealthmonitor(tx *sql.Tx, spec *common.ListSpec) ([]*models.LoadbalancerHealthmonitor, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "loadbalancer_healthmonitor"
	spec.Fields = LoadbalancerHealthmonitorFields
	spec.RefFields = LoadbalancerHealthmonitorRefFields
	result := models.MakeLoadbalancerHealthmonitorSlice()
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
		m, err := scanLoadbalancerHealthmonitor(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowLoadbalancerHealthmonitor shows LoadbalancerHealthmonitor resource
func ShowLoadbalancerHealthmonitor(tx *sql.Tx, uuid string) (*models.LoadbalancerHealthmonitor, error) {
	list, err := ListLoadbalancerHealthmonitor(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateLoadbalancerHealthmonitor updates a resource
func UpdateLoadbalancerHealthmonitor(tx *sql.Tx, uuid string, model *models.LoadbalancerHealthmonitor) error {
	//TODO(nati) support update
	return nil
}

// DeleteLoadbalancerHealthmonitor deletes a resource
func DeleteLoadbalancerHealthmonitor(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerHealthmonitorQuery)
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
