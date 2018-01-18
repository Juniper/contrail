package models

// ServiceAppliance

import "encoding/json"

// ServiceAppliance
type ServiceAppliance struct {
	Annotations                     *KeyValuePairs   `json:"annotations,omitempty"`
	Perms2                          *PermType2       `json:"perms2,omitempty"`
	UUID                            string           `json:"uuid,omitempty"`
	ParentType                      string           `json:"parent_type,omitempty"`
	FQName                          []string         `json:"fq_name,omitempty"`
	ServiceApplianceUserCredentials *UserCredentials `json:"service_appliance_user_credentials,omitempty"`
	ServiceApplianceIPAddress       IpAddressType    `json:"service_appliance_ip_address,omitempty"`
	DisplayName                     string           `json:"display_name,omitempty"`
	ServiceApplianceProperties      *KeyValuePairs   `json:"service_appliance_properties,omitempty"`
	ParentUUID                      string           `json:"parent_uuid,omitempty"`
	IDPerms                         *IdPermsType     `json:"id_perms,omitempty"`

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
		ServiceApplianceProperties: MakeKeyValuePairs(),
		ParentUUID:                 "",
		IDPerms:                    MakeIdPermsType(),
		Perms2:                     MakePermType2(),
		UUID:                       "",
		ParentType:                 "",
		FQName:                     []string{},
		ServiceApplianceUserCredentials: MakeUserCredentials(),
		ServiceApplianceIPAddress:       MakeIpAddressType(),
		DisplayName:                     "",
		Annotations:                     MakeKeyValuePairs(),
	}
}

// MakeServiceApplianceSlice() makes a slice of ServiceAppliance
func MakeServiceApplianceSlice() []*ServiceAppliance {
	return []*ServiceAppliance{}
}
