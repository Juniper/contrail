
package models
// FirewallRuleMatchTagsTypeIdList



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFirewallRuleMatchTagsTypeIdList_tag_type int = iota
)

// FirewallRuleMatchTagsTypeIdList 
type FirewallRuleMatchTagsTypeIdList struct {

    TagType []int `json:"tag_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FirewallRuleMatchTagsTypeIdList) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallRuleMatchTagsTypeIdList makes FirewallRuleMatchTagsTypeIdList
func MakeFirewallRuleMatchTagsTypeIdList() *FirewallRuleMatchTagsTypeIdList{
    return &FirewallRuleMatchTagsTypeIdList{
    //TODO(nati): Apply default
    
            
                TagType: []int{},
            
        
        modified: big.NewInt(0),
    }
}



// MakeFirewallRuleMatchTagsTypeIdListSlice makes a slice of FirewallRuleMatchTagsTypeIdList
func MakeFirewallRuleMatchTagsTypeIdListSlice() []*FirewallRuleMatchTagsTypeIdList {
    return []*FirewallRuleMatchTagsTypeIdList{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FirewallRuleMatchTagsTypeIdList) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FirewallRuleMatchTagsTypeIdList) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallRuleMatchTagsTypeIdList) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FirewallRuleMatchTagsTypeIdList) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallRuleMatchTagsTypeIdList) GetFQName() []string {
    return model.FQName
}

func (model *FirewallRuleMatchTagsTypeIdList) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FirewallRuleMatchTagsTypeIdList) GetParentType() string {
    return model.ParentType
}

func (model *FirewallRuleMatchTagsTypeIdList) GetUuid() string {
    return model.UUID
}

func (model *FirewallRuleMatchTagsTypeIdList) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FirewallRuleMatchTagsTypeIdList) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FirewallRuleMatchTagsTypeIdList) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FirewallRuleMatchTagsTypeIdList) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FirewallRuleMatchTagsTypeIdList) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFirewallRuleMatchTagsTypeIdList_tag_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TagType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TagType as tag_type")
        }
        msg["tag_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FirewallRuleMatchTagsTypeIdList) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FirewallRuleMatchTagsTypeIdList) UpdateReferences() error {
    return nil
}


