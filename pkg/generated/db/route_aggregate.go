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

const insertRouteAggregateQuery = "insert into `route_aggregate` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteRouteAggregateQuery = "delete from `route_aggregate` where uuid = ?"

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

const insertRouteAggregateServiceInstanceQuery = "insert into `ref_route_aggregate_service_instance` (`from`, `to` ,`interface_type`) values (?, ?,?);"

// CreateRouteAggregate inserts RouteAggregate to DB
func (db *DB) createRouteAggregate(
	ctx context.Context,
	request *models.CreateRouteAggregateRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.RouteAggregate
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteAggregateQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertRouteAggregateQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
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

	stmtServiceInstanceRef, err := tx.Prepare(insertRouteAggregateServiceInstanceQuery)
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
		Type:   "route_aggregate",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "route_aggregate", model.UUID, model.GetPerms2().GetShare())
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

		m.UUID = schema.InterfaceToString(value)

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
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "route_aggregate"
	qb.Fields = RouteAggregateFields
	qb.RefFields = RouteAggregateRefFields
	qb.BackRefFields = RouteAggregateBackRefFields
	result := []*models.RouteAggregate{}

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
	deleteQuery := deleteRouteAggregateQuery
	selectQuery := "select count(uuid) from route_aggregate where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	tx := common.GetTransaction(ctx)
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

//CreateRouteAggregate handle a Create API
func (db *DB) CreateRouteAggregate(
	ctx context.Context,
	request *models.CreateRouteAggregateRequest) (*models.CreateRouteAggregateResponse, error) {
	model := request.RouteAggregate
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
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
	if err := common.DoInTransaction(
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
	if err := common.DoInTransaction(
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
	if err := common.DoInTransaction(
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
func (db *DB) ListRouteAggregate(
	ctx context.Context,
	request *models.ListRouteAggregateRequest) (response *models.ListRouteAggregateResponse, err error) {
	if err := common.DoInTransaction(
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
