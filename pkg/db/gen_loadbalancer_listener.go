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

const insertLoadbalancerListenerQuery = "insert into `loadbalancer_listener` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`sni_containers`,`protocol_port`,`protocol`,`default_tls_container`,`connection_limit`,`admin_state`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteLoadbalancerListenerQuery = "delete from `loadbalancer_listener` where uuid = ?"

// LoadbalancerListenerFields is db columns for LoadbalancerListener
var LoadbalancerListenerFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"sni_containers",
	"protocol_port",
	"protocol",
	"default_tls_container",
	"connection_limit",
	"admin_state",
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

// LoadbalancerListenerRefFields is db reference fields for LoadbalancerListener
var LoadbalancerListenerRefFields = map[string][]string{

	"loadbalancer": []string{
	// <schema.Schema Value>

	},
}

// LoadbalancerListenerBackRefFields is db back reference fields for LoadbalancerListener
var LoadbalancerListenerBackRefFields = map[string][]string{}

// LoadbalancerListenerParentTypes is possible parents for LoadbalancerListener
var LoadbalancerListenerParents = []string{

	"project",
}

const insertLoadbalancerListenerLoadbalancerQuery = "insert into `ref_loadbalancer_listener_loadbalancer` (`from`, `to` ) values (?, ?);"

// CreateLoadbalancerListener inserts LoadbalancerListener to DB
// nolint
func (db *DB) createLoadbalancerListener(
	ctx context.Context,
	request *models.CreateLoadbalancerListenerRequest) error {
	tx := GetTransaction(ctx)
	model := request.LoadbalancerListener
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerListenerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLoadbalancerListenerQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		common.MustJSON(model.GetLoadbalancerListenerProperties().GetSniContainers()),
		int(model.GetLoadbalancerListenerProperties().GetProtocolPort()),
		string(model.GetLoadbalancerListenerProperties().GetProtocol()),
		string(model.GetLoadbalancerListenerProperties().GetDefaultTLSContainer()),
		int(model.GetLoadbalancerListenerProperties().GetConnectionLimit()),
		bool(model.GetLoadbalancerListenerProperties().GetAdminState()),
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

	stmtLoadbalancerRef, err := tx.Prepare(insertLoadbalancerListenerLoadbalancerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing LoadbalancerRefs create statement failed")
	}
	defer stmtLoadbalancerRef.Close()
	for _, ref := range model.LoadbalancerRefs {

		_, err = stmtLoadbalancerRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LoadbalancerRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "loadbalancer_listener",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "loadbalancer_listener", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

// nolint
func scanLoadbalancerListener(values map[string]interface{}) (*models.LoadbalancerListener, error) {
	m := models.MakeLoadbalancerListener()

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

	if value, ok := values["sni_containers"]; ok {

		json.Unmarshal(value.([]byte), &m.LoadbalancerListenerProperties.SniContainers)

	}

	if value, ok := values["protocol_port"]; ok {

		m.LoadbalancerListenerProperties.ProtocolPort = common.InterfaceToInt64(value)

	}

	if value, ok := values["protocol"]; ok {

		m.LoadbalancerListenerProperties.Protocol = common.InterfaceToString(value)

	}

	if value, ok := values["default_tls_container"]; ok {

		m.LoadbalancerListenerProperties.DefaultTLSContainer = common.InterfaceToString(value)

	}

	if value, ok := values["connection_limit"]; ok {

		m.LoadbalancerListenerProperties.ConnectionLimit = common.InterfaceToInt64(value)

	}

	if value, ok := values["admin_state"]; ok {

		m.LoadbalancerListenerProperties.AdminState = common.InterfaceToBool(value)

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

	if value, ok := values["ref_loadbalancer"]; ok {
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
			referenceModel := &models.LoadbalancerListenerLoadbalancerRef{}
			referenceModel.UUID = uuid
			m.LoadbalancerRefs = append(m.LoadbalancerRefs, referenceModel)

		}
	}

	return m, nil
}

// ListLoadbalancerListener lists LoadbalancerListener with list spec.
// nolint
func (db *DB) listLoadbalancerListener(ctx context.Context, request *models.ListLoadbalancerListenerRequest) (response *models.ListLoadbalancerListenerResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)
	qb := &ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "loadbalancer_listener"
	qb.Fields = LoadbalancerListenerFields
	qb.RefFields = LoadbalancerListenerRefFields
	qb.BackRefFields = LoadbalancerListenerBackRefFields
	result := []*models.LoadbalancerListener{}

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
		m, err := scanLoadbalancerListener(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListLoadbalancerListenerResponse{
		LoadbalancerListeners: result,
	}
	return response, nil
}

// UpdateLoadbalancerListener updates a resource
// nolint
func (db *DB) updateLoadbalancerListener(
	ctx context.Context,
	request *models.UpdateLoadbalancerListenerRequest,
) error {
	//TODO
	return nil
}

// DeleteLoadbalancerListener deletes a resource
// nolint
func (db *DB) deleteLoadbalancerListener(
	ctx context.Context,
	request *models.DeleteLoadbalancerListenerRequest) error {
	deleteQuery := deleteLoadbalancerListenerQuery
	selectQuery := "select count(uuid) from loadbalancer_listener where uuid = ?"
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

//CreateLoadbalancerListener handle a Create API
// nolint
func (db *DB) CreateLoadbalancerListener(
	ctx context.Context,
	request *models.CreateLoadbalancerListenerRequest) (*models.CreateLoadbalancerListenerResponse, error) {
	model := request.LoadbalancerListener
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createLoadbalancerListener(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_listener",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLoadbalancerListenerResponse{
		LoadbalancerListener: request.LoadbalancerListener,
	}, nil
}

//UpdateLoadbalancerListener handles a Update request.
func (db *DB) UpdateLoadbalancerListener(
	ctx context.Context,
	request *models.UpdateLoadbalancerListenerRequest) (*models.UpdateLoadbalancerListenerResponse, error) {
	model := request.LoadbalancerListener
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateLoadbalancerListener(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_listener",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLoadbalancerListenerResponse{
		LoadbalancerListener: model,
	}, nil
}

//DeleteLoadbalancerListener delete a resource.
func (db *DB) DeleteLoadbalancerListener(ctx context.Context, request *models.DeleteLoadbalancerListenerRequest) (*models.DeleteLoadbalancerListenerResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteLoadbalancerListener(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLoadbalancerListenerResponse{
		ID: request.ID,
	}, nil
}

//GetLoadbalancerListener a Get request.
// nolint
func (db *DB) GetLoadbalancerListener(ctx context.Context, request *models.GetLoadbalancerListenerRequest) (response *models.GetLoadbalancerListenerResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListLoadbalancerListenerRequest{
		Spec: spec,
	}
	var result *models.ListLoadbalancerListenerResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listLoadbalancerListener(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LoadbalancerListeners) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLoadbalancerListenerResponse{
		LoadbalancerListener: result.LoadbalancerListeners[0],
	}
	return response, nil
}

//ListLoadbalancerListener handles a List service Request.
// nolint
func (db *DB) ListLoadbalancerListener(
	ctx context.Context,
	request *models.ListLoadbalancerListenerRequest) (response *models.ListLoadbalancerListenerResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listLoadbalancerListener(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
