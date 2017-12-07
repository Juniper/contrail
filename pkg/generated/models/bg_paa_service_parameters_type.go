package models

// BGPaaServiceParametersType

import "encoding/json"

// BGPaaServiceParametersType
type BGPaaServiceParametersType struct {
	PortStart L4PortType `json:"port_start"`
	PortEnd   L4PortType `json:"port_end"`
}

//  parents relation object

// String returns json representation of the object
func (model *BGPaaServiceParametersType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBGPaaServiceParametersType makes BGPaaServiceParametersType
func MakeBGPaaServiceParametersType() *BGPaaServiceParametersType {
	return &BGPaaServiceParametersType{
		//TODO(nati): Apply default
		PortEnd:   MakeL4PortType(),
		PortStart: MakeL4PortType(),
	}
}

// InterfaceToBGPaaServiceParametersType makes BGPaaServiceParametersType from interface
func InterfaceToBGPaaServiceParametersType(iData interface{}) *BGPaaServiceParametersType {
	data := iData.(map[string]interface{})
	return &BGPaaServiceParametersType{
		PortStart: InterfaceToL4PortType(data["port_start"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"PortStart","GoType":"L4PortType","GoPremitive":false}
		PortEnd: InterfaceToL4PortType(data["port_end"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"PortEnd","GoType":"L4PortType","GoPremitive":false}

	}
}

// InterfaceToBGPaaServiceParametersTypeSlice makes a slice of BGPaaServiceParametersType from interface
func InterfaceToBGPaaServiceParametersTypeSlice(data interface{}) []*BGPaaServiceParametersType {
	list := data.([]interface{})
	result := MakeBGPaaServiceParametersTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToBGPaaServiceParametersType(item))
	}
	return result
}

// MakeBGPaaServiceParametersTypeSlice() makes a slice of BGPaaServiceParametersType
func MakeBGPaaServiceParametersTypeSlice() []*BGPaaServiceParametersType {
	return []*BGPaaServiceParametersType{}
}
