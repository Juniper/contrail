package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLinklocalServiceEntryType makes LinklocalServiceEntryType
// nolint
func MakeLinklocalServiceEntryType() *LinklocalServiceEntryType {
	return &LinklocalServiceEntryType{
		//TODO(nati): Apply default
		IPFabricServiceIP:      []string{},
		LinklocalServiceName:   "",
		LinklocalServiceIP:     "",
		IPFabricServicePort:    0,
		IPFabricDNSServiceName: "",
		LinklocalServicePort:   0,
	}
}

// MakeLinklocalServiceEntryType makes LinklocalServiceEntryType
// nolint
func InterfaceToLinklocalServiceEntryType(i interface{}) *LinklocalServiceEntryType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LinklocalServiceEntryType{
		//TODO(nati): Apply default
		IPFabricServiceIP:      common.InterfaceToStringList(m["ip_fabric_service_ip"]),
		LinklocalServiceName:   common.InterfaceToString(m["linklocal_service_name"]),
		LinklocalServiceIP:     common.InterfaceToString(m["linklocal_service_ip"]),
		IPFabricServicePort:    common.InterfaceToInt64(m["ip_fabric_service_port"]),
		IPFabricDNSServiceName: common.InterfaceToString(m["ip_fabric_DNS_service_name"]),
		LinklocalServicePort:   common.InterfaceToInt64(m["linklocal_service_port"]),
	}
}

// MakeLinklocalServiceEntryTypeSlice() makes a slice of LinklocalServiceEntryType
// nolint
func MakeLinklocalServiceEntryTypeSlice() []*LinklocalServiceEntryType {
	return []*LinklocalServiceEntryType{}
}

// InterfaceToLinklocalServiceEntryTypeSlice() makes a slice of LinklocalServiceEntryType
// nolint
func InterfaceToLinklocalServiceEntryTypeSlice(i interface{}) []*LinklocalServiceEntryType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LinklocalServiceEntryType{}
	for _, item := range list {
		result = append(result, InterfaceToLinklocalServiceEntryType(item))
	}
	return result
}
