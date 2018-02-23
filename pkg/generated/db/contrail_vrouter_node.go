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

const insertContrailVrouterNodeQuery = "insert into `contrail_vrouter_node` (`vrouter_type`,`vrouter_bond_interface_members`,`vrouter_bond_interface`,`uuid`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`default_gateway`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteContrailVrouterNodeQuery = "delete from `contrail_vrouter_node` where uuid = ?"

// ContrailVrouterNodeFields is db columns for ContrailVrouterNode
var ContrailVrouterNodeFields = []string{
   "vrouter_type",
   "vrouter_bond_interface_members",
   "vrouter_bond_interface",
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
   "default_gateway",
   "key_value_pair",
   
}

// ContrailVrouterNodeRefFields is db reference fields for ContrailVrouterNode
var ContrailVrouterNodeRefFields = map[string][]string{
   
    "node": []string{
        // <schema.Schema Value>
        
    },
   
}

// ContrailVrouterNodeBackRefFields is db back reference fields for ContrailVrouterNode
var ContrailVrouterNodeBackRefFields = map[string][]string{
   
}

// ContrailVrouterNodeParentTypes is possible parents for ContrailVrouterNode
var ContrailVrouterNodeParents = []string{
   
}


const insertContrailVrouterNodeNodeQuery = "insert into `ref_contrail_vrouter_node_node` (`from`, `to` ) values (?, ?);"


// CreateContrailVrouterNode inserts ContrailVrouterNode to DB
func CreateContrailVrouterNode(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.CreateContrailVrouterNodeRequest) error {
    model := request.ContrailVrouterNode
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailVrouterNodeQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertContrailVrouterNodeQuery,
    }).Debug("create query")
    _, err = stmt.ExecContext(ctx, string(model.GetVrouterType()),
    string(model.GetVrouterBondInterfaceMembers()),
    string(model.GetVrouterBondInterface()),
    string(model.GetUUID()),
    string(model.GetProvisioningState()),
    string(model.GetProvisioningStartTime()),
    string(model.GetProvisioningProgressStage()),
    int(model.GetProvisioningProgress()),
    string(model.GetProvisioningLog()),
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
    string(model.GetDefaultGateway()),
    common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    stmtNodeRef, err := tx.Prepare(insertContrailVrouterNodeNodeQuery)
	if err != nil {
        return errors.Wrap(err,"preparing NodeRefs create statement failed")
	}
    defer stmtNodeRef.Close()
    for _, ref := range model.NodeRefs {
       
        _, err = stmtNodeRef.ExecContext(ctx, model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"NodeRefs create failed")
        }
    }
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "contrail_vrouter_node",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "contrail_vrouter_node", model.UUID, model.GetPerms2().GetShare())
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanContrailVrouterNode(values map[string]interface{} ) (*models.ContrailVrouterNode, error) {
    m := models.MakeContrailVrouterNode()
    
    if value, ok := values["vrouter_type"]; ok {
        
            
               m.VrouterType = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["vrouter_bond_interface_members"]; ok {
        
            
               m.VrouterBondInterfaceMembers = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["vrouter_bond_interface"]; ok {
        
            
               m.VrouterBondInterface = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["uuid"]; ok {
        
            
               m.UUID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["provisioning_state"]; ok {
        
            
               m.ProvisioningState = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["provisioning_start_time"]; ok {
        
            
               m.ProvisioningStartTime = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["provisioning_progress_stage"]; ok {
        
            
               m.ProvisioningProgressStage = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["provisioning_progress"]; ok {
        
            
               m.ProvisioningProgress = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["provisioning_log"]; ok {
        
            
               m.ProvisioningLog = schema.InterfaceToString(value)
            
        
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
    
    if value, ok := values["default_gateway"]; ok {
        
            
               m.DefaultGateway = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    if value, ok := values["ref_node"]; ok {
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
            referenceModel := &models.ContrailVrouterNodeNodeRef{}
            referenceModel.UUID = uuid
            m.NodeRefs = append(m.NodeRefs, referenceModel)
            
        }
    }
    
    
    return m, nil
}

// ListContrailVrouterNode lists ContrailVrouterNode with list spec.
func ListContrailVrouterNode(ctx context.Context, tx *sql.Tx, request *models.ListContrailVrouterNodeRequest) (response *models.ListContrailVrouterNodeResponse, err error) {
    var rows *sql.Rows
    qb := &common.ListQueryBuilder{}
    qb.Auth = common.GetAuthCTX(ctx) 
    spec := request.Spec
    qb.Spec = spec
    qb.Table = "contrail_vrouter_node"
    qb.Fields = ContrailVrouterNodeFields
    qb.RefFields = ContrailVrouterNodeRefFields
    qb.BackRefFields = ContrailVrouterNodeBackRefFields
    result := []*models.ContrailVrouterNode{}

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
            m, err := scanContrailVrouterNode(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    response = &models.ListContrailVrouterNodeResponse{
       ContrailVrouterNodes: result,
    }
    return response, nil
}

// UpdateContrailVrouterNode updates a resource
func UpdateContrailVrouterNode(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.UpdateContrailVrouterNodeRequest,
    ) error {
    //TODO
    return nil
}

// DeleteContrailVrouterNode deletes a resource
func DeleteContrailVrouterNode(
    ctx context.Context,
    tx *sql.Tx, 
    request *models.DeleteContrailVrouterNodeRequest) error {
    deleteQuery := deleteContrailVrouterNodeQuery
    selectQuery := "select count(uuid) from contrail_vrouter_node where uuid = ?"
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