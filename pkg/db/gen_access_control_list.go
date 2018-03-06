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

// AccessControlListFields is db columns for AccessControlList
var AccessControlListFields = []string{
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
	"access_control_list_hash",
	"dynamic",
	"acl_rule",
}

// AccessControlListRefFields is db reference fields for AccessControlList
var AccessControlListRefFields = map[string][]string{}

// AccessControlListBackRefFields is db back reference fields for AccessControlList
var AccessControlListBackRefFields = map[string][]string{}

// AccessControlListParentTypes is possible parents for AccessControlList
var AccessControlListParents = []string{

	"security_group",

	"virtual_network",
}

// CreateAccessControlList inserts AccessControlList to DB
// nolint
func (db *DB) createAccessControlList(
	ctx context.Context,
	request *models.CreateAccessControlListRequest) error {
	qb := db.queryBuilders["access_control_list"]
	tx := GetTransaction(ctx)
	model := request.AccessControlList
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
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()),
		int(model.GetAccessControlListHash()),
		bool(model.GetAccessControlListEntries().GetDynamic()),
		common.MustJSON(model.GetAccessControlListEntries().GetACLRule()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "access_control_list",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "access_control_list", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanAccessControlList(values map[string]interface{}) (*models.AccessControlList, error) {
	m := models.MakeAccessControlList()

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

	if value, ok := values["access_control_list_hash"]; ok {

		m.AccessControlListHash = common.InterfaceToInt64(value)

	}

	if value, ok := values["dynamic"]; ok {

		m.AccessControlListEntries.Dynamic = common.InterfaceToBool(value)

	}

	if value, ok := values["acl_rule"]; ok {

		json.Unmarshal(value.([]byte), &m.AccessControlListEntries.ACLRule)

	}

	return m, nil
}

// ListAccessControlList lists AccessControlList with list spec.
func (db *DB) listAccessControlList(ctx context.Context, request *models.ListAccessControlListRequest) (response *models.ListAccessControlListResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["access_control_list"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.AccessControlList{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
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
		m, err := scanAccessControlList(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListAccessControlListResponse{
		AccessControlLists: result,
	}
	return response, nil
}

// UpdateAccessControlList updates a resource
func (db *DB) updateAccessControlList(
	ctx context.Context,
	request *models.UpdateAccessControlListRequest,
) error {
	//TODO
	return nil
}

// DeleteAccessControlList deletes a resource
func (db *DB) deleteAccessControlList(
	ctx context.Context,
	request *models.DeleteAccessControlListRequest) error {
	qb := db.queryBuilders["access_control_list"]

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

//CreateAccessControlList handle a Create API
// nolint
func (db *DB) CreateAccessControlList(
	ctx context.Context,
	request *models.CreateAccessControlListRequest) (*models.CreateAccessControlListResponse, error) {
	model := request.AccessControlList
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createAccessControlList(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "access_control_list",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAccessControlListResponse{
		AccessControlList: request.AccessControlList,
	}, nil
}

//UpdateAccessControlList handles a Update request.
func (db *DB) UpdateAccessControlList(
	ctx context.Context,
	request *models.UpdateAccessControlListRequest) (*models.UpdateAccessControlListResponse, error) {
	model := request.AccessControlList
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateAccessControlList(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "access_control_list",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAccessControlListResponse{
		AccessControlList: model,
	}, nil
}

//DeleteAccessControlList delete a resource.
func (db *DB) DeleteAccessControlList(ctx context.Context, request *models.DeleteAccessControlListRequest) (*models.DeleteAccessControlListResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteAccessControlList(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAccessControlListResponse{
		ID: request.ID,
	}, nil
}

//GetAccessControlList a Get request.
func (db *DB) GetAccessControlList(ctx context.Context, request *models.GetAccessControlListRequest) (response *models.GetAccessControlListResponse, err error) {
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
	listRequest := &models.ListAccessControlListRequest{
		Spec: spec,
	}
	var result *models.ListAccessControlListResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listAccessControlList(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.AccessControlLists) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAccessControlListResponse{
		AccessControlList: result.AccessControlLists[0],
	}
	return response, nil
}

//ListAccessControlList handles a List service Request.
// nolint
func (db *DB) ListAccessControlList(
	ctx context.Context,
	request *models.ListAccessControlListRequest) (response *models.ListAccessControlListResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listAccessControlList(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
