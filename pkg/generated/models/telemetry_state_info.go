package models

// TelemetryStateInfo

import "encoding/json"

type TelemetryStateInfo struct {
	Resource   []*TelemetryResourceInfo `json:"resource"`
	ServerPort int                      `json:"server_port"`
	ServerIP   string                   `json:"server_ip"`
}

func (model *TelemetryStateInfo) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeTelemetryStateInfo() *TelemetryStateInfo {
	return &TelemetryStateInfo{
		//TODO(nati): Apply default

		Resource: MakeTelemetryResourceInfoSlice(),

		ServerPort: 0,
		ServerIP:   "",
	}
}

func InterfaceToTelemetryStateInfo(iData interface{}) *TelemetryStateInfo {
	data := iData.(map[string]interface{})
	return &TelemetryStateInfo{
		ServerIP: data["server_ip"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServerIP","GoType":"string"}

		Resource: InterfaceToTelemetryResourceInfoSlice(data["resource"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string"},"path":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Path","GoType":"string"},"rate":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Rate","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TelemetryResourceInfo","CollectionType":"","Column":"","Item":null,"GoName":"Resource","GoType":"TelemetryResourceInfo"},"GoName":"Resource","GoType":"[]*TelemetryResourceInfo"}
		ServerPort: data["server_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServerPort","GoType":"int"}

	}
}

func InterfaceToTelemetryStateInfoSlice(data interface{}) []*TelemetryStateInfo {
	list := data.([]interface{})
	result := MakeTelemetryStateInfoSlice()
	for _, item := range list {
		result = append(result, InterfaceToTelemetryStateInfo(item))
	}
	return result
}

func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
	return []*TelemetryStateInfo{}
}
