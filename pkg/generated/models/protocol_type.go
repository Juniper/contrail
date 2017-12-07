package models

// ProtocolType

import "encoding/json"

// ProtocolType
type ProtocolType struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

//  parents relation object

// String returns json representation of the object
func (model *ProtocolType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeProtocolType makes ProtocolType
func MakeProtocolType() *ProtocolType {
	return &ProtocolType{
		//TODO(nati): Apply default
		Protocol: "",
		Port:     0,
	}
}

// InterfaceToProtocolType makes ProtocolType from interface
func InterfaceToProtocolType(iData interface{}) *ProtocolType {
	data := iData.(map[string]interface{})
	return &ProtocolType{
		Port: data["port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Port","GoType":"int","GoPremitive":true}
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToProtocolTypeSlice makes a slice of ProtocolType from interface
func InterfaceToProtocolTypeSlice(data interface{}) []*ProtocolType {
	list := data.([]interface{})
	result := MakeProtocolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToProtocolType(item))
	}
	return result
}

// MakeProtocolTypeSlice() makes a slice of ProtocolType
func MakeProtocolTypeSlice() []*ProtocolType {
	return []*ProtocolType{}
}
