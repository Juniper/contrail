
package models
// AddressAllocationModeType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type AddressAllocationModeType string

// MakeAddressAllocationModeType makes AddressAllocationModeType
func MakeAddressAllocationModeType() AddressAllocationModeType {
    var data AddressAllocationModeType
    return data
}



// MakeAddressAllocationModeTypeSlice makes a slice of AddressAllocationModeType
func MakeAddressAllocationModeTypeSlice() []AddressAllocationModeType {
    return []AddressAllocationModeType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AddressAllocationModeType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AddressAllocationModeType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AddressAllocationModeType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AddressAllocationModeType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AddressAllocationModeType) GetFQName() []string {
    return model.FQName
}

func (model *AddressAllocationModeType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AddressAllocationModeType) GetParentType() string {
    return model.ParentType
}

func (model *AddressAllocationModeType) GetUuid() string {
    return model.UUID
}

func (model *AddressAllocationModeType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AddressAllocationModeType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AddressAllocationModeType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AddressAllocationModeType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AddressAllocationModeType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *AddressAllocationModeType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AddressAllocationModeType) UpdateReferences() error {
    return nil
}


