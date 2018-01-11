package models

// ServiceAppliance

import "encoding/json"

// ServiceAppliance
type ServiceAppliance struct {
	ParentUUID                      string           `json:"parent_uuid"`
	FQName                          []string         `json:"fq_name"`
	DisplayName                     string           `json:"display_name"`
	Annotations                     *KeyValuePairs   `json:"annotations"`
	Perms2                          *PermType2       `json:"perms2"`
	UUID                            string           `json:"uuid"`
	ServiceApplianceProperties      *KeyValuePairs   `json:"service_appliance_properties"`
	ParentType                      string           `json:"parent_type"`
	IDPerms                         *IdPermsType     `json:"id_perms"`
	ServiceApplianceUserCredentials *UserCredentials `json:"service_appliance_user_credentials"`
	ServiceApplianceIPAddress       IpAddressType    `json:"service_appliance_ip_address"`

	PhysicalInterfaceRefs []*ServiceAppliancePhysicalInterfaceRef `json:"physical_interface_refs"`
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
		Perms2:                          MakePermType2(),
		UUID:                            "",
		ParentUUID:                      "",
		FQName:                          []string{},
		DisplayName:                     "",
		Annotations:                     MakeKeyValuePairs(),
		ServiceApplianceUserCredentials: MakeUserCredentials(),
		ServiceApplianceIPAddress:       MakeIpAddressType(),
		ServiceApplianceProperties:      MakeKeyValuePairs(),
		ParentType:                      "",
		IDPerms:                         MakeIdPermsType(),
	}
}

// MakeServiceApplianceSlice() makes a slice of ServiceAppliance
func MakeServiceApplianceSlice() []*ServiceAppliance {
	return []*ServiceAppliance{}
}
