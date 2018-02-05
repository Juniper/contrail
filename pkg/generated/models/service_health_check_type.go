package models

// ServiceHealthCheckType

import "encoding/json"

// ServiceHealthCheckType
type ServiceHealthCheckType struct {
	HealthCheckType HealthCheckType         `json:"health_check_type,omitempty"`
	HTTPMethod      string                  `json:"http_method,omitempty"`
	Timeout         int                     `json:"timeout,omitempty"`
	URLPath         string                  `json:"url_path,omitempty"`
	DelayUsecs      int                     `json:"delayUsecs,omitempty"`
	Delay           int                     `json:"delay,omitempty"`
	ExpectedCodes   string                  `json:"expected_codes,omitempty"`
	MaxRetries      int                     `json:"max_retries,omitempty"`
	MonitorType     HealthCheckProtocolType `json:"monitor_type,omitempty"`
	TimeoutUsecs    int                     `json:"timeoutUsecs,omitempty"`
	Enabled         bool                    `json:"enabled"`
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
		Delay:           0,
		HealthCheckType: MakeHealthCheckType(),
		HTTPMethod:      "",
		MonitorType:     MakeHealthCheckProtocolType(),
		TimeoutUsecs:    0,
		Enabled:         false,
		ExpectedCodes:   "",
		MaxRetries:      0,
	}
}

// MakeServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
	return []*ServiceHealthCheckType{}
}
