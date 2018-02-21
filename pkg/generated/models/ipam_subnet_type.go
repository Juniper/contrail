
package models
// IpamSubnetType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propIpamSubnetType_dhcp_option_list int = iota
    propIpamSubnetType_subnet_name int = iota
    propIpamSubnetType_enable_dhcp int = iota
    propIpamSubnetType_alloc_unit int = iota
    propIpamSubnetType_dns_nameservers int = iota
    propIpamSubnetType_dns_server_address int = iota
    propIpamSubnetType_subnet int = iota
    propIpamSubnetType_addr_from_start int = iota
    propIpamSubnetType_allocation_pools int = iota
    propIpamSubnetType_last_modified int = iota
    propIpamSubnetType_host_routes int = iota
    propIpamSubnetType_default_gateway int = iota
    propIpamSubnetType_created int = iota
    propIpamSubnetType_subnet_uuid int = iota
)

// IpamSubnetType 
type IpamSubnetType struct {

    SubnetUUID string `json:"subnet_uuid,omitempty"`
    LastModified string `json:"last_modified,omitempty"`
    HostRoutes *RouteTableType `json:"host_routes,omitempty"`
    DefaultGateway IpAddressType `json:"default_gateway,omitempty"`
    Created string `json:"created,omitempty"`
    DNSNameservers []string `json:"dns_nameservers,omitempty"`
    DHCPOptionList *DhcpOptionsListType `json:"dhcp_option_list,omitempty"`
    SubnetName string `json:"subnet_name,omitempty"`
    EnableDHCP bool `json:"enable_dhcp"`
    AllocUnit int `json:"alloc_unit,omitempty"`
    AllocationPools []*AllocationPoolType `json:"allocation_pools,omitempty"`
    DNSServerAddress IpAddressType `json:"dns_server_address,omitempty"`
    Subnet *SubnetType `json:"subnet,omitempty"`
    AddrFromStart bool `json:"addr_from_start"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *IpamSubnetType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeIpamSubnetType makes IpamSubnetType
func MakeIpamSubnetType() *IpamSubnetType{
    return &IpamSubnetType{
    //TODO(nati): Apply default
    EnableDHCP: false,
        AllocUnit: 0,
        DNSNameservers: []string{},
        DHCPOptionList: MakeDhcpOptionsListType(),
        SubnetName: "",
        Subnet: MakeSubnetType(),
        AddrFromStart: false,
        
            
                AllocationPools:  MakeAllocationPoolTypeSlice(),
            
        DNSServerAddress: MakeIpAddressType(),
        DefaultGateway: MakeIpAddressType(),
        Created: "",
        SubnetUUID: "",
        LastModified: "",
        HostRoutes: MakeRouteTableType(),
        
        modified: big.NewInt(0),
    }
}



// MakeIpamSubnetTypeSlice makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
    return []*IpamSubnetType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IpamSubnetType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IpamSubnetType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamSubnetType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IpamSubnetType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamSubnetType) GetFQName() []string {
    return model.FQName
}

func (model *IpamSubnetType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IpamSubnetType) GetParentType() string {
    return model.ParentType
}

func (model *IpamSubnetType) GetUuid() string {
    return model.UUID
}

func (model *IpamSubnetType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IpamSubnetType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IpamSubnetType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IpamSubnetType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IpamSubnetType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propIpamSubnetType_subnet) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Subnet); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Subnet as subnet")
        }
        msg["subnet"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_addr_from_start) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AddrFromStart); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AddrFromStart as addr_from_start")
        }
        msg["addr_from_start"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_allocation_pools) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AllocationPools); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AllocationPools as allocation_pools")
        }
        msg["allocation_pools"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_dns_server_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DNSServerAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DNSServerAddress as dns_server_address")
        }
        msg["dns_server_address"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_default_gateway) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultGateway); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultGateway as default_gateway")
        }
        msg["default_gateway"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_created) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Created); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Created as created")
        }
        msg["created"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_subnet_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SubnetUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SubnetUUID as subnet_uuid")
        }
        msg["subnet_uuid"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_last_modified) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LastModified); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LastModified as last_modified")
        }
        msg["last_modified"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_host_routes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.HostRoutes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: HostRoutes as host_routes")
        }
        msg["host_routes"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_enable_dhcp) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EnableDHCP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EnableDHCP as enable_dhcp")
        }
        msg["enable_dhcp"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_alloc_unit) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AllocUnit); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AllocUnit as alloc_unit")
        }
        msg["alloc_unit"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_dns_nameservers) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DNSNameservers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DNSNameservers as dns_nameservers")
        }
        msg["dns_nameservers"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_dhcp_option_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DHCPOptionList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DHCPOptionList as dhcp_option_list")
        }
        msg["dhcp_option_list"] = &val
    }
    
    if model.modified.Bit(propIpamSubnetType_subnet_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SubnetName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SubnetName as subnet_name")
        }
        msg["subnet_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *IpamSubnetType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IpamSubnetType) UpdateReferences() error {
    return nil
}


