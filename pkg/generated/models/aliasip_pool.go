package models

// AliasIPPool

import "encoding/json"

// AliasIPPool
type AliasIPPool struct {
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`

	AliasIPs []*AliasIP `json:"alias_ips"`
}

// String returns json representation of the object
func (model *AliasIPPool) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAliasIPPool makes AliasIPPool
func MakeAliasIPPool() *AliasIPPool {
	return &AliasIPPool{
		//TODO(nati): Apply default
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
	}
}

// MakeAliasIPPoolSlice() makes a slice of AliasIPPool
func MakeAliasIPPoolSlice() []*AliasIPPool {
	return []*AliasIPPool{}
}
