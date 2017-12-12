package models

// ServiceHealthCheckType

import "encoding/json"

// ServiceHealthCheckType
type ServiceHealthCheckType struct {
	Enabled         bool                    `json:"enabled"`
	ExpectedCodes   string                  `json:"expected_codes"`
	MaxRetries      int                     `json:"max_retries"`
	HealthCheckType HealthCheckType         `json:"health_check_type"`
	DelayUsecs      int                     `json:"delayUsecs"`
	TimeoutUsecs    int                     `json:"timeoutUsecs"`
	Delay           int                     `json:"delay"`
	HTTPMethod      string                  `json:"http_method"`
	Timeout         int                     `json:"timeout"`
	URLPath         string                  `json:"url_path"`
	MonitorType     HealthCheckProtocolType `json:"monitor_type"`
}

// String returns json representation of the object
func (model *ServiceHealthCheckType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceHealthCheckType makes ServiceHealthCheckType
func MakeServiceHealthCheckType() *ServiceHealthCheckType {
	return &ServiceHealthCheckType{
		//TODO(nati): Apply default
		DelayUsecs:      0,
		TimeoutUsecs:    0,
		Delay:           0,
		HTTPMethod:      "",
		Timeout:         0,
		URLPath:         "",
		MonitorType:     MakeHealthCheckProtocolType(),
		Enabled:         false,
		ExpectedCodes:   "",
		MaxRetries:      0,
		HealthCheckType: MakeHealthCheckType(),
	}
}

// InterfaceToServiceHealthCheckType makes ServiceHealthCheckType from interface
func InterfaceToServiceHealthCheckType(iData interface{}) *ServiceHealthCheckType {
	data := iData.(map[string]interface{})
	return &ServiceHealthCheckType{
		ExpectedCodes: data["expected_codes"].(string),

		//{"description":"In case monitor protocol is HTTP, expected return code for HTTP operations like 200 ok.","type":"string"}
		MaxRetries: data["max_retries"].(int),

		//{"description":"Number of failures before declaring health bad","type":"integer"}
		HealthCheckType: InterfaceToHealthCheckType(data["health_check_type"]),

		//{"description":"Health check type, currently only link-local, end-to-end and segment are supported","type":"string","enum":["link-local","end-to-end","segment"]}
		Enabled: data["enabled"].(bool),

		//{"description":"Administratively enable or disable this health check.","type":"boolean"}
		TimeoutUsecs: data["timeoutUsecs"].(int),

		//{"description":"Time in micro seconds to wait for response","type":"integer"}
		Delay: data["delay"].(int),

		//{"description":"Time in seconds at which health check is repeated","type":"integer"}
		HTTPMethod: data["http_method"].(string),

		//{"description":"In case monitor protocol is HTTP, type of http method used like GET, PUT, POST etc","type":"string"}
		Timeout: data["timeout"].(int),

		//{"description":"Time in seconds to wait for response","type":"integer"}
		URLPath: data["url_path"].(string),

		//{"description":"In case monitor protocol is HTTP, URL to be used. In case of ICMP, ip address","type":"string"}
		MonitorType: InterfaceToHealthCheckProtocolType(data["monitor_type"]),

		//{"description":"Protocol used to monitor health, currently only HTTP, ICMP(ping), and BFD are supported","type":"string","enum":["PING","HTTP","BFD"]}
		DelayUsecs: data["delayUsecs"].(int),

		//{"description":"Time in micro seconds at which health check is repeated","type":"integer"}

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
