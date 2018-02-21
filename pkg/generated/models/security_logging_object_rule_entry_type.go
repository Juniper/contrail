
package models
// SecurityLoggingObjectRuleEntryType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propSecurityLoggingObjectRuleEntryType_rule_uuid int = iota
    propSecurityLoggingObjectRuleEntryType_rate int = iota
)

// SecurityLoggingObjectRuleEntryType 
type SecurityLoggingObjectRuleEntryType struct {

    RuleUUID string `json:"rule_uuid,omitempty"`
    Rate int `json:"rate,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *SecurityLoggingObjectRuleEntryType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSecurityLoggingObjectRuleEntryType makes SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryType() *SecurityLoggingObjectRuleEntryType{
    return &SecurityLoggingObjectRuleEntryType{
    //TODO(nati): Apply default
    Rate: 0,
        RuleUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeSecurityLoggingObjectRuleEntryTypeSlice makes a slice of SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryTypeSlice() []*SecurityLoggingObjectRuleEntryType {
    return []*SecurityLoggingObjectRuleEntryType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SecurityLoggingObjectRuleEntryType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SecurityLoggingObjectRuleEntryType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SecurityLoggingObjectRuleEntryType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SecurityLoggingObjectRuleEntryType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SecurityLoggingObjectRuleEntryType) GetFQName() []string {
    return model.FQName
}

func (model *SecurityLoggingObjectRuleEntryType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SecurityLoggingObjectRuleEntryType) GetParentType() string {
    return model.ParentType
}

func (model *SecurityLoggingObjectRuleEntryType) GetUuid() string {
    return model.UUID
}

func (model *SecurityLoggingObjectRuleEntryType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SecurityLoggingObjectRuleEntryType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SecurityLoggingObjectRuleEntryType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SecurityLoggingObjectRuleEntryType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SecurityLoggingObjectRuleEntryType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propSecurityLoggingObjectRuleEntryType_rule_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RuleUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RuleUUID as rule_uuid")
        }
        msg["rule_uuid"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObjectRuleEntryType_rate) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Rate); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Rate as rate")
        }
        msg["rate"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *SecurityLoggingObjectRuleEntryType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SecurityLoggingObjectRuleEntryType) UpdateReferences() error {
    return nil
}


