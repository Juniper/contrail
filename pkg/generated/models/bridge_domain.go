package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	Isid               IsidType                 `json:"isid,omitempty"`
	Perms2             *PermType2               `json:"perms2,omitempty"`
	ParentUUID         string                   `json:"parent_uuid,omitempty"`
	ParentType         string                   `json:"parent_type,omitempty"`
	FQName             []string                 `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType             `json:"id_perms,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
	MacAgingTime       MACAgingTime             `json:"mac_aging_time,omitempty"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled"`
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control,omitempty"`
	DisplayName        string                   `json:"display_name,omitempty"`
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
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		Isid:               MakeIsidType(),
		Perms2:             MakePermType2(),
		ParentUUID:         "",
		MacLimitControl:    MakeMACLimitControlType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		UUID:               "",
		MacAgingTime:       MakeMACAgingTime(),
		MacLearningEnabled: false,
		MacMoveControl:     MakeMACMoveLimitControlType(),
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}
