package models

// AliasIPPool

import "encoding/json"

// AliasIPPool
type AliasIPPool struct {
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`

	AliasIPs []*AliasIP `json:"alias_ips,omitempty"`
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
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeAliasIPPoolSlice() makes a slice of AliasIPPool
func MakeAliasIPPoolSlice() []*AliasIPPool {
	return []*AliasIPPool{}
}
