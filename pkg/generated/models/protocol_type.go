package models

// ProtocolType

import "encoding/json"

type ProtocolType struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

func (model *ProtocolType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeProtocolType() *ProtocolType {
	return &ProtocolType{
		//TODO(nati): Apply default
		Protocol: "",
		Port:     0,
	}
}

func InterfaceToProtocolType(iData interface{}) *ProtocolType {
	data := iData.(map[string]interface{})
	return &ProtocolType{
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"}
		Port: data["port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Port","GoType":"int"}

	}
}

func InterfaceToProtocolTypeSlice(data interface{}) []*ProtocolType {
	list := data.([]interface{})
	result := MakeProtocolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToProtocolType(item))
	}
	return result
}

func MakeProtocolTypeSlice() []*ProtocolType {
	return []*ProtocolType{}
}
