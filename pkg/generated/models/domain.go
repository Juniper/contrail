package models

// Domain

import "encoding/json"

// Domain
type Domain struct {
	Perms2       *PermType2        `json:"perms2"`
	UUID         string            `json:"uuid"`
	ParentUUID   string            `json:"parent_uuid"`
	IDPerms      *IdPermsType      `json:"id_perms"`
	DisplayName  string            `json:"display_name"`
	Annotations  *KeyValuePairs    `json:"annotations"`
	ParentType   string            `json:"parent_type"`
	FQName       []string          `json:"fq_name"`
	DomainLimits *DomainLimitsType `json:"domain_limits"`

	APIAccessLists   []*APIAccessList   `json:"api_access_lists"`
	Namespaces       []*Namespace       `json:"namespaces"`
	Projects         []*Project         `json:"projects"`
	ServiceTemplates []*ServiceTemplate `json:"service_templates"`
	VirtualDNSs      []*VirtualDNS      `json:"virtual_DNSs"`
}

// String returns json representation of the object
func (model *Domain) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDomain makes Domain
func MakeDomain() *Domain {
	return &Domain{
		//TODO(nati): Apply default
		DomainLimits: MakeDomainLimitsType(),
		ParentType:   "",
		FQName:       []string{},
		Annotations:  MakeKeyValuePairs(),
		Perms2:       MakePermType2(),
		UUID:         "",
		ParentUUID:   "",
		IDPerms:      MakeIdPermsType(),
		DisplayName:  "",
	}
}

// MakeDomainSlice() makes a slice of Domain
func MakeDomainSlice() []*Domain {
	return []*Domain{}
}
