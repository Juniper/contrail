package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	FQName                                []string       `json:"fq_name,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
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
		CapacityDrives:                        "",
		IDPerms:                               MakeIdPermsType(),
		ProvisioningProgress:                  0,
		StorageManagementBondInterfaceMembers: "",
		ParentUUID:                            "",
		ProvisioningStartTime:                 "",
		ProvisioningState:                     "",
		FQName:                                []string{},
		Perms2:                                MakePermType2(),
		InternalapiBondInterfaceMembers: "",
		PerformanceDrives:               "",
		UUID:                            "",
		ParentType:                      "",
		DisplayName:                     "",
		Annotations:                     MakeKeyValuePairs(),
		ProvisioningLog:                 "",
		ProvisioningProgressStage:       "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
