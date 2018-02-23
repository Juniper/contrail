package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeContrailStorageNode makes ContrailStorageNode
func MakeContrailStorageNode() *ContrailStorageNode{
    return &ContrailStorageNode{
    //TODO(nati): Apply default
    ProvisioningLog: "",
        ProvisioningProgress: 0,
        ProvisioningProgressStage: "",
        ProvisioningStartTime: "",
        ProvisioningState: "",
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        JournalDrives: "",
        OsdDrives: "",
        StorageAccessBondInterfaceMembers: "",
        StorageBackendBondInterfaceMembers: "",
        
    }
}

// MakeContrailStorageNode makes ContrailStorageNode
func InterfaceToContrailStorageNode(i interface{}) *ContrailStorageNode{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ContrailStorageNode{
    //TODO(nati): Apply default
    ProvisioningLog: schema.InterfaceToString(m["provisioning_log"]),
        ProvisioningProgress: schema.InterfaceToInt64(m["provisioning_progress"]),
        ProvisioningProgressStage: schema.InterfaceToString(m["provisioning_progress_stage"]),
        ProvisioningStartTime: schema.InterfaceToString(m["provisioning_start_time"]),
        ProvisioningState: schema.InterfaceToString(m["provisioning_state"]),
        UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        JournalDrives: schema.InterfaceToString(m["journal_drives"]),
        OsdDrives: schema.InterfaceToString(m["osd_drives"]),
        StorageAccessBondInterfaceMembers: schema.InterfaceToString(m["storage_access_bond_interface_members"]),
        StorageBackendBondInterfaceMembers: schema.InterfaceToString(m["storage_backend_bond_interface_members"]),
        
    }
}

// MakeContrailStorageNodeSlice() makes a slice of ContrailStorageNode
func MakeContrailStorageNodeSlice() []*ContrailStorageNode {
    return []*ContrailStorageNode{}
}

// InterfaceToContrailStorageNodeSlice() makes a slice of ContrailStorageNode
func InterfaceToContrailStorageNodeSlice(i interface{}) []*ContrailStorageNode {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ContrailStorageNode{}
    for _, item := range list {
        result = append(result, InterfaceToContrailStorageNode(item) )
    }
    return result
}



