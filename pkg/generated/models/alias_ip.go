package models

// AliasIP

import "encoding/json"

// AliasIP
type AliasIP struct {
	AliasIPAddress       IpAddressType       `json:"alias_ip_address,omitempty"`
	AliasIPAddressFamily IpAddressFamilyType `json:"alias_ip_address_family,omitempty"`
	DisplayName          string              `json:"display_name,omitempty"`
	Perms2               *PermType2          `json:"perms2,omitempty"`
	ParentUUID           string              `json:"parent_uuid,omitempty"`
	Annotations          *KeyValuePairs      `json:"annotations,omitempty"`
	UUID                 string              `json:"uuid,omitempty"`
	ParentType           string              `json:"parent_type,omitempty"`
	FQName               []string            `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType        `json:"id_perms,omitempty"`

	ProjectRefs                 []*AliasIPProjectRef                 `json:"project_refs,omitempty"`
	VirtualMachineInterfaceRefs []*AliasIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
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
		IDPerms:              MakeIdPermsType(),
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		ParentType:           "",
		FQName:               []string{},
		ParentUUID:           "",
		AliasIPAddress:       MakeIpAddressType(),
		AliasIPAddressFamily: MakeIpAddressFamilyType(),
		DisplayName:          "",
		Perms2:               MakePermType2(),
	}
}

// MakeAliasIPSlice() makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
	return []*AliasIP{}
}
