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

// InterfaceToDomainLimitsType makes DomainLimitsType from interface
func InterfaceToDomainLimitsType(iData interface{}) *DomainLimitsType {
	data := iData.(map[string]interface{})
	return &DomainLimitsType{
		SecurityGroupLimit: data["security_group_limit"].(int),

		//{"description":"Maximum number of security groups allowed in this domain","type":"integer"}
		ProjectLimit: data["project_limit"].(int),

		//{"description":"Maximum number of projects allowed in this domain","type":"integer"}
		VirtualNetworkLimit: data["virtual_network_limit"].(int),

		//{"description":"Maximum number of virtual networks allowed in this domain","type":"integer"}

	}
}

// InterfaceToDomainLimitsTypeSlice makes a slice of DomainLimitsType from interface
func InterfaceToDomainLimitsTypeSlice(data interface{}) []*DomainLimitsType {
	list := data.([]interface{})
	result := MakeDomainLimitsTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDomainLimitsType(item))
	}
	return result
}

// MakeDomainLimitsTypeSlice() makes a slice of DomainLimitsType
func MakeDomainLimitsTypeSlice() []*DomainLimitsType {
	return []*DomainLimitsType{}
}
