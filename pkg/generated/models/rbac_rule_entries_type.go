
package models
// RbacRuleEntriesType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRbacRuleEntriesType_rbac_rule int = iota
)

// RbacRuleEntriesType 
type RbacRuleEntriesType struct {

    RbacRule []*RbacRuleType `json:"rbac_rule,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *RbacRuleEntriesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRbacRuleEntriesType makes RbacRuleEntriesType
func MakeRbacRuleEntriesType() *RbacRuleEntriesType{
    return &RbacRuleEntriesType{
    //TODO(nati): Apply default
    
            
                RbacRule:  MakeRbacRuleTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeRbacRuleEntriesTypeSlice makes a slice of RbacRuleEntriesType
func MakeRbacRuleEntriesTypeSlice() []*RbacRuleEntriesType {
    return []*RbacRuleEntriesType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RbacRuleEntriesType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *RbacRuleEntriesType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *RbacRuleEntriesType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *RbacRuleEntriesType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *RbacRuleEntriesType) GetFQName() []string {
    return model.FQName
}

func (model *RbacRuleEntriesType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RbacRuleEntriesType) GetParentType() string {
    return model.ParentType
}

func (model *RbacRuleEntriesType) GetUuid() string {
    return model.UUID
}

func (model *RbacRuleEntriesType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RbacRuleEntriesType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RbacRuleEntriesType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RbacRuleEntriesType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RbacRuleEntriesType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRbacRuleEntriesType_rbac_rule) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RbacRule); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RbacRule as rbac_rule")
        }
        msg["rbac_rule"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RbacRuleEntriesType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RbacRuleEntriesType) UpdateReferences() error {
    return nil
}


