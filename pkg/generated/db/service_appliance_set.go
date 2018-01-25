package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertServiceApplianceSetQuery = "insert into `service_appliance_set` (`uuid`,`key_value_pair`,`service_appliance_ha_mode`,`service_appliance_driver`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteServiceApplianceSetQuery = "delete from `service_appliance_set` where uuid = ?"

// ServiceApplianceSetFields is db columns for ServiceApplianceSet
var ServiceApplianceSetFields = []string{
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
   
}

// ServiceApplianceSetRefFields is db reference fields for ServiceApplianceSet
var ServiceApplianceSetRefFields = map[string][]string{
   
}

// ServiceApplianceSetBackRefFields is db back reference fields for ServiceApplianceSet
var ServiceApplianceSetBackRefFields = map[string][]string{
   
   "service_appliance": []string{
        "uuid",
        "username",
        "password",
        "key_value_pair",
        "service_appliance_ip_address",
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
   
}

// ServiceApplianceSetParentTypes is possible parents for ServiceApplianceSet
var ServiceApplianceSetParents = []string{
   
   "global_system_config",
   
}



// CreateServiceApplianceSet inserts ServiceApplianceSet to DB
func CreateServiceApplianceSet(tx *sql.Tx, model *models.ServiceApplianceSet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceSetQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertServiceApplianceSetQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    common.MustJSON(model.ServiceApplianceSetProperties.KeyValuePair),
    string(model.ServiceApplianceHaMode),
    string(model.ServiceApplianceDriver),
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
        Type: "service_appliance_set",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "service_appliance_set", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanServiceApplianceSet(values map[string]interface{} ) (*models.ServiceApplianceSet, error) {
    m := models.MakeServiceApplianceSet()
    
    if value, ok := values["uuid"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.UUID = castedValue
            

        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.ServiceApplianceSetProperties.KeyValuePair)
        
    }
    
    if value, ok := values["service_appliance_ha_mode"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ServiceApplianceHaMode = castedValue
            

        
    }
    
    if value, ok := values["service_appliance_driver"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ServiceApplianceDriver = castedValue
            

        
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
    
    if value, ok := values["annotations_key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    
    if value, ok := values["backref_service_appliance"]; ok {
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
            childModel := models.MakeServiceAppliance()
            m.ServiceAppliances = append(m.ServiceAppliances, childModel)

            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.UUID = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["username"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ServiceApplianceUserCredentials.Username = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["password"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ServiceApplianceUserCredentials.Password = castedValue
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.ServiceApplianceProperties.KeyValuePair)
                
                }
            
                if propertyValue, ok := childResourceMap["service_appliance_ip_address"]; ok && propertyValue != nil {
                
                    
                        castedValue := common.InterfaceToString(propertyValue)
                    
                    
                        childModel.ServiceApplianceIPAddress = models.IpAddressType(castedValue)
                    
                
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
    
    return m, nil
}

// ListServiceApplianceSet lists ServiceApplianceSet with list spec.
func ListServiceApplianceSet(tx *sql.Tx, spec *common.ListSpec) ([]*models.ServiceApplianceSet, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "service_appliance_set"
    spec.Fields = ServiceApplianceSetFields
    spec.RefFields = ServiceApplianceSetRefFields
    spec.BackRefFields = ServiceApplianceSetBackRefFields
    result := models.MakeServiceApplianceSetSlice()

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
            m, err := scanServiceApplianceSet(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateServiceApplianceSet updates a resource
func UpdateServiceApplianceSet(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateServiceApplianceSetQuery = "update `service_appliance_set` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateServiceApplianceSetQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ServiceApplianceSetProperties.KeyValuePair" , "."); ok {
        updateServiceApplianceSetQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ServiceApplianceHaMode" , "."); ok {
        updateServiceApplianceSetQuery += "`service_appliance_ha_mode` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ServiceApplianceDriver" , "."); ok {
        updateServiceApplianceSetQuery += "`service_appliance_driver` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateServiceApplianceSetQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateServiceApplianceSetQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateServiceApplianceSetQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateServiceApplianceSetQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateServiceApplianceSetQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateServiceApplianceSetQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateServiceApplianceSetQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateServiceApplianceSetQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateServiceApplianceSetQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateServiceApplianceSetQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateServiceApplianceSetQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateServiceApplianceSetQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateServiceApplianceSetQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateServiceApplianceSetQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateServiceApplianceSetQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateServiceApplianceSetQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateServiceApplianceSetQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateServiceApplianceSetQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateServiceApplianceSetQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateServiceApplianceSetQuery += "`annotations_key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateServiceApplianceSetQuery += ","
    }
    
    updateServiceApplianceSetQuery =
    updateServiceApplianceSetQuery[:len(updateServiceApplianceSetQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateServiceApplianceSetQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateServiceApplianceSetQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "service_appliance_set", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteServiceApplianceSet deletes a resource
func DeleteServiceApplianceSet(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteServiceApplianceSetQuery
    selectQuery := "select count(uuid) from service_appliance_set where uuid = ?"
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