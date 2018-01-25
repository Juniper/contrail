package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertContrailAnalyticsDatabaseNodeRoleQuery = "insert into `contrail_analytics_database_node_role` (`uuid`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteContrailAnalyticsDatabaseNodeRoleQuery = "delete from `contrail_analytics_database_node_role` where uuid = ?"

// ContrailAnalyticsDatabaseNodeRoleFields is db columns for ContrailAnalyticsDatabaseNodeRole
var ContrailAnalyticsDatabaseNodeRoleFields = []string{
   "uuid",
   "provisioning_state",
   "provisioning_start_time",
   "provisioning_progress_stage",
   "provisioning_progress",
   "provisioning_log",
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

// ContrailAnalyticsDatabaseNodeRoleRefFields is db reference fields for ContrailAnalyticsDatabaseNodeRole
var ContrailAnalyticsDatabaseNodeRoleRefFields = map[string][]string{
   
}

// ContrailAnalyticsDatabaseNodeRoleBackRefFields is db back reference fields for ContrailAnalyticsDatabaseNodeRole
var ContrailAnalyticsDatabaseNodeRoleBackRefFields = map[string][]string{
   
}

// ContrailAnalyticsDatabaseNodeRoleParentTypes is possible parents for ContrailAnalyticsDatabaseNodeRole
var ContrailAnalyticsDatabaseNodeRoleParents = []string{
   
}



// CreateContrailAnalyticsDatabaseNodeRole inserts ContrailAnalyticsDatabaseNodeRole to DB
func CreateContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, model *models.ContrailAnalyticsDatabaseNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsDatabaseNodeRoleQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertContrailAnalyticsDatabaseNodeRoleQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    string(model.ProvisioningState),
    string(model.ProvisioningStartTime),
    string(model.ProvisioningProgressStage),
    int(model.ProvisioningProgress),
    string(model.ProvisioningLog),
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
        Type: "contrail_analytics_database_node_role",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "contrail_analytics_database_node_role", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanContrailAnalyticsDatabaseNodeRole(values map[string]interface{} ) (*models.ContrailAnalyticsDatabaseNodeRole, error) {
    m := models.MakeContrailAnalyticsDatabaseNodeRole()
    
    if value, ok := values["uuid"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.UUID = castedValue
            

        
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
    
    
    
    return m, nil
}

// ListContrailAnalyticsDatabaseNodeRole lists ContrailAnalyticsDatabaseNodeRole with list spec.
func ListContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, spec *common.ListSpec) ([]*models.ContrailAnalyticsDatabaseNodeRole, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "contrail_analytics_database_node_role"
    spec.Fields = ContrailAnalyticsDatabaseNodeRoleFields
    spec.RefFields = ContrailAnalyticsDatabaseNodeRoleRefFields
    spec.BackRefFields = ContrailAnalyticsDatabaseNodeRoleBackRefFields
    result := models.MakeContrailAnalyticsDatabaseNodeRoleSlice()

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
            m, err := scanContrailAnalyticsDatabaseNodeRole(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateContrailAnalyticsDatabaseNodeRole updates a resource
func UpdateContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateContrailAnalyticsDatabaseNodeRoleQuery = "update `contrail_analytics_database_node_role` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningState" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`provisioning_state` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningStartTime" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`provisioning_start_time` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningProgressStage" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`provisioning_progress_stage` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningProgress" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`provisioning_progress` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningLog" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`provisioning_log` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateContrailAnalyticsDatabaseNodeRoleQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateContrailAnalyticsDatabaseNodeRoleQuery += ","
    }
    
    updateContrailAnalyticsDatabaseNodeRoleQuery =
    updateContrailAnalyticsDatabaseNodeRoleQuery[:len(updateContrailAnalyticsDatabaseNodeRoleQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateContrailAnalyticsDatabaseNodeRoleQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateContrailAnalyticsDatabaseNodeRoleQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "contrail_analytics_database_node_role", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteContrailAnalyticsDatabaseNodeRole deletes a resource
func DeleteContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteContrailAnalyticsDatabaseNodeRoleQuery
    selectQuery := "select count(uuid) from contrail_analytics_database_node_role where uuid = ?"
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