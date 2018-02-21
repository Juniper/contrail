
package models
// QosIdForwardingClassPair



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propQosIdForwardingClassPair_key int = iota
    propQosIdForwardingClassPair_forwarding_class_id int = iota
)

// QosIdForwardingClassPair 
type QosIdForwardingClassPair struct {

    Key int `json:"key,omitempty"`
    ForwardingClassID ForwardingClassId `json:"forwarding_class_id,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *QosIdForwardingClassPair) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeQosIdForwardingClassPair makes QosIdForwardingClassPair
func MakeQosIdForwardingClassPair() *QosIdForwardingClassPair{
    return &QosIdForwardingClassPair{
    //TODO(nati): Apply default
    Key: 0,
        ForwardingClassID: MakeForwardingClassId(),
        
        modified: big.NewInt(0),
    }
}



// MakeQosIdForwardingClassPairSlice makes a slice of QosIdForwardingClassPair
func MakeQosIdForwardingClassPairSlice() []*QosIdForwardingClassPair {
    return []*QosIdForwardingClassPair{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *QosIdForwardingClassPair) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *QosIdForwardingClassPair) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *QosIdForwardingClassPair) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *QosIdForwardingClassPair) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *QosIdForwardingClassPair) GetFQName() []string {
    return model.FQName
}

func (model *QosIdForwardingClassPair) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *QosIdForwardingClassPair) GetParentType() string {
    return model.ParentType
}

func (model *QosIdForwardingClassPair) GetUuid() string {
    return model.UUID
}

func (model *QosIdForwardingClassPair) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *QosIdForwardingClassPair) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *QosIdForwardingClassPair) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *QosIdForwardingClassPair) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *QosIdForwardingClassPair) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propQosIdForwardingClassPair_key) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Key); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Key as key")
        }
        msg["key"] = &val
    }
    
    if model.modified.Bit(propQosIdForwardingClassPair_forwarding_class_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ForwardingClassID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ForwardingClassID as forwarding_class_id")
        }
        msg["forwarding_class_id"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *QosIdForwardingClassPair) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *QosIdForwardingClassPair) UpdateReferences() error {
    return nil
}


