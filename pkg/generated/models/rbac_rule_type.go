
package models
// RbacRuleType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRbacRuleType_rule_object int = iota
    propRbacRuleType_rule_perms int = iota
    propRbacRuleType_rule_field int = iota
)

// RbacRuleType 
type RbacRuleType struct {

    RuleObject string `json:"rule_object,omitempty"`
    RulePerms []*RbacPermType `json:"rule_perms,omitempty"`
    RuleField string `json:"rule_field,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *RbacRuleType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRbacRuleType makes RbacRuleType
func MakeRbacRuleType() *RbacRuleType{
    return &RbacRuleType{
    //TODO(nati): Apply default
    RuleField: "",
        RuleObject: "",
        
            
                RulePerms:  MakeRbacPermTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeRbacRuleTypeSlice makes a slice of RbacRuleType
func MakeRbacRuleTypeSlice() []*RbacRuleType {
    return []*RbacRuleType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RbacRuleType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *RbacRuleType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *RbacRuleType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *RbacRuleType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *RbacRuleType) GetFQName() []string {
    return model.FQName
}

func (model *RbacRuleType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RbacRuleType) GetParentType() string {
    return model.ParentType
}

func (model *RbacRuleType) GetUuid() string {
    return model.UUID
}

func (model *RbacRuleType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RbacRuleType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RbacRuleType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RbacRuleType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RbacRuleType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRbacRuleType_rule_object) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RuleObject); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RuleObject as rule_object")
        }
        msg["rule_object"] = &val
    }
    
    if model.modified.Bit(propRbacRuleType_rule_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RulePerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RulePerms as rule_perms")
        }
        msg["rule_perms"] = &val
    }
    
    if model.modified.Bit(propRbacRuleType_rule_field) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RuleField); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RuleField as rule_field")
        }
        msg["rule_field"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RbacRuleType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RbacRuleType) UpdateReferences() error {
    return nil
}


