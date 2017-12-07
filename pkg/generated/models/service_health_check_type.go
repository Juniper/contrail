package models

// ServiceHealthCheckType

import "encoding/json"

// ServiceHealthCheckType
type ServiceHealthCheckType struct {
	URLPath         string                  `json:"url_path"`
	MonitorType     HealthCheckProtocolType `json:"monitor_type"`
	ExpectedCodes   string                  `json:"expected_codes"`
	Timeout         int                     `json:"timeout"`
	Enabled         bool                    `json:"enabled"`
	Delay           int                     `json:"delay"`
	MaxRetries      int                     `json:"max_retries"`
	HealthCheckType HealthCheckType         `json:"health_check_type"`
	HTTPMethod      string                  `json:"http_method"`
	DelayUsecs      int                     `json:"delayUsecs"`
	TimeoutUsecs    int                     `json:"timeoutUsecs"`
}

//  parents relation object

// String returns json representation of the object
func (model *ServiceHealthCheckType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceHealthCheckType makes ServiceHealthCheckType
func MakeServiceHealthCheckType() *ServiceHealthCheckType {
	return &ServiceHealthCheckType{
		//TODO(nati): Apply default
		ExpectedCodes:   "",
		Timeout:         0,
		URLPath:         "",
		MonitorType:     MakeHealthCheckProtocolType(),
		HealthCheckType: MakeHealthCheckType(),
		HTTPMethod:      "",
		DelayUsecs:      0,
		TimeoutUsecs:    0,
		Enabled:         false,
		Delay:           0,
		MaxRetries:      0,
	}
}

// InterfaceToServiceHealthCheckType makes ServiceHealthCheckType from interface
func InterfaceToServiceHealthCheckType(iData interface{}) *ServiceHealthCheckType {
	data := iData.(map[string]interface{})
	return &ServiceHealthCheckType{
		DelayUsecs: data["delayUsecs"].(int),

		//{"Title":"","Description":"Time in micro seconds at which health check is repeated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DelayUsecs","GoType":"int","GoPremitive":true}
		TimeoutUsecs: data["timeoutUsecs"].(int),

		//{"Title":"","Description":"Time in micro seconds to wait for response","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TimeoutUsecs","GoType":"int","GoPremitive":true}
		Enabled: data["enabled"].(bool),

		//{"Title":"","Description":"Administratively enable or disable this health check.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Enabled","GoType":"bool","GoPremitive":true}
		Delay: data["delay"].(int),

		//{"Title":"","Description":"Time in seconds at which health check is repeated","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Delay","GoType":"int","GoPremitive":true}
		MaxRetries: data["max_retries"].(int),

		//{"Title":"","Description":"Number of failures before declaring health bad","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MaxRetries","GoType":"int","GoPremitive":true}
		HealthCheckType: InterfaceToHealthCheckType(data["health_check_type"]),

		//{"Title":"","Description":"Health check type, currently only link-local, end-to-end and segment are supported","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["link-local","end-to-end","segment"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/HealthCheckType","CollectionType":"","Column":"","Item":null,"GoName":"HealthCheckType","GoType":"HealthCheckType","GoPremitive":false}
		HTTPMethod: data["http_method"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, type of http method used like GET, PUT, POST etc","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"HTTPMethod","GoType":"string","GoPremitive":true}
		ExpectedCodes: data["expected_codes"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, expected return code for HTTP operations like 200 ok.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ExpectedCodes","GoType":"string","GoPremitive":true}
		Timeout: data["timeout"].(int),

		//{"Title":"","Description":"Time in seconds to wait for response","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Timeout","GoType":"int","GoPremitive":true}
		URLPath: data["url_path"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, URL to be used. In case of ICMP, ip address","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"URLPath","GoType":"string","GoPremitive":true}
		MonitorType: InterfaceToHealthCheckProtocolType(data["monitor_type"]),

		//{"Title":"","Description":"Protocol used to monitor health, currently only HTTP, ICMP(ping), and BFD are supported","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["PING","HTTP","BFD"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/HealthCheckProtocolType","CollectionType":"","Column":"","Item":null,"GoName":"MonitorType","GoType":"HealthCheckProtocolType","GoPremitive":false}

	}
}

// InterfaceToServiceHealthCheckTypeSlice makes a slice of ServiceHealthCheckType from interface
func InterfaceToServiceHealthCheckTypeSlice(data interface{}) []*ServiceHealthCheckType {
	list := data.([]interface{})
	result := MakeServiceHealthCheckTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceHealthCheckType(item))
	}
	return result
}

// MakeServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
	return []*ServiceHealthCheckType{}
}
