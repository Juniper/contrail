package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
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
func CreateOsImage(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateOsImageRequest) error {
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
	_, err = stmt.ExecContext(ctx, string(model.Visibility),
		string(model.UUID),
		string(model.UpdatedAt),
		string(model.Tags),
		string(model.Status),
		int(model.Size_),
		bool(model.Protected),
		string(model.Property),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		string(model.Owner),
		string(model.Name),
		int(model.MinRAM),
		int(model.MinDisk),
		string(model.Location),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		string(model.ID),
		common.MustJSON(model.FQName),
		string(model.File),
		string(model.DisplayName),
		string(model.DiskFormat),
		string(model.CreatedAt),
		string(model.ContainerFormat),
		string(model.Checksum),
		common.MustJSON(model.Annotations.KeyValuePair))
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
	err = common.CreateSharing(tx, "os_image", model.UUID, model.Perms2.Share)
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

		castedValue := common.InterfaceToString(value)

		m.Visibility = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["updated_at"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UpdatedAt = castedValue

	}

	if value, ok := values["tags"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Tags = castedValue

	}

	if value, ok := values["status"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Status = castedValue

	}

	if value, ok := values["size"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Size_ = castedValue

	}

	if value, ok := values["protected"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.Protected = castedValue

	}

	if value, ok := values["property"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Property = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Owner = castedValue

	}

	if value, ok := values["name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Name = castedValue

	}

	if value, ok := values["min_ram"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MinRAM = castedValue

	}

	if value, ok := values["min_disk"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MinDisk = castedValue

	}

	if value, ok := values["location"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Location = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["file"]; ok {

		castedValue := common.InterfaceToString(value)

		m.File = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["disk_format"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DiskFormat = castedValue

	}

	if value, ok := values["created_at"]; ok {

		castedValue := common.InterfaceToString(value)

		m.CreatedAt = castedValue

	}

	if value, ok := values["container_format"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ContainerFormat = castedValue

	}

	if value, ok := values["checksum"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Checksum = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListOsImage lists OsImage with list spec.
func ListOsImage(ctx context.Context, tx *sql.Tx, request *models.ListOsImageRequest) (response *models.ListOsImageResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "os_image"
	qb.Fields = OsImageFields
	qb.RefFields = OsImageRefFields
	qb.BackRefFields = OsImageBackRefFields
	result := models.MakeOsImageSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
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
func UpdateOsImage(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateOsImageRequest,
) error {
	//TODO
	return nil
}

// DeleteOsImage deletes a resource
func DeleteOsImage(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteOsImageRequest) error {
	deleteQuery := deleteOsImageQuery
	selectQuery := "select count(uuid) from os_image where uuid = ?"
	var err error
	var count int
	uuid := request.ID
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
