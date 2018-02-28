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

const insertAnalyticsNodeQuery = "insert into `analytics_node` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`,`analytics_node_ip_address`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteAnalyticsNodeQuery = "delete from `analytics_node` where uuid = ?"

// AnalyticsNodeFields is db columns for AnalyticsNode
var AnalyticsNodeFields = []string{
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
	"analytics_node_ip_address",
}

// AnalyticsNodeRefFields is db reference fields for AnalyticsNode
var AnalyticsNodeRefFields = map[string][]string{}

// AnalyticsNodeBackRefFields is db back reference fields for AnalyticsNode
var AnalyticsNodeBackRefFields = map[string][]string{}

// AnalyticsNodeParentTypes is possible parents for AnalyticsNode
var AnalyticsNodeParents = []string{

	"global_system_config",
}

// CreateAnalyticsNode inserts AnalyticsNode to DB
// nolint
func (db *DB) createAnalyticsNode(
	ctx context.Context,
	request *models.CreateAnalyticsNodeRequest) error {
	tx := GetTransaction(ctx)
	model := request.AnalyticsNode
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAnalyticsNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertAnalyticsNodeQuery,
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
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()),
		string(model.GetAnalyticsNodeIPAddress()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "analytics_node",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "analytics_node", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

// nolint
func scanAnalyticsNode(values map[string]interface{}) (*models.AnalyticsNode, error) {
	m := models.MakeAnalyticsNode()

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

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["analytics_node_ip_address"]; ok {

		m.AnalyticsNodeIPAddress = common.InterfaceToString(value)

	}

	return m, nil
}

// ListAnalyticsNode lists AnalyticsNode with list spec.
// nolint
func (db *DB) listAnalyticsNode(ctx context.Context, request *models.ListAnalyticsNodeRequest) (response *models.ListAnalyticsNodeResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)
	qb := &ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "analytics_node"
	qb.Fields = AnalyticsNodeFields
	qb.RefFields = AnalyticsNodeRefFields
	qb.BackRefFields = AnalyticsNodeBackRefFields
	result := []*models.AnalyticsNode{}

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
		m, err := scanAnalyticsNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListAnalyticsNodeResponse{
		AnalyticsNodes: result,
	}
	return response, nil
}

// UpdateAnalyticsNode updates a resource
// nolint
func (db *DB) updateAnalyticsNode(
	ctx context.Context,
	request *models.UpdateAnalyticsNodeRequest,
) error {
	//TODO
	return nil
}

// DeleteAnalyticsNode deletes a resource
// nolint
func (db *DB) deleteAnalyticsNode(
	ctx context.Context,
	request *models.DeleteAnalyticsNodeRequest) error {
	deleteQuery := deleteAnalyticsNodeQuery
	selectQuery := "select count(uuid) from analytics_node where uuid = ?"
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

//CreateAnalyticsNode handle a Create API
// nolint
func (db *DB) CreateAnalyticsNode(
	ctx context.Context,
	request *models.CreateAnalyticsNodeRequest) (*models.CreateAnalyticsNodeResponse, error) {
	model := request.AnalyticsNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createAnalyticsNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "analytics_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAnalyticsNodeResponse{
		AnalyticsNode: request.AnalyticsNode,
	}, nil
}

//UpdateAnalyticsNode handles a Update request.
func (db *DB) UpdateAnalyticsNode(
	ctx context.Context,
	request *models.UpdateAnalyticsNodeRequest) (*models.UpdateAnalyticsNodeResponse, error) {
	model := request.AnalyticsNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateAnalyticsNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "analytics_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAnalyticsNodeResponse{
		AnalyticsNode: model,
	}, nil
}

//DeleteAnalyticsNode delete a resource.
func (db *DB) DeleteAnalyticsNode(ctx context.Context, request *models.DeleteAnalyticsNodeRequest) (*models.DeleteAnalyticsNodeResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteAnalyticsNode(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAnalyticsNodeResponse{
		ID: request.ID,
	}, nil
}

//GetAnalyticsNode a Get request.
// nolint
func (db *DB) GetAnalyticsNode(ctx context.Context, request *models.GetAnalyticsNodeRequest) (response *models.GetAnalyticsNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListAnalyticsNodeRequest{
		Spec: spec,
	}
	var result *models.ListAnalyticsNodeResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listAnalyticsNode(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.AnalyticsNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAnalyticsNodeResponse{
		AnalyticsNode: result.AnalyticsNodes[0],
	}
	return response, nil
}

//ListAnalyticsNode handles a List service Request.
// nolint
func (db *DB) ListAnalyticsNode(
	ctx context.Context,
	request *models.ListAnalyticsNodeRequest) (response *models.ListAnalyticsNodeResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listAnalyticsNode(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
