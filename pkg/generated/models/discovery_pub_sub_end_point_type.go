package models

// DiscoveryPubSubEndPointType

import "encoding/json"

// DiscoveryPubSubEndPointType
type DiscoveryPubSubEndPointType struct {
	EpPrefix  *SubnetType `json:"ep_prefix"`
	EpVersion string      `json:"ep_version"`
	EpID      string      `json:"ep_id"`
	EpType    string      `json:"ep_type"`
}

// String returns json representation of the object
func (model *DiscoveryPubSubEndPointType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointType() *DiscoveryPubSubEndPointType {
	return &DiscoveryPubSubEndPointType{
		//TODO(nati): Apply default
		EpPrefix:  MakeSubnetType(),
		EpVersion: "",
		EpID:      "",
		EpType:    "",
	}
}

// InterfaceToDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType from interface
func InterfaceToDiscoveryPubSubEndPointType(iData interface{}) *DiscoveryPubSubEndPointType {
	data := iData.(map[string]interface{})
	return &DiscoveryPubSubEndPointType{
		EpVersion: data["ep_version"].(string),

		//{"description":"All  servers or clients whose version match this version","type":"string"}
		EpID: data["ep_id"].(string),

		//{"description":"Specific service or client which is set of one.","type":"string"}
		EpType: data["ep_type"].(string),

		//{"description":"Type of service or client","type":"string"}
		EpPrefix: InterfaceToSubnetType(data["ep_prefix"]),

		//{"description":"All  servers or clients whose ip match this prefix","type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}

	}
}

// InterfaceToDiscoveryPubSubEndPointTypeSlice makes a slice of DiscoveryPubSubEndPointType from interface
func InterfaceToDiscoveryPubSubEndPointTypeSlice(data interface{}) []*DiscoveryPubSubEndPointType {
	list := data.([]interface{})
	result := MakeDiscoveryPubSubEndPointTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDiscoveryPubSubEndPointType(item))
	}
	return result
}

// MakeDiscoveryPubSubEndPointTypeSlice() makes a slice of DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointTypeSlice() []*DiscoveryPubSubEndPointType {
	return []*DiscoveryPubSubEndPointType{}
}
