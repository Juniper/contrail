package models

// InstanceIP

import "encoding/json"

// InstanceIP
type InstanceIP struct {
	ServiceInstanceIP     bool                `json:"service_instance_ip"`
	DisplayName           string              `json:"display_name,omitempty"`
	InstanceIPSecondary   bool                `json:"instance_ip_secondary"`
	ParentType            string              `json:"parent_type,omitempty"`
	Annotations           *KeyValuePairs      `json:"annotations,omitempty"`
	ServiceHealthCheckIP  bool                `json:"service_health_check_ip"`
	SecondaryIPTrackingIP *SubnetType         `json:"secondary_ip_tracking_ip,omitempty"`
	InstanceIPLocalIP     bool                `json:"instance_ip_local_ip"`
	IDPerms               *IdPermsType        `json:"id_perms,omitempty"`
	Perms2                *PermType2          `json:"perms2,omitempty"`
	InstanceIPMode        AddressMode         `json:"instance_ip_mode,omitempty"`
	SubnetUUID            string              `json:"subnet_uuid,omitempty"`
	UUID                  string              `json:"uuid,omitempty"`
	ParentUUID            string              `json:"parent_uuid,omitempty"`
	FQName                []string            `json:"fq_name,omitempty"`
	InstanceIPAddress     IpAddressType       `json:"instance_ip_address,omitempty"`
	InstanceIPFamily      IpAddressFamilyType `json:"instance_ip_family,omitempty"`

	NetworkIpamRefs             []*InstanceIPNetworkIpamRef             `json:"network_ipam_refs,omitempty"`
	VirtualNetworkRefs          []*InstanceIPVirtualNetworkRef          `json:"virtual_network_refs,omitempty"`
	VirtualMachineInterfaceRefs []*InstanceIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	PhysicalRouterRefs          []*InstanceIPPhysicalRouterRef          `json:"physical_router_refs,omitempty"`
	VirtualRouterRefs           []*InstanceIPVirtualRouterRef           `json:"virtual_router_refs,omitempty"`

	FloatingIPs []*FloatingIP `json:"floating_ips,omitempty"`
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
		ServiceInstanceIP:     false,
		DisplayName:           "",
		ServiceHealthCheckIP:  false,
		SecondaryIPTrackingIP: MakeSubnetType(),
		InstanceIPSecondary:   false,
		ParentType:            "",
		Annotations:           MakeKeyValuePairs(),
		InstanceIPMode:        MakeAddressMode(),
		SubnetUUID:            "",
		InstanceIPLocalIP:     false,
		IDPerms:               MakeIdPermsType(),
		Perms2:                MakePermType2(),
		InstanceIPAddress:     MakeIpAddressType(),
		InstanceIPFamily:      MakeIpAddressFamilyType(),
		UUID:                  "",
		ParentUUID:            "",
		FQName:                []string{},
	}
}

// MakeInstanceIPSlice() makes a slice of InstanceIP
func MakeInstanceIPSlice() []*InstanceIP {
	return []*InstanceIP{}
}
