
package models
// FirewallRuleMatchTagsType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFirewallRuleMatchTagsType_tag_list int = iota
)

// FirewallRuleMatchTagsType 
type FirewallRuleMatchTagsType struct {

    TagList []string `json:"tag_list,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FirewallRuleMatchTagsType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallRuleMatchTagsType makes FirewallRuleMatchTagsType
func MakeFirewallRuleMatchTagsType() *FirewallRuleMatchTagsType{
    return &FirewallRuleMatchTagsType{
    //TODO(nati): Apply default
    TagList: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeFirewallRuleMatchTagsTypeSlice makes a slice of FirewallRuleMatchTagsType
func MakeFirewallRuleMatchTagsTypeSlice() []*FirewallRuleMatchTagsType {
    return []*FirewallRuleMatchTagsType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FirewallRuleMatchTagsType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FirewallRuleMatchTagsType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallRuleMatchTagsType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FirewallRuleMatchTagsType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallRuleMatchTagsType) GetFQName() []string {
    return model.FQName
}

func (model *FirewallRuleMatchTagsType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FirewallRuleMatchTagsType) GetParentType() string {
    return model.ParentType
}

func (model *FirewallRuleMatchTagsType) GetUuid() string {
    return model.UUID
}

func (model *FirewallRuleMatchTagsType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FirewallRuleMatchTagsType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FirewallRuleMatchTagsType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FirewallRuleMatchTagsType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FirewallRuleMatchTagsType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFirewallRuleMatchTagsType_tag_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TagList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TagList as tag_list")
        }
        msg["tag_list"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FirewallRuleMatchTagsType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FirewallRuleMatchTagsType) UpdateReferences() error {
    return nil
}


