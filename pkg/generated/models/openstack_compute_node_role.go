package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                 string         `json:"display_name,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	ParentUUID                  string         `json:"parent_uuid,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
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
		VrouterType:                 "",
		UUID:                        "",
		FQName:                      []string{},
		Annotations:                 MakeKeyValuePairs(),
		ProvisioningProgressStage:   "",
		VrouterBondInterface:        "",
		ParentUUID:                  "",
		Perms2:                      MakePermType2(),
		ProvisioningState:           "",
		ProvisioningProgress:        0,
		DefaultGateway:              "",
		ParentType:                  "",
		VrouterBondInterfaceMembers: "",
		IDPerms:                     MakeIdPermsType(),
		DisplayName:                 "",
		ProvisioningStartTime:       "",
		ProvisioningLog:             "",
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
