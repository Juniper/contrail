
package models
// DiscoveryPubSubEndPointType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propDiscoveryPubSubEndPointType_ep_version int = iota
    propDiscoveryPubSubEndPointType_ep_id int = iota
    propDiscoveryPubSubEndPointType_ep_type int = iota
    propDiscoveryPubSubEndPointType_ep_prefix int = iota
)

// DiscoveryPubSubEndPointType 
type DiscoveryPubSubEndPointType struct {

    EpPrefix *SubnetType `json:"ep_prefix,omitempty"`
    EpVersion string `json:"ep_version,omitempty"`
    EpID string `json:"ep_id,omitempty"`
    EpType string `json:"ep_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *DiscoveryPubSubEndPointType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeDiscoveryPubSubEndPointType makes DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointType() *DiscoveryPubSubEndPointType{
    return &DiscoveryPubSubEndPointType{
    //TODO(nati): Apply default
    EpVersion: "",
        EpID: "",
        EpType: "",
        EpPrefix: MakeSubnetType(),
        
        modified: big.NewInt(0),
    }
}



// MakeDiscoveryPubSubEndPointTypeSlice makes a slice of DiscoveryPubSubEndPointType
func MakeDiscoveryPubSubEndPointTypeSlice() []*DiscoveryPubSubEndPointType {
    return []*DiscoveryPubSubEndPointType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DiscoveryPubSubEndPointType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *DiscoveryPubSubEndPointType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *DiscoveryPubSubEndPointType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *DiscoveryPubSubEndPointType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *DiscoveryPubSubEndPointType) GetFQName() []string {
    return model.FQName
}

func (model *DiscoveryPubSubEndPointType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DiscoveryPubSubEndPointType) GetParentType() string {
    return model.ParentType
}

func (model *DiscoveryPubSubEndPointType) GetUuid() string {
    return model.UUID
}

func (model *DiscoveryPubSubEndPointType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DiscoveryPubSubEndPointType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DiscoveryPubSubEndPointType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DiscoveryPubSubEndPointType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DiscoveryPubSubEndPointType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propDiscoveryPubSubEndPointType_ep_version) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EpVersion); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EpVersion as ep_version")
        }
        msg["ep_version"] = &val
    }
    
    if model.modified.Bit(propDiscoveryPubSubEndPointType_ep_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EpID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EpID as ep_id")
        }
        msg["ep_id"] = &val
    }
    
    if model.modified.Bit(propDiscoveryPubSubEndPointType_ep_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EpType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EpType as ep_type")
        }
        msg["ep_type"] = &val
    }
    
    if model.modified.Bit(propDiscoveryPubSubEndPointType_ep_prefix) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EpPrefix); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EpPrefix as ep_prefix")
        }
        msg["ep_prefix"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *DiscoveryPubSubEndPointType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DiscoveryPubSubEndPointType) UpdateReferences() error {
    return nil
}


