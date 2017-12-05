package models

// DiscoveryPubSubEndPointType

import "encoding/json"

type DiscoveryPubSubEndPointType struct {
	EpID      string      `json:"ep_id"`
	EpType    string      `json:"ep_type"`
	EpPrefix  *SubnetType `json:"ep_prefix"`
	EpVersion string      `json:"ep_version"`
}

func (model *DiscoveryPubSubEndPointType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeDiscoveryPubSubEndPointType() *DiscoveryPubSubEndPointType {
	return &DiscoveryPubSubEndPointType{
		//TODO(nati): Apply default
		EpVersion: "",
		EpID:      "",
		EpType:    "",
		EpPrefix:  MakeSubnetType(),
	}
}

func InterfaceToDiscoveryPubSubEndPointType(iData interface{}) *DiscoveryPubSubEndPointType {
	data := iData.(map[string]interface{})
	return &DiscoveryPubSubEndPointType{
		EpVersion: data["ep_version"].(string),

		//{"Title":"","Description":"All  servers or clients whose version match this version","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpVersion","GoType":"string"}
		EpID: data["ep_id"].(string),

		//{"Title":"","Description":"Specific service or client which is set of one.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpID","GoType":"string"}
		EpType: data["ep_type"].(string),

		//{"Title":"","Description":"Type of service or client","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpType","GoType":"string"}
		EpPrefix: InterfaceToSubnetType(data["ep_prefix"]),

		//{"Title":"","Description":"All  servers or clients whose ip match this prefix","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"EpPrefix","GoType":"SubnetType"}

	}
}

func InterfaceToDiscoveryPubSubEndPointTypeSlice(data interface{}) []*DiscoveryPubSubEndPointType {
	list := data.([]interface{})
	result := MakeDiscoveryPubSubEndPointTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDiscoveryPubSubEndPointType(item))
	}
	return result
}

func MakeDiscoveryPubSubEndPointTypeSlice() []*DiscoveryPubSubEndPointType {
	return []*DiscoveryPubSubEndPointType{}
}
