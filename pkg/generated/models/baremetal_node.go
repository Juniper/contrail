package models

// BaremetalNode

import "encoding/json"

// BaremetalNode
type BaremetalNode struct {
	IpmiAddress   string         `json:"ipmi_address,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	IpmiUsername  string         `json:"ipmi_username,omitempty"`
	DiskGB        int            `json:"disk_gb,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	CPUArch       string         `json:"cpu_arch,omitempty"`
	MemoryMB      int            `json:"memory_mb,omitempty"`
	DeployKernel  string         `json:"deploy_kernel,omitempty"`
	Name          string         `json:"name,omitempty"`
	IpmiPassword  string         `json:"ipmi_password,omitempty"`
	CPUCount      int            `json:"cpu_count,omitempty"`
	DeployRamdisk string         `json:"deploy_ramdisk,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
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
		CPUArch:       "",
		MemoryMB:      0,
		DeployKernel:  "",
		ParentUUID:    "",
		Perms2:        MakePermType2(),
		Name:          "",
		IpmiPassword:  "",
		CPUCount:      0,
		DeployRamdisk: "",
		IpmiAddress:   "",
		ParentType:    "",
		UUID:          "",
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
		IpmiUsername:  "",
		DiskGB:        0,
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
	}
}

// MakeBaremetalNodeSlice() makes a slice of BaremetalNode
func MakeBaremetalNodeSlice() []*BaremetalNode {
	return []*BaremetalNode{}
}
