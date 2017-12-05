package models

// ServiceInterfaceTag

import "encoding/json"

type ServiceInterfaceTag struct {
	InterfaceType ServiceInterfaceType `json:"interface_type"`
}

func (model *ServiceInterfaceTag) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeServiceInterfaceTag() *ServiceInterfaceTag {
	return &ServiceInterfaceTag{
		//TODO(nati): Apply default
		InterfaceType: MakeServiceInterfaceType(),
	}
}

func InterfaceToServiceInterfaceTag(iData interface{}) *ServiceInterfaceTag {
	data := iData.(map[string]interface{})
	return &ServiceInterfaceTag{
		InterfaceType: InterfaceToServiceInterfaceType(data["interface_type"]),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceInterfaceType","CollectionType":"","Column":"interface_type","Item":null,"GoName":"InterfaceType","GoType":"ServiceInterfaceType"}

	}
}

func InterfaceToServiceInterfaceTagSlice(data interface{}) []*ServiceInterfaceTag {
	list := data.([]interface{})
	result := MakeServiceInterfaceTagSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceInterfaceTag(item))
	}
	return result
}

func MakeServiceInterfaceTagSlice() []*ServiceInterfaceTag {
	return []*ServiceInterfaceTag{}
}
