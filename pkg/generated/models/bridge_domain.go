package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	MacAgingTime       MACAgingTime             `json:"mac_aging_time,omitempty"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled,omitempty"`
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control,omitempty"`
	DisplayName        string                   `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs           `json:"annotations,omitempty"`
	Perms2             *PermType2               `json:"perms2,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
	ParentType         string                   `json:"parent_type,omitempty"`
	FQName             []string                 `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType             `json:"id_perms,omitempty"`
	Isid               IsidType                 `json:"isid,omitempty"`
	ParentUUID         string                   `json:"parent_uuid,omitempty"`
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
		Isid:               MakeIsidType(),
		ParentUUID:         "",
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		UUID:               "",
		MacAgingTime:       MakeMACAgingTime(),
		MacLearningEnabled: false,
		MacMoveControl:     MakeMACMoveLimitControlType(),
		MacLimitControl:    MakeMACLimitControlType(),
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}
