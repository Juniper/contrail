package models

// Domain

import "encoding/json"

// Domain
type Domain struct {
	Annotations  *KeyValuePairs    `json:"annotations,omitempty"`
	FQName       []string          `json:"fq_name,omitempty"`
	IDPerms      *IdPermsType      `json:"id_perms,omitempty"`
	DisplayName  string            `json:"display_name,omitempty"`
	Perms2       *PermType2        `json:"perms2,omitempty"`
	UUID         string            `json:"uuid,omitempty"`
	ParentUUID   string            `json:"parent_uuid,omitempty"`
	ParentType   string            `json:"parent_type,omitempty"`
	DomainLimits *DomainLimitsType `json:"domain_limits,omitempty"`

	APIAccessLists   []*APIAccessList   `json:"api_access_lists,omitempty"`
	Namespaces       []*Namespace       `json:"namespaces,omitempty"`
	Projects         []*Project         `json:"projects,omitempty"`
	ServiceTemplates []*ServiceTemplate `json:"service_templates,omitempty"`
	VirtualDNSs      []*VirtualDNS      `json:"virtual_DNSs,omitempty"`
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
		FQName:       []string{},
		IDPerms:      MakeIdPermsType(),
		DisplayName:  "",
		Perms2:       MakePermType2(),
		UUID:         "",
		ParentUUID:   "",
		ParentType:   "",
		Annotations:  MakeKeyValuePairs(),
	}
}

// MakeDomainSlice() makes a slice of Domain
func MakeDomainSlice() []*Domain {
	return []*Domain{}
}
