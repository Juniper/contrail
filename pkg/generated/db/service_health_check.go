package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceHealthCheckQuery = "insert into `service_health_check` (`uuid`,`url_path`,`timeoutUsecs`,`timeout`,`monitor_type`,`max_retries`,`http_method`,`health_check_type`,`expected_codes`,`enabled`,`delayUsecs`,`delay`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceHealthCheckQuery = "update `service_health_check` set `uuid` = ?,`url_path` = ?,`timeoutUsecs` = ?,`timeout` = ?,`monitor_type` = ?,`max_retries` = ?,`http_method` = ?,`health_check_type` = ?,`expected_codes` = ?,`enabled` = ?,`delayUsecs` = ?,`delay` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteServiceHealthCheckQuery = "delete from `service_health_check` where uuid = ?"

// ServiceHealthCheckFields is db columns for ServiceHealthCheck
var ServiceHealthCheckFields = []string{
	"uuid",
	"url_path",
	"timeoutUsecs",
	"timeout",
	"monitor_type",
	"max_retries",
	"http_method",
	"health_check_type",
	"expected_codes",
	"enabled",
	"delayUsecs",
	"delay",
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
}

// ServiceHealthCheckRefFields is db reference fields for ServiceHealthCheck
var ServiceHealthCheckRefFields = map[string][]string{

	"service_instance": {
		// <common.Schema Value>
		"interface_type",
	},
}

// ServiceHealthCheckBackRefFields is db back reference fields for ServiceHealthCheck
var ServiceHealthCheckBackRefFields = map[string][]string{}

const insertServiceHealthCheckServiceInstanceQuery = "insert into `ref_service_health_check_service_instance` (`from`, `to` ,`interface_type`) values (?, ?,?);"

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
	_, err = stmt.Exec(string(model.UUID),
		string(model.ServiceHealthCheckProperties.URLPath),
		int(model.ServiceHealthCheckProperties.TimeoutUsecs),
		int(model.ServiceHealthCheckProperties.Timeout),
		string(model.ServiceHealthCheckProperties.MonitorType),
		int(model.ServiceHealthCheckProperties.MaxRetries),
		string(model.ServiceHealthCheckProperties.HTTPMethod),
		string(model.ServiceHealthCheckProperties.HealthCheckType),
		string(model.ServiceHealthCheckProperties.ExpectedCodes),
		bool(model.ServiceHealthCheckProperties.Enabled),
		int(model.ServiceHealthCheckProperties.DelayUsecs),
		int(model.ServiceHealthCheckProperties.Delay),
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
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceInstanceRef, err := tx.Prepare(insertServiceHealthCheckServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceInstanceRefs create statement failed")
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeServiceInterfaceTag()
		}

		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.InterfaceType))
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

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["url_path"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceHealthCheckProperties.URLPath = castedValue

	}

	if value, ok := values["timeoutUsecs"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.TimeoutUsecs = castedValue

	}

	if value, ok := values["timeout"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.Timeout = castedValue

	}

	if value, ok := values["monitor_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceHealthCheckProperties.MonitorType = models.HealthCheckProtocolType(castedValue)

	}

	if value, ok := values["max_retries"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.MaxRetries = castedValue

	}

	if value, ok := values["http_method"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceHealthCheckProperties.HTTPMethod = castedValue

	}

	if value, ok := values["health_check_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceHealthCheckProperties.HealthCheckType = models.HealthCheckType(castedValue)

	}

	if value, ok := values["expected_codes"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceHealthCheckProperties.ExpectedCodes = castedValue

	}

	if value, ok := values["enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceHealthCheckProperties.Enabled = castedValue

	}

	if value, ok := values["delayUsecs"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.DelayUsecs = castedValue

	}

	if value, ok := values["delay"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ServiceHealthCheckProperties.Delay = castedValue

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

	if value, ok := values["ref_service_instance"]; ok {
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
			referenceModel := &models.ServiceHealthCheckServiceInstanceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

			attr := models.MakeServiceInterfaceTag()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListServiceHealthCheck lists ServiceHealthCheck with list spec.
func ListServiceHealthCheck(tx *sql.Tx, spec *common.ListSpec) ([]*models.ServiceHealthCheck, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "service_health_check"
	if spec.Fields == nil {
		spec.Fields = ServiceHealthCheckFields
	}
	spec.RefFields = ServiceHealthCheckRefFields
	spec.BackRefFields = ServiceHealthCheckBackRefFields
	result := models.MakeServiceHealthCheckSlice()
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
		m, err := scanServiceHealthCheck(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateServiceHealthCheck updates a resource
func UpdateServiceHealthCheck(tx *sql.Tx, uuid string, model *models.ServiceHealthCheck) error {
	//TODO(nati) support update
	return nil
}

// DeleteServiceHealthCheck deletes a resource
func DeleteServiceHealthCheck(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	query := deleteServiceHealthCheckQuery
	var err error

	if auth.IsAdmin() {
		_, err = tx.Exec(query, uuid)
	} else {
		query += " and owner = ?"
		_, err = tx.Exec(query, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
