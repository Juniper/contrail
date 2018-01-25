package db

import (
    "database/sql"
    "encoding/json"
    "github.com/Juniper/contrail/pkg/common"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/pkg/errors"

    log "github.com/sirupsen/logrus"
)

const insertLoadbalancerMemberQuery = "insert into `loadbalancer_member` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`weight`,`status_description`,`status`,`protocol_port`,`admin_state`,`address`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteLoadbalancerMemberQuery = "delete from `loadbalancer_member` where uuid = ?"

// LoadbalancerMemberFields is db columns for LoadbalancerMember
var LoadbalancerMemberFields = []string{
   "uuid",
   "share",
   "owner_access",
   "owner",
   "global_access",
   "parent_uuid",
   "parent_type",
   "weight",
   "status_description",
   "status",
   "protocol_port",
   "admin_state",
   "address",
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

// LoadbalancerMemberRefFields is db reference fields for LoadbalancerMember
var LoadbalancerMemberRefFields = map[string][]string{
   
}

// LoadbalancerMemberBackRefFields is db back reference fields for LoadbalancerMember
var LoadbalancerMemberBackRefFields = map[string][]string{
   
}

// LoadbalancerMemberParentTypes is possible parents for LoadbalancerMember
var LoadbalancerMemberParents = []string{
   
   "loadbalancer_pool",
   
}



// CreateLoadbalancerMember inserts LoadbalancerMember to DB
func CreateLoadbalancerMember(tx *sql.Tx, model *models.LoadbalancerMember) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerMemberQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertLoadbalancerMemberQuery,
    }).Debug("create query")
    _, err = stmt.Exec(string(model.UUID),
    common.MustJSON(model.Perms2.Share),
    int(model.Perms2.OwnerAccess),
    string(model.Perms2.Owner),
    int(model.Perms2.GlobalAccess),
    string(model.ParentUUID),
    string(model.ParentType),
    int(model.LoadbalancerMemberProperties.Weight),
    string(model.LoadbalancerMemberProperties.StatusDescription),
    string(model.LoadbalancerMemberProperties.Status),
    int(model.LoadbalancerMemberProperties.ProtocolPort),
    bool(model.LoadbalancerMemberProperties.AdminState),
    string(model.LoadbalancerMemberProperties.Address),
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
        Type: "loadbalancer_member",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "loadbalancer_member", model.UUID, model.Perms2.Share)
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanLoadbalancerMember(values map[string]interface{} ) (*models.LoadbalancerMember, error) {
    m := models.MakeLoadbalancerMember()
    
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
    
    if value, ok := values["weight"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.LoadbalancerMemberProperties.Weight = castedValue
            

        
    }
    
    if value, ok := values["status_description"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.LoadbalancerMemberProperties.StatusDescription = castedValue
            

        
    }
    
    if value, ok := values["status"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.LoadbalancerMemberProperties.Status = castedValue
            

        
    }
    
    if value, ok := values["protocol_port"]; ok {
        
            
                castedValue := common.InterfaceToInt(value)
            
            
                m.LoadbalancerMemberProperties.ProtocolPort = castedValue
            

        
    }
    
    if value, ok := values["admin_state"]; ok {
        
            
                castedValue := common.InterfaceToBool(value)
            
            
                m.LoadbalancerMemberProperties.AdminState = castedValue
            

        
    }
    
    if value, ok := values["address"]; ok {
        
            
                castedValue := common.InterfaceToString(value)
            
            
                m.LoadbalancerMemberProperties.Address = models.IpAddressType(castedValue)
            

        
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

// ListLoadbalancerMember lists LoadbalancerMember with list spec.
func ListLoadbalancerMember(tx *sql.Tx, spec *common.ListSpec) ([]*models.LoadbalancerMember, error) {
    var rows *sql.Rows
    var err error
    //TODO (check input)
    spec.Table = "loadbalancer_member"
    spec.Fields = LoadbalancerMemberFields
    spec.RefFields = LoadbalancerMemberRefFields
    spec.BackRefFields = LoadbalancerMemberBackRefFields
    result := models.MakeLoadbalancerMemberSlice()

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
            m, err := scanLoadbalancerMember(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    return result, nil
}

// UpdateLoadbalancerMember updates a resource
func UpdateLoadbalancerMember(tx *sql.Tx, uuid string, model map[string]interface{}) error {
    // Prepare statement for updating data
    var updateLoadbalancerMemberQuery = "update `loadbalancer_member` set "

    updatedValues := make([]interface{}, 0)
    
    if value, ok := common.GetValueByPath(model, ".UUID" , "."); ok {
        updateLoadbalancerMemberQuery += "`uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Share" , "."); ok {
        updateLoadbalancerMemberQuery += "`share` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess" , "."); ok {
        updateLoadbalancerMemberQuery += "`owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.Owner" , "."); ok {
        updateLoadbalancerMemberQuery += "`owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess" , "."); ok {
        updateLoadbalancerMemberQuery += "`global_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentUUID" , "."); ok {
        updateLoadbalancerMemberQuery += "`parent_uuid` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".ParentType" , "."); ok {
        updateLoadbalancerMemberQuery += "`parent_type` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerMemberProperties.Weight" , "."); ok {
        updateLoadbalancerMemberQuery += "`weight` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerMemberProperties.StatusDescription" , "."); ok {
        updateLoadbalancerMemberQuery += "`status_description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerMemberProperties.Status" , "."); ok {
        updateLoadbalancerMemberQuery += "`status` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerMemberProperties.ProtocolPort" , "."); ok {
        updateLoadbalancerMemberQuery += "`protocol_port` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerMemberProperties.AdminState" , "."); ok {
        updateLoadbalancerMemberQuery += "`admin_state` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".LoadbalancerMemberProperties.Address" , "."); ok {
        updateLoadbalancerMemberQuery += "`address` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible" , "."); ok {
        updateLoadbalancerMemberQuery += "`user_visible` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess" , "."); ok {
        updateLoadbalancerMemberQuery += "`permissions_owner_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner" , "."); ok {
        updateLoadbalancerMemberQuery += "`permissions_owner` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess" , "."); ok {
        updateLoadbalancerMemberQuery += "`other_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess" , "."); ok {
        updateLoadbalancerMemberQuery += "`group_access` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group" , "."); ok {
        updateLoadbalancerMemberQuery += "`group` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified" , "."); ok {
        updateLoadbalancerMemberQuery += "`last_modified` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Enable" , "."); ok {
        updateLoadbalancerMemberQuery += "`enable` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToBool(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Description" , "."); ok {
        updateLoadbalancerMemberQuery += "`description` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Creator" , "."); ok {
        updateLoadbalancerMemberQuery += "`creator` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".IDPerms.Created" , "."); ok {
        updateLoadbalancerMemberQuery += "`created` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".FQName" , "."); ok {
        updateLoadbalancerMemberQuery += "`fq_name` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".DisplayName" , "."); ok {
        updateLoadbalancerMemberQuery += "`display_name` = ?"
        
            updatedValues = append(updatedValues, common.InterfaceToString(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair" , "."); ok {
        updateLoadbalancerMemberQuery += "`key_value_pair` = ?"
        
            updatedValues = append(updatedValues, common.MustJSON(value))
        
        updateLoadbalancerMemberQuery += ","
    }
    
    updateLoadbalancerMemberQuery =
    updateLoadbalancerMemberQuery[:len(updateLoadbalancerMemberQuery)-1] + " where `uuid` = ? ;"
    updatedValues = append(updatedValues, string(uuid))
    stmt, err := tx.Prepare(updateLoadbalancerMemberQuery)
    if err != nil {
        return errors.Wrap(err, "preparing update statement failed")
    }
    defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": updateLoadbalancerMemberQuery,
    }).Debug("update query")
    _, err = stmt.Exec( updatedValues... )
    if err != nil {
        return errors.Wrap(err, "update failed")
    }

    
    share, ok := common.GetValueByPath(model, ".Perms2.Share" , ".")
    if ok {
        err = common.UpdateSharing(tx, "loadbalancer_member", string(uuid), share.([]interface{}))
        if err != nil {
            return err
        }
    }

    log.WithFields(log.Fields{
        "model": model,
    }).Debug("updated")
    return err
}

// DeleteLoadbalancerMember deletes a resource
func DeleteLoadbalancerMember(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
    deleteQuery := deleteLoadbalancerMemberQuery
    selectQuery := "select count(uuid) from loadbalancer_member where uuid = ?"
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