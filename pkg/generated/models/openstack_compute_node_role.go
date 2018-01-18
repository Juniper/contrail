package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	DisplayName                 string         `json:"display_name,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	ProvisioningState           string         `json:"provisioning_state,omitempty"`
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	ParentUUID                  string         `json:"parent_uuid,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
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
		ProvisioningStartTime:       "",
		Annotations:                 MakeKeyValuePairs(),
		Perms2:                      MakePermType2(),
		UUID:                        "",
		DisplayName:                 "",
		ProvisioningLog:             "",
		VrouterType:                 "",
		VrouterBondInterfaceMembers: "",
		ParentUUID:                  "",
		ParentType:                  "",
		FQName:                      []string{},
		ProvisioningState:           "",
		DefaultGateway:              "",
		IDPerms:                     MakeIdPermsType(),
		ProvisioningProgress:        0,
		ProvisioningProgressStage:   "",
		VrouterBondInterface:        "",
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
