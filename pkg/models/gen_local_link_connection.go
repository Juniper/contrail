package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLocalLinkConnection makes LocalLinkConnection
// nolint
func MakeLocalLinkConnection() *LocalLinkConnection {
	return &LocalLinkConnection{
		//TODO(nati): Apply default
		SwitchID:   "",
		PortID:     "",
		SwitchInfo: "",
	}
}

// MakeLocalLinkConnection makes LocalLinkConnection
// nolint
func InterfaceToLocalLinkConnection(i interface{}) *LocalLinkConnection {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LocalLinkConnection{
		//TODO(nati): Apply default
		SwitchID:   common.InterfaceToString(m["switch_id"]),
		PortID:     common.InterfaceToString(m["port_id"]),
		SwitchInfo: common.InterfaceToString(m["switch_info"]),
	}
}

// MakeLocalLinkConnectionSlice() makes a slice of LocalLinkConnection
// nolint
func MakeLocalLinkConnectionSlice() []*LocalLinkConnection {
	return []*LocalLinkConnection{}
}

// InterfaceToLocalLinkConnectionSlice() makes a slice of LocalLinkConnection
// nolint
func InterfaceToLocalLinkConnectionSlice(i interface{}) []*LocalLinkConnection {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LocalLinkConnection{}
	for _, item := range list {
		result = append(result, InterfaceToLocalLinkConnection(item))
	}
	return result
}
