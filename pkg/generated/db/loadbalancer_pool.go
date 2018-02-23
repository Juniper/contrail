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

const insertLoadbalancerPoolQuery = "insert into `loadbalancer_pool` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`loadbalancer_pool_provider`,`subnet_id`,`status_description`,`status`,`session_persistence`,`protocol`,`persistence_cookie_name`,`loadbalancer_method`,`admin_state`,`key_value_pair`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteLoadbalancerPoolQuery = "delete from `loadbalancer_pool` where uuid = ?"

// LoadbalancerPoolFields is db columns for LoadbalancerPool
var LoadbalancerPoolFields = []string{
   "uuid",
   "share",
   "owner_access",
   "owner",
   "global_access",
   "parent_uuid",
   "parent_type",
   "loadbalancer_pool_provider",
   "subnet_id",
   "status_description",
   "status",
   "session_persistence",
   "protocol",
   "persistence_cookie_name",
   "loadbalancer_method",
   "admin_state",
   "key_value_pair",
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
   "annotations_key_value_pair",
   
}

// LoadbalancerPoolRefFields is db reference fields for LoadbalancerPool
var LoadbalancerPoolRefFields = map[string][]string{
   
    "loadbalancer_healthmonitor": []string{
        // <schema.Schema Value>
        
    },
   
    "service_appliance_set": []string{
        // <schema.Schema Value>
        
    },
   
    "virtual_machine_interface": []string{
        // <schema.Schema Value>
        
    },
   
    "loadbalancer_listener": []string{
        // <schema.Schema Value>
        
    },
   
    "service_instance": []string{
        // <schema.Schema Value>
        
    },
   
}

// LoadbalancerPoolBackRefFields is db back reference fields for LoadbalancerPool
var LoadbalancerPoolBackRefFields = map[string][]string{
   
   
   "loadbalancer_member": []string{
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
        
   },
   
}

// LoadbalancerPoolParentTypes is possible parents for LoadbalancerPool
var LoadbalancerPoolParents = []string{
   
   "project",
   
}


