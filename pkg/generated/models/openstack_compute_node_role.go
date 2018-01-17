package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	DisplayName                 string         `json:"display_name,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	ParentUUID                  string         `json:"parent_uuid,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState           string         `json:"provisioning_state,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
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
		ParentUUID:                  "",
		FQName:                      []string{},
		ProvisioningStartTime:       "",
		ProvisioningState:           "",
		DefaultGateway:              "",
		VrouterType:                 "",
		IDPerms:                     MakeIdPermsType(),
		Annotations:                 MakeKeyValuePairs(),
		VrouterBondInterface:        "",
		VrouterBondInterfaceMembers: "",
		ProvisioningLog:             "",
		ProvisioningProgressStage:   "",
		DisplayName:                 "",
		UUID:                        "",
		ProvisioningProgress:        0,
		Perms2:                      MakePermType2(),
		ParentType:                  "",
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
