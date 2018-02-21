
package models
// IpamType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propIpamType_dhcp_option_list int = iota
    propIpamType_host_routes int = iota
    propIpamType_cidr_block int = iota
    propIpamType_ipam_method int = iota
    propIpamType_ipam_dns_method int = iota
    propIpamType_ipam_dns_server int = iota
)

// IpamType 
type IpamType struct {

    IpamMethod IpamMethodType `json:"ipam_method,omitempty"`
    IpamDNSMethod IpamDnsMethodType `json:"ipam_dns_method,omitempty"`
    IpamDNSServer *IpamDnsAddressType `json:"ipam_dns_server,omitempty"`
    DHCPOptionList *DhcpOptionsListType `json:"dhcp_option_list,omitempty"`
    HostRoutes *RouteTableType `json:"host_routes,omitempty"`
    CidrBlock *SubnetType `json:"cidr_block,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *IpamType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeIpamType makes IpamType
func MakeIpamType() *IpamType{
    return &IpamType{
    //TODO(nati): Apply default
    IpamMethod: MakeIpamMethodType(),
        IpamDNSMethod: MakeIpamDnsMethodType(),
        IpamDNSServer: MakeIpamDnsAddressType(),
        DHCPOptionList: MakeDhcpOptionsListType(),
        HostRoutes: MakeRouteTableType(),
        CidrBlock: MakeSubnetType(),
        
        modified: big.NewInt(0),
    }
}



// MakeIpamTypeSlice makes a slice of IpamType
func MakeIpamTypeSlice() []*IpamType {
    return []*IpamType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IpamType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IpamType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IpamType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamType) GetFQName() []string {
    return model.FQName
}

func (model *IpamType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IpamType) GetParentType() string {
    return model.ParentType
}

func (model *IpamType) GetUuid() string {
    return model.UUID
}

func (model *IpamType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IpamType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IpamType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IpamType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IpamType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propIpamType_ipam_dns_server) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IpamDNSServer); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IpamDNSServer as ipam_dns_server")
        }
        msg["ipam_dns_server"] = &val
    }
    
    if model.modified.Bit(propIpamType_dhcp_option_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DHCPOptionList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DHCPOptionList as dhcp_option_list")
        }
        msg["dhcp_option_list"] = &val
    }
    
    if model.modified.Bit(propIpamType_host_routes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.HostRoutes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: HostRoutes as host_routes")
        }
        msg["host_routes"] = &val
    }
    
    if model.modified.Bit(propIpamType_cidr_block) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.CidrBlock); err != nil {
            return nil, errors.Wrap(err, "Marshal of: CidrBlock as cidr_block")
        }
        msg["cidr_block"] = &val
    }
    
    if model.modified.Bit(propIpamType_ipam_method) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IpamMethod); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IpamMethod as ipam_method")
        }
        msg["ipam_method"] = &val
    }
    
    if model.modified.Bit(propIpamType_ipam_dns_method) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IpamDNSMethod); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IpamDNSMethod as ipam_dns_method")
        }
        msg["ipam_dns_method"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *IpamType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IpamType) UpdateReferences() error {
    return nil
}


