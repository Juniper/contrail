
package models
// LongLivedGracefulRestartTimeType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type LongLivedGracefulRestartTimeType int

// MakeLongLivedGracefulRestartTimeType makes LongLivedGracefulRestartTimeType
func MakeLongLivedGracefulRestartTimeType() LongLivedGracefulRestartTimeType {
    var data LongLivedGracefulRestartTimeType
    return data
}



// MakeLongLivedGracefulRestartTimeTypeSlice makes a slice of LongLivedGracefulRestartTimeType
func MakeLongLivedGracefulRestartTimeTypeSlice() []LongLivedGracefulRestartTimeType {
    return []LongLivedGracefulRestartTimeType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LongLivedGracefulRestartTimeType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LongLivedGracefulRestartTimeType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LongLivedGracefulRestartTimeType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LongLivedGracefulRestartTimeType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LongLivedGracefulRestartTimeType) GetFQName() []string {
    return model.FQName
}

func (model *LongLivedGracefulRestartTimeType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LongLivedGracefulRestartTimeType) GetParentType() string {
    return model.ParentType
}

func (model *LongLivedGracefulRestartTimeType) GetUuid() string {
    return model.UUID
}

func (model *LongLivedGracefulRestartTimeType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LongLivedGracefulRestartTimeType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LongLivedGracefulRestartTimeType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LongLivedGracefulRestartTimeType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LongLivedGracefulRestartTimeType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *LongLivedGracefulRestartTimeType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LongLivedGracefulRestartTimeType) UpdateReferences() error {
    return nil
}


