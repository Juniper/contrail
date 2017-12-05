package models

// PortMap

import "encoding/json"

type PortMap struct {
	SRCPort  int    `json:"src_port"`
	Protocol string `json:"protocol"`
	DSTPort  int    `json:"dst_port"`
}

func (model *PortMap) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakePortMap() *PortMap {
	return &PortMap{
		//TODO(nati): Apply default
		SRCPort:  0,
		Protocol: "",
		DSTPort:  0,
	}
}

func InterfaceToPortMap(iData interface{}) *PortMap {
	data := iData.(map[string]interface{})
	return &PortMap{
		SRCPort: data["src_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SRCPort","GoType":"int"}
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"}
		DSTPort: data["dst_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DSTPort","GoType":"int"}

	}
}

func InterfaceToPortMapSlice(data interface{}) []*PortMap {
	list := data.([]interface{})
	result := MakePortMapSlice()
	for _, item := range list {
		result = append(result, InterfaceToPortMap(item))
	}
	return result
}

func MakePortMapSlice() []*PortMap {
	return []*PortMap{}
}
