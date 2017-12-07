package models

// ServiceApplianceInterfaceType

import "encoding/json"

// ServiceApplianceInterfaceType
type ServiceApplianceInterfaceType struct {
	InterfaceType ServiceInterfaceType `json:"interface_type"`
}

//  parents relation object

// String returns json representation of the object
func (model *ServiceApplianceInterfaceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceApplianceInterfaceType makes ServiceApplianceInterfaceType
func MakeServiceApplianceInterfaceType() *ServiceApplianceInterfaceType {
	return &ServiceApplianceInterfaceType{
		//TODO(nati): Apply default
		InterfaceType: MakeServiceInterfaceType(),
	}
}

// InterfaceToServiceApplianceInterfaceType makes ServiceApplianceInterfaceType from interface
func InterfaceToServiceApplianceInterfaceType(iData interface{}) *ServiceApplianceInterfaceType {
	data := iData.(map[string]interface{})
	return &ServiceApplianceInterfaceType{
		InterfaceType: InterfaceToServiceInterfaceType(data["interface_type"]),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceInterfaceType","CollectionType":"","Column":"interface_type","Item":null,"GoName":"InterfaceType","GoType":"ServiceInterfaceType","GoPremitive":false}

	}
}

// InterfaceToServiceApplianceInterfaceTypeSlice makes a slice of ServiceApplianceInterfaceType from interface
func InterfaceToServiceApplianceInterfaceTypeSlice(data interface{}) []*ServiceApplianceInterfaceType {
	list := data.([]interface{})
	result := MakeServiceApplianceInterfaceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceApplianceInterfaceType(item))
	}
	return result
}

// MakeServiceApplianceInterfaceTypeSlice() makes a slice of ServiceApplianceInterfaceType
func MakeServiceApplianceInterfaceTypeSlice() []*ServiceApplianceInterfaceType {
	return []*ServiceApplianceInterfaceType{}
}
