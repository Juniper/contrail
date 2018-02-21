
package models
// LinklocalServicesTypes



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLinklocalServicesTypes_linklocal_service_entry int = iota
)

// LinklocalServicesTypes 
type LinklocalServicesTypes struct {

    LinklocalServiceEntry []*LinklocalServiceEntryType `json:"linklocal_service_entry,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *LinklocalServicesTypes) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLinklocalServicesTypes makes LinklocalServicesTypes
func MakeLinklocalServicesTypes() *LinklocalServicesTypes{
    return &LinklocalServicesTypes{
    //TODO(nati): Apply default
    
            
                LinklocalServiceEntry:  MakeLinklocalServiceEntryTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeLinklocalServicesTypesSlice makes a slice of LinklocalServicesTypes
func MakeLinklocalServicesTypesSlice() []*LinklocalServicesTypes {
    return []*LinklocalServicesTypes{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LinklocalServicesTypes) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LinklocalServicesTypes) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LinklocalServicesTypes) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LinklocalServicesTypes) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LinklocalServicesTypes) GetFQName() []string {
    return model.FQName
}

func (model *LinklocalServicesTypes) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LinklocalServicesTypes) GetParentType() string {
    return model.ParentType
}

func (model *LinklocalServicesTypes) GetUuid() string {
    return model.UUID
}

func (model *LinklocalServicesTypes) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LinklocalServicesTypes) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LinklocalServicesTypes) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LinklocalServicesTypes) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LinklocalServicesTypes) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLinklocalServicesTypes_linklocal_service_entry) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LinklocalServiceEntry); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LinklocalServiceEntry as linklocal_service_entry")
        }
        msg["linklocal_service_entry"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LinklocalServicesTypes) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LinklocalServicesTypes) UpdateReferences() error {
    return nil
}


