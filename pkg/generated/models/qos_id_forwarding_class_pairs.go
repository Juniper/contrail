
package models
// QosIdForwardingClassPairs



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propQosIdForwardingClassPairs_qos_id_forwarding_class_pair int = iota
)

// QosIdForwardingClassPairs 
type QosIdForwardingClassPairs struct {

    QosIDForwardingClassPair []*QosIdForwardingClassPair `json:"qos_id_forwarding_class_pair,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *QosIdForwardingClassPairs) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeQosIdForwardingClassPairs makes QosIdForwardingClassPairs
func MakeQosIdForwardingClassPairs() *QosIdForwardingClassPairs{
    return &QosIdForwardingClassPairs{
    //TODO(nati): Apply default
    
            
                QosIDForwardingClassPair:  MakeQosIdForwardingClassPairSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeQosIdForwardingClassPairsSlice makes a slice of QosIdForwardingClassPairs
func MakeQosIdForwardingClassPairsSlice() []*QosIdForwardingClassPairs {
    return []*QosIdForwardingClassPairs{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *QosIdForwardingClassPairs) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *QosIdForwardingClassPairs) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *QosIdForwardingClassPairs) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *QosIdForwardingClassPairs) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *QosIdForwardingClassPairs) GetFQName() []string {
    return model.FQName
}

func (model *QosIdForwardingClassPairs) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *QosIdForwardingClassPairs) GetParentType() string {
    return model.ParentType
}

func (model *QosIdForwardingClassPairs) GetUuid() string {
    return model.UUID
}

func (model *QosIdForwardingClassPairs) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *QosIdForwardingClassPairs) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *QosIdForwardingClassPairs) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *QosIdForwardingClassPairs) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *QosIdForwardingClassPairs) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propQosIdForwardingClassPairs_qos_id_forwarding_class_pair) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.QosIDForwardingClassPair); err != nil {
            return nil, errors.Wrap(err, "Marshal of: QosIDForwardingClassPair as qos_id_forwarding_class_pair")
        }
        msg["qos_id_forwarding_class_pair"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *QosIdForwardingClassPairs) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *QosIdForwardingClassPairs) UpdateReferences() error {
    return nil
}


