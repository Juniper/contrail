
package models
// DiscoveryServiceAssignmentType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propDiscoveryServiceAssignmentType_subscriber int = iota
    propDiscoveryServiceAssignmentType_publisher int = iota
)

// DiscoveryServiceAssignmentType 
type DiscoveryServiceAssignmentType struct {

    Subscriber []*DiscoveryPubSubEndPointType `json:"subscriber,omitempty"`
    Publisher *DiscoveryPubSubEndPointType `json:"publisher,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *DiscoveryServiceAssignmentType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeDiscoveryServiceAssignmentType makes DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentType() *DiscoveryServiceAssignmentType{
    return &DiscoveryServiceAssignmentType{
    //TODO(nati): Apply default
    
            
                Subscriber:  MakeDiscoveryPubSubEndPointTypeSlice(),
            
        Publisher: MakeDiscoveryPubSubEndPointType(),
        
        modified: big.NewInt(0),
    }
}



// MakeDiscoveryServiceAssignmentTypeSlice makes a slice of DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentTypeSlice() []*DiscoveryServiceAssignmentType {
    return []*DiscoveryServiceAssignmentType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DiscoveryServiceAssignmentType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *DiscoveryServiceAssignmentType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *DiscoveryServiceAssignmentType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *DiscoveryServiceAssignmentType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *DiscoveryServiceAssignmentType) GetFQName() []string {
    return model.FQName
}

func (model *DiscoveryServiceAssignmentType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DiscoveryServiceAssignmentType) GetParentType() string {
    return model.ParentType
}

func (model *DiscoveryServiceAssignmentType) GetUuid() string {
    return model.UUID
}

func (model *DiscoveryServiceAssignmentType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DiscoveryServiceAssignmentType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DiscoveryServiceAssignmentType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DiscoveryServiceAssignmentType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DiscoveryServiceAssignmentType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propDiscoveryServiceAssignmentType_subscriber) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Subscriber); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Subscriber as subscriber")
        }
        msg["subscriber"] = &val
    }
    
    if model.modified.Bit(propDiscoveryServiceAssignmentType_publisher) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Publisher); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Publisher as publisher")
        }
        msg["publisher"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *DiscoveryServiceAssignmentType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DiscoveryServiceAssignmentType) UpdateReferences() error {
    return nil
}


