package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address"`
	Annotations                  *KeyValuePairs       `json:"annotations"`
	Perms2                       *PermType2           `json:"perms2"`
	UUID                         string               `json:"uuid"`
	FQName                       []string             `json:"fq_name"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction"`
	DisplayName                  string               `json:"display_name"`
	IDPerms                      *IdPermsType         `json:"id_perms"`

	// project <utils.Reference Value>
	ProjectRefs []*FloatingIPProjectRef
	// virtual_machine_interface <utils.Reference Value>
	VirtualMachineInterfaceRefs []*FloatingIPVirtualMachineInterfaceRef

	InstanceIPs     []*FloatingIPInstanceIP
	FloatingIPPools []*FloatingIPFloatingIPPool
}

// FloatingIPProjectRef references each other
type FloatingIPProjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// FloatingIPVirtualMachineInterfaceRef references each other
type FloatingIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// FloatingIP parents relation object

type FloatingIPInstanceIP struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

type FloatingIPFloatingIPPool struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *FloatingIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFloatingIP makes FloatingIP
func MakeFloatingIP() *FloatingIP {
	return &FloatingIP{
		//TODO(nati): Apply default
		FloatingIPAddress:            MakeIpAddressType(),
		FloatingIPFixedIPAddress:     MakeIpAddressType(),
		Annotations:                  MakeKeyValuePairs(),
		FloatingIPIsVirtualIP:        false,
		FloatingIPPortMappings:       MakePortMappings(),
		FloatingIPPortMappingsEnable: false,
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
		DisplayName:                  "",
		Perms2:                       MakePermType2(),
		UUID:                         "",
		FQName:                       []string{},
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		IDPerms:                      MakeIdPermsType(),
	}
}

// InterfaceToFloatingIP makes FloatingIP from interface
func InterfaceToFloatingIP(iData interface{}) *FloatingIP {
	data := iData.(map[string]interface{})
	return &FloatingIP{
		FloatingIPIsVirtualIP: data["floating_ip_is_virtual_ip"].(bool),

		//{"Title":"","Description":"This floating ip is used as virtual ip (VIP) in case of LBaaS.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"floating_ip_is_virtual_ip","Item":null,"GoName":"FloatingIPIsVirtualIP","GoType":"bool","GoPremitive":true}
		FloatingIPAddress: InterfaceToIpAddressType(data["floating_ip_address"]),

		//{"Title":"","Description":"Floating ip address.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"floating_ip_address","Item":null,"GoName":"FloatingIPAddress","GoType":"IpAddressType","GoPremitive":false}
		FloatingIPFixedIPAddress: InterfaceToIpAddressType(data["floating_ip_fixed_ip_address"]),

		//{"Title":"","Description":"This floating is tracking given fixed ip of the interface. The given fixed ip is used in 1:1 NAT.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"floating_ip_fixed_ip_address","Item":null,"GoName":"FloatingIPFixedIPAddress","GoType":"IpAddressType","GoPremitive":false}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair","GoPremitive":false},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs","GoPremitive":false}
		FloatingIPAddressFamily: InterfaceToIpAddressFamilyType(data["floating_ip_address_family"]),

		//{"Title":"","Description":"Ip address family of the floating ip, IpV4 or IpV6","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["v4","v6"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressFamilyType","CollectionType":"","Column":"floating_ip_address_family","Item":null,"GoName":"FloatingIPAddressFamily","GoType":"IpAddressFamilyType","GoPremitive":false}
		FloatingIPPortMappings: InterfaceToPortMappings(data["floating_ip_port_mappings"]),

		//{"Title":"","Description":"List of PortMaps for this floating-ip.","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"port_mappings":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"dst_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DSTPort","GoType":"int","GoPremitive":true},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string","GoPremitive":true},"src_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SRCPort","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortMap","CollectionType":"","Column":"","Item":null,"GoName":"PortMappings","GoType":"PortMap","GoPremitive":false},"GoName":"PortMappings","GoType":"[]*PortMap","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortMappings","CollectionType":"list","Column":"floating_ip_port_mappings","Item":null,"GoName":"FloatingIPPortMappings","GoType":"PortMappings","GoPremitive":false}
		FloatingIPPortMappingsEnable: data["floating_ip_port_mappings_enable"].(bool),

		//{"Title":"","Description":"If it is false, floating-ip Nat is done for all Ports. If it is true, floating-ip Nat is done to the list of PortMaps.","SQL":"bool","Default":false,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"floating_ip_port_mappings_enable","Item":null,"GoName":"FloatingIPPortMappingsEnable","GoType":"bool","GoPremitive":true}
		FloatingIPTrafficDirection: InterfaceToTrafficDirectionType(data["floating_ip_traffic_direction"]),

		//{"Title":"","Description":"Specifies direction of traffic for the floating-ip","SQL":"varchar(255)","Default":"both","Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["ingress","egress","both"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TrafficDirectionType","CollectionType":"","Column":"floating_ip_traffic_direction","Item":null,"GoName":"FloatingIPTrafficDirection","GoType":"TrafficDirectionType","GoPremitive":false}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string","GoPremitive":true}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType","GoPremitive":false},"GoName":"Share","GoType":"[]*ShareType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2","GoPremitive":false}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string","GoPremitive":true}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string","GoPremitive":true},"GoName":"FQName","GoType":"[]string","GoPremitive":true}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string","GoPremitive":true},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string","GoPremitive":true},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string","GoPremitive":true},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string","GoPremitive":true},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"permissions_owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"permissions_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType","GoPremitive":false},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType","GoPremitive":false}

	}
}

// InterfaceToFloatingIPSlice makes a slice of FloatingIP from interface
func InterfaceToFloatingIPSlice(data interface{}) []*FloatingIP {
	list := data.([]interface{})
	result := MakeFloatingIPSlice()
	for _, item := range list {
		result = append(result, InterfaceToFloatingIP(item))
	}
	return result
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
