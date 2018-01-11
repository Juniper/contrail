package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	DefaultGateway              string         `json:"default_gateway"`
	VrouterType                 string         `json:"vrouter_type"`
	ParentUUID                  string         `json:"parent_uuid"`
	ProvisioningLog             string         `json:"provisioning_log"`
	ProvisioningProgress        int            `json:"provisioning_progress"`
	DisplayName                 string         `json:"display_name"`
	Perms2                      *PermType2     `json:"perms2"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members"`
	UUID                        string         `json:"uuid"`
	ParentType                  string         `json:"parent_type"`
	FQName                      []string       `json:"fq_name"`
	IDPerms                     *IdPermsType   `json:"id_perms"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage"`
	ProvisioningStartTime       string         `json:"provisioning_start_time"`
	Annotations                 *KeyValuePairs `json:"annotations"`
	ProvisioningState           string         `json:"provisioning_state"`
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
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Perms2:      MakePermType2(),
		VrouterBondInterfaceMembers: "",
		UUID: "",
		ProvisioningStartTime:     "",
		VrouterBondInterface:      "",
		ProvisioningProgressStage: "",
		Annotations:               MakeKeyValuePairs(),
		ProvisioningState:         "",
		ParentUUID:                "",
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		DefaultGateway:            "",
		VrouterType:               "",
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
