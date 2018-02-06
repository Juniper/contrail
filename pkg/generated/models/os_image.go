package models

// OsImage

// OsImage
//proteus:generate
type OsImage struct {
	UUID            string         `json:"uuid,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	Name            string         `json:"name,omitempty"`
	Owner           string         `json:"owner,omitempty"`
	ID              string         `json:"id,omitempty"`
	Size_           int            `json:"size,omitempty"`
	Status          string         `json:"status,omitempty"`
	Location        string         `json:"location,omitempty"`
	File            string         `json:"file,omitempty"`
	Checksum        string         `json:"checksum,omitempty"`
	CreatedAt       string         `json:"created_at,omitempty"`
	UpdatedAt       string         `json:"updated_at,omitempty"`
	ContainerFormat string         `json:"container_format,omitempty"`
	DiskFormat      string         `json:"disk_format,omitempty"`
	Protected       bool           `json:"protected"`
	Visibility      string         `json:"visibility,omitempty"`
	Property        string         `json:"property,omitempty"`
	MinDisk         int            `json:"min_disk,omitempty"`
	MinRAM          int            `json:"min_ram,omitempty"`
	Tags            string         `json:"tags,omitempty"`
}

// MakeOsImage makes OsImage
func MakeOsImage() *OsImage {
	return &OsImage{
		//TODO(nati): Apply default
		UUID:            "",
		ParentUUID:      "",
		ParentType:      "",
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		Name:            "",
		Owner:           "",
		ID:              "",
		Size_:           0,
		Status:          "",
		Location:        "",
		File:            "",
		Checksum:        "",
		CreatedAt:       "",
		UpdatedAt:       "",
		ContainerFormat: "",
		DiskFormat:      "",
		Protected:       false,
		Visibility:      "",
		Property:        "",
		MinDisk:         0,
		MinRAM:          0,
		Tags:            "",
	}
}

// MakeOsImageSlice() makes a slice of OsImage
func MakeOsImageSlice() []*OsImage {
	return []*OsImage{}
}
