package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	ParentUUID                  string         `json:"parent_uuid,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningState           string         `json:"provisioning_state,omitempty"`
	DisplayName                 string         `json:"display_name,omitempty"`
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
		DisplayName:                 "",
		Annotations:                 MakeKeyValuePairs(),
		Perms2:                      MakePermType2(),
		IDPerms:                     MakeIdPermsType(),
		ProvisioningState:           "",
		VrouterBondInterface:        "",
		VrouterType:                 "",
		ParentUUID:                  "",
		UUID:                        "",
		ParentType:                  "",
		ProvisioningProgressStage:   "",
		ProvisioningStartTime:       "",
		DefaultGateway:              "",
		VrouterBondInterfaceMembers: "",
		FQName:               []string{},
		ProvisioningLog:      "",
		ProvisioningProgress: 0,
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
