package models

// DomainLimitsType

import "encoding/json"

// DomainLimitsType
type DomainLimitsType struct {
	ProjectLimit        int `json:"project_limit"`
	VirtualNetworkLimit int `json:"virtual_network_limit"`
	SecurityGroupLimit  int `json:"security_group_limit"`
}

// String returns json representation of the object
func (model *DomainLimitsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDomainLimitsType makes DomainLimitsType
func MakeDomainLimitsType() *DomainLimitsType {
	return &DomainLimitsType{
		//TODO(nati): Apply default
		SecurityGroupLimit:  0,
		ProjectLimit:        0,
		VirtualNetworkLimit: 0,
	}
}

// MakeDomainLimitsTypeSlice() makes a slice of DomainLimitsType
func MakeDomainLimitsTypeSlice() []*DomainLimitsType {
	return []*DomainLimitsType{}
}
