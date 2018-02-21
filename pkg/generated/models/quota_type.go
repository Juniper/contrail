
package models
// QuotaType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propQuotaType_route_table int = iota
    propQuotaType_virtual_DNS_record int = iota
    propQuotaType_security_group_rule int = iota
    propQuotaType_instance_ip int = iota
    propQuotaType_defaults int = iota
    propQuotaType_loadbalancer_pool int = iota
    propQuotaType_subnet int = iota
    propQuotaType_virtual_DNS int = iota
    propQuotaType_floating_ip int = iota
    propQuotaType_floating_ip_pool int = iota
    propQuotaType_loadbalancer_healthmonitor int = iota
    propQuotaType_virtual_ip int = iota
    propQuotaType_security_group int = iota
    propQuotaType_network_policy int = iota
    propQuotaType_logical_router int = iota
    propQuotaType_service_instance int = iota
    propQuotaType_loadbalancer_member int = iota
    propQuotaType_virtual_machine_interface int = iota
    propQuotaType_virtual_router int = iota
    propQuotaType_network_ipam int = iota
    propQuotaType_service_template int = iota
    propQuotaType_bgp_router int = iota
    propQuotaType_access_control_list int = iota
    propQuotaType_global_vrouter_config int = iota
    propQuotaType_security_logging_object int = iota
    propQuotaType_virtual_network int = iota
)

