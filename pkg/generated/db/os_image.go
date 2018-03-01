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

const insertOsImageQuery = "insert into `os_image` (`visibility`,`uuid`,`updated_at`,`tags`,`status`,`size`,`protected`,`property`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`_owner`,`name`,`min_ram`,`min_disk`,`location`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`id`,`fq_name`,`file`,`display_name`,`disk_format`,`created_at`,`container_format`,`checksum`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteOsImageQuery = "delete from `os_image` where uuid = ?"

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
func (db *DB) createOsImage(
	ctx context.Context,
	request *models.CreateOsImageRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.OsImage
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOsImageQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertOsImageQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetVisibility()),
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
		string(model.GetChecksum()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "os_image",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "os_image", model.UUID, model.GetPerms2().GetShare())
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

		m.Visibility = schema.InterfaceToString(value)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["updated_at"]; ok {

		m.UpdatedAt = schema.InterfaceToString(value)

	}

	if value, ok := values["tags"]; ok {

		m.Tags = schema.InterfaceToString(value)

	}

	if value, ok := values["status"]; ok {

		m.Status = schema.InterfaceToString(value)

	}

	if value, ok := values["size"]; ok {

		m.Size_ = schema.InterfaceToInt64(value)

	}

	if value, ok := values["protected"]; ok {

		m.Protected = schema.InterfaceToBool(value)

	}

	if value, ok := values["property"]; ok {

		m.Property = schema.InterfaceToString(value)

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

	if value, ok := values["_owner"]; ok {

		m.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["name"]; ok {

		m.Name = schema.InterfaceToString(value)

	}

	if value, ok := values["min_ram"]; ok {

		m.MinRAM = schema.InterfaceToInt64(value)

	}

	if value, ok := values["min_disk"]; ok {

		m.MinDisk = schema.InterfaceToInt64(value)

	}

	if value, ok := values["location"]; ok {

		m.Location = schema.InterfaceToString(value)

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

	if value, ok := values["id"]; ok {

		m.ID = schema.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["file"]; ok {

		m.File = schema.InterfaceToString(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["disk_format"]; ok {

		m.DiskFormat = schema.InterfaceToString(value)

	}

	if value, ok := values["created_at"]; ok {

		m.CreatedAt = schema.InterfaceToString(value)

	}

	if value, ok := values["container_format"]; ok {

		m.ContainerFormat = schema.InterfaceToString(value)

	}

	if value, ok := values["checksum"]; ok {

		m.Checksum = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListOsImage lists OsImage with list spec.
func (db *DB) listOsImage(ctx context.Context, request *models.ListOsImageRequest) (response *models.ListOsImageResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "os_image"
	qb.Fields = OsImageFields
	qb.RefFields = OsImageRefFields
	qb.BackRefFields = OsImageBackRefFields
	result := []*models.OsImage{}

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
	deleteQuery := deleteOsImageQuery
	selectQuery := "select count(uuid) from os_image where uuid = ?"
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

//CreateOsImage handle a Create API
func (db *DB) CreateOsImage(
	ctx context.Context,
	request *models.CreateOsImageRequest) (*models.CreateOsImageResponse, error) {
	model := request.OsImage
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
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
	if err := common.DoInTransaction(
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
	if err := common.DoInTransaction(
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
	if err := common.DoInTransaction(
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
func (db *DB) ListOsImage(
	ctx context.Context,
	request *models.ListOsImageRequest) (response *models.ListOsImageResponse, err error) {
	if err := common.DoInTransaction(
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
