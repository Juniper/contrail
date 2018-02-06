package models

// Server

// Server
//proteus:generate
type Server struct {
	UUID        string                   `json:"uuid,omitempty"`
	ParentUUID  string                   `json:"parent_uuid,omitempty"`
	ParentType  string                   `json:"parent_type,omitempty"`
	FQName      []string                 `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType             `json:"id_perms,omitempty"`
	DisplayName string                   `json:"display_name,omitempty"`
	Annotations *KeyValuePairs           `json:"annotations,omitempty"`
	Perms2      *PermType2               `json:"perms2,omitempty"`
	Created     string                   `json:"created,omitempty"`
	HostId      string                   `json:"hostId,omitempty"`
	ID          string                   `json:"id,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Image       *OpenStackImageProperty  `json:"image,omitempty"`
	Flavor      *OpenStackFlavorProperty `json:"flavor,omitempty"`
	Addresses   *OpenStackAddress        `json:"addresses,omitempty"`
	AccessIPv4  string                   `json:"accessIPv4,omitempty"`
	AccessIPv6  string                   `json:"accessIPv6,omitempty"`
	ConfigDrive bool                     `json:"config_drive"`
	Progress    int                      `json:"progress,omitempty"`
	Status      string                   `json:"status,omitempty"`
	HostStatus  string                   `json:"host_status,omitempty"`
	TenantID    string                   `json:"tenant_id,omitempty"`
	Updated     string                   `json:"updated,omitempty"`
	UserID      int                      `json:"user_id,omitempty"`
	Locked      bool                     `json:"locked"`
}

// MakeServer makes Server
func MakeServer() *Server {
	return &Server{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		Created:     "",
		HostId:      "",
		ID:          "",
		Name:        "",
		Image:       MakeOpenStackImageProperty(),
		Flavor:      MakeOpenStackFlavorProperty(),
		Addresses:   MakeOpenStackAddress(),
		AccessIPv4:  "",
		AccessIPv6:  "",
		ConfigDrive: false,
		Progress:    0,
		Status:      "",
		HostStatus:  "",
		TenantID:    "",
		Updated:     "",
		UserID:      0,
		Locked:      false,
	}
}

// MakeServerSlice() makes a slice of Server
func MakeServerSlice() []*Server {
	return []*Server{}
}
