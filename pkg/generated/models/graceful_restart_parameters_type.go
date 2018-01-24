package models

// GracefulRestartParametersType

import "encoding/json"

// GracefulRestartParametersType
type GracefulRestartParametersType struct {
	Enable               bool                             `json:"enable"`
	EndOfRibTimeout      EndOfRibTimeType                 `json:"end_of_rib_timeout,omitempty"`
	BGPHelperEnable      bool                             `json:"bgp_helper_enable"`
	XMPPHelperEnable     bool                             `json:"xmpp_helper_enable"`
	RestartTime          GracefulRestartTimeType          `json:"restart_time,omitempty"`
	LongLivedRestartTime LongLivedGracefulRestartTimeType `json:"long_lived_restart_time,omitempty"`
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
		EndOfRibTimeout:      MakeEndOfRibTimeType(),
		BGPHelperEnable:      false,
		XMPPHelperEnable:     false,
		RestartTime:          MakeGracefulRestartTimeType(),
		LongLivedRestartTime: MakeLongLivedGracefulRestartTimeType(),
		Enable:               false,
	}
}

// MakeGracefulRestartParametersTypeSlice() makes a slice of GracefulRestartParametersType
func MakeGracefulRestartParametersTypeSlice() []*GracefulRestartParametersType {
	return []*GracefulRestartParametersType{}
}
