
package models
// LoadbalancerMemberType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancerMemberType_admin_state int = iota
    propLoadbalancerMemberType_address int = iota
    propLoadbalancerMemberType_protocol_port int = iota
    propLoadbalancerMemberType_status int = iota
    propLoadbalancerMemberType_status_description int = iota
    propLoadbalancerMemberType_weight int = iota
)

// LoadbalancerMemberType 
type LoadbalancerMemberType struct {

    Status string `json:"status,omitempty"`
    StatusDescription string `json:"status_description,omitempty"`
    Weight int `json:"weight,omitempty"`
    AdminState bool `json:"admin_state"`
    Address IpAddressType `json:"address,omitempty"`
    ProtocolPort int `json:"protocol_port,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *LoadbalancerMemberType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancerMemberType makes LoadbalancerMemberType
func MakeLoadbalancerMemberType() *LoadbalancerMemberType{
    return &LoadbalancerMemberType{
    //TODO(nati): Apply default
    Status: "",
        StatusDescription: "",
        Weight: 0,
        AdminState: false,
        Address: MakeIpAddressType(),
        ProtocolPort: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerMemberTypeSlice makes a slice of LoadbalancerMemberType
func MakeLoadbalancerMemberTypeSlice() []*LoadbalancerMemberType {
    return []*LoadbalancerMemberType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerMemberType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LoadbalancerMemberType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerMemberType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LoadbalancerMemberType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerMemberType) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerMemberType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerMemberType) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerMemberType) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerMemberType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerMemberType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerMemberType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerMemberType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerMemberType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancerMemberType_status_description) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StatusDescription); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StatusDescription as status_description")
        }
        msg["status_description"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMemberType_weight) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Weight); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Weight as weight")
        }
        msg["weight"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMemberType_admin_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AdminState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AdminState as admin_state")
        }
        msg["admin_state"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMemberType_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Address); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Address as address")
        }
        msg["address"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMemberType_protocol_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProtocolPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProtocolPort as protocol_port")
        }
        msg["protocol_port"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMemberType_status) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Status); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Status as status")
        }
        msg["status"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LoadbalancerMemberType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerMemberType) UpdateReferences() error {
    return nil
}


