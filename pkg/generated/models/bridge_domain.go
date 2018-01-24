package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	MacAgingTime       MACAgingTime             `json:"mac_aging_time,omitempty"`
	Isid               IsidType                 `json:"isid,omitempty"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control,omitempty"`
	ParentType         string                   `json:"parent_type,omitempty"`
	DisplayName        string                   `json:"display_name,omitempty"`
	Perms2             *PermType2               `json:"perms2,omitempty"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled"`
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
	ParentUUID         string                   `json:"parent_uuid,omitempty"`
	FQName             []string                 `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType             `json:"id_perms,omitempty"`
	Annotations        *KeyValuePairs           `json:"annotations,omitempty"`
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
		DisplayName:        "",
		Perms2:             MakePermType2(),
		MacAgingTime:       MakeMACAgingTime(),
		Isid:               MakeIsidType(),
		MacLimitControl:    MakeMACLimitControlType(),
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		Annotations:        MakeKeyValuePairs(),
		MacLearningEnabled: false,
		MacMoveControl:     MakeMACMoveLimitControlType(),
		UUID:               "",
		ParentUUID:         "",
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}
