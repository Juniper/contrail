package models

// MACLimitControlType

import "encoding/json"

type MACLimitControlType struct {
	MacLimit       int                      `json:"mac_limit"`
	MacLimitAction MACLimitExceedActionType `json:"mac_limit_action"`
}

func (model *MACLimitControlType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeMACLimitControlType() *MACLimitControlType {
	return &MACLimitControlType{
		//TODO(nati): Apply default
		MacLimit:       0,
		MacLimitAction: MakeMACLimitExceedActionType(),
	}
}

func InterfaceToMACLimitControlType(iData interface{}) *MACLimitControlType {
	data := iData.(map[string]interface{})
	return &MACLimitControlType{
		MacLimit: data["mac_limit"].(int),

		//{"Title":"","Description":"Number of MACs that can be learnt","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MacLimit","GoType":"int"}
		MacLimitAction: InterfaceToMACLimitExceedActionType(data["mac_limit_action"]),

		//{"Title":"","Description":"Action to be taken when MAC limit exceeds","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["log","alarm","shutdown","drop"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitExceedActionType","CollectionType":"","Column":"","Item":null,"GoName":"MacLimitAction","GoType":"MACLimitExceedActionType"}

	}
}

func InterfaceToMACLimitControlTypeSlice(data interface{}) []*MACLimitControlType {
	list := data.([]interface{})
	result := MakeMACLimitControlTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACLimitControlType(item))
	}
	return result
}

func MakeMACLimitControlTypeSlice() []*MACLimitControlType {
	return []*MACLimitControlType{}
}
