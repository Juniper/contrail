package models

// ServiceHealthCheckType

import "encoding/json"

// ServiceHealthCheckType
type ServiceHealthCheckType struct {
	HealthCheckType HealthCheckType         `json:"health_check_type"`
	MonitorType     HealthCheckProtocolType `json:"monitor_type"`
	Delay           int                     `json:"delay"`
	ExpectedCodes   string                  `json:"expected_codes"`
	Enabled         bool                    `json:"enabled"`
	MaxRetries      int                     `json:"max_retries"`
	HTTPMethod      string                  `json:"http_method"`
	Timeout         int                     `json:"timeout"`
	URLPath         string                  `json:"url_path"`
	DelayUsecs      int                     `json:"delayUsecs"`
	TimeoutUsecs    int                     `json:"timeoutUsecs"`
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
		Timeout:         0,
		URLPath:         "",
		DelayUsecs:      0,
		TimeoutUsecs:    0,
		Enabled:         false,
		MaxRetries:      0,
		HTTPMethod:      "",
		Delay:           0,
		ExpectedCodes:   "",
		HealthCheckType: MakeHealthCheckType(),
		MonitorType:     MakeHealthCheckProtocolType(),
	}
}

// InterfaceToServiceHealthCheckType makes ServiceHealthCheckType from interface
func InterfaceToServiceHealthCheckType(iData interface{}) *ServiceHealthCheckType {
	data := iData.(map[string]interface{})
	return &ServiceHealthCheckType{
		Enabled: data["enabled"].(bool),

		//{"description":"Administratively enable or disable this health check.","type":"boolean"}
		MaxRetries: data["max_retries"].(int),

		//{"description":"Number of failures before declaring health bad","type":"integer"}
		HTTPMethod: data["http_method"].(string),

		//{"description":"In case monitor protocol is HTTP, type of http method used like GET, PUT, POST etc","type":"string"}
		Timeout: data["timeout"].(int),

		//{"description":"Time in seconds to wait for response","type":"integer"}
		URLPath: data["url_path"].(string),

		//{"description":"In case monitor protocol is HTTP, URL to be used. In case of ICMP, ip address","type":"string"}
		DelayUsecs: data["delayUsecs"].(int),

		//{"description":"Time in micro seconds at which health check is repeated","type":"integer"}
		TimeoutUsecs: data["timeoutUsecs"].(int),

		//{"description":"Time in micro seconds to wait for response","type":"integer"}
		HealthCheckType: InterfaceToHealthCheckType(data["health_check_type"]),

		//{"description":"Health check type, currently only link-local, end-to-end and segment are supported","type":"string","enum":["link-local","end-to-end","segment"]}
		MonitorType: InterfaceToHealthCheckProtocolType(data["monitor_type"]),

		//{"description":"Protocol used to monitor health, currently only HTTP, ICMP(ping), and BFD are supported","type":"string","enum":["PING","HTTP","BFD"]}
		Delay: data["delay"].(int),

		//{"description":"Time in seconds at which health check is repeated","type":"integer"}
		ExpectedCodes: data["expected_codes"].(string),

		//{"description":"In case monitor protocol is HTTP, expected return code for HTTP operations like 200 ok.","type":"string"}

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
