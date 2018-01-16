package models

// Project

import "encoding/json"

// Project
type Project struct {
	ParentType   string         `json:"parent_type,omitempty"`
	FQName       []string       `json:"fq_name,omitempty"`
	IDPerms      *IdPermsType   `json:"id_perms,omitempty"`
	Annotations  *KeyValuePairs `json:"annotations,omitempty"`
	Perms2       *PermType2     `json:"perms2,omitempty"`
	VxlanRouting bool           `json:"vxlan_routing,omitempty"`
	AlarmEnable  bool           `json:"alarm_enable,omitempty"`
	ParentUUID   string         `json:"parent_uuid,omitempty"`
	DisplayName  string         `json:"display_name,omitempty"`
	Quota        *QuotaType     `json:"quota,omitempty"`
	UUID         string         `json:"uuid,omitempty"`

	AliasIPPoolRefs          []*ProjectAliasIPPoolRef          `json:"alias_ip_pool_refs,omitempty"`
	NamespaceRefs            []*ProjectNamespaceRef            `json:"namespace_refs,omitempty"`
	ApplicationPolicySetRefs []*ProjectApplicationPolicySetRef `json:"application_policy_set_refs,omitempty"`
	FloatingIPPoolRefs       []*ProjectFloatingIPPoolRef       `json:"floating_ip_pool_refs,omitempty"`

	AddressGroups              []*AddressGroup              `json:"address_groups,omitempty"`
	Alarms                     []*Alarm                     `json:"alarms,omitempty"`
	APIAccessLists             []*APIAccessList             `json:"api_access_lists,omitempty"`
	ApplicationPolicySets      []*ApplicationPolicySet      `json:"application_policy_sets,omitempty"`
	BGPAsAServices             []*BGPAsAService             `json:"bgp_as_a_services,omitempty"`
	BGPVPNs                    []*BGPVPN                    `json:"bgpvpns,omitempty"`
	FirewallPolicys            []*FirewallPolicy            `json:"firewall_policys,omitempty"`
	FirewallRules              []*FirewallRule              `json:"firewall_rules,omitempty"`
	InterfaceRouteTables       []*InterfaceRouteTable       `json:"interface_route_tables,omitempty"`
	LoadbalancerHealthmonitors []*LoadbalancerHealthmonitor `json:"loadbalancer_healthmonitors,omitempty"`
	LoadbalancerListeners      []*LoadbalancerListener      `json:"loadbalancer_listeners,omitempty"`
	LoadbalancerPools          []*LoadbalancerPool          `json:"loadbalancer_pools,omitempty"`
	Loadbalancers              []*Loadbalancer              `json:"loadbalancers,omitempty"`
	LogicalRouters             []*LogicalRouter             `json:"logical_routers,omitempty"`
	NetworkIpams               []*NetworkIpam               `json:"network_ipams,omitempty"`
	NetworkPolicys             []*NetworkPolicy             `json:"network_policys,omitempty"`
	QosConfigs                 []*QosConfig                 `json:"qos_configs,omitempty"`
	RouteAggregates            []*RouteAggregate            `json:"route_aggregates,omitempty"`
	RouteTables                []*RouteTable                `json:"route_tables,omitempty"`
	RoutingPolicys             []*RoutingPolicy             `json:"routing_policys,omitempty"`
	SecurityGroups             []*SecurityGroup             `json:"security_groups,omitempty"`
	SecurityLoggingObjects     []*SecurityLoggingObject     `json:"security_logging_objects,omitempty"`
	ServiceGroups              []*ServiceGroup              `json:"service_groups,omitempty"`
	ServiceHealthChecks        []*ServiceHealthCheck        `json:"service_health_checks,omitempty"`
	ServiceInstances           []*ServiceInstance           `json:"service_instances,omitempty"`
	Tags                       []*Tag                       `json:"tags,omitempty"`
	Users                      []*User                      `json:"users,omitempty"`
	VirtualIPs                 []*VirtualIP                 `json:"virtual_ips,omitempty"`
	VirtualMachineInterfaces   []*VirtualMachineInterface   `json:"virtual_machine_interfaces,omitempty"`
	VirtualNetworks            []*VirtualNetwork            `json:"virtual_networks,omitempty"`
}

// ProjectAliasIPPoolRef references each other
type ProjectAliasIPPoolRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// ProjectNamespaceRef references each other
type ProjectNamespaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *SubnetType
}

// ProjectApplicationPolicySetRef references each other
type ProjectApplicationPolicySetRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// ProjectFloatingIPPoolRef references each other
type ProjectFloatingIPPoolRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *Project) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeProject makes Project
func MakeProject() *Project {
	return &Project{
		//TODO(nati): Apply default
		Quota:        MakeQuotaType(),
		UUID:         "",
		ParentUUID:   "",
		DisplayName:  "",
		VxlanRouting: false,
		AlarmEnable:  false,
		ParentType:   "",
		FQName:       []string{},
		IDPerms:      MakeIdPermsType(),
		Annotations:  MakeKeyValuePairs(),
		Perms2:       MakePermType2(),
	}
}

// MakeProjectSlice() makes a slice of Project
func MakeProjectSlice() []*Project {
	return []*Project{}
}
