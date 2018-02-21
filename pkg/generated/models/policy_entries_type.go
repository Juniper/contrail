
package models
// PolicyEntriesType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPolicyEntriesType_policy_rule int = iota
)

// PolicyEntriesType 
type PolicyEntriesType struct {

    PolicyRule []*PolicyRuleType `json:"policy_rule,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PolicyEntriesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePolicyEntriesType makes PolicyEntriesType
func MakePolicyEntriesType() *PolicyEntriesType{
    return &PolicyEntriesType{
    //TODO(nati): Apply default
    
            
                PolicyRule:  MakePolicyRuleTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakePolicyEntriesTypeSlice makes a slice of PolicyEntriesType
func MakePolicyEntriesTypeSlice() []*PolicyEntriesType {
    return []*PolicyEntriesType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PolicyEntriesType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PolicyEntriesType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PolicyEntriesType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PolicyEntriesType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PolicyEntriesType) GetFQName() []string {
    return model.FQName
}

func (model *PolicyEntriesType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PolicyEntriesType) GetParentType() string {
    return model.ParentType
}

func (model *PolicyEntriesType) GetUuid() string {
    return model.UUID
}

func (model *PolicyEntriesType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PolicyEntriesType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PolicyEntriesType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PolicyEntriesType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PolicyEntriesType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPolicyEntriesType_policy_rule) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PolicyRule); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PolicyRule as policy_rule")
        }
        msg["policy_rule"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PolicyEntriesType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PolicyEntriesType) UpdateReferences() error {
    return nil
}


