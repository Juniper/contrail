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

const insertSecurityLoggingObjectQuery = "insert into `security_logging_object` (`uuid`,`rule`,`security_logging_object_rate`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteSecurityLoggingObjectQuery = "delete from `security_logging_object` where uuid = ?"

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

const insertSecurityLoggingObjectNetworkPolicyQuery = "insert into `ref_security_logging_object_network_policy` (`from`, `to` ,`rule`) values (?, ?,?);"

const insertSecurityLoggingObjectSecurityGroupQuery = "insert into `ref_security_logging_object_security_group` (`from`, `to` ,`rule`) values (?, ?,?);"

// CreateSecurityLoggingObject inserts SecurityLoggingObject to DB
func CreateSecurityLoggingObject(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateSecurityLoggingObjectRequest) error {
	model := request.SecurityLoggingObject
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSecurityLoggingObjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertSecurityLoggingObjectQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
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
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtNetworkPolicyRef, err := tx.Prepare(insertSecurityLoggingObjectNetworkPolicyQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NetworkPolicyRefs create statement failed")
	}
	defer stmtNetworkPolicyRef.Close()
	for _, ref := range model.NetworkPolicyRefs {

		if ref.Attr == nil {
			ref.Attr = &models.SecurityLoggingObjectRuleListType{}
		}

		_, err = stmtNetworkPolicyRef.ExecContext(ctx, model.UUID, ref.UUID, common.MustJSON(ref.Attr.GetRule()))
		if err != nil {
			return errors.Wrap(err, "NetworkPolicyRefs create failed")
		}
	}

	stmtSecurityGroupRef, err := tx.Prepare(insertSecurityLoggingObjectSecurityGroupQuery)
	if err != nil {
		return errors.Wrap(err, "preparing SecurityGroupRefs create statement failed")
	}
	defer stmtSecurityGroupRef.Close()
	for _, ref := range model.SecurityGroupRefs {

		if ref.Attr == nil {
			ref.Attr = &models.SecurityLoggingObjectRuleListType{}
		}

		_, err = stmtSecurityGroupRef.ExecContext(ctx, model.UUID, ref.UUID, common.MustJSON(ref.Attr.GetRule()))
		if err != nil {
			return errors.Wrap(err, "SecurityGroupRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "security_logging_object",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "security_logging_object", model.UUID, model.GetPerms2().GetShare())
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

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["rule"]; ok {

		json.Unmarshal(value.([]byte), &m.SecurityLoggingObjectRules.Rule)

	}

	if value, ok := values["security_logging_object_rate"]; ok {

		m.SecurityLoggingObjectRate = schema.InterfaceToInt64(value)

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

	if value, ok := values["ref_security_group"]; ok {
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
			referenceModel := &models.SecurityLoggingObjectSecurityGroupRef{}
			referenceModel.UUID = uuid
			m.SecurityGroupRefs = append(m.SecurityGroupRefs, referenceModel)

			attr := models.MakeSecurityLoggingObjectRuleListType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_network_policy"]; ok {
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
func ListSecurityLoggingObject(ctx context.Context, tx *sql.Tx, request *models.ListSecurityLoggingObjectRequest) (response *models.ListSecurityLoggingObjectResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "security_logging_object"
	qb.Fields = SecurityLoggingObjectFields
	qb.RefFields = SecurityLoggingObjectRefFields
	qb.BackRefFields = SecurityLoggingObjectBackRefFields
	result := []*models.SecurityLoggingObject{}

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
func UpdateSecurityLoggingObject(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateSecurityLoggingObjectRequest,
) error {
	//TODO
	return nil
}

// DeleteSecurityLoggingObject deletes a resource
func DeleteSecurityLoggingObject(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteSecurityLoggingObjectRequest) error {
	deleteQuery := deleteSecurityLoggingObjectQuery
	selectQuery := "select count(uuid) from security_logging_object where uuid = ?"
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
