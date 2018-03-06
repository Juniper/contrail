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
// nolint
func (db *DB) createServer(
	ctx context.Context,
	request *models.CreateServerRequest) error {
	qb := db.queryBuilders["server"]
	tx := GetTransaction(ctx)
	model := request.Server
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		int(model.GetUserID()),
		string(model.GetUpdated()),
		string(model.GetTenantID()),
		string(model.GetStatus()),
		int(model.GetProgress()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetName()),
		bool(model.GetLocked()),
		string(model.GetImage().GetLinks().GetType()),
		string(model.GetImage().GetLinks().GetRel()),
		string(model.GetImage().GetLinks().GetHref()),
		string(model.GetImage().GetID()),
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
		string(model.GetHostStatus()),
		string(model.GetHostId()),
		common.MustJSON(model.GetFQName()),
		string(model.GetFlavor().GetLinks().GetType()),
		string(model.GetFlavor().GetLinks().GetRel()),
		string(model.GetFlavor().GetLinks().GetHref()),
		string(model.GetFlavor().GetID()),
		string(model.GetDisplayName()),
		string(model.GetCreated()),
		bool(model.GetConfigDrive()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()),
		string(model.GetAddresses().GetAddr()),
		string(model.GetAccessIPv6()),
		string(model.GetAccessIPv4()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "server",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "server", model.UUID, model.GetPerms2().GetShare())
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

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["user_id"]; ok {

		m.UserID = common.InterfaceToInt64(value)

	}

	if value, ok := values["updated"]; ok {

		m.Updated = common.InterfaceToString(value)

	}

	if value, ok := values["tenant_id"]; ok {

		m.TenantID = common.InterfaceToString(value)

	}

	if value, ok := values["status"]; ok {

		m.Status = common.InterfaceToString(value)

	}

	if value, ok := values["progress"]; ok {

		m.Progress = common.InterfaceToInt64(value)

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

	if value, ok := values["name"]; ok {

		m.Name = common.InterfaceToString(value)

	}

	if value, ok := values["locked"]; ok {

		m.Locked = common.InterfaceToBool(value)

	}

	if value, ok := values["type"]; ok {

		m.Image.Links.Type = common.InterfaceToString(value)

	}

	if value, ok := values["rel"]; ok {

		m.Image.Links.Rel = common.InterfaceToString(value)

	}

	if value, ok := values["href"]; ok {

		m.Image.Links.Href = common.InterfaceToString(value)

	}

	if value, ok := values["id"]; ok {

		m.Image.ID = common.InterfaceToString(value)

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

	if value, ok := values["_id"]; ok {

		m.ID = common.InterfaceToString(value)

	}

	if value, ok := values["host_status"]; ok {

		m.HostStatus = common.InterfaceToString(value)

	}

	if value, ok := values["hostId"]; ok {

		m.HostId = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["links_type"]; ok {

		m.Flavor.Links.Type = common.InterfaceToString(value)

	}

	if value, ok := values["links_rel"]; ok {

		m.Flavor.Links.Rel = common.InterfaceToString(value)

	}

	if value, ok := values["links_href"]; ok {

		m.Flavor.Links.Href = common.InterfaceToString(value)

	}

	if value, ok := values["flavor_id"]; ok {

		m.Flavor.ID = common.InterfaceToString(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["_created"]; ok {

		m.Created = common.InterfaceToString(value)

	}

	if value, ok := values["config_drive"]; ok {

		m.ConfigDrive = common.InterfaceToBool(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["addr"]; ok {

		m.Addresses.Addr = common.InterfaceToString(value)

	}

	if value, ok := values["accessIPv6"]; ok {

		m.AccessIPv6 = common.InterfaceToString(value)

	}

	if value, ok := values["accessIPv4"]; ok {

		m.AccessIPv4 = common.InterfaceToString(value)

	}

	return m, nil
}

// ListServer lists Server with list spec.
func (db *DB) listServer(ctx context.Context, request *models.ListServerRequest) (response *models.ListServerResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["server"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.Server{}

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
func (db *DB) updateServer(
	ctx context.Context,
	request *models.UpdateServerRequest,
) error {
	//TODO
	return nil
}

// DeleteServer deletes a resource
func (db *DB) deleteServer(
	ctx context.Context,
	request *models.DeleteServerRequest) error {
	qb := db.queryBuilders["server"]

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

//CreateServer handle a Create API
// nolint
func (db *DB) CreateServer(
	ctx context.Context,
	request *models.CreateServerRequest) (*models.CreateServerResponse, error) {
	model := request.Server
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createServer(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "server",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServerResponse{
		Server: request.Server,
	}, nil
}

//UpdateServer handles a Update request.
func (db *DB) UpdateServer(
	ctx context.Context,
	request *models.UpdateServerRequest) (*models.UpdateServerResponse, error) {
	model := request.Server
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateServer(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "server",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServerResponse{
		Server: model,
	}, nil
}

//DeleteServer delete a resource.
func (db *DB) DeleteServer(ctx context.Context, request *models.DeleteServerRequest) (*models.DeleteServerResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteServer(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServerResponse{
		ID: request.ID,
	}, nil
}

//GetServer a Get request.
func (db *DB) GetServer(ctx context.Context, request *models.GetServerRequest) (response *models.GetServerResponse, err error) {
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
	listRequest := &models.ListServerRequest{
		Spec: spec,
	}
	var result *models.ListServerResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listServer(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Servers) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServerResponse{
		Server: result.Servers[0],
	}
	return response, nil
}

//ListServer handles a List service Request.
// nolint
func (db *DB) ListServer(
	ctx context.Context,
	request *models.ListServerRequest) (response *models.ListServerResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listServer(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
