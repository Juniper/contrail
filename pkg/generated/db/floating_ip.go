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

const insertFloatingIPQuery = "insert into `floating_ip` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`floating_ip_traffic_direction`,`port_mappings`,`floating_ip_port_mappings_enable`,`floating_ip_is_virtual_ip`,`floating_ip_fixed_ip_address`,`floating_ip_address_family`,`floating_ip_address`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteFloatingIPQuery = "delete from `floating_ip` where uuid = ?"

// FloatingIPFields is db columns for FloatingIP
var FloatingIPFields = []string{
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
   "floating_ip_traffic_direction",
   "port_mappings",
   "floating_ip_port_mappings_enable",
   "floating_ip_is_virtual_ip",
   "floating_ip_fixed_ip_address",
   "floating_ip_address_family",
   "floating_ip_address",
   "display_name",
   "key_value_pair",
   
}

// FloatingIPRefFields is db reference fields for FloatingIP
var FloatingIPRefFields = map[string][]string{
   
    "project": []string{
        // <schema.Schema Value>
        
    },
   
    "virtual_machine_interface": []string{
        // <schema.Schema Value>
        
    },
   
}

// FloatingIPBackRefFields is db back reference fields for FloatingIP
var FloatingIPBackRefFields = map[string][]string{
   
}

// FloatingIPParentTypes is possible parents for FloatingIP
var FloatingIPParents = []string{
   
   "instance_ip",
   
   "floating_ip_pool",
   
}


const insertFloatingIPProjectQuery = "insert into `ref_floating_ip_project` (`from`, `to` ) values (?, ?);"

const insertFloatingIPVirtualMachineInterfaceQuery = "insert into `ref_floating_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"


// CreateFloatingIP inserts FloatingIP to DB
func CreateFloatingIP(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.CreateFloatingIPRequest) error {
    model := request.FloatingIP
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFloatingIPQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertFloatingIPQuery,
    }).Debug("create query")
    _, err = stmt.ExecContext(ctx, string(model.GetUUID()),
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
    string(model.GetFloatingIPTrafficDirection()),
    common.MustJSON(model.GetFloatingIPPortMappings().GetPortMappings()),
    bool(model.GetFloatingIPPortMappingsEnable()),
    bool(model.GetFloatingIPIsVirtualIP()),
    string(model.GetFloatingIPFixedIPAddress()),
    string(model.GetFloatingIPAddressFamily()),
    string(model.GetFloatingIPAddress()),
    string(model.GetDisplayName()),
    common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    stmtProjectRef, err := tx.Prepare(insertFloatingIPProjectQuery)
	if err != nil {
        return errors.Wrap(err,"preparing ProjectRefs create statement failed")
	}
    defer stmtProjectRef.Close()
    for _, ref := range model.ProjectRefs {
       
        _, err = stmtProjectRef.ExecContext(ctx, model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"ProjectRefs create failed")
        }
    }
    
    stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertFloatingIPVirtualMachineInterfaceQuery)
	if err != nil {
        return errors.Wrap(err,"preparing VirtualMachineInterfaceRefs create statement failed")
	}
    defer stmtVirtualMachineInterfaceRef.Close()
    for _, ref := range model.VirtualMachineInterfaceRefs {
       
        _, err = stmtVirtualMachineInterfaceRef.ExecContext(ctx, model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"VirtualMachineInterfaceRefs create failed")
        }
    }
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "floating_ip",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "floating_ip", model.UUID, model.GetPerms2().GetShare())
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanFloatingIP(values map[string]interface{} ) (*models.FloatingIP, error) {
    m := models.MakeFloatingIP()
    
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
    
    if value, ok := values["floating_ip_traffic_direction"]; ok {
        
            
               m.FloatingIPTrafficDirection = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["port_mappings"]; ok {
        
            json.Unmarshal(value.([]byte), &m.FloatingIPPortMappings.PortMappings)
        
    }
    
    if value, ok := values["floating_ip_port_mappings_enable"]; ok {
        
            
               m.FloatingIPPortMappingsEnable = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["floating_ip_is_virtual_ip"]; ok {
        
            
               m.FloatingIPIsVirtualIP = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["floating_ip_fixed_ip_address"]; ok {
        
            
               m.FloatingIPFixedIPAddress = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["floating_ip_address_family"]; ok {
        
            
               m.FloatingIPAddressFamily = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["floating_ip_address"]; ok {
        
            
               m.FloatingIPAddress = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["display_name"]; ok {
        
            
               m.DisplayName = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    if value, ok := values["ref_project"]; ok {
        var references []interface{}
        stringValue := schema.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &references )
        for _, reference := range references {
            referenceMap, ok := reference.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := schema.InterfaceToString(referenceMap["to"])
            if uuid == "" {
                continue
            }
            referenceModel := &models.FloatingIPProjectRef{}
            referenceModel.UUID = uuid
            m.ProjectRefs = append(m.ProjectRefs, referenceModel)
            
        }
    }
    
    if value, ok := values["ref_virtual_machine_interface"]; ok {
        var references []interface{}
        stringValue := schema.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &references )
        for _, reference := range references {
            referenceMap, ok := reference.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := schema.InterfaceToString(referenceMap["to"])
            if uuid == "" {
                continue
            }
            referenceModel := &models.FloatingIPVirtualMachineInterfaceRef{}
            referenceModel.UUID = uuid
            m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)
            
        }
    }
    
    
    return m, nil
}

// ListFloatingIP lists FloatingIP with list spec.
func ListFloatingIP(ctx context.Context, tx *sql.Tx, request *models.ListFloatingIPRequest) (response *models.ListFloatingIPResponse, err error) {
    var rows *sql.Rows
    qb := &common.ListQueryBuilder{}
    qb.Auth = common.GetAuthCTX(ctx) 
    spec := request.Spec
    qb.Spec = spec
    qb.Table = "floating_ip"
    qb.Fields = FloatingIPFields
    qb.RefFields = FloatingIPRefFields
    qb.BackRefFields = FloatingIPBackRefFields
    result := []*models.FloatingIP{}

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
            m, err := scanFloatingIP(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    response = &models.ListFloatingIPResponse{
       FloatingIPs: result,
    }
    return response, nil
}

// UpdateFloatingIP updates a resource
func UpdateFloatingIP(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.UpdateFloatingIPRequest,
    ) error {
    //TODO
    return nil
}

// DeleteFloatingIP deletes a resource
func DeleteFloatingIP(
    ctx context.Context,
    tx *sql.Tx, 
    request *models.DeleteFloatingIPRequest) error {
    deleteQuery := deleteFloatingIPQuery
    selectQuery := "select count(uuid) from floating_ip where uuid = ?"
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