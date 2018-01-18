package models

// AliasIP

import "encoding/json"

// AliasIP
type AliasIP struct {
	UUID                 string              `json:"uuid,omitempty"`
	ParentType           string              `json:"parent_type,omitempty"`
	Annotations          *KeyValuePairs      `json:"annotations,omitempty"`
	AliasIPAddress       IpAddressType       `json:"alias_ip_address,omitempty"`
	AliasIPAddressFamily IpAddressFamilyType `json:"alias_ip_address_family,omitempty"`
	Perms2               *PermType2          `json:"perms2,omitempty"`
	ParentUUID           string              `json:"parent_uuid,omitempty"`
	FQName               []string            `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType        `json:"id_perms,omitempty"`
	DisplayName          string              `json:"display_name,omitempty"`

	ProjectRefs                 []*AliasIPProjectRef                 `json:"project_refs,omitempty"`
	VirtualMachineInterfaceRefs []*AliasIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
}

// AliasIPProjectRef references each other
type AliasIPProjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// AliasIPVirtualMachineInterfaceRef references each other
type AliasIPVirtualMachineInterfaceRef struct {
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
		AliasIPAddressFamily: MakeIpAddressFamilyType(),
		Perms2:               MakePermType2(),
		ParentUUID:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		AliasIPAddress:       MakeIpAddressType(),
		ParentType:           "",
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
	}
}

// MakeAliasIPSlice() makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
	return []*AliasIP{}
}
