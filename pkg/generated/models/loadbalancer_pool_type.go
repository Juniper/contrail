package models

// LoadbalancerPoolType

import "encoding/json"

// LoadbalancerPoolType
type LoadbalancerPoolType struct {
	PersistenceCookieName string                   `json:"persistence_cookie_name"`
	StatusDescription     string                   `json:"status_description"`
	LoadbalancerMethod    LoadbalancerMethodType   `json:"loadbalancer_method"`
	Status                string                   `json:"status"`
	Protocol              LoadbalancerProtocolType `json:"protocol"`
	SubnetID              UuidStringType           `json:"subnet_id"`
	SessionPersistence    SessionPersistenceType   `json:"session_persistence"`
	AdminState            bool                     `json:"admin_state"`
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
		AdminState:            false,
		PersistenceCookieName: "",
		StatusDescription:     "",
		LoadbalancerMethod:    MakeLoadbalancerMethodType(),
		Status:                "",
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
		SessionPersistence:    MakeSessionPersistenceType(),
	}
}

// MakeLoadbalancerPoolTypeSlice() makes a slice of LoadbalancerPoolType
func MakeLoadbalancerPoolTypeSlice() []*LoadbalancerPoolType {
	return []*LoadbalancerPoolType{}
}
