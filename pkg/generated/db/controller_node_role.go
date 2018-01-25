package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertControllerNodeRoleQuery = "insert into `controller_node_role` (`uuid`,`storage_management_bond_interface_members`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`performance_drives`,`parent_uuid`,`parent_type`,`internalapi_bond_interface_members`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`capacity_drives`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteControllerNodeRoleQuery = "delete from `controller_node_role` where uuid = ?"

// ControllerNodeRoleFields is db columns for ControllerNodeRole
var ControllerNodeRoleFields = []string{
   "uuid",
   "storage_management_bond_interface_members",
   "provisioning_state",
   "provisioning_start_time",
   "provisioning_progress_stage",
   "provisioning_progress",
   "provisioning_log",
   "share",
   "owner_access",
   "owner",
   "global_access",
   "performance_drives",
   "parent_uuid",
   "parent_type",
   "internalapi_bond_interface_members",
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
   "capacity_drives",
   "key_value_pair",
   
}

// ControllerNodeRoleRefFields is db reference fields for ControllerNodeRole
var ControllerNodeRoleRefFields = map[string][]string{
   
}

// ControllerNodeRoleBackRefFields is db back reference fields for ControllerNodeRole
var ControllerNodeRoleBackRefFields = map[string][]string{
   
}

// ControllerNodeRoleParentTypes is possible parents for ControllerNodeRole
var ControllerNodeRoleParents = []string{
   
}



// CreateControllerNodeRole inserts ControllerNodeRole to DB
func CreateControllerNodeRole(tx *sql.Tx, model *models.ControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertControllerNodeRoleQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertControllerNodeRoleQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    string(model.StorageManagementBondInterfaceMembers),
    string(model.ProvisioningState),
    string(model.ProvisioningStartTime),
    string(model.ProvisioningProgressStage),
    int(model.ProvisioningProgress),
    string(model.ProvisioningLog),
    common.MustJSON(model.Perms2.Share),
    int(model.Perms2.OwnerAccess),
    string(model.Perms2.Owner),
    int(model.Perms2.GlobalAccess),
    string(model.PerformanceDrives),
    string(model.ParentUUID),
    string(model.ParentType),
    string(model.InternalapiBondInterfaceMembers),
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
    string(model.CapacityDrives),
    common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "controller_node_role",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "controller_node_role", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanControllerNodeRole(values map[string]interface{} ) (*models.ControllerNodeRole, error) {
    m := models.MakeControllerNodeRole()
    
    if value, ok := values["uuid"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.UUID = castedValue
            

        
    }
    
    if value, ok := values["storage_management_bond_interface_members"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.StorageManagementBondInterfaceMembers = castedValue
            

        
    }
    
    if value, ok := values["provisioning_state"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ProvisioningState = castedValue
            

        
    }
    
    if value, ok := values["provisioning_start_time"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ProvisioningStartTime = castedValue
            

        
    }
    
    if value, ok := values["provisioning_progress_stage"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ProvisioningProgressStage = castedValue
            

        
    }
    
    if value, ok := values["provisioning_progress"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.ProvisioningProgress = castedValue
            

        
    }
    
    if value, ok := values["provisioning_log"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ProvisioningLog = castedValue
            

        
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
    
    if value, ok := values["performance_drives"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.PerformanceDrives = castedValue
            

        
    }
    
    if value, ok := values["parent_uuid"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ParentUUID = castedValue
            

        
    }
    
    if value, ok := values["parent_type"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.ParentType = castedValue
            

        
    }
    
    if value, ok := values["internalapi_bond_interface_members"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.InternalapiBondInterfaceMembers = castedValue
            

        
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
    
    if value, ok := values["capacity_drives"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.CapacityDrives = castedValue
            

        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    
    return m, nil
}

// ListControllerNodeRole lists ControllerNodeRole with list spec.
func ListControllerNodeRole(tx *sql.Tx, spec *common.ListSpec) ([]*models.ControllerNodeRole, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "controller_node_role"
    spec.Fields = ControllerNodeRoleFields
    spec.RefFields = ControllerNodeRoleRefFields
    spec.BackRefFields = ControllerNodeRoleBackRefFields
    result := models.MakeControllerNodeRoleSlice()

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
            m, err := scanControllerNodeRole(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateControllerNodeRole updates a resource
func UpdateControllerNodeRole(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateControllerNodeRoleQuery = "update `controller_node_role` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateControllerNodeRoleQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".StorageManagementBondInterfaceMembers" , "."); ok {
        updateControllerNodeRoleQuery += "`storage_management_bond_interface_members` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningState" , "."); ok {
        updateControllerNodeRoleQuery += "`provisioning_state` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningStartTime" , "."); ok {
        updateControllerNodeRoleQuery += "`provisioning_start_time` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningProgressStage" , "."); ok {
        updateControllerNodeRoleQuery += "`provisioning_progress_stage` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningProgress" , "."); ok {
        updateControllerNodeRoleQuery += "`provisioning_progress` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningLog" , "."); ok {
        updateControllerNodeRoleQuery += "`provisioning_log` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateControllerNodeRoleQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateControllerNodeRoleQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateControllerNodeRoleQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateControllerNodeRoleQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".PerformanceDrives" , "."); ok {
        updateControllerNodeRoleQuery += "`performance_drives` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateControllerNodeRoleQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateControllerNodeRoleQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".InternalapiBondInterfaceMembers" , "."); ok {
        updateControllerNodeRoleQuery += "`internalapi_bond_interface_members` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateControllerNodeRoleQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateControllerNodeRoleQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateControllerNodeRoleQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateControllerNodeRoleQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateControllerNodeRoleQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateControllerNodeRoleQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateControllerNodeRoleQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateControllerNodeRoleQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateControllerNodeRoleQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateControllerNodeRoleQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateControllerNodeRoleQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateControllerNodeRoleQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateControllerNodeRoleQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".CapacityDrives" , "."); ok {
        updateControllerNodeRoleQuery += "`capacity_drives` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateControllerNodeRoleQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateControllerNodeRoleQuery += ","
    }
    
    updateControllerNodeRoleQuery =
    updateControllerNodeRoleQuery[:len(updateControllerNodeRoleQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateControllerNodeRoleQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateControllerNodeRoleQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "controller_node_role", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteControllerNodeRole deletes a resource
func DeleteControllerNodeRole(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteControllerNodeRoleQuery
    selectQuery := "select count(uuid) from controller_node_role where uuid = ?"
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