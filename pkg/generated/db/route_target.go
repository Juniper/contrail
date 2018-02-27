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

const insertRouteTargetQuery = "insert into `route_target` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteRouteTargetQuery = "delete from `route_target` where uuid = ?"

// RouteTargetFields is db columns for RouteTarget
var RouteTargetFields = []string{
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

// RouteTargetRefFields is db reference fields for RouteTarget
var RouteTargetRefFields = map[string][]string{}

// RouteTargetBackRefFields is db back reference fields for RouteTarget
var RouteTargetBackRefFields = map[string][]string{}

// RouteTargetParentTypes is possible parents for RouteTarget
var RouteTargetParents = []string{}

// CreateRouteTarget inserts RouteTarget to DB
func (db *DB) createRouteTarget(
	ctx context.Context,
	request *models.CreateRouteTargetRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.RouteTarget
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteTargetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertRouteTargetQuery,
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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "route_target",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "route_target", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanRouteTarget(values map[string]interface{}) (*models.RouteTarget, error) {
	m := models.MakeRouteTarget()

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

	return m, nil
}

// ListRouteTarget lists RouteTarget with list spec.
func (db *DB) listRouteTarget(ctx context.Context, request *models.ListRouteTargetRequest) (response *models.ListRouteTargetResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "route_target"
	qb.Fields = RouteTargetFields
	qb.RefFields = RouteTargetRefFields
	qb.BackRefFields = RouteTargetBackRefFields
	result := []*models.RouteTarget{}

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
		m, err := scanRouteTarget(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListRouteTargetResponse{
		RouteTargets: result,
	}
	return response, nil
}

// UpdateRouteTarget updates a resource
func (db *DB) updateRouteTarget(
	ctx context.Context,
	request *models.UpdateRouteTargetRequest,
) error {
	//TODO
	return nil
}

// DeleteRouteTarget deletes a resource
func (db *DB) deleteRouteTarget(
	ctx context.Context,
	request *models.DeleteRouteTargetRequest) error {
	deleteQuery := deleteRouteTargetQuery
	selectQuery := "select count(uuid) from route_target where uuid = ?"
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

//CreateRouteTarget handle a Create API
func (db *DB) CreateRouteTarget(
	ctx context.Context,
	request *models.CreateRouteTargetRequest) (*models.CreateRouteTargetResponse, error) {
	model := request.RouteTarget
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createRouteTarget(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_target",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateRouteTargetResponse{
		RouteTarget: request.RouteTarget,
	}, nil
}

//UpdateRouteTarget handles a Update request.
func (db *DB) UpdateRouteTarget(
	ctx context.Context,
	request *models.UpdateRouteTargetRequest) (*models.UpdateRouteTargetResponse, error) {
	model := request.RouteTarget
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateRouteTarget(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_target",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateRouteTargetResponse{
		RouteTarget: model,
	}, nil
}

//DeleteRouteTarget delete a resource.
func (db *DB) DeleteRouteTarget(ctx context.Context, request *models.DeleteRouteTargetRequest) (*models.DeleteRouteTargetResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteRouteTarget(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteRouteTargetResponse{
		ID: request.ID,
	}, nil
}

//GetRouteTarget a Get request.
func (db *DB) GetRouteTarget(ctx context.Context, request *models.GetRouteTargetRequest) (response *models.GetRouteTargetResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListRouteTargetRequest{
		Spec: spec,
	}
	var result *models.ListRouteTargetResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listRouteTarget(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.RouteTargets) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetRouteTargetResponse{
		RouteTarget: result.RouteTargets[0],
	}
	return response, nil
}

//ListRouteTarget handles a List service Request.
func (db *DB) ListRouteTarget(
	ctx context.Context,
	request *models.ListRouteTargetRequest) (response *models.ListRouteTargetResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listRouteTarget(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
