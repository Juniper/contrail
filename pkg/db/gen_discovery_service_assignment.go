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

const insertDiscoveryServiceAssignmentQuery = "insert into `discovery_service_assignment` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteDiscoveryServiceAssignmentQuery = "delete from `discovery_service_assignment` where uuid = ?"

// DiscoveryServiceAssignmentFields is db columns for DiscoveryServiceAssignment
var DiscoveryServiceAssignmentFields = []string{
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

// DiscoveryServiceAssignmentRefFields is db reference fields for DiscoveryServiceAssignment
var DiscoveryServiceAssignmentRefFields = map[string][]string{}

// DiscoveryServiceAssignmentBackRefFields is db back reference fields for DiscoveryServiceAssignment
var DiscoveryServiceAssignmentBackRefFields = map[string][]string{

	"dsa_rule": []string{
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
		"subscriber",
		"ep_version",
		"ep_type",
		"ip_prefix_len",
		"ip_prefix",
		"ep_id",
		"display_name",
		"key_value_pair",
	},
}

// DiscoveryServiceAssignmentParentTypes is possible parents for DiscoveryServiceAssignment
var DiscoveryServiceAssignmentParents = []string{}

// CreateDiscoveryServiceAssignment inserts DiscoveryServiceAssignment to DB
// nolint
func (db *DB) createDiscoveryServiceAssignment(
	ctx context.Context,
	request *models.CreateDiscoveryServiceAssignmentRequest) error {
	tx := GetTransaction(ctx)
	model := request.DiscoveryServiceAssignment
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDiscoveryServiceAssignmentQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertDiscoveryServiceAssignmentQuery,
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

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "discovery_service_assignment",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "discovery_service_assignment", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

// nolint
func scanDiscoveryServiceAssignment(values map[string]interface{}) (*models.DiscoveryServiceAssignment, error) {
	m := models.MakeDiscoveryServiceAssignment()

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

	if value, ok := values["backref_dsa_rule"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeDsaRule()
			m.DsaRules = append(m.DsaRules, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["subscriber"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.DsaRuleEntry.Subscriber)

			}

			if propertyValue, ok := childResourceMap["ep_version"]; ok && propertyValue != nil {

				childModel.DsaRuleEntry.Publisher.EpVersion = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ep_type"]; ok && propertyValue != nil {

				childModel.DsaRuleEntry.Publisher.EpType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ip_prefix_len"]; ok && propertyValue != nil {

				childModel.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ip_prefix"]; ok && propertyValue != nil {

				childModel.DsaRuleEntry.Publisher.EpPrefix.IPPrefix = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ep_id"]; ok && propertyValue != nil {

				childModel.DsaRuleEntry.Publisher.EpID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListDiscoveryServiceAssignment lists DiscoveryServiceAssignment with list spec.
// nolint
func (db *DB) listDiscoveryServiceAssignment(ctx context.Context, request *models.ListDiscoveryServiceAssignmentRequest) (response *models.ListDiscoveryServiceAssignmentResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)
	qb := &ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "discovery_service_assignment"
	qb.Fields = DiscoveryServiceAssignmentFields
	qb.RefFields = DiscoveryServiceAssignmentRefFields
	qb.BackRefFields = DiscoveryServiceAssignmentBackRefFields
	result := []*models.DiscoveryServiceAssignment{}

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
		m, err := scanDiscoveryServiceAssignment(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListDiscoveryServiceAssignmentResponse{
		DiscoveryServiceAssignments: result,
	}
	return response, nil
}

// UpdateDiscoveryServiceAssignment updates a resource
// nolint
func (db *DB) updateDiscoveryServiceAssignment(
	ctx context.Context,
	request *models.UpdateDiscoveryServiceAssignmentRequest,
) error {
	//TODO
	return nil
}

// DeleteDiscoveryServiceAssignment deletes a resource
// nolint
func (db *DB) deleteDiscoveryServiceAssignment(
	ctx context.Context,
	request *models.DeleteDiscoveryServiceAssignmentRequest) error {
	deleteQuery := deleteDiscoveryServiceAssignmentQuery
	selectQuery := "select count(uuid) from discovery_service_assignment where uuid = ?"
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

//CreateDiscoveryServiceAssignment handle a Create API
// nolint
func (db *DB) CreateDiscoveryServiceAssignment(
	ctx context.Context,
	request *models.CreateDiscoveryServiceAssignmentRequest) (*models.CreateDiscoveryServiceAssignmentResponse, error) {
	model := request.DiscoveryServiceAssignment
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createDiscoveryServiceAssignment(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "discovery_service_assignment",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateDiscoveryServiceAssignmentResponse{
		DiscoveryServiceAssignment: request.DiscoveryServiceAssignment,
	}, nil
}

//UpdateDiscoveryServiceAssignment handles a Update request.
func (db *DB) UpdateDiscoveryServiceAssignment(
	ctx context.Context,
	request *models.UpdateDiscoveryServiceAssignmentRequest) (*models.UpdateDiscoveryServiceAssignmentResponse, error) {
	model := request.DiscoveryServiceAssignment
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateDiscoveryServiceAssignment(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "discovery_service_assignment",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateDiscoveryServiceAssignmentResponse{
		DiscoveryServiceAssignment: model,
	}, nil
}

//DeleteDiscoveryServiceAssignment delete a resource.
func (db *DB) DeleteDiscoveryServiceAssignment(ctx context.Context, request *models.DeleteDiscoveryServiceAssignmentRequest) (*models.DeleteDiscoveryServiceAssignmentResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteDiscoveryServiceAssignment(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteDiscoveryServiceAssignmentResponse{
		ID: request.ID,
	}, nil
}

//GetDiscoveryServiceAssignment a Get request.
// nolint
func (db *DB) GetDiscoveryServiceAssignment(ctx context.Context, request *models.GetDiscoveryServiceAssignmentRequest) (response *models.GetDiscoveryServiceAssignmentResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListDiscoveryServiceAssignmentRequest{
		Spec: spec,
	}
	var result *models.ListDiscoveryServiceAssignmentResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listDiscoveryServiceAssignment(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.DiscoveryServiceAssignments) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetDiscoveryServiceAssignmentResponse{
		DiscoveryServiceAssignment: result.DiscoveryServiceAssignments[0],
	}
	return response, nil
}

//ListDiscoveryServiceAssignment handles a List service Request.
// nolint
func (db *DB) ListDiscoveryServiceAssignment(
	ctx context.Context,
	request *models.ListDiscoveryServiceAssignmentRequest) (response *models.ListDiscoveryServiceAssignmentResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listDiscoveryServiceAssignment(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
