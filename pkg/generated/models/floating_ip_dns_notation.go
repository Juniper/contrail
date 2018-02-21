
package models
// FloatingIpDnsNotation


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type FloatingIpDnsNotation string

// MakeFloatingIpDnsNotation makes FloatingIpDnsNotation
func MakeFloatingIpDnsNotation() FloatingIpDnsNotation {
    var data FloatingIpDnsNotation
    return data
}



// MakeFloatingIpDnsNotationSlice makes a slice of FloatingIpDnsNotation
func MakeFloatingIpDnsNotationSlice() []FloatingIpDnsNotation {
    return []FloatingIpDnsNotation{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FloatingIpDnsNotation) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FloatingIpDnsNotation) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FloatingIpDnsNotation) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FloatingIpDnsNotation) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FloatingIpDnsNotation) GetFQName() []string {
    return model.FQName
}

func (model *FloatingIpDnsNotation) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FloatingIpDnsNotation) GetParentType() string {
    return model.ParentType
}

func (model *FloatingIpDnsNotation) GetUuid() string {
    return model.UUID
}

func (model *FloatingIpDnsNotation) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FloatingIpDnsNotation) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FloatingIpDnsNotation) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FloatingIpDnsNotation) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FloatingIpDnsNotation) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *FloatingIpDnsNotation) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FloatingIpDnsNotation) UpdateReferences() error {
    return nil
}


