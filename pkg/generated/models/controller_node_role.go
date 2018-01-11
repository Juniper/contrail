package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	ParentUUID                            string         `json:"parent_uuid"`
	ParentType                            string         `json:"parent_type"`
	FQName                                []string       `json:"fq_name"`
	IDPerms                               *IdPermsType   `json:"id_perms"`
	CapacityDrives                        string         `json:"capacity_drives"`
	Perms2                                *PermType2     `json:"perms2"`
	ProvisioningLog                       string         `json:"provisioning_log"`
	Annotations                           *KeyValuePairs `json:"annotations"`
	DisplayName                           string         `json:"display_name"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members"`
	PerformanceDrives                     string         `json:"performance_drives"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members"`
	UUID                                  string         `json:"uuid"`
	ProvisioningState                     string         `json:"provisioning_state"`
	ProvisioningProgress                  int            `json:"provisioning_progress"`
}

// String returns json representation of the object
func (model *ControllerNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeControllerNodeRole makes ControllerNodeRole
func MakeControllerNodeRole() *ControllerNodeRole {
	return &ControllerNodeRole{
		//TODO(nati): Apply default
		ParentUUID:                            "",
		ParentType:                            "",
		FQName:                                []string{},
		IDPerms:                               MakeIdPermsType(),
		CapacityDrives:                        "",
		Perms2:                                MakePermType2(),
		ProvisioningLog:                       "",
		Annotations:                           MakeKeyValuePairs(),
		DisplayName:                           "",
		ProvisioningStartTime:                 "",
		ProvisioningState:                     "",
		ProvisioningProgress:                  0,
		ProvisioningProgressStage:             "",
		InternalapiBondInterfaceMembers:       "",
		PerformanceDrives:                     "",
		StorageManagementBondInterfaceMembers: "",
		UUID: "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
