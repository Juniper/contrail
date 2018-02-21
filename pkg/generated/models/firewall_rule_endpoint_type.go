
package models
// FirewallRuleEndpointType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFirewallRuleEndpointType_virtual_network int = iota
    propFirewallRuleEndpointType_any int = iota
    propFirewallRuleEndpointType_address_group int = iota
    propFirewallRuleEndpointType_subnet int = iota
    propFirewallRuleEndpointType_tags int = iota
    propFirewallRuleEndpointType_tag_ids int = iota
)

// FirewallRuleEndpointType 
type FirewallRuleEndpointType struct {

    AddressGroup string `json:"address_group,omitempty"`
    Subnet *SubnetType `json:"subnet,omitempty"`
    Tags []string `json:"tags,omitempty"`
    TagIds []int `json:"tag_ids,omitempty"`
    VirtualNetwork string `json:"virtual_network,omitempty"`
    Any bool `json:"any"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FirewallRuleEndpointType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallRuleEndpointType makes FirewallRuleEndpointType
func MakeFirewallRuleEndpointType() *FirewallRuleEndpointType{
    return &FirewallRuleEndpointType{
    //TODO(nati): Apply default
    Subnet: MakeSubnetType(),
        Tags: []string{},
        
            
                TagIds: []int{},
            
        VirtualNetwork: "",
        Any: false,
        AddressGroup: "",
        
        modified: big.NewInt(0),
    }
}



// MakeFirewallRuleEndpointTypeSlice makes a slice of FirewallRuleEndpointType
func MakeFirewallRuleEndpointTypeSlice() []*FirewallRuleEndpointType {
    return []*FirewallRuleEndpointType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FirewallRuleEndpointType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FirewallRuleEndpointType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallRuleEndpointType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FirewallRuleEndpointType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallRuleEndpointType) GetFQName() []string {
    return model.FQName
}

func (model *FirewallRuleEndpointType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FirewallRuleEndpointType) GetParentType() string {
    return model.ParentType
}

func (model *FirewallRuleEndpointType) GetUuid() string {
    return model.UUID
}

func (model *FirewallRuleEndpointType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FirewallRuleEndpointType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FirewallRuleEndpointType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FirewallRuleEndpointType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FirewallRuleEndpointType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFirewallRuleEndpointType_virtual_network) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualNetwork); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualNetwork as virtual_network")
        }
        msg["virtual_network"] = &val
    }
    
    if model.modified.Bit(propFirewallRuleEndpointType_any) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Any); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Any as any")
        }
        msg["any"] = &val
    }
    
    if model.modified.Bit(propFirewallRuleEndpointType_address_group) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AddressGroup); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AddressGroup as address_group")
        }
        msg["address_group"] = &val
    }
    
    if model.modified.Bit(propFirewallRuleEndpointType_subnet) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Subnet); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Subnet as subnet")
        }
        msg["subnet"] = &val
    }
    
    if model.modified.Bit(propFirewallRuleEndpointType_tags) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Tags); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Tags as tags")
        }
        msg["tags"] = &val
    }
    
    if model.modified.Bit(propFirewallRuleEndpointType_tag_ids) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TagIds); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TagIds as tag_ids")
        }
        msg["tag_ids"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FirewallRuleEndpointType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FirewallRuleEndpointType) UpdateReferences() error {
    return nil
}


