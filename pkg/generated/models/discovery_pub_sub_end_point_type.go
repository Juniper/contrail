package models

// DiscoveryPubSubEndPointType

import "encoding/json"

// DiscoveryPubSubEndPointType
type DiscoveryPubSubEndPointType struct {
	EpType    string      `json:"ep_type"`
	EpPrefix  *SubnetType `json:"ep_prefix"`
	EpVersion string      `json:"ep_version"`
	EpID      string      `json:"ep_id"`
}

//  parents relation object

// String returns json representation of the object
func (model *DiscoveryPubSubEndPointType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointType() *DiscoveryPubSubEndPointType {
	return &DiscoveryPubSubEndPointType{
		//TODO(nati): Apply default
		EpID:      "",
		EpType:    "",
		EpPrefix:  MakeSubnetType(),
		EpVersion: "",
	}
}

// InterfaceToDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType from interface
func InterfaceToDiscoveryPubSubEndPointType(iData interface{}) *DiscoveryPubSubEndPointType {
	data := iData.(map[string]interface{})
	return &DiscoveryPubSubEndPointType{
		EpVersion: data["ep_version"].(string),

		//{"Title":"","Description":"All  servers or clients whose version match this version","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpVersion","GoType":"string","GoPremitive":true}
		EpID: data["ep_id"].(string),

		//{"Title":"","Description":"Specific service or client which is set of one.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpID","GoType":"string","GoPremitive":true}
		EpType: data["ep_type"].(string),

		//{"Title":"","Description":"Type of service or client","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpType","GoType":"string","GoPremitive":true}
		EpPrefix: InterfaceToSubnetType(data["ep_prefix"]),

		//{"Title":"","Description":"All  servers or clients whose ip match this prefix","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"EpPrefix","GoType":"SubnetType","GoPremitive":false}

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
