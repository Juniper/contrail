package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertGlobalSystemConfigQuery = "insert into `global_system_config` (`uuid`,`statlist`,`plugin_property`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`mac_move_time_window`,`mac_move_limit_action`,`mac_move_limit`,`mac_limit_action`,`mac_limit`,`mac_aging_time`,`subnet`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`ibgp_auto_mesh`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`end_of_rib_timeout`,`graceful_restart_parameters_enable`,`bgp_helper_enable`,`fq_name`,`display_name`,`config_version`,`port_start`,`port_end`,`bgp_always_compare_med`,`autonomous_system`,`key_value_pair`,`alarm_enable`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteGlobalSystemConfigQuery = "delete from `global_system_config` where uuid = ?"

// GlobalSystemConfigFields is db columns for GlobalSystemConfig
var GlobalSystemConfigFields = []string{
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
   
}

// GlobalSystemConfigRefFields is db reference fields for GlobalSystemConfig
var GlobalSystemConfigRefFields = map[string][]string{
   
    "bgp_router": []string{
        // <common.Schema Value>
        
    },
   
}

// GlobalSystemConfigBackRefFields is db back reference fields for GlobalSystemConfig
var GlobalSystemConfigBackRefFields = map[string][]string{
   
   "alarm": []string{
        "uve_key",
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
        "alarm_severity",
        "or_list",
        
   },
   
   "analytics_node": []string{
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
        "analytics_node_ip_address",
        
   },
   
   "api_access_list": []string{
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
        "rbac_rule",
        "key_value_pair",
        
   },
   
   "config_node": []string{
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
        "config_node_ip_address",
        "key_value_pair",
        
   },
   
   "database_node": []string{
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
        "database_node_ip_address",
        "key_value_pair",
        
   },
   
   "global_qos_config": []string{
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
        "dns",
        "control",
        "analytics",
        "key_value_pair",
        
   },
   
   "global_vrouter_config": []string{
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
        
   },
   
   "physical_router": []string{
        "uuid",
        "server_port",
        "server_ip",
        "resource",
        "physical_router_vnc_managed",
        "physical_router_vendor_name",
        "username",
        "password",
        "version",
        "v3_security_name",
        "v3_security_level",
        "v3_security_engine_id",
        "v3_privacy_protocol",
        "v3_privacy_password",
        "v3_engine_time",
        "v3_engine_id",
        "v3_engine_boots",
        "v3_context_engine_id",
        "v3_context",
        "v3_authentication_protocol",
        "v3_authentication_password",
        "v2_community",
        "timeout",
        "retries",
        "local_port",
        "physical_router_snmp",
        "physical_router_role",
        "physical_router_product_name",
        "physical_router_management_ip",
        "physical_router_loopback_ip",
        "physical_router_lldp",
        "service_port",
        "physical_router_image_uri",
        "physical_router_dataplane_ip",
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
   
   "service_appliance_set": []string{
        "uuid",
        "key_value_pair",
        "service_appliance_ha_mode",
        "service_appliance_driver",
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
        "annotations_key_value_pair",
        
   },
   
   "virtual_router": []string{
        "virtual_router_type",
        "virtual_router_ip_address",
        "virtual_router_dpdk_enabled",
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
        
   },
   
}

// GlobalSystemConfigParentTypes is possible parents for GlobalSystemConfig
var GlobalSystemConfigParents = []string{
   
   "config_root",
   
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
    _, err = stmt.Exec(string(model.UUID),
    common.MustJSON(model.UserDefinedLogStatistics.Statlist),
    common.MustJSON(model.PluginTuning.PluginProperty),
    common.MustJSON(model.Perms2.Share),
    int(model.Perms2.OwnerAccess),
    string(model.Perms2.Owner),
    int(model.Perms2.GlobalAccess),
    string(model.ParentUUID),
    string(model.ParentType),
    int(model.MacMoveControl.MacMoveTimeWindow),
    string(model.MacMoveControl.MacMoveLimitAction),
    int(model.MacMoveControl.MacMoveLimit),
    string(model.MacLimitControl.MacLimitAction),
    int(model.MacLimitControl.MacLimit),
    int(model.MacAgingTime),
    common.MustJSON(model.IPFabricSubnets.Subnet),
    bool(model.IDPerms.UserVisible),
    int(model.IDPerms.Permissions.OwnerAccess),
    string(model.IDPerms.Permissions.Owner),
    int(model.IDPerms.Permissions.OtherAccess),
    int(model.IDPerms.Permissions.GroupAccess),
    string(model.IDPerms.Permissions.Group),
    string(model.IDPerms.LastModified),
    bool(model.IDPerms.Enable),
    string(model.IDPerms.Description),
    string(model.IDPerms.Creator),
    string(model.IDPerms.Created),
    bool(model.IbgpAutoMesh),
    bool(model.GracefulRestartParameters.XMPPHelperEnable),
    int(model.GracefulRestartParameters.RestartTime),
    int(model.GracefulRestartParameters.LongLivedRestartTime),
    int(model.GracefulRestartParameters.EndOfRibTimeout),
    bool(model.GracefulRestartParameters.Enable),
    bool(model.GracefulRestartParameters.BGPHelperEnable),
    common.MustJSON(model.FQName),
    string(model.DisplayName),
    string(model.ConfigVersion),
    int(model.BgpaasParameters.PortStart),
    int(model.BgpaasParameters.PortEnd),
    bool(model.BGPAlwaysCompareMed),
    int(model.AutonomousSystem),
    common.MustJSON(model.Annotations.KeyValuePair),
    bool(model.AlarmEnable))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    stmtBGPRouterRef, err := tx.Prepare(insertGlobalSystemConfigBGPRouterQuery)
	if err != nil {
        return errors.Wrap(err,"preparing BGPRouterRefs create statement failed")
	}
    defer stmtBGPRouterRef.Close()
    for _, ref := range model.BGPRouterRefs {
       
        _, err = stmtBGPRouterRef.Exec(model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"BGPRouterRefs create failed")
        }
    }
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "global_system_config",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "global_system_config", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanGlobalSystemConfig(values map[string]interface{} ) (*models.GlobalSystemConfig, error) {
    m := models.MakeGlobalSystemConfig()
    
    if value, ok := values["uuid"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.UUID = castedValue
            

        
    }
    
    if value, ok := values["statlist"]; ok {
        
            json.Unmarshal(value.([]byte), &m.UserDefinedLogStatistics.Statlist)
        
    }
    
    if value, ok := values["plugin_property"]; ok {
        
            json.Unmarshal(value.([]byte), &m.PluginTuning.PluginProperty)
        
    }
    
    if value, ok := values["share"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Perms2.Share)
        
    }
    
    if value, ok := values["owner_access"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.Perms2.OwnerAccess = models.AccessType(castedValue)
            

        
    }
    
    if value, ok := values["owner"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.Perms2.Owner = castedValue
            

        
    }
    
    if value, ok := values["global_access"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.Perms2.GlobalAccess = models.AccessType(castedValue)
            

        
    }
    
    if value, ok := values["parent_uuid"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ParentUUID = castedValue
            

        
    }
    
    if value, ok := values["parent_type"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ParentType = castedValue
            

        
    }
    
    if value, ok := values["mac_move_time_window"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.MacMoveControl.MacMoveTimeWindow = models.MACMoveTimeWindow(castedValue)
            

        
    }
    
    if value, ok := values["mac_move_limit_action"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.MacMoveControl.MacMoveLimitAction = models.MACLimitExceedActionType(castedValue)
            

        
    }
    
    if value, ok := values["mac_move_limit"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.MacMoveControl.MacMoveLimit = castedValue
            

        
    }
    
    if value, ok := values["mac_limit_action"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.MacLimitControl.MacLimitAction = models.MACLimitExceedActionType(castedValue)
            

        
    }
    
    if value, ok := values["mac_limit"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.MacLimitControl.MacLimit = castedValue
            

        
    }
    
    if value, ok := values["mac_aging_time"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.MacAgingTime = models.MACAgingTime(castedValue)
            

        
    }
    
    if value, ok := values["subnet"]; ok {
        
            json.Unmarshal(value.([]byte), &m.IPFabricSubnets.Subnet)
        
    }
    
    if value, ok := values["user_visible"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.IDPerms.UserVisible = castedValue
            

        
    }
    
    if value, ok := values["permissions_owner_access"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
            

        
    }
    
    if value, ok := values["permissions_owner"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.IDPerms.Permissions.Owner = castedValue
            

        
    }
    
    if value, ok := values["other_access"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
            

        
    }
    
    if value, ok := values["group_access"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
            

        
    }
    
    if value, ok := values["group"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.IDPerms.Permissions.Group = castedValue
            

        
    }
    
    if value, ok := values["last_modified"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.IDPerms.LastModified = castedValue
            

        
    }
    
    if value, ok := values["enable"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.IDPerms.Enable = castedValue
            

        
    }
    
    if value, ok := values["description"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.IDPerms.Description = castedValue
            

        
    }
    
    if value, ok := values["creator"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.IDPerms.Creator = castedValue
            

        
    }
    
    if value, ok := values["created"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.IDPerms.Created = castedValue
            

        
    }
    
    if value, ok := values["ibgp_auto_mesh"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.IbgpAutoMesh = castedValue
            

        
    }
    
    if value, ok := values["xmpp_helper_enable"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.GracefulRestartParameters.XMPPHelperEnable = castedValue
            

        
    }
    
    if value, ok := values["restart_time"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.GracefulRestartParameters.RestartTime = models.GracefulRestartTimeType(castedValue)
            

        
    }
    
    if value, ok := values["long_lived_restart_time"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.GracefulRestartParameters.LongLivedRestartTime = models.LongLivedGracefulRestartTimeType(castedValue)
            

        
    }
    
    if value, ok := values["end_of_rib_timeout"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.GracefulRestartParameters.EndOfRibTimeout = models.EndOfRibTimeType(castedValue)
            

        
    }
    
    if value, ok := values["graceful_restart_parameters_enable"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.GracefulRestartParameters.Enable = castedValue
            

        
    }
    
    if value, ok := values["bgp_helper_enable"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.GracefulRestartParameters.BGPHelperEnable = castedValue
            

        
    }
    
    if value, ok := values["fq_name"]; ok {
        
            json.Unmarshal(value.([]byte), &m.FQName)
        
    }
    
    if value, ok := values["display_name"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.DisplayName = castedValue
            

        
    }
    
    if value, ok := values["config_version"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ConfigVersion = castedValue
            

        
    }
    
    if value, ok := values["port_start"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.BgpaasParameters.PortStart = models.L4PortType(castedValue)
            

        
    }
    
    if value, ok := values["port_end"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.BgpaasParameters.PortEnd = models.L4PortType(castedValue)
            

        
    }
    
    if value, ok := values["bgp_always_compare_med"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.BGPAlwaysCompareMed = castedValue
            

        
    }
    
    if value, ok := values["autonomous_system"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.AutonomousSystem = models.AutonomousSystemType(castedValue)
            

        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    if value, ok := values["alarm_enable"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.AlarmEnable = castedValue
            

        
    }
    
    
    if value, ok := values["ref_bgp_router"]; ok {
        var references []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &references )
        for _, reference := range references {
            referenceMap, ok := reference.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(referenceMap["to"])
            if uuid == "" {
                continue
            }
            referenceModel := &models.GlobalSystemConfigBGPRouterRef{}
            referenceModel.UUID = uuid
            m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)
            
        }
    }
    
    
    if value, ok := values["backref_alarm"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeAlarm()
            m.Alarms = append(m.Alarms, childModel)

            
                if propertyValue, ok := childResourceMap["uve_key"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.UveKeys.UveKey)
                
                }
            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
                if propertyValue, ok := childResourceMap["alarm_severity"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.AlarmSeverity = models.AlarmSeverity(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["or_list"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.AlarmRules.OrList)
                
                }
            
        }
    }
    
    if value, ok := values["backref_analytics_node"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeAnalyticsNode()
            m.AnalyticsNodes = append(m.AnalyticsNodes, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
                if propertyValue, ok := childResourceMap["analytics_node_ip_address"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.AnalyticsNodeIPAddress = models.IpAddressType(castedValue)
                    
                
                }
            
        }
    }
    
    if value, ok := values["backref_api_access_list"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeAPIAccessList()
            m.APIAccessLists = append(m.APIAccessLists, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["rbac_rule"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.APIAccessListEntries.RbacRule)
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
        }
    }
    
    if value, ok := values["backref_config_node"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeConfigNode()
            m.ConfigNodes = append(m.ConfigNodes, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["config_node_ip_address"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ConfigNodeIPAddress = models.IpAddressType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
        }
    }
    
    if value, ok := values["backref_database_node"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeDatabaseNode()
            m.DatabaseNodes = append(m.DatabaseNodes, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["database_node_ip_address"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DatabaseNodeIPAddress = models.IpAddressType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
        }
    }
    
    if value, ok := values["backref_global_qos_config"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeGlobalQosConfig()
            m.GlobalQosConfigs = append(m.GlobalQosConfigs, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["dns"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.ControlTrafficDSCP.DNS = models.DscpValueType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["control"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.ControlTrafficDSCP.Control = models.DscpValueType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["analytics"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.ControlTrafficDSCP.Analytics = models.DscpValueType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
        }
    }
    
    if value, ok := values["backref_global_vrouter_config"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeGlobalVrouterConfig()
            m.GlobalVrouterConfigs = append(m.GlobalVrouterConfigs, childModel)

            
                if propertyValue, ok := childResourceMap["vxlan_network_identifier_mode"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.VxlanNetworkIdentifierMode = models.VxlanNetworkIdentifierModeType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["linklocal_service_entry"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.LinklocalServices.LinklocalServiceEntry)
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["forwarding_mode"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ForwardingMode = models.ForwardingModeType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["flow_export_rate"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.FlowExportRate = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["flow_aging_timeout"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FlowAgingTimeoutList.FlowAgingTimeout)
                
                }
            
                if propertyValue, ok := childResourceMap["encapsulation"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.EncapsulationPriorities.Encapsulation)
                
                }
            
                if propertyValue, ok := childResourceMap["enable_security_logging"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.EnableSecurityLogging = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["source_port"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.EcmpHashingIncludeFields.SourcePort = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["source_ip"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.EcmpHashingIncludeFields.SourceIP = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["ip_protocol"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.EcmpHashingIncludeFields.IPProtocol = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["hashing_configured"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.EcmpHashingIncludeFields.HashingConfigured = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["destination_port"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.EcmpHashingIncludeFields.DestinationPort = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["destination_ip"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.EcmpHashingIncludeFields.DestinationIP = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
        }
    }
    
    if value, ok := values["backref_physical_router"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakePhysicalRouter()
            m.PhysicalRouters = append(m.PhysicalRouters, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["server_port"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.TelemetryInfo.ServerPort = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["server_ip"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.TelemetryInfo.ServerIP = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["resource"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.TelemetryInfo.Resource)
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_vnc_managed"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.PhysicalRouterVNCManaged = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_vendor_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterVendorName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["username"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterUserCredentials.Username = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["password"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterUserCredentials.Password = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["version"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.Version = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_security_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3SecurityName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_security_level"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3SecurityLevel = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_security_engine_id"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3SecurityEngineID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_privacy_protocol"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3PrivacyProtocol = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_privacy_password"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3PrivacyPassword = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_engine_time"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3EngineTime = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_engine_id"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3EngineID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_engine_boots"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3EngineBoots = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_context_engine_id"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3ContextEngineID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_context"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3Context = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_authentication_protocol"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v3_authentication_password"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V3AuthenticationPassword = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["v2_community"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.V2Community = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["timeout"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.Timeout = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["retries"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.Retries = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["local_port"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMPCredentials.LocalPort = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_snmp"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.PhysicalRouterSNMP = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_role"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterRole = models.PhysicalRouterRole(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_product_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterProductName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_management_ip"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterManagementIP = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_loopback_ip"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterLoopbackIP = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_lldp"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.PhysicalRouterLLDP = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["service_port"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.PhysicalRouterJunosServicePorts.ServicePort)
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_image_uri"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterImageURI = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["physical_router_dataplane_ip"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.PhysicalRouterDataplaneIP = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
        }
    }
    
    if value, ok := values["backref_service_appliance_set"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeServiceApplianceSet()
            m.ServiceApplianceSets = append(m.ServiceApplianceSets, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.ServiceApplianceSetProperties.KeyValuePair)
                
                }
            
                if propertyValue, ok := childResourceMap["service_appliance_ha_mode"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ServiceApplianceHaMode = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["service_appliance_driver"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ServiceApplianceDriver = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["annotations_key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
        }
    }
    
    if value, ok := values["backref_virtual_router"]; ok {
        var childResources []interface{}
        stringValue := common.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := common.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeVirtualRouter()
            m.VirtualRouters = append(m.VirtualRouters, childModel)

            
                if propertyValue, ok := childResourceMap["virtual_router_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.VirtualRouterType = models.VirtualRouterType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["virtual_router_ip_address"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.VirtualRouterIPAddress = models.IpAddressType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["virtual_router_dpdk_enabled"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.VirtualRouterDPDKEnabled = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Perms2.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Perms2.GlobalAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentUUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ParentType = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.UserVisible = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Owner = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Permissions.Group = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.LastModified = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.IDPerms.Enable = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Description = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Creator = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.IDPerms.Created = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
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
    spec.BackRefFields = GlobalSystemConfigBackRefFields
    result := models.MakeGlobalSystemConfigSlice()

    if spec.ParentFQName != nil {
        parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
        if err != nil {
            return nil, errors.Wrap(err, "can't find parents")
        }
        spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
    }

    query := spec.BuildQuery()
    columns := spec.Columns
    values := spec.Values
    log.WithFields(log.Fields{
        "listSpec": spec,
        "query": query,
    }).Debug("select query")
    rows, err = tx.Query(query, values...)
    if err != nil {
        return nil, errors.Wrap(err,"select query failed")
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
            m, err := scanGlobalSystemConfig(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateGlobalSystemConfig updates a resource
func UpdateGlobalSystemConfig(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateGlobalSystemConfigQuery = "update `global_system_config` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateGlobalSystemConfigQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".UserDefinedLogStatistics.Statlist" , "."); ok {
        updateGlobalSystemConfigQuery += "`statlist` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".PluginTuning.PluginProperty" , "."); ok {
        updateGlobalSystemConfigQuery += "`plugin_property` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateGlobalSystemConfigQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateGlobalSystemConfigQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateGlobalSystemConfigQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateGlobalSystemConfigQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateGlobalSystemConfigQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateGlobalSystemConfigQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".MacMoveControl.MacMoveTimeWindow" , "."); ok {
        updateGlobalSystemConfigQuery += "`mac_move_time_window` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".MacMoveControl.MacMoveLimitAction" , "."); ok {
        updateGlobalSystemConfigQuery += "`mac_move_limit_action` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".MacMoveControl.MacMoveLimit" , "."); ok {
        updateGlobalSystemConfigQuery += "`mac_move_limit` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".MacLimitControl.MacLimitAction" , "."); ok {
        updateGlobalSystemConfigQuery += "`mac_limit_action` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".MacLimitControl.MacLimit" , "."); ok {
        updateGlobalSystemConfigQuery += "`mac_limit` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".MacAgingTime" , "."); ok {
        updateGlobalSystemConfigQuery += "`mac_aging_time` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IPFabricSubnets.Subnet" , "."); ok {
        updateGlobalSystemConfigQuery += "`subnet` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateGlobalSystemConfigQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateGlobalSystemConfigQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateGlobalSystemConfigQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateGlobalSystemConfigQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateGlobalSystemConfigQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateGlobalSystemConfigQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateGlobalSystemConfigQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateGlobalSystemConfigQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateGlobalSystemConfigQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateGlobalSystemConfigQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateGlobalSystemConfigQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IbgpAutoMesh" , "."); ok {
        updateGlobalSystemConfigQuery += "`ibgp_auto_mesh` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".GracefulRestartParameters.XMPPHelperEnable" , "."); ok {
        updateGlobalSystemConfigQuery += "`xmpp_helper_enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".GracefulRestartParameters.RestartTime" , "."); ok {
        updateGlobalSystemConfigQuery += "`restart_time` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".GracefulRestartParameters.LongLivedRestartTime" , "."); ok {
        updateGlobalSystemConfigQuery += "`long_lived_restart_time` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".GracefulRestartParameters.EndOfRibTimeout" , "."); ok {
        updateGlobalSystemConfigQuery += "`end_of_rib_timeout` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".GracefulRestartParameters.Enable" , "."); ok {
        updateGlobalSystemConfigQuery += "`graceful_restart_parameters_enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".GracefulRestartParameters.BGPHelperEnable" , "."); ok {
        updateGlobalSystemConfigQuery += "`bgp_helper_enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateGlobalSystemConfigQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateGlobalSystemConfigQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ConfigVersion" , "."); ok {
        updateGlobalSystemConfigQuery += "`config_version` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".BgpaasParameters.PortStart" , "."); ok {
        updateGlobalSystemConfigQuery += "`port_start` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".BgpaasParameters.PortEnd" , "."); ok {
        updateGlobalSystemConfigQuery += "`port_end` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".BGPAlwaysCompareMed" , "."); ok {
        updateGlobalSystemConfigQuery += "`bgp_always_compare_med` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".AutonomousSystem" , "."); ok {
        updateGlobalSystemConfigQuery += "`autonomous_system` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateGlobalSystemConfigQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".AlarmEnable" , "."); ok {
        updateGlobalSystemConfigQuery += "`alarm_enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateGlobalSystemConfigQuery += ","
    }
    
    updateGlobalSystemConfigQuery =
    updateGlobalSystemConfigQuery[:len(updateGlobalSystemConfigQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateGlobalSystemConfigQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateGlobalSystemConfigQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
        if value, ok := common.GetValueByPath(model, "BGPRouterRefs" , "."); ok {
            for _, ref := range value.([]interface{}) {
                refQuery := ""
                refValues := make([]interface{}, 0)
                refKeys := make([]string, 0)
                refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
                if !ok {
                    return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
                }
                
                refValues = append(refValues, uuid)
                refValues = append(refValues, refUUID)
                operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
                switch operation {
                    case common.ADD:
                        refQuery = "insert into `ref_global_system_config_bgp_router` ("
                        values := "values(" 
                        for _, value := range refKeys {
                            refQuery += "`" + value + "`, "
                            values += "?,"
                        }
                        refQuery += "`from`, `to`) "
                        values += "?,?);"
                        refQuery += values
                    case common.UPDATE:
                        refQuery = "update `ref_global_system_config_bgp_router` set "
                        if len(refKeys) == 0 {
                            return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref BGPRouterRefs")
                        } 
                        for _, value := range refKeys {
                            refQuery += "`" + value + "` = ?,"
                        }
                        refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
                    case common.DELETE:
                        refQuery = "delete from `ref_global_system_config_bgp_router` where `from` = ? AND `to`= ?;"
                        refValues = refValues[len(refValues)-2:]
                    default:
                        return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
                }
                stmt, err := tx.Prepare(refQuery)
                if err != nil {
                    return errors.Wrap(err, "preparing BGPRouterRefs update statement failed")
                }
                _, err = stmt.Exec( refValues... )
                if err != nil {
                    return errors.Wrap(err, "BGPRouterRefs update failed")
                }
            }
        }
    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "global_system_config", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteGlobalSystemConfig deletes a resource
func DeleteGlobalSystemConfig(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteGlobalSystemConfigQuery
    selectQuery := "select count(uuid) from global_system_config where uuid = ?"
    var err error
    var count int

    if auth.IsAdmin() {
        row := tx.QueryRow(selectQuery, uuid)
        if err != nil {
            return errors.Wrap(err, "not found")
        }
        row.Scan(&count)
        if count == 0 {
           return errors.New("Not found")
        }
        _, err = tx.Exec(deleteQuery, uuid)
    }else{
        deleteQuery += " and owner = ?"
        selectQuery += " and owner = ?"
        row := tx.QueryRow(selectQuery, uuid, auth.ProjectID() )
        if err != nil {
            return errors.Wrap(err, "not found")
        }
        row.Scan(&count)
        if count == 0 {
           return errors.New("Not found")
        }
        _, err = tx.Exec(deleteQuery, uuid, auth.ProjectID() )
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