const insertLoadbalancerPoolServiceApplianceSetQuery = "insert into `ref_loadbalancer_pool_service_appliance_set` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolVirtualMachineInterfaceQuery = "insert into `ref_loadbalancer_pool_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolLoadbalancerListenerQuery = "insert into `ref_loadbalancer_pool_loadbalancer_listener` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolServiceInstanceQuery = "insert into `ref_loadbalancer_pool_service_instance` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolLoadbalancerHealthmonitorQuery = "insert into `ref_loadbalancer_pool_loadbalancer_healthmonitor` (`from`, `to` ) values (?, ?);"


// CreateLoadbalancerPool inserts LoadbalancerPool to DB
func CreateLoadbalancerPool(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.CreateLoadbalancerPoolRequest) error {
    model := request.LoadbalancerPool
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerPoolQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertLoadbalancerPoolQuery,
    }).Debug("create query")
    _, err = stmt.ExecContext(ctx, string(model.GetUUID()),
    common.MustJSON(model.GetPerms2().GetShare()),
    int(model.GetPerms2().GetOwnerAccess()),
    string(model.GetPerms2().GetOwner()),
    int(model.GetPerms2().GetGlobalAccess()),
    string(model.GetParentUUID()),
    string(model.GetParentType()),
    string(model.GetLoadbalancerPoolProvider()),
    string(model.GetLoadbalancerPoolProperties().GetSubnetID()),
    string(model.GetLoadbalancerPoolProperties().GetStatusDescription()),
    string(model.GetLoadbalancerPoolProperties().GetStatus()),
    string(model.GetLoadbalancerPoolProperties().GetSessionPersistence()),
    string(model.GetLoadbalancerPoolProperties().GetProtocol()),
    string(model.GetLoadbalancerPoolProperties().GetPersistenceCookieName()),
    string(model.GetLoadbalancerPoolProperties().GetLoadbalancerMethod()),
    bool(model.GetLoadbalancerPoolProperties().GetAdminState()),
    common.MustJSON(model.GetLoadbalancerPoolCustomAttributes().GetKeyValuePair()),
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
    
    stmtServiceInstanceRef, err := tx.Prepare(insertLoadbalancerPoolServiceInstanceQuery)
	if err != nil {
        return errors.Wrap(err,"preparing ServiceInstanceRefs create statement failed")
	}
    defer stmtServiceInstanceRef.Close()
    for _, ref := range model.ServiceInstanceRefs {
       
        _, err = stmtServiceInstanceRef.ExecContext(ctx, model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"ServiceInstanceRefs create failed")
        }
    }
    
    stmtLoadbalancerHealthmonitorRef, err := tx.Prepare(insertLoadbalancerPoolLoadbalancerHealthmonitorQuery)
	if err != nil {
        return errors.Wrap(err,"preparing LoadbalancerHealthmonitorRefs create statement failed")
	}
    defer stmtLoadbalancerHealthmonitorRef.Close()
    for _, ref := range model.LoadbalancerHealthmonitorRefs {
       
        _, err = stmtLoadbalancerHealthmonitorRef.ExecContext(ctx, model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"LoadbalancerHealthmonitorRefs create failed")
        }
    }
    
    stmtServiceApplianceSetRef, err := tx.Prepare(insertLoadbalancerPoolServiceApplianceSetQuery)
	if err != nil {
        return errors.Wrap(err,"preparing ServiceApplianceSetRefs create statement failed")
	}
    defer stmtServiceApplianceSetRef.Close()
    for _, ref := range model.ServiceApplianceSetRefs {
       
        _, err = stmtServiceApplianceSetRef.ExecContext(ctx, model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"ServiceApplianceSetRefs create failed")
        }
    }
    
    stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLoadbalancerPoolVirtualMachineInterfaceQuery)
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
    
    stmtLoadbalancerListenerRef, err := tx.Prepare(insertLoadbalancerPoolLoadbalancerListenerQuery)
	if err != nil {
        return errors.Wrap(err,"preparing LoadbalancerListenerRefs create statement failed")
	}
    defer stmtLoadbalancerListenerRef.Close()
    for _, ref := range model.LoadbalancerListenerRefs {
       
        _, err = stmtLoadbalancerListenerRef.ExecContext(ctx, model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"LoadbalancerListenerRefs create failed")
        }
    }
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "loadbalancer_pool",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "loadbalancer_pool", model.UUID, model.GetPerms2().GetShare())
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanLoadbalancerPool(values map[string]interface{} ) (*models.LoadbalancerPool, error) {
    m := models.MakeLoadbalancerPool()
    
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
    
    if value, ok := values["loadbalancer_pool_provider"]; ok {
        
            
               m.LoadbalancerPoolProvider = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["subnet_id"]; ok {
        
            
               m.LoadbalancerPoolProperties.SubnetID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["status_description"]; ok {
        
            
               m.LoadbalancerPoolProperties.StatusDescription = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["status"]; ok {
        
            
               m.LoadbalancerPoolProperties.Status = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["session_persistence"]; ok {
        
            
               m.LoadbalancerPoolProperties.SessionPersistence = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["protocol"]; ok {
        
            
               m.LoadbalancerPoolProperties.Protocol = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["persistence_cookie_name"]; ok {
        
            
               m.LoadbalancerPoolProperties.PersistenceCookieName = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["loadbalancer_method"]; ok {
        
            
               m.LoadbalancerPoolProperties.LoadbalancerMethod = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["admin_state"]; ok {
        
            
               m.LoadbalancerPoolProperties.AdminState = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.LoadbalancerPoolCustomAttributes.KeyValuePair)
        
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
    
    if value, ok := values["annotations_key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    if value, ok := values["ref_service_appliance_set"]; ok {
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
            referenceModel := &models.LoadbalancerPoolServiceApplianceSetRef{}
            referenceModel.UUID = uuid
            m.ServiceApplianceSetRefs = append(m.ServiceApplianceSetRefs, referenceModel)
            
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
            referenceModel := &models.LoadbalancerPoolVirtualMachineInterfaceRef{}
            referenceModel.UUID = uuid
            m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)
            
        }
    }
    
    if value, ok := values["ref_loadbalancer_listener"]; ok {
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
            referenceModel := &models.LoadbalancerPoolLoadbalancerListenerRef{}
            referenceModel.UUID = uuid
            m.LoadbalancerListenerRefs = append(m.LoadbalancerListenerRefs, referenceModel)
            
        }
    }
    
    if value, ok := values["ref_service_instance"]; ok {
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
            referenceModel := &models.LoadbalancerPoolServiceInstanceRef{}
            referenceModel.UUID = uuid
            m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)
            
        }
    }
    
    if value, ok := values["ref_loadbalancer_healthmonitor"]; ok {
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
            referenceModel := &models.LoadbalancerPoolLoadbalancerHealthmonitorRef{}
            referenceModel.UUID = uuid
            m.LoadbalancerHealthmonitorRefs = append(m.LoadbalancerHealthmonitorRefs, referenceModel)
            
        }
    }
    
    
    
    if value, ok := values["backref_loadbalancer_member"]; ok {
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
            childModel := models.MakeLoadbalancerMember()
            m.LoadbalancerMembers = append(m.LoadbalancerMembers, childModel)

            
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
            
                if propertyValue, ok := childResourceMap["weight"]; ok && propertyValue != nil {
                
                    
                        childModel.LoadbalancerMemberProperties.Weight = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["status_description"]; ok && propertyValue != nil {
                
                    
                        childModel.LoadbalancerMemberProperties.StatusDescription = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["status"]; ok && propertyValue != nil {
                
                    
                        childModel.LoadbalancerMemberProperties.Status = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["protocol_port"]; ok && propertyValue != nil {
                
                    
                        childModel.LoadbalancerMemberProperties.ProtocolPort = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["admin_state"]; ok && propertyValue != nil {
                
                    
                        childModel.LoadbalancerMemberProperties.AdminState = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["address"]; ok && propertyValue != nil {
                
                    
                        childModel.LoadbalancerMemberProperties.Address = schema.InterfaceToString(propertyValue)
                    
                
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

// ListLoadbalancerPool lists LoadbalancerPool with list spec.
func ListLoadbalancerPool(ctx context.Context, tx *sql.Tx, request *models.ListLoadbalancerPoolRequest) (response *models.ListLoadbalancerPoolResponse, err error) {
    var rows *sql.Rows
    qb := &common.ListQueryBuilder{}
    qb.Auth = common.GetAuthCTX(ctx) 
    spec := request.Spec
    qb.Spec = spec
    qb.Table = "loadbalancer_pool"
    qb.Fields = LoadbalancerPoolFields
    qb.RefFields = LoadbalancerPoolRefFields
    qb.BackRefFields = LoadbalancerPoolBackRefFields
    result := []*models.LoadbalancerPool{}

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
            m, err := scanLoadbalancerPool(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    response = &models.ListLoadbalancerPoolResponse{
       LoadbalancerPools: result,
    }
    return response, nil
}

// UpdateLoadbalancerPool updates a resource
func UpdateLoadbalancerPool(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.UpdateLoadbalancerPoolRequest,
    ) error {
    //TODO
    return nil
}

// DeleteLoadbalancerPool deletes a resource
func DeleteLoadbalancerPool(
    ctx context.Context,
    tx *sql.Tx, 
    request *models.DeleteLoadbalancerPoolRequest) error {
    deleteQuery := deleteLoadbalancerPoolQuery
    selectQuery := "select count(uuid) from loadbalancer_pool where uuid = ?"
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