package models

// GracefulRestartParametersType

import "encoding/json"

// GracefulRestartParametersType
type GracefulRestartParametersType struct {
	Enable               bool                             `json:"enable"`
	EndOfRibTimeout      EndOfRibTimeType                 `json:"end_of_rib_timeout"`
	BGPHelperEnable      bool                             `json:"bgp_helper_enable"`
	XMPPHelperEnable     bool                             `json:"xmpp_helper_enable"`
	RestartTime          GracefulRestartTimeType          `json:"restart_time"`
	LongLivedRestartTime LongLivedGracefulRestartTimeType `json:"long_lived_restart_time"`
}

//  parents relation object

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
		XMPPHelperEnable: data["xmpp_helper_enable"].(bool),

		//{"Title":"","Description":"Enable GR Helper mode for XMPP peers (agents) in contrail-control","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"XMPPHelperEnable","GoType":"bool","GoPremitive":true}
		RestartTime: InterfaceToGracefulRestartTimeType(data["restart_time"]),

		//{"Title":"","Description":"Time (in seconds) taken by the restarting speaker to get back to stable state","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4095,"Ref":"types.json#/definitions/GracefulRestartTimeType","CollectionType":"","Column":"","Item":null,"GoName":"RestartTime","GoType":"GracefulRestartTimeType","GoPremitive":false}
		LongLivedRestartTime: InterfaceToLongLivedGracefulRestartTimeType(data["long_lived_restart_time"]),

		//{"Title":"","Description":"Extended Time (in seconds) taken by the restarting speaker after restart-time to get back to stable state","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":16777215,"Ref":"types.json#/definitions/LongLivedGracefulRestartTimeType","CollectionType":"","Column":"","Item":null,"GoName":"LongLivedRestartTime","GoType":"LongLivedGracefulRestartTimeType","GoPremitive":false}
		Enable: data["enable"].(bool),

		//{"Title":"","Description":"Enable/Disable knob for all GR parameters to take effect","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true}
		EndOfRibTimeout: InterfaceToEndOfRibTimeType(data["end_of_rib_timeout"]),

		//{"Title":"","Description":"Maximum time (in seconds) to wait for EndOfRib reception/transmission","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":4095,"Ref":"types.json#/definitions/EndOfRibTimeType","CollectionType":"","Column":"","Item":null,"GoName":"EndOfRibTimeout","GoType":"EndOfRibTimeType","GoPremitive":false}
		BGPHelperEnable: data["bgp_helper_enable"].(bool),

		//{"Title":"","Description":"Enable GR Helper mode for BGP peers in contrail-control","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"BGPHelperEnable","GoType":"bool","GoPremitive":true}

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
