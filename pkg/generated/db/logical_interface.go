package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertLogicalInterfaceQuery = "insert into `logical_interface` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`logical_interface_vlan_tag`,`logical_interface_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteLogicalInterfaceQuery = "delete from `logical_interface` where uuid = ?"

// LogicalInterfaceFields is db columns for LogicalInterface
var LogicalInterfaceFields = []string{
   "uuid",
   "share",
   "owner_access",
   "owner",
   "global_access",
   "parent_uuid",
   "parent_type",
   "logical_interface_vlan_tag",
   "logical_interface_type",
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

// LogicalInterfaceRefFields is db reference fields for LogicalInterface
var LogicalInterfaceRefFields = map[string][]string{
   
    "virtual_machine_interface": []string{
        // <common.Schema Value>
        
    },
   
}

// LogicalInterfaceBackRefFields is db back reference fields for LogicalInterface
var LogicalInterfaceBackRefFields = map[string][]string{
   
}

// LogicalInterfaceParentTypes is possible parents for LogicalInterface
var LogicalInterfaceParents = []string{
   
   "physical_router",
   
   "physical_interface",
   
}


const insertLogicalInterfaceVirtualMachineInterfaceQuery = "insert into `ref_logical_interface_virtual_machine_interface` (`from`, `to` ) values (?, ?);"


// CreateLogicalInterface inserts LogicalInterface to DB
func CreateLogicalInterface(tx *sql.Tx, model *models.LogicalInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLogicalInterfaceQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertLogicalInterfaceQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    common.MustJSON(model.Perms2.Share),
    int(model.Perms2.OwnerAccess),
    string(model.Perms2.Owner),
    int(model.Perms2.GlobalAccess),
    string(model.ParentUUID),
    string(model.ParentType),
    int(model.LogicalInterfaceVlanTag),
    string(model.LogicalInterfaceType),
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
    
    stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLogicalInterfaceVirtualMachineInterfaceQuery)
	if err != nil {
        return errors.Wrap(err,"preparing VirtualMachineInterfaceRefs create statement failed")
	}
    defer stmtVirtualMachineInterfaceRef.Close()
    for _, ref := range model.VirtualMachineInterfaceRefs {
       
        _, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"VirtualMachineInterfaceRefs create failed")
        }
    }
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "logical_interface",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "logical_interface", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanLogicalInterface(values map[string]interface{} ) (*models.LogicalInterface, error) {
    m := models.MakeLogicalInterface()
    
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
    
    if value, ok := values["logical_interface_vlan_tag"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.LogicalInterfaceVlanTag = castedValue
            

        
    }
    
    if value, ok := values["logical_interface_type"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.LogicalInterfaceType = models.LogicalInterfaceType(castedValue)
            

        
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
    
    
    if value, ok := values["ref_virtual_machine_interface"]; ok {
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
            referenceModel := &models.LogicalInterfaceVirtualMachineInterfaceRef{}
            referenceModel.UUID = uuid
            m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)
            
        }
    }
    
    
    return m, nil
}

// ListLogicalInterface lists LogicalInterface with list spec.
func ListLogicalInterface(tx *sql.Tx, spec *common.ListSpec) ([]*models.LogicalInterface, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "logical_interface"
    spec.Fields = LogicalInterfaceFields
    spec.RefFields = LogicalInterfaceRefFields
    spec.BackRefFields = LogicalInterfaceBackRefFields
    result := models.MakeLogicalInterfaceSlice()

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
            m, err := scanLogicalInterface(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateLogicalInterface updates a resource
func UpdateLogicalInterface(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateLogicalInterfaceQuery = "update `logical_interface` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateLogicalInterfaceQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateLogicalInterfaceQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateLogicalInterfaceQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateLogicalInterfaceQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateLogicalInterfaceQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateLogicalInterfaceQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateLogicalInterfaceQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LogicalInterfaceVlanTag" , "."); ok {
        updateLogicalInterfaceQuery += "`logical_interface_vlan_tag` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LogicalInterfaceType" , "."); ok {
        updateLogicalInterfaceQuery += "`logical_interface_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateLogicalInterfaceQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateLogicalInterfaceQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateLogicalInterfaceQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateLogicalInterfaceQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateLogicalInterfaceQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateLogicalInterfaceQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateLogicalInterfaceQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateLogicalInterfaceQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateLogicalInterfaceQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateLogicalInterfaceQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateLogicalInterfaceQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateLogicalInterfaceQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateLogicalInterfaceQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateLogicalInterfaceQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateLogicalInterfaceQuery += ","
    }
    
    updateLogicalInterfaceQuery =
    updateLogicalInterfaceQuery[:len(updateLogicalInterfaceQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateLogicalInterfaceQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateLogicalInterfaceQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
        if value, ok := common.GetValueByPath(model, "VirtualMachineInterfaceRefs" , "."); ok {
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
                        refQuery = "insert into `ref_logical_interface_virtual_machine_interface` ("
                        values := "values(" 
                        for _, value := range refKeys {
                            refQuery += "`" + value + "`, "
                            values += "?,"
                        }
                        refQuery += "`from`, `to`) "
                        values += "?,?);"
                        refQuery += values
                    case common.UPDATE:
                        refQuery = "update `ref_logical_interface_virtual_machine_interface` set "
                        if len(refKeys) == 0 {
                            return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref VirtualMachineInterfaceRefs")
                        } 
                        for _, value := range refKeys {
                            refQuery += "`" + value + "` = ?,"
                        }
                        refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
                    case common.DELETE:
                        refQuery = "delete from `ref_logical_interface_virtual_machine_interface` where `from` = ? AND `to`= ?;"
                        refValues = refValues[len(refValues)-2:]
                    default:
                        return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
                }
                stmt, err := tx.Prepare(refQuery)
                if err != nil {
                    return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs update statement failed")
                }
                _, err = stmt.Exec( refValues... )
                if err != nil {
                    return errors.Wrap(err, "VirtualMachineInterfaceRefs update failed")
                }
            }
        }
    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "logical_interface", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteLogicalInterface deletes a resource
func DeleteLogicalInterface(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteLogicalInterfaceQuery
    selectQuery := "select count(uuid) from logical_interface where uuid = ?"
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