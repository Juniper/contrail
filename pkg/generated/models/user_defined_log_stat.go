package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeUserDefinedLogStat makes UserDefinedLogStat
func MakeUserDefinedLogStat() *UserDefinedLogStat {
	return &UserDefinedLogStat{
		//TODO(nati): Apply default
		Pattern: "",
		Name:    "",
	}
}

// MakeUserDefinedLogStat makes UserDefinedLogStat
func InterfaceToUserDefinedLogStat(i interface{}) *UserDefinedLogStat {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &UserDefinedLogStat{
		//TODO(nati): Apply default
		Pattern: schema.InterfaceToString(m["pattern"]),
		Name:    schema.InterfaceToString(m["name"]),
	}
}

// MakeUserDefinedLogStatSlice() makes a slice of UserDefinedLogStat
func MakeUserDefinedLogStatSlice() []*UserDefinedLogStat {
	return []*UserDefinedLogStat{}
}

// InterfaceToUserDefinedLogStatSlice() makes a slice of UserDefinedLogStat
func InterfaceToUserDefinedLogStatSlice(i interface{}) []*UserDefinedLogStat {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*UserDefinedLogStat{}
	for _, item := range list {
		result = append(result, InterfaceToUserDefinedLogStat(item))
	}
	return result
}
