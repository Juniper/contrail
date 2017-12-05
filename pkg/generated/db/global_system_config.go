package db

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"strings"
)

const insertGlobalSystemConfigQuery = "insert into `global_system_config` (`fq_name`,`alarm_enable`,`ibgp_auto_mesh`,`bgp_always_compare_med`,`mac_limit`,`mac_limit_action`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`config_version`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`user_defined_log_statistics`,`display_name`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`plugin_property`,`subnet`,`key_value_pair`,`port_start`,`port_end`,`mac_aging_time`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`graceful_restart_parameters_enable`,`end_of_rib_timeout`,`bgp_helper_enable`,`autonomous_system`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalSystemConfigQuery = "update `global_system_config` set `fq_name` = ?,`alarm_enable` = ?,`ibgp_auto_mesh` = ?,`bgp_always_compare_med` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`config_version` = ?,`mac_move_time_window` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`user_defined_log_statistics` = ?,`display_name` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`plugin_property` = ?,`subnet` = ?,`key_value_pair` = ?,`port_start` = ?,`port_end` = ?,`mac_aging_time` = ?,`xmpp_helper_enable` = ?,`restart_time` = ?,`long_lived_restart_time` = ?,`graceful_restart_parameters_enable` = ?,`end_of_rib_timeout` = ?,`bgp_helper_enable` = ?,`autonomous_system` = ?,`uuid` = ?;"
const deleteGlobalSystemConfigQuery = "delete from `global_system_config` where uuid = ?"
const listGlobalSystemConfigQuery = "select `global_system_config`.`fq_name`,`global_system_config`.`alarm_enable`,`global_system_config`.`ibgp_auto_mesh`,`global_system_config`.`bgp_always_compare_med`,`global_system_config`.`mac_limit`,`global_system_config`.`mac_limit_action`,`global_system_config`.`created`,`global_system_config`.`creator`,`global_system_config`.`user_visible`,`global_system_config`.`last_modified`,`global_system_config`.`other_access`,`global_system_config`.`group`,`global_system_config`.`group_access`,`global_system_config`.`owner`,`global_system_config`.`owner_access`,`global_system_config`.`enable`,`global_system_config`.`description`,`global_system_config`.`config_version`,`global_system_config`.`mac_move_time_window`,`global_system_config`.`mac_move_limit`,`global_system_config`.`mac_move_limit_action`,`global_system_config`.`user_defined_log_statistics`,`global_system_config`.`display_name`,`global_system_config`.`share`,`global_system_config`.`perms2_owner`,`global_system_config`.`perms2_owner_access`,`global_system_config`.`global_access`,`global_system_config`.`plugin_property`,`global_system_config`.`subnet`,`global_system_config`.`key_value_pair`,`global_system_config`.`port_start`,`global_system_config`.`port_end`,`global_system_config`.`mac_aging_time`,`global_system_config`.`xmpp_helper_enable`,`global_system_config`.`restart_time`,`global_system_config`.`long_lived_restart_time`,`global_system_config`.`graceful_restart_parameters_enable`,`global_system_config`.`end_of_rib_timeout`,`global_system_config`.`bgp_helper_enable`,`global_system_config`.`autonomous_system`,`global_system_config`.`uuid` from `global_system_config`"
const showGlobalSystemConfigQuery = "select `global_system_config`.`fq_name`,`global_system_config`.`alarm_enable`,`global_system_config`.`ibgp_auto_mesh`,`global_system_config`.`bgp_always_compare_med`,`global_system_config`.`mac_limit`,`global_system_config`.`mac_limit_action`,`global_system_config`.`created`,`global_system_config`.`creator`,`global_system_config`.`user_visible`,`global_system_config`.`last_modified`,`global_system_config`.`other_access`,`global_system_config`.`group`,`global_system_config`.`group_access`,`global_system_config`.`owner`,`global_system_config`.`owner_access`,`global_system_config`.`enable`,`global_system_config`.`description`,`global_system_config`.`config_version`,`global_system_config`.`mac_move_time_window`,`global_system_config`.`mac_move_limit`,`global_system_config`.`mac_move_limit_action`,`global_system_config`.`user_defined_log_statistics`,`global_system_config`.`display_name`,`global_system_config`.`share`,`global_system_config`.`perms2_owner`,`global_system_config`.`perms2_owner_access`,`global_system_config`.`global_access`,`global_system_config`.`plugin_property`,`global_system_config`.`subnet`,`global_system_config`.`key_value_pair`,`global_system_config`.`port_start`,`global_system_config`.`port_end`,`global_system_config`.`mac_aging_time`,`global_system_config`.`xmpp_helper_enable`,`global_system_config`.`restart_time`,`global_system_config`.`long_lived_restart_time`,`global_system_config`.`graceful_restart_parameters_enable`,`global_system_config`.`end_of_rib_timeout`,`global_system_config`.`bgp_helper_enable`,`global_system_config`.`autonomous_system`,`global_system_config`.`uuid` from `global_system_config` where uuid = ?"

