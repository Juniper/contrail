
package models
// DnsRecordOrderType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type DnsRecordOrderType string

// MakeDnsRecordOrderType makes DnsRecordOrderType
func MakeDnsRecordOrderType() DnsRecordOrderType {
    var data DnsRecordOrderType
    return data
}



// MakeDnsRecordOrderTypeSlice makes a slice of DnsRecordOrderType
func MakeDnsRecordOrderTypeSlice() []DnsRecordOrderType {
    return []DnsRecordOrderType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DnsRecordOrderType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *DnsRecordOrderType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *DnsRecordOrderType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *DnsRecordOrderType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *DnsRecordOrderType) GetFQName() []string {
    return model.FQName
}

func (model *DnsRecordOrderType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DnsRecordOrderType) GetParentType() string {
    return model.ParentType
}

func (model *DnsRecordOrderType) GetUuid() string {
    return model.UUID
}

func (model *DnsRecordOrderType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DnsRecordOrderType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DnsRecordOrderType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DnsRecordOrderType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DnsRecordOrderType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *DnsRecordOrderType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DnsRecordOrderType) UpdateReferences() error {
    return nil
}


