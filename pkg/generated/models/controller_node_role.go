package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
//proteus:generate
type ControllerNodeRole struct {
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	FQName                                []string       `json:"fq_name,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
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
		ProvisioningLog:                       "",
		ProvisioningProgress:                  0,
		ProvisioningProgressStage:             "",
		ProvisioningStartTime:                 "",
		ProvisioningState:                     "",
		UUID:                                  "",
		ParentUUID:                            "",
		ParentType:                            "",
		FQName:                                []string{},
		IDPerms:                               MakeIdPermsType(),
		DisplayName:                           "",
		Annotations:                           MakeKeyValuePairs(),
		Perms2:                                MakePermType2(),
		CapacityDrives:                        "",
		InternalapiBondInterfaceMembers:       "",
		PerformanceDrives:                     "",
		StorageManagementBondInterfaceMembers: "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
