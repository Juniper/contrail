package models

// FlowAgingTimeout

import "encoding/json"

// FlowAgingTimeout
type FlowAgingTimeout struct {
	TimeoutInSeconds int    `json:"timeout_in_seconds"`
	Protocol         string `json:"protocol"`
	Port             int    `json:"port"`
}

//  parents relation object

// String returns json representation of the object
func (model *FlowAgingTimeout) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFlowAgingTimeout makes FlowAgingTimeout
func MakeFlowAgingTimeout() *FlowAgingTimeout {
	return &FlowAgingTimeout{
		//TODO(nati): Apply default
		TimeoutInSeconds: 0,
		Protocol:         "",
		Port:             0,
	}
}

// InterfaceToFlowAgingTimeout makes FlowAgingTimeout from interface
func InterfaceToFlowAgingTimeout(iData interface{}) *FlowAgingTimeout {
	data := iData.(map[string]interface{})
	return &FlowAgingTimeout{
		TimeoutInSeconds: data["timeout_in_seconds"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TimeoutInSeconds","GoType":"int","GoPremitive":true}
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string","GoPremitive":true}
		Port: data["port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Port","GoType":"int","GoPremitive":true}

	}
}

// InterfaceToFlowAgingTimeoutSlice makes a slice of FlowAgingTimeout from interface
func InterfaceToFlowAgingTimeoutSlice(data interface{}) []*FlowAgingTimeout {
	list := data.([]interface{})
	result := MakeFlowAgingTimeoutSlice()
	for _, item := range list {
		result = append(result, InterfaceToFlowAgingTimeout(item))
	}
	return result
}

// MakeFlowAgingTimeoutSlice() makes a slice of FlowAgingTimeout
func MakeFlowAgingTimeoutSlice() []*FlowAgingTimeout {
	return []*FlowAgingTimeout{}
}
