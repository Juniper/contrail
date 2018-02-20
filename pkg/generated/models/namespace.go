package models

// Namespace

// Namespace
//proteus:generate
type Namespace struct {
	UUID          string         `json:"uuid,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
	NamespaceCidr *SubnetType    `json:"namespace_cidr,omitempty"`
}

// MakeNamespace makes Namespace
func MakeNamespace() *Namespace {
	return &Namespace{
		//TODO(nati): Apply default
		UUID:          "",
		ParentUUID:    "",
		ParentType:    "",
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
		NamespaceCidr: MakeSubnetType(),
	}
}

// MakeNamespaceSlice() makes a slice of Namespace
func MakeNamespaceSlice() []*Namespace {
	return []*Namespace{}
}
