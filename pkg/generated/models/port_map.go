package models

// PortMap

import "encoding/json"

// PortMap
type PortMap struct {
	SRCPort  int    `json:"src_port"`
	Protocol string `json:"protocol"`
	DSTPort  int    `json:"dst_port"`
}

//  parents relation object

// String returns json representation of the object
func (model *PortMap) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePortMap makes PortMap
func MakePortMap() *PortMap {
	return &PortMap{
		//TODO(nati): Apply default
		SRCPort:  0,
		Protocol: "",
		DSTPort:  0,
	}
}

// InterfaceToPortMap makes PortMap from interface
func InterfaceToPortMap(iData interface{}) *PortMap {
	data := iData.(map[string]interface{})
	return &PortMap{
		SRCPort: data["src_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SRCPort","GoType":"int","GoPremitive":true}
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string","GoPremitive":true}
		DSTPort: data["dst_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DSTPort","GoType":"int","GoPremitive":true}

	}
}

// InterfaceToPortMapSlice makes a slice of PortMap from interface
func InterfaceToPortMapSlice(data interface{}) []*PortMap {
	list := data.([]interface{})
	result := MakePortMapSlice()
	for _, item := range list {
		result = append(result, InterfaceToPortMap(item))
	}
	return result
}

// MakePortMapSlice() makes a slice of PortMap
func MakePortMapSlice() []*PortMap {
	return []*PortMap{}
}
