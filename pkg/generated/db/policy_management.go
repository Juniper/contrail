package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertPolicyManagementQuery = "insert into `policy_management` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deletePolicyManagementQuery = "delete from `policy_management` where uuid = ?"

// PolicyManagementFields is db columns for PolicyManagement
var PolicyManagementFields = []string{
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

// PolicyManagementRefFields is db reference fields for PolicyManagement
var PolicyManagementRefFields = map[string][]string{
   
}

// PolicyManagementBackRefFields is db back reference fields for PolicyManagement
var PolicyManagementBackRefFields = map[string][]string{
   
   "address_group": []string{
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
        "subnet",
        
   },
   
   "application_policy_set": []string{
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
        
   },
   
   "firewall_policy": []string{
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
   
   "firewall_rule": []string{
        "uuid",
        "start_port",
        "end_port",
        "protocol_id",
        "protocol",
        "dst_ports_start_port",
        "dst_ports_end_port",
        "share",
        "owner_access",
        "owner",
        "global_access",
        "parent_uuid",
        "parent_type",
        "tag_list",
        "tag_type",
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
        "virtual_network",
        "tags",
        "tag_ids",
        "ip_prefix_len",
        "ip_prefix",
        "any",
        "address_group",
        "endpoint_1_virtual_network",
        "endpoint_1_tags",
        "endpoint_1_tag_ids",
        "subnet_ip_prefix_len",
        "subnet_ip_prefix",
        "endpoint_1_any",
        "endpoint_1_address_group",
        "display_name",
        "direction",
        "key_value_pair",
        "simple_action",
        "qos_action",
        "udp_port",
        "vtep_dst_mac_address",
        "vtep_dst_ip_address",
        "vni",
        "routing_instance",
        "nic_assisted_mirroring_vlan",
        "nic_assisted_mirroring",
        "nh_mode",
        "juniper_header",
        "encapsulation",
        "analyzer_name",
        "analyzer_mac_address",
        "analyzer_ip_address",
        "log",
        "gateway_name",
        "assign_routing_instance",
        "apply_service",
        "alert",
        
   },
   
   "service_group": []string{
        "uuid",
        "firewall_service",
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

// PolicyManagementParentTypes is possible parents for PolicyManagement
var PolicyManagementParents = []string{
   
}



// CreatePolicyManagement inserts PolicyManagement to DB
func CreatePolicyManagement(tx *sql.Tx, model *models.PolicyManagement) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPolicyManagementQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertPolicyManagementQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    common.MustJSON(model.Perms2.Share),
    int(model.Perms2.OwnerAccess),
    string(model.Perms2.Owner),
    int(model.Perms2.GlobalAccess),
    string(model.ParentUUID),
    string(model.ParentType),
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
    common.MustJSON(model.FQName),
    string(model.DisplayName),
    common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "policy_management",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "policy_management", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanPolicyManagement(values map[string]interface{} ) (*models.PolicyManagement, error) {
    m := models.MakePolicyManagement()
    
    if value, ok := values["uuid"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.UUID = castedValue
            

        
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
    
    if value, ok := values["fq_name"]; ok {
        
            json.Unmarshal(value.([]byte), &m.FQName)
        
    }
    
    if value, ok := values["display_name"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.DisplayName = castedValue
            

        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    
    if value, ok := values["backref_address_group"]; ok {
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
            childModel := models.MakeAddressGroup()
            m.AddressGroups = append(m.AddressGroups, childModel)

            
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
            
                if propertyValue, ok := childResourceMap["subnet"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.AddressGroupPrefix.Subnet)
                
                }
            
        }
    }
    
    if value, ok := values["backref_application_policy_set"]; ok {
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
            childModel := models.MakeApplicationPolicySet()
            m.ApplicationPolicySets = append(m.ApplicationPolicySets, childModel)

            
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
            
                if propertyValue, ok := childResourceMap["all_applications"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.AllApplications = castedValue
                    
                
                }
            
        }
    }
    
    if value, ok := values["backref_firewall_policy"]; ok {
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
            childModel := models.MakeFirewallPolicy()
            m.FirewallPolicys = append(m.FirewallPolicys, childModel)

            
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
    
    if value, ok := values["backref_firewall_rule"]; ok {
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
            childModel := models.MakeFirewallRule()
            m.FirewallRules = append(m.FirewallRules, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["start_port"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Service.SRCPorts.StartPort = models.L4PortType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["end_port"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Service.SRCPorts.EndPort = models.L4PortType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["protocol_id"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Service.ProtocolID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["protocol"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Service.Protocol = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["dst_ports_start_port"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Service.DSTPorts.StartPort = models.L4PortType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["dst_ports_end_port"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Service.DSTPorts.EndPort = models.L4PortType(castedValue)
                    
                
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
            
                if propertyValue, ok := childResourceMap["tag_list"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.MatchTags.TagList)
                
                }
            
                if propertyValue, ok := childResourceMap["tag_type"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.MatchTagTypes.TagType)
                
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
            
                if propertyValue, ok := childResourceMap["virtual_network"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Endpoint2.VirtualNetwork = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["tags"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Endpoint2.Tags)
                
                }
            
                if propertyValue, ok := childResourceMap["tag_ids"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Endpoint2.TagIds)
                
                }
            
                if propertyValue, ok := childResourceMap["ip_prefix_len"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Endpoint2.Subnet.IPPrefixLen = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["ip_prefix"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Endpoint2.Subnet.IPPrefix = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["any"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.Endpoint2.Any = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["address_group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Endpoint2.AddressGroup = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["endpoint_1_virtual_network"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Endpoint1.VirtualNetwork = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["endpoint_1_tags"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Endpoint1.Tags)
                
                }
            
                if propertyValue, ok := childResourceMap["endpoint_1_tag_ids"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Endpoint1.TagIds)
                
                }
            
                if propertyValue, ok := childResourceMap["subnet_ip_prefix_len"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.Endpoint1.Subnet.IPPrefixLen = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["subnet_ip_prefix"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Endpoint1.Subnet.IPPrefix = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["endpoint_1_any"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.Endpoint1.Any = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["endpoint_1_address_group"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Endpoint1.AddressGroup = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.DisplayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["direction"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.Direction = models.FirewallRuleDirectionType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
                if propertyValue, ok := childResourceMap["simple_action"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.SimpleAction = models.SimpleActionType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["qos_action"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.QosAction = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["udp_port"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.UDPPort = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["vtep_dst_mac_address"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["vtep_dst_ip_address"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["vni"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.StaticNHHeader.Vni = models.VxlanNetworkIdentifierType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["routing_instance"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.RoutingInstance = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["nic_assisted_mirroring_vlan"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToInt(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.NicAssistedMirroringVlan = models.VlanIdType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["nic_assisted_mirroring"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.NicAssistedMirroring = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["nh_mode"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.NHMode = models.NHModeType(castedValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["juniper_header"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.JuniperHeader = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["encapsulation"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.Encapsulation = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["analyzer_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.AnalyzerName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["analyzer_mac_address"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.AnalyzerMacAddress = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["analyzer_ip_address"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.MirrorTo.AnalyzerIPAddress = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["log"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.ActionList.Log = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["gateway_name"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.GatewayName = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["assign_routing_instance"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ActionList.AssignRoutingInstance = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["apply_service"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.ActionList.ApplyService)
                
                }
            
                if propertyValue, ok := childResourceMap["alert"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToBool(propertyValue)
                    
                    
                        childModel.ActionList.Alert = castedValue
                    
                
                }
            
        }
    }
    
    if value, ok := values["backref_service_group"]; ok {
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
            childModel := models.MakeServiceGroup()
            m.ServiceGroups = append(m.ServiceGroups, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["firewall_service"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.ServiceGroupFirewallServiceList.FirewallService)
                
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

// ListPolicyManagement lists PolicyManagement with list spec.
func ListPolicyManagement(tx *sql.Tx, spec *common.ListSpec) ([]*models.PolicyManagement, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "policy_management"
    spec.Fields = PolicyManagementFields
    spec.RefFields = PolicyManagementRefFields
    spec.BackRefFields = PolicyManagementBackRefFields
    result := models.MakePolicyManagementSlice()

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
            m, err := scanPolicyManagement(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdatePolicyManagement updates a resource
func UpdatePolicyManagement(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updatePolicyManagementQuery = "update `policy_management` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updatePolicyManagementQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updatePolicyManagementQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updatePolicyManagementQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updatePolicyManagementQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updatePolicyManagementQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updatePolicyManagementQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updatePolicyManagementQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updatePolicyManagementQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updatePolicyManagementQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updatePolicyManagementQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updatePolicyManagementQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updatePolicyManagementQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updatePolicyManagementQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updatePolicyManagementQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updatePolicyManagementQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updatePolicyManagementQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updatePolicyManagementQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updatePolicyManagementQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updatePolicyManagementQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updatePolicyManagementQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updatePolicyManagementQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updatePolicyManagementQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updatePolicyManagementQuery += ","
    }
    
    updatePolicyManagementQuery =
    updatePolicyManagementQuery[:len(updatePolicyManagementQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updatePolicyManagementQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updatePolicyManagementQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "policy_management", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeletePolicyManagement deletes a resource
func DeletePolicyManagement(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deletePolicyManagementQuery
    selectQuery := "select count(uuid) from policy_management where uuid = ?"
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