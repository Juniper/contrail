package models

// LoadbalancerHealthmonitorType

import "encoding/json"

type LoadbalancerHealthmonitorType struct {
	MonitorType   HealthmonitorType `json:"monitor_type"`
	Delay         int               `json:"delay"`
	ExpectedCodes string            `json:"expected_codes"`
	MaxRetries    int               `json:"max_retries"`
	HTTPMethod    string            `json:"http_method"`
	AdminState    bool              `json:"admin_state"`
	Timeout       int               `json:"timeout"`
	URLPath       string            `json:"url_path"`
}

func (model *LoadbalancerHealthmonitorType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeLoadbalancerHealthmonitorType() *LoadbalancerHealthmonitorType {
	return &LoadbalancerHealthmonitorType{
		//TODO(nati): Apply default
		Delay:         0,
		ExpectedCodes: "",
		MaxRetries:    0,
		HTTPMethod:    "",
		AdminState:    false,
		Timeout:       0,
		URLPath:       "",
		MonitorType:   MakeHealthmonitorType(),
	}
}

func InterfaceToLoadbalancerHealthmonitorType(iData interface{}) *LoadbalancerHealthmonitorType {
	data := iData.(map[string]interface{})
	return &LoadbalancerHealthmonitorType{
		ExpectedCodes: data["expected_codes"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, expected return code for HTTP operations like 200 ok.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ExpectedCodes","GoType":"string"}
		MaxRetries: data["max_retries"].(int),

		//{"Title":"","Description":"Number of failures before declaring health bad","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MaxRetries","GoType":"int"}
		HTTPMethod: data["http_method"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, type of http method used like GET, PUT, POST etc","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"HTTPMethod","GoType":"string"}
		AdminState: data["admin_state"].(bool),

		//{"Title":"","Description":"Administratively up or dowm.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AdminState","GoType":"bool"}
		Timeout: data["timeout"].(int),

		//{"Title":"","Description":"Time in seconds to wait for response","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Timeout","GoType":"int"}
		URLPath: data["url_path"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, URL to be used. In case of ICMP, ip address","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"URLPath","GoType":"string"}
		MonitorType: InterfaceToHealthmonitorType(data["monitor_type"]),

		//{"Title":"","Description":"Protocol used to monitor health, PING, HTTP, HTTPS or TCP","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["PING","TCP","HTTP","HTTPS"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/HealthmonitorType","CollectionType":"","Column":"","Item":null,"GoName":"MonitorType","GoType":"HealthmonitorType"}
		Delay: data["delay"].(int),

		//{"Title":"","Description":"Time in seconds  at which health check is repeated","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Delay","GoType":"int"}

	}
}

func InterfaceToLoadbalancerHealthmonitorTypeSlice(data interface{}) []*LoadbalancerHealthmonitorType {
	list := data.([]interface{})
	result := MakeLoadbalancerHealthmonitorTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerHealthmonitorType(item))
	}
	return result
}

func MakeLoadbalancerHealthmonitorTypeSlice() []*LoadbalancerHealthmonitorType {
	return []*LoadbalancerHealthmonitorType{}
}
