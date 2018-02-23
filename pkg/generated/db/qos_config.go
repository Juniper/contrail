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

const insertQosConfigQuery = "insert into `qos_config` (`qos_id_forwarding_class_pair`,`uuid`,`qos_config_type`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`mpls_exp_entries_qos_id_forwarding_class_pair`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`dscp_entries_qos_id_forwarding_class_pair`,`display_name`,`default_forwarding_class_id`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteQosConfigQuery = "delete from `qos_config` where uuid = ?"

// QosConfigFields is db columns for QosConfig
var QosConfigFields = []string{
	"qos_id_forwarding_class_pair",
	"uuid",
	"qos_config_type",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"mpls_exp_entries_qos_id_forwarding_class_pair",
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
	"dscp_entries_qos_id_forwarding_class_pair",
	"display_name",
	"default_forwarding_class_id",
	"key_value_pair",
}

// QosConfigRefFields is db reference fields for QosConfig
var QosConfigRefFields = map[string][]string{

	"global_system_config": []string{
		// <schema.Schema Value>

	},
}

// QosConfigBackRefFields is db back reference fields for QosConfig
var QosConfigBackRefFields = map[string][]string{}

// QosConfigParentTypes is possible parents for QosConfig
var QosConfigParents = []string{

	"global_qos_config",

	"project",
}

const insertQosConfigGlobalSystemConfigQuery = "insert into `ref_qos_config_global_system_config` (`from`, `to` ) values (?, ?);"

// CreateQosConfig inserts QosConfig to DB
func CreateQosConfig(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateQosConfigRequest) error {
	model := request.QosConfig
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertQosConfigQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, common.MustJSON(model.GetVlanPriorityEntries().GetQosIDForwardingClassPair()),
		string(model.GetUUID()),
		string(model.GetQosConfigType()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		common.MustJSON(model.GetMPLSExpEntries().GetQosIDForwardingClassPair()),
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
		common.MustJSON(model.GetDSCPEntries().GetQosIDForwardingClassPair()),
		string(model.GetDisplayName()),
		int(model.GetDefaultForwardingClassID()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtGlobalSystemConfigRef, err := tx.Prepare(insertQosConfigGlobalSystemConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing GlobalSystemConfigRefs create statement failed")
	}
	defer stmtGlobalSystemConfigRef.Close()
	for _, ref := range model.GlobalSystemConfigRefs {

		_, err = stmtGlobalSystemConfigRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "GlobalSystemConfigRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "qos_config",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "qos_config", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanQosConfig(values map[string]interface{}) (*models.QosConfig, error) {
	m := models.MakeQosConfig()

	if value, ok := values["qos_id_forwarding_class_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.VlanPriorityEntries.QosIDForwardingClassPair)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["qos_config_type"]; ok {

		m.QosConfigType = schema.InterfaceToString(value)

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

	if value, ok := values["mpls_exp_entries_qos_id_forwarding_class_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.MPLSExpEntries.QosIDForwardingClassPair)

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

	if value, ok := values["dscp_entries_qos_id_forwarding_class_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.DSCPEntries.QosIDForwardingClassPair)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["default_forwarding_class_id"]; ok {

		m.DefaultForwardingClassID = schema.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_global_system_config"]; ok {
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
			referenceModel := &models.QosConfigGlobalSystemConfigRef{}
			referenceModel.UUID = uuid
			m.GlobalSystemConfigRefs = append(m.GlobalSystemConfigRefs, referenceModel)

		}
	}

	return m, nil
}

// ListQosConfig lists QosConfig with list spec.
func ListQosConfig(ctx context.Context, tx *sql.Tx, request *models.ListQosConfigRequest) (response *models.ListQosConfigResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "qos_config"
	qb.Fields = QosConfigFields
	qb.RefFields = QosConfigRefFields
	qb.BackRefFields = QosConfigBackRefFields
	result := []*models.QosConfig{}

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
		m, err := scanQosConfig(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListQosConfigResponse{
		QosConfigs: result,
	}
	return response, nil
}

// UpdateQosConfig updates a resource
func UpdateQosConfig(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateQosConfigRequest,
) error {
	//TODO
	return nil
}

// DeleteQosConfig deletes a resource
func DeleteQosConfig(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteQosConfigRequest) error {
	deleteQuery := deleteQosConfigQuery
	selectQuery := "select count(uuid) from qos_config where uuid = ?"
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
