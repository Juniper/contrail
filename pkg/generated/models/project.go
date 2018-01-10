package models

// Project

import "encoding/json"

// Project
type Project struct {
	DisplayName  string         `json:"display_name"`
	Perms2       *PermType2     `json:"perms2"`
	UUID         string         `json:"uuid"`
	IDPerms      *IdPermsType   `json:"id_perms"`
	VxlanRouting bool           `json:"vxlan_routing"`
	AlarmEnable  bool           `json:"alarm_enable"`
	Quota        *QuotaType     `json:"quota"`
	Annotations  *KeyValuePairs `json:"annotations"`
	ParentUUID   string         `json:"parent_uuid"`
	ParentType   string         `json:"parent_type"`
	FQName       []string       `json:"fq_name"`

	NamespaceRefs            []*ProjectNamespaceRef            `json:"namespace_refs"`
	ApplicationPolicySetRefs []*ProjectApplicationPolicySetRef `json:"application_policy_set_refs"`
	FloatingIPPoolRefs       []*ProjectFloatingIPPoolRef       `json:"floating_ip_pool_refs"`
	AliasIPPoolRefs          []*ProjectAliasIPPoolRef          `json:"alias_ip_pool_refs"`

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
		IDPerms:      MakeIdPermsType(),
		VxlanRouting: false,
		AlarmEnable:  false,
		Quota:        MakeQuotaType(),
		Annotations:  MakeKeyValuePairs(),
		ParentUUID:   "",
		ParentType:   "",
		FQName:       []string{},
		DisplayName:  "",
		Perms2:       MakePermType2(),
		UUID:         "",
	}
}

// InterfaceToProject makes Project from interface
func InterfaceToProject(iData interface{}) *Project {
	data := iData.(map[string]interface{})
	return &Project{
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		VxlanRouting: data["vxlan_routing"].(bool),

		//{"description":"When this knob is enabled for a project, an internal system VN (VN-Int) is created for every logical router in the project.","type":"boolean"}
		AlarmEnable: data["alarm_enable"].(bool),

		//{"description":"Flag to enable/disable alarms configured under global-system-config. True, if not set.","type":"boolean"}
		Quota: InterfaceToQuotaType(data["quota"]),

		//{"description":"Max instances limits for various objects under project.","type":"object","properties":{"access_control_list":{"type":"integer"},"bgp_router":{"type":"integer"},"defaults":{"type":"integer"},"floating_ip":{"type":"integer"},"floating_ip_pool":{"type":"integer"},"global_vrouter_config":{"type":"integer"},"instance_ip":{"type":"integer"},"loadbalancer_healthmonitor":{"type":"integer"},"loadbalancer_member":{"type":"integer"},"loadbalancer_pool":{"type":"integer"},"logical_router":{"type":"integer"},"network_ipam":{"type":"integer"},"network_policy":{"type":"integer"},"route_table":{"type":"integer"},"security_group":{"type":"integer"},"security_group_rule":{"type":"integer"},"security_logging_object":{"type":"integer"},"service_instance":{"type":"integer"},"service_template":{"type":"integer"},"subnet":{"type":"integer"},"virtual_DNS":{"type":"integer"},"virtual_DNS_record":{"type":"integer"},"virtual_ip":{"type":"integer"},"virtual_machine_interface":{"type":"integer"},"virtual_network":{"type":"integer"},"virtual_router":{"type":"integer"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToProjectSlice makes a slice of Project from interface
func InterfaceToProjectSlice(data interface{}) []*Project {
	list := data.([]interface{})
	result := MakeProjectSlice()
	for _, item := range list {
		result = append(result, InterfaceToProject(item))
	}
	return result
}

// MakeProjectSlice() makes a slice of Project
func MakeProjectSlice() []*Project {
	return []*Project{}
}
