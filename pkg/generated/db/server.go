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

const insertServerQuery = "insert into `server` (`uuid`,`user_id`,`updated`,`tenant_id`,`status`,`progress`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`name`,`locked`,`type`,`rel`,`href`,`id`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`_id`,`host_status`,`hostId`,`fq_name`,`links_type`,`links_rel`,`links_href`,`flavor_id`,`display_name`,`_created`,`config_drive`,`key_value_pair`,`addr`,`accessIPv6`,`accessIPv4`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteServerQuery = "delete from `server` where uuid = ?"

// ServerFields is db columns for Server
var ServerFields = []string{
	"uuid",
	"user_id",
	"updated",
	"tenant_id",
	"status",
	"progress",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"name",
	"locked",
	"type",
	"rel",
	"href",
	"id",
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
	"_id",
	"host_status",
	"hostId",
	"fq_name",
	"links_type",
	"links_rel",
	"links_href",
	"flavor_id",
	"display_name",
	"_created",
	"config_drive",
	"key_value_pair",
	"addr",
	"accessIPv6",
	"accessIPv4",
}

// ServerRefFields is db reference fields for Server
var ServerRefFields = map[string][]string{}

// ServerBackRefFields is db back reference fields for Server
var ServerBackRefFields = map[string][]string{}

// ServerParentTypes is possible parents for Server
var ServerParents = []string{}

// CreateServer inserts Server to DB
func CreateServer(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateServerRequest) error {
	model := request.Server
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServerQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.UUID),
		int(model.UserID),
		string(model.Updated),
		string(model.TenantID),
		string(model.Status),
		int(model.Progress),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		string(model.Name),
		bool(model.Locked),
		string(model.Image.Links.Type),
		string(model.Image.Links.Rel),
		string(model.Image.Links.Href),
		string(model.Image.ID),
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
		string(model.HostStatus),
		string(model.HostId),
		common.MustJSON(model.FQName),
		string(model.Flavor.Links.Type),
		string(model.Flavor.Links.Rel),
		string(model.Flavor.Links.Href),
		string(model.Flavor.ID),
		string(model.DisplayName),
		string(model.Created),
		bool(model.ConfigDrive),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.Addresses.Addr),
		string(model.AccessIPv6),
		string(model.AccessIPv4))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "server",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "server", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServer(values map[string]interface{}) (*models.Server, error) {
	m := models.MakeServer()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["user_id"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.UserID = castedValue

	}

	if value, ok := values["updated"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Updated = castedValue

	}

	if value, ok := values["tenant_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.TenantID = castedValue

	}

	if value, ok := values["status"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Status = castedValue

	}

	if value, ok := values["progress"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Progress = castedValue

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

	if value, ok := values["name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Name = castedValue

	}

	if value, ok := values["locked"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.Locked = castedValue

	}

	if value, ok := values["type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Image.Links.Type = castedValue

	}

	if value, ok := values["rel"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Image.Links.Rel = castedValue

	}

	if value, ok := values["href"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Image.Links.Href = castedValue

	}

	if value, ok := values["id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Image.ID = castedValue

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

	if value, ok := values["_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ID = castedValue

	}

	if value, ok := values["host_status"]; ok {

		castedValue := common.InterfaceToString(value)

		m.HostStatus = castedValue

	}

	if value, ok := values["hostId"]; ok {

		castedValue := common.InterfaceToString(value)

		m.HostId = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["links_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Flavor.Links.Type = castedValue

	}

	if value, ok := values["links_rel"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Flavor.Links.Rel = castedValue

	}

	if value, ok := values["links_href"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Flavor.Links.Href = castedValue

	}

	if value, ok := values["flavor_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Flavor.ID = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["_created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Created = castedValue

	}

	if value, ok := values["config_drive"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ConfigDrive = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["addr"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Addresses.Addr = castedValue

	}

	if value, ok := values["accessIPv6"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AccessIPv6 = castedValue

	}

	if value, ok := values["accessIPv4"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AccessIPv4 = castedValue

	}

	return m, nil
}

// ListServer lists Server with list spec.
func ListServer(ctx context.Context, tx *sql.Tx, request *models.ListServerRequest) (response *models.ListServerResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "server"
	qb.Fields = ServerFields
	qb.RefFields = ServerRefFields
	qb.BackRefFields = ServerBackRefFields
	result := models.MakeServerSlice()

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
		m, err := scanServer(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListServerResponse{
		Servers: result,
	}
	return response, nil
}

// UpdateServer updates a resource
func UpdateServer(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateServerRequest,
) error {
	//TODO
	return nil
}

// DeleteServer deletes a resource
func DeleteServer(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteServerRequest) error {
	deleteQuery := deleteServerQuery
	selectQuery := "select count(uuid) from server where uuid = ?"
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
