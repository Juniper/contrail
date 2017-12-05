package models

// JunosServicePorts

import "encoding/json"

type JunosServicePorts struct {
	ServicePort []string `json:"service_port"`
}

func (model *JunosServicePorts) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeJunosServicePorts() *JunosServicePorts {
	return &JunosServicePorts{
		//TODO(nati): Apply default
		ServicePort: []string{},
	}
}

func InterfaceToJunosServicePorts(iData interface{}) *JunosServicePorts {
	data := iData.(map[string]interface{})
	return &JunosServicePorts{
		ServicePort: data["service_port"].([]string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServicePort","GoType":"string"},"GoName":"ServicePort","GoType":"[]string"}

	}
}

func InterfaceToJunosServicePortsSlice(data interface{}) []*JunosServicePorts {
	list := data.([]interface{})
	result := MakeJunosServicePortsSlice()
	for _, item := range list {
		result = append(result, InterfaceToJunosServicePorts(item))
	}
	return result
}

func MakeJunosServicePortsSlice() []*JunosServicePorts {
	return []*JunosServicePorts{}
}
