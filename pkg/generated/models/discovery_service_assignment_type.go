package models

// DiscoveryServiceAssignmentType

import "encoding/json"

// DiscoveryServiceAssignmentType
type DiscoveryServiceAssignmentType struct {
	Subscriber []*DiscoveryPubSubEndPointType `json:"subscriber"`
	Publisher  *DiscoveryPubSubEndPointType   `json:"publisher"`
}

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

		Subscriber: InterfaceToDiscoveryPubSubEndPointTypeSlice(data["subscriber"]),

		//{"description":"subscriber set","type":"array","item":{"type":"object","properties":{"ep_id":{"type":"string"},"ep_prefix":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"ep_type":{"type":"string"},"ep_version":{"type":"string"}}}}
		Publisher: InterfaceToDiscoveryPubSubEndPointType(data["publisher"]),

		//{"description":"Publisher set","type":"object","properties":{"ep_id":{"type":"string"},"ep_prefix":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"ep_type":{"type":"string"},"ep_version":{"type":"string"}}}

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
