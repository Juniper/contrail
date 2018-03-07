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

// RouteAggregateFields is db columns for RouteAggregate
var RouteAggregateFields = []string{
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

// RouteAggregateRefFields is db reference fields for RouteAggregate
var RouteAggregateRefFields = map[string][]string{

	"service_instance": []string{
		// <schema.Schema Value>
		"interface_type",
	},
}

// RouteAggregateBackRefFields is db back reference fields for RouteAggregate
var RouteAggregateBackRefFields = map[string][]string{}

// RouteAggregateParentTypes is possible parents for RouteAggregate
var RouteAggregateParents = []string{

	"project",
}

// CreateRouteAggregate inserts RouteAggregate to DB
// nolint
func (db *DB) createRouteAggregate(
	ctx context.Context,
	request *models.CreateRouteAggregateRequest) error {
	qb := db.queryBuilders["route_aggregate"]
	tx := GetTransaction(ctx)
	model := request.RouteAggregate
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

	for _, ref := range model.ServiceInstanceRefs {

		if ref.Attr == nil {
			ref.Attr = &models.ServiceInterfaceTag{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_instance"), model.UUID, ref.UUID, string(ref.Attr.GetInterfaceType()))
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "route_aggregate",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "route_aggregate", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanRouteAggregate(values map[string]interface{}) (*models.RouteAggregate, error) {
	m := models.MakeRouteAggregate()

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
			referenceModel := &models.RouteAggregateServiceInstanceRef{}
			referenceModel.UUID = uuid
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

			attr := models.MakeServiceInterfaceTag()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListRouteAggregate lists RouteAggregate with list spec.
func (db *DB) listRouteAggregate(ctx context.Context, request *models.ListRouteAggregateRequest) (response *models.ListRouteAggregateResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["route_aggregate"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.RouteAggregate{}

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
		m, err := scanRouteAggregate(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListRouteAggregateResponse{
		RouteAggregates: result,
	}
	return response, nil
}

// UpdateRouteAggregate updates a resource
func (db *DB) updateRouteAggregate(
	ctx context.Context,
	request *models.UpdateRouteAggregateRequest,
) error {
	//TODO
	return nil
}

// DeleteRouteAggregate deletes a resource
func (db *DB) deleteRouteAggregate(
	ctx context.Context,
	request *models.DeleteRouteAggregateRequest) error {
	qb := db.queryBuilders["route_aggregate"]

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

//CreateRouteAggregate handle a Create API
// nolint
func (db *DB) CreateRouteAggregate(
	ctx context.Context,
	request *models.CreateRouteAggregateRequest) (*models.CreateRouteAggregateResponse, error) {
	model := request.RouteAggregate
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createRouteAggregate(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_aggregate",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateRouteAggregateResponse{
		RouteAggregate: request.RouteAggregate,
	}, nil
}

//UpdateRouteAggregate handles a Update request.
func (db *DB) UpdateRouteAggregate(
	ctx context.Context,
	request *models.UpdateRouteAggregateRequest) (*models.UpdateRouteAggregateResponse, error) {
	model := request.RouteAggregate
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateRouteAggregate(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_aggregate",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateRouteAggregateResponse{
		RouteAggregate: model,
	}, nil
}

//DeleteRouteAggregate delete a resource.
func (db *DB) DeleteRouteAggregate(ctx context.Context, request *models.DeleteRouteAggregateRequest) (*models.DeleteRouteAggregateResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteRouteAggregate(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteRouteAggregateResponse{
		ID: request.ID,
	}, nil
}

//GetRouteAggregate a Get request.
func (db *DB) GetRouteAggregate(ctx context.Context, request *models.GetRouteAggregateRequest) (response *models.GetRouteAggregateResponse, err error) {
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
	listRequest := &models.ListRouteAggregateRequest{
		Spec: spec,
	}
	var result *models.ListRouteAggregateResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listRouteAggregate(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.RouteAggregates) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetRouteAggregateResponse{
		RouteAggregate: result.RouteAggregates[0],
	}
	return response, nil
}

//ListRouteAggregate handles a List service Request.
// nolint
func (db *DB) ListRouteAggregate(
	ctx context.Context,
	request *models.ListRouteAggregateRequest) (response *models.ListRouteAggregateResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listRouteAggregate(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
