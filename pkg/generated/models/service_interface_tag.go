package models

// ServiceInterfaceTag

import "encoding/json"

// ServiceInterfaceTag
type ServiceInterfaceTag struct {
	InterfaceType ServiceInterfaceType `json:"interface_type"`
}

// String returns json representation of the object
func (model *ServiceInterfaceTag) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceInterfaceTag makes ServiceInterfaceTag
func MakeServiceInterfaceTag() *ServiceInterfaceTag {
	return &ServiceInterfaceTag{
		//TODO(nati): Apply default
		InterfaceType: MakeServiceInterfaceType(),
	}
}

// InterfaceToServiceInterfaceTag makes ServiceInterfaceTag from interface
func InterfaceToServiceInterfaceTag(iData interface{}) *ServiceInterfaceTag {
	data := iData.(map[string]interface{})
	return &ServiceInterfaceTag{
		InterfaceType: InterfaceToServiceInterfaceType(data["interface_type"]),

		//{"type":"string"}

	}
}

// InterfaceToServiceInterfaceTagSlice makes a slice of ServiceInterfaceTag from interface
func InterfaceToServiceInterfaceTagSlice(data interface{}) []*ServiceInterfaceTag {
	list := data.([]interface{})
	result := MakeServiceInterfaceTagSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceInterfaceTag(item))
	}
	return result
}

// MakeServiceInterfaceTagSlice() makes a slice of ServiceInterfaceTag
func MakeServiceInterfaceTagSlice() []*ServiceInterfaceTag {
	return []*ServiceInterfaceTag{}
}
