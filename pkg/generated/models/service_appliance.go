package models

// ServiceAppliance

import "encoding/json"

// ServiceAppliance
type ServiceAppliance struct {
	ParentType                      string           `json:"parent_type,omitempty"`
	Perms2                          *PermType2       `json:"perms2,omitempty"`
	ServiceApplianceProperties      *KeyValuePairs   `json:"service_appliance_properties,omitempty"`
	UUID                            string           `json:"uuid,omitempty"`
	ParentUUID                      string           `json:"parent_uuid,omitempty"`
	IDPerms                         *IdPermsType     `json:"id_perms,omitempty"`
	DisplayName                     string           `json:"display_name,omitempty"`
	Annotations                     *KeyValuePairs   `json:"annotations,omitempty"`
	ServiceApplianceUserCredentials *UserCredentials `json:"service_appliance_user_credentials,omitempty"`
	ServiceApplianceIPAddress       IpAddressType    `json:"service_appliance_ip_address,omitempty"`
	FQName                          []string         `json:"fq_name,omitempty"`

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
		ServiceApplianceIPAddress: MakeIpAddressType(),
		FQName:                          []string{},
		IDPerms:                         MakeIdPermsType(),
		DisplayName:                     "",
		Annotations:                     MakeKeyValuePairs(),
		ServiceApplianceUserCredentials: MakeUserCredentials(),
		UUID:       "",
		ParentUUID: "",
		ParentType: "",
		Perms2:     MakePermType2(),
		ServiceApplianceProperties: MakeKeyValuePairs(),
	}
}

// MakeServiceApplianceSlice() makes a slice of ServiceAppliance
func MakeServiceApplianceSlice() []*ServiceAppliance {
	return []*ServiceAppliance{}
}
