package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	AnalyzerMacAddress       string              `json:"analyzer_mac_address"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring"`
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan"`
	RoutingInstance          string              `json:"routing_instance"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address"`
	Encapsulation            string              `json:"encapsulation"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header"`
	AnalyzerName             string              `json:"analyzer_name"`
	NHMode                   NHModeType          `json:"nh_mode"`
	JuniperHeader            bool                `json:"juniper_header"`
	UDPPort                  int                 `json:"udp_port"`
}

// String returns json representation of the object
func (model *MirrorActionType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMirrorActionType makes MirrorActionType
func MakeMirrorActionType() *MirrorActionType {
	return &MirrorActionType{
		//TODO(nati): Apply default
		AnalyzerName:             "",
		NHMode:                   MakeNHModeType(),
		JuniperHeader:            false,
		UDPPort:                  0,
		StaticNHHeader:           MakeStaticMirrorNhType(),
		NicAssistedMirroring:     false,
		NicAssistedMirroringVlan: MakeVlanIdType(),
		RoutingInstance:          "",
		AnalyzerIPAddress:        "",
		Encapsulation:            "",
		AnalyzerMacAddress:       "",
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
