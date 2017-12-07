package models

// LoadbalancerHealthmonitorType

import "encoding/json"

// LoadbalancerHealthmonitorType
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

//  parents relation object

// String returns json representation of the object
func (model *LoadbalancerHealthmonitorType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerHealthmonitorType makes LoadbalancerHealthmonitorType
func MakeLoadbalancerHealthmonitorType() *LoadbalancerHealthmonitorType {
	return &LoadbalancerHealthmonitorType{
		//TODO(nati): Apply default
		MaxRetries:    0,
		HTTPMethod:    "",
		AdminState:    false,
		Timeout:       0,
		URLPath:       "",
		MonitorType:   MakeHealthmonitorType(),
		Delay:         0,
		ExpectedCodes: "",
	}
}

// InterfaceToLoadbalancerHealthmonitorType makes LoadbalancerHealthmonitorType from interface
func InterfaceToLoadbalancerHealthmonitorType(iData interface{}) *LoadbalancerHealthmonitorType {
	data := iData.(map[string]interface{})
	return &LoadbalancerHealthmonitorType{
		Delay: data["delay"].(int),

		//{"Title":"","Description":"Time in seconds  at which health check is repeated","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Delay","GoType":"int","GoPremitive":true}
		ExpectedCodes: data["expected_codes"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, expected return code for HTTP operations like 200 ok.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ExpectedCodes","GoType":"string","GoPremitive":true}
		MaxRetries: data["max_retries"].(int),

		//{"Title":"","Description":"Number of failures before declaring health bad","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MaxRetries","GoType":"int","GoPremitive":true}
		HTTPMethod: data["http_method"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, type of http method used like GET, PUT, POST etc","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"HTTPMethod","GoType":"string","GoPremitive":true}
		AdminState: data["admin_state"].(bool),

		//{"Title":"","Description":"Administratively up or dowm.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AdminState","GoType":"bool","GoPremitive":true}
		Timeout: data["timeout"].(int),

		//{"Title":"","Description":"Time in seconds to wait for response","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Timeout","GoType":"int","GoPremitive":true}
		URLPath: data["url_path"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, URL to be used. In case of ICMP, ip address","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"URLPath","GoType":"string","GoPremitive":true}
		MonitorType: InterfaceToHealthmonitorType(data["monitor_type"]),

		//{"Title":"","Description":"Protocol used to monitor health, PING, HTTP, HTTPS or TCP","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["PING","TCP","HTTP","HTTPS"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/HealthmonitorType","CollectionType":"","Column":"","Item":null,"GoName":"MonitorType","GoType":"HealthmonitorType","GoPremitive":false}

	}
}

// InterfaceToLoadbalancerHealthmonitorTypeSlice makes a slice of LoadbalancerHealthmonitorType from interface
func InterfaceToLoadbalancerHealthmonitorTypeSlice(data interface{}) []*LoadbalancerHealthmonitorType {
	list := data.([]interface{})
	result := MakeLoadbalancerHealthmonitorTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerHealthmonitorType(item))
	}
	return result
}

// MakeLoadbalancerHealthmonitorTypeSlice() makes a slice of LoadbalancerHealthmonitorType
func MakeLoadbalancerHealthmonitorTypeSlice() []*LoadbalancerHealthmonitorType {
	return []*LoadbalancerHealthmonitorType{}
}
