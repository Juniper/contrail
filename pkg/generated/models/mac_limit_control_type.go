package models

// MACLimitControlType

import "encoding/json"

// MACLimitControlType
type MACLimitControlType struct {
	MacLimit       int                      `json:"mac_limit"`
	MacLimitAction MACLimitExceedActionType `json:"mac_limit_action"`
}

// String returns json representation of the object
func (model *MACLimitControlType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMACLimitControlType makes MACLimitControlType
func MakeMACLimitControlType() *MACLimitControlType {
	return &MACLimitControlType{
		//TODO(nati): Apply default
		MacLimit:       0,
		MacLimitAction: MakeMACLimitExceedActionType(),
	}
}

// InterfaceToMACLimitControlType makes MACLimitControlType from interface
func InterfaceToMACLimitControlType(iData interface{}) *MACLimitControlType {
	data := iData.(map[string]interface{})
	return &MACLimitControlType{
		MacLimit: data["mac_limit"].(int),

		//{"description":"Number of MACs that can be learnt","type":"integer"}
		MacLimitAction: InterfaceToMACLimitExceedActionType(data["mac_limit_action"]),

		//{"description":"Action to be taken when MAC limit exceeds","type":"string","enum":["log","alarm","shutdown","drop"]}

	}
}

// InterfaceToMACLimitControlTypeSlice makes a slice of MACLimitControlType from interface
func InterfaceToMACLimitControlTypeSlice(data interface{}) []*MACLimitControlType {
	list := data.([]interface{})
	result := MakeMACLimitControlTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACLimitControlType(item))
	}
	return result
}

// MakeMACLimitControlTypeSlice() makes a slice of MACLimitControlType
func MakeMACLimitControlTypeSlice() []*MACLimitControlType {
	return []*MACLimitControlType{}
}
