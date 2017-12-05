package models

// FlowAgingTimeout

import "encoding/json"

type FlowAgingTimeout struct {
	TimeoutInSeconds int    `json:"timeout_in_seconds"`
	Protocol         string `json:"protocol"`
	Port             int    `json:"port"`
}

func (model *FlowAgingTimeout) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFlowAgingTimeout() *FlowAgingTimeout {
	return &FlowAgingTimeout{
		//TODO(nati): Apply default
		Protocol:         "",
		Port:             0,
		TimeoutInSeconds: 0,
	}
}

func InterfaceToFlowAgingTimeout(iData interface{}) *FlowAgingTimeout {
	data := iData.(map[string]interface{})
	return &FlowAgingTimeout{
		TimeoutInSeconds: data["timeout_in_seconds"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TimeoutInSeconds","GoType":"int"}
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"}
		Port: data["port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Port","GoType":"int"}

	}
}

func InterfaceToFlowAgingTimeoutSlice(data interface{}) []*FlowAgingTimeout {
	list := data.([]interface{})
	result := MakeFlowAgingTimeoutSlice()
	for _, item := range list {
		result = append(result, InterfaceToFlowAgingTimeout(item))
	}
	return result
}

func MakeFlowAgingTimeoutSlice() []*FlowAgingTimeout {
	return []*FlowAgingTimeout{}
}
