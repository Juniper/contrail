package models

// ServiceAppliance

import "encoding/json"

// ServiceAppliance
type ServiceAppliance struct {
	ServiceApplianceIPAddress       IpAddressType    `json:"service_appliance_ip_address,omitempty"`
	Perms2                          *PermType2       `json:"perms2,omitempty"`
	ParentType                      string           `json:"parent_type,omitempty"`
	IDPerms                         *IdPermsType     `json:"id_perms,omitempty"`
	DisplayName                     string           `json:"display_name,omitempty"`
	FQName                          []string         `json:"fq_name,omitempty"`
	ServiceApplianceUserCredentials *UserCredentials `json:"service_appliance_user_credentials,omitempty"`
	ServiceApplianceProperties      *KeyValuePairs   `json:"service_appliance_properties,omitempty"`
	Annotations                     *KeyValuePairs   `json:"annotations,omitempty"`
	UUID                            string           `json:"uuid,omitempty"`
	ParentUUID                      string           `json:"parent_uuid,omitempty"`

	PhysicalInterfaceRefs []*ServiceAppliancePhysicalInterfaceRef `json:"physical_interface_refs,omitempty"`
}

// ServiceAppliancePhysicalInterfaceRef references each other
type ServiceAppliancePhysicalInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceApplianceInterfaceType
}

// String returns json representation of the object
func (model *ServiceAppliance) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceAppliance makes ServiceAppliance
func MakeServiceAppliance() *ServiceAppliance {
	return &ServiceAppliance{
		//TODO(nati): Apply default
		ServiceApplianceUserCredentials: MakeUserCredentials(),
		ServiceApplianceProperties:      MakeKeyValuePairs(),
		Annotations:                     MakeKeyValuePairs(),
		UUID:                            "",
		ParentUUID:                      "",
		FQName:                          []string{},
		ServiceApplianceIPAddress: MakeIpAddressType(),
		Perms2:      MakePermType2(),
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
	}
}

// MakeServiceApplianceSlice() makes a slice of ServiceAppliance
func MakeServiceApplianceSlice() []*ServiceAppliance {
	return []*ServiceAppliance{}
}
