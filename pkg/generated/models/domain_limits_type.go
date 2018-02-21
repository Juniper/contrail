
package models
// DomainLimitsType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propDomainLimitsType_virtual_network_limit int = iota
    propDomainLimitsType_security_group_limit int = iota
    propDomainLimitsType_project_limit int = iota
)

// DomainLimitsType 
type DomainLimitsType struct {

    ProjectLimit int `json:"project_limit,omitempty"`
    VirtualNetworkLimit int `json:"virtual_network_limit,omitempty"`
    SecurityGroupLimit int `json:"security_group_limit,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *DomainLimitsType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeDomainLimitsType makes DomainLimitsType
func MakeDomainLimitsType() *DomainLimitsType{
    return &DomainLimitsType{
    //TODO(nati): Apply default
    ProjectLimit: 0,
        VirtualNetworkLimit: 0,
        SecurityGroupLimit: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeDomainLimitsTypeSlice makes a slice of DomainLimitsType
func MakeDomainLimitsTypeSlice() []*DomainLimitsType {
    return []*DomainLimitsType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DomainLimitsType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *DomainLimitsType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *DomainLimitsType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *DomainLimitsType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *DomainLimitsType) GetFQName() []string {
    return model.FQName
}

func (model *DomainLimitsType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DomainLimitsType) GetParentType() string {
    return model.ParentType
}

func (model *DomainLimitsType) GetUuid() string {
    return model.UUID
}

func (model *DomainLimitsType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DomainLimitsType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DomainLimitsType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DomainLimitsType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DomainLimitsType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propDomainLimitsType_project_limit) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProjectLimit); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProjectLimit as project_limit")
        }
        msg["project_limit"] = &val
    }
    
    if model.modified.Bit(propDomainLimitsType_virtual_network_limit) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualNetworkLimit); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualNetworkLimit as virtual_network_limit")
        }
        msg["virtual_network_limit"] = &val
    }
    
    if model.modified.Bit(propDomainLimitsType_security_group_limit) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecurityGroupLimit); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecurityGroupLimit as security_group_limit")
        }
        msg["security_group_limit"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *DomainLimitsType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DomainLimitsType) UpdateReferences() error {
    return nil
}


