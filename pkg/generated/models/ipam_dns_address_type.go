
package models
// IpamDnsAddressType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propIpamDnsAddressType_tenant_dns_server_address int = iota
    propIpamDnsAddressType_virtual_dns_server_name int = iota
)

// IpamDnsAddressType 
type IpamDnsAddressType struct {

    VirtualDNSServerName string `json:"virtual_dns_server_name,omitempty"`
    TenantDNSServerAddress *IpAddressesType `json:"tenant_dns_server_address,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *IpamDnsAddressType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeIpamDnsAddressType makes IpamDnsAddressType
func MakeIpamDnsAddressType() *IpamDnsAddressType{
    return &IpamDnsAddressType{
    //TODO(nati): Apply default
    VirtualDNSServerName: "",
        TenantDNSServerAddress: MakeIpAddressesType(),
        
        modified: big.NewInt(0),
    }
}



// MakeIpamDnsAddressTypeSlice makes a slice of IpamDnsAddressType
func MakeIpamDnsAddressTypeSlice() []*IpamDnsAddressType {
    return []*IpamDnsAddressType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IpamDnsAddressType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IpamDnsAddressType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamDnsAddressType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IpamDnsAddressType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamDnsAddressType) GetFQName() []string {
    return model.FQName
}

func (model *IpamDnsAddressType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IpamDnsAddressType) GetParentType() string {
    return model.ParentType
}

func (model *IpamDnsAddressType) GetUuid() string {
    return model.UUID
}

func (model *IpamDnsAddressType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IpamDnsAddressType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IpamDnsAddressType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IpamDnsAddressType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IpamDnsAddressType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propIpamDnsAddressType_tenant_dns_server_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TenantDNSServerAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TenantDNSServerAddress as tenant_dns_server_address")
        }
        msg["tenant_dns_server_address"] = &val
    }
    
    if model.modified.Bit(propIpamDnsAddressType_virtual_dns_server_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualDNSServerName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualDNSServerName as virtual_dns_server_name")
        }
        msg["virtual_dns_server_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *IpamDnsAddressType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IpamDnsAddressType) UpdateReferences() error {
    return nil
}


