package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState           string         `json:"provisioning_state,omitempty"`
	ParentUUID                  string         `json:"parent_uuid,omitempty"`
	DisplayName                 string         `json:"display_name,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
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
		DefaultGateway:              "",
		VrouterBondInterface:        "",
		VrouterBondInterfaceMembers: "",
		ParentType:                  "",
		ProvisioningLog:             "",
		IDPerms:                     MakeIdPermsType(),
		ProvisioningProgress:        0,
		ProvisioningState:           "",
		ParentUUID:                  "",
		DisplayName:                 "",
		Perms2:                      MakePermType2(),
		UUID:                        "",
		ProvisioningProgressStage: "",
		VrouterType:               "",
		FQName:                    []string{},
		Annotations:               MakeKeyValuePairs(),
		ProvisioningStartTime:     "",
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
