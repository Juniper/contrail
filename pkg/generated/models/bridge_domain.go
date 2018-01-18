package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
	ParentUUID         string                   `json:"parent_uuid,omitempty"`
	ParentType         string                   `json:"parent_type,omitempty"`
	Annotations        *KeyValuePairs           `json:"annotations,omitempty"`
	MacAgingTime       MACAgingTime             `json:"mac_aging_time,omitempty"`
	Isid               IsidType                 `json:"isid,omitempty"`
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
	Perms2             *PermType2               `json:"perms2,omitempty"`
	DisplayName        string                   `json:"display_name,omitempty"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled"`
	FQName             []string                 `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType             `json:"id_perms,omitempty"`
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
		MacMoveControl:     MakeMACMoveLimitControlType(),
		MacLimitControl:    MakeMACLimitControlType(),
		UUID:               "",
		ParentUUID:         "",
		ParentType:         "",
		Annotations:        MakeKeyValuePairs(),
		MacAgingTime:       MakeMACAgingTime(),
		Isid:               MakeIsidType(),
		Perms2:             MakePermType2(),
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		MacLearningEnabled: false,
		FQName:             []string{},
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}
