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

// DsaRuleFields is db columns for DsaRule
var DsaRuleFields = []string{
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
	"subscriber",
	"ep_version",
	"ep_type",
	"ip_prefix_len",
	"ip_prefix",
	"ep_id",
	"display_name",
	"configuration_version",
	"key_value_pair",
}

// DsaRuleRefFields is db reference fields for DsaRule
var DsaRuleRefFields = map[string][]string{}

// DsaRuleBackRefFields is db back reference fields for DsaRule
var DsaRuleBackRefFields = map[string][]string{}

// DsaRuleParentTypes is possible parents for DsaRule
var DsaRuleParents = []string{

	"discovery_service_assignment",
}

// CreateDsaRule inserts DsaRule to DB
// nolint
func (db *DB) createDsaRule(
	ctx context.Context,
	request *models.CreateDsaRuleRequest) error {
	qb := db.queryBuilders["dsa_rule"]
	tx := GetTransaction(ctx)
	model := request.DsaRule
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
		common.MustJSON(model.GetDsaRuleEntry().GetSubscriber()),
		string(model.GetDsaRuleEntry().GetPublisher().GetEpVersion()),
		string(model.GetDsaRuleEntry().GetPublisher().GetEpType()),
		int(model.GetDsaRuleEntry().GetPublisher().GetEpPrefix().GetIPPrefixLen()),
		string(model.GetDsaRuleEntry().GetPublisher().GetEpPrefix().GetIPPrefix()),
		string(model.GetDsaRuleEntry().GetPublisher().GetEpID()),
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
		Type:   "dsa_rule",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "dsa_rule", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanDsaRule(values map[string]interface{}) (*models.DsaRule, error) {
	m := models.MakeDsaRule()

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

	if value, ok := values["subscriber"]; ok {

		json.Unmarshal(value.([]byte), &m.DsaRuleEntry.Subscriber)

	}

	if value, ok := values["ep_version"]; ok {

		m.DsaRuleEntry.Publisher.EpVersion = common.InterfaceToString(value)

	}

	if value, ok := values["ep_type"]; ok {

		m.DsaRuleEntry.Publisher.EpType = common.InterfaceToString(value)

	}

	if value, ok := values["ip_prefix_len"]; ok {

		m.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen = common.InterfaceToInt64(value)

	}

	if value, ok := values["ip_prefix"]; ok {

		m.DsaRuleEntry.Publisher.EpPrefix.IPPrefix = common.InterfaceToString(value)

	}

	if value, ok := values["ep_id"]; ok {

		m.DsaRuleEntry.Publisher.EpID = common.InterfaceToString(value)

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

// ListDsaRule lists DsaRule with list spec.
func (db *DB) listDsaRule(ctx context.Context, request *models.ListDsaRuleRequest) (response *models.ListDsaRuleResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["dsa_rule"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.DsaRule{}

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
		m, err := scanDsaRule(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListDsaRuleResponse{
		DsaRules: result,
	}
	return response, nil
}

// UpdateDsaRule updates a resource
func (db *DB) updateDsaRule(
	ctx context.Context,
	request *models.UpdateDsaRuleRequest,
) error {
	//TODO
	return nil
}

// DeleteDsaRule deletes a resource
func (db *DB) deleteDsaRule(
	ctx context.Context,
	request *models.DeleteDsaRuleRequest) error {
	qb := db.queryBuilders["dsa_rule"]

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

//CreateDsaRule handle a Create API
// nolint
func (db *DB) CreateDsaRule(
	ctx context.Context,
	request *models.CreateDsaRuleRequest) (*models.CreateDsaRuleResponse, error) {
	model := request.DsaRule
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createDsaRule(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dsa_rule",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateDsaRuleResponse{
		DsaRule: request.DsaRule,
	}, nil
}

//UpdateDsaRule handles a Update request.
func (db *DB) UpdateDsaRule(
	ctx context.Context,
	request *models.UpdateDsaRuleRequest) (*models.UpdateDsaRuleResponse, error) {
	model := request.DsaRule
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateDsaRule(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dsa_rule",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateDsaRuleResponse{
		DsaRule: model,
	}, nil
}

//DeleteDsaRule delete a resource.
func (db *DB) DeleteDsaRule(ctx context.Context, request *models.DeleteDsaRuleRequest) (*models.DeleteDsaRuleResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteDsaRule(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteDsaRuleResponse{
		ID: request.ID,
	}, nil
}

//GetDsaRule a Get request.
func (db *DB) GetDsaRule(ctx context.Context, request *models.GetDsaRuleRequest) (response *models.GetDsaRuleResponse, err error) {
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
	listRequest := &models.ListDsaRuleRequest{
		Spec: spec,
	}
	var result *models.ListDsaRuleResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listDsaRule(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.DsaRules) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetDsaRuleResponse{
		DsaRule: result.DsaRules[0],
	}
	return response, nil
}

//ListDsaRule handles a List service Request.
// nolint
func (db *DB) ListDsaRule(
	ctx context.Context,
	request *models.ListDsaRuleRequest) (response *models.ListDsaRuleResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listDsaRule(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
