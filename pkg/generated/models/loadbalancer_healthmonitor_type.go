package models

// LoadbalancerHealthmonitorType

import "encoding/json"

// LoadbalancerHealthmonitorType
type LoadbalancerHealthmonitorType struct {
	HTTPMethod    string            `json:"http_method"`
	AdminState    bool              `json:"admin_state"`
	Timeout       int               `json:"timeout"`
	URLPath       string            `json:"url_path"`
	MonitorType   HealthmonitorType `json:"monitor_type"`
	Delay         int               `json:"delay"`
	ExpectedCodes string            `json:"expected_codes"`
	MaxRetries    int               `json:"max_retries"`
}

// String returns json representation of the object
func (model *LoadbalancerHealthmonitorType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerHealthmonitorType makes LoadbalancerHealthmonitorType
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

// InterfaceToLoadbalancerHealthmonitorType makes LoadbalancerHealthmonitorType from interface
func InterfaceToLoadbalancerHealthmonitorType(iData interface{}) *LoadbalancerHealthmonitorType {
	data := iData.(map[string]interface{})
	return &LoadbalancerHealthmonitorType{
		URLPath: data["url_path"].(string),

		//{"description":"In case monitor protocol is HTTP, URL to be used. In case of ICMP, ip address","type":"string"}
		MonitorType: InterfaceToHealthmonitorType(data["monitor_type"]),

		//{"description":"Protocol used to monitor health, PING, HTTP, HTTPS or TCP","type":"string","enum":["PING","TCP","HTTP","HTTPS"]}
		Delay: data["delay"].(int),

		//{"description":"Time in seconds  at which health check is repeated","type":"integer"}
		ExpectedCodes: data["expected_codes"].(string),

		//{"description":"In case monitor protocol is HTTP, expected return code for HTTP operations like 200 ok.","type":"string"}
		MaxRetries: data["max_retries"].(int),

		//{"description":"Number of failures before declaring health bad","type":"integer"}
		HTTPMethod: data["http_method"].(string),

		//{"description":"In case monitor protocol is HTTP, type of http method used like GET, PUT, POST etc","type":"string"}
		AdminState: data["admin_state"].(bool),

		//{"description":"Administratively up or dowm.","type":"boolean"}
		Timeout: data["timeout"].(int),

		//{"description":"Time in seconds to wait for response","type":"integer"}

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
