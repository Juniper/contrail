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

const insertServiceConnectionModuleQuery = "insert into `service_connection_module` (`uuid`,`service_type`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`e2_service`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteServiceConnectionModuleQuery = "delete from `service_connection_module` where uuid = ?"

// ServiceConnectionModuleFields is db columns for ServiceConnectionModule
var ServiceConnectionModuleFields = []string{
	"uuid",
	"service_type",
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
	"e2_service",
	"display_name",
	"key_value_pair",
}

// ServiceConnectionModuleRefFields is db reference fields for ServiceConnectionModule
var ServiceConnectionModuleRefFields = map[string][]string{

	"service_object": []string{
	// <schema.Schema Value>

	},
}

// ServiceConnectionModuleBackRefFields is db back reference fields for ServiceConnectionModule
var ServiceConnectionModuleBackRefFields = map[string][]string{}

// ServiceConnectionModuleParentTypes is possible parents for ServiceConnectionModule
var ServiceConnectionModuleParents = []string{}

const insertServiceConnectionModuleServiceObjectQuery = "insert into `ref_service_connection_module_service_object` (`from`, `to` ) values (?, ?);"

// CreateServiceConnectionModule inserts ServiceConnectionModule to DB
// nolint
func (db *DB) createServiceConnectionModule(
	ctx context.Context,
	request *models.CreateServiceConnectionModuleRequest) error {
	tx := GetTransaction(ctx)
	model := request.ServiceConnectionModule
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceConnectionModuleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceConnectionModuleQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		string(model.GetServiceType()),
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
		string(model.GetE2Service()),
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceObjectRef, err := tx.Prepare(insertServiceConnectionModuleServiceObjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceObjectRefs create statement failed")
	}
	defer stmtServiceObjectRef.Close()
	for _, ref := range model.ServiceObjectRefs {

		_, err = stmtServiceObjectRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceObjectRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "service_connection_module",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "service_connection_module", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

// nolint
func scanServiceConnectionModule(values map[string]interface{}) (*models.ServiceConnectionModule, error) {
	m := models.MakeServiceConnectionModule()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["service_type"]; ok {

		m.ServiceType = common.InterfaceToString(value)

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

	if value, ok := values["e2_service"]; ok {

		m.E2Service = common.InterfaceToString(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_service_object"]; ok {
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
			referenceModel := &models.ServiceConnectionModuleServiceObjectRef{}
			referenceModel.UUID = uuid
			m.ServiceObjectRefs = append(m.ServiceObjectRefs, referenceModel)

		}
	}

	return m, nil
}

// ListServiceConnectionModule lists ServiceConnectionModule with list spec.
// nolint
func (db *DB) listServiceConnectionModule(ctx context.Context, request *models.ListServiceConnectionModuleRequest) (response *models.ListServiceConnectionModuleResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)
	qb := &ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "service_connection_module"
	qb.Fields = ServiceConnectionModuleFields
	qb.RefFields = ServiceConnectionModuleRefFields
	qb.BackRefFields = ServiceConnectionModuleBackRefFields
	result := []*models.ServiceConnectionModule{}

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
		m, err := scanServiceConnectionModule(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListServiceConnectionModuleResponse{
		ServiceConnectionModules: result,
	}
	return response, nil
}

// UpdateServiceConnectionModule updates a resource
// nolint
func (db *DB) updateServiceConnectionModule(
	ctx context.Context,
	request *models.UpdateServiceConnectionModuleRequest,
) error {
	//TODO
	return nil
}

// DeleteServiceConnectionModule deletes a resource
// nolint
func (db *DB) deleteServiceConnectionModule(
	ctx context.Context,
	request *models.DeleteServiceConnectionModuleRequest) error {
	deleteQuery := deleteServiceConnectionModuleQuery
	selectQuery := "select count(uuid) from service_connection_module where uuid = ?"
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

//CreateServiceConnectionModule handle a Create API
// nolint
func (db *DB) CreateServiceConnectionModule(
	ctx context.Context,
	request *models.CreateServiceConnectionModuleRequest) (*models.CreateServiceConnectionModuleResponse, error) {
	model := request.ServiceConnectionModule
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createServiceConnectionModule(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceConnectionModuleResponse{
		ServiceConnectionModule: request.ServiceConnectionModule,
	}, nil
}

//UpdateServiceConnectionModule handles a Update request.
func (db *DB) UpdateServiceConnectionModule(
	ctx context.Context,
	request *models.UpdateServiceConnectionModuleRequest) (*models.UpdateServiceConnectionModuleResponse, error) {
	model := request.ServiceConnectionModule
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateServiceConnectionModule(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceConnectionModuleResponse{
		ServiceConnectionModule: model,
	}, nil
}

//DeleteServiceConnectionModule delete a resource.
func (db *DB) DeleteServiceConnectionModule(ctx context.Context, request *models.DeleteServiceConnectionModuleRequest) (*models.DeleteServiceConnectionModuleResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteServiceConnectionModule(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceConnectionModuleResponse{
		ID: request.ID,
	}, nil
}

//GetServiceConnectionModule a Get request.
// nolint
func (db *DB) GetServiceConnectionModule(ctx context.Context, request *models.GetServiceConnectionModuleRequest) (response *models.GetServiceConnectionModuleResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListServiceConnectionModuleRequest{
		Spec: spec,
	}
	var result *models.ListServiceConnectionModuleResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listServiceConnectionModule(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceConnectionModules) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceConnectionModuleResponse{
		ServiceConnectionModule: result.ServiceConnectionModules[0],
	}
	return response, nil
}

//ListServiceConnectionModule handles a List service Request.
// nolint
func (db *DB) ListServiceConnectionModule(
	ctx context.Context,
	request *models.ListServiceConnectionModuleRequest) (response *models.ListServiceConnectionModuleResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listServiceConnectionModule(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
