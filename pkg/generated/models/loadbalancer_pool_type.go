package models

// LoadbalancerPoolType

import "encoding/json"

// LoadbalancerPoolType
type LoadbalancerPoolType struct {
	AdminState            bool                     `json:"admin_state,omitempty"`
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
	LoadbalancerMethod    LoadbalancerMethodType   `json:"loadbalancer_method,omitempty"`
	Status                string                   `json:"status,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
	SessionPersistence    SessionPersistenceType   `json:"session_persistence,omitempty"`
}

// String returns json representation of the object
func (model *LoadbalancerPoolType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerPoolType makes LoadbalancerPoolType
func MakeLoadbalancerPoolType() *LoadbalancerPoolType {
	return &LoadbalancerPoolType{
		//TODO(nati): Apply default
		SessionPersistence:    MakeSessionPersistenceType(),
		AdminState:            false,
		PersistenceCookieName: "",
		StatusDescription:     "",
		LoadbalancerMethod:    MakeLoadbalancerMethodType(),
		Status:                "",
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
	}
}

// MakeLoadbalancerPoolTypeSlice() makes a slice of LoadbalancerPoolType
func MakeLoadbalancerPoolTypeSlice() []*LoadbalancerPoolType {
	return []*LoadbalancerPoolType{}
}
