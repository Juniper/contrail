
package models
// LoadbalancerPoolType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancerPoolType_status int = iota
    propLoadbalancerPoolType_protocol int = iota
    propLoadbalancerPoolType_subnet_id int = iota
    propLoadbalancerPoolType_session_persistence int = iota
    propLoadbalancerPoolType_admin_state int = iota
    propLoadbalancerPoolType_persistence_cookie_name int = iota
    propLoadbalancerPoolType_status_description int = iota
    propLoadbalancerPoolType_loadbalancer_method int = iota
)

// LoadbalancerPoolType 
type LoadbalancerPoolType struct {

    Protocol LoadbalancerProtocolType `json:"protocol,omitempty"`
    SubnetID UuidStringType `json:"subnet_id,omitempty"`
    SessionPersistence SessionPersistenceType `json:"session_persistence,omitempty"`
    AdminState bool `json:"admin_state"`
    PersistenceCookieName string `json:"persistence_cookie_name,omitempty"`
    StatusDescription string `json:"status_description,omitempty"`
    LoadbalancerMethod LoadbalancerMethodType `json:"loadbalancer_method,omitempty"`
    Status string `json:"status,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *LoadbalancerPoolType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancerPoolType makes LoadbalancerPoolType
func MakeLoadbalancerPoolType() *LoadbalancerPoolType{
    return &LoadbalancerPoolType{
    //TODO(nati): Apply default
    SubnetID: MakeUuidStringType(),
        SessionPersistence: MakeSessionPersistenceType(),
        AdminState: false,
        PersistenceCookieName: "",
        StatusDescription: "",
        LoadbalancerMethod: MakeLoadbalancerMethodType(),
        Status: "",
        Protocol: MakeLoadbalancerProtocolType(),
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerPoolTypeSlice makes a slice of LoadbalancerPoolType
func MakeLoadbalancerPoolTypeSlice() []*LoadbalancerPoolType {
    return []*LoadbalancerPoolType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerPoolType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LoadbalancerPoolType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerPoolType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LoadbalancerPoolType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerPoolType) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerPoolType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerPoolType) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerPoolType) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerPoolType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerPoolType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerPoolType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerPoolType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerPoolType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancerPoolType_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPoolType_subnet_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SubnetID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SubnetID as subnet_id")
        }
        msg["subnet_id"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPoolType_session_persistence) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SessionPersistence); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SessionPersistence as session_persistence")
        }
        msg["session_persistence"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPoolType_admin_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AdminState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AdminState as admin_state")
        }
        msg["admin_state"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPoolType_persistence_cookie_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PersistenceCookieName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PersistenceCookieName as persistence_cookie_name")
        }
        msg["persistence_cookie_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPoolType_status_description) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StatusDescription); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StatusDescription as status_description")
        }
        msg["status_description"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPoolType_loadbalancer_method) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerMethod); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerMethod as loadbalancer_method")
        }
        msg["loadbalancer_method"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPoolType_status) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Status); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Status as status")
        }
        msg["status"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LoadbalancerPoolType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerPoolType) UpdateReferences() error {
    return nil
}


