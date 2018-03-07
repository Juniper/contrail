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

// ProviderAttachmentFields is db columns for ProviderAttachment
var ProviderAttachmentFields = []string{
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

// ProviderAttachmentRefFields is db reference fields for ProviderAttachment
var ProviderAttachmentRefFields = map[string][]string{

	"virtual_router": []string{
	// <schema.Schema Value>

	},
}

// ProviderAttachmentBackRefFields is db back reference fields for ProviderAttachment
var ProviderAttachmentBackRefFields = map[string][]string{}

// ProviderAttachmentParentTypes is possible parents for ProviderAttachment
var ProviderAttachmentParents = []string{}

// CreateProviderAttachment inserts ProviderAttachment to DB
// nolint
func (db *DB) createProviderAttachment(
	ctx context.Context,
	request *models.CreateProviderAttachmentRequest) error {
	qb := db.queryBuilders["provider_attachment"]
	tx := GetTransaction(ctx)
	model := request.ProviderAttachment
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

	for _, ref := range model.VirtualRouterRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_router"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualRouterRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "provider_attachment",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "provider_attachment", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanProviderAttachment(values map[string]interface{}) (*models.ProviderAttachment, error) {
	m := models.MakeProviderAttachment()

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

	if value, ok := values["ref_virtual_router"]; ok {
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
			referenceModel := &models.ProviderAttachmentVirtualRouterRef{}
			referenceModel.UUID = uuid
			m.VirtualRouterRefs = append(m.VirtualRouterRefs, referenceModel)

		}
	}

	return m, nil
}

// ListProviderAttachment lists ProviderAttachment with list spec.
func (db *DB) listProviderAttachment(ctx context.Context, request *models.ListProviderAttachmentRequest) (response *models.ListProviderAttachmentResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["provider_attachment"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.ProviderAttachment{}

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
		m, err := scanProviderAttachment(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListProviderAttachmentResponse{
		ProviderAttachments: result,
	}
	return response, nil
}

// UpdateProviderAttachment updates a resource
func (db *DB) updateProviderAttachment(
	ctx context.Context,
	request *models.UpdateProviderAttachmentRequest,
) error {
	//TODO
	return nil
}

// DeleteProviderAttachment deletes a resource
func (db *DB) deleteProviderAttachment(
	ctx context.Context,
	request *models.DeleteProviderAttachmentRequest) error {
	qb := db.queryBuilders["provider_attachment"]

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

//CreateProviderAttachment handle a Create API
// nolint
func (db *DB) CreateProviderAttachment(
	ctx context.Context,
	request *models.CreateProviderAttachmentRequest) (*models.CreateProviderAttachmentResponse, error) {
	model := request.ProviderAttachment
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createProviderAttachment(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "provider_attachment",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateProviderAttachmentResponse{
		ProviderAttachment: request.ProviderAttachment,
	}, nil
}

//UpdateProviderAttachment handles a Update request.
func (db *DB) UpdateProviderAttachment(
	ctx context.Context,
	request *models.UpdateProviderAttachmentRequest) (*models.UpdateProviderAttachmentResponse, error) {
	model := request.ProviderAttachment
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateProviderAttachment(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "provider_attachment",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateProviderAttachmentResponse{
		ProviderAttachment: model,
	}, nil
}

//DeleteProviderAttachment delete a resource.
func (db *DB) DeleteProviderAttachment(ctx context.Context, request *models.DeleteProviderAttachmentRequest) (*models.DeleteProviderAttachmentResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteProviderAttachment(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteProviderAttachmentResponse{
		ID: request.ID,
	}, nil
}

//GetProviderAttachment a Get request.
func (db *DB) GetProviderAttachment(ctx context.Context, request *models.GetProviderAttachmentRequest) (response *models.GetProviderAttachmentResponse, err error) {
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
	listRequest := &models.ListProviderAttachmentRequest{
		Spec: spec,
	}
	var result *models.ListProviderAttachmentResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listProviderAttachment(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ProviderAttachments) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetProviderAttachmentResponse{
		ProviderAttachment: result.ProviderAttachments[0],
	}
	return response, nil
}

//ListProviderAttachment handles a List service Request.
// nolint
func (db *DB) ListProviderAttachment(
	ctx context.Context,
	request *models.ListProviderAttachmentRequest) (response *models.ListProviderAttachmentResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listProviderAttachment(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
