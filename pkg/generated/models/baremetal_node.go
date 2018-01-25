package models

// BaremetalNode

// BaremetalNode
//proteus:generate
type BaremetalNode struct {
	UUID          string         `json:"uuid,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
	Name          string         `json:"name,omitempty"`
	IpmiAddress   string         `json:"ipmi_address,omitempty"`
	IpmiUsername  string         `json:"ipmi_username,omitempty"`
	IpmiPassword  string         `json:"ipmi_password,omitempty"`
	CPUCount      int            `json:"cpu_count,omitempty"`
	CPUArch       string         `json:"cpu_arch,omitempty"`
	DiskGB        int            `json:"disk_gb,omitempty"`
	MemoryMB      int            `json:"memory_mb,omitempty"`
	DeployKernel  string         `json:"deploy_kernel,omitempty"`
	DeployRamdisk string         `json:"deploy_ramdisk,omitempty"`
}

// MakeBaremetalNode makes BaremetalNode
func MakeBaremetalNode() *BaremetalNode {
	return &BaremetalNode{
		//TODO(nati): Apply default
		UUID:          "",
		ParentUUID:    "",
		ParentType:    "",
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
		Name:          "",
		IpmiAddress:   "",
		IpmiUsername:  "",
		IpmiPassword:  "",
		CPUCount:      0,
		CPUArch:       "",
		DiskGB:        0,
		MemoryMB:      0,
		DeployKernel:  "",
		DeployRamdisk: "",
	}
}

// MakeBaremetalNodeSlice() makes a slice of BaremetalNode
func MakeBaremetalNodeSlice() []*BaremetalNode {
	return []*BaremetalNode{}
}
