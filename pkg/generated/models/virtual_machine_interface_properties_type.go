package models

// VirtualMachineInterfacePropertiesType

import "encoding/json"

// VirtualMachineInterfacePropertiesType
type VirtualMachineInterfacePropertiesType struct {
	SubInterfaceVlanTag  int                  `json:"sub_interface_vlan_tag"`
	LocalPreference      int                  `json:"local_preference"`
	InterfaceMirror      *InterfaceMirrorType `json:"interface_mirror"`
	ServiceInterfaceType ServiceInterfaceType `json:"service_interface_type"`
}

// String returns json representation of the object
func (model *VirtualMachineInterfacePropertiesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualMachineInterfacePropertiesType makes VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesType() *VirtualMachineInterfacePropertiesType {
	return &VirtualMachineInterfacePropertiesType{
		//TODO(nati): Apply default
		ServiceInterfaceType: MakeServiceInterfaceType(),
		SubInterfaceVlanTag:  0,
		LocalPreference:      0,
		InterfaceMirror:      MakeInterfaceMirrorType(),
	}
}

// InterfaceToVirtualMachineInterfacePropertiesType makes VirtualMachineInterfacePropertiesType from interface
func InterfaceToVirtualMachineInterfacePropertiesType(iData interface{}) *VirtualMachineInterfacePropertiesType {
	data := iData.(map[string]interface{})
	return &VirtualMachineInterfacePropertiesType{
		ServiceInterfaceType: InterfaceToServiceInterfaceType(data["service_interface_type"]),

		//{"description":"This interface belongs to Service Instance and is tagged as left, right or other","type":"string"}
		SubInterfaceVlanTag: data["sub_interface_vlan_tag"].(int),

		//{"description":"802.1Q VLAN tag to be used if this interface is sub-interface for some other interface.","type":"integer"}
		LocalPreference: data["local_preference"].(int),

		//{"description":"BGP route local preference for routes representing this interface, higher value is higher preference","type":"integer"}
		InterfaceMirror: InterfaceToInterfaceMirrorType(data["interface_mirror"]),

		//{"description":"Interface Mirror configuration","type":"object","properties":{"mirror_to":{"type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}},"traffic_direction":{"default":"both","type":"string","enum":["ingress","egress","both"]}}}

	}
}

// InterfaceToVirtualMachineInterfacePropertiesTypeSlice makes a slice of VirtualMachineInterfacePropertiesType from interface
func InterfaceToVirtualMachineInterfacePropertiesTypeSlice(data interface{}) []*VirtualMachineInterfacePropertiesType {
	list := data.([]interface{})
	result := MakeVirtualMachineInterfacePropertiesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualMachineInterfacePropertiesType(item))
	}
	return result
}

// MakeVirtualMachineInterfacePropertiesTypeSlice() makes a slice of VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesTypeSlice() []*VirtualMachineInterfacePropertiesType {
	return []*VirtualMachineInterfacePropertiesType{}
}
