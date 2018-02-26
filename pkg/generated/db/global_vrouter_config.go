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

const insertGlobalVrouterConfigQuery = "insert into `global_vrouter_config` (`vxlan_network_identifier_mode`,`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`linklocal_service_entry`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`forwarding_mode`,`flow_export_rate`,`flow_aging_timeout`,`encapsulation`,`enable_security_logging`,`source_port`,`source_ip`,`ip_protocol`,`hashing_configured`,`destination_port`,`destination_ip`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteGlobalVrouterConfigQuery = "delete from `global_vrouter_config` where uuid = ?"

// GlobalVrouterConfigFields is db columns for GlobalVrouterConfig
var GlobalVrouterConfigFields = []string{
	"vxlan_network_identifier_mode",
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"linklocal_service_entry",
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
	"forwarding_mode",
	"flow_export_rate",
	"flow_aging_timeout",
	"encapsulation",
	"enable_security_logging",
	"source_port",
	"source_ip",
	"ip_protocol",
	"hashing_configured",
	"destination_port",
	"destination_ip",
	"display_name",
	"key_value_pair",
}

// GlobalVrouterConfigRefFields is db reference fields for GlobalVrouterConfig
var GlobalVrouterConfigRefFields = map[string][]string{}

// GlobalVrouterConfigBackRefFields is db back reference fields for GlobalVrouterConfig
var GlobalVrouterConfigBackRefFields = map[string][]string{

	"security_logging_object": []string{
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
	},
}

// GlobalVrouterConfigParentTypes is possible parents for GlobalVrouterConfig
var GlobalVrouterConfigParents = []string{

	"global_system_config",
}

// CreateGlobalVrouterConfig inserts GlobalVrouterConfig to DB
func CreateGlobalVrouterConfig(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateGlobalVrouterConfigRequest) error {
	model := request.GlobalVrouterConfig
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalVrouterConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertGlobalVrouterConfigQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetVxlanNetworkIdentifierMode()),
		string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		common.MustJSON(model.GetLinklocalServices().GetLinklocalServiceEntry()),
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
		string(model.GetForwardingMode()),
		int(model.GetFlowExportRate()),
		common.MustJSON(model.GetFlowAgingTimeoutList().GetFlowAgingTimeout()),
		common.MustJSON(model.GetEncapsulationPriorities().GetEncapsulation()),
		bool(model.GetEnableSecurityLogging()),
		bool(model.GetEcmpHashingIncludeFields().GetSourcePort()),
		bool(model.GetEcmpHashingIncludeFields().GetSourceIP()),
		bool(model.GetEcmpHashingIncludeFields().GetIPProtocol()),
		bool(model.GetEcmpHashingIncludeFields().GetHashingConfigured()),
		bool(model.GetEcmpHashingIncludeFields().GetDestinationPort()),
		bool(model.GetEcmpHashingIncludeFields().GetDestinationIP()),
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "global_vrouter_config",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "global_vrouter_config", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanGlobalVrouterConfig(values map[string]interface{}) (*models.GlobalVrouterConfig, error) {
	m := models.MakeGlobalVrouterConfig()

	if value, ok := values["vxlan_network_identifier_mode"]; ok {

		m.VxlanNetworkIdentifierMode = schema.InterfaceToString(value)

	}

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

	if value, ok := values["linklocal_service_entry"]; ok {

		json.Unmarshal(value.([]byte), &m.LinklocalServices.LinklocalServiceEntry)

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

	if value, ok := values["forwarding_mode"]; ok {

		m.ForwardingMode = schema.InterfaceToString(value)

	}

	if value, ok := values["flow_export_rate"]; ok {

		m.FlowExportRate = schema.InterfaceToInt64(value)

	}

	if value, ok := values["flow_aging_timeout"]; ok {

		json.Unmarshal(value.([]byte), &m.FlowAgingTimeoutList.FlowAgingTimeout)

	}

	if value, ok := values["encapsulation"]; ok {

		json.Unmarshal(value.([]byte), &m.EncapsulationPriorities.Encapsulation)

	}

	if value, ok := values["enable_security_logging"]; ok {

		m.EnableSecurityLogging = schema.InterfaceToBool(value)

	}

	if value, ok := values["source_port"]; ok {

		m.EcmpHashingIncludeFields.SourcePort = schema.InterfaceToBool(value)

	}

	if value, ok := values["source_ip"]; ok {

		m.EcmpHashingIncludeFields.SourceIP = schema.InterfaceToBool(value)

	}

	if value, ok := values["ip_protocol"]; ok {

		m.EcmpHashingIncludeFields.IPProtocol = schema.InterfaceToBool(value)

	}

	if value, ok := values["hashing_configured"]; ok {

		m.EcmpHashingIncludeFields.HashingConfigured = schema.InterfaceToBool(value)

	}

	if value, ok := values["destination_port"]; ok {

		m.EcmpHashingIncludeFields.DestinationPort = schema.InterfaceToBool(value)

	}

	if value, ok := values["destination_ip"]; ok {

		m.EcmpHashingIncludeFields.DestinationIP = schema.InterfaceToBool(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["backref_security_logging_object"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeSecurityLoggingObject()
			m.SecurityLoggingObjects = append(m.SecurityLoggingObjects, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["rule"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.SecurityLoggingObjectRules.Rule)

			}

			if propertyValue, ok := childResourceMap["security_logging_object_rate"]; ok && propertyValue != nil {

				childModel.SecurityLoggingObjectRate = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListGlobalVrouterConfig lists GlobalVrouterConfig with list spec.
func ListGlobalVrouterConfig(ctx context.Context, tx *sql.Tx, request *models.ListGlobalVrouterConfigRequest) (response *models.ListGlobalVrouterConfigResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "global_vrouter_config"
	qb.Fields = GlobalVrouterConfigFields
	qb.RefFields = GlobalVrouterConfigRefFields
	qb.BackRefFields = GlobalVrouterConfigBackRefFields
	result := []*models.GlobalVrouterConfig{}

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
		m, err := scanGlobalVrouterConfig(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListGlobalVrouterConfigResponse{
		GlobalVrouterConfigs: result,
	}
	return response, nil
}

// UpdateGlobalVrouterConfig updates a resource
func UpdateGlobalVrouterConfig(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateGlobalVrouterConfigRequest,
) error {
	//TODO
	return nil
}

// DeleteGlobalVrouterConfig deletes a resource
func DeleteGlobalVrouterConfig(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteGlobalVrouterConfigRequest) error {
	deleteQuery := deleteGlobalVrouterConfigQuery
	selectQuery := "select count(uuid) from global_vrouter_config where uuid = ?"
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
