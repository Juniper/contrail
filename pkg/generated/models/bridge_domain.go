package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	MacLearningEnabled bool                     `json:"mac_learning_enabled"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control,omitempty"`
	ParentUUID         string                   `json:"parent_uuid,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
	DisplayName        string                   `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs           `json:"annotations,omitempty"`
	MacAgingTime       MACAgingTime             `json:"mac_aging_time,omitempty"`
	Isid               IsidType                 `json:"isid,omitempty"`
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
	ParentType         string                   `json:"parent_type,omitempty"`
	FQName             []string                 `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType             `json:"id_perms,omitempty"`
	Perms2             *PermType2               `json:"perms2,omitempty"`
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
		MacLearningEnabled: false,
		MacLimitControl:    MakeMACLimitControlType(),
		ParentUUID:         "",
		UUID:               "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		MacAgingTime:       MakeMACAgingTime(),
		Isid:               MakeIsidType(),
		MacMoveControl:     MakeMACMoveLimitControlType(),
		ParentType:         "",
		Perms2:             MakePermType2(),
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}
