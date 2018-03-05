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

// OsImageFields is db columns for OsImage
var OsImageFields = []string{
	"visibility",
	"uuid",
	"updated_at",
	"tags",
	"status",
	"size",
	"protected",
	"property",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"_owner",
	"name",
	"min_ram",
	"min_disk",
	"location",
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
	"id",
	"fq_name",
	"file",
	"display_name",
	"disk_format",
	"created_at",
	"container_format",
	"configuration_version",
	"checksum",
	"key_value_pair",
}

// OsImageRefFields is db reference fields for OsImage
var OsImageRefFields = map[string][]string{}

// OsImageBackRefFields is db back reference fields for OsImage
var OsImageBackRefFields = map[string][]string{}

// OsImageParentTypes is possible parents for OsImage
var OsImageParents = []string{}

// CreateOsImage inserts OsImage to DB
// nolint
func (db *DB) createOsImage(
	ctx context.Context,
	request *models.CreateOsImageRequest) error {
	qb := db.queryBuilders["os_image"]
	tx := GetTransaction(ctx)
	model := request.OsImage
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetVisibility()),
		string(model.GetUUID()),
		string(model.GetUpdatedAt()),
		string(model.GetTags()),
		string(model.GetStatus()),
		int(model.GetSize_()),
		bool(model.GetProtected()),
		string(model.GetProperty()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetOwner()),
		string(model.GetName()),
		int(model.GetMinRAM()),
		int(model.GetMinDisk()),
		string(model.GetLocation()),
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
		string(model.GetID()),
		common.MustJSON(model.GetFQName()),
		string(model.GetFile()),
		string(model.GetDisplayName()),
		string(model.GetDiskFormat()),
		string(model.GetCreatedAt()),
		string(model.GetContainerFormat()),
		int(model.GetConfigurationVersion()),
		string(model.GetChecksum()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "os_image",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "os_image", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanOsImage(values map[string]interface{}) (*models.OsImage, error) {
	m := models.MakeOsImage()

	if value, ok := values["visibility"]; ok {

		m.Visibility = common.InterfaceToString(value)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["updated_at"]; ok {

		m.UpdatedAt = common.InterfaceToString(value)

	}

	if value, ok := values["tags"]; ok {

		m.Tags = common.InterfaceToString(value)

	}

	if value, ok := values["status"]; ok {

		m.Status = common.InterfaceToString(value)

	}

	if value, ok := values["size"]; ok {

		m.Size_ = common.InterfaceToInt64(value)

	}

	if value, ok := values["protected"]; ok {

		m.Protected = common.InterfaceToBool(value)

	}

	if value, ok := values["property"]; ok {

		m.Property = common.InterfaceToString(value)

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

	if value, ok := values["_owner"]; ok {

		m.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["name"]; ok {

		m.Name = common.InterfaceToString(value)

	}

	if value, ok := values["min_ram"]; ok {

		m.MinRAM = common.InterfaceToInt64(value)

	}

	if value, ok := values["min_disk"]; ok {

		m.MinDisk = common.InterfaceToInt64(value)

	}

	if value, ok := values["location"]; ok {

		m.Location = common.InterfaceToString(value)

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

	if value, ok := values["id"]; ok {

		m.ID = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["file"]; ok {

		m.File = common.InterfaceToString(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["disk_format"]; ok {

		m.DiskFormat = common.InterfaceToString(value)

	}

	if value, ok := values["created_at"]; ok {

		m.CreatedAt = common.InterfaceToString(value)

	}

	if value, ok := values["container_format"]; ok {

		m.ContainerFormat = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["checksum"]; ok {

		m.Checksum = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListOsImage lists OsImage with list spec.
func (db *DB) listOsImage(ctx context.Context, request *models.ListOsImageRequest) (response *models.ListOsImageResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["os_image"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.OsImage{}

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
		m, err := scanOsImage(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListOsImageResponse{
		OsImages: result,
	}
	return response, nil
}

// UpdateOsImage updates a resource
func (db *DB) updateOsImage(
	ctx context.Context,
	request *models.UpdateOsImageRequest,
) error {
	//TODO
	return nil
}

// DeleteOsImage deletes a resource
func (db *DB) deleteOsImage(
	ctx context.Context,
	request *models.DeleteOsImageRequest) error {
	qb := db.queryBuilders["os_image"]

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

//CreateOsImage handle a Create API
// nolint
func (db *DB) CreateOsImage(
	ctx context.Context,
	request *models.CreateOsImageRequest) (*models.CreateOsImageResponse, error) {
	model := request.OsImage
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createOsImage(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "os_image",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateOsImageResponse{
		OsImage: request.OsImage,
	}, nil
}

//UpdateOsImage handles a Update request.
func (db *DB) UpdateOsImage(
	ctx context.Context,
	request *models.UpdateOsImageRequest) (*models.UpdateOsImageResponse, error) {
	model := request.OsImage
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateOsImage(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "os_image",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateOsImageResponse{
		OsImage: model,
	}, nil
}

//DeleteOsImage delete a resource.
func (db *DB) DeleteOsImage(ctx context.Context, request *models.DeleteOsImageRequest) (*models.DeleteOsImageResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteOsImage(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteOsImageResponse{
		ID: request.ID,
	}, nil
}

//GetOsImage a Get request.
func (db *DB) GetOsImage(ctx context.Context, request *models.GetOsImageRequest) (response *models.GetOsImageResponse, err error) {
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
	listRequest := &models.ListOsImageRequest{
		Spec: spec,
	}
	var result *models.ListOsImageResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listOsImage(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.OsImages) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetOsImageResponse{
		OsImage: result.OsImages[0],
	}
	return response, nil
}

//ListOsImage handles a List service Request.
// nolint
func (db *DB) ListOsImage(
	ctx context.Context,
	request *models.ListOsImageRequest) (response *models.ListOsImageResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listOsImage(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
