package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
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
// nolint
func (db *DB) createServiceHealthCheck(
	ctx context.Context,
	request *models.CreateServiceHealthCheckRequest) error {
	tx := GetTransaction(ctx)
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

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "service_health_check",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "service_health_check", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

// nolint
func scanServiceHealthCheck(values map[string]interface{}) (*models.ServiceHealthCheck, error) {
	m := models.MakeServiceHealthCheck()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["url_path"]; ok {

		m.ServiceHealthCheckProperties.URLPath = common.InterfaceToString(value)

	}

	if value, ok := values["timeoutUsecs"]; ok {

		m.ServiceHealthCheckProperties.TimeoutUsecs = common.InterfaceToInt64(value)

	}

	if value, ok := values["timeout"]; ok {

		m.ServiceHealthCheckProperties.Timeout = common.InterfaceToInt64(value)

	}

	if value, ok := values["monitor_type"]; ok {

		m.ServiceHealthCheckProperties.MonitorType = common.InterfaceToString(value)

	}

	if value, ok := values["max_retries"]; ok {

		m.ServiceHealthCheckProperties.MaxRetries = common.InterfaceToInt64(value)

	}

	if value, ok := values["http_method"]; ok {

		m.ServiceHealthCheckProperties.HTTPMethod = common.InterfaceToString(value)

	}

	if value, ok := values["health_check_type"]; ok {

		m.ServiceHealthCheckProperties.HealthCheckType = common.InterfaceToString(value)

	}

	if value, ok := values["expected_codes"]; ok {

		m.ServiceHealthCheckProperties.ExpectedCodes = common.InterfaceToString(value)

	}

	if value, ok := values["enabled"]; ok {

		m.ServiceHealthCheckProperties.Enabled = common.InterfaceToBool(value)

	}

	if value, ok := values["delayUsecs"]; ok {

		m.ServiceHealthCheckProperties.DelayUsecs = common.InterfaceToInt64(value)

	}

	if value, ok := values["delay"]; ok {

		m.ServiceHealthCheckProperties.Delay = common.InterfaceToInt64(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = common.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = common.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = common.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = common.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = common.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = common.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

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
			uuid := common.InterfaceToString(referenceMap["to"])
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
// nolint
func (db *DB) listServiceHealthCheck(ctx context.Context, request *models.ListServiceHealthCheckRequest) (response *models.ListServiceHealthCheckResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)
	qb := &ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "service_health_check"
	qb.Fields = ServiceHealthCheckFields
	qb.RefFields = ServiceHealthCheckRefFields
	qb.BackRefFields = ServiceHealthCheckBackRefFields
	result := []*models.ServiceHealthCheck{}

	if spec.ParentFQName != nil {
		parentMetaData, err := GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
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
// nolint
func (db *DB) updateServiceHealthCheck(
	ctx context.Context,
	request *models.UpdateServiceHealthCheckRequest,
) error {
	//TODO
	return nil
}

// DeleteServiceHealthCheck deletes a resource
// nolint
func (db *DB) deleteServiceHealthCheck(
	ctx context.Context,
	request *models.DeleteServiceHealthCheckRequest) error {
	deleteQuery := deleteServiceHealthCheckQuery
	selectQuery := "select count(uuid) from service_health_check where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	tx := GetTransaction(ctx)
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

	err = DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateServiceHealthCheck handle a Create API
// nolint
func (db *DB) CreateServiceHealthCheck(
	ctx context.Context,
	request *models.CreateServiceHealthCheckRequest) (*models.CreateServiceHealthCheckResponse, error) {
	model := request.ServiceHealthCheck
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createServiceHealthCheck(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_health_check",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceHealthCheckResponse{
		ServiceHealthCheck: request.ServiceHealthCheck,
	}, nil
}

//UpdateServiceHealthCheck handles a Update request.
func (db *DB) UpdateServiceHealthCheck(
	ctx context.Context,
	request *models.UpdateServiceHealthCheckRequest) (*models.UpdateServiceHealthCheckResponse, error) {
	model := request.ServiceHealthCheck
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateServiceHealthCheck(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_health_check",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceHealthCheckResponse{
		ServiceHealthCheck: model,
	}, nil
}

//DeleteServiceHealthCheck delete a resource.
func (db *DB) DeleteServiceHealthCheck(ctx context.Context, request *models.DeleteServiceHealthCheckRequest) (*models.DeleteServiceHealthCheckResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteServiceHealthCheck(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceHealthCheckResponse{
		ID: request.ID,
	}, nil
}

//GetServiceHealthCheck a Get request.
// nolint
func (db *DB) GetServiceHealthCheck(ctx context.Context, request *models.GetServiceHealthCheckRequest) (response *models.GetServiceHealthCheckResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListServiceHealthCheckRequest{
		Spec: spec,
	}
	var result *models.ListServiceHealthCheckResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listServiceHealthCheck(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceHealthChecks) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceHealthCheckResponse{
		ServiceHealthCheck: result.ServiceHealthChecks[0],
	}
	return response, nil
}

//ListServiceHealthCheck handles a List service Request.
// nolint
func (db *DB) ListServiceHealthCheck(
	ctx context.Context,
	request *models.ListServiceHealthCheckRequest) (response *models.ListServiceHealthCheckResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listServiceHealthCheck(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
