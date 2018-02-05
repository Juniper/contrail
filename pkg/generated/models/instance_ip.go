package models

// InstanceIP

import "encoding/json"

// InstanceIP
type InstanceIP struct {
	InstanceIPLocalIP     bool                `json:"instance_ip_local_ip"`
	FQName                []string            `json:"fq_name,omitempty"`
	ParentType            string              `json:"parent_type,omitempty"`
	SecondaryIPTrackingIP *SubnetType         `json:"secondary_ip_tracking_ip,omitempty"`
	InstanceIPAddress     IpAddressType       `json:"instance_ip_address,omitempty"`
	SubnetUUID            string              `json:"subnet_uuid,omitempty"`
	InstanceIPFamily      IpAddressFamilyType `json:"instance_ip_family,omitempty"`
	ServiceInstanceIP     bool                `json:"service_instance_ip"`
	InstanceIPSecondary   bool                `json:"instance_ip_secondary"`
	ServiceHealthCheckIP  bool                `json:"service_health_check_ip"`
	DisplayName           string              `json:"display_name,omitempty"`
	Annotations           *KeyValuePairs      `json:"annotations,omitempty"`
	Perms2                *PermType2          `json:"perms2,omitempty"`
	InstanceIPMode        AddressMode         `json:"instance_ip_mode,omitempty"`
	IDPerms               *IdPermsType        `json:"id_perms,omitempty"`
	UUID                  string              `json:"uuid,omitempty"`
	ParentUUID            string              `json:"parent_uuid,omitempty"`

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
		ServiceHealthCheckIP:  false,
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		InstanceIPMode:        MakeAddressMode(),
		IDPerms:               MakeIdPermsType(),
		UUID:                  "",
		ParentUUID:            "",
		InstanceIPLocalIP:     false,
		FQName:                []string{},
		ParentType:            "",
		ServiceInstanceIP:     false,
		InstanceIPSecondary:   false,
		SecondaryIPTrackingIP: MakeSubnetType(),
		InstanceIPAddress:     MakeIpAddressType(),
		SubnetUUID:            "",
		InstanceIPFamily:      MakeIpAddressFamilyType(),
	}
}

// MakeInstanceIPSlice() makes a slice of InstanceIP
func MakeInstanceIPSlice() []*InstanceIP {
	return []*InstanceIP{}
}
