
package models
// IpamSubnets



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propIpamSubnets_subnets int = iota
)

// IpamSubnets 
type IpamSubnets struct {

    Subnets []*IpamSubnetType `json:"subnets,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *IpamSubnets) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeIpamSubnets makes IpamSubnets
func MakeIpamSubnets() *IpamSubnets{
    return &IpamSubnets{
    //TODO(nati): Apply default
    
            
                Subnets:  MakeIpamSubnetTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeIpamSubnetsSlice makes a slice of IpamSubnets
func MakeIpamSubnetsSlice() []*IpamSubnets {
    return []*IpamSubnets{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IpamSubnets) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IpamSubnets) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamSubnets) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IpamSubnets) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamSubnets) GetFQName() []string {
    return model.FQName
}

func (model *IpamSubnets) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IpamSubnets) GetParentType() string {
    return model.ParentType
}

func (model *IpamSubnets) GetUuid() string {
    return model.UUID
}

func (model *IpamSubnets) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IpamSubnets) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IpamSubnets) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IpamSubnets) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IpamSubnets) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propIpamSubnets_subnets) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Subnets); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Subnets as subnets")
        }
        msg["subnets"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *IpamSubnets) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IpamSubnets) UpdateReferences() error {
    return nil
}


