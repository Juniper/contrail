
package models
// VrfAssignRuleType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVrfAssignRuleType_routing_instance int = iota
    propVrfAssignRuleType_match_condition int = iota
    propVrfAssignRuleType_vlan_tag int = iota
    propVrfAssignRuleType_ignore_acl int = iota
)

// VrfAssignRuleType 
type VrfAssignRuleType struct {

    RoutingInstance string `json:"routing_instance,omitempty"`
    MatchCondition *MatchConditionType `json:"match_condition,omitempty"`
    VlanTag int `json:"vlan_tag,omitempty"`
    IgnoreACL bool `json:"ignore_acl"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VrfAssignRuleType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVrfAssignRuleType makes VrfAssignRuleType
func MakeVrfAssignRuleType() *VrfAssignRuleType{
    return &VrfAssignRuleType{
    //TODO(nati): Apply default
    VlanTag: 0,
        IgnoreACL: false,
        RoutingInstance: "",
        MatchCondition: MakeMatchConditionType(),
        
        modified: big.NewInt(0),
    }
}



// MakeVrfAssignRuleTypeSlice makes a slice of VrfAssignRuleType
func MakeVrfAssignRuleTypeSlice() []*VrfAssignRuleType {
    return []*VrfAssignRuleType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VrfAssignRuleType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VrfAssignRuleType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VrfAssignRuleType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VrfAssignRuleType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VrfAssignRuleType) GetFQName() []string {
    return model.FQName
}

func (model *VrfAssignRuleType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VrfAssignRuleType) GetParentType() string {
    return model.ParentType
}

func (model *VrfAssignRuleType) GetUuid() string {
    return model.UUID
}

func (model *VrfAssignRuleType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VrfAssignRuleType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VrfAssignRuleType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VrfAssignRuleType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VrfAssignRuleType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVrfAssignRuleType_match_condition) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MatchCondition); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MatchCondition as match_condition")
        }
        msg["match_condition"] = &val
    }
    
    if model.modified.Bit(propVrfAssignRuleType_vlan_tag) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VlanTag); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VlanTag as vlan_tag")
        }
        msg["vlan_tag"] = &val
    }
    
    if model.modified.Bit(propVrfAssignRuleType_ignore_acl) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IgnoreACL); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IgnoreACL as ignore_acl")
        }
        msg["ignore_acl"] = &val
    }
    
    if model.modified.Bit(propVrfAssignRuleType_routing_instance) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RoutingInstance); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RoutingInstance as routing_instance")
        }
        msg["routing_instance"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VrfAssignRuleType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VrfAssignRuleType) UpdateReferences() error {
    return nil
}


