package models

// UserDefinedLogStat

import "encoding/json"

// UserDefinedLogStat
type UserDefinedLogStat struct {
	Pattern string `json:"pattern"`
	Name    string `json:"name"`
}

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

		//{"description":"Perl type regular expression pattern to match","type":"string"}
		Name: data["name"].(string),

		//{"description":"Name of the stat","type":"string"}

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
