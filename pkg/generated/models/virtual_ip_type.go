
package models
// VirtualIpType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualIpType_status_description int = iota
    propVirtualIpType_protocol int = iota
    propVirtualIpType_persistence_cookie_name int = iota
    propVirtualIpType_connection_limit int = iota
    propVirtualIpType_persistence_type int = iota
    propVirtualIpType_address int = iota
    propVirtualIpType_status int = iota
    propVirtualIpType_subnet_id int = iota
    propVirtualIpType_admin_state int = iota
    propVirtualIpType_protocol_port int = iota
)

// VirtualIpType 
type VirtualIpType struct {

    Status string `json:"status,omitempty"`
    SubnetID UuidStringType `json:"subnet_id,omitempty"`
    AdminState bool `json:"admin_state"`
    ProtocolPort int `json:"protocol_port,omitempty"`
    StatusDescription string `json:"status_description,omitempty"`
    Protocol LoadbalancerProtocolType `json:"protocol,omitempty"`
    PersistenceCookieName string `json:"persistence_cookie_name,omitempty"`
    ConnectionLimit int `json:"connection_limit,omitempty"`
    PersistenceType SessionPersistenceType `json:"persistence_type,omitempty"`
    Address IpAddressType `json:"address,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VirtualIpType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualIpType makes VirtualIpType
func MakeVirtualIpType() *VirtualIpType{
    return &VirtualIpType{
    //TODO(nati): Apply default
    PersistenceType: MakeSessionPersistenceType(),
        Address: MakeIpAddressType(),
        StatusDescription: "",
        Protocol: MakeLoadbalancerProtocolType(),
        PersistenceCookieName: "",
        ConnectionLimit: 0,
        Status: "",
        SubnetID: MakeUuidStringType(),
        AdminState: false,
        ProtocolPort: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualIpTypeSlice makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
    return []*VirtualIpType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualIpType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VirtualIpType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualIpType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VirtualIpType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualIpType) GetFQName() []string {
    return model.FQName
}

func (model *VirtualIpType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualIpType) GetParentType() string {
    return model.ParentType
}

func (model *VirtualIpType) GetUuid() string {
    return model.UUID
}

func (model *VirtualIpType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualIpType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualIpType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualIpType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualIpType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualIpType_subnet_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SubnetID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SubnetID as subnet_id")
        }
        msg["subnet_id"] = &val
    }
    
    if model.modified.Bit(propVirtualIpType_admin_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AdminState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AdminState as admin_state")
        }
        msg["admin_state"] = &val
    }
    
    if model.modified.Bit(propVirtualIpType_protocol_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProtocolPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProtocolPort as protocol_port")
        }
        msg["protocol_port"] = &val
    }
    
    if model.modified.Bit(propVirtualIpType_status) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Status); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Status as status")
        }
        msg["status"] = &val
    }
    
    if model.modified.Bit(propVirtualIpType_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propVirtualIpType_persistence_cookie_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PersistenceCookieName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PersistenceCookieName as persistence_cookie_name")
        }
        msg["persistence_cookie_name"] = &val
    }
    
    if model.modified.Bit(propVirtualIpType_connection_limit) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ConnectionLimit); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ConnectionLimit as connection_limit")
        }
        msg["connection_limit"] = &val
    }
    
    if model.modified.Bit(propVirtualIpType_persistence_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PersistenceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PersistenceType as persistence_type")
        }
        msg["persistence_type"] = &val
    }
    
    if model.modified.Bit(propVirtualIpType_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Address); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Address as address")
        }
        msg["address"] = &val
    }
    
    if model.modified.Bit(propVirtualIpType_status_description) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StatusDescription); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StatusDescription as status_description")
        }
        msg["status_description"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualIpType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualIpType) UpdateReferences() error {
    return nil
}


