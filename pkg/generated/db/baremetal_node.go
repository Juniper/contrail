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

const insertBaremetalNodeQuery = "insert into `baremetal_node` (`uuid`,`updated_at`,`target_provision_state`,`target_power_state`,`provision_state`,`power_state`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`name`,`maintenance_reason`,`maintenance`,`last_error`,`instance_uuid`,`vcpus`,`swap_mb`,`root_gb`,`nova_host_id`,`memory_mb`,`local_gb`,`image_source`,`display_name`,`capabilities`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`ipmi_username`,`ipmi_password`,`ipmi_address`,`deploy_ramdisk`,`deploy_kernel`,`_display_name`,`created_at`,`console_enabled`,`bm_properties_memory_mb`,`disk_gb`,`cpu_count`,`cpu_arch`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteBaremetalNodeQuery = "delete from `baremetal_node` where uuid = ?"

// BaremetalNodeFields is db columns for BaremetalNode
var BaremetalNodeFields = []string{
   "uuid",
   "updated_at",
   "target_provision_state",
   "target_power_state",
   "provision_state",
   "power_state",
   "share",
   "owner_access",
   "owner",
   "global_access",
   "parent_uuid",
   "parent_type",
   "name",
   "maintenance_reason",
   "maintenance",
   "last_error",
   "instance_uuid",
   "vcpus",
   "swap_mb",
   "root_gb",
   "nova_host_id",
   "memory_mb",
   "local_gb",
   "image_source",
   "display_name",
   "capabilities",
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
   "ipmi_username",
   "ipmi_password",
   "ipmi_address",
   "deploy_ramdisk",
   "deploy_kernel",
   "_display_name",
   "created_at",
   "console_enabled",
   "bm_properties_memory_mb",
   "disk_gb",
   "cpu_count",
   "cpu_arch",
   "key_value_pair",
   
}

// BaremetalNodeRefFields is db reference fields for BaremetalNode
var BaremetalNodeRefFields = map[string][]string{
   
}

// BaremetalNodeBackRefFields is db back reference fields for BaremetalNode
var BaremetalNodeBackRefFields = map[string][]string{
   
}

// BaremetalNodeParentTypes is possible parents for BaremetalNode
var BaremetalNodeParents = []string{
   
}



