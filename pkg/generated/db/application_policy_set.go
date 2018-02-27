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

const insertApplicationPolicySetQuery = "insert into `application_policy_set` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`,`all_applications`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteApplicationPolicySetQuery = "delete from `application_policy_set` where uuid = ?"

// ApplicationPolicySetFields is db columns for ApplicationPolicySet
var ApplicationPolicySetFields = []string{
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
	"all_applications",
}

// ApplicationPolicySetRefFields is db reference fields for ApplicationPolicySet
var ApplicationPolicySetRefFields = map[string][]string{

	"firewall_policy": []string{
		// <schema.Schema Value>
		"sequence",
	},

	"global_vrouter_config": []string{
	// <schema.Schema Value>

	},
}

// ApplicationPolicySetBackRefFields is db back reference fields for ApplicationPolicySet
var ApplicationPolicySetBackRefFields = map[string][]string{}

// ApplicationPolicySetParentTypes is possible parents for ApplicationPolicySet
var ApplicationPolicySetParents = []string{

	"project",

	"policy_management",
}

const insertApplicationPolicySetGlobalVrouterConfigQuery = "insert into `ref_application_policy_set_global_vrouter_config` (`from`, `to` ) values (?, ?);"

const insertApplicationPolicySetFirewallPolicyQuery = "insert into `ref_application_policy_set_firewall_policy` (`from`, `to` ,`sequence`) values (?, ?,?);"

// CreateApplicationPolicySet inserts ApplicationPolicySet to DB
func CreateApplicationPolicySet(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateApplicationPolicySetRequest) error {
	model := request.ApplicationPolicySet
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertApplicationPolicySetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertApplicationPolicySetQuery,
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
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()),
		bool(model.GetAllApplications()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtFirewallPolicyRef, err := tx.Prepare(insertApplicationPolicySetFirewallPolicyQuery)
	if err != nil {
		return errors.Wrap(err, "preparing FirewallPolicyRefs create statement failed")
	}
	defer stmtFirewallPolicyRef.Close()
	for _, ref := range model.FirewallPolicyRefs {

		if ref.Attr == nil {
			ref.Attr = &models.FirewallSequence{}
		}

		_, err = stmtFirewallPolicyRef.ExecContext(ctx, model.UUID, ref.UUID, string(ref.Attr.GetSequence()))
		if err != nil {
			return errors.Wrap(err, "FirewallPolicyRefs create failed")
		}
	}

	stmtGlobalVrouterConfigRef, err := tx.Prepare(insertApplicationPolicySetGlobalVrouterConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing GlobalVrouterConfigRefs create statement failed")
	}
	defer stmtGlobalVrouterConfigRef.Close()
	for _, ref := range model.GlobalVrouterConfigRefs {

		_, err = stmtGlobalVrouterConfigRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "GlobalVrouterConfigRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "application_policy_set",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "application_policy_set", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanApplicationPolicySet(values map[string]interface{}) (*models.ApplicationPolicySet, error) {
	m := models.MakeApplicationPolicySet()

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

	if value, ok := values["all_applications"]; ok {

		m.AllApplications = schema.InterfaceToBool(value)

	}

	if value, ok := values["ref_firewall_policy"]; ok {
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
			referenceModel := &models.ApplicationPolicySetFirewallPolicyRef{}
			referenceModel.UUID = uuid
			m.FirewallPolicyRefs = append(m.FirewallPolicyRefs, referenceModel)

			attr := models.MakeFirewallSequence()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_global_vrouter_config"]; ok {
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
			referenceModel := &models.ApplicationPolicySetGlobalVrouterConfigRef{}
			referenceModel.UUID = uuid
			m.GlobalVrouterConfigRefs = append(m.GlobalVrouterConfigRefs, referenceModel)

		}
	}

	return m, nil
}

// ListApplicationPolicySet lists ApplicationPolicySet with list spec.
func ListApplicationPolicySet(ctx context.Context, tx *sql.Tx, request *models.ListApplicationPolicySetRequest) (response *models.ListApplicationPolicySetResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "application_policy_set"
	qb.Fields = ApplicationPolicySetFields
	qb.RefFields = ApplicationPolicySetRefFields
	qb.BackRefFields = ApplicationPolicySetBackRefFields
	result := []*models.ApplicationPolicySet{}

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
		m, err := scanApplicationPolicySet(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListApplicationPolicySetResponse{
		ApplicationPolicySets: result,
	}
	return response, nil
}

// UpdateApplicationPolicySet updates a resource
func UpdateApplicationPolicySet(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateApplicationPolicySetRequest,
) error {
	//TODO
	return nil
}

// DeleteApplicationPolicySet deletes a resource
func DeleteApplicationPolicySet(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteApplicationPolicySetRequest) error {
	deleteQuery := deleteApplicationPolicySetQuery
	selectQuery := "select count(uuid) from application_policy_set where uuid = ?"
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
