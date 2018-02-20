package models

// Flavor

// Flavor
//proteus:generate
type Flavor struct {
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	Name        string         `json:"name,omitempty"`
	Disk        int            `json:"disk,omitempty"`
	Vcpus       int            `json:"vcpus,omitempty"`
	RAM         int            `json:"ram,omitempty"`
	ID          string         `json:"id,omitempty"`
	Property    string         `json:"property,omitempty"`
	RXTXFactor  int            `json:"rxtx_factor,omitempty"`
	Swap        int            `json:"swap,omitempty"`
	IsPublic    bool           `json:"is_public"`
	Ephemeral   int            `json:"ephemeral,omitempty"`
	Links       *OpenStackLink `json:"links,omitempty"`
}

// MakeFlavor makes Flavor
func MakeFlavor() *Flavor {
	return &Flavor{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		Name:        "",
		Disk:        0,
		Vcpus:       0,
		RAM:         0,
		ID:          "",
		Property:    "",
		RXTXFactor:  0,
		Swap:        0,
		IsPublic:    false,
		Ephemeral:   0,
		Links:       MakeOpenStackLink(),
	}
}

// MakeFlavorSlice() makes a slice of Flavor
func MakeFlavorSlice() []*Flavor {
	return []*Flavor{}
}
