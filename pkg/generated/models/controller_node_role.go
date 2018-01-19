package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	FQName                                []string       `json:"fq_name,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
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
		ProvisioningProgress:            0,
		InternalapiBondInterfaceMembers: "",
		FQName:                                []string{},
		DisplayName:                           "",
		Annotations:                           MakeKeyValuePairs(),
		ProvisioningState:                     "",
		CapacityDrives:                        "",
		StorageManagementBondInterfaceMembers: "",
		UUID:                      "",
		ParentUUID:                "",
		ProvisioningLog:           "",
		Perms2:                    MakePermType2(),
		IDPerms:                   MakeIdPermsType(),
		ProvisioningStartTime:     "",
		PerformanceDrives:         "",
		ParentType:                "",
		ProvisioningProgressStage: "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
