package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	AnalyzerIPAddress        string              `json:"analyzer_ip_address,omitempty"`
	Encapsulation            string              `json:"encapsulation,omitempty"`
	AnalyzerMacAddress       string              `json:"analyzer_mac_address,omitempty"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring"`
	NHMode                   NHModeType          `json:"nh_mode,omitempty"`
	JuniperHeader            bool                `json:"juniper_header"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header,omitempty"`
	RoutingInstance          string              `json:"routing_instance,omitempty"`
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan,omitempty"`
	AnalyzerName             string              `json:"analyzer_name,omitempty"`
	UDPPort                  int                 `json:"udp_port,omitempty"`
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
		UDPPort:                  0,
		RoutingInstance:          "",
		NicAssistedMirroring:     false,
		NHMode:                   MakeNHModeType(),
		JuniperHeader:            false,
		StaticNHHeader:           MakeStaticMirrorNhType(),
		AnalyzerIPAddress:        "",
		Encapsulation:            "",
		AnalyzerMacAddress:       "",
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
