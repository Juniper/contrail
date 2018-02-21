
package models
// FirewallRuleDirectionType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type FirewallRuleDirectionType string

// MakeFirewallRuleDirectionType makes FirewallRuleDirectionType
func MakeFirewallRuleDirectionType() FirewallRuleDirectionType {
    var data FirewallRuleDirectionType
    return data
}



// MakeFirewallRuleDirectionTypeSlice makes a slice of FirewallRuleDirectionType
func MakeFirewallRuleDirectionTypeSlice() []FirewallRuleDirectionType {
    return []FirewallRuleDirectionType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FirewallRuleDirectionType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FirewallRuleDirectionType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallRuleDirectionType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FirewallRuleDirectionType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallRuleDirectionType) GetFQName() []string {
    return model.FQName
}

func (model *FirewallRuleDirectionType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FirewallRuleDirectionType) GetParentType() string {
    return model.ParentType
}

func (model *FirewallRuleDirectionType) GetUuid() string {
    return model.UUID
}

func (model *FirewallRuleDirectionType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FirewallRuleDirectionType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FirewallRuleDirectionType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FirewallRuleDirectionType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FirewallRuleDirectionType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *FirewallRuleDirectionType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FirewallRuleDirectionType) UpdateReferences() error {
    return nil
}


