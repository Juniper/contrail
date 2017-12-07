package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertGlobalSystemConfigQuery = "insert into `global_system_config` (`config_version`,`uuid`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`alarm_enable`,`mac_aging_time`,`user_defined_log_statistics`,`subnet`,`autonomous_system`,`mac_limit`,`mac_limit_action`,`display_name`,`port_end`,`port_start`,`plugin_property`,`ibgp_auto_mesh`,`bgp_always_compare_med`,`restart_time`,`long_lived_restart_time`,`graceful_restart_parameters_enable`,`end_of_rib_timeout`,`bgp_helper_enable`,`xmpp_helper_enable`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalSystemConfigQuery = "update `global_system_config` set `config_version` = ?,`uuid` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`mac_move_time_window` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`alarm_enable` = ?,`mac_aging_time` = ?,`user_defined_log_statistics` = ?,`subnet` = ?,`autonomous_system` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`display_name` = ?,`port_end` = ?,`port_start` = ?,`plugin_property` = ?,`ibgp_auto_mesh` = ?,`bgp_always_compare_med` = ?,`restart_time` = ?,`long_lived_restart_time` = ?,`graceful_restart_parameters_enable` = ?,`end_of_rib_timeout` = ?,`bgp_helper_enable` = ?,`xmpp_helper_enable` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`fq_name` = ?;"
const deleteGlobalSystemConfigQuery = "delete from `global_system_config` where uuid = ?"

// GlobalSystemConfigFields is db columns for GlobalSystemConfig
var GlobalSystemConfigFields = []string{
	"config_version",
	"uuid",
	"mac_move_limit",
	"mac_move_limit_action",
	"mac_move_time_window",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"alarm_enable",
	"mac_aging_time",
	"user_defined_log_statistics",
	"subnet",
	"autonomous_system",
	"mac_limit",
	"mac_limit_action",
	"display_name",
	"port_end",
	"port_start",
	"plugin_property",
	"ibgp_auto_mesh",
	"bgp_always_compare_med",
	"restart_time",
	"long_lived_restart_time",
	"graceful_restart_parameters_enable",
	"end_of_rib_timeout",
	"bgp_helper_enable",
	"xmpp_helper_enable",
	"key_value_pair",
	"global_access",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"fq_name",
}

// GlobalSystemConfigRefFields is db reference fields for GlobalSystemConfig
var GlobalSystemConfigRefFields = map[string][]string{

	"bgp_router": {
	// <common.Schema Value>

	},
}

const insertGlobalSystemConfigBGPRouterQuery = "insert into `ref_global_system_config_bgp_router` (`from`, `to` ) values (?, ?);"

