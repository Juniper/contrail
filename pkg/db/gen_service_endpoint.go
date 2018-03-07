// nolint
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

// ServiceEndpointFields is db columns for ServiceEndpoint
var ServiceEndpointFields = []string{
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
	"configuration_version",
	"key_value_pair",
}

// ServiceEndpointRefFields is db reference fields for ServiceEndpoint
var ServiceEndpointRefFields = map[string][]string{

	"service_connection_module": []string{
	// <schema.Schema Value>

	},

	"physical_router": []string{
	// <schema.Schema Value>

	},

	"service_object": []string{
	// <schema.Schema Value>

	},
}

// ServiceEndpointBackRefFields is db back reference fields for ServiceEndpoint
var ServiceEndpointBackRefFields = map[string][]string{}

// ServiceEndpointParentTypes is possible parents for ServiceEndpoint
var ServiceEndpointParents = []string{}

// CreateServiceEndpoint inserts ServiceEndpoint to DB
// nolint
func (db *DB) createServiceEndpoint(
	ctx context.Context,
	request *models.CreateServiceEndpointRequest) error {
	qb := db.queryBuilders["service_endpoint"]
	tx := GetTransaction(ctx)
	model := request.ServiceEndpoint
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
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
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		log.WithFields(log.Fields{
			"model": model,
			"err":   err}).Debug("create failed")
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.ServiceConnectionModuleRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_connection_module"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceConnectionModuleRefs create failed")
		}
	}

	for _, ref := range model.PhysicalRouterRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("physical_router"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs create failed")
		}
	}

	for _, ref := range model.ServiceObjectRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_object"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceObjectRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "service_endpoint",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "service_endpoint", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServiceEndpoint(values map[string]interface{}) (*models.ServiceEndpoint, error) {
	m := models.MakeServiceEndpoint()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

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

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

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
			referenceModel := &models.ServiceEndpointServiceObjectRef{}
			referenceModel.UUID = uuid
			m.ServiceObjectRefs = append(m.ServiceObjectRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_connection_module"]; ok {
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
			referenceModel := &models.ServiceEndpointServiceConnectionModuleRef{}
			referenceModel.UUID = uuid
			m.ServiceConnectionModuleRefs = append(m.ServiceConnectionModuleRefs, referenceModel)

		}
	}

	if value, ok := values["ref_physical_router"]; ok {
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
			referenceModel := &models.ServiceEndpointPhysicalRouterRef{}
			referenceModel.UUID = uuid
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	return m, nil
}

// ListServiceEndpoint lists ServiceEndpoint with list spec.
func (db *DB) listServiceEndpoint(ctx context.Context, request *models.ListServiceEndpointRequest) (response *models.ListServiceEndpointResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["service_endpoint"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.ServiceEndpoint{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
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
		m, err := scanServiceEndpoint(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListServiceEndpointResponse{
		ServiceEndpoints: result,
	}
	return response, nil
}

// UpdateServiceEndpoint updates a resource
func (db *DB) updateServiceEndpoint(
	ctx context.Context,
	request *models.UpdateServiceEndpointRequest,
) error {
	//TODO
	return nil
}

// DeleteServiceEndpoint deletes a resource
func (db *DB) deleteServiceEndpoint(
	ctx context.Context,
	request *models.DeleteServiceEndpointRequest) error {
	qb := db.queryBuilders["service_endpoint"]

	selectQuery := qb.SelectForDeleteQuery()
	deleteQuery := qb.DeleteQuery()

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

	err = db.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateServiceEndpoint handle a Create API
// nolint
func (db *DB) CreateServiceEndpoint(
	ctx context.Context,
	request *models.CreateServiceEndpointRequest) (*models.CreateServiceEndpointResponse, error) {
	model := request.ServiceEndpoint
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createServiceEndpoint(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_endpoint",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceEndpointResponse{
		ServiceEndpoint: request.ServiceEndpoint,
	}, nil
}

//UpdateServiceEndpoint handles a Update request.
func (db *DB) UpdateServiceEndpoint(
	ctx context.Context,
	request *models.UpdateServiceEndpointRequest) (*models.UpdateServiceEndpointResponse, error) {
	model := request.ServiceEndpoint
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateServiceEndpoint(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_endpoint",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceEndpointResponse{
		ServiceEndpoint: model,
	}, nil
}

//DeleteServiceEndpoint delete a resource.
func (db *DB) DeleteServiceEndpoint(ctx context.Context, request *models.DeleteServiceEndpointRequest) (*models.DeleteServiceEndpointResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteServiceEndpoint(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceEndpointResponse{
		ID: request.ID,
	}, nil
}

//GetServiceEndpoint a Get request.
func (db *DB) GetServiceEndpoint(ctx context.Context, request *models.GetServiceEndpointRequest) (response *models.GetServiceEndpointResponse, err error) {
	spec := &models.ListSpec{
		Limit:  1,
		Detail: true,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListServiceEndpointRequest{
		Spec: spec,
	}
	var result *models.ListServiceEndpointResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listServiceEndpoint(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceEndpoints) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceEndpointResponse{
		ServiceEndpoint: result.ServiceEndpoints[0],
	}
	return response, nil
}

//ListServiceEndpoint handles a List service Request.
// nolint
func (db *DB) ListServiceEndpoint(
	ctx context.Context,
	request *models.ListServiceEndpointRequest) (response *models.ListServiceEndpointResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listServiceEndpoint(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
