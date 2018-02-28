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

const insertNetworkPolicyQuery = "insert into `network_policy` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`policy_rule`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteNetworkPolicyQuery = "delete from `network_policy` where uuid = ?"

// NetworkPolicyFields is db columns for NetworkPolicy
var NetworkPolicyFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"policy_rule",
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

// NetworkPolicyRefFields is db reference fields for NetworkPolicy
var NetworkPolicyRefFields = map[string][]string{}

// NetworkPolicyBackRefFields is db back reference fields for NetworkPolicy
var NetworkPolicyBackRefFields = map[string][]string{}

// NetworkPolicyParentTypes is possible parents for NetworkPolicy
var NetworkPolicyParents = []string{

	"project",
}

// CreateNetworkPolicy inserts NetworkPolicy to DB
// nolint
func (db *DB) createNetworkPolicy(
	ctx context.Context,
	request *models.CreateNetworkPolicyRequest) error {
	tx := GetTransaction(ctx)
	model := request.NetworkPolicy
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkPolicyQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertNetworkPolicyQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		common.MustJSON(model.GetNetworkPolicyEntries().GetPolicyRule()),
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

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "network_policy",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "network_policy", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

// nolint
func scanNetworkPolicy(values map[string]interface{}) (*models.NetworkPolicy, error) {
	m := models.MakeNetworkPolicy()

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

	if value, ok := values["policy_rule"]; ok {

		json.Unmarshal(value.([]byte), &m.NetworkPolicyEntries.PolicyRule)

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

	return m, nil
}

// ListNetworkPolicy lists NetworkPolicy with list spec.
// nolint
func (db *DB) listNetworkPolicy(ctx context.Context, request *models.ListNetworkPolicyRequest) (response *models.ListNetworkPolicyResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)
	qb := &ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "network_policy"
	qb.Fields = NetworkPolicyFields
	qb.RefFields = NetworkPolicyRefFields
	qb.BackRefFields = NetworkPolicyBackRefFields
	result := []*models.NetworkPolicy{}

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
		m, err := scanNetworkPolicy(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListNetworkPolicyResponse{
		NetworkPolicys: result,
	}
	return response, nil
}

// UpdateNetworkPolicy updates a resource
// nolint
func (db *DB) updateNetworkPolicy(
	ctx context.Context,
	request *models.UpdateNetworkPolicyRequest,
) error {
	//TODO
	return nil
}

// DeleteNetworkPolicy deletes a resource
// nolint
func (db *DB) deleteNetworkPolicy(
	ctx context.Context,
	request *models.DeleteNetworkPolicyRequest) error {
	deleteQuery := deleteNetworkPolicyQuery
	selectQuery := "select count(uuid) from network_policy where uuid = ?"
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

//CreateNetworkPolicy handle a Create API
// nolint
func (db *DB) CreateNetworkPolicy(
	ctx context.Context,
	request *models.CreateNetworkPolicyRequest) (*models.CreateNetworkPolicyResponse, error) {
	model := request.NetworkPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createNetworkPolicy(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_policy",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateNetworkPolicyResponse{
		NetworkPolicy: request.NetworkPolicy,
	}, nil
}

//UpdateNetworkPolicy handles a Update request.
func (db *DB) UpdateNetworkPolicy(
	ctx context.Context,
	request *models.UpdateNetworkPolicyRequest) (*models.UpdateNetworkPolicyResponse, error) {
	model := request.NetworkPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateNetworkPolicy(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_policy",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateNetworkPolicyResponse{
		NetworkPolicy: model,
	}, nil
}

//DeleteNetworkPolicy delete a resource.
func (db *DB) DeleteNetworkPolicy(ctx context.Context, request *models.DeleteNetworkPolicyRequest) (*models.DeleteNetworkPolicyResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteNetworkPolicy(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteNetworkPolicyResponse{
		ID: request.ID,
	}, nil
}

//GetNetworkPolicy a Get request.
// nolint
func (db *DB) GetNetworkPolicy(ctx context.Context, request *models.GetNetworkPolicyRequest) (response *models.GetNetworkPolicyResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListNetworkPolicyRequest{
		Spec: spec,
	}
	var result *models.ListNetworkPolicyResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listNetworkPolicy(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.NetworkPolicys) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetNetworkPolicyResponse{
		NetworkPolicy: result.NetworkPolicys[0],
	}
	return response, nil
}

//ListNetworkPolicy handles a List service Request.
// nolint
func (db *DB) ListNetworkPolicy(
	ctx context.Context,
	request *models.ListNetworkPolicyRequest) (response *models.ListNetworkPolicyResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listNetworkPolicy(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
