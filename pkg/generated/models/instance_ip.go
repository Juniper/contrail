package models

// InstanceIP

import "encoding/json"

// InstanceIP
type InstanceIP struct {
	SecondaryIPTrackingIP *SubnetType         `json:"secondary_ip_tracking_ip,omitempty"`
	InstanceIPAddress     IpAddressType       `json:"instance_ip_address,omitempty"`
	InstanceIPMode        AddressMode         `json:"instance_ip_mode,omitempty"`
	SubnetUUID            string              `json:"subnet_uuid,omitempty"`
	InstanceIPLocalIP     bool                `json:"instance_ip_local_ip"`
	InstanceIPSecondary   bool                `json:"instance_ip_secondary"`
	Perms2                *PermType2          `json:"perms2,omitempty"`
	ServiceHealthCheckIP  bool                `json:"service_health_check_ip"`
	ParentUUID            string              `json:"parent_uuid,omitempty"`
	IDPerms               *IdPermsType        `json:"id_perms,omitempty"`
	UUID                  string              `json:"uuid,omitempty"`
	ParentType            string              `json:"parent_type,omitempty"`
	DisplayName           string              `json:"display_name,omitempty"`
	Annotations           *KeyValuePairs      `json:"annotations,omitempty"`
	InstanceIPFamily      IpAddressFamilyType `json:"instance_ip_family,omitempty"`
	FQName                []string            `json:"fq_name,omitempty"`
	ServiceInstanceIP     bool                `json:"service_instance_ip"`

	VirtualRouterRefs           []*InstanceIPVirtualRouterRef           `json:"virtual_router_refs,omitempty"`
	NetworkIpamRefs             []*InstanceIPNetworkIpamRef             `json:"network_ipam_refs,omitempty"`
	VirtualNetworkRefs          []*InstanceIPVirtualNetworkRef          `json:"virtual_network_refs,omitempty"`
	VirtualMachineInterfaceRefs []*InstanceIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	PhysicalRouterRefs          []*InstanceIPPhysicalRouterRef          `json:"physical_router_refs,omitempty"`

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
		InstanceIPFamily:      MakeIpAddressFamilyType(),
		ServiceInstanceIP:     false,
		FQName:                []string{},
		Perms2:                MakePermType2(),
		ServiceHealthCheckIP:  false,
		SecondaryIPTrackingIP: MakeSubnetType(),
		InstanceIPAddress:     MakeIpAddressType(),
		InstanceIPMode:        MakeAddressMode(),
		SubnetUUID:            "",
		InstanceIPLocalIP:     false,
		InstanceIPSecondary:   false,
		UUID:                  "",
		ParentUUID:            "",
		IDPerms:               MakeIdPermsType(),
		Annotations:           MakeKeyValuePairs(),
		ParentType:            "",
		DisplayName:           "",
	}
}

// MakeInstanceIPSlice() makes a slice of InstanceIP
func MakeInstanceIPSlice() []*InstanceIP {
	return []*InstanceIP{}
}
