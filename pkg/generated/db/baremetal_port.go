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

const insertBaremetalPortQuery = "insert into `baremetal_port` (`uuid`,`updated_at`,`pxe_enabled`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`node`,`mac_address`,`switch_info`,`switch_id`,`port_id`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`created_at`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteBaremetalPortQuery = "delete from `baremetal_port` where uuid = ?"

// BaremetalPortFields is db columns for BaremetalPort
var BaremetalPortFields = []string{
   "uuid",
   "updated_at",
   "pxe_enabled",
   "share",
   "owner_access",
   "owner",
   "global_access",
   "parent_uuid",
   "parent_type",
   "node",
   "mac_address",
   "switch_info",
   "switch_id",
   "port_id",
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
   "created_at",
   "key_value_pair",
   
}

// BaremetalPortRefFields is db reference fields for BaremetalPort
var BaremetalPortRefFields = map[string][]string{
   
}

// BaremetalPortBackRefFields is db back reference fields for BaremetalPort
var BaremetalPortBackRefFields = map[string][]string{
   
}

// BaremetalPortParentTypes is possible parents for BaremetalPort
var BaremetalPortParents = []string{
   
}



// CreateBaremetalPort inserts BaremetalPort to DB
func CreateBaremetalPort(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.CreateBaremetalPortRequest) error {
    model := request.BaremetalPort
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBaremetalPortQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertBaremetalPortQuery,
    }).Debug("create query")
    _, err = stmt.ExecContext(ctx, string(model.GetUUID()),
    string(model.GetUpdatedAt()),
    bool(model.GetPxeEnabled()),
    common.MustJSON(model.GetPerms2().GetShare()),
    int(model.GetPerms2().GetOwnerAccess()),
    string(model.GetPerms2().GetOwner()),
    int(model.GetPerms2().GetGlobalAccess()),
    string(model.GetParentUUID()),
    string(model.GetParentType()),
    string(model.GetNode()),
    string(model.GetMacAddress()),
    string(model.GetLocalLinkConnection().GetSwitchInfo()),
    string(model.GetLocalLinkConnection().GetSwitchID()),
    string(model.GetLocalLinkConnection().GetPortID()),
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
    string(model.GetCreatedAt()),
    common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "baremetal_port",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "baremetal_port", model.UUID, model.GetPerms2().GetShare())
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanBaremetalPort(values map[string]interface{} ) (*models.BaremetalPort, error) {
    m := models.MakeBaremetalPort()
    
    if value, ok := values["uuid"]; ok {
        
            
               m.UUID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["updated_at"]; ok {
        
            
               m.UpdatedAt = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["pxe_enabled"]; ok {
        
            
               m.PxeEnabled = schema.InterfaceToBool(value)
            
        
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
    
    if value, ok := values["node"]; ok {
        
            
               m.Node = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["mac_address"]; ok {
        
            
               m.MacAddress = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["switch_info"]; ok {
        
            
               m.LocalLinkConnection.SwitchInfo = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["switch_id"]; ok {
        
            
               m.LocalLinkConnection.SwitchID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["port_id"]; ok {
        
            
               m.LocalLinkConnection.PortID = schema.InterfaceToString(value)
            
        
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
    
    if value, ok := values["created_at"]; ok {
        
            
               m.CreatedAt = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    
    return m, nil
}

// ListBaremetalPort lists BaremetalPort with list spec.
func ListBaremetalPort(ctx context.Context, tx *sql.Tx, request *models.ListBaremetalPortRequest) (response *models.ListBaremetalPortResponse, err error) {
    var rows *sql.Rows
    qb := &common.ListQueryBuilder{}
    qb.Auth = common.GetAuthCTX(ctx) 
    spec := request.Spec
    qb.Spec = spec
    qb.Table = "baremetal_port"
    qb.Fields = BaremetalPortFields
    qb.RefFields = BaremetalPortRefFields
    qb.BackRefFields = BaremetalPortBackRefFields
    result := []*models.BaremetalPort{}

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
            m, err := scanBaremetalPort(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    response = &models.ListBaremetalPortResponse{
       BaremetalPorts: result,
    }
    return response, nil
}

// UpdateBaremetalPort updates a resource
func UpdateBaremetalPort(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.UpdateBaremetalPortRequest,
    ) error {
    //TODO
    return nil
}

// DeleteBaremetalPort deletes a resource
func DeleteBaremetalPort(
    ctx context.Context,
    tx *sql.Tx, 
    request *models.DeleteBaremetalPortRequest) error {
    deleteQuery := deleteBaremetalPortQuery
    selectQuery := "select count(uuid) from baremetal_port where uuid = ?"
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