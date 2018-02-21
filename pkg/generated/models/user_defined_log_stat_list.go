
package models
// UserDefinedLogStatList



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propUserDefinedLogStatList_statlist int = iota
)

// UserDefinedLogStatList 
type UserDefinedLogStatList struct {

    Statlist []*UserDefinedLogStat `json:"statlist,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *UserDefinedLogStatList) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeUserDefinedLogStatList makes UserDefinedLogStatList
func MakeUserDefinedLogStatList() *UserDefinedLogStatList{
    return &UserDefinedLogStatList{
    //TODO(nati): Apply default
    
            
                Statlist:  MakeUserDefinedLogStatSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeUserDefinedLogStatListSlice makes a slice of UserDefinedLogStatList
func MakeUserDefinedLogStatListSlice() []*UserDefinedLogStatList {
    return []*UserDefinedLogStatList{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *UserDefinedLogStatList) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *UserDefinedLogStatList) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *UserDefinedLogStatList) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *UserDefinedLogStatList) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *UserDefinedLogStatList) GetFQName() []string {
    return model.FQName
}

func (model *UserDefinedLogStatList) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *UserDefinedLogStatList) GetParentType() string {
    return model.ParentType
}

func (model *UserDefinedLogStatList) GetUuid() string {
    return model.UUID
}

func (model *UserDefinedLogStatList) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *UserDefinedLogStatList) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *UserDefinedLogStatList) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *UserDefinedLogStatList) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *UserDefinedLogStatList) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propUserDefinedLogStatList_statlist) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Statlist); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Statlist as statlist")
        }
        msg["statlist"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *UserDefinedLogStatList) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *UserDefinedLogStatList) UpdateReferences() error {
    return nil
}


