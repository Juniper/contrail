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

const insertConfigRootQuery = "insert into `config_root` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteConfigRootQuery = "delete from `config_root` where uuid = ?"

// ConfigRootFields is db columns for ConfigRoot
var ConfigRootFields = []string{
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

// ConfigRootRefFields is db reference fields for ConfigRoot
var ConfigRootRefFields = map[string][]string{

	"tag": []string{
	// <schema.Schema Value>

	},
}

// ConfigRootBackRefFields is db back reference fields for ConfigRoot
var ConfigRootBackRefFields = map[string][]string{

	"domain": []string{
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
		"virtual_network_limit",
		"security_group_limit",
		"project_limit",
		"display_name",
		"key_value_pair",
	},

	"global_system_config": []string{
		"uuid",
		"statlist",
		"plugin_property",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"mac_move_time_window",
		"mac_move_limit_action",
		"mac_move_limit",
		"mac_limit_action",
		"mac_limit",
		"mac_aging_time",
		"subnet",
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
		"ibgp_auto_mesh",
		"xmpp_helper_enable",
		"restart_time",
		"long_lived_restart_time",
		"end_of_rib_timeout",
		"graceful_restart_parameters_enable",
		"bgp_helper_enable",
		"fq_name",
		"display_name",
		"config_version",
		"port_start",
		"port_end",
		"bgp_always_compare_med",
		"autonomous_system",
		"key_value_pair",
		"alarm_enable",
	},

	"tag": []string{
		"uuid",
		"tag_value",
		"tag_type_name",
		"tag_id",
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

// ConfigRootParentTypes is possible parents for ConfigRoot
var ConfigRootParents = []string{}

const insertConfigRootTagQuery = "insert into `ref_config_root_tag` (`from`, `to` ) values (?, ?);"

// CreateConfigRoot inserts ConfigRoot to DB
func (db *DB) createConfigRoot(
	ctx context.Context,
	request *models.CreateConfigRootRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.ConfigRoot
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertConfigRootQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertConfigRootQuery,
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

	stmtTagRef, err := tx.Prepare(insertConfigRootTagQuery)
	if err != nil {
		return errors.Wrap(err, "preparing TagRefs create statement failed")
	}
	defer stmtTagRef.Close()
	for _, ref := range model.TagRefs {

		_, err = stmtTagRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "TagRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "config_root",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "config_root", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanConfigRoot(values map[string]interface{}) (*models.ConfigRoot, error) {
	m := models.MakeConfigRoot()

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

	if value, ok := values["ref_tag"]; ok {
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
			referenceModel := &models.ConfigRootTagRef{}
			referenceModel.UUID = uuid
			m.TagRefs = append(m.TagRefs, referenceModel)

		}
	}

	if value, ok := values["backref_domain"]; ok {
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
			childModel := models.MakeDomain()
			m.Domains = append(m.Domains, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

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

			if propertyValue, ok := childResourceMap["virtual_network_limit"]; ok && propertyValue != nil {

				childModel.DomainLimits.VirtualNetworkLimit = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["security_group_limit"]; ok && propertyValue != nil {

				childModel.DomainLimits.SecurityGroupLimit = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["project_limit"]; ok && propertyValue != nil {

				childModel.DomainLimits.ProjectLimit = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_global_system_config"]; ok {
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
			childModel := models.MakeGlobalSystemConfig()
			m.GlobalSystemConfigs = append(m.GlobalSystemConfigs, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["statlist"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.UserDefinedLogStatistics.Statlist)

			}

			if propertyValue, ok := childResourceMap["plugin_property"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.PluginTuning.PluginProperty)

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

			if propertyValue, ok := childResourceMap["mac_move_time_window"]; ok && propertyValue != nil {

				childModel.MacMoveControl.MacMoveTimeWindow = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_move_limit_action"]; ok && propertyValue != nil {

				childModel.MacMoveControl.MacMoveLimitAction = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_move_limit"]; ok && propertyValue != nil {

				childModel.MacMoveControl.MacMoveLimit = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_limit_action"]; ok && propertyValue != nil {

				childModel.MacLimitControl.MacLimitAction = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_limit"]; ok && propertyValue != nil {

				childModel.MacLimitControl.MacLimit = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_aging_time"]; ok && propertyValue != nil {

				childModel.MacAgingTime = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["subnet"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.IPFabricSubnets.Subnet)

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

			if propertyValue, ok := childResourceMap["ibgp_auto_mesh"]; ok && propertyValue != nil {

				childModel.IbgpAutoMesh = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["xmpp_helper_enable"]; ok && propertyValue != nil {

				childModel.GracefulRestartParameters.XMPPHelperEnable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["restart_time"]; ok && propertyValue != nil {

				childModel.GracefulRestartParameters.RestartTime = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["long_lived_restart_time"]; ok && propertyValue != nil {

				childModel.GracefulRestartParameters.LongLivedRestartTime = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["end_of_rib_timeout"]; ok && propertyValue != nil {

				childModel.GracefulRestartParameters.EndOfRibTimeout = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["graceful_restart_parameters_enable"]; ok && propertyValue != nil {

				childModel.GracefulRestartParameters.Enable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["bgp_helper_enable"]; ok && propertyValue != nil {

				childModel.GracefulRestartParameters.BGPHelperEnable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["config_version"]; ok && propertyValue != nil {

				childModel.ConfigVersion = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["port_start"]; ok && propertyValue != nil {

				childModel.BgpaasParameters.PortStart = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["port_end"]; ok && propertyValue != nil {

				childModel.BgpaasParameters.PortEnd = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["bgp_always_compare_med"]; ok && propertyValue != nil {

				childModel.BGPAlwaysCompareMed = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["autonomous_system"]; ok && propertyValue != nil {

				childModel.AutonomousSystem = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

			if propertyValue, ok := childResourceMap["alarm_enable"]; ok && propertyValue != nil {

				childModel.AlarmEnable = schema.InterfaceToBool(propertyValue)

			}

		}
	}

	if value, ok := values["backref_tag"]; ok {
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
			childModel := models.MakeTag()
			m.Tags = append(m.Tags, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["tag_value"]; ok && propertyValue != nil {

				childModel.TagValue = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["tag_type_name"]; ok && propertyValue != nil {

				childModel.TagTypeName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["tag_id"]; ok && propertyValue != nil {

				childModel.TagID = schema.InterfaceToString(propertyValue)

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

// ListConfigRoot lists ConfigRoot with list spec.
func (db *DB) listConfigRoot(ctx context.Context, request *models.ListConfigRootRequest) (response *models.ListConfigRootResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "config_root"
	qb.Fields = ConfigRootFields
	qb.RefFields = ConfigRootRefFields
	qb.BackRefFields = ConfigRootBackRefFields
	result := []*models.ConfigRoot{}

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
		m, err := scanConfigRoot(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListConfigRootResponse{
		ConfigRoots: result,
	}
	return response, nil
}

// UpdateConfigRoot updates a resource
func (db *DB) updateConfigRoot(
	ctx context.Context,
	request *models.UpdateConfigRootRequest,
) error {
	//TODO
	return nil
}

// DeleteConfigRoot deletes a resource
func (db *DB) deleteConfigRoot(
	ctx context.Context,
	request *models.DeleteConfigRootRequest) error {
	deleteQuery := deleteConfigRootQuery
	selectQuery := "select count(uuid) from config_root where uuid = ?"
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

//CreateConfigRoot handle a Create API
func (db *DB) CreateConfigRoot(
	ctx context.Context,
	request *models.CreateConfigRootRequest) (*models.CreateConfigRootResponse, error) {
	model := request.ConfigRoot
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createConfigRoot(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_root",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateConfigRootResponse{
		ConfigRoot: request.ConfigRoot,
	}, nil
}

//UpdateConfigRoot handles a Update request.
func (db *DB) UpdateConfigRoot(
	ctx context.Context,
	request *models.UpdateConfigRootRequest) (*models.UpdateConfigRootResponse, error) {
	model := request.ConfigRoot
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateConfigRoot(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_root",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateConfigRootResponse{
		ConfigRoot: model,
	}, nil
}

//DeleteConfigRoot delete a resource.
func (db *DB) DeleteConfigRoot(ctx context.Context, request *models.DeleteConfigRootRequest) (*models.DeleteConfigRootResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteConfigRoot(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteConfigRootResponse{
		ID: request.ID,
	}, nil
}

//GetConfigRoot a Get request.
func (db *DB) GetConfigRoot(ctx context.Context, request *models.GetConfigRootRequest) (response *models.GetConfigRootResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListConfigRootRequest{
		Spec: spec,
	}
	var result *models.ListConfigRootResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listConfigRoot(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ConfigRoots) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetConfigRootResponse{
		ConfigRoot: result.ConfigRoots[0],
	}
	return response, nil
}

//ListConfigRoot handles a List service Request.
func (db *DB) ListConfigRoot(
	ctx context.Context,
	request *models.ListConfigRootRequest) (response *models.ListConfigRootResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listConfigRoot(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
