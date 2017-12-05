package models

// GracefulRestartParametersType

import "encoding/json"

type GracefulRestartParametersType struct {
	LongLivedRestartTime LongLivedGracefulRestartTimeType `json:"long_lived_restart_time"`
	Enable               bool                             `json:"enable"`
	EndOfRibTimeout      EndOfRibTimeType                 `json:"end_of_rib_timeout"`
	BGPHelperEnable      bool                             `json:"bgp_helper_enable"`
	XMPPHelperEnable     bool                             `json:"xmpp_helper_enable"`
	RestartTime          GracefulRestartTimeType          `json:"restart_time"`
}

func (model *GracefulRestartParametersType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

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

func InterfaceToGracefulRestartParametersType(iData interface{}) *GracefulRestartParametersType {
	data := iData.(map[string]interface{})
	return &GracefulRestartParametersType{
		RestartTime: InterfaceToGracefulRestartTimeType(data["restart_time"]),

		//{"Title":"","Description":"Time (in seconds) taken by the restarting speaker to get back to stable state","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4095,"Ref":"types.json#/definitions/GracefulRestartTimeType","CollectionType":"","Column":"","Item":null,"GoName":"RestartTime","GoType":"GracefulRestartTimeType"}
		LongLivedRestartTime: InterfaceToLongLivedGracefulRestartTimeType(data["long_lived_restart_time"]),

		//{"Title":"","Description":"Extended Time (in seconds) taken by the restarting speaker after restart-time to get back to stable state","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":16777215,"Ref":"types.json#/definitions/LongLivedGracefulRestartTimeType","CollectionType":"","Column":"","Item":null,"GoName":"LongLivedRestartTime","GoType":"LongLivedGracefulRestartTimeType"}
		Enable: data["enable"].(bool),

		//{"Title":"","Description":"Enable/Disable knob for all GR parameters to take effect","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Enable","GoType":"bool"}
		EndOfRibTimeout: InterfaceToEndOfRibTimeType(data["end_of_rib_timeout"]),

		//{"Title":"","Description":"Maximum time (in seconds) to wait for EndOfRib reception/transmission","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4095,"Ref":"types.json#/definitions/EndOfRibTimeType","CollectionType":"","Column":"","Item":null,"GoName":"EndOfRibTimeout","GoType":"EndOfRibTimeType"}
		BGPHelperEnable: data["bgp_helper_enable"].(bool),

		//{"Title":"","Description":"Enable GR Helper mode for BGP peers in contrail-control","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"BGPHelperEnable","GoType":"bool"}
		XMPPHelperEnable: data["xmpp_helper_enable"].(bool),

		//{"Title":"","Description":"Enable GR Helper mode for XMPP peers (agents) in contrail-control","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"XMPPHelperEnable","GoType":"bool"}

	}
}

func InterfaceToGracefulRestartParametersTypeSlice(data interface{}) []*GracefulRestartParametersType {
	list := data.([]interface{})
	result := MakeGracefulRestartParametersTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToGracefulRestartParametersType(item))
	}
	return result
}

func MakeGracefulRestartParametersTypeSlice() []*GracefulRestartParametersType {
	return []*GracefulRestartParametersType{}
}
