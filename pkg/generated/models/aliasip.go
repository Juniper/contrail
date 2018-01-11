package models

// AliasIP

import "encoding/json"

// AliasIP
type AliasIP struct {
	AliasIPAddressFamily IpAddressFamilyType `json:"alias_ip_address_family"`
	ParentUUID           string              `json:"parent_uuid"`
	ParentType           string              `json:"parent_type"`
	FQName               []string            `json:"fq_name"`
	DisplayName          string              `json:"display_name"`
	AliasIPAddress       IpAddressType       `json:"alias_ip_address"`
	Perms2               *PermType2          `json:"perms2"`
	UUID                 string              `json:"uuid"`
	IDPerms              *IdPermsType        `json:"id_perms"`
	Annotations          *KeyValuePairs      `json:"annotations"`

	ProjectRefs                 []*AliasIPProjectRef                 `json:"project_refs"`
	VirtualMachineInterfaceRefs []*AliasIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
}

// AliasIPVirtualMachineInterfaceRef references each other
type AliasIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// AliasIPProjectRef references each other
type AliasIPProjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *AliasIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAliasIP makes AliasIP
func MakeAliasIP() *AliasIP {
	return &AliasIP{
		//TODO(nati): Apply default
		Annotations:          MakeKeyValuePairs(),
		AliasIPAddress:       MakeIpAddressType(),
		Perms2:               MakePermType2(),
		UUID:                 "",
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		AliasIPAddressFamily: MakeIpAddressFamilyType(),
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
	}
}

// MakeAliasIPSlice() makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
	return []*AliasIP{}
}
