package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeUserDefinedLogStat makes UserDefinedLogStat
// nolint
func MakeUserDefinedLogStat() *UserDefinedLogStat {
	return &UserDefinedLogStat{
		//TODO(nati): Apply default
		Pattern: "",
		Name:    "",
	}
}

// MakeUserDefinedLogStat makes UserDefinedLogStat
// nolint
func InterfaceToUserDefinedLogStat(i interface{}) *UserDefinedLogStat {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &UserDefinedLogStat{
		//TODO(nati): Apply default
		Pattern: common.InterfaceToString(m["pattern"]),
		Name:    common.InterfaceToString(m["name"]),
	}
}

// MakeUserDefinedLogStatSlice() makes a slice of UserDefinedLogStat
// nolint
func MakeUserDefinedLogStatSlice() []*UserDefinedLogStat {
	return []*UserDefinedLogStat{}
}

// InterfaceToUserDefinedLogStatSlice() makes a slice of UserDefinedLogStat
// nolint
func InterfaceToUserDefinedLogStatSlice(i interface{}) []*UserDefinedLogStat {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*UserDefinedLogStat{}
	for _, item := range list {
		result = append(result, InterfaceToUserDefinedLogStat(item))
	}
	return result
}
