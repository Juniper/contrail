package models

// PhysicalRouterRole

type PhysicalRouterRole string

// MakePhysicalRouterRole makes PhysicalRouterRole
func MakePhysicalRouterRole() PhysicalRouterRole {
	var data PhysicalRouterRole
	return data
}

// InterfaceToPhysicalRouterRole makes PhysicalRouterRole from interface
func InterfaceToPhysicalRouterRole(data interface{}) PhysicalRouterRole {
	return data.(PhysicalRouterRole)
}

// InterfaceToPhysicalRouterRoleSlice makes a slice of PhysicalRouterRole from interface
func InterfaceToPhysicalRouterRoleSlice(data interface{}) []PhysicalRouterRole {
	list := data.([]interface{})
	result := MakePhysicalRouterRoleSlice()
	for _, item := range list {
		result = append(result, InterfaceToPhysicalRouterRole(item))
	}
	return result
}

// MakePhysicalRouterRoleSlice() makes a slice of PhysicalRouterRole
func MakePhysicalRouterRoleSlice() []PhysicalRouterRole {
	return []PhysicalRouterRole{}
}
