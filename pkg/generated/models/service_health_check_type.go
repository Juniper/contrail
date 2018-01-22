package models

// ServiceHealthCheckType

import "encoding/json"

// ServiceHealthCheckType
type ServiceHealthCheckType struct {
	DelayUsecs      int                     `json:"delayUsecs,omitempty"`
	ExpectedCodes   string                  `json:"expected_codes,omitempty"`
	MaxRetries      int                     `json:"max_retries,omitempty"`
	HealthCheckType HealthCheckType         `json:"health_check_type,omitempty"`
	HTTPMethod      string                  `json:"http_method,omitempty"`
	Timeout         int                     `json:"timeout,omitempty"`
	TimeoutUsecs    int                     `json:"timeoutUsecs,omitempty"`
	Enabled         bool                    `json:"enabled"`
	Delay           int                     `json:"delay,omitempty"`
	URLPath         string                  `json:"url_path,omitempty"`
	MonitorType     HealthCheckProtocolType `json:"monitor_type,omitempty"`
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
		MonitorType:     MakeHealthCheckProtocolType(),
		TimeoutUsecs:    0,
		Enabled:         false,
		Delay:           0,
		URLPath:         "",
		HTTPMethod:      "",
		Timeout:         0,
		DelayUsecs:      0,
		ExpectedCodes:   "",
		MaxRetries:      0,
		HealthCheckType: MakeHealthCheckType(),
	}
}

// MakeServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
	return []*ServiceHealthCheckType{}
}
