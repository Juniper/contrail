
package models
// DhcpOptionType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propDhcpOptionType_dhcp_option_value int = iota
    propDhcpOptionType_dhcp_option_value_bytes int = iota
    propDhcpOptionType_dhcp_option_name int = iota
)

// DhcpOptionType 
type DhcpOptionType struct {

    DHCPOptionValue string `json:"dhcp_option_value,omitempty"`
    DHCPOptionValueBytes string `json:"dhcp_option_value_bytes,omitempty"`
    DHCPOptionName string `json:"dhcp_option_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *DhcpOptionType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeDhcpOptionType makes DhcpOptionType
func MakeDhcpOptionType() *DhcpOptionType{
    return &DhcpOptionType{
    //TODO(nati): Apply default
    DHCPOptionValue: "",
        DHCPOptionValueBytes: "",
        DHCPOptionName: "",
        
        modified: big.NewInt(0),
    }
}



// MakeDhcpOptionTypeSlice makes a slice of DhcpOptionType
func MakeDhcpOptionTypeSlice() []*DhcpOptionType {
    return []*DhcpOptionType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DhcpOptionType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *DhcpOptionType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *DhcpOptionType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *DhcpOptionType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *DhcpOptionType) GetFQName() []string {
    return model.FQName
}

func (model *DhcpOptionType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DhcpOptionType) GetParentType() string {
    return model.ParentType
}

func (model *DhcpOptionType) GetUuid() string {
    return model.UUID
}

func (model *DhcpOptionType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DhcpOptionType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DhcpOptionType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DhcpOptionType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DhcpOptionType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propDhcpOptionType_dhcp_option_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DHCPOptionName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DHCPOptionName as dhcp_option_name")
        }
        msg["dhcp_option_name"] = &val
    }
    
    if model.modified.Bit(propDhcpOptionType_dhcp_option_value) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DHCPOptionValue); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DHCPOptionValue as dhcp_option_value")
        }
        msg["dhcp_option_value"] = &val
    }
    
    if model.modified.Bit(propDhcpOptionType_dhcp_option_value_bytes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DHCPOptionValueBytes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DHCPOptionValueBytes as dhcp_option_value_bytes")
        }
        msg["dhcp_option_value_bytes"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *DhcpOptionType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DhcpOptionType) UpdateReferences() error {
    return nil
}


