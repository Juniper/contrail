package models

// BaremetalProperties

// BaremetalProperties
//proteus:generate
type BaremetalProperties struct {
	CPUCount int    `json:"cpu_count,omitempty"`
	CPUArch  string `json:"cpu_arch,omitempty"`
	DiskGB   int    `json:"disk_gb,omitempty"`
	MemoryMB int    `json:"memory_mb,omitempty"`
}

// MakeBaremetalProperties makes BaremetalProperties
func MakeBaremetalProperties() *BaremetalProperties {
	return &BaremetalProperties{
		//TODO(nati): Apply default
		CPUCount: 0,
		CPUArch:  "",
		DiskGB:   0,
		MemoryMB: 0,
	}
}

// MakeBaremetalPropertiesSlice() makes a slice of BaremetalProperties
func MakeBaremetalPropertiesSlice() []*BaremetalProperties {
	return []*BaremetalProperties{}
}
