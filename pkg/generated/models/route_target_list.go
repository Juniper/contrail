package models

// RouteTargetList

import "encoding/json"

type RouteTargetList struct {
	RouteTarget []string `json:"route_target"`
}

func (model *RouteTargetList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeRouteTargetList() *RouteTargetList {
	return &RouteTargetList{
		//TODO(nati): Apply default
		RouteTarget: []string{},
	}
}

func InterfaceToRouteTargetList(iData interface{}) *RouteTargetList {
	data := iData.(map[string]interface{})
	return &RouteTargetList{
		RouteTarget: data["route_target"].([]string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RouteTarget","GoType":"string"},"GoName":"RouteTarget","GoType":"[]string"}

	}
}

func InterfaceToRouteTargetListSlice(data interface{}) []*RouteTargetList {
	list := data.([]interface{})
	result := MakeRouteTargetListSlice()
	for _, item := range list {
		result = append(result, InterfaceToRouteTargetList(item))
	}
	return result
}

func MakeRouteTargetListSlice() []*RouteTargetList {
	return []*RouteTargetList{}
}
