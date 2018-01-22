package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	NHMode                   NHModeType          `json:"nh_mode,omitempty"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header,omitempty"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address,omitempty"`
	AnalyzerMacAddress       string              `json:"analyzer_mac_address,omitempty"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring"`
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan,omitempty"`
	AnalyzerName             string              `json:"analyzer_name,omitempty"`
	RoutingInstance          string              `json:"routing_instance,omitempty"`
	Encapsulation            string              `json:"encapsulation,omitempty"`
	JuniperHeader            bool                `json:"juniper_header"`
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
		AnalyzerMacAddress:       "",
		NicAssistedMirroring:     false,
		NicAssistedMirroringVlan: MakeVlanIdType(),
		AnalyzerName:             "",
		NHMode:                   MakeNHModeType(),
		StaticNHHeader:           MakeStaticMirrorNhType(),
		AnalyzerIPAddress:        "",
		JuniperHeader:            false,
		UDPPort:                  0,
		RoutingInstance:          "",
		Encapsulation:            "",
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
