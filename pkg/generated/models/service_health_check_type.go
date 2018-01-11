package models

// ServiceHealthCheckType

import "encoding/json"

// ServiceHealthCheckType
type ServiceHealthCheckType struct {
	Enabled         bool                    `json:"enabled"`
	Delay           int                     `json:"delay"`
	ExpectedCodes   string                  `json:"expected_codes"`
	MaxRetries      int                     `json:"max_retries"`
	URLPath         string                  `json:"url_path"`
	MonitorType     HealthCheckProtocolType `json:"monitor_type"`
	DelayUsecs      int                     `json:"delayUsecs"`
	HealthCheckType HealthCheckType         `json:"health_check_type"`
	HTTPMethod      string                  `json:"http_method"`
	Timeout         int                     `json:"timeout"`
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
		TimeoutUsecs:    0,
		HealthCheckType: MakeHealthCheckType(),
		HTTPMethod:      "",
		ExpectedCodes:   "",
		MaxRetries:      0,
		URLPath:         "",
		MonitorType:     MakeHealthCheckProtocolType(),
		DelayUsecs:      0,
		Enabled:         false,
		Delay:           0,
	}
}

// MakeServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
	return []*ServiceHealthCheckType{}
}
