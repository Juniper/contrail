package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan,omitempty"`
	NHMode                   NHModeType          `json:"nh_mode,omitempty"`
	JuniperHeader            bool                `json:"juniper_header"`
	RoutingInstance          string              `json:"routing_instance,omitempty"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address,omitempty"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring"`
	AnalyzerName             string              `json:"analyzer_name,omitempty"`
	UDPPort                  int                 `json:"udp_port,omitempty"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header,omitempty"`
	Encapsulation            string              `json:"encapsulation,omitempty"`
	AnalyzerMacAddress       string              `json:"analyzer_mac_address,omitempty"`
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
		AnalyzerIPAddress:        "",
		NicAssistedMirroring:     false,
		NicAssistedMirroringVlan: MakeVlanIdType(),
		NHMode:             MakeNHModeType(),
		JuniperHeader:      false,
		RoutingInstance:    "",
		AnalyzerMacAddress: "",
		AnalyzerName:       "",
		UDPPort:            0,
		StaticNHHeader:     MakeStaticMirrorNhType(),
		Encapsulation:      "",
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
