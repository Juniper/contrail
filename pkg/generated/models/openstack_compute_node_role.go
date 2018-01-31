package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	ParentUUID                  string         `json:"parent_uuid,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
	DisplayName                 string         `json:"display_name,omitempty"`
	ProvisioningState           string         `json:"provisioning_state,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
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
		UUID: "",
		ProvisioningStartTime:       "",
		VrouterBondInterfaceMembers: "",
		ProvisioningProgress:        0,
		ParentUUID:                  "",
		ParentType:                  "",
		FQName:                      []string{},
		IDPerms:                     MakeIdPermsType(),
		DefaultGateway:              "",
		VrouterType:                 "",
		Annotations:                 MakeKeyValuePairs(),
		Perms2:                      MakePermType2(),
		DisplayName:                 "",
		ProvisioningState:           "",
		ProvisioningLog:             "",
		ProvisioningProgressStage:   "",
		VrouterBondInterface:        "",
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
