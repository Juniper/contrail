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

// SecurityLoggingObjectFields is db columns for SecurityLoggingObject
var SecurityLoggingObjectFields = []string{
	"uuid",
	"rule",
	"security_logging_object_rate",
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
	"configuration_version",
	"key_value_pair",
}

// SecurityLoggingObjectRefFields is db reference fields for SecurityLoggingObject
var SecurityLoggingObjectRefFields = map[string][]string{

	"security_group": []string{
		// <schema.Schema Value>
		"rule",
	},

	"network_policy": []string{
		// <schema.Schema Value>
		"rule",
	},
}

// SecurityLoggingObjectBackRefFields is db back reference fields for SecurityLoggingObject
var SecurityLoggingObjectBackRefFields = map[string][]string{}

// SecurityLoggingObjectParentTypes is possible parents for SecurityLoggingObject
var SecurityLoggingObjectParents = []string{

	"project",

	"global_vrouter_config",
}

// CreateSecurityLoggingObject inserts SecurityLoggingObject to DB
// nolint
func (db *DB) createSecurityLoggingObject(
	ctx context.Context,
	request *models.CreateSecurityLoggingObjectRequest) error {
	qb := db.queryBuilders["security_logging_object"]
	tx := GetTransaction(ctx)
	model := request.SecurityLoggingObject
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		common.MustJSON(model.GetSecurityLoggingObjectRules().GetRule()),
		int(model.GetSecurityLoggingObjectRate()),
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
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.NetworkPolicyRefs {

		if ref.Attr == nil {
			ref.Attr = &models.SecurityLoggingObjectRuleListType{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("network_policy"), model.UUID, ref.UUID, common.MustJSON(ref.Attr.GetRule()))
		if err != nil {
			return errors.Wrap(err, "NetworkPolicyRefs create failed")
		}
	}

	for _, ref := range model.SecurityGroupRefs {

		if ref.Attr == nil {
			ref.Attr = &models.SecurityLoggingObjectRuleListType{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("security_group"), model.UUID, ref.UUID, common.MustJSON(ref.Attr.GetRule()))
		if err != nil {
			return errors.Wrap(err, "SecurityGroupRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "security_logging_object",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "security_logging_object", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanSecurityLoggingObject(values map[string]interface{}) (*models.SecurityLoggingObject, error) {
	m := models.MakeSecurityLoggingObject()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["rule"]; ok {

		json.Unmarshal(value.([]byte), &m.SecurityLoggingObjectRules.Rule)

	}

	if value, ok := values["security_logging_object_rate"]; ok {

		m.SecurityLoggingObjectRate = common.InterfaceToInt64(value)

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

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_security_group"]; ok {
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
			referenceModel := &models.SecurityLoggingObjectSecurityGroupRef{}
			referenceModel.UUID = uuid
			m.SecurityGroupRefs = append(m.SecurityGroupRefs, referenceModel)

			attr := models.MakeSecurityLoggingObjectRuleListType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_network_policy"]; ok {
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
			referenceModel := &models.SecurityLoggingObjectNetworkPolicyRef{}
			referenceModel.UUID = uuid
			m.NetworkPolicyRefs = append(m.NetworkPolicyRefs, referenceModel)

			attr := models.MakeSecurityLoggingObjectRuleListType()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListSecurityLoggingObject lists SecurityLoggingObject with list spec.
func (db *DB) listSecurityLoggingObject(ctx context.Context, request *models.ListSecurityLoggingObjectRequest) (response *models.ListSecurityLoggingObjectResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["security_logging_object"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.SecurityLoggingObject{}

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
		m, err := scanSecurityLoggingObject(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListSecurityLoggingObjectResponse{
		SecurityLoggingObjects: result,
	}
	return response, nil
}

// UpdateSecurityLoggingObject updates a resource
func (db *DB) updateSecurityLoggingObject(
	ctx context.Context,
	request *models.UpdateSecurityLoggingObjectRequest,
) error {
	//TODO
	return nil
}

// DeleteSecurityLoggingObject deletes a resource
func (db *DB) deleteSecurityLoggingObject(
	ctx context.Context,
	request *models.DeleteSecurityLoggingObjectRequest) error {
	qb := db.queryBuilders["security_logging_object"]

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

//CreateSecurityLoggingObject handle a Create API
// nolint
func (db *DB) CreateSecurityLoggingObject(
	ctx context.Context,
	request *models.CreateSecurityLoggingObjectRequest) (*models.CreateSecurityLoggingObjectResponse, error) {
	model := request.SecurityLoggingObject
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createSecurityLoggingObject(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_logging_object",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateSecurityLoggingObjectResponse{
		SecurityLoggingObject: request.SecurityLoggingObject,
	}, nil
}

//UpdateSecurityLoggingObject handles a Update request.
func (db *DB) UpdateSecurityLoggingObject(
	ctx context.Context,
	request *models.UpdateSecurityLoggingObjectRequest) (*models.UpdateSecurityLoggingObjectResponse, error) {
	model := request.SecurityLoggingObject
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateSecurityLoggingObject(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_logging_object",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateSecurityLoggingObjectResponse{
		SecurityLoggingObject: model,
	}, nil
}

//DeleteSecurityLoggingObject delete a resource.
func (db *DB) DeleteSecurityLoggingObject(ctx context.Context, request *models.DeleteSecurityLoggingObjectRequest) (*models.DeleteSecurityLoggingObjectResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteSecurityLoggingObject(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteSecurityLoggingObjectResponse{
		ID: request.ID,
	}, nil
}

//GetSecurityLoggingObject a Get request.
func (db *DB) GetSecurityLoggingObject(ctx context.Context, request *models.GetSecurityLoggingObjectRequest) (response *models.GetSecurityLoggingObjectResponse, err error) {
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
	listRequest := &models.ListSecurityLoggingObjectRequest{
		Spec: spec,
	}
	var result *models.ListSecurityLoggingObjectResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listSecurityLoggingObject(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.SecurityLoggingObjects) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetSecurityLoggingObjectResponse{
		SecurityLoggingObject: result.SecurityLoggingObjects[0],
	}
	return response, nil
}

//ListSecurityLoggingObject handles a List service Request.
// nolint
func (db *DB) ListSecurityLoggingObject(
	ctx context.Context,
	request *models.ListSecurityLoggingObjectRequest) (response *models.ListSecurityLoggingObjectResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listSecurityLoggingObject(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
