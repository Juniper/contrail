package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeMirrorActionType makes MirrorActionType
// nolint
func MakeMirrorActionType() *MirrorActionType {
	return &MirrorActionType{
		//TODO(nati): Apply default
		NicAssistedMirroringVlan: 0,
		AnalyzerName:             "",
		NHMode:                   "",
		JuniperHeader:            false,
		UDPPort:                  0,
		RoutingInstance:          "",
		StaticNHHeader:           MakeStaticMirrorNhType(),
		AnalyzerIPAddress:        "",
		Encapsulation:            "",
		AnalyzerMacAddress:       "",
		NicAssistedMirroring:     false,
	}
}

// MakeMirrorActionType makes MirrorActionType
// nolint
func InterfaceToMirrorActionType(i interface{}) *MirrorActionType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &MirrorActionType{
		//TODO(nati): Apply default
		NicAssistedMirroringVlan: common.InterfaceToInt64(m["nic_assisted_mirroring_vlan"]),
		AnalyzerName:             common.InterfaceToString(m["analyzer_name"]),
		NHMode:                   common.InterfaceToString(m["nh_mode"]),
		JuniperHeader:            common.InterfaceToBool(m["juniper_header"]),
		UDPPort:                  common.InterfaceToInt64(m["udp_port"]),
		RoutingInstance:          common.InterfaceToString(m["routing_instance"]),
		StaticNHHeader:           InterfaceToStaticMirrorNhType(m["static_nh_header"]),
		AnalyzerIPAddress:        common.InterfaceToString(m["analyzer_ip_address"]),
		Encapsulation:            common.InterfaceToString(m["encapsulation"]),
		AnalyzerMacAddress:       common.InterfaceToString(m["analyzer_mac_address"]),
		NicAssistedMirroring:     common.InterfaceToBool(m["nic_assisted_mirroring"]),
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
// nolint
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}

// InterfaceToMirrorActionTypeSlice() makes a slice of MirrorActionType
// nolint
func InterfaceToMirrorActionTypeSlice(i interface{}) []*MirrorActionType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*MirrorActionType{}
	for _, item := range list {
		result = append(result, InterfaceToMirrorActionType(item))
	}
	return result
}
