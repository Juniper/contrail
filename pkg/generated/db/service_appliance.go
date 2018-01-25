package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertServiceApplianceQuery = "insert into `service_appliance` (`uuid`,`username`,`password`,`key_value_pair`,`service_appliance_ip_address`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteServiceApplianceQuery = "delete from `service_appliance` where uuid = ?"

// ServiceApplianceFields is db columns for ServiceAppliance
var ServiceApplianceFields = []string{
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
   
}

// ServiceApplianceRefFields is db reference fields for ServiceAppliance
var ServiceApplianceRefFields = map[string][]string{
   
    "physical_interface": []string{
        // <common.Schema Value>
        "interface_type",
        
    },
   
}

// ServiceApplianceBackRefFields is db back reference fields for ServiceAppliance
var ServiceApplianceBackRefFields = map[string][]string{
   
}

// ServiceApplianceParentTypes is possible parents for ServiceAppliance
var ServiceApplianceParents = []string{
   
   "service_appliance_set",
   
}


const insertServiceAppliancePhysicalInterfaceQuery = "insert into `ref_service_appliance_physical_interface` (`from`, `to` ,`interface_type`) values (?, ?,?);"


// CreateServiceAppliance inserts ServiceAppliance to DB
func CreateServiceAppliance(tx *sql.Tx, model *models.ServiceAppliance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertServiceApplianceQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    string(model.ServiceApplianceUserCredentials.Username),
    string(model.ServiceApplianceUserCredentials.Password),
    common.MustJSON(model.ServiceApplianceProperties.KeyValuePair),
    string(model.ServiceApplianceIPAddress),
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
    
    stmtPhysicalInterfaceRef, err := tx.Prepare(insertServiceAppliancePhysicalInterfaceQuery)
	if err != nil {
        return errors.Wrap(err,"preparing PhysicalInterfaceRefs create statement failed")
	}
    defer stmtPhysicalInterfaceRef.Close()
    for _, ref := range model.PhysicalInterfaceRefs {
       
       if ref.Attr == nil {
           ref.Attr = models.MakeServiceApplianceInterfaceType()
       }
       
        _, err = stmtPhysicalInterfaceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.InterfaceType))
	    if err != nil {
            return errors.Wrap(err,"PhysicalInterfaceRefs create failed")
        }
    }
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "service_appliance",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "service_appliance", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanServiceAppliance(values map[string]interface{} ) (*models.ServiceAppliance, error) {
    m := models.MakeServiceAppliance()
    
    if value, ok := values["uuid"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.UUID = castedValue
            

        
    }
    
    if value, ok := values["username"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ServiceApplianceUserCredentials.Username = castedValue
            

        
    }
    
    if value, ok := values["password"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ServiceApplianceUserCredentials.Password = castedValue
            

        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.ServiceApplianceProperties.KeyValuePair)
        
    }
    
    if value, ok := values["service_appliance_ip_address"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ServiceApplianceIPAddress = models.IpAddressType(castedValue)
            

        
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
    
    
    if value, ok := values["ref_physical_interface"]; ok {
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
            referenceModel := &models.ServiceAppliancePhysicalInterfaceRef{}
            referenceModel.UUID = uuid
            m.PhysicalInterfaceRefs = append(m.PhysicalInterfaceRefs, referenceModel)
            
            attr := models.MakeServiceApplianceInterfaceType()
            referenceModel.Attr = attr
            
            
        }
    }
    
    
    return m, nil
}

// ListServiceAppliance lists ServiceAppliance with list spec.
func ListServiceAppliance(tx *sql.Tx, spec *common.ListSpec) ([]*models.ServiceAppliance, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "service_appliance"
    spec.Fields = ServiceApplianceFields
    spec.RefFields = ServiceApplianceRefFields
    spec.BackRefFields = ServiceApplianceBackRefFields
    result := models.MakeServiceApplianceSlice()

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
            m, err := scanServiceAppliance(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateServiceAppliance updates a resource
func UpdateServiceAppliance(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateServiceApplianceQuery = "update `service_appliance` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateServiceApplianceQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ServiceApplianceUserCredentials.Username" , "."); ok {
        updateServiceApplianceQuery += "`username` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ServiceApplianceUserCredentials.Password" , "."); ok {
        updateServiceApplianceQuery += "`password` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ServiceApplianceProperties.KeyValuePair" , "."); ok {
        updateServiceApplianceQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ServiceApplianceIPAddress" , "."); ok {
        updateServiceApplianceQuery += "`service_appliance_ip_address` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateServiceApplianceQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateServiceApplianceQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateServiceApplianceQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateServiceApplianceQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateServiceApplianceQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateServiceApplianceQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateServiceApplianceQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateServiceApplianceQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateServiceApplianceQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateServiceApplianceQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateServiceApplianceQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateServiceApplianceQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateServiceApplianceQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateServiceApplianceQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateServiceApplianceQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateServiceApplianceQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateServiceApplianceQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateServiceApplianceQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateServiceApplianceQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateServiceApplianceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateServiceApplianceQuery += "`annotations_key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateServiceApplianceQuery += ","
    }
    
    updateServiceApplianceQuery =
    updateServiceApplianceQuery[:len(updateServiceApplianceQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateServiceApplianceQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateServiceApplianceQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
        if value, ok := common.GetValueByPath(model, "PhysicalInterfaceRefs" , "."); ok {
            for _, ref := range value.([]interface{}) {
                refQuery := ""
                refValues := make([]interface{}, 0)
                refKeys := make([]string, 0)
                refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
                if !ok {
                    return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
                }
                
                    attrValues, ok := common.GetValueByPath(ref.(map[string]interface{}), "Attr", ".")
                    if ok {
                        
                        if value, ok := common.GetValueByPath(attrValues.(map[string]interface{}), ".InterfaceType" , "."); ok {
                            refKeys = append(refKeys, "interface_type")
                            
                                refValues = append(refValues, common.InterfaceToString(value))
                            
                        }
                        
                    }
                
                refValues = append(refValues, uuid)
                refValues = append(refValues, refUUID)
                operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
                switch operation {
                    case common.ADD:
                        refQuery = "insert into `ref_service_appliance_physical_interface` ("
                        values := "values(" 
                        for _, value := range refKeys {
                            refQuery += "`" + value + "`, "
                            values += "?,"
                        }
                        refQuery += "`from`, `to`) "
                        values += "?,?);"
                        refQuery += values
                    case common.UPDATE:
                        refQuery = "update `ref_service_appliance_physical_interface` set "
                        if len(refKeys) == 0 {
                            return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref PhysicalInterfaceRefs")
                        } 
                        for _, value := range refKeys {
                            refQuery += "`" + value + "` = ?,"
                        }
                        refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
                    case common.DELETE:
                        refQuery = "delete from `ref_service_appliance_physical_interface` where `from` = ? AND `to`= ?;"
                        refValues = refValues[len(refValues)-2:]
                    default:
                        return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
                }
                stmt, err := tx.Prepare(refQuery)
                if err != nil {
                    return errors.Wrap(err, "preparing PhysicalInterfaceRefs update statement failed")
                }
                _, err = stmt.Exec( refValues... )
                if err != nil {
                    return errors.Wrap(err, "PhysicalInterfaceRefs update failed")
                }
            }
        }
    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "service_appliance", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteServiceAppliance deletes a resource
func DeleteServiceAppliance(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteServiceApplianceQuery
    selectQuery := "select count(uuid) from service_appliance where uuid = ?"
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