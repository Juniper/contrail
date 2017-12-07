package models

// DomainLimitsType

import "encoding/json"

// DomainLimitsType
type DomainLimitsType struct {
	ProjectLimit        int `json:"project_limit"`
	VirtualNetworkLimit int `json:"virtual_network_limit"`
	SecurityGroupLimit  int `json:"security_group_limit"`
}

//  parents relation object

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
		ProjectLimit: data["project_limit"].(int),

		//{"Title":"","Description":"Maximum number of projects allowed in this domain","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProjectLimit","GoType":"int","GoPremitive":true}
		VirtualNetworkLimit: data["virtual_network_limit"].(int),

		//{"Title":"","Description":"Maximum number of virtual networks allowed in this domain","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetworkLimit","GoType":"int","GoPremitive":true}
		SecurityGroupLimit: data["security_group_limit"].(int),

		//{"Title":"","Description":"Maximum number of security groups allowed in this domain","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroupLimit","GoType":"int","GoPremitive":true}

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
