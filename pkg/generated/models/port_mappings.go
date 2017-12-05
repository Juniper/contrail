package models

// PortMappings

import "encoding/json"

type PortMappings struct {
	PortMappings []*PortMap `json:"port_mappings"`
}

func (model *PortMappings) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakePortMappings() *PortMappings {
	return &PortMappings{
		//TODO(nati): Apply default

		PortMappings: MakePortMapSlice(),
	}
}

func InterfaceToPortMappings(iData interface{}) *PortMappings {
	data := iData.(map[string]interface{})
	return &PortMappings{

		PortMappings: InterfaceToPortMapSlice(data["port_mappings"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"dst_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DSTPort","GoType":"int"},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"},"src_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SRCPort","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortMap","CollectionType":"","Column":"","Item":null,"GoName":"PortMappings","GoType":"PortMap"},"GoName":"PortMappings","GoType":"[]*PortMap"}

	}
}

func InterfaceToPortMappingsSlice(data interface{}) []*PortMappings {
	list := data.([]interface{})
	result := MakePortMappingsSlice()
	for _, item := range list {
		result = append(result, InterfaceToPortMappings(item))
	}
	return result
}

func MakePortMappingsSlice() []*PortMappings {
	return []*PortMappings{}
}
