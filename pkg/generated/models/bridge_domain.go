package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	ParentType         string                   `json:"parent_type,omitempty"`
	FQName             []string                 `json:"fq_name,omitempty"`
	MacAgingTime       MACAgingTime             `json:"mac_aging_time,omitempty"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control,omitempty"`
	Perms2             *PermType2               `json:"perms2,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
	ParentUUID         string                   `json:"parent_uuid,omitempty"`
	IDPerms            *IdPermsType             `json:"id_perms,omitempty"`
	DisplayName        string                   `json:"display_name,omitempty"`
	Isid               IsidType                 `json:"isid,omitempty"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled,omitempty"`
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
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
		Isid:               MakeIsidType(),
		MacLearningEnabled: false,
		MacMoveControl:     MakeMACMoveLimitControlType(),
		Annotations:        MakeKeyValuePairs(),
		ParentUUID:         "",
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		MacAgingTime:       MakeMACAgingTime(),
		MacLimitControl:    MakeMACLimitControlType(),
		Perms2:             MakePermType2(),
		UUID:               "",
		ParentType:         "",
		FQName:             []string{},
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}
