package models

// InstanceIP

import "encoding/json"

// InstanceIP
type InstanceIP struct {
	SubnetUUID            string              `json:"subnet_uuid"`
	InstanceIPFamily      IpAddressFamilyType `json:"instance_ip_family"`
	InstanceIPLocalIP     bool                `json:"instance_ip_local_ip"`
	InstanceIPSecondary   bool                `json:"instance_ip_secondary"`
	InstanceIPAddress     IpAddressType       `json:"instance_ip_address"`
	InstanceIPMode        AddressMode         `json:"instance_ip_mode"`
	ServiceInstanceIP     bool                `json:"service_instance_ip"`
	Annotations           *KeyValuePairs      `json:"annotations"`
	Perms2                *PermType2          `json:"perms2"`
	SecondaryIPTrackingIP *SubnetType         `json:"secondary_ip_tracking_ip"`
	DisplayName           string              `json:"display_name"`
	ParentType            string              `json:"parent_type"`
	IDPerms               *IdPermsType        `json:"id_perms"`
	UUID                  string              `json:"uuid"`
	ParentUUID            string              `json:"parent_uuid"`
	FQName                []string            `json:"fq_name"`
	ServiceHealthCheckIP  bool                `json:"service_health_check_ip"`

	VirtualNetworkRefs          []*InstanceIPVirtualNetworkRef          `json:"virtual_network_refs"`
	VirtualMachineInterfaceRefs []*InstanceIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
	PhysicalRouterRefs          []*InstanceIPPhysicalRouterRef          `json:"physical_router_refs"`
	VirtualRouterRefs           []*InstanceIPVirtualRouterRef           `json:"virtual_router_refs"`
	NetworkIpamRefs             []*InstanceIPNetworkIpamRef             `json:"network_ipam_refs"`

	FloatingIPs []*FloatingIP `json:"floating_ips"`
}

// InstanceIPNetworkIpamRef references each other
type InstanceIPNetworkIpamRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// InstanceIPVirtualNetworkRef references each other
type InstanceIPVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// InstanceIPVirtualMachineInterfaceRef references each other
type InstanceIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// InstanceIPPhysicalRouterRef references each other
type InstanceIPPhysicalRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// InstanceIPVirtualRouterRef references each other
type InstanceIPVirtualRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *InstanceIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeInstanceIP makes InstanceIP
func MakeInstanceIP() *InstanceIP {
	return &InstanceIP{
		//TODO(nati): Apply default
		ServiceHealthCheckIP:  false,
		UUID:                  "",
		ParentUUID:            "",
		FQName:                []string{},
		InstanceIPAddress:     MakeIpAddressType(),
		SubnetUUID:            "",
		InstanceIPFamily:      MakeIpAddressFamilyType(),
		InstanceIPLocalIP:     false,
		InstanceIPSecondary:   false,
		SecondaryIPTrackingIP: MakeSubnetType(),
		InstanceIPMode:        MakeAddressMode(),
		ServiceInstanceIP:     false,
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		ParentType:            "",
	}
}

// MakeInstanceIPSlice() makes a slice of InstanceIP
func MakeInstanceIPSlice() []*InstanceIP {
	return []*InstanceIP{}
}
