package models

// LoadbalancerPoolType

// LoadbalancerPoolType
//proteus:generate
type LoadbalancerPoolType struct {
	Status                string                   `json:"status,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
	SessionPersistence    SessionPersistenceType   `json:"session_persistence,omitempty"`
	AdminState            bool                     `json:"admin_state"`
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
	LoadbalancerMethod    LoadbalancerMethodType   `json:"loadbalancer_method,omitempty"`
}

// MakeLoadbalancerPoolType makes LoadbalancerPoolType
func MakeLoadbalancerPoolType() *LoadbalancerPoolType {
	return &LoadbalancerPoolType{
		//TODO(nati): Apply default
		Status:                "",
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
		SessionPersistence:    MakeSessionPersistenceType(),
		AdminState:            false,
		PersistenceCookieName: "",
		StatusDescription:     "",
		LoadbalancerMethod:    MakeLoadbalancerMethodType(),
	}
}

// MakeLoadbalancerPoolTypeSlice() makes a slice of LoadbalancerPoolType
func MakeLoadbalancerPoolTypeSlice() []*LoadbalancerPoolType {
	return []*LoadbalancerPoolType{}
}