// CreateBaremetalNode inserts BaremetalNode to DB
func CreateBaremetalNode(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.CreateBaremetalNodeRequest) error {
    model := request.BaremetalNode
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBaremetalNodeQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertBaremetalNodeQuery,
    }).Debug("create query")
    _, err = stmt.ExecContext(ctx, string(model.GetUUID()),
    string(model.GetUpdatedAt()),
    string(model.GetTargetProvisionState()),
    string(model.GetTargetPowerState()),
    string(model.GetProvisionState()),
    string(model.GetPowerState()),
    common.MustJSON(model.GetPerms2().GetShare()),
    int(model.GetPerms2().GetOwnerAccess()),
    string(model.GetPerms2().GetOwner()),
    int(model.GetPerms2().GetGlobalAccess()),
    string(model.GetParentUUID()),
    string(model.GetParentType()),
    string(model.GetName()),
    string(model.GetMaintenanceReason()),
    bool(model.GetMaintenance()),
    string(model.GetLastError()),
    string(model.GetInstanceUUID()),
    string(model.GetInstanceInfo().GetVcpus()),
    string(model.GetInstanceInfo().GetSwapMB()),
    string(model.GetInstanceInfo().GetRootGB()),
    string(model.GetInstanceInfo().GetNovaHostID()),
    string(model.GetInstanceInfo().GetMemoryMB()),
    string(model.GetInstanceInfo().GetLocalGB()),
    string(model.GetInstanceInfo().GetImageSource()),
    string(model.GetInstanceInfo().GetDisplayName()),
    string(model.GetInstanceInfo().GetCapabilities()),
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
    string(model.GetDriverInfo().GetIpmiUsername()),
    string(model.GetDriverInfo().GetIpmiPassword()),
    string(model.GetDriverInfo().GetIpmiAddress()),
    string(model.GetDriverInfo().GetDeployRamdisk()),
    string(model.GetDriverInfo().GetDeployKernel()),
    string(model.GetDisplayName()),
    string(model.GetCreatedAt()),
    bool(model.GetConsoleEnabled()),
    int(model.GetBMProperties().GetMemoryMB()),
    int(model.GetBMProperties().GetDiskGB()),
    int(model.GetBMProperties().GetCPUCount()),
    string(model.GetBMProperties().GetCPUArch()),
    common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "baremetal_node",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "baremetal_node", model.UUID, model.GetPerms2().GetShare())
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanBaremetalNode(values map[string]interface{} ) (*models.BaremetalNode, error) {
    m := models.MakeBaremetalNode()
    
    if value, ok := values["uuid"]; ok {
        
            
               m.UUID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["updated_at"]; ok {
        
            
               m.UpdatedAt = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["target_provision_state"]; ok {
        
            
               m.TargetProvisionState = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["target_power_state"]; ok {
        
            
               m.TargetPowerState = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["provision_state"]; ok {
        
            
               m.ProvisionState = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["power_state"]; ok {
        
            
               m.PowerState = schema.InterfaceToString(value)
            
        
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
    
    if value, ok := values["name"]; ok {
        
            
               m.Name = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["maintenance_reason"]; ok {
        
            
               m.MaintenanceReason = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["maintenance"]; ok {
        
            
               m.Maintenance = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["last_error"]; ok {
        
            
               m.LastError = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["instance_uuid"]; ok {
        
            
               m.InstanceUUID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["vcpus"]; ok {
        
            
               m.InstanceInfo.Vcpus = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["swap_mb"]; ok {
        
            
               m.InstanceInfo.SwapMB = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["root_gb"]; ok {
        
            
               m.InstanceInfo.RootGB = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["nova_host_id"]; ok {
        
            
               m.InstanceInfo.NovaHostID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["memory_mb"]; ok {
        
            
               m.InstanceInfo.MemoryMB = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["local_gb"]; ok {
        
            
               m.InstanceInfo.LocalGB = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["image_source"]; ok {
        
            
               m.InstanceInfo.ImageSource = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["display_name"]; ok {
        
            
               m.InstanceInfo.DisplayName = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["capabilities"]; ok {
        
            
               m.InstanceInfo.Capabilities = schema.InterfaceToString(value)
            
        
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
    
    if value, ok := values["ipmi_username"]; ok {
        
            
               m.DriverInfo.IpmiUsername = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["ipmi_password"]; ok {
        
            
               m.DriverInfo.IpmiPassword = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["ipmi_address"]; ok {
        
            
               m.DriverInfo.IpmiAddress = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["deploy_ramdisk"]; ok {
        
            
               m.DriverInfo.DeployRamdisk = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["deploy_kernel"]; ok {
        
            
               m.DriverInfo.DeployKernel = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["_display_name"]; ok {
        
            
               m.DisplayName = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["created_at"]; ok {
        
            
               m.CreatedAt = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["console_enabled"]; ok {
        
            
               m.ConsoleEnabled = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["bm_properties_memory_mb"]; ok {
        
            
               m.BMProperties.MemoryMB = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["disk_gb"]; ok {
        
            
               m.BMProperties.DiskGB = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["cpu_count"]; ok {
        
            
               m.BMProperties.CPUCount = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["cpu_arch"]; ok {
        
            
               m.BMProperties.CPUArch = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    
    return m, nil
}

// ListBaremetalNode lists BaremetalNode with list spec.
func ListBaremetalNode(ctx context.Context, tx *sql.Tx, request *models.ListBaremetalNodeRequest) (response *models.ListBaremetalNodeResponse, err error) {
    var rows *sql.Rows
    qb := &common.ListQueryBuilder{}
    qb.Auth = common.GetAuthCTX(ctx) 
    spec := request.Spec
    qb.Spec = spec
    qb.Table = "baremetal_node"
    qb.Fields = BaremetalNodeFields
    qb.RefFields = BaremetalNodeRefFields
    qb.BackRefFields = BaremetalNodeBackRefFields
    result := []*models.BaremetalNode{}

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
            m, err := scanBaremetalNode(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    response = &models.ListBaremetalNodeResponse{
       BaremetalNodes: result,
    }
    return response, nil
}

// UpdateBaremetalNode updates a resource
func UpdateBaremetalNode(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.UpdateBaremetalNodeRequest,
    ) error {
    //TODO
    return nil
}

// DeleteBaremetalNode deletes a resource
func DeleteBaremetalNode(
    ctx context.Context,
    tx *sql.Tx, 
    request *models.DeleteBaremetalNodeRequest) error {
    deleteQuery := deleteBaremetalNodeQuery
    selectQuery := "select count(uuid) from baremetal_node where uuid = ?"
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