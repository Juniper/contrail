
package models
// VirtualRouterNetworkIpamType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualRouterNetworkIpamType_subnet int = iota
    propVirtualRouterNetworkIpamType_allocation_pools int = iota
)

// VirtualRouterNetworkIpamType 
type VirtualRouterNetworkIpamType struct {

    Subnet []*SubnetType `json:"subnet,omitempty"`
    AllocationPools []*AllocationPoolType `json:"allocation_pools,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VirtualRouterNetworkIpamType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamType() *VirtualRouterNetworkIpamType{
    return &VirtualRouterNetworkIpamType{
    //TODO(nati): Apply default
    
            
                Subnet:  MakeSubnetTypeSlice(),
            
        
            
                AllocationPools:  MakeAllocationPoolTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualRouterNetworkIpamTypeSlice makes a slice of VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamTypeSlice() []*VirtualRouterNetworkIpamType {
    return []*VirtualRouterNetworkIpamType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualRouterNetworkIpamType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VirtualRouterNetworkIpamType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualRouterNetworkIpamType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VirtualRouterNetworkIpamType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualRouterNetworkIpamType) GetFQName() []string {
    return model.FQName
}

func (model *VirtualRouterNetworkIpamType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualRouterNetworkIpamType) GetParentType() string {
    return model.ParentType
}

func (model *VirtualRouterNetworkIpamType) GetUuid() string {
    return model.UUID
}

func (model *VirtualRouterNetworkIpamType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualRouterNetworkIpamType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualRouterNetworkIpamType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualRouterNetworkIpamType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualRouterNetworkIpamType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualRouterNetworkIpamType_subnet) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Subnet); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Subnet as subnet")
        }
        msg["subnet"] = &val
    }
    
    if model.modified.Bit(propVirtualRouterNetworkIpamType_allocation_pools) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AllocationPools); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AllocationPools as allocation_pools")
        }
        msg["allocation_pools"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualRouterNetworkIpamType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualRouterNetworkIpamType) UpdateReferences() error {
    return nil
}


