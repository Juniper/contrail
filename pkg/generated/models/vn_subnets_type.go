
package models
// VnSubnetsType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVnSubnetsType_ipam_subnets int = iota
    propVnSubnetsType_host_routes int = iota
)

// VnSubnetsType 
type VnSubnetsType struct {

    IpamSubnets []*IpamSubnetType `json:"ipam_subnets,omitempty"`
    HostRoutes *RouteTableType `json:"host_routes,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VnSubnetsType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVnSubnetsType makes VnSubnetsType
func MakeVnSubnetsType() *VnSubnetsType{
    return &VnSubnetsType{
    //TODO(nati): Apply default
    
            
                IpamSubnets:  MakeIpamSubnetTypeSlice(),
            
        HostRoutes: MakeRouteTableType(),
        
        modified: big.NewInt(0),
    }
}



// MakeVnSubnetsTypeSlice makes a slice of VnSubnetsType
func MakeVnSubnetsTypeSlice() []*VnSubnetsType {
    return []*VnSubnetsType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VnSubnetsType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VnSubnetsType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VnSubnetsType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VnSubnetsType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VnSubnetsType) GetFQName() []string {
    return model.FQName
}

func (model *VnSubnetsType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VnSubnetsType) GetParentType() string {
    return model.ParentType
}

func (model *VnSubnetsType) GetUuid() string {
    return model.UUID
}

func (model *VnSubnetsType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VnSubnetsType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VnSubnetsType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VnSubnetsType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VnSubnetsType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVnSubnetsType_host_routes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.HostRoutes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: HostRoutes as host_routes")
        }
        msg["host_routes"] = &val
    }
    
    if model.modified.Bit(propVnSubnetsType_ipam_subnets) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IpamSubnets); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IpamSubnets as ipam_subnets")
        }
        msg["ipam_subnets"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VnSubnetsType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VnSubnetsType) UpdateReferences() error {
    return nil
}


