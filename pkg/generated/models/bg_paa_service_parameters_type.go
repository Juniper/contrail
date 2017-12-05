package models

// BGPaaServiceParametersType

import "encoding/json"

type BGPaaServiceParametersType struct {
	PortStart L4PortType `json:"port_start"`
	PortEnd   L4PortType `json:"port_end"`
}

func (model *BGPaaServiceParametersType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeBGPaaServiceParametersType() *BGPaaServiceParametersType {
	return &BGPaaServiceParametersType{
		//TODO(nati): Apply default
		PortStart: MakeL4PortType(),
		PortEnd:   MakeL4PortType(),
	}
}

func InterfaceToBGPaaServiceParametersType(iData interface{}) *BGPaaServiceParametersType {
	data := iData.(map[string]interface{})
	return &BGPaaServiceParametersType{
		PortStart: InterfaceToL4PortType(data["port_start"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"PortStart","GoType":"L4PortType"}
		PortEnd: InterfaceToL4PortType(data["port_end"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"PortEnd","GoType":"L4PortType"}

	}
}

func InterfaceToBGPaaServiceParametersTypeSlice(data interface{}) []*BGPaaServiceParametersType {
	list := data.([]interface{})
	result := MakeBGPaaServiceParametersTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToBGPaaServiceParametersType(item))
	}
	return result
}

func MakeBGPaaServiceParametersTypeSlice() []*BGPaaServiceParametersType {
	return []*BGPaaServiceParametersType{}
}
