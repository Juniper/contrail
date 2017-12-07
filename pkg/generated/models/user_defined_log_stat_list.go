package models

// UserDefinedLogStatList

import "encoding/json"

// UserDefinedLogStatList
type UserDefinedLogStatList struct {
	Statlist []*UserDefinedLogStat `json:"statlist"`
}

//  parents relation object

// String returns json representation of the object
func (model *UserDefinedLogStatList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeUserDefinedLogStatList makes UserDefinedLogStatList
func MakeUserDefinedLogStatList() *UserDefinedLogStatList {
	return &UserDefinedLogStatList{
		//TODO(nati): Apply default

		Statlist: MakeUserDefinedLogStatSlice(),
	}
}

// InterfaceToUserDefinedLogStatList makes UserDefinedLogStatList from interface
func InterfaceToUserDefinedLogStatList(iData interface{}) *UserDefinedLogStatList {
	data := iData.(map[string]interface{})
	return &UserDefinedLogStatList{

		Statlist: InterfaceToUserDefinedLogStatSlice(data["statlist"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string","GoPremitive":true},"pattern":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Pattern","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UserDefinedLogStat","CollectionType":"","Column":"","Item":null,"GoName":"Statlist","GoType":"UserDefinedLogStat","GoPremitive":false},"GoName":"Statlist","GoType":"[]*UserDefinedLogStat","GoPremitive":true}

	}
}

// InterfaceToUserDefinedLogStatListSlice makes a slice of UserDefinedLogStatList from interface
func InterfaceToUserDefinedLogStatListSlice(data interface{}) []*UserDefinedLogStatList {
	list := data.([]interface{})
	result := MakeUserDefinedLogStatListSlice()
	for _, item := range list {
		result = append(result, InterfaceToUserDefinedLogStatList(item))
	}
	return result
}

// MakeUserDefinedLogStatListSlice() makes a slice of UserDefinedLogStatList
func MakeUserDefinedLogStatListSlice() []*UserDefinedLogStatList {
	return []*UserDefinedLogStatList{}
}
