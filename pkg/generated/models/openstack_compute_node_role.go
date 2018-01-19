package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
	ProvisioningState           string         `json:"provisioning_state,omitempty"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
	DisplayName                 string         `json:"display_name,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	ParentUUID                  string         `json:"parent_uuid,omitempty"`
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
		VrouterBondInterfaceMembers: "",
		VrouterType:                 "",
		IDPerms:                     MakeIdPermsType(),
		ProvisioningProgress:        0,
		ProvisioningState:           "",
		ParentType:                  "",
		FQName:                      []string{},
		Annotations:                 MakeKeyValuePairs(),
		Perms2:                      MakePermType2(),
		ProvisioningLog:             "",
		ProvisioningStartTime:       "",
		DefaultGateway:              "",
		VrouterBondInterface:        "",
		ParentUUID:                  "",
		DisplayName:                 "",
		UUID:                        "",
		ProvisioningProgressStage: "",
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
