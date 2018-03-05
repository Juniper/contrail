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
	"configuration_version",
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
		"configuration_version",
		"key_value_pair",
	},
}

// GlobalVrouterConfigParentTypes is possible parents for GlobalVrouterConfig
var GlobalVrouterConfigParents = []string{

	"global_system_config",
}

// CreateGlobalVrouterConfig inserts GlobalVrouterConfig to DB
// nolint
func (db *DB) createGlobalVrouterConfig(
	ctx context.Context,
	request *models.CreateGlobalVrouterConfigRequest) error {
	qb := db.queryBuilders["global_vrouter_config"]
	tx := GetTransaction(ctx)
	model := request.GlobalVrouterConfig
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetVxlanNetworkIdentifierMode()),
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
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "global_vrouter_config",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "global_vrouter_config", model.UUID, model.GetPerms2().GetShare())
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

		m.VxlanNetworkIdentifierMode = common.InterfaceToString(value)

	}

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

	if value, ok := values["linklocal_service_entry"]; ok {

		json.Unmarshal(value.([]byte), &m.LinklocalServices.LinklocalServiceEntry)

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

	if value, ok := values["forwarding_mode"]; ok {

		m.ForwardingMode = common.InterfaceToString(value)

	}

	if value, ok := values["flow_export_rate"]; ok {

		m.FlowExportRate = common.InterfaceToInt64(value)

	}

	if value, ok := values["flow_aging_timeout"]; ok {

		json.Unmarshal(value.([]byte), &m.FlowAgingTimeoutList.FlowAgingTimeout)

	}

	if value, ok := values["encapsulation"]; ok {

		json.Unmarshal(value.([]byte), &m.EncapsulationPriorities.Encapsulation)

	}

	if value, ok := values["enable_security_logging"]; ok {

		m.EnableSecurityLogging = common.InterfaceToBool(value)

	}

	if value, ok := values["source_port"]; ok {

		m.EcmpHashingIncludeFields.SourcePort = common.InterfaceToBool(value)

	}

	if value, ok := values["source_ip"]; ok {

		m.EcmpHashingIncludeFields.SourceIP = common.InterfaceToBool(value)

	}

	if value, ok := values["ip_protocol"]; ok {

		m.EcmpHashingIncludeFields.IPProtocol = common.InterfaceToBool(value)

	}

	if value, ok := values["hashing_configured"]; ok {

		m.EcmpHashingIncludeFields.HashingConfigured = common.InterfaceToBool(value)

	}

	if value, ok := values["destination_port"]; ok {

		m.EcmpHashingIncludeFields.DestinationPort = common.InterfaceToBool(value)

	}

	if value, ok := values["destination_ip"]; ok {

		m.EcmpHashingIncludeFields.DestinationIP = common.InterfaceToBool(value)

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

	if value, ok := values["backref_security_logging_object"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeSecurityLoggingObject()
			m.SecurityLoggingObjects = append(m.SecurityLoggingObjects, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["rule"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.SecurityLoggingObjectRules.Rule)

			}

			if propertyValue, ok := childResourceMap["security_logging_object_rate"]; ok && propertyValue != nil {

				childModel.SecurityLoggingObjectRate = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListGlobalVrouterConfig lists GlobalVrouterConfig with list spec.
func (db *DB) listGlobalVrouterConfig(ctx context.Context, request *models.ListGlobalVrouterConfigRequest) (response *models.ListGlobalVrouterConfigResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["global_vrouter_config"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.GlobalVrouterConfig{}

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
func (db *DB) updateGlobalVrouterConfig(
	ctx context.Context,
	request *models.UpdateGlobalVrouterConfigRequest,
) error {
	//TODO
	return nil
}

// DeleteGlobalVrouterConfig deletes a resource
func (db *DB) deleteGlobalVrouterConfig(
	ctx context.Context,
	request *models.DeleteGlobalVrouterConfigRequest) error {
	qb := db.queryBuilders["global_vrouter_config"]

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

//CreateGlobalVrouterConfig handle a Create API
// nolint
func (db *DB) CreateGlobalVrouterConfig(
	ctx context.Context,
	request *models.CreateGlobalVrouterConfigRequest) (*models.CreateGlobalVrouterConfigResponse, error) {
	model := request.GlobalVrouterConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createGlobalVrouterConfig(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_vrouter_config",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateGlobalVrouterConfigResponse{
		GlobalVrouterConfig: request.GlobalVrouterConfig,
	}, nil
}

//UpdateGlobalVrouterConfig handles a Update request.
func (db *DB) UpdateGlobalVrouterConfig(
	ctx context.Context,
	request *models.UpdateGlobalVrouterConfigRequest) (*models.UpdateGlobalVrouterConfigResponse, error) {
	model := request.GlobalVrouterConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateGlobalVrouterConfig(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_vrouter_config",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateGlobalVrouterConfigResponse{
		GlobalVrouterConfig: model,
	}, nil
}

//DeleteGlobalVrouterConfig delete a resource.
func (db *DB) DeleteGlobalVrouterConfig(ctx context.Context, request *models.DeleteGlobalVrouterConfigRequest) (*models.DeleteGlobalVrouterConfigResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteGlobalVrouterConfig(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteGlobalVrouterConfigResponse{
		ID: request.ID,
	}, nil
}

//GetGlobalVrouterConfig a Get request.
func (db *DB) GetGlobalVrouterConfig(ctx context.Context, request *models.GetGlobalVrouterConfigRequest) (response *models.GetGlobalVrouterConfigResponse, err error) {
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
	listRequest := &models.ListGlobalVrouterConfigRequest{
		Spec: spec,
	}
	var result *models.ListGlobalVrouterConfigResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listGlobalVrouterConfig(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.GlobalVrouterConfigs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetGlobalVrouterConfigResponse{
		GlobalVrouterConfig: result.GlobalVrouterConfigs[0],
	}
	return response, nil
}

//ListGlobalVrouterConfig handles a List service Request.
// nolint
func (db *DB) ListGlobalVrouterConfig(
	ctx context.Context,
	request *models.ListGlobalVrouterConfigRequest) (response *models.ListGlobalVrouterConfigResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listGlobalVrouterConfig(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
