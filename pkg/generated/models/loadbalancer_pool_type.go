package models

// LoadbalancerPoolType

import "encoding/json"

// LoadbalancerPoolType
type LoadbalancerPoolType struct {
	Status                string                   `json:"status"`
	Protocol              LoadbalancerProtocolType `json:"protocol"`
	SubnetID              UuidStringType           `json:"subnet_id"`
	SessionPersistence    SessionPersistenceType   `json:"session_persistence"`
	AdminState            bool                     `json:"admin_state"`
	PersistenceCookieName string                   `json:"persistence_cookie_name"`
	StatusDescription     string                   `json:"status_description"`
	LoadbalancerMethod    LoadbalancerMethodType   `json:"loadbalancer_method"`
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
		PersistenceCookieName: "",
		StatusDescription:     "",
		LoadbalancerMethod:    MakeLoadbalancerMethodType(),
		Status:                "",
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
		SessionPersistence:    MakeSessionPersistenceType(),
		AdminState:            false,
	}
}

// InterfaceToLoadbalancerPoolType makes LoadbalancerPoolType from interface
func InterfaceToLoadbalancerPoolType(iData interface{}) *LoadbalancerPoolType {
	data := iData.(map[string]interface{})
	return &LoadbalancerPoolType{
		SubnetID: InterfaceToUuidStringType(data["subnet_id"]),

		//{"description":"UUID of the subnet from where the members of the pool are reachable.","type":"string"}
		SessionPersistence: InterfaceToSessionPersistenceType(data["session_persistence"]),

		//{"description":"Method for persistence. HTTP_COOKIE, SOURCE_IP or APP_COOKIE.","type":"string","enum":["SOURCE_IP","HTTP_COOKIE","APP_COOKIE"]}
		AdminState: data["admin_state"].(bool),

		//{"description":"Administrative up or down","type":"boolean"}
		PersistenceCookieName: data["persistence_cookie_name"].(string),

		//{"description":"To Be Added","type":"string"}
		StatusDescription: data["status_description"].(string),

		//{"description":"Operating status description for this loadbalancer pool.","type":"string"}
		LoadbalancerMethod: InterfaceToLoadbalancerMethodType(data["loadbalancer_method"]),

		//{"description":"Load balancing method ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP","type":"string","enum":["ROUND_ROBIN","LEAST_CONNECTIONS","SOURCE_IP"]}
		Status: data["status"].(string),

		//{"description":"Operating status for this loadbalancer pool.","type":"string"}
		Protocol: InterfaceToLoadbalancerProtocolType(data["protocol"]),

		//{"description":"IP protocol string like http, https or tcp.","type":"string","enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"]}

	}
}

// InterfaceToLoadbalancerPoolTypeSlice makes a slice of LoadbalancerPoolType from interface
func InterfaceToLoadbalancerPoolTypeSlice(data interface{}) []*LoadbalancerPoolType {
	list := data.([]interface{})
	result := MakeLoadbalancerPoolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerPoolType(item))
	}
	return result
}

// MakeLoadbalancerPoolTypeSlice() makes a slice of LoadbalancerPoolType
func MakeLoadbalancerPoolTypeSlice() []*LoadbalancerPoolType {
	return []*LoadbalancerPoolType{}
}
