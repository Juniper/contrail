package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control"`
	FQName             []string                 `json:"fq_name"`
	IDPerms            *IdPermsType             `json:"id_perms"`
	DisplayName        string                   `json:"display_name"`
	MacAgingTime       MACAgingTime             `json:"mac_aging_time"`
	Isid               IsidType                 `json:"isid"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled"`
	Annotations        *KeyValuePairs           `json:"annotations"`
	Perms2             *PermType2               `json:"perms2"`
	ParentType         string                   `json:"parent_type"`
	UUID               string                   `json:"uuid"`
	ParentUUID         string                   `json:"parent_uuid"`
}

// String returns json representation of the object
func (model *BridgeDomain) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBridgeDomain makes BridgeDomain
func MakeBridgeDomain() *BridgeDomain {
	return &BridgeDomain{
		//TODO(nati): Apply default
		ParentType:         "",
		UUID:               "",
		ParentUUID:         "",
		DisplayName:        "",
		MacAgingTime:       MakeMACAgingTime(),
		Isid:               MakeIsidType(),
		MacLearningEnabled: false,
		MacMoveControl:     MakeMACMoveLimitControlType(),
		MacLimitControl:    MakeMACLimitControlType(),
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}
