package models

// BridgeDomain

// BridgeDomain
//proteus:generate
type BridgeDomain struct {
	UUID               string                   `json:"uuid,omitempty"`
	ParentUUID         string                   `json:"parent_uuid,omitempty"`
	ParentType         string                   `json:"parent_type,omitempty"`
	FQName             []string                 `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType             `json:"id_perms,omitempty"`
	DisplayName        string                   `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs           `json:"annotations,omitempty"`
	Perms2             *PermType2               `json:"perms2,omitempty"`
	MacAgingTime       MACAgingTime             `json:"mac_aging_time,omitempty"`
	Isid               IsidType                 `json:"isid,omitempty"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled"`
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control,omitempty"`
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
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		MacAgingTime:       MakeMACAgingTime(),
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
