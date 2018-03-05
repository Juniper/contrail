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

// LoadbalancerMemberFields is db columns for LoadbalancerMember
var LoadbalancerMemberFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"weight",
	"status_description",
	"status",
	"protocol_port",
	"admin_state",
	"address",
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

// LoadbalancerMemberRefFields is db reference fields for LoadbalancerMember
var LoadbalancerMemberRefFields = map[string][]string{}

// LoadbalancerMemberBackRefFields is db back reference fields for LoadbalancerMember
var LoadbalancerMemberBackRefFields = map[string][]string{}

// LoadbalancerMemberParentTypes is possible parents for LoadbalancerMember
var LoadbalancerMemberParents = []string{

	"loadbalancer_pool",
}

// CreateLoadbalancerMember inserts LoadbalancerMember to DB
// nolint
func (db *DB) createLoadbalancerMember(
	ctx context.Context,
	request *models.CreateLoadbalancerMemberRequest) error {
	qb := db.queryBuilders["loadbalancer_member"]
	tx := GetTransaction(ctx)
	model := request.LoadbalancerMember
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		int(model.GetLoadbalancerMemberProperties().GetWeight()),
		string(model.GetLoadbalancerMemberProperties().GetStatusDescription()),
		string(model.GetLoadbalancerMemberProperties().GetStatus()),
		int(model.GetLoadbalancerMemberProperties().GetProtocolPort()),
		bool(model.GetLoadbalancerMemberProperties().GetAdminState()),
		string(model.GetLoadbalancerMemberProperties().GetAddress()),
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
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "loadbalancer_member",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "loadbalancer_member", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanLoadbalancerMember(values map[string]interface{}) (*models.LoadbalancerMember, error) {
	m := models.MakeLoadbalancerMember()

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

	if value, ok := values["weight"]; ok {

		m.LoadbalancerMemberProperties.Weight = common.InterfaceToInt64(value)

	}

	if value, ok := values["status_description"]; ok {

		m.LoadbalancerMemberProperties.StatusDescription = common.InterfaceToString(value)

	}

	if value, ok := values["status"]; ok {

		m.LoadbalancerMemberProperties.Status = common.InterfaceToString(value)

	}

	if value, ok := values["protocol_port"]; ok {

		m.LoadbalancerMemberProperties.ProtocolPort = common.InterfaceToInt64(value)

	}

	if value, ok := values["admin_state"]; ok {

		m.LoadbalancerMemberProperties.AdminState = common.InterfaceToBool(value)

	}

	if value, ok := values["address"]; ok {

		m.LoadbalancerMemberProperties.Address = common.InterfaceToString(value)

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

// ListLoadbalancerMember lists LoadbalancerMember with list spec.
func (db *DB) listLoadbalancerMember(ctx context.Context, request *models.ListLoadbalancerMemberRequest) (response *models.ListLoadbalancerMemberResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["loadbalancer_member"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.LoadbalancerMember{}

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
		m, err := scanLoadbalancerMember(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListLoadbalancerMemberResponse{
		LoadbalancerMembers: result,
	}
	return response, nil
}

// UpdateLoadbalancerMember updates a resource
func (db *DB) updateLoadbalancerMember(
	ctx context.Context,
	request *models.UpdateLoadbalancerMemberRequest,
) error {
	//TODO
	return nil
}

// DeleteLoadbalancerMember deletes a resource
func (db *DB) deleteLoadbalancerMember(
	ctx context.Context,
	request *models.DeleteLoadbalancerMemberRequest) error {
	qb := db.queryBuilders["loadbalancer_member"]

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

//CreateLoadbalancerMember handle a Create API
// nolint
func (db *DB) CreateLoadbalancerMember(
	ctx context.Context,
	request *models.CreateLoadbalancerMemberRequest) (*models.CreateLoadbalancerMemberResponse, error) {
	model := request.LoadbalancerMember
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createLoadbalancerMember(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_member",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLoadbalancerMemberResponse{
		LoadbalancerMember: request.LoadbalancerMember,
	}, nil
}

//UpdateLoadbalancerMember handles a Update request.
func (db *DB) UpdateLoadbalancerMember(
	ctx context.Context,
	request *models.UpdateLoadbalancerMemberRequest) (*models.UpdateLoadbalancerMemberResponse, error) {
	model := request.LoadbalancerMember
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateLoadbalancerMember(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_member",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLoadbalancerMemberResponse{
		LoadbalancerMember: model,
	}, nil
}

//DeleteLoadbalancerMember delete a resource.
func (db *DB) DeleteLoadbalancerMember(ctx context.Context, request *models.DeleteLoadbalancerMemberRequest) (*models.DeleteLoadbalancerMemberResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteLoadbalancerMember(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLoadbalancerMemberResponse{
		ID: request.ID,
	}, nil
}

//GetLoadbalancerMember a Get request.
func (db *DB) GetLoadbalancerMember(ctx context.Context, request *models.GetLoadbalancerMemberRequest) (response *models.GetLoadbalancerMemberResponse, err error) {
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
	listRequest := &models.ListLoadbalancerMemberRequest{
		Spec: spec,
	}
	var result *models.ListLoadbalancerMemberResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listLoadbalancerMember(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LoadbalancerMembers) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLoadbalancerMemberResponse{
		LoadbalancerMember: result.LoadbalancerMembers[0],
	}
	return response, nil
}

//ListLoadbalancerMember handles a List service Request.
// nolint
func (db *DB) ListLoadbalancerMember(
	ctx context.Context,
	request *models.ListLoadbalancerMemberRequest) (response *models.ListLoadbalancerMemberResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listLoadbalancerMember(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
