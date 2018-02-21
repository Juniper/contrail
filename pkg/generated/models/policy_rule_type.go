
package models
// PolicyRuleType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPolicyRuleType_protocol int = iota
    propPolicyRuleType_action_list int = iota
    propPolicyRuleType_dst_ports int = iota
    propPolicyRuleType_application int = iota
    propPolicyRuleType_src_ports int = iota
    propPolicyRuleType_rule_sequence int = iota
    propPolicyRuleType_direction int = iota
    propPolicyRuleType_dst_addresses int = iota
    propPolicyRuleType_created int = iota
    propPolicyRuleType_rule_uuid int = iota
    propPolicyRuleType_last_modified int = iota
    propPolicyRuleType_ethertype int = iota
    propPolicyRuleType_src_addresses int = iota
)

// PolicyRuleType 
type PolicyRuleType struct {

    SRCPorts []*PortType `json:"src_ports,omitempty"`
    Protocol string `json:"protocol,omitempty"`
    ActionList *ActionListType `json:"action_list,omitempty"`
    DSTPorts []*PortType `json:"dst_ports,omitempty"`
    Application []string `json:"application,omitempty"`
    LastModified string `json:"last_modified,omitempty"`
    Ethertype EtherType `json:"ethertype,omitempty"`
    SRCAddresses []*AddressType `json:"src_addresses,omitempty"`
    RuleSequence *SequenceType `json:"rule_sequence,omitempty"`
    Direction DirectionType `json:"direction,omitempty"`
    DSTAddresses []*AddressType `json:"dst_addresses,omitempty"`
    Created string `json:"created,omitempty"`
    RuleUUID string `json:"rule_uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PolicyRuleType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePolicyRuleType makes PolicyRuleType
func MakePolicyRuleType() *PolicyRuleType{
    return &PolicyRuleType{
    //TODO(nati): Apply default
    Created: "",
        RuleUUID: "",
        LastModified: "",
        Ethertype: MakeEtherType(),
        
            
                SRCAddresses:  MakeAddressTypeSlice(),
            
        RuleSequence: MakeSequenceType(),
        Direction: MakeDirectionType(),
        
            
                DSTAddresses:  MakeAddressTypeSlice(),
            
        
            
                DSTPorts:  MakePortTypeSlice(),
            
        Application: []string{},
        
            
                SRCPorts:  MakePortTypeSlice(),
            
        Protocol: "",
        ActionList: MakeActionListType(),
        
        modified: big.NewInt(0),
    }
}



// MakePolicyRuleTypeSlice makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
    return []*PolicyRuleType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PolicyRuleType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PolicyRuleType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PolicyRuleType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PolicyRuleType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PolicyRuleType) GetFQName() []string {
    return model.FQName
}

func (model *PolicyRuleType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PolicyRuleType) GetParentType() string {
    return model.ParentType
}

func (model *PolicyRuleType) GetUuid() string {
    return model.UUID
}

func (model *PolicyRuleType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PolicyRuleType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PolicyRuleType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PolicyRuleType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PolicyRuleType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPolicyRuleType_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_action_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ActionList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ActionList as action_list")
        }
        msg["action_list"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_dst_ports) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DSTPorts); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DSTPorts as dst_ports")
        }
        msg["dst_ports"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_application) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Application); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Application as application")
        }
        msg["application"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_src_ports) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SRCPorts); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SRCPorts as src_ports")
        }
        msg["src_ports"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_direction) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Direction); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Direction as direction")
        }
        msg["direction"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_dst_addresses) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DSTAddresses); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DSTAddresses as dst_addresses")
        }
        msg["dst_addresses"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_created) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Created); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Created as created")
        }
        msg["created"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_rule_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RuleUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RuleUUID as rule_uuid")
        }
        msg["rule_uuid"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_last_modified) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LastModified); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LastModified as last_modified")
        }
        msg["last_modified"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_ethertype) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Ethertype); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Ethertype as ethertype")
        }
        msg["ethertype"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_src_addresses) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SRCAddresses); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SRCAddresses as src_addresses")
        }
        msg["src_addresses"] = &val
    }
    
    if model.modified.Bit(propPolicyRuleType_rule_sequence) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RuleSequence); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RuleSequence as rule_sequence")
        }
        msg["rule_sequence"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PolicyRuleType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PolicyRuleType) UpdateReferences() error {
    return nil
}


