package models

// ServiceAppliance

// ServiceAppliance
//proteus:generate
type ServiceAppliance struct {
	UUID                            string           `json:"uuid,omitempty"`
	ParentUUID                      string           `json:"parent_uuid,omitempty"`
	ParentType                      string           `json:"parent_type,omitempty"`
	FQName                          []string         `json:"fq_name,omitempty"`
	IDPerms                         *IdPermsType     `json:"id_perms,omitempty"`
	DisplayName                     string           `json:"display_name,omitempty"`
	Annotations                     *KeyValuePairs   `json:"annotations,omitempty"`
	Perms2                          *PermType2       `json:"perms2,omitempty"`
	ServiceApplianceUserCredentials *UserCredentials `json:"service_appliance_user_credentials,omitempty"`
	ServiceApplianceIPAddress       IpAddressType    `json:"service_appliance_ip_address,omitempty"`
	ServiceApplianceProperties      *KeyValuePairs   `json:"service_appliance_properties,omitempty"`

	PhysicalInterfaceRefs []*ServiceAppliancePhysicalInterfaceRef `json:"physical_interface_refs,omitempty"`
}

// ServiceAppliancePhysicalInterfaceRef references each other
type ServiceAppliancePhysicalInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceApplianceInterfaceType
}

// MakeServiceAppliance makes ServiceAppliance
func MakeServiceAppliance() *ServiceAppliance {
	return &ServiceAppliance{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ServiceApplianceUserCredentials: MakeUserCredentials(),
		ServiceApplianceIPAddress:       MakeIpAddressType(),
		ServiceApplianceProperties:      MakeKeyValuePairs(),
	}
}

// MakeServiceApplianceSlice() makes a slice of ServiceAppliance
func MakeServiceApplianceSlice() []*ServiceAppliance {
	return []*ServiceAppliance{}
}
