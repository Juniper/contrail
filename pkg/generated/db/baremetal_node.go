package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertBaremetalNodeQuery = "insert into `baremetal_node` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`name`,`memory_mb`,`ipmi_username`,`ipmi_password`,`ipmi_address`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`disk_gb`,`deploy_ramdisk`,`deploy_kernel`,`cpu_count`,`cpu_arch`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteBaremetalNodeQuery = "delete from `baremetal_node` where uuid = ?"

// BaremetalNodeFields is db columns for BaremetalNode
var BaremetalNodeFields = []string{
   "uuid",
   "share",
   "owner_access",
   "owner",
   "global_access",
   "parent_uuid",
   "parent_type",
   "name",
   "memory_mb",
   "ipmi_username",
   "ipmi_password",
   "ipmi_address",
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
   "disk_gb",
   "deploy_ramdisk",
   "deploy_kernel",
   "cpu_count",
   "cpu_arch",
   "key_value_pair",
   
}

// BaremetalNodeRefFields is db reference fields for BaremetalNode
var BaremetalNodeRefFields = map[string][]string{
   
}

// BaremetalNodeBackRefFields is db back reference fields for BaremetalNode
var BaremetalNodeBackRefFields = map[string][]string{
   
}

// BaremetalNodeParentTypes is possible parents for BaremetalNode
var BaremetalNodeParents = []string{
   
}



// CreateBaremetalNode inserts BaremetalNode to DB
func CreateBaremetalNode(tx *sql.Tx, model *models.BaremetalNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBaremetalNodeQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertBaremetalNodeQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    common.MustJSON(model.Perms2.Share),
    int(model.Perms2.OwnerAccess),
    string(model.Perms2.Owner),
    int(model.Perms2.GlobalAccess),
    string(model.ParentUUID),
    string(model.ParentType),
    string(model.Name),
    int(model.MemoryMB),
    string(model.IpmiUsername),
    string(model.IpmiPassword),
    string(model.IpmiAddress),
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
    int(model.DiskGB),
    string(model.DeployRamdisk),
    string(model.DeployKernel),
    int(model.CPUCount),
    string(model.CPUArch),
    common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "baremetal_node",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "baremetal_node", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanBaremetalNode(values map[string]interface{} ) (*models.BaremetalNode, error) {
    m := models.MakeBaremetalNode()
    
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
    
    if value, ok := values["name"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.Name = castedValue
            

        
    }
    
    if value, ok := values["memory_mb"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.MemoryMB = castedValue
            

        
    }
    
    if value, ok := values["ipmi_username"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.IpmiUsername = castedValue
            

        
    }
    
    if value, ok := values["ipmi_password"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.IpmiPassword = castedValue
            

        
    }
    
    if value, ok := values["ipmi_address"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.IpmiAddress = castedValue
            

        
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
    
    if value, ok := values["disk_gb"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.DiskGB = castedValue
            

        
    }
    
    if value, ok := values["deploy_ramdisk"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.DeployRamdisk = castedValue
            

        
    }
    
    if value, ok := values["deploy_kernel"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.DeployKernel = castedValue
            

        
    }
    
    if value, ok := values["cpu_count"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.CPUCount = castedValue
            

        
    }
    
    if value, ok := values["cpu_arch"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.CPUArch = castedValue
            

        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    
    return m, nil
}

// ListBaremetalNode lists BaremetalNode with list spec.
func ListBaremetalNode(tx *sql.Tx, spec *common.ListSpec) ([]*models.BaremetalNode, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "baremetal_node"
    spec.Fields = BaremetalNodeFields
    spec.RefFields = BaremetalNodeRefFields
    spec.BackRefFields = BaremetalNodeBackRefFields
    result := models.MakeBaremetalNodeSlice()

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
            m, err := scanBaremetalNode(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateBaremetalNode updates a resource
func UpdateBaremetalNode(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateBaremetalNodeQuery = "update `baremetal_node` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateBaremetalNodeQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateBaremetalNodeQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateBaremetalNodeQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateBaremetalNodeQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateBaremetalNodeQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateBaremetalNodeQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateBaremetalNodeQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Name" , "."); ok {
        updateBaremetalNodeQuery += "`name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".MemoryMB" , "."); ok {
        updateBaremetalNodeQuery += "`memory_mb` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IpmiUsername" , "."); ok {
        updateBaremetalNodeQuery += "`ipmi_username` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IpmiPassword" , "."); ok {
        updateBaremetalNodeQuery += "`ipmi_password` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IpmiAddress" , "."); ok {
        updateBaremetalNodeQuery += "`ipmi_address` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateBaremetalNodeQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateBaremetalNodeQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateBaremetalNodeQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateBaremetalNodeQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateBaremetalNodeQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateBaremetalNodeQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateBaremetalNodeQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateBaremetalNodeQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateBaremetalNodeQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateBaremetalNodeQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateBaremetalNodeQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateBaremetalNodeQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateBaremetalNodeQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DiskGB" , "."); ok {
        updateBaremetalNodeQuery += "`disk_gb` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DeployRamdisk" , "."); ok {
        updateBaremetalNodeQuery += "`deploy_ramdisk` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DeployKernel" , "."); ok {
        updateBaremetalNodeQuery += "`deploy_kernel` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".CPUCount" , "."); ok {
        updateBaremetalNodeQuery += "`cpu_count` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".CPUArch" , "."); ok {
        updateBaremetalNodeQuery += "`cpu_arch` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateBaremetalNodeQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateBaremetalNodeQuery += ","
    }
    
    updateBaremetalNodeQuery =
    updateBaremetalNodeQuery[:len(updateBaremetalNodeQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateBaremetalNodeQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateBaremetalNodeQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "baremetal_node", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteBaremetalNode deletes a resource
func DeleteBaremetalNode(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteBaremetalNodeQuery
    selectQuery := "select count(uuid) from baremetal_node where uuid = ?"
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