package models

// DomainLimitsType

import "encoding/json"

// DomainLimitsType
type DomainLimitsType struct {
	ProjectLimit        int `json:"project_limit,omitempty"`
	VirtualNetworkLimit int `json:"virtual_network_limit,omitempty"`
	SecurityGroupLimit  int `json:"security_group_limit,omitempty"`
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
		ProjectLimit:        0,
		VirtualNetworkLimit: 0,
		SecurityGroupLimit:  0,
	}
}

// MakeDomainLimitsTypeSlice() makes a slice of DomainLimitsType
func MakeDomainLimitsTypeSlice() []*DomainLimitsType {
	return []*DomainLimitsType{}
}
