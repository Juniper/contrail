package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan,omitempty"`
	AnalyzerName             string              `json:"analyzer_name,omitempty"`
	NHMode                   NHModeType          `json:"nh_mode,omitempty"`
	RoutingInstance          string              `json:"routing_instance,omitempty"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address,omitempty"`
	JuniperHeader            bool                `json:"juniper_header"`
	UDPPort                  int                 `json:"udp_port,omitempty"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header,omitempty"`
	Encapsulation            string              `json:"encapsulation,omitempty"`
	AnalyzerMacAddress       string              `json:"analyzer_mac_address,omitempty"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring"`
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
		NicAssistedMirroringVlan: MakeVlanIdType(),
		AnalyzerName:             "",
		NHMode:                   MakeNHModeType(),
		RoutingInstance:          "",
		AnalyzerIPAddress:        "",
		NicAssistedMirroring:     false,
		JuniperHeader:            false,
		UDPPort:                  0,
		StaticNHHeader:           MakeStaticMirrorNhType(),
		Encapsulation:            "",
		AnalyzerMacAddress:       "",
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
