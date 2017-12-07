package models

// UserDefinedLogStat

import "encoding/json"

// UserDefinedLogStat
type UserDefinedLogStat struct {
	Pattern string `json:"pattern"`
	Name    string `json:"name"`
}

//  parents relation object

// String returns json representation of the object
func (model *UserDefinedLogStat) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeUserDefinedLogStat makes UserDefinedLogStat
func MakeUserDefinedLogStat() *UserDefinedLogStat {
	return &UserDefinedLogStat{
		//TODO(nati): Apply default
		Pattern: "",
		Name:    "",
	}
}

// InterfaceToUserDefinedLogStat makes UserDefinedLogStat from interface
func InterfaceToUserDefinedLogStat(iData interface{}) *UserDefinedLogStat {
	data := iData.(map[string]interface{})
	return &UserDefinedLogStat{
		Pattern: data["pattern"].(string),

		//{"Title":"","Description":"Perl type regular expression pattern to match","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Pattern","GoType":"string","GoPremitive":true}
		Name: data["name"].(string),

		//{"Title":"","Description":"Name of the stat","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToUserDefinedLogStatSlice makes a slice of UserDefinedLogStat from interface
func InterfaceToUserDefinedLogStatSlice(data interface{}) []*UserDefinedLogStat {
	list := data.([]interface{})
	result := MakeUserDefinedLogStatSlice()
	for _, item := range list {
		result = append(result, InterfaceToUserDefinedLogStat(item))
	}
	return result
}

// MakeUserDefinedLogStatSlice() makes a slice of UserDefinedLogStat
func MakeUserDefinedLogStatSlice() []*UserDefinedLogStat {
	return []*UserDefinedLogStat{}
}
