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

// BridgeDomainFields is db columns for BridgeDomain
var BridgeDomainFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"mac_move_time_window",
	"mac_move_limit_action",
	"mac_move_limit",
	"mac_limit_action",
	"mac_limit",
	"mac_learning_enabled",
	"mac_aging_time",
	"isid",
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

// BridgeDomainRefFields is db reference fields for BridgeDomain
var BridgeDomainRefFields = map[string][]string{}

// BridgeDomainBackRefFields is db back reference fields for BridgeDomain
var BridgeDomainBackRefFields = map[string][]string{}

// BridgeDomainParentTypes is possible parents for BridgeDomain
var BridgeDomainParents = []string{

	"virtual_network",
}

// CreateBridgeDomain inserts BridgeDomain to DB
// nolint
func (db *DB) createBridgeDomain(
	ctx context.Context,
	request *models.CreateBridgeDomainRequest) error {
	qb := db.queryBuilders["bridge_domain"]
	tx := GetTransaction(ctx)
	model := request.BridgeDomain
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		int(model.GetMacMoveControl().GetMacMoveTimeWindow()),
		string(model.GetMacMoveControl().GetMacMoveLimitAction()),
		int(model.GetMacMoveControl().GetMacMoveLimit()),
		string(model.GetMacLimitControl().GetMacLimitAction()),
		int(model.GetMacLimitControl().GetMacLimit()),
		bool(model.GetMacLearningEnabled()),
		int(model.GetMacAgingTime()),
		int(model.GetIsid()),
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

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "bridge_domain",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "bridge_domain", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanBridgeDomain(values map[string]interface{}) (*models.BridgeDomain, error) {
	m := models.MakeBridgeDomain()

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

	if value, ok := values["mac_move_time_window"]; ok {

		m.MacMoveControl.MacMoveTimeWindow = common.InterfaceToInt64(value)

	}

	if value, ok := values["mac_move_limit_action"]; ok {

		m.MacMoveControl.MacMoveLimitAction = common.InterfaceToString(value)

	}

	if value, ok := values["mac_move_limit"]; ok {

		m.MacMoveControl.MacMoveLimit = common.InterfaceToInt64(value)

	}

	if value, ok := values["mac_limit_action"]; ok {

		m.MacLimitControl.MacLimitAction = common.InterfaceToString(value)

	}

	if value, ok := values["mac_limit"]; ok {

		m.MacLimitControl.MacLimit = common.InterfaceToInt64(value)

	}

	if value, ok := values["mac_learning_enabled"]; ok {

		m.MacLearningEnabled = common.InterfaceToBool(value)

	}

	if value, ok := values["mac_aging_time"]; ok {

		m.MacAgingTime = common.InterfaceToInt64(value)

	}

	if value, ok := values["isid"]; ok {

		m.Isid = common.InterfaceToInt64(value)

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

	return m, nil
}

// ListBridgeDomain lists BridgeDomain with list spec.
func (db *DB) listBridgeDomain(ctx context.Context, request *models.ListBridgeDomainRequest) (response *models.ListBridgeDomainResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["bridge_domain"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.BridgeDomain{}

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
		m, err := scanBridgeDomain(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListBridgeDomainResponse{
		BridgeDomains: result,
	}
	return response, nil
}

// UpdateBridgeDomain updates a resource
func (db *DB) updateBridgeDomain(
	ctx context.Context,
	request *models.UpdateBridgeDomainRequest,
) error {
	//TODO
	return nil
}

// DeleteBridgeDomain deletes a resource
func (db *DB) deleteBridgeDomain(
	ctx context.Context,
	request *models.DeleteBridgeDomainRequest) error {
	qb := db.queryBuilders["bridge_domain"]

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

//CreateBridgeDomain handle a Create API
// nolint
func (db *DB) CreateBridgeDomain(
	ctx context.Context,
	request *models.CreateBridgeDomainRequest) (*models.CreateBridgeDomainResponse, error) {
	model := request.BridgeDomain
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createBridgeDomain(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBridgeDomainResponse{
		BridgeDomain: request.BridgeDomain,
	}, nil
}

//UpdateBridgeDomain handles a Update request.
func (db *DB) UpdateBridgeDomain(
	ctx context.Context,
	request *models.UpdateBridgeDomainRequest) (*models.UpdateBridgeDomainResponse, error) {
	model := request.BridgeDomain
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateBridgeDomain(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBridgeDomainResponse{
		BridgeDomain: model,
	}, nil
}

//DeleteBridgeDomain delete a resource.
func (db *DB) DeleteBridgeDomain(ctx context.Context, request *models.DeleteBridgeDomainRequest) (*models.DeleteBridgeDomainResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteBridgeDomain(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBridgeDomainResponse{
		ID: request.ID,
	}, nil
}

//GetBridgeDomain a Get request.
func (db *DB) GetBridgeDomain(ctx context.Context, request *models.GetBridgeDomainRequest) (response *models.GetBridgeDomainResponse, err error) {
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
	listRequest := &models.ListBridgeDomainRequest{
		Spec: spec,
	}
	var result *models.ListBridgeDomainResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listBridgeDomain(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BridgeDomains) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBridgeDomainResponse{
		BridgeDomain: result.BridgeDomains[0],
	}
	return response, nil
}

//ListBridgeDomain handles a List service Request.
// nolint
func (db *DB) ListBridgeDomain(
	ctx context.Context,
	request *models.ListBridgeDomainRequest) (response *models.ListBridgeDomainResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listBridgeDomain(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
