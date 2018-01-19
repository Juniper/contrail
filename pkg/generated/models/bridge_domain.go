package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control,omitempty"`
	FQName             []string                 `json:"fq_name,omitempty"`
	Annotations        *KeyValuePairs           `json:"annotations,omitempty"`
	Perms2             *PermType2               `json:"perms2,omitempty"`
	Isid               IsidType                 `json:"isid,omitempty"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled"`
	ParentUUID         string                   `json:"parent_uuid,omitempty"`
	ParentType         string                   `json:"parent_type,omitempty"`
	IDPerms            *IdPermsType             `json:"id_perms,omitempty"`
	DisplayName        string                   `json:"display_name,omitempty"`
	MacAgingTime       MACAgingTime             `json:"mac_aging_time,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
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
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		MacAgingTime:       MakeMACAgingTime(),
		UUID:               "",
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		Isid:               MakeIsidType(),
		MacLearningEnabled: false,
		MacMoveControl:     MakeMACMoveLimitControlType(),
		MacLimitControl:    MakeMACLimitControlType(),
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}
