package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	Annotations                 *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage,omitempty"`
	VrouterType                 string         `json:"vrouter_type,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	FQName                      []string       `json:"fq_name,omitempty"`
	ProvisioningState           string         `json:"provisioning_state,omitempty"`
	DefaultGateway              string         `json:"default_gateway,omitempty"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface,omitempty"`
	ProvisioningStartTime       string         `json:"provisioning_start_time,omitempty"`
	ParentType                  string         `json:"parent_type,omitempty"`
	IDPerms                     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                 string         `json:"display_name,omitempty"`
	ProvisioningLog             string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress        int            `json:"provisioning_progress,omitempty"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members,omitempty"`
	Perms2                      *PermType2     `json:"perms2,omitempty"`
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
		VrouterType: "",
		UUID:        "",
		FQName:      []string{},
		ProvisioningProgressStage:   "",
		DefaultGateway:              "",
		VrouterBondInterface:        "",
		ProvisioningStartTime:       "",
		ProvisioningState:           "",
		DisplayName:                 "",
		ProvisioningLog:             "",
		ProvisioningProgress:        0,
		VrouterBondInterfaceMembers: "",
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
