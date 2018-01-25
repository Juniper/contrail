package models

// Domain

// Domain
//proteus:generate
type Domain struct {
	UUID         string            `json:"uuid,omitempty"`
	ParentUUID   string            `json:"parent_uuid,omitempty"`
	ParentType   string            `json:"parent_type,omitempty"`
	FQName       []string          `json:"fq_name,omitempty"`
	IDPerms      *IdPermsType      `json:"id_perms,omitempty"`
	DisplayName  string            `json:"display_name,omitempty"`
	Annotations  *KeyValuePairs    `json:"annotations,omitempty"`
	Perms2       *PermType2        `json:"perms2,omitempty"`
	DomainLimits *DomainLimitsType `json:"domain_limits,omitempty"`

	APIAccessLists []*APIAccessList `json:"api_access_lists,omitempty"`

	Namespaces []*Namespace `json:"namespaces,omitempty"`

	Projects []*Project `json:"projects,omitempty"`

	ServiceTemplates []*ServiceTemplate `json:"service_templates,omitempty"`

	VirtualDNSs []*VirtualDNS `json:"virtual_DNSs,omitempty"`
}

// MakeDomain makes Domain
func MakeDomain() *Domain {
	return &Domain{
		//TODO(nati): Apply default
		UUID:         "",
		ParentUUID:   "",
		ParentType:   "",
		FQName:       []string{},
		IDPerms:      MakeIdPermsType(),
		DisplayName:  "",
		Annotations:  MakeKeyValuePairs(),
		Perms2:       MakePermType2(),
		DomainLimits: MakeDomainLimitsType(),
	}
}

// MakeDomainSlice() makes a slice of Domain
func MakeDomainSlice() []*Domain {
	return []*Domain{}
}
