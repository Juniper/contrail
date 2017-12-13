package models

// ControlTrafficDscpType

import "encoding/json"

// ControlTrafficDscpType
type ControlTrafficDscpType struct {
	Analytics DscpValueType `json:"analytics"`
	DNS       DscpValueType `json:"dns"`
	Control   DscpValueType `json:"control"`
}

// String returns json representation of the object
func (model *ControlTrafficDscpType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeControlTrafficDscpType makes ControlTrafficDscpType
func MakeControlTrafficDscpType() *ControlTrafficDscpType {
	return &ControlTrafficDscpType{
		//TODO(nati): Apply default
		Control:   MakeDscpValueType(),
		Analytics: MakeDscpValueType(),
		DNS:       MakeDscpValueType(),
	}
}

// InterfaceToControlTrafficDscpType makes ControlTrafficDscpType from interface
func InterfaceToControlTrafficDscpType(iData interface{}) *ControlTrafficDscpType {
	data := iData.(map[string]interface{})
	return &ControlTrafficDscpType{
		Analytics: InterfaceToDscpValueType(data["analytics"]),

		//{"description":"DSCP value for traffic towards analytics","type":"integer","minimum":0,"maximum":63}
		DNS: InterfaceToDscpValueType(data["dns"]),

		//{"description":"DSCP value for DNS traffic","type":"integer","minimum":0,"maximum":63}
		Control: InterfaceToDscpValueType(data["control"]),

		//{"description":"DSCP value for control protocols traffic","type":"integer","minimum":0,"maximum":63}

	}
}

// InterfaceToControlTrafficDscpTypeSlice makes a slice of ControlTrafficDscpType from interface
func InterfaceToControlTrafficDscpTypeSlice(data interface{}) []*ControlTrafficDscpType {
	list := data.([]interface{})
	result := MakeControlTrafficDscpTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToControlTrafficDscpType(item))
	}
	return result
}

// MakeControlTrafficDscpTypeSlice() makes a slice of ControlTrafficDscpType
func MakeControlTrafficDscpTypeSlice() []*ControlTrafficDscpType {
	return []*ControlTrafficDscpType{}
}
