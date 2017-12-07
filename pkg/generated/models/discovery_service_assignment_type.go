package models

// DiscoveryServiceAssignmentType

import "encoding/json"

// DiscoveryServiceAssignmentType
type DiscoveryServiceAssignmentType struct {
	Subscriber []*DiscoveryPubSubEndPointType `json:"subscriber"`
	Publisher  *DiscoveryPubSubEndPointType   `json:"publisher"`
}

//  parents relation object

// String returns json representation of the object
func (model *DiscoveryServiceAssignmentType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDiscoveryServiceAssignmentType makes DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentType() *DiscoveryServiceAssignmentType {
	return &DiscoveryServiceAssignmentType{
		//TODO(nati): Apply default

		Subscriber: MakeDiscoveryPubSubEndPointTypeSlice(),

		Publisher: MakeDiscoveryPubSubEndPointType(),
	}
}

// InterfaceToDiscoveryServiceAssignmentType makes DiscoveryServiceAssignmentType from interface
func InterfaceToDiscoveryServiceAssignmentType(iData interface{}) *DiscoveryServiceAssignmentType {
	data := iData.(map[string]interface{})
	return &DiscoveryServiceAssignmentType{
		Publisher: InterfaceToDiscoveryPubSubEndPointType(data["publisher"]),

		//{"Title":"","Description":"Publisher set","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"ep_id":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpID","GoType":"string","GoPremitive":true},"ep_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"EpPrefix","GoType":"SubnetType","GoPremitive":false},"ep_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpType","GoType":"string","GoPremitive":true},"ep_version":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpVersion","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DiscoveryPubSubEndPointType","CollectionType":"","Column":"","Item":null,"GoName":"Publisher","GoType":"DiscoveryPubSubEndPointType","GoPremitive":false}

		Subscriber: InterfaceToDiscoveryPubSubEndPointTypeSlice(data["subscriber"]),

		//{"Title":"","Description":"subscriber set","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ep_id":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpID","GoType":"string","GoPremitive":true},"ep_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"EpPrefix","GoType":"SubnetType","GoPremitive":false},"ep_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpType","GoType":"string","GoPremitive":true},"ep_version":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EpVersion","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DiscoveryPubSubEndPointType","CollectionType":"","Column":"","Item":null,"GoName":"Subscriber","GoType":"DiscoveryPubSubEndPointType","GoPremitive":false},"GoName":"Subscriber","GoType":"[]*DiscoveryPubSubEndPointType","GoPremitive":true}

	}
}

// InterfaceToDiscoveryServiceAssignmentTypeSlice makes a slice of DiscoveryServiceAssignmentType from interface
func InterfaceToDiscoveryServiceAssignmentTypeSlice(data interface{}) []*DiscoveryServiceAssignmentType {
	list := data.([]interface{})
	result := MakeDiscoveryServiceAssignmentTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDiscoveryServiceAssignmentType(item))
	}
	return result
}

// MakeDiscoveryServiceAssignmentTypeSlice() makes a slice of DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentTypeSlice() []*DiscoveryServiceAssignmentType {
	return []*DiscoveryServiceAssignmentType{}
}
