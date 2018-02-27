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

const insertRoutingPolicyQuery = "insert into `routing_policy` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteRoutingPolicyQuery = "delete from `routing_policy` where uuid = ?"

// RoutingPolicyFields is db columns for RoutingPolicy
var RoutingPolicyFields = []string{
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

// RoutingPolicyRefFields is db reference fields for RoutingPolicy
var RoutingPolicyRefFields = map[string][]string{

	"service_instance": []string{
		// <schema.Schema Value>
		"right_sequence",
		"left_sequence",
	},
}

// RoutingPolicyBackRefFields is db back reference fields for RoutingPolicy
var RoutingPolicyBackRefFields = map[string][]string{}

// RoutingPolicyParentTypes is possible parents for RoutingPolicy
var RoutingPolicyParents = []string{

	"project",
}

const insertRoutingPolicyServiceInstanceQuery = "insert into `ref_routing_policy_service_instance` (`from`, `to` ,`right_sequence`,`left_sequence`) values (?, ?,?,?);"

// CreateRoutingPolicy inserts RoutingPolicy to DB
func (db *DB) createRoutingPolicy(
	ctx context.Context,
	request *models.CreateRoutingPolicyRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.RoutingPolicy
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRoutingPolicyQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertRoutingPolicyQuery,
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

	stmtServiceInstanceRef, err := tx.Prepare(insertRoutingPolicyServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceInstanceRefs create statement failed")
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {

		if ref.Attr == nil {
			ref.Attr = &models.RoutingPolicyServiceInstanceType{}
		}

		_, err = stmtServiceInstanceRef.ExecContext(ctx, model.UUID, ref.UUID, string(ref.Attr.GetRightSequence()),
			string(ref.Attr.GetLeftSequence()))
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "routing_policy",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "routing_policy", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanRoutingPolicy(values map[string]interface{}) (*models.RoutingPolicy, error) {
	m := models.MakeRoutingPolicy()

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

	if value, ok := values["ref_service_instance"]; ok {
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
			referenceModel := &models.RoutingPolicyServiceInstanceRef{}
			referenceModel.UUID = uuid
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

			attr := models.MakeRoutingPolicyServiceInstanceType()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListRoutingPolicy lists RoutingPolicy with list spec.
func (db *DB) listRoutingPolicy(ctx context.Context, request *models.ListRoutingPolicyRequest) (response *models.ListRoutingPolicyResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "routing_policy"
	qb.Fields = RoutingPolicyFields
	qb.RefFields = RoutingPolicyRefFields
	qb.BackRefFields = RoutingPolicyBackRefFields
	result := []*models.RoutingPolicy{}

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
		m, err := scanRoutingPolicy(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListRoutingPolicyResponse{
		RoutingPolicys: result,
	}
	return response, nil
}

// UpdateRoutingPolicy updates a resource
func (db *DB) updateRoutingPolicy(
	ctx context.Context,
	request *models.UpdateRoutingPolicyRequest,
) error {
	//TODO
	return nil
}

// DeleteRoutingPolicy deletes a resource
func (db *DB) deleteRoutingPolicy(
	ctx context.Context,
	request *models.DeleteRoutingPolicyRequest) error {
	deleteQuery := deleteRoutingPolicyQuery
	selectQuery := "select count(uuid) from routing_policy where uuid = ?"
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

//CreateRoutingPolicy handle a Create API
func (db *DB) CreateRoutingPolicy(
	ctx context.Context,
	request *models.CreateRoutingPolicyRequest) (*models.CreateRoutingPolicyResponse, error) {
	model := request.RoutingPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createRoutingPolicy(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateRoutingPolicyResponse{
		RoutingPolicy: request.RoutingPolicy,
	}, nil
}

//UpdateRoutingPolicy handles a Update request.
func (db *DB) UpdateRoutingPolicy(
	ctx context.Context,
	request *models.UpdateRoutingPolicyRequest) (*models.UpdateRoutingPolicyResponse, error) {
	model := request.RoutingPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateRoutingPolicy(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateRoutingPolicyResponse{
		RoutingPolicy: model,
	}, nil
}

//DeleteRoutingPolicy delete a resource.
func (db *DB) DeleteRoutingPolicy(ctx context.Context, request *models.DeleteRoutingPolicyRequest) (*models.DeleteRoutingPolicyResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteRoutingPolicy(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteRoutingPolicyResponse{
		ID: request.ID,
	}, nil
}

//GetRoutingPolicy a Get request.
func (db *DB) GetRoutingPolicy(ctx context.Context, request *models.GetRoutingPolicyRequest) (response *models.GetRoutingPolicyResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListRoutingPolicyRequest{
		Spec: spec,
	}
	var result *models.ListRoutingPolicyResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listRoutingPolicy(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.RoutingPolicys) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetRoutingPolicyResponse{
		RoutingPolicy: result.RoutingPolicys[0],
	}
	return response, nil
}

//ListRoutingPolicy handles a List service Request.
func (db *DB) ListRoutingPolicy(
	ctx context.Context,
	request *models.ListRoutingPolicyRequest) (response *models.ListRoutingPolicyResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listRoutingPolicy(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
