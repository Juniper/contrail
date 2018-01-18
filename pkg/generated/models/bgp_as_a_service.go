package models

// BGPAsAService

import "encoding/json"

// BGPAsAService
type BGPAsAService struct {
	BgpaasIpv4MappedIpv6Nexthop      bool                 `json:"bgpaas_ipv4_mapped_ipv6_nexthop"`
	UUID                             string               `json:"uuid,omitempty"`
	BgpaasShared                     bool                 `json:"bgpaas_shared"`
	IDPerms                          *IdPermsType         `json:"id_perms,omitempty"`
	Perms2                           *PermType2           `json:"perms2,omitempty"`
	BgpaasSessionAttributes          string               `json:"bgpaas_session_attributes,omitempty"`
	AutonomousSystem                 AutonomousSystemType `json:"autonomous_system,omitempty"`
	ParentUUID                       string               `json:"parent_uuid,omitempty"`
	ParentType                       string               `json:"parent_type,omitempty"`
	DisplayName                      string               `json:"display_name,omitempty"`
	BgpaasIPAddress                  IpAddressType        `json:"bgpaas_ip_address,omitempty"`
	FQName                           []string             `json:"fq_name,omitempty"`
	Annotations                      *KeyValuePairs       `json:"annotations,omitempty"`
	BgpaasSuppressRouteAdvertisement bool                 `json:"bgpaas_suppress_route_advertisement"`

	VirtualMachineInterfaceRefs []*BGPAsAServiceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	ServiceHealthCheckRefs      []*BGPAsAServiceServiceHealthCheckRef      `json:"service_health_check_refs,omitempty"`
}

// BGPAsAServiceServiceHealthCheckRef references each other
type BGPAsAServiceServiceHealthCheckRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// BGPAsAServiceVirtualMachineInterfaceRef references each other
type BGPAsAServiceVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *BGPAsAService) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBGPAsAService makes BGPAsAService
func MakeBGPAsAService() *BGPAsAService {
	return &BGPAsAService{
		//TODO(nati): Apply default
		BgpaasIPAddress:                  MakeIpAddressType(),
		AutonomousSystem:                 MakeAutonomousSystemType(),
		ParentUUID:                       "",
		ParentType:                       "",
		DisplayName:                      "",
		BgpaasSuppressRouteAdvertisement: false,
		FQName:                      []string{},
		Annotations:                 MakeKeyValuePairs(),
		BgpaasShared:                false,
		BgpaasIpv4MappedIpv6Nexthop: false,
		UUID: "",
		BgpaasSessionAttributes: "",
		IDPerms:                 MakeIdPermsType(),
		Perms2:                  MakePermType2(),
	}
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
	return []*BGPAsAService{}
}
