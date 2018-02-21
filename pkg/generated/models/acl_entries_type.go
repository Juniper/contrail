
package models
// AclEntriesType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAclEntriesType_dynamic int = iota
    propAclEntriesType_acl_rule int = iota
)

// AclEntriesType 
type AclEntriesType struct {

    Dynamic bool `json:"dynamic"`
    ACLRule []*AclRuleType `json:"acl_rule,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AclEntriesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAclEntriesType makes AclEntriesType
func MakeAclEntriesType() *AclEntriesType{
    return &AclEntriesType{
    //TODO(nati): Apply default
    Dynamic: false,
        
            
                ACLRule:  MakeAclRuleTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeAclEntriesTypeSlice makes a slice of AclEntriesType
func MakeAclEntriesTypeSlice() []*AclEntriesType {
    return []*AclEntriesType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AclEntriesType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AclEntriesType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AclEntriesType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AclEntriesType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AclEntriesType) GetFQName() []string {
    return model.FQName
}

func (model *AclEntriesType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AclEntriesType) GetParentType() string {
    return model.ParentType
}

func (model *AclEntriesType) GetUuid() string {
    return model.UUID
}

func (model *AclEntriesType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AclEntriesType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AclEntriesType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AclEntriesType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AclEntriesType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAclEntriesType_dynamic) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Dynamic); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Dynamic as dynamic")
        }
        msg["dynamic"] = &val
    }
    
    if model.modified.Bit(propAclEntriesType_acl_rule) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ACLRule); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ACLRule as acl_rule")
        }
        msg["acl_rule"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AclEntriesType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AclEntriesType) UpdateReferences() error {
    return nil
}


