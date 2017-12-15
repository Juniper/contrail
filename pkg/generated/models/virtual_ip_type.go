package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	PersistenceCookieName string                   `json:"persistence_cookie_name"`
	ConnectionLimit       int                      `json:"connection_limit"`
	AdminState            bool                     `json:"admin_state"`
	ProtocolPort          int                      `json:"protocol_port"`
	Status                string                   `json:"status"`
	Protocol              LoadbalancerProtocolType `json:"protocol"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type"`
	Address               IpAddressType            `json:"address"`
	StatusDescription     string                   `json:"status_description"`
	SubnetID              UuidStringType           `json:"subnet_id"`
}

// String returns json representation of the object
func (model *VirtualIpType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualIpType makes VirtualIpType
func MakeVirtualIpType() *VirtualIpType {
	return &VirtualIpType{
		//TODO(nati): Apply default
		AdminState:            false,
		ProtocolPort:          0,
		Status:                "",
		Protocol:              MakeLoadbalancerProtocolType(),
		PersistenceCookieName: "",
		ConnectionLimit:       0,
		StatusDescription:     "",
		SubnetID:              MakeUuidStringType(),
		PersistenceType:       MakeSessionPersistenceType(),
		Address:               MakeIpAddressType(),
	}
}

// InterfaceToVirtualIpType makes VirtualIpType from interface
func InterfaceToVirtualIpType(iData interface{}) *VirtualIpType {
	data := iData.(map[string]interface{})
	return &VirtualIpType{
		Status: data["status"].(string),

		//{"description":"Operating status for this virtual ip.","type":"string"}
		Protocol: InterfaceToLoadbalancerProtocolType(data["protocol"]),

		//{"description":"IP protocol string like http, https or tcp.","type":"string","enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"]}
		PersistenceCookieName: data["persistence_cookie_name"].(string),

		//{"description":"Set this string if the relation of client and server(pool member) need to persist.","type":"string"}
		ConnectionLimit: data["connection_limit"].(int),

		//{"description":"Maximum number of concurrent connections","type":"integer"}
		AdminState: data["admin_state"].(bool),

		//{"description":"Administrative up or down.","type":"boolean"}
		ProtocolPort: data["protocol_port"].(int),

		//{"description":"Layer 4 protocol destination port.","type":"integer"}
		StatusDescription: data["status_description"].(string),

		//{"description":"Operating status description this virtual ip.","type":"string"}
		SubnetID: InterfaceToUuidStringType(data["subnet_id"]),

		//{"description":"UUID of subnet in which to allocate the Virtual IP.","type":"string"}
		PersistenceType: InterfaceToSessionPersistenceType(data["persistence_type"]),

		//{"description":"Method for persistence. HTTP_COOKIE, SOURCE_IP or APP_COOKIE.","type":"string","enum":["SOURCE_IP","HTTP_COOKIE","APP_COOKIE"]}
		Address: InterfaceToIpAddressType(data["address"]),

		//{"description":"IP address automatically allocated by system.","type":"string"}

	}
}

// InterfaceToVirtualIpTypeSlice makes a slice of VirtualIpType from interface
func InterfaceToVirtualIpTypeSlice(data interface{}) []*VirtualIpType {
	list := data.([]interface{})
	result := MakeVirtualIpTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualIpType(item))
	}
	return result
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
