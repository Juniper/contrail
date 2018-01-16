package models

// InstanceIP

import "encoding/json"

// InstanceIP
type InstanceIP struct {
	ParentType            string              `json:"parent_type,omitempty"`
	SecondaryIPTrackingIP *SubnetType         `json:"secondary_ip_tracking_ip,omitempty"`
	InstanceIPAddress     IpAddressType       `json:"instance_ip_address,omitempty"`
	InstanceIPSecondary   bool                `json:"instance_ip_secondary,omitempty"`
	ParentUUID            string              `json:"parent_uuid,omitempty"`
	IDPerms               *IdPermsType        `json:"id_perms,omitempty"`
	InstanceIPFamily      IpAddressFamilyType `json:"instance_ip_family,omitempty"`
	InstanceIPLocalIP     bool                `json:"instance_ip_local_ip,omitempty"`
	ServiceInstanceIP     bool                `json:"service_instance_ip,omitempty"`
	UUID                  string              `json:"uuid,omitempty"`
	FQName                []string            `json:"fq_name,omitempty"`
	DisplayName           string              `json:"display_name,omitempty"`
	ServiceHealthCheckIP  bool                `json:"service_health_check_ip,omitempty"`
	SubnetUUID            string              `json:"subnet_uuid,omitempty"`
	Perms2                *PermType2          `json:"perms2,omitempty"`
	InstanceIPMode        AddressMode         `json:"instance_ip_mode,omitempty"`
	Annotations           *KeyValuePairs      `json:"annotations,omitempty"`

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
		FQName:                []string{},
		DisplayName:           "",
		ServiceHealthCheckIP:  false,
		SubnetUUID:            "",
		ServiceInstanceIP:     false,
		UUID:                  "",
		InstanceIPMode:        MakeAddressMode(),
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		SecondaryIPTrackingIP: MakeSubnetType(),
		InstanceIPAddress:     MakeIpAddressType(),
		ParentType:            "",
		IDPerms:               MakeIdPermsType(),
		InstanceIPFamily:      MakeIpAddressFamilyType(),
		InstanceIPLocalIP:     false,
		InstanceIPSecondary:   false,
		ParentUUID:            "",
	}
}

// MakeInstanceIPSlice() makes a slice of InstanceIP
func MakeInstanceIPSlice() []*InstanceIP {
	return []*InstanceIP{}
}
