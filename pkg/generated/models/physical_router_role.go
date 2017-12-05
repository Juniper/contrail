package models

// PhysicalRouterRole

type PhysicalRouterRole string

func MakePhysicalRouterRole() PhysicalRouterRole {
	var data PhysicalRouterRole
	return data
}

func InterfaceToPhysicalRouterRole(data interface{}) PhysicalRouterRole {
	return data.(PhysicalRouterRole)
}

func InterfaceToPhysicalRouterRoleSlice(data interface{}) []PhysicalRouterRole {
	list := data.([]interface{})
	result := MakePhysicalRouterRoleSlice()
	for _, item := range list {
		result = append(result, InterfaceToPhysicalRouterRole(item))
	}
	return result
}

func MakePhysicalRouterRoleSlice() []PhysicalRouterRole {
	return []PhysicalRouterRole{}
}
