package models

// DiscoveryServiceAssignmentType

import "encoding/json"

type DiscoveryServiceAssignmentType struct {
	Subscriber []*DiscoveryPubSubEndPointType `json:"subscriber"`
	Publisher  *DiscoveryPubSubEndPointType   `json:"publisher"`
}

func (model *DiscoveryServiceAssignmentType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeDiscoveryServiceAssignmentType() *DiscoveryServiceAssignmentType {
	return &DiscoveryServiceAssignmentType{
		//TODO(nati): Apply default
		Publisher: MakeDiscoveryPubSubEndPointType(),

		Subscriber: MakeDiscoveryPubSubEndPointTypeSlice(),
	}
}

func InterfaceToDiscoveryServiceAssignmentType(iData interface{}) *DiscoveryServiceAssignmentType {
	data := iData.(map[string]interface{})
	return &DiscoveryServiceAssignmentType{

		Subscriber: InterfaceToDiscoveryPubSubEndPointTypeSlice(data["subscriber"]),

		//{"Title":"","Description":"subscriber set","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ep_id":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpID","GoType":"string"},"ep_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"EpPrefix","GoType":"SubnetType"},"ep_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpType","GoType":"string"},"ep_version":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpVersion","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DiscoveryPubSubEndPointType","CollectionType":"","Column":"","Item":null,"GoName":"Subscriber","GoType":"DiscoveryPubSubEndPointType"},"GoName":"Subscriber","GoType":"[]*DiscoveryPubSubEndPointType"}
		Publisher: InterfaceToDiscoveryPubSubEndPointType(data["publisher"]),

		//{"Title":"","Description":"Publisher set","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"ep_id":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpID","GoType":"string"},"ep_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"EpPrefix","GoType":"SubnetType"},"ep_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpType","GoType":"string"},"ep_version":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpVersion","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DiscoveryPubSubEndPointType","CollectionType":"","Column":"","Item":null,"GoName":"Publisher","GoType":"DiscoveryPubSubEndPointType"}

	}
}

func InterfaceToDiscoveryServiceAssignmentTypeSlice(data interface{}) []*DiscoveryServiceAssignmentType {
	list := data.([]interface{})
	result := MakeDiscoveryServiceAssignmentTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDiscoveryServiceAssignmentType(item))
	}
	return result
}

func MakeDiscoveryServiceAssignmentTypeSlice() []*DiscoveryServiceAssignmentType {
	return []*DiscoveryServiceAssignmentType{}
}
