package models

// UserDefinedLogStatList

import "encoding/json"

type UserDefinedLogStatList struct {
	Statlist []*UserDefinedLogStat `json:"statlist"`
}

func (model *UserDefinedLogStatList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeUserDefinedLogStatList() *UserDefinedLogStatList {
	return &UserDefinedLogStatList{
		//TODO(nati): Apply default

		Statlist: MakeUserDefinedLogStatSlice(),
	}
}

func InterfaceToUserDefinedLogStatList(iData interface{}) *UserDefinedLogStatList {
	data := iData.(map[string]interface{})
	return &UserDefinedLogStatList{

		Statlist: InterfaceToUserDefinedLogStatSlice(data["statlist"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string"},"pattern":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Pattern","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UserDefinedLogStat","CollectionType":"","Column":"","Item":null,"GoName":"Statlist","GoType":"UserDefinedLogStat"},"GoName":"Statlist","GoType":"[]*UserDefinedLogStat"}

	}
}

func InterfaceToUserDefinedLogStatListSlice(data interface{}) []*UserDefinedLogStatList {
	list := data.([]interface{})
	result := MakeUserDefinedLogStatListSlice()
	for _, item := range list {
		result = append(result, InterfaceToUserDefinedLogStatList(item))
	}
	return result
}

func MakeUserDefinedLogStatListSlice() []*UserDefinedLogStatList {
	return []*UserDefinedLogStatList{}
}