const insertGlobalSystemConfigBGPRouterQuery = "insert into `ref_global_system_config_bgp_router` (`from`, `to` ) values (?, ?);"

func CreateGlobalSystemConfig(tx *sql.Tx, model *models.GlobalSystemConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalSystemConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		bool(model.AlarmEnable),
		bool(model.IbgpAutoMesh),
		bool(model.BGPAlwaysCompareMed),
		int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.ConfigVersion),
		int(model.MacMoveControl.MacMoveTimeWindow),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		utils.MustJSON(model.UserDefinedLogStatistics),
		string(model.DisplayName),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.PluginTuning.PluginProperty),
		utils.MustJSON(model.IPFabricSubnets.Subnet),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.BgpaasParameters.PortStart),
		int(model.BgpaasParameters.PortEnd),
		int(model.MacAgingTime),
		bool(model.GracefulRestartParameters.XMPPHelperEnable),
		int(model.GracefulRestartParameters.RestartTime),
		int(model.GracefulRestartParameters.LongLivedRestartTime),
		bool(model.GracefulRestartParameters.Enable),
		int(model.GracefulRestartParameters.EndOfRibTimeout),
		bool(model.GracefulRestartParameters.BGPHelperEnable),
		int(model.AutonomousSystem),
		string(model.UUID))

	stmtBGPRouterRef, err := tx.Prepare(insertGlobalSystemConfigBGPRouterQuery)
	if err != nil {
		return err
	}
	defer stmtBGPRouterRef.Close()
	for _, ref := range model.BGPRouterRefs {
		_, err = stmtBGPRouterRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanGlobalSystemConfig(rows *sql.Rows) (*models.GlobalSystemConfig, error) {
	m := models.MakeGlobalSystemConfig()

	var jsonFQName string

	var jsonUserDefinedLogStatistics string

	var jsonPerms2Share string

	var jsonPluginTuningPluginProperty string

	var jsonIPFabricSubnetsSubnet string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&jsonFQName,
		&m.AlarmEnable,
		&m.IbgpAutoMesh,
		&m.BGPAlwaysCompareMed,
		&m.MacLimitControl.MacLimit,
		&m.MacLimitControl.MacLimitAction,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.ConfigVersion,
		&m.MacMoveControl.MacMoveTimeWindow,
		&m.MacMoveControl.MacMoveLimit,
		&m.MacMoveControl.MacMoveLimitAction,
		&jsonUserDefinedLogStatistics,
		&m.DisplayName,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPluginTuningPluginProperty,
		&jsonIPFabricSubnetsSubnet,
		&jsonAnnotationsKeyValuePair,
		&m.BgpaasParameters.PortStart,
		&m.BgpaasParameters.PortEnd,
		&m.MacAgingTime,
		&m.GracefulRestartParameters.XMPPHelperEnable,
		&m.GracefulRestartParameters.RestartTime,
		&m.GracefulRestartParameters.LongLivedRestartTime,
		&m.GracefulRestartParameters.Enable,
		&m.GracefulRestartParameters.EndOfRibTimeout,
		&m.GracefulRestartParameters.BGPHelperEnable,
		&m.AutonomousSystem,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonUserDefinedLogStatistics), &m.UserDefinedLogStatistics)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonPluginTuningPluginProperty), &m.PluginTuning.PluginProperty)

	json.Unmarshal([]byte(jsonIPFabricSubnetsSubnet), &m.IPFabricSubnets.Subnet)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildGlobalSystemConfigWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["mac_limit_action"]; ok {
		results = append(results, "mac_limit_action = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["config_version"]; ok {
		results = append(results, "config_version = ?")
		values = append(values, value)
	}

	if value, ok := where["mac_move_limit_action"]; ok {
		results = append(results, "mac_move_limit_action = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListGlobalSystemConfig(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.GlobalSystemConfig, error) {
	result := models.MakeGlobalSystemConfigSlice()
	whereQuery, values := buildGlobalSystemConfigWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listGlobalSystemConfigQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanGlobalSystemConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowGlobalSystemConfig(tx *sql.Tx, uuid string) (*models.GlobalSystemConfig, error) {
	rows, err := tx.Query(showGlobalSystemConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanGlobalSystemConfig(rows)
	}
	return nil, nil
}

func UpdateGlobalSystemConfig(tx *sql.Tx, uuid string, model *models.GlobalSystemConfig) error {
	return nil
}

func DeleteGlobalSystemConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteGlobalSystemConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
