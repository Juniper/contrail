package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceHealthCheckQuery = "insert into `service_health_check` (`uuid`,`url_path`,`timeoutUsecs`,`timeout`,`monitor_type`,`max_retries`,`http_method`,`health_check_type`,`expected_codes`,`enabled`,`delayUsecs`,`delay`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
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

	"service_instance": []string{
		// <schema.Schema Value>
		"interface_type",
	},
}

// ServiceHealthCheckBackRefFields is db back reference fields for ServiceHealthCheck
var ServiceHealthCheckBackRefFields = map[string][]string{}

// ServiceHealthCheckParentTypes is possible parents for ServiceHealthCheck
var ServiceHealthCheckParents = []string{

	"project",
}

const insertServiceHealthCheckServiceInstanceQuery = "insert into `ref_service_health_check_service_instance` (`from`, `to` ,`interface_type`) values (?, ?,?);"

// CreateServiceHealthCheck inserts ServiceHealthCheck to DB
func CreateServiceHealthCheck(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateServiceHealthCheckRequest) error {
	model := request.ServiceHealthCheck
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
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		string(model.GetServiceHealthCheckProperties().GetURLPath()),
		int(model.GetServiceHealthCheckProperties().GetTimeoutUsecs()),
		int(model.GetServiceHealthCheckProperties().GetTimeout()),
		string(model.GetServiceHealthCheckProperties().GetMonitorType()),
		int(model.GetServiceHealthCheckProperties().GetMaxRetries()),
		string(model.GetServiceHealthCheckProperties().GetHTTPMethod()),
		string(model.GetServiceHealthCheckProperties().GetHealthCheckType()),
		string(model.GetServiceHealthCheckProperties().GetExpectedCodes()),
		bool(model.GetServiceHealthCheckProperties().GetEnabled()),
		int(model.GetServiceHealthCheckProperties().GetDelayUsecs()),
		int(model.GetServiceHealthCheckProperties().GetDelay()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetIDPerms().GetUserVisible()),
		int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
		string(model.GetIDPerms().GetPermissions().GetOwner()),
		int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
		int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
		string(model.GetIDPerms().GetPermissions().GetGroup()),
		string(model.GetIDPerms().GetLastModified()),
		bool(model.GetIDPerms().GetEnable()),
		string(model.GetIDPerms().GetDescription()),
		string(model.GetIDPerms().GetCreator()),
		string(model.GetIDPerms().GetCreated()),
		common.MustJSON(model.GetFQName()),
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
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
			ref.Attr = &models.ServiceInterfaceTag{}
		}

		_, err = stmtServiceInstanceRef.ExecContext(ctx, model.UUID, ref.UUID, string(ref.Attr.GetInterfaceType()))
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "service_health_check",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "service_health_check", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServiceHealthCheck(values map[string]interface{}) (*models.ServiceHealthCheck, error) {
	m := models.MakeServiceHealthCheck()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["url_path"]; ok {

		m.ServiceHealthCheckProperties.URLPath = schema.InterfaceToString(value)

	}

	if value, ok := values["timeoutUsecs"]; ok {

		m.ServiceHealthCheckProperties.TimeoutUsecs = schema.InterfaceToInt64(value)

	}

	if value, ok := values["timeout"]; ok {

		m.ServiceHealthCheckProperties.Timeout = schema.InterfaceToInt64(value)

	}

	if value, ok := values["monitor_type"]; ok {

		m.ServiceHealthCheckProperties.MonitorType = schema.InterfaceToString(value)

	}

	if value, ok := values["max_retries"]; ok {

		m.ServiceHealthCheckProperties.MaxRetries = schema.InterfaceToInt64(value)

	}

	if value, ok := values["http_method"]; ok {

		m.ServiceHealthCheckProperties.HTTPMethod = schema.InterfaceToString(value)

	}

	if value, ok := values["health_check_type"]; ok {

		m.ServiceHealthCheckProperties.HealthCheckType = schema.InterfaceToString(value)

	}

	if value, ok := values["expected_codes"]; ok {

		m.ServiceHealthCheckProperties.ExpectedCodes = schema.InterfaceToString(value)

	}

	if value, ok := values["enabled"]; ok {

		m.ServiceHealthCheckProperties.Enabled = schema.InterfaceToBool(value)

	}

	if value, ok := values["delayUsecs"]; ok {

		m.ServiceHealthCheckProperties.DelayUsecs = schema.InterfaceToInt64(value)

	}

	if value, ok := values["delay"]; ok {

		m.ServiceHealthCheckProperties.Delay = schema.InterfaceToInt64(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = schema.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = schema.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = schema.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = schema.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = schema.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = schema.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = schema.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = schema.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = schema.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_service_instance"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ServiceHealthCheckServiceInstanceRef{}
			referenceModel.UUID = uuid
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

			attr := models.MakeServiceInterfaceTag()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListServiceHealthCheck lists ServiceHealthCheck with list spec.
func ListServiceHealthCheck(ctx context.Context, tx *sql.Tx, request *models.ListServiceHealthCheckRequest) (response *models.ListServiceHealthCheckResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "service_health_check"
	qb.Fields = ServiceHealthCheckFields
	qb.RefFields = ServiceHealthCheckRefFields
	qb.BackRefFields = ServiceHealthCheckBackRefFields
	result := []*models.ServiceHealthCheck{}

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = common.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}

	query := qb.BuildQuery()
	columns := qb.Columns
	values := qb.Values
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.QueryContext(ctx, query, values...)
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
	response = &models.ListServiceHealthCheckResponse{
		ServiceHealthChecks: result,
	}
	return response, nil
}

// UpdateServiceHealthCheck updates a resource
func UpdateServiceHealthCheck(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateServiceHealthCheckRequest,
) error {
	//TODO
	return nil
}

// DeleteServiceHealthCheck deletes a resource
func DeleteServiceHealthCheck(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteServiceHealthCheckRequest) error {
	deleteQuery := deleteServiceHealthCheckQuery
	selectQuery := "select count(uuid) from service_health_check where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	auth := common.GetAuthCTX(ctx)
	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
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
