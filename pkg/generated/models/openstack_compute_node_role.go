package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	ParentUUID                  string         `json:"parent_uuid,omitempty"`
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                 string         `json:"display_name,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState           string         `json:"provisioning_state,omitempty"`
}

// String returns json representation of the object
func (model *OpenstackComputeNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeOpenstackComputeNodeRole makes OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRole() *OpenstackComputeNodeRole {
	return &OpenstackComputeNodeRole{
		//TODO(nati): Apply default
		VrouterType:          "",
		ParentUUID:           "",
		Annotations:          MakeKeyValuePairs(),
		ProvisioningLog:      "",
		VrouterBondInterface: "",
		Perms2:               MakePermType2(),
		UUID:                 "",
		ProvisioningStartTime:       "",
		FQName:                      []string{},
		IDPerms:                     MakeIdPermsType(),
		DisplayName:                 "",
		ProvisioningProgressStage:   "",
		ProvisioningState:           "",
		DefaultGateway:              "",
		VrouterBondInterfaceMembers: "",
		ParentType:                  "",
		ProvisioningProgress:        0,
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
