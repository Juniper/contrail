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

const insertNetworkDeviceConfigQuery = "insert into `network_device_config` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteNetworkDeviceConfigQuery = "delete from `network_device_config` where uuid = ?"

// NetworkDeviceConfigFields is db columns for NetworkDeviceConfig
var NetworkDeviceConfigFields = []string{
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

// NetworkDeviceConfigRefFields is db reference fields for NetworkDeviceConfig
var NetworkDeviceConfigRefFields = map[string][]string{

	"physical_router": []string{
	// <schema.Schema Value>

	},
}

// NetworkDeviceConfigBackRefFields is db back reference fields for NetworkDeviceConfig
var NetworkDeviceConfigBackRefFields = map[string][]string{}

// NetworkDeviceConfigParentTypes is possible parents for NetworkDeviceConfig
var NetworkDeviceConfigParents = []string{}

const insertNetworkDeviceConfigPhysicalRouterQuery = "insert into `ref_network_device_config_physical_router` (`from`, `to` ) values (?, ?);"

// CreateNetworkDeviceConfig inserts NetworkDeviceConfig to DB
func (db *DB) createNetworkDeviceConfig(
	ctx context.Context,
	request *models.CreateNetworkDeviceConfigRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.NetworkDeviceConfig
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkDeviceConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertNetworkDeviceConfigQuery,
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

	stmtPhysicalRouterRef, err := tx.Prepare(insertNetworkDeviceConfigPhysicalRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PhysicalRouterRefs create statement failed")
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {

		_, err = stmtPhysicalRouterRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "network_device_config",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "network_device_config", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanNetworkDeviceConfig(values map[string]interface{}) (*models.NetworkDeviceConfig, error) {
	m := models.MakeNetworkDeviceConfig()

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

	if value, ok := values["ref_physical_router"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.NetworkDeviceConfigPhysicalRouterRef{}
			referenceModel.UUID = uuid
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	return m, nil
}

// ListNetworkDeviceConfig lists NetworkDeviceConfig with list spec.
func (db *DB) listNetworkDeviceConfig(ctx context.Context, request *models.ListNetworkDeviceConfigRequest) (response *models.ListNetworkDeviceConfigResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "network_device_config"
	qb.Fields = NetworkDeviceConfigFields
	qb.RefFields = NetworkDeviceConfigRefFields
	qb.BackRefFields = NetworkDeviceConfigBackRefFields
	result := []*models.NetworkDeviceConfig{}

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
		m, err := scanNetworkDeviceConfig(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListNetworkDeviceConfigResponse{
		NetworkDeviceConfigs: result,
	}
	return response, nil
}

// UpdateNetworkDeviceConfig updates a resource
func (db *DB) updateNetworkDeviceConfig(
	ctx context.Context,
	request *models.UpdateNetworkDeviceConfigRequest,
) error {
	//TODO
	return nil
}

// DeleteNetworkDeviceConfig deletes a resource
func (db *DB) deleteNetworkDeviceConfig(
	ctx context.Context,
	request *models.DeleteNetworkDeviceConfigRequest) error {
	deleteQuery := deleteNetworkDeviceConfigQuery
	selectQuery := "select count(uuid) from network_device_config where uuid = ?"
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

//CreateNetworkDeviceConfig handle a Create API
func (db *DB) CreateNetworkDeviceConfig(
	ctx context.Context,
	request *models.CreateNetworkDeviceConfigRequest) (*models.CreateNetworkDeviceConfigResponse, error) {
	model := request.NetworkDeviceConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createNetworkDeviceConfig(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_device_config",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateNetworkDeviceConfigResponse{
		NetworkDeviceConfig: request.NetworkDeviceConfig,
	}, nil
}

//UpdateNetworkDeviceConfig handles a Update request.
func (db *DB) UpdateNetworkDeviceConfig(
	ctx context.Context,
	request *models.UpdateNetworkDeviceConfigRequest) (*models.UpdateNetworkDeviceConfigResponse, error) {
	model := request.NetworkDeviceConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateNetworkDeviceConfig(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_device_config",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateNetworkDeviceConfigResponse{
		NetworkDeviceConfig: model,
	}, nil
}

//DeleteNetworkDeviceConfig delete a resource.
func (db *DB) DeleteNetworkDeviceConfig(ctx context.Context, request *models.DeleteNetworkDeviceConfigRequest) (*models.DeleteNetworkDeviceConfigResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteNetworkDeviceConfig(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteNetworkDeviceConfigResponse{
		ID: request.ID,
	}, nil
}

//GetNetworkDeviceConfig a Get request.
func (db *DB) GetNetworkDeviceConfig(ctx context.Context, request *models.GetNetworkDeviceConfigRequest) (response *models.GetNetworkDeviceConfigResponse, err error) {
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
	listRequest := &models.ListNetworkDeviceConfigRequest{
		Spec: spec,
	}
	var result *models.ListNetworkDeviceConfigResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listNetworkDeviceConfig(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.NetworkDeviceConfigs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetNetworkDeviceConfigResponse{
		NetworkDeviceConfig: result.NetworkDeviceConfigs[0],
	}
	return response, nil
}

//ListNetworkDeviceConfig handles a List service Request.
func (db *DB) ListNetworkDeviceConfig(
	ctx context.Context,
	request *models.ListNetworkDeviceConfigRequest) (response *models.ListNetworkDeviceConfigResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listNetworkDeviceConfig(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
