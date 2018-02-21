
package models
// FlowAgingTimeoutList



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFlowAgingTimeoutList_flow_aging_timeout int = iota
)

// FlowAgingTimeoutList 
type FlowAgingTimeoutList struct {

    FlowAgingTimeout []*FlowAgingTimeout `json:"flow_aging_timeout,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FlowAgingTimeoutList) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFlowAgingTimeoutList makes FlowAgingTimeoutList
func MakeFlowAgingTimeoutList() *FlowAgingTimeoutList{
    return &FlowAgingTimeoutList{
    //TODO(nati): Apply default
    
            
                FlowAgingTimeout:  MakeFlowAgingTimeoutSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeFlowAgingTimeoutListSlice makes a slice of FlowAgingTimeoutList
func MakeFlowAgingTimeoutListSlice() []*FlowAgingTimeoutList {
    return []*FlowAgingTimeoutList{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FlowAgingTimeoutList) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FlowAgingTimeoutList) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FlowAgingTimeoutList) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FlowAgingTimeoutList) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FlowAgingTimeoutList) GetFQName() []string {
    return model.FQName
}

func (model *FlowAgingTimeoutList) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FlowAgingTimeoutList) GetParentType() string {
    return model.ParentType
}

func (model *FlowAgingTimeoutList) GetUuid() string {
    return model.UUID
}

func (model *FlowAgingTimeoutList) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FlowAgingTimeoutList) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FlowAgingTimeoutList) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FlowAgingTimeoutList) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FlowAgingTimeoutList) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFlowAgingTimeoutList_flow_aging_timeout) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FlowAgingTimeout); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FlowAgingTimeout as flow_aging_timeout")
        }
        msg["flow_aging_timeout"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FlowAgingTimeoutList) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FlowAgingTimeoutList) UpdateReferences() error {
    return nil
}


