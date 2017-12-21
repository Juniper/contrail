package models

// GracefulRestartParametersType

import "encoding/json"

// GracefulRestartParametersType
type GracefulRestartParametersType struct {
	BGPHelperEnable      bool                             `json:"bgp_helper_enable"`
	XMPPHelperEnable     bool                             `json:"xmpp_helper_enable"`
	RestartTime          GracefulRestartTimeType          `json:"restart_time"`
	LongLivedRestartTime LongLivedGracefulRestartTimeType `json:"long_lived_restart_time"`
	Enable               bool                             `json:"enable"`
	EndOfRibTimeout      EndOfRibTimeType                 `json:"end_of_rib_timeout"`
}

// String returns json representation of the object
func (model *GracefulRestartParametersType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeGracefulRestartParametersType makes GracefulRestartParametersType
func MakeGracefulRestartParametersType() *GracefulRestartParametersType {
	return &GracefulRestartParametersType{
		//TODO(nati): Apply default
		BGPHelperEnable:      false,
		XMPPHelperEnable:     false,
		RestartTime:          MakeGracefulRestartTimeType(),
		LongLivedRestartTime: MakeLongLivedGracefulRestartTimeType(),
		Enable:               false,
		EndOfRibTimeout:      MakeEndOfRibTimeType(),
	}
}

// InterfaceToGracefulRestartParametersType makes GracefulRestartParametersType from interface
func InterfaceToGracefulRestartParametersType(iData interface{}) *GracefulRestartParametersType {
	data := iData.(map[string]interface{})
	return &GracefulRestartParametersType{
		Enable: data["enable"].(bool),

		//{"description":"Enable/Disable knob for all GR parameters to take effect","type":"boolean"}
		EndOfRibTimeout: InterfaceToEndOfRibTimeType(data["end_of_rib_timeout"]),

		//{"description":"Maximum time (in seconds) to wait for EndOfRib reception/transmission","type":"integer","minimum":0,"maximum":4095}
		BGPHelperEnable: data["bgp_helper_enable"].(bool),

		//{"description":"Enable GR Helper mode for BGP peers in contrail-control","type":"boolean"}
		XMPPHelperEnable: data["xmpp_helper_enable"].(bool),

		//{"description":"Enable GR Helper mode for XMPP peers (agents) in contrail-control","type":"boolean"}
		RestartTime: InterfaceToGracefulRestartTimeType(data["restart_time"]),

		//{"description":"Time (in seconds) taken by the restarting speaker to get back to stable state","type":"integer","minimum":0,"maximum":4095}
		LongLivedRestartTime: InterfaceToLongLivedGracefulRestartTimeType(data["long_lived_restart_time"]),

		//{"description":"Extended Time (in seconds) taken by the restarting speaker after restart-time to get back to stable state","type":"integer","minimum":0,"maximum":16777215}

	}
}

// InterfaceToGracefulRestartParametersTypeSlice makes a slice of GracefulRestartParametersType from interface
func InterfaceToGracefulRestartParametersTypeSlice(data interface{}) []*GracefulRestartParametersType {
	list := data.([]interface{})
	result := MakeGracefulRestartParametersTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToGracefulRestartParametersType(item))
	}
	return result
}

// MakeGracefulRestartParametersTypeSlice() makes a slice of GracefulRestartParametersType
func MakeGracefulRestartParametersTypeSlice() []*GracefulRestartParametersType {
	return []*GracefulRestartParametersType{}
}
