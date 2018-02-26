package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeLocalLinkConnection makes LocalLinkConnection
func MakeLocalLinkConnection() *LocalLinkConnection {
	return &LocalLinkConnection{
		//TODO(nati): Apply default
		SwitchID:   "",
		PortID:     "",
		SwitchInfo: "",
	}
}

// MakeLocalLinkConnection makes LocalLinkConnection
func InterfaceToLocalLinkConnection(i interface{}) *LocalLinkConnection {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LocalLinkConnection{
		//TODO(nati): Apply default
		SwitchID:   schema.InterfaceToString(m["switch_id"]),
		PortID:     schema.InterfaceToString(m["port_id"]),
		SwitchInfo: schema.InterfaceToString(m["switch_info"]),
	}
}

// MakeLocalLinkConnectionSlice() makes a slice of LocalLinkConnection
func MakeLocalLinkConnectionSlice() []*LocalLinkConnection {
	return []*LocalLinkConnection{}
}

// InterfaceToLocalLinkConnectionSlice() makes a slice of LocalLinkConnection
func InterfaceToLocalLinkConnectionSlice(i interface{}) []*LocalLinkConnection {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LocalLinkConnection{}
	for _, item := range list {
		result = append(result, InterfaceToLocalLinkConnection(item))
	}
	return result
}
