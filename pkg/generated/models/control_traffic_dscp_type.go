package models

// ControlTrafficDscpType

import "encoding/json"

// ControlTrafficDscpType
type ControlTrafficDscpType struct {
	Control   DscpValueType `json:"control"`
	Analytics DscpValueType `json:"analytics"`
	DNS       DscpValueType `json:"dns"`
}

//  parents relation object

// String returns json representation of the object
func (model *ControlTrafficDscpType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeControlTrafficDscpType makes ControlTrafficDscpType
func MakeControlTrafficDscpType() *ControlTrafficDscpType {
	return &ControlTrafficDscpType{
		//TODO(nati): Apply default
		DNS:       MakeDscpValueType(),
		Control:   MakeDscpValueType(),
		Analytics: MakeDscpValueType(),
	}
}

// InterfaceToControlTrafficDscpType makes ControlTrafficDscpType from interface
func InterfaceToControlTrafficDscpType(iData interface{}) *ControlTrafficDscpType {
	data := iData.(map[string]interface{})
	return &ControlTrafficDscpType{
		Control: InterfaceToDscpValueType(data["control"]),

		//{"Title":"","Description":"DSCP value for control protocols traffic","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":63,"Ref":"types.json#/definitions/DscpValueType","CollectionType":"","Column":"","Item":null,"GoName":"Control","GoType":"DscpValueType","GoPremitive":false}
		Analytics: InterfaceToDscpValueType(data["analytics"]),

		//{"Title":"","Description":"DSCP value for traffic towards analytics","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":63,"Ref":"types.json#/definitions/DscpValueType","CollectionType":"","Column":"","Item":null,"GoName":"Analytics","GoType":"DscpValueType","GoPremitive":false}
		DNS: InterfaceToDscpValueType(data["dns"]),

		//{"Title":"","Description":"DSCP value for DNS traffic","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":63,"Ref":"types.json#/definitions/DscpValueType","CollectionType":"","Column":"","Item":null,"GoName":"DNS","GoType":"DscpValueType","GoPremitive":false}

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
