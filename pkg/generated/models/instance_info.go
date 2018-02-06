package models

// InstanceInfo

// InstanceInfo
//proteus:generate
type InstanceInfo struct {
	DisplayName  string `json:"display_name,omitempty"`
	ImageSource  string `json:"image_source,omitempty"`
	LocalGB      string `json:"local_gb,omitempty"`
	MemoryMB     string `json:"memory_mb,omitempty"`
	NovaHostID   string `json:"nova_host_id,omitempty"`
	RootGB       string `json:"root_gb,omitempty"`
	SwapMB       string `json:"swap_mb,omitempty"`
	Vcpus        string `json:"vcpus,omitempty"`
	Capabilities string `json:"capabilities,omitempty"`
}

// MakeInstanceInfo makes InstanceInfo
func MakeInstanceInfo() *InstanceInfo {
	return &InstanceInfo{
		//TODO(nati): Apply default
		DisplayName:  "",
		ImageSource:  "",
		LocalGB:      "",
		MemoryMB:     "",
		NovaHostID:   "",
		RootGB:       "",
		SwapMB:       "",
		Vcpus:        "",
		Capabilities: "",
	}
}

// MakeInstanceInfoSlice() makes a slice of InstanceInfo
func MakeInstanceInfoSlice() []*InstanceInfo {
	return []*InstanceInfo{}
}
