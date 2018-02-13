package models

// RouteType

import "encoding/json"

// RouteType
//proteus:generate
type RouteType struct {
	Prefix              string               `json:"prefix,omitempty"`
	NextHop             string               `json:"next_hop,omitempty"`
	CommunityAttributes *CommunityAttributes `json:"community_attributes,omitempty"`
	NextHopType         RouteNextHopType     `json:"next_hop_type,omitempty"`
}

// String returns json representation of the object
func (model *RouteType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteType makes RouteType
func MakeRouteType() *RouteType {
	return &RouteType{
		//TODO(nati): Apply default
		Prefix:              "",
		NextHop:             "",
		CommunityAttributes: MakeCommunityAttributes(),
		NextHopType:         MakeRouteNextHopType(),
	}
}

// MakeRouteTypeSlice() makes a slice of RouteType
func MakeRouteTypeSlice() []*RouteType {
	return []*RouteType{}
}
