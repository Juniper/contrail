package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualDNSQuery = "insert into `virtual_DNS` (`reverse_resolution`,`record_order`,`next_virtual_DNS`,`floating_ip_record`,`external_visible`,`dynamic_records_from_client`,`domain_name`,`default_ttl_seconds`,`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteVirtualDNSQuery = "delete from `virtual_DNS` where uuid = ?"

// VirtualDNSFields is db columns for VirtualDNS
var VirtualDNSFields = []string{
   "reverse_resolution",
   "record_order",
   "next_virtual_DNS",
   "floating_ip_record",
   "external_visible",
   "dynamic_records_from_client",
   "domain_name",
   "default_ttl_seconds",
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

// VirtualDNSRefFields is db reference fields for VirtualDNS
var VirtualDNSRefFields = map[string][]string{
   
}

// VirtualDNSBackRefFields is db back reference fields for VirtualDNS
var VirtualDNSBackRefFields = map[string][]string{
   
   
   "virtual_DNS_record": []string{
        "record_type",
        "record_ttl_seconds",
        "record_name",
        "record_mx_preference",
        "record_data",
        "record_class",
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
   
}

// VirtualDNSParentTypes is possible parents for VirtualDNS
var VirtualDNSParents = []string{
   
   "domain",
   
}



// CreateVirtualDNS inserts VirtualDNS to DB
func CreateVirtualDNS(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.CreateVirtualDNSRequest) error {
    model := request.VirtualDNS
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualDNSQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertVirtualDNSQuery,
    }).Debug("create query")
    _, err = stmt.ExecContext(ctx, bool(model.GetVirtualDNSData().GetReverseResolution()),
    string(model.GetVirtualDNSData().GetRecordOrder()),
    string(model.GetVirtualDNSData().GetNextVirtualDNS()),
    string(model.GetVirtualDNSData().GetFloatingIPRecord()),
    bool(model.GetVirtualDNSData().GetExternalVisible()),
    bool(model.GetVirtualDNSData().GetDynamicRecordsFromClient()),
    string(model.GetVirtualDNSData().GetDomainName()),
    int(model.GetVirtualDNSData().GetDefaultTTLSeconds()),
    string(model.GetUUID()),
    common.MustJSON(model.GetPerms2().GetShare()),
    int(model.GetPerms2().GetOwnerAccess()),
    string(model.GetPerms2().GetOwner()),
    int(model.GetPerms2().GetGlobalAccess()),
    string(model.GetParentUUID()),
    string(model.GetParentType()),
    bool(model.GetIDPerms().GetUserVisible()),
    int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
    string(model.GetIDPerms().GetPermissions().GetOwner()),
    int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
    int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
    string(model.GetIDPerms().GetPermissions().GetGroup()),
    string(model.GetIDPerms().GetLastModified()),
    bool(model.GetIDPerms().GetEnable()),
    string(model.GetIDPerms().GetDescription()),
    string(model.GetIDPerms().GetCreator()),
    string(model.GetIDPerms().GetCreated()),
    common.MustJSON(model.GetFQName()),
    string(model.GetDisplayName()),
    common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "virtual_DNS",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "virtual_DNS", model.UUID, model.GetPerms2().GetShare())
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanVirtualDNS(values map[string]interface{} ) (*models.VirtualDNS, error) {
    m := models.MakeVirtualDNS()
    
    if value, ok := values["reverse_resolution"]; ok {
        
            
               m.VirtualDNSData.ReverseResolution = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["record_order"]; ok {
        
            
               m.VirtualDNSData.RecordOrder = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["next_virtual_DNS"]; ok {
        
            
               m.VirtualDNSData.NextVirtualDNS = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["floating_ip_record"]; ok {
        
            
               m.VirtualDNSData.FloatingIPRecord = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["external_visible"]; ok {
        
            
               m.VirtualDNSData.ExternalVisible = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["dynamic_records_from_client"]; ok {
        
            
               m.VirtualDNSData.DynamicRecordsFromClient = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["domain_name"]; ok {
        
            
               m.VirtualDNSData.DomainName = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["default_ttl_seconds"]; ok {
        
            
               m.VirtualDNSData.DefaultTTLSeconds = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["uuid"]; ok {
        
            
               m.UUID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["share"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Perms2.Share)
        
    }
    
    if value, ok := values["owner_access"]; ok {
        
            
               m.Perms2.OwnerAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["owner"]; ok {
        
            
               m.Perms2.Owner = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["global_access"]; ok {
        
            
               m.Perms2.GlobalAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["parent_uuid"]; ok {
        
            
               m.ParentUUID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["parent_type"]; ok {
        
            
               m.ParentType = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["user_visible"]; ok {
        
            
               m.IDPerms.UserVisible = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["permissions_owner_access"]; ok {
        
            
               m.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["permissions_owner"]; ok {
        
            
               m.IDPerms.Permissions.Owner = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["other_access"]; ok {
        
            
               m.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["group_access"]; ok {
        
            
               m.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["group"]; ok {
        
            
               m.IDPerms.Permissions.Group = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["last_modified"]; ok {
        
            
               m.IDPerms.LastModified = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["enable"]; ok {
        
            
               m.IDPerms.Enable = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["description"]; ok {
        
            
               m.IDPerms.Description = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["creator"]; ok {
        
            
               m.IDPerms.Creator = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["created"]; ok {
        
            
               m.IDPerms.Created = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["fq_name"]; ok {
        
            json.Unmarshal(value.([]byte), &m.FQName)
        
    }
    
    if value, ok := values["display_name"]; ok {
        
            
               m.DisplayName = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    
    
    if value, ok := values["backref_virtual_DNS_record"]; ok {
        var childResources []interface{}
        stringValue := schema.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := schema.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeVirtualDNSRecord()
            m.VirtualDNSRecords = append(m.VirtualDNSRecords, childModel)

            
                if propertyValue, ok := childResourceMap["record_type"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualDNSRecordData.RecordType = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["record_ttl_seconds"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualDNSRecordData.RecordTTLSeconds = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["record_name"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualDNSRecordData.RecordName = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["record_mx_preference"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualDNSRecordData.RecordMXPreference = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["record_data"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualDNSRecordData.RecordData = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["record_class"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualDNSRecordData.RecordClass = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        childModel.UUID = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        childModel.ParentUUID = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        childModel.ParentType = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        childModel.DisplayName = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
        }
    }
    
    return m, nil
}

// ListVirtualDNS lists VirtualDNS with list spec.
func ListVirtualDNS(ctx context.Context, tx *sql.Tx, request *models.ListVirtualDNSRequest) (response *models.ListVirtualDNSResponse, err error) {
    var rows *sql.Rows
    qb := &common.ListQueryBuilder{}
    qb.Auth = common.GetAuthCTX(ctx) 
    spec := request.Spec
    qb.Spec = spec
    qb.Table = "virtual_DNS"
    qb.Fields = VirtualDNSFields
    qb.RefFields = VirtualDNSRefFields
    qb.BackRefFields = VirtualDNSBackRefFields
    result := []*models.VirtualDNS{}

    if spec.ParentFQName != nil {
       parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
       if err != nil {
           return nil, errors.Wrap(err, "can't find parents")
       }
       spec.Filters = common.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
    }

    query := qb.BuildQuery()
    columns := qb.Columns
    values := qb.Values
    log.WithFields(log.Fields{
        "listSpec": spec,
        "query": query,
    }).Debug("select query")
    rows, err = tx.QueryContext(ctx, query, values...)
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
            m, err := scanVirtualDNS(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    response = &models.ListVirtualDNSResponse{
       VirtualDNSs: result,
    }
    return response, nil
}

// UpdateVirtualDNS updates a resource
func UpdateVirtualDNS(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.UpdateVirtualDNSRequest,
    ) error {
    //TODO
    return nil
}

// DeleteVirtualDNS deletes a resource
func DeleteVirtualDNS(
    ctx context.Context,
    tx *sql.Tx, 
    request *models.DeleteVirtualDNSRequest) error {
    deleteQuery := deleteVirtualDNSQuery
    selectQuery := "select count(uuid) from virtual_DNS where uuid = ?"
    var err error
    var count int
    uuid := request.ID
    auth := common.GetAuthCTX(ctx)
    if auth.IsAdmin() {
        row := tx.QueryRowContext(ctx, selectQuery, uuid)
        if err != nil {
            return errors.Wrap(err, "not found")
        }
        row.Scan(&count)
        if count == 0 {
           return errors.New("Not found")
        }
        _, err = tx.ExecContext(ctx, deleteQuery, uuid)
    }else{
        deleteQuery += " and owner = ?"
        selectQuery += " and owner = ?"
        row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID() )
        if err != nil {
            return errors.Wrap(err, "not found")
        }
        row.Scan(&count)
        if count == 0 {
           return errors.New("Not found")
        }
        _, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID() )
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