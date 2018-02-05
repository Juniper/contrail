package models

// BaremetalNode

import "encoding/json"

// BaremetalNode
type BaremetalNode struct {
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	IpmiUsername  string         `json:"ipmi_username,omitempty"`
	DiskGB        int            `json:"disk_gb,omitempty"`
	MemoryMB      int            `json:"memory_mb,omitempty"`
	DeployRamdisk string         `json:"deploy_ramdisk,omitempty"`
	Name          string         `json:"name,omitempty"`
	CPUArch       string         `json:"cpu_arch,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
	IpmiAddress   string         `json:"ipmi_address,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	IpmiPassword  string         `json:"ipmi_password,omitempty"`
	CPUCount      int            `json:"cpu_count,omitempty"`
	DeployKernel  string         `json:"deploy_kernel,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
}

// String returns json representation of the object
func (model *BaremetalNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBaremetalNode makes BaremetalNode
func MakeBaremetalNode() *BaremetalNode {
	return &BaremetalNode{
		//TODO(nati): Apply default
		IpmiPassword:  "",
		CPUCount:      0,
		DeployKernel:  "",
		FQName:        []string{},
		DisplayName:   "",
		ParentUUID:    "",
		IpmiUsername:  "",
		DiskGB:        0,
		MemoryMB:      0,
		DeployRamdisk: "",
		Annotations:   MakeKeyValuePairs(),
		UUID:          "",
		Name:          "",
		CPUArch:       "",
		IDPerms:       MakeIdPermsType(),
		Perms2:        MakePermType2(),
		IpmiAddress:   "",
		ParentType:    "",
	}
}

// MakeBaremetalNodeSlice() makes a slice of BaremetalNode
func MakeBaremetalNodeSlice() []*BaremetalNode {
	return []*BaremetalNode{}
}
