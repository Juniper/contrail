package models

// Project

import "encoding/json"

// Project
type Project struct {
	AlarmEnable  bool           `json:"alarm_enable"`
	Annotations  *KeyValuePairs `json:"annotations"`
	ParentUUID   string         `json:"parent_uuid"`
	UUID         string         `json:"uuid"`
	ParentType   string         `json:"parent_type"`
	VxlanRouting bool           `json:"vxlan_routing"`
	Quota        *QuotaType     `json:"quota"`
	FQName       []string       `json:"fq_name"`
	IDPerms      *IdPermsType   `json:"id_perms"`
	DisplayName  string         `json:"display_name"`
	Perms2       *PermType2     `json:"perms2"`

	FloatingIPPoolRefs       []*ProjectFloatingIPPoolRef       `json:"floating_ip_pool_refs"`
	AliasIPPoolRefs          []*ProjectAliasIPPoolRef          `json:"alias_ip_pool_refs"`
	NamespaceRefs            []*ProjectNamespaceRef            `json:"namespace_refs"`
	ApplicationPolicySetRefs []*ProjectApplicationPolicySetRef `json:"application_policy_set_refs"`

	AddressGroups              []*AddressGroup              `json:"address_groups"`
	Alarms                     []*Alarm                     `json:"alarms"`
	APIAccessLists             []*APIAccessList             `json:"api_access_lists"`
	ApplicationPolicySets      []*ApplicationPolicySet      `json:"application_policy_sets"`
	BGPAsAServices             []*BGPAsAService             `json:"bgp_as_a_services"`
	BGPVPNs                    []*BGPVPN                    `json:"bgpvpns"`
	FirewallPolicys            []*FirewallPolicy            `json:"firewall_policys"`
	FirewallRules              []*FirewallRule              `json:"firewall_rules"`
	InterfaceRouteTables       []*InterfaceRouteTable       `json:"interface_route_tables"`
	LoadbalancerHealthmonitors []*LoadbalancerHealthmonitor `json:"loadbalancer_healthmonitors"`
	LoadbalancerListeners      []*LoadbalancerListener      `json:"loadbalancer_listeners"`
	LoadbalancerPools          []*LoadbalancerPool          `json:"loadbalancer_pools"`
	Loadbalancers              []*Loadbalancer              `json:"loadbalancers"`
	LogicalRouters             []*LogicalRouter             `json:"logical_routers"`
	NetworkIpams               []*NetworkIpam               `json:"network_ipams"`
	NetworkPolicys             []*NetworkPolicy             `json:"network_policys"`
	QosConfigs                 []*QosConfig                 `json:"qos_configs"`
	RouteAggregates            []*RouteAggregate            `json:"route_aggregates"`
	RouteTables                []*RouteTable                `json:"route_tables"`
	RoutingPolicys             []*RoutingPolicy             `json:"routing_policys"`
	SecurityGroups             []*SecurityGroup             `json:"security_groups"`
	SecurityLoggingObjects     []*SecurityLoggingObject     `json:"security_logging_objects"`
	ServiceGroups              []*ServiceGroup              `json:"service_groups"`
	ServiceHealthChecks        []*ServiceHealthCheck        `json:"service_health_checks"`
	ServiceInstances           []*ServiceInstance           `json:"service_instances"`
	Tags                       []*Tag                       `json:"tags"`
	Users                      []*User                      `json:"users"`
	VirtualIPs                 []*VirtualIP                 `json:"virtual_ips"`
	VirtualMachineInterfaces   []*VirtualMachineInterface   `json:"virtual_machine_interfaces"`
	VirtualNetworks            []*VirtualNetwork            `json:"virtual_networks"`
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

// ProjectAliasIPPoolRef references each other
type ProjectAliasIPPoolRef struct {
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
		DisplayName:  "",
		Perms2:       MakePermType2(),
		UUID:         "",
		ParentType:   "",
		VxlanRouting: false,
		Quota:        MakeQuotaType(),
		FQName:       []string{},
		IDPerms:      MakeIdPermsType(),
		AlarmEnable:  false,
		Annotations:  MakeKeyValuePairs(),
		ParentUUID:   "",
	}
}

// MakeProjectSlice() makes a slice of Project
func MakeProjectSlice() []*Project {
	return []*Project{}
}
