package models

// LoadbalancerHealthmonitorType

import "encoding/json"

// LoadbalancerHealthmonitorType
type LoadbalancerHealthmonitorType struct {
	AdminState    bool              `json:"admin_state"`
	Timeout       int               `json:"timeout"`
	URLPath       string            `json:"url_path"`
	MonitorType   HealthmonitorType `json:"monitor_type"`
	Delay         int               `json:"delay"`
	ExpectedCodes string            `json:"expected_codes"`
	MaxRetries    int               `json:"max_retries"`
	HTTPMethod    string            `json:"http_method"`
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
		Timeout:       0,
		URLPath:       "",
		MonitorType:   MakeHealthmonitorType(),
		Delay:         0,
		ExpectedCodes: "",
		MaxRetries:    0,
		HTTPMethod:    "",
		AdminState:    false,
	}
}

// MakeLoadbalancerHealthmonitorTypeSlice() makes a slice of LoadbalancerHealthmonitorType
func MakeLoadbalancerHealthmonitorTypeSlice() []*LoadbalancerHealthmonitorType {
	return []*LoadbalancerHealthmonitorType{}
}
