package models

// RouteAggregate

import "encoding/json"

// RouteAggregate
type RouteAggregate struct {
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`

	ServiceInstanceRefs []*RouteAggregateServiceInstanceRef `json:"service_instance_refs"`
}

// RouteAggregateServiceInstanceRef references each other
type RouteAggregateServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceInterfaceTag
}

// String returns json representation of the object
func (model *RouteAggregate) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteAggregate makes RouteAggregate
func MakeRouteAggregate() *RouteAggregate {
	return &RouteAggregate{
		//TODO(nati): Apply default
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
	}
}

// MakeRouteAggregateSlice() makes a slice of RouteAggregate
func MakeRouteAggregateSlice() []*RouteAggregate {
	return []*RouteAggregate{}
}
