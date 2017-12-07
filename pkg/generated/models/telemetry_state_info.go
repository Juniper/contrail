package models

// TelemetryStateInfo

import "encoding/json"

// TelemetryStateInfo
type TelemetryStateInfo struct {
	Resource   []*TelemetryResourceInfo `json:"resource"`
	ServerPort int                      `json:"server_port"`
	ServerIP   string                   `json:"server_ip"`
}

//  parents relation object

// String returns json representation of the object
func (model *TelemetryStateInfo) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeTelemetryStateInfo makes TelemetryStateInfo
func MakeTelemetryStateInfo() *TelemetryStateInfo {
	return &TelemetryStateInfo{
		//TODO(nati): Apply default

		Resource: MakeTelemetryResourceInfoSlice(),

		ServerPort: 0,
		ServerIP:   "",
	}
}

// InterfaceToTelemetryStateInfo makes TelemetryStateInfo from interface
func InterfaceToTelemetryStateInfo(iData interface{}) *TelemetryStateInfo {
	data := iData.(map[string]interface{})
	return &TelemetryStateInfo{

		Resource: InterfaceToTelemetryResourceInfoSlice(data["resource"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string","GoPremitive":true},"path":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Path","GoType":"string","GoPremitive":true},"rate":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Rate","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TelemetryResourceInfo","CollectionType":"","Column":"","Item":null,"GoName":"Resource","GoType":"TelemetryResourceInfo","GoPremitive":false},"GoName":"Resource","GoType":"[]*TelemetryResourceInfo","GoPremitive":true}
		ServerPort: data["server_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServerPort","GoType":"int","GoPremitive":true}
		ServerIP: data["server_ip"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServerIP","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToTelemetryStateInfoSlice makes a slice of TelemetryStateInfo from interface
func InterfaceToTelemetryStateInfoSlice(data interface{}) []*TelemetryStateInfo {
	list := data.([]interface{})
	result := MakeTelemetryStateInfoSlice()
	for _, item := range list {
		result = append(result, InterfaceToTelemetryStateInfo(item))
	}
	return result
}

// MakeTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
	return []*TelemetryStateInfo{}
}
