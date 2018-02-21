
package models
// DnsRecordClassType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type DnsRecordClassType string

// MakeDnsRecordClassType makes DnsRecordClassType
func MakeDnsRecordClassType() DnsRecordClassType {
    var data DnsRecordClassType
    return data
}



// MakeDnsRecordClassTypeSlice makes a slice of DnsRecordClassType
func MakeDnsRecordClassTypeSlice() []DnsRecordClassType {
    return []DnsRecordClassType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DnsRecordClassType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *DnsRecordClassType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *DnsRecordClassType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *DnsRecordClassType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *DnsRecordClassType) GetFQName() []string {
    return model.FQName
}

func (model *DnsRecordClassType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DnsRecordClassType) GetParentType() string {
    return model.ParentType
}

func (model *DnsRecordClassType) GetUuid() string {
    return model.UUID
}

func (model *DnsRecordClassType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DnsRecordClassType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DnsRecordClassType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DnsRecordClassType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DnsRecordClassType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *DnsRecordClassType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DnsRecordClassType) UpdateReferences() error {
    return nil
}


