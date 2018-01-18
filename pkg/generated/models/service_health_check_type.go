package models

// ServiceHealthCheckType

import "encoding/json"

// ServiceHealthCheckType
type ServiceHealthCheckType struct {
	Delay           int                     `json:"delay,omitempty"`
	MaxRetries      int                     `json:"max_retries,omitempty"`
	HealthCheckType HealthCheckType         `json:"health_check_type,omitempty"`
	Timeout         int                     `json:"timeout,omitempty"`
	MonitorType     HealthCheckProtocolType `json:"monitor_type,omitempty"`
	Enabled         bool                    `json:"enabled"`
	TimeoutUsecs    int                     `json:"timeoutUsecs,omitempty"`
	ExpectedCodes   string                  `json:"expected_codes,omitempty"`
	HTTPMethod      string                  `json:"http_method,omitempty"`
	URLPath         string                  `json:"url_path,omitempty"`
	DelayUsecs      int                     `json:"delayUsecs,omitempty"`
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
		ExpectedCodes:   "",
		HTTPMethod:      "",
		URLPath:         "",
		Enabled:         false,
		Delay:           0,
		MaxRetries:      0,
		HealthCheckType: MakeHealthCheckType(),
		Timeout:         0,
		MonitorType:     MakeHealthCheckProtocolType(),
	}
}

// MakeServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
	return []*ServiceHealthCheckType{}
}
