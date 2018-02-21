
package models
// FloatingIpPoolSubnetType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFloatingIpPoolSubnetType_subnet_uuid int = iota
)

// FloatingIpPoolSubnetType 
type FloatingIpPoolSubnetType struct {

    SubnetUUID []string `json:"subnet_uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FloatingIpPoolSubnetType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFloatingIpPoolSubnetType makes FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetType() *FloatingIpPoolSubnetType{
    return &FloatingIpPoolSubnetType{
    //TODO(nati): Apply default
    SubnetUUID: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeFloatingIpPoolSubnetTypeSlice makes a slice of FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetTypeSlice() []*FloatingIpPoolSubnetType {
    return []*FloatingIpPoolSubnetType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FloatingIpPoolSubnetType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FloatingIpPoolSubnetType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FloatingIpPoolSubnetType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FloatingIpPoolSubnetType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FloatingIpPoolSubnetType) GetFQName() []string {
    return model.FQName
}

func (model *FloatingIpPoolSubnetType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FloatingIpPoolSubnetType) GetParentType() string {
    return model.ParentType
}

func (model *FloatingIpPoolSubnetType) GetUuid() string {
    return model.UUID
}

func (model *FloatingIpPoolSubnetType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FloatingIpPoolSubnetType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FloatingIpPoolSubnetType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FloatingIpPoolSubnetType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FloatingIpPoolSubnetType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFloatingIpPoolSubnetType_subnet_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SubnetUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SubnetUUID as subnet_uuid")
        }
        msg["subnet_uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FloatingIpPoolSubnetType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FloatingIpPoolSubnetType) UpdateReferences() error {
    return nil
}


