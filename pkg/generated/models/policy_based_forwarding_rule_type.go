
package models
// PolicyBasedForwardingRuleType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPolicyBasedForwardingRuleType_vlan_tag int = iota
    propPolicyBasedForwardingRuleType_src_mac int = iota
    propPolicyBasedForwardingRuleType_service_chain_address int = iota
    propPolicyBasedForwardingRuleType_dst_mac int = iota
    propPolicyBasedForwardingRuleType_protocol int = iota
    propPolicyBasedForwardingRuleType_ipv6_service_chain_address int = iota
    propPolicyBasedForwardingRuleType_direction int = iota
    propPolicyBasedForwardingRuleType_mpls_label int = iota
)

// PolicyBasedForwardingRuleType 
type PolicyBasedForwardingRuleType struct {

    MPLSLabel int `json:"mpls_label,omitempty"`
    VlanTag int `json:"vlan_tag,omitempty"`
    SRCMac string `json:"src_mac,omitempty"`
    ServiceChainAddress string `json:"service_chain_address,omitempty"`
    DSTMac string `json:"dst_mac,omitempty"`
    Protocol string `json:"protocol,omitempty"`
    Ipv6ServiceChainAddress IpAddressType `json:"ipv6_service_chain_address,omitempty"`
    Direction TrafficDirectionType `json:"direction,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PolicyBasedForwardingRuleType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePolicyBasedForwardingRuleType makes PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleType() *PolicyBasedForwardingRuleType{
    return &PolicyBasedForwardingRuleType{
    //TODO(nati): Apply default
    VlanTag: 0,
        SRCMac: "",
        ServiceChainAddress: "",
        DSTMac: "",
        Protocol: "",
        Ipv6ServiceChainAddress: MakeIpAddressType(),
        Direction: MakeTrafficDirectionType(),
        MPLSLabel: 0,
        
        modified: big.NewInt(0),
    }
}



// MakePolicyBasedForwardingRuleTypeSlice makes a slice of PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleTypeSlice() []*PolicyBasedForwardingRuleType {
    return []*PolicyBasedForwardingRuleType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PolicyBasedForwardingRuleType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PolicyBasedForwardingRuleType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PolicyBasedForwardingRuleType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PolicyBasedForwardingRuleType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PolicyBasedForwardingRuleType) GetFQName() []string {
    return model.FQName
}

func (model *PolicyBasedForwardingRuleType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PolicyBasedForwardingRuleType) GetParentType() string {
    return model.ParentType
}

func (model *PolicyBasedForwardingRuleType) GetUuid() string {
    return model.UUID
}

func (model *PolicyBasedForwardingRuleType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PolicyBasedForwardingRuleType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PolicyBasedForwardingRuleType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PolicyBasedForwardingRuleType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PolicyBasedForwardingRuleType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPolicyBasedForwardingRuleType_service_chain_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceChainAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceChainAddress as service_chain_address")
        }
        msg["service_chain_address"] = &val
    }
    
    if model.modified.Bit(propPolicyBasedForwardingRuleType_dst_mac) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DSTMac); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DSTMac as dst_mac")
        }
        msg["dst_mac"] = &val
    }
    
    if model.modified.Bit(propPolicyBasedForwardingRuleType_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propPolicyBasedForwardingRuleType_ipv6_service_chain_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Ipv6ServiceChainAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Ipv6ServiceChainAddress as ipv6_service_chain_address")
        }
        msg["ipv6_service_chain_address"] = &val
    }
    
    if model.modified.Bit(propPolicyBasedForwardingRuleType_direction) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Direction); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Direction as direction")
        }
        msg["direction"] = &val
    }
    
    if model.modified.Bit(propPolicyBasedForwardingRuleType_mpls_label) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MPLSLabel); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MPLSLabel as mpls_label")
        }
        msg["mpls_label"] = &val
    }
    
    if model.modified.Bit(propPolicyBasedForwardingRuleType_vlan_tag) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VlanTag); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VlanTag as vlan_tag")
        }
        msg["vlan_tag"] = &val
    }
    
    if model.modified.Bit(propPolicyBasedForwardingRuleType_src_mac) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SRCMac); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SRCMac as src_mac")
        }
        msg["src_mac"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PolicyBasedForwardingRuleType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PolicyBasedForwardingRuleType) UpdateReferences() error {
    return nil
}


