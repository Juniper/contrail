package models

// RouteAggregate

import "encoding/json"

// RouteAggregate
type RouteAggregate struct {
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`

	ServiceInstanceRefs []*RouteAggregateServiceInstanceRef `json:"service_instance_refs,omitempty"`
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
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
	}
}

// MakeRouteAggregateSlice() makes a slice of RouteAggregate
func MakeRouteAggregateSlice() []*RouteAggregate {
	return []*RouteAggregate{}
}
