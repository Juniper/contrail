
package models
// Project



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propProject_fq_name int = iota
    propProject_display_name int = iota
    propProject_annotations int = iota
    propProject_uuid int = iota
    propProject_id_perms int = iota
    propProject_perms2 int = iota
    propProject_parent_uuid int = iota
    propProject_parent_type int = iota
    propProject_vxlan_routing int = iota
    propProject_alarm_enable int = iota
    propProject_quota int = iota
)

// Project 
type Project struct {

    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    VxlanRouting bool `json:"vxlan_routing"`
    AlarmEnable bool `json:"alarm_enable"`
    Quota *QuotaType `json:"quota,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`

    AliasIPPoolRefs []*ProjectAliasIPPoolRef `json:"alias_ip_pool_refs,omitempty"`
    NamespaceRefs []*ProjectNamespaceRef `json:"namespace_refs,omitempty"`
    ApplicationPolicySetRefs []*ProjectApplicationPolicySetRef `json:"application_policy_set_refs,omitempty"`
    FloatingIPPoolRefs []*ProjectFloatingIPPoolRef `json:"floating_ip_pool_refs,omitempty"`

    AddressGroups []*AddressGroup `json:"address_groups,omitempty"`
    Alarms []*Alarm `json:"alarms,omitempty"`
    APIAccessLists []*APIAccessList `json:"api_access_lists,omitempty"`
    ApplicationPolicySets []*ApplicationPolicySet `json:"application_policy_sets,omitempty"`
    BGPAsAServices []*BGPAsAService `json:"bgp_as_a_services,omitempty"`
    BGPVPNs []*BGPVPN `json:"bgpvpns,omitempty"`
    FirewallPolicys []*FirewallPolicy `json:"firewall_policys,omitempty"`
    FirewallRules []*FirewallRule `json:"firewall_rules,omitempty"`
    InterfaceRouteTables []*InterfaceRouteTable `json:"interface_route_tables,omitempty"`
    LoadbalancerHealthmonitors []*LoadbalancerHealthmonitor `json:"loadbalancer_healthmonitors,omitempty"`
    LoadbalancerListeners []*LoadbalancerListener `json:"loadbalancer_listeners,omitempty"`
    LoadbalancerPools []*LoadbalancerPool `json:"loadbalancer_pools,omitempty"`
    Loadbalancers []*Loadbalancer `json:"loadbalancers,omitempty"`
    LogicalRouters []*LogicalRouter `json:"logical_routers,omitempty"`
    NetworkIpams []*NetworkIpam `json:"network_ipams,omitempty"`
    NetworkPolicys []*NetworkPolicy `json:"network_policys,omitempty"`
    QosConfigs []*QosConfig `json:"qos_configs,omitempty"`
    RouteAggregates []*RouteAggregate `json:"route_aggregates,omitempty"`
    RouteTables []*RouteTable `json:"route_tables,omitempty"`
    RoutingPolicys []*RoutingPolicy `json:"routing_policys,omitempty"`
    SecurityGroups []*SecurityGroup `json:"security_groups,omitempty"`
    SecurityLoggingObjects []*SecurityLoggingObject `json:"security_logging_objects,omitempty"`
    ServiceGroups []*ServiceGroup `json:"service_groups,omitempty"`
    ServiceHealthChecks []*ServiceHealthCheck `json:"service_health_checks,omitempty"`
    ServiceInstances []*ServiceInstance `json:"service_instances,omitempty"`
    Tags []*Tag `json:"tags,omitempty"`
    Users []*User `json:"users,omitempty"`
    VirtualIPs []*VirtualIP `json:"virtual_ips,omitempty"`
    VirtualMachineInterfaces []*VirtualMachineInterface `json:"virtual_machine_interfaces,omitempty"`
    VirtualNetworks []*VirtualNetwork `json:"virtual_networks,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// ProjectApplicationPolicySetRef references each other
type ProjectApplicationPolicySetRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// ProjectFloatingIPPoolRef references each other
type ProjectFloatingIPPoolRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// ProjectAliasIPPoolRef references each other
type ProjectAliasIPPoolRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// ProjectNamespaceRef references each other
type ProjectNamespaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *SubnetType
    
}


// String returns json representation of the object
func (model *Project) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeProject makes Project
func MakeProject() *Project{
    return &Project{
    //TODO(nati): Apply default
    VxlanRouting: false,
        AlarmEnable: false,
        Quota: MakeQuotaType(),
        IDPerms: MakeIdPermsType(),
        Perms2: MakePermType2(),
        ParentUUID: "",
        ParentType: "",
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeProjectSlice makes a slice of Project
func MakeProjectSlice() []*Project {
    return []*Project{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Project) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[domain:0xc4202e4780])
    fqn := []string{}
    
    fqn = Domain{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *Project) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-domain", "_", "-", -1)
}

func (model *Project) GetDefaultName() string {
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *Project) GetType() string {
    return strings.Replace("project", "_", "-", -1)
}

func (model *Project) GetFQName() []string {
    return model.FQName
}

func (model *Project) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Project) GetParentType() string {
    return model.ParentType
}

func (model *Project) GetUuid() string {
    return model.UUID
}

func (model *Project) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Project) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Project) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Project) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Project) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propProject_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propProject_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propProject_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propProject_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propProject_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propProject_vxlan_routing) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VxlanRouting); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VxlanRouting as vxlan_routing")
        }
        msg["vxlan_routing"] = &val
    }
    
    if model.modified.Bit(propProject_alarm_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AlarmEnable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AlarmEnable as alarm_enable")
        }
        msg["alarm_enable"] = &val
    }
    
    if model.modified.Bit(propProject_quota) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Quota); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Quota as quota")
        }
        msg["quota"] = &val
    }
    
    if model.modified.Bit(propProject_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propProject_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propProject_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Project) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Project) UpdateReferences() error {
    return nil
}


