
package models
// SecurityLoggingObjectRuleListType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propSecurityLoggingObjectRuleListType_rule int = iota
)

// SecurityLoggingObjectRuleListType 
type SecurityLoggingObjectRuleListType struct {

    Rule []*SecurityLoggingObjectRuleEntryType `json:"rule,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *SecurityLoggingObjectRuleListType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListType() *SecurityLoggingObjectRuleListType{
    return &SecurityLoggingObjectRuleListType{
    //TODO(nati): Apply default
    
            
                Rule:  MakeSecurityLoggingObjectRuleEntryTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeSecurityLoggingObjectRuleListTypeSlice makes a slice of SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListTypeSlice() []*SecurityLoggingObjectRuleListType {
    return []*SecurityLoggingObjectRuleListType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SecurityLoggingObjectRuleListType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SecurityLoggingObjectRuleListType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SecurityLoggingObjectRuleListType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SecurityLoggingObjectRuleListType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SecurityLoggingObjectRuleListType) GetFQName() []string {
    return model.FQName
}

func (model *SecurityLoggingObjectRuleListType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SecurityLoggingObjectRuleListType) GetParentType() string {
    return model.ParentType
}

func (model *SecurityLoggingObjectRuleListType) GetUuid() string {
    return model.UUID
}

func (model *SecurityLoggingObjectRuleListType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SecurityLoggingObjectRuleListType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SecurityLoggingObjectRuleListType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SecurityLoggingObjectRuleListType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SecurityLoggingObjectRuleListType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propSecurityLoggingObjectRuleListType_rule) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Rule); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Rule as rule")
        }
        msg["rule"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *SecurityLoggingObjectRuleListType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SecurityLoggingObjectRuleListType) UpdateReferences() error {
    return nil
}


