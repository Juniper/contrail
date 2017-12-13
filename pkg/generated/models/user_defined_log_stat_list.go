package models

// UserDefinedLogStatList

import "encoding/json"

// UserDefinedLogStatList
type UserDefinedLogStatList struct {
	Statlist []*UserDefinedLogStat `json:"statlist"`
}

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

		//{"type":"array","item":{"type":"object","properties":{"name":{"type":"string"},"pattern":{"type":"string"}}}}

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
