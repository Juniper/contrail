
package models
// DhcpOptionsListType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propDhcpOptionsListType_dhcp_option int = iota
)

// DhcpOptionsListType 
type DhcpOptionsListType struct {

    DHCPOption []*DhcpOptionType `json:"dhcp_option,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *DhcpOptionsListType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeDhcpOptionsListType makes DhcpOptionsListType
func MakeDhcpOptionsListType() *DhcpOptionsListType{
    return &DhcpOptionsListType{
    //TODO(nati): Apply default
    
            
                DHCPOption:  MakeDhcpOptionTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeDhcpOptionsListTypeSlice makes a slice of DhcpOptionsListType
func MakeDhcpOptionsListTypeSlice() []*DhcpOptionsListType {
    return []*DhcpOptionsListType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DhcpOptionsListType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *DhcpOptionsListType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *DhcpOptionsListType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *DhcpOptionsListType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *DhcpOptionsListType) GetFQName() []string {
    return model.FQName
}

func (model *DhcpOptionsListType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DhcpOptionsListType) GetParentType() string {
    return model.ParentType
}

func (model *DhcpOptionsListType) GetUuid() string {
    return model.UUID
}

func (model *DhcpOptionsListType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DhcpOptionsListType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DhcpOptionsListType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DhcpOptionsListType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DhcpOptionsListType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propDhcpOptionsListType_dhcp_option) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DHCPOption); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DHCPOption as dhcp_option")
        }
        msg["dhcp_option"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *DhcpOptionsListType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DhcpOptionsListType) UpdateReferences() error {
    return nil
}


