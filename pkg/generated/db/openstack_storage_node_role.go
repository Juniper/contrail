package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertOpenstackStorageNodeRoleQuery = "insert into `openstack_storage_node_role` (`uuid`,`storage_backend_bond_interface_members`,`storage_access_bond_interface_members`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`osd_drives`,`journal_drives`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteOpenstackStorageNodeRoleQuery = "delete from `openstack_storage_node_role` where uuid = ?"

// OpenstackStorageNodeRoleFields is db columns for OpenstackStorageNodeRole
var OpenstackStorageNodeRoleFields = []string{
   "uuid",
   "storage_backend_bond_interface_members",
   "storage_access_bond_interface_members",
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
   "osd_drives",
   "journal_drives",
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

// OpenstackStorageNodeRoleRefFields is db reference fields for OpenstackStorageNodeRole
var OpenstackStorageNodeRoleRefFields = map[string][]string{
   
}

// OpenstackStorageNodeRoleBackRefFields is db back reference fields for OpenstackStorageNodeRole
var OpenstackStorageNodeRoleBackRefFields = map[string][]string{
   
}

// OpenstackStorageNodeRoleParentTypes is possible parents for OpenstackStorageNodeRole
var OpenstackStorageNodeRoleParents = []string{
   
}



// CreateOpenstackStorageNodeRole inserts OpenstackStorageNodeRole to DB
func CreateOpenstackStorageNodeRole(tx *sql.Tx, model *models.OpenstackStorageNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackStorageNodeRoleQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertOpenstackStorageNodeRoleQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    string(model.StorageBackendBondInterfaceMembers),
    string(model.StorageAccessBondInterfaceMembers),
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
    string(model.OsdDrives),
    string(model.JournalDrives),
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
        Type: "openstack_storage_node_role",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "openstack_storage_node_role", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanOpenstackStorageNodeRole(values map[string]interface{} ) (*models.OpenstackStorageNodeRole, error) {
    m := models.MakeOpenstackStorageNodeRole()
    
    if value, ok := values["uuid"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.UUID = castedValue
            

        
    }
    
    if value, ok := values["storage_backend_bond_interface_members"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.StorageBackendBondInterfaceMembers = castedValue
            

        
    }
    
    if value, ok := values["storage_access_bond_interface_members"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.StorageAccessBondInterfaceMembers = castedValue
            

        
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
    
    if value, ok := values["osd_drives"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.OsdDrives = castedValue
            

        
    }
    
    if value, ok := values["journal_drives"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.JournalDrives = castedValue
            

        
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

// ListOpenstackStorageNodeRole lists OpenstackStorageNodeRole with list spec.
func ListOpenstackStorageNodeRole(tx *sql.Tx, spec *common.ListSpec) ([]*models.OpenstackStorageNodeRole, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "openstack_storage_node_role"
    spec.Fields = OpenstackStorageNodeRoleFields
    spec.RefFields = OpenstackStorageNodeRoleRefFields
    spec.BackRefFields = OpenstackStorageNodeRoleBackRefFields
    result := models.MakeOpenstackStorageNodeRoleSlice()

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
            m, err := scanOpenstackStorageNodeRole(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateOpenstackStorageNodeRole updates a resource
func UpdateOpenstackStorageNodeRole(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateOpenstackStorageNodeRoleQuery = "update `openstack_storage_node_role` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".StorageBackendBondInterfaceMembers" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`storage_backend_bond_interface_members` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".StorageAccessBondInterfaceMembers" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`storage_access_bond_interface_members` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningState" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`provisioning_state` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningStartTime" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`provisioning_start_time` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningProgressStage" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`provisioning_progress_stage` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningProgress" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`provisioning_progress` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ProvisioningLog" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`provisioning_log` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".OsdDrives" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`osd_drives` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".JournalDrives" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`journal_drives` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateOpenstackStorageNodeRoleQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateOpenstackStorageNodeRoleQuery += ","
    }
    
    updateOpenstackStorageNodeRoleQuery =
    updateOpenstackStorageNodeRoleQuery[:len(updateOpenstackStorageNodeRoleQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateOpenstackStorageNodeRoleQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateOpenstackStorageNodeRoleQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "openstack_storage_node_role", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteOpenstackStorageNodeRole deletes a resource
func DeleteOpenstackStorageNodeRole(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteOpenstackStorageNodeRoleQuery
    selectQuery := "select count(uuid) from openstack_storage_node_role where uuid = ?"
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