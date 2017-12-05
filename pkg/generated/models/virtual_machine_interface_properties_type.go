package models

// VirtualMachineInterfacePropertiesType

import "encoding/json"

type VirtualMachineInterfacePropertiesType struct {
	SubInterfaceVlanTag  int                  `json:"sub_interface_vlan_tag"`
	LocalPreference      int                  `json:"local_preference"`
	InterfaceMirror      *InterfaceMirrorType `json:"interface_mirror"`
	ServiceInterfaceType ServiceInterfaceType `json:"service_interface_type"`
}

func (model *VirtualMachineInterfacePropertiesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeVirtualMachineInterfacePropertiesType() *VirtualMachineInterfacePropertiesType {
	return &VirtualMachineInterfacePropertiesType{
		//TODO(nati): Apply default
		InterfaceMirror:      MakeInterfaceMirrorType(),
		ServiceInterfaceType: MakeServiceInterfaceType(),
		SubInterfaceVlanTag:  0,
		LocalPreference:      0,
	}
}

func InterfaceToVirtualMachineInterfacePropertiesType(iData interface{}) *VirtualMachineInterfacePropertiesType {
	data := iData.(map[string]interface{})
	return &VirtualMachineInterfacePropertiesType{
		SubInterfaceVlanTag: data["sub_interface_vlan_tag"].(int),

		//{"Title":"","Description":"802.1Q VLAN tag to be used if this interface is sub-interface for some other interface.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SubInterfaceVlanTag","GoType":"int"}
		LocalPreference: data["local_preference"].(int),

		//{"Title":"","Description":"BGP route local preference for routes representing this interface, higher value is higher preference","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LocalPreference","GoType":"int"}
		InterfaceMirror: InterfaceToInterfaceMirrorType(data["interface_mirror"]),

		//{"Title":"","Description":"Interface Mirror configuration","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"mirror_to":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"analyzer_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerIPAddress","GoType":"string"},"analyzer_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerMacAddress","GoType":"string"},"analyzer_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerName","GoType":"string"},"encapsulation":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Encapsulation","GoType":"string"},"juniper_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JuniperHeader","GoType":"bool"},"nh_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dynamic","static"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/NHModeType","CollectionType":"","Column":"","Item":null,"GoName":"NHMode","GoType":"NHModeType"},"nic_assisted_mirroring":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroring","GoType":"bool"},"nic_assisted_mirroring_vlan":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroringVlan","GoType":"VlanIdType"},"routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoutingInstance","GoType":"string"},"static_nh_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"vni":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType"},"vtep_dst_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string"},"vtep_dst_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/StaticMirrorNhType","CollectionType":"","Column":"","Item":null,"GoName":"StaticNHHeader","GoType":"StaticMirrorNhType"},"udp_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UDPPort","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MirrorActionType","CollectionType":"","Column":"","Item":null,"GoName":"MirrorTo","GoType":"MirrorActionType"},"traffic_direction":{"Title":"","Description":"","SQL":"","Default":"both","Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["ingress","egress","both"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TrafficDirectionType","CollectionType":"","Column":"","Item":null,"GoName":"TrafficDirection","GoType":"TrafficDirectionType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/InterfaceMirrorType","CollectionType":"","Column":"","Item":null,"GoName":"InterfaceMirror","GoType":"InterfaceMirrorType"}
		ServiceInterfaceType: InterfaceToServiceInterfaceType(data["service_interface_type"]),

		//{"Title":"","Description":"This interface belongs to Service Instance and is tagged as left, right or other","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceInterfaceType","GoType":"ServiceInterfaceType"}

	}
}

func InterfaceToVirtualMachineInterfacePropertiesTypeSlice(data interface{}) []*VirtualMachineInterfacePropertiesType {
	list := data.([]interface{})
	result := MakeVirtualMachineInterfacePropertiesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualMachineInterfacePropertiesType(item))
	}
	return result
}

func MakeVirtualMachineInterfacePropertiesTypeSlice() []*VirtualMachineInterfacePropertiesType {
	return []*VirtualMachineInterfacePropertiesType{}
}
