package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState           string         `json:"provisioning_state,omitempty"`
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID                  string         `json:"parent_uuid,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
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
		Annotations:                 MakeKeyValuePairs(),
		ParentUUID:                  "",
		ProvisioningStartTime:       "",
		VrouterBondInterface:        "",
		IDPerms:                     MakeIdPermsType(),
		DisplayName:                 "",
		VrouterBondInterfaceMembers: "",
		Perms2:                    MakePermType2(),
		FQName:                    []string{},
		UUID:                      "",
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		ProvisioningState:         "",
		DefaultGateway:            "",
		VrouterType:               "",
		ParentType:                "",
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
