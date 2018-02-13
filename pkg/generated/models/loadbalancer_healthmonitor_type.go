package models

// LoadbalancerHealthmonitorType

import "encoding/json"

// LoadbalancerHealthmonitorType
//proteus:generate
type LoadbalancerHealthmonitorType struct {
	Delay         int               `json:"delay,omitempty"`
	ExpectedCodes string            `json:"expected_codes,omitempty"`
	MaxRetries    int               `json:"max_retries,omitempty"`
	HTTPMethod    string            `json:"http_method,omitempty"`
	AdminState    bool              `json:"admin_state"`
	Timeout       int               `json:"timeout,omitempty"`
	URLPath       string            `json:"url_path,omitempty"`
	MonitorType   HealthmonitorType `json:"monitor_type,omitempty"`
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

// MakeLoadbalancerHealthmonitorTypeSlice() makes a slice of LoadbalancerHealthmonitorType
func MakeLoadbalancerHealthmonitorTypeSlice() []*LoadbalancerHealthmonitorType {
	return []*LoadbalancerHealthmonitorType{}
}