// QuotaType 
type QuotaType struct {

    NetworkPolicy int `json:"network_policy,omitempty"`
    LogicalRouter int `json:"logical_router,omitempty"`
    ServiceInstance int `json:"service_instance,omitempty"`
    LoadbalancerMember int `json:"loadbalancer_member,omitempty"`
    VirtualMachineInterface int `json:"virtual_machine_interface,omitempty"`
    VirtualRouter int `json:"virtual_router,omitempty"`
    NetworkIpam int `json:"network_ipam,omitempty"`
    ServiceTemplate int `json:"service_template,omitempty"`
    BGPRouter int `json:"bgp_router,omitempty"`
    AccessControlList int `json:"access_control_list,omitempty"`
    GlobalVrouterConfig int `json:"global_vrouter_config,omitempty"`
    SecurityLoggingObject int `json:"security_logging_object,omitempty"`
    VirtualNetwork int `json:"virtual_network,omitempty"`
    RouteTable int `json:"route_table,omitempty"`
    VirtualDNSRecord int `json:"virtual_DNS_record,omitempty"`
    SecurityGroupRule int `json:"security_group_rule,omitempty"`
    InstanceIP int `json:"instance_ip,omitempty"`
    Defaults int `json:"defaults,omitempty"`
    LoadbalancerPool int `json:"loadbalancer_pool,omitempty"`
    Subnet int `json:"subnet,omitempty"`
    VirtualDNS int `json:"virtual_DNS,omitempty"`
    FloatingIP int `json:"floating_ip,omitempty"`
    FloatingIPPool int `json:"floating_ip_pool,omitempty"`
    LoadbalancerHealthmonitor int `json:"loadbalancer_healthmonitor,omitempty"`
    VirtualIP int `json:"virtual_ip,omitempty"`
    SecurityGroup int `json:"security_group,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *QuotaType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeQuotaType makes QuotaType
func MakeQuotaType() *QuotaType{
    return &QuotaType{
    //TODO(nati): Apply default
    ServiceTemplate: 0,
        BGPRouter: 0,
        AccessControlList: 0,
        GlobalVrouterConfig: 0,
        SecurityLoggingObject: 0,
        VirtualNetwork: 0,
        VirtualRouter: 0,
        NetworkIpam: 0,
        SecurityGroupRule: 0,
        InstanceIP: 0,
        Defaults: 0,
        RouteTable: 0,
        VirtualDNSRecord: 0,
        VirtualDNS: 0,
        FloatingIP: 0,
        FloatingIPPool: 0,
        LoadbalancerHealthmonitor: 0,
        VirtualIP: 0,
        SecurityGroup: 0,
        LoadbalancerPool: 0,
        Subnet: 0,
        ServiceInstance: 0,
        LoadbalancerMember: 0,
        VirtualMachineInterface: 0,
        NetworkPolicy: 0,
        LogicalRouter: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeQuotaTypeSlice makes a slice of QuotaType
func MakeQuotaTypeSlice() []*QuotaType {
    return []*QuotaType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *QuotaType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *QuotaType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *QuotaType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *QuotaType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *QuotaType) GetFQName() []string {
    return model.FQName
}

func (model *QuotaType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *QuotaType) GetParentType() string {
    return model.ParentType
}

func (model *QuotaType) GetUuid() string {
    return model.UUID
}

func (model *QuotaType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *QuotaType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *QuotaType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *QuotaType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *QuotaType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propQuotaType_virtual_DNS) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualDNS); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualDNS as virtual_DNS")
        }
        msg["virtual_DNS"] = &val
    }
    
    if model.modified.Bit(propQuotaType_floating_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIP as floating_ip")
        }
        msg["floating_ip"] = &val
    }
    
    if model.modified.Bit(propQuotaType_floating_ip_pool) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPPool); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPPool as floating_ip_pool")
        }
        msg["floating_ip_pool"] = &val
    }
    
    if model.modified.Bit(propQuotaType_loadbalancer_healthmonitor) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerHealthmonitor); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerHealthmonitor as loadbalancer_healthmonitor")
        }
        msg["loadbalancer_healthmonitor"] = &val
    }
    
    if model.modified.Bit(propQuotaType_virtual_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualIP as virtual_ip")
        }
        msg["virtual_ip"] = &val
    }
    
    if model.modified.Bit(propQuotaType_security_group) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecurityGroup); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecurityGroup as security_group")
        }
        msg["security_group"] = &val
    }
    
    if model.modified.Bit(propQuotaType_loadbalancer_pool) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerPool); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerPool as loadbalancer_pool")
        }
        msg["loadbalancer_pool"] = &val
    }
    
    if model.modified.Bit(propQuotaType_subnet) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Subnet); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Subnet as subnet")
        }
        msg["subnet"] = &val
    }
    
    if model.modified.Bit(propQuotaType_service_instance) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceInstance); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceInstance as service_instance")
        }
        msg["service_instance"] = &val
    }
    
    if model.modified.Bit(propQuotaType_loadbalancer_member) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerMember); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerMember as loadbalancer_member")
        }
        msg["loadbalancer_member"] = &val
    }
    
    if model.modified.Bit(propQuotaType_virtual_machine_interface) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterface); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterface as virtual_machine_interface")
        }
        msg["virtual_machine_interface"] = &val
    }
    
    if model.modified.Bit(propQuotaType_network_policy) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NetworkPolicy); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NetworkPolicy as network_policy")
        }
        msg["network_policy"] = &val
    }
    
    if model.modified.Bit(propQuotaType_logical_router) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LogicalRouter); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LogicalRouter as logical_router")
        }
        msg["logical_router"] = &val
    }
    
    if model.modified.Bit(propQuotaType_service_template) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceTemplate); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceTemplate as service_template")
        }
        msg["service_template"] = &val
    }
    
    if model.modified.Bit(propQuotaType_bgp_router) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BGPRouter); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BGPRouter as bgp_router")
        }
        msg["bgp_router"] = &val
    }
    
    if model.modified.Bit(propQuotaType_access_control_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AccessControlList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AccessControlList as access_control_list")
        }
        msg["access_control_list"] = &val
    }
    
    if model.modified.Bit(propQuotaType_global_vrouter_config) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GlobalVrouterConfig); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GlobalVrouterConfig as global_vrouter_config")
        }
        msg["global_vrouter_config"] = &val
    }
    
    if model.modified.Bit(propQuotaType_security_logging_object) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecurityLoggingObject); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecurityLoggingObject as security_logging_object")
        }
        msg["security_logging_object"] = &val
    }
    
    if model.modified.Bit(propQuotaType_virtual_network) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualNetwork); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualNetwork as virtual_network")
        }
        msg["virtual_network"] = &val
    }
    
    if model.modified.Bit(propQuotaType_virtual_router) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualRouter); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualRouter as virtual_router")
        }
        msg["virtual_router"] = &val
    }
    
    if model.modified.Bit(propQuotaType_network_ipam) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NetworkIpam); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NetworkIpam as network_ipam")
        }
        msg["network_ipam"] = &val
    }
    
    if model.modified.Bit(propQuotaType_security_group_rule) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecurityGroupRule); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecurityGroupRule as security_group_rule")
        }
        msg["security_group_rule"] = &val
    }
    
    if model.modified.Bit(propQuotaType_instance_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InstanceIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InstanceIP as instance_ip")
        }
        msg["instance_ip"] = &val
    }
    
    if model.modified.Bit(propQuotaType_defaults) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Defaults); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Defaults as defaults")
        }
        msg["defaults"] = &val
    }
    
    if model.modified.Bit(propQuotaType_route_table) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RouteTable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RouteTable as route_table")
        }
        msg["route_table"] = &val
    }
    
    if model.modified.Bit(propQuotaType_virtual_DNS_record) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualDNSRecord); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualDNSRecord as virtual_DNS_record")
        }
        msg["virtual_DNS_record"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *QuotaType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *QuotaType) UpdateReferences() error {
    return nil
}


