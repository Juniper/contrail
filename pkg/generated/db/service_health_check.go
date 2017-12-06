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

const insertServiceHealthCheckQuery = "insert into `service_health_check` (`delayUsecs`,`enabled`,`max_retries`,`timeout`,`url_path`,`timeoutUsecs`,`delay`,`health_check_type`,`http_method`,`expected_codes`,`monitor_type`,`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceHealthCheckQuery = "update `service_health_check` set `delayUsecs` = ?,`enabled` = ?,`max_retries` = ?,`timeout` = ?,`url_path` = ?,`timeoutUsecs` = ?,`delay` = ?,`health_check_type` = ?,`http_method` = ?,`expected_codes` = ?,`monitor_type` = ?,`display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?;"
const deleteServiceHealthCheckQuery = "delete from `service_health_check` where uuid = ?"

// ServiceHealthCheckFields is db columns for ServiceHealthCheck
var ServiceHealthCheckFields = []string{
	"delayUsecs",
	"enabled",
	"max_retries",
	"timeout",
	"url_path",
	"timeoutUsecs",
	"delay",
	"health_check_type",
	"http_method",
	"expected_codes",
	"monitor_type",
	"display_name",
	"key_value_pair",
	"owner",
	"owner_access",
	"global_access",
	"share",
	"uuid",
	"fq_name",
	"last_modified",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
}

// ServiceHealthCheckRefFields is db reference fields for ServiceHealthCheck
var ServiceHealthCheckRefFields = map[string][]string{

	"service_instance": {
	// <utils.Schema Value>

	},
}

const insertServiceHealthCheckServiceInstanceQuery = "insert into `ref_service_health_check_service_instance` (`from`, `to` ) values (?, ?);"

// CreateServiceHealthCheck inserts ServiceHealthCheck to DB
func CreateServiceHealthCheck(tx *sql.Tx, model *models.ServiceHealthCheck) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceHealthCheckQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceHealthCheckQuery,
	}).Debug("create query")
	_, err = stmt.Exec(int(model.ServiceHealthCheckProperties.DelayUsecs),
		bool(model.ServiceHealthCheckProperties.Enabled),
		int(model.ServiceHealthCheckProperties.MaxRetries),
		int(model.ServiceHealthCheckProperties.Timeout),
		string(model.ServiceHealthCheckProperties.URLPath),
		int(model.ServiceHealthCheckProperties.TimeoutUsecs),
		int(model.ServiceHealthCheckProperties.Delay),
		string(model.ServiceHealthCheckProperties.HealthCheckType),
		string(model.ServiceHealthCheckProperties.HTTPMethod),
		string(model.ServiceHealthCheckProperties.ExpectedCodes),
		string(model.ServiceHealthCheckProperties.MonitorType),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceInstanceRef, err := tx.Prepare(insertServiceHealthCheckServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceInstanceRefs create statement failed")
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanServiceHealthCheck(values map[string]interface{}) (*models.ServiceHealthCheck, error) {
	m := models.MakeServiceHealthCheck()

	if value, ok := values["delayUsecs"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.DelayUsecs = castedValue

	}

	if value, ok := values["enabled"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.ServiceHealthCheckProperties.Enabled = castedValue

	}

	if value, ok := values["max_retries"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.MaxRetries = castedValue

	}

	if value, ok := values["timeout"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.Timeout = castedValue

	}

	if value, ok := values["url_path"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceHealthCheckProperties.URLPath = castedValue

	}

	if value, ok := values["timeoutUsecs"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.TimeoutUsecs = castedValue

	}

	if value, ok := values["delay"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.Delay = castedValue

	}

	if value, ok := values["health_check_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceHealthCheckProperties.HealthCheckType = models.HealthCheckType(castedValue)

	}

	if value, ok := values["http_method"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceHealthCheckProperties.HTTPMethod = castedValue

	}

	if value, ok := values["expected_codes"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceHealthCheckProperties.ExpectedCodes = castedValue

	}

	if value, ok := values["monitor_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceHealthCheckProperties.MonitorType = models.HealthCheckProtocolType(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

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

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

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

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

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

	if value, ok := values["ref_service_instance"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.ServiceHealthCheckServiceInstanceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

			attr := models.MakeServiceInterfaceTag()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListServiceHealthCheck lists ServiceHealthCheck with list spec.
func ListServiceHealthCheck(tx *sql.Tx, spec *db.ListSpec) ([]*models.ServiceHealthCheck, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "service_health_check"
	spec.Fields = ServiceHealthCheckFields
	spec.RefFields = ServiceHealthCheckRefFields
	result := models.MakeServiceHealthCheckSlice()
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
		m, err := scanServiceHealthCheck(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowServiceHealthCheck shows ServiceHealthCheck resource
func ShowServiceHealthCheck(tx *sql.Tx, uuid string) (*models.ServiceHealthCheck, error) {
	list, err := ListServiceHealthCheck(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateServiceHealthCheck updates a resource
func UpdateServiceHealthCheck(tx *sql.Tx, uuid string, model *models.ServiceHealthCheck) error {
	//TODO(nati) support update
	return nil
}

// DeleteServiceHealthCheck deletes a resource
func DeleteServiceHealthCheck(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceHealthCheckQuery)
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
