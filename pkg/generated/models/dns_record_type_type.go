
package models
// DnsRecordTypeType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type DnsRecordTypeType string

// MakeDnsRecordTypeType makes DnsRecordTypeType
func MakeDnsRecordTypeType() DnsRecordTypeType {
    var data DnsRecordTypeType
    return data
}



// MakeDnsRecordTypeTypeSlice makes a slice of DnsRecordTypeType
func MakeDnsRecordTypeTypeSlice() []DnsRecordTypeType {
    return []DnsRecordTypeType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DnsRecordTypeType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *DnsRecordTypeType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *DnsRecordTypeType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *DnsRecordTypeType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *DnsRecordTypeType) GetFQName() []string {
    return model.FQName
}

func (model *DnsRecordTypeType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DnsRecordTypeType) GetParentType() string {
    return model.ParentType
}

func (model *DnsRecordTypeType) GetUuid() string {
    return model.UUID
}

func (model *DnsRecordTypeType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DnsRecordTypeType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DnsRecordTypeType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DnsRecordTypeType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DnsRecordTypeType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *DnsRecordTypeType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DnsRecordTypeType) UpdateReferences() error {
    return nil
}


