
package models
// ShareType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propShareType_tenant int = iota
    propShareType_tenant_access int = iota
)

// ShareType 
type ShareType struct {

    TenantAccess AccessType `json:"tenant_access,omitempty"`
    Tenant string `json:"tenant,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ShareType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeShareType makes ShareType
func MakeShareType() *ShareType{
    return &ShareType{
    //TODO(nati): Apply default
    TenantAccess: MakeAccessType(),
        Tenant: "",
        
        modified: big.NewInt(0),
    }
}



// MakeShareTypeSlice makes a slice of ShareType
func MakeShareTypeSlice() []*ShareType {
    return []*ShareType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ShareType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ShareType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ShareType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ShareType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ShareType) GetFQName() []string {
    return model.FQName
}

func (model *ShareType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ShareType) GetParentType() string {
    return model.ParentType
}

func (model *ShareType) GetUuid() string {
    return model.UUID
}

func (model *ShareType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ShareType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ShareType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ShareType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ShareType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propShareType_tenant_access) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TenantAccess); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TenantAccess as tenant_access")
        }
        msg["tenant_access"] = &val
    }
    
    if model.modified.Bit(propShareType_tenant) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Tenant); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Tenant as tenant")
        }
        msg["tenant"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ShareType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ShareType) UpdateReferences() error {
    return nil
}


