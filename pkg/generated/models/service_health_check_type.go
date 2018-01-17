package models

// ServiceHealthCheckType

import "encoding/json"

// ServiceHealthCheckType
type ServiceHealthCheckType struct {
	HTTPMethod      string                  `json:"http_method,omitempty"`
	TimeoutUsecs    int                     `json:"timeoutUsecs,omitempty"`
	Enabled         bool                    `json:"enabled,omitempty"`
	Delay           int                     `json:"delay,omitempty"`
	ExpectedCodes   string                  `json:"expected_codes,omitempty"`
	URLPath         string                  `json:"url_path,omitempty"`
	MonitorType     HealthCheckProtocolType `json:"monitor_type,omitempty"`
	DelayUsecs      int                     `json:"delayUsecs,omitempty"`
	MaxRetries      int                     `json:"max_retries,omitempty"`
	HealthCheckType HealthCheckType         `json:"health_check_type,omitempty"`
	Timeout         int                     `json:"timeout,omitempty"`
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
		MonitorType:     MakeHealthCheckProtocolType(),
		DelayUsecs:      0,
		MaxRetries:      0,
		HealthCheckType: MakeHealthCheckType(),
		ExpectedCodes:   "",
		HTTPMethod:      "",
		TimeoutUsecs:    0,
		Enabled:         false,
		Delay:           0,
	}
}

// MakeServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
	return []*ServiceHealthCheckType{}
}
