package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
	ParentUUID         string                   `json:"parent_uuid,omitempty"`
	ParentType         string                   `json:"parent_type,omitempty"`
	FQName             []string                 `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType             `json:"id_perms,omitempty"`
	Isid               IsidType                 `json:"isid,omitempty"`
	Perms2             *PermType2               `json:"perms2,omitempty"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled"`
	DisplayName        string                   `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs           `json:"annotations,omitempty"`
	MacAgingTime       MACAgingTime             `json:"mac_aging_time,omitempty"`
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
		UUID:               "",
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		Isid:               MakeIsidType(),
		MacMoveControl:     MakeMACMoveLimitControlType(),
		MacLimitControl:    MakeMACLimitControlType(),
		Perms2:             MakePermType2(),
		Annotations:        MakeKeyValuePairs(),
		MacAgingTime:       MakeMACAgingTime(),
		MacLearningEnabled: false,
		DisplayName:        "",
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}
