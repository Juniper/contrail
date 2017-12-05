package models

// PortType

import "encoding/json"

type PortType struct {
	EndPort   L4PortType `json:"end_port"`
	StartPort L4PortType `json:"start_port"`
}

func (model *PortType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakePortType() *PortType {
	return &PortType{
		//TODO(nati): Apply default
		StartPort: MakeL4PortType(),
		EndPort:   MakeL4PortType(),
	}
}

func InterfaceToPortType(iData interface{}) *PortType {
	data := iData.(map[string]interface{})
	return &PortType{
		EndPort: InterfaceToL4PortType(data["end_port"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"}
		StartPort: InterfaceToL4PortType(data["start_port"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}

	}
}

func InterfaceToPortTypeSlice(data interface{}) []*PortType {
	list := data.([]interface{})
	result := MakePortTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPortType(item))
	}
	return result
}

func MakePortTypeSlice() []*PortType {
	return []*PortType{}
}
