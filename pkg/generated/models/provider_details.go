package models

// ProviderDetails

import "encoding/json"

// ProviderDetails
type ProviderDetails struct {
	SegmentationID  VlanIdType `json:"segmentation_id"`
	PhysicalNetwork string     `json:"physical_network"`
}

// String returns json representation of the object
func (model *ProviderDetails) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeProviderDetails makes ProviderDetails
func MakeProviderDetails() *ProviderDetails {
	return &ProviderDetails{
		//TODO(nati): Apply default
		SegmentationID:  MakeVlanIdType(),
		PhysicalNetwork: "",
	}
}

// InterfaceToProviderDetails makes ProviderDetails from interface
func InterfaceToProviderDetails(iData interface{}) *ProviderDetails {
	data := iData.(map[string]interface{})
	return &ProviderDetails{
		SegmentationID: InterfaceToVlanIdType(data["segmentation_id"]),

		//{"type":"integer","minimum":1,"maximum":4094}
		PhysicalNetwork: data["physical_network"].(string),

		//{"type":"string"}

	}
}

// InterfaceToProviderDetailsSlice makes a slice of ProviderDetails from interface
func InterfaceToProviderDetailsSlice(data interface{}) []*ProviderDetails {
	list := data.([]interface{})
	result := MakeProviderDetailsSlice()
	for _, item := range list {
		result = append(result, InterfaceToProviderDetails(item))
	}
	return result
}

// MakeProviderDetailsSlice() makes a slice of ProviderDetails
func MakeProviderDetailsSlice() []*ProviderDetails {
	return []*ProviderDetails{}
}