// CreateGlobalSystemConfig inserts GlobalSystemConfig to DB
func CreateGlobalSystemConfig(tx *sql.Tx, model *models.GlobalSystemConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalSystemConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertGlobalSystemConfigQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.ConfigVersion),
		string(model.UUID),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.MacMoveControl.MacMoveTimeWindow),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		bool(model.AlarmEnable),
		int(model.MacAgingTime),
		common.MustJSON(model.UserDefinedLogStatistics),
		common.MustJSON(model.IPFabricSubnets.Subnet),
		int(model.AutonomousSystem),
		int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		string(model.DisplayName),
		int(model.BgpaasParameters.PortEnd),
		int(model.BgpaasParameters.PortStart),
		common.MustJSON(model.PluginTuning.PluginProperty),
		bool(model.IbgpAutoMesh),
		bool(model.BGPAlwaysCompareMed),
		int(model.GracefulRestartParameters.RestartTime),
		int(model.GracefulRestartParameters.LongLivedRestartTime),
		bool(model.GracefulRestartParameters.Enable),
		int(model.GracefulRestartParameters.EndOfRibTimeout),
		bool(model.GracefulRestartParameters.BGPHelperEnable),
		bool(model.GracefulRestartParameters.XMPPHelperEnable),
		common.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		common.MustJSON(model.FQName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtBGPRouterRef, err := tx.Prepare(insertGlobalSystemConfigBGPRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing BGPRouterRefs create statement failed")
	}
	defer stmtBGPRouterRef.Close()
	for _, ref := range model.BGPRouterRefs {

		_, err = stmtBGPRouterRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "BGPRouterRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanGlobalSystemConfig(values map[string]interface{}) (*models.GlobalSystemConfig, error) {
	m := models.MakeGlobalSystemConfig()

	if value, ok := values["config_version"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ConfigVersion = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["mac_move_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveLimit = castedValue

	}

	if value, ok := values["mac_move_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacMoveControl.MacMoveLimitAction = models.MACLimitExceedActionType(castedValue)

	}

	if value, ok := values["mac_move_time_window"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveTimeWindow = models.MACMoveTimeWindow(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["alarm_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.AlarmEnable = castedValue

	}

	if value, ok := values["mac_aging_time"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacAgingTime = models.MACAgingTime(castedValue)

	}

	if value, ok := values["user_defined_log_statistics"]; ok {

		json.Unmarshal(value.([]byte), &m.UserDefinedLogStatistics)

	}

	if value, ok := values["subnet"]; ok {

		json.Unmarshal(value.([]byte), &m.IPFabricSubnets.Subnet)

	}

	if value, ok := values["autonomous_system"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.AutonomousSystem = models.AutonomousSystemType(castedValue)

	}

	if value, ok := values["mac_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacLimitControl.MacLimit = castedValue

	}

	if value, ok := values["mac_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacLimitControl.MacLimitAction = models.MACLimitExceedActionType(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["port_end"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.BgpaasParameters.PortEnd = models.L4PortType(castedValue)

	}

	if value, ok := values["port_start"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.BgpaasParameters.PortStart = models.L4PortType(castedValue)

	}

	if value, ok := values["plugin_property"]; ok {

		json.Unmarshal(value.([]byte), &m.PluginTuning.PluginProperty)

	}

	if value, ok := values["ibgp_auto_mesh"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IbgpAutoMesh = castedValue

	}

	if value, ok := values["bgp_always_compare_med"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.BGPAlwaysCompareMed = castedValue

	}

	if value, ok := values["restart_time"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.GracefulRestartParameters.RestartTime = models.GracefulRestartTimeType(castedValue)

	}

	if value, ok := values["long_lived_restart_time"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.GracefulRestartParameters.LongLivedRestartTime = models.LongLivedGracefulRestartTimeType(castedValue)

	}

	if value, ok := values["graceful_restart_parameters_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.GracefulRestartParameters.Enable = castedValue

	}

	if value, ok := values["end_of_rib_timeout"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.GracefulRestartParameters.EndOfRibTimeout = models.EndOfRibTimeType(castedValue)

	}

	if value, ok := values["bgp_helper_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.GracefulRestartParameters.BGPHelperEnable = castedValue

	}

	if value, ok := values["xmpp_helper_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.GracefulRestartParameters.XMPPHelperEnable = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["ref_bgp_router"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.GlobalSystemConfigBGPRouterRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)

		}
	}

	return m, nil
}

// ListGlobalSystemConfig lists GlobalSystemConfig with list spec.
func ListGlobalSystemConfig(tx *sql.Tx, spec *common.ListSpec) ([]*models.GlobalSystemConfig, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "global_system_config"
	spec.Fields = GlobalSystemConfigFields
	spec.RefFields = GlobalSystemConfigRefFields
	result := models.MakeGlobalSystemConfigSlice()
	query, columns, values := common.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
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
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanGlobalSystemConfig(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowGlobalSystemConfig shows GlobalSystemConfig resource
func ShowGlobalSystemConfig(tx *sql.Tx, uuid string) (*models.GlobalSystemConfig, error) {
	list, err := ListGlobalSystemConfig(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateGlobalSystemConfig updates a resource
func UpdateGlobalSystemConfig(tx *sql.Tx, uuid string, model *models.GlobalSystemConfig) error {
	//TODO(nati) support update
	return nil
}

// DeleteGlobalSystemConfig deletes a resource
func DeleteGlobalSystemConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteGlobalSystemConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
