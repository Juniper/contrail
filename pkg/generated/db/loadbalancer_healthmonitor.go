package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertLoadbalancerHealthmonitorQuery = "insert into `loadbalancer_healthmonitor` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`url_path`,`timeout`,`monitor_type`,`max_retries`,`http_method`,`expected_codes`,`delay`,`admin_state`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteLoadbalancerHealthmonitorQuery = "delete from `loadbalancer_healthmonitor` where uuid = ?"

// LoadbalancerHealthmonitorFields is db columns for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorFields = []string{
   "uuid",
   "share",
   "owner_access",
   "owner",
   "global_access",
   "parent_uuid",
   "parent_type",
   "url_path",
   "timeout",
   "monitor_type",
   "max_retries",
   "http_method",
   "expected_codes",
   "delay",
   "admin_state",
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

// LoadbalancerHealthmonitorRefFields is db reference fields for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorRefFields = map[string][]string{
   
}

// LoadbalancerHealthmonitorBackRefFields is db back reference fields for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorBackRefFields = map[string][]string{
   
}

// LoadbalancerHealthmonitorParentTypes is possible parents for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorParents = []string{
   
   "project",
   
}



// CreateLoadbalancerHealthmonitor inserts LoadbalancerHealthmonitor to DB
func CreateLoadbalancerHealthmonitor(tx *sql.Tx, model *models.LoadbalancerHealthmonitor) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerHealthmonitorQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertLoadbalancerHealthmonitorQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    common.MustJSON(model.Perms2.Share),
    int(model.Perms2.OwnerAccess),
    string(model.Perms2.Owner),
    int(model.Perms2.GlobalAccess),
    string(model.ParentUUID),
    string(model.ParentType),
    string(model.LoadbalancerHealthmonitorProperties.URLPath),
    int(model.LoadbalancerHealthmonitorProperties.Timeout),
    string(model.LoadbalancerHealthmonitorProperties.MonitorType),
    int(model.LoadbalancerHealthmonitorProperties.MaxRetries),
    string(model.LoadbalancerHealthmonitorProperties.HTTPMethod),
    string(model.LoadbalancerHealthmonitorProperties.ExpectedCodes),
    int(model.LoadbalancerHealthmonitorProperties.Delay),
    bool(model.LoadbalancerHealthmonitorProperties.AdminState),
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
        Type: "loadbalancer_healthmonitor",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "loadbalancer_healthmonitor", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanLoadbalancerHealthmonitor(values map[string]interface{} ) (*models.LoadbalancerHealthmonitor, error) {
    m := models.MakeLoadbalancerHealthmonitor()
    
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
    
    if value, ok := values["url_path"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.LoadbalancerHealthmonitorProperties.URLPath = castedValue
            

        
    }
    
    if value, ok := values["timeout"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.LoadbalancerHealthmonitorProperties.Timeout = castedValue
            

        
    }
    
    if value, ok := values["monitor_type"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.LoadbalancerHealthmonitorProperties.MonitorType = models.HealthmonitorType(castedValue)
            

        
    }
    
    if value, ok := values["max_retries"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.LoadbalancerHealthmonitorProperties.MaxRetries = castedValue
            

        
    }
    
    if value, ok := values["http_method"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.LoadbalancerHealthmonitorProperties.HTTPMethod = castedValue
            

        
    }
    
    if value, ok := values["expected_codes"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.LoadbalancerHealthmonitorProperties.ExpectedCodes = castedValue
            

        
    }
    
    if value, ok := values["delay"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.LoadbalancerHealthmonitorProperties.Delay = castedValue
            

        
    }
    
    if value, ok := values["admin_state"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.LoadbalancerHealthmonitorProperties.AdminState = castedValue
            

        
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

// ListLoadbalancerHealthmonitor lists LoadbalancerHealthmonitor with list spec.
func ListLoadbalancerHealthmonitor(tx *sql.Tx, spec *common.ListSpec) ([]*models.LoadbalancerHealthmonitor, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "loadbalancer_healthmonitor"
    spec.Fields = LoadbalancerHealthmonitorFields
    spec.RefFields = LoadbalancerHealthmonitorRefFields
    spec.BackRefFields = LoadbalancerHealthmonitorBackRefFields
    result := models.MakeLoadbalancerHealthmonitorSlice()

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
            m, err := scanLoadbalancerHealthmonitor(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateLoadbalancerHealthmonitor updates a resource
func UpdateLoadbalancerHealthmonitor(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateLoadbalancerHealthmonitorQuery = "update `loadbalancer_healthmonitor` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerHealthmonitorProperties.URLPath" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`url_path` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerHealthmonitorProperties.Timeout" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`timeout` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerHealthmonitorProperties.MonitorType" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`monitor_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerHealthmonitorProperties.MaxRetries" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`max_retries` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerHealthmonitorProperties.HTTPMethod" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`http_method` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerHealthmonitorProperties.ExpectedCodes" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`expected_codes` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerHealthmonitorProperties.Delay" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`delay` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerHealthmonitorProperties.AdminState" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`admin_state` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateLoadbalancerHealthmonitorQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateLoadbalancerHealthmonitorQuery += ","
    }
    
    updateLoadbalancerHealthmonitorQuery =
    updateLoadbalancerHealthmonitorQuery[:len(updateLoadbalancerHealthmonitorQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateLoadbalancerHealthmonitorQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateLoadbalancerHealthmonitorQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "loadbalancer_healthmonitor", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteLoadbalancerHealthmonitor deletes a resource
func DeleteLoadbalancerHealthmonitor(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteLoadbalancerHealthmonitorQuery
    selectQuery := "select count(uuid) from loadbalancer_healthmonitor where uuid = ?"
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